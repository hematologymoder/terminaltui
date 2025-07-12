package views

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/portfolio-tui/data"
	"github.com/portfolio-tui/styles"
)

func RenderExperience(experiences []data.Experience, education []data.Education, width, height int) string {
	var b strings.Builder

	b.WriteString(styles.SubsectionStyle.Render("🌸 experience 🌸"))
	b.WriteString("\n\n")

	for i, exp := range experiences {
		b.WriteString(styles.ProjectTitleStyle.Render(exp.Role))
		b.WriteString("\n")
		b.WriteString(styles.KeyStyle.Render(exp.Company))
		b.WriteString("\n")

		dateStr := exp.StartDate.Format("jan 2006")
		if exp.EndDate != nil {
			dateStr += " - " + exp.EndDate.Format("jan 2006")
		} else {
			dateStr += " - present"
		}
		b.WriteString(styles.DescStyle.Render(dateStr + " | " + exp.Location))
		b.WriteString("\n\n")

		if len(exp.Achievements) > 0 {
			achievementsToShow := exp.Achievements
			if len(achievementsToShow) > 3 {
				achievementsToShow = achievementsToShow[:3]
			}
			
			for _, achievement := range achievementsToShow {
				b.WriteString(fmt.Sprintf("  • %s\n", achievement))
			}
			
			if len(exp.Achievements) > 3 {
				b.WriteString(styles.DescStyle.Render(fmt.Sprintf("  + %d more achievements...", len(exp.Achievements)-3)))
				b.WriteString("\n")
			}
		}

		if i < len(experiences)-1 {
			separator := strings.Repeat("─", 50)
			b.WriteString("\n" + styles.DescStyle.Render(separator) + "\n\n")
		}
	}

	if len(education) > 0 {
		b.WriteString("\n\n")
		b.WriteString(styles.SubsectionStyle.Render("📚 education"))
		b.WriteString("\n\n")
		
		for _, edu := range education {
			b.WriteString(styles.ProjectTitleStyle.Render(edu.Institution))
			b.WriteString("\n")
			b.WriteString(styles.DescStyle.Render(fmt.Sprintf("%s | %d-%d | %s", edu.Field, edu.StartYear, edu.EndYear, edu.Location)))
			b.WriteString("\n")
		}
	}

	b.WriteString("\n\n")

	helpText := styles.HelpStyle.Render(
		styles.KeyStyle.Render("←→") + " sections | " +
		styles.KeyStyle.Render("esc") + " back",
	)
	
	b.WriteString(lipgloss.PlaceHorizontal(width, lipgloss.Center, helpText))

	return b.String()
}