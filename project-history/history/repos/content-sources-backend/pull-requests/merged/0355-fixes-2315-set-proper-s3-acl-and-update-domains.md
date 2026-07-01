---
type: pull_request
number: 355
title: "Fixes 2315: set proper s3 acl and update domains"
state: merged
author: jlsherrill
created: 2023-08-09T20:37:47Z
updated: 2023-08-10T18:43:04Z
closed: 2023-08-10T18:36:55Z
merged: 2023-08-10T18:36:55Z
base_branch: main
head_branch: 2315_2
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/355
---

# Pull Request #355: Fixes 2315: set proper s3 acl and update domains

**Author**: @jlsherrill
**Created**: August 09, 2023 at 08:37 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `2315_2`

## Description

## Summary
previously domains configured with s3 were not configured with an acl that was correct.  This fixes that, but some domains may already be created improperly.  This adds the ability to update s3 storage settings in pulp if they have changed

## Testing steps
A bit hard to truly test.

1.  Create an s3 bucket on an amazon account. (default config is fine)
2.  within IAM in s3, create a user with full s3 permissions (AmazonS3FullAccess policy)
3. create an access key for them
4. on a fresh db, on 'main' branch, configure access to this account:
 ```
  pulp:
    server: http://localhost:8080
    username: admin
    password: password
    storage_type: object #object or local
    custom_repo_objects:
      url: https://s3.us-east-1.amazonaws.com/
      access_key: "MyKey"
      secret_key: "MySecret"
      name: test-pulp
      region: us-east-1
```
5.  try to create and snapshot a repo, this will error
6. switch to this branch, and start the server.
7. create a new repo in the same account. It should snapshot with no issue.

---

## Discussion

### Comment by @jlsherrill on August 09, 2023 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-2315

---

## Reviews

### Review by @rverdile - Approved on August 10, 2023 at 05:53 PM UTC

worked!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/355*
