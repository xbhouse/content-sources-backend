---
type: pull_request
number: 59
title: "Fixes 94: repo validation api"
state: merged
author: jlsherrill
created: 2022-07-27T12:20:13Z
updated: 2022-08-08T17:18:02Z
closed: 2022-08-08T17:18:02Z
merged: 2022-08-08T17:18:02Z
base_branch: main
head_branch: validation
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/59
---

# Pull Request #59: Fixes 94: repo validation api

**Author**: @jlsherrill
**Created**: July 27, 2022 at 12:20 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `validation`

## Description

*No description provided*

---

## Discussion

### Comment by @jlsherrill on July 27, 2022 at 12:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-94

### Comment by @jlsherrill on July 27, 2022 at 06:02 PM UTC

example usage:
```
 curl  -v -X POST  -H "`./scripts/header.sh 1 1`"  http://localhost:8000/api/content_sources/v1/repository_parameters/validate/  -H "Content-Type: application/json"    -d '[{"name": "foo", "url": "https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/" }]'
```

If you want to test a timeout, you can run this in one terminal:
```
nc -lkp 4000
```

and then try to point a repo to it:
```
curl  -v -X POST  -H "`./scripts/header.sh 1 1`"  http://localhost:8000/api/content_sources/v1/repository_parameters/validate/  -H "Content-Type: application/json"    -d '[{"name": "foo", "url": "https://localhost:4000/fake/" }]'
```

### Comment by @mshriver on July 29, 2022 at 09:35 PM UTC

[test]

### Comment by @rverdile on August 04, 2022 at 07:12 PM UTC

Sure I'll look

---

## Reviews

### Review by @Andrewgdewar - Approved on August 04, 2022 at 06:55 PM UTC

This is a "functional" ACK, I have this running smoothly with [my front-end branch](https://github.com/content-services/content-sources-frontend/pull/9) 
If someone with a more discerning eye would like to take a look at the code (which looks fine to me), please do. @rverdile  😉 

### Review by @rverdile - Commented on August 04, 2022 at 09:21 PM UTC

### Review by @rverdile - Commented on August 04, 2022 at 09:23 PM UTC

### Review by @rverdile - Commented on August 04, 2022 at 09:33 PM UTC

### Review by @rverdile - Commented on August 04, 2022 at 09:36 PM UTC

### Review by @jlsherrill - Commented on August 05, 2022 at 03:37 PM UTC

### Review by @jlsherrill - Commented on August 05, 2022 at 04:03 PM UTC

### Review by @jlsherrill - Commented on August 05, 2022 at 04:04 PM UTC

### Review by @rverdile - Commented on August 05, 2022 at 04:58 PM UTC

### Review by @rverdile - Approved on August 05, 2022 at 05:06 PM UTC

nice! code looks good

you can decide if you want RepoMD or RepoMd or Repomd :). I just wanted to mention it, but it's not a big deal.

### Review by @swadeley - Commented on August 08, 2022 at 08:29 AM UTC

### Review by @swadeley - Approved on August 08, 2022 at 05:16 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/59*
