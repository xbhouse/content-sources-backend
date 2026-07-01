---
type: pull_request
number: 956
title: "HMS-5287: Add playwright and github action to backend"
state: merged
author: Andrewgdewar
created: 2025-01-28T18:26:05Z
updated: 2025-02-05T20:00:17Z
closed: 2025-02-05T20:00:17Z
merged: 2025-02-05T20:00:16Z
base_branch: main
head_branch: playwright-POC
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/956
---

# Pull Request #956: HMS-5287: Add playwright and github action to backend

**Author**: @Andrewgdewar
**Created**: January 28, 2025 at 06:26 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `playwright-POC`

## Description

## Summary

No code changes.

Added playwright

Added CICD pipeline for playwright integration.

## Testing steps



---

## Discussion

### Comment by @jlsherrill on January 30, 2025 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-5287

### Comment by @jlsherrill on February 04, 2025 at 09:22 PM UTC

why use _playwright-tests  instead of  playwright-tests or ./tests/playwright/  ?

### Comment by @Andrewgdewar on February 05, 2025 at 03:15 PM UTC

> why use _playwright-tests instead of playwright-tests or ./tests/playwright/ ?

Just wanted to make it the same across front-end/back-end and had no outside input, I can adjust this should we want to move it. 
I'd like to keep the node modules one layer deep for simplicity.

---

## Reviews

### Review by @dominikvagner - Dismissed on February 03, 2025 at 02:16 PM UTC

few small comments _(+Matej's suggestions from the FE PR?)_ otherwise looks good and works! 🎉 nice to see the `.nvmrc` ✨ 
good job 👏🏼

_(also personally I am not that well versed in GH workflows so maybe someone else could also have look) 😅_ 

### Review by @Andrewgdewar - Commented on February 03, 2025 at 04:19 PM UTC

### Review by @Andrewgdewar - Commented on February 03, 2025 at 04:19 PM UTC

### Review by @Andrewgdewar - Commented on February 03, 2025 at 04:19 PM UTC

### Review by @Andrewgdewar - Commented on February 03, 2025 at 05:38 PM UTC

### Review by @swadeley - Commented on February 03, 2025 at 08:44 PM UTC

### Review by @swadeley - Commented on February 03, 2025 at 08:46 PM UTC

### Review by @swadeley - Commented on February 03, 2025 at 08:48 PM UTC

### Review by @dominikvagner - Commented on February 04, 2025 at 07:46 AM UTC

nice updates, more/better docs are always appreciated 😌 
but the `.DS_Store` files also need be removed from the commit along with ignoring them, so run this `git rm --cached "*.DS_Store"`, commit, and will be good to go 😄

### Review by @swadeley - Commented on February 04, 2025 at 04:04 PM UTC

### Review by @Andrewgdewar - Commented on February 04, 2025 at 05:10 PM UTC

### Review by @xbhouse - Commented on February 04, 2025 at 08:25 PM UTC

### Review by @xbhouse - Commented on February 04, 2025 at 08:31 PM UTC

### Review by @jlsherrill - Commented on February 04, 2025 at 09:28 PM UTC

### Review by @jlsherrill - Commented on February 04, 2025 at 09:29 PM UTC

### Review by @jlsherrill - Commented on February 04, 2025 at 09:47 PM UTC

### Review by @Andrewgdewar - Commented on February 05, 2025 at 04:27 PM UTC

### Review by @jlsherrill - Approved on February 05, 2025 at 07:04 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/956*
