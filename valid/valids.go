package valid

import "regexp"

const (
	PATTERN_MAIL          = "^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$"
	PATTERN_IPV4          = "^(((\\d{1,2})|(1\\d{2})|(2[0-4]\\d)|(25[0-5]))\\.){3}((\\d{1,2})|(1\\d{2})|(2[0-4]\\d)|(25[0-5]))$"
	PATTERN_VALID_IPV6    = "^([\\dA-Fa-f]{0,4}:){7}[\\dA-Fa-f]{0,4}$"
	PATTERN_VALID_CHINESE = "^[\u4e00-\u9fa5]+$"
)

//验证邮箱格式
func ValidMail(mail string) bool {
	r, e := regexp.MatchString(PATTERN_MAIL, mail)
	if e != nil {
		return false
	}
	return r
}

//是否为IPV4
func ValidIpv4(ipv4 string) bool {
	r, e := regexp.MatchString(PATTERN_IPV4, ipv4)
	if e != nil {
		return false
	}
	return r
}

//是否为IPV6
func ValidIpv6(ipv6 string) bool {
	r, e := regexp.MatchString(PATTERN_VALID_IPV6, ipv6)
	if e != nil {
		return false
	}
	return r
}

//是否为IP
func ValidIp(ip string) bool {
	r1 := ValidIpv4(ip)
	r2 := ValidIpv6(ip)
	return r1 || r2
}

//是否为汉语、汉字
func ValidChinese(ch string) bool {
	r, e := regexp.MatchString(PATTERN_VALID_CHINESE, ch)
	if e != nil {
		return false
	}
	return r
}
