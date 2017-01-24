package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Hotel struct {
	item.Item

	Name        string `json:"name"`
	Description string `json:"description"`
	BlockCode   string `json:"block_code"`
	RoomRate    int    `json:"room_rate"`
	Phone       string `json:"phone"`
	PhotoURL    string `json:"photo_u_r_l"`
	Onsale      bool   `json:"onsale"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
}

func (h *Hotel) String() string {
	return h.Name
}

type HotelListResult struct {
	Data []Hotel `json:"data"`
}

// MarshalEditor writes a buffer of html to edit a Hotel within the CMS
// and implements editor.Editable
func (h *Hotel) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(h,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Hotel field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", h, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.Textarea("Description", h, map[string]string{
				"label":       "Description",
				"type":        "text",
				"placeholder": "Enter the Description here",
			}),
		},
		editor.Field{
			View: editor.Input("BlockCode", h, map[string]string{
				"label":       "BlockCode",
				"type":        "text",
				"placeholder": "Enter the BlockCode here",
			}),
		},
		editor.Field{
			View: editor.Input("RoomRate", h, map[string]string{
				"label":       "Room Rate",
				"type":        "text",
				"placeholder": "Enter the Room Rate here",
			}),
		},
		editor.Field{
			View: editor.Input("Phone", h, map[string]string{
				"label":       "Phone",
				"type":        "text",
				"placeholder": "Enter the Phone here",
			}),
		},
		editor.Field{
			View: editor.File("PhotoURL", h, map[string]string{
				"label":       "Photo",
				"type":        "text",
				"placeholder": "Enter the Photo here",
			}),
		},
		editor.Field{
			View: editor.Checkbox("Onsale", h, map[string]string{}, map[string]string{
				"true": "On Sale",
			}),
		},
		editor.Field{
			View: editor.Input("Address", h, map[string]string{
				"label":       "Address",
				"type":        "text",
				"placeholder": "Enter the Address here",
			}),
		},
		editor.Field{
			View: editor.Input("City", h, map[string]string{
				"label":       "City",
				"type":        "text",
				"placeholder": "Enter the City here",
			}),
		},
		editor.Field{
			View: editor.Input("State", h, map[string]string{
				"label":       "State",
				"type":        "text",
				"placeholder": "Enter the State here",
			}),
		},
		editor.Field{
			View: editor.Input("Zip", h, map[string]string{
				"label":       "Zip",
				"type":        "text",
				"placeholder": "Enter the Zip here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Hotel editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Hotel"] = func() interface{} { return new(Hotel) }
}
