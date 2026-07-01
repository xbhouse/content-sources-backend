---
type: pull_request
number: 1287
title: "POC: package systems"
state: merged
author: psegedy
created: 2023-08-10T17:36:46Z
updated: 2023-08-21T13:05:47Z
closed: 2023-08-21T13:05:47Z
merged: 2023-08-21T13:05:47Z
base_branch: poc
head_branch: poc_package_systems
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1287
---

# Pull Request #1287: POC: package systems

**Author**: @psegedy
**Created**: August 10, 2023 at 05:36 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `poc` ← **Head**: `poc_package_systems`

## Description

TODO:
- [x] add system_package2 test_data

original query - `(cost=1.40..26.06 rows=1 width=292) (actual time=2.213..3.124 rows=145 loops=1)`
```sql
SELECT
    count(*) over () as total,
    ih.tags as tags_str,
    ih.groups as groups_str,
    sp.inventory_id as id,
    sp.display_name as display_name,
    p.evra as installed_evra,
    spkg.latest_evra as available_evra,
    (update_status(spkg.update_data) = 'Installable') as updatable,
    null as Tags,
    bl.name as baseline_name,
    sp.baseline_uptodate as baseline_uptodate,
    bl.id as baseline_id,
    ih.system_profile -> 'operating_system' ->> 'name' || ' ' || coalesce(
        ih.system_profile -> 'operating_system' ->> 'major' || '.' || (ih.system_profile -> 'operating_system' ->> 'minor'),
        ''
    ) as os,
    ih.system_profile -> 'rhsm' ->> 'version' as rhsm,
    update_status(spkg.update_data) as update_status,
    Groups as Groups
FROM
    system_platform sp
    JOIN inventory.hosts ih ON ih.id = sp.inventory_id
    JOIN system_package spkg on spkg.system_id = sp.id
    AND spkg.rh_account_id = 50
    JOIN package p on p.id = spkg.package_id
    JOIN package_name pn on pn.id = spkg.name_id
    LEFT JOIN baseline bl ON sp.baseline_id = bl.id
    AND sp.rh_account_id = bl.rh_account_id
WHERE
    sp.rh_account_id = 50
    AND sp.stale = false
    AND pn.name = 'kernel'
    AND spkg.package_id in (2200, 2500);
```

new query - `(cost=1.97..38.07 rows=1 width=360) (actual time=4.045..4.123 rows=145 loops=1)`
```sql
SELECT
    count(*) over () as total,
    ih.tags as tags_str,
    ih.groups as groups_str,
    sp.inventory_id as id,
    sp.display_name as display_name,
    p.evra as installed_evra,
    null as AvailableEVRA,
    (spkg.installable_id IS NOT NULL) as updatable,
    null as Tags,
    bl.name as baseline_name,
    sp.baseline_uptodate as baseline_uptodate,
    spkg.installable_id as installable_id,
    spkg.applicable_id as applicable_id,
    pi.evra as installable_evra,
    pa.evra as applicable_evra,
    bl.id as baseline_id,
    ih.system_profile -> 'operating_system' ->> 'name' || ' ' || coalesce(
        ih.system_profile -> 'operating_system' ->> 'major' || '.' || (
            ih.system_profile -> 'operating_system' ->> 'minor'
        ),
        ''
    ) as os,
    ih.system_profile -> 'rhsm' ->> 'version' as rhsm,
    CASE
        WHEN spkg.installable_id is not null THEN 'Installable'
        WHEN spkg.applicable_id is not null THEN 'Applicable'
        ELSE 'None'
    END as update_status,
    Groups as Groups
FROM
    system_platform sp
    JOIN inventory.hosts ih ON ih.id = sp.inventory_id
    JOIN system_package2 spkg on spkg.system_id = sp.id
    AND spkg.rh_account_id = 50
    JOIN package p on p.id = spkg.package_id
    JOIN package_name pn on pn.id = spkg.name_id
    LEFT JOIN package pi ON pi.id = spkg.installable_id
    LEFT JOIN package pa ON pa.id = spkg.applicable_id
    LEFT JOIN baseline bl ON sp.baseline_id = bl.id
    AND sp.rh_account_id = bl.rh_account_id
WHERE
    sp.rh_account_id = 50
    AND sp.stale = false
    AND pn.name = 'kernel'
    AND spkg.package_id in (2200, 2500);
```
## Secure Coding Practices Checklist GitHub Link
- https://github.com/RedHatInsights/secure-coding-checklist

## Secure Coding Checklist
- [x] Input Validation
- [x] Output Encoding
- [x] Authentication and Password Management
- [x] Session Management
- [x] Access Control
- [x] Cryptographic Practices
- [x] Error Handling and Logging
- [x] Data Protection
- [x] Communication Security
- [x] System Configuration
- [x] Database Security
- [x] File Management
- [x] Memory Management
- [x] General Coding Practices


---

## Discussion

### Comment by @jira-linking on August 10, 2023 at 05:42 PM UTC

Commits missing Jira IDs:
f00cd514a72ddf2842106b23f616ff7d76bfd274
0eb06ec283810843bf55df7c3303192c7570a5c8


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1287*
