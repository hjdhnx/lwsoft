// +build windows
//运行时请带上参数GOARCH=386
package main

import (
	"./mylwsoft"
	"./mywin32api"
	"fmt"
	"github.com/lxn/win"
	"strconv"
	"strings"
	"time"
)

var state string

func runall(lw *mylwsoft.LwSoft, task string) {
	switch task {
	case "主线":
		a := lw.FindMultiColorEx(0, 485, 43, 526, "B82210-000000", "-2|5|E6C391-000000,3|-5|FCDA72-000000", 1, 0)
		b := strings.Split(a, ",")
		if len(b) == 3 {
			_, x, y := b[0], b[1], b[2]
			xpd, _ := strconv.Atoi(x)
			ypd, _ := strconv.Atoi(y)
			ret := lw.MoveClick(xpd, ypd)
			if ret == 1 {
				fmt.Println("点击了一次主线，接下来挑战")
				state = "挑战"
			} else {
				runall(lw, task)
			}
		}
	case "挑战":
		if lw.FindMultiColor(170, 546, 209, 568, "F0A91E-000000", "3|3|693000-000000,6|1|F7AF1F-000000", 1, 0) == 1 {
			fmt.Println("点击了一次挑战，接下来查询熔炼")
			state = "前往熔炼"
		}
	case "前往熔炼":
		if lw.FindMultiColor(163, 402, 222, 422, "B26B11-000000", "8|-2|FBB320-000000,5|-1|74380C-000000", 1, 0) == 1 {
			fmt.Println("点击了一次前往熔炼，接下来准备熔炼")
			state = "熔炼"
		} else {
			fmt.Println("无需熔炼，挑战中，完毕后继续主线")
			state = "主线"
		}
	case "熔炼":
		if lw.FindMultiColor(263, 496, 299, 516, "C98715-000000", "6|3|D28913-000000,7|1|FEB621-000000", 1, 0) == 1 {
			fmt.Println("点击了一次熔炼，接下熔炼3次")
			for i := 0; i < 3; i++ {
				lw.FindMultiColor(263, 496, 299, 516, "C98715-000000", "6|3|D28913-000000,7|1|FEB621-000000", 1, 0)
				time.Sleep(1000 * time.Millisecond)
			}
			fmt.Println("熔炼完毕，返回并准备挑战")
			lw.FindMultiColor(324, 543, 357, 576, "B16A3E-000000", "6|-6|B36B3B-000000,9|13|EDD6B8-000000", 1, 0)
			state = "挑战"
		}
	case "返回":
		lw.FindMultiColor(324, 543, 357, 576, "B16A3E-000000", "6|-6|B36B3B-000000,9|13|EDD6B8-000000", 1, 0)
	default:
		fmt.Println("未知的任务类型")
	}
}

func main() {
	//vcl.ShowMessage("hello  world")
	mywin32api.SleepIntS(1)
	win.GetWindow(win.HWND(0), 5)
	//mylwsoft.MessageBox(0,"系统提示","这是复杂的弹窗",win.MB_YESNO)
	ret := mywin32api.Confirm("系统提示", "是否停止程序?")
	if ret == 6 {
		fmt.Println("你点了确定,程序结束")
		return
	} else if ret == 7 {
		fmt.Println("你点了取消,输出hello world")
	}
	mywin32api.MessageBoxE("系统提示", "这是简单的弹窗\n你好,世界！")
	fmt.Println(ret)
	//ole.CoInitializeEx(0, 0)
	//defer ole.CoUninitialize()
	lw, err := mylwsoft.NewLwSoft()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("版本号 version:", lw.Ver())
	//hwndStr := lw.EnumWindow(0, "TheRender", "RenderWindow", 2)
	hwndStr := lw.EnumWindow("雷电模拟器-1", "")
	fmt.Println(hwndStr)
	arr := strings.Split(hwndStr, ",")
	if len(arr) < 1 {
		fmt.Println("没有雷电多开器窗口!")
		fmt.Println()
		//var a = [...]int{}
		return
	}
	fmt.Println("句柄数组:", arr)
	//fmt.Println("句柄 hwnd:", lw.FindWindow("TheRender","RenderWindow"))
	if len(arr) == 1 {
		hwnd, _ := strconv.Atoi(arr[0])
		if hwnd == 0 {
			fmt.Println("句柄查找失败")
			return
		}
		lw.GetWindowVertex(hwnd)
		lw.MoveWindow(hwnd, int(lw.X()), int(lw.Y()), 380, 710)
		//上面两条是初始化竖屏模拟器尺寸的

		hwnd = int(mywin32api.GetWindow(hwnd))
		lw.BindWindow(hwnd, 5, 1, 1, 0, 0)
		fmt.Println("绑定成功,句柄:", hwnd)
		state = "主线"
		for hwnd != 0 {
			//fmt.Println(state)
			runall(lw, state)
			time.Sleep(1500 * time.Millisecond)
		}
	}
}
