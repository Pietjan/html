package html

import (
	"fmt"
	"io"
	"slices"
)

var singleLineElements = []string{
	metaElement, linkElement, titleElement, spanElement, h1Element, h2Element, h3Element, h4Element, h5Element, h6Element, tdElement, thElement, preElement,
}

var inlineElements = []string{
	aElement, abbrElement, bElement, bdiElement, bdoElement, brElement, citeElement, codeElement, dataElement,
	delElement, dfnElement, emElement, iElement, insElement, kbdElement, markElement, qElement, rpElement,
	rtElement, rubyElement, sElement, sampElement, smallElement, spanElement, strongElement,
	subElement, supElement, timeElement, uElement, varElement, wbrElement,
}

func Render(w io.Writer, elements ...Element) {
	for _, element := range elements {
		element.Render(w)
	}
}

func ElementWriter(w io.Writer, element Element) {
	if element.Tag == "" && len(element.Nodes) == 0 && element.Content == "" {
		return
	}

	if element.Tag == htmlElement {
		fmt.Fprintf(w, "<!DOCTYPE html>\n")
	}

	if element.Tag != "" {
		fmt.Fprintf(w, "<%s", element.Tag)
		AttributesWriter(w, element.Attributes)
		fmt.Fprintf(w, ">")
	}

	if slices.Contains([]string{metaElement, linkElement}, element.Tag) {
		fmt.Fprintf(w, "\n")
		return
	}

	singleLine := slices.Contains(singleLineElements, element.Tag) || len(element.Nodes) == 0
	if len(element.Nodes) == 0 && element.Content == "" || singleLine {
		fmt.Fprintf(w, "%s</%s>", element.Content, element.Tag)
		if !slices.Contains(inlineElements, element.Tag) {
			fmt.Fprintf(w, "\n")
		}
		return
	}

	fmt.Fprintf(w, "\n")

	if element.Content != "" {
		fmt.Fprintf(IndentWriter(w, Space(2)), "%s\n", element.Content)
	}

	for _, node := range element.Nodes {
		ElementWriter(IndentWriter(w, Space(2)), node)
	}

	if element.Tag != "" {
		fmt.Fprintf(w, "</%s>\n", element.Tag)
	}
}

func AttributesWriter(w io.Writer, attributes map[string]string) {
	var keys []string
	for k := range attributes {
		keys = append(keys, k)
	}

	slices.Sort(keys)

	for _, key := range keys {
		fmt.Fprintf(w, ` %s`, key)
		if val := attributes[key]; val != "" {
			fmt.Fprintf(w, `="%s"`, val)
		}
	}
}
