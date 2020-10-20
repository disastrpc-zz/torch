package tui

import (
	"gitlab.com/tslocum/cview"
)

// Tui represents the terminal view
type Tui struct {
	In  *cview.InputField
	Tv  *cview.TextView
	Flx *cview.Flex
	App *cview.Application
}

// Init returns pointer to Tui struct which contains ui components, also set many UI components to correct status
func Init() (view *Tui) {

	app := cview.NewApplication().
		EnableMouse(true)

	in := initInput()
	tv := initText(app)
	flx := initFlex(in, tv)

	tv.SetBackgroundColor(696969)

	in.SetLabelColor(377369)

	view = &Tui{
		in,
		tv,
		flx,
		app,
	}

	return view
}

// InitInput returns InputField primitive
func initInput() *cview.InputField {
	return cview.NewInputField().
		SetPlaceholder(" server command").
		SetFieldWidth(0).
		SetFieldBackgroundColor(323232)
}

// InitText returns TextView primitive
func initText(app *cview.Application) *cview.TextView {
	return cview.NewTextView().
		SetScrollable(true).
		SetChangedFunc(func() {
			app.Draw()
		})
}

// InitFlex returns Flex with in and tv added
func initFlex(in *cview.InputField, tv *cview.TextView) *cview.Flex {
	return cview.NewFlex().
		AddItem(tv, 0, 2, true).
		AddItem(in, 1, 1, true).
		SetDirection(cview.FlexRow)
}
