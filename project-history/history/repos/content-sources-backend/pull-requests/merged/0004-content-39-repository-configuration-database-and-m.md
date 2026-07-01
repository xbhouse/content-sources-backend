---
type: pull_request
number: 4
title: "CONTENT-39: repository configuration database and model"
state: merged
author: rverdile
created: 2022-05-12T19:08:30Z
updated: 2022-05-19T18:45:06Z
closed: 2022-05-19T18:45:06Z
merged: 2022-05-19T18:45:06Z
base_branch: main
head_branch: my-changes
labels: []
url: https://github.com/content-services/content-sources-backend/pull/4
---

# Pull Request #4: CONTENT-39: repository configuration database and model

**Author**: @rverdile
**Created**: May 12, 2022 at 07:08 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `my-changes`

## Description

If you want to try to use this, you have to create a psql database and set the environment variables for connecting (see Makefile for names).

I made the DB connection a global variable in main just as an example. I don't know how we actually want to handle that, but I'm assuming we don't want to open new connection every time we access the database.


To-do List:

- [x] Add force rollback for failed migrations
- [x] Add template for migrations, so creating a new automatically includes transaction syntax
- [x] Remove repository auth configuration ID
- [x] Remove "name" flag in favor of dbmigrate new NAME
- [x] Move TestMain to its own file
- [ ] [Skip] Integrate with clowder for database configuration
- [x] Handle local env variables for database configuration

---

## Reviews

### Review by @jlsherrill - Commented on May 16, 2022 at 01:02 PM UTC

### Review by @jlsherrill - Commented on May 16, 2022 at 01:03 PM UTC

### Review by @jlsherrill - Commented on May 16, 2022 at 01:09 PM UTC

### Review by @rverdile - Commented on May 16, 2022 at 01:55 PM UTC

### Review by @jlsherrill - Commented on May 18, 2022 at 05:12 PM UTC

### Review by @jlsherrill - Commented on May 18, 2022 at 05:15 PM UTC

### Review by @jlsherrill - Commented on May 18, 2022 at 05:18 PM UTC

### Review by @rverdile - Commented on May 18, 2022 at 05:47 PM UTC

### Review by @jlsherrill - Commented on May 18, 2022 at 05:59 PM UTC

### Review by @jlsherrill - Commented on May 18, 2022 at 06:00 PM UTC

### Review by @rverdile - Commented on May 18, 2022 at 06:02 PM UTC

### Review by @rverdile - Commented on May 18, 2022 at 06:06 PM UTC

### Review by @jlsherrill - Commented on May 19, 2022 at 01:25 PM UTC

### Review by @jlsherrill - Commented on May 19, 2022 at 01:26 PM UTC

### Review by @jlsherrill - Commented on May 19, 2022 at 01:37 PM UTC

### Review by @rverdile - Commented on May 19, 2022 at 02:14 PM UTC

### Review by @jlsherrill - Commented on May 19, 2022 at 02:26 PM UTC

### Review by @jlsherrill - Approved on May 19, 2022 at 06:00 PM UTC

ACK

![dance](https://media.giphy.com/media/hxc32veg6tbqg/giphy.gif)





---

*Archived from: https://github.com/content-services/content-sources-backend/pull/4*
