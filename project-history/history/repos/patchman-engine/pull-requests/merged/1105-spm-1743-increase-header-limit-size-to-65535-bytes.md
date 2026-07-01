---
type: pull_request
number: 1105
title: "SPM-1743 increase header limit size to 65535 bytes"
state: merged
author: tkasparek
created: 2022-09-15T11:54:51Z
updated: 2022-09-15T13:54:12Z
closed: 2022-09-15T13:54:12Z
merged: 2022-09-15T13:54:12Z
base_branch: master
head_branch: header_size
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1105
---

# Pull Request #1105: SPM-1743 increase header limit size to 65535 bytes

**Author**: @tkasparek
**Created**: September 15, 2022 at 11:54 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `header_size`

## Description

*No description provided*

---

## Discussion

### Comment by @jira-linking on September 15, 2022 at 11:54 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1743


### Comment by @codecov-commenter on September 15, 2022 at 12:03 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1105?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **62.87**% // Head: **62.95**% // Increases project coverage by **`+0.08%`** :tada:
> Coverage data is based on head [(`323f85e`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1105?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`146efec`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/146efeca7ce7bae151b4bc75c056227eb7c7fde5?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 80.43% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1105      +/-   ##
==========================================
+ Coverage   62.87%   62.95%   +0.08%     
==========================================
  Files          97       97              
  Lines        5535     5548      +13     
==========================================
+ Hits         3480     3493      +13     
+ Misses       1611     1606       -5     
- Partials      444      449       +5     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.95% <80.43%> (+0.08%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1105?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1105/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `9.96% <0.00%> (ø)` | |
| [manager/controllers/baseline\_create.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1105/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9jcmVhdGUuZ28=) | `75.80% <ø> (ø)` | |
| [manager/controllers/baseline\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1105/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9kZXRhaWwuZ28=) | `88.09% <ø> (ø)` | |
| [manager/kafka/kafka.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1105/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9rYWZrYS9rYWZrYS5nbw==) | `68.88% <ø> (ø)` | |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1105/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.41% <76.19%> (+0.25%)` | :arrow_up: |
| [base/utils/gin.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1105/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9naW4uZ28=) | `21.05% <100.00%> (ø)` | |
| [evaluator/evaluate\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1105/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Fkdmlzb3JpZXMuZ28=) | `76.69% <100.00%> (ø)` | |
| [evaluator/evaluate\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1105/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | `78.26% <100.00%> (ø)` | |
| [evaluator/notifications.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1105/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL25vdGlmaWNhdGlvbnMuZ28=) | `70.90% <100.00%> (ø)` | |
| [manager/controllers/baseline\_delete.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1105/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9kZWxldGUuZ28=) | `71.87% <100.00%> (ø)` | |
| ... and [5 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1105/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1105?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on September 15, 2022 at 12:49 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1105*
