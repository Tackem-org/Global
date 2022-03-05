package helpers

import (
	"regexp"
	"strings"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
)

func CheckPath(path string) bool {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.helpers.CheckPath")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] path=%s", path)
	if strings.HasPrefix(path, "/") {
		path = strings.Replace(path, "/", "", 1)
	}
	s := strings.Split(path, "/")
	if len(s) == 0 {
		return true
	}
	startCount := strings.Count(path, "{")
	endCount := strings.Count(path, "}")
	if startCount == 0 && endCount == 0 {
		return true
	}
	for _, part := range s {
		v, _ := CheckPathPart(part)
		if !v {
			return false
		}
	}
	return true
}

func CheckPathPart(part string) (ok bool, isVarData []string) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.helpers.CheckPathPart")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] part=%s", part)
	startCount := strings.Count(part, "{")
	endCount := strings.Count(part, "}")
	if startCount == 0 && endCount == 0 {
		return true, nil
	}

	if startCount%2 != 0 || endCount%2 != 0 {
		logging.Warningf("Bad Path Part [%s] - Bad Bracket Setup", part)
		return false, nil
	}

	if startCount > 2 || endCount > 2 {
		logging.Warningf("Bad Path Part [%s] - Can only have 1 variable for each path part", part)
		return false, nil
	}

	if !strings.HasPrefix(part, "{{") {
		logging.Warningf("Bad Path Part [%s] - Prefix Brackets {{ not at start of section", part)
		return false, nil
	}

	if !strings.HasSuffix(part, "}}") {
		logging.Warningf("Bad Path Part [%s] - Suffix Brackets }} not at end of section", part)
		return false, nil
	}

	splitPart := strings.Split(strings.ReplaceAll(strings.ReplaceAll(part, "{", ""), "}", ""), ":")

	if len(splitPart) != 2 {
		logging.Warningf("Bad Path Part [%s] - Part not in correct format should be {{[number|string]:[valiable name]}}", part)
		return false, nil
	}

	if matched, _ := regexp.Match(`number|string`, []byte(splitPart[0])); !matched {
		logging.Warningf("Bad Path Part [%s] - variable Type not 'number' or 'string'", part)
		return false, nil
	}

	if matched, _ := regexp.Match(`[a-zA-Z0-9]`, []byte(splitPart[1])); !matched {
		logging.Warningf("Bad Path Part [%s] - variable value has no name", part)
		return false, nil
	}

	return true, splitPart
}
