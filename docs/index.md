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
}
```
