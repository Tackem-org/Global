package system

import (
	"embed"
	"regexp"
	"strconv"
	"strings"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/structs"
)

func WebSetup(fileSystemIn *embed.FS) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.WebSetup(fileSystemIn *embed.FS)]")
	pagesData = make(map[string]func(in *structs.WebRequest) (*structs.WebReturn, error))
	adminPagesData = make(map[string]func(in *structs.WebRequest) (*structs.WebReturn, error))
	webSocketData = make(map[string]func(in *WebSocketRequest) (*WebSocketReturn, error))
	fileSystem = fileSystemIn
}

func WebAddPath(path string, call func(in *structs.WebRequest) (*structs.WebReturn, error)) bool {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.WebAddPath(path string, call func(in *structs.WebRequest) (*structs.WebReturn, error)) bool] {path=%s}", path)
	if strings.Contains(path, "static") {
		logging.Warningf("Adding %s to remoteWeb Failed - cannot use static in the name", path)
		return false
	}
	if _, exists := pagesData[path]; exists {
		logging.Warningf("Adding %s to remoteWeb Failed - Path already exists", path)
		return false
	}

	for _, part := range strings.Split(path, "/") {
		if !checkPathPart(part) {
			logging.Warningf("Adding %s to remoteWeb Failed: Part format Bad %s", path, part)
			return false
		}
	}

	pagesData[path] = call
	return true
}

func WebAddAdminPath(path string, call func(in *structs.WebRequest) (*structs.WebReturn, error)) bool {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.WebAdminAddPath(path string, call func(in *structs.WebRequest) (*structs.WebReturn, error)) bool] {path=%s}", path)
	if strings.Contains(path, "static") {
		logging.Warningf("Adding %s to remoteWeb Failed - cannot use static in the name", path)
		return false
	}
	if _, exists := adminPagesData[path]; exists {
		logging.Warningf("Adding %s to remoteWeb Failed - Path already exists", path)
		return false
	}

	for _, part := range strings.Split(path, "/") {
		if !checkPathPart(part) {
			logging.Warningf("Adding %s to remoteWeb Failed: Part format Bad %s", path, part)
			return false
		}
	}

	adminPagesData[path] = call
	return true
}

func WebAddWebSocket(path string, call func(in *WebSocketRequest) (*WebSocketReturn, error)) bool {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.WebAddWebSocket(path string, call func(in *structs.WebRequest) (*structs.WebReturn, error)) bool] {path=%s}", path)

	if !strings.HasSuffix(path, ".ws") {
		logging.Warningf("Adding Web Socket %s to remoteWeb Failed - missing \".ws\" Suffix", path)
		return false
	}
	if _, exists := webSocketData[path]; exists {
		logging.Warningf("Adding Web Socket %s to remoteWeb Failed - Path already exists", path)
		return false
	}

	startCount := strings.Count(path, "{")
	endCount := strings.Count(path, "}")
	if startCount != 0 || endCount != 0 {
		logging.Warningf("Adding Web Socket %s to remoteWeb Failed - Path Cannot use Variables", path)
		return false
	}
	webSocketData[path] = call
	return true
}

func WebAddAdminWebSocket(path string, call func(in *WebSocketRequest) (*WebSocketReturn, error)) bool {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.WebAddWebSocket(path string, call func(in *structs.WebRequest) (*structs.WebReturn, error)) bool] {path=%s}", path)

	if !strings.HasSuffix(path, ".ws") {
		logging.Warningf("Adding Web Socket %s to remoteWeb Failed - missing \".ws\" Suffix", path)
		return false
	}
	if _, exists := adminWebSocketData[path]; exists {
		logging.Warningf("Adding Web Socket %s to remoteWeb Failed - Path already exists", path)
		return false
	}

	startCount := strings.Count(path, "{")
	endCount := strings.Count(path, "}")
	if startCount != 0 || endCount != 0 {
		logging.Warningf("Adding Web Socket %s to remoteWeb Failed - Path Cannot use Variables", path)
		return false
	}
	webSocketData[path] = call
	return true
}

func WebRemovePath(path string) bool {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.WebRemovePath(path string) bool] {path=%s}", path)
	if _, exists := pagesData[path]; !exists {
		logging.Warningf("Removing %s from remoteWeb Failed - path not found", path)
		return false
	}

	delete(pagesData, path)
	return true
}

func WebRemoveAdminPath(path string) bool {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.WebRemoveAdminPath(path string) bool] {path=%s}", path)
	if _, exists := adminPagesData[path]; !exists {
		logging.Warningf("Removing %s from remoteWeb Failed - path not found", path)
		return false
	}

	delete(adminPagesData, path)
	return true
}

func WebRemoveWebSocket(path string) bool {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.WebRemoveWebSocket(path string) bool] {path=%s}", path)
	if _, exists := webSocketData[path]; !exists {
		logging.Warningf("Removing Web Socket %s from remoteWeb Failed - path not found", path)
		return false
	}

	delete(webSocketData, path)
	return true
}

func WebRemoveAdminWebSocket(path string) bool {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.WebRemoveWebSocket(path string) bool] {path=%s}", path)
	if _, exists := adminWebSocketData[path]; !exists {
		logging.Warningf("Removing Web Socket %s from remoteWeb Failed - path not found", path)
		return false
	}

	delete(webSocketData, path)
	return true
}

func checkPathPart(part string) bool {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.checkPathPart(part string) bool] {path=%s}", part)
	startCount := strings.Count(part, "{")
	endCount := strings.Count(part, "}")
	if startCount == 0 && endCount == 0 {
		return true
	}

	if startCount%2 != 0 || endCount%2 != 0 {
		logging.Warningf("Bad Path Part [%s] - Bad Bracket Setup", part)
		return false
	}
	if startCount == 2 || endCount == 2 {
		if !strings.HasPrefix(part, "{{") {
			logging.Warningf("Bad Path Part [%s] - Prefix Brackets {{ not at start of section", part)
			return false
		}

		if !strings.HasSuffix(part, "}}") {
			logging.Warningf("Bad Path Part [%s] - Suffix Brackets }} not at end of section", part)
			return false
		}

		splitPart := strings.Split(strings.ReplaceAll(strings.ReplaceAll(part, "{", ""), "}", ""), ":")

		if len(splitPart) != 2 {
			logging.Warningf("Bad Path Part [%s] - Part not in correct format should be {{[number|string]:[valiable name]}}", part)
			return false
		}

		if matched, _ := regexp.Match(`number|string`, []byte(splitPart[0])); !matched {
			logging.Warningf("Bad Path Part [%s] - variable Type not 'number' or 'string'", part)
			return false
		}

		if matched, _ := regexp.Match(`[a-zA-Z0-9]`, []byte(splitPart[1])); !matched {
			logging.Warningf("Bad Path Part [%s] - variable value has no name", part)
			return false
		}
	} else {
		// varCount := startCount / 2
		logging.Info("TODO MORE SPECIAL CHECK OF PATH PART FOR MULTI VARIABLES")
	}
	return true
}

func getPathVariables(path string, section *map[string]func(in *structs.WebRequest) (*structs.WebReturn, error)) (string, *map[string]interface{}) {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.getPathVariables(path string, section *map[string]func(in *structs.WebRequest) (*structs.WebReturn, error)) (string, *map[string]interface{})] {path=%s}", path)
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	returnData := make(map[string]interface{})
	reString := regexp.MustCompile(`^[a-zA-Z0-9]`)
	for key := range *section {
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
						}
					} else if splitPart[0] == "string" {
						if matched := reString.MatchString(pathPart); matched {
							returnData[splitPart[1]] = pathPart
						} else {
							match = false
							break
						}
					}
				} else {
					if keyParts[index] != pathPart {
						match = false
						break
					}
				}
			} else if strings.Count(keyParts[index], "{{") > 1 && strings.Count(keyParts[index], "}}") > 1 {
				// varCount := strings.Count(keyParts[index], "{{")
				logging.Info("TODO MORE SPECIAL CHECK OF PATH PART FOR MULTI VARIABLES")
				logging.Infof("%s:%s", pathPart, keyParts)
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
