# kilosubspace

`kilosubspace` enables using [Subspace](https://github.com/subspacecommunity/subspace) as a UI to define and manage peers for [Kilo](https://github.com/squat/kilo).

[![Build Status](https://travis-ci.org/squat/kilosubspace.svg?branch=master)](https://travis-ci.org/squat/kilosubspace)
[![Go Report Card](https://goreportcard.com/badge/github.com/squat/kilosubspace)](https://goreportcard.com/report/github.com/squat/kilosubspace)

## Getting Started

To run `kilosubspace`, first [install Kilo](https://github.com/squat/kilo#installing-on-kubernetes).
Next, edit the included manifest and set the `SUBSPACE_HTTP_HOST` and `SUBSPACE_LISTENPORT` variables to match the endpoint for one of the nodes in the Kilo mesh.

*Note*: the host may be either an IP address or DNS name but must match the endpoint of an existing node; peers will use this host:port combination to connect to the mesh; the host will also be used as the domain for SAML and cookies.

Finally, deploy `kilosubspace` to the cluster:

```shell
kubectl apply -f https://raw.githubusercontent.com/squat/kilosubspace/master/manifests/kilosubspace.yaml
```

## Usage

[embedmd]:# (tmp/help.txt)
```txt
Use Kilo as a backend for Subspace

Usage:
  kilosubspace [flags]
  kilosubspace [command]

Available Commands:
  findkey     Find the WireGuard public key corresponding to an endpoint.
  help        Help about any command
  set         Set the WireGuard configuration for a peer.

Flags:
  -h, --help                help for kilosubspace
      --kubeconfig string   Path to kubeconfig. (default "/home/squat/src/onseu2019/assets/auth/kubeconfig")
      --wg string           Path to wg binary.

Use "kilosubspace [command] --help" for more information about a command.
```
