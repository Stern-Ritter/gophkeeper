package client

import (
	"github.com/rivo/tview"
)

// selectView sets the given view as the root of the application and focuses on it.
// This function is useful for switching between different views in application.
func selectView(app *tview.Application, view tview.Primitive) {
	app.SetRoot(view, true).SetFocus(view)
}

// setTableHeader sets the header row of a tview.Table with the given column names.
func setTableHeader(table *tview.Table, columns []string) {
	for idx, column := range columns {
		table.SetCell(0, idx, tview.NewTableCell(column))
	}
}

// getColumnCounter returns a closure that increments and returns a column counter.
func getColumnCounter() func() int {
	currentColumn := -1
	return func() int {
		currentColumn++
		return currentColumn
	}
}

// newClickableCell creates a new tview.TableCell that is clickable.
// The cell executes the provided handler function when clicked.
func newClickableCell(text string, handler func()) *tview.TableCell {
	cell := tview.NewTableCell(text).
		SetAlign(tview.AlignCenter).
		SetClickedFunc(func() bool {
			handler()
			return false
		})

	return cell
}

// stopApp stops the tview application.
func stopApp(app *tview.Application) {
	app.Stop()
}

// clearForm clear form fields state
func clearForm(form *tview.Form) {
	for i := 0; i < form.GetFormItemCount(); i++ {
		item := form.GetFormItem(i)
		switch item := item.(type) {
		case *tview.InputField:
			item.SetText("")
		case *tview.TextView:
			item.SetText("")
		}
	}
}
