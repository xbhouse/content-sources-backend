---
type: pull_request
number: 653
title: "Fixes 4044: fix introspection of large repos"
state: merged
author: jlsherrill
created: 2024-04-29T16:04:09Z
updated: 2024-04-30T08:00:24Z
closed: 2024-04-30T07:33:13Z
merged: 2024-04-30T07:33:13Z
base_branch: main
head_branch: 4044
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/653
---

# Pull Request #653: Fixes 4044: fix introspection of large repos

**Author**: @jlsherrill
**Created**: April 29, 2024 at 04:04 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4044`

## Description

## Summary
Previously introspecting large repos resulted in an error

## Testing steps

Create a repo with URL https://mirror.web-ster.com/fedora/releases/40/Everything/x86_64/os/
Disable snapshotting to speed it up.
it will fail without this PR

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @swadeley on April 29, 2024 at 04:21 PM UTC

Hi, note to self, with current code:
```
Type
    introspect

Error
    error introspecting https://mirror.web-ster.com/fedora/releases/40/Everything/x86_64/os/: failed retrieving existing checksum in rpms: extended protocol limited to 65535 parameters
```

```
Type
    snapshot
Error
    reason: Killed by signal 9.. 
```

### Comment by @jlsherrill on April 29, 2024 at 04:30 PM UTC

https://issues.redhat.com/browse/HMS-4044

### Comment by @swadeley on April 30, 2024 at 07:32 AM UTC

Hi
```

Type
    introspect
Status
    completed
Queued At
    30 Apr 2024 - 08:31:00
Started At
    30 Apr 2024 - 08:31:00
Finished At
    30 Apr 2024 - 08:31:27
```

---

## Reviews

### Review by @jlsherrill - Commented on April 29, 2024 at 07:59 PM UTC

### Review by @xbhouse - Approved on April 29, 2024 at 08:37 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/653*
