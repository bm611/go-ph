package main

import (
	"fmt"

	"github.com/bm611/go-ph/cmd"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	cmd.Execute()

	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingTop(1).
		PaddingLeft(4).
		Width(22).
		PaddingBottom(1)

	fmt.Println(style.Render("Hello, kitty"))

}
