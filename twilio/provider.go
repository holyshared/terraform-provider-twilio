package twilio

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/twilio/twilio-go"
)

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	accountSid := d.Get("accountSid").(string)
	authToken := d.Get("authToken").(string)

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
			"accountSid": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_ACCOUNT_SID", nil),
			},
			"authToken": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_AUTH_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"twilio_chat_services": dataSourceChatServices(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}
