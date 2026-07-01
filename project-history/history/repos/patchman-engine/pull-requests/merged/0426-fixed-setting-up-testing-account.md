---
type: pull_request
number: 426
title: "fixed setting up testing account"
state: merged
author: josef-hak
created: 2020-11-20T14:31:00Z
updated: 2020-11-27T15:22:19Z
closed: 2020-11-23T12:07:36Z
merged: 2020-11-23T12:07:36Z
base_branch: master
head_branch: fix-test-account
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/426
---

# Pull Request #426: fixed setting up testing account

**Author**: @josef-hak
**Created**: November 20, 2020 at 02:31 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-test-account`

## Description

- after updating testing accounts in test_data.sql it's needed to update account setting into the env variable ACCOUNT_NUMER in `./dev/scripts/env.sh`

---

## Reviews

### Review by @semtexzv - Approved on November 23, 2020 at 12:07 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/426*
