---
type: pull_request
number: 633
title: "Fixes 3939: create/update RH repos as public"
state: merged
author: jlsherrill
created: 2024-04-11T15:39:37Z
updated: 2024-04-12T07:25:26Z
closed: 2024-04-12T07:25:26Z
merged: 2024-04-12T07:25:26Z
base_branch: main
head_branch: 3939
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/633
---

# Pull Request #633: Fixes 3939: create/update RH repos as public

**Author**: @jlsherrill
**Created**: April 11, 2024 at 03:39 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `3939`

## Description

## Summary
Create all red hat repos as 'public' so they can be searched just like public repos.

## Testing steps
In master, revert this change:
https://github.com/content-services/content-sources-backend/commit/542f8f32f1c2770609b5bb2c77b62c78ff72a069
then run:
```make compose-clean compose-up repos-import```

Now, verify that some RH repos are not public:

```
make db-cli-connect
# select origin, url, public from repositories;
```

now check out this PR and re-run ```make repos-import``` and rerun the query:

```
make db-cli-connect
# select origin, url, public from repositories;
```

all RH repos should now have public  set to true.

For QE, this will be hard to fully test outside of stage, as the new deployments have all the RH repos in external_repos.json, so they will be public on new deployments.

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on April 11, 2024 at 04:09 PM UTC

https://issues.redhat.com/browse/HMS-3939

---

## Reviews

### Review by @xbhouse - Approved on April 11, 2024 at 05:46 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/633*
