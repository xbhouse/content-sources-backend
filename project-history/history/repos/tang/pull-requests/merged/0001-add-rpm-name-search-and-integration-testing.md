---
type: pull_request
number: 1
title: "Add rpm name search and integration testing"
state: merged
author: rverdile
created: 2023-12-08T22:40:47Z
updated: 2024-01-03T14:22:49Z
closed: 2024-01-03T14:22:45Z
merged: 2024-01-03T14:22:45Z
base_branch: main
head_branch: rpm-search
labels: []
url: https://github.com/content-services/tang/pull/1
---

# Pull Request #1: Add rpm name search and integration testing

**Author**: @rverdile
**Created**: December 08, 2023 at 10:40 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `rpm-search`

## Description

### Summary
Adds
- An exported interface, Tangy, that that includes the method `RpmRepositoryVersionPackageSearch()`, to search RPMs by name given a list of repository hrefs.
- Tangy includes database and logging configuration with zerolog.
- 3 internal (not exported) packages used for testing/development purposes.
- Make commands to start up a pulp server with podman-compose.
- A CI integration test for searching RPMs by name
- A CI test for linting
- A README with instructions on to how use this repository


---

## Discussion

### Comment by @jlsherrill on December 20, 2023 at 10:06 PM UTC

Also, would it make sense to include a Mock of the interface here? Instead of in our application.

### Comment by @rverdile on December 21, 2023 at 04:53 PM UTC

updated

---

## Reviews

### Review by @jlsherrill - Commented on December 13, 2023 at 02:01 PM UTC

### Review by @rverdile - Commented on December 13, 2023 at 08:27 PM UTC

### Review by @jlsherrill - Commented on December 19, 2023 at 03:25 PM UTC

### Review by @jlsherrill - Commented on December 19, 2023 at 04:10 PM UTC

### Review by @jlsherrill - Commented on December 19, 2023 at 09:11 PM UTC

### Review by @jlsherrill - Commented on December 19, 2023 at 10:02 PM UTC

### Review by @jlsherrill - Commented on December 21, 2023 at 04:58 PM UTC

### Review by @rverdile - Commented on December 21, 2023 at 06:33 PM UTC

### Review by @jlsherrill - Commented on December 21, 2023 at 09:12 PM UTC

### Review by @jlsherrill - Commented on December 21, 2023 at 09:13 PM UTC

### Review by @jlsherrill - Approved on January 02, 2024 at 06:14 PM UTC

---

*Archived from: https://github.com/content-services/tang/pull/1*
