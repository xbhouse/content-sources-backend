---
type: pull_request
number: 1066
title: "Build: update go packages"
state: closed
author: dominikvagner
created: 2025-04-08T12:23:14Z
updated: 2025-04-09T07:16:44Z
closed: 2025-04-09T07:16:44Z
base_branch: main
head_branch: update-pkgs
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1066
---

# Pull Request #1066: Build: update go packages

**Author**: @dominikvagner
**Created**: April 08, 2025 at 12:23 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `update-pkgs`

## Description

## Summary
This PR combines all the currently open update PRs into 1. After this is merged those can be closed. 😅 

> Combines PR/updates to go packages: 
> - Update module github.com/klauspost/compress to v1.18.0
> - Update module github.com/fsnotify/fsnotify to v1.9.0
> - Update module github.com/cloudflare/circl to v1.6.0
> - Update module google.golang.org/protobuf to v1.36.6
> - Update github.com/rcrowley/go-metrics digest to 65e299d
> - Update module github.com/mattn/go-colorable to v0.1.14
> - Update module github.com/aws/smithy-go to v1.22.3
> - Update go-openapi packages

Closes #1048
Closes #1049 
Closes #1052 
Closes #1060
Closes #1061
Closes #1062
Closes #1063
Closes #1064

---

## Discussion

### Comment by @rverdile on April 08, 2025 at 07:10 PM UTC

Based on the discussion today, since these are all indirect dependencies, I think we can close this? :sweat_smile: 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1066*
