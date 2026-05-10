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

func BuildVersionLabel(version contract.PkgVersion) string {
	if version.Semver == nil {
		return version.Label + NonSemverSymbol
	}

	return version.Label
}
