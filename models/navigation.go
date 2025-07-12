package models

type Section int

const (
	SectionWelcome Section = iota
	SectionProjects
	SectionSkills
	SectionExperience
	SectionContact
	SectionExit
)

var SectionNames = map[Section]string{
	SectionWelcome:    "♡ welcome ♡",
	SectionProjects:   "✨ projects ✨",
	SectionSkills:     "💖 skills 💖",
	SectionExperience: "🌸 experience 🌸",
	SectionContact:    "💌 contact 💌",
	SectionExit:       "👋 exit 👋",
}

type NavigationState struct {
	CurrentSection   Section
	SelectedItem     int
	ProjectDetailView bool
	ProjectIndex     int
}

func NewNavigationState() NavigationState {
	return NavigationState{
		CurrentSection:   SectionWelcome,
		SelectedItem:     0,
		ProjectDetailView: false,
		ProjectIndex:     0,
	}
}