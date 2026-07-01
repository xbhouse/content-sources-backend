---
type: pull_request
number: 261
title: "Fixes 1672: support multiple gpg keys"
state: merged
author: rverdile
created: 2023-04-28T14:01:51Z
updated: 2023-05-01T14:30:15Z
closed: 2023-05-01T14:16:51Z
merged: 2023-05-01T14:16:51Z
base_branch: main
head_branch: two-gpg
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/261
---

# Pull Request #261: Fixes 1672: support multiple gpg keys

**Author**: @rverdile
**Created**: April 28, 2023 at 02:01 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `two-gpg`

## Description

## Summary
Previously, if there were multiple keys in a gpg key file, only the first key was read. Now multiple keys can be read.

## Testing steps
1. In UI, create this repo: https://dl.google.com/linux/chrome/rpm/stable/x86_64
2. With this gpg key: https://dl.google.com/linux/linux_signing_key.pub
3. With metadata verification set to true
4. Should not return a validation error. Previously, would return the error `Error validating signature: openpgp: signature made by unknown entity. Is this the correct GPG Key?`

---

## Discussion

### Comment by @jlsherrill on April 28, 2023 at 02:30 PM UTC

https://issues.redhat.com/browse/HMS-1672

### Comment by @mayurilahane on April 28, 2023 at 09:40 PM UTC

Just deployed the pod with these changes in it.
And as mentioned:

_"Should not return a validation error. Previously, would return the error Error validating signature: openpgp: signature made by an unknown entity. Is this the correct GPG Key?"_

It is **not** giving any error for this url and gpg key with metadata verification true


### Comment by @mayurilahane on April 28, 2023 at 09:41 PM UTC

![Screenshot from 2023-04-28 17-37-05](https://user-images.githubusercontent.com/21276218/235258999-106ffadf-f63c-45bf-8ea1-361b63e07d0e.png)


---

## Reviews

### Review by @jlsherrill - Approved on April 28, 2023 at 04:13 PM UTC

### Review by @rverdile - Commented on April 28, 2023 at 08:59 PM UTC

### Review by @rverdile - Commented on April 28, 2023 at 09:03 PM UTC

### Review by @mayurilahane - Commented on April 28, 2023 at 10:30 PM UTC

### Review by @mayurilahane - Commented on May 01, 2023 at 02:16 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/261*
