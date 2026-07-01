---
type: pull_request
number: 42
title: "Remove unused status_text"
state: merged
author: semtexzv
created: 2020-01-03T12:01:25Z
updated: 2021-03-16T09:26:45Z
closed: 2020-01-03T13:55:19Z
merged: 2020-01-03T13:55:19Z
base_branch: master
head_branch: status_text
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/42
---

# Pull Request #42: Remove unused status_text

**Author**: @semtexzv
**Created**: January 03, 2020 at 12:01 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `status_text`

## Description

This field in the database is currently not used, and its use in the vulnerabilities application is different than what the name implies. In vulnerabilities, the `status_text` is used to store the `Justification` from the UI, and is not a cache for values stored in status_id, as the name would imply.

Patchman does not have this functionality, and there are no plans for adding it.

---

## Reviews

### Review by @josef-hak - Approved on January 03, 2020 at 01:55 PM UTC

Ok, thanks.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/42*
