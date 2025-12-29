package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

// WindowsTerminalTheme represents the structure of Windows Terminal JSON color schemes
type WindowsTerminalTheme struct {
	Name                string `json:"name"`
	Foreground          string `json:"foreground"`
	Background          string `json:"background"`
	CursorColor         string `json:"cursorColor"`
	Black               string `json:"black"`
	Red                 string `json:"red"`
	Green               string `json:"green"`
	Yellow              string `json:"yellow"`
	Blue                string `json:"blue"`
	Purple              string `json:"purple"` // Maps to Magenta
	Cyan                string `json:"cyan"`
	White               string `json:"white"`
	BrightBlack         string `json:"brightBlack"`
	BrightRed           string `json:"brightRed"`
	BrightGreen         string `json:"brightGreen"`
	BrightYellow        string `json:"brightYellow"`
	BrightBlue          string `json:"brightBlue"`
	BrightPurple        string `json:"brightPurple"` // Maps to BrightMagenta
	BrightCyan          string `json:"brightCyan"`
	BrightWhite         string `json:"brightWhite"`
	SelectionBackground string `json:"selectionBackground"` // Not used in our Theme struct
}

// Theme represents our output theme format
type Theme struct {
	Name          string
	Foreground    string
	Background    string
	Cursor        string
	Black         string
	Red           string
	Green         string
	Yellow        string
	Blue          string
	Magenta       string
	Cyan          string
	White         string
	BrightBlack   string
	BrightRed     string
	BrightGreen   string
	BrightYellow  string
	BrightBlue    string
	BrightMagenta string
	BrightCyan    string
	BrightWhite   string
}

func main() {
	log.Println("Generating themes from iTerm2-Color-Schemes...")

	// Clone the repository to a temporary directory
	tmpDir, err := os.MkdirTemp("", "iterm2-color-schemes-*")
	if err != nil {
		log.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	log.Printf("Cloning repository to %s...", tmpDir)
	cmd := exec.Command("git", "clone", "--depth=1", "https://github.com/mbadolato/iTerm2-Color-Schemes.git", tmpDir)
	if output, err := cmd.CombinedOutput(); err != nil {
		log.Fatalf("Failed to clone repository: %v\n%s", err, output)
	}

	// Parse themes from JSON files
	schemesDir := filepath.Join(tmpDir, "windowsterminal")
	themes, err := parseThemes(schemesDir)
	if err != nil {
		log.Fatalf("Failed to parse themes: %v", err)
	}

	log.Printf("Parsed %d themes", len(themes))

	// Generate Go code
	if err := generateCode(themes); err != nil {
		log.Fatalf("Failed to generate code: %v", err)
	}

	log.Println("Successfully generated generated_themes.go")
}

func parseThemes(schemesDir string) ([]Theme, error) {
	files, err := filepath.Glob(filepath.Join(schemesDir, "*.json"))
	if err != nil {
		return nil, fmt.Errorf("failed to list JSON files: %w", err)
	}

	var themes []Theme
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			log.Printf("Warning: failed to read %s: %v", file, err)
			continue
		}

		var wtTheme WindowsTerminalTheme
		if err := json.Unmarshal(data, &wtTheme); err != nil {
			log.Printf("Warning: failed to parse %s: %v", file, err)
			continue
		}

		// Use filename as name if name field is empty
		name := wtTheme.Name
		if name == "" {
			name = strings.TrimSuffix(filepath.Base(file), ".json")
		}

		theme := Theme{
			Name:          name,
			Foreground:    normalizeColor(wtTheme.Foreground),
			Background:    normalizeColor(wtTheme.Background),
			Cursor:        normalizeColor(wtTheme.CursorColor),
			Black:         normalizeColor(wtTheme.Black),
			Red:           normalizeColor(wtTheme.Red),
			Green:         normalizeColor(wtTheme.Green),
			Yellow:        normalizeColor(wtTheme.Yellow),
			Blue:          normalizeColor(wtTheme.Blue),
			Magenta:       normalizeColor(wtTheme.Purple), // Purple -> Magenta
			Cyan:          normalizeColor(wtTheme.Cyan),
			White:         normalizeColor(wtTheme.White),
			BrightBlack:   normalizeColor(wtTheme.BrightBlack),
			BrightRed:     normalizeColor(wtTheme.BrightRed),
			BrightGreen:   normalizeColor(wtTheme.BrightGreen),
			BrightYellow:  normalizeColor(wtTheme.BrightYellow),
			BrightBlue:    normalizeColor(wtTheme.BrightBlue),
			BrightMagenta: normalizeColor(wtTheme.BrightPurple), // BrightPurple -> BrightMagenta
			BrightCyan:    normalizeColor(wtTheme.BrightCyan),
			BrightWhite:   normalizeColor(wtTheme.BrightWhite),
		}

		themes = append(themes, theme)
	}

	// Sort themes by name for consistent output
	sort.Slice(themes, func(i, j int) bool {
		return themes[i].Name < themes[j].Name
	})

	return themes, nil
}

// normalizeColor ensures color is properly formatted (lowercase hex with #)
func normalizeColor(color string) string {
	if color == "" {
		return "#000000" // Default to black
	}
	// Ensure it starts with # and is lowercase
	color = strings.ToLower(color)
	if !strings.HasPrefix(color, "#") {
		color = "#" + color
	}
	return color
}

func generateCode(themes []Theme) error {
	tmpl := `// Code generated by go generate; DO NOT EDIT.

package themes

var allThemes = map[string]*Theme{
{{- range . }}
	"{{ .Name }}": {
		Name:          "{{ .Name }}",
		Foreground:    "{{ .Foreground }}",
		Background:    "{{ .Background }}",
		Cursor:        "{{ .Cursor }}",
		Black:         "{{ .Black }}",
		Red:           "{{ .Red }}",
		Green:         "{{ .Green }}",
		Yellow:        "{{ .Yellow }}",
		Blue:          "{{ .Blue }}",
		Magenta:       "{{ .Magenta }}",
		Cyan:          "{{ .Cyan }}",
		White:         "{{ .White }}",
		BrightBlack:   "{{ .BrightBlack }}",
		BrightRed:     "{{ .BrightRed }}",
		BrightGreen:   "{{ .BrightGreen }}",
		BrightYellow:  "{{ .BrightYellow }}",
		BrightBlue:    "{{ .BrightBlue }}",
		BrightMagenta: "{{ .BrightMagenta }}",
		BrightCyan:    "{{ .BrightCyan }}",
		BrightWhite:   "{{ .BrightWhite }}",
	},
{{- end }}
}
`

	t, err := template.New("themes").Parse(tmpl)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Generate code into buffer
	var buf bytes.Buffer
	if err := t.Execute(&buf, themes); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	// Format the generated code
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("failed to format generated code: %w", err)
	}

	// Write formatted code to file
	if err := os.WriteFile("generated_themes.go", formatted, 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return nil
}
