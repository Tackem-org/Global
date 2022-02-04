package web

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Master/config"
	"github.com/Tackem-org/Master/data"

	"github.com/schollz/websocket"
	"github.com/schollz/websocket/wsjson"

	pb "github.com/Tackem-org/Proto/pb/config"
	pbc "github.com/Tackem-org/Proto/pb/registration"
)

func adminConfigWebsocketHandler(w http.ResponseWriter, r *http.Request) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.adminConfigWebsocketHandler(w http.ResponseWriter, r *http.Request)]")

	baseWebsocketHandler(w, r, func(ctx context.Context, w http.ResponseWriter, r *http.Request, c *websocket.Conn, v map[string]interface{}) bool {
		command := v["command"].(string)
		if command == "setConfig" {
			key := v["key"].(string)
			ci := data.GetConfigItemByKey(key)
			if ci == nil {
				logging.Warningf("Attempted to Change A Key With No Data: [%s]", key)
				wsjson.Write(ctx, c, map[string]string{
					"command": "ConfigSaveFailed",
					"key":     key,
					"error":   "Key Data Not Found",
				})
			} else if !config.InConfig(key) {
				logging.Warningf("Attempted to Change A Key not in the Config: [%s]", key)
				wsjson.Write(ctx, c, map[string]string{
					"command": "ConfigSaveFailed",
					"key":     key,
					"error":   "Key Config Not Found",
				})
			}
			pass, wrongType := true, false
			switch ci.GetType() {
			case pb.ValueType_Bool:
				if val, ok := v["value"].(bool); ok {
					pass = config.SetBool(key, val)
				} else {
					wrongType = true
				}
			case pb.ValueType_Duration:
				if str, ok := v["value"].(string); ok {
					val, err := helpers.StringToDuration(str)
					if err != nil {
						wsjson.Write(ctx, c, map[string]string{
							"command": "ConfigSaveFailed",
							"key":     key,
							"error":   err.Error(),
						})
						return true
					}
					pass = config.SetDuration(key, val)
				} else {
					wrongType = true
				}
			case pb.ValueType_Float64:
				if str, ok := v["value"].(string); ok {
					val, err := strconv.ParseFloat(str, 64)
					if err != nil {
						wsjson.Write(ctx, c, map[string]string{
							"command": "ConfigSaveFailed",
							"key":     key,
							"error":   err.Error(),
						})
						return true
					}
					pass = config.SetFloat64(key, val)
				} else {
					wrongType = true
				}
			case pb.ValueType_Int:
				if str, ok := v["value"].(string); ok {
					val, err := strconv.Atoi(str)
					if err != nil {
						wsjson.Write(ctx, c, map[string]string{
							"command": "ConfigSaveFailed",
							"key":     key,
							"error":   err.Error(),
						})
						return true
					}
					pass = config.SetInt(key, int(val))
				} else {
					wrongType = true
				}
			case pb.ValueType_Int32:
				if str, ok := v["value"].(string); ok {
					val, err := strconv.Atoi(str)
					if err != nil {
						wsjson.Write(ctx, c, map[string]string{
							"command": "ConfigSaveFailed",
							"key":     key,
							"error":   err.Error(),
						})
						return true
					}
					pass = config.SetInt32(key, int32(val))
				} else {
					wrongType = true
				}
			case pb.ValueType_Int64:
				if str, ok := v["value"].(string); ok {
					val, err := strconv.Atoi(str)
					if err != nil {
						wsjson.Write(ctx, c, map[string]string{
							"command": "ConfigSaveFailed",
							"key":     key,
							"error":   err.Error(),
						})
						return true
					}
					pass = config.SetInt64(key, int64(val))
				} else {
					wrongType = true
				}
			case pb.ValueType_IntSlice:
				if in, ok := v["value"].([]interface{}); ok {
					is, err := helpers.InterfaceSliceToIntSlice(in)
					if err != nil {
						wsjson.Write(ctx, c, map[string]string{
							"command": "ConfigSaveFailed",
							"key":     key,
							"error":   err.Error(),
						})
						return true
					}
					pass = config.SetIntSlice(key, is)
				} else {
					wrongType = true
				}
			case pb.ValueType_String:
				if str, ok := v["value"].(string); ok {
					pass = config.SetString(key, str)
				} else {
					wrongType = true
				}
			case pb.ValueType_StringSlice:
				if val, ok := v["value"].(string); ok {
					if ci.InputType == pbc.InputType_TextArea {
						pass = config.SetStringSlice(key, strings.Split(val, "\n"))
					} else {
						pass = config.SetStringSlice(key, strings.Split(val, ","))
					}
				} else if in, ok := v["value"].([]interface{}); ok {
					ss := helpers.InterfaceSliceToStringSlice(in)
					pass = config.SetStringSlice(key, ss)
				} else {
					wrongType = true
				}
			case pb.ValueType_Time:
				if val, ok := v["value"].(string); ok {
					var t time.Time
					var err error
					if inputType, ok := v["inputType"].(string); ok {
						switch inputType {
						case "date":
							t, err = time.Parse("2006-01-02", val)
						case "datetime-local":
							t, err = time.Parse("2006-01-02T15:04", val)
						case "time":
							t, err = time.Parse("15:04", val)
						default:
							t, err = time.Parse("2006-01-02T15:04:05", val)
						}
					} else {
						t, err = time.Parse("2006-01-02T15:04:05", val)
					}
					if err != nil {
						wsjson.Write(ctx, c, map[string]string{
							"command": "ConfigSaveFailed",
							"key":     key,
							"error":   err.Error(),
						})
						return true
					}
					pass = config.SetTime(key, t)
				} else {
					wrongType = true
				}
			case pb.ValueType_Uint:
				if val, ok := v["value"].(string); ok {
					val, err := strconv.ParseUint(val, 10, 64)
					if err != nil {
						wsjson.Write(ctx, c, map[string]string{
							"command": "ConfigSaveFailed",
							"key":     key,
							"error":   err.Error(),
						})
						return true
					}
					pass = config.SetUint(key, uint(val))
				} else {
					wrongType = true
				}
			case pb.ValueType_Uint32:
				if val, ok := v["value"].(string); ok {
					val, err := strconv.ParseUint(val, 10, 32)
					if err != nil {
						wsjson.Write(ctx, c, map[string]string{
							"command": "ConfigSaveFailed",
							"key":     key,
							"error":   err.Error(),
						})
						return true
					}
					pass = config.SetUint32(key, uint32(val))
				} else {
					wrongType = true
				}
			case pb.ValueType_Uint64:
				if val, ok := v["value"].(string); ok {
					val, err := strconv.ParseUint(val, 10, 64)
					if err != nil {
						wsjson.Write(ctx, c, map[string]string{
							"command": "ConfigSaveFailed",
							"key":     key,
							"error":   err.Error(),
						})
						return true
					}
					pass = config.SetUint64(key, uint64(val))
				} else {
					wrongType = true
				}
			}
			if wrongType {
				logging.Warningf("Attempted to Change A Key With Wrong Type: [%s] expecting %s Got %T", key, ci.Type.String(), v["value"])
				wsjson.Write(ctx, c, map[string]string{
					"command": "ConfigSaveFailed",
					"key":     key,
					"error":   "Wrong Data Type",
				})
			} else if pass {
				if ci.RestartRequired {
					config.SetRestartRequired(ci.Key)
				}
				wsjson.Write(ctx, c, map[string]string{
					"command": "ConfigSaved",
					"key":     key,
				})
			} else {
				logging.Errorf("Failed to change Config Var: %s", key)
				wsjson.Write(ctx, c, map[string]string{
					"command": "ConfigSaveFailed",
					"key":     key,
					"error":   "Config Failed To save Valid data",
				})
			}
			return true
		}
		return false
	})
}
