package html

type Attributes map[string]string

// Merge merges b into a, overwriting any existing keys
func (a Attributes) Merge(b Attributes) {
	for k, v := range b {
		a[k] = v
	}
}

var AttributeNames = append(globalAttributes, otherAttributes...)

var globalAttributes = []string{
	accesskeyAttribute,
	autocapitalizeAttribute,
	autofocusAttribute,
	classAttribute,
	contenteditableAttribute,
	dirAttribute,
	draggableAttribute,
	enterkeyhintAttribute,
	exportpartsAttribute,
	hiddenAttribute,
	idAttribute,
	inertAttribute,
	inputmodeAttribute,
	isAttribute,
	itemidAttribute,
	itempropAttribute,
	itemrefAttribute,
	itemscopeAttribute,
	itemtypeAttribute,
	langAttribute,
	nonceAttribute,
	partAttribute,
	popoverAttribute,
	slotAttribute,
	spellcheckAttribute,
	styleAttribute,
	tabindexAttribute,
	titleAttribute,
	translateAttribute,
	writingsuggestionsAttribute,
}

var otherAttributes = []string{
	acceptAttribute,
	autocompleteAttribute,
	captureAttribute,
	crossoriginAttribute,
	dirnameAttribute,
	disabledAttribute,
	elementtimingAttribute,
	forAttribute,
	maxAttribute,
	maxlengthAttribute,
	minAttribute,
	minlengthAttribute,
	multipleAttribute,
	patternAttribute,
	placeholderAttribute,
	readonlyAttribute,
	relAttribute,
	requiredAttribute,
	sizeAttribute,
	stepAttribute,
	charsetAttribute,
	contentAttribute,
}

const (
	charsetAttribute            = "charset"
	contentAttribute            = "content"
	accesskeyAttribute          = "accesskey"
	autocapitalizeAttribute     = "autocapitalize"
	autofocusAttribute          = "autofocus"
	classAttribute              = "class"
	contenteditableAttribute    = "contenteditable"
	dirAttribute                = "dir"
	draggableAttribute          = "draggable"
	enterkeyhintAttribute       = "enterkeyhint"
	exportpartsAttribute        = "exportparts"
	hiddenAttribute             = "hidden"
	idAttribute                 = "id"
	inertAttribute              = "inert"
	inputmodeAttribute          = "inputmode"
	isAttribute                 = "is"
	itemidAttribute             = "itemid"
	itempropAttribute           = "itemprop"
	itemrefAttribute            = "itemref"
	itemscopeAttribute          = "itemscope"
	itemtypeAttribute           = "itemtype"
	langAttribute               = "lang"
	nonceAttribute              = "nonce"
	partAttribute               = "part"
	popoverAttribute            = "popover"
	slotAttribute               = "slot"
	spellcheckAttribute         = "spellcheck"
	styleAttribute              = "style"
	tabindexAttribute           = "tabindex"
	titleAttribute              = "title"
	translateAttribute          = "translate"
	writingsuggestionsAttribute = "writingsuggestions"
	acceptAttribute             = "accept"
	autocompleteAttribute       = "autocomplete"
	captureAttribute            = "capture"
	crossoriginAttribute        = "crossorigin"
	dirnameAttribute            = "dirname"
	disabledAttribute           = "disabled"
	elementtimingAttribute      = "elementtiming"
	forAttribute                = "for"
	maxAttribute                = "max"
	maxlengthAttribute          = "maxlength"
	minAttribute                = "min"
	minlengthAttribute          = "minlength"
	multipleAttribute           = "multiple"
	patternAttribute            = "pattern"
	placeholderAttribute        = "placeholder"
	readonlyAttribute           = "readonly"
	relAttribute                = "rel"
	requiredAttribute           = "required"
	sizeAttribute               = "size"
	stepAttribute               = "step"
)

const (
	alertRole            = "alert"
	alertdialogRole      = "alertdialog"
	applicationRole      = "application"
	articleRole          = "article"
	bannerRole           = "banner"
	buttonRole           = "button"
	cellRole             = "cell"
	checkboxRole         = "checkbox"
	columnheaderRole     = "columnheader"
	comboboxRole         = "combobox"
	commandRole          = "command"
	commentRole          = "comment"
	complementaryRole    = "complementary"
	compositeRole        = "composite"
	contentinfoRole      = "contentinfo"
	definitionRole       = "definition"
	dialogRole           = "dialog"
	documentRole         = "document"
	feedRole             = "feed"
	figureRole           = "figure"
	formRole             = "form"
	genericRole          = "generic"
	gridRole             = "grid"
	gridcellRole         = "gridcell"
	groupRole            = "group"
	headingRole          = "heading"
	imgRole              = "img"
	inputRole            = "input"
	landmarkRole         = "landmark"
	linkRole             = "link"
	listRole             = "list"
	listboxRole          = "listbox"
	listitemRole         = "listitem"
	logRole              = "log"
	mainRole             = "main"
	markRole             = "mark"
	marqueeRole          = "marquee"
	mathRole             = "math"
	menuRole             = "menu"
	menubarRole          = "menubar"
	menuitemRole         = "menuitem"
	menuitemcheckboxRole = "menuitemcheckbox"
	menuitemradioRole    = "menuitemradio"
	meterRole            = "meter"
	navigationRole       = "navigation"
	noneRole             = "none"
	noteRole             = "note"
	optionRole           = "option"
	presentationRole     = "presentation"
	progressbarRole      = "progressbar"
	radioRole            = "radio"
	radiogroupRole       = "radiogroup"
	rangeRole            = "range"
	regionRole           = "region"
	roletypeRole         = "roletype"
	rowRole              = "row"
	rowgroupRole         = "rowgroup"
	rowheaderRole        = "rowheader"
	scrollbarRole        = "scrollbar"
	searchRole           = "search"
	searchboxRole        = "searchbox"
	sectionRole          = "section"
	sectionheadRole      = "sectionhead"
	selectRole           = "select"
	separatorRole        = "separator"
	sliderRole           = "slider"
	spinbuttonRole       = "spinbutton"
	statusRole           = "status"
	structureRole        = "structure"
	suggestionRole       = "suggestion"
	switchRole           = "switch"
	tabRole              = "tab"
	tableRole            = "table"
	tablistRole          = "tablist"
	tabpanelRole         = "tabpanel"
	termRole             = "term"
	textboxRole          = "textbox"
	timerRole            = "timer"
	toolbarRole          = "toolbar"
	tooltipRole          = "tooltip"
	treeRole             = "tree"
	treegridRole         = "treegrid"
	treeitemRole         = "treeitem"
	widgetRole           = "widget"
	windowRole           = "window"
)
