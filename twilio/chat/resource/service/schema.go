package service

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	tv "terraform-provider-twilio/twilio/validation"
)

var supportWebHookEvents = tv.ListOfMatchString([]string{
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
})

var Schema = map[string]*schema.Schema{
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
						unWarns, unErrs := validation.ListOfUniqueStrings(val, key)
						for _, w := range unWarns {
							warns = append(warns, w)
						}
						for _, e := range unErrs {
							errs = append(errs, e)
						}

						whWarns, whErrs := supportWebHookEvents(val, key)
						for _, w := range whWarns {
							warns = append(warns, w)
						}
						for _, e := range whErrs {
							errs = append(errs, e)
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
}
