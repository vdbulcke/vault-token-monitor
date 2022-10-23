# Install 

You can find the pre-compiled binaries on the release page [https://github.com/vdbulcke/vault-token-monitor/releases](https://github.com/vdbulcke/vault-token-monitor/releases)




## Getting Latest Version 


```sh
TAG=$(curl https://api.github.com/repos/vdbulcke/vault-token-monitor/releases/latest  |jq .tag_name -r )
VERSION=$(echo $TAG | cut -d 'v' -f 2)
```

!!! info
    You will need `jq` and `curl` in your `PATH`

## Docker

See the [Packages Page](https://github.com/vdbulcke/vault-token-monitor/pkgs/container/vault-token-monitor) to find the latest docker image.

```bash
docker run --rm  -v /path/to/your/config.yaml:/app/config.yaml:z -p 9000:9000 ghcr.io/vdbulcke/vault-token-monitor:v${TAG}
```



## MacOS 

=== "Intel"
    1. Download the binary  from the [releases](https://github.com/vdbulcke/vault-token-monitor/releases) page:
      ```sh
      curl -LO "https://github.com/vdbulcke/vault-token-monitor/releases/download/${TAG}/vault-token-monitor_${VERSION}_Darwin_x86_64.tar.gz"
      
      ```
    1. Extract Binary:
      ```sh
      tar xzf "vault-token-monitor_${VERSION}_Darwin_x86_64.tar.gz"
      ```
    1. Check Version: 
      ```sh
      ./vault-token-monitor version
      ```
    1. Install in your `PATH`: 
      ```sh
      sudo install vault-token-monitor /usr/local/bin/
      ```
      Or
      ```sh
      sudo mv vault-token-monitor /usr/local/bin/
      ```

=== "ARM (M1)"
    1. Download the binary  from the [releases](https://github.com/vdbulcke/vault-token-monitor/releases) page:
      ```sh
      curl -LO "https://github.com/vdbulcke/vault-token-monitor/releases/download/${TAG}/vault-token-monitor_${VERSION}_Darwin_amr64.tar.gz"
      
      ```
    1. Extract Binary:
      ```sh
      tar xzf "vault-token-monitor_${VERSION}_Darwin_amr64.tar.gz"
      ```
    1. Check Version: 
      ```sh
      ./vault-token-monitor version
      ```
    1. Install in your `PATH`: 
      ```sh
      sudo install vault-token-monitor /usr/local/bin/
      ```
      Or
      ```sh
      sudo mv vault-token-monitor /usr/local/bin/
      ```
=== "Universal Binary"

    1. Download the binary  from the [releases](https://github.com/vdbulcke/vault-token-monitor/releases) page:
      ```sh
      curl -LO "https://github.com/vdbulcke/vault-token-monitor/releases/download/${TAG}/vault-token-monitor_${VERSION}_Darwin_all.tar.gz"
      
      ```
    1. Extract Binary:
      ```sh
      tar xzf "vault-token-monitor_${VERSION}_Darwin_all.tar.gz"
      ```
    1. Check Version: 
      ```sh
      ./vault-token-monitor version
      ```
    1. Install in your `PATH`: 
      ```sh
      sudo install vault-token-monitor /usr/local/bin/
      ```
      Or
      ```sh
      sudo mv vault-token-monitor /usr/local/bin/
      ```



## Linux 


=== "Intel"
    1. Download the binary  from the [releases](https://github.com/vdbulcke/vault-token-monitor/releases) page:
      ```sh
      curl -LO "https://github.com/vdbulcke/vault-token-monitor/releases/download/${TAG}/vault-token-monitor_${VERSION}_Linux_x86_64.tar.gz"
      
      ```
    1. Extract Binary:
      ```sh
      tar xzf "vault-token-monitor_${VERSION}_Linux_x86_64.tar.gz"
      ```
    1. Check Version: 
      ```sh
      ./vault-token-monitor version
      ```
    1. Install in your `PATH`: 
      ```sh
      sudo install vault-token-monitor /usr/local/bin/
      ```
      Or
      ```sh
      sudo mv vault-token-monitor /usr/local/bin/
      ```

=== "ARM"
    1. Download the binary  from the [releases](https://github.com/vdbulcke/vault-token-monitor/releases) page:
      ```sh
      curl -LO "https://github.com/vdbulcke/vault-token-monitor/releases/download/${TAG}/vault-token-monitor_${VERSION}_Linux_amr64.tar.gz"
      
      ```
    1. Extract Binary:
      ```sh
      tar xzf "vault-token-monitor_${VERSION}_Linux_amr64.tar.gz"
      ```
    1. Check Version: 
      ```sh
      ./vault-token-monitor version
      ```
    1. Install in your `PATH`: 
      ```sh
      sudo install vault-token-monitor /usr/local/bin/
      ```
      Or
      ```sh
      sudo mv vault-token-monitor /usr/local/bin/
      ```
      
## Windows 


=== "Intel"
    1. Download the binary `vault-token-monitor_[VERSION]_Windows_x86_64.zip`  from the [releases](https://github.com/vdbulcke/vault-token-monitor/releases) page
     
    1. Unzip the Binary

    1. Check Version: 
      ```sh
      ./vault-token-monitor.exe version
      ```

