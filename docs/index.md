---
page_title: "Provider: Twilio"
subcategory: ""
description: |-
  Terraform provider for interacting with Twilio API.
---

# Twilio Provider

Twilio Provider can manage Twilio configuration.

Use the navigation to the left to read about the available resources.

## Example Usage

Do not keep your authentication password in HCL for production environments, use Terraform environment variables.

```terraform
provider "twilio" {
  account_sid = var.account_sid
  auth_token = var.auth_token
}
```

## Schema

### Optional

- **account_sid** (String) Username to authenticate to Twilio API
- **auth_token** (String) Auth token to authenticate to Twilio API
