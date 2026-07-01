---
type: pull_request
number: 85
title: "Fixes 164: properly set versions when de-dupeing"
state: merged
author: jlsherrill
created: 2022-08-24T18:11:03Z
updated: 2022-08-25T16:04:24Z
closed: 2022-08-25T09:44:46Z
merged: 2022-08-25T09:44:46Z
base_branch: main
head_branch: 164
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/85
---

# Pull Request #85: Fixes 164: properly set versions when de-dupeing

**Author**: @jlsherrill
**Created**: August 24, 2022 at 06:11 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `164`

## Description

on hooks you have to call SetColumn, you cannot just
update the attribute directly, even if it works for creation

---

## Discussion

### Comment by @jlsherrill on August 24, 2022 at 06:12 PM UTC

Create a repo:
```
curl -X POST http://localhost:8000/api/content_sources/v1/repositories/   -H "`scripts/header.sh 1 1`"  -d '{"url": "http://foo.com/bar", "versions": ["7"], "name": "foo"}'  -H "Content-Type: application/json"   -v
```

Then update it with duplicated versions (setting the right UUID):
```
 1137  curl -X PUT http://localhost:8000/api/content_sources/v1/repositories/29126567-ff05-4608-9bf4-4630148be298   -H "`scripts/header.sh 1 1`"  -d '{"url": "http://foo.com/bazr", "distribution_versions": ["7", "7"], "name": "foo2"}'  -H "Content-Type: application/json"   -v
```

Then fetch the repo:
```
curl -X GET http://localhost:8000/api/content_sources/v1/repositories/29126567-ff05-4608-9bf4-4630148be298  -H "`scripts/header.sh 1 1`" 
```

You shouldn't see duplicated versions, only "7"

### Comment by @jlsherrill on August 24, 2022 at 06:22 PM UTC

fixes this iqe test:

tests.test_repositories.test_create_update_distro_version_is_unique (from pytest) 

```
AssertionError: assert ['8', '8'] == ['8']
  Left contains one more item: '8'
  Full diff:
  - ['8']
  + ['8', '8']
```

### Comment by @jlsherrill on August 24, 2022 at 06:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-164

### Comment by @mshriver on August 24, 2022 at 06:44 PM UTC

Test failures aren't related to the intent of this PR, and are already being addressed.

The intended test failure has been addressed as demonstrated by the results in ci.int.devshift.
```14:15:33 tests/test_repositories.py::test_create_update_distro_version_is_unique PASSED [ 96%]```


---

## Reviews

### Review by @mshriver - Approved on August 24, 2022 at 06:44 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/85*
