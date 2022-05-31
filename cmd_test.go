package main

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestGetDescriptionLookup(t *testing.T) {
	desc1 := "d1"
	mode1 := "m1"
	mode2 := "m2"
	emptyString := ""
	t.Run("getDescriptionLookup", func(t *testing.T) {
		tables := map[string]Table{
			"p1": []Column{
				{
					Description: &desc1,
					Mode:        &mode1,
					Name:        "n1",
					Type:        "t1",
				},
				{
					Description: &emptyString,
					Mode:        &mode2,
					Name:        "n2",
					Type:        "t2",
				},
			},
			"p2": []Column{
				{
					Description: &emptyString,
					Mode:        &mode1,
					Name:        "n1",
					Type:        "t1",
				},
				{
					Description: &emptyString,
					Mode:        &mode2,
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
	desc1 := "d1"
	mode1 := "m1"
	mode2 := "m2"
	emptyString := ""

	t.Run("fillInDescriptions", func(t *testing.T) {
		tables := map[string]Table{
			"p1": []Column{
				{
					Description: &desc1,
					Mode:        &mode1,
					Name:        "n1",
					Type:        "t1",
				},
				{
					Description: &emptyString,
					Mode:        &mode2,
					Name:        "n2",
					Type:        "t2",
				},
			},
			"p2": []Column{
				{
					Description: &emptyString,
					Mode:        &mode1,
					Name:        "n1",
					Type:        "t1",
				},
				{
					Description: &emptyString,
					Mode:        &mode2,
					Name:        "n2",
					Type:        "t2",
				},
			},
		}
		expected := map[string]Table{
			"p2": []Column{
				{
					Description: &desc1,
					Mode:        &mode1,
					Name:        "n1",
					Type:        "t1",
				},
				{
					Description: &emptyString,
					Mode:        &mode2,
					Name:        "n2",
					Type:        "t2",
				},
			},
		}
		descriptionLookup := map[string]*string{
			"n1": &desc1,
		}
		result := fillInDescriptions(tables, descriptionLookup)
		assert.Equal(t, result["n1"], expected["n1"], "Fills in descriptions correctly")
		assert.Equal(t, len(result), 2, "Fills in descriptions correctly")
	})
}
