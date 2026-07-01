---
type: pull_request
number: 1154
title: "Build: append RHEL feature manually when importing repos"
state: merged
author: xbhouse
created: 2025-07-24T17:25:37Z
updated: 2025-07-25T13:31:24Z
closed: 2025-07-24T18:28:13Z
merged: 2025-07-24T18:28:13Z
base_branch: main
head_branch: layered-repos-fix
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1154
---

# Pull Request #1154: Build: append RHEL feature manually when importing repos

**Author**: @xbhouse
**Created**: July 24, 2025 at 05:25 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `layered-repos-fix`

## Description

## Summary

The prod feature service doesn't include RHEL-OS-x86_64 yet as a feature, so we can't include that in the request.

## Testing steps

1. Set the feature service server in the clients section of the config to the prod server and grab the certs from vault and set them here too: 

```
  feature_service:
    server: https://feature.api.redhat.com/features/v1
    client_cert:
    client_key:
    ca_cert:
    client_cert_path: /home/bhouse/certs/cp_prod_client.cert
    client_key_path: /home/bhouse/certs/cp_prod_client.key
    ca_cert_path:
```
2. Also set options.feature_filter in the config to `["RHEL-HA-x86_64", "OPENSHIFT-OCP-x86_64"]`
3. On a fresh backend, run make repos-import
4. You should see the layered repos (openshift and HA) in the response when listing repos either via API or UI (using your prod account) 
    



---

## Reviews

### Review by @jlsherrill - Approved on July 24, 2025 at 06:14 PM UTC

ACK tested pointing at production, hit the error you saw, updated to this and removed that feature and was able to see HA and other repos returned via the api 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1154*
