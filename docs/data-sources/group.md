# seeweb\_group

Use this data source to get information about a specific [group][1].

## Example Usage

```hcl
data "seeweb_group" "example" {
  id = 12
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) A unique number ID.

## Attributes Reference

* `name` - A unique group name.
* `notes` - A human-readable friendly description.
* `enabled` - The current status of a group.

[1]: https://docs.seeweb.it/ecs/api/#list-all-groups
