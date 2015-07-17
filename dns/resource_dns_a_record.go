package dns

import (
	"net"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDnsARecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceDnsARecordCreate,
		Update: resourceDnsARecordUpdate,
		Read:   resourceDnsARecordRead,
		Delete: resourceDnsARecordDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

	d.Set("addrs", addrs)
	return nil
}

func resourceDnsARecordUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceDnsARecordRead(d, meta)
}

func resourceDnsARecordDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
