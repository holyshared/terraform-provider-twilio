package fcm

import (
	"context"

	tw "github.com/twilio/twilio-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func deleteContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tw.RestClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := client.ChatV2.DeleteCredential(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
