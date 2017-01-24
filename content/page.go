package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Page struct {
	item.Item
	Address  string `json:"address"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Content  string `json:"content"`
}
type PageListResult struct {
	Data []Page `json:"data"`
}

func (p *Page) String() string {
	return p.Address
}

func (p *Page) ItemSlug() string {
	return p.Address
}

// MarshalEditor writes a buffer of html to edit a Page within the CMS
// and implements editor.Editable
func (p *Page) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(p,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Page field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Address", p, map[string]string{
				"label":       "Address",
				"type":        "text",
				"placeholder": "Enter the page slug:  eg: /about/{{coc}}",
			}),
		},
		editor.Field{
			View: editor.Input("Title", p, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.Input("Subtitle", p, map[string]string{
				"label":       "Subtitle",
				"type":        "text",
				"placeholder": "Enter the Subtitle here",
			}),
		},
		editor.Field{
			View: editor.Textarea("Content", p, map[string]string{
				"label":       "Content",
				"type":        "text",
				"placeholder": "Enter the Content here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Page editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Page"] = func() interface{} { return new(Page) }
}
