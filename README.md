# Vault Token Monitoring Server

`vault-token-monitor`  is a monitoring server that can expose your Vault accessor tokens TTL as prometheus metrics where you can build dashboards and alert policies. 


Moreover `vault-token-monitor`  can also auto-renew token when the TTL is bellowed configurable thresholds.

## Features

- Lookup and expose Vault token TTL as prometheus metrics
- Auto Renew tokens when TTL is below threshold
- Sample [Grafana Dashboards ](./tutorial/grafana-terraform-config/dashboards/)

## Install 

Follow [install doc](https://vdbulcke.github.io/vault-token-monitor/install/) to install binaries. 

Docker images can be found on [ghcr.io/vdbulcke/vault-token-monitor](https://github.com/vdbulcke/vault-token-monitor/pkgs/container/vault-token-monitor)


## Documentation

Full documentation can be found [here](https://vdbulcke.github.io/vault-token-monitor/)