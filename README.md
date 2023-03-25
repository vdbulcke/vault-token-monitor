# Vault Token Monitoring Server

`vault-token-monitor`  is a monitoring server that can expose your Vault accessor tokens TTL as prometheus metrics where you can build dashboards and alert policies. 


Moreover `vault-token-monitor`  can also auto-renew token when the TTL is bellowed configurable thresholds.

## Features

- Lookup and expose Vault token TTL as prometheus metrics
- Auto Renew tokens when TTL is below threshold
- Sample [Grafana Dashboards ](./tutorial/grafana-terraform-config/dashboards/)

[Changelog](./CHANGELOG.md)

## Install 

Follow [install doc](https://vdbulcke.github.io/vault-token-monitor/install/) to install binaries. 

Docker images can be found on [ghcr.io/vdbulcke/vault-token-monitor](https://github.com/vdbulcke/vault-token-monitor/pkgs/container/vault-token-monitor)



### Validate Signature With Cosign

Make sure you have `cosign` installed locally (see [Cosign Install](https://docs.sigstore.dev/cosign/installation/)).


Then you can use the `./verify_signature.sh` in this repo: 

```bash
./verify_signature.sh PATH_TO_DOWNLOADED_ARCHIVE TAG_VERSION
```
for example
```bash
$ ./verify_signature.sh  ~/Downloads/vault-token-monitor_0.2.0_Linux_x86_64.tar.gz v0.2.0

Checking Signature for version: v0.2.0
Verified OK

```

## Run 

```bash
vault-token-monitor server --config example/lab.yaml
```

## Documentation

Full documentation can be found [here](https://vdbulcke.github.io/vault-token-monitor/)

See CLI documentation [here](./doc/vault-token-monitor.md).
