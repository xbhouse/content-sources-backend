---
type: pull_request
number: 497
title: "Fixes 3250: provide filter for excluding loaded RH repos"
state: merged
author: jlsherrill
created: 2023-12-06T16:56:07Z
updated: 2023-12-13T15:30:31Z
closed: 2023-12-13T15:01:28Z
merged: 2023-12-13T15:01:28Z
base_branch: main
head_branch: HMS-3250
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/497
---

# Pull Request #497: Fixes 3250: provide filter for excluding loaded RH repos

**Author**: @jlsherrill
**Created**: December 06, 2023 at 04:56 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `HMS-3250`

## Description

## Summary

Provide an env variable that can be used to exclude some or all of the red hat repos from being loaded.  

## Testing steps

In dev,
In between each test, a ```make compose-clean compose-up``` is needed to clear out what is there.
With backend and frontend running:
```
go run ./cmd/external-repos/main.go import
```

should still import all red hat repos (can check in the UI).  
```
OPTIONS_REPOSITORY_IMPORT_FILTER="small" go run ./cmd/external-repos/main.go import
```

should import only 1 small ansible redhat repo

```
OPTIONS_REPOSITORY_IMPORT_FILTER="none" go run ./cmd/external-repos/main.go import
```
should cause no red hat repos to be imported.

For QE, you can use with bonfire: 

--set-parameter content-sources-backend/OPTIONS_REPOSITORY_IMPORT_FILTER=small
--set-parameter content-sources-backend/OPTIONS_REPOSITORY_IMPORT_FILTER=none

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on December 06, 2023 at 05:00 PM UTC

https://issues.redhat.com/browse/HMS-3250

### Comment by @jlsherrill on December 06, 2023 at 05:22 PM UTC

also, if you do test with ephemeral, you'll need the deployment.yaml changes here

### Comment by @swadeley on December 07, 2023 at 01:33 AM UTC

Hi, I confirm:

`--set-parameter content-sources-backend/OPTIONS_REPOSITORY_IMPORT_FILTER=small` results in just the ansible repo being deployed.

---

## Reviews

### Review by @jlsherrill - Commented on December 06, 2023 at 05:00 PM UTC

### Review by @rverdile - Approved on December 07, 2023 at 04:38 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/497*
