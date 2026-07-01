---
type: pull_request
number: 363
title: "Fixes 2367: talk to pulp api directly"
state: merged
author: jlsherrill
created: 2023-08-16T15:36:05Z
updated: 2023-09-04T14:00:32Z
closed: 2023-08-16T19:20:07Z
merged: 2023-08-16T19:20:07Z
base_branch: main
head_branch: 2367
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/363
---

# Pull Request #363: Fixes 2367: talk to pulp api directly

**Author**: @jlsherrill
**Created**: August 16, 2023 at 03:36 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `2367`

## Description

## Summary
do not go through the web container


## Testing steps
Tests pass?

If you want to further test, test a snapshot in ephemeral and verify the task doesn't fail with this deployment.yaml


---

## Discussion

### Comment by @jlsherrill on August 16, 2023 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-2367

### Comment by @jlsherrill on August 16, 2023 at 04:09 PM UTC

updating stage here:  https://gitlab.cee.redhat.com/service/app-interface/-/merge_requests/77682

so you can /lgtm once this is merged

### Comment by @swadeley on August 16, 2023 at 06:17 PM UTC

Hi

re comment 0

not through the web container refers to this diff:
```
-    value: "http://cs-pulp-web-svc:24880/"
+    value: "http://cs-pulp-api-svc:24817/"
```

### Comment by @swadeley on August 16, 2023 at 07:20 PM UTC

Hi

all API tests pass, made a repo and checked it makes a snapshot.

---

## Reviews

### Review by @Andrewgdewar - Approved on August 16, 2023 at 03:46 PM UTC

ack

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/363*
