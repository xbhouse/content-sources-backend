---
type: pull_request
number: 643
title: "Fix message processing, add more lenient datetime parser"
state: merged
author: semtexzv
created: 2021-04-28T10:31:44Z
updated: 2021-04-28T12:50:46Z
closed: 2021-04-28T12:50:46Z
merged: 2021-04-28T12:50:46Z
base_branch: master
head_branch: datetime-parser
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/643
---

# Pull Request #643: Fix message processing, add more lenient datetime parser

**Author**: @semtexzv
**Created**: April 28, 2021 at 10:31 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `datetime-parser`

## Description

We should accept datetime formats with optional fields. Go doesn't like what the platform is sending us, so let's add multiple formats for parsing. 

Also added logging and reporting error that occur when parsing messages.
## Secure Coding Practices Checklist GitHub Link
- https://github.com/RedHatInsights/secure-coding-checklist

## Secure Coding Checklist
- [x] Input Validation
- [ ] Output Encoding
- [ ] Authentication and Password Management
- [ ] Session Management
- [ ] Access Control
- [ ] Cryptographic Practices
- [x] Error Handling and Logging
- [ ] Data Protection
- [ ] Communication Security
- [ ] System Configuration
- [ ] Database Security
- [ ] File Management
- [ ] Memory Management
- [ ] General Coding Practices


---

## Reviews

### Review by @MichaelMraka - Approved on April 28, 2021 at 11:20 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/643*
