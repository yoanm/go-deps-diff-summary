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
