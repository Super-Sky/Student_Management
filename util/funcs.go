package util

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

func GetRandStr(length int) string {
	var (
		sideCar = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0',
			'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
			'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	)
	for i := 0; i < 3; i++ {
		var pwd = make([]byte, length)
		for j := 0; j < length; j++ {
			index := r.Int() % len(sideCar)
			pwd[j] = sideCar[index]
		}
		return string(pwd)
	}
	return ""
}

//类似python的in
func In(target string, StrArray []string) bool {
	sort.Strings(StrArray)
	index := sort.SearchStrings(StrArray, target)
	if index < len(StrArray) && StrArray[index] == target {
		return true
	}
	return false
}

// IsNum 字符串是否为数字
func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// TwoDimensionalToString 二维数组转字符串two-dimensional
func TwoDimensionalToString(listData [][]string) string {
	var (
		result string
	)
	for i := 0; i < len(listData); i++ {
		str := strings.Join(listData[i], "/")
		var build strings.Builder
		build.WriteString(result)
		build.WriteString(";")
		build.WriteString(str)
		result = build.String()
	}
	return result
}

func GetMd5pro(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func Base64Encode(message string) string {
	return base64.StdEncoding.EncodeToString([]byte(message))
}

func Base64DeEncode(encodedMessage string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encodedMessage)
	if err != nil {
		return "", err
	}
	return string(data), err
}
