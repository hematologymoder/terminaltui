package views

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/portfolio-tui/data"
	"github.com/portfolio-tui/styles"
)

func RenderSkills(categories []data.SkillCategory, selectedIndex int, width, height int) string {
	var b strings.Builder

	b.WriteString(styles.SubsectionStyle.Render("💖 skills 💖"))
	b.WriteString("\n\n")

	for i, category := range categories {
		categoryName := category.Name
		if i == selectedIndex {
			b.WriteString(styles.ListItemSelectedStyle.Render(fmt.Sprintf("▶ %s", categoryName)))
		} else {
			b.WriteString(styles.SubsectionStyle.Render(fmt.Sprintf("  %s", categoryName)))
		}
		b.WriteString("\n")

		skillsToShow := category.Skills
		if len(skillsToShow) > 4 {
			skillsToShow = skillsToShow[:4]
		}
		
		for _, skill := range skillsToShow {
			skillName := skill.Name
			if len(skillName) > 20 {
				skillName = skillName[:17] + "..."
			}
			
			skillLine := fmt.Sprintf("    %-20s %s %dy", 
				skillName, 
				renderCompactSkillLevel(skill.Level),
				skill.Years)
			b.WriteString(skillLine)
			b.WriteString("\n")
		}
		
		if len(category.Skills) > 4 {
			b.WriteString(styles.DescStyle.Render(fmt.Sprintf("    + %d more skills...", len(category.Skills)-4)))
			b.WriteString("\n")
		}
		
		if i < len(categories)-1 {
			separator := strings.Repeat("─", 40)
			b.WriteString(styles.DescStyle.Render("  " + separator))
			b.WriteString("\n")
		}
		b.WriteString("\n")
	}

	legend := styles.DescStyle.Render("skill levels: " +
		renderCompactSkillLevel(1) + " basic  " +
		renderCompactSkillLevel(3) + " good  " +
		renderCompactSkillLevel(5) + " expert")
	b.WriteString(lipgloss.PlaceHorizontal(width, lipgloss.Center, legend))
	b.WriteString("\n\n")

	helpText := styles.HelpStyle.Render(
		styles.KeyStyle.Render("↑↓") + " navigate | " +
		styles.KeyStyle.Render("esc") + " back",
	)
	
	b.WriteString(lipgloss.PlaceHorizontal(width, lipgloss.Center, helpText))

	return b.String()
}

func renderCompactSkillLevel(level int) string {
	filled := ""
	empty := ""
	
	for i := 0; i < level; i++ {
		filled += "●"
	}
	for i := level; i < 5; i++ {
		empty += "○"
	}
	
	return styles.ProgressBarStyle.Render(filled) + styles.ProgressEmptyStyle.Render(empty)
}