# Seeweb Provider

[Seeweb](https://www.seeweb.it/) Cloud Server Shared CPU is the on demand cloud service ideal for testing, development, staging and for who needs to create and destroy per-per-use cloud machines quickly and with cost optimization.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Seeweb provider
terraform {
  required_providers {
    seeweb = {
      source  = "seeweb/seeweb"
    }
  }
}

provider "seeweb" {
  token = var.seeweb_token # Or Provider expects to find it in $SEEWEB_TOKEN
}

```

## Argument Reference

The following arguments are supported:

* `token` - (Required) The v2 authorization token. It can also be sourced from the SEEWEB_TOKEN environment variable. See [API Documentation](https://docs.seeweb.it/ecs/api/#api-authentication) for more information.
* `api_url_override` - (Optional) It can be used to set a custom proxy endpoint as Seeweb client api url overriding the default one.
