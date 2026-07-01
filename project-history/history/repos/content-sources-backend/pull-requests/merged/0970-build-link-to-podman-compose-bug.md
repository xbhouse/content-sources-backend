---
type: pull_request
number: 970
title: "Build: Link to podman-compose bug"
state: merged
author: swadeley
created: 2025-02-10T10:30:53Z
updated: 2025-09-08T07:31:31Z
closed: 2025-03-12T15:16:44Z
merged: 2025-03-12T15:16:44Z
base_branch: main
head_branch: swadeley/podman-compose_ver_3_bug
labels: []
url: https://github.com/content-services/content-sources-backend/pull/970
---

# Pull Request #970: Build: Link to podman-compose bug

**Author**: @swadeley
**Created**: February 10, 2025 at 10:30 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/podman-compose_ver_3_bug`

## Description

## Summary

Instructions should mention podman-compose bug in 1.3.0

## Testing steps



---

## Discussion

### Comment by @swadeley on February 10, 2025 at 10:37 AM UTC

@Andrewgdewar . on line 87: Ensure the correct [node version]

Can we say what version is the "correct" version? Or is the correct version whatever it finds in the config? So its  more of an enabling step.

EDIT:

Doc[1]  says `Latest version of Node.js 18, 20 or 22.`

EDIT 2: I see now, its just the valuse in the `.nvmrc` config file that needs to be used. I think the best is just to say `Apply the node version from the `.nvmrc` file:`  because stating the user must use the correct version just raises more questions. Just tell people what to do.

[1] https://playwright.dev/docs/intro#system-requirements

---

## Reviews

### Review by @swadeley - Commented on February 10, 2025 at 12:45 PM UTC

### Review by @Andrewgdewar - Approved on March 12, 2025 at 03:11 PM UTC

Readme changes! woot!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/970*
