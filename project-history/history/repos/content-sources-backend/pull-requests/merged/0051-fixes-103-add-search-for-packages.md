---
type: pull_request
number: 51
title: "Fixes 103: add search for packages"
state: merged
author: avisiedo
created: 2022-07-13T09:09:27Z
updated: 2022-07-25T12:37:13Z
closed: 2022-07-25T12:37:13Z
merged: 2022-07-25T12:37:13Z
base_branch: main
head_branch: hmscontent-103
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/51
---

# Pull Request #51: Fixes 103: add search for packages

**Author**: @avisiedo
**Created**: July 13, 2022 at 09:09 AM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `hmscontent-103`

## Description

- Add `Public` column with `false` as default value.
- Add handler for the `/rpms/names` route.
- Implement the logic to list the suggested packages given a query and a list of urls.
- Add unit tests.


---

## Discussion

### Comment by @jlsherrill on July 13, 2022 at 09:15 AM UTC

https://issues.redhat.com/browse/HMSCONTENT-103

### Comment by @jlsherrill on July 18, 2022 at 06:42 PM UTC

Built and pushed image for 1c975f8ab8904f361a178f0373d988a38ef432ee

### Comment by @avisiedo on July 19, 2022 at 11:40 AM UTC

Rebase once https://github.com/content-services/content-sources-backend/pull/38 has been merged as it uses fields created at there

### Comment by @jlsherrill on July 19, 2022 at 08:20 PM UTC

https://issues.redhat.com/browse/14709253

### Comment by @jlsherrill on July 20, 2022 at 08:33 PM UTC

This is looking really good!  Mostly minor comments.  I commented out the repository_configurations join and got results!  Awesome work!

```
curl -X POST  -H 'Content-Type: application/json'  http://localhost:8000/api/content_sources/v1.0/rpms/names   -d '{"urls":["https://packages.cloud.google.com/yum/repos/cloud-sdk-el8-x86_64"], "search":"goog"}'    -H "`./scripts/header.sh  1 1 `"
[{"package_name":"google-cloud-cli","summary":"Google Cloud CLI"},{"package_name":"google-cloud-cli-anthos-auth","summary":"Google Cloud CLI Anthos Auth"},{"package_name":"google-cloud-cli-app-engine-go","summary":"Google Cloud CLI App Engine Go"},{"package_name":"google-cloud-cli-app-engine-grpc","summary":"Google Cloud CLI App Engine gRPC"},{"package_name":"google-cloud-cli-app-engine-java","summary":"Google Cloud CLI App Engine Java"},{"package_name":"google-cloud-cli-app-engine-python","summary":"Google Cloud CLI App Engine Python"},{"package_name":"google-cloud-cli-app-engine-python-extras","summary":"Google Cloud CLI App Engine Python extra libraries"},{"package_name":"google-cloud-cli-bigtable-emulator","summary":"Google Cloud CLI Bigtable Emulator"},{"package_name":"google-cloud-cli-cbt","summary":"Google Cloud CLI Cloud Bigtable"},{"package_name":"google-cloud-cli-cloud-build-local","summary":"Google Cloud CLI Cloud Build Local Builder"},{"package_name":"google-cloud-cli-cloud-run-proxy","summary":"Google Cloud CLI Cloud Run Proxy"},{"package_name":"google-cloud-cli-config-connector","summary":"Google Cloud CLI Config Connector"},{"package_name":"google-cloud-cli-datalab","summary":"Google Cloud CLI Datalab"},{"package_name":"google-cloud-cli-datastore-emulator","summary":"Google Cloud CLI Datastore Emulator"},{"package_name":"google-cloud-cli-firestore-emulator","summary":"Google Cloud CLI Firestore Emulator"},{"package_name":"google-cloud-cli-gke-gcloud-auth-plugin","summary":"Google Cloud CLI GKE Auth Plugin"},{"package_name":"google-cloud-cli-harbourbridge","summary":"Google Cloud CLI Harbourbridge"},{"package_name":"google-cloud-cli-kpt","summary":"Google Cloud CLI kpt"},{"package_name":"google-cloud-cli-kubectl-oidc","summary":"Google Cloud CLI Kubectl OIDC"},{"package_name":"google-cloud-cli-local-extract","summary":"Google Cloud CLI Local Extract"}]

```

### Comment by @swadeley on July 25, 2022 at 08:00 AM UTC

@avisiedo Hello, rebase please.

---

## Reviews

### Review by @jlsherrill - Commented on July 20, 2022 at 08:22 PM UTC

### Review by @jlsherrill - Commented on July 20, 2022 at 08:23 PM UTC

### Review by @jlsherrill - Commented on July 20, 2022 at 08:27 PM UTC

### Review by @jlsherrill - Commented on July 20, 2022 at 08:31 PM UTC

### Review by @jlsherrill - Commented on July 20, 2022 at 08:34 PM UTC

### Review by @rverdile - Commented on July 22, 2022 at 01:34 PM UTC

### Review by @rverdile - Commented on July 22, 2022 at 01:38 PM UTC

### Review by @rverdile - Commented on July 22, 2022 at 01:43 PM UTC

### Review by @rverdile - Commented on July 22, 2022 at 01:52 PM UTC

It looks like repository_rpm is missing handler tests, even on main branch. Might be good to add that here.

### Review by @avisiedo - Commented on July 22, 2022 at 07:00 PM UTC

### Review by @avisiedo - Commented on July 22, 2022 at 07:29 PM UTC

### Review by @avisiedo - Commented on July 22, 2022 at 07:37 PM UTC

### Review by @avisiedo - Commented on July 25, 2022 at 08:28 AM UTC

### Review by @avisiedo - Commented on July 25, 2022 at 08:33 AM UTC

### Review by @swadeley - Approved on July 25, 2022 at 12:36 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/51*
