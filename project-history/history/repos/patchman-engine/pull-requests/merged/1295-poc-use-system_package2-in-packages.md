---
type: pull_request
number: 1295
title: "POC: use system_package2 in /packages"
state: merged
author: psegedy
created: 2023-08-21T13:38:53Z
updated: 2023-08-21T13:54:13Z
closed: 2023-08-21T13:54:13Z
merged: 2023-08-21T13:54:13Z
base_branch: poc
head_branch: poc_packages
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1295
---

# Pull Request #1295: POC: use system_package2 in /packages

**Author**: @psegedy
**Created**: August 21, 2023 at 01:38 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `poc` ← **Head**: `poc_packages`

## Description

original query `(cost=21664.34..21672.64 rows=200 width=66) (actual time=177.243..181.107 rows=300 loops=1)`
```sql
SELECT
    count(*) over () as total,
    pn.name as name,
    pn.summary as summary,
    res.systems_installed as systems_installed,
    res.systems_installable as systems_updatable,
    res.systems_installable as systems_installable,
    res.systems_applicable as systems_applicable
FROM
    package_name pn
    JOIN (
        SELECT
            spkg.name_id as name_id,
            count(*) as systems_installed,
            count(*) filter (
                where
                    update_status(spkg.update_data) = 'Installable'
            ) as systems_installable,
            count(*) filter (
                where
                    update_status(spkg.update_data) != 'None'
            ) as systems_applicable
        FROM
            system_package spkg
        WHERE
            spkg.rh_account_id = 37
            AND spkg.system_id IN (
                SELECT
                    sp.id
                FROM
                    system_platform sp
                    JOIN inventory.hosts ih ON ih.id = sp.inventory_id
                WHERE
                    sp.rh_account_id = 37
                    AND (
                        sp.stale = false
                        AND sp.packages_installed > 0
                    )
            )
        GROUP BY
            "spkg"."name_id"
    ) res ON res.name_id = pn.id;
```

query with system_package2 `(cost=5265.89..5274.19 rows=200 width=66) (actual time=146.491..151.353 rows=300 loops=1)`
```sql
SELECT
    count(*) over () as total,
    pn.name as name,
    pn.summary as summary,
    res.systems_installed as systems_installed,
    res.systems_installable as systems_updatable,
    res.systems_installable as systems_installable,
    res.systems_applicable as systems_applicable
FROM
    package_name pn
    JOIN (
        SELECT
            spkg.name_id as name_id,
            count(*) as systems_installed,
            count(*) filter (
                where
                    spkg.installable_id is not null
            ) as systems_installable,
            count(*) filter (
                where
                    spkg.installable_id is not null
                    or spkg.applicable_id is not null
            ) as systems_applicable
        FROM
            system_package2 spkg
        WHERE
            spkg.rh_account_id = 37
            AND spkg.system_id IN (
                SELECT
                    sp.id
                FROM
                    system_platform sp
                    JOIN inventory.hosts ih ON ih.id = sp.inventory_id
                WHERE
                    sp.rh_account_id = 37
                    AND (
                        sp.stale = false
                        AND sp.packages_installed > 0
                    )
            )
        GROUP BY
            "spkg"."name_id"
    ) res ON res.name_id = pn.id;
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

### Comment by @jira-linking on August 21, 2023 at 01:38 PM UTC

Commits missing Jira IDs:
51d14a4e3c18d9c4055e32b430b07ce5a51f4909


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1295*
