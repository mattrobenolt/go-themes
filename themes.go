//go:generate go run generator/main.go

// Package themes provides 450+ terminal color schemes from iTerm2-Color-Schemes
// as ready-to-use Go structs. Perfect for TUI applications built with frameworks
// like Lipgloss and Bubble Tea.
//
// All color values are provided as lowercase hex strings (e.g., "#282a36") that
// work directly with most TUI libraries without conversion.
//
// # Basic Usage
//
//	theme, err := themes.GetTheme("Dracula")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(theme.Foreground) // #f8f8f2
//
// # Using with Lipgloss
//
//	theme, _ := themes.GetTheme("Nord")
//	style := lipgloss.NewStyle().
//		Foreground(lipgloss.Color(theme.Foreground)).
//		Background(lipgloss.Color(theme.Background))
//
// # Listing Themes
//
//	themes := themes.ListThemes()
//	fmt.Printf("Available themes: %d\n", len(themes))
package themes

import (
	"errors"
	"strings"
)

// Theme represents a color scheme with all its color values as hex strings.
// All color values are in hex format (e.g., "#282a36") and work directly with
// TUI frameworks like Lipgloss.
type Theme struct {
	Name       string
	Foreground string
	Background string
	Cursor     string
	// ANSI colors 0-7
	Black   string
	Red     string
	Green   string
	Yellow  string
	Blue    string
	Magenta string
	Cyan    string
	White   string
	// ANSI bright colors 8-15
	BrightBlack   string
	BrightRed     string
	BrightGreen   string
	BrightYellow  string
	BrightBlue    string
	BrightMagenta string
	BrightCyan    string
	BrightWhite   string
}

// ErrThemeNotFound is returned when a requested theme doesn't exist.
var ErrThemeNotFound = errors.New("theme not found")

// GetTheme returns a theme by name (case-insensitive).
// Returns ErrThemeNotFound if the theme doesn't exist.
func GetTheme(name string) (*Theme, error) {
	nameLower := strings.ToLower(name)
	for themeName, theme := range allThemes {
		if strings.ToLower(themeName) == nameLower {
			return theme, nil
		}
	}
	return nil, ErrThemeNotFound
}

// ListThemes returns a sorted list of all available theme names.
func ListThemes() []string {
	themes := make([]string, 0, len(allThemes))
	for name := range allThemes {
		themes = append(themes, name)
	}
	return themes
}

// GetAllThemes returns a map of all available themes keyed by name.
func GetAllThemes() map[string]*Theme {
	return allThemes
}
