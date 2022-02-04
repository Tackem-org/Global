package web

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"

	"github.com/schollz/websocket"
	"github.com/schollz/websocket/wsjson"
)

func baseWebsocketHandler(w http.ResponseWriter, r *http.Request, systemFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request, c *websocket.Conn, v map[string]interface{}) bool) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.baseWebsocketHandler(w http.ResponseWriter, r *http.Request, systemFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request, c *websocket.Conn, v map[string]interface{}) bool)]")

	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		logging.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Hour*120000)
	defer cancel()

	for {
		var v map[string]interface{}
		err = wsjson.Read(ctx, c, &v)
		if err != nil {
			if websocket.CloseStatus(err) == websocket.StatusNoStatusRcvd {
				break
			}
			logging.Error(err.Error())
			c.Close(websocket.StatusInternalError, "internal error")
			break
		}

		command := v["command"].(string)
		if command == "close" || command == "quit" || command == "exit" {
			break
		}

		if systemFunc(ctx, w, r, c, v) {
			continue
		}

		logging.Error(fmt.Sprintf("ERROR web socket unknown command received: %v\n", v))
	}

	if websocket.CloseStatus(err) == websocket.StatusGoingAway {
		err = nil
	}
	c.Close(websocket.StatusNormalClosure, "")

}
