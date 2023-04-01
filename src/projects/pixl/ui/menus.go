package ui

import (
	"errors"
	"image"
	"image/png"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"zerotomastery.io/pixl/util"
)

func saveFileDialog(app *AppInit) {
	dialog.ShowFileSave(func(file fyne.URIWriteCloser, e error) {
		if file == nil {
			return
		}

		err := png.Encode(file, app.PixlCanvas.PixelData)
		if err != nil {
			dialog.ShowError(err, app.PixlWindow)
			return
		}

		app.State.SetFilePath(file.URI().Path())
	}, app.PixlWindow)
}

func BuildSaveAsMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save As...", func() {
		saveFileDialog(app)
	})
}

func BuildSaveMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save", func() {
		if app.State.FilePath == "" {
			// open save as dialog
			saveFileDialog(app)
			return
		}

		tryClose := func(fh *os.File) {
			err := fh.Close()
			if err != nil {
				// failed to close, show error
				dialog.ShowError(err, app.PixlWindow)
			}
		}

		file, err := os.Create(app.State.FilePath)
		defer tryClose(file)

		if err != nil {
			dialog.ShowError(err, app.PixlWindow)
			return
		}

		// save image
		err = png.Encode(file, app.PixlCanvas.PixelData)
		if err != nil {
			dialog.ShowError(err, app.PixlWindow)
			return
		}
	})
}

func BuildNewMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("New", func() {
		sizeValidator := func(s string) error {
			value, err := strconv.Atoi(s)
			if err != nil {
				return errors.New("must be a positive integer")
			}
			if value <= 0 {
				return errors.New("must be > 0")
			}
			return nil
		}

		widthEntry := widget.NewEntry()
		widthEntry.Validator = sizeValidator

		heightEntry := widget.NewEntry()
		heightEntry.Validator = sizeValidator

		widthFormEntry := widget.NewFormItem("Width", widthEntry)
		heightFormEntry := widget.NewFormItem("Height", heightEntry)

		formItems := []*widget.FormItem{widthFormEntry, heightFormEntry}

		dialog.ShowForm("New Image", "Create", "Cancel", formItems, func(ok bool) {
			if ok {
				pixelWidth := 0
				if widthEntry.Validate() != nil {
					dialog.ShowError(errors.New("invalid width"), app.PixlWindow)
					return
				}
				pixelWidth, _ = strconv.Atoi(widthEntry.Text)

				pixelHeight := 0
				if heightEntry.Validate() != nil {
					dialog.ShowError(errors.New("invalid height"), app.PixlWindow)
					return
				}
				pixelHeight, _ = strconv.Atoi(heightEntry.Text)

				app.PixlCanvas.NewDrawing(pixelWidth, pixelHeight)
			}
		}, app.PixlWindow)
	})
}

func BuildOpenMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Open", func() {
		dialog.ShowFileOpen(func(file fyne.URIReadCloser, e error) {
			if file == nil {
				return
			}

			defer file.Close()

			// load image
			image, _, err := image.Decode(file)
			if err != nil {
				dialog.ShowError(err, app.PixlWindow)
				return
			}

			app.PixlCanvas.LoadImage(image)
			app.State.SetFilePath(file.URI().Path())

			imageColors := util.GetImageColors(image)
			i := 0
			for c := range imageColors {
				if i == len(app.Swatches) {
					// no more swatches available
					break
				}
				app.Swatches[i].SetColor(c)
				i++
			}
		}, app.PixlWindow)
	})
}

func BuildMenus(app *AppInit) *fyne.Menu {
	return fyne.NewMenu(
		"File",
		BuildNewMenu(app),
		BuildOpenMenu(app),
		BuildSaveMenu(app),
		BuildSaveAsMenu(app),
	)
}

func SetupMenus(app *AppInit) {
	menus := BuildMenus(app)
	mainMenu := fyne.NewMainMenu(menus)
	app.PixlWindow.SetMainMenu(mainMenu)
}
