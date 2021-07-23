package service

import (
	"context"
	"time"

	tw "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/chat/v2"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func applyNotificationsToParams(params *openapi.UpdateServiceParams, settings map[string]interface{}) *openapi.UpdateServiceParams {
	if logEnabled, ok := settings["log_enabled"].(bool); ok {
		params.NotificationsLogEnabled = &logEnabled
	}

	if templates, ok := settings["new_message"].([]interface{}); ok {
		settings := templates[0].(map[string]interface{})
		if v, ok := settings["enabled"].(bool); ok {
			params.SetNotificationsNewMessageEnabled(v)
		}
		if v, ok := settings["template"].(string); ok {
			params.SetNotificationsNewMessageTemplate(v)
		}
		if v, ok := settings["sound"].(string); ok {
			params.SetNotificationsNewMessageSound(v)
		}
		if v, ok := settings["badge_count_enabled"].(bool); ok {
			params.SetNotificationsNewMessageBadgeCountEnabled(v)
		}
	}

	if templates, ok := settings["invited_to_channel"].([]interface{}); ok {
		settings := templates[0].(map[string]interface{})
		if v, ok := settings["enabled"].(bool); ok {
			params.SetNotificationsInvitedToChannelEnabled(v)
		}
		if v, ok := settings["template"].(string); ok {
			params.SetNotificationsInvitedToChannelTemplate(v)
		}
		if v, ok := settings["sound"].(string); ok {
			params.SetNotificationsInvitedToChannelSound(v)
		}
	}

	if templates, ok := settings["added_to_channel"].([]interface{}); ok {
		settings := templates[0].(map[string]interface{})
		if v, ok := settings["enabled"].(bool); ok {
			params.SetNotificationsAddedToChannelEnabled(v)
		}
		if v, ok := settings["template"].(string); ok {
			params.SetNotificationsAddedToChannelTemplate(v)
		}
		if v, ok := settings["sound"].(string); ok {
			params.SetNotificationsAddedToChannelSound(v)
		}
	}

	if templates, ok := settings["removed_from_channel"].([]interface{}); ok {
		settings := templates[0].(map[string]interface{})
		if v, ok := settings["enabled"].(bool); ok {
			params.SetNotificationsRemovedFromChannelEnabled(v)
		}
		if v, ok := settings["template"].(string); ok {
			params.SetNotificationsRemovedFromChannelTemplate(v)
		}
		if v, ok := settings["sound"].(string); ok {
			params.SetNotificationsRemovedFromChannelSound(v)
		}
	}

	return params
}

func updateContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tw.RestClient)

	params := &openapi.UpdateServiceParams{}

	if d.HasChange("friendly_name") {
		if friendlyName, ok := d.Get("friendly_name").(string); ok {
			params.SetFriendlyName(friendlyName)
		}
	}

	if d.HasChange("roles") {
		roles := d.Get("roles").([]interface{})
		if len(roles) > 0 {
			settings := roles[0].(map[string]interface{})
			if v, ok := settings["default_service_role"].(string); ok {
				params.SetDefaultServiceRoleSid(v)
			}
			if v, ok := settings["default_channel_role"].(string); ok {
				params.SetDefaultChannelRoleSid(v)
			}
			if v, ok := settings["default_channel_creator_role"].(string); ok {
				params.SetDefaultChannelCreatorRoleSid(v)
			}
		}
	}

	if d.HasChange("limits") {
		limits := d.Get("limits").([]interface{})
		if len(limits) > 0 {
			settings := limits[0].(map[string]interface{})
			if v, ok := settings["channel_members"].(int); ok {
				params.SetLimitsChannelMembers(v)
			}
			if v, ok := settings["user_channels"].(int); ok {
				params.SetLimitsUserChannels(v)
			}
		}
	}

	if d.HasChange("additional_settings") {
		additionalSettings := d.Get("additional_settings").([]interface{})

		if len(additionalSettings) > 0 {
			settings := additionalSettings[0].(map[string]interface{})
			if v, ok := settings["reachability_enabled"].(bool); ok {
				params.SetReachabilityEnabled(v)
			}
			if v, ok := settings["read_status_enabled"].(bool); ok {
				params.SetReadStatusEnabled(v)
			}
			if v, ok := settings["consumption_report_interval"].(int); ok {
				params.SetConsumptionReportInterval(v)
			}
			if v, ok := settings["typing_indicator_timeout"].(int); ok {
				params.SetTypingIndicatorTimeout(v)
			}
			if v, ok := settings["pre_webhook_retry_count"].(int); ok {
				params.SetPreWebhookRetryCount(v)
			}
			if v, ok := settings["post_webhook_retry_count"].(int); ok {
				params.SetPostWebhookRetryCount(v)
			}
		}
	}

	if d.HasChange("webhooks") {
		webhooks := d.Get("webhooks").([]interface{})

		if len(webhooks) > 0 {
			settings := webhooks[0].(map[string]interface{})

			if v, ok := settings["metheventsod"].([]interface{}); ok {
				watchEvents := []string{}
				for _, e := range v {
					watchEvents = append(watchEvents, e.(string))
				}
				params.SetWebhookFilters(watchEvents)
			}

			if v, ok := settings["method"].(string); ok {
				params.SetWebhookMethod(v)
			}
			if v, ok := settings["pre_hook_url"].(string); ok {
				params.SetPreWebhookUrl(v)
			}
			if v, ok := settings["post_hook_url"].(string); ok {
				params.SetPostWebhookUrl(v)
			}
		}
	}

	if d.HasChange("notifications") {
		if notifications, ok := d.Get("notifications").([]interface{}); ok && len(notifications) > 0 {
			settings := notifications[0].(map[string]interface{})
			params = applyNotificationsToParams(params, settings)
		}
	}

	res, err := client.ChatV2.UpdateService(d.Id(), params)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("date_updated", res.DateUpdated.Format(time.RFC3339)); err != nil {
		return diag.FromErr(err)
	}

	return readContext(ctx, d, m)
}
