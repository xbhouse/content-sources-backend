---
type: pull_request
number: 4
title: "update to new pulp href"
state: merged
author: rverdile
created: 2024-01-23T16:44:43Z
updated: 2024-01-23T19:37:02Z
closed: 2024-01-23T19:37:00Z
merged: 2024-01-23T19:36:59Z
base_branch: main
head_branch: fix-href
labels: []
url: https://github.com/content-services/tang/pull/4
---

# Pull Request #4: update to new pulp href

**Author**: @rverdile
**Created**: January 23, 2024 at 04:44 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `fix-href`

## Description

Changes to pulp href, where api is preprended to the path, breaks the href parsing function. This shifts the indexes over one to account for that.

And update zest version

---

## Discussion

### Comment by @jlsherrill on January 23, 2024 at 05:00 PM UTC

probably want to update https://github.com/content-services/tang/blob/main/compose_files/pulp/docker-compose.yml  to use       PULP_API_ROOT: /api/pulp/

### Comment by @rverdile on January 23, 2024 at 06:32 PM UTC

thanks!

---

## Reviews

### Review by @jlsherrill - Approved on January 23, 2024 at 07:27 PM UTC

---

*Archived from: https://github.com/content-services/tang/pull/4*
