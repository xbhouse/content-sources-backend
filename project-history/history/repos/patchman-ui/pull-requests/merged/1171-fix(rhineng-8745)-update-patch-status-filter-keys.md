---
type: pull_request
number: 1171
title: "fix(RHINENG-8745): update patch status filter keys"
state: merged
author: mkholjuraev
created: 2024-03-08T19:03:07Z
updated: 2024-03-13T13:32:38Z
closed: 2024-03-13T13:07:11Z
merged: 2024-03-13T13:07:11Z
base_branch: master
head_branch: fix-system_updatable
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1171
---

# Pull Request #1171: fix(RHINENG-8745): update patch status filter keys

**Author**: @mkholjuraev
**Created**: March 08, 2024 at 07:03 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-system_updatable`

## Description

# Description

Associated Jira ticket: # [(issue)](https://issues.redhat.com/browse/RHINENG-8745)

This fixes the broken packages page by updating the Patch status filter key from system_updatable which is obsolete to the new systems_applicable


# How to test the PR
1. Visit the packages page
2. Observe that the page works as expected and the table is filtered by **Systems with patches available** by default
3. From the filters dropdown find **Patch status** filter
4. Play with it, make sure it works correclt

# Before the change


# After the change


# Dependent work link


# Checklist:

- [ ] The commit message has the Jira ticket linked
- [ ] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @mkholjuraev on March 13, 2024 at 01:32 PM UTC

:tada: This PR is included in version 1.67.2 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.67.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @AsToNlele - Approved on March 13, 2024 at 11:39 AM UTC

LGTM, thanks for the fix :smile: 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1171*
