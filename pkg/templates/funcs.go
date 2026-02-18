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
			return getEmoji(name, noMoji)
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

// getEmoji returns an emoji or text alternative based on the noMoji flag.
func getEmoji(name string, noMoji bool) string {
	if noMoji {
		switch name {
		case "check":
			return "[YES]"
		case "cross":
			return "[NO]"
		case "critical":
			return "[CRIT]"
		case "high":
			return "[HIGH]"
		case "medium":
			return "[MED]"
		case "low":
			return "[LOW]"
		case "clean":
			return "[OK]"
		default:
			return ""
		}
	}

	switch name {
	case "whale":
		return "\U0001F433 "
	case "gear":
		return "\u2699\uFE0F "
	case "shield":
		return "\U0001F6E1\uFE0F "
	case "tag":
		return "\U0001F3F7\uFE0F "
	case "search":
		return "\U0001F50D "
	case "down":
		return "\U0001F447 "
	case "package":
		return "\U0001F4E6 "
	case "check":
		return "\u2705"
	case "cross":
		return "\u274C"
	case "critical":
		return "\U0001F6D1"
	case "high":
		return "\U0001F7E0"
	case "medium":
		return "\U0001F7E1"
	case "low":
		return "\U0001F535"
	case "clean":
		return "\U0001F7E2"
	default:
		return ""
	}
}
