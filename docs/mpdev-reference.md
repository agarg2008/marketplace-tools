## Setup

TODO: Give instructions for installing mpdev.

## Mpdev resources

The `mpdev` tool creates resources from yaml templates. The mpdev resource schema
is inspired by the schema for kubernetes resources, where a resource type is
uniquely specified by a `kind` and `apiVersion`.

Currently, the `mpdev` tool supports two types of resources,
[`DeploymentManagerAutogenTemplate`](https://pkg.go.dev/github.com/GoogleCloudPlatform/marketplace-tools/mpdev/internal/apply?tab=doc#DeploymentManagerAutogenTemplate)
and
[`DeploymentManagerTemplate`](https://pkg.go.dev/github.com/GoogleCloudPlatform/marketplace-tools/mpdev/internal/apply?tab=doc#DeploymentManagerTemplate).
See the 
[deployment manager guide](./deployment-manager-guide.md) for how to configure
these resources.

## Commands

### Start from preconfigured mpdev template

The `pkg get` command downloads a preconfigured mpdev template. `mpdev pkg` is
a wrapper around 
[`kpt pkg`](https://googlecontainertools.github.io/kpt/reference/pkg/get]).

```bash
mpdev pkg get https://github.com/marketplace-tools.git/examples/deployment-manager/autogen/singlevm dir
```

### Customize an mpdev template

The `cfg set` command can be used to programmatically customize values in a
preconfigured mpdev template.
`mpdev cfg` is a wrapper around
[`kpt cfg`](https://googlecontainertools.github.io/kpt/reference/cfg/set]).

```bash
mpdev cfg set dir projectId YourProject
```

### Generate mpdev resources

The `apply` command creates resources from the mpdev template.

```bash
mpdev apply -f dir/configurations.yaml
```

The `dry-run` option can be used to verify that your environment is setup
correctly to create the resources in the configuration files. This depends on
the mpdev resource but can include proper `gcp` permissions, and whether 
`docker` is installed.

```bash
mpdev apply --dry-run -f dir/configurations.yaml
```