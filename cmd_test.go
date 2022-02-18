package main

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestGetDescriptionLookup(t *testing.T) {
	t.Run("getDescriptionLookup", func(t *testing.T) {
		tables := map[string]Table{
			"p1": []Column{
				{
					Description: "d1",
					Mode:        "m1",
					Name:        "n1",
					Type:        "t1",
				},
				{
					Description: "",
					Mode:        "m2",
					Name:        "n2",
					Type:        "t2",
				},
			},
			"p2": []Column{
				{
					Description: "",
					Mode:        "m1",
					Name:        "n1",
					Type:        "t1",
				},
				{
					Description: "",
					Mode:        "m2",
					Name:        "n2",
					Type:        "t2",
				},
			},
		}
		expected := fmt.Sprint(map[string]string{
			"n1": "d1"})
		result := fmt.Sprint(getDescriptionLookup(tables))
		assert.Equal(t, result, expected, "Gets correct description lookup")
		if result != expected {
			t.Errorf("%s did not match %s", result, expected)
		}
	})
}
func TestFillinDescriptions(t *testing.T) {
	t.Run("fillInDescriptions", func(t *testing.T) {
		tables := map[string]Table{
			"p1": []Column{
				{
					Description: "d1",
					Mode:        "m1",
					Name:        "n1",
					Type:        "t1",
				},
				{
					Description: "",
					Mode:        "m2",
					Name:        "n2",
					Type:        "t2",
				},
			},
			"p2": []Column{
				{
					Description: "",
					Mode:        "m1",
					Name:        "n1",
					Type:        "t1",
				},
				{
					Description: "",
					Mode:        "m2",
					Name:        "n2",
					Type:        "t2",
				},
			},
		}
		expected := fmt.Sprint(map[string]Table{
			"p2": []Column{
				{
					Description: "d1",
					Mode:        "m1",
					Name:        "n1",
					Type:        "t1",
				},
				{
					Description: "",
					Mode:        "m2",
					Name:        "n2",
					Type:        "t2",
				},
			},
		})
		descriptionLookup := map[string]string{
			"n1": "d1",
		}
		result := fmt.Sprint(fillInDescriptions(tables, descriptionLookup))
		assert.Equal(t, result, expected, "Fills in descriptions correctly")
		if result != expected {
			t.Errorf("%s did not match %s", result, expected)
		}
	})
}
