---
type: pull_request
number: 841
title: "Fixes 4740: select correct snapshots for date selection"
state: merged
author: jlsherrill
created: 2024-10-08T19:34:04Z
updated: 2024-10-11T14:57:49Z
closed: 2024-10-11T14:57:45Z
merged: 2024-10-11T14:57:45Z
base_branch: main
head_branch: query
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/841
---

# Pull Request #841: Fixes 4740: select correct snapshots for date selection

**Author**: @jlsherrill
**Created**: October 08, 2024 at 07:34 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `query`

## Description

## Summary

Previously this function was written to return snapshots closest to the date, which isn't what we want.  We want it to return the newest snapshot prior to that date if it exists, if not the oldest snapshots after that date.

## Testing steps
It might work best to test this on a fresh/clean db

1.  Create a repository and let it snapshot:

```
####
POST http://localhost:8000/api/content-sources/v1.0/repositories/
x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==
//x-Rh-Insights-Request-Id: 9d229d34-7219-405d-ba3d-8f99cf7c36ae
Content-Type: application/json

{
  "name": "test57",
  "url": "http://jlsherrill.fedorapeople.org/fake-repos/needed-errata/",
  "gpg_key": "",
  "snapshot": true
}
```

2.  Update the snapshot to the previous day:
```
update snapshots set created_at = '2024-10-09T1130';
```

3.  Update the repo with a new URL and let it snapshot:
```
####
PATCH http://localhost:8000/api/content-sources/v1.0/repositories/48b45085-5366-4b55-9270-183198a887be
x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==
//x-Rh-Insights-Request-Id: 9d229d34-7219-405d-ba3d-8f99cf7c36ae
Content-Type: application/json

{
  "url": "http://yum.theforeman.org/pulpcore/3.16/el7/x86_64/"
}
```


So now you have two snapshots, one from 10-09 and one from 10-10  (or whatever today is).  You can now use the for_date api to test this change:


4.   
```
####
POST http://localhost:8000/api/content-sources/v1.0/snapshots/for_date/
x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==
Content-Type: application/json

{
  "repository_uuids": ["48b45085-5366-4b55-9270-183198a887be"],
  "date": "2024-10-10T15:34:58Z"
}
```


Before this change, you would get whichever snapshot is Closest to the date you specified.  So if you used a date 1 hour before the 2nd snapshot, you'd get the 2nd snapshot which wasn't correct.  The goal is to get the newest snapshot that existed at the time of the date you specifeid. 

Try specifying a date:
1.  prior to the first snapshot - should return the 1st snapshot
2. different times after the first snapshot but before the 2nd (test closer to the first one and closer to the 2nd) -  All should return the 1st snapshot
3. after the 2nd snapshot - should return the 2nd snapshot.

---

## Discussion

### Comment by @jlsherrill on October 08, 2024 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-4740

### Comment by @jlsherrill on October 11, 2024 at 02:21 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Commented on October 10, 2024 at 07:46 PM UTC

looks good! also tested with multiple repos

### Review by @rverdile - Approved on October 11, 2024 at 02:57 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/841*
