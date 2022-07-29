# seeweb\_plans

Use this data source to get information about [plans][1] available.

## Example Usage

```hcl
data "seeweb_plans" "example" {}
```

## Attributes Reference

* `plans` - List of plans with their data.

### Plans (`plans`) supports the following:

* `id` - A unique number ID.
* `name` - Plan name.
* `cpu` - Number of CPU.
* `ram` - Memory sixe in MB.
* `disk` - Disk size in MB.
* `hourly_price` - Hourly price in Euro.
* `montly_price` - Montly price in Euro.
* `windows` - So windows family, true or false.
* `available` - Plan available, true or false.
* `available_regions` - Available regions.

The `available_regions` block contains the following attributes:

* `id` - A unique number ID.
* `location` - Region location.
* `description` - Region description.


[1]: https://docs.seeweb.it/ecs/api/#list-all-plans
