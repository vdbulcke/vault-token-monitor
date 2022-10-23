# Server Configuration 

You can find a complete example of the client configuration in [example/config.yaml](https://github.com/vdbulcke/vault-token-monitor/blob/main/example/config.yaml).


## Sample Config

```yaml
---


##
## Vault Config
##
### The Vault API Address
### (mandatory)
vault_address: "http://127.0.0.1:8200"
### Vault Token 
###   having the permission to lookupand renew  other 
###   accessor tokens
### (mandatory)
vault_token: "hvs.CAESINEj-..."

### TLS Setting  (optional) 
### Default false
# skip_tls_validation: true
### Path to a PEM encoded CA file
# vault_ca_pem_file: /path/to/ca.pem


##
## Prometheus
## 
### Listening port 
### (mandatory)
prometheus_listening_port: 9000

##
## Scheduler
##
### Period duration at which the scheduler
### will lookup token to check for TTL, and
### for auto renew threshold.
###
###
### Allowed Unit format for duration "s", "m", "h".
### (mandatory)
scheduler_period_duration: 1m

##
## Vault Accessor Token list
##
### (mandatory)
vault_accessor_token_list: 
  ## A Vault token accessor to monitor
  ## (mandatory)
  - token_accessor: "T1hr2GKaWgNrriLQLsQRuf1d"

  ## A Vault token accessor to monitor
  ## and to renew when TTL of token is lower than 
  ## the auto_renew_threshold_duration:
  ### Allowed Unit format for duration "s", "m", "h".
  ## (optional)
  - 
    token_accessor: "F3JwTAZeuAH5KoOdxEXTWwIT"
    auto_renew_threshold_duration: 3m

```

## Vault Config

!!!important
    Mandatory Configuration settings

### Vault Address 

The Vault API address of your Vault Server.

```yaml
vault_address: "http://127.0.0.1:8200"
```

### Vault Token

The Vault token used to authenticate to your Vault Server, and lookup and renew other accessor token. 

```yaml
### Vault Token 
###   having the permission to lookupand renew  other 
###   accessor tokens
### (mandatory)
vault_token: "hvs.CAESINEj-..."

```

Here is an example of a Hashicorp Vault policy (`global/automated-token-renew`) for that token: 
```hcl

## Allow tokens to look up their own properties
## for vault cli login
path "auth/token/lookup-self" {
    capabilities = ["read"]
}


## Allow list of accessors token
path "/auth/token/accessors" {
  capabilities = [ "read", "list"]
}

## Allow accessor tokens lookup 
path "/auth/token/lookup-accessor" {
  capabilities = [ "read", "update"]
}


## Allow renew of accessors token 

path "/auth/token/renew-accessor" {
  capabilities = [ "read", "update"]
}
```

You can generate the token as a periodic and orphan token with:
```bash
vault token create  -policy=global/automated-token-renew -no-default-policy -orphan -period=24h
```

### TLS Configuration

!!!info
    Optional settings

If your Vault Server is using self-signed certificates or a non public Root CA, you can use one of those settings: 

```yaml
### TLS Setting  (optional) 
### Default false
skip_tls_validation: true
### Path to a PEM encoded CA file
vault_ca_pem_file: /path/to/ca.pem
```




## Prometheus Metric Server

!!!Important
    Mandatory setting

```yaml
##
## Prometheus
## 
### Listening port 
### (mandatory)
prometheus_listening_port: 9000
```

## Scheduler Period

!!!Important
    Mandatory setting

The monitoring server will lookup the configured Vault accessor token at this frequency.

```yaml
##
## Scheduler
##
### Period duration at which the scheduler
### will lookup token to check for TTL, and
### for auto renew threshold.
###
###
### Allowed Unit format for duration "s", "m", "h".
### (mandatory)
scheduler_period_duration: 4h
```

## Vault Accessor Token List

!!!Important
    Mandatory setting

This is where you can define Vault accessor tokens to be monitored. The scheduler will lookup the access token's TTL at the configured (`scheduler_period_duration`) frequency, and update the prometheus metric to reflect the TTL expiration time. 

Additionally, you can configure an auto renew threshold (`auto_renew_threshold_duration`) for the accessor token. If the accessor token TTL is lower than the configured threshold then the monitoring server will renew the token. 

```yaml
##
## Vault Accessor Token list
##
### (mandatory)
vault_accessor_token_list: 
  ## A Vault token accessor to monitor
  ## (mandatory)
  - token_accessor: "T1hr2GKaWgNrriLQLsQRuf1d"

  ## A Vault token accessor to monitor
  ## and to renew when TTL of token is lower than 
  ## the auto_renew_threshold_duration:
  ### Allowed Unit format for duration "s", "m", "h".
  ## (optional)
  - 
    token_accessor: "F3JwTAZeuAH5KoOdxEXTWwIT"
    auto_renew_threshold_duration: 24h
```

!!!Warning
    Make sure that your `auto_renew_threshold_duration` are greater than the `scheduler_period_duration` 