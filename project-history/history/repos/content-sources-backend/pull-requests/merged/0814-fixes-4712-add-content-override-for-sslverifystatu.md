---
type: pull_request
number: 814
title: "Fixes 4712: add content override for sslverifystatus"
state: merged
author: jlsherrill
created: 2024-09-12T17:10:34Z
updated: 2024-09-13T17:35:58Z
closed: 2024-09-13T17:35:46Z
merged: 2024-09-13T17:35:46Z
base_branch: main
head_branch: 4712
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/814
---

# Pull Request #814: Fixes 4712: add content override for sslverifystatus

**Author**: @jlsherrill
**Created**: September 12, 2024 at 05:10 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4712`

## Description

## Summary

By default clients registered to stage/prod have  sslsslverifystatus set to 1, which works fine for cdn.redhat.com, but not for s3.   This overrides that setting for each content set.

## Testing steps

Register a system and associate it to a template: https://github.com/content-services/content-sources-backend/blob/main/docs/register_client.md

on the client run 'subscription-manager refresh'

```
# grep sslverifystatus /etc/yum.repos.d/redhat.repo 
sslverifystatus = 0
```



---

## Discussion

### Comment by @jlsherrill on September 12, 2024 at 05:30 PM UTC

https://issues.redhat.com/browse/HMS-4712

### Comment by @jlsherrill on September 13, 2024 at 05:35 PM UTC

apologies, this will have to be qe'd in stage!

---

## Reviews

### Review by @Andrewgdewar - Approved on September 13, 2024 at 05:33 PM UTC

Was able to reproduce the expected outcome! Ack!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/814*
