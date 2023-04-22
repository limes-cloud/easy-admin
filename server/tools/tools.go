package tools

import (
	"crypto/md5"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math"
	"math/rand"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func RandomCode(len int) string {
	//初始化随机数
	rand.Seed(time.Now().Unix())
	var code = rand.Intn(int(math.Pow10(len)) - int(math.Pow10(len-1)))
	return strconv.Itoa(code + int(math.Pow10(len-1)))
}

func ParsePwd(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CompareHashPwd(p1, p2 string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p1), []byte(p2)) == nil
}

type ListType interface {
	~string | ~int | ~int64 | ~[]byte | ~rune | ~float64
}

func InList[ListType comparable](list []ListType, val ListType) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetReflectElem(val reflect.Value) reflect.Value {
	for val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}

func HideStar(str string) (result string) {
	if str == "" {
		return "***"
	}

	if strings.Contains(str, "@") {
		// 邮箱
		res := strings.Split(str, "@")
		if len(res[0]) < 3 {
			resString := "***"
			result = resString + "@" + res[1]
		} else {
			res2 := Substr2(str, 0, 3)
			resString := res2 + "***"
			result = resString + "@" + res[1]
		}
		return result
	} else {
		reg := `^1[0-9]\d{9}$`
		rgx := regexp.MustCompile(reg)
		mobileMatch := rgx.MatchString(str)
		if mobileMatch {
			// 手机号
			result = Substr2(str, 0, 3) + "****" + Substr2(str, 7, 11)
		} else {
			nameRune := []rune(str)
			lens := len(nameRune)
			if lens <= 1 {
				result = "***"
			} else if lens == 2 {
				result = string(nameRune[:1]) + "*"
			} else if lens == 3 {
				result = string(nameRune[:1]) + "*" + string(nameRune[2:3])
			} else if lens == 4 {
				result = string(nameRune[:1]) + "**" + string(nameRune[lens-1:lens])
			} else if lens > 4 {
				result = string(nameRune[:2]) + "***" + string(nameRune[lens-2:lens])
			}
		}
		return
	}
}

func Substr2(str string, start int, end int) string {
	rs := []rune(str)
	return string(rs[start:end])
}

func Md5(id any) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprint(id))))
}
