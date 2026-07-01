---
type: pull_request
number: 753
title: "Fixes 4339,4371: add uploads API"
state: merged
author: jlsherrill
created: 2024-07-23T21:22:01Z
updated: 2024-08-08T12:06:31Z
closed: 2024-08-07T19:49:37Z
merged: 2024-08-07T19:49:37Z
base_branch: main
head_branch: 4339
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/753
---

# Pull Request #753: Fixes 4339,4371: add uploads API

**Author**: @jlsherrill
**Created**: July 23, 2024 at 09:22 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4339`

## Description

## Summary

## Testing steps

1.  Create a repository in the UI or Api
2.   use the ./scripts/upload.py script to upload some repositories
```
IDENTITY_HEADER="eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==" python ./scripts/uploads.py      "8695b221-bf5d-4f5b-b544-43e563e0f558"  ~/rpms/crow-0.8-1.noarch.rpm   ~/rpms/foo-1.2.3.noarch.rpm
```

try with large RPMs, many rpms, etc...
Note, you'll need to make sure you use the same org in the identity header as you've used to create the repo

4. Fetch the repository and check the package_count, fetch the repository rpms  (/repositories/UUID/rpms)  
5. check that a snapshot is made each time, and that the snapshots show the rpms that should be there
6. access the repo in pulp for the snapshots and verify expected rpms are present  

7.  Re-run the command with rpms that have already been uploaded, (or a mix of some that have and haven't), if there are no new rpms, no snapshot should be made

8.  If you want run the command with  -f to finalize the uploads before adding them.  Note that this will fail if the rpms are already present in the org (we likely won't be using this method going forward)

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on July 23, 2024 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-4339

### Comment by @jlsherrill on July 24, 2024 at 06:18 PM UTC

/retest

### Comment by @jlsherrill on July 26, 2024 at 07:38 PM UTC

/retest

### Comment by @jlsherrill on August 01, 2024 at 01:20 PM UTC

/retest

### Comment by @jlsherrill on August 02, 2024 at 06:04 PM UTC

feedback should now be addressed

### Comment by @jlsherrill on August 05, 2024 at 03:33 PM UTC

added upload permissions to this PR in a new commit

### Comment by @jlsherrill on August 05, 2024 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-4371

### Comment by @jlsherrill on August 06, 2024 at 12:45 PM UTC

apparently the previous pr that tried to prevent creation/update of upload repos with snapshot set to false wasn't working properly.  I've adjusted here and added better tests. 

In addition i added a catch at the start of the task so it won't panic

---

## Reviews

### Review by @rverdile - Commented on August 01, 2024 at 02:32 PM UTC

### Review by @rverdile - Commented on August 05, 2024 at 07:33 PM UTC

Looks really good overall. 

I see one issue where, if I create an upload repository and set snapshot to false, I get a panic trying to upload to that repository. 

It still panics if I set snapshot=true after and retry that upload

### Review by @rverdile - Approved on August 06, 2024 at 02:54 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/753*
