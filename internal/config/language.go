package config

import "strings"

const (
	DefaultLanguage = "zh-CN"
	EnglishLanguage = "en"
)

func NormalizeLanguage(value string) string {
	normalized := strings.ToLower(strings.TrimSpace(value))
	switch normalized {
	case "en", "en-us", "en_us":
		return EnglishLanguage
	case "zh", "zh-cn", "zh_cn", "zh-hans", "zh_hans":
		return DefaultLanguage
	default:
		if normalized == "" {
			return DefaultLanguage
		}
	}
	return DefaultLanguage
}

func GetLanguage() string {
	cfg := ReadConfig()
	return NormalizeLanguage(cfg.System.Language)
}
