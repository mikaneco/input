package main

import (
	"github.com/mikaneco/inputbox/view"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.NewWithID("jp.input.app")
	w := a.NewWindow("Inputs")
	w.Resize(fyne.NewSize(400, 100))

	tabs := container.NewAppTabs(
		container.NewTabItem("Post", view.PostView(a)),
		container.NewTabItem("Slack Setting", view.SettingSlackView(a)),
		container.NewTabItem("API Setting", view.SettingApiView(a)),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabs)

	w.ShowAndRun()
}
