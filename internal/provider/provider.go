package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/lyr-7d1h/terraform-provider-k3s/internal/clusterit"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown
}

func Provider(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"public_ssh_key": {
					Type: schema.TypeString,
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				// "scaffolding_data_source": dataSourceScaffolding(),
			},
			ResourcesMap: map[string]*schema.Resource{
				// "scaffolding_resource": resourceScaffolding(),
			},
			ConfigureContextFunc: configure,
		}

		return p
	}
}

func configure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	public_ssh_key_path := d.Get("public_ssh_key").(string)


	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	k, err := clusterit.New(public_ssh_key_path)

	if err != nil {
		diags = append(diags, diag.Diagnostic{  Severity: diag.Error,  Summary:  "Unable to create k3s client",  Detail:   err.Error(),})
		return nil, diags 
	}


	return k, diags
}