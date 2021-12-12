package remoteWebSystem

// needs to accept a map[string]func([inputs])
import (
	"embed"
	"fmt"
	"regexp"
	"strings"

	"github.com/Tackem-org/Global/logging"
)

// pb.UnimplementedRemoteWebServer
var (
	pagesData map[string]func(in *WebRequest) *WebReturn
	fs        *embed.FS
)

type WebRequest struct {
	FullPath      string
	CleanPath     string
	QueryParams   map[string]interface{}
	Post          map[string]interface{}
	PathVariables map[string]interface{}
}

type WebReturn struct {
	FilePath string
	PageData map[string]interface{}
}

func Setup(fsIn *embed.FS) {
	pagesData = make(map[string]func(in *WebRequest) *WebReturn)
	fs = fsIn
}

func NewServer() *RemoteWebSystem {
	return &RemoteWebSystem{
		pages: &pagesData,
		fs:    fs,
	}
}

func AddPath(path string, call func(in *WebRequest) *WebReturn) bool {
	logging.Info(fmt.Sprintf("Adding %s to remoteWeb", path))
	if _, exists := pagesData[path]; exists {
		logging.Warning(fmt.Sprintf("Adding %s to remoteWeb Failed", path))
		return false
	}

	for _, part := range strings.Split(path, "/") {
		if !checkPathPart(part) {
			logging.Warning(fmt.Sprintf("Adding %s to remoteWeb Failed: %s", path, part))
			return false
		}
	}

	pagesData[path] = call
	return true
}

func RemovePath(path string) bool {
	logging.Info(fmt.Sprintf("Removing %s from remoteWeb", path))
	if _, exists := pagesData[path]; !exists {
		logging.Warning(fmt.Sprintf("Removing %s from remoteWeb Failed - path not found", path))
		return false
	}

	delete(pagesData, path)
	return true
}

func checkPathPart(part string) bool {
	startCount := strings.Count(part, "{")
	endCount := strings.Count(part, "}")
	if startCount == 0 && endCount == 0 {
		// if matched, _ := regexp.Match(`[a-zA-Z0-9-/]`, []byte(part)); !matched {
		// 	logging.Warning(fmt.Sprintf("Bad Path Part [%s] - Bad Characters in Path you should only use alphanumeric and hyphen", part))
		// 	return false
		// }
		return true
	}

	if startCount != 2 || endCount != 2 {
		logging.Warning(fmt.Sprintf("Bad Path Part [%s] - Bad Bracket Setup", part))
		return false
	}
	if !strings.HasPrefix(part, "{{") {
		logging.Warning(fmt.Sprintf("Bad Path Part [%s] - Prefix Brackets {{ not at start of section", part))
		return false
	}

	if !strings.HasSuffix(part, "}}") {
		logging.Warning(fmt.Sprintf("Bad Path Part [%s] - Suffix Brackets }} not at end of section", part))
		return false
	}

	splitPart := strings.Split(strings.ReplaceAll(strings.ReplaceAll(part, "{", ""), "}", ""), ":")

	if len(splitPart) != 2 {
		logging.Warning(fmt.Sprintf("Bad Path Part [%s] - Part not in correct format should be {{[number|string]:[valiable name]}}", part))
		return false
	}

	if matched, _ := regexp.Match(`number|string`, []byte(splitPart[0])); !matched {
		logging.Warning(fmt.Sprintf("Bad Path Part [%s] - variable Type not 'number' or 'string'", part))
		return false
	}

	if matched, _ := regexp.Match(`[a-zA-Z0-9]`, []byte(splitPart[1])); !matched {
		logging.Warning(fmt.Sprintf("Bad Path Part [%s] - variable value has no name", part))
		return false
	}

	return true
}

func GetPathVariables(path string) (string, *map[string]interface{}) {
	returnData := new(map[string]interface{})

	// for key := range pagesData {
	// 	if key == path {
	// 		return path, returnData
	// 	}
	// 	if strings.Contains(key, "{{") && strings.Contains(key, "}}") {
	// 		//HERE
	// 	}

	// }

	// splitPath := strings.Split(path, "/")
	// for _, part := range splitPath {
	// 	//HERE
	// 	//for each key in Pages check for "{{}}" if not then do a simple compare and pass back the url if matched

	// 	// if strings.HasPrefix(part, "{{") && strings.HasSuffix(part, "}}") {
	// 	// 	splitPart := strings.Split(strings.ReplaceAll(strings.ReplaceAll(part, "{", ""), "}", ""), ":")
	// 	// }

	// }
	return path, returnData
}
