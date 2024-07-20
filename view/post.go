package view

import (
	"log"

	"github.com/mikaneco/inputbox/api"
	"github.com/mikaneco/inputbox/slack"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func PostView(a fyne.App) fyne.CanvasObject {
	leading := widget.NewLabel("Set note and press enter")
	entry := widget.NewMultiLineEntry()
	slackCheck := widget.NewCheck("Send to Slack", func(b bool) {
		a.Preferences().SetString("sendSlack", "true")
	})

	apiCheck := widget.NewCheck("Send to API", func(b bool) {
		a.Preferences().SetString("sendApi", "true")
	})
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Input", Widget: entry},
		},
		OnCancel: func() {
			entry.SetText("")
		},
		OnSubmit: func() {
			leading.SetText("Sending request...")

			go func() {
				if a.Preferences().String("sendApi") == "true" {
					api := api.NewApiServer(
						a.Preferences().String("endpoint"),
						a.Preferences().String("apiKey"),
					)
					log.Println("api", entry.Text)
					err := api.SendMessage(entry.Text)

					if err != nil {
						leading.SetText("Error: " + err.Error())
					} else {
						leading.SetText("Request sent")
						entry.SetText("")
					}
				}

				if a.Preferences().String("sendSlack") == "true" {
					s := slack.NewSlack(
						a.Preferences().String("token"),
						a.Preferences().String("channel"),
					)
					log.Println("slack", entry.Text)
					err := s.SendMessage(entry.Text)

					if err != nil {
						leading.SetText("Error: " + err.Error())
					} else {
						leading.SetText("Request sent")
						entry.SetText("")
					}
				}
			}()
		},
	}

	return container.NewVBox(leading, slackCheck, apiCheck, form)
}
