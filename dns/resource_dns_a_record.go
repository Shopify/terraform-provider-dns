package dns

import (
	"net"
	"sort"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDnsARecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceDnsARecordCreate,
		Read:   resourceDnsARecordRead,
		Delete: resourceDnsARecordDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optionally sort A records in alphabetical order.
			// This is helpful when a name uses round-robin DNS, which may
			// sort records with multiple addresses in a non-deterministic order.
			// This random sorting can cause flapping in terraform plans, where
			// the changes in sort order cause dependent resources to update
			// despite having no real change in the set of addresses.
			"sort": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},

			"addrs": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
		},
	}
}

func resourceDnsARecordCreate(d *schema.ResourceData, meta interface{}) error {
	d.SetId(d.Get("name").(string))
	return resourceDnsARecordRead(d, meta)
}

func resourceDnsARecordRead(d *schema.ResourceData, meta interface{}) error {
	addrs, err := net.LookupHost(d.Id())
	if err != nil {
		return err
	}

	sortingEnabled := d.Get("sort").(bool)

	if sortingEnabled {
		sort.Strings(addrs)
	}

	d.Set("addrs", addrs)
	return nil
}

func resourceDnsARecordDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
