// +build windows
//运行时请带上参数GOARCH=386

package mylwsoft

import "github.com/go-ole/go-ole/oleutil"

func (com *LwSoft) SetUAC(enable int) int32 { //0关 1开
	ret, _ := oleutil.CallMethod(com.lw, "SetUAC", enable)
	return int32(ret.Val)
}
