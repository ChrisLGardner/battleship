package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	_ "embed"

	"github.com/chrislgardner/battleship/package/battleship"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	appTitle string = `[yellow]Battleships`
	//go:embed loading.txt
	asciiLoading string

	app              = tview.NewApplication()
	pages            = tview.NewPages()
	userInputControl = newUserInputControl()
)

func main() {

	battleship.NewGame(8, 8)

	frontText, frontFlex := newLoadingPage(userInputControl)

	pages.AddPage("front", frontFlex, true, true)
	pages.AddPage("main", newContentPage(), true, false)

	if err := app.SetRoot(pages, true).SetFocus(frontText).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func newLoadingPage(userInput tview.Primitive) (textview tview.Primitive, flex tview.Primitive) {
	frontTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetChangedFunc(func() {
			app.Draw()
		}).
		SetTextAlign(tview.AlignCenter).
		SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEnter {
				pages.SwitchToPage("main")
				app.SetFocus(userInput)
			}
		})

	go func() {
		for _, word := range strings.Split(asciiLoading, "\n") {
			fmt.Fprintf(frontTextView, "%s\n", word)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	frontTextView.
		SetBorder(true).
		SetBorderAttributes(tcell.AttrBold).
		SetBorderColor(tcell.ColorPurple)

	frontFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(newFPHeader(appTitle, ""), 0, 1, false).
		AddItem(frontTextView, 0, 6, true)

	return frontTextView, frontFlex
}

// Set the header for the app
func newFPHeader(title string, text string) tview.Primitive {
	header := tview.NewTextView().SetText(text).
		SetTextAlign(1).
		SetDynamicColors(true)
	header.SetBorder(true).
		SetBorderAttributes(tcell.AttrBold).
		SetBorderColor(tcell.ColorPurple).
		SetTitle(title)
	return header
}

func newContentPage() tview.Primitive {
	middle := tview.NewFlex().
		AddItem(newMainText(), 0, 3, false)
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(newFPHeader(appTitle, ""), 0, 1, false).
		AddItem(middle, 0, 8, false).
		AddItem(userInputControl, 0, 2, true)
	return flex
}

func newMainText() tview.Primitive {
	landingBody := tview.NewTextView().
		SetWordWrap(true).
		SetText(GetBoard())
	landingBody.SetBorder(true).
		SetBorderAttributes(tcell.AttrBold).
		SetBorderColor(tcell.ColorPurple).
		SetTitle("[green]Game Board")

	return landingBody
}

func newUserInputControl() tview.Primitive {
	input := tview.NewInputField().
		SetLabel("Input player move (x,y):").
		SetFieldWidth(10) //.
		// SetAcceptanceFunc(func(textToCheck string, lastChar rune) bool {
		// 	if char, err := strconv.Atoi(string(lastChar)); err != nil || char >= len(battleship.Board) || lastChar != ',' {
		// 		return false
		// 	}
		// 	return true
		// })

	// input.SetDoneFunc(func(key tcell.Key) {
	// 	if key == tcell.KeyEnter {
	// 		PlayerFire(input.GetText())
	// 	}
	// })

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		k := event.Key()
		if k == tcell.KeyEnter && input.HasFocus() {
			PlayerFire(input.GetText())
		}

		return event
	})
	return input

	// flex := tview.NewFlex().
	// 	AddItem(input, 0, 1, true).
	// 	SetBorder(true).
	// 	SetBorderAttributes(tcell.AttrBold).
	// 	SetBorderColor(tcell.ColorPurple).
	// 	SetTitle("[green]Input")

	// return flex
}

func PlayerFire(input string) {

	splitInput := strings.Split(input, ",")
	targetX, _ := strconv.Atoi(splitInput[0])
	targetY, _ := strconv.Atoi(splitInput[1])

	battleship.PlayerMove(battleship.Position{X: targetX, Y: targetY})
	fmt.Println(battleship.Board)

}

func GetBoard() string {
	var sb strings.Builder

	for i, row := range battleship.Board {
		if i == 0 {
			sb.WriteString(" ")
			for k := 0; k < len(row); k++ {
				sb.WriteString(fmt.Sprintf("  %d  ", k))
			}
			sb.WriteString("\n")
		}
		sb.WriteString(fmt.Sprint(i))
		for _, col := range row {
			if col != 'M' || col != 'H' {
				sb.WriteString(" [ ] ")
			} else {
				sb.WriteString(fmt.Sprintf(" [%s] ", string(col)))
			}
		}
		sb.WriteString("\n")
	}

	sb.WriteString("\n")
	sb.WriteString("H marks a hit, M marks a miss")

	return sb.String()
}
