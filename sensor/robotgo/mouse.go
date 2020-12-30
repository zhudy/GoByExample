//代码参考：https://github.com/go-vgo/robotgo/blob/master/examples
//$ go get github.com/go-vgo/robotgo
//报错: fatal error: X11/keysym.h: No such file or directory
//$ apt search Xlib.h
//sudo apt-get install libx11-dev ................. for X11/Xlib.h
//sudo apt-get install mesa-common-dev........ for GL/glx.h
//sudo apt-get install libglu1-mesa-dev ..... for GL/glu.h
//sudo apt-get install libxrandr-dev ........... for X11/extensions/Xrandr.h
//sudo apt-get install libxi-dev ................... for X11/extensions/XInput.h
//$ sudo apt-get install libx11-xcb-dev  自己加的
//$ sudo apt-get install libxkbcommon-x11-dev  自己加的
/*改用apt-get来查缺失文件
$ sudo apt-get install apt-file
$ sudo apt-file update
$ apt-file search "X11/extensions/XTest.h"	这个好使！
报告：libxtst-dev: /usr/include/X11/extensions/XTest.h
$ sudo apt install libxtst-dev
$ apt-file search "png.h"
$ sudo apt install libpng-dev
*/

package main

import (
	"github.com/go-vgo/robotgo"
)

func main() {
  robotgo.ScrollMouse(10, "up")
  robotgo.MouseClick("left", true)
  robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)
}
