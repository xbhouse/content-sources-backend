---
type: pull_request
number: 951
title: "HMS-5293: add list feature content"
state: merged
author: rverdile
created: 2025-01-23T14:34:29Z
updated: 2025-02-04T16:18:28Z
closed: 2025-02-04T16:18:24Z
merged: 2025-02-04T16:18:23Z
base_branch: main
head_branch: list-feature-content
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/951
---

# Pull Request #951: HMS-5293: add list feature content

**Author**: @rverdile
**Created**: January 23, 2025 at 02:34 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `list-feature-content`

## Description

## Summary
This adds an API to list the content if a feature from the feature service API

## Testing steps
1. Add the correct certs path to the clients.feature_service option in the config.yaml
2. Use the `GET /api/content-sources/v1/admin/features/` endpoint list the features
3. Choose one and use the name with `GET /api/content-sources/v1/admin/features/:name/content/` to list the content
4. Check that the fields match what is described by the JIRA ticket



---

## Discussion

### Comment by @jlsherrill on January 23, 2025 at 03:00 PM UTC

https://issues.redhat.com/browse/HMS-5293

### Comment by @jlsherrill on January 27, 2025 at 08:20 PM UTC

overall looked good.  There are some issues with the older release (such as RHEL 4,5.etc..) with distribution & arch not being correct,  but IDK if need to worry about it right now.  

I think we'll probably just wanna adjust the logic as we go based on what we need.   It may change some with RHEL 10 and we may need to adjust regardless



---

## Reviews

### Review by @jlsherrill - Commented on January 27, 2025 at 02:17 PM UTC

### Review by @rverdile - Commented on January 28, 2025 at 04:42 PM UTC

### Review by @jlsherrill - Approved on February 03, 2025 at 07:13 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/951*
