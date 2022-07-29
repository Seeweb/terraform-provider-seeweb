# seeweb\_action

Use this data source to get information about a specific [action][1].

## Example Usage

```hcl
data "seeweb_action" "example" {
  id = 10535
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) A unique numeric ID that can be used to identify and reference an action.

## Attributes Reference

* `status` - The current status of the action. This can be "in-progress", "completed", or "error".
* `resource` - A unique identifier for the resource that the action is associated with* `plan` - The current action plan.
* `user` - A unique identifier for the account that the action is associated with.
* `created_at` - A time value given in ISO8601 combined date and time format that represents when the action was created.
* `started_at` - A time value given in ISO8601 combined date and time format that represents when the action was initiated.
* `completed_at` - A time value given in ISO8601 combined date and time format that represents when the action was completed.
* `type` - This is the type of action that the object represents.
* `resource_type` - The type of resource that the action is associated with.
* `progress` - A value that represent the percentage of completation.


[1]: https://docs.seeweb.it/ecs/api/#list-all-actions
