#!/bin/bash
set -e
set -x
set -o pipefail

# podman run -d   --network host --name prometheus  -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml:z  prom/prometheus:latest --config.file=/etc/prometheus/prometheus.yml
podman run -d   --network host --name grafana  -e "GF_SECURITY_ADMIN_PASSWORD=password" grafana/grafana:latest