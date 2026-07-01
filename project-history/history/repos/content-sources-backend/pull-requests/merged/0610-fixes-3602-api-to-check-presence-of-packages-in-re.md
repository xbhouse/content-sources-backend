---
type: pull_request
number: 610
title: "Fixes 3602: api to check presence of packages in repos"
state: merged
author: xbhouse
created: 2024-03-18T14:20:14Z
updated: 2024-03-22T14:30:26Z
closed: 2024-03-22T14:27:01Z
merged: 2024-03-22T14:27:01Z
base_branch: main
head_branch: detect-presence-rpms
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/610
---

# Pull Request #610: Fixes 3602: api to check presence of packages in repos

**Author**: @xbhouse
**Created**: March 18, 2024 at 02:20 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `detect-presence-rpms`

## Description

## Summary

- Adds api to check if one or more package names exist in a set of repositories (given by URL and/or UUID)

## Testing steps

- Create/introspect a few repositories
- Check for package names that exist in those repos with this endpoint and sample request: 

`POST /rpms/presence`

```
{
  "rpm_names": ["giraffe", "lion", "elephant", "6tunnel", "zebra"],
  "urls": [
    "https://fedorapeople.org/groups/katello/fakerepos/zoo/",
    "https://fedorapeople.org/groups/katello/fakerepos/zoo2/"
  ],
  "uuids": [ "5fc43beb-975f-44c2-98b0-584640d7ea55"
  ]
}
```

Response should return found and missing packages:

```
{
    "found": [
        "6tunnel",
        "elephant",
        "giraffe",
        "lion"
    ],
    "missing": [
        "zebra"
    ]
}
```

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on March 18, 2024 at 02:30 PM UTC

https://issues.redhat.com/browse/HMS-3602

### Comment by @swadeley on March 19, 2024 at 09:50 AM UTC

/retest

### Comment by @swadeley on March 19, 2024 at 11:25 AM UTC

Hi

lgtm


```
In [12]: app.content_sources.rest_client.rpms_api.detect_rpm({
    ...:   "rpm_names": ["bear", "camel", "giraffe", "lion", "zebra"],
    ...:   "urls": [
    ...:     "https://jlsherrill.fedorapeople.org/fake-repos/revision/one/",
    ...:     "https://jlsherrill.fedorapeople.org/fake-repos/revision/two/"
    ...:   ],
    ...:   "uuids": [ "45d1e4de-ef89-4809-8831-e0179882fced"
    ...:   ]
    ...: })
2024-03-19 11:24:16.255 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[12]: {'found': ['bear', 'camel'], 'missing': ['giraffe', 'lion', 'zebra']}
```

### Comment by @mayurilahane on March 21, 2024 at 01:35 PM UTC

/retest

---

## Reviews

### Review by @xbhouse - Commented on March 18, 2024 at 02:47 PM UTC

### Review by @rverdile - Commented on March 19, 2024 at 02:25 PM UTC

### Review by @rverdile - Commented on March 19, 2024 at 02:38 PM UTC

this looks good! just going to play around with it a bit

### Review by @xbhouse - Commented on March 19, 2024 at 04:36 PM UTC

### Review by @rverdile - Commented on March 20, 2024 at 03:06 PM UTC

### Review by @rverdile - Commented on March 20, 2024 at 03:09 PM UTC

### Review by @rverdile - Commented on March 20, 2024 at 03:15 PM UTC

### Review by @xbhouse - Commented on March 20, 2024 at 04:03 PM UTC

### Review by @rverdile - Commented on March 20, 2024 at 04:58 PM UTC

### Review by @rverdile - Commented on March 20, 2024 at 05:11 PM UTC

### Review by @xbhouse - Commented on March 20, 2024 at 05:36 PM UTC

### Review by @xbhouse - Commented on March 20, 2024 at 05:40 PM UTC

### Review by @rverdile - Commented on March 20, 2024 at 08:17 PM UTC

### Review by @rverdile - Approved on March 20, 2024 at 09:12 PM UTC

nice job!!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/610*
