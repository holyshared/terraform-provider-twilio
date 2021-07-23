---
page_title: "twilio_chat_service Resource - terraform-provider-twilio"
subcategory: ""
description:  "Twilio Programmable Chat configurations"
---

## Example Usage

```terraform
resource "twilio_chat_service" "dev" {
  friendly_name = "my-service"
  limits {
    user_channels = 250
    channel_members = 100
  }
  additional_settings {
    reachability_enabled = true
    read_status_enabled = true
    consumption_report_interval = 10
    typing_indicator_timeout = 5
    pre_webhook_retry_count = 1
    post_webhook_retry_count = 1
  }
  webhooks {
    events = ["onMessageSend", "onMessageSent"]
    method = "POST"
    pre_hook_url = "https://example.com"
    post_hook_url = "https://example.com"
  }
  notifications {
    log_enabled = true
    new_message {
      enabled = true
      template = "$${CHANNEL};$${USER}: $${MESSAGE}"
      sound = "default"
      badge_count_enabled = true
    }
    invited_to_channel {
      enabled = true
      template = "$${USER} has invited you to join the channel $${CHANNEL}"
      sound = "default"
    }
    added_to_channel {
      enabled = true
      template = "You have been added to channel $${CHANNEL} by $${USER}"
      sound = "default"
    }
    removed_from_channel {
      enabled = true
      template = "$${USER} has removed you from the channel $${CHANNEL}"
      sound = "default"
    }
  }
}
```
