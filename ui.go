package main

import (
	"fmt"
	"log"
	"strconv"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var tabSlice []string

func createUI(quesList []string, description []string, allAnswers map[int][]solution, acceptedAnswer map[int]solution) {
	makeRange := func(number int) []string {
		if number == 0 {
			return []string{"0"}
		}
		temp := make([]string, number)
		for i := range temp {
			temp[i] = strconv.Itoa(i + 1)
		}
		return temp
	}

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	width, height := ui.TerminalDimensions()

	header := widgets.NewParagraph()
	header.Text = "Press q to Exit, BackSpace to Return to Last Screen, Right Left to Shift Focus , Enter to open Question,Tab to Shift Right Answers , Ctrl-Z to move Left"
	header.SetRect(0, height-4, width, height)
	header.Border = false
	header.WrapText = true
	header.TextStyle.Bg = ui.ColorBlue

	quesBox := NewParagraph()

	quesBox.Title = "Question Description"
	quesBox.SetRect(width/2, 4, width, height-4)
	quesBox.start = 0
	quesBox.end = 44
	quesBox.WrapText = true
	quesBox.BorderStyle.Fg = ui.ColorYellow

	answerBox := NewParagraph()

	answerBox.Title = "Solution"
	answerBox.start = 0
	answerBox.end = 44
	answerBox.SetRect(35, 5, 80, 15)
	answerBox.BorderStyle.Fg = ui.ColorYellow

	quesListBox := widgets.NewList()
	quesListBox.Title = "Question List"
	quesListBox.Rows = quesList
	quesListBox.TextStyle = ui.NewStyle(ui.ColorYellow)
	quesListBox.WrapText = true
	quesListBox.SetRect(0, 4, width, height-4)

	tabpane := widgets.NewTabPane(tabSlice...)
	tabpane.SetRect(0, 1, 80, 4)
	tabpane.Border = true
	ui.Render(quesListBox, header)

	renderTab := func(quesNum int) {
		width, height = ui.TerminalDimensions()
		quesBox.SetRect(width/2, 4, width, height-4)
		quesListBox.SetRect(0, 4, width/2, height-4)

		quesBox.Text = allAnswers[quesNum][0].description

		ui.Render(quesBox, header, quesListBox)
	}
	scrollText := func(flag int, Box *Paragraph) {
		if flag == 1 {
			Box.start = Box.start + 2
			Box.end = Box.end + 2
		} else {
			Box.start = Box.start - 2
			Box.end = Box.end - 2
		}

		ui.Render(Box)
	}
	updateAnswer := func(ques int, quesNum int) {
		width, height := ui.TerminalDimensions()
		header.SetRect(0, height-4, width, height)
		quesListBox.SetRect(0, 0, 0, 0)
		answerBox.SetRect(width/2, 4, width, height-4)
		quesBox.SetRect(0, 4, width/2, height-4)

		if (acceptedAnswer[ques] == solution{}) {
			if len(allAnswers[ques]) == 1 {
				answerBox.Text = "Sorry no answer available"
			} else {
				quesNum = quesNum + 1
				answerBox.Text = allAnswers[ques][quesNum].description
				answerBox.Title = "Solution - " + allAnswers[ques][quesNum].upvotes + " upvotes"
			}
		} else {
			if quesNum == 0 {
				answerBox.Text = acceptedAnswer[ques].description
				answerBox.Title = "Solution (Accepted) - " + acceptedAnswer[ques].upvotes + " upvotes"
			} else {
				answerBox.Text = allAnswers[ques][quesNum].description
				answerBox.Title = "Solution - " + allAnswers[ques][quesNum].upvotes + " upvotes"
			}
		}
		ui.Render(quesBox, quesListBox, header, answerBox)

	}

	showAnswer := func(quesNum int) {
		width, height := ui.TerminalDimensions()

		quesListBox.SetRect(0, 0, 0, 0)
		answerBox.SetRect(width/2, 4, width, height-4)
		quesBox.SetRect(0, 4, width/2, height-4)
		quesBox.Text = allAnswers[quesNum][0].description
		quesBox.Title = fmt.Sprintf("Question %d", quesNum)
		header.SetRect(0, height-4, width, height)
		if (acceptedAnswer[quesNum] == solution{} && len(allAnswers[quesNum]) > 1) {
			tabSlice = makeRange(len(allAnswers[quesNum]) - 1)
			answerBox.Text = allAnswers[quesNum][1].description
			answerBox.Title = "Solution - " + allAnswers[quesNum][1].upvotes + " upvotes"
		} else if (len(allAnswers[quesNum]) == 1 && acceptedAnswer[quesNum] == solution{}) {

			answerBox.Text = "Sorry no answer available"
		} else {
			tabSlice = makeRange(len(allAnswers[quesNum]))
			answerBox.Text = acceptedAnswer[quesNum].description
			answerBox.Title = "Solution (Accepted)- " + acceptedAnswer[quesNum].upvotes + " upvotes"
		}

		tabpane.TabNames = tabSlice
		ui.Render(quesBox, header, quesListBox, answerBox, tabpane)
	}
	previousKey := ""
	focus := 0
	window := 0
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>", "<MouseWheelDown>":

			if focus == 0 && window == 0 {

				quesListBox.ScrollDown()
				renderTab(quesListBox.SelectedRow)
			}
			if focus == 1 && window == 1 {
				scrollText(1, quesBox)
			}
			if focus == 2 && window == 1 {
				scrollText(1, answerBox)
			}

		case "k", "<Up>", "<MouseWheelUp>":

			if focus == 0 && window == 0 {

				quesListBox.ScrollUp()
				renderTab(quesListBox.SelectedRow)
			}
			if focus == 1 && window == 1 {
				scrollText(2, quesBox)
			}
			if focus == 2 && window == 1 {
				scrollText(2, answerBox)
			}

		case "<Right>":
			if window == 1 {
				focus = 2
				answerBox.BorderStyle.Fg = ui.ColorWhite
				quesBox.BorderStyle.Fg = ui.ColorYellow
				ui.Render(answerBox, quesBox)

			}
		case "<Left>":
			if window == 1 {
				focus = 1
				answerBox.BorderStyle.Fg = ui.ColorYellow
				quesBox.BorderStyle.Fg = ui.ColorWhite
				ui.Render(answerBox, quesBox)

			}
		case "<Enter>":
			ui.Clear()
			window = 1
			focus = 2

			showAnswer(quesListBox.SelectedRow)
		case "<Backspace>":
			ui.Clear()
			tabpane.ActiveTabIndex = 0
			window = 0
			focus = 0
			renderTab(quesListBox.SelectedRow)

		case "<Tab>":
			if window == 1 {
				tabpane.FocusRight()

				ui.Render(tabpane)
				updateAnswer(quesListBox.SelectedRow, tabpane.ActiveTabIndex)
				focus = 2
			}
		case "<C-z>":
			tabpane.FocusLeft()

			ui.Render(tabpane, header)
			updateAnswer(quesListBox.SelectedRow, tabpane.ActiveTabIndex)
		case "<C-d>":
			quesListBox.ScrollHalfPageDown()
		case "<C-u>":
			quesListBox.ScrollHalfPageUp()
		case "<C-f>":
			quesListBox.ScrollPageDown()
		case "<C-b>":
			quesListBox.ScrollPageUp()
		case "g":
			if previousKey == "g" {
				quesListBox.ScrollTop()
			}
		case "<Home>":
			quesListBox.ScrollTop()
		case "G", "<End>":
			quesListBox.ScrollBottom()
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}

		ui.Render(quesListBox)
	}
}
