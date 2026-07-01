---
type: pull_request
number: 23
title: "Fixes 3539: introspection fails if comps file not detected"
state: merged
author: xbhouse
created: 2024-02-06T22:16:33Z
updated: 2024-02-12T07:09:18Z
closed: 2024-02-09T15:40:06Z
merged: 2024-02-09T15:40:06Z
base_branch: master
head_branch: fix-comps-detection
labels: []
url: https://github.com/content-services/yummy/pull/23
---

# Pull Request #23: Fixes 3539: introspection fails if comps file not detected

**Author**: @xbhouse
**Created**: February 06, 2024 at 10:16 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-comps-detection`

## Description

## Summary

Handles case where a repository's comps has no file extension and /or are in either `group_gz` or `group`

## Testing steps

Package groups and environments can be parsed from these repositories: 

http://yum.theforeman.org/pulpcore/3.16/el7/x86_64/ 
http://yum.theforeman.org/pulpcore/3.16/el8/x86_64/

(if there are any other testing steps I can take or repositories to check please let me know) 

---

## Discussion

### Comment by @swadeley on February 12, 2024 at 07:09 AM UTC

https://issues.redhat.com/browse/HMS-3539

---

## Reviews

### Review by @rverdile - Commented on February 08, 2024 at 05:10 PM UTC

### Review by @rverdile - Commented on February 08, 2024 at 05:16 PM UTC

### Review by @xbhouse - Commented on February 08, 2024 at 05:17 PM UTC

### Review by @xbhouse - Commented on February 08, 2024 at 05:23 PM UTC

### Review by @rverdile - Commented on February 08, 2024 at 06:20 PM UTC

### Review by @rverdile - Approved on February 08, 2024 at 09:12 PM UTC

good job! 

---

*Archived from: https://github.com/content-services/yummy/pull/23*
