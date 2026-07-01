---
type: pull_request
number: 23
title: "Openapi clients"
state: merged
author: semtexzv
created: 2019-12-02T10:18:15Z
updated: 2019-12-05T15:52:03Z
closed: 2019-12-05T15:52:03Z
merged: 2019-12-05T15:52:03Z
base_branch: master
head_branch: openapi-clients
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/23
---

# Pull Request #23: Openapi clients

**Author**: @semtexzv
**Created**: December 02, 2019 at 10:18 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `openapi-clients`

## Description

- Adds code for handling the 3scale identity strings
- Removes old `hosts` handling code & benchmarks
- Adds scripts to generate OpenApi clients into the _generated file

Goal of this PR is not to finish the whole evaluation flow, but to add OpenAPI generated clients and make listener able to call inventory and VMaaS.

---

## Discussion

### Comment by @semtexzv on December 04, 2019 at 12:32 PM UTC

Moved the clients to separate repo. After we have created a dedicated repo in the RedHatIsights org, I'll move the clients there.

//edit: Clients are now in the https://github.com/RedHatInsights/patchman-clients repo.

---

## Reviews

### Review by @josef-hak - Changes Requested on December 03, 2019 at 02:29 PM UTC

### Review by @semtexzv - Commented on December 04, 2019 at 12:25 PM UTC

### Review by @semtexzv - Commented on December 04, 2019 at 12:26 PM UTC

### Review by @semtexzv - Commented on December 04, 2019 at 12:26 PM UTC

### Review by @semtexzv - Commented on December 05, 2019 at 09:48 AM UTC

### Review by @josef-hak - Changes Requested on December 05, 2019 at 10:11 AM UTC

Cool. Quite a lot of work done. Just some questions and notes.

### Review by @semtexzv - Commented on December 05, 2019 at 12:16 PM UTC

### Review by @semtexzv - Commented on December 05, 2019 at 12:17 PM UTC

### Review by @semtexzv - Commented on December 05, 2019 at 12:18 PM UTC

### Review by @semtexzv - Commented on December 05, 2019 at 12:18 PM UTC

### Review by @josef-hak - Commented on December 05, 2019 at 12:18 PM UTC

### Review by @josef-hak - Commented on December 05, 2019 at 12:20 PM UTC

### Review by @josef-hak - Changes Requested on December 05, 2019 at 02:28 PM UTC

### Review by @josef-hak - Commented on December 05, 2019 at 02:51 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/23*
