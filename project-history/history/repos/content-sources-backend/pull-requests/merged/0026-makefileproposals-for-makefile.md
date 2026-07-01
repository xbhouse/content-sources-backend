---
type: pull_request
number: 26
title: "makefile:proposals for Makefile"
state: merged
author: avisiedo
created: 2022-06-07T13:24:26Z
updated: 2022-06-10T14:55:25Z
closed: 2022-06-10T14:55:25Z
merged: 2022-06-10T14:55:25Z
base_branch: main
head_branch: proposal-makefile-changes
labels: []
url: https://github.com/content-services/content-sources-backend/pull/26
---

# Pull Request #26: makefile:proposals for Makefile

**Author**: @avisiedo
**Created**: June 07, 2022 at 01:24 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `proposal-makefile-changes`

## Description

A set of changes based on small snippets I had from other repos:

- Set default values for necessary variables in `mk/variables.mk`.
- Add generic rules to build golang binaries at `cmd/dir/**`, and other helper rules `tidy` and `vendor`.
- Add generic rule to generate diagrams from plantuml files at `docs/` directory.
- Install swag locally to the repository and add dependency in `openapi` rule.
- Add rules to start/stop postgres from a container (useful for developing from the workstation).
- Add generic rules to build and push container images.
- Set `help` as default rule, and generate help for the rules from the comments.

Signed-off-by: Alejandro Visiedo <avisiedo@redhat.com>

---

## Discussion

### Comment by @jlsherrill on June 07, 2022 at 08:32 PM UTC

i'm in favor of most all of this!   I do have some questions about the change to use './vendor' though.  Could you explain a bit more about that reasoning?  It doesn't look like you've checked it in, so i'm imagining you don't intend it to be checked in (which i wouldn't want).  What purpose is it serving?  

### Comment by @rverdile on June 07, 2022 at 08:59 PM UTC

I like these changes!

### Comment by @avisiedo on June 08, 2022 at 10:24 AM UTC

> i'm in favor of most all of this! I do have some questions about the change to use './vendor' though. Could you explain a bit more about that reasoning? It doesn't look like you've checked it in, so i'm imagining you don't intend it to be checked in (which i wouldn't want). What purpose is it serving?

My thoughts about vendor (that could be wrong ;) ):

- The origin was to keep the dependencies attached to the
  project repository and keep that dependencies under
  version control.
- In the practice I have seen it was evoking
  PR which are difficult to review because even a small
  change could evoke big changes in the PR (updated
  dependencies that was evoking to download the whole
  library referenced into go.mod file at the vendor
  directory). Personally I see this as the biggest downside.
- Then, why I am making use of it? If I understand correctly,
  golang download the source code of external libraries at
  `$GOPATH/src`, but for one library only one version is
  enabled, which means if a repository is using `v1.4.5` and
  other is using `v1.5.3`, this could evoke problems when
  building code in different repos.
  - It could impacts working from a workstation, where in a
    moment we are using this repository, and a few hours later
    we could be using another golang repository which
    dependencies could differs. This in the past has evoked to me
    a lot of headhakes with golang dependencies, having to
    remove the current `$GOPATH` directory and recreate it again.
    Or maybe in the past I did wrong and something that I am
    missing take me to reset the golang environment to restore
    the situation :) 
- Said the above, and if I am not wrong on my thoughts, the benefit
  of using the `vendor/` directory but not include it
  into the repository are that we keep our dependencies with no
  impact to/from other repositories, and as it is not included
  into the repository we avoid the noise that could be evoked
  when updating `vendor/` directory. So we would have a
  `vendor/` directory when working with the repo, but that
  directory is not committed to the repository, but keep the
  dependencies in that directory instead of the
  `$GOPATH/src` directory, isolating our dependencies when
  working with the repository.

If I were a wrong on my understanding about some part of golang (likely
as I see myself a learner on this language yet), please do not hesitate
to comment it so that I can correct my thoughts and learn about it.

### Comment by @avisiedo on June 08, 2022 at 10:29 AM UTC

I have updated the rule that install swag to make it more isolated.

### Comment by @jlsherrill on June 08, 2022 at 05:38 PM UTC

I *think* think that this point:
 
> Then, why I am making use of it? If I understand correctly,
> golang download the source code of external libraries at
> $GOPATH/src, but for one library only one version is
> enabled, which means if a repository is using v1.4.5 and
> other is using v1.5.3, this could evoke problems when
> building code in different repos.

Isn't actually true.  I have multiple versions of the same package in my $GOPATH/pkg.  Its very possible that that was true in the past (as it seems like Go modules have improved a lot in recent tiems).   From my understand (and i'm still very new to Go) that go.mod will store known good dependency versions (and go.sum includes sub-dependency versions as well), so keeping them all in ./vendor/ doesn't really buy you much version 'locking'.  You can replicate the same exact set of dependencies other installs.  Then if you want to update all your dependencies, you can simply run 'go get -u ./...'   && 'go mod tidy'

### Comment by @avisiedo on June 08, 2022 at 06:52 PM UTC

Thanks for the clarification. I think then that vendor/ is not providing benefit; I will review the PR and I will remove it to avoid noise into the code. Thanks for the comment.

### Comment by @avisiedo on June 09, 2022 at 04:36 AM UTC

I have added a couple of changes:
- Alias rule 'install-swag' to install the swag binary in `release/`.
- Rule 'printvars' that can be useful to know the values assigned to the variables.

### Comment by @avisiedo on June 09, 2022 at 08:33 AM UTC

NOT MERGE - I have just realized the container does not build with this changes.

Fixed: the container is being built locally, it likely will work in the pipeline when merged.

### Comment by @avisiedo on June 10, 2022 at 07:09 AM UTC

**update**: I have added a rule 'db-cli-connect' to open a postgres client cli in the container.

---

## Reviews

### Review by @jlsherrill - Commented on June 07, 2022 at 08:45 PM UTC

### Review by @jlsherrill - Commented on June 07, 2022 at 08:50 PM UTC

### Review by @avisiedo - Commented on June 08, 2022 at 08:28 AM UTC

### Review by @avisiedo - Commented on June 08, 2022 at 08:34 AM UTC

### Review by @jlsherrill - Commented on June 08, 2022 at 08:02 PM UTC

### Review by @avisiedo - Commented on June 09, 2022 at 04:00 AM UTC

### Review by @avisiedo - Commented on June 09, 2022 at 11:23 AM UTC

### Review by @jlsherrill - Commented on June 09, 2022 at 03:52 PM UTC

### Review by @jlsherrill - Commented on June 09, 2022 at 05:57 PM UTC

### Review by @avisiedo - Commented on June 10, 2022 at 05:08 AM UTC

### Review by @avisiedo - Commented on June 10, 2022 at 08:02 AM UTC

### Review by @jlsherrill - Commented on June 10, 2022 at 02:05 PM UTC

### Review by @jlsherrill - Approved on June 10, 2022 at 02:54 PM UTC

nice work!  This is awesome!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/26*
