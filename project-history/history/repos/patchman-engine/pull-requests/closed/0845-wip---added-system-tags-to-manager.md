---
type: pull_request
number: 845
title: "wip - added system tags to manager"
state: closed
author: mkholjuraev
created: 2021-10-13T14:19:54Z
updated: 2021-11-12T08:56:54Z
closed: 2021-11-12T08:56:54Z
base_branch: master
head_branch: add-tags-into-api
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/845
---

# Pull Request #845: wip - added system tags to manager

**Author**: @mkholjuraev
**Created**: October 13, 2021 at 02:19 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `add-tags-into-api`

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

### Comment by @mkholjuraev on October 13, 2021 at 02:51 PM UTC

This is a draft PR to add tags info into /systems endpoint as additional info. Following is the format that I want to send as a response. I have created separate SystemTags and used them to select tags from the database. I am binding the subquery to fetch the tags from the database into existing  querySystems.
```
tags: [
            {
              "key": "string",
              "namespace": "string",
              "value": "string"
            }
        ],
```

 In the openApi docs I can see that the format is correct 
 
 
![image](https://user-images.githubusercontent.com/59481011/137152851-1c77ffb2-4b43-4a18-87db-bcc16e6efdb9.png)


However, tests and openApi request call to /systems are failing with the following error.

![image](https://user-images.githubusercontent.com/59481011/137153689-6f84fc99-ef08-4b1b-b1fc-24ee943c80fb.png)

I believe that there is a problem with db connection. It would be a great help if someone can have a look and give a hint so that I I can solve the issue on my own for learning :)

### Comment by @josef-hak on November 12, 2021 at 08:56 AM UTC

Resolved in #854 

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/845*
