# seeweb\_server

A [server][1] represents something you monitor (like a web server, email server, or database server). It is a container for related incidents that associates them with escalation policies.


## Example Usage

```hcl
resource "seeweb_server" "example" {
  plan        = "ECS1"
  location       = "it-fr2"
  image       = "centos-7"
  notes       = "created using terraform"
}
```

## Argument Reference

The following arguments are supported:

  * `plan` - (Required) the server plan. (*Forces new resource*)
  * `location` - (Required) location identifier. (*Forces new resource*)
  * `image` - (Required) The image name of a public or private template. (*Forces new resource*)
  * `notes` - (Required) The human-readable string you wish to use to display server name.
  * `ssh_key` - (Optional) Public key label. (*Forces new resource*)
  * `group` - (Optional) The group name where the server belongs.

## Attributes Reference

* `ipv4` - The server public IPv4.
* `ipv6` - The server public IPv6.
* `plan_size` - The server plan configuration sizes.
* `so` - The template image used to deploy this server.
* `creation_date` - A time value given in ISO8601 combined date and time format that represents when the server was created.
* `deletion_date` - A time value given in ISO8601 combined date and time format that represents when the server was deleted.
* `active_flag` - A flag value that shows if the server is active.
* `status` - The server status: Booted, Booting, Deleting, Deleted.
* `api_version` - The server API version.
* `user` - The server account username.
* `virttype` - The virtualization engine name.

### Plan Size (`plan_size`) supports the following:

* `core` - Number of CPU cores.
* `ram` - Memory size in MB.
* `disk` - Disk size in MB.

## Import

Servers can be imported using the `id`, e.g.

```
$ terraform import seeweb_server.main PLBP09X
```

[1]: https://docs.seeweb.it/ecs/api/#create-new-server
