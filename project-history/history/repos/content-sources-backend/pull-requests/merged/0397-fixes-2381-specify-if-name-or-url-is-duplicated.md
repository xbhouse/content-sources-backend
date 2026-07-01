---
type: pull_request
number: 397
title: "Fixes 2381: specify if name or URL is duplicated"
state: merged
author: rverdile
created: 2023-09-21T15:22:47Z
updated: 2023-09-29T15:00:32Z
closed: 2023-09-29T14:53:29Z
merged: 2023-09-29T14:53:29Z
base_branch: main
head_branch: dupe-field
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/397
---

# Pull Request #397: Fixes 2381: specify if name or URL is duplicated

**Author**: @rverdile
**Created**: September 21, 2023 at 03:22 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `dupe-field`

## Description

## Summary
When creating a repository, if a name or URL is already being used, the error message should say which field is duplicated. In a previous change, we renamed constraints that were being used to determine which field is duplicated. I've updated the function with the new constraint names.

## Testing steps
1. Create a repository
2. Create a repository again with the same name.
3. Check the error message. It should tell you "name" is the duplicated field.
4. Create a repository again with the a different name, but same URL.
5. Check the error message. It should tell you "URL" is the duplicated field.


---

## Discussion

### Comment by @jlsherrill on September 21, 2023 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-2381

### Comment by @jlsherrill on September 25, 2023 at 07:01 PM UTC

This needs a rebase still

### Comment by @swadeley on September 28, 2023 at 10:36 AM UTC

Hi

tests OK

`HTTP response body: {"errors":[{"status":400,"title":"Error creating repository","detail":"Repository with this URL already belongs to organization"}]}`


`HTTP response body: {"errors":[{"status":400,"title":"Error creating repository","detail":"Repository with this name already belongs to organization"}]}
`

lgtm

### Comment by @rverdile on September 28, 2023 at 02:55 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Commented on September 21, 2023 at 03:24 PM UTC

### Review by @jlsherrill - Commented on September 25, 2023 at 07:09 PM UTC

### Review by @rverdile - Commented on September 26, 2023 at 06:24 PM UTC

### Review by @jlsherrill - Commented on September 28, 2023 at 11:48 AM UTC

### Review by @jlsherrill - Commented on September 28, 2023 at 11:49 AM UTC

### Review by @jlsherrill - Commented on September 28, 2023 at 11:51 AM UTC

### Review by @rverdile - Commented on September 28, 2023 at 01:32 PM UTC

### Review by @jlsherrill - Commented on September 28, 2023 at 01:47 PM UTC

### Review by @jlsherrill - Commented on September 28, 2023 at 01:47 PM UTC

### Review by @jlsherrill - Commented on September 28, 2023 at 02:23 PM UTC

### Review by @rverdile - Commented on September 28, 2023 at 02:28 PM UTC

### Review by @jlsherrill - Approved on September 28, 2023 at 02:28 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/397*
