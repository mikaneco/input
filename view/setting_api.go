package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func SettingApiView(a fyne.App) fyne.CanvasObject {
	leading := widget.NewLabel("Set endpoint and API key")
	endpoint := widget.NewEntry()
	key := widget.NewEntry()

	endpoint.SetText(a.Preferences().String("endpoint"))
	key.SetText(a.Preferences().String("apiKey"))

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Endpoint", Widget: endpoint},
			{Text: "Key", Widget: key},
		},
		OnCancel: func() {
			endpoint.SetText(a.Preferences().String("endpoint"))
			key.SetText(a.Preferences().String("apiKey"))
		},
		OnSubmit: func() {
			a.Preferences().SetString("endpoint", endpoint.Text)
			a.Preferences().SetString("apiKey", key.Text)
		},
		SubmitText: "Save",
	}

	return container.NewVBox(leading, form)
}
