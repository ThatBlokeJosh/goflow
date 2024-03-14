package main

import (
	"bytes"
	"fmt"
	"goflow/utils"
	"html"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var textTab *tview.TextView = view()
var app *tview.Application = tview.NewApplication()
var page int = 1
var search string = ""

func input() *tview.InputField {
  input := tview.NewInputField().SetLabel(" >>> ")

  input.SetFieldBackgroundColor(tcell.ColorNames["black"])
  input.SetBorder(true)
  input.SetFieldWidth(50)
  
  input = input.SetDoneFunc(func(key tcell.Key) {
    if key == tcell.KeyEnter {
	if input.GetText() == "->" {
		page++
	} else if input.GetText() == "<-" {
		if page > 1 {
			page--
		}
	} else {
		search = input.GetText()		
		page = 1
		textTab.Clear()
		textTab.Write([]byte("[blue]Loading..."))
	}
	items := utils.Search(search, page)
	textTab.Clear()
	DrawItems(items)
    } 
  })

  return input
}


func formatChatMessage(item utils.Item) []byte {
	var buf bytes.Buffer
		
	buf.WriteString(fmt.Sprintf(" [blue]%s[white]\n \t[ ", html.UnescapeString(item.Title)))
	for _, tag := range item.Tags {
		buf.WriteString(fmt.Sprintf("[yellow]%s[white] ", tag))
	} 
	buf.WriteString(fmt.Sprintf("]\n \t[purple]Answered[white]: [red]%s[white]", strconv.FormatBool(item.Answered)))
	buf.WriteString(fmt.Sprintf("\n \t[orange]>[white] [green]%s[white]\n", item.Link))
	buf.WriteRune('\n')

	return buf.Bytes()
}
func DrawItems(items utils.Items) {
	textTab.Clear()
	if len(items.Items) == 0 {
		textTab.Write([]byte("[blue]No questions for this topic"))
		return
	}
	for _, item := range items.Items {
		textTab.Write(formatChatMessage(item))
	}
}

func view() *tview.TextView {
  view := tview.NewTextView().SetDynamicColors(true)

  view.SetBorder(true)
  view.SetTitle("GoFlow")
  view.SetTitleAlign(tview.AlignLeft)

  return view
}



func layout() {
	app := tview.NewApplication()
	grid := tview.NewGrid()
	grid.SetRows(0, 5).AddItem(textTab, 0, 0, 1, 4, 0, 0, false).AddItem(input(), 1, 0, 1, 4, 0, 0, true)
	if err := app.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}
}

func main() {
  layout()
}
