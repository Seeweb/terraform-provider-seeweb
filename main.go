package main

import (
	"github.com/Seeweb/terraform-provider/seeweb"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: seeweb.Provider,
	})
}
