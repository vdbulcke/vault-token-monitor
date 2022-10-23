

## https://registry.terraform.io/providers/grafana/grafana/latest/docs/resources/data_source
resource "grafana_data_source" "prometheus" {
  type = "prometheus"
  name = "prometheus"
  url  = "http://localhost:9090"

  is_default = true
  uid = "QOkjnnNVk"
  json_data_encoded = jsonencode({
    http_method     = "POST"
  })
}


## https://registry.terraform.io/providers/grafana/grafana/latest/docs/resources/dashboard
resource "grafana_dashboard" "metrics" {
  # config_json = file("${path.module}/dashboards/vault-token-monitor.json")
  config_json = file("${path.module}/dashboards/vault-token-monitor-exact.json")
}