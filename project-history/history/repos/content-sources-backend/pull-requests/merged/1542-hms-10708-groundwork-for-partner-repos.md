---
type: pull_request
number: 1542
title: "HMS-10708: groundwork for partner repos"
state: merged
author: TenSt
created: 2026-06-17T14:02:02Z
updated: 2026-06-24T13:33:18Z
closed: 2026-06-24T13:33:18Z
merged: 2026-06-24T13:33:18Z
base_branch: main
head_branch: stepan/HMS-10708-groundwork-for-partner-repos
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1542
---

# Pull Request #1542: HMS-10708: groundwork for partner repos

**Author**: @TenSt
**Created**: June 17, 2026 at 02:02 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `stepan/HMS-10708-groundwork-for-partner-repos`

## Description

## Summary
This PR:
- Adds `partner` and `published` DB columns with indexes for partner repository visibility
- Exposes `partner` and `published` on repository/snapshot API responses and `?partner=` filter field
- Maps new model fields through `ModelToApiFields` and `SnapshotModelToApi`
- Introduces `partner_visibility.go` helpers (`IsForeignPartnerView`, `ForeignPartnerVisibleExpr`, `HasPublishedSnapshot`) for next user stories
- Adds unit and DAO tests for partner visibility predicates and published-snapshot checks
- Updates the API client and docs

NOTE: The api client for playwright and docs are generated with the script, so they can be skimmed.

## Testing steps
Run the unit tests:
```
CONFIG_PATH="$(pwd)/configs/" go test ./pkg/dao/ -run 'TestIsForeignPartnerView|TestPartnerVisibilitySuite' -count=1 -v
```

---

## Discussion

### Comment by @xbhouse on June 17, 2026 at 02:30 PM UTC

https://issues.redhat.com/browse/HMS-10708

---

## Reviews

### Review by @Starle21 - Approved on June 24, 2026 at 11:25 AM UTC

I picked Stepan's brains on this PR that were not clear to me, here goes a few notes:
- this PR builds a minimal foundation for the partner repos feature that other tickets will follow on without needing to worry about modelling data
- focuses on adding fields to the model of `RepositoryConfiguration` and `Snapshot`
- we talked about why "partner repo" is modeled with the `partner` flag on `RepositoryConfiguration` and not as a new type in the `origin` of `Repository` 
- - partner repo is only of type `upload` for the moment, but in the near future (with `COPR` repositories) other repository types might also be extended to provide "partner" functionality
- - so partner repo is not really a new "type" of a `Repository`, more like a property of all the repository types
-  partner repo logic is not wired and used up at the moment - eg. `FilterData` or logic in `partner_visibility`, since these are only the model foundations and logic will be added later
-  the design of the data model for partner repos (the partner flag for `RepositoryConfiguration` and the `published` flag for `Snapshot`) is minimal so no bigger changes are needed in the codebase

Thank you Stepan for the explanations! 🎉 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1542*
