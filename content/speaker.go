package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Speaker struct {
	item.Item

	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Title         string `json:"title"`
	Company       string `json:"company"`
	Bio           string `json:"bio"`
	Website       string `json:"website"`
	PhotoURL      string `json:"photo_u_r_l"`
	Twitter       string `json:"twitter"`
	Linkedin      string `json:"linkedin"`
	Facebook      string `json:"facebook"`
	Googleplus    string `json:"googleplus"`
	Github        string `json:"github"`
	ContactPhone  string `json:"contact_phone"`
	ContactEmail  string `json:"contact_email"`
	Address       string `json:"address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Zip           string `json:"zip"`
	Country       string `json:"country"`
	TalkType      string `json:"talk_type"`
	Presentations []int  `json:"presentations"`
}

func (s *Speaker) String() string {
	return s.FirstName + " " + s.LastName
}

type SpeakerListResult struct {
	Data []Speaker `json:"data"`
}

// MarshalEditor writes a buffer of html to edit a Speaker within the CMS
// and implements editor.Editable
func (s *Speaker) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(s,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Speaker field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("FirstName", s, map[string]string{
				"label":       "First Name",
				"type":        "text",
				"placeholder": "Enter the First Name here",
			}),
		},
		editor.Field{
			View: editor.Input("LastName", s, map[string]string{
				"label":       "Last Name",
				"type":        "text",
				"placeholder": "Enter the Last Name here",
			}),
		},
		editor.Field{
			View: editor.Textarea("Bio", s, map[string]string{
				"label":       "Bio",
				"type":        "text",
				"placeholder": "Enter the Bio here",
			}),
		},
		editor.Field{
			View: editor.Input("Website", s, map[string]string{
				"label":       "Website",
				"type":        "text",
				"placeholder": "Enter the speaker's website here",
			}),
		},
		editor.Field{
			View: editor.File("PhotoURL", s, map[string]string{
				"label":       "Photo",
				"type":        "text",
				"placeholder": "Enter the PhotoURL here",
			}),
		},
		editor.Field{
			View: editor.Input("Twitter", s, map[string]string{
				"label":       "Twitter",
				"type":        "text",
				"placeholder": "Enter the Twitter URL here, no",
			}),
		},
		editor.Field{
			View: editor.Input("Linkedin", s, map[string]string{
				"label":       "Linkedin",
				"type":        "text",
				"placeholder": "Enter the Linkedin URL here",
			}),
		},
		editor.Field{
			View: editor.Input("Facebook", s, map[string]string{
				"label":       "Facebook",
				"type":        "text",
				"placeholder": "Enter the Facebook URL here",
			}),
		},
		editor.Field{
			View: editor.Input("Googleplus", s, map[string]string{
				"label":       "Googleplus",
				"type":        "text",
				"placeholder": "Enter the Googleplus url here",
			}),
		},

		editor.Field{
			View: editor.Input("Github", s, map[string]string{
				"label":       "Github",
				"type":        "text",
				"placeholder": "Enter the Github url here",
			}),
		},
		editor.Field{
			View: editor.Input("ContactPhone", s, map[string]string{
				"label":       "Contact Phone",
				"type":        "text",
				"placeholder": "Enter the Contact Phone here",
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
			View: editor.Input("Address", s, map[string]string{
				"label":       "Address",
				"type":        "text",
				"placeholder": "Enter the Address here",
			}),
		},
		editor.Field{
			View: editor.Input("City", s, map[string]string{
				"label":       "City",
				"type":        "text",
				"placeholder": "Enter the City here",
			}),
		},
		editor.Field{
			View: editor.Input("State", s, map[string]string{
				"label":       "State",
				"type":        "text",
				"placeholder": "Enter the State here",
			}),
		},
		editor.Field{
			View: editor.Input("Zip", s, map[string]string{
				"label":       "Zip",
				"type":        "text",
				"placeholder": "Enter the Zip here",
			}),
		},
		editor.Field{
			View: editor.Input("Country", s, map[string]string{
				"label":       "Country",
				"type":        "text",
				"placeholder": "Enter the Country here",
			}),
		},
		editor.Field{
			View: editor.Select("TalkType", s, map[string]string{
				"label":       "TalkType",
				"type":        "text",
				"placeholder": "Enter the Talk Type here",
			}, map[string]string{
				"Keynote":  "Keynote",
				"Tutorial": "Tutorial",
				"Workshop": "Workshop",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Speaker editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Speaker"] = func() interface{} { return new(Speaker) }
}
