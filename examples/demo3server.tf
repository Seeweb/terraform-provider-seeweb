#Example for creating three servers: load balancer, webserver, database

provider "seeweb" {} # Expecting Seeweb auth token in env var $SEEWEB_TOKEN

resource "seeweb_server" "server-lb" {
  plan     = "ECS1"
  location = "it-fr2"
  image    = "centos-7"
  notes    = "Demo 3 server - LB"
}

resource "seeweb_server" "server-web" {
  plan     = "ECS2"
  location = "it-fr2"
  image    = "centos-7"
  notes    = "Demo 3 server - WEB"
}

resource "seeweb_server" "server-db" {
  plan     = "ECS2"
  location = "it-fr2"
  image    = "centos-7"
  notes    = "Demo 3 server - DATABASE"
}
