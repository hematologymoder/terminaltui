package views

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/portfolio-tui/data"
	"github.com/portfolio-tui/styles"
)

const contactArt = `
  ╭─────────────────────────────────────╮
  │        💌 let's chat! 💌            │
  ╰─────────────────────────────────────╯
`

func RenderContact(contact data.Contact, width, height int) string {
	var b strings.Builder

	b.WriteString(styles.SubsectionStyle.Render("💌 contact 💌"))
	b.WriteString("\n\n")

	centeredArt := lipgloss.PlaceHorizontal(width, lipgloss.Center, styles.ASCIIStyle.Render(contactArt))
	b.WriteString(centeredArt)
	b.WriteString("\n")

	contactBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(styles.Primary).
		Padding(1, 3).
		Width(50)

	var contactContent strings.Builder

	contactContent.WriteString(styles.SubsectionStyle.Render("📧 email:"))
	contactContent.WriteString("\n")
	contactContent.WriteString("   " + styles.KeyStyle.Render(contact.Email))
	contactContent.WriteString("\n\n")

	if contact.GitHub != "" {
		contactContent.WriteString(styles.SubsectionStyle.Render("💻 github:"))
		contactContent.WriteString("\n")
		contactContent.WriteString("   " + styles.KeyStyle.Render(contact.GitHub))
		contactContent.WriteString("\n\n")
	}

	if contact.Website != "" {
		contactContent.WriteString(styles.SubsectionStyle.Render("🌐 website:"))
		contactContent.WriteString("\n")
		contactContent.WriteString("   " + styles.KeyStyle.Render(contact.Website))
		contactContent.WriteString("\n\n")
	}

	contactContent.WriteString(styles.SubsectionStyle.Render("📍 location:"))
	contactContent.WriteString("\n")
	contactContent.WriteString("   " + styles.KeyStyle.Render(contact.Location))

	renderedBox := contactBox.Render(contactContent.String())
	centeredBox := lipgloss.PlaceHorizontal(width, lipgloss.Center, renderedBox)
	b.WriteString(centeredBox)
	b.WriteString("\n\n")

	messageBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(styles.Secondary).
		Padding(1, 2).
		Width(60)
	
	message := "hi there! \n\n" +
		"i love chatting about programming, endocrinology, or really anything. \n" +
		"feel free to reach out through any of the above or my signal adenine.24 \n\n" +
		"always excited to meet new people!"
	
	renderedMessage := messageBox.Render(message)
	centeredMessage := lipgloss.PlaceHorizontal(width, lipgloss.Center, renderedMessage)
	b.WriteString(centeredMessage)
	b.WriteString("\n\n")

	helpText := styles.HelpStyle.Render(
		styles.KeyStyle.Render("←→") + " sections | " +
		styles.KeyStyle.Render("esc") + " back",
	)
	
	b.WriteString(lipgloss.PlaceHorizontal(width, lipgloss.Center, helpText))

	return b.String()
}
