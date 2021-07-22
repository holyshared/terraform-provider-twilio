package service

import (
	"context"
	"time"

	tw "github.com/twilio/twilio-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ReadContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
