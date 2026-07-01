---
type: pull_request
number: 1174
title: "change(RemediationButton):RHINENG-3123 text for no advisory systems"
state: merged
author: darkeriossss
created: 2024-03-14T11:45:02Z
updated: 2024-04-03T13:00:31Z
closed: 2024-04-03T11:19:18Z
merged: 2024-04-03T11:19:18Z
base_branch: master
head_branch: no-avisories-remediation
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1174
---

# Pull Request #1174: change(RemediationButton):RHINENG-3123 text for no advisory systems

**Author**: @darkeriossss
**Created**: March 14, 2024 at 11:45 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `no-avisories-remediation`

## Description

Change wording in NoDataModal when a system with no advisories is remediated

Associated Jira ticket: RHINENG-3123

# How to test 

1. Run this PR with https://github.com/RedHatInsights/insights-remediations-frontend/pull/408
2. Go to Patch systems page
3. Select system with no advisories available
4. The modal should have changed wording

# Before the change
![image](https://github.com/RedHatInsights/insights-remediations-frontend/assets/62351699/3c5b4fcc-992e-43bf-931a-476e1f9777e6)

# After the change
![image](https://github.com/RedHatInsights/patchman-ui/assets/62351699/16c8fadd-2f1d-4c29-94ef-f2326ee36230)

# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [x] Screenshots before and after the change are added
- [x] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @darkeriossss on March 14, 2024 at 12:18 PM UTC

/retest

### Comment by @mkholjuraev on March 25, 2024 at 12:45 PM UTC

@xmicha82 can we merge this PR as the remediation one has been already merged? It would help us to move the card to QA verification, creating more dev-capacity. 

### Comment by @johnsonm325 on March 28, 2024 at 04:02 PM UTC

@xmicha82 pinging you again on this.

### Comment by @mkholjuraev on April 02, 2024 at 11:52 AM UTC

/retest

### Comment by @mkholjuraev on April 03, 2024 at 01:00 PM UTC

:tada: This PR is included in version 1.67.3 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.67.3)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on March 15, 2024 at 11:36 AM UTC

LGTM! tested with remediation PR

### Review by @darkeriossss - Commented on March 19, 2024 at 01:08 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1174*
