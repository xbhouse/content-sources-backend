---
type: pull_request
number: 1585
title: "fix(deps): update module github.com/prometheus/client_golang to v1.21.0"
state: merged
author: red-hat-konflux
created: 2025-02-23T12:56:53Z
updated: 2026-04-06T13:12:12Z
closed: 2025-02-24T12:32:08Z
merged: 2025-02-24T12:32:08Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-prometheus-client_golang-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1585
---

# Pull Request #1585: fix(deps): update module github.com/prometheus/client_golang to v1.21.0

**Author**: @red-hat-konflux
**Created**: February 23, 2025 at 12:56 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-prometheus-client_golang-1.x`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [github.com/prometheus/client_golang](https://redirect.github.com/prometheus/client_golang) | require | minor | `v1.20.5` -> `v1.21.0` |

---

### Release Notes

<details>
<summary>prometheus/client_golang (github.com/prometheus/client_golang)</summary>

### [`v1.21.0`](https://redirect.github.com/prometheus/client_golang/releases/tag/v1.21.0): / 2025-02-19

[Compare Source](https://redirect.github.com/prometheus/client_golang/compare/v1.20.5...v1.21.0)

:warning: This release contains potential breaking change if you upgrade `github.com/prometheus/common` to 0.62+ together with client_golang (and depend on the strict, legacy validation for the label names). New common version [changes `model.NameValidationScheme` global variable](https://redirect.github.com/prometheus/common/pull/724), which relaxes the validation of label names and metric name, allowing all UTF-8 characters. Typically, this should not break any user, unless your test or usage expects strict certain names to panic/fail on client_golang metric registration, gathering or scrape. In case of problems change `model.NameValidationScheme` to old `model.LegacyValidation` value in your project `init` function. :warning:

-   \[BUGFIX] gocollector: Fix help message for runtime/metric metrics. [#&#8203;1583](https://redirect.github.com/prometheus/client_golang/issues/1583)
-   \[BUGFIX] prometheus: Fix `Desc.String()` method for no labels case. [#&#8203;1687](https://redirect.github.com/prometheus/client_golang/issues/1687)
-   \[PERF] prometheus: Optimize popular `prometheus.BuildFQName` function; now up to 30% faster. [#&#8203;1665](https://redirect.github.com/prometheus/client_golang/issues/1665)
-   \[PERF] prometheus: Optimize `Inc`, `Add` and `Observe` cumulative metrics; now up to 50% faster under high concurrent contention. [#&#8203;1661](https://redirect.github.com/prometheus/client_golang/issues/1661)
-   \[CHANGE] Upgrade prometheus/common to 0.62.0 which changes `model.NameValidationScheme` global variable. [#&#8203;1712](https://redirect.github.com/prometheus/client_golang/issues/1712)
-   \[CHANGE] Add support for Go 1.23. [#&#8203;1602](https://redirect.github.com/prometheus/client_golang/issues/1602)
-   \[FEATURE] process_collector: Add support for Darwin systems. [#&#8203;1600](https://redirect.github.com/prometheus/client_golang/issues/1600) [#&#8203;1616](https://redirect.github.com/prometheus/client_golang/issues/1616) [#&#8203;1625](https://redirect.github.com/prometheus/client_golang/issues/1625) [#&#8203;1675](https://redirect.github.com/prometheus/client_golang/issues/1675) [#&#8203;1715](https://redirect.github.com/prometheus/client_golang/issues/1715)
-   \[FEATURE] api: Add ability to invoke `CloseIdleConnections` on api.Client using `api.Client.(CloseIdler).CloseIdleConnections()` casting. [#&#8203;1513](https://redirect.github.com/prometheus/client_golang/issues/1513)
-   \[FEATURE] promhttp: Add `promhttp.HandlerOpts.EnableOpenMetricsTextCreatedSamples` option to create OpenMetrics \_created lines. Not recommended unless you want to use opt-in Created Timestamp feature. Community works on OpenMetrics 2.0 format that should make those lines obsolete (they increase cardinality significantly). [#&#8203;1408](https://redirect.github.com/prometheus/client_golang/issues/1408)
-   \[FEATURE] prometheus: Add `NewConstNativeHistogram` function. [#&#8203;1654](https://redirect.github.com/prometheus/client_golang/issues/1654)

<details>
<summary> All commits </summary>
* Merge release-1.20 to main by @&#8203;bwplotka in https://github.com/prometheus/client_golang/pull/1582
* gocollector: Tiny fix for help message with runtime/metrics source. by @&#8203;bwplotka in https://github.com/prometheus/client_golang/pull/1583
* ci: bump dagger to the latest version by @&#8203;marcosnils in https://github.com/prometheus/client_golang/pull/1588
* Merge release-1.20 back to main by @&#8203;ArthurSens in https://github.com/prometheus/client_golang/pull/1593
* Update linting by @&#8203;SuperQ in https://github.com/prometheus/client_golang/pull/1603
* Update supported Go versions by @&#8203;SuperQ in https://github.com/prometheus/client_golang/pull/1602
* build(deps): bump golang.org/x/sys from 0.22.0 to 0.24.0 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1611
* build(deps): bump github.com/prometheus/common from 0.55.0 to 0.57.0 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1612
* changed the name of all variables with min/max name by @&#8203;parthlaw in https://github.com/prometheus/client_golang/pull/1606
* Update Dagger and build. by @&#8203;SuperQ in https://github.com/prometheus/client_golang/pull/1610
* build(deps): bump github/codeql-action from 3.25.15 to 3.26.6 in the github-actions group across 1 directory by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1614
* examples: Improved GoCollector example. by @&#8203;bwplotka in https://github.com/prometheus/client_golang/pull/1589
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1615
* process_collector: fill in most statistics on macOS by @&#8203;mharbison72 in https://github.com/prometheus/client_golang/pull/1600
* :zap: http client defer CloseIdleConnections by @&#8203;cuisongliu in https://github.com/prometheus/client_golang/pull/1513
* Set allow-utf-8 in Format during tests to avoid escaping. by @&#8203;ywwg in https://github.com/prometheus/client_golang/pull/1618
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1622
* Merge Release 1.20 back to main by @&#8203;ArthurSens in https://github.com/prometheus/client_golang/pull/1627
* examples: Add custom labels example by @&#8203;ying-jeanne in https://github.com/prometheus/client_golang/pull/1626
* Refactor default runtime metrics tests for Go collector so that default runtime metric set autogenerates by @&#8203;vesari in https://github.com/prometheus/client_golang/pull/1631
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1628
* process_xxx_memory statistics for macOS (cgo) by @&#8203;mharbison72 in https://github.com/prometheus/client_golang/pull/1616
* build(deps): bump github.com/klauspost/compress from 1.17.9 to 1.17.10 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1633
* build(deps): bump golang.org/x/sys from 0.24.0 to 0.25.0 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1632
* process_collector: Add Platform-Specific Describe for processCollector by @&#8203;ying-jeanne in https://github.com/prometheus/client_golang/pull/1625
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1635
* build(deps): bump the github-actions group with 4 updates by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1634
* Optionally print OM created lines by @&#8203;ArthurSens in https://github.com/prometheus/client_golang/pull/1408
* process_collector: merge wasip1 and js into a single implementation by @&#8203;ying-jeanne in https://github.com/prometheus/client_golang/pull/1644
* Merge release 1.20 to main by @&#8203;bwplotka in https://github.com/prometheus/client_golang/pull/1647
* Add Arianna as maintainer 💪 by @&#8203;ArthurSens in https://github.com/prometheus/client_golang/pull/1651
* test add headers round tripper by @&#8203;Manask322 in https://github.com/prometheus/client_golang/pull/1657
* build(deps): bump github.com/klauspost/compress from 1.17.10 to 1.17.11 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1668
* build(deps): bump golang.org/x/sys from 0.25.0 to 0.26.0 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1669
* build(deps): bump github.com/prometheus/common from 0.59.1 to 0.60.1 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1667
* build(deps): bump google.golang.org/protobuf from 1.34.2 to 1.35.1 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1670
* Optimize BuildFQName function by @&#8203;jkroepke in https://github.com/prometheus/client_golang/pull/1665
* fix: use injected now() instead of time.Now() in summary methods by @&#8203;imorph in https://github.com/prometheus/client_golang/pull/1672
* process_collector: avoid a compiler warning on macOS (fixes #&#8203;1660) by @&#8203;mharbison72 in https://github.com/prometheus/client_golang/pull/1675
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1674
* build(deps): bump the github-actions group across 1 directory with 3 updates by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1678
* [chore]: enable perfsprint linter by @&#8203;mmorel-35 in https://github.com/prometheus/client_golang/pull/1676
* Duplicate of #&#8203;1662 by @&#8203;imorph in https://github.com/prometheus/client_golang/pull/1673
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1679
* chore: enable usestdlibvars linter by @&#8203;mmorel-35 in https://github.com/prometheus/client_golang/pull/1680
* Add: exponential backoff for CAS operations on floats by @&#8203;imorph in https://github.com/prometheus/client_golang/pull/1661
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1683
* [1617] Add ConstnativeHistogram by @&#8203;shivanthzen in https://github.com/prometheus/client_golang/pull/1654
* fix: replace fmt.Errorf with errors.New by @&#8203;kakkoyun in https://github.com/prometheus/client_golang/pull/1689
* Add codeowners by @&#8203;kakkoyun in https://github.com/prometheus/client_golang/pull/1688
* fix: add very small delay between observations in `TestHistogramAtomicObserve` by @&#8203;imorph in https://github.com/prometheus/client_golang/pull/1691
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1692
* Fix: handle nil variableLabels in Desc.String() method and add tests for nil label values by @&#8203;kakkoyun in https://github.com/prometheus/client_golang/pull/1687
* examples: Follow best practices and established naming conventions by @&#8203;lilic in https://github.com/prometheus/client_golang/pull/1650
* setup OSSF Scorecard workflow by @&#8203;mmorel-35 in https://github.com/prometheus/client_golang/pull/1432
* build(deps): bump google.golang.org/protobuf from 1.35.1 to 1.35.2 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1697
* build(deps): bump golang.org/x/sys from 0.26.0 to 0.27.0 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1696
* build(deps): bump the github-actions group with 5 updates by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1695
* update links to openmetrics to reference the v1.0.0 release by @&#8203;dashpole in https://github.com/prometheus/client_golang/pull/1699
* build(deps): bump google.golang.org/protobuf from 1.35.2 to 1.36.1 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1706
* build(deps): bump golang.org/x/sys from 0.27.0 to 0.28.0 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1705
* build(deps): bump the github-actions group with 5 updates by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1707
* build(deps): bump github.com/prometheus/common from 0.60.1 to 0.61.0 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1704
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1703
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1708
* Upgrade to prometheus/common 0.62.0 with breaking change by @&#8203;bwplotka in https://github.com/prometheus/client_golang/pull/1712
* build(deps): bump golang.org/x/net from 0.26.0 to 0.33.0 in /tutorials/whatsup by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1713
* docs: Add RELEASE.md for the release process by @&#8203;kakkoyun in https://github.com/prometheus/client_golang/pull/1690
* tutorials/whatsup: Updated deps by @&#8203;bwplotka in https://github.com/prometheus/client_golang/pull/1716
* process collector: Fixed pedantic registry failures on darwin with cgo. by @&#8203;bwplotka in https://github.com/prometheus/client_golang/pull/1715
* Revert "ci: daggerize test and lint pipelines (#&#8203;1534)" by @&#8203;bwplotka in https://github.com/prometheus/client_golang/pull/1717
* Cut 1.21.0-rc.0 by @&#8203;bwplotka in https://github.com/prometheus/client_golang/pull/1718
* Cut 1.21 by @&#8203;bwplotka in https://github.com/prometheus/client_golang/pull/1737
</details>

#### New Contributors
* @&#8203;parthlaw made their first contributi[https://github.com/prometheus/client_golang/pull/1606](https://redirect.github.com/prometheus/client_golang/pull/1606)l/1606
* @&#8203;mharbison72 made their first contributi[https://github.com/prometheus/client_golang/pull/1600](https://redirect.github.com/prometheus/client_golang/pull/1600)l/1600
* @&#8203;cuisongliu made their first contributi[https://github.com/prometheus/client_golang/pull/1513](https://redirect.github.com/prometheus/client_golang/pull/1513)l/1513
* @&#8203;ying-jeanne made their first contributi[https://github.com/prometheus/client_golang/pull/1626](https://redirect.github.com/prometheus/client_golang/pull/1626)l/1626
* @&#8203;Manask322 made their first contributi[https://github.com/prometheus/client_golang/pull/1657](https://redirect.github.com/prometheus/client_golang/pull/1657)l/1657
* @&#8203;jkroepke made their first contributi[https://github.com/prometheus/client_golang/pull/1665](https://redirect.github.com/prometheus/client_golang/pull/1665)l/1665
* @&#8203;imorph made their first contributi[https://github.com/prometheus/client_golang/pull/1672](https://redirect.github.com/prometheus/client_golang/pull/1672)l/1672
* @&#8203;mmorel-35 made their first contributi[https://github.com/prometheus/client_golang/pull/1676](https://redirect.github.com/prometheus/client_golang/pull/1676)l/1676
* @&#8203;shivanthzen made their first contributi[https://github.com/prometheus/client_golang/pull/1654](https://redirect.github.com/prometheus/client_golang/pull/1654)l/1654
* @&#8203;dashpole made their first contributi[https://github.com/prometheus/client_golang/pull/1699](https://redirect.github.com/prometheus/client_golang/pull/1699)l/1699

**Full Changelog**: https://github.com/prometheus/client_golang/compare/v1.20.5...v1.21.0

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

### Comment by @red-hat-konflux on February 23, 2025 at 12:56 PM UTC

### ℹ Artifact update notice

##### File name: go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 1 additional dependency was updated


Details:


| **Package**                    | **Change**             |
| :----------------------------- | :--------------------- |
| `github.com/prometheus/common` | `v0.61.0` -> `v0.62.0` |

### Comment by @jira-linking on February 23, 2025 at 12:56 PM UTC

Commits missing Jira IDs:
82dedc10b7142c864b8de52a9973781a5ee33e83


### Comment by @codecov-commenter on February 23, 2025 at 01:01 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1585?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 57.97%. Comparing base ([`a717fed`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/a717fed5ead64993492c49c5847fa2da1b13b848?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`82dedc1`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/82dedc10b7142c864b8de52a9973781a5ee33e83?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 994 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1585   +/-   ##
=======================================
  Coverage   57.97%   57.97%           
=======================================
  Files         143      143           
  Lines       11170    11170           
=======================================
  Hits         6476     6476           
  Misses       4112     4112           
  Partials      582      582           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1585/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1585/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `57.97% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1585?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1585*
