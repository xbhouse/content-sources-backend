---
type: pull_request
number: 676
title: "Fixes 4054: privatize repos without trailing slash"
state: merged
author: jlsherrill
created: 2024-05-22T19:48:22Z
updated: 2024-05-29T08:00:32Z
closed: 2024-05-29T07:35:34Z
merged: 2024-05-29T07:35:34Z
base_branch: main
head_branch: 4054
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/676
---

# Pull Request #676: Fixes 4054: privatize repos without trailing slash

**Author**: @jlsherrill
**Created**: May 22, 2024 at 07:48 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4054`

## Description

## Summary

This marks all public repos that are missing a trailing slash as not public.  This should cause them to be deleted if there are no repo configs pointing to them.  

## Testing steps

Create a duplicate:  
`make repos-import`
`make db-cli-connect`

```
insert into repositories (uuid, url, public) values (gen_random_uuid(), 'https://cdn.redhat.com/content/dist/rhel8/8.8/x86_64/appstream/os', true);
```

Now run the up migration.  It should no longer be public:

```
select count(*) from repositories where url = 'https://cdn.redhat.com/content/dist/rhel8/8.8/x86_64/appstream/os'
 and public = true;

```

0

```
select count(*) from repositories where url = 'https://cdn.redhat.com/content/dist/rhel8/8.8/x86_64/appstream/os'
 and public = false;
```
1


## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on May 22, 2024 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-4054

### Comment by @swadeley on May 28, 2024 at 07:43 AM UTC

/retest

---

## Reviews

### Review by @xbhouse - Approved on May 28, 2024 at 07:19 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/676*
