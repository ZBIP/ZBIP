terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

variable "digitalocean_token" {}

provider "digitalocean" {
  token = var.digitalocean_token
}

data "digitalocean_app" "existing_app" {
  app_id = "3de9aff3-00fe-4574-8045-72435ee8e248"
}

resource "digitalocean_app" "zbip-app" {
app_id = data.digitalocean_app.existing_app.app_id
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
