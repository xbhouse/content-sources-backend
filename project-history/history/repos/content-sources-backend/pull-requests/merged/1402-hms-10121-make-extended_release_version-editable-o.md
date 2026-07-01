---
type: pull_request
number: 1402
title: "HMS-10121: make extended_release_version editable on templates"
state: merged
author: katarinazaprazna
created: 2026-03-03T14:22:59Z
updated: 2026-03-10T16:19:31Z
closed: 2026-03-10T16:19:31Z
merged: 2026-03-10T16:19:31Z
base_branch: main
head_branch: templates-edit-minor-version
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1402
---

# Pull Request #1402: HMS-10121: make extended_release_version editable on templates

**Author**: @katarinazaprazna
**Created**: March 03, 2026 at 02:22 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `templates-edit-minor-version`

## Description

## Summary

Allows updating the `extended_release_version` field on templates via PATCH/PUT endpoints

## Testing steps

---

## Discussion

### Comment by @xbhouse on March 03, 2026 at 02:30 PM UTC

https://issues.redhat.com/browse/HMS-10121

### Comment by @rverdile on March 03, 2026 at 03:32 PM UTC

If we add `--pull=newer` [here](https://github.com/content-services/content-sources-backend/blob/main/mk/swag.mk#L33-L34), it makes it so if there's a new image it will pull the new one

### Comment by @katarinazaprazna on March 04, 2026 at 10:22 PM UTC

Hey @rverdile. I've added the Docker option, but it looks like the OpenAPI docs have changed quite a bit. Are you seeing the same output on your end when you run the command?

### Comment by @rverdile on March 05, 2026 at 08:30 PM UTC

@katarinazaprazna that's expected!

### Comment by @rverdile on March 05, 2026 at 09:32 PM UTC

I think you need to re-run `make openapi`

---

## Reviews

### Review by @rverdile - Commented on March 05, 2026 at 09:32 PM UTC

### Review by @rverdile - Approved on March 10, 2026 at 04:14 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1402*
