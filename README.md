# K3S - Terraform Provider

Terraform provider for setting up a k3s cluster on bare metal machines.

## Support

It will try to be as headless as possible abstracting away any complexities.

_**NOTE:**_ Because it tries to do everything itself it might make undesired decisions and might not work on all systems.

Supported OS's:

- Armbian
- Ubuntu

Supported devices:

- Odroid HC4
- Odroid MC1
- Raspberry Pi B+

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
- [Go](https://golang.org/doc/install) >= 1.15

## Using the provider

Example:

```tf
provider "k3s" {
  # Used for setting up ssh connections to servers
  pub_ssh_key = "~/.ssh/id_rsa.pub"

  # Server not yet using ssh key defined try connecting using default username and passwords
  extra_default_ssh_credentials = {
    "username" = "password"
  }
}

locals {
  servers = ["192.168.1.6", "192.168.1.7", "192.168.1.8"]
  servers = ["192.168.1.15", "55.150.23.57"]
}

resource "k3s_server" "server" {
  for_each = local.servers
  ip = each.value

  # K3S cli arguments (https://rancher.com/docs/k3s/latest/en/installation/install-options/)
  flags = ["--disable traefik"]

  # Kubernetes Labels (https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)
  labels = []
}

resource "k3s_server" "agent" {
  for_each = local.agents

  ip = each.value
  labels = []
}
```
