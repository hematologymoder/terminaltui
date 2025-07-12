package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/portfolio-tui/models"
	"github.com/portfolio-tui/views"
)

type appWithView struct {
	models.App
}

func (a appWithView) Init() tea.Cmd {
	return a.App.Init()
}

func (a appWithView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	updatedApp, cmd := a.App.Update(msg)
	return appWithView{App: updatedApp}, cmd
}

func (a appWithView) View() string {
	return views.RenderView(a.App)
}

func main() {
	m := models.NewApp()
	app := appWithView{App: m}

	p := tea.NewProgram(app, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println("Thanks for visiting!")
}