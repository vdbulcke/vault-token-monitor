
variable "grafana_auth_password" {
  description = "The Grafana admin password"
  type = string
  sensitive = true
  default = "password"
}

variable "grafana_auth_user" {
  description = "The Grafana admin user"
  type = string
  default = "admin"
}