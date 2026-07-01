---
type: pull_request
number: 1175
title: "chore: refactor advisory systems table"
state: merged
author: mkholjuraev
created: 2024-03-14T15:57:50Z
updated: 2024-05-13T12:01:56Z
closed: 2024-04-17T15:19:38Z
merged: 2024-04-17T15:19:38Z
base_branch: master
head_branch: refactor-advisor-systems
labels: ["released", "refactor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1175
---

# Pull Request #1175: chore: refactor advisory systems table

**Author**: @mkholjuraev
**Created**: March 14, 2024 at 03:57 PM UTC
**Status**: Merged
**Labels**: `released`, `refactor`
**Base**: `master` ← **Head**: `refactor-advisor-systems`

## Description

# Description

Associated Jira ticket: # (issue)

This PR is intended to refactor the advisory systems table to make the code cleaner. The refactor only divides existing code to 2 components AdvisorySystems and AdvisorySystemsTable. AdvisorySystems is a wrapper around the table and the remediation wizard. AdvisorySystemsTable is solely for the table functionality.

It also has the tests rewritten


# How to test the PR

Please include steps to test your PR.

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

### Comment by @Fewwy on March 14, 2024 at 04:19 PM UTC

It looks good, I ran the tests and they are passing. Do we need to update anything in them?

### Comment by @mkholjuraev on March 14, 2024 at 06:14 PM UTC

Nope, the PR is ready to be reviewed and merged if everything is fine

### Comment by @mkholjuraev on March 15, 2024 at 07:58 AM UTC

You were right, somehow I did not push the commit for tests rewrite

### Comment by @mkholjuraev on April 03, 2024 at 09:06 AM UTC

@Fewwy I completely forgot about this PR. Would you mind having another look?

### Comment by @mkholjuraev on May 13, 2024 at 12:01 PM UTC

:tada: This PR is included in version 1.67.4 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.67.4)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @gkarat - Commented on April 15, 2024 at 11:49 AM UTC

Hey, everything works well. Though, I have noticed an increase in network requests number (mostly, /system_profile/... and /tags requests). Can you please take a look and tell if it is expected? If not, can we make sure the performance doesn't downgrade after the PR changes? 

Before:

![Screenshot 2024-04-15 at 13 46 19](https://github.com/RedHatInsights/patchman-ui/assets/31385370/fb6e8128-2c28-4975-9cd5-8e04a7dae92d)

![Screenshot 2024-04-15 at 13 46 31](https://github.com/RedHatInsights/patchman-ui/assets/31385370/406aecf6-e59c-4365-ad21-e5d71eb1147b)

After:

![Screenshot 2024-04-15 at 13 46 29](https://github.com/RedHatInsights/patchman-ui/assets/31385370/89bb7cfd-5fc0-4ca4-88d1-d7a02277dd97)

![Screenshot 2024-04-15 at 13 46 15](https://github.com/RedHatInsights/patchman-ui/assets/31385370/cdc82331-57c4-4001-a6c4-6fa43a7ffead)


### Review by @gkarat - Approved on April 17, 2024 at 02:27 PM UTC

No longer reproducible ^, I think it was not caused by your PR. Let's have it merged, thanks @mkholjuraev 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1175*
