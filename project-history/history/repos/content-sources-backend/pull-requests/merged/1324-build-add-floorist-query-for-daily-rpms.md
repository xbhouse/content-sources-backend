---
type: pull_request
number: 1324
title: "Build: add floorist query for daily rpms"
state: merged
author: jlsherrill
created: 2025-12-03T19:00:29Z
updated: 2025-12-04T16:12:54Z
closed: 2025-12-04T16:12:54Z
merged: 2025-12-04T16:12:54Z
base_branch: main
head_branch: rpm_query
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1324
---

# Pull Request #1324: Build: add floorist query for daily rpms

**Author**: @jlsherrill
**Created**: December 03, 2025 at 07:00 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `rpm_query`

## Description

## Summary

We've been requested to gather the list of rpms in this group of red hat repositories and upload via floorist

## Testing steps

 make repos-import-rhel9
make run  

# let the repos introspect

make db-cli-connect

```
           select rpms.name, rpms.version, rpms.release, rpms.arch, rpms.epoch, rc.name as repo_name, CURRENT_DATE as query_date
             from rpms
             inner join repositories_rpms on rpms.uuid = repositories_rpms.rpm_uuid
             inner join repositories on repositories_rpms.repository_uuid = repositories.uuid
             inner join repository_configurations rc on repositories.uuid = rc.repository_uuid
             where rc.org_id = '-1'
             and rc.name in (
               'Red Hat Enterprise Linux 10 for x86_64 - AppStream (RPMs)',
               'Red Hat Enterprise Linux 9 for x86_64 - AppStream (RPMs)',
               'Red Hat Enterprise Linux 8 for x86_64 - AppStream (RPMs)',
               'Red Hat Enterprise Linux 10 for x86_64 - BaseOS (RPMs)',
               'Red Hat Enterprise Linux 9 for x86_64 - BaseOS (RPMs)',
               'Red Hat Enterprise Linux 8 for x86_64 - BaseOS (RPMs)'
             );
```

and see that you get some results



---

## Reviews

### Review by @rverdile - Approved on December 03, 2025 at 08:46 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1324*
