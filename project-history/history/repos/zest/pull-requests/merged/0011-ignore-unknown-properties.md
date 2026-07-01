---
type: pull_request
number: 11
title: "ignore unknown properties"
state: merged
author: jlsherrill
created: 2024-03-06T14:34:13Z
updated: 2024-03-06T18:25:11Z
closed: 2024-03-06T18:25:11Z
merged: 2024-03-06T18:25:11Z
base_branch: main
head_branch: unknown
labels: []
url: https://github.com/content-services/zest/pull/11
---

# Pull Request #11: ignore unknown properties

**Author**: @jlsherrill
**Created**: March 06, 2024 at 02:34 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `unknown`

## Description

This ensures that if a new property is introduced in pulp, we won't error on older bindings that aren't expecting it.   I tested it locally and it seemed to work as promised.

---

*Archived from: https://github.com/content-services/zest/pull/11*
