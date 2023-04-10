# Chroma Client for Go
[![License - MIT](https://img.shields.io/github/license/CSXL/solus?style=for-the-badge)](LICENSE)
![GitHub Actions Build Status](https://img.shields.io/github/actions/workflow/status/CSXL/go-chroma/push.yml?logo=github&style=for-the-badge)
[![Package Documentation](https://img.shields.io/badge/-reference-blue?labelColor=grey&style=for-the-badge&logo=go)](https://pkg.go.dev/github.com/CSXL/go-chroma)

This is a generated client for the [Chroma](https://github.com/chroma-core/chroma)
embeddings database. It was generated using Deepmap's [oapi-codegen](github.com/deepmap/oapi-codegen).

See the [Go Packages](https://pkg.go.dev/github.com/CSXL/go-chroma) page for the generated Go client documentation. See [Chroma's website](https://www.trychroma.com/) for more information on their project.

## Generate command
```bash
oapi-codegen -package chroma -generate types,client -o chroma.gen.go openapi_spec.json
```
