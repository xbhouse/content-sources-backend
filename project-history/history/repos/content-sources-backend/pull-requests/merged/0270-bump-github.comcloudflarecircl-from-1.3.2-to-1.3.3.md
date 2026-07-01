---
type: pull_request
number: 270
title: "Bump github.com/cloudflare/circl from 1.3.2 to 1.3.3"
state: merged
author: dependabot
created: 2023-05-11T20:47:26Z
updated: 2023-05-17T08:21:03Z
closed: 2023-05-17T08:20:54Z
merged: 2023-05-17T08:20:53Z
base_branch: main
head_branch: dependabot/go_modules/github.com/cloudflare/circl-1.3.3
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/270
---

# Pull Request #270: Bump github.com/cloudflare/circl from 1.3.2 to 1.3.3

**Author**: @dependabot
**Created**: May 11, 2023 at 08:47 PM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/github.com/cloudflare/circl-1.3.3`

## Description

Bumps [github.com/cloudflare/circl](https://github.com/cloudflare/circl) from 1.3.2 to 1.3.3.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/cloudflare/circl/releases">github.com/cloudflare/circl's releases</a>.</em></p>
<blockquote>
<h2>CIRCL v1.3.3</h2>
<h2>New Features</h2>
<ul>
<li><a href="https://ascon.iaik.tugraz.at/">ASCON</a> light-weight authenticated encryption.</li>
<li>Hybrid KEM for HPKE based on Kyber and X25519.</li>
<li>CIRCL can be compiled both as static and dynamic linking modes.</li>
</ul>
<h2>Security</h2>
<ul>
<li>Fixes error-handling on rand readers.</li>
</ul>
<h2>What's Changed</h2>
<ul>
<li>Use untyped consts for Kyber params by <a href="https://github.com/tmthrgd"><code>@​tmthrgd</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/398">cloudflare/circl#398</a></li>
<li>zk/dl: adds prefixed labels and updates nomenclature. by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/396">cloudflare/circl#396</a></li>
<li>Bumping Go version. by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/399">cloudflare/circl#399</a></li>
<li>kem: add P-256 + Kyber768Draft00 hybrid by <a href="https://github.com/bwesterb"><code>@​bwesterb</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/402">cloudflare/circl#402</a></li>
<li>ckem: pass xof to elliptic.GenerateKey directly by <a href="https://github.com/bwesterb"><code>@​bwesterb</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/403">cloudflare/circl#403</a></li>
<li>Adding Ascon, an AEAD lightweight cipher. by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/400">cloudflare/circl#400</a></li>
<li>Add Ascon-80pq to cipher\ascon by <a href="https://github.com/dhcgn"><code>@​dhcgn</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/404">cloudflare/circl#404</a></li>
<li>ascon: update formulas and check for API compatibility by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/406">cloudflare/circl#406</a></li>
<li>all: enables dynamic linking, removes R15 is clobbered by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/407">cloudflare/circl#407</a></li>
<li>ascon: Removes table of constants. by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/408">cloudflare/circl#408</a></li>
<li>tkn20: prevent panics on key gen errors by <a href="https://github.com/tmthrgd"><code>@​tmthrgd</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/409">cloudflare/circl#409</a></li>
<li>expander,tkn20: remove superfluous Reset calls by <a href="https://github.com/tmthrgd"><code>@​tmthrgd</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/410">cloudflare/circl#410</a></li>
<li>Updating stdlib crypto library. by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/413">cloudflare/circl#413</a></li>
<li>Reduce x/crypto and x/sys versions to match Go 1.20 by <a href="https://github.com/Lekensteyn"><code>@​Lekensteyn</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/414">cloudflare/circl#414</a></li>
<li>Make ascon cipher go routine safe by <a href="https://github.com/enj"><code>@​enj</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/416">cloudflare/circl#416</a></li>
<li>tkn20,kyber,x25519,x448: plug constant-time leaks by <a href="https://github.com/tmthrgd"><code>@​tmthrgd</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/411">cloudflare/circl#411</a></li>
<li>Check for crypto/rand errors and ReadFull io.Readers by <a href="https://github.com/bwesterb"><code>@​bwesterb</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/417">cloudflare/circl#417</a></li>
<li>Fix encapsulation seed size by <a href="https://github.com/chris-wood"><code>@​chris-wood</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/419">cloudflare/circl#419</a></li>
<li>Add X25519Kyber768Draft00 experimental HPKE KEM by <a href="https://github.com/chris-wood"><code>@​chris-wood</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/421">cloudflare/circl#421</a></li>
<li>hpke: Adding NonceSize function to AEAD. by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/424">cloudflare/circl#424</a></li>
<li>hpke: Address always nil parameter. by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/425">cloudflare/circl#425</a></li>
<li>hpke: update and move xyber768d00 test vectors by <a href="https://github.com/bwesterb"><code>@​bwesterb</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/426">cloudflare/circl#426</a></li>
<li>hpke: fix encapsulation seed in test for xyber by <a href="https://github.com/bwesterb"><code>@​bwesterb</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/428">cloudflare/circl#428</a></li>
<li>Remove scalar sha3 amd64 assembly by <a href="https://github.com/bwesterb"><code>@​bwesterb</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/429">cloudflare/circl#429</a></li>
<li>Add HPKE benchmarks by <a href="https://github.com/chris-wood"><code>@​chris-wood</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/434">cloudflare/circl#434</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/tmthrgd"><code>@​tmthrgd</code></a> made their first contribution in <a href="https://redirect.github.com/cloudflare/circl/pull/398">cloudflare/circl#398</a></li>
<li><a href="https://github.com/dhcgn"><code>@​dhcgn</code></a> made their first contribution in <a href="https://redirect.github.com/cloudflare/circl/pull/404">cloudflare/circl#404</a></li>
<li><a href="https://github.com/Lekensteyn"><code>@​Lekensteyn</code></a> made their first contribution in <a href="https://redirect.github.com/cloudflare/circl/pull/414">cloudflare/circl#414</a></li>
<li><a href="https://github.com/enj"><code>@​enj</code></a> made their first contribution in <a href="https://redirect.github.com/cloudflare/circl/pull/416">cloudflare/circl#416</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/cloudflare/circl/compare/v1.3.2...v1.3.3">https://github.com/cloudflare/circl/compare/v1.3.2...v1.3.3</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/cloudflare/circl/commit/3bef500f2b925f150815a360b90081021e082939"><code>3bef500</code></a> Releasing CIRCL v1.3.3</li>
<li><a href="https://github.com/cloudflare/circl/commit/4002bafcebdd3b32974f70cf86a4682f82d9b3b5"><code>4002baf</code></a> Add HPKE benchmarks</li>
<li><a href="https://github.com/cloudflare/circl/commit/795540340d5c79e5768a0135741cd7c3e5f7de93"><code>7955403</code></a> Remove scalar sha3 amd64 assembly</li>
<li><a href="https://github.com/cloudflare/circl/commit/aef72508ab9bf35177e84ae23f94170f4546b63e"><code>aef7250</code></a> hpke: fix encapsulation seed in test for xyber</li>
<li><a href="https://github.com/cloudflare/circl/commit/808526a555262691f406ceed2ac1e4e7421faf96"><code>808526a</code></a> hpke: update and move xyber768d00 test vectors</li>
<li><a href="https://github.com/cloudflare/circl/commit/c7845aa1035e0b2d0397663e0adc283fd16af50a"><code>c7845aa</code></a> Address always nil parameter.</li>
<li><a href="https://github.com/cloudflare/circl/commit/2475a3f4a6255da8795b2a8f0ec7e71e3ee6d37e"><code>2475a3f</code></a> Adding NonceSize function to AEAD.</li>
<li><a href="https://github.com/cloudflare/circl/commit/eaec71f4cccf05035481b034b8ce9dc8755118ec"><code>eaec71f</code></a> Add X25519Kyber768Draft00 experimental HPKE KEM</li>
<li><a href="https://github.com/cloudflare/circl/commit/f0db2881a9618356223ed31090cdb33feb1e6d23"><code>f0db288</code></a> Fix encapsulation seed size</li>
<li><a href="https://github.com/cloudflare/circl/commit/f4c0e87526ec17305e8a573f1c58acedc5539a92"><code>f4c0e87</code></a> Update go-ristretto dep</li>
<li>Additional commits viewable in <a href="https://github.com/cloudflare/circl/compare/v1.3.2...v1.3.3">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=github.com/cloudflare/circl&package-manager=go_modules&previous-version=1.3.2&new-version=1.3.3)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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
- `@dependabot ignore this major version` will close this PR and stop Dependabot creating any more for this major version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this minor version` will close this PR and stop Dependabot creating any more for this minor version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this dependency` will close this PR and stop Dependabot creating any more for this dependency (unless you reopen the PR or upgrade to it yourself)
You can disable automated security fix PRs for this repo from the [Security Alerts page](https://github.com/content-services/content-sources-backend/network/alerts).

</details>

---

## Discussion

### Comment by @app-sre-bot on May 11, 2023 at 08:47 PM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on May 11, 2023 at 09:13 PM UTC

[test]

### Comment by @swadeley on May 15, 2023 at 06:45 AM UTC

Hi, test failed because title failed to pass match regex - /(Fixes |Refs )

### Comment by @jlsherrill on May 15, 2023 at 12:36 PM UTC

@swadeley there's two iqe test failures too:

```
[tests.test_repositories.test_repository_crud[ViaREST]](https://ci.int.devshift.net/job/content-sources-backend-pr-check-main/14/testReport/junit/tests/test_repositories/test_repository_crud_ViaREST_/)
    [tests.test_repositories.test_rbac_for_repos[ViaREST]](https://ci.int.devshift.net/job/content-sources-backend-pr-check-main/14/testReport/junit/tests/test_repositories/test_rbac_for_repos_ViaREST_/)
```

### Comment by @swadeley on May 16, 2023 at 12:46 PM UTC

Hi,
for RBAC fail:
"Set my default user to be hms-qa (!170) · Merge requests · insights-qe / iqe-content-sources-plugin · GitLab"
for CRUD fail:
"Change introspection exception arg (!172) · Merge requests · insights-qe / iqe-content-sources-plugin · GitLab"

### Comment by @swadeley on May 16, 2023 at 05:09 PM UTC

/retest

---

## Reviews

### Review by @swadeley - Approved on May 16, 2023 at 05:18 PM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/270*
