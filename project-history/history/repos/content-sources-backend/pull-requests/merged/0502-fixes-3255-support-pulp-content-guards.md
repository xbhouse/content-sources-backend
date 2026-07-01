---
type: pull_request
number: 502
title: "Fixes 3255: support pulp content guards"
state: merged
author: jlsherrill
created: 2023-12-13T01:39:36Z
updated: 2024-01-03T12:32:04Z
closed: 2024-01-03T12:32:01Z
merged: 2024-01-03T12:32:01Z
base_branch: main
head_branch: 3255
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/502
---

# Pull Request #502: Fixes 3255: support pulp content guards

**Author**: @jlsherrill
**Created**: December 13, 2023 at 01:39 AM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `3255`

## Description

## Summary

When snapshotting a custom repository, we check to see if there are three content guards created in the org:
1.  content guard that checks for org_id on the identity header
2.  content guard that checks for a specific custom_dn on the identity header (used by turnpike)
3. content guard that Composites these together

If any of these 3 are not created, or not configured correctly, they will be updated or created.  The created pulp distribution is then associated with the content guard.

## Testing steps

1. Set config values for `guard_subject_dn` and `custom_repo_content_guards`
```
  pulp:
    download_policy: on_demand
    server: http://localhost:8080
    username: admin
    password: password
    storage_type: local #object or local
    guard_subject_dn: "some_dn"
    custom_repo_content_guards: true
```

you can change 'some_dn' to whatever you want, but remember what you set it to!

5. Import a small redhat repo:
``` 
OPTIONS_REPOSITORY_IMPORT_FILTER=small go run cmd/external-repos/main.go import
```

6.  Run the server:  ```make run```

3.  Trigger sync:  ``` go run cmd/external-repos/main.go nightly-jobs```
7.  create a custom repo with snapshotting, in the UI and let it snapshot
8.  Grab the org id from the db (ignore -1):  

`make db-cli-connect`  
```
select org_id from where repository_configurations;
```

9.  Fetch the Red Hat repo without anything specified.  Navigate to the red hat repos, you should have one repo, list snapshots and grab the repo config.  Grab the URL for the repository.  For local development, change the hostname and port to localhost:8080. for example:

http://localhost:8080/pulp/content/abcde/fghjkl/

You should be able to curl that url and get valid html (200 Response), as well as fetch its ./repodata/repomd.xml
```
curl http://localhost:8080/pulp/content/abcde/fghjkl/repodata/repomd.xml
```

10. Get unauthorized when fetching the custom repo.  Similar to the redhat repo, grab the snapshot URL  and change the host/port and try to curl it, you should get an unauthorized response
11.  Now fetch using a cert with a proper org id with your org_id:
```
curl  http://localhost:8080/pulp/content/abcde/fghjkl/repodata/repomd.xml   -H "`./scripts/header.sh $ORG_ID 123 jdoe`"
```
12.  Now fetch using a 'turnpike' header:

```
curl  http://localhost:8080/pulp/content/abcde/fghjkl/repodata/repomd.xml   -H "`./scripts/turnpike_header.sh $some_dn`"
```
replacing some_dn with whatever you put in your config at step 1

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on December 13, 2023 at 02:00 AM UTC

https://issues.redhat.com/browse/HMS-3255

### Comment by @xbhouse on December 20, 2023 at 05:01 PM UTC

just testing functionality. getting the expected 200 with the RH repo. but i'm seeing 403s when fetching the custom repo using both my org id and the DN. i have `guard_subject_dn: "test_dn"` in the config. maybe there's something i missed:

`curl -v http://localhost:8080/pulp/content/a4178861/34e04ea8-4999-444c-9405-026e23d58a44/56ca6e62-1890-4089-b887-488eb04addf2/repodata/repomd.xml -H "./scripts/header.sh 16799052 11491789 bryttaniehouse"`

`< HTTP/1.1 403 Access forbidden. Refusals: ["Guard: 'org_id_guard', HREF: '/pulp/a4178861/api/v3/contentguards/core/header/018c8818-4bf3-791a-8bfb-c191f97009bd/', class: '<class 'pulpcore.app.models.publication.HeaderContentGuard'>', denial: [Access denied.].", "Guard: 'turnpike_guard', HREF: '/pulp/a4178861/api/v3/contentguards/core/header/018c8818-4ef7-7a44-b1b7-46e7b2632ced/', class: '<class 'pulpcore.app.models.publication.HeaderContentGuard'>', denial: [Access denied.]."`

`curl -v http://localhost:8080/pulp/content/a4178861/34e04ea8-4999-444c-9405-026e23d58a44/56ca6e62-1890-4089-b887-488eb04addf2/repodata/repomd.xml -H "./scripts/turnpike_header.sh $test_dn"`

`< HTTP/1.1 403 Access forbidden. Refusals: ["Guard: 'org_id_guard', HREF: '/pulp/a4178861/api/v3/contentguards/core/header/018c8818-4bf3-791a-8bfb-c191f97009bd/', class: '<class 'pulpcore.app.models.publication.HeaderContentGuard'>', denial: [Access denied.].", "Guard: 'turnpike_guard', HREF: '/pulp/a4178861/api/v3/contentguards/core/header/018c8818-4ef7-7a44-b1b7-46e7b2632ced/', class: '<class 'pulpcore.app.models.publication.HeaderContentGuard'>', denial: [Access denied.]."`

### Comment by @xbhouse on December 20, 2023 at 07:32 PM UTC

ok works as expected with these commands :) (not formatted correctly in the comment but the backticks are needed around the header, d'oh!):

curl -v http://localhost:8080/pulp/content/a4178861/34e04ea8-4999-444c-9405-026e23d58a44/56ca6e62-1890-4089-b887-488eb04addf2/repodata/repomd.xml -H "`./scripts/turnpike_header.sh test_dn`"

curl -v http://localhost:8080/pulp/content/a4178861/34e04ea8-4999-444c-9405-026e23d58a44/56ca6e62-1890-4089-b887-488eb04addf2/repodata/repomd.xml -H "`./scripts/header.sh 16799052 11491786 11491786`"

and integration tests pass when the server isn't running. thank you for the help lol @jlsherrill 

### Comment by @xbhouse on December 20, 2023 at 07:32 PM UTC

i've just tested functionality and the code looks fine to me, but i think it would be good if someone else reviewed as well @Andrewgdewar @rverdile 

### Comment by @jlsherrill on January 02, 2024 at 06:14 PM UTC

added!

### Comment by @jlsherrill on January 03, 2024 at 12:31 PM UTC

merging as this is actually no-qe-needed, bot hasn't updated yet

---

## Reviews

### Review by @rverdile - Commented on December 21, 2023 at 08:30 PM UTC

This looks good. I'll just reiterate my one thought, which was to add a bool on the snapshot so it's easier to tell which snapshots have content guards.

### Review by @rverdile - Approved on January 02, 2024 at 07:30 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/502*
