package util

import (
	"fmt"
	"net"
	"regexp"
	"time"
)

//RandomStr 获取一个随机字符串
func RandomStr() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// LocalIP 获取机器的IP
func LocalIP() string {
	info, _ := net.InterfaceAddrs()
	for _, addr := range info {
		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}
	return "0.0.0.0"
}

func IfFunc(flag bool, a interface{}, fun func() string) interface{} {
	if flag {
		return a
	} else {
		return fun()
	}
}

func If(flag bool, a interface{}, b interface{}) interface{} {
	if flag {
		return a
	} else {
		return b
	}
}

/** 校验 ip 有效性。 */
func IsIP(ip string) (b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}
	return true
}
