package themes_test

import (
	"testing"

	"go.withmatt.com/themes"
)

func TestGetTheme(t *testing.T) {
	// Test getting a known theme
	theme, err := themes.GetTheme("Dracula")
	if err != nil {
		t.Fatalf("Failed to get Dracula theme: %v", err)
	}

	if theme.Name != "Dracula" {
		t.Errorf("Expected name 'Dracula', got '%s'", theme.Name)
	}

	if theme.Background != "#282a36" {
		t.Errorf("Expected background '#282a36', got '%s'", theme.Background)
	}

	if theme.Foreground != "#f8f8f2" {
		t.Errorf("Expected foreground '#f8f8f2', got '%s'", theme.Foreground)
	}
}

func TestGetThemeCaseInsensitive(t *testing.T) {
	// Test case-insensitive lookup
	theme1, err1 := themes.GetTheme("Dracula")
	theme2, err2 := themes.GetTheme("dracula")
	theme3, err3 := themes.GetTheme("DRACULA")

	if err1 != nil || err2 != nil || err3 != nil {
		t.Fatal("Case-insensitive lookup failed")
	}

	if theme1.Name != theme2.Name || theme2.Name != theme3.Name {
		t.Error("Case-insensitive lookups returned different themes")
	}
}

func TestGetThemeNotFound(t *testing.T) {
	_, err := themes.GetTheme("NonExistentTheme12345")
	if err != themes.ErrThemeNotFound {
		t.Errorf("Expected ErrThemeNotFound, got %v", err)
	}
}

func TestListThemes(t *testing.T) {
	themeList := themes.ListThemes()
	if len(themeList) == 0 {
		t.Fatal("ListThemes returned empty list")
	}

	// We should have hundreds of themes
	if len(themeList) < 400 {
		t.Errorf("Expected at least 400 themes, got %d", len(themeList))
	}

	t.Logf("Found %d themes", len(themeList))
}

func TestGetAllThemes(t *testing.T) {
	allThemes := themes.GetAllThemes()
	if len(allThemes) == 0 {
		t.Fatal("GetAllThemes returned empty map")
	}

	// Verify we can access themes from the map
	if theme, exists := allThemes["Dracula"]; !exists {
		t.Error("Dracula theme not found in GetAllThemes")
	} else if theme.Name != "Dracula" {
		t.Errorf("Theme name mismatch: %s", theme.Name)
	}
}
