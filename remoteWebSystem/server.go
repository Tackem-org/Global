package remoteWebSystem

// needs to accept a map[string]func([inputs])
//Need to somehow send and recieve requests for files in static folder.
//Look at some kind of custom handler on Master that will then use GPRC to request the file on the other side and
//return it to then be served to the system.
import (
	"context"
	"embed"
	"encoding/json"
	iofs "io/fs"
	"net/http"
	"strings"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/registerService"
	pb "github.com/Tackem-org/Proto/pb/remoteweb"
)

type RemoteWebSystem struct {
	pages *map[string]func(in *WebRequest) *WebReturn
	fs    *embed.FS
	pb.UnimplementedRemoteWebServer
}

func (r *RemoteWebSystem) Page(ctx context.Context, in *pb.PageRequest) (*pb.PageResponse, error) {
	cleanPath := in.GetPath()
	adminPage := false
	if strings.HasPrefix(cleanPath, "admin/") {
		adminPage = true
		cleanPath = strings.Replace(cleanPath, "admin/", "", 1)
	}

	serviceType := registerService.Data.GetServiceType()
	if serviceType != "service" {
		if strings.HasPrefix(cleanPath, serviceType+"/") {
			cleanPath = strings.Replace(cleanPath, serviceType+"/", "", 1)
		}
	}
	serviceName := registerService.Data.GetServiceName()
	if strings.HasPrefix(cleanPath, serviceName+"/") {
		cleanPath = strings.Replace(cleanPath, serviceName+"/", "", 1)
	}

	if adminPage {
		cleanPath = "admin/" + cleanPath
	}

	if cleanPath == "" {
		cleanPath = "/"
	}

	logging.Info("[GPRC Remote Web System Page Request] " + cleanPath)

	webRequest := WebRequest{
		FullPath:  in.GetPath(),
		CleanPath: cleanPath,
	}

	json.Unmarshal([]byte(in.GetQueryParamsJson()), &webRequest.QueryParams)
	json.Unmarshal([]byte(in.GetPostJson()), &webRequest.Post)

	pagesKey, pathVariables := GetPathVariables(cleanPath)
	webRequest.PathVariables = *pathVariables

	if call, exists := (*r.pages)[pagesKey]; exists {
		returnData := call(&webRequest)

		templateHtml, err := r.fs.ReadFile("pages/" + returnData.FilePath + ".html")
		if err != nil {
			logging.Error("[GPRC Remote Web System Page Request] " + in.GetPath() + ":" + err.Error())

		}

		var pageData []byte
		pageData, err = json.Marshal(returnData.PageData)
		if err != nil {
			logging.Fatal(err)
		}

		return &pb.PageResponse{
			StatusCode:        http.StatusOK,
			TemplateHtml:      string(templateHtml),
			PageVariablesJson: string(pageData),
		}, nil
	}

	return &pb.PageResponse{
		StatusCode:   http.StatusNotFound,
		ErrorMessage: "Page Not Found",
	}, nil
}

func (r *RemoteWebSystem) File(ctx context.Context, in *pb.FileRequest) (*pb.FileResponse, error) {
	logging.Info("[GPRC Remote Web System File Request] " + in.GetPath())

	path := strings.Split(in.GetPath(), "/static/")[1]
	data, err := fs.ReadFile(path)
	if err != nil {
		logging.Error("[GPRC Remote Web System File Request] " + in.GetPath() + ":" + err.Error())
		sc := http.StatusInternalServerError
		em := "Internal Error"
		switch err.(type) {
		case *iofs.PathError:
			if path[len(path)-1:] == "/" {
				sc = http.StatusForbidden
				em = "Path is a Directory Access Forbidden"
			} else {
				sc = http.StatusNotFound
				em = "File Not Found"
			}
		default:
		}

		return &pb.FileResponse{
			StatusCode:   uint32(sc),
			ErrorMessage: em,
		}, err
	}
	return &pb.FileResponse{
		StatusCode:   http.StatusOK,
		ErrorMessage: "",
		File:         data,
	}, nil

}
