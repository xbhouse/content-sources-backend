---
type: pull_request
number: 819
title: "Build: Update pre-commit hook golangci-lint to v1.61.0"
state: merged
author: red-hat-konflux
created: 2024-09-17T12:23:48Z
updated: 2024-09-18T13:41:22Z
closed: 2024-09-18T13:41:18Z
merged: 2024-09-18T13:41:18Z
base_branch: main
head_branch: konflux/mintmaker/main/golangci-golangci-lint-1.x
labels: []
url: https://github.com/content-services/content-sources-backend/pull/819
---

# Pull Request #819: Build: Update pre-commit hook golangci-lint to v1.61.0

**Author**: @red-hat-konflux
**Created**: September 17, 2024 at 12:23 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/golangci-golangci-lint-1.x`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [golangci/golangci-lint](https://togithub.com/golangci/golangci-lint) | repository | minor | `v1.43.0` -> `v1.61.0` |

Note: The `pre-commit` manager in Renovate is not supported by the `pre-commit` maintainers or community. Please do not report any problems there, instead [create a Discussion in the Renovate repository](https://togithub.com/renovatebot/renovate/discussions/new) if you have any questions.

---

### Release Notes

<details>
<summary>golangci/golangci-lint (golangci/golangci-lint)</summary>

### [`v1.61.0`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1610)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.60.3...v1.61.0)

1.  Enhancements
    -   Add `junit-xml-extended` format
    -   Exclude Swagger Codegen files by default
2.  Updated linters
    -   `dupword`: from 0.0.14 to 0.1.1
    -   `fatcontext`: from 0.4.0 to 0.5.2
    -   `gci`: from 0.13.4 to 0.13.5 (new option `no-lex-order`)
    -   `go-ruleguard`: from 0.4.2 to [`0fe6f58`](https://togithub.com/golangci/golangci-lint/commit/0fe6f58b47b1) (fix panic with custom linters)
    -   `godot`: from 1.4.16 to 1.4.17
    -   `gomodguard`: from 1.3.3 to 1.3.5
    -   `gosec`: disable temporarily `G407`
    -   `gosec`: from [`ab3f6c1`](https://togithub.com/golangci/golangci-lint/commit/ab3f6c1c83a0) to 2.21.2 (partially fix `G115`)
    -   `intrange`: from 0.1.2 to 0.2.0
    -   `nolintlint`: remove the empty line in the directive replacement
3.  Misc.
    -   Improve runtime version parsing
4.  Documentation
    -   Add additional info about `typecheck`

### [`v1.60.3`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1603)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.60.2...v1.60.3)

1.  Updated linters
    -   `gosec`: from [`81cda2f`](https://togithub.com/golangci/golangci-lint/commit/81cda2f91fbe) to [`ab3f6c1`](https://togithub.com/golangci/golangci-lint/commit/ab3f6c1c83a0) (fix `G115` false positives)
2.  Misc.
    -   Check that the Go version use to build is greater or equals to the Go version of the project

### [`v1.60.2`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1602)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.60.1...v1.60.2)

1.  Updated linters

-   `gofmt`: update to HEAD (go1.22)
-   `gofumpt`: from 0.6.0 to 0.7.0
-   `gosec`: fix G602 analyzer
-   `gosec`: from [`5f0084e`](https://togithub.com/golangci/golangci-lint/commit/5f0084eb01a9) to [`81cda2f`](https://togithub.com/golangci/golangci-lint/commit/81cda2f91fbe) (adds `G115`, `G405`, `G406`, `G506`, `G507`)
-   `staticcheck`: from 0.5.0 to 0.5.1
-   `staticcheck`: propagate Go version
-   `wrapcheck`: from 2.8.3 to 2.9.0
-   ⚠️ `exportloopref`: deprecation

### [`v1.60.1`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1601)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.60.0...v1.60.1)

1.  Updated linters
    -   `errorlint`: from 1.5.2 to 1.6.0
    -   `exhaustruct`: from 3.2.0 to 3.3.0 (recognize custom error values in return)
    -   `fatcontext`: from 0.2.2 to 0.4.0 (fix false positives for context stored in structs)
    -   `gocognit`: from 1.1.2 to 1.1.3
    -   `gomodguard`: from 1.3.2 to 1.3.3
    -   `govet` (`printf`): report non-constant format, no args
    -   `lll`: advertise max line length instead of just reporting failure
    -   `revive`: from 1.3.7 to 1.3.9 (new rule: `comments-density`)
    -   `sloglint`: from 0.7.1 to 0.7.2
    -   `spancheck`: from 0.6.1 to 0.6.2
    -   `staticcheck`: from 0.4.7 to 0.5.0
    -   `tenv`: from 1.7.1 to 1.10.0 (remove reports on fuzzing)
    -   `testifylint`: from 1.3.1 to 1.4.3 (new options: `formatter`, `suite-broken-parallel`, `suite-subtest-run`)
    -   `tparallel`: from 0.3.1 to 0.3.2
    -   `usestdlibvars`: from 1.26.0 to 1.27.0 (fix false-positive with number used inside a mathematical operations)
    -   `wsl`: from 4.2.1 to 4.4.1
    -   ️⚠️ `unused`: remove `exported-is-used` option
2.  Fixes
    -   SARIF: sanitize level property
    -   ️⚠️ `typecheck` issues should never be ignored
3.  Documentation
    -   Add link on linter without configuration
    -   Remove 'trusted by' page
    -   `wsl` update documentation of the configuration
4.  misc.
    -   🎉 go1.23 support

### [`v1.60.0`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1600)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.59.1...v1.60.0)

Cancelled due to a CI problem.

### [`v1.59.1`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1591)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.59.0...v1.59.1)

1.  Updated linters
    -   `errorlint`: from 1.5.1 to 1.5.2
    -   `gomnd`: deprecated configuration compatibility
    -   `intrange`: add `style` preset
    -   `misspell`: from 0.5.1 to 0.6.0
    -   `sloglint`: from 0.7.0 to 0.7.1
    -   `testifylint`: from 1.3.0 to 1.3.1
    -   `unparam`: bump to HEAD
    -   `usestdlibvars`: from 1.25.0 to 1.26.0
2.  Fixes
    -   SARIF: init empty result slice
    -   SARIF: issue column >= 1
3.  Documentation
    -   `revive`: update documentation of the configuration

### [`v1.59.0`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1590)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.58.2...v1.59.0)

1.  Enhancements
    -   Add SARIF output format
    -   Allow the analysis of generated files (`issues.exclude-generated: disable`)
2.  Updated linters
    -   `errcheck`: fix deprecation warning
    -   `go-critic`: from 0.11.3 to 0.11.4
    -   `gosec`: from 2.20.0 to [`5f0084e`](https://togithub.com/golangci/golangci-lint/commit/5f0084eb01a9) (fix G601 and G113 performance issues)
    -   `sloglint`: from 0.6.0 to 0.7.0 (new option `forbidden-keys`)
    -   `testifylint`: from 1.2.0 to 1.3.0 (new checker `negative-positive` and new option `go-require.ignore-http-handlers`)
3.  Misc.
    -   ️️⚠️ Deprecate `github-action` output format
    -   ️️⚠️ Deprecate `issues.exclude-generated-strict` option (replaced by `issues.exclude-generated: strict`)
    -   ️️⚠️ Add warning about disabled and deprecated linters (level 2)

### [`v1.58.2`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1582)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.58.1...v1.58.2)

1.  Updated linters
    -   `canonicalheader`: from 1.0.6 to 1.1.1
    -   `gosec`: from 2.19.0 to 2.20.0
    -   `musttag`: from 0.12.1 to 0.12.2
    -   `nilnil`: from 0.1.8 to 0.1.9
2.  Documentation
    -   Improve integrations and install pages

### [`v1.58.1`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1581)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.58.0...v1.58.1)

1.  Updated linters
    -   `tagalign`: from 1.3.3 to 1.3.4
    -   `protogetter`: from 0.3.5 to 0.3.6
    -   `gochecknoinits`: fix analyzer name
2.  Fixes
    -   Restores previous `gihub-actions` output format (removes GitHub Action problem matchers)

### [`v1.58.0`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1580)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.57.2...v1.58.0)

1.  New linters
    -   `fatcontext`: https://github.com/Crocmagnon/fatcontext
    -   `canonicalheader`: https://github.com/lasiar/canonicalheader
2.  Updated linters
    -   `copyloopvar`: from 1.0.10 to 1.1.0 (`ignore-alias` is replaced by `check-alias` with the opposite behavior)
    -   `decorder`: from 0.4.1 to 0.4.2
    -   `errname`: from 0.1.12 to 0.1.13
    -   `errorlint`: from 1.4.8 to 1.5.1 (new options `allowed-errors` and `allowed-errors-wildcard`)
    -   `execinquery`: deprecate linter ⚠️
    -   `gci`: from 0.12.3 to 0.13.4 (new section `localModule`)
    -   `gocritic`: from 0.11.2 to 0.11.3
    -   `spancheck`: from 0.5.3 to 0.6.1
    -   `goerr113` is replaced by `err113` ⚠️
    -   `gomnd` is replaced by `mnd` ⚠️
    -   `gomodguard`: from 1.3.1 to 1.3.2
    -   `grouper`: from 1.1.1 to 1.1.2
    -   `intrange`: from 0.1.1 to 0.1.2
    -   `mirror`: from 1.1.0 to 1.2.0
    -   `misspell`: from 0.4.1 to 0.5.1
    -   `musttag`: from 0.9.0 to 0.12.1
    -   `nilnil`: from 0.1.7 to 0.1.8
    -   `nonamedreturns`: from 1.0.4 to 1.0.5
    -   `promlinter`: from 0.2.0 to 0.3.0
    -   `sloglint`: from 0.5.0 to 0.6.0
    -   `unparam`: bump to HEAD ([`063aff9`](https://togithub.com/golangci/golangci-lint/commit/063aff900ca150b80930c8de76f11d7e6488222f))
    -   `whitespace`: from 0.1.0 to 0.1.1
3.  Enhancements
    -   Speed up "fast" linters when only "fast" linters are run: between 40% and 80% faster at first run (i.e. without cache)
4.  Fixes
    -   Use version with module plugins
    -   Skip `go.mod` report inside autogenerated processor
    -   Keep only `typecheck` issues when needed
    -   Don't hide `typecheck` errors inside diff processor
5.  Misc.
    -   ⚠️ log an error when using previously deprecated linters ([Linter Deprecation Cycle](https://golangci-lint.run/product/roadmap/#linter-deprecation-cycle))
        -   [`deadcode`](https://togithub.com/remyoudompheng/go-misc/tree/master/deadcode): deprecated since v1.49.0 (2022-08-23).
        -   [`exhaustivestruct`](https://togithub.com/mbilski/exhaustivestruct): deprecated since v1.46.0 (2022-05-08).
        -   [`golint`](https://togithub.com/golang/lint): deprecated since v1.41.0 (2021-06-15).
        -   [`ifshort`](https://togithub.com/esimonov/ifshort): deprecated since v1.48.0 (2022-08-04).
        -   [`interfacer`](https://togithub.com/mvdan/interfacer): deprecated since v1.38.0 (2021-03-03).
        -   [`maligned`](https://togithub.com/mdempsky/maligned): deprecated since v1.38.0 (2021-03-03).
        -   [`nosnakecase`](https://togithub.com/sivchari/nosnakecase): deprecated since v1.48.0 (2022-08-04).
        -   [`scopelint`](https://togithub.com/kyoh86/scopelint): deprecated since v1.39.0 (2021-03-25).
        -   [`structcheck`](https://togithub.com/opennota/check): deprecated since v1.49.0 (2022-08-23).
        -   [`varcheck`](https://togithub.com/opennota/check): deprecated since v1.49.0 (2022-08-23).
    -   ⚠️ Deprecate usage of linter alternative names
    -   Remove help display on errors with `config verify` command
    -   Add `pre-commit` hook to run `config verify`
    -   Improve `github-action` output
6.  Documentation
    -   Remove deprecated Atom from Editor Integrations

GitHub Action (v5.1.0) for golangci-lint:

-   supports for `pull`, `pull_request_target`, and `merge_group` events with the option `only-new-issues`.
-   ️️⚠️ `skip-pkg-cache` and `skip-build-cache` have been removed because the cache related to Go itself is already handled by `actions/setup-go`.
-   with golangci-lint v1.58, the file information (path and position) will be displayed on the log.

### [`v1.57.2`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1572)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.57.1...v1.57.2)

1.  Updated linters
    -   `contextcheck`: from 1.1.4 to 1.1.5
    -   `copyloopvar`: from 1.0.8 to 1.0.10
    -   `ginkgolinter`: from 0.16.1 to 0.16.2
    -   `goconst`: from 1.7.0 to 1.7.1
    -   `gomoddirectives`: from 0.2.3 to 0.2.4
    -   `intrange`: from 0.1.0 to 0.1.1
2.  Misc.
    -   Display warnings on deprecated linter options
    -   Fix missing `colored-tab` output format
    -   Fix TeamCity `inspectionType` service message
3.  Documentation
    -   Remove invalid example about mixing files and directory
    -   Improve linters page

### [`v1.57.1`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1571)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.57.0...v1.57.1)

1.  Fixes
    -   Ignore issues with invalid position (e.g. `contextcheck`).

### [`v1.57.0`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1570)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.56.2...v1.57.0)

1.  New linters
    -   `copyloopvar`: https://github.com/karamaru-alpha/copyloopvar
    -   `intrange`: https://github.com/ckaznocha/intrange
2.  Updated linters
    -   `dupword`: from 0.0.13 to 0.0.14
    -   `gci`: from 0.12.1 to 0.12.3
    -   `ginkgolinter`: from 0.15.2 to 0.16.1 (new option `force-expect-to`, `validate-async-intervals`, and `forbid-spec-pollution`)
    -   `go-critic`: from 0.11.1 to 0.11.2
    -   `go-critic`: support of `enable-all` and `disable-all` options
    -   `go-spancheck`: from 0.5.2 to 0.5.3
    -   `gomodguard`: from 1.3.0 to 1.3.1
    -   `govet`: deprecation of `check-shadowing` ⚠️
    -   `govet`: disable temporarily `httpresponse` because of a bug [https://github.com/golang/go/issues/66259](https://togithub.com/golang/go/issues/66259)
    -   `misspell`: add `extra-words`
    -   `musttag`: from 0.8.0 to 0.9.0
    -   `nakedret`: from 2.0.2 to 2.0.4
    -   `paralleltest`: from 1.0.9 to 1.0.10
    -   `perfsprint`: from 0.6.0 to 0.7.1 (new option `strconcat`)
    -   `protogetter`: from 0.3.4 to 0.3.5
    -   `revive`: add `exclude` option
    -   `sloglint`: from 0.4.0 to 0.5.0 (new option `no-global`)
    -   `staticcheck`: from 0.4.6 to 0.4.7
    -   `testifylint`: from 1.1.2 to 1.2.0 (new option `bool-compare`)
    -   `unconvert`: to HEAD (new options `fast-math` and `safe`)
    -   `wrapcheck`: from 2.8.1 to 2.8.3
    -   Disable `copyloopvar` and `intrange` on Go < 1.22
3.  Enhancements
    -   🧩New custom linters system https://golangci-lint.run/plugins/module-plugins/
    -   Allow running only a specific linter without modifying the file configuration (`--enable-only`)
    -   Allow custom sort order for the reports (`output.sort-order`)
    -   Automatically adjust the maximum concurrency to the container CPU quota if `run.concurrency=0`
    -   Add `config verify` command to check the configuration against the JSON Schema
    -   Option to strictly follow Go generated file convention (`issues.exclude-generated-strict`)
    -   Syntax to not override `severity` from linters (`@linter`)
    -   Use severities from `gosec`
    -   Create automatically directory related to `output.formats.path`
    -   Use the first issue without inline on `mergeLineIssues` on multiple issues
4.  Misc.
    -   ⚠️ Inactivate deprecated linters (`deadcode`, `exhaustivestruct`, `golint`, `ifshort`, `interfacer`, `maligned`, `nosnakecase`, `scopelint`, `structcheck`, `varcheck`)
    -   ⚠️ Deprecated CLI flags have been removed (deprecated since 2018)
    -   ⚠️ Move `show-stats` option from `run` to `output` configuration section
    -   ⚠️ Replace `run.skip-xxx` options by `issues.exclude-xxx` options
    -   ⚠️ Replace `output.format` by `output.formats` with a new file configuration syntax
    -   Internal rewrite of the CLI
    -   Improve 'no go files to analyze' message
    -   Use `GOTOOLCHAIN=auto` inside the Docker images
5.  Documentation
    -   ⚠️ Define the linter deprecation cycle https://golangci-lint.run/product/roadmap/#linter-deprecation-cycle
    -   🎉Use information from the previous release to create linter pages
    -   Publish JSON schema on https://golangci-lint.run/jsonschema/golangci.jsonschema.json
    -   Reorganize documentation pages
    -   Add an explanation about the configuration file inside golangci-lint repository

**⚠️ Important ⚠️**

1.  Deprecated linters are inactivated, you still need to disable them if you are using `enable-all`.
2.  Deprecated CLI flags (about linter settings and `deadline`) have been removed.

### [`v1.56.2`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1562)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.56.1...v1.56.2)

1.  updated linters
    -   `go-critic`: from 0.11.0 to 0.11.1
    -   `gosec`: from 2.18.2 to 2.19.0
    -   `testifylint`: from 1.1.1 to 1.1.2
    -   `usestdlibvars`: from 1.24.0 to 1.25.0
    -   `wsl`: from 4.2.0 to 4.2.1
2.  misc.
    -   Fix missing version in Docker image
3.  Documentation
    -   Explain the limitation of `new-from-rev` and `new-from-patch`

### [`v1.56.1`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1561)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.56.0...v1.56.1)

1.  updated linters
    -   `errcheck`: from 1.6.3 to 1.7.0
    -   `govet`: disable `loopclosure` with go1.22
    -   `revive`: from 1.3.6 to 1.3.7
    -   `testifylint`: from 1.1.0 to 1.1.1

### [`v1.56.0`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1560)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.55.2...v1.56.0)

1.  new linters
    -   `spancheck`: https://github.com/jjti/go-spancheck
2.  updated linters
    -   `depguard`: from 2.1.0 to 2.2.0
    -   `exhaustive`: from 0.11.0 to 0.12.0
    -   `exhaustruct`: from 3.1.0 to 3.2.0
    -   `gci`: from 0.11.2 to 0.12.1
    -   `ginkgolinter`: from 0.14.1 to 0.15.2
    -   `go-check-sumtype`: from 0.1.3 to 0.1.4
    -   `go-critic`: from 0.9.0 to 0.11.0
    -   `go-errorlint`: from 1.4.5 to 1.4.8
    -   `go-spancheck`: from 0.4.2 to 0.5.2
    -   `goconst`: from 1.6.0 to 1.7.0
    -   `godot`: from 1.4.15 to 1.4.16
    -   `gofumpt`: from 0.5.0 to 0.6.0
    -   `inamedparam`: from 0.1.2 to 0.1.3
    -   `ineffassign`: from 0.0.0-20230610083614-0e73809eb601 to 0.1.0
    -   `ireturn`: from 0.2.2 to 0.3.0
    -   `misspell`: add mode option
    -   `musttag`: from v0.7.2 to v0.8.0
    -   `paralleltest`: from 1.0.8 to 1.0.9
    -   `perfsprint`: from 0.2.0 to 0.6.0
    -   `protogetter`: from 0.2.3 to 0.3.4
    -   `revive`: from 1.3.4 to 1.3.6
    -   `sloglint`: add static-msg option
    -   `sloglint`: from 0.1.2 to 0.4.0
    -   `testifylint`: from 0.2.3 to 1.1.0
    -   `unparam`: from [`2022122`](https://togithub.com/golangci/golangci-lint/commit/20221223090309)-7455f1af531d to [`2024010`](https://togithub.com/golangci/golangci-lint/commit/20240104100049)-c549a3470d14
    -   `whitespace`: update after moving to the `analysis` package
    -   `wsl`: from 3.4.0 to 4.2.0
    -   `zerologlint`: from 0.1.3 to 0.1.5
3.  misc.
    -   🎉 go1.22 support
    -   Implement stats per linter with a flag
    -   Make versioning inside Docker image consistent with binaries
    -   Parse Go RC version
4.  Documentation
    -   Fix `noctx` description
    -   Add missing fields to `.golangci.reference.yml`
    -   Improve `.golangci.reference.yml` defaults
    -   `typecheck`: improve FAQ
    -   `exhaustruct`: note that struct regular expressions are expected to match the entire `package/name/structname`
    -   `wrapcheck`: adjust `ignoreSigs` to new defaults

**Important**

`testifylint` has [breaking changes](https://togithub.com/Antonboom/testifylint/releases/tag/v1.0.0) about enabling/disabling checks:

-   If you were using the option `enable` with a filtered list of checks, you should either add `disable-all: true` (1) or use `disable` field (2).

    ```yml
    ```

### Example (1)

      testifylint:
        disable-all: true
        enable:
          - bool-compare
          - compares
          - empty
          - error-is-as
          - error-nil
          - expected-actual
          - go-require
          - float-compare
          - len
          - nil-compare
          - require-error

### - suite-dont-use-pkg

          - suite-extra-assert-call
          - suite-thelper
    ```

    ```yml

### Example (2)

      testifylint:
        disable:
          - suite-dont-use-pkg
    ```

### [`v1.55.2`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1552)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.55.1...v1.55.2)

1.  updated linters
    -   `ireturn`: from 0.2.1 to 0.2.2
    -   `ginkgolinter`: from 0.14.0 to 0.14.1

### [`v1.55.1`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1551)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.55.0...v1.55.1)

1.  updated linters
    -   `gosec`: from 2.18.1 to 2.18.2
2.  misc.
    -   `revgrep`: from v0.5.0 to v0.5.2 (support git < 2.41.0)
    -   output: convert backslashes to forward slashes for GitHub Action annotations printer

### [`v1.55.0`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1550)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.54.2...v1.55.0)

1.  new linters
    -   `gochecksumtype`: https://github.com/alecthomas/go-check-sumtype
    -   `inamedparam`: https://github.com/macabu/inamedparam
    -   `perfsprint`: https://github.com/catenacyber/perfsprint
    -   `protogetter`: https://github.com/ghostiam/protogetter
    -   `sloglint`: https://github.com/go-simpler/sloglint
    -   `testifylint`: https://github.com/Antonboom/testifylint
2.  updated linters
    -   `bidichk`: from 0.2.4 to 0.2.7
    -   `decorder`: from 0.4.0 to 0.4.1
    -   `dupword`: from 0.0.12 to 0.0.13
    -   `errchkjson`: from 0.3.1 to 0.3.6
    -   `gci`: from 0.11.0 to 0.11.2
    -   `ginkgolinter`: from 0.13.5 to 0.14.0
    -   `go-errorlint`: from 1.4.4 to 1.4.5
    -   `gocognit`: from 1.0.7 to 1.1.0
    -   `goconst`: from 1.5.1 to 1.6.0
    -   `godot`: from 1.4.14 to 1.4.15
    -   `gofmt`: update to HEAD
    -   `goimports`: update to HEAD
    -   `gosec`: from 2.17.0 to 2.18.1
    -   `gosmopolitan`: from 1.2.1 to 1.2.2
    -   `govet`: add `appends` analyzer
    -   `ireturn`: from 0.2.0 to 0.2.1
    -   `protogetter`: from 0.2.2 to 0.2.3
    -   `revgrep`: from [`745bb2f`](https://togithub.com/golangci/golangci-lint/commit/745bb2f7c2e6) to v0.5.0
    -   `revive`: from 1.3.2 to 1.3.4
    -   `sqlclosecheck`: from 0.4.0 to 0.5.1
    -   `staticcheck`: from 0.4.5 to 0.4.6
    -   `tagalign`: from 1.3.2 to 1.3.3
    -   `unused`: support passing in options
3.  misc.
    -   Add a pre-commit hook to check all files
4.  Documentation
    -   add source options to exclude-rules docs
    -   `gosec`: add G602 to includes/excludes inside .golangci.reference.yml

### [`v1.54.2`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1542)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.54.1...v1.54.2)

1.  updated linters:
    -   `errname`: from 0.1.10 to 0.1.12
    -   `ginkgolinter`: from 0.13.3 to 0.13.5
    -   `go-errorlint`: from 1.4.3 to 1.4.4
    -   `godot`: from 1.4.11 to 1.4.14
    -   `gosec`: from 2.16.0 to 2.17.0
    -   `musttag`: from 0.7.1 to 0.7.2
    -   `nilnil`: from 0.1.5 to 0.1.7
    -   `staticcheck`: from 0.4.3 to 0.4.5
    -   `usestdlibvars`: from 1.23.0 to 1.24.0
    -   `govet`: add missing `directive` and `slog` passes

### [`v1.54.1`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1541)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.54.0...v1.54.1)

1.  updated linters:
    -   `go-critic`:  from 0.8.2 to 0.9.0
2.  misc.
    -   plugin: temporarily hide warning about using plugins using the old API

### [`v1.54.0`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1540)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.53.3...v1.54.0)

1.  updated linters:
    -   `decorder`: from 0.2.3 to 0.4.0
    -   `dupword`: from 0.0.11 to 0.0.12
    -   `errorlint`: from 1.4.2 to 1.4.3
    -   `exhaustruct`: from 2.3.0 to 3.1.0
    -   `forbidigo`: from 1.5.3 to 1.6.0
    -   `funlen`: from 0.0.3 to 0.1.0
    -   `gci`: from 0.10.1 to 0.11.0
    -   `ginkgolinter`: from 0.12.1 to 0.13.3
    -   `go-critic`: from 0.8.1 to 0.8.2
    -   `go-errorlint`: from 1.4.2 to 1.4.3
    -   `go-exhaustruct`: from 2.3.0 to 3.1.0
    -   `gocognit`: from 1.0.6 to 1.0.7
    -   `gocritic`: from 0.8.1 to 0.8.2
    -   `gofmt`: autofix missing newline at EOF
    -   `misspell`: 0.4.0 to 0.4.1
    -   `musttag`: from 0.7.0 to 0.7.1
    -   `paralleltest`: from 1.0.7 to 1.0.8
    -   `tagalign`: from 1.2.2 to 1.3.2
    -   `typecheck`: explain it and remove it from the linter list
    -   `zerologlint`: from 0.1.2 to 0.1.3
2.  misc.
    -   🎉 go1.21 support
    -   plugin: include custom linters in `enable-all`
    -   plugin: allow to use settings for plugins
3.  Documentation
    -   Add linter descriptions.

**Important**

`ruleguard` (a "rule" inside `gocritic`) was disabled in this release (v1.54.0) and was enabled again in the next release (v1.54.1).

`exhaustruct` has breaking changes with regular expressions, more details [here](https://togithub.com/GaijinEntertainment/go-exhaustruct/releases/tag/v3.0.0).

### [`v1.53.3`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1533)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.53.2...v1.53.3)

1.  updated linters:
    -   `cyclop`: from 1.2.0 to 1.2.1
    -   `exhaustive`: from 0.10.0 to 0.11.0
    -   `forbidigo`: from 1.5.1 to 1.5.3
    -   `ginkgolinter`: from 0.12.2 to 0.12.1
    -   `ineffassign`: bump to HEAD
    -   `nakedret`: from 2.0.1 to 2.0.2
    -   `zerologlint`: from 0.1.1 to 0.1.2
2.  misc.
    -   codeclimate: reduce memory allocation
    -   support illumos/amd64

### [`v1.53.2`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1532)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.53.1...v1.53.2)

1.  updated linters
    -   `depguard`: from v2.0.1 to 2.1.0
2.  misc.
    -   `depguard`: throw error only when the linter is called

### [`v1.53.1`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1531)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.53.0...v1.53.1)

1.  misc.
    -   `depguard`: fix GOROOT detection
    -   `importas`: fix duplication detection when aliases use regular expression replacement pattern

### [`v1.53.0`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1530)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.52.2...v1.53.0)

1.  new linters
    -   `gosmopolitan`: https://github.com/xen0n/gosmopolitan
    -   `mirror`: https://github.com/butuzov/mirror
    -   `tagalign`: https://github.com/4meepo/tagalign
    -   `zerologlint`: https://github.com/ykadowak/zerologlint
2.  updated linters
    -   `bodyclose`: bump to HEAD
    -   `containedctx`: from 1.0.2 to 1.0.3
    -   `depguard`: migrate to v2
    -   `errname`: from 0.1.9 to 0.1.10
    -   `exhaustive`: from 0.9.5 to 0.10.0
    -   `forbidigo`: better support for configuring complex rules
    -   `gci`: improve error message
    -   `ginkgolinter`: add suppress-async-assertion option
    -   `ginkgolinter`: from 0.9.0 to 0.12.0
    -   `go-critic`: from 0.7.0 to 0.8.1
    -   `go-errorlint`: from 1.4.0 to 1.4.2
    -   `gofumpt`: from 0.4.0 to 0.5.0
    -   `gosec`: convert global settings as map with proper key type
    -   `gosec`: from 2.15.0 to 2.16.0
    -   `importas`: detect duplicate alias or package in the configuration
    -   `importas`: fix typo in logs
    -   `ireturn`: from 0.1.1 to 0.2.0
    -   `musttag`: from 0.5.0 to 0.7.0
    -   `nakedret`: to 2.0.1
    -   `nilnil`: from 0.1.3 to 0.1.5
    -   `paralleltest`: from 1.0.6 to 1.0.7
    -   `revive`: from 1.3.1 to 1.3.2
    -   `tagliatelle`: from 0.4.0 to 0.5.0
    -   `usestdlibvars`: fix configuration
3.  misc.
    -   `golang.org/x/tools`: from 0.7.0 to 0.9.2
    -   add loongarch64 to the install script
    -   output: add colored-tab
    -   remove warning when the config file is explicitly stdin
    -   rules: support inverted path match
4.  Documentation
    -   `mnd`: clarify ignore usage examples to be regexps
    -   `tagliatelle`: update documentation
    -   improve features section
    -   update supported Go versions FAQ

### [`v1.52.2`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1522)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.52.1...v1.52.2)

1.  updated linters
    -   `tparallel`: from 0.3.0 to 0.3.1
2.  misc.
    -   fix: pre-commit `require_serial` and `pass_filenames`

### [`v1.52.1`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1521)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.52.0...v1.52.1)

1.  misc.
    -   fix: improve panic management
    -   fix: the combination of --fix and --path-prefix

### [`v1.52.0`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1520)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.51.2...v1.52.0)

1.  updated linters
    -   `asciicheck`: from 0.1.1 to 0.2.0
    -   `bidichk`: from 0.2.3 to 0.2.4
    -   `contextcheck`: from 1.1.3 to 1.1.4
    -   `dupword`: from 0.0.9 to 0.0.11
    -   `durationcheck`: from 0.0.9 to 0.0.10
    -   `errchkjson`: from 0.3.0 to 0.3.1
    -   `errname`: from 0.1.7 to 0.1.9
    -   `forbidigo`: from 1.4.0 to 1.5.1
    -   `gci`: from 0.9.1 to 0.10.1
    -   `ginkgolinter`: from 0.8.1 to 0.9.0
    -   `go-critic`: from 0.6.7 to 0.7.0
    -   `go-errorlint`: from 1.1.0  to 1.4.0
    -   `godox`: bump to HEAD
    -   `lll`: skip go command
    -   `loggercheck`: from 0.9.3 to 0.9.4
    -   `musttag`: from 0.4.5 to 0.5.0
    -   `nilnil`: from 0.1.1 to 0.1.3
    -   `noctx`: from 0.0.1 to 0.0.2
    -   `revive`: from 1.2.5 to 1.3.1
    -   `rowserrcheck`: remove limitation related to generics support
    -   `staticcheck`: from 0.4.2 to 0.4.3
    -   `testpackage`: from 1.1.0 to 1.1.1
    -   `tparallel`: from 0.2.1 to 0.3.0
    -   `wastedassign`: remove limitation related to generics support
    -   `wrapcheck`: from 2.8.0 to 2.8.1
2.  misc.
    -   Add TeamCity output format
    -   Consider path prefix when matching path patterns
    -   Add Go version to version information
3.  Documentation
    -   Add Tekton in Trusted By page
    -   Clarify that custom linters are not enabled by default
    -   Remove description for deprecated "go" option

### [`v1.51.2`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1512)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.51.1...v1.51.2)

1.  updated linters
    -   `forbidigo`: from 1.3.0 to 1.4.0
    -   `gci`: from 0.9.0 to 0.9.1
    -   `go-critic`: from 0.6.5 to 0.6.7
    -   `go-errorlint`: from 1.0.6 to 1.1.0
    -   `gosec`: from 2.14.0 to 2.15.0
    -   `musttag`: from 0.4.4 to 0.4.5
    -   `staticcheck`: from 0.4.0 to 0.4.2
    -   `tools`: from 0.5.0 to 0.6.0
    -   `usestdlibvars`: from 1.21.1 to 1.23.0
    -   `wsl`: from 3.3.0 to 3.4.0
    -   `govet`: enable `timeformat` by default
2.  misc.
    -   fix: cache status size calculation
    -   add new source archive
3.  Documentation
    -   Improve installation section
    -   Replace links to godoc.org with pkg.go.dev

### [`v1.51.1`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1511)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.51.0...v1.51.1)

1.  updated linters
    -   `ginkgolinter`: from 0.7.1 to 0.8.1
    -   `ineffassign`: bump to HEAD
    -   `musttag`: from 0.4.3 to 0.4.4
    -   `sqlclosecheck`: from 0.3.0 to 0.4.0
    -   `staticcheck`: bump to v0.4.0
    -   `wastedassign`: from 2.0.6 to 2.0.7
    -   `wrapcheck`: from 2.7.0 to 2.8.0

### [`v1.51.0`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1510)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.50.1...v1.51.0)

1.  new linters
    -   `ginkgolinter`: https://github.com/nunnatsa/ginkgolinter
    -   `musttag`: https://github.com/tmzane/musttag
    -   `gocheckcompilerdirectives`: https://github.com/leighmcculloch/gocheckcompilerdirectives
2.  updated linters
    -   `bodyclose`: to HEAD
    -   `dupword`: from 0.0.7 to 0.0.9
    -   `errcheck`: from 1.6.2 to 1.6.3
    -   `exhaustive`: from 0.8.3 to 0.9.5
    -   `exportloopref`: from 0.1.8 to 0.1.11
    -   `gci`: from 0.8.1 to 0.9.0
    -   `ginkgolinter`: from 0.6.0 to 0.7.1
    -   `go-errorlint`: from 1.0.5 to 1.0.6
    -   `go-ruleguard`: from 0.3.21 to 0.3.22
    -   `gocheckcompilerdirectives`: from 1.1.0 to 1.2.1
    -   `gochecknoglobals`: from 0.1.0 to 0.2.1
    -   `gomodguard`: from 1.2.4 to 1.3.0
    -   `gosec`: from 2.13.1 to 2.14.0
    -   `govet`: Add `timeformat` to analysers
    -   `grouper`: from 1.1.0 to 1.1.1
    -   `musttag`: from 0.4.1 to 0.4.3
    -   `revive`: from 1.2.4 to 1.2.5
    -   `tagliatelle`: from 0.3.1 to 0.4.0
    -   `tenv`: from 1.7.0 to 1.7.1
    -   `unparam`: bump to HEAD
    -   `usestdlibvars`: from 1.20.0 to 1.21.1
    -   `wsl`: fix `force-err-cuddling` flag
3.  misc.
    -   go1.20 support
    -   remove deprecated linters from presets
    -   Build NetBSD binaries
    -   Build loong64 binaries
4.  Documentation
    -   `goimport`: improve documentation for local-prefixes
    -   `gomnd`: add missing always ignored functions
    -   `nolint`: fix typo
    -   `tagliatelle` usage typo
    -   add note about binary requirement for plugin
    -   cache preserving and colored output on docker runs
    -   improve documentation about debugging.
    -   improve Editor Integration section
    -   More specific default cache directory
    -   update output example to use valid checkstyle example; add json example

### [`v1.50.1`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1501)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.50.0...v1.50.1)

1.  updated linters
    -   `contextcheck`: from 1.1.2 to 1.1.3
    -   `go-mnd`: from 2.5.0 to 2.5.1
    -   `wrapcheck`: from 2.6.2 to 2.7.0
    -   `revive`: fix configuration parsing
    -   `lll`: skip imports
2.  misc.
    -   windows: remove redundant character escape '/'
    -   code-climate: add default severity

### [`v1.50.0`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1500)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.49.0...v1.50.0)

1.  new linters
    -   `dupword`: https://github.com/Abirdcfly/dupword
    -   `testableexamples`: https://github.com/maratori/testableexamples
2.  updated linters
    -   `contextcheck`: change owner
    -   `contextcheck`: from 1.0.6 to 1.1.2
    -   `depguard`: from 1.1.0 to 1.1.1
    -   `exhaustive`: add missing config
    -   `exhaustive`: from 0.8.1 to 0.8.3
    -   `gci`: from 0.6.3 to 0.8.0
    -   `go-critic`: from 0.6.4 to 0.6.5
    -   `go-errorlint`: from 1.0.2 to 1.0.5
    -   `go-reassign`: v0.1.2 to v0.2.0
    -   `gofmt`: add option `rewrite-rules`
    -   `gofumpt` from 0.3.1 to 0.4.0
    -   `goimports`: update to HEAD
    -   `interfacebloat`: fix configuration loading
    -   `logrlint`: rename `logrlint` to `loggercheck`
    -   `paralleltest`: add tests of the ignore-missing option
    -   `revive`: from 1.2.3 to 1.2.4
    -   `usestdlibvars`: from 1.13.0 to 1.20.0
    -   `wsl`: support all configs and update docs
3.  misc.
    -   Normalize `exclude-rules` paths for Windows
    -   add riscv64 to the install script
4.  Documentation
    -   cli: remove reference to old service

### [`v1.49.0`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1490)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.48.0...v1.49.0)

IMPORTANT: `varcheck` and `deadcode` has been removed of default linters.

1.  new linters
    -   `interfacebloat`: https://github.com/sashamelentyev/interfacebloat
    -   `logrlint`: https://github.com/timonwong/logrlint
    -   `reassign`: https://github.com/curioswitch/go-reassign
2.  updated linters
    -   `go-colorable`: from 0.1.12 to 0.1.13
    -   `go-critic`: from 0.6.3 to 0.6.4
    -   `go-errorlint`: from 1.0.0 to 1.0.2
    -   `go-exhaustruct`: from 2.2.2 to 2.3.0
    -   `gopsutil`: from 3.22.6 to 3.22.7
    -   `gosec`: from 2.12.0 to 2.13.1
    -   `revive`: from 1.2.1 to 1.2.3
    -   `usestdlibvars`: from 1.8.0 to 1.13.0
    -   `contextcheck`: from v1.0.4 to v1.0.6 && re-enable
    -   `nosnakecase`: This linter is deprecated.
    -   `varcheck`: This linter is deprecated use `unused` instead.
    -   `deadcode`: This linter is deprecated use `unused` instead.
    -   `structcheck`: This linter is deprecated use `unused` instead.
3.  documentation
    -   `revive`: fix wrong URL
    -   Add a section about default exclusions
    -   `usestdlibvars`: fix typo in documentation
    -   `nolintlint`: remove allow-leading-space option
    -   Update documentation and assets
4.  misc.
    -   dev: rewrite the internal tests framework
    -   fix: exit early on run --version
    -   fix: set an explicit `GOROOT` in the Docker image for `go-critic`

### [`v1.48.0`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1480)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.47.3...v1.48.0)

1.  new linters
    -   `usestdlibvars`:https://github.com/sashamelentyev/usestdlibvars
2.  updated linters
    -   `contextcheck`: disable linter
    -   `errcheck`: from 1.6.1 to 1.6.2
    -   `gci`: add missing `custom-order` setting
    -   `gci`: from 0.5.0 to 0.6.0
    -   `ifshort`: deprecate linter
    -   `nolint`: drop allow-leading-space option and add "nolint:all"
    -   `revgrep`: bump to HEAD
3.  documentation
    -   remove outdated info on source install
4.  misc
    -   go1.19 support

### [`v1.47.3`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1473)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.47.2...v1.47.3)

1.  updated linters:
    -   remove some go1.18 limitations
    -   `asasalint`: from 0.0.10 to 0.0.11
    -   `decorder`: from 0.2.2 to v0.2.3
    -   `gci`: fix panic with invalid configuration option
    -   `gci`: from 0.4.3 to v0.5.0
    -   `go-exhaustruct`: from 2.2.0 to 2.2.2
    -   `gomodguard`: from 1.2.3 to 1.2.4
    -   `nosnakecase`: from 1.5.0 to 1.7.0
    -   `honnef.co/go/tools`: from 0.3.2 to v0.3.3
2.  misc
    -   cgo: fix linters ignoring CGo files

### [`v1.47.2`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1472)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.47.1...v1.47.2)

1.  updated linters:
    -   `revive`: ignore slow rules

### [`v1.47.1`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1471)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.47.0...v1.47.1)

1.  updated linters:
    -   `gci`: from 0.4.2 to 0.4.3
    -   `gci`: remove the use of stdin
    -   `gci`: fix options display
    -   `tenv`: from 1.6.0 to 1.7.0
    -   `unparam`: bump to HEAD

### [`v1.47.0`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1470)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.46.2...v1.47.0)

1.  new linters:
    -   `asasalint`: https://github.com/alingse/asasalint
    -   `nosnakecase`: https://github.com/sivchari/nosnakecase
2.  updated linters:
    -   `decorder`: from 0.2.1 to 0.2.2
    -   `errcheck`: from 1.6.0 to 1.6.1
    -   `errname`: from 0.1.6 to 0.1.7
    -   `exhaustive`: from 0.7.11 to 0.8.1
    -   `gci`: fix issues and re-enable autofix
    -   `gci`: from 0.3.4 to 0.4.2
    -   `go-exhaustruct`: from 2.1.0 to 2.2.0
    -   `go-ruleguard`: from 0.3.19 to 0.3.21
    -   `gocognit`: from 1.0.5 to 1.0.6
    -   `gocyclo`: from 0.5.1 to 0.6.0
    -   `golang.org/x/tools`: bump to HEAD
    -   `gosec`: allow `global` config
    -   `gosec`: from 2.11.0 to 2.12.0
    -   `nonamedreturns`: from 1.0.1 to 1.0.4
    -   `paralleltest`: from 1.0.3 to 1.0.6
    -   `staticcheck`: fix generics
    -   `staticcheck`: from 0.3.1 to 0.3.2
    -   `tenv`: from 1.5.0 to 1.6.0
    -   `testpackage`: from 1.0.1 to 1.1.0
    -   `thelper`: from 0.6.2 to 0.6.3
    -   `wrapcheck`: from 2.6.1 to 2.6.2
3.  documentation:
    -   add thanks page
    -   add a clear explanation about the `staticcheck` integration.
    -   `depguard`: add `ignore-file-rules`
    -   `depguard`: adjust phrasing
    -   `gocritic`: add `enable` and `disable` ruleguard settings
    -   `gomnd`: fix typo
    -   `gosec`: add configs for all existing rules
    -   `govet`: add settings for `shadow` and `unusedresult`
    -   `thelper`: add `fuzz` config and description
    -   linters: add defaults

### [`v1.46.2`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1462)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.46.1...v1.46.2)

1.  updated linters:
    -   `execinquery`: bump from v1.2.0 to v1.2.1
    -   `errorlint`: bump to v1.0.0
    -   `thelper`: allow to disable one option
2.  documentation:
    -   rename `.golangci.example.yml` to `.golangci.reference.yml`
    -   add `containedctx` linter to the list of available linters

### [`v1.46.1`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1461)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.46.0...v1.46.1)

1.  updated linters:
    -   `execinquery`: bump from v0.6.0 to v0.6.1
2.  documentation:
    -   add missing linters

### [`v1.46.0`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1460)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.45.2...v1.46.0)

1.  new linters:
    -   `execinquery`: https://github.com/lufeee/execinquery
    -   `nonamedreturns`: https://github.com/firefart/nonamedreturns
    -   `nosprintfhostport`: https://github.com/stbenjam/no-sprintf-host-port
    -   `exhaustruct`: https://github.com/GaijinEntertainment/go-exhaustruct
2.  updated linters:
    -   `bidichk`: from 0.2.2 to 0.2.3
    -   `deadcode`: bump to HEAD
    -   `errchkjson`: from 0.2.3 to 0.3.0
    -   `errname`: from 0.1.5 to 0.1.6
    -   `go-critic`: from 0.6.2 to 0.6.3
    -   `gocyclo`: from 0.4.0 to 0.5.1
    -   `gofumpt` from 0.3.0 to 0.3.1
    -   `gomoddirectives`: from 0.2.2 to 0.2.3
    -   `gosec`: from 2.10.0 to 2.11.0
    -   `honnef.co/go/tools`: from 0.2.2to 0.3.1 (go1.18 support)
    -   `nilnil`: from 0.1.0 to 0.1.1
    -   `nonamedreturns`: bump from 1.0.0 to 1.0.1
    -   `predeclared`: from 0.2.1 to 0.2.2
    -   `promlinter`: bump to v0.2.0
    -   `revive`: from 1.1.4 to 1.2.1
    -   `tenv`: from 1.4.7 to 1.5.0
    -   `thelper`: from 0.5.1 to 0.6.2
    -   `unused`: fix false-positive
    -   `varnamelen`: bump to v0.8.0
    -   `wrapcheck`: from 2.5.0 to 2.6.1
    -   `exhaustivestruct`: This linter is deprecated use `exhaustruct` instead.
3.  documentation:
    -   Update "Shell Completion" instruction on Linux
    -   Update FAQ page
4.  misc:
    -   log: enable override coloring based on `CLICOLOR` and `CLICOLOR_FORCE`

### [`v1.45.2`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1452)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.45.1...v1.45.2)

1.  misc:
    -   fix: help command

### [`v1.45.1`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1451)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.45.0...v1.45.1)

1.  updated linters:
    -   `interfacer`: inactivate with go1.18
    -   `govet`: inactivate unsupported analyzers (go1.18)
    -   `depguard`: reduce requirements
    -   `structcheck`: inactivate with go1.18
    -   `varnamelen`: bump from v0.6.0 to v0.6.1
2.  misc:
    -   Automatic Go version detection 🎉 (go1.18)
    -   docker: update base images (go1.18)

### [`v1.45.0`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1450)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.44.2...v1.45.0)

1.  updated linters:
    -   `cobra`: from 1.3.0 to 1.4.0
    -   `containedctx`: from 1.0.1 to 1.0.2
    -   `errcheck`: add an option to remove default exclusions
    -   `gci`: from 0.3.1 to 0.3.2
    -   `go-header`: from 0.4.2 to 0.4.3
    -   `gofumpt`: add module-path setting
    -   `gofumpt`: from 0.2.1 to 0.3.0
    -   `gopsutil`: from 3.22.1 to 3.22.2
    -   `gosec`: from 2.9.6 to 2.10.0
    -   `makezero`: from 1.1.0 to 1.1.1
    -   `revive`: fix default values
    -   `wrapcheck`: from 2.4.0 to 2.5.0
2.  documentation:
    -   docs: add "back to the top" button
    -   docs: add `forbidigo` example that uses comments
    -   docs: improve linters page
3.  misc:
    -   go1.18 support 🎉
    -   Add an option to manage the targeted version of Go
    -   Default to YAML when config file has no extension

### [`v1.44.2`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1442)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.44.1...v1.44.2)

1.  updated linters:
    -   `gci`: bump to HEAD
    -   `gci`: restore defaults for sections
    -   `whitespace`: from 0.0.4 to 0.0.5
2.  documentation:
    -   add link to configuration in the linters list

### [`v1.44.1`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1441)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.44.0...v1.44.1)

1.  updated linters:
    -   `bidichk`: from 0.2.1 to 0.2.2
    -   `errchkjson`: from 0.2.1 to 0.2.3
    -   `thelper`: from 0.5.0 to 0.5.1
    -   `tagliatelle`: from 0.3.0 to 0.3.1
    -   `gopsutil`: from 3.21.12 to 3.22.1
    -   `gci`: from 0.2.9 to 0.3.0
    -   `revive`: from v1.1.3 to v1.1.4
    -   `varnamelen`: from v0.5.0 to v0.6.0
2.  documentation:
    -   linters: improve configuration pages
    -   `decorder`: fix `disable-init-func-first-check: false` elaboration
3.  misc:
    -   fix debug output

### [`v1.44.0`](https://togithub.com/golangci/golangci-lint/blob/HEAD/CHANGELOG.md#v1440)

[Compare Source](https://togithub.com/golangci/golangci-lint/compare/v1.43.0...v1.44.0)

1.  new linters:
    -   `containedctx`: https://github.com/sivchari/containedctx
    -   `decorder`: https://gitlab.com/bosi/decorder
    -   `errchkjson`: https://github.com/breml/errchkjson
    -   `maintidx`: https://github.com/yagipy/maintidx
    -   `grouper`: https://github.com/leonklingele/grouper
2.  updated linters:
    -   `asciicheck`: bump to v0.1.1
    -   `bidichk`: from 0.1.1 to 0.2.1
    -   `bodyclose`: bump to HEAD
    -   `decorder`: from 0.2.0 to 0.2.1
    -   `depguard`: from 1.0.1 to 1.1.0
    -   `errchkjson`: from 0.2.0 to 0.2.1
    -   `errorlint`: bump to HEAD
    -   `exhaustive`: drop deprecated/unused settings
    -   `exhaustive`: from v0.2.3 to 0.7.11
    -   `forbidigo`: from 1.2.0 to 1.3.0
    -   `forcetypeassert`: bump to v0.1.0
    -   `gocritic`: from 0.6.1 to 0.6.2
    -   `gocritic`: support autofix
    -   `gocyclo`: from 0.3.1 to 0.4.0
    -   `godot`: add period option
    -   `gofumpt`: from 0.1.1 to 0.2.1
    -   `gomnd`: from 2.4.0 to 2.5.0
    -   `gomnd`: new configuration
    -   `gosec`: from 2.9.1 to 2.9.6
    -   `ifshort`: from 1.0.3 to 1.0.4
    -   `ineffassign`: bump to HEAD
    -   `makezero`: to v1.1.0
    -   `promlinter`: from v0.1.0 to HEAD
    -   `revive`: fix `enableAllRules`
    -   `revive`: from 1.1.2 to 1.1.3
    -   `staticcheck`: from 0.2.1 to 0.2.2
    -   `tagliatelle`: from 0.2.0 to 0.3.0
    -   `thelper`: from 0.4.0 to 0.5.0
    -   `unparam`: bump to HEAD
    -   `varnamelen`: bump to v0.5.0
    -   `wrapcheck`: update configuration to include `ignoreSignRegexps`
3.  documentation:
    -   linters: improve pages about configuration
    -   improve page about false-positive
    -   `nolintlint`: fix wrong default value in comment
    -   `revive`: add a more detailed configuration
4.  misc:
    -   outputs: Add support for multiple outputs
    -   outputs: Print error text in `<failure>` tag content for more readable JUnit output
    -   outputs: ensure that the Issues key in JSON format is a list
    -   Return error if any linter fails to run
    -   cli: Show deprecated mark in the CLI linters help

</details>

---

### Configuration

📅 **Schedule**: Branch creation - At any time (no schedule defined), Automerge - At any time (no schedule defined).

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

This PR has been generated by [Renovate Bot](https://togithub.com/renovatebot/renovate).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOC41NS4yLXJwbSIsInVwZGF0ZWRJblZlciI6IjM4LjU1LjItcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6W119-->


---

## Discussion

### Comment by @app-sre-bot on September 17, 2024 at 12:24 PM UTC

Can one of the admins verify this patch?

---

## Reviews

### Review by @jlsherrill - Approved on September 18, 2024 at 01:23 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/819*
