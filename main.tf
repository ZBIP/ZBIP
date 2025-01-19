terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

variable "do_token" {}

provider "digitalocean" {
  token = var.do_token
}

resource "digitalocean_app" "zbip-app" {
  spec {
    name   = "zbip-app"
    region = "fra"

    service {
      name               = "go-service"
      instance_count     = 1
      instance_size_slug = "apps-s-1vcpu-1gb"

      git {
        repo_clone_url = "https://github.com/ZBIP/ZBIP.git"
        branch         = "main"
      }
    }
  }
}
