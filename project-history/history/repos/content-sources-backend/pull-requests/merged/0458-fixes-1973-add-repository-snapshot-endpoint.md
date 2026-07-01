---
type: pull_request
number: 458
title: "Fixes 1973: Add repository snapshot endpoint"
state: merged
author: Andrewgdewar
created: 2023-11-02T21:03:17Z
updated: 2024-01-02T14:12:10Z
closed: 2023-12-19T15:52:24Z
merged: 2023-12-19T15:52:24Z
base_branch: main
head_branch: HMS-1973
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/458
---

# Pull Request #458: Fixes 1973: Add repository snapshot endpoint

**Author**: @Andrewgdewar
**Created**: November 02, 2023 at 09:03 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `HMS-1973`

## Description

## Summary

Adds a repository snapshot endpoint for manually triggering snapshots

## Testing steps

**Method 1**

Easiest method of testing would be to change the following lines in the FE > ContentApi.ts

<img width="728" alt="Screenshot 2023-11-02 at 2 54 47 PM" src="https://github.com/content-services/content-sources-backend/assets/38083295/63406c01-33c3-40b4-9cc1-b1c4fd47c821">

And disable this one line in the ContentListTable

<img width="752" alt="Screenshot 2023-11-02 at 2 55 52 PM" src="https://github.com/content-services/content-sources-backend/assets/38083295/1a35011f-ceab-4de1-ac6e-bcdcdfd234b1">

So that when you click the "Introspect Now" button in the kebab, it is actually requesting to sync a snapshot.

**Method 2**
Alternatively call the new endpoint directly: /repositories/{uuid}/snapshot/ [post] 
You'll need to have created a repository.

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate 


---

## Discussion

### Comment by @jlsherrill on November 02, 2023 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-1973

### Comment by @Andrewgdewar on November 03, 2023 at 02:42 PM UTC

/retest

### Comment by @jlsherrill on November 27, 2023 at 02:54 PM UTC

Two small changes, and then this is good.

### Comment by @swadeley on December 03, 2023 at 11:52 PM UTC

Hi @Andrewgdewar , rebase please.

### Comment by @swadeley on December 04, 2023 at 06:32 AM UTC

Hi @Andrewgdewar 

I updated API docs.

I created a repo with `'snapshot': False,`

`In [3]: app.content_sources.rest_client.repositories_api.create_repository(repo)
Out[3]: 
{'account_id': '0369233',
 'content_type': 'rpm',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'last_introspection_error': '',
 'last_introspection_time': '',
 'last_success_introspection_time': '',
 'last_update_introspection_time': '',
 'metadata_verification': False,
 'name': 'fedoratest-snapshot-false',
 'org_id': '3340851',
 'origin': 'external',
 'package_count': 0,
 'snapshot': False,
 'status': 'Pending',
 'url': 'https://stephenw.fedorapeople.org/multirepos/5/repo06/',
 'uuid': 'ca7faa4a-eab8-4fc4-80f4-d46e4c97fc42'}`

I created a snapshot:

`In [7]: app.content_sources.rest_client.repositories_api.create_snapshot('ca7faa4a-eab8-4fc4-80f4-d46e4c97fc42')

In [8]: app.content_sources.rest_client.repositories_api.get_repository('ca7faa4a-eab8-4fc4-80f4-d46e4c97fc42')
2023-12-04 14:24:16.586 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=yIaOKulwjvcCShJyfiTbpIaOsCRmUhmN, params=[]
Out[8]: 
{'account_id': '0369233',
 'content_type': 'rpm',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'last_introspection_error': '',
 'last_introspection_time': '2023-12-04T06:18:24Z',
 'last_snapshot': {'added_counts': {'rpm.package': 1},
                   'content_counts': {'rpm.package': 1},
                   'created_at': '2023-12-04T06:23:54.714368Z',
                   'removed_counts': {},
                   'repository_path': 'ff180fe7/ca7faa4a-eab8-4fc4-80f4-d46e4c97fc42/6f07466d-b03f-4f48-9865-ab164dc0d48b',
                   'url': 'https://env-ephemeral-gbpyh4.<snip>.openshiftapps.com/pulp/content/ff180fe7/ca7faa4a-eab8-4fc4-80f4-d46e4c97fc42/6f07466d-b03f-4f48-9865-ab164dc0d48b/',
                   'uuid': 'fb3fc5e0-182c-43d4-8bfe-5b719c94f1b3'},
 'last_snapshot_task_uuid': '9b3fa3df-86ee-43de-a3db-70d3befe450f',
 'last_snapshot_uuid': 'fb3fc5e0-182c-43d4-8bfe-5b719c94f1b3',
 'last_success_introspection_time': '2023-12-04T06:18:24Z',
 'last_update_introspection_time': '2023-12-04T06:18:24Z',
 'metadata_verification': False,
 'name': 'fedoratest-snapshot-false',
 'org_id': '3340851',
 'origin': 'external',
 'package_count': 1,
 'snapshot': False,
 'status': 'Valid',
 'url': 'https://stephenw.fedorapeople.org/multirepos/5/repo06/',
 'uuid': 'ca7faa4a-eab8-4fc4-80f4-d46e4c97fc42'}

In [9]: 
`

When i check the UI, I see it says "Last snapshot 2 minutes ago", but when I check the kebab menu it says "No snapshots yet".  You could say that is out of scope and fix that later? A front-end ch ange?

### Comment by @swadeley on December 04, 2023 at 06:50 AM UTC

Hi, I created another repo with `'snapshot': True,` . I waited a few mins for the initial snapshot to be created. I then used API to try and create another snapshot but no Admin tasks or changes were made. UI says Last snapshot 15 minutes ago.

EDIT: I see now and admin taks with `message: "Skipping Sync (no change from previous sync)"
code: "sync.was_skipped"`
So, it did try,

### Comment by @Andrewgdewar on December 04, 2023 at 08:24 PM UTC

> When i check the UI, I see it says "Last snapshot 2 minutes ago", but when I check the kebab menu it says "No snapshots yet". You could say that is out of scope and fix that later? A front-end change?

Hmmm the described issue sounds almost like the snapshot was taken but unsuccessful, if the "last_snapshot" variable exists and is not empty, the snapshot list dropdown item should be populated. 
As the last snapshot time stamp is within the "last_snapshot", I'm not sure how this could happen.. hmm will attempt to recreate.

### Comment by @Andrewgdewar on December 18, 2023 at 06:20 PM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Commented on November 06, 2023 at 06:06 PM UTC

### Review by @Andrewgdewar - Commented on November 06, 2023 at 06:29 PM UTC

### Review by @jlsherrill - Commented on November 27, 2023 at 02:32 PM UTC

### Review by @jlsherrill - Commented on November 27, 2023 at 02:54 PM UTC

### Review by @Andrewgdewar - Commented on November 29, 2023 at 04:40 PM UTC

### Review by @jlsherrill - Approved on November 29, 2023 at 04:41 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/458*
