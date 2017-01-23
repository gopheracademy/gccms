package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Event struct {
	item.Item

	Name             string `json:"name"`
	EventSlug        string `json:"event_slug"`
	ShortDescription string `json:"short_description"`
	LongDescription  string `json:"long_description"`
	Date             int64  `json:"date"`
	StartTime        int64  `json:"start_time"`
	EndTime          int64  `json:"end_time"`
	MapURL           string `json:"map_u_r_l"`
	Address          string `json:"address"`
	City             string `json:"city"`
	State            string `json:"state"`
	Zip              string `json:"zip"`
	Phone            string `json:"phone"`
	Email            string `json:"email"`
	PhotoURL         string `json:"photo_url"`
}

func (e *Event) String() string {
	return e.EventSlug
}

type EventListResult struct {
	Data []Event `json:"data"`
}

func (e *Event) ItemSlug() string {
	return e.EventSlug
}

// MarshalEditor writes a buffer of html to edit a Event within the CMS
// and implements editor.Editable
func (e *Event) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(e,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Event field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("EventSlug", e, map[string]string{
				"label":       "EventSlug",
				"type":        "text",
				"placeholder": "Enter the event slug: eg /events/{{slug}}",
			}),
		},
		editor.Field{
			View: editor.Input("Name", e, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.Input("ShortDescription", e, map[string]string{
				"label":       "Short Description",
				"type":        "text",
				"placeholder": "Enter the Short Description here",
			}),
		},
		editor.Field{
			View: editor.Textarea("LongDescription", e, map[string]string{
				"label":       "Long Description",
				"type":        "text",
				"placeholder": "Enter the Long Description here",
			}),
		},
		editor.Field{
			View: editor.Timestamp("Date", e, map[string]string{
				"label":       "Date",
				"type":        "text",
				"placeholder": "Enter the Date here",
			}),
		},
		editor.Field{
			View: editor.Timestamp("StartTime", e, map[string]string{
				"label":       "StartTime",
				"type":        "text",
				"placeholder": "Enter the StartTime here",
			}),
		},
		editor.Field{
			View: editor.Timestamp("EndTime", e, map[string]string{
				"label":       "EndTime",
				"type":        "text",
				"placeholder": "Enter the EndTime here",
			}),
		},
		editor.Field{
			View: editor.Input("MapURL", e, map[string]string{
				"label":       "MapURL",
				"type":        "text",
				"placeholder": "Enter the MapURL here",
			}),
		},
		editor.Field{
			View: editor.Input("Address", e, map[string]string{
				"label":       "Address",
				"type":        "text",
				"placeholder": "Enter the Address here",
			}),
		},
		editor.Field{
			View: editor.Input("City", e, map[string]string{
				"label":       "City",
				"type":        "text",
				"placeholder": "Enter the City here",
			}),
		},
		editor.Field{
			View: editor.Input("State", e, map[string]string{
				"label":       "State",
				"type":        "text",
				"placeholder": "Enter the State here",
			}),
		},
		editor.Field{
			View: editor.Input("Zip", e, map[string]string{
				"label":       "Zip",
				"type":        "text",
				"placeholder": "Enter the Zip here",
			}),
		},
		editor.Field{
			View: editor.Input("Phone", e, map[string]string{
				"label":       "Phone",
				"type":        "text",
				"placeholder": "Enter the Phone here",
			}),
		},
		editor.Field{
			View: editor.Input("Email", e, map[string]string{
				"label":       "Email",
				"type":        "text",
				"placeholder": "Enter the Email here",
			}),
		},
		editor.Field{
			View: editor.File("PhotoURL", e, map[string]string{
				"label":       "Photo",
				"type":        "text",
				"placeholder": "Enter the Photo here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Event editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Event"] = func() interface{} { return new(Event) }
}
