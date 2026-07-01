---
type: pull_request
number: 979
title: "Test: add format and lint checks to PW tests"
state: merged
author: dominikvagner
created: 2025-02-13T22:52:35Z
updated: 2025-02-18T17:37:37Z
closed: 2025-02-18T17:37:37Z
merged: 2025-02-18T17:37:37Z
base_branch: main
head_branch: add-pw-tests-format-lint
labels: []
url: https://github.com/content-services/content-sources-backend/pull/979
---

# Pull Request #979: Test: add format and lint checks to PW tests

**Author**: @dominikvagner
**Created**: February 13, 2025 at 10:52 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `add-pw-tests-format-lint`

## Description

## Summary
This PR adds 2 GH actions to check formatting and linting of our PW tests. 🔍 To do this new scripts are added to our `packages.json` which can be used with `yarn` (few dev dependencies are needed to do this, and also couple of format and lint config files). 🤖⚙️

New `yarn` scripts:
- `yarn lint` - checks for linting warnings and errors
- `yarn lint:fix` - fixes all auto-fixable linting problems
- `yarn format` - auto formats our PW tests code
- `yarn format:check` - check for formatting violations
- `yarn verify` - runs checks for linting, formatting and all PW tests

## Testing steps
1. Run tests on main and on this PR, should be the same outcome.
2. Test out the new scripts/commands and verify their functionality.



---

## Reviews

### Review by @Andrewgdewar - Approved on February 18, 2025 at 05:37 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/979*
