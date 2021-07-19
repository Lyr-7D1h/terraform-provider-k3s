terraform {
  required_providers {
    k3s = {
      version = "0.1"
      source  = "hashicorp.com/lyr-7d1h/k3s"
    }
  }
}

provider "k3s" {
  public_ssh_key = "~/.ssh/id_rsa.pub"
}
