package helpers

import (
	"math/rand"
	"time"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
)

func RandStr(length int) string {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.helpers.RandStr")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] length=%d", length)
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	r := make([]rune, length)
	for i := range r {
		r[i] = letters[rand.Intn(len(letters))]
	}
	key := string(r)
	if key == "" {
		return RandStr(length)
	}
	return key
}
