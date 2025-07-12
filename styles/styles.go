package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	Primary     = lipgloss.Color("#FFB6C1")
	Secondary   = lipgloss.Color("#FF69B4")
	Accent      = lipgloss.Color("#FFC0CB")
	Muted       = lipgloss.Color("#DDA0DD")
	Error       = lipgloss.Color("#DC143C")
	Success     = lipgloss.Color("#98FB98")
	Warning     = lipgloss.Color("#FFB6C1")
	Background  = lipgloss.Color("#FFF0F5")
	Foreground  = lipgloss.Color("#4B0082")

	BaseStyle = lipgloss.NewStyle().
		Background(Background).
		Foreground(Foreground)

	TitleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(Primary).
		MarginBottom(1).
		Padding(0, 1)

	HeaderStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(Secondary).
		BorderStyle(lipgloss.DoubleBorder()).
		BorderForeground(Primary).
		Padding(0, 2).
		MarginBottom(1)

	NavItemStyle = lipgloss.NewStyle().
		Padding(0, 2).
		MarginRight(2)

	NavItemSelectedStyle = NavItemStyle.Copy().
		Foreground(Background).
		Background(Primary).
		Bold(true)

	SectionStyle = lipgloss.NewStyle().
		Padding(1, 2).
		MarginBottom(1)

	SubsectionStyle = lipgloss.NewStyle().
		Foreground(Secondary).
		Bold(true).
		MarginBottom(1)

	ContentStyle = lipgloss.NewStyle().
		Padding(0, 2)

	ParagraphStyle = lipgloss.NewStyle().
		MarginBottom(1).
		Padding(0, 2)

	ListStyle = lipgloss.NewStyle().
		MarginLeft(2)

	ListItemStyle = lipgloss.NewStyle().
		Foreground(Foreground).
		MarginBottom(0)

	ListItemSelectedStyle = lipgloss.NewStyle().
		Foreground(Primary).
		Bold(true).
		MarginBottom(0)

	ProjectTitleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(Primary).
		MarginBottom(0)

	ProjectDescStyle = lipgloss.NewStyle().
		Foreground(Muted).
		MarginLeft(2)

	TechTagStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(Secondary).
		Padding(0, 1).
		MarginRight(1)

	SkillNameStyle = lipgloss.NewStyle().
		Width(20).
		Foreground(Foreground)

	SkillBarStyle = lipgloss.NewStyle().
		Foreground(Primary)

	SkillYearsStyle = lipgloss.NewStyle().
		Foreground(Muted).
		MarginLeft(2)

	ProgressBarStyle = lipgloss.NewStyle().
		Foreground(Primary).
		Background(lipgloss.Color("#FFF0F5"))

	ProgressEmptyStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFE4E1"))

	BorderStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(Muted).
		Padding(1, 2)

	HelpStyle = lipgloss.NewStyle().
		Foreground(Muted).
		Padding(1, 2)

	KeyStyle = lipgloss.NewStyle().
		Foreground(Secondary).
		Bold(true)

	DescStyle = lipgloss.NewStyle().
		Foreground(Muted)

	StatusStyle = lipgloss.NewStyle().
		Foreground(Success).
		Bold(true)

	ASCIIStyle = lipgloss.NewStyle().
		Foreground(Primary).
		Bold(true)
)

func RenderProgressBar(current, total, width int) string {
	if width <= 0 {
		width = 20
	}
	
	filled := (current * width) / total
	if filled > width {
		filled = width
	}
	
	bar := ""
	for i := 0; i < filled; i++ {
		bar += "█"
	}
	for i := filled; i < width; i++ {
		bar += "░"
	}
	
	return ProgressBarStyle.Render(bar)
}

func RenderSkillLevel(level int) string {
	return RenderProgressBar(level, 5, 10)
}