package html

import (
	"strconv"
	"strings"
)

// Document returns an HTML document
func Document(lang string, head, body Element) Element { return HTML(Lang(lang), AddNode(head, body)) }

// Content (text) for an element
func Content(content string) Element { return Element{Content: content} }

// Anchor <a>: The Anchor element
func Anchor(href, content string, options ...func(*Element)) Element {
	return New(aElement, append(options, AddAttribute("href", href), SetContent(content))...)
}

// Abbreviation <abbr>: The Abbreviation element
func Abbreviation(content string, options ...func(*Element)) Element {
	return New(abbrElement, append(options, SetContent(content))...)
}

// Address <address>: The Address element
func Address(options ...func(*Element)) Element { return New(addressElement, options...) }

// Area <area>: The Area element
func Area(options ...func(*Element)) Element { return New(areaElement, options...) }

// Article <article>: The Article element
func Article(options ...func(*Element)) Element { return New(articleElement, options...) }

// Aside <aside>: The Aside element
func Aside(options ...func(*Element)) Element { return New(asideElement, options...) }

// Audio <audio>: The Audio element
func Audio(options ...func(*Element)) Element { return New(audioElement, options...) }

// Bold <b>: The Bring Attention To element
func Bold(content string, options ...func(*Element)) Element {
	return New(bElement, append(options, SetContent(content))...)
}

// Base <base>: The Document Base URL element
func Base(options ...func(*Element)) Element { return New(baseElement, options...) }

// BidirectionalIsolate <bdi>: The Bidirectional Isolate element
func BidirectionalIsolate(content string, options ...func(*Element)) Element {
	return New(bdiElement, append(options, SetContent(content))...)
}

// BidirectionalOverride <bdo>: The Bidirectional Override element
func BidirectionalOverride(content string, options ...func(*Element)) Element {
	return New(bdoElement, append(options, SetContent(content))...)
}

// Blockquote <blockquote>: The Block Quotation element
func Blockquote(content string, options ...func(*Element)) Element {
	return New(blockquoteElement, append(options, SetContent(content))...)
}

// Body <body>: The Document Body element
func Body(nodes ...Element) Element { return New(bodyElement, AddNode(nodes...)) }

// LineBreak <br>: The Line Break element
func LineBreak(options ...func(*Element)) Element { return New(brElement, options...) }

// Button <button type="button">: The Button element
func Button(options ...func(*Element)) Element {
	return New(buttonElement, append([]func(*Element){AddAttribute("type", "button")}, options...)...)
}

// Canvas <canvas>: The Graphics Canvas element
func Canvas(options ...func(*Element)) Element { return New(canvasElement, options...) }

// Caption <caption>: The Table Caption element
func Caption(options ...func(*Element)) Element { return New(captionElement, options...) }

// Citation <cite>: The Citation element
func Citation(options ...func(*Element)) Element { return New(citeElement, options...) }

// Code <code>: The Inline Code element
func Code(options ...func(*Element)) Element { return New(codeElement, options...) }

// Column <col>: The Table Column element
func Column(options ...func(*Element)) Element { return New(colElement, options...) }

// ColumnGroup <colgroup>: The Table Column Group element
func ColumnGroup(options ...func(*Element)) Element { return New(colgroupElement, options...) }

// Data <data>: The Data element
func Data(options ...func(*Element)) Element { return New(dataElement, options...) }

// DataList <datalist>: The HTML Data List element
func DataList(options ...func(*Element)) Element { return New(datalistElement, options...) }

// DescriptionDetails <dd>: The Description Details element
func DescriptionDetails(options ...func(*Element)) Element { return New(ddElement, options...) }

// DeletedText <del>: The Deleted Text element
func DeletedText(content string, options ...func(*Element)) Element {
	return New(delElement, append(options, SetContent(content))...)
}

// Details <details>: The HTML Details element
func Details(options ...func(*Element)) Element { return New(detailsElement, options...) }

// Definition <dfn>: The Definition element
func Definition(content string, options ...func(*Element)) Element {
	return New(dfnElement, append(options, SetContent(content))...)
}

// Dialog <dialog>: The HTML Dialog element
func Dialog(options ...func(*Element)) Element { return New(dialogElement, options...) }

// Div <div>: The Div element
func Div(options ...func(*Element)) Element { return New(divElement, options...) }

// DefinitionList <dl>: The Definition List element
func DefinitionList(options ...func(*Element)) Element { return New(dlElement, options...) }

// DefinitionTerm <dt>: The Definition Term element
func DefinitionTerm(options ...func(*Element)) Element { return New(dtElement, options...) }

// Emphasis <em>: The Emphasis element
func Emphasis(options ...func(*Element)) Element { return New(emElement, options...) }

// Embed <embed>: The Embed element
func Embed(options ...func(*Element)) Element { return New(embedElement, options...) }

// FieldSet <fieldset>: The Field Set element
func FieldSet(options ...func(*Element)) Element { return New(fieldsetElement, options...) }

// FigCaption <figcaption>: The Figure Caption element
func FigCaption(options ...func(*Element)) Element { return New(figcaptionElement, options...) }

// Figure <figure>: The Figure element
func Figure(options ...func(*Element)) Element { return New(figureElement, options...) }

// Footer <footer>: The Footer element
func Footer(options ...func(*Element)) Element { return New(footerElement, options...) }

// Form <form>: The Form element
func Form(options ...func(*Element)) Element { return New(formElement, options...) }

// H1 <h1>: The Heading element
func Heading1(content string, options ...func(*Element)) Element {
	return New(h1Element, append(options, SetContent(content))...)
}

// Heading2 <h2>: The Heading element
func Heading2(content string, options ...func(*Element)) Element {
	return New(h2Element, append(options, SetContent(content))...)
}

// Heading3 <h3>: The Heading element
func Heading3(content string, options ...func(*Element)) Element {
	return New(h3Element, append(options, SetContent(content))...)
}

// Heading4 <h4>: The Heading element
func Heading4(content string, options ...func(*Element)) Element {
	return New(h4Element, append(options, SetContent(content))...)
}

// Heading5 <h5>: The Heading element
func Heading5(content string, options ...func(*Element)) Element {
	return New(h5Element, append(options, SetContent(content))...)
}

// Heading6 <h6>: The Heading element
func Heading6(content string, options ...func(*Element)) Element {
	return New(h6Element, append(options, SetContent(content))...)
}

// Head <head>: The Head element
func Head(nodes ...Element) Element { return New(headElement, AddNode(nodes...)) }

// Header <header>: The Header element
func Header(options ...func(*Element)) Element { return New(headerElement, options...) }

// Heading Group <hgroup>: The Heading Group element
func HeadingGroup(options ...func(*Element)) Element { return New(hgroupElement, options...) }

// HorizontalRule <hr>: The Horizontal Rule element
func HorizontalRule(options ...func(*Element)) Element { return New(hrElement, options...) }

// HTML <html>: The HTML element
func HTML(options ...func(*Element)) Element { return New(htmlElement, options...) }

// Italic <i>: The Emphasis element
func Italic(content string, options ...func(*Element)) Element {
	return New(iElement, append(options, SetContent(content))...)
}

// Inline Frame <iframe>: The HTML Inline Frame element
func InlineFrame(options ...func(*Element)) Element { return New(iframeElement, options...) }

// Image <img>: The HTML Image element
func Image(src, alt string, options ...func(*Element)) Element {
	return New(imgElement, append(options, AddAttribute("src", src), AddAttribute("alt", alt))...)
}

// Inserted <ins>: The Inserted Text element
func Inserted(content string, options ...func(*Element)) Element {
	return New(insElement, append(options, SetContent(content))...)
}

// Keyboard <kbd>: The Keyboard Input element
func Keyboard(content string, options ...func(*Element)) Element {
	return New(kbdElement, append(options, SetContent(content))...)
}

// Label <label>: The Label element
func Label(options ...func(*Element)) Element { return New(labelElement, options...) }

// Legend <legend>: The Legend element
func Legend(options ...func(*Element)) Element { return New(legendElement, options...) }

// ListItem <li>: The List Item element
func ListItem(options ...func(*Element)) Element { return New(liElement, options...) }

// Link <link>: The HTML Link element
func Link(options ...func(*Element)) Element { return New(linkElement, options...) }

// LinkStylesheet <link rel="stylesheet" href="style.css">
func LinkStylesheet(href string, options ...func(*Element)) Element {
	return New(linkElement, append(options, AddAttribute("rel", "stylesheet"), AddAttribute("href", href))...)
}

// Main <main>: The Main element
func Main(options ...func(*Element)) Element { return New(mainElement, options...) }

// Map <map>: The HTML Map element
func Map(options ...func(*Element)) Element { return New(mapElement, options...) }

// Menu <menu>: The HTML Menu element
func Menu(options ...func(*Element)) Element { return New(menuElement, options...) }

// Meta <meta>: The HTML Meta element
func Meta(options ...func(*Element)) Element { return New(metaElement, options...) }

// MetaCharset <meta charset="UTF-8">
func MetaCharset(charset string, options ...func(*Element)) Element {
	return New(metaElement, append(options, AddAttribute("charset", charset))...)
}

// MetaViewport <meta name="viewport" content="width=device-width, initial-scale=1.0">
func MetaViewport(content string, options ...func(*Element)) Element {
	return New(metaElement, append(options, AddAttribute("name", "viewport"), AddAttribute("content", content))...)
}

// Nav <nav>: The HTML Navigation element
func Nav(options ...func(*Element)) Element { return New(navElement, options...) }

// NoScript <noscript>: The HTML noscript element
func NoScript(options ...func(*Element)) Element { return New(noscriptElement, options...) }

// Object <object>: The HTML object element
func Object(options ...func(*Element)) Element { return New(objectElement, options...) }

// OrderedList <ol>: The Ordered List element
func OrderedList(options ...func(*Element)) Element { return New(olElement, options...) }

// OptGroup <optgroup>: The Option Group element
func OptGroup(options ...func(*Element)) Element { return New(optgroupElement, options...) }

// Option <option>: The Option element
func Option(options ...func(*Element)) Element { return New(optionElement, options...) }

// Output <output>: The HTML Output element
func Output(options ...func(*Element)) Element { return New(outputElement, options...) }

// Paragraph <p>: The Paragraph element
func Paragraph(content string, options ...func(*Element)) Element {
	return New(pElement, append(options, SetContent(content))...)
}

// Picture <picture>: The HTML Picture element
func Picture(options ...func(*Element)) Element { return New(pictureElement, options...) }

// Preformatted <pre>: The Preformatted Text element
func Preformatted(options ...func(*Element)) Element { return New(preElement, options...) }

// Progress <progress>: The HTML Progress element
func Progress(options ...func(*Element)) Element { return New(progressElement, options...) }

// Quote <q>: The Quote element
func Quote(content string, options ...func(*Element)) Element {
	return New(qElement, append(options, SetContent(content))...)
}

// RubyParenthesis <rp>: The Ruby Parenthesis element
func RubyParenthesis(content string, options ...func(*Element)) Element {
	return New(rpElement, append(options, SetContent(content))...)
}

// RubyText <rt>: The Ruby Text element
func RubyText(content string, options ...func(*Element)) Element {
	return New(rtElement, append(options, SetContent(content))...)
}

// Ruby <ruby>: The Ruby element
func Ruby(content string, options ...func(*Element)) Element {
	return New(rubyElement, append(options, SetContent(content))...)
}

// Strikethrough <s>: The Strikethrough element
func Strikethrough(content string, options ...func(*Element)) Element {
	return New(sElement, append(options, SetContent(content))...)
}

// Sample <samp>: The Sample element
func Sample(options ...func(*Element)) Element { return New(sampElement, options...) }

// Script <script>: The HTML Script element
func Script(src string, options ...func(*Element)) Element {
	return New(scriptElement, append(options, AddAttribute("src", src))...)
}

// Search <search>: The HTML Search element
func Search(options ...func(*Element)) Element { return New(searchElement, options...) }

// Section <section>: The HTML Section element
func Section(options ...func(*Element)) Element { return New(sectionElement, options...) }

// Select <select>: The HTML Select element
func Select(options ...func(*Element)) Element { return New(selectElement, options...) }

// Slot <slot>: The HTML Slot element
func Slot(options ...func(*Element)) Element { return New(slotElement, options...) }

// Small <small>: The Small Text element
func Small(options ...func(*Element)) Element { return New(smallElement, options...) }

// Source <source>: The HTML Source element
func Source(options ...func(*Element)) Element { return New(sourceElement, options...) }

// Span <span>: The HTML Span element
func Span(options ...func(*Element)) Element { return New(spanElement, options...) }

// Strong <strong>: The Strong Text element
func Strong(options ...func(*Element)) Element { return New(strongElement, options...) }

// Style <style>: The HTML Style element
func Style(options ...func(*Element)) Element { return New(styleElement, options...) }

// Subscript <sub>: The Subscript element
func Subscript(options ...func(*Element)) Element { return New(subElement, options...) }

// Summary <summary>: The HTML Summary element
func Summary(options ...func(*Element)) Element { return New(summaryElement, options...) }

// Superscript <sup>: The Superscript element
func Superscript(options ...func(*Element)) Element { return New(supElement, options...) }

// Table <table>: The HTML Table element
func Table(options ...func(*Element)) Element { return New(tableElement, options...) }

// TableBody <tbody>: The HTML Table Body element
func TableBody(options ...func(*Element)) Element { return New(tbodyElement, options...) }

// TableData <td>: The HTML Table Data element
func TableData(options ...func(*Element)) Element { return New(tdElement, options...) }

// Template <template>: The HTML Template element
func Template(options ...func(*Element)) Element { return New(templateElement, options...) }

// TextArea <textarea>: The HTML Textarea element
func TextArea(options ...func(*Element)) Element { return New(textareaElement, options...) }

// TableFooter <tfoot>: The HTML Table Footer element
func TableFooter(options ...func(*Element)) Element { return New(tfootElement, options...) }

// TableHeader <th>: The HTML Table Header element
func TableHeader(options ...func(*Element)) Element { return New(thElement, options...) }

// TableHead <thead>: The HTML Table Header element
func TableHead(options ...func(*Element)) Element { return New(theadElement, options...) }

// Time <time>: The HTML Time element
func Time(options ...func(*Element)) Element { return New(timeElement, options...) }

// Title <title>: The HTML Title element
func Title(title string, options ...func(*Element)) Element {
	return New(titleElement, append(options, SetContent(title))...)
}

// TableRow <tr>: The HTML Table Row element
func TableRow(options ...func(*Element)) Element { return New(trElement, options...) }

// Track <track>: The HTML Track element
func Track(options ...func(*Element)) Element { return New(trackElement, options...) }

// Underline <u>: The Underline Text element
func Underline(content string, options ...func(*Element)) Element {
	return New(uElement, append(options, SetContent(content))...)
}

// UnorderedList <ul>: The HTML Unordered List element
func UnorderedList(options ...func(*Element)) Element { return New(ulElement, options...) }

// Variable <var>: The Variable Text element
func Variable(options ...func(*Element)) Element { return New(varElement, options...) }

// Video <video>: The HTML Video element
func Video(options ...func(*Element)) Element { return New(videoElement, options...) }

// WordBreak <wbr>: The Word Break Opportunity element
func WordBreak(options ...func(*Element)) Element { return New(wbrElement, options...) }

// InputButton <input type="button">
func InputButton(value string, options ...func(*Element)) Element {
	return New(inputElement, append(options, AddAttribute("type", "button"), AddAttribute("value", value))...)
}

// InputCheckbox <input type="checkbox">
func InputCheckbox(name string, options ...func(*Element)) Element {
	return New(inputElement, append(options, AddAttribute("type", "checkbox"), AddAttribute("name", name))...)
}

// InputColor <input type="color">
func InputColor(name string, options ...func(*Element)) Element {
	return New(inputElement, append(options, AddAttribute("type", "color"), AddAttribute("name", name))...)
}

// InputDate <input type="date">
func InputDate(name string, options ...func(*Element)) Element {
	return New(inputElement, append(options, AddAttribute("type", "date"), AddAttribute("name", name))...)
}

// InputDateTimeLocal <input type="datetime-local">
func InputDateTimeLocal(name string, options ...func(*Element)) Element {
	return New(inputElement, append(options, AddAttribute("type", "datetime-local"), AddAttribute("name", name))...)
}

// InputEmail <input type="email">
func InputEmail(name string, options ...func(*Element)) Element {
	return New(inputElement, append(options, AddAttribute("type", "email"), AddAttribute("name", name))...)
}

// InputFile <input type="file">
func InputFile(name string, options ...func(*Element)) Element {
	return New(inputElement, append(options, AddAttribute("type", "file"), AddAttribute("name", name))...)
}

// InputHidden <input type="hidden">
func InputHidden(name string, options ...func(*Element)) Element {
	return New(inputElement, append(options, AddAttribute("type", "hidden"), AddAttribute("name", name))...)
}

// InputImage <input type="image">
func InputImage(name string, options ...func(*Element)) Element {
	return New(inputElement, append(options, AddAttribute("type", "image"), AddAttribute("name", name))...)
}

// InputMonth <input type="month">
func InputMonth(name string, options ...func(*Element)) Element {
	return New(inputElement, append(options, AddAttribute("type", "month"), AddAttribute("name", name))...)
}

// InputNumber <input type="number">
func InputNumber(name string, options ...func(*Element)) Element {
	return New(inputElement, append(options, AddAttribute("type", "number"), AddAttribute("name", name))...)
}

// InputPassword <input type="password">
func InputPassword(name string, options ...func(*Element)) Element {
	return New(inputElement, append(options, AddAttribute("type", "password"), AddAttribute("name", name))...)
}

// InputRadio <input type="radio">
func InputRadio(name string, options ...func(*Element)) Element {
	return New(inputElement, append(options, AddAttribute("type", "radio"), AddAttribute("name", name))...)
}

// InputRange <input type="range">
func InputRange(name string, options ...func(*Element)) Element {
	return New(inputElement, append(options, AddAttribute("type", "range"), AddAttribute("name", name))...)
}

// InputReset <input type="reset">
func InputReset(name string, options ...func(*Element)) Element {
	return New(inputElement, append(options, AddAttribute("type", "reset"), AddAttribute("name", name))...)
}

// InputSearch <input type="search">
func InputSearch(name string, options ...func(*Element)) Element {
	return New(inputElement, append(options, AddAttribute("type", "search"), AddAttribute("name", name))...)
}

// InputSubmit <input type="submit">
func InputSubmit(value string, options ...func(*Element)) Element {
	return New(inputElement, append(options, AddAttribute("type", "submit"), AddAttribute("name", value))...)
}

// InputTel <input type="tel">
func InputTel(name string, options ...func(*Element)) Element {
	return New(inputElement, append(options, AddAttribute("type", "tel"), AddAttribute("name", name))...)
}

// InputText <input type="text">
func InputText(name string, options ...func(*Element)) Element {
	return New(inputElement, append(options, AddAttribute("type", "text"), AddAttribute("name", name))...)
}

// InputTime <input type="time">
func InputTime(name string, options ...func(*Element)) Element {
	return New(inputElement, append(options, AddAttribute("type", "time"), AddAttribute("name", name))...)
}

// InputURL <input type="url">
func InputURL(name string, options ...func(*Element)) Element {
	return New(inputElement, append(options, AddAttribute("type", "url"), AddAttribute("name", name))...)
}

// InputWeek <input type="week">
func InputWeek(name string, options ...func(*Element)) Element {
	return New(inputElement, append(options, AddAttribute("type", "week"), AddAttribute("name", name))...)
}

// AccessKey sets the accesskey attribute
func AccessKey(value string) func(*Element) { return AddAttribute(accesskeyAttribute, value) }

// AutoCapitalize sets the autocapitalize attribute
func AutoCapitalize(value string) func(*Element) { return AddAttribute(autocapitalizeAttribute, value) }

// Autofocus sets the autofocus attribute
func Autofocus() func(*Element) { return AddAttribute(autofocusAttribute, "") }

// Charset sets the charset attribute
func Charset(value string) func(*Element) { return AddAttribute(charsetAttribute, value) }

// Content sets the content attribute
func ContentAttr(value string) func(*Element) { return AddAttribute(contentAttribute, value) }

// Class sets the class attribute
func Class(className string) func(*Element) { return AddAttribute(classAttribute, className) }

// ContentEditable sets the contenteditable attribute
func ContentEditable(value string) func(*Element) {
	return AddAttribute(contenteditableAttribute, value)
}

// Dir sets the dir attribute
func Dir(value string) func(*Element) { return AddAttribute(dirAttribute, value) }

// Draggable sets the draggable attribute
func Draggable(value bool) func(*Element) {
	return AddAttribute(draggableAttribute, strconv.FormatBool(value))
}

// EnterKeyHint sets the enterkeyhint attribute
func EnterKeyHint(value string) func(*Element) { return AddAttribute(enterkeyhintAttribute, value) }

// exportparts sets the exportparts attribute
func ExportParts(values ...string) func(*Element) {
	return AddAttribute(exportpartsAttribute, strings.Join(values, ", "))
}

// Hidden sets the hidden attribute
func Hidden() func(*Element) { return AddAttribute(hiddenAttribute, "") }

// ID sets the id attribute
func ID(id string) func(*Element) { return AddAttribute(idAttribute, id) }

// Inert sets the inert attribute
func Inert() func(*Element) { return AddAttribute(inertAttribute, "") }

// InputMode sets the inputmode attribute
func InputMode(value string) func(*Element) { return AddAttribute(inputmodeAttribute, value) }

// Is sets the is attribute
func Is(value string) func(*Element) { return AddAttribute(isAttribute, value) }

// ItemID sets the itemid attribute
func ItemID(value string) func(*Element) { return AddAttribute(itemidAttribute, value) }

// ItemProp sets the itemprop attribute
func ItemProp(value string) func(*Element) { return AddAttribute(itempropAttribute, value) }

// ItemRef sets the itemref attribute
func ItemRef(value string) func(*Element) { return AddAttribute(itemrefAttribute, value) }

// ItemScope sets the itemscope attribute
func ItemScope() func(*Element) { return AddAttribute(itemscopeAttribute, "") }

// ItemType sets the itemtype attribute
func ItemType(value string) func(*Element) { return AddAttribute(itemtypeAttribute, value) }

// Lang sets the lang attribute
func Lang(value string) func(*Element) { return AddAttribute(langAttribute, value) }

// Name sets the name attribute
func Name(value string) func(*Element) { return AddAttribute("name", value) }

// Nonce sets the nonce attribute
func Nonce(value string) func(*Element) { return AddAttribute(nonceAttribute, value) }

// Part sets the part attribute
func Part(value string) func(*Element) { return AddAttribute(partAttribute, value) }

// Popover sets the popover attribute, values: "auto" (default) | "manual"
func Popover(value string) func(*Element) { return AddAttribute(popoverAttribute, value) }

// SlotAttr sets the slot attribute
func SlotAttr(value string) func(*Element) { return AddAttribute(slotAttribute, value) }

// SpellCheck sets the spellcheck attribute
func SpellCheck(value bool) func(*Element) {
	return AddAttribute(spellcheckAttribute, strconv.FormatBool(value))
}

// StyleAttr sets the style attribute
func StyleAttr(value string) func(*Element) { return AddAttribute(styleAttribute, value) }

// TabIndex sets the tabindex attribute
func TabIndex(value int) func(*Element) { return AddAttribute(tabindexAttribute, strconv.Itoa(value)) }

// TitleAttr sets the title attribute
func TitleAttr(value string) func(*Element) { return AddAttribute(titleAttribute, value) }

// Translate sets the translate attribute
func Translate(value bool) func(*Element) {
	return AddAttribute(translateAttribute, strconv.FormatBool(value))
}

// WriingSuggestions sets the writingsuggestion attribute
func WritingSuggestions(value bool) func(*Element) {
	return AddAttribute(writingsuggestionsAttribute, strconv.FormatBool(value))
}

// Accept sets the accept attribute
func Accept(value string) func(*Element) { return AddAttribute(acceptAttribute, value) }

// Autocomplete sets the autocomplete attribute
func Autocomplete(value string) func(*Element) { return AddAttribute(autocompleteAttribute, value) }

// Capture sets the capture attribute
func Capture(value string) func(*Element) { return AddAttribute(captureAttribute, value) }

// crossorigin sets the crossorigin attribute
func Crossorigin(value string) func(*Element) { return AddAttribute(crossoriginAttribute, value) }

// dirname sets the dirname attribute
func Dirname(value string) func(*Element) { return AddAttribute(dirnameAttribute, value) }

// disabled sets the disabled attribute
func Disabled() func(*Element) { return AddAttribute(disabledAttribute, "") }

// ElementTiming sets the elementtiming attribute
func ElementTiming(value string) func(*Element) { return AddAttribute(elementtimingAttribute, value) }

// For sets the for attribute
func For(value string) func(*Element) { return AddAttribute(forAttribute, value) }

// Max sets the max attribute
func Max(value string) func(*Element) { return AddAttribute(maxAttribute, value) }

// MaxLength sets the maxlength attribute
func Maxlength(value int) func(*Element) {
	return AddAttribute(maxlengthAttribute, strconv.Itoa(value))
}

// Min sets the min attribute
func Min(value string) func(*Element) { return AddAttribute(minAttribute, value) }

// MinLength sets the minlength attribute
func Minlength(value int) func(*Element) {
	return AddAttribute(minlengthAttribute, strconv.Itoa(value))
}

// Multiple sets the multiple attribute
func Multiple() func(*Element) { return AddAttribute(multipleAttribute, "") }

// Pattern sets the pattern attribute
func Pattern(value string) func(*Element) { return AddAttribute(patternAttribute, value) }

// Placeholder sets the placeholder attribute
func Placeholder(value string) func(*Element) { return AddAttribute(placeholderAttribute, value) }

// Readonly sets the readonly attribute
func Readonly() func(*Element) { return AddAttribute(readonlyAttribute, "") }

// Rel sets the rel attribute
func Rel(value string) func(*Element) { return AddAttribute(relAttribute, value) }

// Required sets the required attribute
func Required() func(*Element) { return AddAttribute(requiredAttribute, "") }

// Size sets the size attribute
func Size(value int) func(*Element) { return AddAttribute(sizeAttribute, strconv.Itoa(value)) }

// Step sets the step attribute
func Step(value int) func(*Element) { return AddAttribute(stepAttribute, strconv.Itoa(value)) }

func DataAttr(key, value string) func(*Element) { return AddAttribute("data-"+key, value) }

// RoleAlert sets role="alert"
func RoleAlert() func(*Element) { return AddAttribute("role", alertRole) }

// RoleAlertDialog sets role="alertdialog"
func RoleAlertDialog() func(*Element) { return AddAttribute("role", alertdialogRole) }

// RoleApplication sets role="application"
func RoleApplication() func(*Element) { return AddAttribute("role", applicationRole) }

// RoleArticle sets role="article"
func RoleArticle() func(*Element) { return AddAttribute("role", articleRole) }

// RoleBanner sets role="banner"
func RoleBanner() func(*Element) { return AddAttribute("role", bannerRole) }

// RoleButton sets role="button"
func RoleButton() func(*Element) { return AddAttribute("role", buttonRole) }

// RoleCell sets role="cell"
func RoleCell() func(*Element) { return AddAttribute("role", cellRole) }

// RoleCheckbox sets role="checkbox"
func RoleCheckbox() func(*Element) { return AddAttribute("role", checkboxRole) }

// RoleColumnheader sets role="columnheader"
func RoleColumnheader() func(*Element) { return AddAttribute("role", columnheaderRole) }

// RoleCombobox sets role="combobox"
func RoleCombobox() func(*Element) { return AddAttribute("role", comboboxRole) }

// RoleCommand sets role="command"
func RoleCommand() func(*Element) { return AddAttribute("role", commandRole) }

// RoleComment sets role="comment"
func RoleComment() func(*Element) { return AddAttribute("role", commentRole) }

// RoleComplementary sets role="complementary"
func RoleComplementary() func(*Element) { return AddAttribute("role", complementaryRole) }

// RoleContentinfo sets role="contentinfo"
func RoleContentinfo() func(*Element) { return AddAttribute("role", contentinfoRole) }

// RoleDefinition sets role="definition"
func RoleDefinition() func(*Element) { return AddAttribute("role", definitionRole) }

// RoleDialog sets role="dialog"
func RoleDialog() func(*Element) { return AddAttribute("role", dialogRole) }

// RoleDocument sets role="document"
func RoleDocument() func(*Element) { return AddAttribute("role", documentRole) }

// RoleFeed sets role="feed"
func RoleFeed() func(*Element) { return AddAttribute("role", feedRole) }

// RoleFigure sets role="figure"
func RoleFigure() func(*Element) { return AddAttribute("role", figureRole) }

// RoleForm sets role="form"
func RoleForm() func(*Element) { return AddAttribute("role", formRole) }

// RoleGrid sets role="grid"
func RoleGrid() func(*Element) { return AddAttribute("role", gridRole) }

// RoleGridCell sets role="gridcell"
func RoleGridCell() func(*Element) { return AddAttribute("role", gridcellRole) }

// RoleGroup sets role="group"
func RoleGroup() func(*Element) { return AddAttribute("role", groupRole) }

// RoleHeading sets role="heading"
func RoleHeading() func(*Element) { return AddAttribute("role", headingRole) }

// RoleImg sets role="img"
func RoleImg() func(*Element) { return AddAttribute("role", imgRole) }

// RoleInput sets role="input"
func RoleInput() func(*Element) { return AddAttribute("role", inputRole) }

// RoleLandmark sets role="landmark"
func RoleLandmark() func(*Element) { return AddAttribute("role", landmarkRole) }

// RoleLink sets role="link"
func RoleLink() func(*Element) { return AddAttribute("role", linkRole) }

// RoleList sets role="list"
func RoleList() func(*Element) { return AddAttribute("role", listRole) }

// RoleListbox sets role="listbox"
func RoleListbox() func(*Element) { return AddAttribute("role", listboxRole) }

// RoleListitem sets role="listitem"
func RoleListitem() func(*Element) { return AddAttribute("role", listitemRole) }

// RoleLog sets role="log"
func RoleLog() func(*Element) { return AddAttribute("role", logRole) }

// RoleMain sets role="main"
func RoleMain() func(*Element) { return AddAttribute("role", mainRole) }

// RoleMarquee sets role="marquee"
func RoleMarquee() func(*Element) { return AddAttribute("role", marqueeRole) }

// RoleMath sets role="math"
func RoleMath() func(*Element) { return AddAttribute("role", mathRole) }

// RoleMenu sets role="menu"
func RoleMenu() func(*Element) { return AddAttribute("role", menuRole) }

// RoleMenubar sets role="menubar"
func RoleMenubar() func(*Element) { return AddAttribute("role", menubarRole) }

// RoleNavigation sets role="navigation"
func RoleNavigation() func(*Element) { return AddAttribute("role", navigationRole) }

// RoleNone sets role="none"
func RoleNone() func(*Element) { return AddAttribute("role", noneRole) }

// RoleNote sets role="note"
func RoleNote() func(*Element) { return AddAttribute("role", noteRole) }

// RoleOption sets role="option"
func RoleOption() func(*Element) { return AddAttribute("role", optionRole) }

// RolePresentation sets role="presentation"
func RolePresentation() func(*Element) { return AddAttribute("role", presentationRole) }

// RoleProgressBar sets role="progressbar"
func RoleProgressBar() func(*Element) { return AddAttribute("role", progressbarRole) }

// RoleRadio sets role="radio"
func RoleRadio() func(*Element) { return AddAttribute("role", radioRole) }

// RoleRadiogroup sets role="radiogroup"
func RoleRadiogroup() func(*Element) { return AddAttribute("role", radiogroupRole) }

// RoleRegion sets role="region"
func RoleRegion() func(*Element) { return AddAttribute("role", regionRole) }

// RoleRow sets role="row"
func RoleRow() func(*Element) { return AddAttribute("role", rowRole) }

// RoleRowgroup sets role="rowgroup"
func RoleRowgroup() func(*Element) { return AddAttribute("role", rowgroupRole) }

// RoleRowheader sets role="rowheader"
func RoleRowheader() func(*Element) { return AddAttribute("role", rowheaderRole) }

// RoleScrollbar sets role="scrollbar"
func RoleScrollbar() func(*Element) { return AddAttribute("role", scrollbarRole) }

// RoleSearch sets role="search"
func RoleSearch() func(*Element) { return AddAttribute("role", searchRole) }

// RoleSearchbox sets role="searchbox"
func RoleSearchbox() func(*Element) { return AddAttribute("role", searchboxRole) }

// RoleSeparator sets role="separator"
func RoleSeparator() func(*Element) { return AddAttribute("role", separatorRole) }

// RoleSlider sets role="slider"
func RoleSlider() func(*Element) { return AddAttribute("role", sliderRole) }

// RoleSpinbutton sets role="spinbutton"
func RoleSpinbutton() func(*Element) { return AddAttribute("role", spinbuttonRole) }

// RoleStatus sets role="status"
func RoleStatus() func(*Element) { return AddAttribute("role", statusRole) }

// RoleSwitch sets role="switch"
func RoleSwitch() func(*Element) { return AddAttribute("role", switchRole) }

// RoleTab sets role="tab"
func RoleTab() func(*Element) { return AddAttribute("role", tabRole) }

// RoleTable sets role="table"
func RoleTable() func(*Element) { return AddAttribute("role", tableRole) }

// RoleTablist sets role="tablist"
func RoleTablist() func(*Element) { return AddAttribute("role", tablistRole) }

// RoleTerm sets role="term"
func RoleTerm() func(*Element) { return AddAttribute("role", termRole) }

// RoleTextbox sets role="textbox"
func RoleTextbox() func(*Element) { return AddAttribute("role", textboxRole) }

// RoleTimer sets role="timer"
func RoleTimer() func(*Element) { return AddAttribute("role", timerRole) }

// RoleToolbar sets role="toolbar"
func RoleToolbar() func(*Element) { return AddAttribute("role", toolbarRole) }

// RoleTooltip sets role="tooltip"
func RoleTooltip() func(*Element) { return AddAttribute("role", tooltipRole) }

// RoleTree sets role="tree"
func RoleTree() func(*Element) { return AddAttribute("role", treeRole) }

// RoleTreegrid sets role="treegrid"
func RoleTreegrid() func(*Element) { return AddAttribute("role", treegridRole) }

// RoleTreeitem sets role="treeitem"
func RoleTreeitem() func(*Element) { return AddAttribute("role", treeitemRole) }

// RoleWidget sets role="widget"
func RoleWidget() func(*Element) { return AddAttribute("role", widgetRole) }

// RoleWindow sets role="window"
func RoleWindow() func(*Element) { return AddAttribute("role", windowRole) }
