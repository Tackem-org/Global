package helpers

import (
	"encoding/json"
	"errors"
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

func InterfaceSliceToIntSlice(in []interface{}) ([]int, error) {
	//TODO make this see what each type is in the slice and deal with it from there
	// instead of only accepting string get it to accept int too
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

func StringToIntSlice(in string) ([]int, error) {
	splitString := strings.Split(in, ",")
	out := make([]int, len(splitString))
	for i, v := range splitString {
		tmp, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		out[i] = tmp
	}
	return out, nil
}

func StringToStringSlice(in string) []string {
	out := strings.Split(in, ",")
	return out
}

func StringToStringMap[T string | []byte](in T) (map[string]interface{}, error) {
	out := map[string]interface{}{}
	err := json.Unmarshal([]byte(in), &out)
	if err != nil {
		return nil, err
	}
	out = mapRecursion(out)
	return out, nil
}

func mapRecursion(out map[string]interface{}) map[string]interface{} {
	for key, val := range out {
		switch x := val.(type) {
		default:
			continue
		case float64:
			if i, err := strconv.Atoi(fmt.Sprintf("%.0f", x)); err == nil && float64(i) == x {
				out[key] = i
			}
		case map[string]interface{}:
			out[key] = mapRecursion(x)
		}
	}

	return out
}

func StringToDuration(in string) (time.Duration, error) {
	if f, err := strconv.ParseFloat(in, 64); err == nil {
		return time.Duration(f), nil
	} else if d, err := str2duration.ParseDuration(in); err == nil {
		return d, nil
	}
	return 0, errors.New("cannot convert value to duration")
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

func MapStringInterfaceToMapStringString(in map[string]interface{}) map[string]string {
	out := map[string]string{}
	for k, v := range in {
		out[k] = fmt.Sprint(v)
	}
	return out
}
