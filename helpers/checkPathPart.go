package helpers

import (
	"regexp"
	"strings"

	"github.com/Tackem-org/Global/logging"
)

func CheckPath(path string) bool {
	if strings.HasPrefix(path, "/") {
		path = strings.Replace(path, "/", "", 1)
	}
	if strings.Count(path, "{") == 0 && strings.Count(path, "}") == 0 {
		return true
	}

	s := strings.Split(path, "/")
	for _, part := range s {
		v, _ := CheckPathPart(part)
		if !v {
			return false
		}
	}
	return true
}

func CheckPathPart(part string) (ok bool, isVarData []string) {
	startCount := strings.Count(part, "{")
	endCount := strings.Count(part, "}")
	if startCount == 0 && endCount == 0 {
		return true, nil
	}

	if startCount%2 != 0 || endCount%2 != 0 {
		logging.Warning("Bad Path Part [%s] - Bad Bracket Setup", part)
		return false, nil
	}

	if startCount > 2 || endCount > 2 {
		logging.Warning("Bad Path Part [%s] - Can only have 1 variable for each path part", part)
		return false, nil
	}

	if !strings.HasPrefix(part, "{{") {
		logging.Warning("Bad Path Part [%s] - Prefix Brackets {{ not at start of section", part)
		return false, nil
	}

	if !strings.HasSuffix(part, "}}") {
		logging.Warning("Bad Path Part [%s] - Suffix Brackets }} not at end of section", part)
		return false, nil
	}

	splitPart := strings.Split(strings.ReplaceAll(strings.ReplaceAll(part, "{", ""), "}", ""), ":")

	if len(splitPart) != 2 {
		logging.Warning("Bad Path Part [%s] - Part not in correct format should be {{[number|string]:[valiable name]}}", part)
		return false, nil
	}
	if splitPart[0] != "number" && splitPart[0] != "string" {
		logging.Warning("Bad Path Part [%s] - variable Type not 'number' or 'string'", part)
		return false, nil
	}

	isAlphaNumeric := regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString
	if !isAlphaNumeric(splitPart[1]) {
		logging.Warning("Bad Path Part [%s] - variable value has no name", part)
		return false, nil
	}

	return true, splitPart
}
