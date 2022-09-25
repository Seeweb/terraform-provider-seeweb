#Example for creating GPU server

provider "seeweb" {} # Expecting Seeweb auth token in env var $SEEWEB_TOKEN

resource "seeweb_server" "GPU-server-001" {
  plan     = "ECS1GPU"
  location = "it-fr2"
  image    = "centos-7"
  notes    = "GPU server 001 reated with Terraform"
}
