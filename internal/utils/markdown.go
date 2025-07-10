package utils

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

var md goldmark.Markdown

func init() {
	md = goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Table,
			extension.Strikethrough,
			extension.TaskList,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
		),
	)
}

// MarkdownToHTML converts markdown content to HTML
func MarkdownToHTML(content string) template.HTML {
	var buf bytes.Buffer
	if err := md.Convert([]byte(content), &buf); err != nil {
		return template.HTML("")
	}
	return template.HTML(buf.String())
}

// GetSummary extracts first paragraph or first N characters as summary
func GetSummary(content string, maxLength int) string {
	// Remove markdown formatting for summary
	plainText := strings.ReplaceAll(content, "#", "")
	plainText = strings.ReplaceAll(plainText, "*", "")
	plainText = strings.ReplaceAll(plainText, "_", "")
	plainText = strings.ReplaceAll(plainText, "[", "")
	plainText = strings.ReplaceAll(plainText, "]", "")
	plainText = strings.ReplaceAll(plainText, "(", "")
	plainText = strings.ReplaceAll(plainText, ")", "")
	
	// Get first paragraph
	paragraphs := strings.Split(plainText, "\n\n")
	summary := paragraphs[0]
	
	// Trim to maxLength if needed
	if len(summary) > maxLength {
		summary = summary[:maxLength] + "..."
	}
	
	return strings.TrimSpace(summary)
} 