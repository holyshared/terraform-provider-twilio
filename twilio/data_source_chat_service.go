package twilio

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceChatServicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	// FIXME: use twilio sdk?
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/chat_services.json", "http://localhost:19090"), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	chatServices := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&chatServices)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("chat_services", chatServices); err != nil {
		return diag.FromErr(err)
	}

	// FIXME: account sid ?
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

/*
[
	{
		sid: string
		friendlyName: string
		dateCreated: stirng
	}
]
*/
func dataSourceChatServices() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceChatServicesRead,
		Schema: map[string]*schema.Schema{
			"chat_services": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"friendlyName": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dateCreated": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}
