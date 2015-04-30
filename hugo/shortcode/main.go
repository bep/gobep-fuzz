package shortcode

import (
	"github.com/spf13/hugo/hugolib"
	"github.com/spf13/hugo/tpl"
	"strings"
)

var t tpl.Template

func init() {
	t := tpl.New()
	t.AddInternalShortcode("inner.html", `Shortcode... {{ with index .Params 0 }}{{ . }}{{ end }}-- {{ with index .Params 1 }}{{ . }}{{ end }}- {{ with .Inner }}{{ . }}{{ end }}`)
	t.AddInternalShortcode("sc1	.html", `sc1: {{ with index .Params 0 }}{{ . }}{{ end }}-- {{ with index .Params 1 }}{{ . }}{{ end }}-`)
	t.AddInternalShortcode("sc2	.html", `sc1: {{ with index .Params 0 }}{{ . }}{{ end }}-- {{ with index .Params 1 }}{{ . }}{{ end }}-`)

}

func Fuzz(data []byte) int {

	score := 0

	// func HandleShortcodes(stringToParse string, page *Page, t tpl.Template) (string, error) {

	p, _ := hugolib.NewPageFrom(strings.NewReader("\n#Hello world"), "foo.md")

	s, err := hugolib.HandleShortcodes(string(data), p, t)

	if err != nil {
		return 0
	}

	if len(s) > 0 {
		score++
	}

	score += strings.Count(s, "HUGOSHORTCODE")

	return score

}
