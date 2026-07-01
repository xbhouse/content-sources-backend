---
type: pull_request
number: 1035
title: "HMS-5397: add job for marking domains by labels "
state: merged
author: dominikvagner
created: 2025-03-18T13:27:03Z
updated: 2025-03-27T08:48:33Z
closed: 2025-03-27T08:48:33Z
merged: 2025-03-27T08:48:33Z
base_branch: main
head_branch: 5397
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1035
---

# Pull Request #1035: HMS-5397: add job for marking domains by labels 

**Author**: @dominikvagner
**Created**: March 18, 2025 at 01:27 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `5397`

## Description

## Summary
This adds a new job that sets a 'content-sources' label on all of our domains if they don't have it. Also makes the domain creation add this label as well. Will be useful for the pulp service team for differentiating our domains from others. 🍊 
For this an update of zest and pulp containers was needed, along with local pulp compose config change. 🔧 

## Testing steps
1. Create any repo with snapshots enabled.
2. List pulp domains and verify that the newly created domain has the `contentsources` label set.
3. Comment out the line **133**: `domain.SetPulpLabels(map[string]string{"contentsources": "true"})` in the file `pkg/clients/pulp_client/domains.go`, rebuild and restart the server.
4. Create any repo with snapshots enabled in a different org from the first one.
5. Check that it doesn't have the label.
6. Run the job `go run ./cmd/jobs/main.go set-domain-label`.
7. Check that it now has that label.

---

## Discussion

### Comment by @jlsherrill on March 18, 2025 at 01:30 PM UTC

https://issues.redhat.com/browse/HMS-5397

### Comment by @jlsherrill on March 21, 2025 at 07:14 PM UTC

couple small things, but otherwise looks good!

---

## Reviews

### Review by @jlsherrill - Commented on March 21, 2025 at 04:20 PM UTC

### Review by @jlsherrill - Commented on March 21, 2025 at 07:13 PM UTC

### Review by @dominikvagner - Commented on March 25, 2025 at 01:23 PM UTC

### Review by @dominikvagner - Commented on March 25, 2025 at 01:25 PM UTC

### Review by @jlsherrill - Approved on March 25, 2025 at 08:27 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1035*
