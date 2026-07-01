---
type: pull_request
number: 817
title: "Fixes 4729: stop adding guard to RH repos"
state: merged
author: jlsherrill
created: 2024-09-16T13:20:04Z
updated: 2024-09-16T22:59:15Z
closed: 2024-09-16T22:59:15Z
merged: 2024-09-16T22:59:15Z
base_branch: main
head_branch: 4729
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/817
---

# Pull Request #817: Fixes 4729: stop adding guard to RH repos

**Author**: @jlsherrill
**Created**: September 16, 2024 at 01:20 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4729`

## Description

## Summary

* stops adding content guards to rh repos (there was a typo where it the orgID check was checking that of the template)
* updates the contentGuardHref on update, previously existing templates wouldn't have been fixed
* Unifies some logic around distributions into a PulpDistributionHelper

## Testing steps

1.  on main, set custom_repo_content_guards to true in your config.yaml
2. run  `make repos-import-rhel9`
3. run  `go run cmd/external-repos/main.go nightly-jobs`  let the snapshots happen
4. create a custom repo for snapshotting (can be https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/)
5. create a template, add the red hat repos and a custom repo
6. check the distribution/content guard assignment:
```
psql --host localhost  pulp -U pulp

select base_path, content_guard_id from core_distribution;
```

```
                                        base_path                                        |           content_guard_id           
-----------------------------------------------------------------------------------------+--------------------------------------
 templates/ab8ac6f2-cf49-4e51-826e-7d57c5c0b9d3/d6424a7d-51e7-4476-a20f-ceb0b428291d     | 0191faf4-c230-7fec-84db-d40c9705a361
 templates/ab8ac6f2-cf49-4e51-826e-7d57c5c0b9d3/content/dist/rhel9/9/x86_64/appstream/os | 0191ee4d-fd30-7910-913c-b78adcd28b4d
 templates/ab8ac6f2-cf49-4e51-826e-7d57c5c0b9d3/content/dist/rhel9/9/x86_64/baseos/os    | 0191ee4d-fd30-7910-913c-b78adcd28b4d

```

note that all 3 repos have a content guard associated, this is incorrect.

7. checkout this PR
8. issue an update the template in the api or ui (which is easier).  This will trigger an update content template task.  No need to change anything
9.  re-run the sql query:

```
psql --host localhost  pulp -U pulp

select base_path, content_guard_id from core_distribution;
```

```
                                        base_path                                        |           content_guard_id           
-----------------------------------------------------------------------------------------+--------------------------------------
 templates/ab8ac6f2-cf49-4e51-826e-7d57c5c0b9d3/d6424a7d-51e7-4476-a20f-ceb0b428291d     | 0191faf4-c230-7fec-84db-d40c9705a361
 templates/ab8ac6f2-cf49-4e51-826e-7d57c5c0b9d3/content/dist/rhel9/9/x86_64/appstream/os | 
 templates/ab8ac6f2-cf49-4e51-826e-7d57c5c0b9d3/content/dist/rhel9/9/x86_64/baseos/os    | 


```

The red hat repos should be 'fixed', and the custom repo should still have its guard associated

---

## Discussion

### Comment by @jlsherrill on September 16, 2024 at 01:30 PM UTC

https://issues.redhat.com/browse/HMS-4729

### Comment by @xbhouse on September 16, 2024 at 08:18 PM UTC

ah will re-approve once tests are finished and pass

### Comment by @jlsherrill on September 16, 2024 at 09:11 PM UTC

clowder deploy timeout 

/retest

---

## Reviews

### Review by @xbhouse - Approved on September 16, 2024 at 08:07 PM UTC

looks good and content guards aren't added to RH repos!

### Review by @xbhouse - Approved on September 16, 2024 at 08:29 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/817*
