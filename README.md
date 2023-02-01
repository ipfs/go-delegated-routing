# ⚠️ This repository has been moved to https://github.com/ipfs/go-libipfs/tree/main/routing/http.

go-delegated-routing
=======================

> Delegated routing Client and Server over Reframe RPC

This package provides delegated routing implementation in Go:
- Client (for IPFS nodes like [Kubo](https://github.com/ipfs/kubo/blob/master/docs/config.md#routingrouters-parameters)),
- Server (for public indexers such as https://cid.contact)

## Documentation

- Go docs: https://pkg.go.dev/github.com/ipfs/go-delegated-routing
- What is Reframe? https://blog.ipfs.tech/2022-09-02-introducing-reframe/
  - Reframe Specs: https://github.com/ipfs/specs/blob/main/reframe/

## Lead Maintainer

🦗🎶

## Generating

Client and Server code can be (re-)generated via:

```console
go generate -v ./...
```

## Contributing

Contributions are welcome! This repository is part of the IPFS project and therefore governed by our [contributing guidelines](https://github.com/ipfs/community/blob/master/CONTRIBUTING.md).

## License

[SPDX-License-Identifier: Apache-2.0 OR MIT](LICENSE.md)
