---
type: pull_request
number: 98
title: "Added re-evaluation using new instance of evaluator"
state: merged
author: josef-hak
created: 2020-02-03T15:17:01Z
updated: 2020-02-06T16:00:11Z
closed: 2020-02-06T13:03:46Z
merged: 2020-02-06T13:03:45Z
base_branch: master
head_branch: re-eval
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/98
---

# Pull Request #98: Added re-evaluation using new instance of evaluator

**Author**: @josef-hak
**Created**: February 03, 2020 at 03:17 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `re-eval`

## Description

- split evaluator into two components:
  - `evaluator_upload` - will evaluate system after new archive upload. Component `listener` will send kafka message for that to `patchman.evaluator.recalc` kafka topic.
  - `evaluator_recalc` - will evaluate systems after vmaas update event. Component `vmaas_sync` will send kafka messages for that (for all systems) to `patchman.evaluator.recalc` kafka topic.
- speed up `docker-compose up` development setup using var `GORUN` and local folder mounting
- updated openshift configs
- added script to `docker-compose.yml` to feed db with testing data 
- updated `README.md` with description how to test whole app

---

## Discussion

### Comment by @codecov-io on February 04, 2020 at 01:02 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/98?src=pr&el=h1) Report
> :exclamation: No coverage uploaded for pull request base (`master@ae92cd4`). [Click here to learn what that means](https://docs.codecov.io/docs/error-reference#section-missing-base-commit).
> The diff coverage is `20.93%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/98/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/98?src=pr&el=tree)

```diff
@@            Coverage Diff            @@
##             master      #98   +/-   ##
=========================================
  Coverage          ?   54.36%           
=========================================
  Files             ?       35           
  Lines             ?     1398           
  Branches          ?        0           
=========================================
  Hits              ?      760           
  Misses            ?      547           
  Partials          ?       91
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/98?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [evaluator/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/98/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL21ldHJpY3MuZ28=) | `0% <0%> (ø)` | |
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/98/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `24.32% <0%> (ø)` | |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/98/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `0% <0%> (ø)` | |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/98/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `39.42% <0%> (ø)` | |
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/98/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `0% <0%> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/98/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `70.49% <69.23%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/98?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/98?src=pr&el=footer). Last update [ae92cd4...5ed6c15](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/98?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @beav on February 04, 2020 at 06:26 PM UTC

Can you give more detail on what this change does? I'm not sure what the re-evaluator's function is

### Comment by @josef-hak on February 04, 2020 at 08:27 PM UTC

@beav ok, I added description and improved docker-compose setup to easier test it

---

## Reviews

### Review by @semtexzv - Approved on February 06, 2020 at 01:02 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/98*
