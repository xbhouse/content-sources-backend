---
type: pull_request
number: 1181
title: "HMS-8904: Test to introspect pulp project repos "
state: closed
author: mayurilahane
created: 2025-08-18T19:35:15Z
updated: 2025-08-25T03:05:43Z
closed: 2025-08-21T13:05:35Z
base_branch: main
head_branch: mlahane/HMS-8904
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1181
---

# Pull Request #1181: HMS-8904: Test to introspect pulp project repos 

**Author**: @mayurilahane
**Created**: August 18, 2025 at 07:35 PM UTC
**Status**: Closed
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `mlahane/HMS-8904`

## Description

## Summary
Migrate test - https://gitlab.cee.redhat.com/insights-qe/iqe-content-sources-plugin/-/blob/master/iqe_content_sources/tests/test_repository_api_only.py?ref_type=heads#L1060

## Testing steps
peer review 



---

## Discussion

### Comment by @jlsherrill on August 18, 2025 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-8904

### Comment by @swadeley on August 20, 2025 at 10:05 AM UTC

@mayurilahane Test steps look like a good match to the old IQE test.

### Comment by @rverdile on August 20, 2025 at 07:52 PM UTC

Should this test be an integration test? It looks like the iqe test was using the "stable sam" account

### Comment by @mayurilahane on August 21, 2025 at 01:06 PM UTC

Moving this test to integration test folder 
raised  
https://github.com/content-services/content-sources-frontend/pull/623


---

## Reviews

### Review by @swadeley - Commented on August 20, 2025 at 09:49 AM UTC

### Review by @swadeley - Commented on August 20, 2025 at 09:56 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1181*
