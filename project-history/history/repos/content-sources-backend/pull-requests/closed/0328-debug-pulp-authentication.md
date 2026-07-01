---
type: pull_request
number: 328
title: "Debug pulp authentication"
state: closed
author: jlsherrill
created: 2023-07-03T17:24:03Z
updated: 2023-07-06T19:28:42Z
closed: 2023-07-06T19:28:41Z
base_branch: main
head_branch: pulp_app
labels: []
url: https://github.com/content-services/content-sources-backend/pull/328
---

# Pull Request #328: Debug pulp authentication

**Author**: @jlsherrill
**Created**: July 03, 2023 at 05:24 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `pulp_app`

## Description

## Summary

Currently debugging authentication errors in ephemeral environments 

Current thought is that the backend app is up but not properly configured for pulp, although there is some evidence this isn't the case.  This change means that our app will not be 'up' until pulp is fully up and running.  

TODO:  make this work without a pulp server!
DO NOT MERGE

## Testing steps


---

## Discussion

### Comment by @swadeley on July 05, 2023 at 07:36 AM UTC

Hi

Where would you expect to see this message:
message = "Pulp DB Connection not available."

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/328*
