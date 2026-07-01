---
type: pull_request
number: 36
title: "Update Dependencies"
state: merged
author: jlsherrill
created: 2022-06-15T18:57:11Z
updated: 2022-06-22T15:07:27Z
closed: 2022-06-22T15:07:24Z
merged: 2022-06-22T15:07:24Z
base_branch: main
head_branch: dep_update
labels: []
url: https://github.com/content-services/content-sources-backend/pull/36
---

# Pull Request #36: Update Dependencies

**Author**: @jlsherrill
**Created**: June 15, 2022 at 06:57 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `dep_update`

## Description

*No description provided*

---

## Discussion

### Comment by @Andrewgdewar on June 17, 2022 at 02:52 PM UTC

Did you delete your go.sum before updating? I find it cleans out old values best that way. Not sure if it's best practice.

### Comment by @jlsherrill on June 21, 2022 at 01:24 PM UTC

good call, updated!

---

## Reviews

### Review by @Andrewgdewar - Commented on June 17, 2022 at 04:14 PM UTC

After pulling this into my branch and "go get ./..." and "go mod tidy" I see a bit of a different output. 
Did you run tidy on this?

### Review by @Andrewgdewar - Approved on June 22, 2022 at 03:03 PM UTC

No changes when I updated! ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/36*
