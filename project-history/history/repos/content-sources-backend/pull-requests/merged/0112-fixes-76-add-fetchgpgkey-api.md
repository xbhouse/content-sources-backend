---
type: pull_request
number: 112
title: "Fixes 76: Add fetchGpgKey API"
state: merged
author: Andrewgdewar
created: 2022-09-21T20:56:20Z
updated: 2022-10-13T15:00:28Z
closed: 2022-10-13T14:31:32Z
merged: 2022-10-13T14:31:32Z
base_branch: main
head_branch: CONTENT-66
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/112
---

# Pull Request #112: Fixes 76: Add fetchGpgKey API

**Author**: @Andrewgdewar
**Created**: September 21, 2022 at 08:56 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `CONTENT-66`

## Description

*No description provided*

---

## Discussion

### Comment by @jlsherrill on September 21, 2022 at 09:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-66

### Comment by @rverdile on September 23, 2022 at 11:57 AM UTC

this should be "Fixes 76", 66 is the story

### Comment by @jlsherrill on September 23, 2022 at 02:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-76

### Comment by @Andrewgdewar on September 28, 2022 at 09:44 PM UTC

Suggestions for a meaningful test for this without actually calling an external URL?  
@rverdile @jlsherrill  

### Comment by @jlsherrill on September 29, 2022 at 03:56 PM UTC

here's some examples:
http://yum.theforeman.org/RPM-GPG-KEY-foreman
http://yum.theforeman.org/releases/1.21/RPM-GPG-KEY-foreman
https://packages.microsoft.com/keys/microsoft.asc

sorry, i misread your comment a bit.  So here:  https://github.com/content-services/content-sources-backend/blob/main/pkg/dao/external_resource_test.go#L35-L67

we do something similar, where we test the dao layer fetching a repomd.xml file.  You could include a gpg key in the source file as a string an serve that out in a test

### Comment by @jlsherrill on October 04, 2022 at 07:00 PM UTC

small wording comment, otherwise looks good!

### Comment by @rverdile on October 07, 2022 at 03:08 PM UTC

I wonder why this didn't get marked "Ready for QE"?

---

## Reviews

### Review by @swadeley - Commented on September 22, 2022 at 07:47 AM UTC

### Review by @Andrewgdewar - Commented on September 22, 2022 at 08:07 PM UTC

### Review by @Andrewgdewar - Commented on September 22, 2022 at 08:07 PM UTC

### Review by @rverdile - Commented on September 22, 2022 at 08:10 PM UTC

### Review by @Andrewgdewar - Commented on September 22, 2022 at 08:15 PM UTC

### Review by @jlsherrill - Commented on September 22, 2022 at 08:24 PM UTC

### Review by @jlsherrill - Commented on September 22, 2022 at 08:27 PM UTC

### Review by @jlsherrill - Commented on September 22, 2022 at 08:28 PM UTC

### Review by @Andrewgdewar - Commented on September 22, 2022 at 08:40 PM UTC

### Review by @Andrewgdewar - Commented on September 22, 2022 at 08:42 PM UTC

### Review by @Andrewgdewar - Commented on September 22, 2022 at 08:42 PM UTC

### Review by @jlsherrill - Commented on September 23, 2022 at 05:59 PM UTC

### Review by @jlsherrill - Commented on September 23, 2022 at 06:02 PM UTC

### Review by @jlsherrill - Commented on September 23, 2022 at 06:02 PM UTC

### Review by @jlsherrill - Commented on September 23, 2022 at 08:42 PM UTC

### Review by @jlsherrill - Commented on September 23, 2022 at 08:47 PM UTC

### Review by @Andrewgdewar - Commented on September 23, 2022 at 09:17 PM UTC

### Review by @Andrewgdewar - Commented on September 26, 2022 at 05:02 PM UTC

### Review by @rverdile - Commented on September 27, 2022 at 03:35 PM UTC

### Review by @rverdile - Commented on September 27, 2022 at 03:42 PM UTC

### Review by @Andrewgdewar - Commented on September 27, 2022 at 06:21 PM UTC

### Review by @Andrewgdewar - Commented on September 27, 2022 at 06:22 PM UTC

### Review by @Andrewgdewar - Commented on September 27, 2022 at 07:05 PM UTC

### Review by @rverdile - Commented on September 28, 2022 at 02:52 PM UTC

### Review by @jlsherrill - Commented on September 28, 2022 at 04:47 PM UTC

### Review by @Andrewgdewar - Commented on September 28, 2022 at 09:34 PM UTC

### Review by @Andrewgdewar - Commented on September 28, 2022 at 09:35 PM UTC

### Review by @Andrewgdewar - Commented on September 28, 2022 at 09:42 PM UTC

### Review by @jlsherrill - Commented on September 30, 2022 at 04:00 PM UTC

### Review by @jlsherrill - Commented on September 30, 2022 at 04:03 PM UTC

### Review by @jlsherrill - Commented on October 04, 2022 at 06:55 PM UTC

### Review by @jlsherrill - Commented on October 04, 2022 at 06:55 PM UTC

### Review by @jlsherrill - Commented on October 04, 2022 at 07:24 PM UTC

### Review by @jlsherrill - Commented on October 04, 2022 at 07:24 PM UTC

### Review by @Andrewgdewar - Commented on October 04, 2022 at 08:01 PM UTC

### Review by @swadeley - Commented on October 05, 2022 at 01:02 PM UTC

### Review by @jlsherrill - Commented on October 05, 2022 at 01:20 PM UTC

### Review by @jlsherrill - Commented on October 06, 2022 at 07:21 PM UTC

### Review by @jlsherrill - Commented on October 06, 2022 at 07:22 PM UTC

### Review by @jlsherrill - Approved on October 06, 2022 at 11:20 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/112*
