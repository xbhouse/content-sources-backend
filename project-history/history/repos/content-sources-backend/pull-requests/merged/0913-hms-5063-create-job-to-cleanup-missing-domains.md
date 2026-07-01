---
type: pull_request
number: 913
title: "HMS-5063: create job to cleanup missing domains"
state: merged
author: xbhouse
created: 2024-12-04T20:25:22Z
updated: 2024-12-19T14:43:56Z
closed: 2024-12-19T14:43:56Z
merged: 2024-12-19T14:43:56Z
base_branch: main
head_branch: 5063
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/913
---

# Pull Request #913: HMS-5063: create job to cleanup missing domains

**Author**: @xbhouse
**Created**: December 04, 2024 at 08:25 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `5063`

## Description

## Summary

Adds a job to clean up the domains in our DB that are not in pulp

## Testing steps

You can use the pulp CLI to test this, setup docs [here](https://pulpproject.org/pulp-cli/)

1. Create 2 repos in different orgs and let them snapshot
2. List domains in pulp via the pulp CLI: `pulp domain list | jq '.[].name'`
3. Delete one of the repos
4. Delete the domain associated to the deleted repo's org ID in pulp via the pulp CLI: `pulp domain destroy --name <domain_name>`
5. List the domains in pulp again to confirm the domains that are left
6. Run this job to clean up the domains in our DB: `go run cmd/jobs/main.go cleanup-missing-domains`
7. List the domains in our DB. They should now match pulp's cs-prefixed domains


---

## Discussion

### Comment by @jlsherrill on December 04, 2024 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-5063

### Comment by @jlsherrill on December 17, 2024 at 06:40 PM UTC

couple of comments, but otherwise looks good!  I'll do a full test after updates

---

## Reviews

### Review by @jlsherrill - Commented on December 17, 2024 at 06:36 PM UTC

### Review by @jlsherrill - Commented on December 17, 2024 at 06:37 PM UTC

### Review by @xbhouse - Commented on December 18, 2024 at 09:12 PM UTC

### Review by @jlsherrill - Commented on December 18, 2024 at 09:15 PM UTC

### Review by @jlsherrill - Commented on December 19, 2024 at 01:54 PM UTC

### Review by @xbhouse - Commented on December 19, 2024 at 02:01 PM UTC

### Review by @jlsherrill - Approved on December 19, 2024 at 02:01 PM UTC

ACK! tested and worked great

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/913*
