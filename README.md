# Semaphore UI Charts

[![Release Build](https://github.com/semaphoreui/charts/actions/workflows/release.yml/badge.svg)](https://github.com/semaphoreui/charts/actions/workflows/release.yml) [![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/cloudhippie)](https://artifacthub.io/packages/search?repo=cloudhippie)

Definition and publishing of Helm charts to install tools used by Cloudhippie or
tools built within our GitHub organization.

## Usage

Make sure you have installed [Helm][helm], after that you can install the charts
repository and search for available charts:

```console
helm repo add cloudhippie https://cloudhippie.github.io/charts
helm search repo cloudhippie
```

## Security

If you find a security issue please contact
[security@semaphoreui.com](mailto:security@semaphoreui.com) first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

-   [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2023 Cloudhippie <info@cloudhippie.de>
```

[helm]: https://helm.sh
