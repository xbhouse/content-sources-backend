---
type: pull_request
number: 595
title: "Fixes 3490: update distributions for create/update templates"
state: merged
author: rverdile
created: 2024-03-05T16:27:23Z
updated: 2024-04-16T15:32:29Z
closed: 2024-04-16T15:32:24Z
merged: 2024-04-16T15:32:24Z
base_branch: main
head_branch: template-distributions
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/595
---

# Pull Request #595: Fixes 3490: update distributions for create/update templates

**Author**: @rverdile
**Created**: March 05, 2024 at 04:27 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `template-distributions`

## Description

## Summary
This creates/updates/deletes distributions in pulp when a template is created or updated (and repositories are added/removed).

See the JIRA card for the breakdown of the distribution paths. In summary, distributions are created under a path associated to each template. And the paths for red hat repos and custom repos are different.

## Testing steps
1. These testing steps are just a starting point, there's a lot of cases to test here.
2. Sync 2 repositories.
3. Create a template and include the first repository.
4. After the template is created, grab the distribution_href from the "templates_repository_configurations" table. Query pulp for that href. Make sure it is created and that the path matches the requirements of the JIRA card.
5. Patch the template, including the first and second repositories in the request. Check that the old repository's distribution persists and that a new one is created for the second repository.
6. Patch the template again, removing the second repository. The distribution for that repository should be deleted.
7. If you have the means to host your own repositories, you'll also want to test changing the date of the template. I'm not sure how to test this otherwise.

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on March 05, 2024 at 04:30 PM UTC

https://issues.redhat.com/browse/HMS-3490

### Comment by @swadeley on March 14, 2024 at 02:28 PM UTC

/retest

### Comment by @jlsherrill on March 21, 2024 at 10:07 AM UTC

few comments, otherwise looks good.  doing some testing

### Comment by @jlsherrill on March 21, 2024 at 01:48 PM UTC

seeing an issue where i tried to simulate having snapshots of different dates by:

1.  creating a custom repo, letting it snapshot
2. change the url and let it snapshot again
3. run this sql command (or similar) to update the first snapshot's created_at date:
```
update snapshots set created_at = '2024-03-19 09:41:24.568133+00' where uuid = 'f7454454-a6cd-47cd-a8d2-c34f34a707cd';
```
4.  Create a template and change the date to the 19th, or the 21st.  I'd expect the contents of the template path for that repo to change, but i don't see them change.  I added some debugging and it seems like the distribution is being updated with the same publication no matter which date i picked.  It always picks the one from the 19th

I thought it might be an issue with the FetchSnapshotsModelByDateAndRepository method, but when i call the api that uses it, that seems fine:

```

####
POST http://localhost:8000/api/content-sources/v1.0/snapshots/for_date/
x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiIxNzY4NDYzOCIsICJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJmb28ifSwiYWNjb3VudF9udW1iZXIiOiIxNzY4NDYzOCIsImludGVybmFsIjp7Im9yZ19pZCI6IjE3Njg0NjM4In19fQo=
Content-Type: application/json

{
  "repository_uuids": ["5079621c-c97c-4120-9aeb-60d68d3dcc29"],
  "date": "2024-03-20"

}
```

```
  {
    "repository_uuid": "5079621c-c97c-4120-9aeb-60d68d3dcc29",
    "is_after": false,
    "match": {
      "uuid": "67c68114-0a6e-4f50-a6f3-5d09e9419eaf",
      "created_at": "2024-03-21T06:08:49.060373-04:00",
      "repository_path": "cf98e9d7/5079621c-c97c-4120-9aeb-60d68d3dcc29/4866867f-3fe0-4c20-99f3-a0b342964bfa",
      "content_counts": {
        "rpm.package": 1
      },
```

### Comment by @jlsherrill on April 08, 2024 at 04:28 PM UTC

/retest

### Comment by @jlsherrill on April 08, 2024 at 05:27 PM UTC

i noticed that if i select a repo that has no snapshots, i get an error:  

```
1:25PM ERR recovered panic in worker with error: runtime error: index out of range [-1] error="runtime error: index out of range [-1]" request_id=vWLyHErsQHhairrylRETYxAZixLBEyKE task_id=c78c02f7-676d-4fd8-a0f9-b258143fde3e task_type=update-template-distributions
```

Its probably a situation we should warn against in the ui, but ideally the task wouldn't error and would just skip that repo

### Comment by @jlsherrill on April 08, 2024 at 06:11 PM UTC

just the issue with the error, otherwise this is working great!

---

## Reviews

### Review by @rverdile - Commented on March 07, 2024 at 07:41 PM UTC

### Review by @rverdile - Commented on March 07, 2024 at 07:44 PM UTC

### Review by @rverdile - Commented on March 07, 2024 at 07:46 PM UTC

### Review by @rverdile - Commented on March 07, 2024 at 07:49 PM UTC

### Review by @jlsherrill - Commented on March 12, 2024 at 12:59 PM UTC

### Review by @jlsherrill - Commented on March 12, 2024 at 01:15 PM UTC

### Review by @jlsherrill - Commented on March 12, 2024 at 05:23 PM UTC

### Review by @jlsherrill - Commented on March 12, 2024 at 05:25 PM UTC

### Review by @jlsherrill - Commented on March 12, 2024 at 05:28 PM UTC

### Review by @jlsherrill - Commented on March 12, 2024 at 05:31 PM UTC

### Review by @jlsherrill - Commented on March 12, 2024 at 06:03 PM UTC

### Review by @jlsherrill - Commented on March 12, 2024 at 06:06 PM UTC

### Review by @jlsherrill - Commented on March 12, 2024 at 06:13 PM UTC

### Review by @rverdile - Commented on March 13, 2024 at 07:34 PM UTC

### Review by @rverdile - Commented on March 13, 2024 at 07:35 PM UTC

### Review by @rverdile - Commented on March 13, 2024 at 07:35 PM UTC

### Review by @rverdile - Commented on March 13, 2024 at 07:36 PM UTC

### Review by @rverdile - Commented on March 13, 2024 at 07:36 PM UTC

### Review by @jlsherrill - Commented on March 21, 2024 at 08:26 AM UTC

### Review by @jlsherrill - Commented on March 21, 2024 at 08:37 AM UTC

### Review by @jlsherrill - Commented on March 21, 2024 at 08:45 AM UTC

### Review by @jlsherrill - Commented on March 21, 2024 at 08:46 AM UTC

### Review by @jlsherrill - Commented on March 21, 2024 at 01:18 PM UTC

### Review by @rverdile - Commented on March 22, 2024 at 07:04 PM UTC

### Review by @rverdile - Commented on March 22, 2024 at 07:12 PM UTC

### Review by @rverdile - Commented on March 22, 2024 at 07:13 PM UTC

### Review by @rverdile - Commented on March 22, 2024 at 08:01 PM UTC

### Review by @jlsherrill - Approved on April 11, 2024 at 03:55 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/595*
