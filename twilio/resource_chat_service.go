package twilio

import (
	"context"
	"time"

	tw "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/chat/v2"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceChatService() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceChatServiceCreate,
		ReadContext:   resourceChatServiceRead,
		UpdateContext: resourceChatServiceUpdate,
		DeleteContext: resourceChatServiceDelete,
		Schema: map[string]*schema.Schema{
			"friendly_name": {
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

func resourceChatServiceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	resourceChatServiceRead(ctx, d, m)

	return diags
}

func resourceChatServiceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tw.RestClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	serviceSid := d.Id()

	res, err := client.ChatV2.FetchService(serviceSid)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("friendly_name", res.FriendlyName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("date_created", res.DateCreated.Format(time.RFC3339)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("date_updated", res.DateUpdated.Format(time.RFC3339)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceChatServiceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tw.RestClient)

	if d.HasChange("friendly_name") {
		friendlyName := d.Get("friendly_name").(string)

		res, err := client.ChatV2.UpdateService(d.Id(), &openapi.UpdateServiceParams{
			FriendlyName: &friendlyName,
		})
		if err != nil {
			return diag.FromErr(err)
		}

		if err := d.Set("date_updated", res.DateUpdated.Format(time.RFC3339)); err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceChatServiceRead(ctx, d, m)
}

func resourceChatServiceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tw.RestClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := client.ChatV2.DeleteService(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
