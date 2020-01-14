package htmlparse

import (
	"bytes"
	"io"

	"golang.org/x/net/html"
)

func GetAttribute(n *html.Node, key string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

// func checkId(n *html.Node, id string) bool {
// 	if n.Type == html.ElementNode {
// 		s, ok := GetAttribute(n, "id")
// 		if ok && s == id {
// 			return true
// 		}
// 	}
// 	return false
// }

// func checkName(n *html.Node, name string) bool {
// 	if n.Type == html.ElementNode {
// 		s, ok := GetAttribute(n, "name")
// 		if ok && s == name {
// 			return true
// 		}
// 	}
// 	return false
// }

func check(n *html.Node, query string, nodeType string) bool {
	if n.Type == html.ElementNode {
		s, ok := GetAttribute(n, nodeType)
		if ok && s == query {
			return true
		}
	}
	return false
}

func traverse(n *html.Node, query string, nodeType string, checkMethod func(*html.Node, string, string) bool) *html.Node {
	if checkMethod(n, query, nodeType) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result := traverse(c, query, nodeType, checkMethod)
		if result != nil {
			return result
		}
	}

	return nil
}

func GetElementById(n *html.Node, id string) *html.Node {
	return traverse(n, id, "id", check)
}

func GetElementByName(n *html.Node, name string) *html.Node {
	return traverse(n, name, "name", check)
}

func RenderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}
