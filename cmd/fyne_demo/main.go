// Package main provides various examples of Fyne API capabilities
package main

import "errors"
import "fmt"

import "github.com/fyne-io/fyne"
import "github.com/fyne-io/fyne/app"
import "github.com/fyne-io/fyne/layout"
import "github.com/fyne-io/fyne/theme"
import "github.com/fyne-io/fyne/dialog"
import "github.com/fyne-io/fyne/widget"

func formApp(app fyne.App) {
	w := app.NewWindow("Form")

	largeText := widget.NewEntry()
	//	largeText.Text = "\n\n\n"

	form := &widget.Form{
		OnCancel: func() {
			w.Close()
		},
		OnSubmit: func() {
			fmt.Println("Form submitted")
		},
	}
	form.Append("Name", widget.NewEntry())
	form.Append("Email", widget.NewEntry())
	form.Append("Message", largeText)
	w.SetContent(form)
	w.Show()
}

func confirmCallback(response bool) {
	fmt.Println("Responded with", response)
}

func main() {
	app := app.New()

	w := app.NewWindow("Fyne Demo")
	entry := widget.NewEntry()
	entry.Text = "Entry"

	w.SetContent(widget.NewVBox(
		widget.NewToolbar(widget.NewToolbarAction(theme.MailComposeIcon(), func() { fmt.Println("New") }),
			widget.NewToolbarSeparator(),
			widget.NewToolbarSpacer(),
			widget.NewToolbarAction(theme.CutIcon(), func() { fmt.Println("Cut") }),
			widget.NewToolbarAction(theme.CopyIcon(), func() { fmt.Println("Copy") }),
			widget.NewToolbarAction(theme.PasteIcon(), func() { fmt.Println("Paste") }),
		),

		widget.NewButton("Apps", func() {
			dialog.ShowInformation("Information", "Example applications have moved to https://github.com/fyne-io/examples", w)
		}),

		widget.NewGroup("Demos",
			widget.NewButton("Canvas", func() { Canvas(app) }),
			widget.NewButton("Icons", func() { Icons(app) }),
			widget.NewButton("Layout", func() { Layout(app) }),
			widget.NewButton("Widgets", func() { Widget(app) }),
			widget.NewButton("Form", func() { formApp(app) }),
		),

		widget.NewGroup("Dialogs",
			widget.NewButton("Info", func() {
				dialog.ShowInformation("Information", "You should know this thing...", w)
			}),
			widget.NewButton("Error", func() {
				err := errors.New("A dummy error message")
				dialog.ShowError(err, w)
			}),
			widget.NewButton("Confirm", func() {
				cnf := dialog.NewConfirm("Confirmation", "Are you enjoying this demo?", confirmCallback, w)
				cnf.SetDismissText("Nah")
				cnf.SetConfirmText("Oh Yes!")
				cnf.Show()
			}),
			widget.NewButton("Custom", func() {
				dialog.ShowCustom("MyDialog", "Nice", widget.NewCheck("Inside a dialog", func(bool) {}), w)
			}),
		),

		layout.NewSpacer(),

		fyne.NewContainerWithLayout(layout.NewGridLayout(2),
			widget.NewButton("Dark", func() {
				fyne.GetSettings().SetTheme("dark")
			}),
			widget.NewButton("Light", func() {
				fyne.GetSettings().SetTheme("light")
			}),
		),
		widget.NewButtonWithIcon("Quit", theme.CancelIcon(), func() {
			app.Quit()
		}),
	))
	w.ShowAndRun()
}
