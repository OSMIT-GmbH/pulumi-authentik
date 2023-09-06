# authentik Resource Provider

The authentik Resource Provider lets you manage [authentik](https://https://goauthentik.io/) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @osmit-gmbh/pulumi-authentik
```

or `yarn`:

```bash
yarn add @osmit-gmbh/pulumi-authentik
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/OSMIT-GmbH/pulumi-authentik/sdk/go/...
```

## Configuration

The following configuration points are available for the `authentik` provider:

- `authentik:apiKey` (environment: `authentik_API_KEY`) - the API key for `authentik`
- `authentik:region` (environment: `authentik_REGION`) - the region in which to deploy resources

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/authentik/api-docs/).
