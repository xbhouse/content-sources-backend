---
type: pull_request
number: 1193
title: "tag security compliance image"
state: merged
author: wcater803
created: 2024-09-11T14:34:51Z
updated: 2024-09-11T14:51:00Z
closed: 2024-09-11T14:51:00Z
merged: 2024-09-11T14:51:00Z
base_branch: security-compliance
head_branch: security-compliance
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1193
---

# Pull Request #1193: tag security compliance image

**Author**: @wcater803
**Created**: September 11, 2024 at 02:34 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `security-compliance` ← **Head**: `security-compliance`

## Description

# Description

Adding an if-statement to build_deploy.sh to properly tag the Container Image generated when building from the Security-Compliance Branch.

```
if [[ "$GIT_BRANCH" == "origin/security-compliance" ]]; then
	export IMAGE_TAG="sc-$(date +%Y%m%d)-$(git rev-parse --short=7 HEAD)"
fi
```

# Checklist:

- [ ] The commit message has the Jira ticket linked
- [ ] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1193*
