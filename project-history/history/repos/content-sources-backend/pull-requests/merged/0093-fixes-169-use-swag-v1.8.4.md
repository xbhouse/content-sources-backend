---
type: pull_request
number: 93
title: "Fixes 169: Use swag v1.8.4"
state: merged
author: rverdile
created: 2022-08-31T12:30:09Z
updated: 2022-09-01T13:50:51Z
closed: 2022-09-01T13:50:51Z
merged: 2022-09-01T13:50:51Z
base_branch: main
head_branch: pin-swag
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/93
---

# Pull Request #93: Fixes 169: Use swag v1.8.4

**Author**: @rverdile
**Created**: August 31, 2022 at 12:30 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `pin-swag`

## Description

For now, to avoid breaking change in v1.8.5

I can't find a simple way to pin package versions with Go. go.mod doesn't work quite like a typical requirements file. It seems the two ways are to download the wanted version manually or keep the wanted version locally and replace the requirement with the local version.

---

## Discussion

### Comment by @jlsherrill on August 31, 2022 at 12:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-169

### Comment by @mshriver on August 31, 2022 at 03:23 PM UTC

Looks pretty sane to me, and I'd agree with limiting the scope of complexity to achieve this since the intent is for it to be temporary until a release including the necessary fix is made.

### Comment by @rverdile on September 01, 2022 at 01:35 PM UTC

@avisiedo qe-testing-needed means that after a developer acks it, qe also needs to ack it, so it was still good for you to review it. Thanks :).

---

## Reviews

### Review by @avisiedo - Dismissed on September 01, 2022 at 08:19 AM UTC

nothing to add; lgtm

### Review by @mshriver - Approved on September 01, 2022 at 01:50 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/93*
