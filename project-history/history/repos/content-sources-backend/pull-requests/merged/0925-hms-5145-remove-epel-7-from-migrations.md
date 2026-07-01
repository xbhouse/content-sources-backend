---
type: pull_request
number: 925
title: "HMS-5145: remove epel 7 from migrations"
state: merged
author: dominikvagner
created: 2024-12-18T07:48:17Z
updated: 2024-12-19T10:25:04Z
closed: 2024-12-19T10:25:04Z
merged: 2024-12-19T10:25:04Z
base_branch: main
head_branch: 5145
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/925
---

# Pull Request #925: HMS-5145: remove epel 7 from migrations

**Author**: @dominikvagner
**Created**: December 18, 2024 at 07:48 AM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `5145`

## Description

## Summary
This PR removes the `epel 7` repo insert from our DB migrations, this is causing errors when newly created deployments (local, emphemeral) are spun up and it's attempted to introspect this repo, because it no longer exists. 

## Testing steps
1. Spin up a new local deployment with: `make compose-clean && make compose-up`
2. Run nightly jobs (`go run cmd/external-repos/main.go nightly-jobs`) and verify there are no errors concerning epel 7 introspection, like:
    `error introspecting https://dl.fedoraproject.org/pub/epel/7/x86_64/: Cannot fetch https://dl.fedoraproject.org/pub/epel/7/x86_64/repodata/repomd.xml: 404`


---

## Discussion

### Comment by @jlsherrill on December 18, 2024 at 08:00 AM UTC

https://issues.redhat.com/browse/HMS-5145

---

## Reviews

### Review by @jlsherrill - Approved on December 18, 2024 at 08:14 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/925*
