---
type: pull_request
number: 1648
title: "Update module github.com/onsi/gomega to v1.37.0"
state: merged
author: red-hat-konflux
created: 2025-04-27T14:27:21Z
updated: 2025-04-28T09:51:35Z
closed: 2025-04-28T09:51:35Z
merged: 2025-04-28T09:51:35Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-onsi-gomega-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1648
---

# Pull Request #1648: Update module github.com/onsi/gomega to v1.37.0

**Author**: @red-hat-konflux
**Created**: April 27, 2025 at 02:27 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-onsi-gomega-1.x`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [github.com/onsi/gomega](https://redirect.github.com/onsi/gomega) | indirect | minor | `v1.16.0` -> `v1.37.0` |

---

### Release Notes

<details>
<summary>onsi/gomega (github.com/onsi/gomega)</summary>

### [`v1.37.0`](https://redirect.github.com/onsi/gomega/releases/tag/v1.37.0)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.36.3...v1.37.0)

#### 1.37.0

##### Features

-   add To/ToNot/NotTo aliases for AsyncAssertion \[[`5666f98`](https://redirect.github.com/onsi/gomega/commit/5666f98)]

### [`v1.36.3`](https://redirect.github.com/onsi/gomega/releases/tag/v1.36.3)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.36.2...v1.36.3)

#### 1.36.3

##### Maintenance

-   bump all the things \[[`adb8b49`](https://redirect.github.com/onsi/gomega/commit/adb8b49)]
-   chore: replace `interface{}` with `any` \[[`7613216`](https://redirect.github.com/onsi/gomega/commit/7613216)]
-   Bump google.golang.org/protobuf from 1.36.1 to 1.36.5 ([#&#8203;822](https://redirect.github.com/onsi/gomega/issues/822)) \[[`9fe5259`](https://redirect.github.com/onsi/gomega/commit/9fe5259)]
-   remove spurious "toolchain" from go.mod ([#&#8203;819](https://redirect.github.com/onsi/gomega/issues/819)) \[[`a0e85b9`](https://redirect.github.com/onsi/gomega/commit/a0e85b9)]
-   Bump golang.org/x/net from 0.33.0 to 0.35.0 ([#&#8203;823](https://redirect.github.com/onsi/gomega/issues/823)) \[[`604a8b1`](https://redirect.github.com/onsi/gomega/commit/604a8b1)]
-   Bump activesupport from 6.0.6.1 to 6.1.7.5 in /docs ([#&#8203;772](https://redirect.github.com/onsi/gomega/issues/772)) \[[`36fbc84`](https://redirect.github.com/onsi/gomega/commit/36fbc84)]
-   Bump github-pages from 231 to 232 in /docs ([#&#8203;778](https://redirect.github.com/onsi/gomega/issues/778)) \[[`ced70d7`](https://redirect.github.com/onsi/gomega/commit/ced70d7)]
-   Bump rexml from 3.2.6 to 3.3.9 in /docs ([#&#8203;788](https://redirect.github.com/onsi/gomega/issues/788)) \[[`c8b4a07`](https://redirect.github.com/onsi/gomega/commit/c8b4a07)]
-   Bump github.com/onsi/ginkgo/v2 from 2.22.1 to 2.22.2 ([#&#8203;812](https://redirect.github.com/onsi/gomega/issues/812)) \[[`06431b9`](https://redirect.github.com/onsi/gomega/commit/06431b9)]
-   Bump webrick from 1.8.1 to 1.9.1 in /docs ([#&#8203;800](https://redirect.github.com/onsi/gomega/issues/800)) \[[`b55a92d`](https://redirect.github.com/onsi/gomega/commit/b55a92d)]
-   Fix typos ([#&#8203;813](https://redirect.github.com/onsi/gomega/issues/813)) \[[`a1d518b`](https://redirect.github.com/onsi/gomega/commit/a1d518b)]

### [`v1.36.2`](https://redirect.github.com/onsi/gomega/releases/tag/v1.36.2)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.36.1...v1.36.2)

#### Maintenance

-   Bump nokogiri from 1.16.3 to 1.16.5 in /docs by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/onsi/gomega/pull/757](https://redirect.github.com/onsi/gomega/pull/757)
-   Bump github.com/onsi/ginkgo/v2 from 2.20.1 to 2.22.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/onsi/gomega/pull/808](https://redirect.github.com/onsi/gomega/pull/808)
-   Bump golang.org/x/net from 0.30.0 to 0.33.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/onsi/gomega/pull/807](https://redirect.github.com/onsi/gomega/pull/807)
-   Bump google.golang.org/protobuf from 1.35.1 to 1.36.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/onsi/gomega/pull/810](https://redirect.github.com/onsi/gomega/pull/810)

### [`v1.36.1`](https://redirect.github.com/onsi/gomega/releases/tag/v1.36.1)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.36.0...v1.36.1)

#### 1.36.1

##### Fixes

-   Fix [https://github.com/onsi/gomega/issues/803](https://redirect.github.com/onsi/gomega/issues/803) \[[`1c6c112`](https://redirect.github.com/onsi/gomega/commit/1c6c112)]
-   resolves [#&#8203;696](https://redirect.github.com/onsi/gomega/issues/696): make HaveField great on pointer receivers given only a non-addressable value \[[`4feb9d7`](https://redirect.github.com/onsi/gomega/commit/4feb9d7)]

### [`v1.36.0`](https://redirect.github.com/onsi/gomega/releases/tag/v1.36.0)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.35.1...v1.36.0)

#### 1.36.0

##### Features

-   new: make collection-related matchers Go 1.23 iterator aware \[[`4c964c6`](https://redirect.github.com/onsi/gomega/commit/4c964c6)]

##### Maintenance

-   Replace min/max helpers with built-in min/max \[[`ece6872`](https://redirect.github.com/onsi/gomega/commit/ece6872)]
-   Fix some typos in docs \[[`8e924d7`](https://redirect.github.com/onsi/gomega/commit/8e924d7)]

### [`v1.35.1`](https://redirect.github.com/onsi/gomega/releases/tag/v1.35.1)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.35.0...v1.35.1)

#### 1.35.1

##### Fixes

-   Export EnforceDefaultTimeoutsWhenUsingContexts and DisableDefaultTimeoutsWhenUsingContext \[[`ca36da1`](https://redirect.github.com/onsi/gomega/commit/ca36da1)]

### [`v1.35.0`](https://redirect.github.com/onsi/gomega/releases/tag/v1.35.0)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.34.2...v1.35.0)

#### 1.35.0

##### Features

-   You can now call `EnforceDefaultTimeoutsWhenUsingContexts()` to have `Eventually` honor the default timeout when passed a context.  (prior to this you had to expclility add a timeout) \[[`e4c4265`](https://redirect.github.com/onsi/gomega/commit/e4c4265)]
-   You can call `StopTrying(message).Successfully()` to abort a `Consistently` early without failure \[[`eeca931`](https://redirect.github.com/onsi/gomega/commit/eeca931)]

##### Fixes

-   Stop memoizing the result of `HaveField` to avoid unexpected errors when used with async assertions. \[[`3bdbc4e`](https://redirect.github.com/onsi/gomega/commit/3bdbc4e)]

##### Maintenance

-   Bump all dependencies \[[`a05a416`](https://redirect.github.com/onsi/gomega/commit/a05a416)]

### [`v1.34.2`](https://redirect.github.com/onsi/gomega/releases/tag/v1.34.2)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.34.1...v1.34.2)

#### 1.34.2

Require Go 1.22+

##### Maintenance

-   bump ginkgo as well \[[`c59c6dc`](https://redirect.github.com/onsi/gomega/commit/c59c6dc)]
-   bump to go 1.22 - remove x/exp dependency \[[`8158b99`](https://redirect.github.com/onsi/gomega/commit/8158b99)]

### [`v1.34.1`](https://redirect.github.com/onsi/gomega/releases/tag/v1.34.1)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.34.0...v1.34.1)

#### 1.34.1

##### Maintenance

-   Use slices from exp/slices to keep golang 1.20 compat \[[`5e71dcd`](https://redirect.github.com/onsi/gomega/commit/5e71dcd)]

### [`v1.34.0`](https://redirect.github.com/onsi/gomega/releases/tag/v1.34.0)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.33.1...v1.34.0)

#### 1.34.0

##### Features

-   Add RoundTripper method to ghttp.Server \[[`c549e0d`](https://redirect.github.com/onsi/gomega/commit/c549e0d)]

##### Fixes

-   fix incorrect handling of nil slices in HaveExactElements (fixes [#&#8203;771](https://redirect.github.com/onsi/gomega/issues/771)) \[[`878940c`](https://redirect.github.com/onsi/gomega/commit/878940c)]
-   issue\_765 - fixed bug in Hopcroft-Karp algorithm \[[`ebadb67`](https://redirect.github.com/onsi/gomega/commit/ebadb67)]

##### Maintenance

-   bump ginkgo \[[`8af2ece`](https://redirect.github.com/onsi/gomega/commit/8af2ece)]
-   Fix typo in docs \[[`123a071`](https://redirect.github.com/onsi/gomega/commit/123a071)]
-   Bump github.com/onsi/ginkgo/v2 from 2.17.2 to 2.17.3 ([#&#8203;756](https://redirect.github.com/onsi/gomega/issues/756)) \[[`0e69083`](https://redirect.github.com/onsi/gomega/commit/0e69083)]
-   Bump google.golang.org/protobuf from 1.33.0 to 1.34.1 ([#&#8203;755](https://redirect.github.com/onsi/gomega/issues/755)) \[[`2675796`](https://redirect.github.com/onsi/gomega/commit/2675796)]
-   Bump golang.org/x/net from 0.24.0 to 0.25.0 ([#&#8203;754](https://redirect.github.com/onsi/gomega/issues/754)) \[[`4160c0f`](https://redirect.github.com/onsi/gomega/commit/4160c0f)]
-   Bump github-pages from 230 to 231 in /docs ([#&#8203;748](https://redirect.github.com/onsi/gomega/issues/748)) \[[`892c303`](https://redirect.github.com/onsi/gomega/commit/892c303)]

### [`v1.33.1`](https://redirect.github.com/onsi/gomega/releases/tag/v1.33.1)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.33.0...v1.33.1)

#### 1.33.1

##### Fixes

-   fix confusing eventually docs \[[`3a66379`](https://redirect.github.com/onsi/gomega/commit/3a66379)]

##### Maintenance

-   Bump github.com/onsi/ginkgo/v2 from 2.17.1 to 2.17.2 \[[`e9bc35a`](https://redirect.github.com/onsi/gomega/commit/e9bc35a)]

### [`v1.33.0`](https://redirect.github.com/onsi/gomega/releases/tag/v1.33.0)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.32.0...v1.33.0)

#### 1.33.0

##### Features

`Receive` not accepts `Receive(<POINTER>, MATCHER>)`, allowing you to pick out a specific value on the channel that satisfies the provided matcher and is stored in the provided pointer.

##### Maintenance

-   Bump github.com/onsi/ginkgo/v2 from 2.15.0 to 2.17.1 ([#&#8203;745](https://redirect.github.com/onsi/gomega/issues/745)) \[[`9999deb`](https://redirect.github.com/onsi/gomega/commit/9999deb)]
-   Bump github-pages from 229 to 230 in /docs ([#&#8203;735](https://redirect.github.com/onsi/gomega/issues/735)) \[[`cb5ff21`](https://redirect.github.com/onsi/gomega/commit/cb5ff21)]
-   Bump golang.org/x/net from 0.20.0 to 0.23.0 ([#&#8203;746](https://redirect.github.com/onsi/gomega/issues/746)) \[[`bac6596`](https://redirect.github.com/onsi/gomega/commit/bac6596)]

### [`v1.32.0`](https://redirect.github.com/onsi/gomega/releases/tag/v1.32.0)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.31.1...v1.32.0)

#### 1.32.0

##### Maintenance

-   Migrate github.com/golang/protobuf to google.golang.org/protobuf \[[`436a197`](https://redirect.github.com/onsi/gomega/commit/436a197)]

    This release drops the deprecated github.com/golang/protobuf and adopts google.golang.org/protobuf.  Care was taken to ensure the release is backwards compatible (thanks [@&#8203;jbduncan](https://redirect.github.com/jbduncan) !).  Please open an issue if you run into one.

-   chore: test with Go 1.22 ([#&#8203;733](https://redirect.github.com/onsi/gomega/issues/733)) \[[`32ef35e`](https://redirect.github.com/onsi/gomega/commit/32ef35e)]

-   Bump golang.org/x/net from 0.19.0 to 0.20.0 ([#&#8203;717](https://redirect.github.com/onsi/gomega/issues/717)) \[[`a0d0387`](https://redirect.github.com/onsi/gomega/commit/a0d0387)]

-   Bump github-pages and jekyll-feed in /docs ([#&#8203;732](https://redirect.github.com/onsi/gomega/issues/732)) \[[`b71e477`](https://redirect.github.com/onsi/gomega/commit/b71e477)]

-   docs: fix typo and broken anchor link to gstruct \[[`f460154`](https://redirect.github.com/onsi/gomega/commit/f460154)]

-   docs: fix HaveEach matcher signature \[[`a2862e4`](https://redirect.github.com/onsi/gomega/commit/a2862e4)]

### [`v1.31.1`](https://redirect.github.com/onsi/gomega/releases/tag/v1.31.1)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.31.0...v1.31.1)

#### 1.31.1

##### Fixes

-   Inverted arguments order of FailureMessage of BeComparableToMatcher \[[`e0dd999`](https://redirect.github.com/onsi/gomega/commit/e0dd999)]
-   Update test in case keeping msg is desired \[[`ad1a367`](https://redirect.github.com/onsi/gomega/commit/ad1a367)]

##### Maintenance

-   Show how to import the format sub package \[[`24e958d`](https://redirect.github.com/onsi/gomega/commit/24e958d)]
-   tidy up go.sum \[[`26661b8`](https://redirect.github.com/onsi/gomega/commit/26661b8)]
-   bump dependencies \[[`bde8f7a`](https://redirect.github.com/onsi/gomega/commit/bde8f7a)]

### [`v1.31.0`](https://redirect.github.com/onsi/gomega/releases/tag/v1.31.0)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.30.0...v1.31.0)

#### 1.31.0

##### Features

-   Async assertions include context cancellation cause if present \[[`121c37f`](https://redirect.github.com/onsi/gomega/commit/121c37f)]

##### Maintenance

-   Bump minimum go version \[[`dee1e3c`](https://redirect.github.com/onsi/gomega/commit/dee1e3c)]
-   docs: fix typo in example usage "occured" -> "occurred" \[[`49005fe`](https://redirect.github.com/onsi/gomega/commit/49005fe)]
-   Bump actions/setup-go from 4 to 5 ([#&#8203;714](https://redirect.github.com/onsi/gomega/issues/714)) \[[`f1c8757`](https://redirect.github.com/onsi/gomega/commit/f1c8757)]
-   Bump github/codeql-action from 2 to 3 ([#&#8203;715](https://redirect.github.com/onsi/gomega/issues/715)) \[[`9836e76`](https://redirect.github.com/onsi/gomega/commit/9836e76)]
-   Bump github.com/onsi/ginkgo/v2 from 2.13.0 to 2.13.2 ([#&#8203;713](https://redirect.github.com/onsi/gomega/issues/713)) \[[`54726f0`](https://redirect.github.com/onsi/gomega/commit/54726f0)]
-   Bump golang.org/x/net from 0.17.0 to 0.19.0 ([#&#8203;711](https://redirect.github.com/onsi/gomega/issues/711)) \[[`df97ecc`](https://redirect.github.com/onsi/gomega/commit/df97ecc)]
-   docs: fix `HaveExactElement` typo ([#&#8203;712](https://redirect.github.com/onsi/gomega/issues/712)) \[[`a672c86`](https://redirect.github.com/onsi/gomega/commit/a672c86)]

### [`v1.30.0`](https://redirect.github.com/onsi/gomega/releases/tag/v1.30.0)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.29.0...v1.30.0)

#### 1.30.0

##### Features

-   BeTrueBecause and BeFalseBecause allow for better failure messages \[[`4da4c7f`](https://redirect.github.com/onsi/gomega/commit/4da4c7f)]

##### Maintenance

-   Bump actions/checkout from 3 to 4 ([#&#8203;694](https://redirect.github.com/onsi/gomega/issues/694)) \[[`6ca6e97`](https://redirect.github.com/onsi/gomega/commit/6ca6e97)]
-   doc: fix type on gleak go doc \[[`f1b8343`](https://redirect.github.com/onsi/gomega/commit/f1b8343)]

### [`v1.29.0`](https://redirect.github.com/onsi/gomega/releases/tag/v1.29.0)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.28.1...v1.29.0)

#### 1.29.0

##### Features

-   MatchError can now take an optional func(error) bool + description \[[`2b39142`](https://redirect.github.com/onsi/gomega/commit/2b39142)]

### [`v1.28.1`](https://redirect.github.com/onsi/gomega/releases/tag/v1.28.1)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.28.0...v1.28.1)

#### 1.28.1

##### Maintenance

-   Bump github.com/onsi/ginkgo/v2 from 2.12.0 to 2.13.0 \[[`635d196`](https://redirect.github.com/onsi/gomega/commit/635d196)]
-   Bump github.com/google/go-cmp from 0.5.9 to 0.6.0 \[[`14f8859`](https://redirect.github.com/onsi/gomega/commit/14f8859)]
-   Bump golang.org/x/net from 0.14.0 to 0.17.0 \[[`d8a6508`](https://redirect.github.com/onsi/gomega/commit/d8a6508)]
-   [#&#8203;703](https://redirect.github.com/onsi/gomega/issues/703) doc(matchers): HaveEach() doc comment updated \[[`2705bdb`](https://redirect.github.com/onsi/gomega/commit/2705bdb)]
-   Minor typos ([#&#8203;699](https://redirect.github.com/onsi/gomega/issues/699)) \[[`375648c`](https://redirect.github.com/onsi/gomega/commit/375648c)]

### [`v1.28.0`](https://redirect.github.com/onsi/gomega/releases/tag/v1.28.0)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.27.10...v1.28.0)

#### 1.28.0

##### Features

-   Add VerifyHost handler to ghttp ([#&#8203;698](https://redirect.github.com/onsi/gomega/issues/698)) \[[`0b03b36`](https://redirect.github.com/onsi/gomega/commit/0b03b36)]

##### Fixes

-   Read Body for Newer Responses in HaveHTTPBodyMatcher ([#&#8203;686](https://redirect.github.com/onsi/gomega/issues/686)) \[[`18d6673`](https://redirect.github.com/onsi/gomega/commit/18d6673)]

##### Maintenance

-   Bump github.com/onsi/ginkgo/v2 from 2.11.0 to 2.12.0 ([#&#8203;693](https://redirect.github.com/onsi/gomega/issues/693)) \[[`55a33f3`](https://redirect.github.com/onsi/gomega/commit/55a33f3)]
-   Typo in matchers.go ([#&#8203;691](https://redirect.github.com/onsi/gomega/issues/691)) \[[`de68e8f`](https://redirect.github.com/onsi/gomega/commit/de68e8f)]
-   Bump commonmarker from 0.23.9 to 0.23.10 in /docs ([#&#8203;690](https://redirect.github.com/onsi/gomega/issues/690)) \[[`ab17f5e`](https://redirect.github.com/onsi/gomega/commit/ab17f5e)]
-   chore: update test matrix for Go 1.21 ([#&#8203;689](https://redirect.github.com/onsi/gomega/issues/689)) \[[`5069017`](https://redirect.github.com/onsi/gomega/commit/5069017)]
-   Bump golang.org/x/net from 0.12.0 to 0.14.0 ([#&#8203;688](https://redirect.github.com/onsi/gomega/issues/688)) \[[`babe25f`](https://redirect.github.com/onsi/gomega/commit/babe25f)]

### [`v1.27.10`](https://redirect.github.com/onsi/gomega/releases/tag/v1.27.10)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.27.9...v1.27.10)

#### 1.27.10

##### Fixes

-   fix: go 1.21 adding goroutine ID to creator+location ([#&#8203;685](https://redirect.github.com/onsi/gomega/issues/685)) \[[`bdc7803`](https://redirect.github.com/onsi/gomega/commit/bdc7803)]

### [`v1.27.9`](https://redirect.github.com/onsi/gomega/releases/tag/v1.27.9)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.27.8...v1.27.9)

#### 1.27.9

##### Fixes

-   Prevent nil-dereference in format.Object for boxed nil error ([#&#8203;681](https://redirect.github.com/onsi/gomega/issues/681)) \[[`3b31fc3`](https://redirect.github.com/onsi/gomega/commit/3b31fc3)]

##### Maintenance

-   Bump golang.org/x/net from 0.11.0 to 0.12.0 ([#&#8203;679](https://redirect.github.com/onsi/gomega/issues/679)) \[[`360849b`](https://redirect.github.com/onsi/gomega/commit/360849b)]
-   chore: use String() instead of fmt.Sprintf ([#&#8203;678](https://redirect.github.com/onsi/gomega/issues/678)) \[[`86f3659`](https://redirect.github.com/onsi/gomega/commit/86f3659)]
-   Bump golang.org/x/net from 0.10.0 to 0.11.0 ([#&#8203;674](https://redirect.github.com/onsi/gomega/issues/674)) \[[`642ead0`](https://redirect.github.com/onsi/gomega/commit/642ead0)]
-   chore: unnecessary use of fmt.Sprintf ([#&#8203;677](https://redirect.github.com/onsi/gomega/issues/677)) \[[`ceb9ca6`](https://redirect.github.com/onsi/gomega/commit/ceb9ca6)]
-   Bump github.com/onsi/ginkgo/v2 from 2.10.0 to 2.11.0 ([#&#8203;675](https://redirect.github.com/onsi/gomega/issues/675)) \[[`a2087d8`](https://redirect.github.com/onsi/gomega/commit/a2087d8)]
-   docs: fix ContainSubstring references ([#&#8203;673](https://redirect.github.com/onsi/gomega/issues/673)) \[[`fc9a89f`](https://redirect.github.com/onsi/gomega/commit/fc9a89f)]
-   Bump github.com/onsi/ginkgo/v2 from 2.9.7 to 2.10.0 ([#&#8203;671](https://redirect.github.com/onsi/gomega/issues/671)) \[[`9076019`](https://redirect.github.com/onsi/gomega/commit/9076019)]

### [`v1.27.8`](https://redirect.github.com/onsi/gomega/releases/tag/v1.27.8)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.27.7...v1.27.8)

#### 1.27.8

##### Fixes

-   HaveExactElement should not call FailureMessage if a submatcher returned an error \[[`096f392`](https://redirect.github.com/onsi/gomega/commit/096f392)]

##### Maintenance

-   Bump github.com/onsi/ginkgo/v2 from 2.9.5 to 2.9.7 ([#&#8203;669](https://redirect.github.com/onsi/gomega/issues/669)) \[[`8884bee`](https://redirect.github.com/onsi/gomega/commit/8884bee)]

### [`v1.27.7`](https://redirect.github.com/onsi/gomega/releases/tag/v1.27.7)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.27.6...v1.27.7)

#### 1.27.7

##### Fixes

-   fix: gcustom.MakeMatcher accepts nil as actual value ([#&#8203;666](https://redirect.github.com/onsi/gomega/issues/666)) \[[`57054d5`](https://redirect.github.com/onsi/gomega/commit/57054d5)]

##### Maintenance

-   update gitignore \[[`05c1bc6`](https://redirect.github.com/onsi/gomega/commit/05c1bc6)]
-   Bump github.com/onsi/ginkgo/v2 from 2.9.4 to 2.9.5 ([#&#8203;663](https://redirect.github.com/onsi/gomega/issues/663)) \[[`7cadcf6`](https://redirect.github.com/onsi/gomega/commit/7cadcf6)]
-   Bump golang.org/x/net from 0.9.0 to 0.10.0 ([#&#8203;662](https://redirect.github.com/onsi/gomega/issues/662)) \[[`b524839`](https://redirect.github.com/onsi/gomega/commit/b524839)]
-   Bump github.com/onsi/ginkgo/v2 from 2.9.2 to 2.9.4 ([#&#8203;661](https://redirect.github.com/onsi/gomega/issues/661)) \[[`5f44694`](https://redirect.github.com/onsi/gomega/commit/5f44694)]
-   Bump commonmarker from 0.23.8 to 0.23.9 in /docs ([#&#8203;657](https://redirect.github.com/onsi/gomega/issues/657)) \[[`05dc99a`](https://redirect.github.com/onsi/gomega/commit/05dc99a)]
-   Bump nokogiri from 1.14.1 to 1.14.3 in /docs ([#&#8203;658](https://redirect.github.com/onsi/gomega/issues/658)) \[[`3a033d1`](https://redirect.github.com/onsi/gomega/commit/3a033d1)]
-   Replace deprecated NewGomegaWithT with NewWithT ([#&#8203;659](https://redirect.github.com/onsi/gomega/issues/659)) \[[`a19238f`](https://redirect.github.com/onsi/gomega/commit/a19238f)]
-   Bump golang.org/x/net from 0.8.0 to 0.9.0 ([#&#8203;656](https://redirect.github.com/onsi/gomega/issues/656)) \[[`29ed041`](https://redirect.github.com/onsi/gomega/commit/29ed041)]
-   Bump actions/setup-go from 3 to 4 ([#&#8203;651](https://redirect.github.com/onsi/gomega/issues/651)) \[[`11b2080`](https://redirect.github.com/onsi/gomega/commit/11b2080)]

### [`v1.27.6`](https://redirect.github.com/onsi/gomega/releases/tag/v1.27.6)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.27.5...v1.27.6)

#### 1.27.6

##### Fixes

-   Allow collections matchers to work correctly when expected has nil elements \[[`60e7cf3`](https://redirect.github.com/onsi/gomega/commit/60e7cf3)]

##### Maintenance

-   updates MatchError godoc comment to also accept a Gomega matcher ([#&#8203;654](https://redirect.github.com/onsi/gomega/issues/654)) \[[`67b869d`](https://redirect.github.com/onsi/gomega/commit/67b869d)]

### [`v1.27.5`](https://redirect.github.com/onsi/gomega/releases/tag/v1.27.5)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.27.4...v1.27.5)

#### 1.27.5

##### Maintenance

-   Bump github.com/onsi/ginkgo/v2 from 2.9.1 to 2.9.2 ([#&#8203;653](https://redirect.github.com/onsi/gomega/issues/653)) \[[`a215021`](https://redirect.github.com/onsi/gomega/commit/a215021)]
-   Bump github.com/go-task/slim-sprig ([#&#8203;652](https://redirect.github.com/onsi/gomega/issues/652)) \[[`a26fed8`](https://redirect.github.com/onsi/gomega/commit/a26fed8)]

### [`v1.27.4`](https://redirect.github.com/onsi/gomega/releases/tag/v1.27.4)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.27.3...v1.27.4)

#### 1.27.4

##### Fixes

-   improve error formatting and remove duplication of error message in Eventually/Consistently \[[`854f075`](https://redirect.github.com/onsi/gomega/commit/854f075)]

##### Maintenance

-   Bump github.com/onsi/ginkgo/v2 from 2.9.0 to 2.9.1 ([#&#8203;650](https://redirect.github.com/onsi/gomega/issues/650)) \[[`ccebd9b`](https://redirect.github.com/onsi/gomega/commit/ccebd9b)]

### [`v1.27.3`](https://redirect.github.com/onsi/gomega/releases/tag/v1.27.3)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.27.2...v1.27.3)

#### 1.27.3

##### Fixes

-   format.Object now always includes err.Error() when passed an error \[[`86d97ef`](https://redirect.github.com/onsi/gomega/commit/86d97ef)]
-   Fix HaveExactElements to work inside ContainElement or other collection matchers ([#&#8203;648](https://redirect.github.com/onsi/gomega/issues/648)) \[[`636757e`](https://redirect.github.com/onsi/gomega/commit/636757e)]

##### Maintenance

-   Bump github.com/golang/protobuf from 1.5.2 to 1.5.3 ([#&#8203;649](https://redirect.github.com/onsi/gomega/issues/649)) \[[`cc16689`](https://redirect.github.com/onsi/gomega/commit/cc16689)]
-   Bump github.com/onsi/ginkgo/v2 from 2.8.4 to 2.9.0 ([#&#8203;646](https://redirect.github.com/onsi/gomega/issues/646)) \[[`e783366`](https://redirect.github.com/onsi/gomega/commit/e783366)]

### [`v1.27.2`](https://redirect.github.com/onsi/gomega/releases/tag/v1.27.2)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.27.1...v1.27.2)

#### 1.27.2

##### Fixes

-   improve poll progress message when polling a consistently that has been passing \[[`28a319b`](https://redirect.github.com/onsi/gomega/commit/28a319b)]

##### Maintenance

-   bump ginkgo
-   remove tools.go hack as Ginkgo 2.8.2 automatically pulls in the cli dependencies \[[`81443b3`](https://redirect.github.com/onsi/gomega/commit/81443b3)]

### [`v1.27.1`](https://redirect.github.com/onsi/gomega/releases/tag/v1.27.1)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.27.0...v1.27.1)

#### 1.27.1

##### Maintenance

-   Bump golang.org/x/net from 0.6.0 to 0.7.0 ([#&#8203;640](https://redirect.github.com/onsi/gomega/issues/640)) \[[`bc686cd`](https://redirect.github.com/onsi/gomega/commit/bc686cd)]

### [`v1.27.0`](https://redirect.github.com/onsi/gomega/releases/tag/v1.27.0)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.26.0...v1.27.0)

#### 1.27.0

##### Features

-   Add HaveExactElements matcher ([#&#8203;634](https://redirect.github.com/onsi/gomega/issues/634)) \[[`9d50783`](https://redirect.github.com/onsi/gomega/commit/9d50783)]
-   update Gomega docs to discuss GinkgoHelper() \[[`be32774`](https://redirect.github.com/onsi/gomega/commit/be32774)]

##### Maintenance

-   Bump github.com/onsi/ginkgo/v2 from 2.8.0 to 2.8.1 ([#&#8203;639](https://redirect.github.com/onsi/gomega/issues/639)) \[[`296a68b`](https://redirect.github.com/onsi/gomega/commit/296a68b)]
-   Bump golang.org/x/net from 0.5.0 to 0.6.0 ([#&#8203;638](https://redirect.github.com/onsi/gomega/issues/638)) \[[`c2b098b`](https://redirect.github.com/onsi/gomega/commit/c2b098b)]
-   Bump github-pages from 227 to 228 in /docs ([#&#8203;636](https://redirect.github.com/onsi/gomega/issues/636)) \[[`a9069ab`](https://redirect.github.com/onsi/gomega/commit/a9069ab)]
-   test: update matrix for Go 1.20 ([#&#8203;635](https://redirect.github.com/onsi/gomega/issues/635)) \[[`6bd25c8`](https://redirect.github.com/onsi/gomega/commit/6bd25c8)]
-   Bump github.com/onsi/ginkgo/v2 from 2.7.0 to 2.8.0 ([#&#8203;631](https://redirect.github.com/onsi/gomega/issues/631)) \[[`5445f8b`](https://redirect.github.com/onsi/gomega/commit/5445f8b)]
-   Bump webrick from 1.7.0 to 1.8.1 in /docs ([#&#8203;630](https://redirect.github.com/onsi/gomega/issues/630)) \[[`03e93bb`](https://redirect.github.com/onsi/gomega/commit/03e93bb)]
-   codeql: add ruby language ([#&#8203;626](https://redirect.github.com/onsi/gomega/issues/626)) \[[`63c7d21`](https://redirect.github.com/onsi/gomega/commit/63c7d21)]
-   dependabot: add bundler package-ecosystem for docs ([#&#8203;625](https://redirect.github.com/onsi/gomega/issues/625)) \[[`d92f963`](https://redirect.github.com/onsi/gomega/commit/d92f963)]

### [`v1.26.0`](https://redirect.github.com/onsi/gomega/releases/tag/v1.26.0)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.25.0...v1.26.0)

#### 1.26.0

##### Features

-   When a polled function returns an error, keep track of the actual and report on the matcher state of the last non-errored actual \[[`21f3090`](https://redirect.github.com/onsi/gomega/commit/21f3090)]
-   improve eventually failure message output \[[`c530fb3`](https://redirect.github.com/onsi/gomega/commit/c530fb3)]

##### Fixes

-   fix several documentation spelling issues \[[`e2eff1f`](https://redirect.github.com/onsi/gomega/commit/e2eff1f)]

### [`v1.25.0`](https://redirect.github.com/onsi/gomega/releases/tag/v1.25.0)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.24.2...v1.25.0)

#### 1.25.0

##### Features

-   add `MustPassRepeatedly(int)` to asyncAssertion ([#&#8203;619](https://redirect.github.com/onsi/gomega/issues/619)) \[[`4509f72`](https://redirect.github.com/onsi/gomega/commit/4509f72)]
-   compare unwrapped errors using DeepEqual ([#&#8203;617](https://redirect.github.com/onsi/gomega/issues/617)) \[[`aaeaa5d`](https://redirect.github.com/onsi/gomega/commit/aaeaa5d)]

##### Maintenance

-   Bump golang.org/x/net from 0.4.0 to 0.5.0 ([#&#8203;614](https://redirect.github.com/onsi/gomega/issues/614)) \[[`c7cfea4`](https://redirect.github.com/onsi/gomega/commit/c7cfea4)]
-   Bump github.com/onsi/ginkgo/v2 from 2.6.1 to 2.7.0 ([#&#8203;615](https://redirect.github.com/onsi/gomega/issues/615)) \[[`71b8adb`](https://redirect.github.com/onsi/gomega/commit/71b8adb)]
-   Docs: Fix typo "MUltiple" -> "Multiple" ([#&#8203;616](https://redirect.github.com/onsi/gomega/issues/616)) \[[`9351dda`](https://redirect.github.com/onsi/gomega/commit/9351dda)]
-   clean up go.sum \[[`cd1dc1d`](https://redirect.github.com/onsi/gomega/commit/cd1dc1d)]

### [`v1.24.2`](https://redirect.github.com/onsi/gomega/releases/tag/v1.24.2)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.24.1...v1.24.2)

#### 1.24.2

##### Fixes

-   Correctly handle assertion failure panics for eventually/consistnetly "g Gomega"s in a goroutine \[[`78f1660`](https://redirect.github.com/onsi/gomega/commit/78f1660)]
-   docs:Fix typo "you an" -> "you can" ([#&#8203;607](https://redirect.github.com/onsi/gomega/issues/607)) \[[`3187c1f`](https://redirect.github.com/onsi/gomega/commit/3187c1f)]
-   fixes issue [#&#8203;600](https://redirect.github.com/onsi/gomega/issues/600) ([#&#8203;606](https://redirect.github.com/onsi/gomega/issues/606)) \[[`808d192`](https://redirect.github.com/onsi/gomega/commit/808d192)]

##### Maintenance

-   Bump golang.org/x/net from 0.2.0 to 0.4.0 ([#&#8203;611](https://redirect.github.com/onsi/gomega/issues/611)) \[[`6ebc0bf`](https://redirect.github.com/onsi/gomega/commit/6ebc0bf)]
-   Bump nokogiri from 1.13.9 to 1.13.10 in /docs ([#&#8203;612](https://redirect.github.com/onsi/gomega/issues/612)) \[[`258cfc8`](https://redirect.github.com/onsi/gomega/commit/258cfc8)]
-   Bump github.com/onsi/ginkgo/v2 from 2.5.0 to 2.5.1 ([#&#8203;609](https://redirect.github.com/onsi/gomega/issues/609)) \[[`e6c3eb9`](https://redirect.github.com/onsi/gomega/commit/e6c3eb9)]

### [`v1.24.1`](https://redirect.github.com/onsi/gomega/compare/v1.24.0...v1.24.1)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.24.0...v1.24.1)

### [`v1.24.0`](https://redirect.github.com/onsi/gomega/releases/tag/v1.24.0)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.23.0...v1.24.0)

#### 1.24.0

##### Features

Introducting [gcustom](https://onsi.github.io/gomega/#gcustom-a-convenient-mechanism-for-buildling-custom-matchers) - a convenient mechanism for building custom matchers.

This is an RC release for `gcustom`.  The external API may be tweaked in response to feedback however it is expected to remain mostly stable.

##### Maintenance

-   Update BeComparableTo documentation \[[`756eaa0`](https://redirect.github.com/onsi/gomega/commit/756eaa0)]

### [`v1.23.0`](https://redirect.github.com/onsi/gomega/releases/tag/v1.23.0)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.22.1...v1.23.0)

#### 1.23.0

##### Features

-   Custom formatting on a per-type basis can be provided using `format.RegisterCustomFormatter()` -- see the docs [here](https://onsi.github.io/gomega/#adjusting-output)

-   Substantial improvement have been made to `StopTrying()`:
    -   Users can now use `StopTrying().Wrap(err)` to wrap errors and `StopTrying().Attach(description, object)` to attach arbitrary objects to the `StopTrying()` error
    -   `StopTrying()` is now always interpreted as a failure.  If you are an early adopter of `StopTrying()` you may need to change your code as the prior version would match against the returned value even if `StopTrying()` was returned.  Going forward the `StopTrying()` api should remain stable.
    -   `StopTrying()` and `StopTrying().Now()` can both be used in matchers - not just polled functions.

-   `TryAgainAfter(duration)` is used like `StopTrying()` but instructs `Eventually` and `Consistently` that the poll should be tried again after the specified duration.  This allows you to dynamically adjust the polling duration.

-   `ctx` can now be passed-in as the first argument to `Eventually` and `Consistently`.

#### Maintenance

-   Bump github.com/onsi/ginkgo/v2 from 2.3.0 to 2.3.1 ([#&#8203;597](https://redirect.github.com/onsi/gomega/issues/597)) \[[`afed901`](https://redirect.github.com/onsi/gomega/commit/afed901)]
-   Bump nokogiri from 1.13.8 to 1.13.9 in /docs ([#&#8203;599](https://redirect.github.com/onsi/gomega/issues/599)) \[[`7c691b3`](https://redirect.github.com/onsi/gomega/commit/7c691b3)]
-   Bump github.com/google/go-cmp from 0.5.8 to 0.5.9 ([#&#8203;587](https://redirect.github.com/onsi/gomega/issues/587)) \[[`ff22665`](https://redirect.github.com/onsi/gomega/commit/ff22665)]

### [`v1.22.1`](https://redirect.github.com/onsi/gomega/releases/tag/v1.22.1)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.22.0...v1.22.1)

#### 1.22.1

#### Fixes

-   When passed a context and no explicit timeout, Eventually will only timeout when the context is cancelled \[[`e5105cf`](https://redirect.github.com/onsi/gomega/commit/e5105cf)]
-   Allow StopTrying() to be wrapped \[[`bf3cba9`](https://redirect.github.com/onsi/gomega/commit/bf3cba9)]

#### Maintenance

-   bump to ginkgo v2.3.0 \[[`c5d5c39`](https://redirect.github.com/onsi/gomega/commit/c5d5c39)]

### [`v1.22.0`](https://redirect.github.com/onsi/gomega/releases/tag/v1.22.0)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.21.1...v1.22.0)

#### 1.22.0

##### Features

Several improvements have been made to `Eventually` and `Consistently` in this and the most recent releases:

-   Eventually and Consistently can take a context.Context \[[`65c01bc`](https://redirect.github.com/onsi/gomega/commit/65c01bc)]
    This enables integration with Ginkgo 2.3.0's interruptible nodes and node timeouts.
-   Eventually and Consistently that are passed a SpecContext can provide reports when an interrupt occurs \[[`0d063c9`](https://redirect.github.com/onsi/gomega/commit/0d063c9)]
-   Eventually/Consistently will forward an attached context to functions that ask for one \[[`e2091c5`](https://redirect.github.com/onsi/gomega/commit/e2091c5)]
-   Eventually/Consistently supports passing arguments to functions via WithArguments() \[[`a2dc7c3`](https://redirect.github.com/onsi/gomega/commit/a2dc7c3)]
-   Eventually and Consistently can now be stopped early with StopTrying(message) and StopTrying(message).Now() \[[`52976bb`](https://redirect.github.com/onsi/gomega/commit/52976bb)]

These improvements are all documented in [Gomega's docs](https://onsi.github.io/gomega/#making-asynchronous-assertions)

### [`v1.21.1`](https://redirect.github.com/onsi/gomega/releases/tag/v1.21.1)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.21.0...v1.21.1)

#### v1.21.1

##### Features

-   Eventually and Consistently that are passed a SpecContext can provide reports when an interrupt occurs \[[`0d063c9`](https://redirect.github.com/onsi/gomega/commit/0d063c9)]

### [`v1.21.0`](https://redirect.github.com/onsi/gomega/releases/tag/v1.21.0)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.20.2...v1.21.0)

#### 1.21.0

##### Features

-   Eventually and Consistently can take a context.Context \[[`65c01bc`](https://redirect.github.com/onsi/gomega/commit/65c01bc)]
    This enables integration with Ginkgo 2.3.0's interruptible nodes and node timeouts.
-   Introduces Eventually.Within.ProbeEvery with tests and documentation ([#&#8203;591](https://redirect.github.com/onsi/gomega/issues/591)) \[[`f633800`](https://redirect.github.com/onsi/gomega/commit/f633800)]
-   New BeKeyOf matcher with documentation and unit tests ([#&#8203;590](https://redirect.github.com/onsi/gomega/issues/590)) \[[`fb586b3`](https://redirect.github.com/onsi/gomega/commit/fb586b3)]

#### Fixes

-   Cover the entire gmeasure suite with leak detection \[[`8c54344`](https://redirect.github.com/onsi/gomega/commit/8c54344)]
-   Fix gmeasure leak \[[`119d4ce`](https://redirect.github.com/onsi/gomega/commit/119d4ce)]
-   Ignore new Ginkgo ProgressSignal goroutine in gleak \[[`ba548e2`](https://redirect.github.com/onsi/gomega/commit/ba548e2)]

#### Maintenance

-   Fixes crashes on newer Ruby 3 installations by upgrading github-pages gem dependency ([#&#8203;596](https://redirect.github.com/onsi/gomega/issues/596)) \[[`12469a0`](https://redirect.github.com/onsi/gomega/commit/12469a0)]

### [`v1.20.2`](https://redirect.github.com/onsi/gomega/releases/tag/v1.20.2)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.20.1...v1.20.2)

#### 1.20.2

#### Fixes

-   label specs that rely on remote access; bump timeout on short-circuit test to make it less flaky \[[`35eeadf`](https://redirect.github.com/onsi/gomega/commit/35eeadf)]
-   gexec: allow more headroom for SIGABRT-related unit tests ([#&#8203;581](https://redirect.github.com/onsi/gomega/issues/581)) \[[`5b78f40`](https://redirect.github.com/onsi/gomega/commit/5b78f40)]
-   Enable reading from a closed gbytes.Buffer ([#&#8203;575](https://redirect.github.com/onsi/gomega/issues/575)) \[[`061fd26`](https://redirect.github.com/onsi/gomega/commit/061fd26)]

#### Maintenance

-   Bump github.com/onsi/ginkgo/v2 from 2.1.5 to 2.1.6 ([#&#8203;583](https://redirect.github.com/onsi/gomega/issues/583)) \[[`55d895b`](https://redirect.github.com/onsi/gomega/commit/55d895b)]
-   Bump github.com/onsi/ginkgo/v2 from 2.1.4 to 2.1.5 ([#&#8203;582](https://redirect.github.com/onsi/gomega/issues/582)) \[[`346de7c`](https://redirect.github.com/onsi/gomega/commit/346de7c)]

### [`v1.20.1`](https://redirect.github.com/onsi/gomega/releases/tag/v1.20.1)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.20.0...v1.20.1)

#### 1.20.1

#### Fixes

-   fix false positive gleaks when using ginkgo -p ([#&#8203;577](https://redirect.github.com/onsi/gomega/issues/577)) \[[`cb46517`](https://redirect.github.com/onsi/gomega/commit/cb46517)]
-   Fix typos in gomega_dsl.go ([#&#8203;569](https://redirect.github.com/onsi/gomega/issues/569)) \[[`5f71ed2`](https://redirect.github.com/onsi/gomega/commit/5f71ed2)]
-   don't panic on Eventually(nil), fixing [#&#8203;555](https://redirect.github.com/onsi/gomega/issues/555) ([#&#8203;567](https://redirect.github.com/onsi/gomega/issues/567)) \[[`9d1186f`](https://redirect.github.com/onsi/gomega/commit/9d1186f)]
-   vet optional description args in assertions, fixing [#&#8203;560](https://redirect.github.com/onsi/gomega/issues/560) ([#&#8203;566](https://redirect.github.com/onsi/gomega/issues/566)) \[[`8e37808`](https://redirect.github.com/onsi/gomega/commit/8e37808)]

#### Maintenance

-   test: add new Go 1.19 to test matrix ([#&#8203;571](https://redirect.github.com/onsi/gomega/issues/571)) \[[`40d7efe`](https://redirect.github.com/onsi/gomega/commit/40d7efe)]
-   Bump tzinfo from 1.2.9 to 1.2.10 in /docs ([#&#8203;564](https://redirect.github.com/onsi/gomega/issues/564)) \[[`5f26371`](https://redirect.github.com/onsi/gomega/commit/5f26371)]

### [`v1.20.0`](https://redirect.github.com/onsi/gomega/releases/tag/v1.20.0)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.19.0...v1.20.0)

#### Features

-   New [`gleak`](https://onsi.github.io/gomega/#codegleakcode-finding-leaked-goroutines) experimental goroutine leak detection package! ([#&#8203;538](https://redirect.github.com/onsi/gomega/issues/538)) \[[`85ba7bc`](https://redirect.github.com/onsi/gomega/commit/85ba7bc)]
-   New `BeComparableTo` matcher([#&#8203;546](https://redirect.github.com/onsi/gomega/issues/546)) that uses `gocmp` to make comparisons \[[`e77ea75`](https://redirect.github.com/onsi/gomega/commit/e77ea75)]
-   New `HaveExistingField` matcher ([#&#8203;553](https://redirect.github.com/onsi/gomega/issues/553)) \[[`fd130e1`](https://redirect.github.com/onsi/gomega/commit/fd130e1)]
-   Document how to wrap Gomega ([#&#8203;539](https://redirect.github.com/onsi/gomega/issues/539)) \[[`56714a4`](https://redirect.github.com/onsi/gomega/commit/56714a4)]

#### Fixes

-   Support pointer receivers in HaveField; fixes [#&#8203;543](https://redirect.github.com/onsi/gomega/issues/543) ([#&#8203;544](https://redirect.github.com/onsi/gomega/issues/544)) \[[`8dab36e`](https://redirect.github.com/onsi/gomega/commit/8dab36e)]

#### Maintenance

-   Bump various dependencies:
    -   Upgrade to yaml.v3 ([#&#8203;556](https://redirect.github.com/onsi/gomega/issues/556)) \[[`f5a83b1`](https://redirect.github.com/onsi/gomega/commit/f5a83b1)]
    -   Bump github/codeql-action from 1 to 2 ([#&#8203;549](https://redirect.github.com/onsi/gomega/issues/549)) \[[`52f5adf`](https://redirect.github.com/onsi/gomega/commit/52f5adf)]
    -   Bump github.com/google/go-cmp from 0.5.7 to 0.5.8 ([#&#8203;551](https://redirect.github.com/onsi/gomega/issues/551)) \[[`5f3942d`](https://redirect.github.com/onsi/gomega/commit/5f3942d)]
    -   Bump nokogiri from 1.13.4 to 1.13.6 in /docs ([#&#8203;554](https://redirect.github.com/onsi/gomega/issues/554)) \[[`eb4b4c2`](https://redirect.github.com/onsi/gomega/commit/eb4b4c2)]
    -   Use latest ginkgo ([#&#8203;535](https://redirect.github.com/onsi/gomega/issues/535)) \[[`1c29028`](https://redirect.github.com/onsi/gomega/commit/1c29028)]
    -   Bump nokogiri from 1.13.3 to 1.13.4 in /docs ([#&#8203;541](https://redirect.github.com/onsi/gomega/issues/541)) \[[`1ce84d5`](https://redirect.github.com/onsi/gomega/commit/1ce84d5)]
    -   Bump actions/setup-go from 2 to 3 ([#&#8203;540](https://redirect.github.com/onsi/gomega/issues/540)) \[[`755485e`](https://redirect.github.com/onsi/gomega/commit/755485e)]
    -   Bump nokogiri from 1.12.5 to 1.13.3 in /docs ([#&#8203;522](https://redirect.github.com/onsi/gomega/issues/522)) \[[`4fbb0dc`](https://redirect.github.com/onsi/gomega/commit/4fbb0dc)]
    -   Bump actions/checkout from 2 to 3 ([#&#8203;526](https://redirect.github.com/onsi/gomega/issues/526)) \[[`ac49202`](https://redirect.github.com/onsi/gomega/commit/ac49202)]

#### 1.19.0

#### Features

-   New [`HaveEach`](https://onsi.github.io/gomega/#haveeachelement-interface) matcher to ensure that each and every element in an `array`, `slice`, or `map` satisfies the passed in matcher. ([#&#8203;523](https://redirect.github.com/onsi/gomega/issues/523)) \[[`9fc2ae2`](https://redirect.github.com/onsi/gomega/commit/9fc2ae2)] ([#&#8203;524](https://redirect.github.com/onsi/gomega/issues/524)) \[[`c8ba582`](https://redirect.github.com/onsi/gomega/commit/c8ba582)]
-   Users can now wrap the `Gomega` interface to implement custom behavior on each assertion. ([#&#8203;521](https://redirect.github.com/onsi/gomega/issues/521)) \[[`1f2e714`](https://redirect.github.com/onsi/gomega/commit/1f2e714)]
-   [`ContainElement`](https://onsi.github.io/gomega/#containelementelement-interface) now accepts an additional pointer argument.  Elements that satisfy the matcher are stored in the pointer enabling developers to easily add subsequent, more detailed, assertions against the matching element. ([#&#8203;527](https://redirect.github.com/onsi/gomega/issues/527)) \[[`1a4e27f`](https://redirect.github.com/onsi/gomega/commit/1a4e27f)]

#### Fixes

-   update RELEASING instructions to match ginkgo \[[`0917cde`](https://redirect.github.com/onsi/gomega/commit/0917cde)]
-   Bump github.com/onsi/ginkgo/v2 from 2.0.0 to 2.1.3 ([#&#8203;519](https://redirect.github.com/onsi/gomega/issues/519)) \[[`49ab4b0`](https://redirect.github.com/onsi/gomega/commit/49ab4b0)]
-   Fix CVE-2021-38561 ([#&#8203;534](https://redirect.github.com/onsi/gomega/issues/534)) \[[`f1b4456`](https://redirect.github.com/onsi/gomega/commit/f1b4456)]
-   Fix max number of samples in experiments on non-64-bit systems. ([#&#8203;528](https://redirect.github.com/onsi/gomega/issues/528)) \[[`1c84497`](https://redirect.github.com/onsi/gomega/commit/1c84497)]
-   Remove dependency on ginkgo v1.16.4 ([#&#8203;530](https://redirect.github.com/onsi/gomega/issues/530)) \[[`4dea8d5`](https://redirect.github.com/onsi/gomega/commit/4dea8d5)]
-   Fix for Go 1.18 ([#&#8203;532](https://redirect.github.com/onsi/gomega/issues/532)) \[[`56d2a29`](https://redirect.github.com/onsi/gomega/commit/56d2a29)]
-   Document precendence of timeouts ([#&#8203;533](https://redirect.github.com/onsi/gomega/issues/533)) \[[`b607941`](https://redirect.github.com/onsi/gomega/commit/b607941)]

### [`v1.19.0`](https://redirect.github.com/onsi/gomega/releases/tag/v1.19.0)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.18.1...v1.19.0)

#### Features

-   New [`HaveEach`](https://onsi.github.io/gomega/#haveeachelement-interface) matcher to ensure that each and every element in an `array`, `slice`, or `map` satisfies the passed in matcher. ([#&#8203;523](https://redirect.github.com/onsi/gomega/issues/523)) \[[`9fc2ae2`](https://redirect.github.com/onsi/gomega/commit/9fc2ae2)] ([#&#8203;524](https://redirect.github.com/onsi/gomega/issues/524)) \[[`c8ba582`](https://redirect.github.com/onsi/gomega/commit/c8ba582)]
-   Users can now wrap the `Gomega` interface to implement custom behavior on each assertion. ([#&#8203;521](https://redirect.github.com/onsi/gomega/issues/521)) \[[`1f2e714`](https://redirect.github.com/onsi/gomega/commit/1f2e714)]
-   [`ContainElement`](https://onsi.github.io/gomega/#containelementelement-interface) now accepts an additional pointer argument.  Elements that satisfy the matcher are stored in the pointer enabling developers to easily add subsequent, more detailed, assertions against the matching element. ([#&#8203;527](https://redirect.github.com/onsi/gomega/issues/527)) \[[`1a4e27f`](https://redirect.github.com/onsi/gomega/commit/1a4e27f)]

#### Fixes

-   update RELEASING instructions to match ginkgo \[[`0917cde`](https://redirect.github.com/onsi/gomega/commit/0917cde)]
-   Bump github.com/onsi/ginkgo/v2 from 2.0.0 to 2.1.3 ([#&#8203;519](https://redirect.github.com/onsi/gomega/issues/519)) \[[`49ab4b0`](https://redirect.github.com/onsi/gomega/commit/49ab4b0)]
-   Fix CVE-2021-38561 ([#&#8203;534](https://redirect.github.com/onsi/gomega/issues/534)) \[[`f1b4456`](https://redirect.github.com/onsi/gomega/commit/f1b4456)]
-   Fix max number of samples in experiments on non-64-bit systems. ([#&#8203;528](https://redirect.github.com/onsi/gomega/issues/528)) \[[`1c84497`](https://redirect.github.com/onsi/gomega/commit/1c84497)]
-   Remove dependency on ginkgo v1.16.4 ([#&#8203;530](https://redirect.github.com/onsi/gomega/issues/530)) \[[`4dea8d5`](https://redirect.github.com/onsi/gomega/commit/4dea8d5)]
-   Fix for Go 1.18 ([#&#8203;532](https://redirect.github.com/onsi/gomega/issues/532)) \[[`56d2a29`](https://redirect.github.com/onsi/gomega/commit/56d2a29)]
-   Document precendence of timeouts ([#&#8203;533](https://redirect.github.com/onsi/gomega/issues/533)) \[[`b607941`](https://redirect.github.com/onsi/gomega/commit/b607941)]

### [`v1.18.1`](https://redirect.github.com/onsi/gomega/releases/tag/v1.18.1)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.18.0...v1.18.1)

-   add pointer support to HaveField matcher ([#&#8203;495](https://redirect.github.com/onsi/gomega/issues/495)) \[[`79e41a3`](https://redirect.github.com/onsi/gomega/commit/79e41a3)]

### [`v1.18.0`](https://redirect.github.com/onsi/gomega/releases/tag/v1.18.0)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.17.0...v1.18.0)

#### Features

-   Docs now live on the master branch in the docs folder which will make for easier PRs.  The docs also use Ginkgo 2.0's new docs html/css/js. \[[`2570272`](https://redirect.github.com/onsi/gomega/commit/2570272)]
-   New HaveValue matcher can handle actuals that are either values (in which case they are passed on unscathed) or pointers (in which case they are indirected).  [Docs here.](https://onsi.github.io/gomega/#working-with-values) ([#&#8203;485](https://redirect.github.com/onsi/gomega/issues/485)) \[[`bdc087c`](https://redirect.github.com/onsi/gomega/commit/bdc087c)]
-   Gmeasure has been declared GA \[[`360db9d`](https://redirect.github.com/onsi/gomega/commit/360db9d)]

#### Fixes

-   Gomega now uses ioutil for Go 1.15 and lower ([#&#8203;492](https://redirect.github.com/onsi/gomega/issues/492)) - official support is only for the most recent two major versions of Go but this will unblock users who need to stay on older unsupported versions of Go. \[[`c29c1c0`](https://redirect.github.com/onsi/gomega/commit/c29c1c0)]

#### Maintenace

-   Remove Travis workflow ([#&#8203;491](https://redirect.github.com/onsi/gomega/issues/491)) \[[`72e6040`](https://redirect.github.com/onsi/gomega/commit/72e6040)]
-   Upgrade to Ginkgo 2.0.0 GA \[[`f383637`](https://redirect.github.com/onsi/gomega/commit/f383637)]
-   chore: fix description of HaveField matcher ([#&#8203;487](https://redirect.github.com/onsi/gomega/issues/487)) \[[`2b4b2c0`](https://redirect.github.com/onsi/gomega/commit/2b4b2c0)]
-   use tools.go to ensure Ginkgo cli dependencies are included \[[`f58a52b`](https://redirect.github.com/onsi/gomega/commit/f58a52b)]
-   remove dockerfile and simplify github actions to match ginkgo's actions \[[`3f8160d`](https://redirect.github.com/onsi/gomega/commit/3f8160d)]

### [`v1.17.0`](https://redirect.github.com/onsi/gomega/releases/tag/v1.17.0)

[Compare Source](https://redirect.github.com/onsi/gomega/compare/v1.16.0...v1.17.0)

#### 1.17.0

##### Features

-   Add HaveField matcher \[[`3a26311`](https://redirect.github.com/onsi/gomega/commit/3a26311)]
-   add Error() assertions on the final error value of multi-return values ([#&#8203;480](https://redirect.github.com/onsi/gomega/issues/480)) \[[`2f96943`](https://redirect.github.com/onsi/gomega/commit/2f96943)]
-   separate out offsets and timeouts ([#&#8203;478](https://redirect.github.com/onsi/gomega/issues/478)) \[[`18a4723`](https://redirect.github.com/onsi/gomega/commit/18a4723)]
-   fix transformation error reporting ([#&#8203;479](https://redirect.github.com/onsi/gomega/issues/479)) \[[`e001fab`](https://redirect.github.com/onsi/gomega/commit/e001fab)]
-   allow transform functions to report errors ([#&#8203;472](https://redirect.github.com/onsi/gomega/issues/472)) \[[`bf93408`](https://redirect.github.com/onsi/gomega/commit/bf93408)]

##### Fixes

Stop using deprecated ioutil package ([#&#8203;467](https://redirect.github.com/onsi/gomega/issues/467)) \[[`07f405d`](https://redirect.github.com/onsi/gomega/commit/07f405d)]

</details>

---

### Configuration

📅 **Schedule**: Branch creation - "after 5am on sunday" in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOS4xNTguMC1ycG0iLCJ1cGRhdGVkSW5WZXIiOiIzOS4xNTguMC1ycG0iLCJ0YXJnZXRCcmFuY2giOiJtYXN0ZXIiLCJsYWJlbHMiOltdfQ==-->


---

## Discussion

### Comment by @jira-linking on April 27, 2025 at 02:27 PM UTC

Commits missing Jira IDs:
a6c8efb95a37b14f00158a15372d4272c7c56cb6


### Comment by @codecov-commenter on April 27, 2025 at 02:32 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1648?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 58.19%. Comparing base [(`1aa08aa`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/1aa08aa94412571600ec83334da23517d8601a38?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`a6c8efb`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/a6c8efb95a37b14f00158a15372d4272c7c56cb6?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1648      +/-   ##
==========================================
- Coverage   58.21%   58.19%   -0.03%     
==========================================
  Files         146      146              
  Lines       11397    11397              
==========================================
- Hits         6635     6632       -3     
- Misses       4175     4178       +3     
  Partials      587      587              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1648/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1648/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.19% <ø> (-0.03%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1648?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1648*
