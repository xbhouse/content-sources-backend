---
type: pull_request
number: 1106
title: "HMS-5913: add package dates from RHEL lifecycle"
state: merged
author: rverdile
created: 2025-05-07T20:13:48Z
updated: 2025-05-14T02:55:47Z
closed: 2025-05-14T02:55:46Z
merged: 2025-05-14T02:55:46Z
base_branch: main
head_branch: pkg-lifecycle-info
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1106
---

# Pull Request #1106: HMS-5913: add package dates from RHEL lifecycle

**Author**: @rverdile
**Created**: May 07, 2025 at 08:13 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `pkg-lifecycle-info`

## Description

## Summary
Adds package streams info to package_sources. Pulls package stream info from the roadmap API and adds it to package sources if the RPM name matches.
    
Also adds RHEL EoL as the end date for any package source that is not in the roadmap API. Uses the end date for the latest released version of the matching RHEL major version.

## Testing steps
To test package streams:
1. Search for "nodejs" in a rhel9 repo
2. Under package sources, you should see both the node 16 module and the 16.14.0 package stream

To test the rhel lifecycle info:
1. When a RHEL package doesn't exist in the roadmap, the RHEL end date is used instead
2. Search with a codeready repo. For example, "help2man" in https://cdn.redhat.com/content/dist/rhel9/9/aarch64/codeready-builder/os/



---

## Discussion

### Comment by @jlsherrill on May 07, 2025 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-5913

### Comment by @jlsherrill on May 13, 2025 at 01:16 PM UTC

/retest

### Comment by @jlsherrill on May 13, 2025 at 05:00 PM UTC

The copilot suggestions i think are overall pretty good (and likely things i would have missed).   But let me know if you disagree. 

The only other thing, is that the package streams look like this:  `"stream": "16.14.0",`  which isn't great, as its different than the other streams, and i think the version in rhel won't always be 16.14.0.  This is more the fault of the underlying data we have, and talking with sam i don' think its something they want to fix.   I think lets merge it leaving it as it is and see how it looks for other packages.  If we decide to, we can do a split on periods and just use the major version (if other packages show simialr things)

EDIT: well and actually for something like mariadb, we'd need the "x.y" to help differentiate:
```
{
  "stream": "10.11",
  "impl": "dnf_module",
  "name": "mariadb"
}
{
  "stream": "10.5.13",
  "impl": "package",
  "name": "mariadb"
}
```

### Comment by @rverdile on May 13, 2025 at 05:31 PM UTC

forgot about the copilot suggestions, added those

### Comment by @swadeley on May 14, 2025 at 02:15 AM UTC

/retest

### Comment by @swadeley on May 14, 2025 at 02:55 AM UTC

Hi, tests are passing and the specific improvement needs testing in stage.

---

## Reviews

### Review by @copilot-pull-request-reviewer - Commented on May 12, 2025 at 05:32 PM UTC

## Pull Request Overview

This PR adds package stream and RHEL lifecycle information to package sources by integrating with the roadmap API. Key changes include:
- Retrieving and appending package stream details for RPMs via new roadmap API calls.
- Adding fallback logic to include RHEL lifecycle end dates if package stream data is not available.
- Updating tests, mocks, and cache interfaces to accommodate the new lifecycle functionality.

### Reviewed Changes

Copilot reviewed 9 out of 9 changed files in this pull request and generated 4 comments.

<details>
<summary>Show a summary per file</summary>

| File                                            | Description                                                        |
|-------------------------------------------------|--------------------------------------------------------------------|
| pkg/dao/rpms_test.go                            | Added test cases for RHEL repo configurations and lifecycle info.  |
| pkg/dao/rpms.go                                 | Updated search logic to add package sources and lifecycle info.    |
| pkg/clients/roadmap_client/roadmap_client_mock.go | Added mocks for lifecycle API calls.                               |
| pkg/clients/roadmap_client/roadmap.go           | Introduced lifecycle response types and functions.                 |
| pkg/clients/roadmap_client/client.go            | Updated client interface for new lifecycle methods.                |
| pkg/cache/redis.go                              | Added cache functions for RHEL lifecycle data.                     |
| pkg/cache/noop.go                               | Added no-op implementations for RHEL lifecycle caching.            |
| pkg/cache/cache_mock.go                          | Added mocks for the new RHEL lifecycle cache functions.            |
| pkg/cache/cache.go                              | Extended cache interface to include lifecycle methods.             |
</details>






### Review by @jlsherrill - Commented on May 12, 2025 at 08:21 PM UTC

### Review by @jlsherrill - Commented on May 12, 2025 at 08:26 PM UTC

### Review by @rverdile - Commented on May 12, 2025 at 08:28 PM UTC

### Review by @jlsherrill - Approved on May 13, 2025 at 06:02 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1106*
