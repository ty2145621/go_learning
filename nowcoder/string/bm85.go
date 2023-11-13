package string

import (
	"regexp"
)

/**
BM85 验证IP地址
题目
题解(71)
讨论(102)
排行
面经new
中等  通过率：27.43%  时间限制：1秒  空间限制：256M
知识点
字符串
描述
编写一个函数来验证输入的字符串是否是有效的 IPv4 或 IPv6 地址
*/

func solve(IP string) string {
	// write code here
	regIpv4 := regexp.MustCompile(`^((\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5])\.){3}(\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5])$`)
	regIpv6 := regexp.MustCompile(`^([0-9a-fA-F]{1,4}:){7}([0-9a-fA-F]{1,4})$`)
	regIpv6Null := regexp.MustCompile(`(^|:)00`)
	if regIpv4.MatchString(IP) {
		return "IPv4"
	}
	if regIpv6.MatchString(IP) && !regIpv6Null.MatchString(IP) {
		return "IPv6"
	}
	return "Neither"
}
