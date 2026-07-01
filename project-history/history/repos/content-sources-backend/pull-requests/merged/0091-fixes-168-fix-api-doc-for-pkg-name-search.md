---
type: pull_request
number: 91
title: "Fixes 168: fix api doc for pkg name search"
state: merged
author: jlsherrill
created: 2022-08-29T19:41:19Z
updated: 2022-08-29T20:42:02Z
closed: 2022-08-29T20:41:35Z
merged: 2022-08-29T20:41:35Z
base_branch: main
head_branch: 168
labels: ["bug", "qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/91
---

# Pull Request #91: Fixes 168: fix api doc for pkg name search

**Author**: @jlsherrill
**Created**: August 29, 2022 at 07:41 PM UTC
**Status**: Merged
**Labels**: `bug`, `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `168`

## Description

*No description provided*

---

## Discussion

### Comment by @jlsherrill on August 29, 2022 at 07:41 PM UTC

regen'ing the docs seemed to update some other stuff... let me see if i can tell why

EDIT: i think one of the previous prs needed a regenerating, but due to the segfault bug in swag we didn't notice

### Comment by @jlsherrill on August 29, 2022 at 08:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-168

### Comment by @mshriver on August 29, 2022 at 08:42 PM UTC

Local testing of the spec updates was successful, merging this and pushing them to plugin repo to unblock initial tests for search endpoint.

---

## Reviews

### Review by @mshriver - Approved on August 29, 2022 at 07:45 PM UTC

@swadeley test for the endpoint uncovered this, thanks @jlsherrill !

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/91*
