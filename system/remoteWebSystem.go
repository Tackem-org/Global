package system

import (
	"embed"
	"regexp"
	"strings"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/structs"
	pb "github.com/Tackem-org/Proto/pb/registration"
)

type PageFunc = func(in *structs.WebRequest) (*structs.WebReturn, error)
type SocketFunc = func(in *WebSocketRequest) (*WebSocketReturn, error)

func WebSetup(fileSystemIn *embed.FS) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.WebSetup(fileSystemIn *embed.FS)]")
	pagesData = make(map[string]PageFunc)
	adminPagesData = make(map[string]PageFunc)
	webSocketData = make(map[string]SocketFunc)
	fileSystem = fileSystemIn
}

func WebAddPath(webLinkItem *pb.WebLinkItem, call PageFunc) bool {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.WebAddPath(webLinkItem *pb.WebLinkItem, call PageFunc) bool] {webLinkItem=%+v}", webLinkItem)
	if strings.Contains(webLinkItem.Path, "static") {
		logging.Warningf("Adding %s to remoteWeb Failed - cannot use static in the name", webLinkItem.Path)
		return false
	}
	if _, exists := pagesData[webLinkItem.Path]; exists {
		logging.Warningf("Adding %s to remoteWeb Failed - Path already exists", webLinkItem.Path)
		return false
	}

	for _, part := range strings.Split(webLinkItem.Path, "/") {
		if !checkPathPart(part) {
			logging.Warningf("Adding %s to remoteWeb Failed: Part format Bad %s", webLinkItem.Path, part)
			return false
		}
	}

	pagesData[webLinkItem.Path] = call
	pagesProtoData = append(pagesProtoData, webLinkItem)
	return true
}

func WebAddAdminPath(webLinkItem *pb.AdminWebLinkItem, call PageFunc) bool {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.WebAdminAddPath(webLinkItem *pb.WebLinkItem, call PageFunc) bool] {webLinkItem=%+v}", webLinkItem)
	if strings.Contains(webLinkItem.Path, "static") {
		logging.Warningf("Adding %s to remoteWeb Failed - cannot use static in the name", webLinkItem.Path)
		return false
	}
	if _, exists := adminPagesData[webLinkItem.Path]; exists {
		logging.Warningf("Adding %s to remoteWeb Failed - Path already exists", webLinkItem.Path)
		return false
	}

	for _, part := range strings.Split(webLinkItem.Path, "/") {
		if !checkPathPart(part) {
			logging.Warningf("Adding %s to remoteWeb Failed: Part format Bad %s", webLinkItem.Path, part)
			return false
		}
	}

	adminPagesData[webLinkItem.Path] = call
	adminPagesProtoData = append(adminPagesProtoData, webLinkItem)
	return true
}

func WebAddWebSocket(webSocketItem *pb.WebSocketItem, call SocketFunc) bool {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.WebAddWebSocket(webSocketItem *pb.WebSocketItem, call SocketFunc) bool] {webSocketItem.Command=%s}", webSocketItem.Command)

	if _, exists := webSocketData[webSocketItem.Command]; exists {
		logging.Warningf("Adding Web Socket %s to remoteWeb Failed - Command already exists", webSocketItem.Command)
		return false
	}

	if matched, _ := regexp.Match(`[a-zA-Z0-9]`, []byte(webSocketItem.Command)); !matched {
		logging.Warningf("Adding Web Socket %s to remoteWeb Failed - Command name not valid", webSocketItem.Command)
		return false
	}

	webSocketData[webSocketItem.Command] = call
	webSocketProtoData = append(webSocketProtoData, webSocketItem)
	return true
}

func WebRemovePath(path string) bool {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.WebRemovePath(path string) bool] {path=%s}", path)
	if _, exists := pagesData[path]; !exists {
		logging.Warningf("Removing %s from remoteWeb Failed - path not found", path)
		return false
	}

	delete(pagesData, path)
	for index, item := range pagesProtoData {
		if item.Path == path {
			pagesProtoData = append(pagesProtoData[:index], pagesProtoData[index+1:]...)
		}
	}
	return true
}

func WebRemoveAdminPath(path string) bool {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.WebRemoveAdminPath(path string) bool] {path=%s}", path)
	if _, exists := adminPagesData[path]; !exists {
		logging.Warningf("Removing %s from remoteWeb Failed - path not found", path)
		return false
	}

	delete(adminPagesData, path)
	for index, item := range adminPagesProtoData {
		if item.Path == path {
			adminPagesProtoData = append(adminPagesProtoData[:index], adminPagesProtoData[index+1:]...)
		}
	}
	return true
}

func WebRemoveWebSocket(command string) bool {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.WebRemoveWebSocket(command string) bool] {command=%s}", command)
	if _, exists := webSocketData[command]; !exists {
		logging.Warningf("Removing Web Socket %s from remoteWeb Failed - command not found", command)
		return false
	}

	delete(webSocketData, command)
	for index, item := range webSocketProtoData {
		if item.Command == command {
			webSocketProtoData = append(webSocketProtoData[:index], webSocketProtoData[index+1:]...)
		}
	}
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
