// +build windows
//运行时请带上参数GOARCH=386

package mylwsoft

import (
	"github.com/go-ole/go-ole/oleutil"
)



func (com *LwSoft) EnumWindow(title ,class string) string {
	ret, _ := oleutil.CallMethod(com.lw, "EnumWindow", title,class)
	return ret.ToString()
}
func (com *LwSoft) FindWindow(title,class string) int32 {
	ret, _ := oleutil.CallMethod(com.lw, "FindWindow",title,class)
	return int32(ret.Val)
}
func (com *LwSoft) FindSonWindow(fhwnd int,title, class string) int32 {
	ret, _ := oleutil.CallMethod(com.lw, "FindSonWindow", fhwnd,title ,class)
	return int32(ret.Val)
}

func (com *LwSoft) SetWindowSize(hwnd,width,height int) int32 {
	ret, _ := oleutil.CallMethod(com.lw, "SetWindowSize",hwnd,width,height)
	return int32(ret.Val)
}
func (com *LwSoft) MoveWindow(hwnd,x,y,width,height int) int32 {
	ret, _ := oleutil.CallMethod(com.lw, "MoveWindow",hwnd,x,y,width,height)
	return int32(ret.Val)
}

func (com *LwSoft) GetWindowVertex(hwnd int) int32 {
	ret, _ := oleutil.CallMethod(com.lw, "GetWindowVertex",hwnd)
	return int32(ret.Val)
}
func (com *LwSoft) X() int32 {
	ret, _ := oleutil.CallMethod(com.lw, "x")
	return int32(ret.Val)
}
func (com *LwSoft) Y() int32 {
	ret, _ := oleutil.CallMethod(com.lw, "y")
	return int32(ret.Val)
}
func (com *LwSoft) BindWindow(hwnd,display,mouse,keypad,added,mode int) int32 {
	ret, _ := oleutil.CallMethod(com.lw, "BindWindow", hwnd,display,mouse,keypad,added,mode)
	return int32(ret.Val)
}

func (com *LwSoft) UnBindWindow() int32 {
	ret, _ := oleutil.CallMethod(com.lw, "UnBindWindow")
	return int32(ret.Val)
}

func (com *LwSoft) LockMove() int32 { //禁止改变尺寸
	ret, _ := oleutil.CallMethod(com.lw, "LockMove")
	return int32(ret.Val)
}

func (com *LwSoft) CreateFoobar(fhwnd,x,y,w,h,canMove,sizeBox int,balckColor,borderColor string) int32 {
	ret, _ := oleutil.CallMethod(com.lw, "CreateFoobar", fhwnd,x,y,w,h,canMove,sizeBox,balckColor,borderColor)
	return int32(ret.Val)
}
func (com *LwSoft) FoobarSetFont(hwnd int,color,balckColor string,size,Format,Bold,Underline,Italic int) int32 {
	ret, _ := oleutil.CallMethod(com.lw, "FoobarSetFont", hwnd,color,balckColor,size,Format,Bold,Underline,Italic)
	return int32(ret.Val)
}
func (com *LwSoft) FoobarAddText(hwnd int,text string) int32 {
	ret, _ := oleutil.CallMethod(com.lw, "FoobarAddText", hwnd,text)
	return int32(ret.Val)
}
func (com *LwSoft) FoobarClearText(hwnd int) int32 {
	ret, _ := oleutil.CallMethod(com.lw, "FoobarClearText", hwnd)
	return int32(ret.Val)
}
func (com *LwSoft) FoobarClose(hwnd int) int32 {
	ret, _ := oleutil.CallMethod(com.lw, "FoobarClose", hwnd)
	return int32(ret.Val)
}