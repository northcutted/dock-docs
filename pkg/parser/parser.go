package parser

import (
	"bufio"
	"os"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// DocItem represents a single documented instruction extracted from a Dockerfile.
type DocItem struct {
	Name        string // e.g., "PORT"
	Value       string // inferred default from the instruction
	Description string // from @description
	Type        string // "ARG", "ENV", "LABEL", "EXPOSE"
	Required    bool   // from @required
}

// Documentation holds all extracted documentation items from a Dockerfile.
type Documentation struct {
	Items []DocItem
}

// FilterByType returns items of a specific type (ARG, ENV, LABEL, EXPOSE).
func (d *Documentation) FilterByType(t string) []DocItem {
	var filtered []DocItem
	for _, item := range d.Items {
		if item.Type == t {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

// Parse reads a Dockerfile and extracts documentation metadata.
func Parse(filename string) (*Documentation, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		// Ignore error on close in defer as we are reading only
		_ = f.Close()
	}()

	result, err := parser.Parse(f)
	if err != nil {
		return nil, err
	}

	doc := &Documentation{
		Items: make([]DocItem, 0),
	}

	for _, node := range result.AST.Children {
		// 1. Parse comments into a list of metadata objects
		metas := parseComments(node)

		var items []DocItem

		switch strings.ToUpper(node.Value) {
		case "ARG":
			items = parseMultiKV(node, "ARG")
		case "ENV":
			items = parseMultiKV(node, "ENV")
		case "LABEL":
			items = parseMultiKV(node, "LABEL")
		case "EXPOSE":
			items = parseExpose(node)
		default:
			continue
		}

		// 2. Merge metadata into items
		// Strategy:
		// - If we have multiple items and multiple metas, map 1:1.
		// - If we have multiple items but only 1 meta, apply meta to first item only?
		//   User spec: "associating the single comment to all of them might look weird".
		//   So yes, map 1:1. Remaining items get no metadata (unless we decide otherwise later).
		for i := range items {
			if i < len(metas) {
				m := metas[i]
				if m.Name != "" {
					items[i].Name = m.Name
				}
				if m.Description != "" {
					items[i].Description = m.Description
				}
				if m.Value != "" {
					// @default tag overrides inferred value
					items[i].Value = m.Value
				}
				if m.Required {
					items[i].Required = true
				}
			}
			doc.Items = append(doc.Items, items[i])
		}
	}

	return doc, nil
}

func parseComments(node *parser.Node) []DocItem {
	if node.PrevComment == nil {
		return nil
	}

	var metas []DocItem
	current := DocItem{}
	seen := make(map[string]bool)
	hasContent := false

	scanner := bufio.NewScanner(strings.NewReader(strings.Join(node.PrevComment, "\n")))
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		cleanLine := strings.TrimPrefix(line, "#")
		cleanLine = strings.TrimSpace(cleanLine)

		var tag string
		if strings.HasPrefix(cleanLine, "@name:") {
			tag = "name"
		} else if strings.HasPrefix(cleanLine, "@description:") {
			tag = "desc"
		} else if strings.HasPrefix(cleanLine, "@default:") {
			tag = "default"
		} else if strings.HasPrefix(cleanLine, "@required:") {
			tag = "required"
		}

		if tag != "" {
			// If we've already seen this tag in the current block, start a new one
			if seen[tag] {
				if hasContent {
					metas = append(metas, current)
				}
				current = DocItem{}
				seen = make(map[string]bool)
				// hasContent = false // This was the ineffectual assignment reported by linter
			}
			seen[tag] = true
			hasContent = true
		}

		switch tag {
		case "name":
			current.Name = strings.TrimSpace(strings.TrimPrefix(cleanLine, "@name:"))
		case "desc":
			current.Description = strings.TrimSpace(strings.TrimPrefix(cleanLine, "@description:"))
		case "default":
			current.Value = strings.TrimSpace(strings.TrimPrefix(cleanLine, "@default:"))
		case "required":
			val := strings.TrimSpace(strings.TrimPrefix(cleanLine, "@required:"))
			if val == "true" {
				current.Required = true
			}
		}
	}

	if hasContent {
		metas = append(metas, current)
	}

	return metas
}

// Helper to strip surrounding quotes
func stripQuotes(s string) string {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}

// parseMultiKV handles ENV, ARG, LABEL which can have multiple key-value pairs (except ARG usually 1)
func parseMultiKV(node *parser.Node, typeStr string) []DocItem {
	var items []DocItem
	if node.Next == nil {
		return items
	}

	// Heuristic based on observation of Buildkit parser (v1.x?):
	// ENV K=V -> K, V, =
	// ENV K V -> K, V, ""
	// ENV K=V K2=V2 -> K, V, =, K2, V2, =

	curr := node.Next
	for curr != nil {
		key := curr.Value
		if key == "" {
			// Should not happen for key?
			curr = curr.Next
			continue
		}

		var val string
		if curr.Next != nil {
			val = curr.Next.Value
			// strip quotes? Buildkit seems to keep quotes for Value node?
			// In debug output: Next.Next Value: "\"/opt/myapp/bin:$PATH\""
			// Yes, quotes are present.
			val = stripQuotes(val)
		}

		item := DocItem{
			Type:  typeStr,
			Name:  key,
			Value: val,
		}
		items = append(items, item)

		// Advance
		if curr.Next != nil {
			if curr.Next.Next != nil {
				// Check separator
				// sep := curr.Next.Next.Value
				// If sep is "" or "=", we skip it.
				curr = curr.Next.Next.Next
			} else {
				// End of chain (Key, Val, nil) - Wait, we saw Key, Val, "".
				// So if Next.Next is nil, maybe we stop?
				curr = nil
			}
		} else {
			curr = nil
		}
	}

	return items
}

func parseExpose(node *parser.Node) []DocItem {
	var items []DocItem
	curr := node.Next
	for curr != nil {
		item := DocItem{Type: "EXPOSE"}
		// EXPOSE 8080 or EXPOSE 80/tcp
		item.Name = curr.Value // Default name is the port
		item.Value = curr.Value
		items = append(items, item)
		curr = curr.Next
	}
	return items
}
