---
type: pull_request
number: 7
title: "refactor to be more object oriented"
state: merged
author: rverdile
created: 2022-11-08T16:05:37Z
updated: 2022-12-05T14:58:07Z
closed: 2022-12-05T14:58:06Z
merged: 2022-12-05T14:58:06Z
base_branch: master
head_branch: refactor
labels: []
url: https://github.com/content-services/yummy/pull/7
---

# Pull Request #7: refactor to be more object oriented

**Author**: @rverdile
**Created**: November 08, 2022 at 04:05 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `refactor`

## Description

Refactors most yummy functions to be methods of a `Repository` struct. Adds `Signature` as method and  `FetchGPGKey` as utility function.

Repository structs now look like this:
```go
type YummySettings struct {
	Client *http.Client
	URL    *string
}

type YumRepository interface {
	Configure(settings YummySettings)
	Packages() (packages []Package, statusCode int, err error)
	Repomd() (repomd *Repomd, statusCode int, err error)
	Signature() (repomdSignature *string, statusCode int, err error)
}

type Repository struct {
	settings        YummySettings
	packages        []Package // Packages repository contains
	repomdSignature *string   // Signature of the repository
	repomd          *Repomd   // Repomd of the repository
}
```

`Packages()`, `Repomd()`, and `Signature()` only make a GET request if those fields on `Repository` are `nil`.

The Repomd struct was changed to this:
```go
// Repomd metadata of the repomd of a repository
type Repomd struct {
	XMLName      xml.Name `xml:"repomd"`
	Data         []Data   `xml:"data"`
	Revision     string   `xml:"revision"`
	RepomdString *string  `xml:"-"`
}
```
`RepomdString` is to replace this usecase: https://github.com/content-services/content-sources-backend/blob/main/pkg/dao/repository_config.go#L501

Todo:

- [x] Better tests
- [x] Update README

---

## Discussion

### Comment by @jlsherrill on November 29, 2022 at 09:23 PM UTC

nice work! this improves test coverage by ~30%, taking yummy up to ~75%

### Comment by @jlsherrill on December 02, 2022 at 12:58 PM UTC

and i guess we should tag this as v1.0.0 once its merged??

---

## Reviews

### Review by @jlsherrill - Commented on November 08, 2022 at 08:56 PM UTC

### Review by @jlsherrill - Commented on November 08, 2022 at 08:58 PM UTC

### Review by @jlsherrill - Commented on November 08, 2022 at 09:01 PM UTC

### Review by @jlsherrill - Commented on November 08, 2022 at 09:15 PM UTC

### Review by @rverdile - Commented on November 09, 2022 at 04:46 PM UTC

### Review by @rverdile - Commented on November 11, 2022 at 08:22 PM UTC

### Review by @rverdile - Commented on November 28, 2022 at 08:06 PM UTC

### Review by @jlsherrill - Commented on November 29, 2022 at 09:24 PM UTC

### Review by @jlsherrill - Approved on December 02, 2022 at 12:52 PM UTC

---

*Archived from: https://github.com/content-services/yummy/pull/7*
