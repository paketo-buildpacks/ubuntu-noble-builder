# Paketo Noble Buildpackless Builder

## `paketobuildpacks/ubuntu-noble-builder:latest`

This builder uses the [Paketo Noble Stack](https://github.com/paketo-buildpacks/ubuntu-noble-base-images) (an Ubuntu Noble
base image) and contains **no buildpacks nor order groups**. To use this
builder, you must specify buildpacks at build time using whatever mechanisms
your CNB platform of choice offers.

For example, with the `pack` CLI, use `--buildpack` as follows:
```bash
pack build dotnet-with-buildpackless-builder \
--buildpack index.docker.io/paketo-buildpacks/dotnet-core \
--builder paketobuildpacks/ubuntu-noble-builder-buildpackless:latest
```

To see which versions of build and run images and the lifecycle are contained
within a given builder version, see the
[Releases](https://github.com/paketo-buildpacks/ubuntu-noble-base-images/releases)
on this repo. This information is also available in the `builder-buildpackless.toml`.

For more information about this builder and how to use it, visit the [Paketo
builder documentation](https://paketo.io/docs/builders/).  To learn about the
stack included in this builder, visit the [Paketo stack
documentation](https://paketo.io/docs/stacks/).

