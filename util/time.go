package util

import "time"

// 获取东八区 Location
func DongBaDistrict() *time.Location {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai")
	return cstSh
}

// 获取UTC Location
func UTC() *time.Location {
	var cstSh, _ = time.LoadLocation("UTC")
	return cstSh
}

func DongBaTime() *time.Time {
	localTime := time.Now().In(DongBaDistrict())
	return &localTime
}
