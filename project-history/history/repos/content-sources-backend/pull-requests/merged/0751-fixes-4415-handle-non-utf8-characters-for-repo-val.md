---
type: pull_request
number: 751
title: "Fixes 4415: handle non-UTF8 characters for repo validate"
state: merged
author: rverdile
created: 2024-07-18T20:52:19Z
updated: 2024-07-19T11:30:19Z
closed: 2024-07-19T11:06:54Z
merged: 2024-07-19T11:06:54Z
base_branch: main
head_branch: validate-500
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/751
---

# Pull Request #751: Fixes 4415: handle non-UTF8 characters for repo validate

**Author**: @rverdile
**Created**: July 18, 2024 at 08:52 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `validate-500`

## Description

## Summary
Returns 400 instead of 500 (for name and URL) when sending non-UTF8 characters in repository validate request

## Testing steps
1. POST `/api/content-sources/v1/repository_parameters/validate/`

with body:
  [
     {       "url": "\u0000/"   }
  ]

2. Do the same for name
3. Previously, this would return a 500 error. Now it should return a 400.

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on July 18, 2024 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-4415

### Comment by @swadeley on July 19, 2024 at 11:06 AM UTC

Hi

```
In [1]: app.content_sources.rest_client.repositories_api.validate_repository_parameters(
   ...:         [
   ...: { "url": "\u0000/" }
   ...: ]
   ...:     )
HTTP response body: {"errors":[{"status":400,"title":"Error validating repository","detail":"Request parameters contain invalid syntax"}]}
```

Same error with name

Thank you

---

## Reviews

### Review by @xbhouse - Approved on July 18, 2024 at 10:07 PM UTC

ack!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/751*
