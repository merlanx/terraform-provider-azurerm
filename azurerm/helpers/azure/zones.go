package azure

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/validate"
)

func SchemaZones() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		ForceNew: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validate.NoEmptyStrings,
		},
	}
}

func SchemaSingleZone() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		ForceNew: true,
		MaxItems: 1,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validate.NoEmptyStrings,
		},
	}
}

func SchemaMultipleZones() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		ForceNew: true,
		MinItems: 1,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validate.NoEmptyStrings,
		},
	}
}

func SchemaZonesComputed() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validate.NoEmptyStrings,
		},
	}
}

func ExpandZones(v []interface{}) *[]string {
	zones := make([]string, 0)
	for _, zone := range v {
		zones = append(zones, zone.(string))
	}
	if len(zones) > 0 {
		return &zones
	} else {
		return nil
	}
}

func FlattenZones(v *[]string) []interface{} {
	zones := make([]interface{}, 0)
	if v == nil {
		return zones
	}

	for _, s := range *v {
		zones = append(zones, s)
	}
	return zones
}
