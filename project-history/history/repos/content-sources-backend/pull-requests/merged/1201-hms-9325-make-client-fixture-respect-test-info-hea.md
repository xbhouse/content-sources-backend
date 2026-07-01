---
type: pull_request
number: 1201
title: "HMS-9325: make client fixture respect test info headers"
state: merged
author: dominikvagner
created: 2025-09-16T12:16:12Z
updated: 2025-09-22T12:27:44Z
closed: 2025-09-22T12:27:44Z
merged: 2025-09-22T12:27:44Z
base_branch: main
head_branch: test-utils-extrahttpheaders
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1201
---

# Pull Request #1201: HMS-9325: make client fixture respect test info headers

**Author**: @dominikvagner
**Created**: September 16, 2025 at 12:16 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `test-utils-extrahttpheaders`

## Description

## Summary
This makes the Playwright client fixture inside of test-utils respect the 'extraHTTPHeaders' passed in from the TestInfo. That will enable setting test wide headers for specific users, which will be helpful for API integration tests.

## Testing steps
Tests pass. 
_This was tested manually with integration tests in our frontend repo. If you wanted to the same, pull the changes from this PR into test-utils in the FE repo and in integration tests that use the client, you should be able to set the `extraHTTPHeaders` for the whole test, as we do it in other tests._


---

## Discussion

### Comment by @jlsherrill on September 16, 2025 at 12:30 PM UTC

https://issues.redhat.com/browse/HMS-9325

### Comment by @rverdile on September 17, 2025 at 01:54 PM UTC

/retest

### Comment by @dominikvagner on September 19, 2025 at 03:12 PM UTC

/retest

---

## Reviews

### Review by @mayurilahane - Approved on September 18, 2025 at 06:33 PM UTC

Tried out these changes locally with integration tests which were using another user and it is working as expected.

Just needed to set test.use() to that specific user to get it working.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1201*
