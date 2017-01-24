package ns1

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func trialResource() *schema.Resource {
	s := map[string]*schema.Schema{
		"a": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			ConflictsWith: []string{"b"},
		},
		"b": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			ConflictsWith: []string{"a"},
		},
		"first": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem:     &schema.Resource{
				Schema: map[string]*schema.Schema{
					"x": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
						ConflictsWith: []string{"first.x"},
					},
					"y": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
						ConflictsWith: []string{"first.y"},
					},
				},
			},
		},
		"second": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem:     &schema.Resource{
				Schema: map[string]*schema.Schema{
					"x": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
						ConflictsWith: []string{"second.x"},
					},
					"y": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
						ConflictsWith: []string{"second.y"},
					},
				},
			},
		},
	}
	s = addPermsSchema(s)
	return &schema.Resource{
		Schema: s,
		Read:   trialRead,
		Importer: &schema.ResourceImporter{State: trialStateFunc},
	}
}

func trialRead(d *schema.ResourceData, meta interface{}) error {
	d.Set("a", "aaa")
	d.Set("first", []interface{}{
		map[string]interface{}{
			"x": "xxx",
		},
	})
	d.Set("second", []interface{}{
		map[string]interface{}{
			"y": "yyy",
		},
	})
	return nil
}

func trialStateFunc(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return []*schema.ResourceData{d}, nil
}
