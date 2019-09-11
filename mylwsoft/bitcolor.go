// +build windows
//运行时请带上参数GOARCH=386

package mylwsoft

import "github.com/go-ole/go-ole/oleutil"

func (com *LwSoft) FindMultiColor(x1, y1, x2, y2 int,first_color ,offset_color string,sim float32, dir int) int32{
	ret, _ := oleutil.CallMethod(com.lw, "FindMultiColor",x1, y1, x2, y2,first_color,offset_color,sim, dir,300,1,5,5)//找到点击偏移5
	return int32(ret.Val)
}

func (com *LwSoft) FindMultiColorEx(x1, y1, x2, y2 int,first_color ,offset_color string,sim float32, dir int) string {
	ret, _ := oleutil.CallMethod(com.lw, "FindMultiColorEx",x1, y1, x2, y2,first_color,offset_color,sim, dir)
	return ret.ToString()
}