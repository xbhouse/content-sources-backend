---
type: pull_request
number: 142
title: "Fixes 200: yummy refactor"
state: merged
author: rverdile
created: 2022-11-08T16:16:57Z
updated: 2022-12-09T19:55:22Z
closed: 2022-12-09T19:55:19Z
merged: 2022-12-09T19:55:19Z
base_branch: main
head_branch: yummy
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/142
---

# Pull Request #142: Fixes 200: yummy refactor

**Author**: @rverdile
**Created**: November 08, 2022 at 04:16 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `yummy`

## Description

Companion to this PR: https://github.com/content-services/yummy/pull/7

Moves the functionality in `dao/external_resources.go` to yummy. Refactors to adjust to yummy changes. 

Feedback welcome, architectural feedback in particular.

To test, you'll need to put this in your go.mod: `replace github.com/content-services/yummy => /path/to/yummy`


TODO

- [x] Fix tests.

---

## Discussion

### Comment by @jlsherrill on November 18, 2022 at 09:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-200

### Comment by @jlsherrill on November 30, 2022 at 01:11 PM UTC

Overall, a big +1, i don't have any major comments

### Comment by @jlsherrill on December 05, 2022 at 05:17 PM UTC

looks like it needs a rebase

### Comment by @rverdile on December 05, 2022 at 08:55 PM UTC

not sure if CI failure here is related

### Comment by @rverdile on December 06, 2022 at 07:17 PM UTC

~~Current~~ Previous test failure is because the `yum.YumRepository` object in our backend is instantiated when the handler is registered.  This means that repeated calls to yummy do not fetch new data because the handler is registered only when the app first loads.

Opened this PR to yummy: https://github.com/content-services/yummy/pull/8, so that we have a method to clear the data.

---

## Reviews

### Review by @rverdile - Commented on November 11, 2022 at 08:25 PM UTC

### Review by @rverdile - Commented on November 11, 2022 at 08:26 PM UTC

### Review by @rverdile - Commented on November 28, 2022 at 07:58 PM UTC

### Review by @jlsherrill - Commented on November 29, 2022 at 09:07 PM UTC

### Review by @jlsherrill - Commented on November 29, 2022 at 09:08 PM UTC

### Review by @jlsherrill - Commented on November 29, 2022 at 09:09 PM UTC

### Review by @jlsherrill - Approved on December 02, 2022 at 01:16 PM UTC

ACK, we'll just need to tag yummy and then update the go.mod here

### Review by @jlsherrill - Commented on December 07, 2022 at 02:28 PM UTC

### Review by @rverdile - Commented on December 07, 2022 at 03:01 PM UTC

### Review by @rverdile - Commented on December 07, 2022 at 04:34 PM UTC

### Review by @jlsherrill - Approved on December 09, 2022 at 05:23 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/142*
