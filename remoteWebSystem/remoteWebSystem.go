package remoteWebSystem

import (
	"embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Tackem-org/Global/logging"
)

var (
	pagesData      map[string]func(in *WebRequest) (*WebReturn, error)
	adminPagesData map[string]func(in *WebRequest) (*WebReturn, error)
	fs             *embed.FS
)

type WebRequest struct {
	FullPath      string
	CleanPath     string
	UserID        uint64
	SessionToken  string
	Method        string
	QueryParams   map[string]interface{}
	Post          map[string]interface{}
	PathVariables map[string]interface{}
}

type WebReturn struct {
	FilePath   string
	PageString string
	PageData   map[string]interface{}
}

func Setup(fsIn *embed.FS) {
	pagesData = make(map[string]func(in *WebRequest) (*WebReturn, error))
	adminPagesData = make(map[string]func(in *WebRequest) (*WebReturn, error))
	fs = fsIn
}

func NewServer() *RemoteWebSystem {
	return &RemoteWebSystem{
		pages:      &pagesData,
		adminPages: &adminPagesData,
		fs:         fs,
	}
}

//TODO NEED TO MAKE A SEPERATE FUNCTION FOR ADDING AN ADMIN SECTION
func AddPath(path string, call func(in *WebRequest) (*WebReturn, error)) bool {
	logging.Info(fmt.Sprintf("Adding %s to remoteWeb", path))
	if strings.Contains(path, "static") {
		logging.Warning(fmt.Sprintf("Adding %s to remoteWeb Failed - cannot use static in the name", path))
		return false
	}
	if _, exists := pagesData[path]; exists {
		logging.Warning(fmt.Sprintf("Adding %s to remoteWeb Failed - Path already exists", path))
		return false
	}

	for _, part := range strings.Split(path, "/") {
		if !checkPathPart(part) {
			logging.Warning(fmt.Sprintf("Adding %s to remoteWeb Failed: Part format Bad %s", path, part))
			return false
		}
	}

	pagesData[path] = call
	return true
}

func AddAdminPath(path string, call func(in *WebRequest) (*WebReturn, error)) bool {
	logging.Info(fmt.Sprintf("Adding %s to remoteWeb [Admin]", path))
	if strings.Contains(path, "static") {
		logging.Warning(fmt.Sprintf("Adding %s to remoteWeb Failed - cannot use static in the name", path))
		return false
	}
	if _, exists := adminPagesData[path]; exists {
		logging.Warning(fmt.Sprintf("Adding %s to remoteWeb Failed - Path already exists", path))
		return false
	}

	for _, part := range strings.Split(path, "/") {
		if !checkPathPart(part) {
			logging.Warning(fmt.Sprintf("Adding %s to remoteWeb Failed: Part format Bad %s", path, part))
			return false
		}
	}

	adminPagesData[path] = call
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

func RemoveAdminPath(path string) bool {
	logging.Info(fmt.Sprintf("Removing %s from remoteWeb [Admin]", path))
	if _, exists := adminPagesData[path]; !exists {
		logging.Warning(fmt.Sprintf("Removing %s from remoteWeb Failed - path not found", path))
		return false
	}

	delete(adminPagesData, path)
	return true
}

func checkPathPart(part string) bool {
	startCount := strings.Count(part, "{")
	endCount := strings.Count(part, "}")
	if startCount == 0 && endCount == 0 {
		return true
	}

	if startCount%2 != 0 || endCount%2 != 0 {
		logging.Warning(fmt.Sprintf("Bad Path Part [%s] - Bad Bracket Setup", part))
		return false
	}
	if startCount == 2 || endCount == 2 {
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
	} else {
		// varCount := startCount / 2
		logging.Info("TODO MORE SPECIAL CHECK OF PATH PART FOR MULTI VARIABLES")
	}
	return true
}

func GetPathVariables(path string) (string, *map[string]interface{}) {
	returnData := make(map[string]interface{})
	reString := regexp.MustCompile(`^[a-zA-Z0-9]`)
	for key := range pagesData {
		if strings.Count(key, "{")+strings.Count(key, "}") == 0 {
			if key == path {
				return key, nil
			}
			continue
		}
		keyParts := strings.Split(key, "/")
		pathParts := strings.Split(path, "/")
		if len(keyParts) != len(pathParts) {
			continue
		}
		match := true
		for index, pathPart := range pathParts {
			if strings.Count(keyParts[index], "{{") == 1 && strings.Count(keyParts[index], "}}") == 1 {
				if strings.HasPrefix(keyParts[index], "{{") && strings.HasSuffix(keyParts[index], "}}") {
					splitPart := strings.Split(strings.ReplaceAll(strings.ReplaceAll(keyParts[index], "{", ""), "}", ""), ":")
					if splitPart[0] == "number" {
						if i, err := strconv.Atoi(pathPart); err == nil {
							if returnData == nil {
								returnData = make(map[string]interface{})
							}
							returnData[splitPart[1]] = i
						} else {
							match = false
							break
							// return "Bad Request Variable", http.StatusBadRequest, nil
						}
					} else if splitPart[0] == "string" {
						if matched := reString.MatchString(pathPart); matched {
							returnData[splitPart[1]] = pathPart
						} else {
							match = false
							break
							// return "Bad Request Variable", http.StatusBadRequest, nil
						}
					}
				} else {
					if keyParts[index] != pathPart {
						match = false
						break
					}
				}
			} else {
				// varCount := strings.Count(keyParts[index], "{{")
				logging.Info("TODO MORE SPECIAL CHECK OF PATH PART FOR MULTI VARIABLES")
			}
		}
		if match {
			return key, &returnData
		} else {
			returnData = make(map[string]interface{})
		}
	}

	return "", nil
}
