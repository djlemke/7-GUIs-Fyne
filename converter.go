package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func createConverter() *fyne.Container {
	fahr := binding.NewFloat()
	celsius := binding.NewFloat()
	fahrEntry := widget.NewEntryWithData(binding.FloatToStringWithFormat(fahr, "%.2f"))
	celsiusEntry := widget.NewEntryWithData(binding.FloatToStringWithFormat(celsius, "%.2f"))

	fahr.Set(32)
	celsius.Set(0)

	fahr.AddListener(binding.NewDataListener(func() {
		f, err := fahr.Get()
		if err != nil {
			log.Fatal("Error getting fahr value!")
		}
		c := (f - 32) * (5.0 / 9.0)
		celsius.Set(c)
	}))

	celsius.AddListener(binding.NewDataListener(func() {
		c, err := celsius.Get()
		if err != nil {
			log.Fatal("Error getting celsius value!")
		}
		f := c*(9.0/5.0) + 32
		fahr.Set(f)
	}))

	/*return container.NewHBox(
		celsiusEntry,
		widget.NewLabel("Celsius ="),
		fahrEntry,
		widget.NewLabel("Fahrenheit"),
	)*/
	return container.NewGridWithColumns(
		2,
		container.NewBorder(nil, nil, nil, widget.NewLabel("Celsius ="), celsiusEntry),
		container.NewBorder(nil, nil, nil, widget.NewLabel("Fahrenheit"), fahrEntry),
	)
}
