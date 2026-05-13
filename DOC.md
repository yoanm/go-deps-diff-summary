# summary

Package summary generates human-readable Markdown summaries of dependency changes.

It transforms dependency diff data from the go-deps-diff library into organized,
categorized Markdown output with visual indicators for package relationships and
change types. The output includes collapsible sections, Unicode symbols for quick
visual scanning, and semantic version tracking.

The GenerateForChanges function is the primary API. It accepts a DiffMap from
go-deps-diff and produces a Markdown string ready for display, documentation,
or file storage.

Example:

```go
changes, err := difflib.Diff(oldPkgs, newPkgs)
if err != nil {
	log.Fatal(err)
}
markdown := summary.GenerateForChanges(changes, "Composer")
fmt.Println(markdown)
```

Helper Functions:

- BuildVersionLabel: Formats versions with semantic version indicators
- GetPackageSymbol: Returns symbols indicating package type/relationship
- GetOperationSymbol: Returns symbols for change operations

## Constants

```golang
const (
    AbandonedSymbol = "💀"
    NonSemverSymbol = "❗"
)
```

```golang
const (
    RootRequirementSymbol      = "🗄️"
    RootDevRequirementSymbol   = "🧰"
    TransitiveDependencySymbol = "🔗️"
)
```

```golang
const (
    SemverComponentUpgradeSymbol     = "🔺"
    SemverComponentDowngradeSymbol   = "🔻"
    SemverComponentSymbol            = "🔹"
    UnknownOperationSymbol           = "❓"
    SemverExtraUpdateOperationSymbol = SemverComponentSymbol +
        "." + SemverComponentSymbol +
        "." + SemverComponentSymbol +
        UnknownOperationSymbol
    RemovalOperationSymbol   = "❌"
    AdditionOperationSymbol  = "➕️"
    NonChangeOperationSymbol = "🟰"
    UnmanagedSymbol          = "❔"
)
```

## Functions

### func [BuildVersionLabel](/labels.go#L44)

`func BuildVersionLabel(version contract.PkgVersion) string`

BuildVersionLabel formats a package version into a display label.
If the version is not semantic, it appends a NonSemverSymbol (❗) indicator.

Parameters:

```go
- version: The package version to format
```

Returns:
A formatted string representation of the version, potentially with a symbol appended
to indicate non-semantic versioning (e.g., git hashes, custom versions).

Example:

```go
label := BuildVersionLabel(contract.PkgVersion{
	Label: "1.2.3",
	Semver: &contract.Semver{...},
})
fmt.Println(label) // Output: 1.2.3
```

### func [GenerateForChanges](/main.go#L64)

`func GenerateForChanges(changes contract.DiffMap, managerName string) string`

GenerateForChanges produces a Markdown-formatted summary of package dependency changes.
It takes a DiffMap containing package changes and organizes them into categorized sections
with appropriate symbols and formatting. The managerName parameter (e.g., "Composer" or "Npm")
is used in the header to identify the package manager context.

Parameters:

```diff
- changes: A contract.DiffMap from go-deps-diff containing all package changes to summarize
- managerName: The name of the package manager (e.g., "Composer", "Npm") for header formatting
```

Returns:
A Markdown-formatted string with sections for additions, removals, updates, and unchanged packages.
Changes are visually distinguished using Unicode symbols and organized in collapsible detail sections.

Example:

```go
changes, err := difflib.Diff(oldPkgs, newPkgs)
if err != nil {
	log.Fatal(err)
}
markdown := GenerateForChanges(changes, "Composer")
fmt.Println(markdown)
```

### func [GetOperationSymbol](/symbols.go#L74)

`func GetOperationSymbol(operation contract.Operation) string`

GetOperationSymbol returns a Unicode symbol representing a package change operation.
The symbol communicates the type of change: additions (➕), removals (❌), upgrades (🔺),
downgrades (🔻), patch updates (🔹), or no change (🟰). For semver updates, nested symbols
indicate which component changed (major/minor/patch).

Parameters:

```go
- operation: The package operation/change to determine the symbol for
```

Returns:
A Unicode emoji string (possibly in HTML sub/sup tags) representing the operation type.

Example:

```go
symbol := GetOperationSymbol(contract.Operation{
	Name: contract.UpgradeOperation,
	SemverType: contract.SemverMajorUpdate,
})
fmt.Println(symbol) // Output: <sub><sup>🔺.🔹.🔹</sup></sub>
```

### func [GetPackageSymbol](/symbols.go#L45)

`func GetPackageSymbol(pkg contract.PkgWrapper) string`

GetPackageSymbol returns a Unicode symbol indicating the package type and relationship.
The symbol communicates whether the package is a root requirement (🗄️), root dev requirement (🧰),
or a transitive dependency (🔗️).

Parameters:

```go
- pkg: The package to determine the symbol for
```

Returns:
A Unicode emoji string representing the package type (🗄️, 🧰, or 🔗️).

Example:

```go
symbol := GetPackageSymbol(pkg)
fmt.Println(symbol) // Output: 🗄️ (for root requirement)
```

## Sub Packages

* [markdown](./markdown): Package markdown provides a builder pattern for generating Markdown and HTML documents with structured formatting, indentation support, and collapsible sections.

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
