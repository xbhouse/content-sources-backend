---
type: pull_request
number: 56
title: "Fixes 134: Set names for init-containers and typo fix"
state: merged
author: Victoremepunto
created: 2022-07-21T09:33:36Z
updated: 2022-07-21T13:22:59Z
closed: 2022-07-21T13:13:32Z
merged: 2022-07-21T13:13:32Z
base_branch: main
head_branch: set-names-for-init-containers
labels: []
url: https://github.com/content-services/content-sources-backend/pull/56
---

# Pull Request #56: Fixes 134: Set names for init-containers and typo fix

**Author**: @Victoremepunto
**Created**: July 21, 2022 at 09:33 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `set-names-for-init-containers`

## Description

The Clowder template require names for initContainers if more than one is specified

```
  STDERR:
The ClowdApp "content-sources" is invalid: 

* spec.Deployments[0].PodSpec.InitContainers[0]: Forbidden: multiple initcontainers must have a name

* spec.Deployments[0].PodSpec.InitContainers[1]: Forbidden: multiple initcontainers must have a name
```


see the Clowder API [spec](https://redhatinsights.github.io/clowder/clowder/dev/api_reference.html#k8s-api-github-com-redhatinsights-clowder-apis-cloud-redhat-com-v1alpha1-initcontainer)



---

## Discussion

### Comment by @jlsherrill on July 21, 2022 at 01:22 PM UTC

Thanks @Victoremepunto !!

---

## Reviews

### Review by @jlsherrill - Approved on July 21, 2022 at 01:10 PM UTC

### Review by @swadeley - Approved on July 21, 2022 at 01:13 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/56*
