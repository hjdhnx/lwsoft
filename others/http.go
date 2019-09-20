// +build windows
//运行时请带上参数GOARCH=386

package others

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func Request(url string) string {
	request, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		return string(body)
	} else {
		return ""
	}
}

func JsonStr2map(jsonString string) map[string]string { //类json字符串转map
	jsonString = strings.Replace(jsonString, "\n", "", -1)
	jsonString = strings.Replace(jsonString, " ", "", -1)
	updateInfo := make(map[string]string)
	key_valueArray := strings.Split(jsonString, ",")
	var key, value string
	for _, key_value := range key_valueArray {
		linkArray := strings.Split(key_value, "':'")
		key = strings.Replace(linkArray[0], "'", "", -1)
		value = strings.Replace(linkArray[1], "'", "", -1)
		updateInfo[key] = value
	}
	return updateInfo
}

func GetMapKeys(m map[string]string) []string { //获取map里面的key
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率较高
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Type(v interface{}) string { //判断变量V的类型
	return fmt.Sprintf("%T", v)
}

func F2s(num float32) string { //小数转字符串
	strNum := strconv.FormatFloat(float64(num), 'E', -1, 32)
	return strNum
}
func S2f(strNum string) float32 { //字符串转小数(非数值字符串会得到0)
	num, _ := strconv.ParseFloat(strNum, 32)
	return float32(num)
}
func I2f(num int64) string { //数值优化显示
	var back string
	n := int(num)
	if n > G10(3) && n < G10(4) {
		back = strconv.Itoa(int(n/G10(3))) + " K"
	} else if n > G10(4) && n < G10(7) {
		back = strconv.Itoa(int(n/G10(4))) + " W"
	} else if n > G10(7) && n < G10(8) {
		back = strconv.Itoa(int(n/G10(7))) + " kw"
	} else if n > G10(8) {
		back = strconv.Itoa(int(n/G10(8))) + " E"
	} else {
		back = strconv.Itoa(n)
	}
	return back
}

func G10(n int) int { //生成10的多次方
	value := math.Pow10(n)
	return int(value)
}
func I2f2(num int64) string { //文件大小优化显示
	var back string
	n := int(num)
	if n > G10(3) && n < G10(6) {
		back = strconv.Itoa(int(n/G10(3))) + " K"
	} else if n > G10(6) && n < G10(9) {
		back = strconv.Itoa(int(n/G10(6))) + " M"
	} else if n > G10(9) && n < G10(12) {
		back = strconv.Itoa(int(n/G10(9))) + " G"
	} else if n > G10(12) {
		back = strconv.Itoa(int(n/G10(12))) + " T"
	} else {
		back = strconv.Itoa(n)
	}
	return back
}

func Httpdownload(url, savepath string) {
	var (
		buf     = make([]byte, 32*1024)
		written int64 //已写入的数据
		lastWtn int64 //用来判断还有没有写入 如果用EOF也可以 用这个可以得到speed
	)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	_, filename, _ := GetFile(url)
	Checkdir(savepath)
	f, err := os.Create(savepath + "/" + filename)
	if err != nil {
		panic(err)
	}
	lenth := res.ContentLength
	defer res.Body.Close()
	lastWtn = -1
	for {
		nr, _ := res.Body.Read(buf)
		if nr > 0 {
			nw, ew := f.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				fmt.Println("下载错误")
				return
			}
			if nr != nw {
				err = io.ErrShortBuffer
				fmt.Println("下载错误")
				return
			}
		}
		if written == lastWtn {
			goto LOOP
		}
		lastWtn = written
		fmt.Printf("\r已下载 %.2f %%  %sb/%sb ", float32(lastWtn)/float32(lenth)*100, I2f2(lastWtn), I2f2(lenth))
	}
LOOP:
	os.Stdout.Sync()
	fmt.Println("\n下载完毕")
}
