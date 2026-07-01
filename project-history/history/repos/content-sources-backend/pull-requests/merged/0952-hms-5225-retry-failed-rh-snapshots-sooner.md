---
type: pull_request
number: 952
title: "HMS-5225: retry failed RH snapshots sooner"
state: merged
author: jlsherrill
created: 2025-01-24T18:15:38Z
updated: 2025-02-06T15:29:33Z
closed: 2025-02-06T15:29:32Z
merged: 2025-02-06T15:29:32Z
base_branch: main
head_branch: 5225
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/952
---

# Pull Request #952: HMS-5225: retry failed RH snapshots sooner

**Author**: @jlsherrill
**Created**: January 24, 2025 at 06:15 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `5225`

## Description

## Summary

Will retry failed RH snapshots at next job run time (currently every hour).  Will retry custom repo snapshots once a day, up to 10 times.  After 10 times, the user will have to manually snaphot the repository.  Once it snapshots cleanly, the count will reset to zero.

This also fixes a bug where non requeable tasks had a next_retry_time filled in.  The only effect of this bug is that you couldn't manually snapshot a repo if the previous snapshot had failed.  

## Testing steps

(I did this on a fresh db, its probably easier to do it this way).

Create a custom repository with an invalid url:

```
POST http://localhost:8000/api/content-sources/v1.0/repositories/
x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==
//x-Rh-Insights-Request-Id: 9d229d34-7219-405d-ba3d-8f99cf7c36ae
Content-Type: application/json

{
  "name": "test16",
  "url": "https://example.com/badurl/",
  "gpg_key": "",
  "snapshot": true
}
```


Fetch the repo and see the failed_snapshot_count increment:
```
GET http://localhost:8000/api/content-sources/v1.0/repositories/?origin=external
x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==
//x-Rh-Insights-Request-Id: 9d229d34-7219-405d-ba3d-8f99cf7c36ae
Content-Type: application/json
```

      "failed_snapshot_count": 1,


Trigger a manual snapshot 
```
POST http://localhost:8000/api/content-sources/v1.0/repositories/04fbfc3f-603c-454b-8f64-355597a1a964/snapshot/
x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==
//x-Rh-Insights-Request-Id: 9d229d34-7219-405d-ba3d-8f99cf7c36ae
Content-Type: application/json
```

This will continue to fail.

Update all the tasks to be queued at one day ago:

```
make db-cli-connect
 update tasks set queued_at = CURRENT_DATE - 1;
```

run the nightly job to trigger it:
```
 go run cmd/external-repos/main.go nightly-jobs 24
```

you should see the one repo get queued for snapshotting each time.  You'll have to re-set the dates in between.   
If you don't reset the queued_at date inbetween, it won't queue the snapshot (feel free to test this).

Once the failed_snapshot_count hits 10, the snapshot should no longer show up queued in the nightly-job.



Now to test red hat repos. Import rhel9 repos:

```
make repos-import-rhel9
```

Comment out the cdn certs in your config.yaml, to cause snapshots to fail:
```
certs:
 #cert_path: "/home/jlsherri/cdncert/cert.pem"
```

restart the server locally

now we need to try to snapshot these repos 10 times:

```
for i in `seq 1 10`; do go run cmd/external-repos/main.go nightly-jobs 24  ; sleep 15; done;
```

After this is done you can check the count:

```
select failed_snapshot_count from repository_configurations where org_id = '-1';
```

Now re-running the nightly job, should trigger the snapshots again, even though we are over the limit:

```
 go run cmd/external-repos/main.go nightly-jobs 24 
```


You can now add the certs back, restart the server and try it again:
```
 go run cmd/external-repos/main.go nightly-jobs 24 
```
and they should still snapshot.

---

## Discussion

### Comment by @jlsherrill on January 24, 2025 at 06:30 PM UTC

https://issues.redhat.com/browse/HMS-5225

### Comment by @swadeley on January 31, 2025 at 09:08 AM UTC

@jlsherrill Please rebase

### Comment by @jlsherrill on February 05, 2025 at 08:43 PM UTC

should be good for api regeneration

### Comment by @swadeley on February 06, 2025 at 01:02 PM UTC

/retest

### Comment by @swadeley on February 06, 2025 at 01:37 PM UTC

> should be good for api regeneration

Hi @jlsherrill 

Something is wrong, tests are failing:
`ApiRepositoryResponse has no attribute 'failed_snapshot_count'`

EDIT:  I think I need another rebase my side.,



### Comment by @swadeley on February 06, 2025 at 01:56 PM UTC

/retest

### Comment by @swadeley on February 06, 2025 at 02:47 PM UTC

/retest

### Comment by @swadeley on February 06, 2025 at 03:28 PM UTC

```
15:26:27 ==== 42 passed, 3 skipped, 26 deselected, 861 warnings in 535.62s (0:08:55) ====
```

---

## Reviews

### Review by @dominikvagner - Approved on January 31, 2025 at 08:01 AM UTC

looks good, thanks for the great testing steps, works as expected 👏🏼 🎉 
approved! ✅ but the IQE API specs need to be updated 😅 

### Review by @swadeley - Commented on January 31, 2025 at 09:07 AM UTC

### Review by @jlsherrill - Commented on February 03, 2025 at 07:21 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/952*
