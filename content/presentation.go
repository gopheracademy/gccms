package content

import (
	"fmt"

	"github.com/bosssauce/reference"
	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Presentation struct {
	item.Item

	Title            string   `json:"title"`
	ShortDescription string   `json:"short_description"`
	LongDescription  string   `json:"long_description"`
	Speakers         []string `json:"speakers"`
	DisplayOrder     int      `json:"display_order"`
	Day		int	`json:"day"`
	Slot		string 	`json:"slot"`
	Track 		string `json:"track"`

}

func (p *Presentation) String() string {
	return p.Title
}

type PresentationListResult struct {
	Data []Presentation `json:"data"`
}

// MarshalEditor writes a buffer of html to edit a Presentation within the CMS
// and implements editor.Editable
func (p *Presentation) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(p,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Presentation field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Title", p, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.Input("ShortDescription", p, map[string]string{
				"label":       "Short Description",
				"type":        "text",
				"placeholder": "Enter the Short Description here",
			}),
		},
		editor.Field{
			View: editor.Textarea("LongDescription", p, map[string]string{
				"label":       "Long Description",
				"type":        "text",
				"placeholder": "Enter the Long Description here",
			}),
		},

		editor.Field{
			View: editor.Input("Day", p, map[string]string{
				"label":       "Day",
				"type":        "int",
				"placeholder": "Enter the Day (1 or 2)",
			}),
		},
		editor.Field{
			View: editor.Input("DisplayOrder", p, map[string]string{
				"label":       "Display Order",
				"type":        "int",
				"placeholder": "Enter the Display Order here",
			}),
		},
		editor.Field{
    			View: reference.Select("Slot", p, map[string]string{
        		"label": "Select Slot",
    			}, "Slot", `{{.description}}`),
		},
editor.Field{
    View: editor.Select("Track", p, map[string]string{
        "label": "Select Track",
    }, map[string]string{
        // "value": "Display Name",
        "Main":     "Mile High Ballroom",
        "MHB1a":    "MHB1a",
        "MHB1d": "MHB1d",
        "MBH2a":     "MHB2a",
    }),
},
		editor.Field{
			//	View: editor.Input("Lessons", m, map[string]string{
			//		"label":       "Lessons",
			//		"type":        "text",
			//		"placeholder": "Enter the Lessons here",
			//	}),
			View: reference.SelectRepeater("Speakers", p, map[string]string{
				"label":       "Speaker",
				"type":        "text",
				"placeholder": "Enter the Speaker here",
			},
				"Speaker",
				"{{ .first_name }} {{ .last_name }}",
			),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Presentation editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Presentation"] = func() interface{} { return new(Presentation) }
}
