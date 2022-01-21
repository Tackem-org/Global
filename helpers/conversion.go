package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	str2duration "github.com/xhit/go-str2duration/v2"
)

func IntSliceToStringSlice(in []int) []string {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[helpers.IntSliceToStringSlice(in []int) []string] {in=%v}", in)
	out := make([]string, len(in))
	for i, v := range in {
		out[i] = fmt.Sprint(v)
	}
	return out
}

func StringSliceToIntSlice(in []string) []int {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[helpers.StringSliceToIntSlice(in []string) []int] {in=%v}", in)
	out := make([]int, len(in))
	for i, v := range in {
		out[i], _ = strconv.Atoi(v)
	}
	return out
}

func InterfaceSliceToStringSlice(in []interface{}) []string {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[helpers.InterfaceSliceToStringSlice(in []interface{}) []string] {in=%v}", in)
	out := make([]string, len(in))
	for i, v := range in {
		out[i] = fmt.Sprint(v)
	}
	return out
}

func InterfaceSliceToIntSlice(in []interface{}) ([]int, error) {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[helpers.InterfaceSliceToIntSlice(in []interface{}) []int] {in=%v}", in)
	out := make([]int, len(in))
	for i, v := range in {
		tmp, err := strconv.Atoi(v.(string))
		if err != nil {
			return nil, err
		}
		out[i] = tmp
	}
	return out, nil
}

func StringToIntSlice(in string) []int {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[helpers.StringToIntSlice(in string) []int] {in=%v}", in)
	tmp := strings.Split(in, ",")
	out := make([]int, len(tmp))
	for i, v := range tmp {
		out[i], _ = strconv.Atoi(v)
	}
	return out
}

func StringToStringSlice(in string) []string {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[helpers.StringToStringSlice(in string) []string] {in=%s}", in)
	out := strings.Split(in, ",")
	return out
}

func StringToStringMap(in string) map[string]interface{} {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[helpers.StringToStringMap(in string) map[string]interface{}] {in=%s}", in)
	out := map[string]interface{}{}
	json.Unmarshal([]byte(in), &out)
	return out
}

func StringToDuration(in string) (time.Duration, error) {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[helpers.StringToDuration(in string) time.Duration] {in=%s}", in)
	if f, err := strconv.ParseFloat(in, 64); err == nil {
		return time.Duration(f), nil
	} else if i, err := strconv.Atoi(in); err == nil {
		return time.Duration(i), nil
	} else if d, err := str2duration.ParseDuration(in); err == nil {
		return d, nil
	}
	return 0, errors.New("cannot convert value to duration")
}

func DurationToString(in time.Duration) string {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[helpers.DurationToString(in time.Duration) string] {in=%s}", in.String())

	f := in.String()
	if !strings.Contains(f, "h") {
		return f
	}

	s := strings.Split(f, "h")
	i, _ := strconv.Atoi(s[0])

	if i < 24 {
		return f
	}

	d, h := i/24, i%24
	if d < 7 {
		return fmt.Sprintf("%dd%dh%s", d, h, s[1])
	}
	w, d := d/7, d%7
	return fmt.Sprintf("%dw%dd%dh%s", w, d, h, s[1])
}

func MapStringInterfaceToMapStringString(in map[string]interface{}) map[string]string {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[helpers.MapStringInterfaceToMapStringString(in map[string]interface{}) map[string]string] {in=%v}", in)
	out := map[string]string{}
	for k, v := range in {
		out[k] = fmt.Sprint(v)
	}
	return out
}
