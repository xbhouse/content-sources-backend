---
type: pull_request
number: 941
title: "HMS-5056: add layered RH repos with feature name"
state: merged
author: xbhouse
created: 2025-01-14T21:12:53Z
updated: 2025-01-16T18:27:56Z
closed: 2025-01-16T18:27:56Z
merged: 2025-01-16T18:27:56Z
base_branch: main
head_branch: 5056
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/941
---

# Pull Request #941: HMS-5056: add layered RH repos with feature name

**Author**: @xbhouse
**Created**: January 14, 2025 at 09:12 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `5056`

## Description

## Summary

- Adds layered RH repos (OCP and HA) to redhat_repos.json and allows importing of RH repos to be controlled by the feature name
- Adds feature_name as a new column on the repository_configurations table and imports this for RH repos. For custom repos, this value will be null

## Testing steps

1. Set `options.feature_filter` in config.yaml to include only the "RHEL-OS-x86_64" feature name (see config.yaml.example)
2. Run `make repos-import && go run cmd/external-repos/main.go nightly-jobs`. Only the 14 pre-existing RH repos should be imported and snapshotted
3. Update `options.feature_filter` to also include the "OPENSHIFT-OCP-x86_64" and "RHEL-HA-x86_64" features and restart the server
4. Run`make repos-import && go run cmd/external-repos/main.go nightly-jobs` again. You should see the 5 additional layered RH repos (OCP and HA) imported and snapshotted
5. If `options.feature_filter` is not set in config.yaml, the default ("RHEL-OS-x86_64") should be used and only the pre-existing RH repos should be imported and snapshotted


---

## Discussion

### Comment by @jlsherrill on January 14, 2025 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-5056

### Comment by @jlsherrill on January 15, 2025 at 08:08 PM UTC

small comment, otherwise looks good!

---

## Reviews

### Review by @jlsherrill - Commented on January 14, 2025 at 09:56 PM UTC

### Review by @jlsherrill - Commented on January 15, 2025 at 03:27 PM UTC

### Review by @xbhouse - Commented on January 15, 2025 at 04:03 PM UTC

### Review by @jlsherrill - Commented on January 15, 2025 at 08:08 PM UTC

### Review by @jlsherrill - Approved on January 16, 2025 at 04:31 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/941*
