---
type: pull_request
number: 529
title: "[WIP] Allow pagination entities nullable in spec"
state: closed
author: digitronik
created: 2021-02-10T11:25:48Z
updated: 2021-04-14T05:19:05Z
closed: 2021-04-14T05:19:05Z
base_branch: master
head_branch: nullable
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/529
---

# Pull Request #529: [WIP] Allow pagination entities nullable in spec

**Author**: @digitronik
**Created**: February 10, 2021 at 11:25 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `nullable`

## Description

## Secure Coding Practices Checklist GitHub Link
- https://github.com/RedHatInsights/secure-coding-checklist

I was trying to generate a client with new OpenAPI 5. facing errors like

```
ApiTypeError: Invalid type for variable 'previous'. Required value type is str and passed type was NoneType at ['received_data']['links']['previous']
```

## Secure Coding Checklist
- [ ] Input Validation
- [ ] Output Encoding
- [ ] Authentication and Password Management
- [ ] Session Management
- [ ] Access Control
- [ ] Cryptographic Practices
- [ ] Error Handling and Logging
- [ ] Data Protection
- [ ] Communication Security
- [ ] System Configuration
- [ ] Database Security
- [ ] File Management
- [ ] Memory Management
- [ ] General Coding Practices


---

## Discussion

### Comment by @josef-hak on April 13, 2021 at 11:22 AM UTC

@digitronik is it still actual or can we close it?

### Comment by @digitronik on April 14, 2021 at 05:19 AM UTC

@Josca my bad... I want to learn it but time :)
closing this as @MichaelMraka  works on openapi5 client. I think it will cover this.


---

## Reviews

### Review by @josef-hak - Commented on February 10, 2021 at 12:25 PM UTC

### Review by @digitronik - Commented on February 10, 2021 at 12:41 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/529*
