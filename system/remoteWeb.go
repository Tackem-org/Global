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
	pb.UnimplementedRemoteWebServer
}

func (r *RemoteWebSystem) Page(ctx context.Context, in *pb.PageRequest) (*pb.PageResponse, error) {
	logging.Debugf(debug.FUNCTIONCALLS|debug.GRPCSERVER, "CALLED:[system.(r *RemoteWebSystem) Page(ctx context.Context, in *pb.PageRequest) (*pb.PageResponse, error)] {in=%v}", in)
	return r.page(ctx, in, &pagesData)
}

func (r *RemoteWebSystem) AdminPage(ctx context.Context, in *pb.PageRequest) (*pb.PageResponse, error) {
	logging.Debugf(debug.FUNCTIONCALLS|debug.GRPCSERVER, "CALLED:[system.(r *RemoteWebSystem) AdminPage(ctx context.Context, in *pb.PageRequest) (*pb.PageResponse, error)] {in=%v}", in)
	return r.page(ctx, in, &adminPagesData)
}

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

func (r *RemoteWebSystem) page(ctx context.Context, in *pb.PageRequest, section *map[string]PageFunc) (*pb.PageResponse, error) {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.(r *RemoteWebSystem) page(ctx context.Context, in *pb.PageRequest, section *map[string]PageFunc) (*pb.PageResponse, error)] {in=%v}", in)
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
	json.Unmarshal([]byte(in.PathParamsJson), &webRequest.PathVariables)
	if call, exists := (*section)[in.BasePath]; exists {
		returnData, err := call(&webRequest)
		if err != nil {
			logging.Errorf("[GRPC Remote Web System Page Request] %s:%s", in.GetPath(), err.Error())
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
		logging.Errorf("[GRPC Remote Web System Page Request] %s: Function returned no Page Filename or html string", in.GetPath())
		return &pb.PageResponse{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "ERROR WITH THE SYSTEM",
		}, nil
	}

	logging.Warningf("[GRPC Remote Web System Page Request] %s: Not found", in.GetPath())
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
		logging.Errorf("[GRPC Remote Web System Page Request] %s:%s", in.GetPath(), err.Error())
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
		logging.Errorf("[GRPC Remote Web System Page Request] %s:%s", in.GetPath(), err.Error())
		return &pb.PageResponse{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "ERROR WITH THE SYSTEM",
		}, nil
	}

	var pageData []byte
	pageData, err = json.Marshal(returnData.PageData)
	if err != nil {
		logging.Errorf("[GRPC Remote Web System Page Request] %s:%s", in.GetPath(), err.Error())
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
	logging.Debugf(debug.FUNCTIONCALLS|debug.GRPCSERVER, "CALLED:[system.(r *RemoteWebSystem) File(returnData *structs.WebReturn, in *pb.FileRequest) (*pb.FileResponse, error)] {in=%v}", in)

	data, err := fileSystem.ReadFile(in.Path)
	if err != nil {
		logging.Errorf("[GRPC Remote Web System File Request] %s:%s", in.GetPath(), err.Error())
		sc := http.StatusInternalServerError
		em := "Internal Error"
		switch err.(type) {
		case *iofs.PathError:
			if in.Path[len(in.Path)-1:] == "/" {
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
	logging.Debugf(debug.FUNCTIONCALLS|debug.GRPCSERVER, "CALLED:[system.(r *RemoteWebSystem) WebSocket(returnData *structs.WebReturn, in *pb.WebSocketRequest) (*pb.WebSocketResponse, error)] {in=%v}", in)

	var d map[string]interface{}
	json.Unmarshal([]byte(in.DataJson), &d)

	webSocketRequest := WebSocketRequest{
		Command: in.Command,
		User:    structs.GetUserData(in.GetUser()),
		Data:    d,
	}
	command := in.Command
	serviceType := RegData().GetServiceType()
	if serviceType != "service" {
		command = strings.Replace(command, serviceType+".", "", 1)
	}
	serviceName := RegData().GetServiceName()
	command = strings.Replace(command, serviceName+".", "", 1)

	if call, exists := webSocketData[command]; exists {
		returnData, err := call(&webSocketRequest)
		if err != nil {
			logging.Errorf("[GRPC Remote Web Socket Request] %s:%s", in.Command, err.Error())
			return &pb.WebSocketResponse{
				StatusCode:   http.StatusInternalServerError,
				ErrorMessage: "ERROR WITH THE SYSTEM",
			}, nil
		}

		returnJson, _ := json.Marshal(returnData.Data)
		return &pb.WebSocketResponse{
			StatusCode:   returnData.StatusCode,
			ErrorMessage: returnData.ErrorMessage,
			TellAll:      returnData.TellAll,
			DataJson:     string(returnJson),
		}, nil
	}
	return &pb.WebSocketResponse{
		StatusCode:   http.StatusNotFound,
		ErrorMessage: "Web Socket Not Found",
	}, nil
}

func (r *RemoteWebSystem) Tasks(ctx context.Context, in *pb.TasksRequest) (*pb.TasksResponse, error) {
	logging.Debugf(debug.FUNCTIONCALLS|debug.GRPCSERVER, "CALLED:[system.(r *RemoteWebSystem) Tasks(returnData *structs.WebReturn, in *pb.TasksRequest) (*pb.TasksResponse, error)] {in=%v}", in)
	logging.Info("Master Has Requested Tasks")
	t := Data.TaskGrabber()
	if len(t) == 0 {
		logging.Info("No Tasks To Send (sending empty list)")
	} else {
		logging.Info("Sending Master Tasks")

	}
	return &pb.TasksResponse{
		Tasks: t,
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
