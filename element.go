package html

import (
	"io"
	"strings"
)

// Element is an HTML element
type Element struct {
	Tag        string
	Content    string
	Attributes Attributes
	Nodes      []Element
}

// Render renders the element
func (e Element) Render(w io.Writer) {
	ElementWriter(w, e)
}

// String returns the string representation of the element
func (e Element) String() string {
	b := strings.Builder{}
	ElementWriter(&b, e)
	return b.String()
}

// AddAttribute sets an attribute, overwriting any existing value with the same key
func AddAttribute(key, value string) func(*Element) {
	return func(e *Element) {
		e.Attributes[key] = value
	}
}

// AddAttributes sets multiple attributes, overwriting any existing values with the same key
func AddAttributes(attributes map[string]string) func(*Element) {
	return func(e *Element) {
		e.Attributes.Merge(attributes)
	}
}

// SetNode sets the nodes of the element, overwriting all existing nodes
func SetNode(nodes ...Element) func(*Element) {
	return func(e *Element) {
		e.Nodes = nodes
	}
}

// AddNode adds nodes to the element
func AddNode(nodes ...Element) func(*Element) {
	return func(e *Element) {
		e.Nodes = append(e.Nodes, nodes...)
	}
}

// SetContent sets the content of the element
func SetContent(content string) func(*Element) {
	return func(e *Element) {
		e.Content = content
	}
}

func New(tag string, options ...func(*Element)) Element {
	e := Element{
		Tag:        tag,
		Attributes: make(map[string]string),
	}

	for _, option := range options {
		option(&e)
	}

	return e
}

var ElementNames = []string{
	aElement,
	abbrElement,
	addressElement,
	areaElement,
	articleElement,
	asideElement,
	audioElement,
	bElement,
	baseElement,
	bdiElement,
	bdoElement,
	blockquoteElement,
	bodyElement,
	brElement,
	buttonElement,
	canvasElement,
	captionElement,
	citeElement,
	codeElement,
	colElement,
	colgroupElement,
	dataElement,
	datalistElement,
	ddElement,
	delElement,
	detailsElement,
	dfnElement,
	dialogElement,
	divElement,
	dlElement,
	dtElement,
	emElement,
	embedElement,
	fieldsetElement,
	figcaptionElement,
	figureElement,
	footerElement,
	formElement,
	h1Element,
	h2Element,
	h3Element,
	h4Element,
	h5Element,
	h6Element,
	headElement,
	headerElement,
	hgroupElement,
	hrElement,
	htmlElement,
	iElement,
	iframeElement,
	imgElement,
	inputElement,
	insElement,
	kbdElement,
	labelElement,
	legendElement,
	liElement,
	linkElement,
	mainElement,
	mapElement,
	markElement,
	menuElement,
	metaElement,
	meterElement,
	navElement,
	noscriptElement,
	objectElement,
	olElement,
	optgroupElement,
	optionElement,
	outputElement,
	pElement,
	pictureElement,
	preElement,
	progressElement,
	qElement,
	rpElement,
	rtElement,
	rubyElement,
	sElement,
	sampElement,
	scriptElement,
	searchElement,
	sectionElement,
	selectElement,
	slotElement,
	smallElement,
	sourceElement,
	spanElement,
	strongElement,
	styleElement,
	subElement,
	summaryElement,
	supElement,
	tableElement,
	tbodyElement,
	tdElement,
	templateElement,
	textareaElement,
	tfootElement,
	thElement,
	theadElement,
	timeElement,
	titleElement,
	trElement,
	trackElement,
	uElement,
	ulElement,
	varElement,
	videoElement,
	wbrElement,
	svgElement,
	pathElement,
}

const (
	aElement          = "a"
	abbrElement       = "abbr"
	addressElement    = "address"
	areaElement       = "area"
	articleElement    = "article"
	asideElement      = "aside"
	audioElement      = "audio"
	bElement          = "b"
	baseElement       = "base"
	bdiElement        = "bdi"
	bdoElement        = "bdo"
	blockquoteElement = "blockquote"
	bodyElement       = "body"
	brElement         = "br"
	buttonElement     = "button"
	canvasElement     = "canvas"
	captionElement    = "caption"
	citeElement       = "cite"
	codeElement       = "code"
	colElement        = "col"
	colgroupElement   = "colgroup"
	dataElement       = "data"
	datalistElement   = "datalist"
	ddElement         = "dd"
	delElement        = "del"
	detailsElement    = "details"
	dfnElement        = "dfn"
	dialogElement     = "dialog"
	divElement        = "div"
	dlElement         = "dl"
	dtElement         = "dt"
	emElement         = "em"
	embedElement      = "embed"
	fieldsetElement   = "fieldset"
	figcaptionElement = "figcaption"
	figureElement     = "figure"
	footerElement     = "footer"
	formElement       = "form"
	h1Element         = "h1"
	h2Element         = "h2"
	h3Element         = "h3"
	h4Element         = "h4"
	h5Element         = "h5"
	h6Element         = "h6"
	headElement       = "head"
	headerElement     = "header"
	hgroupElement     = "hgroup"
	hrElement         = "hr"
	htmlElement       = "html"
	iElement          = "i"
	iframeElement     = "iframe"
	imgElement        = "img"
	inputElement      = "input"
	insElement        = "ins"
	kbdElement        = "kbd"
	labelElement      = "label"
	legendElement     = "legend"
	liElement         = "li"
	linkElement       = "link"
	mainElement       = "main"
	mapElement        = "map"
	markElement       = "mark"
	menuElement       = "menu"
	metaElement       = "meta"
	meterElement      = "meter"
	navElement        = "nav"
	noscriptElement   = "noscript"
	objectElement     = "object"
	olElement         = "ol"
	optgroupElement   = "optgroup"
	optionElement     = "option"
	outputElement     = "output"
	pElement          = "p"
	pictureElement    = "picture"
	preElement        = "pre"
	progressElement   = "progress"
	qElement          = "q"
	rpElement         = "rp"
	rtElement         = "rt"
	rubyElement       = "ruby"
	sElement          = "s"
	sampElement       = "samp"
	scriptElement     = "script"
	searchElement     = "search"
	sectionElement    = "section"
	selectElement     = "select"
	slotElement       = "slot"
	smallElement      = "small"
	sourceElement     = "source"
	spanElement       = "span"
	strongElement     = "strong"
	styleElement      = "style"
	subElement        = "sub"
	summaryElement    = "summary"
	supElement        = "sup"
	tableElement      = "table"
	tbodyElement      = "tbody"
	tdElement         = "td"
	templateElement   = "template"
	textareaElement   = "textarea"
	tfootElement      = "tfoot"
	thElement         = "th"
	theadElement      = "thead"
	timeElement       = "time"
	titleElement      = "title"
	trElement         = "tr"
	trackElement      = "track"
	uElement          = "u"
	ulElement         = "ul"
	varElement        = "var"
	videoElement      = "video"
	wbrElement        = "wbr"
	svgElement        = "svg"
	pathElement       = "path"
)
