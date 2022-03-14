package remoteWeb

import (
	"encoding/json"
	"net/http"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/Global/system/setupData"
	pb "github.com/Tackem-org/Proto/pb/remoteweb"
)

func makeWebRequest(in *pb.PageRequest) *structs.WebRequest {

	user := structs.GetUserData(in.User)
	webRequest := structs.WebRequest{
		Path:     in.Path,
		BasePath: in.BasePath,
		User:     user,
		Method:   in.Method,
	}

	json.Unmarshal([]byte(in.QueryParamsJson), &webRequest.QueryParams)  //TODO Change this to helper.StringToStringMap
	json.Unmarshal([]byte(in.PostJson), &webRequest.Post)                //TODO Change this to helper.StringToStringMap
	json.Unmarshal([]byte(in.PathParamsJson), &webRequest.PathVariables) //TODO Change this to helper.StringToStringMap

	return &webRequest
}

func makePageResponse(in *pb.PageRequest, webReturn *structs.WebReturn, err error) *pb.PageResponse {
	if err != nil {
		logging.Error("[GRPC Remote Web System Page Request] %s:%s", in.GetPath(), err.Error())
		return &pb.PageResponse{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "ERROR WITH THE SYSTEM",
		}
	}

	if webReturn.StatusCode >= 300 || webReturn.StatusCode <= 199 {
		return &pb.PageResponse{
			StatusCode:   webReturn.StatusCode,
			ErrorMessage: webReturn.ErrorMessage,
		}
	}
	if webReturn.FilePath != "" {
		return pageFile(webReturn, in)
	} else if webReturn.PageString != "" {
		return pageString(webReturn, in)
	}

	logging.Error("[GRPC Remote Web System Page Request] %s: Function returned no html Data", in.GetPath())
	return &pb.PageResponse{
		StatusCode:   http.StatusInternalServerError,
		ErrorMessage: "ERROR WITH THE SYSTEM",
	}
}

func pageString(returnData *structs.WebReturn, in *pb.PageRequest) *pb.PageResponse {
	var pageData []byte
	pageData, err := json.Marshal(returnData.PageData)
	if err != nil {
		logging.Error("[GRPC Remote Web System Page Request] %s:%s", in.GetPath(), err.Error())
		return &pb.PageResponse{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "ERROR WITH THE SYSTEM",
		}
	}
	css, js := getBaseCSSandJS(in.Path)
	return &pb.PageResponse{
		StatusCode:        returnData.StatusCode,
		TemplateHtml:      returnData.PageString,
		PageVariablesJson: string(pageData),
		CustomPageName:    returnData.CustomPageName,
		CustomCss:         append(css, returnData.CustomCss...),
		CustomJs:          append(js, returnData.CustomJs...),
	}
}

func pageFile(returnData *structs.WebReturn, in *pb.PageRequest) *pb.PageResponse {
	templateHtml, err := setupData.Data.StaticFS.ReadFile("pages/" + returnData.FilePath + ".html")
	if err != nil {
		logging.Error("[GRPC Remote Web System Page Request] %s:%s", in.GetPath(), err.Error())
		return &pb.PageResponse{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "ERROR WITH THE SYSTEM",
		}
	}

	var pageData []byte
	pageData, err = json.Marshal(returnData.PageData)
	if err != nil {
		logging.Error("[GRPC Remote Web System Page Request] %s:%s", in.GetPath(), err.Error())
		return &pb.PageResponse{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "ERROR WITH THE SYSTEM",
		}
	}
	css, js := getBaseCSSandJS(returnData.FilePath)
	return &pb.PageResponse{
		StatusCode:        returnData.StatusCode,
		TemplateHtml:      string(templateHtml),
		PageVariablesJson: string(pageData),
		CustomPageName:    returnData.CustomPageName,
		CustomCss:         append(css, returnData.CustomCss...),
		CustomJs:          append(js, returnData.CustomJs...),
	}
}

func getBaseCSSandJS(path string) (css []string, js []string) {
	baseurl := ""
	if setupData.Data.ServiceType != "system" {
		baseurl += setupData.Data.ServiceType + "/"
	}
	baseurl += setupData.Data.ServiceName + "/static/"

	cssfile, err := setupData.Data.StaticFS.Open("css/" + setupData.Data.ServiceName + ".css")
	if err == nil {
		css = append(css, baseurl+"css/"+setupData.Data.ServiceName)
	}
	if cssfile != nil {
		cssfile.Close()
	}

	jsfile, err := setupData.Data.StaticFS.Open("js/" + setupData.Data.ServiceName + ".js")
	if err == nil {
		js = append(js, baseurl+"js/"+setupData.Data.ServiceName)
	}
	if jsfile != nil {
		jsfile.Close()
	}

	cssfile, err = setupData.Data.StaticFS.Open("css/" + path + ".css")
	if err == nil {
		css = append(css, baseurl+"css/"+path)
	}
	if cssfile != nil {
		cssfile.Close()
	}

	jsfile, err = setupData.Data.StaticFS.Open("js/" + path + ".js")
	if err == nil {
		js = append(js, baseurl+"js/"+path)
	}
	if jsfile != nil {
		jsfile.Close()
	}
	return
}
