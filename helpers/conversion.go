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
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.helpers.IntSliceToStringSlice")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] in=%v", in)
	out := make([]string, len(in))
	for i, v := range in {
		out[i] = fmt.Sprint(v)
	}
	return out
}

func StringSliceToIntSlice(in []string) []int {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.helpers.StringSliceToIntSlice")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] in=%v", in)
	out := make([]int, len(in))
	for i, v := range in {
		out[i], _ = strconv.Atoi(v)
	}
	return out
}

func InterfaceSliceToStringSlice(in []interface{}) []string {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.helpers.InterfaceSliceToStringSlice")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] in=%v", in)
	out := make([]string, len(in))
	for i, v := range in {
		out[i] = fmt.Sprint(v)
	}
	return out
}

func InterfaceSliceToIntSlice(in []interface{}) ([]int, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.helpers.InterfaceSliceToIntSlice")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] in=%v", in)
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
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.helpers.StringToIntSlice")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] in=%v", in)
	tmp := strings.Split(in, ",")
	out := make([]int, len(tmp))
	for i, v := range tmp {
		out[i], _ = strconv.Atoi(v)
	}
	return out
}

func StringToStringSlice(in string) []string {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.helpers.StringToStringSlice")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] in=%s", in)
	out := strings.Split(in, ",")
	return out
}

func StringToStringMap(in string) map[string]interface{} {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.helpers.StringToStringMap")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] in=%s", in)
	out := map[string]interface{}{}
	json.Unmarshal([]byte(in), &out)
	return out
}

func StringToDuration(in string) (time.Duration, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.helpers.StringToDuration")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] in=%s", in)
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
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.helpers.DurationToString")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] in=%s", in.String())

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
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.helpers.MapStringInterfaceToMapStringString")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] in=%v", in)
	out := map[string]string{}
	for k, v := range in {
		out[k] = fmt.Sprint(v)
	}
	return out
}
