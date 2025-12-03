# Provider HSDP

`provider-hsdp` is a [Crossplane](https://crossplane.io/) provider that is built using [Upjet](https://github.com/crossplane/upjet) code generation tools and exposes XRM-conformant managed resources for the [HSDP (Philips HealthSuite Digital Platform)](https://www.philips.com/a-w/about/innovation/healthsuite-digital-platform.html) API.

## Overview

This provider wraps the [terraform-provider-hsdp](https://github.com/philips-software/terraform-provider-hsdp) to provide Kubernetes-native management of HSDP resources including:

- **IAM (Identity and Access Management)**: Organizations, Users, Groups, Roles, Applications, Services, Clients, etc.
- **DBS (Data Broker Service)**: SQS Subscribers, Topic Subscriptions
- **MDM (Master Data Management)**: Propositions, Data Types
- **Tenant Management**: Tenant Keys

## Getting Started

For more information on getting started with Crossplane providers, see the [Crossplane documentation](https://docs.crossplane.io/latest/packages/providers/).

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/dip-software/provider-dip/issues).
