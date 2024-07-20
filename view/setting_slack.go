package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func SettingSlackView(a fyne.App) fyne.CanvasObject {
	leading := widget.NewLabel("Set Slack channel and token")
	channel := widget.NewEntry()
	token := widget.NewEntry()

	channel.SetText(a.Preferences().String("channel"))
	token.SetText(a.Preferences().String("token"))

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Channel", Widget: channel},
			{Text: "Token", Widget: token},
		},
		OnCancel: func() {
			channel.SetText(a.Preferences().String("channel"))
			token.SetText(a.Preferences().String("token"))
		},
		OnSubmit: func() {
			a.Preferences().SetString("channel", channel.Text)
			a.Preferences().SetString("token", token.Text)
		},
		SubmitText: "Save",
	}

	return container.NewVBox(leading, form)
}
