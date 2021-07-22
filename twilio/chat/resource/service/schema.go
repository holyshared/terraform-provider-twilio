package service

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var supportWebHookEvents = []string{
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

var roles = schema.Schema{
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
}

var limits = schema.Schema{
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
}

var additionalSettings = schema.Schema{
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
}

var webhooks = schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"events": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice(supportWebHookEvents, false),
				},
				Optional: true,
				Computed: false,
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
}

var notificationTemplate = schema.Resource{
	Schema: map[string]*schema.Schema{
		"enabled": {
			Type:     schema.TypeBool,
			Computed: false,
			Optional: true,
		},
		"template": {
			Type:     schema.TypeString,
			Computed: false,
			Optional: true,
		},
		"sound": {
			Type:     schema.TypeString,
			Computed: false,
			Optional: true,
		},
	},
}

var notificationTemplateWithBadgeCount = schema.Resource{
	Schema: map[string]*schema.Schema{
		"enabled": {
			Type:     schema.TypeBool,
			Computed: false,
			Optional: true,
		},
		"template": {
			Type:     schema.TypeString,
			Computed: false,
			Optional: true,
		},
		"sound": {
			Type:     schema.TypeString,
			Computed: false,
			Optional: true,
		},
		"badge_count_enabled": {
			Type:     schema.TypeBool,
			Computed: false,
			Optional: true,
		},
	},
}

var notifications = schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"log_enabled": {
				Type:     schema.TypeBool,
				Computed: false,
				Optional: true,
			},
			"new_message": {
				Type:     schema.TypeList,
				Elem:     &notificationTemplateWithBadgeCount,
				Computed: false,
				Optional: true,
				MaxItems: 1,
			},
			"invited_to_channel": {
				Type:     schema.TypeList,
				Elem:     &notificationTemplate,
				Computed: false,
				Optional: true,
				MaxItems: 1,
			},
			"added_to_channel": {
				Type:     schema.TypeList,
				Elem:     &notificationTemplate,
				Computed: false,
				Optional: true,
				MaxItems: 1,
			},
			"removed_from_channel": {
				Type:     schema.TypeList,
				Elem:     &notificationTemplate,
				Computed: false,
				Optional: true,
				MaxItems: 1,
			},
		},
	},
	Optional: true,
	Computed: false,
	MaxItems: 1,
}

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
	"roles":               &roles,
	"limits":              &limits,
	"additional_settings": &additionalSettings,
	"webhooks":            &webhooks,
	"notifications":       &notifications,
}
