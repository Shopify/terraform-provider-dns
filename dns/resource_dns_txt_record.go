package dns

import (
	"net"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDnsTxtRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceDnsTxtRecordCreate,
		Update: resourceDnsTxtRecordUpdate,
		Read:   resourceDnsTxtRecordRead,
		Delete: resourceDnsTxtRecordDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"update": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"record": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"records": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
		},
	}
}

func resourceDnsTxtRecordCreate(d *schema.ResourceData, meta interface{}) error {
	d.SetId(d.Get("name").(string))
	records, err := net.LookupTXT(d.Id())
	if err != nil {
		return err
	}

	if len(records) > 0 {
		d.Set("record", records[0])
	}
	d.Set("records", records)
	return nil
}

func resourceDnsTxtRecordRead(d *schema.ResourceData, meta interface{}) error {
	// Read if forcing an update, or we haven't yet read any records
	if d.Get("update").(bool) {
		records, err := net.LookupTXT(d.Id())
		if err != nil {
			return err
		}

		if len(records) > 0 {
			d.Set("record", records[0])
		}
		d.Set("records", records)
	}
	return nil
}

func resourceDnsTxtRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceDnsTxtRecordRead(d, meta)
}

func resourceDnsTxtRecordDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
