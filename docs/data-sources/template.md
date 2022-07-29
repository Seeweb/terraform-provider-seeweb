# seeweb\_template

Use this data source to get information about a specific [template][1].

## Example Usage

```hcl
data "seeweb_template" "example" {
  id = 504
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) A unique numeric ID that can be used to identify and reference a template.

## Attributes Reference

* `name` - A unique template name.
* `creation_date` - A time value given in ISO8601 combined date and time format that represents when the template was created.
* `active_flag` - A flag that represents if a template is manageble.
* `status` - The current status of the template. This can be "Creating", "Created", "Deleting" or "Deleted".
* `uuid` - A unique identifier.
* `notes` - A human-readable friendly description.

[1]: https://docs.seeweb.it/ecs/api/#list-all-templates
