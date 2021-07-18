# terraform-provider-twilio

```
terraform {
  required_providers {
    twilio = {
      version = "0.13"
      source = "holyshared/twilio"
    }
  }
}

provider "twilio" {
  account_sid = var.account_sid
  auth_token = var.auth_token
}

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
}
```
