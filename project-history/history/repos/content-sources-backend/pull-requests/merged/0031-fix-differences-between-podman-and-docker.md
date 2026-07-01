---
type: pull_request
number: 31
title: "Fix differences between podman and docker"
state: merged
author: jlsherrill
created: 2022-06-10T19:00:43Z
updated: 2022-06-13T14:02:23Z
closed: 2022-06-13T14:02:18Z
merged: 2022-06-13T14:02:18Z
base_branch: main
head_branch: docker_support
labels: []
url: https://github.com/content-services/content-sources-backend/pull/31
---

# Pull Request #31: Fix differences between podman and docker

**Author**: @jlsherrill
**Created**: June 10, 2022 at 07:00 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `docker_support`

## Description

docker doesn't support 'container exists', so use inspect instead.
Docker and podman use different health text in inspect, so check both

---

## Discussion

### Comment by @jlsherrill on June 10, 2022 at 07:01 PM UTC

@Andrewgdewar if you can check this works for you too, that'd be helpful!

### Comment by @jlsherrill on June 10, 2022 at 07:02 PM UTC

@avisiedo   @Andrewgdewar is using docker on an M1 mac, and ran into some issues, so i think i fixed them here.

### Comment by @avisiedo on June 13, 2022 at 05:51 AM UTC

@jlsherrill @Andrewgdewar I have added some suggestions, I could check in macos, but not with a m1 core; maybe more things could change.

### Comment by @jlsherrill on June 13, 2022 at 12:04 PM UTC

@Andrewgdewar can you give this a spin and db-down  db-clean  and db-up again to make sure it still works for you?

---

## Reviews

### Review by @avisiedo - Commented on June 13, 2022 at 05:26 AM UTC

### Review by @avisiedo - Commented on June 13, 2022 at 05:47 AM UTC

### Review by @jlsherrill - Commented on June 13, 2022 at 10:50 AM UTC

### Review by @avisiedo - Commented on June 13, 2022 at 11:42 AM UTC

### Review by @Andrewgdewar - Approved on June 13, 2022 at 01:57 PM UTC

Confirmed this is working for me. Feel free to get other approvals and/or ping me if further changes are made. ACK.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/31*
