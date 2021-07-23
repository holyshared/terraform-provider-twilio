---
page_title: "twilio_chat_fcm_credential Resource - terraform-provider-twilio"
subcategory: ""
description:  "Server keys of push notifications"
---

## Example Usage

```terraform
resource "twilio_chat_fcm_credential" "terraform_fcm_credential" {
  friendly_name = "terraform-test-fcm-credential"
  secret = var.fcm_secret
}
```

## Argument Reference

- `friendly_name` - (Required) The credential name of push notification
- `secret` - (Required) The server key of Firebase Console 
