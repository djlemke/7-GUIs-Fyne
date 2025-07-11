package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("7 GUIs Fyne")
	w.SetMaster()

	w.SetContent(container.New(layout.NewVBoxLayout(),
		widget.NewButton("Counter", func() {
			log.Println("Counter")
			counter := a.NewWindow("Counter")
			counter.SetContent(createCounter())
			counter.Resize(fyne.NewSize(100, 100))
			counter.Show()
		}),
		widget.NewButton("Temperature Converter", func() {
			log.Println("Temperature Converter")
			converter := a.NewWindow("Temperature Converter")
			converter.SetContent(createConverter())
			converter.Resize(fyne.NewSize(400, 50))
			converter.Show()
		}),
		widget.NewButton("Flight Booker", func() {
			log.Println("Flight Booker")
		}),
		widget.NewButton("Timer", func() {
			log.Println("Timer")
		}),
		widget.NewButton("CRUD", func() {
			log.Println("CRUD")
		}),
		widget.NewButton("Circle Drawer", func() {
			log.Println("Circle Drawer")
		}),
		widget.NewButton("Cells", func() {
			log.Println("Cells")
		}),
	))
	w.Resize(fyne.NewSize(250, 450))
	w.ShowAndRun()
}
