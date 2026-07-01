---
type: pull_request
number: 1394
title: "HMS-5830: replace custom epel with community on import"
state: merged
author: dominikvagner
created: 2026-02-19T12:32:36Z
updated: 2026-02-24T12:26:35Z
closed: 2026-02-24T12:26:35Z
merged: 2026-02-24T12:26:35Z
base_branch: main
head_branch: replace-custom-epel-on-import
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1394
---

# Pull Request #1394: HMS-5830: replace custom epel with community on import

**Author**: @dominikvagner
**Created**: February 19, 2026 at 12:32 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `replace-custom-epel-on-import`

## Description

## Summary
This adds a check for custom EPEL repos to the bulk import DAO method, so that when a custom EPEL repo is detected the community repo with the same URL is returned and the import isn't imported. Also suppresses errors/warnings on import of existing community repos.

## Testing steps
Try importing a custom EPEL and verify it returns a community one.
Try importing an existing community repo and verify there are no errors/warnings, only non-existing repos should error.




---

## Discussion

### Comment by @xbhouse on February 19, 2026 at 01:00 PM UTC

https://issues.redhat.com/browse/HMS-5830

### Comment by @xbhouse on February 19, 2026 at 09:44 PM UTC

i think the logic in the handler will need to be updated too to prevent triggering snapshot or introspect tasks for community repos

### Comment by @dominikvagner on February 23, 2026 at 09:30 AM UTC

> i think the logic in the handler will need to be updated too to prevent triggering snapshot or introspect tasks for community repos

good point! added a skip clause for community repos 👍🏼 

---

## Reviews

### Review by @xbhouse - Approved on February 23, 2026 at 04:08 PM UTC

this looks great, tested locally with the ib frontend following your [setup instructions](https://github.com/content-services/content-sources-backend?tab=readme-ov-file#testing-with-image-builder-frontend) and it works well 🥳  awesome job! 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1394*
