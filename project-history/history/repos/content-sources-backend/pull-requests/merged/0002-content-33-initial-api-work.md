---
type: pull_request
number: 2
title: "CONTENT-33: initial api work"
state: merged
author: jlsherrill
created: 2022-05-10T12:42:09Z
updated: 2022-05-24T16:06:13Z
closed: 2022-05-24T16:06:13Z
merged: 2022-05-24T16:06:13Z
base_branch: main
head_branch: api
labels: []
url: https://github.com/content-services/content-sources-backend/pull/2
---

# Pull Request #2: CONTENT-33: initial api work

**Author**: @jlsherrill
**Created**: May 10, 2022 at 12:42 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `api`

## Description

*No description provided*

---

## Discussion

### Comment by @jlsherrill on May 10, 2022 at 12:45 PM UTC

things left to do:

- [x]  integrate with db models, add additional fields
- [x] add more tests
- [x] resolve TODOs
- [x] integrate api prefix
- [x] add openapispec
- [x] move to app
- [x] service openapi docs

### Comment by @rverdile on May 18, 2022 at 08:40 PM UTC

I've noticed you tend to be explicit with the int size (int64). Do you think this is something we should be doing everywhere?

### Comment by @jlsherrill on May 18, 2022 at 08:53 PM UTC

I'm not 100% sure why i originally went that way, but i do want to describe some quirks around this.  Limit() and Offset() in gorm seem to take in ints, while its Count() seems to return an int64.    It looks like x86_64 systems use 64bits for 'int' so in practice it really doesn't matter, however its possible that converting from int64 to int could lead to a smaller number.  

That said, i think it makes sense to use int and only convert from int64 in the one case we need (the Count() method).  I'm gonna update to do that. 

### Comment by @rverdile on May 23, 2022 at 03:56 PM UTC

How do I install the `swag` command? I can't run `make openapi` without it. Might also be good to add this to the README.

### Comment by @jlsherrill on May 23, 2022 at 09:02 PM UTC

I think thats my last major push.   

---

## Reviews

### Review by @jlsherrill - Commented on May 19, 2022 at 07:07 PM UTC

### Review by @rverdile - Commented on May 19, 2022 at 08:50 PM UTC

Mostly looked at repositories so far.  Haven't looked closely at the seeding or makefile changes yet.

### Review by @jlsherrill - Commented on May 23, 2022 at 01:17 PM UTC

### Review by @rverdile - Commented on May 23, 2022 at 06:56 PM UTC

Another small comment. Everything else so far looks good to me. Is there more pushes on the way?

### Review by @jlsherrill - Commented on May 23, 2022 at 09:00 PM UTC

### Review by @jlsherrill - Commented on May 23, 2022 at 09:02 PM UTC

### Review by @rverdile - Commented on May 24, 2022 at 01:53 PM UTC

### Review by @rverdile - Commented on May 24, 2022 at 02:10 PM UTC

### Review by @rverdile - Commented on May 24, 2022 at 02:24 PM UTC

### Review by @jlsherrill - Commented on May 24, 2022 at 02:30 PM UTC

### Review by @rverdile - Approved on May 24, 2022 at 03:00 PM UTC

pending the small changes, and if there's no big changes, this looks good to me.

### Review by @rverdile - Commented on May 24, 2022 at 03:43 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/2*
