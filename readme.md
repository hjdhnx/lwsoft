# golang 乐玩插件完善封装

# 源码克隆方法
` ` `
git clone https://github.com/hjdhnx/lwsoft.git
` ` `

## mylwsof 包说明
| 包内文件英文名        | 文件中文意思    |  作用  |
| --------    | -----: | :----:  |
| init.go      | 乐玩com初始化模块  |   用于初始化ole库并创建乐玩对象   |
| bitcolor.go      | 找图找色模块  |   用于windows窗口内找图找色   |
| mouse.go        |   鼠标操作模块 |   用于模拟鼠标的操作如鼠标左键单击，移动  |
| window.go        |  窗口操作模块  |   用于窗口设置，窗口绑定等操作  |


## mywin32api 包说明


| 包内文件英文名        | 文件中文意思    |  作用  |
| --------    | -----: | :----:  |
| mouse.go      | win32api的鼠标操作集合  |   调用win32api操作鼠标   |
| window.go      | win32api的窗口操作集合   |   调用win32api窗口   |


## main.go说明

*此文件是基于雷电模拟器里运行的游戏梦幻捉妖师的主线任务示例*