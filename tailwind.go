package tailless

import (
	"strings"
)

type stringMap map[string]string

type tailwindCollection struct {
	Items stringMap
}

func newTailwindCollection() *tailwindCollection {
	collection := tailwindCollection{}
	collection.Items = make(map[string]string)

	initTailwind(&collection)

	return &collection
}

func (t *tailwindCollection) Get(name string) *selectorNode {
	value := t.Items[name]
	if value == "" {
		return nil
	}

	children := make([]node, 0)
	child := newDeclarationNode(value, 0)
	children = append(children, child)

	n := selectorNode{}
	n.Children = children

	return &n
}

func (t *tailwindCollection) Set(name string, node *selectorNode) {

}

func (t *tailwindCollection) Add(name string, value string) {
	t.Items[name] = value
}

func initTailwind(c *tailwindCollection) {
	initColors(c, "text", "color: $1;")
	initColors(c, "bg", "background-color: $1;")

	initSizes(c, "p", "padding: $1;")
	initSizes(c, "px", "padding-left: $1; padding-right: $1;")
	initSizes(c, "py", "padding-top: $1; padding-bottom: $1;")
	initSizes(c, "pl", "padding-left: $1;")
	initSizes(c, "pt", "padding-top: $1;")
	initSizes(c, "pr", "padding-right: $1;")
	initSizes(c, "pb", "padding-bottom: $1;")

	initSizes(c, "m", "margin: $1")
	initSizes(c, "mx", "margin-left: $1; margin-right: $1;")
	c.Add(".mx-auto", "margin-left: auto; margin-right: auto;")
	initSizes(c, "my", "margin-top: $1; margin-bottom: $1;")
	initSizes(c, "ml", "margin-left: $1;")
	initSizes(c, "mt", "margin-top: $1;")
	initSizes(c, "mr", "margin-right: $1;")
	initSizes(c, "mb", "margin-bottom: $1;")

	initSizes(c, "inset", "inset: $1;")
	initSizes(c, "inset-x", "left: $1; right: $1;")
	initSizes(c, "inset-y", "top: $1; bottom: $1;")
	initSizes(c, "top", "top: $1;")
	initSizes(c, "right", "right: $1;")
	initSizes(c, "bottom", "bottom: $1;")
	initSizes(c, "left", "left: $1;")
	initSizes(c, "gap", "gap: $1;")

	initWidthHeight(c, "w", "width: $1;")
	initWidthHeight(c, "h", "height: $1;")
	initWidth(c, "w", "width: $1;")
	c.Add(".h-screen", "height: 100vh;")
	initMinWidth(c, "min-w", "min-width: $1;")
	initMaxWidth(c, "max-w", "max-width: $1;")
	initMinHeight(c, "min-h", "min-height: $1;")
	initMaxHeight(c, "max-h", "max-height: $1;")

	initFontSizes(c, "text", "font-size: $1; line-height: $2;")
	initLeading(c, "leading", "line-height: $1;")
	initColors(c, "text", "color: $1;")
	initTextAlign(c, "text", "text-align: $1;")

	initFontWeights(c, "font", "font-weight: $1;")

	initBorderRadius(c, "rounded", "border-radius: $1;")
	initBorderRadius(c, "rounded-t", "border-top-left-radius: $1; border-top-right-radius: $1;")
	initBorderRadius(c, "rounded-r", "border-top-right-radius: $1; border-bottom-right-radius: $1;")
	initBorderRadius(c, "rounded-b", "border-bottom-left-radius: $1; border-bottom-right-radius: $1;")
	initBorderRadius(c, "rounded-l", "border-top-left-radius: $1; border-bottom-left-radius: $1;")

	initBorderWidth(c, "border", "border-width: $1;")
	initBorderWidth(c, "border-l", "border-left-width: $1;")
	initBorderWidth(c, "border-t", "border-top-width: $1;")
	initBorderWidth(c, "border-r", "border-right-width: $1;")
	initBorderWidth(c, "border-b", "border-bottom-width: $1;")
	initColors(c, "border", "border-color: $1;")
	initBorderStyle(c, "border", "border-style: $1;")

	initOutlineWidth(c, "outline", "outline-width: $1;")
	initColors(c, "outline", "outline-color: $1;")
	initOutlineStyle(c, "outline", "outline-style: $1;")
	initOutlineWidth(c, "outline-offset", "outline-offset: $1;")
	c.Add(".outline-none", "outline: none;")

	c.Add(".underline", "text-decoration: underline;")
	c.Add(".overline", "text-decoration: overline;")
	c.Add(".line-through", "text-decoration: linie-through;")
	c.Add(".no-underline", "text-decoration: none;")
	initColors(c, "decoration", "text-decoration-color: $1;")
	initOverflow(c, "overflow", "overflow: $1;")
	initOverflow(c, "overflow-x", "overflow-x: $1;")
	initOverflow(c, "overflow-y", "overflow-y: $1;")

	initObjectFit(c, "object", "object-fit: $1;")
	initObjectPosition(c, "object", "object-position: $1;")
	initOpacity(c, "opacity", "opacity: $1;")

	initLetterSpacing(c, "tracking", "letter-spacing: $1;")

	c.Add(".truncate", "overflow: hidden; text-overflow: ellipsis; white-space: nowrap;")
	c.Add(".uppercase", "text-transform: uppercase;")
	c.Add(".text-ellipsis", "text-overflow: ellipsis;")
	c.Add(".text-clip", "text-overflow: clip;")

	c.Add(".block", "display: block;")
	c.Add(".inline-block", "display: inline-block;")
	c.Add(".inline", "display: inline;")
	c.Add(".flex", "display: flex;")
	c.Add(".inline-flex", "display: inline-flex;")
	c.Add(".grid", "display: grid")
	c.Add(".inline-grid", "display: inline-grid")
	c.Add(".contents", "display: contents")
	c.Add(".list-item", "display: list-item")
	c.Add(".hidden", "display: none;")

	c.Add(".static", "position: static;")
	c.Add(".fixed", "position: fixed;")
	c.Add(".absolute", "position: absolute;")
	c.Add(".relative", "position: relative;")
	c.Add(".sticky", "position: sticky;")

	c.Add(".float-left", "float: left;")
	c.Add(".float-right", "float: right;")
	c.Add(".float-none", "float: none;")

	c.Add(".flex-row", "flex-direction: row;")
	c.Add(".flex-row-reverse", "flex-direction: row-reverse;")
	c.Add(".flex-col", "flex-direction: column;")
	c.Add(".flex-col-reverse", "flex-direction: column-reverse;")
	initSizes(c, "basis", "flex-basis: $1;")
	c.Add(".flex-wrap", "flex-wrap: wrap;")
	c.Add(".flex-wrap-reverse", "flex-wrap: wrap-reverse;")
	c.Add(".flex-nowrap", "flex-wrap: nowrap;")
	c.Add(".flex-1", "flex: 1 1 0%;")
	c.Add(".flex-auto", "flex: 1 1 auto;")
	c.Add(".flex-initial", "flex: 0 1 auto;")
	c.Add(".flex-none", "flex: none;")
	c.Add(".grow", "flex-grow: 1;")
	c.Add(".grow-0", "flex-grow: 0;")
	c.Add(".shrink", "flex-shrink: 1;")
	c.Add(".shrink-0", "flex-shrink: 0;")
	c.Add(".flex-grow", "flex-grow: 1;")
	c.Add(".flex-grow-0", "flex-grow: 0;")
	c.Add(".flex-shrink", "flex-shrink: 1;")
	c.Add(".flex-shrink-0", "flex-shrink: 0;")

	c.Add(".text-wrap", "text-wrap: wrap;")
	c.Add(".text-nowrap", "text-wrap: nowrap;")
	c.Add(".text-balance", "text-wrap: balance;")
	c.Add(".text-pretty", "text-wrap: pretty;")

	initJustifyContent(c, "justify", "justify-content: $1;")
	initJustifyItems(c, "justify-items", "justify-items: $1;")
	initJustifySelf(c, "justify-self", "justify-self: $1;")
	initAlignContent(c, "content", "align-content: $1;")
	initAlignItems(c, "items", "align-items: $1;")
	initAlignSelf(c, "self", "align-self: $1;")
	initSizes(c, "gap", "gap: $1;")
	initSizes(c, "gap-x", "column-gap: $1;")
	initSizes(c, "gap-y", "row-gap: $1;")
	initZIndex(c, "z", "z-index: $1;")
	initShadow(c, "shadow", "box-shadow: $1;")

	c.Add(".transition-none", "transition-property: none;")
	c.Add(".transition-all", "transition-property: all; transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1); transition-duration: 150ms;")
	c.Add(".transition", "transition-property: color, background-color, border-color, text-decoration-color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter; transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1); transition-duration: 150ms;")
	c.Add(".transition-colors", "transition-property: color, background-color, border-color, text-decoration-color, fill, stroke; transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1); transition-duration: 150ms;")
	c.Add(".transition-opacity", "transition-property: opacity; transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1); transition-duration: 150ms; ")
	c.Add(".transition-shadow", "transition-property: box-shadow; transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1); transition-duration: 150ms;")
	c.Add(".transition-transform", "transition-property: transform; transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1); transition-duration: 150ms;")

	initDurationDelay(c, "duration", "transition-duration: $1;")
	initEase(c, "ease", "transition-timing-function: $1;")
	initDurationDelay(c, "delay", "transition-delay: $1;")

	initScale(c, "scale", "transform: scale($1);")
	initScale(c, "scale-x", "transform: scale-X($1);")
	initScale(c, "scale-y", "transform: scale-Y($1);")
	initRotate(c, "rotate", "transform: rotate($1);")
	initTranslate(c, "translate-x", "transform: translateX($1);")
	initTranslate(c, "translate-y", "transform: translateY($1);")
	initTranslate(c, "-translate-x", "transform: translateX(-$1);")
	initTranslate(c, "-translate-y", "transform: translateY(-$1);")
	initSkew(c, "skew-x", "transform: skewX($1);")
	initSkew(c, "skew-y", "transform: skewY($1);")
	initOrigin(c, "origin", "transform-origin: $1;")

	initCursors(c, "cursor", "cursor: $1;")
	initPointerEvents(c, "pointer-events", "pointer-events: $1;")
}

func initSizes(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("0", "0px")
	s.Set("px", "1px")
	s.Set("0.5", "0.125rem")
	s.Set("1", "0.25rem")
	s.Set("1.5", "0.375rem")
	s.Set("2", "0.5rem")
	s.Set("2.5", "0.625rem")
	s.Set("3", "0.75rem")
	s.Set("3.5", "0.875rem")
	s.Set("4", "1rem")
	s.Set("5", "1.25rem")
	s.Set("6", "1.5rem")
	s.Set("7", "1.75rem")
	s.Set("8", "2rem")
	s.Set("9", "2.25rem")
	s.Set("10", "2.5rem")
	s.Set("11", "2.75rem")
	s.Set("12", "3rem")
	s.Set("14", "3.5rem")
	s.Set("16", "4rem")
	s.Set("20", "5rem")
	s.Set("24", "6rem")
	s.Set("28", "7rem")
	s.Set("32", "8rem")
	s.Set("36", "9rem")
	s.Set("40", "10rem")
	s.Set("44", "11rem")
	s.Set("48", "12rem")
	s.Set("52", "13rem")
	s.Set("56", "14rem")
	s.Set("60", "15rem")
	s.Set("64", "16rem")
	s.Set("72", "18rem")
	s.Set("80", "20rem")
	s.Set("88", "22rem")
	s.Set("96", "24rem")
}

func initColors(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("inherit", "inherit")
	s.Set("current", "currentColor")
	s.Set("transparent", "transparent")
	s.Set("black", "#000000")
	s.Set("white", "#ffffff")

	for name, value := range *colors {
		s.Set(name, value)
	}
}

func initWidthHeight(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("0", "0px")
	s.Set("px", "1px")
	s.Set("0.5", "0.125rem")
	s.Set("1", "0.25rem")
	s.Set("1.5", "0.375rem")
	s.Set("2", "0.5rem")
	s.Set("2.5", "0.625rem")
	s.Set("3", "0.75rem")
	s.Set("3.5", "0.875rem")
	s.Set("4", "1rem")
	s.Set("5", "1.25rem")
	s.Set("6", "1.5rem")
	s.Set("7", "1.75rem")
	s.Set("8", "2rem")
	s.Set("9", "2.25rem")
	s.Set("10", "2.5rem")
	s.Set("11", "2.75rem")
	s.Set("12", "3rem")
	s.Set("14", "3.5rem")
	s.Set("16", "4rem")
	s.Set("20", "5rem")
	s.Set("24", "6rem")
	s.Set("28", "7rem")
	s.Set("32", "8rem")
	s.Set("36", "9rem")
	s.Set("40", "10rem")
	s.Set("44", "11rem")
	s.Set("48", "12rem")
	s.Set("52", "13rem")
	s.Set("56", "14rem")
	s.Set("60", "15rem")
	s.Set("64", "16rem")
	s.Set("72", "18rem")
	s.Set("80", "20rem")
	s.Set("96", "24rem")
	s.Set("auto", "auto")
	s.Set("1/2", "50%")
	s.Set("1/3", "33.333333%")
	s.Set("2/3", "66.666667%")
	s.Set("1/4", "25%")
	s.Set("2/4", "50%")
	s.Set("3/4", "75%")
	s.Set("1/5", "20%")
	s.Set("2/5", "40%")
	s.Set("3/5", "60%")
	s.Set("4/5", "80%")
	s.Set("1/6", "16.666667%")
	s.Set("2/6", "33.333333%")
	s.Set("3/6", "50%")
	s.Set("4/6", "66.666667%")
	s.Set("5/6", "83.333333%")
	s.Set("full", "100%")
	s.Set("min", "min-content")
	s.Set("max", "max-content")
	s.Set("fit", "fit-content")
}

func initWidth(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("1/12", "8.333333%")
	s.Set("2/12", "16.666667%")
	s.Set("3/12", "25%")
	s.Set("4/12", "33.333333%")
	s.Set("5/12", "41.666667%")
	s.Set("6/12", "50%")
	s.Set("7/12", "58.333333%")
	s.Set("8/12", "66.666667%")
	s.Set("9/12", "75%")
	s.Set("10/12", "83.333333%")
	s.Set("11/12", "91.666667%")
	s.Set("screen", "100vw")
}

func initMinWidth(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("0", "0px")
	s.Set("1", "0.25rem")
	s.Set("2", "0.5rem")
	s.Set("3", "0.75rem")
	s.Set("4", "1rem")
	s.Set("5", "1.25rem")
	s.Set("6", "1.5rem")
	s.Set("7", "1.75rem")
	s.Set("8", "2rem")
	s.Set("9", "2.25rem")
	s.Set("10", "2.5rem")
	s.Set("11", "2.75rem")
	s.Set("12", "3rem")
	s.Set("14", "3.5rem")
	s.Set("16", "4rem")
	s.Set("20", "5rem")
	s.Set("24", "6rem")
	s.Set("28", "7rem")
	s.Set("32", "8rem")
	s.Set("36", "9rem")
	s.Set("40", "10rem")
	s.Set("44", "11rem")
	s.Set("48", "12rem")
	s.Set("52", "13rem")
	s.Set("56", "14rem")
	s.Set("60", "15rem")
	s.Set("64", "16rem")
	s.Set("72", "18rem")
	s.Set("80", "20rem")
	s.Set("96", "24rem")

	s.Set("px", "1px")
	s.Set("0.5", "0.125rem")
	s.Set("1.5", "0.375rem")
	s.Set("2.5", "0.625rem")
	s.Set("3.5", "0.875rem")

	s.Set("full", "100%")
	s.Set("min", "min-content")
	s.Set("max", "max-content")
	s.Set("fit", "fit-content")
}

func initMaxWidth(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("0", "0rem")
	s.Set("none", "none")
	s.Set("xs", "20rem")
	s.Set("sm", "24rem")
	s.Set("md", "28rem")
	s.Set("lg", "32rem")
	s.Set("xl", "36rem")
	s.Set("2xl", "42rem")
	s.Set("3xl", "48rem")
	s.Set("4xl", "56rem")
	s.Set("5xl", "64rem")
	s.Set("6xl", "72rem")
	s.Set("7xl", "80rem")
	s.Set("full", "100%")
	s.Set("min", "min-content")
	s.Set("max", "max-content")
	s.Set("fit", "fit-content")
	s.Set("prose", "65ch")
	s.Set("screen-sm", "640px")
	s.Set("screen-md", "768px")
	s.Set("screen-lg", "1024px")
	s.Set("screen-xl", "1280px")
	s.Set("screen-2xl", "1536px")
}

func initMinHeight(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("0", "0px")
	s.Set("1", "0.25rem")
	s.Set("2", "0.5rem")
	s.Set("3", "0.75rem")
	s.Set("4", "1rem")
	s.Set("5", "1.25rem")
	s.Set("6", "1.5rem")
	s.Set("7", "1.75rem")
	s.Set("8", "2rem")
	s.Set("9", "2.25rem")
	s.Set("10", "2.5rem")
	s.Set("11", "2.75rem")
	s.Set("12", "3rem")
	s.Set("14", "3.5rem")
	s.Set("16", "4rem")
	s.Set("20", "5rem")
	s.Set("24", "6rem")
	s.Set("28", "7rem")
	s.Set("32", "8rem")
	s.Set("36", "9rem")
	s.Set("40", "10rem")
	s.Set("44", "11rem")
	s.Set("48", "12rem")
	s.Set("52", "13rem")
	s.Set("56", "14rem")
	s.Set("60", "15rem")
	s.Set("64", "16rem")
	s.Set("72", "18rem")
	s.Set("80", "20rem")
	s.Set("96", "24rem")
	s.Set("full", "100%")
	s.Set("screen", "100vh")
	s.Set("min", "min-content")
	s.Set("max", "max-content")
	s.Set("fit", "fit-content")
}

func initMaxHeight(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("0", "0px")
	s.Set("px", "1px")
	s.Set("0.5", "0.125rem")
	s.Set("1", "0.25rem")
	s.Set("1.5", "0.375rem")
	s.Set("2", "0.5rem")
	s.Set("2.5", "0.625rem")
	s.Set("3", "0.75rem")
	s.Set("3.5", "0.875rem")
	s.Set("4", "1rem")
	s.Set("5", "1.25rem")
	s.Set("6", "1.5rem")
	s.Set("7", "1.75rem")
	s.Set("8", "2rem")
	s.Set("9", "2.25rem")
	s.Set("10", "2.5rem")
	s.Set("11", "2.75rem")
	s.Set("12", "3rem")
	s.Set("14", "3.5rem")
	s.Set("16", "4rem")
	s.Set("20", "5rem")
	s.Set("24", "6rem")
	s.Set("28", "7rem")
	s.Set("32", "8rem")
	s.Set("36", "9rem")
	s.Set("40", "10rem")
	s.Set("44", "11rem")
	s.Set("48", "12rem")
	s.Set("52", "13rem")
	s.Set("56", "14rem")
	s.Set("60", "15rem")
	s.Set("64", "16rem")
	s.Set("72", "18rem")
	s.Set("80", "20rem")
	s.Set("96", "24rem")
	s.Set("none", "none")
	s.Set("full", "100%")
	s.Set("screen", "100vh")
	s.Set("min", "min-content")
	s.Set("max", "max-content")
	s.Set("fit", "fit-content")
}

func initLetterSpacing(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("tighter", "-0.05em")
	s.Set("tight", "-0.025em")
	s.Set("normal", "0em")
	s.Set("wide", "0.025em")
	s.Set("wider", "0.05em")
	s.Set("widest", "0.1em")
}

func initBorderRadius(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("none", "0px")
	s.Set("sm", "0.125rem")
	s.Set("", "0.25rem")
	s.Set("md", "0.375rem")
	s.Set("lg", "0.5rem")
	s.Set("xl", "0.75rem")
	s.Set("2xl", "1rem")
	s.Set("3xl", "1.5rem")
	s.Set("full", "9999px")
}

func initBorderWidth(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("0", "0px")
	s.Set("2", "2px")
	s.Set("4", "4px")
	s.Set("8", "8px")
	s.Set("", "1px")
}

func initBorderStyle(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("solid", "solid")
	s.Set("dashed", "dashed")
	s.Set("dotted", "dotted")
	s.Set("double", "double")
	s.Set("hidden", "hidden")
	s.Set("none", "none")
}

func initFontWeights(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("thin", "100")
	s.Set("extralight", "200")
	s.Set("light", "300")
	s.Set("normal", "400")
	s.Set("medium", "500")
	s.Set("semibold", "600")
	s.Set("bold", "700")
	s.Set("extrabold", "800")
	s.Set("black", "900")
}

func initFontSizes(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set2("xs", "0.75rem", "1rem")
	s.Set2("sm", "0.875rem", "1.25rem")
	s.Set2("base", "1rem", "1.5rem")
	s.Set2("lg", "1.125rem", "1.75rem")
	s.Set2("xl", "1.25rem", "1.75rem")
	s.Set2("2xl", "1.5rem", "2rem")
	s.Set2("3xl", "1.875rem", "2.25rem")
	s.Set2("4xl", "2.25rem", "2.5rem")
	s.Set2("5xl", "3rem", "1")
	s.Set2("6xl", "3.75rem", "1")
	s.Set2("7xl", "4.5rem", "1")
	s.Set2("8xl", "6rem", "1")
	s.Set2("9xl", "8rem", "1")
}

func initLeading(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("3", ".75rem")
	s.Set("4", "1rem")
	s.Set("5", "1.25rem")
	s.Set("6", "1.50rem")
	s.Set("7", "1.75rem")
	s.Set("8", "2rem")
	s.Set("9", "2.25rem")
	s.Set("10", "2.5rem")
	s.Set("none", "1")
	s.Set("tight", "1.25")
	s.Set("snug", "1.375")
	s.Set("normal", "1.5")
	s.Set("relaxed", "1.625")
	s.Set("loose", "2")
}

func initJustifyContent(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("normal", "normal")
	s.Set("start", "flex-start")
	s.Set("end", "flex-end")
	s.Set("center", "center")
	s.Set("between", "space-between")
	s.Set("around", "space-around")
	s.Set("evenly", "space-evenly")
	s.Set("stretch", "stretch")
}

func initJustifyItems(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("start", "start")
	s.Set("end", "end")
	s.Set("center", "center")
	s.Set("stretch", "stretch")
}

func initJustifySelf(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("auto", "auto")
	s.Set("start", "start")
	s.Set("end", "end")
	s.Set("center", "center")
	s.Set("stretch", "stretch")
}

func initAlignContent(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("normal", "normal")
	s.Set("center", "center")
	s.Set("start", "flex-start")
	s.Set("end", "flex-end")
	s.Set("between", "space-between")
	s.Set("around", "space-around")
	s.Set("evenly", "space-evenly")
	s.Set("baseline", "baseline")
	s.Set("stretch", "stretch")
}

func initAlignItems(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("start", "start")
	s.Set("end", "end")
	s.Set("center", "center")
	s.Set("baseline", "baseline")
	s.Set("stretch", "stretch")
}

func initAlignSelf(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("auto", "auto")
	s.Set("start", "start")
	s.Set("end", "end")
	s.Set("center", "center")
	s.Set("stretch", "stretch")
	s.Set("baseline", "baseline")
}

func initShadow(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("sm", "0 1px 2px 0 rgb(0 0 0 / 0.05)")
	s.Set("", "0 1px 3px 0 rgb(0 0 0 / 0.1), 0 1px 2px -1px rgb(0 0 0 / 0.1)")
	s.Set("md", "0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1)")
	s.Set("lg", "0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1)")
	s.Set("xl", "0 20px 25px -5px rgb(0 0 0 / 0.1), 0 8px 10px -6px rgb(0 0 0 / 0.1)")
	s.Set("2xl", "0 25px 50px -12px rgb(0 0 0 / 0.25)")
	s.Set("inner", "inset 0 2px 4px 0 rgb(0 0 0 / 0.05)")
	s.Set("none", "0 0 #0000")
}

func initOverflow(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("auto", "auto")
	s.Set("hidden", "hidden")
	s.Set("clip", "clip")
	s.Set("visible", "visible")
	s.Set("scroll", "scroll")
}

func initObjectFit(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("contain", "contain")
	s.Set("cover", "cover")
	s.Set("fill", "fill")
	s.Set("none", "none")
	s.Set("scale-down", "scale-down")
}

func initObjectPosition(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("bottom", "bottom")
	s.Set("center", "center")
	s.Set("left", "left")
	s.Set("left-bottom", "left bottom")
	s.Set("left-top", "left top")
	s.Set("right", "right")
	s.Set("right-bottom", "right-bottom")
	s.Set("right-top", "right-top")
	s.Set("top", "top")
}

func initOpacity(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("0", "0")
	s.Set("5", "0.05")
	s.Set("10", "0.1")
	s.Set("20", "0.2")
	s.Set("25", "0.25")
	s.Set("30", "0.3")
	s.Set("40", "0.4")
	s.Set("50", "0.5")
	s.Set("60", "0.6")
	s.Set("70", "0.7")
	s.Set("75", "0.75")
	s.Set("80", "0.8")
	s.Set("90", "0.9")
	s.Set("95", "0.95")
	s.Set("100", "1")
}

func initEase(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("linear", "linear")
	s.Set("in", "cubic-bezier(0.4, 0, 1, 1)")
	s.Set("out", "cubic-bezier(0, 0, 0.2, 1)")
	s.Set("in-out", "cubic-bezier(0.4, 0, 0.2, 1)")
}

func initDurationDelay(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("0", "0s")
	s.Set("75", "75ms")
	s.Set("100", "100ms")
	s.Set("150", "150ms")
	s.Set("200", "200ms")
	s.Set("300", "300ms")
	s.Set("500", "500ms")
	s.Set("700", "700ms")
	s.Set("1000", "1000ms")
}

func initScale(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("0", "0")
	s.Set("50", ".5")
	s.Set("75", ".75")
	s.Set("90", ".9")
	s.Set("95", ".95")
	s.Set("100", "1")
	s.Set("105", "1.05")
	s.Set("110", "1.1")
	s.Set("125", "1.25")
	s.Set("150", "1.50")
}

func initRotate(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("0", "0deg")
	s.Set("1", "1deg")
	s.Set("2", "2deg")
	s.Set("3", "3deg")
	s.Set("6", "6deg")
	s.Set("12", "12deg")
	s.Set("45", "45deg")
	s.Set("90", "90deg")
	s.Set("180", "180deg")
}

func initTranslate(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("0", "0px")
	s.Set("px", "1px")
	s.Set("0.5", "0.125rem")
	s.Set("1", "0.25rem")
	s.Set("1.5", "0.375rem")
	s.Set("2", "0.5rem")
	s.Set("2.5", "0.625rem")
	s.Set("3", "0.75rem")
	s.Set("3.5", "0.875rem")
	s.Set("4", "1rem")
	s.Set("5", "1.25rem")
	s.Set("6", "1.5rem")
	s.Set("7", "1.75rem")
	s.Set("8", "2rem")
	s.Set("9", "2.25rem")
	s.Set("10", "2.5rem")
	s.Set("11", "2.75rem")
	s.Set("12", "3rem")
	s.Set("14", "3.5rem")
	s.Set("16", "4rem")
	s.Set("20", "5rem")
	s.Set("24", "6rem")
	s.Set("28", "7rem")
	s.Set("32", "8rem")
	s.Set("36", "9rem")
	s.Set("40", "10rem")
	s.Set("44", "11rem")
	s.Set("48", "12rem")
	s.Set("52", "13rem")
	s.Set("56", "14rem")
	s.Set("60", "15rem")
	s.Set("64", "16rem")
	s.Set("72", "18rem")
	s.Set("80", "20rem")
	s.Set("96", "24rem")
	s.Set("1/2", "50%")
	s.Set("1/3", "33.333333%")
	s.Set("2/3", "66.666667%")
	s.Set("1/4", "25%")
	s.Set("2/4", "50%")
	s.Set("3/4", "75%")
	s.Set("full", "100%")
}

func initSkew(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("0", "0deg")
	s.Set("1", "1deg")
	s.Set("2", "2deg")
	s.Set("3", "3deg")
	s.Set("6", "6deg")
	s.Set("12", "12deg")
}

func initOrigin(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("center", "center")
	s.Set("top", "top")
	s.Set("top-right", "top right")
	s.Set("right", "right")
	s.Set("bottom-right", "bottom right")
	s.Set("bottom", "bottom")
	s.Set("bottom-left", "bottom left")
	s.Set("left", "left")
	s.Set("top-left", "top left")
}

func initZIndex(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("0", "0")
	s.Set("10", "10")
	s.Set("20", "20")
	s.Set("30", "30")
	s.Set("40", "40")
	s.Set("50", "50")
	s.Set("auto", "auto")
}

func initOutlineWidth(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("0", "0px")
	s.Set("1", "1px")
	s.Set("2", "2px")
	s.Set("4", "4px")
	s.Set("8", "8px")
}

func initOutlineStyle(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("none", "!outline: 2px solid transparent; outline-offset: 2px;")
	s.Set("", "solid")
	s.Set("dashed", "dashed")
	s.Set("dotted", "dotted")
	s.Set("double", "double")
}

func initCursors(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("auto", "auto")
	s.Set("default", "default")
	s.Set("pointer", "pointer")
	s.Set("wait", "wait")
	s.Set("text", "text")
	s.Set("move", "move")
	s.Set("help", "help")
	s.Set("not-allowed", "not-allowed")
}

func initPointerEvents(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("none", "none")
	s.Set("auto", "auto")
}

func initTextAlign(c *tailwindCollection, prefix string, template string) {
	s := createHelper(c, prefix, template)

	s.Set("left", "left")
	s.Set("center", "center")
	s.Set("right", "right")
	s.Set("justify", "justify")
	s.Set("start", "start")
	s.Set("end", "end")
}

type helper struct {
	Collection   *tailwindCollection
	NamePrefix   string
	TextTemplate string
}

func createHelper(c *tailwindCollection, namePrefix string, textTemplate string) *helper {
	return &helper{Collection: c, NamePrefix: namePrefix, TextTemplate: textTemplate}
}

func (h *helper) Set(name string, value string) {
	n := "." + h.NamePrefix
	if name != "" {
		n += "-" + name
	}
	h.Collection.Add(n, strings.Replace(h.TextTemplate, "$1", value, -1))
}

func (h *helper) Set2(name string, value1 string, value2 string) {
	str := strings.Replace(h.TextTemplate, "$1", value1, -1)
	str = strings.Replace(str, "$2", value2, -1)
	name = "." + h.NamePrefix + "-" + name
	h.Collection.Add(name, str)
}
