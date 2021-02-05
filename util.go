package distributedid

import (
	"errors"
	"net"
)

var (
	// ErrInvalidIP 无效ip
	ErrInvalidIP = errors.New("invalid ip")
)

func convertNode(ip string) (node int64, err error) {
	if len(ip) == 0 {
		return 0, ErrInvalidIP
	}
	ipv4 := net.ParseIP(ip).To4()

	if len(ipv4) < 4 {
		return 0, ErrInvalidIP
	}
	node = int64(ipv4[2])<<8 + int64(ipv4[3])

	if node == 0 {
		return 0, ErrInvalidIP
	}
	return
}
