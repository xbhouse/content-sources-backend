---
type: pull_request
number: 1213
title: "HMS-8907: migrate test_create_validation"
state: merged
author: swadeley
created: 2025-09-23T13:13:00Z
updated: 2025-10-03T07:15:53Z
closed: 2025-10-03T07:15:53Z
merged: 2025-10-03T07:15:53Z
base_branch: main
head_branch: swadeley/HMS-8907
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1213
---

# Pull Request #1213: HMS-8907: migrate test_create_validation

**Author**: @swadeley
**Created**: September 23, 2025 at 01:13 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `swadeley/HMS-8907`

## Description

## Summary

Add URL whitespace validation
[HMS-8907 migrate test_create_validation](https://issues.redhat.com/browse/HMS-8907)

## Testing steps
tests pass 

Note: Although the original IQE test (link in Jira) was an API test I was asked many months ago to move input validation checks to unit tests.



---

## Discussion

### Comment by @jlsherrill on September 23, 2025 at 01:30 PM UTC

https://issues.redhat.com/browse/HMS-8907

### Comment by @xbhouse on September 24, 2025 at 07:38 PM UTC

i think this could also be included as part of the `Validate repository parameters` playwright API test

### Comment by @swadeley on September 26, 2025 at 08:03 AM UTC

> i think this could also be included as part of the `Validate repository parameters` playwright API test

Hi @xbhouse , I was told some time ago to move validation of inputs to unit tests. Lets see what others think. 

### Comment by @xbhouse on September 26, 2025 at 12:40 PM UTC

> > i think this could also be included as part of the `Validate repository parameters` playwright API test
> 
> Hi @xbhouse , I was told some time ago to move validation of inputs to unit tests. Lets see what others think.

ok, i recall you mentioning that in standup -- thanks for clarifying :) it might be helpful to include that context in the PR summary so reviewers understand the reasoning

---

## Reviews

### Review by @katarinazaprazna - Approved on October 02, 2025 at 07:55 PM UTC

Thanks for your work on this! Approved :) 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1213*
