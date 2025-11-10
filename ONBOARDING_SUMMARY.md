# Provider Equinix Onboarding Summary

This document summarizes the fresh onboarding of provider-equinix following the official upjet guide.

## What Was Done

### 1. Repository Reset
- Removed all old crossplane-runtime based code
- Replaced with upjet-provider-template structure
- Ran `hack/prepare.sh` to configure naming:
  - Provider name: equinix
  - Organization: displague  
  - CRD root group: crossplane.io

### 2. Terraform Provider Configuration
Updated Makefile with:
```makefile
export TERRAFORM_PROVIDER_SOURCE := equinix/equinix
export TERRAFORM_PROVIDER_VERSION := 2.5.0
export TERRAFORM_PROVIDER_DOWNLOAD_NAME := terraform-provider-equinix
export TERRAFORM_DOCS_PATH := docs/resources
```

### 3. Authentication Configuration
Configured `internal/clients/equinix.go` to support:
- `client_id` and `client_secret`
- `auth_token`
- `token`
- Optional configuration: endpoint, request_timeout, max_retries, etc.

### 4. External Name Configuration
Added external name configurations for 11 resources in `config/external_name.go`:

**Fabric Resources:**
- equinix_fabric_connection
- equinix_fabric_cloud_router
- equinix_fabric_network
- equinix_fabric_routing_protocol
- equinix_fabric_service_profile

**Network Edge Resources:**
- equinix_network_acl_template
- equinix_network_bgp
- equinix_network_device
- equinix_network_device_link
- equinix_network_ssh_key
- equinix_network_ssh_user

**Excluded:** All `equinix_metal_*` resources (deprecated)

### 5. Code Generation
- Generated provider schema from Terraform provider
- Generated 22 CRDs (11 cluster-scoped + 11 namespace-scoped)
- Generated controllers for all resources
- Generated API types

### 6. Bug Fixes
- Fixed `SAFEHOST_PLATFORM` variable in Makefile (changed to `HOST_PLATFORM`)
- Fixed ProviderConfig type imports
- Skipped documentation scraper (Equinix docs use different format)

## Generated Resources

### Cluster-Scoped Resources (equinix.crossplane.io)
- fabric.equinix.crossplane.io/CloudRouter
- fabric.equinix.crossplane.io/Connection
- fabric.equinix.crossplane.io/Network
- fabric.equinix.crossplane.io/RoutingProtocol
- fabric.equinix.crossplane.io/ServiceProfile
- network.equinix.crossplane.io/ACLTemplate
- network.equinix.crossplane.io/BGP
- network.equinix.crossplane.io/Device
- network.equinix.crossplane.io/DeviceLink
- network.equinix.crossplane.io/SSHKey
- network.equinix.crossplane.io/SSHUser

### Namespace-Scoped Resources (equinix.m.crossplane.io)
Same 11 resources with namespace-scoped API group

## File Structure

```
provider-equinix/
├── apis/                          # API types
│   ├── cluster/                   # Cluster-scoped resources
│   │   ├── fabric/v1alpha1/      # Fabric resources
│   │   ├── network/v1alpha1/     # Network resources
│   │   ├── v1alpha1/             # ProviderConfig
│   │   └── v1beta1/              # ProviderConfig
│   └── namespaced/               # Namespace-scoped resources
│       ├── fabric/v1alpha1/
│       ├── network/v1alpha1/
│       ├── v1alpha1/
│       └── v1beta1/
├── cmd/
│   ├── generator/                # Code generator
│   └── provider/                 # Provider binary
├── config/                       # Provider configuration
│   ├── external_name.go         # External name configs
│   ├── provider.go              # Provider setup
│   ├── provider-metadata.yaml   # Provider metadata
│   └── schema.json              # Terraform schema
├── internal/
│   ├── clients/                 # Terraform client setup
│   └── controller/              # Controllers
│       ├── cluster/
│       └── namespaced/
├── package/crds/                # Generated CRDs
└── examples/                    # Example manifests
```

## Next Steps

To use this provider:

1. **Install the provider** in your Kubernetes cluster
2. **Create a ProviderConfig** with Equinix credentials
3. **Create managed resources** using the generated CRDs

## Build Commands

```bash
# Generate code
make generate

# Build provider
go build -o bin/provider ./cmd/provider/

# Build generator  
go build -o bin/generator ./cmd/generator/
```

## Notes

- Terraform Provider version: 2.5.0
- Terraform version: 1.5.7 (< 1.6.0 due to license restrictions)
- Go version: 1.24+
- All code compiles successfully
- Both cluster-scoped and namespace-scoped resources available
