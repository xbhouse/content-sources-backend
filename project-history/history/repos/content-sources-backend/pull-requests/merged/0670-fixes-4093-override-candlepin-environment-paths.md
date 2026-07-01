---
type: pull_request
number: 670
title: "Fixes 4093: Override candlepin environment paths"
state: merged
author: jlsherrill
created: 2024-05-15T18:10:40Z
updated: 2024-05-21T20:04:49Z
closed: 2024-05-21T20:04:41Z
merged: 2024-05-21T20:04:41Z
base_branch: main
head_branch: override
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/670
---

# Pull Request #670: Fixes 4093: Override candlepin environment paths

**Author**: @jlsherrill
**Created**: May 15, 2024 at 06:10 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `override`

## Description

for custom repos

## Summary

Adds overriding of custom content baseurls 

## Testing steps

Follow the directions here: https://github.com/content-services/content-sources-backend/blob/main/docs/register_client.md

At the end you should have a client configured to use a  custom repository and a red hat repository.  You should now be able to 'dnf install' a package from the custom repo

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on May 15, 2024 at 06:30 PM UTC

https://issues.redhat.com/browse/HMS-4093

### Comment by @jlsherrill on May 16, 2024 at 08:04 PM UTC

I'm not sure it fully matters, as its ignored if the content isn't in the environment.  It should be deleted when the content is deleted.... which we need to do on repo deletion, i'll file a task.  If you feel like we should be removing them, i can add that

---

## Reviews

### Review by @jlsherrill - Commented on May 16, 2024 at 07:03 PM UTC

### Review by @rverdile - Commented on May 16, 2024 at 07:48 PM UTC

### Review by @rverdile - Commented on May 16, 2024 at 07:53 PM UTC

should this also delete content overrides if a repository is removed from a template?

### Review by @jlsherrill - Commented on May 16, 2024 at 08:01 PM UTC

### Review by @rverdile - Commented on May 21, 2024 at 06:17 PM UTC

### Review by @rverdile - Commented on May 21, 2024 at 06:18 PM UTC

I created a template with RHEL 9 baseos and appstream, but I only see a content override for the cert of appstream. Is that a bug?

```
[
  {
    "created": "2024-05-21T17:30:42+0000",
    "updated": "2024-05-21T17:30:42+0000",
    "name": "sslcacert",
    "contentLabel": "rhel-9-for-x86_64-appstream-rpms",
    "value": " "
  }
]
```

### Review by @jlsherrill - Commented on May 21, 2024 at 06:24 PM UTC

### Review by @rverdile - Approved on May 21, 2024 at 07:23 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/670*
