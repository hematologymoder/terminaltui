package views

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/portfolio-tui/data"
	"github.com/portfolio-tui/lastfm"
	"github.com/portfolio-tui/models"
	"github.com/portfolio-tui/styles"
)

const asciiArt = `
  ╭─────────────────────────────────────╮
  │         emily's portfolio           │
  │                                     │
  ╰─────────────────────────────────────╯
`

func RenderWelcomeWithTyping(app models.App, height int) string {
	content := app.Content
	width := app.Width
	
	if app.Nav.CurrentSection == models.SectionWelcome && !app.TypingDone {
		typedContent := *content
		if app.TypingIndex < len(content.Introduction[0]) {
			typedContent.Introduction = []string{
				content.Introduction[0][:app.TypingIndex] + "▋",
				content.Introduction[1],
				content.Introduction[2],
			}
		}
		return renderWelcomeInternalWithNP(&typedContent, width, height, app.NowPlaying)
	}
	
	return renderWelcomeInternalWithNP(content, width, height, app.NowPlaying)
}

func RenderWelcome(content *data.PortfolioContent, width, height int) string {
	return renderWelcomeInternal(content, width, height)
}

func renderWelcomeInternal(content *data.PortfolioContent, width, height int) string {
	return renderWelcomeInternalWithNP(content, width, height, nil)
}

func renderWelcomeInternalWithNP(content *data.PortfolioContent, width, height int, nowPlaying *lastfm.NowPlaying) string {
	var b strings.Builder

	centeredArt := lipgloss.PlaceHorizontal(width, lipgloss.Center, styles.ASCIIStyle.Render(asciiArt))
	b.WriteString(centeredArt)
	b.WriteString("\n\n")

	nameTitle := fmt.Sprintf("%s - %s", content.Name, content.Title)
	centeredName := lipgloss.PlaceHorizontal(width, lipgloss.Center, styles.TitleStyle.Render(nameTitle))
	b.WriteString(centeredName)
	b.WriteString("\n")

	centeredRole := lipgloss.PlaceHorizontal(width, lipgloss.Center, styles.SubsectionStyle.Render(content.CurrentRole))
	b.WriteString(centeredRole)
	b.WriteString("\n\n")

	introSection := styles.BorderStyle.Width(width - 4).Render(
		strings.Join(content.Introduction, "\n\n"),
	)
	b.WriteString(introSection)
	b.WriteString("\n\n")

	b.WriteString(styles.SubsectionStyle.Render("✨ my highlights ✨"))
	b.WriteString("\n")
	
	for _, highlight := range content.Highlights {
		b.WriteString(styles.ListItemStyle.Render(fmt.Sprintf("  • %s", highlight)))
		b.WriteString("\n")
	}
	b.WriteString("\n")

	if nowPlaying != nil && nowPlaying.IsPlaying {
		nowPlayingBox := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(styles.Primary).
			Padding(0, 1).
			Width(40)
		
		npContent := fmt.Sprintf("🎵 now playing: %s\n   by %s", 
			nowPlaying.Track, 
			nowPlaying.Artist)
			
		if nowPlaying.Album != "" {
			npContent += fmt.Sprintf("\n   from %s", nowPlaying.Album)
		}
		
		renderedNP := nowPlayingBox.Render(npContent)
		b.WriteString(lipgloss.PlaceHorizontal(width, lipgloss.Center, renderedNP))
		b.WriteString("\n\n")
	}

	navHelp := styles.HelpStyle.Render(
		"nav: " +
			styles.KeyStyle.Render("←→") + "/" + styles.KeyStyle.Render("h/l") + " sections | " +
			styles.KeyStyle.Render("↑↓") + "/" + styles.KeyStyle.Render("j/k") + " move | " +
			styles.KeyStyle.Render("enter") + " select | " +
			styles.KeyStyle.Render("?") + " help",
	)
	
	footer := lipgloss.PlaceHorizontal(width, lipgloss.Center, navHelp)
	
	output := b.String()
	fullHeight := height - lipgloss.Height(output) - lipgloss.Height(footer) - 2
	if fullHeight > 0 {
		output += strings.Repeat("\n", fullHeight)
	}
	output += footer

	return output
}
