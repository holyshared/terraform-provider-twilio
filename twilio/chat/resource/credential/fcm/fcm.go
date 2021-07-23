package fcm

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceCredentialService() *schema.Resource {
	return &schema.Resource{
		CreateContext: createContext,
		ReadContext:   readContext,
		UpdateContext: updateContext,
		DeleteContext: deleteContext,
		Schema: map[string]*schema.Schema{
			"friendly_name": {
				Type:     schema.TypeString,
				Computed: false,
				Required: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secret": {
				Type:     schema.TypeString,
				Computed: false,
				Required: true,
			},
			"date_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
