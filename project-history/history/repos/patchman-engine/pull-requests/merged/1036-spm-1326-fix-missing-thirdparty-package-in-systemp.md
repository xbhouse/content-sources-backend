---
type: pull_request
number: 1036
title: "SPM-1326: fix missing thirdparty package in system/packages"
state: merged
author: psegedy
created: 2022-07-22T15:28:49Z
updated: 2022-07-22T15:49:09Z
closed: 2022-07-22T15:49:09Z
merged: 2022-07-22T15:49:08Z
base_branch: master
head_branch: fix_thirdparty
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1036
---

# Pull Request #1036: SPM-1326: fix missing thirdparty package in system/packages

**Author**: @psegedy
**Created**: July 22, 2022 at 03:28 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix_thirdparty`

## Description

Basically the same thing as for https://github.com/RedHatInsights/patchman-engine/pull/1026, there is `""` vs null issue in DB, so we might want to use left join

with `JOIN`
```
-d "{\"query\": \"select pn.name from system_package sp JOIN package p ON sp.package_id = p.id JOIN strings sum ON p.summary_hash = sum.id JOIN package_name pn ON pn.id = sp.name_id where sp.system_id = '1422060' and sp.name_id = '660016426'\"}" -s | jq
{
  "result": [
    [
      "name"
    ]
  ],
  "error": ""
}
```

with LEFT JOIN
```
-d "{\"query\": \"select pn.name from system_package sp JOIN package p ON sp.package_id = p.id LEFT JOIN strings sum ON p.summary_hash = sum.id JOIN package_name pn ON pn.id = sp.name_id where sp.system_id = '1422060' and sp.name_id = '660016426'\"}" -s | jq
{
  "result": [
    [
      "name"
    ],
    [
      "thirdparty"
    ]
  ],
  "error": ""
}
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

### Comment by @jira-linking on July 22, 2022 at 03:28 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1326


---

## Reviews

### Review by @michalslomczynski - Approved on July 22, 2022 at 03:30 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1036*
