package valid

import (
	"strconv"
	"strings"
	"regexp"
	"time"
)

var idCardWeight = [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
var idCardCheckBit = [11]string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}
var idCardProvinces = []int{11, 12, 13, 14, 15, 21, 22, 23, 31, 32, 33, 34, 35, 36, 37, 41, 42, 43, 44, 45, 46, 50, 51, 52, 53, 54, 61, 62, 63, 64, 65, 71, 81, 82}

const ()

/**
1.将前面的身份证号码17位数分别乘以不同的系数。 从第一位到第十七位的系数分别为：7 9 10 5 8 4 2 1 6 3 7 9 10 5 8 4 2
2.将这17位数字和系数相乘的结果相加。
3.用加出来和除以11，看余数是多少？
4.余数只可能有0 1 2 3 4 5 6 7 8 9 10这11个数字。 其分别对应的最后一位身份证的号码为1 0 X 9 8 7 6 5 4 3 2。
5.通过上面得知如果余数是2，就会在身份证的第18位数字上出现罗马数字的Ⅹ。 如果余数是10，身份证的最后一位号码就是2。
*/
func ValidIdcard(idcard string) (bool, string) {
	result := true
	cause := ""
	if len(idcard) != 18 {
		result = false
		cause = "身份证必须为18位"
		return result, cause
	}
	rb,er1:=regexp.MatchString("^\\d{17}(\\d|X){1}$",idcard)
	if(!rb){
		result = false
		cause = "身份证格式错误"
		return result, cause
	}
	if(er1 != nil) {
		result = false
		cause = "身份证匹配错误：" + er1.Error()
		return result, cause
	}

	havePro := false
	for _, v := range idCardProvinces {

		p,_:=strconv.Atoi(idcard[0:2])
		if p==v{
			havePro=true
		}
	}
	if !havePro{
		result = false
		cause = "身份证省份错误"
		return result, cause
	}

	idate:=idcard[6:14]
	itime,er2:=time.Parse("20060102",idate)
	if er2!=nil || itime.After(time.Now()) || itime.AddDate(200,0,0).Before(time.Now()) {
		result = false
		cause = "身份证日期错误"
		return result, cause
	}

	cbt := CalcCheckBit(idcard)
	if strings.Compare(cbt, idcard[17:]) != 0 {
		result = false
		cause = "身份证校验位错误"
	}
	return result, cause
}

//计算身份证校验位
func CalcCheckBit(idcard string) string {
	sum := 0
	for i := 0; i < 17; i++ {
		r1, _ := strconv.Atoi(idcard[i : i+1])
		sum += r1 * idCardWeight[i]
	}
	return idCardCheckBit[sum%11]
}
