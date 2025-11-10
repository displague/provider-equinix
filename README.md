# Provider Equinix

`provider-equinix` is a [Crossplane](https://crossplane.io/) provider for
[Equinix](https://www.equinix.com/) that is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the Equinix API.

## Overview

This provider supports:
- **Fabric Resources**: Cloud Router, Connection, Network, Routing Protocol, Service Profile
- **Network Edge Resources**: ACL Template, BGP, Device, Device Link, SSH Key, SSH User

**Note**: Metal resources (`equinix_metal_*`) are deprecated and not included in this provider.

The provider is generated from the [Equinix Terraform Provider](https://github.com/equinix/terraform-provider-equinix) version **2.5.0**.

## Getting Started

Install the provider by using the following crossplane Provider manifest:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-equinix
spec:
  package: displague/provider-equinix:latest
```

## Available Resources

The provider includes both **cluster-scoped** and **namespace-scoped** variants of all resources.

### Fabric Resources (fabric.equinix.crossplane.io)
- CloudRouter
- Connection
- Network
- RoutingProtocol
- ServiceProfile

### Network Edge Resources (network.equinix.crossplane.io)
- ACLTemplate
- BGP
- Device
- DeviceLink
- SSHKey
- SSHUser

## Configuration

Create a ProviderConfig with your Equinix credentials:

```yaml
apiVersion: equinix.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: equinix-creds
      namespace: crossplane-system
      key: credentials
```

The credentials secret should contain a JSON object with your Equinix API credentials:

```json
{
  "client_id": "your-client-id",
  "client_secret": "your-client-secret",
  "auth_token": "your-auth-token"
}
```

## Developing

### Prerequisites
- Go 1.24+
- Terraform 1.5.x (< 1.6.0 due to license restrictions)
- Kubernetes cluster with Crossplane installed

### Generate Provider Code

```console
make generate
```

### Build Provider Binary

```console
go build -o bin/provider ./cmd/provider/
```

### Run Against a Kubernetes Cluster

```console
make run
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/displague/provider-equinix/issues).

## License

provider-equinix is under the Apache 2.0 license.
