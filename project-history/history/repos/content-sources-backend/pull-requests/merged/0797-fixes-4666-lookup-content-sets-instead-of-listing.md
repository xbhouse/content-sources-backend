---
type: pull_request
number: 797
title: "Fixes 4666: lookup content sets instead of listing"
state: merged
author: jlsherrill
created: 2024-09-04T15:31:04Z
updated: 2024-09-06T14:30:31Z
closed: 2024-09-06T14:05:11Z
merged: 2024-09-06T14:05:11Z
base_branch: main
head_branch: debug_err
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/797
---

# Pull Request #797: Fixes 4666: lookup content sets instead of listing

**Author**: @jlsherrill
**Created**: September 04, 2024 at 03:31 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `debug_err`

## Description

## Summary

The newer version of candlepin limits the number of content sets returned and errors if more than 3000 are returned. 
Candlepin does support  paging, but these options are not exposed in the api spec for reasons.
We really don't need to list them all, so this changes it to fetch the contet sets by label (which is a new feature)

## Testing steps

Create a content template with some subset of rhel 9 repos and a custom repo.  Verify that a client sees just those pieces of content (or check the db and verify that the  cp_environment_content table in candlepin has what is expected



---

## Discussion

### Comment by @jlsherrill on September 05, 2024 at 01:30 PM UTC

https://issues.redhat.com/browse/HMS-4666

### Comment by @jlsherrill on September 06, 2024 at 02:05 PM UTC

merging since openapi was fixed by another PR

---

## Reviews

### Review by @rverdile - Approved on September 05, 2024 at 08:46 PM UTC

looks good, seeing only expected content sets on registered client

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/797*
