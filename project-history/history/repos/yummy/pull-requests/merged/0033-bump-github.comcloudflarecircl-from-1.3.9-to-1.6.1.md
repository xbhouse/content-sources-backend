---
type: pull_request
number: 33
title: "Bump github.com/cloudflare/circl from 1.3.9 to 1.6.1"
state: merged
author: dependabot
created: 2025-06-10T21:33:34Z
updated: 2025-08-12T15:02:14Z
closed: 2025-08-12T15:02:10Z
merged: 2025-08-12T15:02:10Z
base_branch: master
head_branch: dependabot/go_modules/github.com/cloudflare/circl-1.6.1
labels: ["dependencies", "go"]
url: https://github.com/content-services/yummy/pull/33
---

# Pull Request #33: Bump github.com/cloudflare/circl from 1.3.9 to 1.6.1

**Author**: @dependabot
**Created**: June 10, 2025 at 09:33 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `master` ← **Head**: `dependabot/go_modules/github.com/cloudflare/circl-1.6.1`

## Description

Bumps [github.com/cloudflare/circl](https://github.com/cloudflare/circl) from 1.3.9 to 1.6.1.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/cloudflare/circl/releases">github.com/cloudflare/circl's releases</a>.</em></p>
<blockquote>
<h2>CIRCL v1.6.1</h2>
<ul>
<li>Fixes some point checks on the FourQ curve.</li>
<li>Hybrid KEM fails on low-order points.</li>
</ul>
<h3>What's Changed</h3>
<ul>
<li>kem/hybrid: ensure X25519 hybrids fails with low order points by <a href="https://github.com/Lekensteyn"><code>@​Lekensteyn</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/541">cloudflare/circl#541</a></li>
<li>.github: Use native ARM64 builders instead of QEMU by <a href="https://github.com/Lekensteyn"><code>@​Lekensteyn</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/542">cloudflare/circl#542</a></li>
<li>Fixes several errors on twisted Edwards curves. by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/545">cloudflare/circl#545</a></li>
<li>Release v1.6.1 by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/546">cloudflare/circl#546</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/cloudflare/circl/compare/v1.6.0...v1.6.1">https://github.com/cloudflare/circl/compare/v1.6.0...v1.6.1</a></p>
<h2>CIRCL v1.6.0</h2>
<h3>New!</h3>
<ul>
<li><a href="https://github.com/cloudflare/circl/blob/main/vdaf/prio3">Prio3</a> Verifiable Distributed Aggregation Function (<a href="https://datatracker.ietf.org/doc/draft-irtf-cfrg-vdaf/">draft-irtf-cfrg-vdaf</a>).</li>
<li><a href="https://github.com/cloudflare/circl/blob/main/kem/xwing">X-Wing</a>: general-purpose hybrid post-quantum KEM (<a href="https://datatracker.ietf.org/doc/draft-connolly-cfrg-xwing-kem/">draft-connolly-cfrg-xwing-kem</a>)</li>
</ul>
<h3>What's Changed</h3>
<ul>
<li>Add OIDs to ML-DSA by <a href="https://github.com/bwesterb"><code>@​bwesterb</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/519">cloudflare/circl#519</a></li>
<li>Adds Prio3 a set of verifiable distributed aggregation functions. by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/522">cloudflare/circl#522</a></li>
<li>Run semgrep cronjob only in upstream repository. by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/526">cloudflare/circl#526</a></li>
<li>X-Wing PQ/T hybrid by <a href="https://github.com/bwesterb"><code>@​bwesterb</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/471">cloudflare/circl#471</a></li>
<li>ckem: move crypto/elliptic to crypto/ecdh by <a href="https://github.com/MingLLuo"><code>@​MingLLuo</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/529">cloudflare/circl#529</a></li>
<li>hpke: Update HPKE code to use ecdh stdlib package. by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/530">cloudflare/circl#530</a></li>
<li>prio3: Adds polynomial multiplication using NTT by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/532">cloudflare/circl#532</a></li>
<li>Add Prio3 in readme. by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/527">cloudflare/circl#527</a></li>
</ul>
<h3>New Contributors</h3>
<ul>
<li><a href="https://github.com/MingLLuo"><code>@​MingLLuo</code></a> made their first contribution in <a href="https://redirect.github.com/cloudflare/circl/pull/529">cloudflare/circl#529</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/cloudflare/circl/compare/v1.5.0...v1.6.0">https://github.com/cloudflare/circl/compare/v1.5.0...v1.6.0</a></p>
<h1>CIRCL v1.5.0</h1>
<p><strong>New:</strong> ML-DSA, Module-Lattice-based Digital Signature Algorithm.</p>
<h3>What's Changed</h3>
<ul>
<li>kem: add X25519MLKEM768 TLS hybrid KEM by <a href="https://github.com/bwesterb"><code>@​bwesterb</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/510">cloudflare/circl#510</a></li>
<li>Create semgrep.yml by <a href="https://github.com/hrushikeshdeshpande"><code>@​hrushikeshdeshpande</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/514">cloudflare/circl#514</a></li>
<li>repo: Some fixes reported by CodeQL by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/515">cloudflare/circl#515</a></li>
<li>Add ML-DSA (FIPS204) by <a href="https://github.com/bwesterb"><code>@​bwesterb</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/480">cloudflare/circl#480</a></li>
<li>sign/mldsa: Add test for ML-DSA signature verification. by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/517">cloudflare/circl#517</a></li>
<li>Release v1.5.0 by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/518">cloudflare/circl#518</a></li>
</ul>
<h3>New Contributors</h3>
<ul>
<li><a href="https://github.com/hrushikeshdeshpande"><code>@​hrushikeshdeshpande</code></a> made their first contribution in <a href="https://redirect.github.com/cloudflare/circl/pull/514">cloudflare/circl#514</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/cloudflare/circl/compare/v1.4.0...v1.5.0">https://github.com/cloudflare/circl/compare/v1.4.0...v1.5.0</a></p>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/cloudflare/circl/commit/c6d33e35234ebf5c4319d12ae7d77d7d17053e56"><code>c6d33e3</code></a> Release v1.6.1</li>
<li><a href="https://github.com/cloudflare/circl/commit/0c3868ef6fc8ce864bc4104863186afdd2947f14"><code>0c3868e</code></a> curve4q: Shared must fail with low order points.</li>
<li><a href="https://github.com/cloudflare/circl/commit/9fd570dd508eef941d3f42fb94413a899b96d52e"><code>9fd570d</code></a> curve4q: Test showing DH does not fails on identity point.</li>
<li><a href="https://github.com/cloudflare/circl/commit/c988ceba827fe09896e770c152646dded447903d"><code>c988ceb</code></a> fourq: Correctly unmarshalling point.</li>
<li><a href="https://github.com/cloudflare/circl/commit/ef2611dcde7f6d25e31082412bbb30f2a870d133"><code>ef2611d</code></a> fourq: Test showing point unmarshal fails.</li>
<li><a href="https://github.com/cloudflare/circl/commit/05eba44d1a35f979c5f3ac914bcc50c1122e8ced"><code>05eba44</code></a> fourq: Handle the case of Z=0 for IsOnCurve and IsEqual.</li>
<li><a href="https://github.com/cloudflare/circl/commit/eef08780cc3cb9befa20014e65f731391103be6b"><code>eef0878</code></a> fourq: Test showing isEqual and IsOnCurve fail.</li>
<li><a href="https://github.com/cloudflare/circl/commit/2298474ef688938e4a81ca14990b9a11a8677e2a"><code>2298474</code></a> goldilocks; Handling points with z=0.</li>
<li><a href="https://github.com/cloudflare/circl/commit/5a940a111507232035d0b753fbf3068c52d6b8ac"><code>5a940a1</code></a> goldilocks: Test for IsEqual must fail with Z=0</li>
<li><a href="https://github.com/cloudflare/circl/commit/48c3b6a2746a18db4d8b675ab296980514359340"><code>48c3b6a</code></a> ed25519: Fix isEqual to handle points with Z=0.</li>
<li>Additional commits viewable in <a href="https://github.com/cloudflare/circl/compare/v1.3.9...v1.6.1">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=github.com/cloudflare/circl&package-manager=go_modules&previous-version=1.3.9&new-version=1.6.1)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

You can trigger a rebase of this PR by commenting `@dependabot rebase`.

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
You can disable automated security fix PRs for this repo from the [Security Alerts page](https://github.com/content-services/yummy/network/alerts).

</details>

> **Note**
> Automatic rebases have been disabled on this pull request as it has been open for over 30 days.


---

## Reviews

### Review by @jlsherrill - Approved on August 12, 2025 at 03:01 PM UTC

---

*Archived from: https://github.com/content-services/yummy/pull/33*
