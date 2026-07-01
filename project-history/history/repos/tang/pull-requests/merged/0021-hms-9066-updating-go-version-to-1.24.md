---
type: pull_request
number: 21
title: "HMS-9066: Updating go version to 1.24"
state: merged
author: mayurilahane
created: 2025-09-18T04:32:26Z
updated: 2025-10-07T14:09:36Z
closed: 2025-10-07T14:09:36Z
merged: 2025-10-07T14:09:36Z
base_branch: main
head_branch: mlahane/HMS-9066
labels: []
url: https://github.com/content-services/tang/pull/21
---

# Pull Request #21: HMS-9066: Updating go version to 1.24

**Author**: @mayurilahane
**Created**: September 18, 2025 at 04:32 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `mlahane/HMS-9066`

## Description

*No description provided*

---

## Discussion

### Comment by @mayurilahane on September 24, 2025 at 03:45 PM UTC

Getting lint errors which suggest to use the pattern to set verison is 1.24.x 

and i have set it to 1.24 which seems a better option to me 

### Comment by @rverdile on September 24, 2025 at 05:04 PM UTC

It's a little confusing, but the linter errors are actually caused by this line here: https://github.com/content-services/tang/blob/main/.github/workflows/tang-actions.yaml#L23

It needs to be updated to a newer version of golangci-lint. I would say v2.4 to match the backend: https://github.com/content-services/content-sources-backend/blob/main/.github/workflows/content-sources-actions.yml#L51-L53

### Comment by @mayurilahane on September 24, 2025 at 06:56 PM UTC

> It's a little confusing, but the linter errors are actually caused by this line here: https://github.com/content-services/tang/blob/main/.github/workflows/tang-actions.yaml#L23
> 
> It needs to be updated to a newer version of golangci-lint. I would say v2.4 to match the backend: https://github.com/content-services/content-sources-backend/blob/main/.github/workflows/content-sources-actions.yml#L51-L53

how do you figure it out ? that it should match with backend golangci-lint version ?
i thought it will update by itself after `go get -u`  cmd 

### Comment by @rverdile on September 24, 2025 at 07:04 PM UTC

> how do you figure it out ? that it should match with backend golangci-lint version ?
i thought it will update by itself after go get -u cmd

Good question. `go get -u ./...` will only upgrade the dependencies imported in the go.mod. golangci-lint is not a dependency, but a CLI tool for linting. We install and use it locally with `make lint`. so everything probably worked locally. However, the golangci-lint action is a separate github action, independent of what we use locally. We should maybe make some change to the action to match what we use locally, but just bumping the version is good for now :)

### Comment by @rverdile on September 24, 2025 at 07:43 PM UTC

there's a CI error. sounds like you just need to bump the action version
```
Failed to run: Error: invalid version string 'v2.4.0', golangci-lint v2 is not supported by golangci-lint-action v6, you must update to golangci-lint-action v7., Error: invalid version string 'v2.4.0', golangci-lint v2 is not supported by golangci-lint-action v6, you must update to golangci-lint-action v7
```

### Comment by @mayurilahane on September 25, 2025 at 09:07 PM UTC

/retest

### Comment by @marusak on September 30, 2025 at 12:27 PM UTC

> - Bump Go version to 1.24 and golangci-lint to v2.4.0 in GitHub Actions workflow
> - Update .golangci.yaml for compatibility with golangci-lint v2
> - Refactor internal/zestwrapper/rpm.go to use safer HTTP response body closing with error handling and improved logging
> - Switch to math/rand/v2 in pkg/tangy/queries.go for randomness, remove golang.org/x/exp/rand (deprecated) dependency
> - Update dependancies in go.mod and go.sum

There seem to be ~4 different changes? But they are all in the same commit :thinking: Are they strictly connected thus being in one commit?

### Comment by @rverdile on September 30, 2025 at 01:11 PM UTC

> There seem to be ~4 different changes? But they are all in the same commit 🤔 Are they strictly connected thus being in one commit?

@marusak They're all connected to bumping the go version and failures related to that

### Comment by @mayurilahane on September 30, 2025 at 01:25 PM UTC

> > * Bump Go version to 1.24 and golangci-lint to v2.4.0 in GitHub Actions workflow
> > * Update .golangci.yaml for compatibility with golangci-lint v2
> > * Refactor internal/zestwrapper/rpm.go to use safer HTTP response body closing with error handling and improved logging
> > * Switch to math/rand/v2 in pkg/tangy/queries.go for randomness, remove golang.org/x/exp/rand (deprecated) dependency
> > * Update dependancies in go.mod and go.sum
> 
> There seem to be ~4 different changes? But they are all in the same commit 🤔 Are they strictly connected thus being in one commit?

Hey yes, these are connected.
I updated go and go linter version.

Updating linter version, recomended these changes which looked good to me in pkg/tangy/queries.go and  internal/zestwrapper/rpm.go.

Although i need to undo the changes in  internal/zestwrapper/rpm.go as per 
https://github.com/content-services/tang/pull/21#discussion_r2391394979



---

## Reviews

### Review by @rverdile - Commented on September 24, 2025 at 12:02 PM UTC

I would do a dependency bump here: `go get -u ./...`

### Review by @rverdile - Commented on September 30, 2025 at 01:10 PM UTC

### Review by @rverdile - Commented on October 01, 2025 at 04:14 PM UTC

### Review by @rverdile - Approved on October 07, 2025 at 02:07 PM UTC

looks good! going to merge and release a new version. Then you'll be able to update your backend PR

---

*Archived from: https://github.com/content-services/tang/pull/21*
