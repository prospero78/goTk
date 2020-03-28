# goTk -- графика Tk для Go
goTk -- библиотека для Golang Tcl/Tk

	go get github.com/prospero78/goTk


### Установка Tcl/Tk

http://www.tcl-lang.org

* MacOS X, Windows

	https://www.activestate.com/activetcl/downloads

* Ubuntu

	$ sudo apt install tk-dev

* CentOS

	$ sudo yum install tk-devel

### Demo

   https://github.com/prospero78/goTk/demotk

### Sample
```go
package main

import (
	tk "github.com/prospero78/goTk/libtk"
)

type Window struct {
	*tk.Window
}

func NewWindow() *Window {
	mw := &Window{tk.RootWindow()}
	lbl := tk.NewLabel(mw, "Привет, гофер!")
	btn := tk.NewButton(mw, "Quit")
	btn.OnCommand(func() {
		tk.Quit()
	})
	tk.NewVPackLayout(mw).AddWidgets(lbl, tk.NewLayoutSpacer(mw, 0, true), btn)
	mw.ResizeN(300, 200)
	return mw
}

func main() {
	tk.MainLoop(func() {
		mw := NewWindow()
		mw.SetTitle("Пример goTk")
		mw.Center()
		mw.ShowNormal()
	})
}
```