package valid

import (
	"testing"
)

func TestValidIdcard(t *testing.T) {
	pr1 := "37098220070721770X"
	pr2 := "37098220070721770X"
	pe1 := "370982230707217707"
	pe2 := "370982230707217707"
	r1:= ValidIdcard(pr1)
	r2:= ValidIdcard(pr2)
	e1:= ValidIdcard(pe1)
	e2:= ValidIdcard(pe2)
	if !r1 || !r2 {
		t.Error("ValidIdcard()正确身份证未通过。")
	}
	if e1 || e2 {
		t.Error("ValidIdcard()错误身份证通过。")
	}
}
func TestValidChinese(t *testing.T) {
	pr1 := "中国"
	pr2 := "饕餮"
	pe1 := "Ab1"
	pe2 := "中国！"
	r1 := ValidChinese(pr1)
	r2 := ValidChinese(pr2)
	e1 := ValidChinese(pe1)
	e2 := ValidChinese(pe2)
	if !r1 || !r2 {
		t.Error("ValidChinese()正确未通过。")
	}
	if e1 || e2 {
		t.Error("ValidChinese()错误通过。")
	}
}
func TestValidIpv4(t *testing.T) {
	pr1 := "192.168.0.1"
	pr2 := "123.125.114.144"
	pe1 := "192.257.0.3"
	pe2 := "122.125.122.22.3"
	r1 := ValidIpv4(pr1)
	r2 := ValidIpv4(pr2)
	e1 := ValidIpv4(pe1)
	e2 := ValidIpv4(pe2)
	if !r1 || !r2 {
		t.Error("ValidIpv4()正确IPV4未通过。")
	}
	if e1 || e2 {
		t.Error("ValidIpv4()错误IPV4通过。")
	}
}
func TestValidIpv6(t *testing.T) {
	pr1 := "0000:0000:0000:0000:0000:ffff:c0a8:5909"
	pr2 := "2001:0DB8:0000:0000:0000:0000:1428:0000"
	pe1 := "192.257.0.3"
	pe2 := "2001:0DG8:0000:0000:0000:0000:1428:00003"
	r1 := ValidIpv6(pr1)
	r2 := ValidIpv6(pr2)
	e1 := ValidIpv6(pe1)
	e2 := ValidIpv6(pe2)
	if !r1 || !r2 {
		t.Error("ValidIpv6()正确IPV6未通过。")
	}
	if e1 || e2 {
		t.Error("ValidIpv6()错误IPV6通过。")
	}
}
func TestValidIp(t *testing.T) {
	pr1 := "0000:0000:0000:0000:0000:ffff:c0a8:5909"
	pr2 := "123.125.114.144"
	pe1 := "122.125.122.22.3"
	pe2 := "2001:0DG8:0000:0000:0000:0000:1428:00003"
	r1 := ValidIp(pr1)
	r2 := ValidIp(pr2)
	e1 := ValidIp(pe1)
	e2 := ValidIp(pe2)
	if !r1 || !r2 {
		t.Error("ValidIp()正确IP未通过。")
	}
	if e1 || e2 {
		t.Error("ValidIp()错误IP通过。")
	}
}
func TestValidMail(t *testing.T) {
	pr1 := "77558@qq.com"
	pr2 := "zhangF@163.com"
	pe1 := "77885@"
	pe2 := "aadf@1111"
	r1 := ValidMail(pr1)
	r2 := ValidMail(pr2)
	e1 := ValidMail(pe1)
	e2 := ValidMail(pe2)
	if !r1 || !r2 {
		t.Error("ValidMail()正确mail未通过。")
	}
	if e1 || e2 {
		t.Error("ValidMail()错误mail通过。")
	}
}

func BenchmarkValidIdcard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ValidIdcard("37098220070721770X")
	}
}
