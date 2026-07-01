---
type: pull_request
number: 22
title: "CONTENT-67: db migrate in containers"
state: merged
author: jlsherrill
created: 2022-05-31T17:50:56Z
updated: 2022-06-02T15:06:43Z
closed: 2022-06-02T15:06:40Z
merged: 2022-06-02T15:06:40Z
base_branch: main
head_branch: cont_migrate
labels: []
url: https://github.com/content-services/content-sources-backend/pull/22
---

# Pull Request #22: CONTENT-67: db migrate in containers

**Author**: @jlsherrill
**Created**: May 31, 2022 at 05:50 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `cont_migrate`

## Description

This:
* builds the migration files and the migration command into the container 
* configures the deployment to use an init container to migrate the database
* Adjusts the migration command to not error if no changes are detected

---

## Discussion

### Comment by @rverdile on June 01, 2022 at 07:16 PM UTC

Would I test this by trying to deploy it?

### Comment by @jlsherrill on June 01, 2022 at 07:56 PM UTC

You can if you'd like:

```
 bonfire deploy -n $NS --set-image-tag  jlsherri/content-sources=fd5a03268be-test content-sources
```

with my local bonfire config modified to be:

```
# Bonfire deployment configuration

# Defines where to fetch the file that defines application configs
appsFile:
  host: gitlab
  repo: insights-platform/cicd-common
  path: bonfire_configs/ephemeral_apps.yaml

# (optional) define any apps locally. An app defined here with <name> will override config for app
# <name> in above fetched config.
apps:
- name: content-sources
  components:
    - name: content-sources
      host: github
      repo: jlsherrill/content-sources-backend
      path: deployments/deployment.yaml
      parameters:
        - name: ENV_NAME
          required: true
```

Note that repo was changed to point to mine where 'main' has this change.

---

## Reviews

### Review by @rverdile - Commented on June 01, 2022 at 04:09 PM UTC

### Review by @jlsherrill - Commented on June 01, 2022 at 06:01 PM UTC

### Review by @jlsherrill - Commented on June 01, 2022 at 06:03 PM UTC

### Review by @rverdile - Commented on June 01, 2022 at 06:27 PM UTC

### Review by @jlsherrill - Commented on June 01, 2022 at 06:30 PM UTC

### Review by @rverdile - Commented on June 01, 2022 at 06:37 PM UTC

### Review by @rverdile - Approved on June 02, 2022 at 01:24 PM UTC

lgtm. tested deploying and no issues.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/22*
