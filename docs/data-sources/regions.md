# seeweb\_regions

Use this data source to get information about [regions][1] available.

## Example Usage

```hcl
data "seeweb_regions" "example" {}
```

## Attributes Reference

* `regions` - List of regions with their data.

### Regions (`regions`) supports the following:

* `id` - A unique number ID.
* `location` - Region location.
* `description` - Region description.

[1]: https://docs.seeweb.it/ecs/api/#list-all-regions
