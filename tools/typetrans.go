package tools

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	
	"golang.org/x/crypto/bcrypt"
)

//string to int
func StrToInt(index string) int {
	result,_:= strconv.Atoi(index)
	return result
}
//str转int64
func StrToInt64(str string) (int64, error) {
	var (
		num int64
		err error
	)
	num, err = strconv.ParseInt(str, 10, 64)
	if err != nil {
		return num, err
	}
	
	return num, nil
}

//校验hash密码
func CompareHashAndPassword(e string, p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		return false, err
	}
	return true, nil
}

//float64 To string
func Float64ToString(e float64) string {
	return strconv.FormatFloat(e, 'E', -1, 64)
}
//int To string
func IntToString(e int) string {
	return strconv.Itoa(e)
}
//int64 To string
func Int64ToString(e int64) string {
	return strconv.FormatInt(e, 10)
}

//map to json
func MapToJson(input map[string]interface{}) string {
	data,_ := json.Marshal(input)
	return string(data)
}

//[]rune to string
func RuneToStr(r []rune) string {
	return string(r)
}
//string to []rune
func StrToRune(s string) []rune {
	return []rune(s)
}

//bool to string
func BoolToStr(b bool) string {
	//todo :bool to string
	sBool := strconv.FormatBool(true) //方法1
	return sBool
}

//string to bool
func StrToBool(str string) bool {
	//todo :string to bool
	//接受 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False 等字符串；
	//其他形式的字符串会返回错误
	b, _ := strconv.ParseBool("1")
	return b
}

//复杂结构转为string，可转类型包含map、struct
func StructToString(data interface{}) (string, error) {
	d, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	
	return string(d), nil
}

//去除数组中的空元素
func RemovEmpty(str []string) []string {
	var result []string = make([]string, 0)
	for i, _ := range str {
		if str[i] != "" {
			result = append(result, str[i])
		} else {
			continue
		}
	}
	
	return result
}
//检测字符串是否包含汉字,包含为true，不包含为false
func CheckString(str string) bool {
	reg := regexp.MustCompile(`[\p{Han}]+`)
	result := reg.FindAllString(str, -1)
	if len(result) >= 1 {
		return true
	}
	return false
}

//检测字符串是否包含阿拉伯数字,包含为true，不包含为false
func CheckInt(str string) bool {
	reg := regexp.MustCompile(`[0-9]+`)
	result := reg.FindAllString(str, -1)
	if len(result) != 0 {
		return true
	}
	return false
}
//检测数组中是否存在某元素，如存在，返回true,，不存在返回false
func CheckSlcHasStr(a string, b []string) bool {
	for _, v := range b {
		if strings.Contains(v, a) {
			return true
		}
	}
	
	return false
}

//取百分比,n表示取几位小数
func FloatRound(f float64, n int) float64 {
	format := "%." + strconv.Itoa(n) + "f"
	res, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return res
}