package data

import "time"

type Project struct {
	Name        string
	Description string
	LongDesc    string
	Technologies []string
	GitHubURL   string
	Features    []string
	Status      string
	StartDate   time.Time
	EndDate     *time.Time
}

type Skill struct {
	Name       string
	Level      int // 1-5
	Years      int
}

type SkillCategory struct {
	Name   string
	Skills []Skill
}

type Experience struct {
	Company         string
	Role            string
	StartDate       time.Time
	EndDate         *time.Time
	Location        string
	Responsibilities []string
	Achievements    []string
}

type Education struct {
	Institution string
	Degree      string
	Field       string
	StartYear   int
	EndYear     int
	Location    string
}

type Contact struct {
	Email      string
	GitHub     string
	LinkedIn   string
	Website    string
	Location   string
}

type PortfolioContent struct {
	Name            string
	Title           string
	Introduction    []string
	CurrentRole     string
	Highlights      []string
	Projects        []Project
	SkillCategories []SkillCategory
	Experiences     []Experience
	Education       []Education
	Contact         Contact
}

func GetPortfolioContent() *PortfolioContent {
	return &PortfolioContent{
		Name:  "emily",
		Title: "student, developer, endocrinologist",
		Introduction: []string{
			"hiiii!! welcome to my cute little terminal portfolio! i'm emily.",
			"im interested in endocrinology and i write about hrt and nhtc at papers.adenine.xyz",
			"this terminal interface was built with go and bubble tea, dm me if you want a template like this or if you want the repo.",
		},
		CurrentRole: "high school student & self-taught developer",
		Highlights: []string{
			"self-taught endocrinologist",
			"aws certified in 4 areas (cloud practitioner, developer, solutions architect, sysops)",
			"i run and make some very simple websites",
		},
		Projects: []Project{
			{
				Name:        "",
				Description: "",
				LongDesc:    "",
				Technologies: []string{"next.js", "react", "tailwind", "aws lambda", "postgresql"},
				GitHubURL:   "",
				Features: []string{
					"",
					"",
					"",
					"",
					"",
				},
				Status:    "active",
				StartDate: time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC),
			},
			{
				Name:        "papers.adenine.xyz",
				Description: "simple website for hosting papers that i write with one user, me. extremely simple ui, frontend, and backend.",
				LongDesc:    "",
				Technologies: []string{"javascript", "htmx", "postgresql", "css"},
				GitHubURL:   "",
				Features: []string{
					"",
					"",
					"",
					"",
					"",
				},
				Status:    "active",
				StartDate: time.Date(2025, 5, 1, 0, 0, 0, 0, time.UTC),
			},
			{
				Name:        "miellé",
				Description: "pink-themed calorie tracking app",
				LongDesc:    "the cutest calorie tracker you'll ever use! features barcode scanning, nutrition api integration, and a beautiful pink interface. tracks macros, micros, and use something that actually works",
				Technologies: []string{"swift", "swiftui", "supabase", "postgresql"},
				GitHubURL:   "github.com/adenine/mielle",
				Features: []string{
					"barcode scanning for instant nutrition",
					"gorgeous pink gradient ui",
					"cute progress charts and graphs",
					"local food database integration",
				},
				Status:    "completed",
				StartDate: time.Date(2025, 4, 2, 0, 0, 0, 0, time.UTC),
				EndDate:   &[]time.Time{time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)}[0],
			},
		},
		SkillCategories: []SkillCategory{
			{
				Name: "frontend",
				Skills: []Skill{
					{Name: "react/next.js", Level: 5, Years: 3},
					{Name: "svelte/vite", Level: 4, Years: 2},
					{Name: "tailwind css", Level: 5, Years: 3},
					{Name: "typescript", Level: 4, Years: 2},
					{Name: "htmx", Level: 3, Years: 1},
				},
			},
			{
				Name: "backend ",
				Skills: []Skill{
					{Name: "python/flask", Level: 4, Years: 3},
					{Name: "django", Level: 3, Years: 1},
					{Name: "node.js", Level: 4, Years: 2},
					{Name: "postgresql", Level: 4, Years: 2},
					{Name: "mysql", Level: 3, Years: 2},
				},
			},
			{
				Name: "cloud & devops",
				Skills: []Skill{
					{Name: "aws (certified!)", Level: 5, Years: 2},
					{Name: "docker", Level: 4, Years: 2},
					{Name: "kubernetes", Level: 3, Years: 1},
					{Name: "github actions", Level: 4, Years: 2},
					{Name: "vercel/netlify", Level: 5, Years: 3},
				},
			},
			{
				Name: "medical knowledge",
				Skills: []Skill{
					{Name: "endocrinology", Level: 4, Years: 1},
					{Name: "hematology", Level: 2, Years: 2},
					{Name: "emergency medicine", Level: 3, Years: 2},
					{Name: "biotechnology (for stock trading) ", Level: 2, Years: 2},
				},
			},
			{
				Name: "other cool stuff",
				Skills: []Skill{
					{Name: "git/github", Level: 5, Years: 3},
					{Name: "neovim (ha ha)", Level: 4, Years: 3},
					{Name: "stock trading", Level: 3, Years: 3},
					{Name: "hardware/soldering", Level: 3, Years: 2},
				},
			},
		},
		Experiences: []Experience{
			{
				Company:   "my room lmao",
				Role:      "endocrinology researcher",
				StartDate: time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
				Location:  "san francisco, ca",
				Responsibilities: []string{
					"studying endocrinology textbooks and research papers",
					"✍writing clinical-style literature with proper citations",
					"researching hrt protocols and metabolic pathways",
					"analyzing pharmaceutical compounds and mechanisms",
				},
				Achievements: []string{
					"read a lot of books (lame)",
					"wrote two papers, one of which is actualled talked about sometimes",
					"somewhat smart sometimes",
					"i run a fake analysis lab",
				},
			},
			{
				Company:   "independent projects",
				Role:      "full-stack developer",
				StartDate: time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
				Location:  "san francisco, ca",
				Responsibilities: []string{
					"building web applications from scratch",
					"deploying to my own server in my house",
					"i like designing cool interfaces",
					"also i do hackathons sometimes",
				},
				Achievements: []string{
					"placed 1st in two hackathons and won some money and got a cloudflare prize in another",
					"miellé is like the only good thing ive ever done",
					"also i made a cdn",
					"i am kinda good at optimization",
				},
			},
			{
				Company:   "stock trading",
				Role:      "options trader lmao",
				StartDate: time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC),
				Location:  "public",
				Responsibilities: []string{
					"mainly s&p and biotech options ",
					"i read the financial times for my news",
					"also biopharmcatalyst is great",
					"i use public to trade",
				},
				Achievements: []string{
					"somewhat consistent positive returns",
					"i do research i promise this isnt just gambling",
					"this helps for endocrinology a little",
					"i made a bot to buy and sell options fast but i violated FINRA rules",
				},
			},
		},
		Education: []Education{
			{
				Institution: "im not telling you",
				Degree:      "high school student",
				Field:       "general studies ig?? im 15",
				StartYear:   2024,
				EndYear:     2028,
				Location:    "san francisco, ca",
			},
		},
		Contact: Contact{
			Email:    "emily@adenine.dev",
			GitHub:   "github.com/hematologymoder",
			LinkedIn: "",
			Website:  "estrogenizedtwink.com",
			Location: "san francisco, ca",
		},
	}
}
