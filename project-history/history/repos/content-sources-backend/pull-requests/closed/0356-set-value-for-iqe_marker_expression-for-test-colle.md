---
type: pull_request
number: 356
title: "Set value for IQE_MARKER_EXPRESSION for test collection"
state: closed
author: aadhikar
created: 2023-08-10T11:48:16Z
updated: 2023-08-15T00:00:08Z
closed: 2023-08-15T00:00:08Z
base_branch: main
head_branch: marker-value
labels: ["qe-approved"]
url: https://github.com/content-services/content-sources-backend/pull/356
---

# Pull Request #356: Set value for IQE_MARKER_EXPRESSION for test collection

**Author**: @aadhikar
**Created**: August 10, 2023 at 11:48 AM UTC
**Status**: Closed
**Labels**: `qe-approved`
**Base**: `main` ← **Head**: `marker-value`

## Description

*No description provided*

---

## Discussion

### Comment by @app-sre-bot on August 10, 2023 at 11:49 AM UTC

Can one of the admins verify this patch?

### Comment by @mshriver on August 10, 2023 at 12:54 PM UTC

@aadhikar the PR Validation fails because the title isn't matching expected format for this repo, need to prefix with `Refs HMS-####`.

Please put in draft until the iqe-hms-integration MR introducing the marker is merged, otherwise this change would result in no tests getting collected for the pipeline. 

### Comment by @mshriver on August 14, 2023 at 11:57 PM UTC

Going to cover this in #358 

---

## Reviews

### Review by @mshriver - Commented on August 10, 2023 at 12:55 PM UTC

### Review by @mshriver - Changes Requested on August 10, 2023 at 12:55 PM UTC

Need to set parameter default, not envvar default 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/356*
