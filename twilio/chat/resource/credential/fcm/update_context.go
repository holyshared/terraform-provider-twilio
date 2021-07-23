package fcm

import (
	"context"
	"time"

	tw "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/chat/v2"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func updateContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tw.RestClient)

	params := &openapi.UpdateCredentialParams{}

	if d.HasChange("friendly_name") {
		if friendlyName, ok := d.Get("friendly_name").(string); ok {
			params.SetFriendlyName(friendlyName)
		}
	}

	if d.HasChange("secret") {
		secret := d.Get("secret").(string)
		params.SetSecret(secret)
	}

	res, err := client.ChatV2.UpdateCredential(d.Id(), params)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("date_updated", res.DateUpdated.Format(time.RFC3339)); err != nil {
		return diag.FromErr(err)
	}

	return readContext(ctx, d, m)
}
