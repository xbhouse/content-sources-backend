---
type: pull_request
number: 1215
title: "Build: fix instructions for running go integration tests"
state: merged
author: xbhouse
created: 2025-09-23T20:55:25Z
updated: 2025-09-24T15:14:01Z
closed: 2025-09-24T15:14:01Z
merged: 2025-09-24T15:14:01Z
base_branch: main
head_branch: fix-readme
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1215
---

# Pull Request #1215: Build: fix instructions for running go integration tests

**Author**: @xbhouse
**Created**: September 23, 2025 at 08:55 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `fix-readme`

## Description

## Summary

Updates the IP address for the pulp.content host entry needed to run the go integration tests

## Testing steps



---

## Discussion

### Comment by @swadeley on September 24, 2025 at 07:38 AM UTC

Hi @xbhouse ,

 is this to get Pulp content from stage?
 If so, would I not need the IP address from the VPN tunnel endpoint?

Or is this for testing agsint local backend or frontend?
If so, IP addresses do change for those using DNS at home.
Integration tests have been working for me when I run against stage

### Comment by @xbhouse on September 24, 2025 at 11:36 AM UTC

@swadeley these instructions are for running the go integration tests (not playwright). the readme was wrong and was causing some confusion when tests would fail with connection errors. 

i'll clarify this in the readme too!

### Comment by @swadeley on September 24, 2025 at 11:45 AM UTC

> @swadeley these instructions are for running the go integration tests (not playwright). the readme was wrong and was causing some confusion when tests would fail with connection errors.
> 
> i'll clarify this in the readme too!

OK, it is in the "Developing" section , but good to make it clear. Thank you.

### Comment by @rverdile on September 24, 2025 at 12:51 PM UTC

/retest

---

## Reviews

### Review by @swadeley - Approved on September 24, 2025 at 11:45 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1215*
