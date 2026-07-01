---
type: pull_request
number: 21
title: "handle compressed comps"
state: merged
author: xbhouse
created: 2023-12-20T20:53:09Z
updated: 2024-01-04T15:22:52Z
closed: 2024-01-04T15:22:52Z
merged: 2024-01-04T15:22:52Z
base_branch: master
head_branch: compressed-comps
labels: []
url: https://github.com/content-services/yummy/pull/21
---

# Pull Request #21: handle compressed comps

**Author**: @xbhouse
**Created**: December 20, 2023 at 08:53 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `compressed-comps`

## Description

Adds support for compressed comps.xml files. 

Without this fix, introspection of a compressed comps results in an error that contains non-UTF8 characters in the error message which causes [this error](https://issues.redhat.com/browse/HMS-3231) being seen in stage and prod: 

`Fatal error introspecting repository https://fixtures.pulpproject.org/rpm-unsigned/ error="failed to update introspection timestamps: ERROR: invalid byte sequence for encoding \"UTF8\": 0xf4 0x62 0x20 0x28 (SQLSTATE 22021)" request_id=deQEWFigbmWEtXZNkXmUFOpnbpKzjJyc task_id=cb471363-688d-4f71-8635-983bb00cf9f4 task_type=introspect`

---

## Reviews

### Review by @jlsherrill - Commented on December 21, 2023 at 06:41 PM UTC

### Review by @xbhouse - Commented on January 03, 2024 at 08:04 PM UTC

### Review by @jlsherrill - Approved on January 04, 2024 at 03:01 PM UTC

---

*Archived from: https://github.com/content-services/yummy/pull/21*
