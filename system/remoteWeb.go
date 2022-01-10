package system

import (
	"context"
	"encoding/json"
	iofs "io/fs"
	"net/http"
	"strings"

	"github.com/Tackem-org/Global/logging"
	pb "github.com/Tackem-org/Proto/pb/remoteweb"
)

type RemoteWebSystem struct {
	pages      *map[string]func(in *WebRequest) (*WebReturn, error)
	adminPages *map[string]func(in *WebRequest) (*WebReturn, error)
	webSockets *map[string]func(in *WebSocketRequest) (*WebSocketReturn, error)
	pb.UnimplementedRemoteWebServer
}

func NewRemoteWebServer() *RemoteWebSystem {
	return &RemoteWebSystem{
		pages:      &pagesData,
		adminPages: &adminPagesData,
		webSockets: &webSocketData,
	}
}

func (r *RemoteWebSystem) Page(ctx context.Context, in *pb.PageRequest) (*pb.PageResponse, error) {
	logging.Info("[GPRC Remote Web System Page Request] " + in.GetPath())
	return r.page(ctx, in, r.pages)
}

func (r *RemoteWebSystem) AdminPage(ctx context.Context, in *pb.PageRequest) (*pb.PageResponse, error) {
	logging.Info("[GPRC Remote Web System Admin Page Request] " + in.GetPath())
	return r.page(ctx, in, r.adminPages)
}

func cleanPath(in string) (cleanPath string) {
	if strings.HasPrefix(in, "admin/") {
		cleanPath = strings.Replace(in, "admin/", "", 1)
	} else {
		cleanPath = in
	}

	serviceType := RegData().GetServiceType()
	if serviceType != "service" {
		if strings.HasPrefix(cleanPath, serviceType+"/") {
			cleanPath = strings.Replace(cleanPath, serviceType+"/", "", 1)
		}
	}
	serviceName := RegData().GetServiceName()
	if strings.HasPrefix(cleanPath, serviceName+"/") {
		cleanPath = strings.Replace(cleanPath, serviceName+"/", "", 1)
	} else if strings.HasPrefix(cleanPath, serviceName) {
		cleanPath = strings.Replace(cleanPath, serviceName, "", 1)
	}
	return
}
func (r *RemoteWebSystem) page(ctx context.Context, in *pb.PageRequest, section *map[string]func(in *WebRequest) (*WebReturn, error)) (*pb.PageResponse, error) {
	cleanPath := cleanPath(in.Path)
	if cleanPath == "" {
		cleanPath = "/"
	}
	webRequest := WebRequest{
		FullPath:  in.GetPath(),
		CleanPath: cleanPath,
		UserID:    in.GetUserId(),
		Method:    in.GetMethod(),
	}

	json.Unmarshal([]byte(in.GetQueryParamsJson()), &webRequest.QueryParams)
	json.Unmarshal([]byte(in.GetPostJson()), &webRequest.Post)

	pagesKey, pathVariables := getPathVariables(cleanPath, section)
	if pagesKey == "" {
		return &pb.PageResponse{
			StatusCode:   http.StatusNotFound,
			ErrorMessage: "Page Not Found",
		}, nil
	}
	if pathVariables != nil {
		webRequest.PathVariables = *pathVariables
	}
	if call, exists := (*section)[pagesKey]; exists {
		returnData, err := call(&webRequest)
		if err != nil {
			logging.Error("[GPRC Remote Web System Page Request] " + in.GetPath() + ":" + err.Error())
			return &pb.PageResponse{
				StatusCode:   http.StatusInternalServerError,
				ErrorMessage: "ERROR WITH THE SYSTEM",
			}, nil
		}

		if returnData.FilePath != "" {
			return r.pageFile(returnData, in)
		} else if returnData.PageString != "" {
			return r.pageString(returnData, in)
		}
		logging.Error("[GPRC Remote Web System Page Request] " + in.GetPath() + ": Function returned no Page Filename or html string")
		return &pb.PageResponse{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "ERROR WITH THE SYSTEM",
		}, nil
	}

	logging.Warning("[GPRC Remote Web System Page Request] " + in.GetPath() + ": Not found")
	return &pb.PageResponse{
		StatusCode:   http.StatusNotFound,
		ErrorMessage: "Page Not Found",
	}, nil
}

func (r *RemoteWebSystem) pageString(returnData *WebReturn, in *pb.PageRequest) (*pb.PageResponse, error) {
	var pageData []byte
	pageData, err := json.Marshal(returnData.PageData)
	if err != nil {
		logging.Error("[GPRC Remote Web System Page Request] " + in.GetPath() + ":" + err.Error())
		return &pb.PageResponse{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "ERROR WITH THE SYSTEM",
		}, nil
	}
	css, js := getBaseCSSandJS()
	return &pb.PageResponse{
		StatusCode:        http.StatusOK,
		TemplateHtml:      returnData.PageString,
		PageVariablesJson: string(pageData),
		CustomPageName:    returnData.CustomPageName,
		CustomCss:         append(css, returnData.CustomCss...),
		CustomJs:          append(js, returnData.CustomJs...),
	}, nil

}

func (r *RemoteWebSystem) pageFile(returnData *WebReturn, in *pb.PageRequest) (*pb.PageResponse, error) {
	templateHtml, err := fileSystem.ReadFile("pages/" + returnData.FilePath + ".html")
	if err != nil {
		logging.Error("[GPRC Remote Web System Page Request] " + in.GetPath() + ":" + err.Error())
		return &pb.PageResponse{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "ERROR WITH THE SYSTEM",
		}, nil
	}

	var pageData []byte
	pageData, err = json.Marshal(returnData.PageData)
	if err != nil {
		logging.Error("[GPRC Remote Web System Page Request] " + in.GetPath() + ":" + err.Error())
		return &pb.PageResponse{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "ERROR WITH THE SYSTEM",
		}, nil
	}
	css, js := getBaseCSSandJS()
	return &pb.PageResponse{
		StatusCode:        http.StatusOK,
		TemplateHtml:      string(templateHtml),
		PageVariablesJson: string(pageData),
		CustomPageName:    returnData.CustomPageName,
		CustomCss:         append(css, returnData.CustomCss...),
		CustomJs:          append(js, returnData.CustomJs...),
	}, nil
}

func (r *RemoteWebSystem) File(ctx context.Context, in *pb.FileRequest) (*pb.FileResponse, error) {
	logging.Info("[GPRC Remote Web System File Request] " + in.GetPath())

	path := strings.Split(in.GetPath(), "/static/")[1]
	data, err := fileSystem.ReadFile(path)
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
		}, nil
	}
	return &pb.FileResponse{
		StatusCode:   http.StatusOK,
		ErrorMessage: "",
		File:         data,
	}, nil

}

func (r *RemoteWebSystem) WebSocket(ctx context.Context, in *pb.WebSocketRequest) (*pb.WebSocketResponse, error) {
	logging.Info("[GPRC Remote Web Socket Request] " + in.GetPath())

	path := cleanPath(in.Path)
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	var d map[string]interface{}
	json.Unmarshal([]byte(in.DataJson), &d)
	webSocketRequest := WebSocketRequest{
		Path:   path,
		UserID: in.UserId,
		Data:   d,
	}

	if call, exists := (*r.webSockets)[path]; exists {
		returnData, err := call(&webSocketRequest)
		if err != nil {
			logging.Error("[GPRC Remote Web Socket Request] " + in.GetPath() + ":" + err.Error())
			return &pb.WebSocketResponse{
				StatusCode:   http.StatusInternalServerError,
				ErrorMessage: "ERROR WITH THE SYSTEM",
			}, nil
		}

		returnJson, _ := json.Marshal(returnData.Data)
		return &pb.WebSocketResponse{
			StatusCode: http.StatusOK,
			DataJson:   string(returnJson),
		}, nil
	}
	return &pb.WebSocketResponse{
		StatusCode:   http.StatusNotFound,
		ErrorMessage: "Web Socket Not Found",
	}, nil
}

func getBaseCSSandJS() (css []string, js []string) {
	baseurl := ""
	if regData.data.ServiceType != "system" {
		baseurl += regData.data.ServiceType + "/"
	}
	baseurl += regData.data.ServiceName + "/static/"

	cssfile, err := fileSystem.Open("css/" + regData.data.ServiceName + ".css")
	if err == nil {
		css = append(css, baseurl+"css/"+regData.data.ServiceName)
	}
	if cssfile != nil {
		cssfile.Close()
	}

	jsfile, err := fileSystem.Open("js/" + regData.data.ServiceName + ".js")
	if err == nil {
		js = append(js, baseurl+"js/"+regData.data.ServiceName)
	}
	if jsfile != nil {
		jsfile.Close()
	}
	return
}

func (r *RemoteWebSystem) ValidPage(ctx context.Context, in *pb.ValidRequest) (*pb.ValidResponse, error) {
	return r.validPage(ctx, in, r.pages)
}

func (r *RemoteWebSystem) ValidAdminPage(ctx context.Context, in *pb.ValidRequest) (*pb.ValidResponse, error) {
	return r.validPage(ctx, in, r.adminPages)
}

func (r *RemoteWebSystem) validPage(ctx context.Context, in *pb.ValidRequest, section *map[string]func(in *WebRequest) (*WebReturn, error)) (*pb.ValidResponse, error) {
	cleanPath := cleanPath(in.Path)
	if cleanPath == "" {
		cleanPath = "/"
	}
	pagesKey, _ := getPathVariables(cleanPath, section)
	if pagesKey == "" {
		return &pb.ValidResponse{
			Found: false,
		}, nil
	}
	_, exists := (*section)[pagesKey]
	return &pb.ValidResponse{
		Found: exists,
	}, nil

}

func (r *RemoteWebSystem) ValidFile(ctx context.Context, in *pb.ValidRequest) (*pb.ValidResponse, error) {
	path := strings.Split(in.GetPath(), "/static/")[1]
	f, err := fileSystem.Open(path)
	f.Close()
	return &pb.ValidResponse{
		Found: err != nil,
	}, nil

}

func (r *RemoteWebSystem) ValidWebSocket(ctx context.Context, in *pb.ValidRequest) (*pb.ValidResponse, error) {

	path := cleanPath(in.Path)
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	_, exists := (*r.webSockets)[path]
	return &pb.ValidResponse{
		Found: exists,
	}, nil
}
