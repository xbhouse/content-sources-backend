---
type: pull_request
number: 46
title: "Update module github.com/onsi/ginkgo/v2 to v2.27.2"
state: merged
author: red-hat-konflux
created: 2025-11-06T08:19:36Z
updated: 2025-11-07T08:39:36Z
closed: 2025-11-07T08:39:35Z
merged: 2025-11-07T08:39:35Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-onsi-ginkgo-v2-2.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/46
---

# Pull Request #46: Update module github.com/onsi/ginkgo/v2 to v2.27.2

**Author**: @red-hat-konflux
**Created**: November 06, 2025 at 08:19 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-onsi-ginkgo-v2-2.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/onsi/ginkgo/v2](https://redirect.github.com/onsi/ginkgo) | `v2.13.2` -> `v2.27.2` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fonsi%2fginkgo%2fv2/v2.27.2?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fonsi%2fginkgo%2fv2/v2.13.2/v2.27.2?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>onsi/ginkgo (github.com/onsi/ginkgo/v2)</summary>

### [`v2.27.2`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.27.2)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.27.1...v2.27.2)

#### 2.27.2

##### Fixes

- inline automaxprocs to simplify dependencies; this will be removed when Go 1.26 comes out \[[`a69113a`](https://redirect.github.com/onsi/ginkgo/commit/a69113a)]

##### Maintenance

- Fix syntax errors and typo \[[`a99c6e0`](https://redirect.github.com/onsi/ginkgo/commit/a99c6e0)]
- Fix paragraph position error \[[`f993df5`](https://redirect.github.com/onsi/ginkgo/commit/f993df5)]

### [`v2.27.1`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.27.1)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.27.0...v2.27.1)

#### 2.27.1

##### Fixes

- Fix Ginkgo Reporter slice-bounds panic \[[`606c1cb`](https://redirect.github.com/onsi/ginkgo/commit/606c1cb)]
- Bug Fix: Add GinkoTBWrapper.Attr() and GinkoTBWrapper.Output() \[[`a6463b3`](https://redirect.github.com/onsi/ginkgo/commit/a6463b3)]

### [`v2.27.0`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.27.0)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.26.0...v2.27.0)

#### 2.27.0

##### Features

##### Transforming Nodes during Tree Construction

This release adds support for `NodeArgsTransformer`s that can be registered with `AddTreeConstructionNodeArgsTransformer`.

These are called during the tree construction phase as nodes are constructed and can modify the node strings and decorators.  This enables frameworks built on top of Ginkgo to modify Ginkgo nodes and enforce conventions.

Learn more [here](https://onsi.github.io/ginkgo/#advanced-transforming-node-arguments-during-tree-construction).

##### Spec Prioritization

A new `SpecPriority(int)` decorator has been added.  Ginkgo will honor priority when ordering specs, ensuring that higher priority specs start running before lower priority specs

Learn more [here](https://onsi.github.io/ginkgo/#prioritizing-specs).

##### Maintenance

- Bump rexml from 3.4.0 to 3.4.2 in /docs ([#&#8203;1595](https://redirect.github.com/onsi/ginkgo/issues/1595)) \[[`1333dae`](https://redirect.github.com/onsi/ginkgo/commit/1333dae)]
- Bump github.com/gkampitakis/go-snaps from 0.5.14 to 0.5.15 ([#&#8203;1600](https://redirect.github.com/onsi/ginkgo/issues/1600)) \[[`17ae63e`](https://redirect.github.com/onsi/ginkgo/commit/17ae63e)]

### [`v2.26.0`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.26.0)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.25.3...v2.26.0)

#### 2.26.0

##### Features

Ginkgo can now generate json-formatted reports that are compatible with the `go test` json format.  Use `ginkgo --gojson-report=report.go.json`.  This is not intended to be a replacement for Ginkgo's native json format which is more information rich and better models Ginkgo's test structure semantics.

### [`v2.25.3`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.25.3)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.25.2...v2.25.3)

#### 2.25.3

##### Fixes

- emit --github-output group only for progress report itself \[[`f01aed1`](https://redirect.github.com/onsi/ginkgo/commit/f01aed1)]

### [`v2.25.2`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.25.2)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.25.1...v2.25.2)

#### 2.25.2

##### Fixes

Add github output group for progress report content

##### Maintenance

Bump Gomega

### [`v2.25.1`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.25.1)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.25.0...v2.25.1)

#### 2.25.1

##### Fixes

- fix(types): ignore nameless nodes on FullText() \[[`10866d3`](https://redirect.github.com/onsi/ginkgo/commit/10866d3)]
- chore: fix some CodeQL warnings \[[`2e42cff`](https://redirect.github.com/onsi/ginkgo/commit/2e42cff)]

### [`v2.25.0`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.25.0)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.24.0...v2.25.0)

#### 2.25.0

##### `AroundNode`

This release introduces a new decorator to support more complex spec setup usecases.

`AroundNode` registers a function that runs before each individual node.  This is considered a more advanced decorator.

Please read the [docs](https://onsi.github.io/ginkgo/#advanced-around-node) for more information and some examples.

Allowed signatures:

- `AroundNode(func())` - `func` will be called before the node is run.
- `AroundNode(func(ctx context.Context) context.Context)` - `func` can wrap the passed in context and return a new one which will be passed on to the node.
- `AroundNode(func(ctx context.Context, body func(ctx context.Context)))` - `ctx` is the context for the node and `body` is a function that must be called to run the node.  This gives you complete control over what runs before and after the node.

Multiple `AroundNode` decorators can be applied to a single node and they will run in the order they are applied.

Unlike setup nodes like `BeforeEach` and `DeferCleanup`, `AroundNode` is guaranteed to run in the same goroutine as the decorated node.  This is necessary when working with lower-level libraries that must run on a single thread (you can call `runtime.LockOSThread()` in the `AroundNode` to ensure that the node runs on a single thread).

Since `AroundNode` allows you to modify the context you can also use `AroundNode` to implement shared setup that attaches values to the context.

If applied to a container, `AroundNode` will run before every node in the container.  Including setup nodes like `BeforeEach` and `DeferCleanup`.

`AroundNode` can also be applied to `RunSpecs` to run before every node in the suite.  This opens up new mechanisms for instrumenting individual nodes across an entire suite.

### [`v2.24.0`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.24.0)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.23.4...v2.24.0)

#### 2.24.0

##### Features

Specs can now be decorated with (e.g.) `SemVerConstraint("2.1.0")` and `ginkgo --sem-ver-filter="2.1.1"` will only run constrained specs that match the requested version.  Learn more in the docs [here](https://onsi.github.io/ginkgo/#spec-semantic-version-filtering)!  Thanks to [@&#8203;Icarus9913](https://redirect.github.com/Icarus9913) for the PR.

##### Fixes

- remove -o from run command \[[`3f5d379`](https://redirect.github.com/onsi/ginkgo/commit/3f5d379)].  fixes [#&#8203;1582](https://redirect.github.com/onsi/ginkgo/issues/1582)

##### Maintenance

Numerous dependency bumps and documentation fixes

### [`v2.23.4`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.23.4)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.23.3...v2.23.4)

#### 2.23.4

Prior to this release Ginkgo would compute the incorrect number of available CPUs when running with `-p` in a linux container.  Thanks to [@&#8203;emirot](https://redirect.github.com/emirot) for the fix!

##### Features

- Add automaxprocs for using CPUQuota \[[`2b9c428`](https://redirect.github.com/onsi/ginkgo/commit/2b9c428)]

##### Fixes

- clarify gotchas about -vet flag \[[`1f59d07`](https://redirect.github.com/onsi/ginkgo/commit/1f59d07)]

##### Maintenance

- bump dependencies \[[`2d134d5`](https://redirect.github.com/onsi/ginkgo/commit/2d134d5)]

### [`v2.23.3`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.23.3)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.23.2...v2.23.3)

#### 2.23.3

##### Fixes

- allow `-` as a standalone argument \[[`cfcc1a5`](https://redirect.github.com/onsi/ginkgo/commit/cfcc1a5)]
- Bug Fix: Add GinkoTBWrapper.Chdir() and GinkoTBWrapper.Context() \[[`feaf292`](https://redirect.github.com/onsi/ginkgo/commit/feaf292)]
- ignore exit code for symbol test on linux \[[`88e2282`](https://redirect.github.com/onsi/ginkgo/commit/88e2282)]

### [`v2.23.2`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.23.2)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.23.1...v2.23.2)

#### 2.23.2

🎉🎉🎉

At long last, some long-standing performance gaps between `ginkgo` and `go test` have been resolved!

Ginkgo operates by running `go test -c` to generate test binaries, and then running those binaries.  It turns out that the compilation step of `go test -c` is slower than `go test`'s compilation step because `go test` strips out debug symbols (`ldflags=-w`) whereas `go test -c` does not.

Ginkgo now passes the appropriate `ldflags` to `go test -c` when running specs to strip out symbols.  This is only done when it is safe to do so and symbols are preferred when profiling is enabled and when `ginkgo build` is called explicitly.

This, coupled, with the [instructions for disabling XProtect on MacOS](https://onsi.github.io/ginkgo/#if-you-are-running-on-macos) yields a much better performance experience with Ginkgo.

### [`v2.23.1`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.23.1)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.23.0...v2.23.1)

#### 2.23.1

#### 🚨 For users on MacOS 🚨

A long-standing Ginkgo performance issue on MacOS seems to be due to mac's antimalware XProtect.  You can follow the instructions [here](https://onsi.github.io/ginkgo/#if-you-are-running-on-macos) to disable it in your terminal.  Doing so sped up Ginkgo's own test suite from 1m8s to 47s.

##### Fixes

Ginkgo's CLI is now a bit clearer if you pass flags in incorrectly:

- make it clearer that you need to pass a filename to the various profile flags, not an absolute directory \[[`a0e52ff`](https://redirect.github.com/onsi/ginkgo/commit/a0e52ff)]
- emit an error and exit if the ginkgo invocation includes flags after positional arguments \[[`b799d8d`](https://redirect.github.com/onsi/ginkgo/commit/b799d8d)]

This might cause existing CI builds to fail.  If so then it's likely that your CI build was misconfigured and should be corrected.  Open an issue if you need help.

### [`v2.23.0`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.23.0)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.22.2...v2.23.0)

#### 2.23.0

Ginkgo 2.23.0 adds a handful of methods to `GinkgoT()` to make it compatible with the `testing.TB` interface in Go 1.24.  `GinkgoT().Context()`, in particular, is a useful shorthand for generating a new context that will clean itself up in a `DeferCleanup()`.  This has subtle behavior differences from the golang implementation but should make sense in a Ginkgo... um... context.

##### Features

- bump to go 1.24.0 - support new testing.TB methods and add a test to cover testing.TB regressions \[[`37a511b`](https://redirect.github.com/onsi/ginkgo/commit/37a511b)]

##### Fixes

- fix edge case where build -o is pointing at an explicit file, not a directory \[[`7556a86`](https://redirect.github.com/onsi/ginkgo/commit/7556a86)]
- Fix binary paths when precompiling multiple suites. \[[`4df06c6`](https://redirect.github.com/onsi/ginkgo/commit/4df06c6)]

##### Maintenance

- Fix: Correct Markdown list rendering in MIGRATING\_TO\_V2.md \[[`cbcf39a`](https://redirect.github.com/onsi/ginkgo/commit/cbcf39a)]
- docs: fix test workflow badge ([#&#8203;1512](https://redirect.github.com/onsi/ginkgo/issues/1512)) \[[`9b261ff`](https://redirect.github.com/onsi/ginkgo/commit/9b261ff)]
- Bump golang.org/x/net in /integration/\_fixtures/version\_mismatch\_fixture ([#&#8203;1516](https://redirect.github.com/onsi/ginkgo/issues/1516)) \[[`00f19c8`](https://redirect.github.com/onsi/ginkgo/commit/00f19c8)]
- Bump golang.org/x/tools from 0.28.0 to 0.30.0 ([#&#8203;1515](https://redirect.github.com/onsi/ginkgo/issues/1515)) \[[`e98a4df`](https://redirect.github.com/onsi/ginkgo/commit/e98a4df)]
- Bump activesupport from 6.0.6.1 to 6.1.7.5 in /docs ([#&#8203;1504](https://redirect.github.com/onsi/ginkgo/issues/1504)) \[[`60cc4e2`](https://redirect.github.com/onsi/ginkgo/commit/60cc4e2)]
- Bump github-pages from 231 to 232 in /docs ([#&#8203;1447](https://redirect.github.com/onsi/ginkgo/issues/1447)) \[[`fea6f2d`](https://redirect.github.com/onsi/ginkgo/commit/fea6f2d)]
- Bump rexml from 3.2.8 to 3.3.9 in /docs ([#&#8203;1497](https://redirect.github.com/onsi/ginkgo/issues/1497)) \[[`31d7813`](https://redirect.github.com/onsi/ginkgo/commit/31d7813)]
- Bump webrick from 1.8.1 to 1.9.1 in /docs ([#&#8203;1501](https://redirect.github.com/onsi/ginkgo/issues/1501)) \[[`fc3bbd6`](https://redirect.github.com/onsi/ginkgo/commit/fc3bbd6)]
- Code linting ([#&#8203;1500](https://redirect.github.com/onsi/ginkgo/issues/1500)) \[[`aee0d56`](https://redirect.github.com/onsi/ginkgo/commit/aee0d56)]
- change interface{} to any ([#&#8203;1502](https://redirect.github.com/onsi/ginkgo/issues/1502)) \[[`809a710`](https://redirect.github.com/onsi/ginkgo/commit/809a710)]

### [`v2.22.2`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.22.2)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.22.1...v2.22.2)

#### What's Changed

- Bump golang.org/x/net from 0.32.0 to 0.33.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1496](https://redirect.github.com/onsi/ginkgo/pull/1496)
- Bump golang.org/x/crypto from 0.17.0 to 0.31.0 in /ginkgo/performance/\_fixtures/performance\_fixture by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1498](https://redirect.github.com/onsi/ginkgo/pull/1498)
- Bump github.com/onsi/gomega from 1.36.1 to 1.36.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1499](https://redirect.github.com/onsi/ginkgo/pull/1499)

**Full Changelog**: <https://github.com/onsi/ginkgo/compare/v2.22.1...v2.22.2>

### [`v2.22.1`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.22.1)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.22.0...v2.22.1)

#### 2.22.1

##### Fixes

Fix CSV encoding

- Update tests \[[`aab3da6`](https://redirect.github.com/onsi/ginkgo/commit/aab3da6)]
- Properly encode CSV rows \[[`c09df39`](https://redirect.github.com/onsi/ginkgo/commit/c09df39)]
- Add test case for proper csv escaping \[[`96a80fc`](https://redirect.github.com/onsi/ginkgo/commit/96a80fc)]
- Add meta-test \[[`43dad69`](https://redirect.github.com/onsi/ginkgo/commit/43dad69)]

##### Maintenance

- ensure \*.test files are gitignored so we don't accidentally commit compiled tests again \[[`c88c634`](https://redirect.github.com/onsi/ginkgo/commit/c88c634)]
- remove golang.org/x/net/context in favour of stdlib context \[[`4df44bf`](https://redirect.github.com/onsi/ginkgo/commit/4df44bf)]

### [`v2.22.0`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.22.0)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.21.0...v2.22.0)

#### 2.22.0

##### Features

- Add label to serial nodes \[[`0fcaa08`](https://redirect.github.com/onsi/ginkgo/commit/0fcaa08)]

This allows serial tests to be filtered using the `label-filter`

##### Maintenance

Various doc fixes

### [`v2.21.0`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.21.0)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.20.2...v2.21.0)

#### 2.21.0

##### Features

- add support for GINKGO\_TIME\_FORMAT \[[`a69eb39`](https://redirect.github.com/onsi/ginkgo/commit/a69eb39)]
- add GINKGO\_NO\_COLOR to disable colors via environment variables \[[`bcab9c8`](https://redirect.github.com/onsi/ginkgo/commit/bcab9c8)]

##### Fixes

- increase threshold in timeline matcher \[[`e548367`](https://redirect.github.com/onsi/ginkgo/commit/e548367)]
- Fix the document by replacing `SpecsThatWillBeRun` with `SpecsThatWillRun`
  \[[`c2c4d3c`](https://redirect.github.com/onsi/ginkgo/commit/c2c4d3c)]

##### Maintenance

- bump various dependencies \[[`7e65a00`](https://redirect.github.com/onsi/ginkgo/commit/7e65a00)]

### [`v2.20.2`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.20.2)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.20.1...v2.20.2)

#### 2.20.2

Require Go 1.22+

##### Maintenance

- bump go to v1.22 \[[`a671816`](https://redirect.github.com/onsi/ginkgo/commit/a671816)]

### [`v2.20.1`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.20.1)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.20.0...v2.20.1)

#### 2.20.1

##### Fixes

- make BeSpecEvent duration matcher more forgiving \[[`d6f9640`](https://redirect.github.com/onsi/ginkgo/commit/d6f9640)]

### [`v2.20.0`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.20.0)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.19.1...v2.20.0)

#### 2.20.0

##### Features

- Add buildvcs flag \[[`be5ab95`](https://redirect.github.com/onsi/ginkgo/commit/be5ab95)]

##### Maintenance

- Add update-deps to makefile \[[`d303d14`](https://redirect.github.com/onsi/ginkgo/commit/d303d14)]
- bump all dependencies \[[`7a50221`](https://redirect.github.com/onsi/ginkgo/commit/7a50221)]

### [`v2.19.1`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.19.1)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.19.0...v2.19.1)

#### 2.19.1

##### Fixes

- update supported platforms for race conditions \[[`63c8c30`](https://redirect.github.com/onsi/ginkgo/commit/63c8c30)]
- \[build] Allow custom name for binaries. \[[`ff41e27`](https://redirect.github.com/onsi/ginkgo/commit/ff41e27)]

##### Maintenance

- bump gomega \[[`76f4e0c`](https://redirect.github.com/onsi/ginkgo/commit/76f4e0c)]
- Bump rexml from 3.2.6 to 3.2.8 in /docs ([#&#8203;1417](https://redirect.github.com/onsi/ginkgo/issues/1417)) \[[`b69c00d`](https://redirect.github.com/onsi/ginkgo/commit/b69c00d)]
- Bump golang.org/x/sys from 0.20.0 to 0.21.0 ([#&#8203;1425](https://redirect.github.com/onsi/ginkgo/issues/1425)) \[[`f097741`](https://redirect.github.com/onsi/ginkgo/commit/f097741)]

### [`v2.19.0`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.19.0)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.18.0...v2.19.0)

#### 2.19.0

##### Features

[Label Sets](https://onsi.github.io/ginkgo/#label-sets) allow for more expressive and flexible label filtering.

### [`v2.18.0`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.18.0)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.17.3...v2.18.0)

#### 2.18.0

##### Features

- Add --slience-skips and --force-newlines \[[`f010b65`](https://redirect.github.com/onsi/ginkgo/commit/f010b65)]
- fail when no tests were run and --fail-on-empty was set \[[`d80eebe`](https://redirect.github.com/onsi/ginkgo/commit/d80eebe)]

##### Fixes

- Fix table entry context edge case \[[`42013d6`](https://redirect.github.com/onsi/ginkgo/commit/42013d6)]

##### Maintenance

- Bump golang.org/x/tools from 0.20.0 to 0.21.0 ([#&#8203;1406](https://redirect.github.com/onsi/ginkgo/issues/1406)) \[[`fcf1fd7`](https://redirect.github.com/onsi/ginkgo/commit/fcf1fd7)]
- Bump github.com/onsi/gomega from 1.33.0 to 1.33.1 ([#&#8203;1399](https://redirect.github.com/onsi/ginkgo/issues/1399)) \[[`8bb14fd`](https://redirect.github.com/onsi/ginkgo/commit/8bb14fd)]
- Bump golang.org/x/net from 0.24.0 to 0.25.0 ([#&#8203;1407](https://redirect.github.com/onsi/ginkgo/issues/1407)) \[[`04bfad7`](https://redirect.github.com/onsi/ginkgo/commit/04bfad7)]

### [`v2.17.3`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.17.3)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.17.2...v2.17.3)

#### 2.17.3

##### Fixes

`ginkgo watch` now ignores hidden files \[[`bde6e00`](https://redirect.github.com/onsi/ginkgo/commit/bde6e00)]

### [`v2.17.2`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.17.2)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.17.1...v2.17.2)

#### 2.17.2

##### Fixes

- fix: close files \[[`32259c8`](https://redirect.github.com/onsi/ginkgo/commit/32259c8)]
- fix github output log level for skipped specs \[[`780e7a3`](https://redirect.github.com/onsi/ginkgo/commit/780e7a3)]

##### Maintenance

- Bump github.com/google/pprof \[[`d91fe4e`](https://redirect.github.com/onsi/ginkgo/commit/d91fe4e)]
- Bump github.com/go-task/slim-sprig to v3 \[[`8cb662e`](https://redirect.github.com/onsi/ginkgo/commit/8cb662e)]
- Bump golang.org/x/net in /integration/\_fixtures/version\_mismatch\_fixture ([#&#8203;1391](https://redirect.github.com/onsi/ginkgo/issues/1391)) \[[`3134422`](https://redirect.github.com/onsi/ginkgo/commit/3134422)]
- Bump github-pages from 230 to 231 in /docs ([#&#8203;1384](https://redirect.github.com/onsi/ginkgo/issues/1384)) \[[`eca81b4`](https://redirect.github.com/onsi/ginkgo/commit/eca81b4)]
- Bump golang.org/x/tools from 0.19.0 to 0.20.0 ([#&#8203;1383](https://redirect.github.com/onsi/ginkgo/issues/1383)) \[[`760def8`](https://redirect.github.com/onsi/ginkgo/commit/760def8)]
- Bump golang.org/x/net from 0.23.0 to 0.24.0 ([#&#8203;1381](https://redirect.github.com/onsi/ginkgo/issues/1381)) \[[`4ce33f4`](https://redirect.github.com/onsi/ginkgo/commit/4ce33f4)]
- Fix test for gomega version bump \[[`f2fcd97`](https://redirect.github.com/onsi/ginkgo/commit/f2fcd97)]
- Bump github.com/onsi/gomega from 1.30.0 to 1.33.0 ([#&#8203;1390](https://redirect.github.com/onsi/ginkgo/issues/1390)) \[[`fd622d2`](https://redirect.github.com/onsi/ginkgo/commit/fd622d2)]
- Bump golang.org/x/tools from 0.17.0 to 0.19.0 ([#&#8203;1368](https://redirect.github.com/onsi/ginkgo/issues/1368)) \[[`5474a26`](https://redirect.github.com/onsi/ginkgo/commit/5474a26)]
- Bump github-pages from 229 to 230 in /docs ([#&#8203;1359](https://redirect.github.com/onsi/ginkgo/issues/1359)) \[[`e6d1170`](https://redirect.github.com/onsi/ginkgo/commit/e6d1170)]
- Bump google.golang.org/protobuf from 1.28.0 to 1.33.0 ([#&#8203;1374](https://redirect.github.com/onsi/ginkgo/issues/1374)) \[[`7f447b2`](https://redirect.github.com/onsi/ginkgo/commit/7f447b2)]
- Bump golang.org/x/net from 0.20.0 to 0.23.0 ([#&#8203;1380](https://redirect.github.com/onsi/ginkgo/issues/1380)) \[[`f15239a`](https://redirect.github.com/onsi/ginkgo/commit/f15239a)]

### [`v2.17.1`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.17.1)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.17.0...v2.17.1)

#### 2.17.1

##### Fixes

- If the user sets --seed=0, make sure all parallel nodes get the same seed \[[`af0330d`](https://redirect.github.com/onsi/ginkgo/commit/af0330d)]

### [`v2.17.0`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.17.0)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.16.0...v2.17.0)

#### 2.17.0

##### Features

- add `--github-output` for nicer output in github actions \[[`e8a2056`](https://redirect.github.com/onsi/ginkgo/commit/e8a2056)]

##### Maintenance

- fix typo in core\_dsl.go \[[`977bc6f`](https://redirect.github.com/onsi/ginkgo/commit/977bc6f)]
- Fix typo in docs \[[`e297e7b`](https://redirect.github.com/onsi/ginkgo/commit/e297e7b)]

### [`v2.16.0`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.16.0)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.15.0...v2.16.0)

#### 2.16.0

##### Features

- add SpecContext to reporting nodes

##### Fixes

- merge coverages instead of combining them ([#&#8203;1329](https://redirect.github.com/onsi/ginkgo/issues/1329)) ([#&#8203;1340](https://redirect.github.com/onsi/ginkgo/issues/1340)) \[[`23f0cc5`](https://redirect.github.com/onsi/ginkgo/commit/23f0cc5)]
- core\_dsl: disable Getwd() with environment variable ([#&#8203;1357](https://redirect.github.com/onsi/ginkgo/issues/1357)) \[[`cd418b7`](https://redirect.github.com/onsi/ginkgo/commit/cd418b7)]

##### Maintenance

- docs/index.md: Typo \[[`2cebe8d`](https://redirect.github.com/onsi/ginkgo/commit/2cebe8d)]
- fix docs \[[`06de431`](https://redirect.github.com/onsi/ginkgo/commit/06de431)]
- chore: test with Go 1.22 ([#&#8203;1352](https://redirect.github.com/onsi/ginkgo/issues/1352)) \[[`898cba9`](https://redirect.github.com/onsi/ginkgo/commit/898cba9)]
- Bump golang.org/x/tools from 0.16.1 to 0.17.0 ([#&#8203;1336](https://redirect.github.com/onsi/ginkgo/issues/1336)) \[[`17ae120`](https://redirect.github.com/onsi/ginkgo/commit/17ae120)]
- Bump golang.org/x/sys from 0.15.0 to 0.16.0 ([#&#8203;1327](https://redirect.github.com/onsi/ginkgo/issues/1327)) \[[`5a179ed`](https://redirect.github.com/onsi/ginkgo/commit/5a179ed)]
- Bump github.com/go-logr/logr from 1.3.0 to 1.4.1 ([#&#8203;1321](https://redirect.github.com/onsi/ginkgo/issues/1321)) \[[`a1e6b69`](https://redirect.github.com/onsi/ginkgo/commit/a1e6b69)]
- Bump github-pages and jekyll-feed in /docs ([#&#8203;1351](https://redirect.github.com/onsi/ginkgo/issues/1351)) \[[`d52951d`](https://redirect.github.com/onsi/ginkgo/commit/d52951d)]
- Fix docs for handling failures in goroutines ([#&#8203;1339](https://redirect.github.com/onsi/ginkgo/issues/1339)) \[[`4471b2e`](https://redirect.github.com/onsi/ginkgo/commit/4471b2e)]

### [`v2.15.0`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.15.0)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.14.0...v2.15.0)

#### 2.15.0

##### Features

- JUnit reports now interpret Label(owner:X) and set owner to X. \[[`8f3bd70`](https://redirect.github.com/onsi/ginkgo/commit/8f3bd70)]
- include cancellation reason when cancelling spec context \[[`96e915c`](https://redirect.github.com/onsi/ginkgo/commit/96e915c)]

##### Fixes

- emit output of failed go tool cover invocation so users can try to debug things for themselves \[[`c245d09`](https://redirect.github.com/onsi/ginkgo/commit/c245d09)]
- fix outline when using nodot in ginkgo v2 \[[`dca77c8`](https://redirect.github.com/onsi/ginkgo/commit/dca77c8)]
- Document areas where GinkgoT() behaves differently from testing.T \[[`dbaf18f`](https://redirect.github.com/onsi/ginkgo/commit/dbaf18f)]
- bugfix(docs): use Unsetenv instead of Clearenv ([#&#8203;1337](https://redirect.github.com/onsi/ginkgo/issues/1337)) \[[`6f67a14`](https://redirect.github.com/onsi/ginkgo/commit/6f67a14)]

##### Maintenance

- Bump to go 1.20 \[[`4fcd0b3`](https://redirect.github.com/onsi/ginkgo/commit/4fcd0b3)]

### [`v2.14.0`](https://redirect.github.com/onsi/ginkgo/releases/tag/v2.14.0)

[Compare Source](https://redirect.github.com/onsi/ginkgo/compare/v2.13.2...v2.14.0)

#### 2.14.0

##### Features

You can now use `GinkgoTB()` when you need an instance of `testing.TB` to pass to a library.

Prior to this release table testing only supported generating individual `It`s for each test entry.  `DescribeTableSubtree` extends table testing support to entire testing subtrees - under the hood `DescrieTableSubtree` generates a new container for each entry and invokes your function to fill our the container.  See the [docs](https://onsi.github.io/ginkgo/#generating-subtree-tables) to learn more.

- Introduce DescribeTableSubtree \[[`65ec56d`](https://redirect.github.com/onsi/ginkgo/commit/65ec56d)]
- add GinkgoTB() to docs \[[`4a2c832`](https://redirect.github.com/onsi/ginkgo/commit/4a2c832)]
- Add GinkgoTB() function ([#&#8203;1333](https://redirect.github.com/onsi/ginkgo/issues/1333)) \[[`92b6744`](https://redirect.github.com/onsi/ginkgo/commit/92b6744)]

##### Fixes

- Fix typo in internal/suite.go ([#&#8203;1332](https://redirect.github.com/onsi/ginkgo/issues/1332)) \[[`beb9507`](https://redirect.github.com/onsi/ginkgo/commit/beb9507)]
- Fix typo in docs/index.md ([#&#8203;1319](https://redirect.github.com/onsi/ginkgo/issues/1319)) \[[`4ac3a13`](https://redirect.github.com/onsi/ginkgo/commit/4ac3a13)]
- allow wasm to compile with ginkgo present ([#&#8203;1311](https://redirect.github.com/onsi/ginkgo/issues/1311)) \[[`b2e5bc5`](https://redirect.github.com/onsi/ginkgo/commit/b2e5bc5)]

##### Maintenance

- Bump golang.org/x/tools from 0.16.0 to 0.16.1 ([#&#8203;1316](https://redirect.github.com/onsi/ginkgo/issues/1316)) \[[`465a8ec`](https://redirect.github.com/onsi/ginkgo/commit/465a8ec)]
- Bump actions/setup-go from 4 to 5 ([#&#8203;1313](https://redirect.github.com/onsi/ginkgo/issues/1313)) \[[`eab0e40`](https://redirect.github.com/onsi/ginkgo/commit/eab0e40)]
- Bump github/codeql-action from 2 to 3 ([#&#8203;1317](https://redirect.github.com/onsi/ginkgo/issues/1317)) \[[`fbf9724`](https://redirect.github.com/onsi/ginkgo/commit/fbf9724)]
- Bump golang.org/x/crypto ([#&#8203;1318](https://redirect.github.com/onsi/ginkgo/issues/1318)) \[[`3ee80ee`](https://redirect.github.com/onsi/ginkgo/commit/3ee80ee)]
- Bump golang.org/x/tools from 0.14.0 to 0.16.0 ([#&#8203;1306](https://redirect.github.com/onsi/ginkgo/issues/1306)) \[[`123e1d5`](https://redirect.github.com/onsi/ginkgo/commit/123e1d5)]
- Bump github.com/onsi/gomega from 1.29.0 to 1.30.0 ([#&#8203;1297](https://redirect.github.com/onsi/ginkgo/issues/1297)) \[[`558f6e0`](https://redirect.github.com/onsi/ginkgo/commit/558f6e0)]
- Bump golang.org/x/net from 0.17.0 to 0.19.0 ([#&#8203;1307](https://redirect.github.com/onsi/ginkgo/issues/1307)) \[[`84ff7f3`](https://redirect.github.com/onsi/ginkgo/commit/84ff7f3)]

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

To execute skipped test pipelines write comment `/ok-to-test`.

---
### Documentation

Find out how to configure dependency updates in [MintMaker documentation](https://konflux-ci.dev/docs/mintmaker/user/) or see all available configuration options in [Renovate documentation](https://docs.renovatebot.com/configuration-options/).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6W119-->


---

## Discussion

### Comment by @red-hat-konflux on November 07, 2025 at 04:19 AM UTC

### ℹ Artifact update notice

##### File name: rh_identity_transform/go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 2 additional dependencies were updated


Details:


| **Package**                        | **Change**                                                                   |
| :--------------------------------- | :--------------------------------------------------------------------------- |
| `github.com/Masterminds/semver/v3` | `v3.3.0` -> `v3.4.0`                                                         |
| `github.com/google/pprof`          | `v0.0.0-20231212022811-ec68065c825e` -> `v0.0.0-20250403155104-27863c87afa6` |

---

## Reviews

### Review by @Hyperkid123 - Approved on November 07, 2025 at 08:39 AM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/46*
