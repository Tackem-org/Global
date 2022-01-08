package helpers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	str2duration "github.com/xhit/go-str2duration/v2"
)

func IntSliceToStringSlice(in []int) []string {
	out := make([]string, len(in))
	for i, v := range in {
		out[i] = fmt.Sprint(v)
	}
	return out
}

func StringSliceToIntSlice(in []string) []int {
	out := make([]int, len(in))
	for i, v := range in {
		out[i], _ = strconv.Atoi(v)
	}
	return out
}

func InterfaceSliceToStringSlice(in []interface{}) []string {
	out := make([]string, len(in))
	for i, v := range in {
		out[i] = fmt.Sprint(v)
	}
	return out
}

func InterfaceSliceToIntSlice(in []interface{}) []int {
	out := make([]int, len(in))
	for i, v := range in {
		out[i], _ = strconv.Atoi(v.(string))
	}
	return out
}

func StringToIntSlice(in string) []int {
	tmp := strings.Split(in, ",")
	out := make([]int, len(tmp))
	for i, v := range tmp {
		out[i], _ = strconv.Atoi(v)
	}
	return out
}

func StringToStringSlice(in string) []string {
	out := strings.Split(in, ",")
	return out
}

func StringToStringMapString(in string) map[string]string {
	out := map[string]string{}
	json.Unmarshal([]byte(in), &out)
	return out
}

func StringToStringMapStringSlice(in string) map[string][]string {
	out := map[string][]string{}
	json.Unmarshal([]byte(in), &out)
	return out
}

func StringToStringMap(in string) map[string]interface{} {
	out := map[string]interface{}{}
	json.Unmarshal([]byte(in), &out)
	return out
}

func StringToDuration(in string) time.Duration {
	if f, err := strconv.ParseFloat(in, 64); err == nil {
		return time.Duration(f)
	} else if i, err := strconv.Atoi(in); err == nil {
		return time.Duration(i)
	} else if d, err := str2duration.ParseDuration(in); err == nil {
		return d
	}
	return 0
}

func DurationToString(in time.Duration) string {

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
