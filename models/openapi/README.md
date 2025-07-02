# GBFS OpenAPI Specification

OpenAPI specifications for the General Bikeshare Feed Specification (GBFS) to easily create models and client code for existing GBFS systems or server code for your own implementation.

## Usage

### Code generation

Many tools will let you generate models or client and server code direcly from the openapi file, e.g.

* [OpenAPI Generator](https://openapi-generator.tech)
* [Swagger Codegen](https://swagger.io/tools/swagger-codegen/)

or one of [https://tools.openapis.org/categories/code-generators.html](these generators).

### Customization

The `openapi.yaml` can be extended with individual details within the `ìnfo` block. It's also recommended to add a `servers` block with your base root url, e.g.

```
servers:
  - url: https://my.gbfs.org/v1
```
If needed, this can be overriden for a specific path.

## Versions
- v2.3
- v3.0
