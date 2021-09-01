package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("change to Welcome")
		}),
	))

	w.Show()
	a.Run()

	//w.ShowAndRun() // 是w.Show() && a.Run()的快捷方式

}

/**
打包的任务： go get fyne.io/fyne/v2/cmd/fyne

			go build
			fyne package -icon mylogo.png
*/
