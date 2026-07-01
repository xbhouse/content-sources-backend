---
type: pull_request
number: 968
title: "SPM-1550: add org_id populator cji"
state: merged
author: psegedy
created: 2022-06-02T11:35:26Z
updated: 2022-06-03T08:25:47Z
closed: 2022-06-03T08:25:47Z
merged: 2022-06-03T08:25:47Z
base_branch: master
head_branch: org_id_translation
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/968
---

# Pull Request #968: SPM-1550: add org_id populator cji

**Author**: @psegedy
**Created**: June 02, 2022 at 11:35 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `org_id_translation`

## Description

Add org_id populator to translate account numbers to org_ids - https://github.com/RedHatInsights/tenant-utils/blob/main/cmd/org-id-column-populator/README.md

## Secure Coding Practices Checklist GitHub Link
- https://github.com/RedHatInsights/secure-coding-checklist

## Secure Coding Checklist
- [x] Input Validation
- [x] Output Encoding
- [x] Authentication and Password Management
- [x] Session Management
- [x] Access Control
- [x] Cryptographic Practices
- [x] Error Handling and Logging
- [x] Data Protection
- [x] Communication Security
- [x] System Configuration
- [x] Database Security
- [x] File Management
- [x] Memory Management
- [x] General Coding Practices


---

## Discussion

### Comment by @codecov-commenter on June 02, 2022 at 11:42 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/968?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#968](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/968?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (25062f5) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/486c1e1ed53d040254f49b1754cbf3153be002ee?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (486c1e1) will **decrease** coverage by `0.38%`.
> The diff coverage is `45.62%`.

```diff
@@            Coverage Diff             @@
##           master     #968      +/-   ##
==========================================
- Coverage   61.30%   60.91%   -0.39%     
==========================================
  Files          92       94       +2     
  Lines        4975     5102     +127     
==========================================
+ Hits         3050     3108      +58     
- Misses       1518     1585      +67     
- Partials      407      409       +2     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.91% <45.62%> (-0.39%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/968?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/968/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `50.00% <0.00%> (+22.41%)` | :arrow_up: |
| [base/mqueue/payload\_tracker\_event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/968/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvcGF5bG9hZF90cmFja2VyX2V2ZW50Lmdv) | `0.00% <0.00%> (ø)` | |
| [base/mqueue/platform\_event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/968/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvcGxhdGZvcm1fZXZlbnQuZ28=) | `9.52% <9.52%> (ø)` | |
| [base/utils/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/968/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | `31.97% <25.00%> (-0.20%)` | :arrow_down: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/968/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `68.12% <62.50%> (-0.60%)` | :arrow_down: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/968/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `81.45% <97.82%> (+2.64%)` | :arrow_up: |
| [evaluator/notifications.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/968/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL25vdGlmaWNhdGlvbnMuZ28=) | `69.69% <100.00%> (ø)` | |
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/968/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `68.75% <100.00%> (+2.08%)` | :arrow_up: |
| [manager/kafka/kafka.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/968/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9rYWZrYS9rYWZrYS5nbw==) | `68.88% <100.00%> (ø)` | |
| [vmaas\_sync/send\_messages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/968/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9zZW5kX21lc3NhZ2VzLmdv) | `34.61% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/968?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/968?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [2ff439f...25062f5](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/968?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on June 02, 2022 at 02:02 PM UTC

tried in ephemeral with `clowdjobinvocation/populate-org-id-column-1`, it seems to be working
before
```
patchman=# select * from rh_account;
 id |  name   | org_id
----+---------+--------
  1 | 6089719 |
(1 row)
```
after
```
patchman=# select * from rh_account;
 id |  name   |  org_id
----+---------+----------
  1 | 6089719 | 11789772
(1 row)
```

### Comment by @MichaelMraka on June 02, 2022 at 02:20 PM UTC

Is it kind of one-time thing or should it stay permanently in patch?

### Comment by @psegedy on June 03, 2022 at 08:25 AM UTC

Just a one-time thing, we are already storing (and updating) org_ids from API requests and kafka messages. This is meant for updating all rows in DB at once. After that we can drop this Job because we are already storing org_id in db.

---

## Reviews

### Review by @MichaelMraka - Approved on June 02, 2022 at 02:20 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/968*
