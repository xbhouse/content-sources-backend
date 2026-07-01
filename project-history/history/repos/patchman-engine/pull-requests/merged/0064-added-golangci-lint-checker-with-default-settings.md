---
type: pull_request
number: 64
title: "Added golangci-lint checker with default settings"
state: merged
author: josef-hak
created: 2020-01-15T16:30:42Z
updated: 2020-01-17T12:33:25Z
closed: 2020-01-16T08:54:22Z
merged: 2020-01-16T08:54:22Z
base_branch: master
head_branch: code_analys
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/64
---

# Pull Request #64: Added golangci-lint checker with default settings

**Author**: @josef-hak
**Created**: January 15, 2020 at 04:30 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `code_analys`

## Description

- added automatic checking
- fixed lint hints

---

## Discussion

### Comment by @beav on January 15, 2020 at 04:46 PM UTC

@Josca can you check if the linting was executed in the PR tests? i didn't see it but its very possible I missed it when scrolling through.

### Comment by @josef-hak on January 16, 2020 at 08:44 AM UTC

@beav Yes, it was only visible when code analysis failed. It was not ideal. I've added simple [message](https://github.com/RedHatInsights/patchman-engine/pull/64/checks?check_run_id=392736900#step:3:2173) after analysis finishing.

---

## Reviews

### Review by @beav - Commented on January 15, 2020 at 04:43 PM UTC

### Review by @josef-hak - Commented on January 16, 2020 at 08:19 AM UTC

### Review by @semtexzv - Approved on January 16, 2020 at 08:49 AM UTC

Linting failing fails the test run & lint runs as a part of the testing process. :+1: 

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/64*
