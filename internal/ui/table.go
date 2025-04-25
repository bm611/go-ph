package ui

import (
	"strconv"

	"github.com/bm611/go-ph/internal/llm"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var (
	purple      = lipgloss.Color("99")
	brightWhite = lipgloss.Color("255")

	headerStyle = lipgloss.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Center)
	cellStyle   = lipgloss.NewStyle().Padding(0, 1).Foreground(brightWhite)
)

func RenderTable(products []llm.ProductRespType) string {

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(purple)).
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
