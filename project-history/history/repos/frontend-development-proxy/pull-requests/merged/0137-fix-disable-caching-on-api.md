---
type: pull_request
number: 137
title: "fix: disable caching on /api/*"
state: merged
author: xbhouse
created: 2026-05-15T17:29:33Z
updated: 2026-05-18T07:27:55Z
closed: 2026-05-18T07:27:54Z
merged: 2026-05-18T07:27:54Z
base_branch: main
head_branch: disable-api-cache
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/137
---

# Pull Request #137: fix: disable caching on /api/*

**Author**: @xbhouse
**Created**: May 15, 2026 at 05:29 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `disable-api-cache`

## Description

### Description
<!-- Summary of proposed changes: what and why. -->

This disables caching for `/api/*` endpoints, since those will likely return different output based on the token. 
Without this, we are seeing cases where parallel requests with the same URL but different bearer tokens reuse responses and request IDs. 

[RHCLOUD-47766](https://issues.redhat.com/browse/RHCLOUD-47766)

---

### How to test locally
<!-- How can a reviewer exercise this? Special setup, env vars, seed data? -->
<!-- Bug fixes: how to reproduce the original bug and confirm the fix. -->

1. Run any frontend with the proxy without these changes (`fec dev-proxy`) and make 2 parallel curl requests to the same URL with different Authorization headers. This script could help:

```
URL="https://stage.foo.redhat.com:1337/api/rbac/v2/workspaces/?limit=1000&offset=0"

for i in $(seq 1 20); do   
  ( curl --max-time 10 -skg -D /tmp/h1.txt -o /tmp/b1.json -H "Authorization: $USER1" "$URL" ) &   
  ( curl --max-time 10 -skg -D /tmp/h2.txt -o /tmp/b2.json -H "Authorization: $USER2" "$URL" ) &   
  wait   
  id1=$(grep -i '^x-rh-insights-request-id:' /tmp/h1.txt | tr -d '\r');   
  id2=$(grep -i '^x-rh-insights-request-id:' /tmp/h2.txt | tr -d '\r');   
  [ "$id1" = "$id2" ] && echo "iter=$i same request id: $id1" && break 
done
```

You should see requests from different users reuse the same request ID.

2. Build a new image with the changes in this PR, run it, and restart the frontend with `fec static`.

3. Run the script above. There shouldn't be any shared request IDs.

---

### Anything reviewers should know?
<!-- Trade-offs, performance considerations, pre/post merge actions required. -->

This doesn't disable caching for frontend assets and performance doesn't seem to be impacted locally. 
Another option might be to use a cache key that includes the Authorization header (and maybe the cookie) if that's preferred. 

---

### Checklist
- [ ] Tests: new/updated tests cover the change
- [ ] API: spec updated if endpoints changed (or N/A)
- [ ] Migrations: backwards compatible if schema changed (or N/A)
- [ ] Dependencies: no known impact to dependent services
- [ ] Security: reviewed against [secure coding checklist](https://github.com/RedHatInsights/secure-coding-checklist) (or N/A)

### AI disclosure
<!-- If AI tools contributed, note them. E.g.: Assisted by: Claude Code -->

Assisted by: Cursor

[RHCLOUD-47766]: https://redhat.atlassian.net/browse/RHCLOUD-47766?atlOrigin=eyJpIjoiNWRkNTljNzYxNjVmNDY3MDlhMDU5Y2ZhYzA5YTRkZjUiLCJwIjoiZ2l0aHViLWNvbS1KU1cifQ

---

## Reviews

### Review by @Hyperkid123 - Approved on May 18, 2026 at 07:27 AM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/137*
