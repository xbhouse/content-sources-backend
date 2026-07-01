---
type: pull_request
number: 18
title: "handling case of no comps.xml in repo"
state: merged
author: xbhouse
created: 2023-11-08T15:02:26Z
updated: 2023-11-09T15:58:20Z
closed: 2023-11-09T15:58:20Z
merged: 2023-11-09T15:58:20Z
base_branch: master
head_branch: no_comps
labels: []
url: https://github.com/content-services/yummy/pull/18
---

# Pull Request #18: handling case of no comps.xml in repo

**Author**: @xbhouse
**Created**: November 08, 2023 at 03:02 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `no_comps`

## Description

- returns an error only if parsing the comps fails when there is a comps.xml in the repo
- added mock and test for this case

---

## Discussion

### Comment by @xbhouse on November 09, 2023 at 03:48 PM UTC

@jlsherrill thanks! do i need to create a tag before merging? i'm not really sure how to create a tag and release 

### Comment by @jlsherrill on November 09, 2023 at 03:55 PM UTC

ah, its super easy, just:

1.  merge the pr
2. locally, check out master/main, pull the latest changes
3. run `git tag v1.0.8`  (or whatever the version should be)
4. git push upstream v1.0.8    (replacing upstream with whatever your cotnent-services/yummy remote name is).  

Let me know if you have any questions or troubles!

---

## Reviews

### Review by @jlsherrill - Commented on November 09, 2023 at 01:31 PM UTC

### Review by @jlsherrill - Commented on November 09, 2023 at 01:31 PM UTC

### Review by @jlsherrill - Commented on November 09, 2023 at 01:32 PM UTC

### Review by @xbhouse - Commented on November 09, 2023 at 01:58 PM UTC

### Review by @jlsherrill - Commented on November 09, 2023 at 02:06 PM UTC

### Review by @xbhouse - Commented on November 09, 2023 at 02:42 PM UTC

### Review by @jlsherrill - Commented on November 09, 2023 at 02:53 PM UTC

### Review by @jlsherrill - Commented on November 09, 2023 at 03:00 PM UTC

### Review by @jlsherrill - Commented on November 09, 2023 at 03:03 PM UTC

### Review by @xbhouse - Commented on November 09, 2023 at 03:24 PM UTC

### Review by @xbhouse - Commented on November 09, 2023 at 03:29 PM UTC

### Review by @jlsherrill - Commented on November 09, 2023 at 03:29 PM UTC

### Review by @xbhouse - Commented on November 09, 2023 at 03:31 PM UTC

### Review by @xbhouse - Commented on November 09, 2023 at 03:38 PM UTC

### Review by @jlsherrill - Approved on November 09, 2023 at 03:43 PM UTC

ACK! 

---

*Archived from: https://github.com/content-services/yummy/pull/18*
