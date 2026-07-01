---
type: pull_request
number: 1562
title: "chore(deps): bump node-forge from 1.3.3 to 1.4.0"
state: merged
author: dependabot
created: 2026-03-29T02:18:20Z
updated: 2026-04-02T13:45:47Z
closed: 2026-04-02T13:45:39Z
merged: 2026-04-02T13:45:39Z
base_branch: master
head_branch: dependabot/npm_and_yarn/node-forge-1.4.0
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1562
---

# Pull Request #1562: chore(deps): bump node-forge from 1.3.3 to 1.4.0

**Author**: @dependabot
**Created**: March 29, 2026 at 02:18 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/node-forge-1.4.0`

## Description

Bumps [node-forge](https://github.com/digitalbazaar/forge) from 1.3.3 to 1.4.0.
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/digitalbazaar/forge/blob/main/CHANGELOG.md">node-forge's changelog</a>.</em></p>
<blockquote>
<h2>1.4.0 - 2026-03-24</h2>
<h3>Security</h3>
<ul>
<li><strong>HIGH</strong>: Denial of Service in <code>BigInteger.modInverse()</code>
<ul>
<li>A Denial of Service (DoS) vulnerability exists due to an infinite loop in
the <code>BigInteger.modInverse()</code> function (inherited from the bundled jsbn
library). When <code>modInverse()</code> is called with a zero value as input, the
internal Extended Euclidean Algorithm enters an unreachable exit condition,
causing the process to hang indefinitely and consume 100% CPU.</li>
<li>Reported by Kr0emer.</li>
<li>CVE ID: <a href="https://www.cve.org/CVERecord?id=CVE-2026-33891">CVE-2026-33891</a></li>
<li>GHSA ID: <a href="https://github.com/digitalbazaar/forge/security/advisories/GHSA-5m6q-g25r-mvwx">GHSA-5gfm-wpxj-wjgq</a></li>
</ul>
</li>
<li><strong>HIGH</strong>: Signature forgery in RSA-PKCS due to ASN.1 extra field.
<ul>
<li>RSASSA PKCS#1 v1.5 signature verification accepts forged signatures for low
public exponent keys (e=3). Attackers can forge signatures by stuffing
&quot;garbage&quot; bytes within the ASN.1 structure in order to construct a
signature that passes verification, enabling Bleichenbacher style forgery.
This issue is similar to CVE-2022-24771, but adds bytes in an addition
field within the ASN.1 structure, rather than outside of it.</li>
<li>Additionally, forge does not validate that signatures include a minimum of
8 bytes of padding as defined by the specification, providing attackers
additional space to construct Bleichenbacher forgeries.</li>
<li>Reported as part of a U.C. Berkeley security research project by:
<ul>
<li>Austin Chu, Sohee Kim, and Corban Villa.</li>
</ul>
</li>
<li>CVE ID: <a href="https://www.cve.org/CVERecord?id=CVE-2026-33894">CVE-2026-33894</a></li>
<li>GHSA ID: <a href="https://github.com/digitalbazaar/forge/security/advisories/GHSA-ppp5-5v6c-4jwp">GHSA-ppp5-5v6c-4jwp</a></li>
</ul>
</li>
<li><strong>HIGH</strong>: Signature forgery in Ed25519 due to missing S &lt; L check.
<ul>
<li>Ed25519 signature verification accepts forged non-canonical signatures
where the scalar S is not reduced modulo the group order (S &gt;= L). A valid
signature and its S + L variant both verify in forge, while Node.js
crypto.verify (OpenSSL-backed) rejects the S + L variant, as defined by the
specification. This class of signature malleability has been exploited in
practice to bypass authentication and authorization logic (see
CVE-2026-25793, CVE-2022-35961). Applications relying on signature
uniqueness (i.e., dedup by signature bytes, replay tracking, signed-object
canonicalization checks) may be bypassed.</li>
<li>Reported as part of a U.C. Berkeley security research project by:
<ul>
<li>Austin Chu, Sohee Kim, and Corban Villa.</li>
</ul>
</li>
<li>CVE ID: <a href="https://www.cve.org/CVERecord?id=CVE-2026-33895">CVE-2026-33895</a></li>
<li>GHSA ID: <a href="https://github.com/digitalbazaar/forge/security/advisories/GHSA-q67f-28xg-22rw">GHSA-q67f-28xg-22rw</a></li>
</ul>
</li>
<li><strong>HIGH</strong>: <code>basicConstraints</code> bypass in certificate chain verification.
<ul>
<li><code>pki.verifyCertificateChain()</code> does not enforce RFC 5280 <code>basicConstraints</code>
requirements when an intermediate certificate lacks both the
<code>basicConstraints</code> and <code>keyUsage</code> extensions. This allows any leaf
certificate (without these extensions) to act as a CA and sign other
certificates, which node-forge will accept as valid.</li>
<li>Reported by Doruk Tan Ozturk (<a href="https://github.com/peaktwilight"><code>@​peaktwilight</code></a>) - doruk.ch</li>
<li>CVE ID: <a href="https://www.cve.org/CVERecord?id=CVE-2026-33896">CVE-2026-33896</a></li>
<li>GHSA ID: <a href="https://github.com/digitalbazaar/forge/security/advisories/GHSA-2328-f5f3-gj25">GHSA-2328-f5f3-gj25</a></li>
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
<li><a href="https://github.com/digitalbazaar/forge/commit/fa385f92440879601240020f158bed68e444e83a"><code>fa385f9</code></a> Release 1.4.0.</li>
<li><a href="https://github.com/digitalbazaar/forge/commit/07d4e162762ed4fdab5caca9ebf78237fcf85339"><code>07d4e16</code></a> Update changelog.</li>
<li><a href="https://github.com/digitalbazaar/forge/commit/cb90fd92091ee34e4abab3ad0c835eeea3d06c3e"><code>cb90fd9</code></a> Update changelog.</li>
<li><a href="https://github.com/digitalbazaar/forge/commit/963e7c5c7b0f03de1b28a1e5a42a6bafda4cf711"><code>963e7c5</code></a> Add unit test for &quot;pseudonym&quot;</li>
<li><a href="https://github.com/digitalbazaar/forge/commit/f0b6f5b7c5d1c918240e975e0cade4f47d005446"><code>f0b6f5b</code></a> Add pseudonym OID</li>
<li><a href="https://github.com/digitalbazaar/forge/commit/3df48a311d4b53dc6493b7a47a8d07f3669957d9"><code>3df48a3</code></a> Fix missing CVE ID.</li>
<li><a href="https://github.com/digitalbazaar/forge/commit/2e492832fb25227e6b647cbe1ac981c123171e90"><code>2e49283</code></a> Add x509 <code>basicConstraints</code> check.</li>
<li><a href="https://github.com/digitalbazaar/forge/commit/bdecf11571c9f1a487cc0fe72fe78ff6dfa96b85"><code>bdecf11</code></a> Add canonical signature scaler check for S &lt; L.</li>
<li><a href="https://github.com/digitalbazaar/forge/commit/af094e69c60ac5f7b29f2b1957c53ae5e12fd4a0"><code>af094e6</code></a> Add RSA padding and DigestInfo length checks.</li>
<li><a href="https://github.com/digitalbazaar/forge/commit/796eeb1673f6ec636fda02dfc295047d9f7aefe0"><code>796eeb1</code></a> Improve jsbn fix.</li>
<li>Additional commits viewable in <a href="https://github.com/digitalbazaar/forge/compare/v1.3.3...v1.4.0">compare view</a></li>
</ul>
</details>
<br />


---

## Discussion

### Comment by @codecov-commenter on March 29, 2026 at 02:20 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1562?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 75.88%. Comparing base ([`a08d0d6`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/a08d0d69deb64ea8523582130fbabff1dde6be51?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`0034448`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/00344484a6e043293f40a7057a1ee404d14628ec?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 5 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1562   +/-   ##
=======================================
  Coverage   75.88%   75.88%           
=======================================
  Files         103      103           
  Lines        3164     3164           
  Branches      686      686           
=======================================
  Hits         2401     2401           
  Misses        684      684           
  Partials       79       79           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1562?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @swadeley - Approved on April 02, 2026 at 01:33 PM UTC

ACK

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1562*
