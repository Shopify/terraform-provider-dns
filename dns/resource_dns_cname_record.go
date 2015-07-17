package dns

import (
	"net"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDnsCnameRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceDnsCnameRecordCreate,
		Update: resourceDnsCnameRecordUpdate,
		Read:   resourceDnsCnameRecordRead,
		Delete: resourceDnsCnameRecordDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"cname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceDnsCnameRecordCreate(d *schema.ResourceData, meta interface{}) error {
	d.SetId(d.Get("name").(string))
	return resourceDnsCnameRecordRead(d, meta)
}

func resourceDnsCnameRecordRead(d *schema.ResourceData, meta interface{}) error {
	cname, err := net.LookupCNAME(d.Id())
	if err != nil {
		return err
	}

	d.Set("cname", cname)
	return nil
}

func resourceDnsCnameRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceDnsCnameRecordRead(d, meta)
}

func resourceDnsCnameRecordDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
