---
type: pull_request
number: 55
title: "Fixes 49: add introspection command"
state: merged
author: jlsherrill
created: 2022-07-20T14:26:05Z
updated: 2022-07-29T15:58:31Z
closed: 2022-07-29T15:58:31Z
merged: 2022-07-29T15:58:31Z
base_branch: main
head_branch: introspect_2
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/55
---

# Pull Request #55: Fixes 49: add introspection command

**Author**: @jlsherrill
**Created**: July 20, 2022 at 02:26 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `introspect_2`

## Description

This still needs a bit of work, 

* around having the ssl certificates be able to be passed in, and likely set via the config
* possibly look at introspecting all repos, not just public ones
* moar tests

---

## Discussion

### Comment by @jlsherrill on July 20, 2022 at 03:44 PM UTC

Built and pushed image for 479c12510d8e2000e4ec21be27c5206cf673327f

### Comment by @jlsherrill on July 22, 2022 at 12:45 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-49

### Comment by @jlsherrill on July 22, 2022 at 08:20 PM UTC

Built and pushed image for 345225dc5963fa17b30df11682ae83b3000df242

### Comment by @jlsherrill on July 25, 2022 at 08:06 PM UTC

A lot of these methods i think would benefit from a comment, i'm going to do that myself since i wrote a lot of them and push that up.

### Comment by @avisiedo on July 27, 2022 at 10:56 AM UTC

I am adding more tests and checking better some scenarios

### Comment by @lechuk47 on July 27, 2022 at 01:06 PM UTC

/retest

### Comment by @mshriver on July 27, 2022 at 06:51 PM UTC

test

### Comment by @mshriver on July 27, 2022 at 06:51 PM UTC

/test


### Comment by @avisiedo on July 27, 2022 at 08:22 PM UTC

I have to import the repositories before run the command for a specific url. After I select a no redhat repository and I use the command with that url such as:

```sh
make repos-download repos-import
./release/external-repos introspect https://packages.cloud.google.com/yum/repos/google-compute-engine-el9-x86_64-stable
```


### Comment by @jlsherrill on July 28, 2022 at 07:09 PM UTC

I'm good with this i think, will leave the final ack to you @rverdile !

---

## Reviews

### Review by @jlsherrill - Commented on July 25, 2022 at 07:35 PM UTC

### Review by @jlsherrill - Commented on July 25, 2022 at 07:41 PM UTC

### Review by @jlsherrill - Commented on July 25, 2022 at 08:01 PM UTC

### Review by @jlsherrill - Commented on July 25, 2022 at 08:08 PM UTC

### Review by @jlsherrill - Commented on July 26, 2022 at 01:35 PM UTC

### Review by @jlsherrill - Commented on July 26, 2022 at 01:42 PM UTC

### Review by @jlsherrill - Commented on July 26, 2022 at 01:56 PM UTC

### Review by @avisiedo - Commented on July 26, 2022 at 07:07 PM UTC

### Review by @jlsherrill - Commented on July 26, 2022 at 07:30 PM UTC

### Review by @jlsherrill - Commented on July 26, 2022 at 07:31 PM UTC

### Review by @avisiedo - Commented on July 26, 2022 at 07:35 PM UTC

### Review by @jlsherrill - Commented on July 26, 2022 at 07:45 PM UTC

### Review by @jlsherrill - Commented on July 26, 2022 at 07:46 PM UTC

### Review by @jlsherrill - Commented on July 26, 2022 at 07:47 PM UTC

### Review by @jlsherrill - Commented on July 26, 2022 at 07:48 PM UTC

### Review by @jlsherrill - Commented on July 26, 2022 at 08:15 PM UTC

### Review by @jlsherrill - Commented on July 26, 2022 at 08:18 PM UTC

### Review by @jlsherrill - Commented on July 26, 2022 at 08:29 PM UTC

### Review by @rverdile - Commented on July 26, 2022 at 08:35 PM UTC

### Review by @avisiedo - Commented on July 27, 2022 at 04:34 AM UTC

### Review by @avisiedo - Commented on July 27, 2022 at 06:12 AM UTC

### Review by @avisiedo - Commented on July 27, 2022 at 07:43 AM UTC

### Review by @rverdile - Commented on July 27, 2022 at 02:33 PM UTC

### Review by @rverdile - Commented on July 27, 2022 at 02:44 PM UTC

Not sure if I'm using this incorrectly, but the command:  `./release/external-repos introspect https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/` 

fails with `{"level":"panic","error":"record not found","time":"2022-07-27T10:26:06-04:00","message":"Failed to introspect repositories"}`. 

Was able to introspect this with `introspect-all` but not `introspect`

Also, I think the error messages in within these introspect functions should be returning an error message along with the error, so it's easier to find where the error is coming from. I'll comment some examples.

### Review by @rverdile - Commented on July 27, 2022 at 02:55 PM UTC

### Review by @avisiedo - Commented on July 28, 2022 at 06:06 AM UTC

### Review by @jlsherrill - Commented on July 28, 2022 at 12:50 PM UTC

### Review by @jlsherrill - Commented on July 28, 2022 at 01:42 PM UTC

### Review by @jlsherrill - Commented on July 28, 2022 at 02:24 PM UTC

### Review by @jlsherrill - Commented on July 28, 2022 at 02:48 PM UTC

### Review by @jlsherrill - Commented on July 28, 2022 at 03:12 PM UTC

### Review by @jlsherrill - Commented on July 28, 2022 at 03:20 PM UTC

### Review by @avisiedo - Commented on July 28, 2022 at 06:09 PM UTC

### Review by @rverdile - Commented on July 28, 2022 at 07:37 PM UTC

### Review by @rverdile - Commented on July 28, 2022 at 07:41 PM UTC

### Review by @rverdile - Commented on July 28, 2022 at 07:46 PM UTC

Left two small comments, but this is looking good. Commands are working for me and I see all the additional error messages were added.

### Review by @avisiedo - Commented on July 28, 2022 at 09:16 PM UTC

### Review by @rverdile - Approved on July 29, 2022 at 01:27 PM UTC

ack

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/55*
