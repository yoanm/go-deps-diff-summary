# summary

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

### func [BuildVersionLabel](/labels.go#L27)

`func BuildVersionLabel(version contract.PkgVersion) string`

### func [GenerateForChanges](/main.go#L18)

`func GenerateForChanges(changes contract.DiffMap, managerName string) string`

### func [GetOperationSymbol](/symbols.go#L42)

`func GetOperationSymbol(operation contract.Operation) string`

### func [GetPackageSymbol](/symbols.go#L31)

`func GetPackageSymbol(pkg contract.PkgWrapper) string`

## Sub Packages

* [markdown](./markdown)

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
