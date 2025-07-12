package views

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/portfolio-tui/models"
	"github.com/portfolio-tui/styles"
)

func RenderView(app models.App) string {
	if !app.Ready {
		return "Initializing..."
	}

	navBar := renderNavBar(app.Nav.CurrentSection, app.Width)
	
	contentHeight := app.Height - lipgloss.Height(navBar) - 2
	
	var content string
	switch app.Nav.CurrentSection {
	case models.SectionWelcome:
		content = RenderWelcomeWithTyping(app, contentHeight)
	
	case models.SectionProjects:
		if app.Nav.ProjectDetailView {
			project := app.Content.Projects[app.Nav.ProjectIndex]
			content = RenderProjectDetail(project, app.Width, contentHeight)
		} else {
			content = RenderProjectsList(app.Content.Projects, app.Nav.SelectedItem, app.Width, contentHeight)
		}
	
	case models.SectionSkills:
		content = RenderSkills(app.Content.SkillCategories, app.Nav.SelectedItem, app.Width, contentHeight)
	
	case models.SectionExperience:
		content = RenderExperience(app.Content.Experiences, app.Content.Education, app.Width, contentHeight)
	
	case models.SectionContact:
		content = RenderContact(app.Content.Contact, app.Width, contentHeight)
	
	case models.SectionExit:
		content = renderExitConfirmation(app.Width, contentHeight)
	}

	if app.ShowHelp {
		content = renderHelpOverlay(app.Width, app.Height)
	}

	return navBar + "\n" + content
}

func renderNavBar(currentSection models.Section, width int) string {
	sections := []models.Section{
		models.SectionWelcome,
		models.SectionProjects,
		models.SectionSkills,
		models.SectionExperience,
		models.SectionContact,
		models.SectionExit,
	}

	var navItems []string
	for _, section := range sections {
		name := models.SectionNames[section]
		if section == currentSection {
			navItems = append(navItems, styles.NavItemSelectedStyle.Render(name))
		} else {
			navItems = append(navItems, styles.NavItemStyle.Render(name))
		}
	}

	navBar := strings.Join(navItems, "")
	centeredNav := lipgloss.PlaceHorizontal(width, lipgloss.Center, navBar)
	
	border := strings.Repeat("─", width)
	
	return centeredNav + "\n" + styles.DescStyle.Render(border)
}

func renderExitConfirmation(width, height int) string {
	var b strings.Builder

	exitBox := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(styles.Warning).
		Padding(2, 4).
		Width(50)

	exitContent := styles.TitleStyle.Render("aww, leaving already? 🥺") + "\n\n" +
		"thank you so much for visiting! \n" +
		"i hope you enjoyed your stay~ 💕\n\n" +
		styles.KeyStyle.Render("enter") + " to say goodbye\n" +
		styles.KeyStyle.Render("esc") + " to stay a bit longer"

	renderedBox := exitBox.Render(exitContent)
	
	horizontallyCentered := lipgloss.PlaceHorizontal(width, lipgloss.Center, renderedBox)
	verticallyCentered := lipgloss.PlaceVertical(height, lipgloss.Center, horizontallyCentered)
	
	b.WriteString(verticallyCentered)
	
	return b.String()
}

func renderHelpOverlay(width, height int) string {
	helpBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(styles.Primary).
		Padding(2, 4).
		Width(60)

	helpContent := styles.TitleStyle.Render("💝 help & controls 💝") + "\n\n" +
		styles.SubsectionStyle.Render("✨ navigation:") + "\n" +
		fmt.Sprintf("  %s  move up\n", styles.KeyStyle.Render("↑/k")) +
		fmt.Sprintf("  %s  move down\n", styles.KeyStyle.Render("↓/j")) +
		fmt.Sprintf("  %s  previous section\n", styles.KeyStyle.Render("←/h")) +
		fmt.Sprintf("  %s  next section\n", styles.KeyStyle.Render("→/l")) +
		"\n" +
		styles.SubsectionStyle.Render("🌸 actions:") + "\n" +
		fmt.Sprintf("  %s  select/view details\n", styles.KeyStyle.Render("enter")) +
		fmt.Sprintf("  %s  go back\n", styles.KeyStyle.Render("esc")) +
		fmt.Sprintf("  %s    toggle this help\n", styles.KeyStyle.Render("?")) +
		fmt.Sprintf("  %s    quit application\n", styles.KeyStyle.Render("q")) +
		"\n" +
		styles.DescStyle.Render("press any key to close 💕")

	renderedBox := helpBox.Render(helpContent)
	
	horizontallyCentered := lipgloss.PlaceHorizontal(width, lipgloss.Center, renderedBox)
	verticallyCentered := lipgloss.PlaceVertical(height, lipgloss.Center, horizontallyCentered)
	
	return verticallyCentered
}