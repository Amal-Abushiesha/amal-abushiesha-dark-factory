# amal-abushiesha-dark-factory

This repository starts as a Go module. Run the pinned repository checks with:

```sh
task gate
```

The gate builds and analyzes all Go packages, scans them with govulncheck, and
runs race-enabled tests. It requires at least 90% statement coverage for Go
packages. The starter build metadata package has 100% coverage.
