terraform {
  required_providers {
    grafana = {
      source = "grafana/grafana"
      version = "1.30.0"
    }
  }
}

## https://registry.terraform.io/providers/grafana/grafana/latest/docs
provider "grafana" {
  url  = "http://localhost:3000"
  auth = "${var.grafana_auth_user}:${var.grafana_auth_password}"
}