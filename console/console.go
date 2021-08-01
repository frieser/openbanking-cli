package console

import (
	"github.com/pterm/pterm"
	"strings"
)

func BigText(text string) {
	pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString(text)).Render()
}

func Table(data [][]string) {
	pterm.DefaultTable.WithHasHeader().WithData(data).Render()
	pterm.Printo(strings.Repeat(" ", pterm.GetTerminalWidth()))
}

func Task(message string) *pterm.SpinnerPrinter {
	spinner, _ := pterm.DefaultSpinner.Start(message)

	return spinner
}

func HeaderSuccess(message string) {
	pterm.Printo(strings.Repeat(" ", pterm.GetTerminalWidth()))
	pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgGreen)).
		WithMargin(10).WithFullWidth().Println(message)
}

func HeaderError(message string) {
	pterm.Printo(strings.Repeat(" ", pterm.GetTerminalWidth()))
	pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgRed)).
		WithMargin(10).WithFullWidth().Println(message)
}

func Section(message string) {
	pterm.DefaultSection.Println(message)
}
