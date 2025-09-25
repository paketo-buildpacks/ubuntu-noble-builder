# Description

This repository develops, maintains and publishes the following builders.

## Paketo Noble Builder

### `paketobuildpacks/ubuntu-noble-builder`

This builder uses the `paketobuildpacks/ubuntu-noble-build` as build image and `paketobuildpacks/ubuntu-noble-run` as a run image from [Paketo Noble Base Images](https://github.com/paketo-buildpacks/ubuntu-noble-base-images) with buildpacks for Java, Java Native Image, Go, Python, .NET, Node.js, Apache HTTPD, NGINX and Procfile.

To see which versions of build and run images, buildpacks, and the lifecycle that are contained within a given builder version, see the [Releases](https://github.com/paketo-buildpacks/ubuntu-noble-builder/releases) on this repository. This information is also available in the `builders/builder/builder.toml`.

To use this builder, you dont have to specify any buildpacks as they are already included on the `builder.toml`

For example, with the `pack` CLI:

```bash
pack build my-app --builder paketobuildpacks/ubuntu-noble-builder:latest
```

## Paketo Noble Buildpackless Builder

### `paketobuildpacks/ubuntu-noble-builder-buildpackless:latest`

This builder uses the `paketobuildpacks/ubuntu-noble-build` as build image and `paketobuildpacks/ubuntu-noble-run` as a run image from [Paketo Noble Base Images](https://github.com/paketo-buildpacks/ubuntu-noble-base-images) and contains **no buildpacks nor order groups**. To use this builder, you must specify buildpacks at build time using whatever mechanisms your CNB platform of choice offers.

For example, with the `pack` CLI, use `--buildpack` as follows:

```bash
pack build dotnet-with-buildpackless-builder \
--buildpack paketobuildpacks/dotnet-core \
--builder paketobuildpacks/ubuntu-noble-builder-buildpackless:latest
```

To see which versions of build and run images and the lifecycle are contained
within a given builder version, see the
[Releases](https://github.com/paketo-buildpacks/ubuntu-noble-base-images/releases)
on this repo. This information is also available in the `builder-buildpackless/builder.toml`.

For more information about this builder and how to use it, visit the [Paketo
builder documentation](https://paketo.io/docs/builders/). To learn about the
stack included in this builder, visit the [Paketo stack
documentation](https://paketo.io/docs/stacks/).

## More Information

For more information about these builders and how to use them, visit the [Paketo builder documentation](https://paketo.io/docs/builders/).

## Run Tests

To run all smoke tests, run:

```bash
./scripts/smoke.sh
```

## Publish a builder

To publish a builder in a registry, run:

```bash
./scripts/publish \
    --builder-toml-path ./path/to/builder.toml \
    --builder-image-ref <registry-url>:<registry-port>/<builder-image-name>:<image-tag>
```
