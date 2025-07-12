package views

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/portfolio-tui/data"
	"github.com/portfolio-tui/styles"
)

func RenderProjectsList(projects []data.Project, selectedIndex int, width, height int) string {
	var b strings.Builder

	b.WriteString(styles.SubsectionStyle.Render("✨ projects ✨"))
	b.WriteString("\n\n")

	for i, project := range projects {
		var projectView strings.Builder
		
		title := project.Name
		if project.Status == "active" {
			title += " ✨"
		} else {
			title += " ✅"
		}
		
		if i == selectedIndex {
			projectView.WriteString(styles.ListItemSelectedStyle.Render(fmt.Sprintf("▶ %s", title)))
		} else {
			projectView.WriteString(styles.ProjectTitleStyle.Render(fmt.Sprintf("  %s", title)))
		}
		projectView.WriteString("\n")
		
		projectView.WriteString(styles.ProjectDescStyle.Render("    " + project.Description))
		projectView.WriteString("\n")
		
		var techTags []string
		for _, tech := range project.Technologies {
			techTags = append(techTags, styles.TechTagStyle.Render(tech))
		}
		projectView.WriteString("    " + strings.Join(techTags, " "))
		projectView.WriteString("\n")
		
		b.WriteString(projectView.String())
		if i < len(projects)-1 {
			b.WriteString("\n")
		}
	}

	helpText := styles.HelpStyle.Render(
		styles.KeyStyle.Render("↑↓") + " navigate | " +
		styles.KeyStyle.Render("enter") + " details | " +
		styles.KeyStyle.Render("esc") + " back",
	)
	
	b.WriteString("\n\n")
	b.WriteString(lipgloss.PlaceHorizontal(width, lipgloss.Center, helpText))

	return b.String()
}

func RenderProjectDetail(project data.Project, width, height int) string {
	var b strings.Builder

	b.WriteString(styles.TitleStyle.Render("💕 " + project.Name + " 💕"))
	b.WriteString("\n\n")

	if project.Status == "active" {
		b.WriteString(styles.SubsectionStyle.Render("status:"))
		b.WriteString("\n")
		b.WriteString(styles.StatusStyle.Render("✨ active"))
		b.WriteString("\n\n")
	} else {
		b.WriteString(styles.SubsectionStyle.Render("status:"))
		b.WriteString("\n")
		b.WriteString(styles.DescStyle.Render("✅ completed"))
		b.WriteString("\n\n")
	}

	b.WriteString(styles.SubsectionStyle.Render("about:"))
	b.WriteString("\n")
	b.WriteString(project.LongDesc)
	b.WriteString("\n\n")

	b.WriteString(styles.SubsectionStyle.Render("tech stack:"))
	b.WriteString("\n")
	var techTags []string
	for _, tech := range project.Technologies {
		techTags = append(techTags, styles.TechTagStyle.Render(tech))
	}
	b.WriteString(strings.Join(techTags, " "))
	b.WriteString("\n\n")

	b.WriteString(styles.SubsectionStyle.Render("features:"))
	b.WriteString("\n")
	for _, feature := range project.Features {
		b.WriteString(fmt.Sprintf("  • %s\n", feature))
	}
	b.WriteString("\n")

	b.WriteString(styles.SubsectionStyle.Render("github:"))
	b.WriteString("\n")
	b.WriteString(styles.KeyStyle.Render(project.GitHubURL))
	b.WriteString("\n\n")

	dateStr := fmt.Sprintf("started: %s", project.StartDate.Format("jan 2006"))
	if project.EndDate != nil {
		dateStr += fmt.Sprintf(" | completed: %s", project.EndDate.Format("jan 2006"))
	}
	b.WriteString(styles.DescStyle.Render(dateStr))
	b.WriteString("\n\n")

	helpText := styles.HelpStyle.Render(
		styles.KeyStyle.Render("esc") + " back | " +
		styles.KeyStyle.Render("q") + " quit",
	)
	
	b.WriteString(lipgloss.PlaceHorizontal(width, lipgloss.Center, helpText))

	return b.String()
}