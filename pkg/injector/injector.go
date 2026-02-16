package injector

import (
	"fmt"
	"strings"
)

const (
	defaultMarkerBegin = "<!-- BEGIN: dock-docs -->"
	defaultMarkerEnd   = "<!-- END: dock-docs -->"
)

// Inject replaces content between markers in the given file content.
// If markerName is empty, uses the default markers.
// Returns the modified content and an error if markers are missing.
func Inject(fileContent, markerName, newContent string) (string, error) {
	markerBegin := defaultMarkerBegin
	markerEnd := defaultMarkerEnd

	if markerName != "" {
		markerBegin = fmt.Sprintf("<!-- BEGIN: dock-docs:%s -->", markerName)
		markerEnd = fmt.Sprintf("<!-- END: dock-docs:%s -->", markerName)
	}

	startIdx := strings.Index(fileContent, markerBegin)
	endIdx := strings.Index(fileContent, markerEnd)

	if startIdx == -1 || endIdx == -1 {
		return "", fmt.Errorf("markers '%s' and '%s' not found", markerBegin, markerEnd)
	}

	if endIdx <= startIdx {
		return "", fmt.Errorf("end marker found before start marker")
	}

	// Preserve markers
	return fileContent[:startIdx] + markerBegin + "\n" + newContent + "\n" + markerEnd + fileContent[endIdx+len(markerEnd):], nil
}
