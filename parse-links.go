package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Link represents a link in an HTML document.
type Link struct {
	Href string
	Text string
}

// Parse will take an HTML document and return a slice of Links
// extracted from it.
func Parse(r io.Reader) ([]Link, error) {
	htmlTree, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	linkNodes := findLinkNodes(htmlTree)

	var links []Link
	for _, l := range linkNodes {
		links = append(links, extractLink(l))
	}

	return links, nil
}

func findLinkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.DataAtom == atom.A {
		return []*html.Node{n}
	}

	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, findLinkNodes(c)...)
	}
	return ret
}

func extractLink(n *html.Node) Link {
	// the node passed in here should be an <a href=...> tag, we verify this now
	if n.DataAtom != atom.A {
		panic("invalid argument")
	}

	var ret Link

	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}

	ret.Text = extractText(n)

	return ret
}

func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}

	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += extractText(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}
