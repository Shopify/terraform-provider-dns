package dns

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},

		ResourcesMap: map[string]*schema.Resource{
			"dns_a_record":     resourceDnsARecord(),
			"dns_cname_record": resourceDnsCnameRecord(),
			"dns_txt_record":   resourceDnsTxtRecord(),
		},
	}
}
