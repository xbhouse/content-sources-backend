---
type: pull_request
number: 119
title: "Fixes 186: Refactor test suites to run individual test files"
state: merged
author: rverdile
created: 2022-10-03T15:51:45Z
updated: 2022-10-07T13:44:20Z
closed: 2022-10-07T13:44:13Z
merged: 2022-10-07T13:44:13Z
base_branch: main
head_branch: suite-refactor
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/119
---

# Pull Request #119: Fixes 186: Refactor test suites to run individual test files

**Author**: @rverdile
**Created**: October 03, 2022 at 03:51 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `suite-refactor`

## Description

This PR makes it possible (in GoLand) to right-click on a testify-suite-using test file like `dao/repository_config.go`, click "Run", and only run the tests in that file.

This PR should get a review from both a GoLand user and a VSCode user. I'm not sure how it affects VSCode.

Also, I refactored the dao test suite to embedded a new `DaoSuite` type, which deduplicates some code. I did the same thing in the models test suite, but it's a more obvious change there.

---

## Discussion

### Comment by @jlsherrill on October 03, 2022 at 04:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-186

### Comment by @jlsherrill on October 06, 2022 at 05:11 PM UTC

This all looks good and works great in both goland and vscode.  Just one question and then i think its good to go

### Comment by @rverdile on October 06, 2022 at 05:21 PM UTC

> This all looks good and works great in both goland and vscode. Just one question and then i think its good to go

I think I answered that question

---

## Reviews

### Review by @rverdile - Commented on October 03, 2022 at 03:55 PM UTC

### Review by @rverdile - Commented on October 03, 2022 at 06:14 PM UTC

### Review by @jlsherrill - Commented on October 06, 2022 at 03:52 PM UTC

### Review by @rverdile - Commented on October 06, 2022 at 03:58 PM UTC

### Review by @jlsherrill - Commented on October 06, 2022 at 07:17 PM UTC

### Review by @jlsherrill - Commented on October 06, 2022 at 07:18 PM UTC

### Review by @rverdile - Commented on October 07, 2022 at 01:20 PM UTC

### Review by @jlsherrill - Commented on October 07, 2022 at 01:24 PM UTC

### Review by @jlsherrill - Approved on October 07, 2022 at 01:24 PM UTC

### Review by @rverdile - Commented on October 07, 2022 at 01:38 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/119*
