package tui

import "gitlab.com/tslocum/cview"

func initFlex(in *cview.InputField, tv *cview.TextView) *cview.Flex {
	return cview.NewFlex().
		AddItem(tv, 0, 2, true).
		AddItem(in, 1, 1, true).
		SetDirection(cview.FlexRow)
}

// Init returns pointer to cview App interface
func Init() *cview.Application {

	view := cview.NewApplication().
		EnableMouse(true)

	in := cview.NewInputField().
		SetLabel(">").
		SetPlaceholder("server command").
		SetFieldWidth(0)

	tv := cview.NewTextView().
		SetScrollable(true).
		SetChangedFunc(func() {
			view.Draw()
		})

	tv.SetTitle(" Torch Console ").
		SetBorder(true)

	flx := initFlex(in, tv)

	view.SetRoot(flx, true)

	return view

}
