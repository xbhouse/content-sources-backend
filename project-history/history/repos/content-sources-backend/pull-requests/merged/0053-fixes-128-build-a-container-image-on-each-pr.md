---
type: pull_request
number: 53
title: "Fixes 128: build a container image on each pr"
state: merged
author: jlsherrill
created: 2022-07-15T21:03:29Z
updated: 2022-07-19T20:20:30Z
closed: 2022-07-18T13:38:51Z
merged: 2022-07-18T13:38:51Z
base_branch: main
head_branch: pr_image_build
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/53
---

# Pull Request #53: Fixes 128: build a container image on each pr

**Author**: @jlsherrill
**Created**: July 15, 2022 at 09:03 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `pr_image_build`

## Description

*No description provided*

---

## Discussion

### Comment by @jlsherrill on July 15, 2022 at 09:08 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-128

### Comment by @mshriver on July 18, 2022 at 12:03 PM UTC

For the record, the scripts added by #47 will do this as well, building into the cloudservices org via Jenkins. But it requires appsre onboarding and app-interface acceptance, which has been slow and blocking. 

### Comment by @jlsherrill on July 18, 2022 at 01:15 PM UTC

This should work when its merged, i tested on my private account here: https://github.com/jlsherrill/content-sources-backend/runs/7389989851?check_suite_focus=true

### Comment by @jlsherrill on July 18, 2022 at 01:16 PM UTC

and you can see the results here:  https://quay.io/repository/jlsherri/content-sources?tab=tags   pr-1

### Comment by @jlsherrill on July 19, 2022 at 08:20 PM UTC

https://issues.redhat.com/browse/14751782

---

## Reviews

### Review by @swadeley - Approved on July 18, 2022 at 01:38 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/53*
