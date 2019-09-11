// +build windows
//运行时请带上参数GOARCH=386

package mywin32api

import (
	"github.com/lxn/win"
	"strconv"
	"syscall"
)

func GetWindow(hwnd int)int32{
	//value_int,err:=strconv.Atoi(string) string2int
	//str:=strconv.Itoa(value_int)  int2string
	hwnd2, _ := strconv.Atoi(strconv.Itoa(hwnd))
	editHwnd := win.GetWindow(win.HWND(hwnd2), 5)
	return int32(editHwnd)
}

func MessageBox(hwnd int,title,msg string, uType uint32)int32{
	hwnd2, _ := strconv.Atoi(strconv.Itoa(hwnd))
	ret := win.MessageBox(win.HWND(hwnd2), syscall.StringToUTF16Ptr(msg), syscall.StringToUTF16Ptr(title),uType)
	return int32(ret)
}

func Confirm(title,msg string)int32{
	ret := win.MessageBox(win.HWND(0), syscall.StringToUTF16Ptr(msg), syscall.StringToUTF16Ptr(title),win.MB_YESNO|win.MB_ICONINFORMATION)
	return int32(ret)
}

func MessageBoxE(title,msg string)int32{
	ret := win.MessageBox(win.HWND(0), syscall.StringToUTF16Ptr(msg), syscall.StringToUTF16Ptr(title),0)
	return int32(ret)
}
