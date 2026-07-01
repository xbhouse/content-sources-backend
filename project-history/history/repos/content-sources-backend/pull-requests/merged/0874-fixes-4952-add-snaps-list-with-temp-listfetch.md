---
type: pull_request
number: 874
title: "Fixes 4952: add snaps list with temp list/fetch"
state: merged
author: jlsherrill
created: 2024-11-05T21:08:21Z
updated: 2024-11-27T18:30:32Z
closed: 2024-11-27T17:32:12Z
merged: 2024-11-27T17:32:12Z
base_branch: main
head_branch: 4952
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/874
---

# Pull Request #874: Fixes 4952: add snaps list with temp list/fetch

**Author**: @jlsherrill
**Created**: November 05, 2024 at 09:08 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4952`

## Description

## Summary

Adds snapshot list to template list and fetch api calls.  

## Testing steps

1.  Create a repository, let it snapshot
2. Create a template with that repository:
```
####
POST http://localhost:8000/api/content-sources/v1.0/templates/
x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==
Content-Type: application/json

{
  "name":"test3",
  "arch": "x86_64",
  "version": "9",
  "repository_uuids": ["48b89a40-7d3d-4bd8-8596-f20d77c42032"],
  "date": "2024-09-10T15:09:43-04:00"
}
```

3.  Fetch/list the template and see the snapshot list:
```

{
  "data": [
    {
      "uuid": "01bc2b42-b74c-47f4-a9bd-1da2e8b3e941",
      "name": "test3",
      "org_id": "9",
      "description": "",
      "arch": "x86_64",
      "version": "9",
      "date": "2024-09-10T19:09:43Z",
      "repository_uuids": [
        "48b89a40-7d3d-4bd8-8596-f20d77c42032"
      ],
      "snapshots": [
        {
          "uuid": "4c8a35e0-56d4-46bb-b875-2463a954c6cc",
          "created_at": "2024-11-04T13:50:19.182357Z",
          "repository_path": "48e64014/48b89a40-7d3d-4bd8-8596-f20d77c42032/78f2a819-46c8-4188-8403-dfe3f145180d",
          "content_counts": {
            "rpm.advisory": 1235,
            "rpm.package": 7326,
            "rpm.packagecategory": 4,
            "rpm.packageenvironment": 3,
            "rpm.packagegroup": 34,
            "rpm.repo_metadata_file": 1
          },
          "added_counts": {
            "rpm.advisory": 1235,
            "rpm.package": 7326,
            "rpm.packagecategory": 4,
            "rpm.packageenvironment": 3,
            "rpm.packagegroup": 34,
            "rpm.repo_metadata_file": 1
          },
          "removed_counts": {},
          "url": "",
          "repository_name": "",
          "repository_uuid": ""
        }
      ],
      "rhsm_environment_id": "01bc2b42b74c47f4a9bd1da2e8b3e941",
      "created_by": "foo",
      "last_updated_by": "foo",
      "created_at": "2024-11-05T20:10:36.695726Z",
      "updated_at": "2024-11-05T20:10:36.695726Z",
      "use_latest": false,
      "last_update_snapshot_error": "",
      "last_update_task_uuid": "9fe2020f-6499-40b1-9723-43385fccbf22",
      "last_update_task": {
        "uuid": "9fe2020f-6499-40b1-9723-43385fccbf22",
        "status": "failed",
        "created_at": "2024-11-05T20:13:37Z",
        "ended_at": "2024-11-05T20:13:37Z",
        "error": "",
        "org_id": "9",
        "type": "update-template-content",
        "object_type": "template",
        "object_name": "test3",
        "object_uuid": "01bc2b42-b74c-47f4-a9bd-1da2e8b3e941"
      },
      "rhsm_environment_created": true
    }
  ],
  "meta": {
    "limit": 100,
    "offset": 0,
    "count": 1
  },
  "links": {
    "first": "/api/content-sources/v1.0/templates/?limit=100&offset=0",
    "last": "/api/content-sources/v1.0/templates/?limit=100&offset=0"
  }
}
```


One thing to note is that the snapshot list isn't populated at creation time, but instead via the update_content_template task.  I'm not sure we can do much at this time about it.


---

## Discussion

### Comment by @jlsherrill on November 05, 2024 at 09:12 PM UTC

note this includes some changes from https://github.com/content-services/content-sources-backend/pull/865   so that will have to go in first and this be rebased.

### Comment by @jlsherrill on November 05, 2024 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-4952

### Comment by @swadeley on November 13, 2024 at 09:45 AM UTC

Hi @jlsherrill , please rebase.

### Comment by @jlsherrill on November 14, 2024 at 04:04 PM UTC

[test]

### Comment by @swadeley on November 15, 2024 at 12:13 PM UTC

/retest

### Comment by @swadeley on November 15, 2024 at 08:22 PM UTC

/retest

### Comment by @swadeley on November 18, 2024 at 08:55 AM UTC

/retest

### Comment by @jlsherrill on November 18, 2024 at 07:14 PM UTC

[test]

### Comment by @swadeley on November 19, 2024 at 08:36 AM UTC

/retest

### Comment by @swadeley on November 21, 2024 at 08:54 AM UTC

```
08:37:51 go build  -o "/go/src/app/release/candlepin" "cmd/candlepin/main.go"
08:38:17 # github.com/content-services/content-sources-backend/pkg/dao
08:38:17 pkg/dao/templates.go:632:3: undefined: snapshotModelToApi
08:38:17 make: *** [mk/go-rules.mk:26: /go/src/app/release/candlepin] Error 1
08:38:20 Error: building at STEP "RUN make get-deps build": while running runtime: exit status 2
```

### Comment by @swadeley on November 21, 2024 at 08:54 AM UTC

/retest

### Comment by @jlsherrill on November 27, 2024 at 05:00 PM UTC

/retest

### Comment by @jlsherrill on November 27, 2024 at 05:01 PM UTC

ephemeral test:
```
$ curl -u jdoe:FRkJbTAilFzAuVwy  https://ee-shmxpio1.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/  | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1339  100  1339    0     0   6606      0 --:--:-- --:--:-- --:--:--  6628
{
  "data": [
    {
      "uuid": "7b637900-1c96-4b69-a20b-9d8bad43270f",
      "name": "fooz",
      "org_id": "12345",
      "description": "",
      "arch": "x86_64",
      "version": "9",
      "date": "0001-01-01T00:00:00Z",
      "repository_uuids": [
        "3e4403ab-4b40-4b6f-8580-8592f71fafac"
      ],
      "snapshots": [
        {
          "uuid": "26e362ad-2b90-43f7-a8d5-8d404bd29bab",
          "created_at": "2024-11-27T16:55:07.632839Z",
          "repository_path": "cs-redhat/3e4403ab-4b40-4b6f-8580-8592f71fafac/5ca477d8-0765-48a5-be21-6dfcdd70409a",
          "content_counts": {
            "rpm.advisory": 32,
            "rpm.package": 58,
            "rpm.repo_metadata_file": 1
          },
          "added_counts": {
            "rpm.advisory": 32,
            "rpm.package": 58,
            "rpm.repo_metadata_file": 1
          },
          "removed_counts": {},
          "url": "http://pulp-content:8000/api/pulp-content/cs-redhat/3e4403ab-4b40-4b6f-8580-8592f71fafac/5ca477d8-0765-48a5-be21-6dfcdd70409a/",
          "repository_name": "Red Hat Ansible Engine 2 for RHEL 8 x86_64 (RPMs)",
          "repository_uuid": "3e4403ab-4b40-4b6f-8580-8592f71fafac"
        }
      ],
      "rhsm_environment_id": "7b6379001c964b69a20b9d8bad43270f",
      "created_by": "jdoe",
      "last_updated_by": "jdoe",
      "created_at": "2024-11-27T16:59:38.157021Z",
      "updated_at": "2024-11-27T16:59:38.157021Z",
      "use_latest": false,
      "last_update_snapshot_error": "",
      "rhsm_environment_created": false
    }
  ],
  "meta": {
    "limit": 100,
    "offset": 0,
    "count": 1
  },
  "links": {
    "first": "/api/content-sources/v1/templates/?limit=100&offset=0",
    "last": "/api/content-sources/v1/templates/?limit=100&offset=0"
  }
}



```




```
curl -u jdoe:FRkJbTAilFzAuVwy  https://ee-shmxpio1.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/7b637900-1c96-4b69-a20b-9d8bad43270f  | jq


{
  "uuid": "7b637900-1c96-4b69-a20b-9d8bad43270f",
  "name": "fooz",
  "org_id": "12345",
  "description": "",
  "arch": "x86_64",
  "version": "9",
  "date": "0001-01-01T00:00:00Z",
  "repository_uuids": [
    "3e4403ab-4b40-4b6f-8580-8592f71fafac"
  ],
  "snapshots": [
    {
      "uuid": "26e362ad-2b90-43f7-a8d5-8d404bd29bab",
      "created_at": "2024-11-27T16:55:07.632839Z",
      "repository_path": "cs-redhat/3e4403ab-4b40-4b6f-8580-8592f71fafac/5ca477d8-0765-48a5-be21-6dfcdd70409a",
      "content_counts": {
        "rpm.advisory": 32,
        "rpm.package": 58,
        "rpm.repo_metadata_file": 1
      },
      "added_counts": {
        "rpm.advisory": 32,
        "rpm.package": 58,
        "rpm.repo_metadata_file": 1
      },
      "removed_counts": {},
      "url": "http://pulp-content:8000/api/pulp-content/cs-redhat/3e4403ab-4b40-4b6f-8580-8592f71fafac/5ca477d8-0765-48a5-be21-6dfcdd70409a/",
      "repository_name": "Red Hat Ansible Engine 2 for RHEL 8 x86_64 (RPMs)",
      "repository_uuid": "3e4403ab-4b40-4b6f-8580-8592f71fafac"
    }
  ],
  "rhsm_environment_id": "7b6379001c964b69a20b9d8bad43270f",
  "created_by": "jdoe",
  "last_updated_by": "jdoe",
  "created_at": "2024-11-27T16:59:38.157021Z",
  "updated_at": "2024-11-27T16:59:38.157021Z",
  "use_latest": false,
  "last_update_snapshot_error": "",
  "rhsm_environment_created": false
}

```

### Comment by @jlsherrill on November 27, 2024 at 05:32 PM UTC

tests are green, merging

---

## Reviews

### Review by @xbhouse - Approved on November 26, 2024 at 04:16 PM UTC

this looks good! just needs a rebase once [this](https://github.com/content-services/content-sources-backend/pull/906) is merged and an api spec update in iqe

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/874*
