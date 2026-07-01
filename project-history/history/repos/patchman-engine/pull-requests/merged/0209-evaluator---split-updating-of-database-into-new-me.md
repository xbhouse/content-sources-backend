---
type: pull_request
number: 209
title: "Evaluator - split updating of database into new method"
state: merged
author: semtexzv
created: 2020-04-02T10:59:27Z
updated: 2021-03-16T09:28:00Z
closed: 2020-04-02T11:19:18Z
merged: 2020-04-02T11:19:18Z
base_branch: master
head_branch: split-eval
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/209
---

# Pull Request #209: Evaluator - split updating of database into new method

**Author**: @semtexzv
**Created**: April 02, 2020 at 10:59 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `split-eval`

## Description

Move storing of `system_advisories` and `advisory_account_data` into new method. 

Should provide us more granularity,  the label `advisories-processing` will now cover only the processing of reading & processing of the data & label `advisories-store` will cover actual updating of `system_advisories` & `advisory_account_data` 

---

## Reviews

### Review by @josef-hak - Approved on April 02, 2020 at 11:18 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/209*
