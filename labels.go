package summary

import (
	"strings"

	"github.com/yoanm/go-deps-diff/contract"
)

func buildPackageLabel(pkg contract.PkgWrapper) string {
	builder := strings.Builder{}
	// Prepend package type symbol
	builder.WriteString("<sup>" + GetPackageSymbol(pkg) + "</sup>")
	// Add link if available
	if pkg.GetLink() != "" {
		builder.WriteString("<a href=\"" + pkg.GetLink() + "\">" + pkg.GetName() + "</a>")
	} else {
		builder.WriteString(pkg.GetName())
	}
	// Managed abandoned package special case
	if pkg.IsAbandoned() {
		builder.WriteString(AbandonedSymbol)
	}

	return builder.String()
}

// BuildVersionLabel formats a package version into a display label.
// If the version is not semantic, it appends a NonSemverSymbol (❗) indicator.
//
// Parameters:
//   - version: The package version to format
//
// Returns:
// A formatted string representation of the version, potentially with a symbol appended
// to indicate non-semantic versioning (e.g., git hashes, custom versions).
//
// Example:
//
//	label := BuildVersionLabel(contract.PkgVersion{
//		Label: "1.2.3",
//		Semver: &contract.Semver{...},
//	})
//	fmt.Println(label) // Output: 1.2.3
func BuildVersionLabel(version contract.PkgVersion) string {
	if version.Semver == nil {
		return version.Label + NonSemverSymbol
	}

	return version.Label
}
