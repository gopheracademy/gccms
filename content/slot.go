package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Slot struct {
	item.Item

	Number int `json:"number"`
	Description string   `json:"description"`
}

func (p *Slot) String() string {
	return p.Description
}

type SlotListResult struct {
	Data []Slot `json:"data"`
}

// MarshalEditor writes a buffer of html to edit a Slot within the CMS
// and implements editor.Editable
func (p *Slot) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(p,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Slot field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Number", p, map[string]string{
				"label":       "Number",
				"type":        "integer",
				"placeholder": "Enter the Number here",
			}),
		},
		editor.Field{
			View: editor.Input("Description", p, map[string]string{
				"label":       "Description",
				"type":        "text",
				"placeholder": "Enter the Short Description here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Slot editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Slot"] = func() interface{} { return new(Slot) }
}
