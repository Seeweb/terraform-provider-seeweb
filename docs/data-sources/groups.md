# seeweb\_groups

Use this data source to get information about [groups][1] available.

## Example Usage

```hcl
data "seeweb_groups" "example" {}
```

## Attributes Reference

* `groups` - List of groups with their data.

### groups (`groups`) supports the following:

* `id` - A unique number ID.
* `name` - A unique group name.
* `notes` - A human-readable friendly description.
* `enabled` - The current status of a group.


[1]: https://docs.seeweb.it/ecs/api/#list-all-groups