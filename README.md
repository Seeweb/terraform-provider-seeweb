# Terraform Provider for Seeweb

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x
-	[Go](https://golang.org/doc/install) 1.18 (to build the provider plugin)

## Building The Provider

Clone repository to: `$GOPATH/src/github.com/uwtrilogyseaward0m/terraform-provider-seeweb`

```sh
$ mkdir -p $GOPATH/src/github.com/uwtrilogyseaward0m; cd $GOPATH/src/github.com/uwtrilogyseaward0m
$ git clone git@github.com:uwtrilogyseaward0m/terraform-provider-seeweb
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/uwtrilogyseaward0m/terraform-provider-seeweb
$ make build
```
### Testing

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`. Running the acceptance tests requires
that the `SEEWEB_TOKEN` environment variable be set to a valid API Token. 

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

Run a specific subset of tests by name use the `TESTARGS="-run TestName"` option which will run all test functions with "TestName" in their name.

```sh
$ make testacc TESTARGS="-run TestAccSeewebServer"
```
