package main

import (
    "github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	"github.com/coveo/terraform-provider-servicenow/servicenow"
)

func main() {
    plugin.Serve(&plugin.ServeOpts {
        ProviderFunc: func() terraform.ResourceProvider {
            return servicenow.ServiceNowProvider()
        },
    })
}