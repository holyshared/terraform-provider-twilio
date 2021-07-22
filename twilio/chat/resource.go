package chat

import (
	"terraform-provider-twilio/twilio/chat/resource/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceChatService() *schema.Resource {
	return &schema.Resource{
		CreateContext: service.CreateContext,
		ReadContext:   service.ReadContext,
		UpdateContext: service.UpdateContext,
		DeleteContext: service.DeleteContext,
		Schema:        service.Schema,
	}
}
