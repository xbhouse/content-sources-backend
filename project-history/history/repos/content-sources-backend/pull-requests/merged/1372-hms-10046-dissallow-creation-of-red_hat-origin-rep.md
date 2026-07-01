---
type: pull_request
number: 1372
title: "HMS-10046: dissallow creation of red_hat origin repos"
state: merged
author: jlsherrill
created: 2026-01-21T21:29:33Z
updated: 2026-01-27T19:40:52Z
closed: 2026-01-27T19:40:52Z
merged: 2026-01-27T19:40:52Z
base_branch: main
head_branch: origin_red_hat
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1372
---

# Pull Request #1372: HMS-10046: dissallow creation of red_hat origin repos

**Author**: @jlsherrill
**Created**: January 21, 2026 at 09:29 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `origin_red_hat`

## Description

## Summary

Previously the bulk create endpoint allowed users to create repos with origin = red_hat.  This closes that and adds a job to delete any existing ones.

## Testing steps

1.  On main, run this curl command (You can use your own rh-identity if you want to see it in the frontend):

```
curl -X POST --location "http://localhost:8000/api/content-sources/v1.0/repositories/bulk_create/" \
    -H "x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==" \
    -H "Content-Type: application/json" \
    -d '[{
          "name": "test8",
          "gpg_key": "",
          "snapshot": true,
          "origin": "red_hat",
          "url": "https://cdn.redhat.com/content/dist/layered/rhel10/aarch64/fast-datapath/os"
        }]'
```

2.  Check out https://github.com/content-services/content-sources-backend/pull/1368/  and run 'make repos-import'.  This should import the real red hat ones to make sure that works properly with these bad repos in place
3.  Feel free to snapshot those repos with  'make process-repos'  or `go run cmd/external-repos/main.go  snapshot --url $URL`
4. switch to this PR and run   `go run cmd/jobs/main.go  delete-invalid-redhat-repos`
5. Check the database/api/ui to see that the repo was deleted, but the red hat ones remained:
```
 select name, origin from repository_configurations rc inner join repositories r on r.uuid = rc.repository_uuid;
```


---

## Discussion

### Comment by @xbhouse on January 21, 2026 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-10046

---

## Reviews

### Review by @xbhouse - Commented on January 27, 2026 at 06:50 PM UTC

### Review by @xbhouse - Approved on January 27, 2026 at 07:30 PM UTC

looks good! 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1372*
