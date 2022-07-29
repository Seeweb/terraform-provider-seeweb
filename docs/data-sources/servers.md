# seeweb\_servers

Use this data source to get information about [servers][1] available.

## Example Usage

```hcl
data "seeweb_servers" "example" {}
```

## Attributes Reference

* `servers` - List of servers with their data.

### servers (`servers`) supports the following:

* `name` - The server name to use to find a server in the Seeweb API.
* `ipv4` - The server public IPv4.
* `ipv6` - The server public IPv6.
* `plan` - The current server plan.
* `plan_size` - The server plan configuration sizes.
* `location` - An unique identifier for the server region.
* `notes` - A human-readable name for this server.
* `so` - The template image used to deploy this server.
* `creation_date` - A time value given in ISO8601 combined date and time format that represents when the server was created.
* `deletion_date` - A time value given in ISO8601 combined date and time format that represents when the server was deleted.
* `active_flag` - A flag value that shows if the server is active.
* `status` - The server status: Booted, Booting, Deleting, Deleted.
* `api_version` - The server API version.
* `user` - The server account username.
* `virttype` - The virtualization engine name.

### Plan Size (`plan_size`) supports the following:

* `core` - pending.
* `ram` - Memory size in MB.
* `disk` - Disk size in MB.


[1]: https://docs.seeweb.it/ecs/api/#list-all-servers
