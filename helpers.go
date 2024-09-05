package html

import (
	"fmt"
	"strconv"
	"strings"
)

// Document returns an HTML document
func Document(lang string, head, body Element) Element { return HTML(Lang(lang), AddNode(head, body)) }

// Content (text) for an element
func Content(content string) Element { return Element{Content: content} }

// Anchor <a>: The Anchor element
func Anchor(href, content string, options ...func(*Element)) Element {
	return New(aElement, append(options, Attribute("href", href), SetContent(content))...)
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
	return New(buttonElement, append([]func(*Element){Attribute("type", "button")}, options...)...)
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
	return New(imgElement, append(options, Attribute("src", src), Attribute("alt", alt))...)
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
func LinkStyleSheet(href string, options ...func(*Element)) Element {
	return New(linkElement, append(options, Attribute("rel", "stylesheet"), Attribute("href", href))...)
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
	return New(metaElement, append(options, Attribute("charset", charset))...)
}

// MetaViewport <meta name="viewport" content="width=device-width, initial-scale=1.0">
func MetaViewport(content string, options ...func(*Element)) Element {
	return New(metaElement, append(options, Attribute("name", "viewport"), Attribute("content", content))...)
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
	return New(scriptElement, append(options, Attribute("src", src))...)
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
func InputButton(options ...func(*Element)) Element {
	return New(inputElement, append(options, Attribute("type", "button"))...)
}

// InputCheckbox <input type="checkbox">
func InputCheckbox(options ...func(*Element)) Element {
	return New(inputElement, append(options, Attribute("type", "checkbox"))...)
}

// InputColor <input type="color">
func InputColor(options ...func(*Element)) Element {
	return New(inputElement, append(options, Attribute("type", "color"))...)
}

// InputDate <input type="date">
func InputDate(options ...func(*Element)) Element {
	return New(inputElement, append(options, Attribute("type", "date"))...)
}

// InputDateTimeLocal <input type="datetime-local">
func InputDateTimeLocal(options ...func(*Element)) Element {
	return New(inputElement, append(options, Attribute("type", "datetime-local"))...)
}

// InputEmail <input type="email">
func InputEmail(options ...func(*Element)) Element {
	return New(inputElement, append(options, Attribute("type", "email"))...)
}

// InputFile <input type="file">
func InputFile(options ...func(*Element)) Element {
	return New(inputElement, append(options, Attribute("type", "file"))...)
}

// InputHidden <input type="hidden">
func InputHidden(options ...func(*Element)) Element {
	return New(inputElement, append(options, Attribute("type", "hidden"))...)
}

// InputImage <input type="image">
func InputImage(options ...func(*Element)) Element {
	return New(inputElement, append(options, Attribute("type", "image"))...)
}

// InputMonth <input type="month">
func InputMonth(options ...func(*Element)) Element {
	return New(inputElement, append(options, Attribute("type", "month"))...)
}

// InputNumber <input type="number">
func InputNumber(options ...func(*Element)) Element {
	return New(inputElement, append(options, Attribute("type", "number"))...)
}

// InputPassword <input type="password">
func InputPassword(options ...func(*Element)) Element {
	return New(inputElement, append(options, Attribute("type", "password"))...)
}

// InputRadio <input type="radio">
func InputRadio(options ...func(*Element)) Element {
	return New(inputElement, append(options, Attribute("type", "radio"))...)
}

// InputRange <input type="range">
func InputRange(options ...func(*Element)) Element {
	return New(inputElement, append(options, Attribute("type", "range"))...)
}

// InputReset <input type="reset">
func InputReset(options ...func(*Element)) Element {
	return New(inputElement, append(options, Attribute("type", "reset"))...)
}

// InputSearch <input type="search">
func InputSearch(options ...func(*Element)) Element {
	return New(inputElement, append(options, Attribute("type", "search"))...)
}

// InputSubmit <input type="submit">
func InputSubmit(options ...func(*Element)) Element {
	return New(inputElement, append(options, Attribute("type", "submit"))...)
}

// InputTel <input type="tel">
func InputTel(options ...func(*Element)) Element {
	return New(inputElement, append(options, Attribute("type", "tel"))...)
}

// InputText <input type="text">
func InputText(options ...func(*Element)) Element {
	return New(inputElement, append(options, Attribute("type", "text"))...)
}

// InputTime <input type="time">
func InputTime(options ...func(*Element)) Element {
	return New(inputElement, append(options, Attribute("type", "time"))...)
}

// InputURL <input type="url">
func InputURL(options ...func(*Element)) Element {
	return New(inputElement, append(options, Attribute("type", "url"))...)
}

// InputWeek <input type="week">
func InputWeek(options ...func(*Element)) Element {
	return New(inputElement, append(options, Attribute("type", "week"))...)
}

// Value sets the value attribute
func Value(value string) func(*Element) { return Attribute("value", value) }

// AccessKey sets the accesskey attribute
func AccessKey(value string) func(*Element) { return Attribute(accesskeyAttribute, value) }

// AutoCapitalize sets the autocapitalize attribute
func AutoCapitalize(value string) func(*Element) { return Attribute(autocapitalizeAttribute, value) }

// Autofocus sets the autofocus attribute
func Autofocus() func(*Element) { return Attribute(autofocusAttribute, "") }

// Charset sets the charset attribute
func Charset(value string) func(*Element) { return Attribute(charsetAttribute, value) }

// Content sets the content attribute
func ContentAttr(value string) func(*Element) { return Attribute(contentAttribute, value) }

// Class sets the class attribute
func Class(className string) func(*Element) { return Attribute(classAttribute, className) }

// ContentEditable sets the contenteditable attribute
func ContentEditable(value string) func(*Element) {
	return Attribute(contenteditableAttribute, value)
}

// Dir sets the dir attribute
func Dir(value string) func(*Element) { return Attribute(dirAttribute, value) }

// Draggable sets the draggable attribute
func Draggable(value bool) func(*Element) {
	return Attribute(draggableAttribute, strconv.FormatBool(value))
}

// EnterKeyHint sets the enterkeyhint attribute
func EnterKeyHint(value string) func(*Element) { return Attribute(enterkeyhintAttribute, value) }

// exportparts sets the exportparts attribute
func ExportParts(values ...string) func(*Element) {
	return Attribute(exportpartsAttribute, strings.Join(values, ", "))
}

// Hidden sets the hidden attribute
func Hidden() func(*Element) { return Attribute(hiddenAttribute, "") }

// ID sets the id attribute
func ID(id string) func(*Element) { return Attribute(idAttribute, id) }

// Inert sets the inert attribute
func Inert() func(*Element) { return Attribute(inertAttribute, "") }

// InputMode sets the inputmode attribute
func InputMode(value string) func(*Element) { return Attribute(inputmodeAttribute, value) }

// Is sets the is attribute
func Is(value string) func(*Element) { return Attribute(isAttribute, value) }

// ItemID sets the itemid attribute
func ItemID(value string) func(*Element) { return Attribute(itemidAttribute, value) }

// ItemProp sets the itemprop attribute
func ItemProp(value string) func(*Element) { return Attribute(itempropAttribute, value) }

// ItemRef sets the itemref attribute
func ItemRef(value string) func(*Element) { return Attribute(itemrefAttribute, value) }

// ItemScope sets the itemscope attribute
func ItemScope() func(*Element) { return Attribute(itemscopeAttribute, "") }

// ItemType sets the itemtype attribute
func ItemType(value string) func(*Element) { return Attribute(itemtypeAttribute, value) }

// Lang sets the lang attribute
func Lang(value string) func(*Element) { return Attribute(langAttribute, value) }

// Name sets the name attribute
func Name(value string) func(*Element) { return Attribute("name", value) }

// Nonce sets the nonce attribute
func Nonce(value string) func(*Element) { return Attribute(nonceAttribute, value) }

// Part sets the part attribute
func Part(value string) func(*Element) { return Attribute(partAttribute, value) }

// Popover sets the popover attribute, values: "auto" (default) | "manual"
func Popover(value string) func(*Element) { return Attribute(popoverAttribute, value) }

// SlotAttr sets the slot attribute
func SlotAttr(value string) func(*Element) { return Attribute(slotAttribute, value) }

// SpellCheck sets the spellcheck attribute
func SpellCheck(value bool) func(*Element) {
	return Attribute(spellcheckAttribute, strconv.FormatBool(value))
}

// StyleAttr sets the style attribute
func StyleAttr(value string) func(*Element) { return Attribute(styleAttribute, value) }

// TabIndex sets the tabindex attribute
func TabIndex(value int) func(*Element) { return Attribute(tabindexAttribute, strconv.Itoa(value)) }

// TitleAttr sets the title attribute
func TitleAttr(value string) func(*Element) { return Attribute(titleAttribute, value) }

// Translate sets the translate attribute
func Translate(value bool) func(*Element) {
	return Attribute(translateAttribute, strconv.FormatBool(value))
}

// WriingSuggestions sets the writingsuggestion attribute
func WritingSuggestions(value bool) func(*Element) {
	return Attribute(writingsuggestionsAttribute, strconv.FormatBool(value))
}

// Accept sets the accept attribute
func Accept(value string) func(*Element) { return Attribute(acceptAttribute, value) }

// Autocomplete sets the autocomplete attribute
func Autocomplete(value string) func(*Element) { return Attribute(autocompleteAttribute, value) }

// Capture sets the capture attribute
func Capture(value string) func(*Element) { return Attribute(captureAttribute, value) }

// crossorigin sets the crossorigin attribute
func Crossorigin(value string) func(*Element) { return Attribute(crossoriginAttribute, value) }

// dirname sets the dirname attribute
func Dirname(value string) func(*Element) { return Attribute(dirnameAttribute, value) }

// disabled sets the disabled attribute
func Disabled() func(*Element) { return Attribute(disabledAttribute, "") }

// ElementTiming sets the elementtiming attribute
func ElementTiming(value string) func(*Element) { return Attribute(elementtimingAttribute, value) }

// For sets the for attribute
func For(value string) func(*Element) { return Attribute(forAttribute, value) }

// Max sets the max attribute
func Max(value string) func(*Element) { return Attribute(maxAttribute, value) }

// MaxLength sets the maxlength attribute
func Maxlength(value int) func(*Element) {
	return Attribute(maxlengthAttribute, strconv.Itoa(value))
}

// Min sets the min attribute
func Min(value string) func(*Element) { return Attribute(minAttribute, value) }

// MinLength sets the minlength attribute
func Minlength(value int) func(*Element) {
	return Attribute(minlengthAttribute, strconv.Itoa(value))
}

// Multiple sets the multiple attribute
func Multiple() func(*Element) { return Attribute(multipleAttribute, "") }

// Pattern sets the pattern attribute
func Pattern(value string) func(*Element) { return Attribute(patternAttribute, value) }

// Placeholder sets the placeholder attribute
func Placeholder(value string) func(*Element) { return Attribute(placeholderAttribute, value) }

// Readonly sets the readonly attribute
func Readonly() func(*Element) { return Attribute(readonlyAttribute, "") }

// Rel sets the rel attribute
func Rel(value string) func(*Element) { return Attribute(relAttribute, value) }

// Required sets the required attribute
func Required() func(*Element) { return Attribute(requiredAttribute, "") }

// Size sets the size attribute
func Size(value int) func(*Element) { return Attribute(sizeAttribute, strconv.Itoa(value)) }

// Step sets the step attribute
func Step(value int) func(*Element) { return Attribute(stepAttribute, strconv.Itoa(value)) }

func DataAttr(key, value string) func(*Element) { return Attribute("data-"+key, value) }

func Method(value string) func(*Element) { return Attribute("method", value) }

func Action(value string) func(*Element) { return Attribute("action", value) }

// RoleAlert sets role="alert"
func RoleAlert() func(*Element) { return Attribute("role", alertRole) }

// RoleAlertDialog sets role="alertdialog"
func RoleAlertDialog() func(*Element) { return Attribute("role", alertdialogRole) }

// RoleApplication sets role="application"
func RoleApplication() func(*Element) { return Attribute("role", applicationRole) }

// RoleArticle sets role="article"
func RoleArticle() func(*Element) { return Attribute("role", articleRole) }

// RoleBanner sets role="banner"
func RoleBanner() func(*Element) { return Attribute("role", bannerRole) }

// RoleButton sets role="button"
func RoleButton() func(*Element) { return Attribute("role", buttonRole) }

// RoleCell sets role="cell"
func RoleCell() func(*Element) { return Attribute("role", cellRole) }

// RoleCheckbox sets role="checkbox"
func RoleCheckbox() func(*Element) { return Attribute("role", checkboxRole) }

// RoleColumnheader sets role="columnheader"
func RoleColumnheader() func(*Element) { return Attribute("role", columnheaderRole) }

// RoleCombobox sets role="combobox"
func RoleCombobox() func(*Element) { return Attribute("role", comboboxRole) }

// RoleCommand sets role="command"
func RoleCommand() func(*Element) { return Attribute("role", commandRole) }

// RoleComment sets role="comment"
func RoleComment() func(*Element) { return Attribute("role", commentRole) }

// RoleComplementary sets role="complementary"
func RoleComplementary() func(*Element) { return Attribute("role", complementaryRole) }

// RoleContentinfo sets role="contentinfo"
func RoleContentinfo() func(*Element) { return Attribute("role", contentinfoRole) }

// RoleDefinition sets role="definition"
func RoleDefinition() func(*Element) { return Attribute("role", definitionRole) }

// RoleDialog sets role="dialog"
func RoleDialog() func(*Element) { return Attribute("role", dialogRole) }

// RoleDocument sets role="document"
func RoleDocument() func(*Element) { return Attribute("role", documentRole) }

// RoleFeed sets role="feed"
func RoleFeed() func(*Element) { return Attribute("role", feedRole) }

// RoleFigure sets role="figure"
func RoleFigure() func(*Element) { return Attribute("role", figureRole) }

// RoleForm sets role="form"
func RoleForm() func(*Element) { return Attribute("role", formRole) }

// RoleGrid sets role="grid"
func RoleGrid() func(*Element) { return Attribute("role", gridRole) }

// RoleGridCell sets role="gridcell"
func RoleGridCell() func(*Element) { return Attribute("role", gridcellRole) }

// RoleGroup sets role="group"
func RoleGroup() func(*Element) { return Attribute("role", groupRole) }

// RoleHeading sets role="heading"
func RoleHeading() func(*Element) { return Attribute("role", headingRole) }

// RoleImg sets role="img"
func RoleImg() func(*Element) { return Attribute("role", imgRole) }

// RoleInput sets role="input"
func RoleInput() func(*Element) { return Attribute("role", inputRole) }

// RoleLandmark sets role="landmark"
func RoleLandmark() func(*Element) { return Attribute("role", landmarkRole) }

// RoleLink sets role="link"
func RoleLink() func(*Element) { return Attribute("role", linkRole) }

// RoleList sets role="list"
func RoleList() func(*Element) { return Attribute("role", listRole) }

// RoleListbox sets role="listbox"
func RoleListbox() func(*Element) { return Attribute("role", listboxRole) }

// RoleListitem sets role="listitem"
func RoleListitem() func(*Element) { return Attribute("role", listitemRole) }

// RoleLog sets role="log"
func RoleLog() func(*Element) { return Attribute("role", logRole) }

// RoleMain sets role="main"
func RoleMain() func(*Element) { return Attribute("role", mainRole) }

// RoleMarquee sets role="marquee"
func RoleMarquee() func(*Element) { return Attribute("role", marqueeRole) }

// RoleMath sets role="math"
func RoleMath() func(*Element) { return Attribute("role", mathRole) }

// RoleMenu sets role="menu"
func RoleMenu() func(*Element) { return Attribute("role", menuRole) }

// RoleMenubar sets role="menubar"
func RoleMenubar() func(*Element) { return Attribute("role", menubarRole) }

// RoleNavigation sets role="navigation"
func RoleNavigation() func(*Element) { return Attribute("role", navigationRole) }

// RoleNone sets role="none"
func RoleNone() func(*Element) { return Attribute("role", noneRole) }

// RoleNote sets role="note"
func RoleNote() func(*Element) { return Attribute("role", noteRole) }

// RoleOption sets role="option"
func RoleOption() func(*Element) { return Attribute("role", optionRole) }

// RolePresentation sets role="presentation"
func RolePresentation() func(*Element) { return Attribute("role", presentationRole) }

// RoleProgressBar sets role="progressbar"
func RoleProgressBar() func(*Element) { return Attribute("role", progressbarRole) }

// RoleRadio sets role="radio"
func RoleRadio() func(*Element) { return Attribute("role", radioRole) }

// RoleRadiogroup sets role="radiogroup"
func RoleRadiogroup() func(*Element) { return Attribute("role", radiogroupRole) }

// RoleRegion sets role="region"
func RoleRegion() func(*Element) { return Attribute("role", regionRole) }

// RoleRow sets role="row"
func RoleRow() func(*Element) { return Attribute("role", rowRole) }

// RoleRowgroup sets role="rowgroup"
func RoleRowgroup() func(*Element) { return Attribute("role", rowgroupRole) }

// RoleRowheader sets role="rowheader"
func RoleRowheader() func(*Element) { return Attribute("role", rowheaderRole) }

// RoleScrollbar sets role="scrollbar"
func RoleScrollbar() func(*Element) { return Attribute("role", scrollbarRole) }

// RoleSearch sets role="search"
func RoleSearch() func(*Element) { return Attribute("role", searchRole) }

// RoleSearchbox sets role="searchbox"
func RoleSearchbox() func(*Element) { return Attribute("role", searchboxRole) }

// RoleSeparator sets role="separator"
func RoleSeparator() func(*Element) { return Attribute("role", separatorRole) }

// RoleSlider sets role="slider"
func RoleSlider() func(*Element) { return Attribute("role", sliderRole) }

// RoleSpinbutton sets role="spinbutton"
func RoleSpinbutton() func(*Element) { return Attribute("role", spinbuttonRole) }

// RoleStatus sets role="status"
func RoleStatus() func(*Element) { return Attribute("role", statusRole) }

// RoleSwitch sets role="switch"
func RoleSwitch() func(*Element) { return Attribute("role", switchRole) }

// RoleTab sets role="tab"
func RoleTab() func(*Element) { return Attribute("role", tabRole) }

// RoleTable sets role="table"
func RoleTable() func(*Element) { return Attribute("role", tableRole) }

// RoleTablist sets role="tablist"
func RoleTablist() func(*Element) { return Attribute("role", tablistRole) }

// RoleTerm sets role="term"
func RoleTerm() func(*Element) { return Attribute("role", termRole) }

// RoleTextbox sets role="textbox"
func RoleTextbox() func(*Element) { return Attribute("role", textboxRole) }

// RoleTimer sets role="timer"
func RoleTimer() func(*Element) { return Attribute("role", timerRole) }

// RoleToolbar sets role="toolbar"
func RoleToolbar() func(*Element) { return Attribute("role", toolbarRole) }

// RoleTooltip sets role="tooltip"
func RoleTooltip() func(*Element) { return Attribute("role", tooltipRole) }

// RoleTree sets role="tree"
func RoleTree() func(*Element) { return Attribute("role", treeRole) }

// RoleTreegrid sets role="treegrid"
func RoleTreegrid() func(*Element) { return Attribute("role", treegridRole) }

// RoleTreeitem sets role="treeitem"
func RoleTreeitem() func(*Element) { return Attribute("role", treeitemRole) }

// RoleWidget sets role="widget"
func RoleWidget() func(*Element) { return Attribute("role", widgetRole) }

// RoleWindow sets role="window"
func RoleWindow() func(*Element) { return Attribute("role", windowRole) }

func Height(h int) func(*Element) { return Attribute("height", strconv.Itoa(h)) }

func Width(w int) func(*Element) { return Attribute("width", strconv.Itoa(w)) }

func SVG(options ...func(*Element)) Element {
	return New(svgElement, options...)
}

func Path(options ...func(*Element)) Element {
	return New(pathElement, options...)
}

// ViewBox sets the viewBox attribute
func ViewBox(x, y, w, h int) func(*Element) {
	return Attribute("viewBox", fmt.Sprintf("%d %d %d %d", x, y, w, h))
}

// Fill sets the fill attribute
func Fill(fill string) func(*Element) {
	return Attribute("fill", fill)
}

// FillRule sets the fill-rule attribute
func FillRule(rule string) func(*Element) {
	return Attribute("fill-rule", rule)
}

// ClipRule sets the clip-rule attribute
func ClipRule(rule string) func(*Element) {
	return Attribute("clip-rule", rule)
}

// Stroke sets the stroke attribute
func Stroke(stroke string) func(*Element) {
	return Attribute("stroke", stroke)
}

// StrokeWidth sets the stroke-width attribute
func StrokeWidth(width int) func(*Element) {
	return Attribute("stroke-width", strconv.Itoa(width))
}

func D(d string) func(*Element) {
	return Attribute("d", d)
}
