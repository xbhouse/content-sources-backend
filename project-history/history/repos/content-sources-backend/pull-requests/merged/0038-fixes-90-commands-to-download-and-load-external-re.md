---
type: pull_request
number: 38
title: "Fixes 90: commands to download and load external repos"
state: merged
author: jlsherrill
created: 2022-06-17T13:48:03Z
updated: 2022-07-20T13:53:52Z
closed: 2022-07-20T13:53:45Z
merged: 2022-07-20T13:53:45Z
base_branch: main
head_branch: external
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/38
---

# Pull Request #38: Fixes 90: commands to download and load external repos

**Author**: @jlsherrill
**Created**: June 17, 2022 at 01:48 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `external`

## Description

This adds a couple commands:

```
Manage external repo information
  repos-download   Download external repo urls from Image Builder
  repos-import     Import External repo urls from Image Builders into the DB
```

so you can run:

```
rm ./pkg/external_repos/external_repos.json    #clear out the existing cache
make repos-download  
```
the file should be recreated exactly

```
make repos-import
```

check your db and see the repos import:
```
make db-cli-connect


 select * from repositories ;
```

This also adds a "public" field to repositories, to mark repos that are available to all accounts, regardless of presence of repository_configuration

---

## Discussion

### Comment by @jlsherrill on June 28, 2022 at 08:20 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-90

### Comment by @jlsherrill on June 28, 2022 at 08:21 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-90

### Comment by @jlsherrill on June 29, 2022 at 05:51 PM UTC

https://issues.redhat.com/browse/HMSCONTENT--

### Comment by @jlsherrill on July 15, 2022 at 03:09 PM UTC

@avisiedo this is ready for re-review :)

### Comment by @jlsherrill on July 18, 2022 at 06:43 PM UTC

Built and pushed image for 72105ccf4a30e12a3d78929b13656e90f54590b2

### Comment by @jlsherrill on July 19, 2022 at 05:39 PM UTC

I think i addressed everything

### Comment by @jlsherrill on July 19, 2022 at 05:43 PM UTC

for the onConflict, take a look at https://gorm.io/docs/create.html#upsert

### Comment by @jlsherrill on July 19, 2022 at 08:20 PM UTC

https://issues.redhat.com/browse/14695815

---

## Reviews

### Review by @swadeley - Commented on June 27, 2022 at 06:20 PM UTC

just some typos

### Review by @avisiedo - Commented on July 12, 2022 at 12:52 PM UTC

### Review by @avisiedo - Commented on July 12, 2022 at 12:52 PM UTC

### Review by @avisiedo - Commented on July 12, 2022 at 12:59 PM UTC

### Review by @jlsherrill - Commented on July 14, 2022 at 06:13 PM UTC

### Review by @avisiedo - Commented on July 19, 2022 at 10:14 AM UTC

### Review by @avisiedo - Commented on July 19, 2022 at 10:14 AM UTC

### Review by @avisiedo - Commented on July 19, 2022 at 10:22 AM UTC

### Review by @avisiedo - Commented on July 19, 2022 at 10:23 AM UTC

### Review by @avisiedo - Commented on July 19, 2022 at 11:51 AM UTC

### Review by @avisiedo - Commented on July 19, 2022 at 11:54 AM UTC

### Review by @avisiedo - Commented on July 19, 2022 at 12:17 PM UTC

### Review by @jlsherrill - Commented on July 19, 2022 at 12:55 PM UTC

### Review by @avisiedo - Commented on July 19, 2022 at 12:56 PM UTC

### Review by @avisiedo - Changes Requested on July 19, 2022 at 01:12 PM UTC

The majority small changes; only two changes are bigger; and I learned about `clause.OnConflict` in gorm that I didn't know.

After that it LGTM


### Review by @jlsherrill - Commented on July 19, 2022 at 05:37 PM UTC

### Review by @jlsherrill - Commented on July 19, 2022 at 05:40 PM UTC

### Review by @jlsherrill - Commented on July 19, 2022 at 05:42 PM UTC

### Review by @avisiedo - Commented on July 20, 2022 at 10:46 AM UTC

### Review by @avisiedo - Approved on July 20, 2022 at 10:52 AM UTC

### Review by @avisiedo - Commented on July 20, 2022 at 10:56 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/38*
