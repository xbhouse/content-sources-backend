---
type: pull_request
number: 386
title: "Fixes 1581: cancel in-progress snapshot on update"
state: merged
author: rverdile
created: 2023-09-11T15:35:11Z
updated: 2023-10-03T16:39:35Z
closed: 2023-10-03T16:39:34Z
merged: 2023-10-03T16:39:34Z
base_branch: main
head_branch: cancel-tasks
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/386
---

# Pull Request #386: Fixes 1581: cancel in-progress snapshot on update

**Author**: @rverdile
**Created**: September 11, 2023 at 03:35 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `cancel-tasks`

## Description

## Summary
This adds task cancellation to the tasking system and uses it to cancel in-progress snapshots when the URL is updated mid-snapshot.

I also added in a temporary task cancel API endpoint to use for testing while developing.

TODO
- [x] fix test
- [x] double check logging

## Testing steps
1. Create a repository with snapshots enabled. Use a larger one that takes some time to snapshot.
2. Immediately update that repository to use a different URL.
3. A new snapshot task will start for the different URL. The original snapshot will be canceled.
4. When the original snapshot task finishes, the status should be "canceled" and there should be no snapshot in the db or objects in pulp.
5. Query pulp with the hrefs on the original snapshot task payload to check that there are no objects in pulp. The tasks will exist, but they should have no created resources (repository version, publication, distribution).
6. The new snapshot task should be completed as normal.

If you'd like you can use [this gist to help you](https://gist.github.com/rverdile/d2325eaf6f55d248e646ef19ff36f09f).

---

## Discussion

### Comment by @jlsherrill on September 11, 2023 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-1581

### Comment by @jlsherrill on September 25, 2023 at 06:57 PM UTC

Anytime i cancel a task (either manually or via an update to the repo), i end up seeing: 
```
2:56PM ERR recovered panic in worker with error: runtime error: invalid memory address or nil pointer dereference error="runtime error: invalid memory address or nil pointer dereference" request_id=mumVj6Wi1m5JPCM44dORUFE3MQrhUGVs task_id=2b07a141-6503-4598-bdf9-3d567b2c395b task_type=snapshot
```

is that expected?

### Comment by @rverdile on September 25, 2023 at 07:04 PM UTC

> is that expected?

Not exactly. I broke it trying to debug why the integration test is failing here but not locally. You might test without the "debug failing integration test" commit.

### Comment by @jlsherrill on September 25, 2023 at 07:28 PM UTC

yes, you are right, no error without that commit!

### Comment by @swadeley on October 03, 2023 at 04:39 PM UTC

Hi

I created repo and then changed URL

```
Out[2]: 
{'created_at': '2023-10-03T15:15:43Z',
 'ended_at': '2023-10-03T15:16:19Z',
 'error': 'Get '
          '"http://cs-pulp-api-svc:24817/pulp%2Fed1ccac3%2Fapi%2Fv3%2Ftasks%2F018af61c-1751-7b24-b5eb-593452116986%2F": '
          'context canceled',
 'org_id': '3340851',
 'repository_name': '',
 'repository_uuid': '',
 'status': 'canceled',
 'type': 'snapshot',
 'uuid': '416e484c-ca79-4ead-98d9-794bb627c2d2'}
```

Connected to a pod
oc rsh  cs-pulp-content-869b8b99b4-l8bms

I got the cs-pulp-admin password from OpenShift (Workloads > Secrests tab)

```
sh-4.4$ curl  -u admin:<redacted>  "http://cs-pulp-api-svc:24817/pulp%2Fed1ccac3%2Fapi%2Fv3%2Ftasks%2F018af61c-1751-7b24-b5eb-593452116986%2F"
{"pulp_href":"/pulp/ed1ccac3/api/v3/tasks/018af61c-1751-7b24-b5eb-593452116986/","pulp_created":"2023-10-03T15:15:50.782880Z","state":"canceled","name":"pulp_rpm.app.tasks.synchronizing.synchronize","logging_cid":"65a16e77338e4cb391814bc06f3d70c4","created_by":"/pulp/ed1ccac3/api/v3/users/1/","started_at":"2023-10-03T15:15:50.839984Z","finished_at":"2023-10-03T15:16:18.378040Z","error":null,"worker":"/pulp/default/api/v3/workers/018aefe4-c415-793f-9c2f-4187715f1813/","parent_task":null,"child_tasks":[],"task_group":null,"progress_reports":[{"message":"Downloading Metadata Files","code":"sync.downloading.metadata","state":"completed","total":null,"done":6,"suffix":null},{"message":"Downloading Artifacts","code":"sync.downloading.artifacts","state":"running","total":null,"done":0,"suffix":null},{"message":"Associating Content","code":"associating.content","state":"running","total":null,"done":1000,"suffix":null},{"message":"Skipping Packages","code":"sync.skipped.packages","state":"completed","total":0,"done":0,"suffix":null},{"message":"Parsed Packages","code":"sync.parsing.packages","state":"running","total":17753,"done":2508,"suffix":null}],"created_resources":[],"reserved_resources_record":["/pulp/ed1ccac3/api/v3/repositories/rpm/rpm/018af61c-141c-7c8c-b417-214929d86b92/","shared:/pulp/ed1ccac3/api/v3/remotes/rpm/rpm/018af61c-0de9-7c23-9f4f-56a86a67fcd6/","shared:/pulp/default/api/v3/domains/018af61c-01ce-7668-83ce-4fc120b43383/"]}
```

Looks good, no "(repository version, publication, distribution)."

This is the next task:
```
Out[2]: 
{'created_at': '2023-10-03T15:16:12Z',
 'ended_at': '2023-10-03T15:16:31Z',
 'error': '',
 'org_id': '3340851',
 'repository_name': 'cancel in-progress snapshot on update #386',
 'repository_uuid': '26efd3eb-6022-4ea7-ade1-e1ccd6ec29d2',
 'status': 'completed',
 'type': 'snapshot',
 'uuid': 'f6fc179e-18e2-4320-9921-38ea11828943'}
```

Thank you

---

## Reviews

### Review by @jlsherrill - Commented on September 13, 2023 at 07:43 PM UTC

### Review by @jlsherrill - Commented on September 13, 2023 at 07:43 PM UTC

### Review by @rverdile - Commented on September 15, 2023 at 01:23 PM UTC

### Review by @rverdile - Commented on September 15, 2023 at 01:24 PM UTC

### Review by @jlsherrill - Commented on September 21, 2023 at 04:18 PM UTC

### Review by @jlsherrill - Commented on September 21, 2023 at 04:19 PM UTC

### Review by @jlsherrill - Commented on September 21, 2023 at 04:20 PM UTC

### Review by @rverdile - Commented on September 21, 2023 at 06:57 PM UTC

### Review by @rverdile - Commented on September 21, 2023 at 08:10 PM UTC

### Review by @rverdile - Commented on September 21, 2023 at 08:10 PM UTC

### Review by @jlsherrill - Commented on September 25, 2023 at 07:27 PM UTC

### Review by @jlsherrill - Commented on September 28, 2023 at 01:55 PM UTC

### Review by @rverdile - Commented on September 28, 2023 at 02:15 PM UTC

### Review by @jlsherrill - Commented on September 29, 2023 at 04:09 PM UTC

### Review by @rverdile - Commented on September 29, 2023 at 04:56 PM UTC

### Review by @rverdile - Commented on September 29, 2023 at 05:04 PM UTC

### Review by @jlsherrill - Commented on September 29, 2023 at 06:41 PM UTC

### Review by @jlsherrill - Commented on September 29, 2023 at 07:05 PM UTC

### Review by @rverdile - Commented on September 29, 2023 at 07:10 PM UTC

### Review by @jlsherrill - Commented on October 02, 2023 at 03:07 PM UTC

### Review by @jlsherrill - Approved on October 02, 2023 at 03:07 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/386*
