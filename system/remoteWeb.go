package system

import (
	"context"
	"encoding/json"
	iofs "io/fs"
	"net/http"
	"strings"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/structs"
	pb "github.com/Tackem-org/Proto/pb/remoteweb"
)

type RemoteWebSystem struct {
	pages           *map[string]func(in *structs.WebRequest) (*structs.WebReturn, error)
	adminPages      *map[string]func(in *structs.WebRequest) (*structs.WebReturn, error)
	webSockets      *map[string]func(in *WebSocketRequest) (*WebSocketReturn, error)
	adminWebSockets *map[string]func(in *WebSocketRequest) (*WebSocketReturn, error)
	pb.UnimplementedRemoteWebServer
}

func NewRemoteWebServer() *RemoteWebSystem {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.NewRemoteWebServer() *RemoteWebSystem]")
	return &RemoteWebSystem{
		pages:           &pagesData,
		adminPages:      &adminPagesData,
		webSockets:      &webSocketData,
		adminWebSockets: &adminWebSocketData,
	}
}

func (r *RemoteWebSystem) Page(ctx context.Context, in *pb.PageRequest) (*pb.PageResponse, error) {
	logging.Debugf(debug.FUNCTIONCALLS|debug.GPRCSERVER, "CALLED:[system.(r *RemoteWebSystem) Page(ctx context.Context, in *pb.PageRequest) (*pb.PageResponse, error)] {in=%v}", in)
	return r.page(ctx, in, r.pages)
}

func (r *RemoteWebSystem) AdminPage(ctx context.Context, in *pb.PageRequest) (*pb.PageResponse, error) {
	logging.Debugf(debug.FUNCTIONCALLS|debug.GPRCSERVER, "CALLED:[system.(r *RemoteWebSystem) AdminPage(ctx context.Context, in *pb.PageRequest) (*pb.PageResponse, error)] {in=%v}", in)
	return r.page(ctx, in, r.adminPages)
}

//TODO CHANGE HOW THIS DEALS with the new incmming data (getting the path to match and variables from master)
func cleanPath(in string) (cleanPath string) {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.cleanPath(in string) (cleanPath string)] {in=%s}", in)
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

func (r *RemoteWebSystem) page(ctx context.Context, in *pb.PageRequest, section *map[string]func(in *structs.WebRequest) (*structs.WebReturn, error)) (*pb.PageResponse, error) {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.(r *RemoteWebSystem) page(ctx context.Context, in *pb.PageRequest, section *map[string]func(in *structs.WebRequest) (*structs.WebReturn, error)) (*pb.PageResponse, error)] {in=%v}", in)
	cleanPath := cleanPath(in.Path)
	if cleanPath == "" {
		cleanPath = "/"
	}
	user := structs.GetUserData(in.GetUser())
	webRequest := structs.WebRequest{
		FullPath:  in.GetPath(),
		CleanPath: cleanPath,
		User:      user,
		Method:    in.GetMethod(),
	}

	json.Unmarshal([]byte(in.QueryParamsJson), &webRequest.QueryParams)
	json.Unmarshal([]byte(in.PostJson), &webRequest.Post)

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
			logging.Errorf("[GPRC Remote Web System Page Request] %s:%s", in.GetPath(), err.Error())
			return &pb.PageResponse{
				StatusCode:   http.StatusInternalServerError,
				ErrorMessage: "ERROR WITH THE SYSTEM",
			}, nil
		}

		if returnData.StatusCode >= 300 || returnData.StatusCode <= 199 {
			return &pb.PageResponse{
				StatusCode:   returnData.StatusCode,
				ErrorMessage: returnData.ErrorMessage,
			}, nil
		}
		if returnData.FilePath != "" {
			return r.pageFile(returnData, in)
		} else if returnData.PageString != "" {
			return r.pageString(returnData, in)
		}
		logging.Errorf("[GPRC Remote Web System Page Request] %s: Function returned no Page Filename or html string", in.GetPath())
		return &pb.PageResponse{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "ERROR WITH THE SYSTEM",
		}, nil
	}

	logging.Warningf("[GPRC Remote Web System Page Request] %s: Not found", in.GetPath())
	return &pb.PageResponse{
		StatusCode:   http.StatusNotFound,
		ErrorMessage: "Page Not Found",
	}, nil
}

func (r *RemoteWebSystem) pageString(returnData *structs.WebReturn, in *pb.PageRequest) (*pb.PageResponse, error) {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.(r *RemoteWebSystem) pageString(returnData *structs.WebReturn, in *pb.PageRequest) (*pb.PageResponse, error)] {in=%v}", in)
	var pageData []byte
	pageData, err := json.Marshal(returnData.PageData)
	if err != nil {
		logging.Errorf("[GPRC Remote Web System Page Request] %s:%s", in.GetPath(), err.Error())
		return &pb.PageResponse{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "ERROR WITH THE SYSTEM",
		}, nil
	}
	css, js := getBaseCSSandJS(in.Path)
	return &pb.PageResponse{
		StatusCode:        returnData.StatusCode,
		TemplateHtml:      returnData.PageString,
		PageVariablesJson: string(pageData),
		CustomPageName:    returnData.CustomPageName,
		CustomCss:         append(css, returnData.CustomCss...),
		CustomJs:          append(js, returnData.CustomJs...),
	}, nil
}

func (r *RemoteWebSystem) pageFile(returnData *structs.WebReturn, in *pb.PageRequest) (*pb.PageResponse, error) {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.(r *RemoteWebSystem) pageFile(returnData *structs.WebReturn, in *pb.PageRequest) (*pb.PageResponse, error)] {in=%v}", in)
	templateHtml, err := fileSystem.ReadFile("pages/" + returnData.FilePath + ".html")
	if err != nil {
		logging.Errorf("[GPRC Remote Web System Page Request] %s:%s", in.GetPath(), err.Error())
		return &pb.PageResponse{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "ERROR WITH THE SYSTEM",
		}, nil
	}

	var pageData []byte
	pageData, err = json.Marshal(returnData.PageData)
	if err != nil {
		logging.Errorf("[GPRC Remote Web System Page Request] %s:%s", in.GetPath(), err.Error())
		return &pb.PageResponse{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "ERROR WITH THE SYSTEM",
		}, nil
	}
	css, js := getBaseCSSandJS(returnData.FilePath)
	return &pb.PageResponse{
		StatusCode:        returnData.StatusCode,
		TemplateHtml:      string(templateHtml),
		PageVariablesJson: string(pageData),
		CustomPageName:    returnData.CustomPageName,
		CustomCss:         append(css, returnData.CustomCss...),
		CustomJs:          append(js, returnData.CustomJs...),
	}, nil
}

func (r *RemoteWebSystem) File(ctx context.Context, in *pb.FileRequest) (*pb.FileResponse, error) {
	logging.Debugf(debug.FUNCTIONCALLS|debug.GPRCSERVER, "CALLED:[system.(r *RemoteWebSystem) File(returnData *structs.WebReturn, in *pb.FileRequest) (*pb.FileResponse, error)] {in=%v}", in)
	path := strings.Split(in.GetPath(), "/static/")[1]
	data, err := fileSystem.ReadFile(path)
	if err != nil {
		logging.Errorf("[GPRC Remote Web System File Request] %s:%s", in.GetPath(), err.Error())
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
	logging.Debugf(debug.FUNCTIONCALLS|debug.GPRCSERVER, "CALLED:[system.(r *RemoteWebSystem) WebSocket(returnData *structs.WebReturn, in *pb.WebSocketRequest) (*pb.WebSocketResponse, error)] {in=%v}", in)
	return r.webSocket(ctx, in, r.webSockets)
}
func (r *RemoteWebSystem) AdminWebSocket(ctx context.Context, in *pb.WebSocketRequest) (*pb.WebSocketResponse, error) {
	logging.Debugf(debug.FUNCTIONCALLS|debug.GPRCSERVER, "CALLED:[system.(r *RemoteWebSystem) AdminWebSocket(returnData *structs.WebReturn, in *pb.WebSocketRequest) (*pb.WebSocketResponse, error)] {in=%v}", in)
	return r.webSocket(ctx, in, r.adminWebSockets)
}
func (r *RemoteWebSystem) webSocket(ctx context.Context, in *pb.WebSocketRequest, section *map[string]func(in *WebSocketRequest) (*WebSocketReturn, error)) (*pb.WebSocketResponse, error) {
	logging.Debugf(debug.FUNCTIONCALLS|debug.GPRCSERVER, "CALLED:[system.(r *RemoteWebSystem) webSocket(returnData *structs.WebReturn, in *pb.WebSocketRequest, section *map[string]func(in *WebSocketRequest) (*WebSocketReturn, error)) (*pb.WebSocketResponse, error)] {in=%v}", in)
	path := cleanPath(in.Path)
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	var d map[string]interface{}
	json.Unmarshal([]byte(in.DataJson), &d)
	webSocketRequest := WebSocketRequest{
		Path: path,
		User: structs.GetUserData(in.GetUser()),
		Data: d,
	}

	if call, exists := (*section)[path]; exists {
		returnData, err := call(&webSocketRequest)
		if err != nil {
			logging.Errorf("[GPRC Remote Web Socket Request] %s:%s", in.GetPath(), err.Error())
			return &pb.WebSocketResponse{
				StatusCode:   http.StatusInternalServerError,
				ErrorMessage: "ERROR WITH THE SYSTEM",
			}, nil
		}

		returnJson, _ := json.Marshal(returnData.Data)
		return &pb.WebSocketResponse{
			StatusCode:   returnData.StatusCode,
			ErrorMessage: returnData.ErrorMessage,
			DataJson:     string(returnJson),
		}, nil
	}
	return &pb.WebSocketResponse{
		StatusCode:   http.StatusNotFound,
		ErrorMessage: "Web Socket Not Found",
	}, nil
}

func getBaseCSSandJS(path string) (css []string, js []string) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.getBaseCSSandJS() (css []string, js []string)]")
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

	cssfile, err = fileSystem.Open("css/" + path + ".css")
	if err == nil {
		css = append(css, baseurl+"css/"+path)
	}
	if cssfile != nil {
		cssfile.Close()
	}

	jsfile, err = fileSystem.Open("js/" + path + ".js")
	if err == nil {
		js = append(js, baseurl+"js/"+path)
	}
	if jsfile != nil {
		jsfile.Close()
	}
	return
}
