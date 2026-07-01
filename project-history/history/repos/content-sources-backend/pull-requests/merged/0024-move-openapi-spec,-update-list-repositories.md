---
type: pull_request
number: 24
title: "Move openapi spec, update list repositories"
state: merged
author: mshriver
created: 2022-06-02T18:20:35Z
updated: 2022-06-03T12:55:05Z
closed: 2022-06-03T12:54:38Z
merged: 2022-06-03T12:54:38Z
base_branch: main
head_branch: update-repositories-spec
labels: ["bug"]
url: https://github.com/content-services/content-sources-backend/pull/24
---

# Pull Request #24: Move openapi spec, update list repositories

**Author**: @mshriver
**Created**: June 02, 2022 at 06:20 PM UTC
**Status**: Merged
**Labels**: `bug`
**Base**: `main` ← **Head**: `update-repositories-spec`

## Description

Trailing slash is required by the routing, but wasn't included in the spec.  Auto-generated openapi client (IQE) ends up building an invalid URL.

Move from docs to api directory, repository boilerplate indicates that's where it should go.

I've built and run the container locally to validate the repositories endpoint still works for GET. 
I've taken the generated spec and fed it into the IQE openapi client generator to validate the client using the correct URL. 
Trying to deploy that container within ephemeral and make an IQE rest client request against it now. 

---

## Discussion

### Comment by @mshriver on June 02, 2022 at 09:00 PM UTC

The `Format` check is failing on one of the auto-generated files, could use some advice there.

---

## Reviews

### Review by @mshriver - Commented on June 02, 2022 at 08:52 PM UTC

### Review by @jlsherrill - Commented on June 02, 2022 at 09:07 PM UTC

### Review by @jlsherrill - Commented on June 02, 2022 at 09:08 PM UTC

### Review by @mshriver - Commented on June 02, 2022 at 09:17 PM UTC

### Review by @jlsherrill - Commented on June 03, 2022 at 12:53 PM UTC

### Review by @jlsherrill - Commented on June 03, 2022 at 12:53 PM UTC

### Review by @jlsherrill - Approved on June 03, 2022 at 12:54 PM UTC

### Review by @jlsherrill - Commented on June 03, 2022 at 12:55 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/24*
