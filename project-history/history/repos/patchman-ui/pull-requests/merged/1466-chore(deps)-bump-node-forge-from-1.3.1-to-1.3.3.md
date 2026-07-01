---
type: pull_request
number: 1466
title: "chore(deps): bump node-forge from 1.3.1 to 1.3.3"
state: merged
author: dependabot
created: 2026-01-09T20:58:25Z
updated: 2026-02-10T10:43:07Z
closed: 2026-02-10T10:42:55Z
merged: 2026-02-10T10:42:55Z
base_branch: master
head_branch: dependabot/npm_and_yarn/node-forge-1.3.3
labels: ["dependencies", "patch"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1466
---

# Pull Request #1466: chore(deps): bump node-forge from 1.3.1 to 1.3.3

**Author**: @dependabot
**Created**: January 09, 2026 at 08:58 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `patch`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/node-forge-1.3.3`

## Description

Bumps [node-forge](https://github.com/digitalbazaar/forge) from 1.3.1 to 1.3.3.
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/digitalbazaar/forge/blob/main/CHANGELOG.md">node-forge's changelog</a>.</em></p>
<blockquote>
<h2>1.3.3 - 2025-12-02</h2>
<h3>Fixed</h3>
<ul>
<li>[pkcs12] Make digestAlgorithm parameters optional to fix PKCS#12/PFX issues
introduced in 1.3.2.</li>
</ul>
<h2>1.3.2 - 2025-11-25</h2>
<h3>Security</h3>
<ul>
<li><strong>HIGH</strong>: ASN.1 Validator Desynchronization
<ul>
<li>An Interpretation Conflict (CWE-436) vulnerability in node-forge versions
1.3.1 and below enables remote, unauthenticated attackers to craft ASN.1
structures to desynchronize schema validations, yielding a semantic
divergence that may bypass downstream cryptographic verifications and
security decisions.</li>
<li>Reported by Hunter Wodzenski.</li>
<li>CVE ID: <a href="https://www.cve.org/CVERecord?id=CVE-2025-12816">CVE-2025-12816</a></li>
<li>GHSA ID: <a href="https://github.com/digitalbazaar/forge/security/advisories/GHSA-5gfm-wpxj-wjgq">GHSA-5gfm-wpxj-wjgq</a></li>
</ul>
</li>
<li><strong>HIGH</strong>: ASN.1 Unbounded Recursion
<ul>
<li>An Uncontrolled Recursion (CWE-674) vulnerability in node-forge versions
1.3.1 and below enables remote, unauthenticated attackers to craft deep
ASN.1 structures that trigger unbounded recursive parsing. This leads to a
Denial-of-Service (DoS) via stack exhaustion when parsing untrusted DER
inputs.</li>
<li>Reported by Hunter Wodzenski.</li>
<li>CVE ID: <a href="https://www.cve.org/CVERecord?id=CVE-2025-66031">CVE-2025-66031</a></li>
<li>GHSA ID: <a href="https://github.com/digitalbazaar/forge/security/advisories/GHSA-554w-wpv2-vw27">GHSA-554w-wpv2-vw27</a></li>
</ul>
</li>
<li><strong>MODERATE</strong>: ASN.1 OID Integer Truncation
<ul>
<li>An Integer Overflow (CWE-190) vulnerability in node-forge versions 1.3.1
and below enables remote, unauthenticated attackers to craft ASN.1
structures containing OIDs with oversized arcs. These arcs may be decoded
as smaller, trusted OIDs due to 32-bit bitwise truncation, enabling the
bypass of downstream OID-based security decisions.</li>
<li>Reported by Hunter Wodzenski.</li>
<li>CVE ID: <a href="https://www.cve.org/CVERecord?id=CVE-2025-66030">CVE-2025-66030</a></li>
<li>GHSA ID: <a href="https://github.com/digitalbazaar/forge/security/advisories/GHSA-65ch-62r8-g69g">GHSA-65ch-62r8-g69g</a></li>
</ul>
</li>
</ul>
<h3>Fixed</h3>
<ul>
<li>[asn1] Fix for vulnerability identified by CVE-2025-12816 PKCS#12 MAC
verification bypass due to missing macData enforcement and improper
asn1.validate routine.</li>
<li>[asn1] Add <code>fromDer()</code> max recursion depth check.
<ul>
<li>Add a <code>asn1.maxDepth</code> global configurable maximum depth of 256.</li>
<li>Add a <code>asn1.fromDer()</code> per-call <code>maxDepth</code> option.</li>
<li><strong>NOTE</strong>: The default maximum is assumed to be higher than needed for valid
data. If this assumption is false then this could be a breaking change.
Please file an issue if there are use cases that need a higher maximum.</li>
<li><strong>NOTE</strong>: The per-call <code>maxDepth</code> parameter has not been exposed up through
all of the API stack due to the complexities involved. Please file an issue
if there are use cases that require this instead of changing the default</li>
</ul>
</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/digitalbazaar/forge/commit/1cea0aff4901589ae86e314f25782bbe312f9f69"><code>1cea0af</code></a> Release 1.3.3.</li>
<li><a href="https://github.com/digitalbazaar/forge/commit/5265989cf5e54cfe1e27a10d71523007ce0507b1"><code>5265989</code></a> Update changelog.</li>
<li><a href="https://github.com/digitalbazaar/forge/commit/e4f3961406395dd8e985dcf841852ceca73ac3a9"><code>e4f3961</code></a> Fix changelog for release.</li>
<li><a href="https://github.com/digitalbazaar/forge/commit/503979b0295cf633a30199d6bd937f4a222481a0"><code>503979b</code></a> Update changelog.</li>
<li><a href="https://github.com/digitalbazaar/forge/commit/c3b3b32a8c157ac57752934d3af63b5f798b58b8"><code>c3b3b32</code></a> Make digestAlgorithm parameters optional</li>
<li><a href="https://github.com/digitalbazaar/forge/commit/6f70043a6db1abb9f3304f3d432efed3ba50fcca"><code>6f70043</code></a> Update CVE details.</li>
<li><a href="https://github.com/digitalbazaar/forge/commit/f547b0d292745094190ecb250429d21e8804a375"><code>f547b0d</code></a> Start 1.3.3-0.</li>
<li><a href="https://github.com/digitalbazaar/forge/commit/235ad3e70e4fdfdca4fdeb662dfba6588e2c38bd"><code>235ad3e</code></a> Release 1.3.2.</li>
<li><a href="https://github.com/digitalbazaar/forge/commit/25982441171dc9815c87d3d886c5c8a1d092b334"><code>2598244</code></a> Update changelog.</li>
<li><a href="https://github.com/digitalbazaar/forge/commit/0032dd0be8b6fb1b1092ef754d1dde91c10a95ad"><code>0032dd0</code></a> Fix typos.</li>
<li>Additional commits viewable in <a href="https://github.com/digitalbazaar/forge/compare/v1.3.1...v1.3.3">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=node-forge&package-manager=npm_and_yarn&previous-version=1.3.1&new-version=1.3.3)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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
You can disable automated security fix PRs for this repo from the [Security Alerts page](https://github.com/RedHatInsights/patchman-ui/network/alerts).

</details>

> **Note**
> Automatic rebases have been disabled on this pull request as it has been open for over 30 days.


---

## Discussion

### Comment by @codecov-commenter on February 10, 2026 at 10:32 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1466?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.53%. Comparing base ([`87ec4bd`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/87ec4bdb48b1c3c0e5e99802e10cc4341592d4c7?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`9ec1e62`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/9ec1e628c4396299a97ce6cc5267bc4551e4fca6?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1466   +/-   ##
=======================================
  Coverage   62.53%   62.53%           
=======================================
  Files         127      127           
  Lines        3310     3310           
  Branches      899      899           
=======================================
  Hits         2070     2070           
  Misses       1120     1120           
  Partials      120      120           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1466?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @swadeley - Approved on February 10, 2026 at 10:29 AM UTC

ACK

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1466*
