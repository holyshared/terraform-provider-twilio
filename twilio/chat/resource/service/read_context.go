package service

import (
	"context"
	"time"

	tw "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/chat/v2"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var pushNotificationTemplateNames = []string{"new_message", "invited_to_channel", "added_to_channel", "removed_from_channel"}

func rolesFromResponse(cs *openapi.ChatV2Service) *map[string]interface{} {
	setting := map[string]interface{}{}

	if cs.DefaultServiceRoleSid != nil {
		setting["default_service_role"] = *cs.DefaultServiceRoleSid
	}
	if cs.DefaultChannelRoleSid != nil {
		setting["default_channel_role"] = *cs.DefaultChannelRoleSid
	}
	if cs.DefaultChannelCreatorRoleSid != nil {
		setting["default_channel_creator_role"] = *cs.DefaultChannelCreatorRoleSid
	}

	if len(setting) <= 0 {
		return nil
	}

	return &setting
}

func limitsFromResponse(limits map[string]interface{}) map[string]interface{} {
	setting := map[string]interface{}{}

	if v, ok := limits["channel_members"]; ok {
		setting["channel_members"] = v
	}
	if v, ok := limits["user_channels"]; ok {
		setting["user_channels"] = v
	}

	return setting
}

func additionalSettingsFromResponse(cs *openapi.ChatV2Service) *map[string]interface{} {
	setting := map[string]interface{}{}

	if cs.ReachabilityEnabled != nil {
		setting["reachability_enabled"] = *cs.ReachabilityEnabled
	}
	if cs.ReadStatusEnabled != nil {
		setting["read_status_enabled"] = *cs.ReadStatusEnabled
	}
	if cs.ConsumptionReportInterval != nil {
		setting["consumption_report_interval"] = *cs.ConsumptionReportInterval
	}
	if cs.TypingIndicatorTimeout != nil {
		setting["typing_indicator_timeout"] = *cs.TypingIndicatorTimeout
	}
	if cs.PreWebhookRetryCount != nil {
		setting["pre_webhook_retry_count"] = *cs.PreWebhookRetryCount
	}
	if cs.PostWebhookRetryCount != nil {
		setting["post_webhook_retry_count"] = *cs.PostWebhookRetryCount
	}

	if len(setting) <= 0 {
		return nil
	}

	return &setting
}

func webhookFromResponse(cs *openapi.ChatV2Service) *map[string]interface{} {
	webhook := map[string]interface{}{}

	if cs.WebhookFilters != nil {
		webhook["events"] = *cs.WebhookFilters
	}
	if cs.WebhookMethod != nil {
		webhook["method"] = *cs.WebhookMethod
	}
	if cs.PreWebhookUrl != nil {
		webhook["pre_hook_url"] = *cs.PreWebhookUrl
	}
	if cs.PostWebhookUrl != nil {
		webhook["post_hook_url"] = *cs.PostWebhookUrl
	}

	if len(webhook) <= 0 {
		return nil
	}

	return &webhook
}

func templateFromResponse(template map[string]interface{}) map[string]interface{} {
	setting := map[string]interface{}{}
	if v, ok := template["enabled"].(bool); ok {
		setting["enabled"] = v
	}
	if v, ok := template["template"].(string); ok {
		setting["template"] = v
	}
	if v, ok := template["sound"].(string); ok {
		setting["sound"] = v
	}
	if v, ok := template["badge_count_enabled"].(bool); ok {
		setting["badge_count_enabled"] = v
	}
	return setting
}

func notificationsFromResponse(noti map[string]interface{}) map[string]interface{} {
	setting := map[string]interface{}{}

	if v, ok := noti["log_enabled"].(bool); ok {
		setting["log_enabled"] = v
	}

	for _, k := range pushNotificationTemplateNames {
		if v, ok := noti[k].(map[string]interface{}); ok {
			setting[k] = []map[string]interface{}{
				templateFromResponse(v),
			}
		}
	}
	return setting
}

func readContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	if changed, ok := d.Get("roles").([]interface{}); ok {
		if len(changed) > 0 {
			roles := []map[string]interface{}{}
			if ss := rolesFromResponse(res); ss != nil {
				roles = append(roles, *ss)
			}
			if err := d.Set("roles", roles); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	if changed, ok := d.Get("limits").([]interface{}); ok {
		if len(changed) > 0 {
			limits := []map[string]interface{}{}
			if res.Limits != nil {
				limits = append(limits, limitsFromResponse(*res.Limits))
			}
			if err := d.Set("limits", limits); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	if changed, ok := d.Get("additional_settings").([]interface{}); ok {
		if len(changed) > 0 {
			additionalSettings := []map[string]interface{}{}
			if ss := additionalSettingsFromResponse(res); ss != nil {
				additionalSettings = append(additionalSettings, *ss)
			}
			if err := d.Set("additional_settings", additionalSettings); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	if changed, ok := d.Get("webhooks").([]interface{}); ok {
		if len(changed) > 0 {
			webhooks := []map[string]interface{}{}
			if ss := webhookFromResponse(res); ss != nil {
				webhooks = append(webhooks, *ss)
			}
			if err := d.Set("webhooks", webhooks); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	if changed, ok := d.Get("notifications").([]interface{}); ok {
		if len(changed) > 0 {
			notifications := []map[string]interface{}{}
			if res.Notifications != nil {
				notifications = append(notifications, notificationsFromResponse(*res.Notifications))
			}
			if err := d.Set("notifications", notifications); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	return diags
}
