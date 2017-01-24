package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Sponsor struct {
	item.Item

	Name         string `json:"name"`
	Description  string `json:"description"`
	Website      string `json:"website"`
	LogoURL      string `json:"logo_u_r_l"`
	ContactName  string `json:"contact_name"`
	ContactEmail string `json:"contact_email"`
	ContactPhone string `json:"contact_phone"`
	Level        string `json:"level"`
}

func (s *Sponsor) String() string {
	return s.Level + ": " + s.Name
}

type SponsorListResult struct {
	Data []Sponsor `json:"data"`
}

// MarshalEditor writes a buffer of html to edit a Sponsor within the CMS
// and implements editor.Editable
func (s *Sponsor) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(s,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Sponsor field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", s, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.Textarea("Description", s, map[string]string{
				"label":       "Description",
				"type":        "text",
				"placeholder": "Enter the Description here",
			}),
		},
		editor.Field{
			View: editor.Input("Website", s, map[string]string{
				"label":       "Website",
				"type":        "text",
				"placeholder": "Enter the Website here",
			}),
		},
		editor.Field{
			View: editor.File("LogoURL", s, map[string]string{
				"label":       "Logo",
				"type":        "text",
				"placeholder": "Enter the Logo here",
			}),
		},
		editor.Field{
			View: editor.Input("ContactName", s, map[string]string{
				"label":       "Contact Name",
				"type":        "text",
				"placeholder": "Enter the Contact Name here",
			}),
		},
		editor.Field{
			View: editor.Input("ContactEmail", s, map[string]string{
				"label":       "Contact Email",
				"type":        "text",
				"placeholder": "Enter the Contact Email here",
			}),
		},
		editor.Field{
			View: editor.Input("ContactPhone", s, map[string]string{
				"label":       "ContactPhone",
				"type":        "text",
				"placeholder": "Enter the Contact Phone here",
			}),
		},
		editor.Field{
			View: editor.Select("Level", s, map[string]string{
				"label":       "Level",
				"type":        "text",
				"placeholder": "Enter the Level here",
			}, map[string]string{
				"Diamond":  "Diamond",
				"Platinum": "Platinum",
				"Gold":     "Gold",
				"Silver":   "Silver",
				"Bronze":   "Bronze",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Sponsor editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Sponsor"] = func() interface{} { return new(Sponsor) }
}
