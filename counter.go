package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func createCounter() *fyne.Container {
	counter := binding.NewInt()
	str := binding.IntToString(counter)
	label := widget.NewLabelWithData(str)

	counter.Set(0)

	return container.NewVBox(
		label,
		widget.NewButton("Count", func() {
			i, err := counter.Get()
			if err != nil {
				log.Fatal("Error getting counter value!")
			}
			counter.Set(i + 1)
		}))
}
