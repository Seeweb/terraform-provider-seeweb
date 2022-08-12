# Terraform Provider for Seeweb

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x
-	[Go](https://golang.org/doc/install) 1.18 (to build the provider plugin)

## Building The Provider

Clone repository to: `$GOPATH/src/github.com/Seeweb/terraform-provider-seeweb`

```sh
$ mkdir -p $GOPATH/src/github.com/Seeweb; cd $GOPATH/src/github.com/Seeweb
$ git clone git@github.com:Seeweb/terraform-provider-seeweb
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/Seeweb/terraform-provider-seeweb
$ export GOPRIVATE=github.com/Seeweb/* # This step is only necessary if the modules are kept private
$ make build
```

## Using built Provider Locally

In order to test a built version of this Terraform Provider locally, you will to run the following...

```sh
go build -o terraform-provider-seeweb
# The next location might be different in you machine so would need to check first
mv terraform-provider-seeweb ~/.terraform.d/plugins/registry.terraform.io/hashicorp/seeweb/0.0.1/darwin_arm64
```

After that you will be able to require the provider as `seeweb`, like in the following *example*...

## Example

```hcl
provider "seeweb" {} # Expecting Seeweb auth token in env var $SEEWEB_TOKEN

resource "seeweb_server" "testacc" {
  plan     = "ECS1"
  location = "it-fr2"
  image    = "centos-7"
  notes    = "created with Terraform"
}
```

## Testing

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
