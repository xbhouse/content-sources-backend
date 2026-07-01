---
type: pull_request
number: 1013
title: "Build: add PR head refs to github checkout actions"
state: closed
author: dominikvagner
created: 2025-03-11T12:20:39Z
updated: 2025-03-11T17:27:21Z
closed: 2025-03-11T17:27:20Z
base_branch: main
head_branch: fix-checkout-action
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1013
---

# Pull Request #1013: Build: add PR head refs to github checkout actions

**Author**: @dominikvagner
**Created**: March 11, 2025 at 12:20 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `fix-checkout-action`

## Description

## Summary
This adds PR head refs to all GH checkout actions that currently don't have those, this is needed so that the actions run on the same code you are pushing and they don't have different results.

The same thing needed to be done in our FE repo [here](https://github.com/content-services/content-sources-frontend/pull/457). 


---

## Discussion

### Comment by @jlsherrill on March 11, 2025 at 05:27 PM UTC

we discussed this in standup.  We can't actually make this change because pull_request_target doesn't honor the pull request settings, so anyone external could open a PR and the tests will run.  This creates a way for anyone to steal the secrets :/   So we will have to switch back to pull_request which checks out the PR automatically using the checkout action as you'd expect.  I'll go ahead and close this, but let me know if you have any questions!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1013*
