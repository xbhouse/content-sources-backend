---
type: pull_request
number: 659
title: "Fixes 3832: update date type in /snapshots/for_date"
state: merged
author: xbhouse
created: 2024-05-02T20:49:56Z
updated: 2024-05-14T22:01:27Z
closed: 2024-05-14T21:46:28Z
merged: 2024-05-14T21:46:28Z
base_branch: main
head_branch: 3832
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/659
---

# Pull Request #659: Fixes 3832: update date type in /snapshots/for_date

**Author**: @xbhouse
**Created**: May 02, 2024 at 08:49 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `3832`

## Description

## Summary

- Updates the date in the request to /snapshots/for_date to be an alias of type time.Time
- Date can be formatted like either YYYY-MM-DD or 2024-05-02T00:00:00-04:00
- Frontend change for this as well to ensure template creation doesn't fail on the Set up date step, PR [here](https://github.com/content-services/content-sources-frontend/pull/256). <-- this is not needed anymore for this PR since both formats work, but will be once we remove support for YYYY-MM-DD in a following PR

## Testing steps

- Test the /snapshots/for_date endpoint with a request like this, this should work:
```
{
  "date": "2024-04-28T00:00:00-04:00",
  "repository_uuids": ["1510694a-1bfa-479a-b596-4377261f3135"]
}
```
- This should also work:
```
{
  "date": "2024-04-28",
  "repository_uuids": ["1510694a-1bfa-479a-b596-4377261f3135"]
}
```

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on May 02, 2024 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-3832

### Comment by @mayurilahane on May 06, 2024 at 02:00 PM UTC

/retest

### Comment by @mayurilahane on May 06, 2024 at 03:03 PM UTC

![image](https://github.com/content-services/content-sources-backend/assets/21276218/2e5abbe9-32e5-4db1-9edd-f555bf5c0b6a)


### Comment by @mayurilahane on May 06, 2024 at 03:09 PM UTC

And got 400 error 
for the old date type 
![image](https://github.com/content-services/content-sources-backend/assets/21276218/b2ed11d9-ddfb-40ba-bb0e-ff49bd6383f1)


### Comment by @jlsherrill on May 08, 2024 at 11:00 AM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @Andrewgdewar on May 09, 2024 at 09:20 PM UTC

> And got 400 error for the old date type ![image](https://private-user-images.githubusercontent.com/21276218/328225986-b2ed11d9-ddfb-40ba-bb0e-ff49bd6383f1.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MTUyODk4NzgsIm5iZiI6MTcxNTI4OTU3OCwicGF0aCI6Ii8yMTI3NjIxOC8zMjgyMjU5ODYtYjJlZDExZDktZGRmYi00MGJhLWJiMGUtZmY0OWJkNjM4M2YxLnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNDA1MDklMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjQwNTA5VDIxMTkzOFomWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPWVlODRlMjhhMDcyZTc3ZmQ5NWQxNTUwNmU1YzBmMWU0YzkyNTBkZWJlNjI4YWM5MGVkN2RhOTIwNTliOGFmOGImWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0JmFjdG9yX2lkPTAma2V5X2lkPTAmcmVwb19pZD0wIn0.z-dQ9Ww19uqfOtNdigzX1P-ft71epiglG_AILHU3AfY)

This was updated to now allow both date types (temporarily). 

### Comment by @xbhouse on May 14, 2024 at 03:22 PM UTC

> > And got 400 error for the old date type ![image](https://private-user-images.githubusercontent.com/21276218/328225986-b2ed11d9-ddfb-40ba-bb0e-ff49bd6383f1.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MTUyODk4NzgsIm5iZiI6MTcxNTI4OTU3OCwicGF0aCI6Ii8yMTI3NjIxOC8zMjgyMjU5ODYtYjJlZDExZDktZGRmYi00MGJhLWJiMGUtZmY0OWJkNjM4M2YxLnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNDA1MDklMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjQwNTA5VDIxMTkzOFomWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPWVlODRlMjhhMDcyZTc3ZmQ5NWQxNTUwNmU1YzBmMWU0YzkyNTBkZWJlNjI4YWM5MGVkN2RhOTIwNTliOGFmOGImWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0JmFjdG9yX2lkPTAma2V5X2lkPTAmcmVwb19pZD0wIn0.z-dQ9Ww19uqfOtNdigzX1P-ft71epiglG_AILHU3AfY)
> 
> This was updated to now allow both date types (temporarily).

@mayurilahane ^

### Comment by @mayurilahane on May 14, 2024 at 06:27 PM UTC

/retest

### Comment by @mayurilahane on May 14, 2024 at 09:45 PM UTC

![image](https://github.com/content-services/content-sources-backend/assets/21276218/170ceeba-efef-44c5-a633-1f143a4408dc)


### Comment by @mayurilahane on May 14, 2024 at 09:45 PM UTC

![image](https://github.com/content-services/content-sources-backend/assets/21276218/ce7ae94f-0805-444b-bc6b-37e855fa9fb8)


### Comment by @mayurilahane on May 14, 2024 at 09:46 PM UTC

LGTM!

---

## Reviews

### Review by @Andrewgdewar - Approved on May 09, 2024 at 09:23 PM UTC

Tested with both date formats, working well.
Ensure that @croissanne is told that we switched to this date-type. 
She originally wanted to use RFC 3339 as the format format for her blueprints api snapshot date, we should likely get her to switch her api over to that date as well; making updates to the image-builder front-end snapshots integration to support the above.

### Review by @rverdile - Commented on May 10, 2024 at 01:52 PM UTC

### Review by @rverdile - Commented on May 10, 2024 at 01:56 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/659*
