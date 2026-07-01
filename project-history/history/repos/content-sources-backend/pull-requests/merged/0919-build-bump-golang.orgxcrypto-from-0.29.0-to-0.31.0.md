---
type: pull_request
number: 919
title: "Build: Bump golang.org/x/crypto from 0.29.0 to 0.31.0"
state: merged
author: dependabot
created: 2024-12-12T14:03:30Z
updated: 2024-12-13T02:03:58Z
closed: 2024-12-13T02:03:50Z
merged: 2024-12-13T02:03:50Z
base_branch: main
head_branch: dependabot/go_modules/golang.org/x/crypto-0.31.0
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/919
---

# Pull Request #919: Build: Bump golang.org/x/crypto from 0.29.0 to 0.31.0

**Author**: @dependabot
**Created**: December 12, 2024 at 02:03 PM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/golang.org/x/crypto-0.31.0`

## Description

Bumps [golang.org/x/crypto](https://github.com/golang/crypto) from 0.29.0 to 0.31.0.
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/golang/crypto/commit/b4f1988a35dee11ec3e05d6bf3e90b695fbd8909"><code>b4f1988</code></a> ssh: make the public key cache a 1-entry FIFO cache</li>
<li><a href="https://github.com/golang/crypto/commit/7042ebcbe097f305ba3a93f9a22b4befa4b83d29"><code>7042ebc</code></a> openpgp/clearsign: just use rand.Reader in tests</li>
<li><a href="https://github.com/golang/crypto/commit/3e90321ac7bcee3d924ed63ed3ad97be2079cb56"><code>3e90321</code></a> go.mod: update golang.org/x dependencies</li>
<li><a href="https://github.com/golang/crypto/commit/8c4e668694ccbaa1be4785da7e7a40f2ef93152b"><code>8c4e668</code></a> x509roots/fallback: update bundle</li>
<li>See full diff in <a href="https://github.com/golang/crypto/compare/v0.29.0...v0.31.0">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=golang.org/x/crypto&package-manager=go_modules&previous-version=0.29.0&new-version=0.31.0)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

Dependabot will resolve any conflicts with this PR as long as you don't alter it yourself. You can also trigger a rebase manually by commenting `@dependabot rebase`.

[//]: # (dependabot-automerge-start)
[//]: # (dependabot-automerge-end)

---

<details>
<summary>Dependabot commands and options</summary>
<br />

You can trigger Dependabot actions by commenting on this PR:
- `@dependabot rebase` will rebase this PR
- `@dependabot recreate` will recreate this PR, overwriting any edits that have been made to it
- `@dependabot merge` will merge this PR after your CI passes on it
- `@dependabot squash and merge` will squash and merge this PR after your CI passes on it
- `@dependabot cancel merge` will cancel a previously requested merge and block automerging
- `@dependabot reopen` will reopen this PR if it is closed
- `@dependabot close` will close this PR and stop Dependabot recreating it. You can achieve the same result by closing it manually
- `@dependabot show <dependency name> ignore conditions` will show all of the ignore conditions of the specified dependency
- `@dependabot ignore this major version` will close this PR and stop Dependabot creating any more for this major version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this minor version` will close this PR and stop Dependabot creating any more for this minor version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this dependency` will close this PR and stop Dependabot creating any more for this dependency (unless you reopen the PR or upgrade to it yourself)
You can disable automated security fix PRs for this repo from the [Security Alerts page](https://github.com/content-services/content-sources-backend/network/alerts).

</details>

---

## Reviews

### Review by @swadeley - Approved on December 13, 2024 at 02:03 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/919*
