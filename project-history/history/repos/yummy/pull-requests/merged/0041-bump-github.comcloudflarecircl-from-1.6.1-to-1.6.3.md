---
type: pull_request
number: 41
title: "Bump github.com/cloudflare/circl from 1.6.1 to 1.6.3"
state: merged
author: dependabot
created: 2026-02-25T19:31:36Z
updated: 2026-05-14T19:35:41Z
closed: 2026-05-14T19:35:08Z
merged: 2026-05-14T19:35:08Z
base_branch: master
head_branch: dependabot/go_modules/github.com/cloudflare/circl-1.6.3
labels: ["dependencies", "go"]
url: https://github.com/content-services/yummy/pull/41
---

# Pull Request #41: Bump github.com/cloudflare/circl from 1.6.1 to 1.6.3

**Author**: @dependabot
**Created**: February 25, 2026 at 07:31 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `master` ← **Head**: `dependabot/go_modules/github.com/cloudflare/circl-1.6.3`

## Description

Bumps [github.com/cloudflare/circl](https://github.com/cloudflare/circl) from 1.6.1 to 1.6.3.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/cloudflare/circl/releases">github.com/cloudflare/circl's releases</a>.</em></p>
<blockquote>
<h2>CIRCL v1.6.3</h2>
<p>Fix a bug on ecc/p384 scalar multiplication.</p>
<h3>What's Changed</h3>
<ul>
<li>sign/mldsa: Check opts for nil value  by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/582">cloudflare/circl#582</a></li>
<li>ecc/p384: Point addition must handle point doubling case. by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/583">cloudflare/circl#583</a></li>
<li>Release CIRCL v1.6.3 by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/584">cloudflare/circl#584</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/cloudflare/circl/compare/v1.6.2...v1.6.3">https://github.com/cloudflare/circl/compare/v1.6.2...v1.6.3</a></p>
<h2>CIRCL v1.6.2</h2>
<ul>
<li>New SLH-DSA, improvements in ML-DSA for arm64.</li>
<li>Tested compilation on WASM.</li>
</ul>
<h2>What's Changed</h2>
<ul>
<li>Optimize pairing product computation by moving exponentiations to G1. by <a href="https://github.com/dfaranha"><code>@​dfaranha</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/547">cloudflare/circl#547</a></li>
<li>sign: Adding SLH-DSA signature by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/512">cloudflare/circl#512</a></li>
<li>Update code generators to CIRCL v1.6.1. by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/548">cloudflare/circl#548</a></li>
<li>ML-DSA: Add preliminary Wycheproof test vectors by <a href="https://github.com/bwesterb"><code>@​bwesterb</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/552">cloudflare/circl#552</a></li>
<li>go fmt by <a href="https://github.com/bwesterb"><code>@​bwesterb</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/554">cloudflare/circl#554</a></li>
<li>gz-compressing test vectors, use of HexBytes and ReadGzip functions. by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/555">cloudflare/circl#555</a></li>
<li>group: Removes use of elliptic Marshal and Unmarshal functions. by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/556">cloudflare/circl#556</a></li>
<li>Support encoding/decoding ML-DSA private keys (as long as they contain seeds) by <a href="https://github.com/bwesterb"><code>@​bwesterb</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/559">cloudflare/circl#559</a></li>
<li>Update to golangci-lint v2 by <a href="https://github.com/bwesterb"><code>@​bwesterb</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/560">cloudflare/circl#560</a></li>
<li>Preparation for ARM64 Implementation of poly operations for dilithium package. by <a href="https://github.com/elementrics"><code>@​elementrics</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/562">cloudflare/circl#562</a></li>
<li>prepare power2Round for custom implementations in assembly by <a href="https://github.com/elementrics"><code>@​elementrics</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/564">cloudflare/circl#564</a></li>
<li>ARM64 implementation for poly.PackLe16 by <a href="https://github.com/elementrics"><code>@​elementrics</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/563">cloudflare/circl#563</a></li>
<li>add arm64 version of polyMulBy2toD by <a href="https://github.com/elementrics"><code>@​elementrics</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/565">cloudflare/circl#565</a></li>
<li>add arm64 version of polySub by <a href="https://github.com/elementrics"><code>@​elementrics</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/566">cloudflare/circl#566</a></li>
<li>group: add byteLen method for short groups and RandomScalar uses rand.Int by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/568">cloudflare/circl#568</a></li>
<li>add arm64 version of poly.Add/Sub by <a href="https://github.com/elementrics"><code>@​elementrics</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/572">cloudflare/circl#572</a></li>
<li>group: Adding cryptobyte marshaling to scalars by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/569">cloudflare/circl#569</a></li>
<li>Bumping up to Go1.25 by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/574">cloudflare/circl#574</a></li>
<li>ci: Including WASM compilation. by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/577">cloudflare/circl#577</a></li>
<li>Revert to using package-declared HPKE errors for shortkem instead of standard library errors by <a href="https://github.com/harshiniwho"><code>@​harshiniwho</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/578">cloudflare/circl#578</a></li>
<li>Release v1.6.2 by <a href="https://github.com/armfazh"><code>@​armfazh</code></a> in <a href="https://redirect.github.com/cloudflare/circl/pull/579">cloudflare/circl#579</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/dfaranha"><code>@​dfaranha</code></a> made their first contribution in <a href="https://redirect.github.com/cloudflare/circl/pull/547">cloudflare/circl#547</a></li>
<li><a href="https://github.com/elementrics"><code>@​elementrics</code></a> made their first contribution in <a href="https://redirect.github.com/cloudflare/circl/pull/562">cloudflare/circl#562</a></li>
<li><a href="https://github.com/harshiniwho"><code>@​harshiniwho</code></a> made their first contribution in <a href="https://redirect.github.com/cloudflare/circl/pull/578">cloudflare/circl#578</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/cloudflare/circl/compare/v1.6.1...v1.6.2">https://github.com/cloudflare/circl/compare/v1.6.1...v1.6.2</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/cloudflare/circl/commit/24ae53c5d6f7fe18203adc125ba3ed76a38703e1"><code>24ae53c</code></a> Release CIRCL v1.6.3</li>
<li><a href="https://github.com/cloudflare/circl/commit/581020bd4a836b8ce7bd4e414ba2884c07dbc906"><code>581020b</code></a> Rename method to oddMultiplesProjective.</li>
<li><a href="https://github.com/cloudflare/circl/commit/12209a4566605692a8402594e367a5aed5148460"><code>12209a4</code></a> Removing unused cmov for jacobian points.</li>
<li><a href="https://github.com/cloudflare/circl/commit/fcba359f4178645d2c9e50f29ab6966337da4b95"><code>fcba359</code></a> ecc/p384: use of complete projective formulas for scalar multiplication.</li>
<li><a href="https://github.com/cloudflare/circl/commit/5e1bae8d8c2df4e717c2c5c2d5b5d60b629b2ac6"><code>5e1bae8</code></a> ecc/p384: handle point doubling in point addition with Jacobian coordinates.</li>
<li><a href="https://github.com/cloudflare/circl/commit/341604685ff97e8f7440ae4b4711ba1c118c648c"><code>3416046</code></a> Check opts for nil value.</li>
<li><a href="https://github.com/cloudflare/circl/commit/a763d47a6dce43d1f4f7b697d1d7810463a526f6"><code>a763d47</code></a> Release CIRCL v1.6.2</li>
<li><a href="https://github.com/cloudflare/circl/commit/3c70bf9ad53b681fbe5ba6067e454a86549fee8a"><code>3c70bf9</code></a> Bump x/crypto x/sys dependencies.</li>
<li><a href="https://github.com/cloudflare/circl/commit/3f0f15b2bfe67bad81a35e8aec81ae42ca78349d"><code>3f0f15b</code></a> Revert to using package-declared HPKE errors for shortkem instead of standard...</li>
<li><a href="https://github.com/cloudflare/circl/commit/23491bd573cf29b6f567057a158203a2c9dfa30d"><code>23491bd</code></a> Adding generic Power2Round method.</li>
<li>Additional commits viewable in <a href="https://github.com/cloudflare/circl/compare/v1.6.1...v1.6.3">compare view</a></li>
</ul>
</details>
<br />


> **Note**
> Automatic rebases have been disabled on this pull request as it has been open for over 30 days.


---

## Discussion

### Comment by @rverdile on March 04, 2026 at 05:46 PM UTC

@dependabot rebase

### Comment by @rverdile on May 14, 2026 at 04:46 PM UTC

@dependabot recreate

---

## Reviews

### Review by @rverdile - Approved on May 14, 2026 at 07:35 PM UTC

---

*Archived from: https://github.com/content-services/yummy/pull/41*
