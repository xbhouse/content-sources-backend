---
type: pull_request
number: 1380
title: "HMS-9959: support extended release versions for templates"
state: merged
author: rverdile
created: 2026-01-26T15:57:51Z
updated: 2026-02-12T17:42:08Z
closed: 2026-02-12T17:42:05Z
merged: 2026-02-12T17:42:05Z
base_branch: main
head_branch: release-templates
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1380
---

# Pull Request #1380: HMS-9959: support extended release versions for templates

**Author**: @rverdile
**Created**: January 26, 2026 at 03:57 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `release-templates`

## Description

## Summary
Adds support for extended release versions to templates

## Testing steps
1. Create a template with an extended release version
2. Make sure to set the extended_release_version and extended_release
3. Register the template to system: https://github.com/content-services/content-sources-backend/blob/main/docs/register_client.md
4. refresh subscription manager
5. Set the release version on the system e.g. `echo "9.6" | sudo tee /etc/dnf/vars/releasever`
6. Try downloading a package e.g. `dnf install httpd`



---

## Discussion

### Comment by @xbhouse on January 26, 2026 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-9959

### Comment by @rverdile on February 10, 2026 at 06:19 PM UTC

updated

---

## Reviews

### Review by @dominikvagner - Changes Requested on February 06, 2026 at 01:26 PM UTC

2 small comments, otherwise it looks great! ✨ 
tested it with a system and everything works as expected! 🚀 

### Review by @rverdile - Commented on February 06, 2026 at 07:19 PM UTC

### Review by @rverdile - Commented on February 06, 2026 at 07:25 PM UTC

### Review by @dominikvagner - Approved on February 12, 2026 at 04:23 PM UTC

awesome work! thx! ✨ 

ack! ✅ 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1380*
