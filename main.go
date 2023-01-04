package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"fyne.io/fyne/v2/data/binding"
)

func loadJsonData() []string {
	fmt.Println("Loading data from JSON file")

	input, _ := ioutil.ReadFile("data.json")
	var data []string
	json.Unmarshal(input, &data)

	return data
}

func saveJsonData(data binding.StringList) {
	fmt.Println("Saving data to JSON file")
	d, _ := data.Get()
	jsonData, _ := json.Marshal(d)
	ioutil.WriteFile("data.json", jsonData, 0644)

}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("List Data")

	loadedData := loadJsonData()

	data := binding.NewStringList()
	data.Set(loadedData)

	defer saveJsonData(data)

	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	list.OnSelected = func(id widget.ListItemID) {
		list.Unselect(id)
		d, _ := data.GetValue(id)
		w := myApp.NewWindow("Edit Data")

		itemName := widget.NewEntry()
		itemName.Text = d

		updateData := widget.NewButton("Update", func() {
			data.SetValue(id, itemName.Text)
			w.Close()
		})

		cancel := widget.NewButton("Cancel", func() {
			w.Close()
		})

		deleteData := widget.NewButton("Delete", func() {
			var newData []string
			dt, _ := data.Get()

			for index, item := range dt {
				if index != id {
					newData = append(newData, item)
				}
			}

			data.Set(newData)

			w.Close()
		})

		w.SetContent(container.New(layout.NewVBoxLayout(), itemName, updateData, deleteData, cancel))
		w.Resize(fyne.NewSize(400, 200))
		w.CenterOnScreen()
		w.Show()

	}

	add := widget.NewButton("Add", func() {
		w := myApp.NewWindow("Add Data")

		itemName := widget.NewEntry()

		addData := widget.NewButton("Add", func() {
			data.Append(itemName.Text)
			w.Close()
		})

		cancel := widget.NewButton("Cancel", func() {
			w.Close()
		})

		w.SetContent(container.New(layout.NewVBoxLayout(), itemName, addData, cancel))
		w.Resize(fyne.NewSize(400, 200))
		w.CenterOnScreen()
		w.Show()

	})

	exit := widget.NewButton("Quit", func() {

		myWindow.Close()
	})

	myWindow.SetContent(container.NewBorder(nil, container.New(layout.NewVBoxLayout(), add, exit), nil, nil, list))
	myWindow.Resize(fyne.NewSize(400, 600))
	myWindow.SetMaster()
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()

}