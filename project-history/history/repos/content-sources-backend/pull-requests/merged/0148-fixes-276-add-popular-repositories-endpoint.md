---
type: pull_request
number: 148
title: "Fixes 276: Add \"Popular repositories\" endpoint"
state: merged
author: Andrewgdewar
created: 2022-11-24T03:16:28Z
updated: 2022-12-14T14:06:12Z
closed: 2022-12-14T02:43:30Z
merged: 2022-12-14T02:43:30Z
base_branch: main
head_branch: CONTENT-276
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/148
---

# Pull Request #148: Fixes 276: Add "Popular repositories" endpoint

**Author**: @Andrewgdewar
**Created**: November 24, 2022 at 03:16 AM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `CONTENT-276`

## Description

Adds the "/common_repositories/" endpoint with 4 repositories in the common_repositories.cfg for now.

- [x] Still needs tests

---

## Discussion

### Comment by @jlsherrill on November 24, 2022 at 03:30 AM UTC

https://issues.redhat.com/browse/HMSCONTENT-276

### Comment by @avisiedo on November 24, 2022 at 01:57 PM UTC

I suggest to run 'make openapi' to re-generate the `api/openapi.json` and `api/docs.go` files.

### Comment by @jlsherrill on November 28, 2022 at 09:09 PM UTC

overall this is looking great!

### Comment by @jlsherrill on December 09, 2022 at 06:08 PM UTC

Oh and it'd probably help qe to write a few curl commands to test

### Comment by @jlsherrill on December 09, 2022 at 06:13 PM UTC

tested everything else its good to go! just one test addition request

### Comment by @swadeley on December 13, 2022 at 05:18 PM UTC

/retest

---

## Reviews

### Review by @swadeley - Commented on November 24, 2022 at 08:50 AM UTC

### Review by @avisiedo - Dismissed on November 24, 2022 at 02:18 PM UTC

Thanks for the changes :+1: a few suggestions in the pr review.

I could not replay the linter hints from the macos environment; I will try in x86 environment to see if I get something different; very likely I have something wrong in my environment.

If the data is going to be totally static, maybe embed the content in a variable from an external yaml file could be nice (https://pkg.go.dev/embed); on this point, in regarding to the GpgKey field, it could be easier to write the gpg key in a yaml file than in a json file (this reference could be useful: https://yaml-multiline.info)

### Review by @Andrewgdewar - Commented on November 24, 2022 at 09:41 PM UTC

### Review by @Andrewgdewar - Commented on November 24, 2022 at 09:42 PM UTC

### Review by @jlsherrill - Commented on November 28, 2022 at 08:49 PM UTC

### Review by @jlsherrill - Commented on November 28, 2022 at 08:55 PM UTC

### Review by @jlsherrill - Commented on November 28, 2022 at 08:55 PM UTC

### Review by @jlsherrill - Commented on November 28, 2022 at 08:58 PM UTC

### Review by @jlsherrill - Commented on November 28, 2022 at 08:59 PM UTC

### Review by @jlsherrill - Commented on November 28, 2022 at 09:08 PM UTC

### Review by @rverdile - Commented on November 29, 2022 at 04:58 PM UTC

### Review by @rverdile - Commented on November 29, 2022 at 04:59 PM UTC

### Review by @rverdile - Commented on November 29, 2022 at 04:59 PM UTC

### Review by @Andrewgdewar - Commented on December 01, 2022 at 08:44 PM UTC

### Review by @Andrewgdewar - Commented on December 01, 2022 at 09:02 PM UTC

### Review by @Andrewgdewar - Commented on December 02, 2022 at 04:10 PM UTC

### Review by @jlsherrill - Commented on December 05, 2022 at 05:44 PM UTC

### Review by @jlsherrill - Commented on December 05, 2022 at 05:46 PM UTC

### Review by @jlsherrill - Commented on December 05, 2022 at 05:47 PM UTC

### Review by @Andrewgdewar - Commented on December 05, 2022 at 10:06 PM UTC

### Review by @Andrewgdewar - Commented on December 05, 2022 at 10:07 PM UTC

### Review by @Andrewgdewar - Commented on December 05, 2022 at 10:14 PM UTC

### Review by @swadeley - Commented on December 08, 2022 at 09:44 AM UTC

### Review by @jlsherrill - Commented on December 09, 2022 at 05:44 PM UTC

### Review by @Andrewgdewar - Commented on December 09, 2022 at 09:39 PM UTC

### Review by @jlsherrill - Approved on December 09, 2022 at 09:53 PM UTC

### Review by @mshriver - Commented on December 13, 2022 at 06:19 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/148*
