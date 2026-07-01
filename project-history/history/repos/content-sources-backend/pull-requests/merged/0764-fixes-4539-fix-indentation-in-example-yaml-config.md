---
type: pull_request
number: 764
title: "Fixes 4539: fix indentation in example yaml config"
state: merged
author: dominikvagner
created: 2024-08-06T15:42:26Z
updated: 2024-08-07T12:40:14Z
closed: 2024-08-07T12:40:14Z
merged: 2024-08-07T12:40:14Z
base_branch: main
head_branch: 4539
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/764
---

# Pull Request #764: Fixes 4539: fix indentation in example yaml config

**Author**: @dominikvagner
**Created**: August 06, 2024 at 03:42 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4539`

## Description

## Summary
There is a wrong indentation in the `configs/config.yaml.example` file which when copied to remove `.example` extension doesn't work and causes `make compose-up` fail. 

## Testing steps
Fixed indentation and it works now after copying the file.


---

## Discussion

### Comment by @app-sre-bot on August 06, 2024 at 03:43 PM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on August 06, 2024 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-4539

---

## Reviews

### Review by @xbhouse - Approved on August 07, 2024 at 12:24 PM UTC

nice!! 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/764*
