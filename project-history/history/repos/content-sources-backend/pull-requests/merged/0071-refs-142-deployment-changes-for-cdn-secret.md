---
type: pull_request
number: 71
title: "Refs 142: deployment changes for cdn secret"
state: merged
author: jlsherrill
created: 2022-08-03T21:10:38Z
updated: 2022-08-10T17:00:15Z
closed: 2022-08-08T13:26:52Z
merged: 2022-08-08T13:26:52Z
base_branch: main
head_branch: 142_2
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/71
---

# Pull Request #71: Refs 142: deployment changes for cdn secret

**Author**: @jlsherrill
**Created**: August 03, 2022 at 09:10 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `142_2`

## Description

*No description provided*

---

## Discussion

### Comment by @jlsherrill on August 03, 2022 at 09:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-142

### Comment by @jlsherrill on August 03, 2022 at 09:30 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @avisiedo on August 05, 2022 at 06:42 PM UTC

- Checked that the secret exists.

```sh
oc get secrets/content-sources-certs
NAME                    TYPE     DATA   AGE
content-sources-certs   Opaque   1      37m
```

- Checked that the env var contain the secret value in a rsh opened into the container:

```sh
oc rsh pod/POD_NAME
env | grep RH_CDN_CERT_PAIR
RH_CDN_CERT_PAIR=THE_EXPECTED_VALUE
```


---

## Reviews

### Review by @avisiedo - Approved on August 05, 2022 at 06:46 PM UTC

I have checked the secret and the content of the variable into the pod deployed exists and match.

LGTM

### Review by @swadeley - Approved on August 08, 2022 at 01:26 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/71*
