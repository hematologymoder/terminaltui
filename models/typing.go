package models

import (
	"time"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/portfolio-tui/lastfm"
)

type TypingMsg struct {
	Index int
}

func StartTyping() tea.Cmd {
	return tea.Tick(time.Millisecond*50, func(t time.Time) tea.Msg {
		return TypingMsg{Index: 0}
	})
}

func NextChar(index int) tea.Cmd {
	return tea.Tick(time.Millisecond*50, func(t time.Time) tea.Msg {
		return TypingMsg{Index: index}
	})
}

type NowPlayingMsg struct {
	NowPlaying *lastfm.NowPlaying
	Error      error
}

func FetchNowPlaying(client *lastfm.Client) tea.Cmd {
	if client == nil {
		return nil
	}
	
	return func() tea.Msg {
		np, err := client.GetNowPlaying()
		return NowPlayingMsg{
			NowPlaying: np,
			Error:      err,
		}
	}
}

func RefreshNowPlaying(client *lastfm.Client) tea.Cmd {
	return tea.Tick(time.Second*30, func(t time.Time) tea.Msg {
		return FetchNowPlaying(client)()
	})
}