---
type: pull_request
number: 1296
title: "POC: use system_package2 in system_packages"
state: merged
author: psegedy
created: 2023-08-21T14:34:54Z
updated: 2023-08-21T15:06:02Z
closed: 2023-08-21T15:06:02Z
merged: 2023-08-21T15:06:02Z
base_branch: poc
head_branch: poc_system_packages
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1296
---

# Pull Request #1296: POC: use system_package2 in system_packages

**Author**: @psegedy
**Created**: August 21, 2023 at 02:34 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `poc` ← **Head**: `poc_system_packages`

## Description

original query - `(cost=243.58..2128.00 rows=684 width=231) (actual time=27.508..32.832 rows=888 loops=1)`
```sql
SELECT
    pn.name as name,
    p.evra as evra,
    sum.value as summary,
    descr.value as description,
    (update_status(spkg.update_data) = 'Installable') as updatable,
    update_status(spkg.update_data) as update_status,
    spkg.update_data as updates,
    count(*) over () as total
FROM
    system_platform sp
    JOIN inventory.hosts ih ON ih.id = sp.inventory_id
    JOIN system_package spkg on spkg.system_id = sp.id
    AND spkg.rh_account_id = 37
    JOIN package p on p.id = spkg.package_id
    JOIN package_name pn on pn.id = spkg.name_id
    LEFT JOIN strings AS descr ON p.description_hash = descr.id
    LEFT JOIN strings AS sum ON p.summary_hash = sum.id
WHERE
    sp.rh_account_id = 37
    AND sp.inventory_id = '05b19ade-549f-45c6-ab35-942d81a4bd22' :: uuid;
```

with system_package2 - `(cost=592.08..1945.82 rows=684 width=213) (actual time=23.163..23.470 rows=888 loops=1)`
```sql
SELECT
    pn.name as name,
    p.evra as evra,
    sum.value as summary,
    descr.value as description,
    (spkg.installable_id is not null) as updatable,
    CASE
        WHEN spkg.installable_id is not null THEN 'Installable'
        WHEN spkg.applicable_id is not null THEN 'Applicable'
        ELSE 'None'
    END as update_status,
    null as updates,
    pi.evra,
    pa.evra,
    count(*) over () as total
FROM
    system_platform sp
    JOIN inventory.hosts ih ON ih.id = sp.inventory_id
    JOIN system_package2 spkg on spkg.system_id = sp.id
    AND spkg.rh_account_id = 37
    JOIN package p on p.id = spkg.package_id
    JOIN package_name pn on pn.id = spkg.name_id
    LEFT JOIN package pi ON pi.id = spkg.installable_id
    LEFT JOIN package pa ON pa.id = spkg.applicable_id
    LEFT JOIN strings AS descr ON p.description_hash = descr.id
    LEFT JOIN strings AS sum ON p.summary_hash = sum.id
WHERE
    sp.rh_account_id = 37
    AND sp.inventory_id = '05b19ade-549f-45c6-ab35-942d81a4bd22' :: uuid;
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

### Comment by @jira-linking on August 21, 2023 at 02:34 PM UTC

Commits missing Jira IDs:
440567e15246e1491a46aa766f4ef04790bf2f9f


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1296*
