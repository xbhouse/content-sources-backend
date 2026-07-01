---
type: pull_request
number: 394
title: "Update app build to be compatible with chrome 2.0"
state: merged
author: Hyperkid123
created: 2021-01-18T13:59:50Z
updated: 2021-01-25T13:21:17Z
closed: 2021-01-25T13:13:04Z
merged: 2021-01-25T13:13:04Z
base_branch: master
head_branch: migrate-webpack-5
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/394
---

# Pull Request #394: Update app build to be compatible with chrome 2.0

**Author**: @Hyperkid123
**Created**: January 18, 2021 at 01:59 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `migrate-webpack-5`

## Description

This PR makes the Patchman UI compatible with Chrome 2.0. The application will still work with current chrome, but in order to deploy chrome 2.0, all apps that are using the inventory components, have to be migrated first due to a change in the inventory distribution system. This PR is slightly bigger than usual due to some old node packages which are no longer supported in current version of webpack.

Please review this PR thoroughly. The whole build process is now altered and we have to make sure it still works perfectly.

### Changes
- update dependencies
  - there were a lot of dependencies that are no longer compatible with webpack 5. I took the liberty of updating all dev-dependencies and FEC dependencies
  - I have also noticed that the lock file was generated with npm@7+. I just want to point out that most of the FEC dependencies  and PF react-core (next release fixes it) are not yet compatible with npm 7 so I reccomend would avoid using it for now
- Use updated FEC config to create chrome 2.0 required build
- Use chrome 2.0 compatible inventory components
  - current system does not work with chrome 2.0 inventory which is using federated modules to distribute Inventory components
- Fixed async tests in AdvisoryDetail components
  -  async components were not waiting for the async lifecycle to run which lead to false positive/negative results based on in which order are those tests run.
  - Now tests cases are using unique instances of the component where necessary and are waiting for the instance async lifecycle to complete to get the correct test results

cc @jiridostal 

---

## Discussion

### Comment by @jiridostal on January 20, 2021 at 09:50 AM UTC

@Hyperkid123 Unfortunately, there is an issue with system detail page. It triggers some error that caused the whole component tree to be recreated from error boundary and gets in a loop, causing a stuck browser... :-(


### Comment by @Hyperkid123 on January 20, 2021 at 10:20 AM UTC

@jiridostal it is actually supposed to load the ErrorComponent. That is a fallback until chrome 2.0 is available and it uses the current inventory. It's by design.

However, the loading loop is not. I will look at why is it happening.

### Comment by @Hyperkid123 on January 20, 2021 at 12:52 PM UTC

@jiridostal turns out that I've made a silly mistake. I've assigned the `onLoad` function prop to the wrong element. Instead of the `InventoryDetailHead` to your `Header` component.

I will publish the fix and also deal with the conflicts.

### Comment by @codecov-io on January 20, 2021 at 12:59 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/394?src=pr&el=h1) Report
> Merging [#394](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/394?src=pr&el=desc) (4c72d3d) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/059785a9c49f1f390abfb0da911b5adb3930a645?el=desc) (059785a) will **decrease** coverage by `0.03%`.
> The diff coverage is `44.82%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/394/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/394?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #394      +/-   ##
==========================================
- Coverage   72.85%   72.81%   -0.04%     
==========================================
  Files          69       72       +3     
  Lines        1234     1214      -20     
  Branches      160      160              
==========================================
- Hits          899      884      -15     
+ Misses        280      276       -4     
+ Partials       55       54       -1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/394?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [src/App.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/394/diff?src=pr&el=tree#diff-c3JjL0FwcC5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/AppEntry.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/394/diff?src=pr&el=tree#diff-c3JjL0FwcEVudHJ5Lmpz) | `0.00% <0.00%> (ø)` | |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/394/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/SmartComponents/SystemDetail/InventoryPage.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/394/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5UGFnZS5qcw==) | `75.00% <0.00%> (-25.00%)` | :arrow_down: |
| [src/bootstrap-dev.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/394/diff?src=pr&el=tree#diff-c3JjL2Jvb3RzdHJhcC1kZXYuanM=) | `0.00% <0.00%> (ø)` | |
| [src/bootstrap.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/394/diff?src=pr&el=tree#diff-c3JjL2Jvb3RzdHJhcC5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/entry-dev.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/394/diff?src=pr&el=tree#diff-c3JjL2VudHJ5LWRldi5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/entry.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/394/diff?src=pr&el=tree#diff-c3JjL2VudHJ5Lmpz) | `0.00% <0.00%> (ø)` | |
| [src/store/index.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/394/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL2luZGV4Lmpz) | `4.00% <0.00%> (ø)` | |
| [src/Routes.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/394/diff?src=pr&el=tree#diff-c3JjL1JvdXRlcy5qcw==) | `34.28% <50.00%> (ø)` | |
| ... and [8 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/394/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/394?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/394?src=pr&el=footer). Last update [059785a...4c72d3d](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/394?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on January 25, 2021 at 01:21 PM UTC

:tada: This PR is included in version 1.8.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.8.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.8.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/394*
