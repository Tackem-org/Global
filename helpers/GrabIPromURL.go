package helpers

import (
	"fmt"
	"net"
)

var LookUPIP func(host string) ([]net.IP, error) = net.LookupIP

func GrabIPFromURL(in string) string {
	if addr := net.ParseIP(in); addr != nil {
		return in
	}

	addr, _ := LookUPIP(in)
	return fmt.Sprint(addr[0])

}
