# golang 乐玩插件完善封装

# 源码克隆方法
>` ` `
git clone https://github.com/hjdhnx/lwsoft.git
` ` `

## mylwsof 包说明
| 包内文件英文名        | 文件中文意思    |  作用  |
| --------    | -----: | :----:  |
| init.go      | 乐玩com初始化模块  |   用于初始化ole库并创建乐玩对象   |
| bitcolor.go      | 找图找色模块  |   用于windows窗口内找图找色   |
| mouse.go        |   鼠标操作模块 |   用于模拟鼠标的操作如鼠标左键单击，移动  |
| window.go        |  窗口操作模块  |   用于窗口设置，窗口绑定等操作  |
| system.go        |  系统操作模块  |   封装的乐玩系列系统操作如UAC控制  |


## mywin32api 包说明


| 包内文件英文名        | 文件中文意思    |  作用  |
| --------    | -----: | :----:  |
| mouse.go      | win32api的鼠标操作集合  |   调用win32api操作鼠标   |
| window.go      | win32api的窗口操作集合   |   调用win32api窗口   |

## others 包说明


| 包内文件英文名        | 文件中文意思    |  作用  |
| --------    | -----: | :----:  |
| mouse.go      | win32api的鼠标操作集合  |   调用win32api操作鼠标   |
| window.go      | win32api的窗口操作集合   |   调用win32api窗口   |


## dll 包说明

| 包内文件英文名        | 文件中文意思    |  作用  |
| --------    | -----: | :----:  |
| lw9.09.dll      | 乐玩最新com组件  |   深度封状底层模拟操作   |
| reg.bat      | 注册乐玩com批处理命令  |   快速注册com到系统   |

## tools 包说明

| 包内文件英文名        | 文件中文意思    |  作用  |
| --------    | -----: | :----:  |
| 乐玩助手.exe      | 乐玩助手标准版  |   图色可视化测试工具   |
| 乐玩助手9.09.exe      | 乐玩助手9.09加强版  |   图色可视化测试工具   |
| 偏色计算器.exe      | 偏色计算器  |   计算图片图色进行二值化   |
| 接口查看器.exe      | 接口查看器  |   查看com接口函数并支持可视化注册与卸载   |


## 跟目录文件说明

| 包内文件英文名        | 文件中文意思    |  作用  |
| --------    | -----: | :----:  |
| main.exe      | 主程序的二进制文件  |   将main.go便于成main.exe所存放的二进制文件   |

## main.go说明

*此文件是基于雷电模拟器里运行的游戏梦幻捉妖师的主线任务示例*