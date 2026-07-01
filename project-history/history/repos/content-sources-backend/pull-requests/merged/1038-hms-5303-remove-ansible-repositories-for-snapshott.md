---
type: pull_request
number: 1038
title: "HMS-5303: remove ansible repositories for snapshotting"
state: merged
author: jlsherrill
created: 2025-03-19T12:54:29Z
updated: 2025-03-28T20:21:05Z
closed: 2025-03-28T20:20:51Z
merged: 2025-03-28T20:20:51Z
base_branch: main
head_branch: 5303
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1038
---

# Pull Request #1038: HMS-5303: remove ansible repositories for snapshotting

**Author**: @jlsherrill
**Created**: March 19, 2025 at 12:54 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `5303`

## Description

## Summary

This adds support for removing Red Hat repos for syncing.  (Currently they stay for introspection).  

Some changes needed to support this

* within the delete-repository-snapshots task, move the candlepin processing prior to the pulp processing.  This is due to some models being deleted during the pulp processing (it also makes sense to do things in reverse upon deletion).
* Looking up the red hat content ids from their labels
* Looking up templates to remove the repo from across all orgs, not just templates with the same org as the repository

## Testing steps

setup, i'd recommend on a clean db.  On Main, run:

```
make compose-clean compose-up
make repos-import
```
start the server:
```
make run
```
snapshot all the repos:
```
go run cmd/external-repos/main.go process-repos
```

Create a template with RHEL 8, adding the ansible repo to the template.



Check the pulp db:
```
psql --host localhost --port 5432  -U pulp
select base_path from core_distribution;
```

look for:

```
 templates/341aaa4f-6050-4599-99d8-01e81bf4bf5c/content/dist/rhel8/8/x86_64/baseos/os
 templates/341aaa4f-6050-4599-99d8-01e81bf4bf5c/content/dist/rhel8/8/x86_64/appstream/os
 templates/341aaa4f-6050-4599-99d8-01e81bf4bf5c/content/dist/rhel8/8/x86_64/ansible/os
```

Check the candlepin db:
```
make db-cli-connect
\connect candlepin
 select * from cp_environment_content;
```
You should have 3 pieces of content.

Register RHEL 8 system and assign it to the template using https://github.com/content-services/content-sources-backend/blob/main/docs/register_client.md  and verify you can see the 3 repos on the client (via yum repolist).

Now checkout this pr, restart the server:
```
make run
```
Now reimport the red hat repos which will trigger the deletion:

```
make repos-import
```

re-run the pulp and candlepin queries and you should see the ansible repo distribution no longer there in pulp, and only 2 content objects in the environment in candlepin

On the client, run 'subscription-manager refresh', and verify that you only see 2 repos configured on the client


---

## Discussion

### Comment by @jlsherrill on March 19, 2025 at 01:00 PM UTC

https://issues.redhat.com/browse/HMS-5303

### Comment by @xbhouse on March 24, 2025 at 07:07 PM UTC

i couldn't test this fully with a registered client because apparently RHEL8 VMs [aren't compatible](https://github.com/lima-vm/lima/issues/841) with Mac M-series? 😞 @rverdile or @dominikvagner when either of you get a chance, could you verify the full flow works with a registered client?

this mostly looks good to me up to the point i could test though, just a couple questions 😄 

### Comment by @jlsherrill on March 27, 2025 at 01:21 PM UTC

Adding the removal of the introspected repo.  The way it works, is that we just mark it as not public.  Since there should be no repo configs, it should get cleaned up automatically.  You can test this by:

on main run:
``` make repos-import ```
```
make db-cli-connect

update repositories set created_at = '2024-01-01'; 
select url from repositories where url ilike '%ansible%'; 
```

should see the ansible repos

Switch to this PR, run:
```
make repos-import
make process-repos
```
that should trigger the cleanup.  Since the repo has no repo-configs, is not public, and older than a week, it should be deleted:

```
select url from repositories where url ilike '%ansible%'; 
```


---

## Reviews

### Review by @xbhouse - Commented on March 24, 2025 at 05:47 PM UTC

### Review by @xbhouse - Commented on March 24, 2025 at 05:56 PM UTC

### Review by @jlsherrill - Commented on March 26, 2025 at 07:52 PM UTC

### Review by @rverdile - Commented on March 26, 2025 at 08:09 PM UTC

I tested with a rhel8 VM, works!

### Review by @xbhouse - Approved on March 28, 2025 at 04:28 PM UTC

lgtm! nice job!!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1038*
