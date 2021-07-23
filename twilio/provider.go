package twilio

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/twilio/twilio-go"

	chat "terraform-provider-twilio/twilio/chat"
)

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	accountSid := d.Get("account_sid").(string)
	authToken := d.Get("auth_token").(string)

	var diags diag.Diagnostics

	if (accountSid == "") || (authToken == "") {
		return nil, diag.FromErr(fmt.Errorf("Error: %s", "accountSid and authToken is required"))
	}

	client := twilio.NewRestClientWithParams(accountSid, authToken, twilio.RestClientParams{})

	return client, diags
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_ACCOUNT_SID", nil),
			},
			"auth_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_AUTH_TOKEN", nil),
			},
		},
		ResourcesMap:         chat.ResourcesMap,
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}
