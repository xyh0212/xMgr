package x

import (
	"net"
	"strings"
)

func GetPubIP() (string, *Error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		xerr := XErr(err)
		LogErr(xerr)
		return "", xerr
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", nil
}

func IsLanIP(strIp string) bool {
	if strings.HasPrefix(strIp, "192.168.") {
		return true
	}
	if strings.HasPrefix(strIp, "127.0.0.") {
		return true
	}
	return false
}
