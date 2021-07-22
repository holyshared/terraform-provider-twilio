package twilio

import (
	"context"
	"fmt"
	"time"

	tw "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/chat/v2"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var webHookEvents = []string{
	"onMessageSend",
	"onMessageUpdate",
	"onMessageRemove",
	"onMediaMessageSend",
	"onChannelAdd",
	"onChannelUpdate",
	"onChannelDestroy",
	"onMemberAdd",
	"onMemberUpdate",
	"onMemberRemove",
	"onUserUpdate",
	"onMessageSent",
	"onMessageUpdated",
	"onMessageRemoved",
	"onMediaMessageSent",
	"onChannelAdded",
	"onChannelUpdated",
	"onChannelDestroyed",
	"onMemberAdded",
	"onMemberUpdated",
	"onMemberRemoved",
	"onUserAdded",
	"onUserUpdated",
}

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
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_channel_creator_role": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: false,
						},
						"default_channel_role": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: false,
						},
						"default_service_role": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: false,
						},
					},
				},
				Optional: true,
				Computed: false,
				MaxItems: 1,
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
			"webhooks": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"events": {
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Optional: true,
							Computed: false,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								warns, errs = validation.ListOfUniqueStrings(val, key)

								v, ok := val.([]interface{})
								if !ok { // type error
									errs = append(errs, fmt.Errorf("expected type of %q to be List", key))
									return warns, errs
								}

								for _, e := range v {
									if _, eok := e.(string); !eok {
										errs = append(errs, fmt.Errorf("expected %q to only contain string elements, found :%v", key, e))
										return warns, errs
									}
								}

								for _, sv := range v {
									find := false
									for _, tv := range webHookEvents {
										if sv.(string) == tv {
											find = true
											break
										}
										if !find {
											errs = append(errs, fmt.Errorf("expected %q to event names, found %v", key, sv))
											return warns, errs
										}
									}
								}

								return warns, errs
							},
						},
						"method": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     false,
							ValidateFunc: validation.StringInSlice([]string{"GET", "POST"}, false),
						},
						"pre_hook_url": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     false,
							ValidateFunc: validation.IsURLWithHTTPorHTTPS,
						},
						"post_hook_url": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     false,
							ValidateFunc: validation.IsURLWithHTTPorHTTPS,
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

	roles := []map[string]interface{}{{
		"default_service_role":         *res.DefaultServiceRoleSid,
		"default_channel_role":         *res.DefaultChannelRoleSid,
		"default_channel_creator_role": *res.DefaultChannelCreatorRoleSid,
	}}

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

	webhook := map[string]interface{}{}
	if res.WebhookFilters != nil {
		webhook["events"] = *res.WebhookFilters
	}
	if res.WebhookMethod != nil {
		webhook["method"] = *res.WebhookMethod
	}
	if res.PreWebhookUrl != nil {
		webhook["pre_hook_url"] = *res.PreWebhookUrl
	}
	if res.PostWebhookUrl != nil {
		webhook["post_hook_url"] = *res.PostWebhookUrl
	}

	webhooks := []map[string]interface{}{webhook}
	if err := d.Set("webhooks", webhooks); err != nil {
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
		roles := d.Get("roles").([]interface{})
		if len(roles) > 0 {
			settings := roles[0].(map[string]interface{})

			serviceRole, hasServiceRole := settings["default_service_role"]
			if hasServiceRole {
				val := serviceRole.(string)
				params.DefaultServiceRoleSid = &val
			}

			channelRole, hasChannelRole := settings["default_channel_role"]
			if hasChannelRole {
				val := channelRole.(string)
				params.DefaultChannelRoleSid = &val
			}

			channelCreator, hasChannelCreatorRole := settings["default_channel_creator_role"]
			if hasChannelCreatorRole {
				val := channelCreator.(string)
				params.DefaultChannelCreatorRoleSid = &val
			}
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

	if d.HasChange("webhooks") {
		webhooks := d.Get("webhooks").([]interface{})

		if len(webhooks) > 0 {
			settings := webhooks[0].(map[string]interface{})

			events, hasEvents := settings["events"]
			if hasEvents {
				val := events.([]interface{})
				watchEvents := []string{}
				for _, e := range val {
					watchEvents = append(watchEvents, e.(string))
				}
				params.WebhookFilters = &watchEvents
			}

			method, hasMethod := settings["method"]
			if hasMethod {
				val := method.(string)
				params.WebhookMethod = &val
			}

			preHookURL, hasPreHookURL := settings["pre_hook_url"]
			if hasPreHookURL {
				val := preHookURL.(string)
				params.PreWebhookUrl = &val
			}

			postHookURL, hasPostHookURL := settings["post_hook_url"]
			if hasPostHookURL {
				val := postHookURL.(string)
				params.PostWebhookUrl = &val
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
