# seeweb\_templates

Use this data source to get information about [templates][1] available.

## Example Usage

```hcl
data "seeweb_templates" "example" {}
```

## Attributes Reference

* `templates` - List of templates with their data.

### templates (`templates`) supports the following:

* `id` - A unique numeric ID that can be used to identify and reference a template.
* `creation_date` - A time value given in ISO8601 combined date and time format that represents when the template was created.
* `active_flag` - A flag that represents if a template is manageble.
* `status` - The current status of the template. This can be "Creating", "Created", "Deleting" or "Deleted".
* `uuid` - A unique identifier.
* `notes` - A human-readable friendly description.


[1]: https://docs.seeweb.it/ecs/api/#list-all-templates
