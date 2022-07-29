# seeweb\_group

A [group][1] is a label that can be applied to a server in order to better organize your infrastructure.


## Example Usage

```hcl
resource "seeweb_group" "example" {
  notes       = "created using terraform"
  password       = "secret"
}
```

## Argument Reference

The following arguments are supported:

  * `notes` - (Required) A human-readable friendly description.
  * `password` - (Required) A private password used to make a subview of the infrastructure.

## Attributes Reference

* `id` - A unique number ID.
* `name` - A unique group name.
* `notes` - A human-readable friendly description.
* `enabled` - The current status of a group.

## Import

groups can be imported using the `id`, e.g.

```
$ terraform import seeweb_group.main PLBP09X
```

[1]: https://docs.seeweb.it/ecs/api/#create-a-new-group