---
type: pull_request
number: 486
title: "Fixes 2969:  fetch (POST) snapshots for content_templates"
state: merged
author: Andrewgdewar
created: 2023-11-23T23:45:24Z
updated: 2023-12-19T16:00:35Z
closed: 2023-12-19T15:53:16Z
merged: 2023-12-19T15:53:16Z
base_branch: main
head_branch: HMS-2969
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/486
---

# Pull Request #486: Fixes 2969:  fetch (POST) snapshots for content_templates

**Author**: @Andrewgdewar
**Created**: November 23, 2023 at 11:45 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `HMS-2969`

## Description

## Summary

Adds the following API: 
POST /repositories/snapshots/for_date/

Body
`{     repository_uuids: ['abcde'],   date: “2023-01-01” }`

Response: 
```
[
   {
      repository_uuid: "abcde"   
      is_after: true,
      match: { SNAPHOT_JSON    }     
   }  
]
```

## Testing steps

- Create a few repositories, wait for them to snapshot. 
- Hit the above endpoint using your selected tool, filling in the repository_uuids and targeted time (UTC).

Additional information: 
- If the date of the nearest snap is greater than 24 hours past the specified date the "is_true" flag will be false.
- If no match is found, the "is_after" defaults to false.
- The Date will be compared as 24 hours inclusive. IE, if you snapshotted yesterday, and then now, and set the "Date" to yesterday's date, you'd likely still return that snap as a match without the "is_after" flag being switched. 
- Currently I am not surfacing not_found errors if the user has some matches, ie if you had 3 repository_uuids, and only had snaps for the first 2, the third (regardless of cause of not found) will just return an empty "match" and false "is_after" 

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on November 24, 2023 at 10:00 PM UTC

https://issues.redhat.com/browse/HMS-2969

### Comment by @Andrewgdewar on December 18, 2023 at 06:20 PM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Commented on November 27, 2023 at 06:53 PM UTC

### Review by @Andrewgdewar - Commented on November 29, 2023 at 05:04 PM UTC

### Review by @Andrewgdewar - Commented on November 29, 2023 at 05:05 PM UTC

### Review by @Andrewgdewar - Commented on November 29, 2023 at 10:33 PM UTC

### Review by @Andrewgdewar - Commented on November 30, 2023 at 09:52 PM UTC

### Review by @Andrewgdewar - Commented on December 06, 2023 at 07:20 PM UTC

### Review by @jlsherrill - Approved on December 11, 2023 at 06:57 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/486*
