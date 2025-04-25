package ui

import (
	"strconv"

	"github.com/bm611/go-ph/internal/llm"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var (
	// Soft pastel colors
	pastelPink   = lipgloss.Color("#FFB6C1") // Light pink
	pastelPurple = lipgloss.Color("#CBC3E3") // Light purple
	pastelBlue   = lipgloss.Color("#CCCCFF") // Light blue
	brightWhite  = lipgloss.Color("255")

	headerStyle = lipgloss.NewStyle().Foreground(pastelPurple).Bold(true).Align(lipgloss.Center)
	cellStyle   = lipgloss.NewStyle().Padding(0, 1).Foreground(brightWhite)
)

func RenderTable(products []llm.ProductRespType) string {

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderRow(true).
		BorderColumn(true).
		BorderStyle(lipgloss.NewStyle().Foreground(pastelBlue)).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return headerStyle
			default:
				return cellStyle
			}
		}).
		Headers("RANK", "NAME", "DESCRIPTION")

	for _, p := range products {
		t.Row(
			strconv.Itoa(p.Rank), // Convert int Rank to string
			p.Name,
			p.Description,
		)
	}

	return t.String()
}
