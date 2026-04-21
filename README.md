# promptviser

LLM Prompt Adviser

### Client config

    cat ~/.config/promptviser/config.yaml

```yaml
---
clients:
  local_wfe:
    host: https://0.0.0.0:7880
    tls:
      cert: /tmp/promptviser/certs/promptviser_admin.pem
      key: /tmp/promptviser/certs/promptviser_admin.key
      trusted_ca: /tmp/promptviser/certs/trusty_root_ca.pem
    request:
      retry_limit: 3
      timeout: 6s
    storage_folder: ~/.config/promptviser
```

## Usage

```sh
Usage: promptviserctl <command> [flags]

CLI tool for promptviser service

Flags:
  -h, --help                                       Show context-sensitive help.
  -s, --server=STRING                              Address of the remote server to connect. Use PROMPTVISER_SERVER environment to override
  -D, --debug                                      Enable debug mode
      --o=STRING                                   Print output format: json|yaml
      --cfg="~/.config/promptviser/config.yaml"    Configuration file
      --storage=STRING                             flag specifies to override default location: ~/.config/promptviser. Use PROMPTVISER_STORAGE environment to override
  -H, --http                                       Use HTTP client
      --timeout=6                                  Connection timeout
  -c, --cert=STRING                                Client certificate file for mTLS
  -k, --cert-key=STRING                            Client certificate key for mTLS
  -r, --trusted-ca=STRING                          Trusted CA store for server TLS

Commands:
  version    print remote server version
  server     print remote server status
  submit     Submit data for analysis

Run "promptviserctl <command> --help" for more information on a command.
```
