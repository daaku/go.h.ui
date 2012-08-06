package ui

import (
	"crypto/rand"
	"fmt"
	"github.com/daaku/go.h"
	"io"
	"log"
)

func MakeID(prefix string) string {
	b := make([]byte, 8)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		log.Fatal(err)
	}
	if prefix == "" {
		return fmt.Sprintf("%x", b)
	}
	return fmt.Sprintf("%s_%x", prefix, b)
}

type Checkbox struct {
	Label       h.HTML
	Name        string
	Checked     bool
	Description h.HTML
	Value       interface{}
}

func (i *Checkbox) HTML() (h.HTML, error) {
	value := i.Value
	if value == nil {
		value = "1"
	}
	id := MakeID(i.Name)
	return &h.Div{
		Class: "control-group",
		Inner: &h.Frag{
			&h.Label{
				Class: "control-label",
				For:   id,
				Inner: i.Label,
			},
			&h.Div{
				Class: "controls",
				Inner: &h.Label{
					Class: "checkbox",
					Inner: &h.Frag{
						&h.Input{
							Type:    "checkbox",
							ID:      id,
							Name:    i.Name,
							Checked: i.Checked,
							Value:   fmt.Sprint(value),
						},
						i.Description,
					},
				},
			},
		},
	}, nil
}

type ToggleGroup struct {
	Label h.HTML
	Inner h.HTML
}

func (g *ToggleGroup) HTML() (h.HTML, error) {
	return &h.Table{
		Class: "table table-striped table-condensed",
		Inner: &h.Frag{
			&h.Thead{
				Inner: &h.Tr{
					Inner: &h.Frag{
						&h.Th{
							Inner: h.String("On"),
						},
						&h.Th{
							Inner: h.String("Off"),
						},
						&h.Th{
							Inner: g.Label,
						},
					},
				},
			},
			&h.Tbody{
				Inner: g.Inner,
			},
		},
	}, nil
}

type ToggleItem struct {
	Name        string
	Checked     bool
	Description h.HTML
}

func (i *ToggleItem) HTML() (h.HTML, error) {
	return &h.Tr{
		Inner: &h.Frag{
			&h.Td{
				Inner: &h.Input{
					Class:   "radio",
					Type:    "radio",
					Checked: i.Checked,
					Name:    i.Name,
					Value:   "1",
				},
			},
			&h.Td{
				Inner: &h.Input{
					Class:   "radio",
					Type:    "radio",
					Checked: !i.Checked,
					Name:    i.Name,
					Value:   "0",
				},
			},
			&h.Td{
				Inner: i.Description,
			},
		},
	}, nil
}

type TextInput struct {
	Type  string
	Label h.HTML
	Name  string
	Value interface{}
}

func (i *TextInput) HTML() (h.HTML, error) {
	t := i.Type
	if t == "" {
		t = "text"
	}
	id := MakeID(i.Name)
	return &h.Div{
		Class: "control-group",
		Inner: &h.Frag{
			&h.Label{
				Class: "control-label",
				For:   id,
				Inner: i.Label,
			},
			&h.Div{
				Class: "controls",
				Inner: &h.Frag{
					&h.Input{
						Type:  t,
						ID:    id,
						Name:  i.Name,
						Value: fmt.Sprint(i.Value),
					},
				},
			},
		},
	}, nil
}
