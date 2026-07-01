---
type: pull_request
number: 72
title: "Fixes 126: Add openapi docs and generation"
state: merged
author: jlsherrill
created: 2022-08-04T16:08:50Z
updated: 2022-08-23T18:26:58Z
closed: 2022-08-22T19:30:33Z
merged: 2022-08-22T19:30:33Z
base_branch: main
head_branch: openapi_html
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/72
---

# Pull Request #72: Fixes 126: Add openapi docs and generation

**Author**: @jlsherrill
**Created**: August 04, 2022 at 04:08 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `openapi_html`

## Description

added a make task to generate markdown api docs
markdown can be checked in to main, served via https://content-services.github.io/content-sources-backend/

---

## Discussion

### Comment by @jlsherrill on August 04, 2022 at 04:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-126

### Comment by @jlsherrill on August 16, 2022 at 02:19 PM UTC

i'll look!

### Comment by @jlsherrill on August 16, 2022 at 04:04 PM UTC

I think i got them all, but let me know @rverdile if i missed any

### Comment by @ezr-ondrej on August 23, 2022 at 06:26 PM UTC

We're using redoc btw... that's quite a nice tool as well: https://redocly.github.io/redoc/?url=https://raw.githubusercontent.com/RHEnVision/provisioning-backend/main/api/openapi.gen.json

---

## Reviews

### Review by @avisiedo - Commented on August 05, 2022 at 01:05 PM UTC

### Review by @jlsherrill - Commented on August 05, 2022 at 03:27 PM UTC

### Review by @rverdile - Commented on August 09, 2022 at 07:50 PM UTC

You don't necessarily have to change this here, but should we be putting comments/descriptions for every field in the API?

For example, in RepositoryResponse, the description for `name` is blank. The meaning of name is intuitive, so maybe it doesn't need a description. But then there's a blank field, and maybe it's better to be consistent.  

I can go through and point out all the blanks if we care about that.

### Review by @rverdile - Commented on August 17, 2022 at 07:56 PM UTC

### Review by @rverdile - Commented on August 17, 2022 at 07:56 PM UTC

### Review by @rverdile - Commented on August 17, 2022 at 07:57 PM UTC

### Review by @rverdile - Commented on August 17, 2022 at 07:57 PM UTC

### Review by @rverdile - Commented on August 17, 2022 at 07:58 PM UTC

### Review by @rverdile - Commented on August 17, 2022 at 07:58 PM UTC

looks like you got most of them. but I found a couple more empty descriptions

### Review by @swadeley - Approved on August 22, 2022 at 01:15 PM UTC

ACK with some suggestions

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/72*
