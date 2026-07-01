---
type: pull_request
number: 342
title: "Fixes 2188: Add more detail in snapshot listing"
state: merged
author: Andrewgdewar
created: 2023-08-01T15:27:44Z
updated: 2023-09-04T14:00:46Z
closed: 2023-08-29T09:49:56Z
merged: 2023-08-29T09:49:56Z
base_branch: main
head_branch: HMS-2188
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/342
---

# Pull Request #342: Fixes 2188: Add more detail in snapshot listing

**Author**: @Andrewgdewar
**Created**: August 01, 2023 at 03:27 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `HMS-2188`

## Description

## Summary
Adds more detail to Snapshots and Repository response: 

- Adds added_counts & removed_counts to the snapshots table/SnapshotResponse

- Adds last_snapshot_uuid and last_snapshot to the repository_configurations table/RepositoryResponse

## Testing steps
- Add a repository with snapshotting enabled. 
- Fetch/list/edit said repository after it has finished snapshotting, one should see the above new fields on the response.
- Call the snapshot list api, one should see the associated fields above on the response.


---

## Discussion

### Comment by @jlsherrill on August 01, 2023 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-2188

### Comment by @jlsherrill on August 04, 2023 at 02:03 PM UTC

You'll need to rebase and update your migrations to a later timestamp.  The easiest way is:
```
go run cmd/dbmigrate/main.go new my_migration_name
mv db/migrations/20230717141727_my_migration_name.down.sql db/migrations/20230804092757_my_migration_name.down.sql 
mv db/migrations/20230717141727_my_migration_name.up.sql db/migrations/20230804092757_my_migration_name.up.sql 
```

### Comment by @jlsherrill on August 16, 2023 at 09:33 PM UTC

The failure is due to:
* RepoConfig dao fetchRepoConfig()  calling .Preload("LastSnapshot")
* Update calling fetchRepoConfig() and then passing it to gorm: ``` if err := tx.Model(&repoConfig).Updates(repoConfig.MapForUpdate()).Error; err != nil { ```

The documentation indicates you can omit those associations on the Updates() call by using .Omit(): https://gorm.io/docs/associations.html#Skip-Auto-Create-x2F-Update

however, i couldn't get this to work at all.  Maybe you'll have better luck?

Another option is to make fetchRepoConfig()  have an parameter to load associations or not.  Most things doing a fetch would want to, but something like Update() wouldn't.  



### Comment by @swadeley on August 29, 2023 at 08:46 AM UTC

/retest

### Comment by @swadeley on August 29, 2023 at 09:40 AM UTC

Hi

Looks good:
In [2]: app.content_sources.rest_client.repositories_api.list_repositories()
Out[2]: 
{'data': [{'account_id': '0369233',
           'distribution_arch': 'any',
           'distribution_versions': ['any'],
           'failed_introspections_count': 0,
           'gpg_key': '',
           'last_introspection_error': '',
           'last_introspection_time': '2023-08-29T09:18:16Z',
           **'last_snapshot'**: {'added_counts': {'rpm.package': 1},
                             'content_counts': {'rpm.package': 1},
                             'created_at': '2023-08-29T09:18:34.323282Z',
                             'removed_counts': {},
                             'repository_path': ''},
           '**last_snapshot_uuid**': 'add2ca04-e88c-4e38-b59c-ba0b1b7c5ab8',
           'last_success_introspection_time': '2023-08-29T09:18:16Z',
           'last_update_introspection_time': '2023-08-29T09:18:16Z',
           'metadata_verification': False,
           'name': 'fedoratest-1',
           'org_id': '3340851',
           'package_count': 1,
           'snapshot': True,
           'status': 'Valid',
           'url': 'https://stephenw.fedorapeople.org/multirepos/5/repo01/',
           'uuid': 'f92d4f80-3ab6-44ae-b056-d00a1ed7bc09'}],
 'links': {'first': '/api/content-sources/v1/repositories/?limit=100&offset=0',
           'last': '/api/content-sources/v1/repositories/?limit=100&offset=0'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}


### Comment by @swadeley on August 29, 2023 at 09:45 AM UTC

and with GET

In [5]: app.content_sources.rest_client.repositories_api.get_repository('f92d4f80-3ab6-44ae-b056-d00a1ed7bc09')
2023-08-29 10:43:27.383 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=KR44MzIVbH31A877EEkvx6M7Ef7F07Ba, params=[]
Out[5]: 
{'account_id': '0369233',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'last_introspection_error': '',
 'last_introspection_time': '2023-08-29T09:18:16Z',
 '**last_snapshot**': {'added_counts': {'rpm.package': 1},
                   'content_counts': {'rpm.package': 1},
                   'created_at': '2023-08-29T09:18:34.323282Z',
                   'removed_counts': {},
                   'repository_path': ''},
 '**last_snapshot_uuid**': 'add2ca04-e88c-4e38-b59c-ba0b1b7c5ab8',
 'last_success_introspection_time': '2023-08-29T09:18:16Z',
 'last_update_introspection_time': '2023-08-29T09:18:16Z',
 'metadata_verification': False,
 'name': 'fedoratest-1',
 'org_id': '3340851',
 'package_count': 1,
 'snapshot': True,
 'status': 'Valid',
 'url': 'https://stephenw.fedorapeople.org/multirepos/5/repo01/',
 'uuid': 'f92d4f80-3ab6-44ae-b056-d00a1ed7bc09'}


---

## Reviews

### Review by @jlsherrill - Commented on August 16, 2023 at 08:43 PM UTC

### Review by @jlsherrill - Commented on August 24, 2023 at 07:31 PM UTC

### Review by @jlsherrill - Commented on August 24, 2023 at 07:32 PM UTC

### Review by @jlsherrill - Approved on August 28, 2023 at 08:06 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/342*
