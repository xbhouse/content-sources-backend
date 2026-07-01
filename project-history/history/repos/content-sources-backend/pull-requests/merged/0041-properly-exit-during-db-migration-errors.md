---
type: pull_request
number: 41
title: "Properly exit during db migration errors"
state: merged
author: jlsherrill
created: 2022-06-22T17:57:19Z
updated: 2022-06-27T16:42:13Z
closed: 2022-06-27T16:23:49Z
merged: 2022-06-27T16:23:49Z
base_branch: main
head_branch: migrate_error
labels: []
url: https://github.com/content-services/content-sources-backend/pull/41
---

# Pull Request #41: Properly exit during db migration errors

**Author**: @jlsherrill
**Created**: June 22, 2022 at 05:57 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `migrate_error`

## Description

*No description provided*

---

## Discussion

### Comment by @jlsherrill on June 23, 2022 at 04:32 PM UTC

From the zerolog docs:  
`
It is very important to note that when using the zerolog chaining API, as shown above (log.Info().Msg("hello world"), the chain must have either the Msg or Msgf method call. If you forget to add either of these, the log will not occur and there is no compile time error to alert you of this.`

---

## Reviews

### Review by @rverdile - Approved on June 27, 2022 at 01:52 PM UTC

Looks good and didn't find any other places missing "Msg"

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/41*
