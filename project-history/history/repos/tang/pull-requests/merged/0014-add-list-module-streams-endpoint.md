---
type: pull_request
number: 14
title: "Add list module Streams endpoint"
state: merged
author: Andrewgdewar
created: 2024-11-14T19:10:44Z
updated: 2025-01-06T21:58:27Z
closed: 2024-12-19T20:42:56Z
merged: 2024-12-19T20:42:56Z
base_branch: main
head_branch: HMS-4932
labels: []
url: https://github.com/content-services/tang/pull/14
---

# Pull Request #14: Add list module Streams endpoint

**Author**: @Andrewgdewar
**Created**: November 14, 2024 at 07:10 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `HMS-4932`

## Description

This adds the "RpmRepositoryVersionModuleStreams" method. 

That takes the following inputs
```
{
  "uuids": [
    "9d164ac6-b866-4808-af1f-847fbf588658",
    "36e449b9-3d61-43e8-a495-6ac332fdd3a8"
  ],
  "rpm_names": [],
  "search": ""
}

```
and returns a list of modules with their streams:

```
{
 data:[
  {
    "module_name": "swig",
    "streams": [
      {
        "Name": "swig",
        "Stream": "4.1",
        "Context": "rhel9",
        "Arch": "x86_64",
        "Version": "9020020221219141850",
        "Description": "Simplified Wrapper ... etc",
        "Profiles": {
          "common": [
            "swig"
          ],
          "complete": [
            "swig",
            "swig-doc",
            "swig-gdb"
          ]
        }
      }
    ]
  }
]
}
```

## Additional changes
- Added a mockery make command.

---

## Discussion

### Comment by @rverdile on December 02, 2024 at 06:11 PM UTC

Changes look good, but since justin [suggested](https://github.com/content-services/content-sources-backend/pull/897#discussion_r1865876713) to add some more fields to the response, I'll review again once that's done

### Comment by @Andrewgdewar on December 09, 2024 at 08:03 PM UTC

> it looks like pagination was removed, but I think we still want that, right?

No, unfortunately we decided we don't. There is a hard limit of 5000 items if that helps 😓 

---

## Reviews

### Review by @rverdile - Commented on November 22, 2024 at 06:46 PM UTC

A few comments, but looks really good overall. I tested it and everything seems to be working.

### Review by @Andrewgdewar - Commented on November 22, 2024 at 08:11 PM UTC

### Review by @Andrewgdewar - Commented on November 22, 2024 at 08:12 PM UTC

### Review by @Andrewgdewar - Commented on November 22, 2024 at 08:12 PM UTC

### Review by @Andrewgdewar - Commented on November 22, 2024 at 08:13 PM UTC

### Review by @rverdile - Commented on December 09, 2024 at 07:42 PM UTC

it looks like pagination was removed, but I think we still want that, right?

### Review by @rverdile - Commented on December 10, 2024 at 03:06 PM UTC

### Review by @rverdile - Commented on December 10, 2024 at 03:07 PM UTC

### Review by @jlsherrill - Commented on December 12, 2024 at 09:53 PM UTC

### Review by @Andrewgdewar - Commented on December 16, 2024 at 06:31 PM UTC

### Review by @jlsherrill - Commented on December 16, 2024 at 08:11 PM UTC

### Review by @jlsherrill - Approved on December 17, 2024 at 06:12 PM UTC

ack from me, @rverdile may want one final look?

### Review by @rverdile - Approved on December 17, 2024 at 08:00 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/tang/pull/14*
