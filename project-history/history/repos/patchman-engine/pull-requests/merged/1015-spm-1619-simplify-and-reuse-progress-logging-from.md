---
type: pull_request
number: 1015
title: "SPM-1619: simplify and reuse progress logging from package cache"
state: merged
author: MichaelMraka
created: 2022-07-08T15:17:46Z
updated: 2022-07-11T12:27:03Z
closed: 2022-07-11T12:27:03Z
merged: 2022-07-11T12:27:03Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1015
---

# Pull Request #1015: SPM-1619: simplify and reuse progress logging from package cache

**Author**: @MichaelMraka
**Created**: July 08, 2022 at 03:17 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr1`

## Description

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

### Comment by @jira-linking on July 08, 2022 at 03:17 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1619


### Comment by @psegedy on July 11, 2022 at 12:26 PM UTC

```
time="2022-07-11T12:26:01Z" level=info msg="Advisory detail cache preload" gorutineID=10 progress %=19
time="2022-07-11T12:26:01Z" level=info msg="Advisory detail cache preload" gorutineID=11 progress %=19
time="2022-07-11T12:26:03Z" level=info msg="Advisory detail cache preload" gorutineID=10 progress %=38
time="2022-07-11T12:26:03Z" level=info msg="Advisory detail cache preload" gorutineID=11 progress %=39
time="2022-07-11T12:26:05Z" level=info msg="Advisory detail cache preload" gorutineID=10 progress %=56
time="2022-07-11T12:26:05Z" level=info msg="Advisory detail cache preload" gorutineID=11 progress %=57
```
LGTM, unit test failure is unrelated, I'll fix that

---

## Reviews

### Review by @psegedy - Commented on July 11, 2022 at 09:17 AM UTC

I tried that in ephemeral and it logs the progress twice with every ticker
```
time="2022-07-11T09:11:29Z" level=info msg="Advisory detail cache preload" progress %=17
time="2022-07-11T09:11:29Z" level=info msg="Advisory detail cache preload" progress %=18
time="2022-07-11T09:11:31Z" level=info msg="Advisory detail cache preload" progress %=34
time="2022-07-11T09:11:31Z" level=info msg="Advisory detail cache preload" progress %=35
time="2022-07-11T09:11:33Z" level=info msg="Advisory detail cache preload" progress %=52
time="2022-07-11T09:11:33Z" level=info msg="Advisory detail cache preload" progress %=52
time="2022-07-11T09:11:35Z" level=info msg="Advisory detail cache preload" progress %=70
time="2022-07-11T09:11:35Z" level=info msg="Advisory detail cache preload" progress %=70
time="2022-07-11T09:11:37Z" level=info msg="Advisory detail cache preload" progress %=87
time="2022-07-11T09:11:37Z" level=info msg="Advisory detail cache preload" progress %=87
```
IMHO the cause is that manager is running in multiple processes or threads (2 in ephemeral)

### Review by @psegedy - Approved on July 11, 2022 at 12:26 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1015*
