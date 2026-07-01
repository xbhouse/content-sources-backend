---
type: pull_request
number: 221
title: "Refs 533: Add repository introspection endpoint"
state: merged
author: rverdile
created: 2023-02-23T19:58:57Z
updated: 2023-04-14T19:04:21Z
closed: 2023-04-05T18:58:35Z
merged: 2023-04-05T18:58:35Z
base_branch: main
head_branch: introspect-endpoint
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/221
---

# Pull Request #221: Refs 533: Add repository introspection endpoint

**Author**: @rverdile
**Created**: February 23, 2023 at 07:58 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `introspect-endpoint`

## Description

## Summary
This will also have a frontend PR, but I'll put this up as a draft in the meantime for judgement.

- Adds `/repository/{uuid}/introspect` endpoint to introspect a repository associated with the given repository configuration UUID.
- Adds `failed_introspection_count` to the `repositories` table, which represents the number of consecutive failed introspections and is used to skip introspection if a certain limit is reached.
- Adds a `reset_count` boolean param to this endpoint that will reset the `failed_introspections_count` column back to zero.
- Changed format for timestamps returned by API to be RFC3339.
## Testing steps
1. Migrate DB
2. Set the `OPTIONS_INTROSPECT_API_TIME_LIMIT_SEC` environment variable to 10. This should make it should you cannot introspect a repository more than once every 10 seconds. You should also test setting this  option in config.yaml, `options.introspect_api_time_limit_sec`.
3. Introspect an invalid repository with the endpoint, without passing the reset_count param. The failed introspections count should increase. You can verify this by fetching the repository from its GET endpoint.
4. Introspect an invalid repository and pass the reset_count param. The failed introspections count should become 1 after introspection fails.
6. Introspect a valid repository that was failing previously. The failed introspections count should reset to 0.
7. Introspect a repository over the FailedIntrospectionsLimit (20), using the CLI, without using the `--force` flag. It should be skipped. Doing that again, with the `--force` flag, should return an error. 


---

## Discussion

### Comment by @jlsherrill on February 23, 2023 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-533

### Comment by @jlsherrill on February 23, 2023 at 08:00 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @jlsherrill on February 23, 2023 at 09:27 PM UTC

we'll need to make sure that the metrics from here https://github.com/content-services/content-sources-backend/pull/220/files#diff-f171c549d2735774b148a59f7f39861d8e58abe0a6f81735ea7cf933749c72b0R53-R77  get updated depending on which pr goes in first.  

### Comment by @rverdile on March 22, 2023 at 03:00 PM UTC

I haven't addressed your comments yet. This latest push is for the Retrying status addition.

### Comment by @jlsherrill on March 22, 2023 at 03:33 PM UTC

i *think* you need to add an rbac entry here: https://github.com/content-services/content-sources-backend/blob/main/pkg/middleware/rbac_map.go#L22-L35

I filed this task to try to improve this:  https://issues.redhat.com/browse/HMS-1486

### Comment by @jlsherrill on March 22, 2023 at 08:12 PM UTC

Oh and we also need to adjust https://github.com/content-services/content-sources-backend/blob/main/pkg/dao/metrics.go#L59-L89  to ignore repos that are beyond the 'failed count', as we don't want to include them in the stats

### Comment by @rverdile on March 31, 2023 at 06:13 PM UTC

From [this discussion](https://github.com/content-services/content-sources-frontend/pull/104#discussion_r1154611700) on the front PR, I'm thinking it may have been unnecessary for me to add a 'Retrying' status.

I don't think I had another reason for adding it. I can't think of one at the moment, other than maybe it will be useful at some point to handle them separately. If it makes sense to remove it, I'll remove it.

### Comment by @rverdile on April 03, 2023 at 02:25 PM UTC

Removed "Retrying" status

### Comment by @swadeley on April 05, 2023 at 06:48 PM UTC

Hi

I see the failed introspections count  (introspections, plural):

```
Out[3]: 
{'data': [{'account_id': '0369233',
           'distribution_arch': 'any',
           'distribution_versions': ['any'],
           'failed_introspections_count': 2,
           'gpg_key': '',
           'last_introspection_error': 'Cannot fetch '
                                       'https://fixtures.pulpproject.org/rpm-invalid-rpm/repodata/repomd.xml: '
                                       '404',
           'last_introspection_time': '2023-04-05 16:04:47.668374 +0000 UTC',
           'last_success_introspection_time': '',
           'last_update_introspection_time': '',
           'metadata_verification': False,
           'name': 'rpm-invalid-rpm',
           'org_id': '3340851',
           'package_count': 0,
           'status': 'Invalid',
           'url': 'https://fixtures.pulpproject.org/rpm-invalid-rpm/',
           'uuid': '1a0a8537-d286-4612-9209-5a852b1cf4c5'}],
```

### Comment by @swadeley on April 05, 2023 at 06:58 PM UTC

Hi,
 the code says `"/repositories/{uuid}/introspect/"` not `/repository/{uuid}/introspect ` as per comment 0

Anyway,

I had :
`last_introspection_time': '2023-04-05 16:04:47.668374 +0000 UTC',`

and this:
`In [8]: app.content_sources.rest_client.repositories_api.introspect('1a0a8537-d286-4612-9209-5a852b1cf4c5')
2023-04-05 19:55:02.624 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=TEObNNjI64UjTHSwJkJuhxNzQve6nOMk, params=[]`

Gave me:

` 'last_introspection_time': '2023-04-05 18:55:02.786693 +0000 UTC',`

---

## Reviews

### Review by @jlsherrill - Commented on February 23, 2023 at 09:28 PM UTC

### Review by @rverdile - Commented on March 20, 2023 at 02:03 PM UTC

### Review by @rverdile - Commented on March 20, 2023 at 02:09 PM UTC

### Review by @jlsherrill - Commented on March 22, 2023 at 12:45 PM UTC

### Review by @jlsherrill - Commented on March 22, 2023 at 01:05 PM UTC

### Review by @jlsherrill - Commented on March 22, 2023 at 01:18 PM UTC

### Review by @rverdile - Commented on March 22, 2023 at 03:08 PM UTC

### Review by @rverdile - Commented on March 22, 2023 at 03:12 PM UTC

### Review by @jlsherrill - Commented on March 27, 2023 at 06:59 PM UTC

### Review by @jlsherrill - Commented on March 27, 2023 at 07:21 PM UTC

### Review by @rverdile - Commented on March 27, 2023 at 08:20 PM UTC

### Review by @rverdile - Commented on March 28, 2023 at 06:09 PM UTC

### Review by @jlsherrill - Approved on April 05, 2023 at 12:50 PM UTC

### Review by @rverdile - Commented on April 05, 2023 at 03:46 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/221*
