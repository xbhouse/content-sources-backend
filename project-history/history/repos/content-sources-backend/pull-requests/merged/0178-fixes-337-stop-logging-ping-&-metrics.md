---
type: pull_request
number: 178
title: "Fixes 337: stop logging ping & metrics"
state: merged
author: jlsherrill
created: 2023-01-17T13:19:35Z
updated: 2023-01-17T14:54:17Z
closed: 2023-01-17T14:54:17Z
merged: 2023-01-17T14:54:17Z
base_branch: main
head_branch: 337
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/178
---

# Pull Request #178: Fixes 337: stop logging ping & metrics

**Author**: @jlsherrill
**Created**: January 17, 2023 at 01:19 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `337`

## Description

*No description provided*

---

## Discussion

### Comment by @jlsherrill on January 17, 2023 at 01:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-337

### Comment by @jlsherrill on January 17, 2023 at 02:54 PM UTC

actually i kinda prefer the logging struct there, as its part of the configuration structure as a whole

---

## Reviews

### Review by @rverdile - Approved on January 17, 2023 at 02:28 PM UTC

it might also make sense to move the logging struct (https://github.com/content-services/content-sources-backend/blob/a77d175c28932954f7e26d4386ebc759abb33b41/pkg/config/config.go#L45-L48) to `logging.go`, but tested and lgtm.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/178*
