package main

import (
	"log"
	"net/http"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("jp.input.app")
	w := a.NewWindow("Inputs")
	w.Resize(fyne.NewSize(400, 100))

	leadingInput := widget.NewLabel("Set input and press enter")
	entry := widget.NewMultiLineEntry()
	progress := widget.NewProgressBar()
	infinite := widget.NewProgressBarInfinite()
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Input", Widget: entry},
		},
		OnCancel: func() {
			entry.SetText("")
		},
		OnSubmit: func() {
			leadingInput.SetText("Sending request...")
			progress.SetValue(0)
			infinite.Start()

			go func() {
				err := sendRequest(
					a.Preferences().String("endpoint"),
					a.Preferences().String("apiKey"),
					entry.Text,
				)

				if err != nil {
					leadingInput.SetText("Error: " + err.Error())
				} else {
					leadingInput.SetText("Request sent")
					entry.SetText("")
				}
			}()

			progress.SetValue(1)
			infinite.Stop()
		},
	}

	inputGroup := container.NewVBox(leadingInput, form)

	configLeading := widget.NewLabel("Set endpoint and API key")
	endpoint := widget.NewEntry()
	apiKey := widget.NewEntry()

	endpoint.SetText(a.Preferences().String("endpoint"))
	apiKey.SetText(a.Preferences().String("apiKey"))

	settingForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Endpoint", Widget: endpoint},
			{Text: "Key", Widget: apiKey},
		},
		OnCancel: func() {
			endpoint.SetText(a.Preferences().String("endpoint"))
			apiKey.SetText(a.Preferences().String("apiKey"))
		},
		OnSubmit: func() {
			a.Preferences().SetString("endpoint", endpoint.Text)
			a.Preferences().SetString("apiKey", apiKey.Text)
		},
		SubmitText: "Save",
	}

	configGroup := container.NewVBox(configLeading, settingForm)

	tabs := container.NewAppTabs(
		container.NewTabItem("Input", inputGroup),
		container.NewTabItem("Config", configGroup),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabs)

	w.ShowAndRun()
}

func sendRequest(endpoint, apiKey, input string) error {
	log.Println("Sending request to", endpoint, "with", input)
	req, err := http.NewRequest(
		"POST",
		endpoint,
		strings.NewReader(`{"input":"`+input+`"}`),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
