package main

import (
	"errors"

	"github.com/jaytaylor/html2text"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

type Renderer interface {
	render(doc string) (string, error)
}

type textRenderer struct {
}

func (r *textRenderer) render(doc string) (string, error) {
	return doc, nil
}

func newTextRenderer() Renderer {
	return &textRenderer{}
}

type htmlRenderer struct {
}

func (r *htmlRenderer) render(doc string) (string, error) {
	html := bluemonday.UGCPolicy().Sanitize(doc)
	return html2text.FromString(html)
}

func newHtmlRenderer() Renderer {
	return &htmlRenderer{}
}

type mdRenderer struct {
	htmlRenderer Renderer
}

func (r *mdRenderer) render(doc string) (string, error) {
	htmlBytes := blackfriday.MarkdownCommon([]byte(doc))
	return r.htmlRenderer.render(string(htmlBytes[:]))
}

func newMdRenderer(htmlRenderer Renderer) Renderer {
	return &mdRenderer{htmlRenderer}
}

var tr Renderer = newTextRenderer()
var hr Renderer = newHtmlRenderer()
var mr Renderer = newMdRenderer(hr)

func NewRenderer(docType DocType) (Renderer, error) {
	if docType == TEXT {
		return tr, nil
	} else if docType == HTML {
		return hr, nil
	} else if docType == MARKDOWN {
		return mr, nil
	}
	return nil, errors.New("Invalid document type")
}
