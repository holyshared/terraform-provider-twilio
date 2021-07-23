package chat

import (
	"terraform-provider-twilio/twilio/chat/resource/credential/fcm"
	"terraform-provider-twilio/twilio/chat/resource/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var ResourcesMap = map[string]*schema.Resource{
	"twilio_chat_service":        service.ResourceCredentialService(),
	"twilio_chat_fcm_credential": fcm.ResourceCredentialService(),
}
