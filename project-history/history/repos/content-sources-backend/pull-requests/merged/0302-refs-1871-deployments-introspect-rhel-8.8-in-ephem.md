---
type: pull_request
number: 302
title: "Refs 1871: Deployments: Introspect rhel-8.8 in ephemeral"
state: merged
author: jrusz
created: 2023-06-01T13:06:53Z
updated: 2023-06-01T13:54:40Z
closed: 2023-06-01T13:54:40Z
merged: 2023-06-01T13:54:40Z
base_branch: main
head_branch: ephemeral-introspection
labels: []
url: https://github.com/content-services/content-sources-backend/pull/302
---

# Pull Request #302: Refs 1871: Deployments: Introspect rhel-8.8 in ephemeral

**Author**: @jrusz
**Created**: June 01, 2023 at 01:06 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `ephemeral-introspection`

## Description

We started building rhel-8.8 images and so we now need to also update the repository introspection in ephemeral environments.

Related to https://github.com/content-services/content-sources-backend/pull/301

## Summary

## Testing steps


---

## Discussion

### Comment by @app-sre-bot on June 01, 2023 at 01:07 PM UTC

Can one of the admins verify this patch?

### Comment by @mshriver on June 01, 2023 at 01:16 PM UTC

/test

### Comment by @jlsherrill on June 01, 2023 at 01:30 PM UTC

https://issues.redhat.com/browse/HMS-1871

### Comment by @jlsherrill on June 01, 2023 at 01:54 PM UTC

Lets merge this for now and continue to discuss a  more permanent solution.

---

## Reviews

### Review by @mshriver - Commented on June 01, 2023 at 01:16 PM UTC

I don't think this is strictly necessary (there are a lot more repos to introspect) but we're using 8.8 in tests that end up failing because its not introspected yet. 

### Review by @mshriver - Approved on June 01, 2023 at 01:49 PM UTC

I don't think this is strictly necessary (there are a lot more repos to introspect) but we're using 8.8 in tests that end up failing because its not introspected yet. 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/302*
