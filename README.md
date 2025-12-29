# go-themes

Terminal color schemes from [iTerm2-Color-Schemes](https://github.com/mbadolato/iTerm2-Color-Schemes) packaged as Go structs. Includes 450+ themes like Dracula, Nord, Gruvbox, Monokai, etc.

Themes are generated at build time and compiled into your binary, so there's no runtime parsing overhead. Colors are provided as hex strings that work directly with Lipgloss and most TUI libraries.

## Installation

```bash
go get go.withmatt.com/themes
```

## Usage

```go
theme, err := themes.GetTheme("Dracula")
if err != nil {
    log.Fatal(err)
}

fmt.Println(theme.Foreground)  // #f8f8f2
fmt.Println(theme.Background)  // #282a36
```

With Lipgloss:

```go
theme, _ := themes.GetTheme("Nord")

style := lipgloss.NewStyle().
    Foreground(lipgloss.Color(theme.Foreground)).
    Background(lipgloss.Color(theme.Background))

fmt.Println(style.Render("Hello, Nord!"))
```

List all themes:

```go
for _, name := range themes.ListThemes() {
    fmt.Println(name)
}
```

## API

```go
// Get a theme by name (case-insensitive)
func GetTheme(name string) (*Theme, error)

// List all theme names
func ListThemes() []string

// Get all themes as a map
func GetAllThemes() map[string]*Theme
```

The `Theme` struct contains `Foreground`, `Background`, `Cursor`, and all 16 ANSI colors (`Black`, `Red`, `Green`, `Yellow`, `Blue`, `Magenta`, `Cyan`, `White`, and their `Bright*` variants). All colors are lowercase hex strings like `#282a36`.

## Updating

To regenerate themes from the latest upstream:

```bash
go generate ./...
```

This clones [iTerm2-Color-Schemes](https://github.com/mbadolato/iTerm2-Color-Schemes), parses the JSON files, and generates `generated_themes.go`.

## License

MIT License. Color schemes are from [iTerm2-Color-Schemes](https://github.com/mbadolato/iTerm2-Color-Schemes) by [@mbadolato](https://github.com/mbadolato), also MIT licensed.
