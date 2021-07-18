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
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"channel_members": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: false,
						},
						"user_channels": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: false,
						},
					},
				},
				Optional: true,
				Computed: false,
				MaxItems: 1,
			},
			"additional_settings": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"reachability_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: false,
						},
						"read_status_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: false,
						},
						"consumption_report_interval": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: false,
						},
						"typing_indicator_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: false,
						},
						"pre_webhook_retry_count": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: false,
						},
						"post_webhook_retry_count": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: false,
						},
					},
				},
				Optional: true,
				Computed: false,
				MaxItems: 1,
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

	limits := []map[string]interface{}{{
		"channel_members": (*res.Limits)["channel_members"],
		"user_channels":   (*res.Limits)["user_channels"],
	}}

	if err := d.Set("limits", limits); err != nil {
		return diag.FromErr(err)
	}

	settigns := []map[string]interface{}{{
		"reachability_enabled":        *res.ReachabilityEnabled,
		"read_status_enabled":         *res.ReadStatusEnabled,
		"consumption_report_interval": *res.ConsumptionReportInterval,
		"typing_indicator_timeout":    *res.TypingIndicatorTimeout,
		"pre_webhook_retry_count":     *res.PreWebhookRetryCount,
		"post_webhook_retry_count":    *res.PostWebhookRetryCount,
	}}

	if err := d.Set("additional_settings", settigns); err != nil {
		return diag.FromErr(err)
	}

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
		limits := d.Get("limits").([]interface{})
		if len(limits) > 0 {
			settings := limits[0].(map[string]interface{})

			channelMembers, hasChannelMembers := settings["channel_members"]
			if hasChannelMembers {
				val := channelMembers.(int)
				params.LimitsChannelMembers = &val
			}

			userChannels, hasUserChannels := settings["user_channels"]
			if hasUserChannels {
				val := userChannels.(int)
				params.LimitsUserChannels = &val
			}
		}
	}

	if d.HasChange("additional_settings") {
		additionalSettings := d.Get("additional_settings").([]interface{})

		if len(additionalSettings) > 0 {
			settings := additionalSettings[0].(map[string]interface{})

			reachabilityEnabled, hasReachabilityEnabled := settings["reachability_enabled"]
			if hasReachabilityEnabled {
				val := reachabilityEnabled.(bool)
				params.ReachabilityEnabled = &val
			}

			readStatusEnabled, hasReadStatusEnabled := settings["read_status_enabled"]
			if hasReadStatusEnabled {
				val := readStatusEnabled.(bool)
				params.ReadStatusEnabled = &val
			}

			consumptionReportInterval, hasConsumptionReportInterval := settings["consumption_report_interval"]
			if hasConsumptionReportInterval {
				val := consumptionReportInterval.(int)
				params.ConsumptionReportInterval = &val
			}

			typingIndicatorTimeout, hasTypingIndicatorTimeout := settings["typing_indicator_timeout"]
			if hasTypingIndicatorTimeout {
				val := typingIndicatorTimeout.(int)
				params.TypingIndicatorTimeout = &val
			}

			preWebhookRetryCount, hasPreWebhookRetryCount := settings["pre_webhook_retry_count"]
			if hasPreWebhookRetryCount {
				val := preWebhookRetryCount.(int)
				params.PreWebhookRetryCount = &val
			}

			postWebhookRetryCount, hasPostWebhookRetryCount := settings["post_webhook_retry_count"]
			if hasPostWebhookRetryCount {
				val := postWebhookRetryCount.(int)
				params.PostWebhookRetryCount = &val
			}
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
