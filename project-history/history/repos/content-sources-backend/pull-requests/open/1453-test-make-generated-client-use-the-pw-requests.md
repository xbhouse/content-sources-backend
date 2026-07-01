---
type: pull_request
number: 1453
title: "Test: make generated client use the PW requests "
state: open
author: dominikvagner
created: 2026-04-08T14:52:20Z
updated: 2026-06-25T09:35:06Z
base_branch: main
head_branch: ts-client-playwright-network-logs
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1453
---

# Pull Request #1453: Test: make generated client use the PW requests 

**Author**: @dominikvagner
**Created**: April 08, 2026 at 02:52 PM UTC
**Status**: Open
**Labels**: None
**Base**: `main` ← **Head**: `ts-client-playwright-network-logs`

## Description

## Summary
This changes the way how the generated typescript client, we use for
testing, sends the requests. Previously we used the default or undici
for this, but that meant we couldn't see those requests in the PW
traces, making it hard to debug.
We can utilize the built-in Playwright request fixture and API context
to do this, by providing a custom "FetchAPI" to the generated client.
Also makes the client dynamically pick-up fresh credentials per request
by default.

And makes a few smaller tweaks in the first 2 commits:
> This fixes an issue with a wrong describe used in `
> auth.setup.ts`,
> changes the project name of API tests to something appropriate and
> quiets unnecessary logs.
> 
> This fixes incorrect docs and makes the interval longer as 10 ms was way
> too spammy.

## Testing steps
All tests pass and the network logs in Playwright traces display the requests made by the client.

---

## Discussion

### Comment by @dominikvagner on April 09, 2026 at 03:37 PM UTC

/retest

### Comment by @dominikvagner on April 10, 2026 at 07:25 AM UTC

/retest

### Comment by @rverdile on May 27, 2026 at 03:20 PM UTC

@dominikvagner did you wanna merge this? :)

---

## Reviews

### Review by @katarinazaprazna - Approved on April 10, 2026 at 04:53 PM UTC

Thanks for the changes! This is super helpful. Nice work 🙌

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1453*
