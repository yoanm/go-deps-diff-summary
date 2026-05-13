# summary<br/><sup><sub>`depsdiff` summary library</sub></sup>

[![License](https://img.shields.io/github/license/yoanm/go-deps-diff-summary.svg)](https://github.com/yoanm/go-deps-diff-summary)
[![Code size](https://img.shields.io/github/languages/code-size/yoanm/go-deps-diff-summary.svg)](https://github.com/yoanm/go-deps-diff-summary)
[![Go Reference](https://pkg.go.dev/badge/github.com/yoanm/go-deps-diff-summary.svg)](https://pkg.go.dev/github.com/yoanm/go-deps-diff-summary)

![Dependabot Status](https://flat.badgen.net/github/dependabot/yoanm/go-deps-diff-summary)
![Last commit](https://badgen.net/github/last-commit/yoanm/go-deps-diff-summary)

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/ebeacd3a91a74fef8a8ed4ea879ede72)](https://app.codacy.com/gh/yoanm/go-deps-diff-summary/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/yoanm/go-deps-diff-summary?)](https://goreportcard.com/report/github.com/yoanm/go-deps-diff-summary)

[![CI](https://github.com/yoanm/go-deps-diff-summary/actions/workflows/CI.yml/badge.svg?branch=master)](https://github.com/yoanm/go-deps-diff-summary/actions/workflows/CI.yml)
[![codecov](https://codecov.io/gh/yoanm/go-deps-diff-summary/branch/master/graph/badge.svg?token=NHdwEBUFK5)](https://codecov.io/gh/yoanm/go-deps-diff-summary)

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/yoanm/go-deps-diff-summary)


A Go library for generating human-readable Markdown summaries of dependency changes detected by the [go-deps-diff](https://github.com/yoanm/go-deps-diff) library.

## Overview

`summary` is a Go module that transforms dependency diff data into formatted Markdown output. It takes a `DiffMap` from go-deps-diff and produces organized, categorized summaries with visual indicators for package relationships and change types.

## Installation

```bash
go get github.com/yoanm/go-deps-diff-summary
```

## Quick Start

The library provides one main entry point for generating summaries:

### Generate a Markdown Summary

```go
package main

import (
    "fmt"
    "log"

    "github.com/yoanm/go-deps-diff/contract"
    summary "github.com/yoanm/go-deps-diff-summary"
)

func main() {
    // Assuming you have a DiffMap from go-deps-diff
    // This would typically come from comparing two dependency trees
    changes := contract.DiffMap{
        "package-name": &contract.PackageChange{
            Operation: contract.Operation{
                Name: contract.UpgradeOperation,
                SemverType: contract.SemverMinorUpdate,
            },
            Package: /* package info */,
            PreviousVersion: contract.PkgVersion{Label: "1.0.0"},
        },
    }
    
    // Generate Markdown summary
    markdown := summary.GenerateForChanges(changes, "Composer")
    
    // Output or save the markdown
    fmt.Println(markdown)
}
```

## Features

- ✅ Generates structured Markdown with collapsible sections
- ✅ Organizes changes by type (additions, removals, updates, unchanged)
- ✅ Categorizes packages by role (prod requirements, dev requirements, transitive dependencies)
- ✅ Uses Unicode symbols for quick visual identification
- ✅ Detects and marks abandoned packages
- ✅ Tracks semantic version changes (major/minor/patch updates vs downgrades)
- ✅ Supports non-semantic versions (commit hashes, custom versions)
- ✅ Generates HTML tables for detailed package information
- ✅ Integrates seamlessly with go-deps-diff output
- ✅ Comprehensive error handling through go-deps-diff contracts

## How It Works

1. **Input**: Accepts a `contract.DiffMap` from go-deps-diff containing all detected changes
2. **Categorization**: Organizes packages by change type (additions, removals, updates, unchanged) and category (prod, dev, transitive)
3. **Formatting**: Applies visual styling with symbols, colors, and hierarchical structure
4. **Output**: Returns a single Markdown string ready for display, documentation, or file storage

## API Reference

### Main Functions

#### `GenerateForChanges(changes contract.DiffMap, managerName string) string`

Generates a Markdown summary of package changes.

**Parameters:**
- `changes`: A `contract.DiffMap` containing package changes from go-deps-diff
- `managerName`: Name of the package manager (e.g., "Composer", "Npm") used in headers

**Returns:**
- `string`: Formatted Markdown summary with organized sections and symbols

**Example:**
```go
markdown := summary.GenerateForChanges(changes, "Composer")
fmt.Println(markdown)
```

#### `BuildVersionLabel(version contract.PkgVersion) string`

Formats a package version into a display label with appropriate indicators.

**Parameters:**
- `version`: A `contract.PkgVersion` to format

**Returns:**
- `string`: Formatted version label, with `❗` appended if non-semantic

**Example:**
```go
label := summary.BuildVersionLabel(pkgVersion)
// May return: "1.2.3" or "abc1234❗" for non-semantic versions
```

#### `GetPackageSymbol(pkg contract.PkgWrapper) string`

Returns a Unicode symbol indicating the package type.

**Parameters:**
- `pkg`: A `contract.PkgWrapper` representing the package

**Returns:**
- `string`: Unicode emoji - `🗄️` (root requirement), `🧰` (root dev), or `🔗️` (transitive)

**Example:**
```go
symbol := summary.GetPackageSymbol(pkg)
// Returns one of: 🗄️, 🧰, 🔗️
```

#### `GetOperationSymbol(operation contract.Operation) string`

Returns a Unicode symbol representing the change operation.

**Parameters:**
- `operation`: A `contract.Operation` describing the change

**Returns:**
- `string`: Unicode emoji representing the operation (may include HTML sub/sup tags)

**Symbol Reference:**
- `➕` - Package added
- `❌` - Package removed
- `🔺` - Major version upgrade
- `🔻` - Major version downgrade
- `🔹.🔺.🔹` - Minor version upgrade
- `🔹.🔹.🔺` - Patch version upgrade
- `🟰` - No change
- `❓` - Unknown operation

**Example:**
```go
symbol := summary.GetOperationSymbol(operation)
// May return: ➕, ❌, or nested version update symbols
```

## Symbol Reference

The library uses Unicode symbols to represent package states and changes visually:

### Package Type Symbols
- `🗄️` - Root requirement (explicitly required in project)
- `🧰` - Root dev requirement (explicitly required for development)
- `🔗️` - Transitive dependency (required by other packages)
- `💀` - Abandoned package (attached to package name)

### Operation Symbols
- `➕` - Addition (new package)
- `❌` - Removal (package deleted)
- `🔺` - Major/Minor/Patch version change (upgrade)
- `🔻` - Major/Minor/Patch version change (downgrade)
- `🟰` - No change
- `❓` - Unknown operation
- `❔` - Unmanaged case

## Examples

### Generate and Save Markdown

```go
changes, err := difflib.Diff(oldDeps, newDeps)
if err != nil {
    log.Fatal(err)
}

// Generate summary
markdown := summary.GenerateForChanges(changes, "Composer")

// Save to file
err = os.WriteFile("dependency-changes.md", []byte(markdown), 0644)
if err != nil {
    log.Fatal(err)
}

fmt.Println("Summary saved to dependency-changes.md")
```

## Subpackages

### `markdown`

The `markdown` package provides internal utilities for building Markdown output:

- `Builder` - Main builder for constructing Markdown documents
- `NewBuilder()` - Creates a new Builder instance
- Support for headers, tables, details/collapsible sections, and HTML elements

This is a package used by the summary module to format output.

## Testing

```bash
make test
```

## Related Projects

- [go-deps-diff](https://github.com/yoanm/go-deps-diff) - Generic dependency comparison library that powers this tool
- [go-composer-diff](https://github.com/yoanm/go-composer-diff) - PHP Composer comparison library

## License

See LICENSE file for details.
