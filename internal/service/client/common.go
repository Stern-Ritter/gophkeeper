package client

import (
	"github.com/rivo/tview"
)

// SelectView sets the given view as the root of the application and focuses on it.
// This function is useful for switching between different views in application.
func (c *ClientImpl) SelectView(view tview.Primitive) {
	selectView(c.app, view)
}

// UpdateDraw forces the application to redraw its content.
// This method should be called after updating views or their content
// to ensure the changes are rendered on the screen.
func (c *ClientImpl) UpdateDraw() {
	c.app.Draw()
}

// QueueUpdateDraw queues a function to update the application's UI.
// The provided render function will be executed during the next rendering cycle.
func (c *ClientImpl) QueueUpdateDraw(render func()) {
	c.app.QueueUpdateDraw(render)
}

// selectView sets the given view as the root of the application and focuses on it.
// This function is useful for switching between different views in application.
func selectView(app Application, view tview.Primitive) {
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
func stopApp(app Application) {
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
