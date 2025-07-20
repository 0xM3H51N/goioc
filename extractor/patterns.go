package extractor

import "regexp"

func GetPatterns() map[string]*regexp.Regexp {
	return map[string]*regexp.Regexp{
		"ip":     regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}\b`),
		"url":    regexp.MustCompile(`\bhttps?://[^\s/$.?#].[^\s]*`),
		"md5":    regexp.MustCompile(`\b[a-fA-F0-9]{32}\b`),
		"sha256": regexp.MustCompile(`\b[a-fA-F0-9]{64}\b`),
		"email":  regexp.MustCompile(`\b[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}\b`),
	}
}
