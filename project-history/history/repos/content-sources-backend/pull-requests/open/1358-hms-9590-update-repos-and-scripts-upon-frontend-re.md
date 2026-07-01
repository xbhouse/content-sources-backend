---
type: pull_request
number: 1358
title: "HMS-9590: update repos and scripts upon frontend refactor"
state: open
author: Starle21
created: 2026-01-09T08:17:38Z
updated: 2026-03-06T10:59:28Z
base_branch: main
head_branch: HMS-9590
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1358
---

# Pull Request #1358: HMS-9590: update repos and scripts upon frontend refactor

**Author**: @Starle21
**Created**: January 09, 2026 at 08:17 AM UTC
**Status**: Open
**Labels**: None
**Base**: `main` ← **Head**: `HMS-9590`

## Description

Due to the frontend refactor in frontend PR #763,
there was a need to update playwright tests
and the accompanied repositories and the script
that are used for local playwright tests.

Now, `repos-minimal` script imports `hardcoded` repositories
next to the original `epel10` and `small` repositories.

A new script which counts
the duration how long it takes to import
and snapshot or introspect the required repos
for testing has been added.



---

## Discussion

### Comment by @xbhouse on January 09, 2026 at 08:30 AM UTC

https://issues.redhat.com/browse/HMS-9590

---

## Reviews

### Review by @dominikvagner - Commented on January 09, 2026 at 09:10 AM UTC

1 comment for the new command, I think we should update the old one instead 💭 

and this PR is also missing an update to the GH workflow for CI where the command is used _(we also run frontend tests there, so those would break after the FE changes)_ 👀 
one other thing is, if the `small` codeready repo would be removed from the setup, some of our backend PW tests would break
_(also in the helpers inside test-utils, we have a constant with the URL for the small testing repo, that might also need to be changed, if the repo isn't going to be available)_ 😅 

### Review by @Starle21 - Commented on January 09, 2026 at 10:29 AM UTC

### Review by @dominikvagner - Commented on January 12, 2026 at 10:20 AM UTC

### Review by @Starle21 - Commented on March 03, 2026 at 07:52 AM UTC

### Review by @dominikvagner - Approved on March 06, 2026 at 10:59 AM UTC

ack! ✅ 🚀 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1358*
