// +build windows
//运行时请带上参数GOARCH=386

package mylwsoft

import (
	"github.com/go-ole/go-ole/oleutil"
	"time"
)

func (com *LwSoft) MoveTo(x,y int) int32 {
	ret, _ := oleutil.CallMethod(com.lw, "MoveTo", x,y)
	return int32(ret.Val)
}
func (com *LwSoft) LeftClick() int32 {
	ret, _ := oleutil.CallMethod(com.lw, "LeftClick")
	return int32(ret.Val)
}
func (com *LwSoft) MoveClick(x,y int) int32 {
	ret, _ := oleutil.CallMethod(com.lw, "MoveTo", x,y)
	time.Sleep(300 * time.Millisecond)
	ret, _ = oleutil.CallMethod(com.lw, "LeftClick")
	time.Sleep(300 * time.Millisecond)
	ret, _ = oleutil.CallMethod(com.lw, "LeftClick")
	return int32(ret.Val)
}
