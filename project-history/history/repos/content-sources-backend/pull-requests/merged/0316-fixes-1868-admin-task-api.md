---
type: pull_request
number: 316
title: "Fixes 1868: Admin Task API"
state: merged
author: dpang314
created: 2023-06-22T20:51:06Z
updated: 2023-07-14T19:45:23Z
closed: 2023-07-14T19:45:23Z
merged: 2023-07-14T19:45:22Z
base_branch: main
head_branch: HMS-1868-admin-tasks
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/316
---

# Pull Request #316: Fixes 1868: Admin Task API

**Author**: @dpang314
**Created**: June 22, 2023 at 08:51 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `HMS-1868-admin-tasks`

## Description

## Summary

Adds backend API for admin task interface that allows a user to view all tasks across all organizations/accounts. It is implemented as a feature that can be enabled/disabled and restricted to specific accounts in the config.

The API allows filtering by org_id, account_id, and statuses. 

The api struct has a copy of a zest struct in order to properly convert the pulp data to JSON.

The endpoint is not included in the OpenAPI docs.

## Testing steps

Add 
```
features:
  admin_tasks:
    enabled: true
    users: ["adminUser"]
```
to the config.yaml

To see whether the feature is enabled/accessible, you can make a request to the features API:

```
curl localhost:8000/api/content-sources/v1.0/features/  -H "`./scripts/header.sh 1 adminUser`"
```

To list tasks:
```
curl "http://localhost:8000/api/content-sources/v1/admin/tasks/?offset=0&limit=20"  -H "`./scripts/header.sh 1 adminUser`" 
```

To fetch a task:
```
curl "http://localhost:8000/api/content-sources/v1/admin/tasks/{uuid}"  -H "`./scripts/header.sh 1 adminUser`" 
```



---

## Discussion

### Comment by @jlsherrill on June 22, 2023 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-1868

### Comment by @jlsherrill on July 05, 2023 at 06:40 PM UTC

This needs a rebase, but everything looks good and works well!

### Comment by @swadeley on July 10, 2023 at 03:19 PM UTC

/retest

### Comment by @swadeley on July 11, 2023 at 06:52 AM UTC

Hi @dpang314 

Please rebase, because:

I see you have commit for "Fixes 1624" from Date:   Tue Jun 13 11:12:54 , but there are commit from after that date: https://github.com/content-services/content-sources-backend/pull/309/commits

and we have problems with pulp pods not starting when testing this PR.

Thank you

### Comment by @jlsherrill on July 14, 2023 at 02:32 PM UTC

[retest]

### Comment by @jlsherrill on July 14, 2023 at 02:33 PM UTC

/retest

### Comment by @swadeley on July 14, 2023 at 07:27 PM UTC

Hi

testing with:
      parameters:
        FEATURES_ADMIN_TASKS_ENABLED: "True"
        FEATURES_SNAPSHOTS_ENABLED: "True"


```
In [1]: app.content_sources.rest_client.features_api.list_features()
<snip>
Out[1]: 
{'admintasks': {'accessible': True, 'enabled': True},
 'snapshots': {'accessible': True, 'enabled': False}}
```

Here are the two tasks after a repo is created with snapshot=true:
```
content-sources-backend]$ curl  -u "$BASIC_AUTH" "https://${URL}/api/content-sources/v1.0/admin/tasks/?limit=6&offset=0" | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1916  100  1916    0     0   2162      0 --:--:-- --:--:-- --:--:--  2162
{
  "data": [
    {
      "uuid": "d5c7adeb-66b4-4f72-a25d-d9045767bb80",
      "status": "completed",
      "typename": "introspect",
      "queued_at": "2023-07-14T19:03:39Z",
      "started_at": "2023-07-14T19:03:39Z",
      "finished_at": "2023-07-14T19:03:39Z",
      "error": "",
      "org_id": "3340851",
      "account_id": "0369233",
      "payload": null,
      "pulp": {}
    },
    {
      "uuid": "1ac37a45-51ed-480c-bcaf-fdfc9dffd6f2",
      "status": "completed",
      "typename": "snapshot",
      "queued_at": "2023-07-14T19:03:39Z",
      "started_at": "2023-07-14T19:03:39Z",
      "finished_at": "2023-07-14T19:04:08Z",
      "error": "",
      "org_id": "3340851",
      "account_id": "0369233",
      "payload": null,
      "pulp": {}
    },
```

The first task where the repo is introspected:
```
content-sources-backend]$ curl  -u "$BASIC_AUTH" "https://${URL}/api/content-sources/v1.0/admin/tasks/{"d5c7adeb-66b4-4f72-a25d-d9045767bb80"}" | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   352  100   352    0     0    599      0 --:--:-- --:--:-- --:--:--   600
{
  "uuid": "d5c7adeb-66b4-4f72-a25d-d9045767bb80",
  "status": "completed",
  "typename": "introspect",
  "queued_at": "2023-07-14T19:03:39Z",
  "started_at": "2023-07-14T19:03:39Z",
  "finished_at": "2023-07-14T19:03:39Z",
  "error": "",
  "org_id": "3340851",
  "account_id": "0369233",
  "payload": {
    "Url": "https://stephenw.fedorapeople.org/multirepos/5/repo02/",
    "Force": true
  },
  "pulp": {}
}
```

The second task where the snapshot is made:
```
 content-sources-backend]$ curl  -u "$BASIC_AUTH" "https://${URL}/api/content-sources/v1.0/admin/tasks/{"1ac37a45-51ed-480c-bcaf-fdfc9dffd6f2"}" | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  2639    0  2639    0     0    486      0 --:--:--  0:00:05 --:--:--   534
{
  "uuid": "1ac37a45-51ed-480c-bcaf-fdfc9dffd6f2",
  "status": "completed",
  "typename": "snapshot",
  "queued_at": "2023-07-14T19:03:39Z",
  "started_at": "2023-07-14T19:03:39Z",
  "finished_at": "2023-07-14T19:04:08Z",
  "error": "",
  "org_id": "3340851",
  "account_id": "0369233",
  "payload": {
    "SyncTaskHref": "/pulp/api/v3/tasks/018955c9-b6e6-7a7b-8aa1-5a0a39bde143/",
    "SnapshotIdent": null,
    "PublicationTaskHref": "/pulp/api/v3/tasks/018955c9-cf6e-760f-8e42-02dbc6c9df51/",
    "DistributionTaskHref": null
  },
  "pulp": {
    "sync": {
      "pulp_href": "/pulp/api/v3/tasks/018955c9-b6e6-7a7b-8aa1-5a0a39bde143/",
      "pulp_created": "2023-07-14T19:03:50.323528Z",
      "state": "completed",
      "name": "pulp_rpm.app.tasks.synchronizing.synchronize",
      "logging_cid": "ea9504d2e73a4996be344abbc14dd1de",
      "started_at": "2023-07-14T19:03:50.446732Z",
      "finished_at": "2023-07-14T19:03:51.756474Z",
      "worker": "/pulp/api/v3/workers/01895528-d7c0-71d2-9eae-261375b1bf3c/",
      "progress_reports": [
        {
          "message": "Downloading Metadata Files",
          "code": "sync.downloading.metadata",
          "state": "completed",
          "done": 4,
          "suffix": null
        },
        {
          "message": "Downloading Artifacts",
          "code": "sync.downloading.artifacts",
          "state": "completed",
          "done": 0,
          "suffix": null
        },
        {
          "message": "Associating Content",
          "code": "associating.content",
          "state": "completed",
          "done": 1,
          "suffix": null
        },
        {
          "message": "Skipping Packages",
          "code": "sync.skipped.packages",
          "state": "completed",
          "total": 0,
          "done": 0,
          "suffix": null
        },
        {
          "message": "Parsed Packages",
          "code": "sync.parsing.packages",
          "state": "completed",
          "total": 1,
          "done": 1,
          "suffix": null
        },
        {
          "message": "Un-Associating Content",
          "code": "unassociating.content",
          "state": "completed",
          "done": 0,
          "suffix": null
        }
      ],
      "created_resources": [
        "/pulp/api/v3/repositories/rpm/rpm/018955c9-af6f-7dc5-ab6b-d78f2d1b4cda/versions/1/"
      ],
      "reserved_resources_record": [
        "/pulp/api/v3/repositories/rpm/rpm/018955c9-af6f-7dc5-ab6b-d78f2d1b4cda/",
        "shared:/pulp/api/v3/remotes/rpm/rpm/018955c9-9b81-7470-bddd-611d275499db/"
      ]
    },
    "publication": {
      "pulp_href": "/pulp/api/v3/tasks/018955c9-cf6e-760f-8e42-02dbc6c9df51/",
      "pulp_created": "2023-07-14T19:03:56.527134Z",
      "state": "completed",
      "name": "pulp_rpm.app.tasks.publishing.publish",
      "logging_cid": "c3247470e6a74e57a884d234c040da28",
      "started_at": "2023-07-14T19:03:56.653339Z",
      "finished_at": "2023-07-14T19:03:58.551824Z",
      "worker": "/pulp/api/v3/workers/01895528-d7c0-71d2-9eae-261375b1bf3c/",
      "progress_reports": [
        {
          "message": "Generating repository metadata",
          "code": "publish.generating_metadata",
          "state": "completed",
          "total": 1,
          "done": 1,
          "suffix": null
        }
      ],
      "created_resources": [
        "/pulp/api/v3/publications/rpm/rpm/018955c9-cff8-7d0f-beeb-f9f4826b49d8/"
      ],
      "reserved_resources_record": [
        "shared:/pulp/api/v3/repositories/rpm/rpm/018955c9-af6f-7dc5-ab6b-d78f2d1b4cda/"
      ]
    }
  }
}
```


### Comment by @swadeley on July 14, 2023 at 07:32 PM UTC

The Custom Repositories page:

**
![image](https://github.com/content-services/content-sources-backend/assets/11247237/3fb7fc32-8d9c-46db-8940-8f143d24b55a)
**

### Comment by @swadeley on July 14, 2023 at 07:33 PM UTC

The Task Details (when you click View Details)

![image](https://github.com/content-services/content-sources-backend/assets/11247237/785a8921-8abc-4f9c-93aa-5e9d54da2fa6)


### Comment by @swadeley on July 14, 2023 at 07:33 PM UTC

LGTM


thank you

---

## Reviews

### Review by @jlsherrill - Commented on June 23, 2023 at 04:37 PM UTC

### Review by @rverdile - Commented on June 23, 2023 at 04:47 PM UTC

### Review by @jlsherrill - Commented on June 23, 2023 at 04:58 PM UTC

### Review by @dpang314 - Commented on June 23, 2023 at 05:51 PM UTC

### Review by @rverdile - Commented on June 23, 2023 at 05:56 PM UTC

### Review by @rverdile - Commented on June 23, 2023 at 06:15 PM UTC

### Review by @dpang314 - Commented on June 23, 2023 at 06:21 PM UTC

### Review by @jlsherrill - Commented on July 03, 2023 at 08:27 PM UTC

### Review by @jlsherrill - Commented on July 03, 2023 at 08:35 PM UTC

### Review by @jlsherrill - Commented on July 14, 2023 at 02:28 PM UTC

### Review by @swadeley - Approved on July 14, 2023 at 07:45 PM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/316*
