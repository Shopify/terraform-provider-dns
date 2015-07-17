package main

import (
	"github.com/hashicorp/terraform/plugin"

	"github.com/Shopify/terraform-provider-dns/dns"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: dns.Provider,
	})
}
