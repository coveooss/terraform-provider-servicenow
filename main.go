package main

import (
	"github.com/coveooss/terraform-provider-servicenow/servicenow"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return servicenow.Provider()
		},
	})
}
