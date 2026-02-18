package templates

import (
	"strings"
	"text/template"
)

// GetFuncMap returns the template function map used by all templates.
// It includes both standard helper functions and rendering-specific functions.
func GetFuncMap(noMoji bool) template.FuncMap {
	return template.FuncMap{
		// Map/slice helpers (existing)
		"index": safeMapIndex,
		"join":  strings.Join,

		// String helpers
		"lower":      strings.ToLower,
		"upper":      strings.ToUpper,
		"trim":       strings.TrimSpace,
		"contains":   strings.Contains,
		"hasPrefix":  strings.HasPrefix,
		"hasSuffix":  strings.HasSuffix,
		"replace":    strings.ReplaceAll,
		"default":    defaultValue,
		"jsonEscape": jsonEscape,

		// Emoji helper
		"emoji": func(name string) string {
			return GetEmoji(name, noMoji)
		},
	}
}

// safeMapIndex safely accesses a map[string]int, returning 0 for missing keys.
func safeMapIndex(m map[string]int, k string) int {
	if v, ok := m[k]; ok {
		return v
	}
	return 0
}

// defaultValue returns the fallback if val is empty.
func defaultValue(fallback string, val string) string {
	if val == "" {
		return fallback
	}
	return val
}

// jsonEscape escapes a string for safe use inside JSON string values.
func jsonEscape(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `"`, `\"`)
	s = strings.ReplaceAll(s, "\n", `\n`)
	s = strings.ReplaceAll(s, "\r", `\r`)
	s = strings.ReplaceAll(s, "\t", `\t`)
	return s
}

var emojiMap = map[string]string{
	"whale":    "\U0001F433 ",
	"gear":     "\u2699\uFE0F ",
	"shield":   "\U0001F6E1\uFE0F ",
	"tag":      "\U0001F3F7\uFE0F ",
	"search":   "\U0001F50D ",
	"down":     "\U0001F447 ",
	"package":  "\U0001F4E6 ",
	"check":    "\u2705",
	"cross":    "\u274C",
	"critical": "\U0001F6D1",
	"high":     "\U0001F7E0",
	"medium":   "\U0001F7E1",
	"low":      "\U0001F535",
	"clean":    "\U0001F7E2",
}

var noMojiMap = map[string]string{
	"check":    "[YES]",
	"cross":    "[NO]",
	"critical": "[CRIT]",
	"high":     "[HIGH]",
	"medium":   "[MED]",
	"low":      "[LOW]",
	"clean":    "[OK]",
}

// GetEmoji returns an emoji or text alternative based on the noMoji flag.
func GetEmoji(name string, noMoji bool) string {
	if noMoji {
		return noMojiMap[name]
	}
	return emojiMap[name]
}
