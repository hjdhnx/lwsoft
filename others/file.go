// +build windows
//运行时请带上参数GOARCH=386

package others

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//创建新文本文件
func Newfile(path, content string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("文件创建失败：", err)
	}
	defer file.Close()
	file.WriteString(content)
}

//读文本文件内容
func Readfile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("文件打开失败：", err)
	}
	defer file.Close()
	buf := make([]byte, 1024)
	n, _ := file.Read(buf)
	return string(buf[:n])
}

//运行程序
func CallEXE(programeName string, canshu string) error {
	fmt.Println("CallEXE开始执行  " + programeName + " " + canshu)
	arg := []string{canshu}
	fmt.Println("---"+programeName+"---", arg)
	strPath := programeName
	cmd := exec.Command(strPath, arg...)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// 获取当前exe运行目录
func GetCurrentPath() string {
	s, _ := exec.LookPath(os.Args[0])
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

//注册com
func Regsvr32(dllpath string) string {
	command := "cd /d %~dp0\nregsvr32 " + Abspath(dllpath) + " /s"
	return command
}

//相对路径转绝对路径
func Abspath(path string) string {
	pathA, _ := filepath.Abs(path)
	return pathA
}

//注册乐玩
func Regsvrlw(dllpath, batpath string) error {
	content := Regsvr32(dllpath)
	Newfile(batpath, content)
	err := CallEXE(Abspath(batpath), "")
	if err != nil {
		return err
	} else {
		return nil
	}
}
func Put() {
	fmt.Printf("任意键继续...")
	input := ""
	fmt.Scanf("%s", &input)
}
