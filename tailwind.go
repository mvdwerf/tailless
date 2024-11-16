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

	s.Set("slate-50", "#f8fafc")
	s.Set("slate-100", "#f1f5f9")
	s.Set("slate-200", "#e2e8f0")
	s.Set("slate-300", "#cbd5e1")
	s.Set("slate-400", "#94a3b8")
	s.Set("slate-500", "#64748b")
	s.Set("slate-600", "#475569")
	s.Set("slate-700", "#334155")
	s.Set("slate-800", "#1e293b")
	s.Set("slate-900", "#0f172a")
	s.Set("slate-950", "#020617")

	s.Set("gray-50", "#f9fafb")
	s.Set("gray-100", "#f3f4f6")
	s.Set("gray-200", "#e5e7eb")
	s.Set("gray-300", "#d1d5db")
	s.Set("gray-400", "#9ca3af")
	s.Set("gray-500", "#6b7280")
	s.Set("gray-600", "#4b5563")
	s.Set("gray-700", "#374151")
	s.Set("gray-800", "#1f2937")
	s.Set("gray-900", "#111827")
	s.Set("gray-950", "#030712")

	s.Set("zinc-50", "#fafafa")
	s.Set("zinc-100", "#f4f4f5")
	s.Set("zinc-200", "#e4e4e7")
	s.Set("zinc-300", "#d4d4d8")
	s.Set("zinc-400", "#a1a1aa")
	s.Set("zinc-500", "#71717a")
	s.Set("zinc-600", "#52525b")
	s.Set("zinc-700", "#3f3f46")
	s.Set("zinc-800", "#27272a")
	s.Set("zinc-900", "#18181b")
	s.Set("zinc-950", "#09090b")

	s.Set("neutral-50", "#fafafa")
	s.Set("neutral-100", "#f5f5f5")
	s.Set("neutral-200", "#e5e5e5")
	s.Set("neutral-300", "#d4d4d4")
	s.Set("neutral-400", "#a3a3a3")
	s.Set("neutral-500", "#737373")
	s.Set("neutral-600", "#525252")
	s.Set("neutral-700", "#404040")
	s.Set("neutral-800", "#262626")
	s.Set("neutral-900", "#171717")
	s.Set("neutral-950", "#0a0a0a")

	s.Set("stone-50", "#fafaf9")
	s.Set("stone-100", "#f5f5f4")
	s.Set("stone-200", "#e7e5e4")
	s.Set("stone-300", "#d6d3d1")
	s.Set("stone-400", "#a8a29e")
	s.Set("stone-500", "#78716c")
	s.Set("stone-600", "#57534e")
	s.Set("stone-700", "#44403c")
	s.Set("stone-800", "#292524")
	s.Set("stone-900", "#1c1917")
	s.Set("stone-950", "#0c0a09")

	s.Set("red-50", "#fef2f2")
	s.Set("red-100", "#fee2e2")
	s.Set("red-200", "#fecaca")
	s.Set("red-300", "#fca5a5")
	s.Set("red-400", "#f87171")
	s.Set("red-500", "#ef4444")
	s.Set("red-600", "#dc2626")
	s.Set("red-700", "#b91c1c")
	s.Set("red-800", "#991b1b")
	s.Set("red-900", "#7f1d1d")
	s.Set("red-950", "#450a0a")

	s.Set("orange-50", "#fff7ed")
	s.Set("orange-100", "#ffedd5")
	s.Set("orange-200", "#fed7aa")
	s.Set("orange-300", "#fdba74")
	s.Set("orange-400", "#fb923c")
	s.Set("orange-500", "#f97316")
	s.Set("orange-600", "#ea580c")
	s.Set("orange-700", "#c2410c")
	s.Set("orange-800", "#9a3412")
	s.Set("orange-900", "#7c2d12")
	s.Set("orange-950", "#431407")

	s.Set("amber-50", "#fffbeb")
	s.Set("amber-100", "#fef3c7")
	s.Set("amber-200", "#fde68a")
	s.Set("amber-300", "#fcd34d")
	s.Set("amber-400", "#fbbf24")
	s.Set("amber-500", "#f59e0b")
	s.Set("amber-600", "#d97706")
	s.Set("amber-700", "#b45309")
	s.Set("amber-800", "#92400e")
	s.Set("amber-900", "#78350f")
	s.Set("amber-950", "#451a03")

	s.Set("yellow-50", "#fefce8")
	s.Set("yellow-100", "#fef9c3")
	s.Set("yellow-200", "#fef08a")
	s.Set("yellow-300", "#fde047")
	s.Set("yellow-400", "#facc15")
	s.Set("yellow-500", "#eab308")
	s.Set("yellow-600", "#ca8a04")
	s.Set("yellow-700", "#a16207")
	s.Set("yellow-800", "#854d0e")
	s.Set("yellow-900", "#713f12")
	s.Set("yellow-950", "#422006")

	s.Set("lime-50", "#f7fee7")
	s.Set("lime-100", "#ecfccb")
	s.Set("lime-200", "#d9f99d")
	s.Set("lime-300", "#bef264")
	s.Set("lime-400", "#a3e635")
	s.Set("lime-500", "#84cc16")
	s.Set("lime-600", "#65a30d")
	s.Set("lime-700", "#4d7c0f")
	s.Set("lime-800", "#3f6212")
	s.Set("lime-900", "#365314")
	s.Set("lime-950", "#1a2e05")

	s.Set("green-50", "#f0fdf4")
	s.Set("green-100", "#dcfce7")
	s.Set("green-200", "#bbf7d0")
	s.Set("green-300", "#86efac")
	s.Set("green-400", "#4ade80")
	s.Set("green-500", "#22c55e")
	s.Set("green-600", "#16a34a")
	s.Set("green-700", "#15803d")
	s.Set("green-800", "#166534")
	s.Set("green-900", "#14532d")
	s.Set("green-950", "#052e16")

	s.Set("emerald-50", "#ecfdf5")
	s.Set("emerald-100", "#d1fae5")
	s.Set("emerald-200", "#a7f3d0")
	s.Set("emerald-300", "#6ee7b7")
	s.Set("emerald-400", "#34d399")
	s.Set("emerald-500", "#10b981")
	s.Set("emerald-600", "#059669")
	s.Set("emerald-700", "#047857")
	s.Set("emerald-800", "#065f46")
	s.Set("emerald-900", "#064e3b")
	s.Set("emerald-950", "#022c22")

	s.Set("teal-50", "#f0fdfa")
	s.Set("teal-100", "#ccfbf1")
	s.Set("teal-200", "#99f6e4")
	s.Set("teal-300", "#5eead4")
	s.Set("teal-400", "#2dd4bf")
	s.Set("teal-500", "#14b8a6")
	s.Set("teal-600", "#0d9488")
	s.Set("teal-700", "#0f766e")
	s.Set("teal-800", "#115e59")
	s.Set("teal-900", "#134e4a")
	s.Set("teal-950", "#042f2e")

	s.Set("cyan-50", "#ecfeff")
	s.Set("cyan-100", "#cffafe")
	s.Set("cyan-200", "#a5f3fc")
	s.Set("cyan-300", "#67e8f9")
	s.Set("cyan-400", "#22d3ee")
	s.Set("cyan-500", "#06b6d4")
	s.Set("cyan-600", "#0891b2")
	s.Set("cyan-700", "#0e7490")
	s.Set("cyan-800", "#155e75")
	s.Set("cyan-900", "#164e63")
	s.Set("cyan-950", "#083344")

	s.Set("sky-50", "#f0f9ff")
	s.Set("sky-100", "#e0f2fe")
	s.Set("sky-200", "#bae6fd")
	s.Set("sky-300", "#7dd3fc")
	s.Set("sky-400", "#38bdf8")
	s.Set("sky-500", "#0ea5e9")
	s.Set("sky-600", "#0284c7")
	s.Set("sky-700", "#0369a1")
	s.Set("sky-800", "#075985")
	s.Set("sky-900", "#0c4a6e")
	s.Set("sky-950", "#082f49")

	s.Set("blue-50", "#eff6ff")
	s.Set("blue-100", "#dbeafe")
	s.Set("blue-200", "#bfdbfe")
	s.Set("blue-300", "#93c5fd")
	s.Set("blue-400", "#60a5fa")
	s.Set("blue-500", "#3b82f6")
	s.Set("blue-600", "#2563eb")
	s.Set("blue-700", "#1d4ed8")
	s.Set("blue-800", "#1e40af")
	s.Set("blue-900", "#1e3a8a")
	s.Set("blue-950", "#172554")

	s.Set("indigo-50", "#eef2ff")
	s.Set("indigo-100", "#e0e7ff")
	s.Set("indigo-200", "#c7d2fe")
	s.Set("indigo-300", "#a5b4fc")
	s.Set("indigo-400", "#818cf8")
	s.Set("indigo-500", "#6366f1")
	s.Set("indigo-600", "#4f46e5")
	s.Set("indigo-700", "#4338ca")
	s.Set("indigo-800", "#3730a3")
	s.Set("indigo-900", "#312e81")
	s.Set("indigo-950", "#1e1b4b")

	s.Set("violet-50", "#f5f3ff")
	s.Set("violet-100", "#ede9fe")
	s.Set("violet-200", "#ddd6fe")
	s.Set("violet-300", "#c4b5fd")
	s.Set("violet-400", "#a78bfa")
	s.Set("violet-500", "#8b5cf6")
	s.Set("violet-600", "#7c3aed")
	s.Set("violet-700", "#6d28d9")
	s.Set("violet-800", "#5b21b6")
	s.Set("violet-900", "#4c1d95")
	s.Set("violet-950", "#2e1065")

	s.Set("purple-50", "#faf5ff")
	s.Set("purple-100", "#f3e8ff")
	s.Set("purple-200", "#e9d5ff")
	s.Set("purple-300", "#d8b4fe")
	s.Set("purple-400", "#c084fc")
	s.Set("purple-500", "#a855f7")
	s.Set("purple-600", "#9333ea")
	s.Set("purple-700", "#7e22ce")
	s.Set("purple-800", "#6b21a8")
	s.Set("purple-900", "#581c87")
	s.Set("purple-950", "#3b0764")

	s.Set("fuchsia-50", "#fdf4ff")
	s.Set("fuchsia-100", "#fae8ff")
	s.Set("fuchsia-200", "#f5d0fe")
	s.Set("fuchsia-300", "#f0abfc")
	s.Set("fuchsia-400", "#e879f9")
	s.Set("fuchsia-500", "#d946ef")
	s.Set("fuchsia-600", "#c026d3")
	s.Set("fuchsia-700", "#a21caf")
	s.Set("fuchsia-800", "#86198f")
	s.Set("fuchsia-900", "#701a75")
	s.Set("fuchsia-950", "#4a044e")

	s.Set("pink-50", "#fdf2f8")
	s.Set("pink-100", "#fce7f3")
	s.Set("pink-200", "#fbcfe8")
	s.Set("pink-300", "#f9a8d4")
	s.Set("pink-400", "#f472b6")
	s.Set("pink-500", "#ec4899")
	s.Set("pink-600", "#db2777")
	s.Set("pink-700", "#be185d")
	s.Set("pink-800", "#9d174d")
	s.Set("pink-900", "#831843")
	s.Set("pink-950", "#500724")

	s.Set("rose-50", "#fff1f2")
	s.Set("rose-100", "#ffe4e6")
	s.Set("rose-200", "#fecdd3")
	s.Set("rose-300", "#fda4af")
	s.Set("rose-400", "#fb7185")
	s.Set("rose-500", "#f43f5e")
	s.Set("rose-600", "#e11d48")
	s.Set("rose-700", "#be123c")
	s.Set("rose-800", "#9f1239")
	s.Set("rose-900", "#881337")
	s.Set("rose-950", "#4c0519")
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
