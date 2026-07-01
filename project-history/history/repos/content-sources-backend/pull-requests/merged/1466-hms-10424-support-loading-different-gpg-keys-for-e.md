---
type: pull_request
number: 1466
title: "HMS-10424: Support loading different gpg keys for ext release"
state: merged
author: jlsherrill
created: 2026-04-15T19:20:49Z
updated: 2026-04-16T15:11:48Z
closed: 2026-04-16T15:11:48Z
merged: 2026-04-16T15:11:48Z
base_branch: main
head_branch: HMS-10424
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1466
---

# Pull Request #1466: HMS-10424: Support loading different gpg keys for ext release

**Author**: @jlsherrill
**Created**: April 15, 2026 at 07:20 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `HMS-10424`

## Description

## Summary

Supports loading a separate gpg key file for extended support repos.  It is based off the minor version.  If its not set for a given minor version, the major version's gpg key is used.

## Testing steps

On main, `make compose-up`

set in config.yaml:

```
options:
  feature_filter: ["RHEL-OS-x86_64", "OPENSHIFT-OCP-x86_64", "RHEL-HA-x86_64", "RHEL-EUS-x86_64"]
```
```
features:
  extended_release_repos:
    enabled: true
```
 'make repos-import'

check the gpg keys:

```
make db-cli-connect

select length(gpg_key) from repository_configurations rc inner join repositories r on rc.repository_uuid = r.uuid  where r.url  ilike '%eus%';
```

switch to this branch and run  `make repos-import`

re-check the gpg keys, and see that they have changed for eus repos:

```
select length(gpg_key) from repository_configurations rc inner join repositories r on rc.repository_uuid = r.uuid  where r.url  ilike '%eus%';
```


Thist just checks the lenght, you might also spot check to make sure it matches what is intended.

```
 select gpg_key from repository_configurations rc inner join repositories r on rc.repository_uuid = r.uuid  where r.url = 'https://cdn.redhat.com/content/eus/rhel9/9.6/x86_64/supplementary/os/';
```




---

## Discussion

### Comment by @xbhouse on April 15, 2026 at 07:30 PM UTC

https://issues.redhat.com/browse/HMS-10424

---

## Reviews

### Review by @jlsherrill - Commented on April 15, 2026 at 07:23 PM UTC

### Review by @jlsherrill - Commented on April 15, 2026 at 07:25 PM UTC

### Review by @TenSt - Approved on April 16, 2026 at 08:58 AM UTC

lgtm!

### Review by @croissanne - Commented on April 16, 2026 at 02:12 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1466*
