package service

import (
	"context"

	tw "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/chat/v2"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CreateContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tw.RestClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	friendlyName := d.Get("friendly_name").(string)

	res, err := client.ChatV2.CreateService(&openapi.CreateServiceParams{
		FriendlyName: &friendlyName,
	})

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*res.Sid)

	ReadContext(ctx, d, m)

	return diags
}
