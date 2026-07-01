---
type: pull_request
number: 689
title: "ubi8 is now available on docker.io"
state: closed
author: MichaelMraka
created: 2021-06-01T08:57:12Z
updated: 2021-06-02T19:19:58Z
closed: 2021-06-02T19:19:58Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/689
---

# Pull Request #689: ubi8 is now available on docker.io

**Author**: @MichaelMraka
**Created**: June 01, 2021 at 08:57 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `pr1`

## Description

 let's use it in development and simplify our setup

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

### Comment by @MichaelMraka on June 02, 2021 at 07:19 PM UTC

Unfortunately it looks like you can get it from docker.io but then need to register it via subscription-manager (if not on already registered host).

Closing this PR for now.

---

## Reviews

### Review by @josef-hak - Changes Requested on June 02, 2021 at 02:09 PM UTC

Great if it will be possible to use rhel image from dockerhub, but it seems there are not modules supported:
~~~
This system is not registered to Red Hat Subscription Management. You can use subscription-manager to register.

Red Hat Universal Base Image 8 (RPMs) - BaseOS  2.1 MB/s | 786 kB     00:00    
Red Hat Universal Base Image 8 (RPMs) - AppStre  25 MB/s | 7.4 MB     00:00    
Red Hat Universal Base Image 8 (RPMs) - CodeRea 129 kB/s |  15 kB     00:00    
Error: Problems in request:
missing groups or modules: postgresql:12
The command '/bin/sh -c dnf module -y enable postgresql:12 &&     dnf install -y go-toolset postgresql git-core diffutils' returned a non-zero code: 1
Service 'platform' failed to build : Build failed
~~~

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/689*
