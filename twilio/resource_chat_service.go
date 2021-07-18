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
			"roles": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Computed: false,
				},
				Optional: true,
				Computed: false,
			},
			"limits": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
					Computed: false,
				},
				Optional: true,
				Computed: false,
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

	roles := map[string]interface{}{
		"default_service_role":         *res.DefaultServiceRoleSid,
		"default_channel_role":         *res.DefaultChannelRoleSid,
		"default_channel_creator_role": *res.DefaultChannelCreatorRoleSid,
	}

	if err := d.Set("roles", roles); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("limits", *res.Limits); err != nil {
		return diag.FromErr(err)
	}

	/*
		// Service Limits
		//	res.Limits[]
		//	"channel_members": 100,
		//	"user_channels": 250

		// Additional Settings
		res.ReachabilityEnabled
		res.ReadStatusEnabled
		res.ConsumptionReportInterval
		res.TypingIndicatorTimeout
		res.PreWebhookRetryCount
		res.PostWebhookRetryCount
	*/

	/*
		res.WebhookFilters
		res.WebhookMethod
		res.PreWebhookUrl
		res.PostWebhookUrl
	*/

	return diags
}

func resourceChatServiceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tw.RestClient)

	params := &openapi.UpdateServiceParams{}

	if d.HasChange("friendly_name") {
		friendlyName := d.Get("friendly_name").(string)
		params.FriendlyName = &friendlyName
	}

	if d.HasChange("roles") {
		roles := d.Get("roles").(map[string]interface{})

		_, hasServiceRole := roles["default_service_role"]
		if hasServiceRole {
			params.DefaultServiceRoleSid = roles["default_service_role"].(*string)
		}

		_, hasChannelRole := roles["default_channel_role"]
		if hasChannelRole {
			params.DefaultChannelRoleSid = roles["default_channel_role"].(*string)
		}

		_, hasChannelCreatorRole := roles["default_channel_creator_role"]
		if hasChannelCreatorRole {
			params.DefaultChannelCreatorRoleSid = roles["default_channel_creator_role"].(*string)
		}
	}

	if d.HasChange("limits") {
		limits := d.Get("limits").(map[string]interface{})
		_, hasChannelMembers := limits["channel_members"]
		if hasChannelMembers {
			params.LimitsChannelMembers = limits["channel_members"].(*int)
		}

		_, hasUserChannels := limits["user_channels"]
		if hasUserChannels {
			params.LimitsUserChannels = limits["user_channels"].(*int)
		}
	}

	res, err := client.ChatV2.UpdateService(d.Id(), params)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("date_updated", res.DateUpdated.Format(time.RFC3339)); err != nil {
		return diag.FromErr(err)
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
