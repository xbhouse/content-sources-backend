---
type: pull_request
number: 80
title: "Fixes 144: remove unused rpm entries more efficiently"
state: merged
author: jlsherrill
created: 2022-08-16T20:34:59Z
updated: 2022-09-19T13:39:28Z
closed: 2022-09-19T13:39:28Z
merged: 2022-09-19T13:39:28Z
base_branch: main
head_branch: 144
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/80
---

# Pull Request #80: Fixes 144: remove unused rpm entries more efficiently

**Author**: @jlsherrill
**Created**: August 16, 2022 at 08:34 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `144`

## Description

only once after each introspect-all.  Also makes a test a tad more readable IMO

---

## Discussion

### Comment by @jlsherrill on August 16, 2022 at 09:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-144

### Comment by @jlsherrill on August 25, 2022 at 04:05 PM UTC

you could test this by creating some repo, introspecting all repos, get rpm count from the db `select count(*) from rpms`, delete the repo, run introspect-all again, check the count again.  It should decrease.

### Comment by @jlsherrill on August 26, 2022 at 05:06 PM UTC

> i think you left the output directory in here by mistake

![giphy](https://user-images.githubusercontent.com/395077/186956548-84a24669-4c1d-460e-a7e7-3dede2fa2e89.gif)



### Comment by @jlsherrill on August 26, 2022 at 05:08 PM UTC

my plans for secretly inlining a python api client are ruined!

### Comment by @avisiedo on September 15, 2022 at 10:09 AM UTC

Reminder about the change to cover the scenario where no orphaned need to be deleted.

It needs rebase

### Comment by @jlsherrill on September 15, 2022 at 02:35 PM UTC

[test]

### Comment by @jlsherrill on September 15, 2022 at 02:35 PM UTC

/retest

### Comment by @jlsherrill on September 15, 2022 at 02:59 PM UTC

@avisiedo should be good now

### Comment by @avisiedo on September 16, 2022 at 10:33 PM UTC

lgtm 👍 

---

## Reviews

### Review by @rverdile - Commented on August 26, 2022 at 03:56 PM UTC

i think you left the output directory in here by mistake

### Review by @avisiedo - Changes Requested on August 29, 2022 at 05:14 PM UTC

a couple of comments only; I have reviewed and the rest lgtm; I liked the the code deduplication :) 

### Review by @jlsherrill - Commented on September 07, 2022 at 03:53 PM UTC

### Review by @avisiedo - Commented on September 08, 2022 at 05:51 PM UTC

### Review by @avisiedo - Commented on September 08, 2022 at 05:54 PM UTC

### Review by @avisiedo - Changes Requested on September 08, 2022 at 06:02 PM UTC

I have reviewed the two comments; the one related with delete records to avoid collision with rpm checksums does not make sense, sorry for the noise; the second change is to cover when no orphaned records need to be deleted.

### Review by @jlsherrill - Commented on September 12, 2022 at 06:03 PM UTC

### Review by @avisiedo - Approved on September 16, 2022 at 10:33 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/80*
