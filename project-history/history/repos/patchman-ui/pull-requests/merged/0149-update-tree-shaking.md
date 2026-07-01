---
type: pull_request
number: 149
title: "Update tree shaking"
state: merged
author: Hyperkid123
created: 2020-06-04T11:37:13Z
updated: 2020-06-08T11:20:46Z
closed: 2020-06-05T11:41:17Z
merged: 2020-06-05T11:41:17Z
base_branch: master
head_branch: update-tree-shaking
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/149
---

# Pull Request #149: Update tree shaking

**Author**: @Hyperkid123
**Created**: June 04, 2020 at 11:37 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `update-tree-shaking`

## Description

closes #148 

**Please check the whole app as wrong imports may cause a runtime crash.** I tried to click through everything but I a not familiar with the application so I might have missed something.

### Changes
- added bundle analyzer to enable precise build debugging
  - you can run the analyzer with `npm run analyze` / `yarn analyze`
- allowed multiple vendors chunks for js files
  - before all files from node_modules were consolidated into one huge chunk
  - now all lazy-loaded modules have their own vendor chunks which can be shared between multiple modules
- manually changed import paths of FCE to use absolute paths to CJS bundles
- added babel transform imports plugin and configured it for pattern fly assets
  - this will automatically create correct import paths
  - **I strongly recommend doing the same for the FCE packages you can use the PF config as a baseline. Its very easy to forget to use correct import path**
- removed async component in favor of `React.lazy`
  - they do the same thing, no need to have two redundant implementations
  - Just an observation, I would look at how you handle the loading of lazy chunks and create some nice loaders to avoid the ugly blank screens between loads
- use pre-defined inventory component dependencies for PF core and icons

## Results

### Before
Build size of the APP using `build:prod` only JS (no css, no .map, no images):
```bash
  
                                                 js/Advisories.js   2.05 KiB       1  [emitted]               Advisories
    js/Advisories~AdvisoryyPage~InventoryPage~Register~Systems.js   47.1 KiB       0  [emitted]
                                              js/AdvisoryyPage.js   46.2 KiB       2  [emitted]                 
                                                        js/App.js   51.8 KiB       3  [emitted]             
                                              js/InventoryPage.js   1.93 KiB       4  [emitted]                   
                                                   js/Register.js  379 bytes       5  [emitted]               
                                                    js/Systems.js    6.4 KiB       6  [emitted]               
                                                     js/vendor.js    4.7 MiB       7  [emitted]        [big]  

``` 
**Around 5 MiB** total size parsed

### After 
Build size of the APP using `build:prod` only JS (no css, no .map, no images):
```
                                                     js/Advisories.js   2.05 KiB       2  [emitted]               
                                                  js/AdvisoryyPage.js   10.2 KiB       3  [emitted]            
                                                            js/App.js   1.35 MiB       4  [emitted]        [big]  
                                                  js/InventoryPage.js   8.74 KiB       5  [emitted]            
                                                       js/Register.js  379 bytes       6  [emitted]               
                                                        js/Systems.js    3.8 KiB       7  [emitted]               
                                   js/defaultVendors~AdvisoryyPage.js   37.9 KiB       8  [emitted]
js/default~Advisories~AdvisoryyPage~InventoryPage~Register~Systems.js   19.4 KiB       0  [emitted]
                                  js/default~AdvisoryyPage~Systems.js   9.54 KiB       1  [emitted]    
```
**1.44 MiB** total size parsed 

There is still one duplication of PF react-table packages, but that comes from frontend components. I will check that one after I am finished with this PR. It should fix itself in one of the future releases of frontend components.

cc @jiridostal 

---

## Discussion

### Comment by @codecov-commenter on June 04, 2020 at 11:47 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/149?src=pr&el=h1) Report
> Merging [#149](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/149?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/8f2cdfb3d8f1855bf469ddc42bd19f9bfe54913f&el=desc) will **increase** coverage by `0.19%`.
> The diff coverage is `71.42%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/149/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/149?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #149      +/-   ##
==========================================
+ Coverage   37.16%   37.36%   +0.19%     
==========================================
  Files          51       50       -1     
  Lines         826      819       -7     
  Branches       83       82       -1     
==========================================
- Hits          307      306       -1     
+ Misses        473      468       -5     
+ Partials       46       45       -1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/149?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [src/App.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/149/diff?src=pr&el=tree#diff-c3JjL0FwcC5qcw==) | `0.00% <ø> (ø)` | |
| [...ionalComponents/AdvisoriesTable/AdvisoriesTable.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/149/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9BZHZpc29yaWVzVGFibGUvQWR2aXNvcmllc1RhYmxlLmpz) | `0.00% <ø> (ø)` | |
| [...omponents/AdvisoriesTable/AdvisoriesTableAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/149/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9BZHZpc29yaWVzVGFibGUvQWR2aXNvcmllc1RhYmxlQXNzZXRzLmpz) | `0.00% <ø> (ø)` | |
| [...ntationalComponents/AdvisoriesTable/TableFooter.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/149/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9BZHZpc29yaWVzVGFibGUvVGFibGVGb290ZXIuanM=) | `0.00% <ø> (ø)` | |
| [...ationalComponents/AdvisoryHeader/AdvisoryHeader.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/149/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9BZHZpc29yeUhlYWRlci9BZHZpc29yeUhlYWRlci5qcw==) | `0.00% <ø> (ø)` | |
| [...sentationalComponents/Filters/PublishDateFilter.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/149/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL1B1Ymxpc2hEYXRlRmlsdGVyLmpz) | `0.00% <ø> (ø)` | |
| [src/SmartComponents/Advisories/Advisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/149/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yaWVzL0Fkdmlzb3JpZXMuanM=) | `0.00% <ø> (ø)` | |
| [...SmartComponents/AffectedSystems/AffectedSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/149/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZmZlY3RlZFN5c3RlbXMvQWZmZWN0ZWRTeXN0ZW1zLmpz) | `0.00% <ø> (ø)` | |
| [src/SmartComponents/Remediation/Remediation.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/149/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9SZW1lZGlhdGlvbi9SZW1lZGlhdGlvbi5qcw==) | `0.00% <ø> (ø)` | |
| [...rc/SmartComponents/Remediation/RemediationModal.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/149/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9SZW1lZGlhdGlvbi9SZW1lZGlhdGlvbk1vZGFsLmpz) | `0.00% <ø> (ø)` | |
| ... and [10 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/149/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/149?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/149?src=pr&el=footer). Last update [8f2cdfb...fca823b](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/149?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @karelhala on June 05, 2020 at 09:32 AM UTC

ping @jiridostal 

### Comment by @jiridostal on June 05, 2020 at 11:05 AM UTC

That is a lot of changes, give me a few moment to check it.

### Comment by @jiridostal on June 05, 2020 at 11:41 AM UTC

Thanks for that, everything seems to work as expected!

### Comment by @jiridostal on June 08, 2020 at 11:20 AM UTC

:tada: This PR is included in version 0.12.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v0.12.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/0.12.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/149*
