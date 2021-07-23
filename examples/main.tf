terraform {
  required_providers {
    twilio = {
      version = "0.1.4"
      source = "holyshared/twilio"
    }
  }
}

provider "twilio" {
  account_sid = var.account_sid
  auth_token = var.auth_token
}

resource "twilio_chat_service" "terraform_dev" {
  friendly_name = "terraform-dev-1"
}

resource "twilio_chat_service" "terraform_test" {
  friendly_name = "terraform-test-1"
  roles {
    default_service_role = "RL69a3adbc16d9489093fe084f4c798e30"
    default_channel_role = "RL52ecdd95b762485db2363ccff1b82a44"
    default_channel_creator_role = "RL5c67ab91b0b94a2cb95e5c53a263f1d1"
  }
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

resource "twilio_chat_fcm_credential" "terraform_fcm_credential" {
  friendly_name = "terraform-test-fcm-credential"
  secret = var.fcm_secret
}
