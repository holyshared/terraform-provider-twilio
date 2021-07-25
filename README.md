# terraform-provider-twilio

## How to use

Add provider to your main.tf

```terraform
terraform {
  required_providers {
    twilio = {
      version = "0.1.5"
      source = "holyshared/twilio"
    }
  }
}

provider "twilio" {
  account_sid = var.account_sid
  auth_token = var.auth_token
}
```

After the addition is completed, the initialization is performed next.

```shell
terraform init
```
