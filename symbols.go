package summary

import "github.com/yoanm/go-deps-diff/contract"

const (
	AbandonedSymbol = "💀"
	NonSemverSymbol = "❗"
)

const (
	RootRequirementSymbol      = "🗄️"
	RootDevRequirementSymbol   = "🧰"
	TransitiveDependencySymbol = "🔗️"
)

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

// GetPackageSymbol returns a Unicode symbol indicating the package type and relationship.
// The symbol communicates whether the package is a root requirement (🗄️), root dev requirement (🧰),
// or a transitive dependency (🔗️).
//
// Parameters:
//   - pkg: The package to determine the symbol for
//
// Returns:
// A Unicode emoji string representing the package type (🗄️, 🧰, or 🔗️).
//
// Example:
//
//	symbol := GetPackageSymbol(pkg)
//	fmt.Println(symbol) // Output: 🗄️ (for root requirement)
func GetPackageSymbol(pkg contract.PkgWrapper) string {
	switch {
	case pkg.IsRootRequirement():
		return RootRequirementSymbol
	case pkg.IsRootDevRequirement():
		return RootDevRequirementSymbol
	default:
		return TransitiveDependencySymbol
	}
}

// GetOperationSymbol returns a Unicode symbol representing a package change operation.
// The symbol communicates the type of change: additions (➕), removals (❌), upgrades (🔺),
// downgrades (🔻), patch updates (🔹), or no change (🟰). For semver updates, nested symbols
// indicate which component changed (major/minor/patch).
//
// Parameters:
//   - operation: The package operation/change to determine the symbol for
//
// Returns:
// A Unicode emoji string (possibly in HTML sub/sup tags) representing the operation type.
//
// Example:
//
//	symbol := GetOperationSymbol(contract.Operation{
//		Name: contract.UpgradeOperation,
//		SemverType: contract.SemverMajorUpdate,
//	})
//	fmt.Println(symbol) // Output: <sub><sup>🔺.🔹.🔹</sup></sub>
func GetOperationSymbol(operation contract.Operation) string {
	switch operation.Name {
	case contract.UnknownUpdateOperation:
		if operation.SemverType == contract.SemverExtraUpdate {
			return "<sub><sup>" + SemverExtraUpdateOperationSymbol + "</sup></sub>"
		}

		return UnknownOperationSymbol
	case contract.UpgradeOperation:
		return getUpdateOperationSymbol(operation, false)
	case contract.DowngradeOperation:
		return getUpdateOperationSymbol(operation, true)
	case contract.RemovalOperation:
		return RemovalOperationSymbol
	case contract.AdditionOperation:
		return AdditionOperationSymbol
	case contract.NoChangeOperation:
		return NonChangeOperationSymbol
	}

	return UnmanagedSymbol
}

func getUpdateOperationSymbol(operation contract.Operation, isDowngrade bool) string {
	emojiUpdated := SemverComponentUpgradeSymbol
	if isDowngrade {
		emojiUpdated = SemverComponentDowngradeSymbol
	}

	switch operation.SemverType { //nolint:exhaustive // Only those cases can be managed, fallback to unknown otherwise
	case contract.SemverMajorUpdate:
		return "<sub><sup>" + emojiUpdated + "." + SemverComponentSymbol + "." + SemverComponentSymbol + "</sup></sub>"
	case contract.SemverMinorUpdate:
		return "<sub><sup>" + SemverComponentSymbol + "." + emojiUpdated + "." + SemverComponentSymbol + "</sup></sub>"
	case contract.SemverPatchUpdate:
		return "<sub><sup>" + SemverComponentSymbol + "." + SemverComponentSymbol + "." + emojiUpdated + "</sup></sub>"
	}

	return UnmanagedSymbol
}
