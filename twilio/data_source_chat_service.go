package twilio

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tw "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/chat/v2"
)

func dataSourceChatServicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tw.RestClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	res, err := client.ChatV2.ListService(&openapi.ListServiceParams{})

	if err != nil {
		return diag.FromErr(err)
	}

	chatServices := make([]map[string]interface{}, 0)

	for _, s := range res.Services {
		chatServices = append(chatServices, map[string]interface{}{
			"sid":          s.Sid,
			"friendlyName": s.FriendlyName,
			"dateCreated":  s.DateCreated.Format("2006-01-02 15:04:05"),
		})
	}

	if err := d.Set("chat_services", chatServices); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(client.Client.AccountSid())

	return diags
}

/*
[
	{
		sid: string
		friendlyName: string
		dateCreated: stirng
	}
]
*/
func dataSourceChatServices() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceChatServicesRead,
		Schema: map[string]*schema.Schema{
			"chat_services": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"friendlyName": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dateCreated": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}
