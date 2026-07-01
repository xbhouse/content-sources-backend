---
type: pull_request
number: 86
title: "Fixes 143: Remove null fields from api responses"
state: merged
author: rverdile
created: 2022-08-24T19:11:26Z
updated: 2022-08-26T19:30:24Z
closed: 2022-08-26T19:21:14Z
merged: 2022-08-26T19:21:14Z
base_branch: main
head_branch: bulk-create-error
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/86
---

# Pull Request #86: Fixes 143: Remove null fields from api responses

**Author**: @rverdile
**Created**: August 24, 2022 at 07:11 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `bulk-create-error`

## Description

There were three fields I could find that were returning null:
1. The bulk create error field, when there's no error.
~2. The bulk create repository field, when there is an error.~ leaving out because it's not causing a CI failure and it's not a simple fix
3. The distribution versions field after a successful creation, if no versions were specified.

I changed these to:
1. empty string
~2. empty object~
3. empty array

---

## Discussion

### Comment by @jlsherrill on August 24, 2022 at 07:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-143

### Comment by @swadeley on August 25, 2022 at 10:23 AM UTC

Hi

Rebase please.

the ci.int.devshift check for this PR only failed on `test_create_update_distro_version_is_unique`, which is fixed by https://github.com/content-services/content-sources-backend/pull/85, so thats a good.

### Comment by @mshriver on August 25, 2022 at 10:49 AM UTC

Because this contains openapi spec changes, we won't see the test pass in CI - the client generation is not dynamic. @swadeley can pull the openapi spec into our plugin ahead of merging this once we're confident in the content (above discussions about scope of nil values resolved).

### Comment by @jlsherrill on August 25, 2022 at 04:07 PM UTC

this needs  a rebase.

### Comment by @mshriver on August 25, 2022 at 11:35 PM UTC

/retest

### Comment by @rverdile on August 26, 2022 at 01:45 PM UTC

failure looks unrelated: https://github.com/swaggo/swag/issues/1309

---

## Reviews

### Review by @rverdile - Commented on August 24, 2022 at 07:17 PM UTC

### Review by @rverdile - Commented on August 24, 2022 at 07:18 PM UTC

### Review by @jlsherrill - Commented on August 24, 2022 at 07:41 PM UTC

### Review by @jlsherrill - Commented on August 24, 2022 at 07:50 PM UTC

### Review by @rverdile - Commented on August 24, 2022 at 07:51 PM UTC

### Review by @rverdile - Commented on August 24, 2022 at 07:54 PM UTC

### Review by @jlsherrill - Commented on August 24, 2022 at 08:31 PM UTC

### Review by @jlsherrill - Commented on August 24, 2022 at 09:05 PM UTC

### Review by @jlsherrill - Commented on August 24, 2022 at 09:08 PM UTC

### Review by @rverdile - Commented on August 25, 2022 at 01:31 PM UTC

### Review by @rverdile - Commented on August 25, 2022 at 02:52 PM UTC

### Review by @jlsherrill - Commented on August 25, 2022 at 05:18 PM UTC

### Review by @rverdile - Commented on August 25, 2022 at 06:26 PM UTC

### Review by @rverdile - Commented on August 25, 2022 at 06:54 PM UTC

### Review by @rverdile - Commented on August 26, 2022 at 12:17 PM UTC

### Review by @jlsherrill - Approved on August 26, 2022 at 02:52 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/86*
