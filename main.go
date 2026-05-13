// Package summary generates human-readable Markdown summaries of dependency changes.
//
// It transforms dependency diff data from the go-deps-diff library into organized,
// categorized Markdown output with visual indicators for package relationships and
// change types. The output includes collapsible sections, Unicode symbols for quick
// visual scanning, and semantic version tracking.
//
// Main Entry Point:
//
// The GenerateForChanges function is the primary API. It accepts a DiffMap from
// go-deps-diff and produces a Markdown string ready for display, documentation,
// or file storage.
//
// Example:
//
//	changes, err := difflib.Diff(oldPkgs, newPkgs)
//	if err != nil {
//		log.Fatal(err)
//	}
//	markdown := summary.GenerateForChanges(changes, "Go")
//	fmt.Println(markdown)
//
// Helper Functions:
//
// - BuildVersionLabel: Formats versions with semantic version indicators
// - GetPackageSymbol: Returns symbols indicating package type/relationship
// - GetOperationSymbol: Returns symbols for change operations
//
// For more information, see: https://github.com/yoanm/go-deps-diff-summary
package summary

import (
	"fmt"
	"log/slog"
	"strconv"

	"github.com/yoanm/go-deps-diff/contract"

	"github.com/yoanm/go-deps-diff-summary/markdown"
)

const (
	sectionHeaderLevel  = 2
	categoryHeaderLevel = 3
)

// GenerateForChanges produces a Markdown-formatted summary of package dependency changes.
// It takes a DiffMap containing package changes and organizes them into categorized sections
// with appropriate symbols and formatting. The managerName parameter (e.g., "Go" or "Rust")
// is used in the header to identify the package manager context.
//
// Parameters:
//   - changes: A contract.DiffMap from go-deps-diff containing all package changes to summarize
//   - managerName: The name of the package manager (e.g., "Go", "Rust") for header formatting
//
// Returns:
// A Markdown-formatted string with sections for additions, removals, updates, and unchanged packages.
// Changes are visually distinguished using Unicode symbols and organized in collapsible detail sections.
//
// Example:
//
//	changes, err := difflib.Diff(oldPkgs, newPkgs)
//	if err != nil {
//		log.Fatal(err)
//	}
//	markdown := GenerateForChanges(changes, "Go")
//	fmt.Println(markdown)
func GenerateForChanges(changes contract.DiffMap, managerName string) string {
	builder := markdown.NewBuilder()

	builder.Header(fmt.Sprintf("🔎 %s packages 🔍", managerName), 1, 0)

	inOrderMapIteratorHelper[markdownSection, categoriesMap](
		buildSectionsMap(changes),
		getSectionsOrder(),
		func(sectionName markdownSection, categoriesMap categoriesMap) {
			processSection(builder, categoriesMap, sectionName)
		},
	)

	processSummaryCaption(builder)

	return builder.String()
}

func processSection(builder *markdown.Builder, categoriesMap categoriesMap, sectionName markdownSection) {
	slog.Debug("Processing section: " + string(sectionName))
	builder.Header(
		getSectionHeaderFor(sectionName)+"<br/><sub><sup>"+getSectionDescriptionFor(sectionName)+"</sub></sup>",
		sectionHeaderLevel,
		0,
	)

	inOrderMapIteratorHelper[markdownCategory, subCategoriesMap](
		categoriesMap,
		getCategoriesOrder(),
		func(categoryName markdownCategory, subCategoriesMap subCategoriesMap) {
			openedDetails := categoryName == prodUsageCategory &&
				(cautionSection == sectionName || warningSection == sectionName || importantSection == sectionName)

			processCategory(builder, subCategoriesMap, categoryName, openedDetails)
		},
	)

	builder.WriteLine("<hr/>", 0)
}

func processCategory(
	builder *markdown.Builder,
	subCategoriesMap subCategoriesMap,
	categoryName markdownCategory,
	openedDetails bool,
) {
	slog.Debug("Processing category: " + string(categoryName))

	noChangePkgList, otherChangePkgList := splitItemList(subCategoriesMap)

	builder.Header(getCategoryHeaderFor(categoryName), categoryHeaderLevel, 0)
	builder.Details(
		buildSectionSummaryMrk(subCategoriesMap),
		func(builder *markdown.Builder, indentDepth int) {
			if len(otherChangePkgList) > 0 {
				processPkgList(builder, otherChangePkgList, guessShortestPkgRowMode(otherChangePkgList), indentDepth)
			}

			if len(noChangePkgList) > 0 {
				if !openedDetails && len(otherChangePkgList) == 0 {
					// Remove closed details if there is only packages without change
					processPkgList(builder, noChangePkgList, versionOnlyPkgRowMode, indentDepth)
				} else {
					builder.Details(
						"Unchanged packages 🟰<sup>"+strconv.Itoa(len(noChangePkgList))+"</sup>",
						func(builder *markdown.Builder, indentDepth int) {
							processPkgList(builder, noChangePkgList, versionOnlyPkgRowMode, indentDepth)
						},
						false,
						indentDepth,
					)
				}
			}
		},
		openedDetails,
		0,
	)
}

func processPkgList(builder *markdown.Builder, pkgList pkgList, tableMode pkgRowMode, indentDepth int) {
	builder.HTMLTable(
		func(yield func([]string) bool) {
			for _, item := range pkgList {
				if !yield(buildItemMrkRowCells(item, tableMode)) {
					return
				}
			}
		},
		indentDepth,
	)
}

func buildItemMrkRowCells(item *contract.PackageChange, tableMode pkgRowMode) []string {
	cellList := []string{
		buildPackageNameHTMLCell(item.Package),
	}

	pkgVersionCell := buildPackageVersionHTMLCell(item.Package.GetVersion())

	switch tableMode {
	case versionOnlyPkgRowMode:
		cellList = append(cellList, pkgVersionCell)

	case withOperationPkgRowMode:
		operationCell := buildOperationHTMLCell(item.Operation, 0)
		if item.Operation.Name != contract.AdditionOperation {
			cellList = append(cellList, pkgVersionCell, operationCell)
		} else {
			cellList = append(cellList, operationCell, pkgVersionCell)
		}

	case fullPkgRowMode:
		cellList = buildItemMrkFullPkgRowCells(item, cellList, pkgVersionCell)

	default:
		panic("Unmanaged table mode:" + strconv.Itoa(int(tableMode)))
	}

	return cellList
}

func buildItemMrkFullPkgRowCells(item *contract.PackageChange, cellList []string, pkgVersionCell string) []string {
	if item.Operation.Name != contract.AdditionOperation { // Version will be printed at the end for added package !
		switch item.Operation.Name {
		case contract.UnknownUpdateOperation, contract.UpgradeOperation, contract.DowngradeOperation:
			cellList = append(cellList, buildPackageVersionHTMLCell(item.PreviousVersion))
		default:
			cellList = append(cellList, pkgVersionCell)
		}
	}

	colspan := 0
	if contract.NoChangeOperation == item.Operation.Name ||
		contract.AdditionOperation == item.Operation.Name ||
		contract.RemovalOperation == item.Operation.Name {
		colspan = 2
	}

	cellList = append(cellList, buildOperationHTMLCell(item.Operation, colspan))

	switch item.Operation.Name { //nolint:exhaustive // Only those cases should be handled here !
	case contract.AdditionOperation, contract.UnknownUpdateOperation, contract.UpgradeOperation, contract.DowngradeOperation: //nolint:lll // Meaningless here
		cellList = append(cellList, pkgVersionCell)
	}

	return cellList
}

func buildOperationHTMLCell(operation contract.Operation, colspan int) string {
	opColspanDirective := ""
	if colspan > 1 {
		opColspanDirective = fmt.Sprintf(" colspan=\"%d\"", colspan)
	}

	return "<td align=\"center\"" + opColspanDirective + ">" + GetOperationSymbol(operation) + "</td>"
}

func buildPackageVersionHTMLCell(version contract.PkgVersion) string {
	return "<td align=\"right\">" + BuildVersionLabel(version) + "</td>"
}

func buildPackageNameHTMLCell(pkg contract.PkgWrapper) string {
	return "<td align=\"left\">" + buildPackageLabel(pkg) + "</td>"
}
