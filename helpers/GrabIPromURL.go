package helpers

import (
	"fmt"
	"net"
)

func GrabIPFromURL(in string) string {
	if addr := net.ParseIP(in); addr != nil {
		return in
	}

	addr, err := net.LookupIP(in)

	if err == nil && len(addr) > 0 {
		return fmt.Sprint(addr[0])
	}

	return ""
}
