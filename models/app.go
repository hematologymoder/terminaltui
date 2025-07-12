package models

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/portfolio-tui/config"
	"github.com/portfolio-tui/data"
	"github.com/portfolio-tui/lastfm"
)

type KeyMap struct {
	Up     key.Binding
	Down   key.Binding
	Left   key.Binding
	Right  key.Binding
	Enter  key.Binding
	Back   key.Binding
	Quit   key.Binding
	Help   key.Binding
}

var DefaultKeyMap = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "down"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "right"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc", "backspace"),
		key.WithHelp("esc", "back"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
}

type App struct {
	Nav        NavigationState
	Keys       KeyMap
	Help       help.Model
	Content    *data.PortfolioContent
	Width      int
	Height     int
	ShowHelp   bool
	Ready      bool
	TypingIndex int
	TypingDone  bool
	LastFM      *lastfm.Client
	NowPlaying  *lastfm.NowPlaying
	Config      *config.Config
}

func NewApp() App {
	cfg, _ := config.Load()
	
	app := App{
		Nav:      NewNavigationState(),
		Keys:     DefaultKeyMap,
		Help:     help.New(),
		Content:  data.GetPortfolioContent(),
		ShowHelp: false,
		Ready:    false,
		TypingIndex: 0,
		TypingDone:  false,
		Config:      cfg,
	}
	
	if cfg != nil && cfg.LastFM.APIKey != "" && cfg.LastFM.Username != "" {
		app.LastFM = lastfm.NewClient(cfg.LastFM.APIKey, cfg.LastFM.Username)
	}
	
	return app
}

func (m App) Init() tea.Cmd {
	return tea.Batch(
		StartTyping(),
		FetchNowPlaying(m.LastFM),
	)
}

func (m App) Update(msg tea.Msg) (App, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.Ready = true
		m.Help.Width = msg.Width

	case TypingMsg:
		if m.Nav.CurrentSection == SectionWelcome && !m.TypingDone {
			introText := m.Content.Introduction[0]
			if msg.Index < len(introText) {
				m.TypingIndex = msg.Index + 1
				return m, NextChar(m.TypingIndex)
			} else {
				m.TypingDone = true
			}
		}
		return m, nil
		
	case NowPlayingMsg:
		if msg.Error == nil {
			m.NowPlaying = msg.NowPlaying
		}
		return m, RefreshNowPlaying(m.LastFM)

	case tea.KeyMsg:
		if !m.TypingDone {
			m.TypingDone = true
			m.TypingIndex = len(m.Content.Introduction[0])
		}
		
		switch {
		case key.Matches(msg, m.Keys.Quit):
			if m.Nav.CurrentSection == SectionExit || (!m.Nav.ProjectDetailView && m.Nav.CurrentSection == SectionWelcome) {
				return m, tea.Quit
			}
			m.Nav.CurrentSection = SectionExit
			return m, nil

		case key.Matches(msg, m.Keys.Help):
			m.ShowHelp = !m.ShowHelp
			return m, nil

		case key.Matches(msg, m.Keys.Back):
			if m.Nav.ProjectDetailView {
				m.Nav.ProjectDetailView = false
			} else if m.Nav.CurrentSection != SectionWelcome {
				m.Nav.CurrentSection = SectionWelcome
			}
			return m, nil

		case key.Matches(msg, m.Keys.Up):
			if m.Nav.ProjectDetailView {
				return m, nil
			}
			
			switch m.Nav.CurrentSection {
			case SectionWelcome:
				if m.Nav.CurrentSection > 0 {
					m.Nav.CurrentSection--
				}
			case SectionProjects:
				if m.Nav.SelectedItem > 0 {
					m.Nav.SelectedItem--
				}
			case SectionSkills:
				if m.Nav.SelectedItem > 0 {
					m.Nav.SelectedItem--
				}
			}

		case key.Matches(msg, m.Keys.Down):
			if m.Nav.ProjectDetailView {
				return m, nil
			}
			
			switch m.Nav.CurrentSection {
			case SectionWelcome:
				if int(m.Nav.CurrentSection) < len(SectionNames)-1 {
					m.Nav.CurrentSection++
				}
			case SectionProjects:
				if m.Nav.SelectedItem < len(m.Content.Projects)-1 {
					m.Nav.SelectedItem++
				}
			case SectionSkills:
				if m.Nav.SelectedItem < len(m.Content.SkillCategories)-1 {
					m.Nav.SelectedItem++
				}
			}

		case key.Matches(msg, m.Keys.Left):
			if int(m.Nav.CurrentSection) > 0 {
				m.Nav.CurrentSection--
				m.Nav.SelectedItem = 0
			}

		case key.Matches(msg, m.Keys.Right):
			if int(m.Nav.CurrentSection) < len(SectionNames)-1 {
				m.Nav.CurrentSection++
				m.Nav.SelectedItem = 0
			}

		case key.Matches(msg, m.Keys.Enter):
			switch m.Nav.CurrentSection {
			case SectionExit:
				return m, tea.Quit
			case SectionProjects:
				if !m.Nav.ProjectDetailView {
					m.Nav.ProjectDetailView = true
					m.Nav.ProjectIndex = m.Nav.SelectedItem
				}
			}
		}
	}

	return m, nil
}

