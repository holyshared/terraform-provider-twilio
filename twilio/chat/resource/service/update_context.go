package service

import (
	"context"
	"time"

	tw "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/chat/v2"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func UpdateContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	return ReadContext(ctx, d, m)
}
