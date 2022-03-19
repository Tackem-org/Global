package remoteWeb_test

import (
	"errors"
	"io/fs"
	"net/http"
	"testing"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/Global/system/grpcSystem/servers/remoteWeb"
	"github.com/Tackem-org/Global/system/setupData"
	pb "github.com/Tackem-org/Proto/pb/remoteweb"
	"github.com/stretchr/testify/assert"
)

type MockEmbed struct{}

var mockEmbedFiles []string = []string{
	"css/found.css",
	"js/found.js",
	"css/test2.css",
	"js/test3.js",
	"pages/found.html",
}

func (me *MockEmbed) Open(name string) (fs.File, error) {
	for _, v := range mockEmbedFiles {
		if name == v {
			return &MockFile{}, nil
		}
	}
	return nil, errors.New("FAIL")
}

func (me *MockEmbed) ReadFile(name string) ([]byte, error) {
	for _, v := range mockEmbedFiles {
		if name == v {
			return []byte("FOUND"), nil
		}
	}
	return nil, errors.New("FAIL")
}

type MockFile struct{}

func (mf *MockFile) Stat() (fs.FileInfo, error) {
	return nil, nil
}
func (mf *MockFile) Read([]byte) (int, error) {
	return 0, nil
}
func (mf *MockFile) Close() error {
	return nil
}

func TestMakeWebRequest(t *testing.T) {
	in := &pb.PageRequest{
		Path:            "",
		BasePath:        "",
		User:            &pb.UserData{},
		Method:          "",
		QueryParamsJson: "{}",
		PostJson:        "{}",
		PathParamsJson:  "{}",
	}

	expected := &structs.WebRequest{
		Path:          "",
		BasePath:      "",
		User:          &structs.UserData{},
		SessionToken:  "",
		Method:        "",
		QueryParams:   map[string]interface{}{},
		Post:          map[string]interface{}{},
		PathVariables: map[string]interface{}{},
	}
	result := remoteWeb.MakeWebRequest(in)
	assert.Equal(t, expected, result)
}

func TestGetBaseCSSandJS(t *testing.T) {
	tests := []struct {
		serviceType      string
		serviceName      string
		path             string
		expectedCSSCount int
		expectedJSCount  int
	}{
		{"test", "test", "found", 1, 1},
		{"system", "test", "found", 1, 1},
		{"system", "test2", "found", 2, 1},
		{"system", "test3", "found", 1, 2},
	}

	for _, test := range tests {
		setupData.Data = &setupData.SetupData{
			ServiceType: test.serviceType,
			ServiceName: test.serviceName,
			StaticFS:    &MockEmbed{},
		}

		css, js := remoteWeb.GetBaseCSSandJS(test.path)
		assert.Equal(t, test.expectedCSSCount, len(css))
		assert.Equal(t, test.expectedJSCount, len(js))
	}
}

func TestPageFile(t *testing.T) {
	setupData.Data = &setupData.SetupData{
		ServiceType: "service",
		ServiceName: "test",
		StaticFS:    &MockEmbed{},
	}

	returnData1 := &structs.WebReturn{
		FilePath: "missing",
	}
	in := &pb.PageRequest{}
	response1 := remoteWeb.PageFile(returnData1, in)
	assert.NotNil(t, response1)
	assert.Equal(t, http.StatusInternalServerError, int(response1.StatusCode))

	returnData2 := &structs.WebReturn{
		StatusCode: http.StatusOK,
		FilePath:   "found",
	}
	response2 := remoteWeb.PageFile(returnData2, in)
	assert.NotNil(t, response2)
	assert.Equal(t, http.StatusOK, int(response2.StatusCode))

}

func TestPageString(t *testing.T) {
	setupData.Data = &setupData.SetupData{
		ServiceType: "service",
		ServiceName: "test",
		StaticFS:    &MockEmbed{},
	}
	returnData1 := &structs.WebReturn{
		StatusCode: http.StatusOK,
		PageString: "html test",
	}
	in := &pb.PageRequest{}
	response1 := remoteWeb.PageString(returnData1, in)
	assert.Equal(t, http.StatusOK, int(response1.StatusCode))
	assert.Equal(t, returnData1.PageString, response1.TemplateHtml)
}

func TestMakePageResponse(t *testing.T) {
	setupData.Data = &setupData.SetupData{
		ServiceType: "service",
		ServiceName: "test",
		StaticFS:    &MockEmbed{},
	}
	response1 := remoteWeb.MakePageResponse(&pb.PageRequest{}, &structs.WebReturn{}, errors.New("TEST"))
	assert.NotNil(t, response1)
	assert.Equal(t, http.StatusInternalServerError, int(response1.StatusCode))

	response2 := remoteWeb.MakePageResponse(&pb.PageRequest{}, &structs.WebReturn{StatusCode: http.StatusBadGateway}, nil)
	assert.NotNil(t, response2)
	assert.Equal(t, http.StatusBadGateway, int(response2.StatusCode))

	response3 := remoteWeb.MakePageResponse(&pb.PageRequest{}, &structs.WebReturn{StatusCode: http.StatusOK}, nil)
	assert.NotNil(t, response3)
	assert.Equal(t, http.StatusInternalServerError, int(response3.StatusCode))

	//TODO Make this do the last two lines
}
