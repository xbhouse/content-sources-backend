---
type: pull_request
number: 717
title: "Fixes 2246: fix conflict on guard creation"
state: merged
author: jlsherrill
created: 2024-06-24T18:56:23Z
updated: 2024-06-27T14:28:16Z
closed: 2024-06-27T14:28:12Z
merged: 2024-06-27T14:28:12Z
base_branch: main
head_branch: 2246
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/717
---

# Pull Request #717: Fixes 2246: fix conflict on guard creation

**Author**: @jlsherrill
**Created**: June 24, 2024 at 06:56 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `2246`

## Description

## Summary

After reviewing the domain creation logic to handle race conditions on create, i determined the current method was best, but could use some simplification by use of errorWithResponseBody
while testing, however, i discovered that the content_Guard creation did not properly handle race conditions.

## Testing steps

1.  turn on content guard creation by setting clients -> pulp -> 'custom_repo_content_guards' to true in config.yaml
2. run the server
3. run:
```
curl -H "`./scripts/header.sh $RANDOM $RANDOM`"  http://localhost:8000/api/content-sources/v1/repositories/bulk_create/  -H "content-type: application/json"  -X POST -d '[{"name":"abc1", "url":"https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/", "snapshot": true},{"name":"abc2", "url":"https://jlsherrill.fedorapeople.org/fake-repos/signed/", "snapshot": true}, {"name":"abc3", "url":"https://jlsherrill.fedorapeople.org/fake-repos/really-empty/", "snapshot": true}, {"name":"abc4", "url":"https://jlsherrill.fedorapeople.org/fake-repos/empty/", "snapshot": true}]'
```

This will try to create 3 repos with snapshotting in a random org.  Monitor the backend logs and verify no error occurs from the tasks

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on June 24, 2024 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-2246

### Comment by @jlsherrill on June 25, 2024 at 07:56 PM UTC

/retest

---

## Reviews

### Review by @xbhouse - Approved on June 25, 2024 at 06:45 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/717*
