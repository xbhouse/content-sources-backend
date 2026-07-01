---
type: pull_request
number: 140
title: "Fixes 209: update base container image"
state: merged
author: jlsherrill
created: 2022-11-03T18:45:18Z
updated: 2022-11-07T13:44:04Z
closed: 2022-11-07T13:44:00Z
merged: 2022-11-07T13:44:00Z
base_branch: main
head_branch: 209
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/140
---

# Pull Request #140: Fixes 209: update base container image

**Author**: @jlsherrill
**Created**: November 03, 2022 at 06:45 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `209`

## Description

and swith to ubi8 micro

---

## Discussion

### Comment by @jlsherrill on November 03, 2022 at 07:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-209

### Comment by @jlsherrill on November 04, 2022 at 12:47 AM UTC

looked at using micro, but didn't have a trusted ca store... so syncing didn't work :)

### Comment by @avisiedo on November 04, 2022 at 01:05 PM UTC

LGTM

- image built and pushed to quay.io
- [Clair](https://github.com/quay/clair) scanner report security scan passed.

---

## Reviews

### Review by @avisiedo - Approved on November 04, 2022 at 01:05 PM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/140*
