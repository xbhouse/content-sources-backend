---
type: pull_request
number: 947
title: "Build: Bump the go group with 4 updates"
state: merged
author: dependabot
created: 2025-01-20T04:37:05Z
updated: 2025-01-20T17:10:26Z
closed: 2025-01-20T17:10:19Z
merged: 2025-01-20T17:10:19Z
base_branch: main
head_branch: dependabot/go_modules/go-55219f93cb
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/947
---

# Pull Request #947: Build: Bump the go group with 4 updates

**Author**: @dependabot
**Created**: January 20, 2025 at 04:37 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-55219f93cb`

## Description

Bumps the go group with 4 updates: [github.com/ProtonMail/go-crypto](https://github.com/ProtonMail/go-crypto), [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2), [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) and [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2).

Updates `github.com/ProtonMail/go-crypto` from 1.1.4 to 1.1.5
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/ProtonMail/go-crypto/releases">github.com/ProtonMail/go-crypto's releases</a>.</em></p>
<blockquote>
<h2>Release v1.1.5</h2>
<h2>What's Changed</h2>
<ul>
<li>Check binding signature details against primary key by <a href="https://github.com/twiss"><code>@​twiss</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/264">ProtonMail/go-crypto#264</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/ProtonMail/go-crypto/compare/v1.1.4...v1.1.5">https://github.com/ProtonMail/go-crypto/compare/v1.1.4...v1.1.5</a></p>
<h2>Release v1.1.5-proton</h2>
<h2>What's Changed</h2>
<p>This release is v1.1.5 with support for the following non-standardized features:</p>
<ul>
<li>Presistent symmetric keys <a href="https://www.ietf.org/archive/id/draft-ietf-openpgp-persistent-symmetric-keys-00.html">draft-ietf-openpgp-persistent-symmetric-keys-00</a></li>
<li>Automatic forwarding <a href="https://www.ietf.org/archive/id/draft-wussler-openpgp-forwarding-00.html">draft-wussler-openpgp-forwarding-00</a></li>
<li>Post-quantum algorithms <a href="https://datatracker.ietf.org/doc/draft-ietf-openpgp-pqc/">draft-ietf-openpgp-pqc</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/d703f494c90d2fa80fdbbaf33c084a7225411b55"><code>d703f49</code></a> Check binding signature details against primary key (<a href="https://redirect.github.com/ProtonMail/go-crypto/issues/264">#264</a>)</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/72cacd5a68e18e0178a7ac1d1e06209fec8b9f98"><code>72cacd5</code></a> ci: Update openpgp-interop-test-analyzer to v2.1.0 (<a href="https://redirect.github.com/ProtonMail/go-crypto/issues/243">#243</a>)</li>
<li>See full diff in <a href="https://github.com/ProtonMail/go-crypto/compare/v1.1.4...v1.1.5">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.32.8 to 1.33.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8fac7d48b4707a8fdd9cb23b34b928fc83e38777"><code>8fac7d4</code></a> Release 2025-01-15</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9eff8acbf1cc8bb294171b969a8e2803bf235a07"><code>9eff8ac</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f47aa72637a4492512213811a314dda0b9c9a189"><code>f47aa72</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5934dd36b8ecbc42a321216e8da721e1258ae8a8"><code>5934dd3</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6636822440828c3eebaacfe9a182c9eb47895236"><code>6636822</code></a> feat: flexible checksum updates (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2808">#2808</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4ffbb7cbe6bfeebb50a773d45df57d3fd126cafb"><code>4ffbb7c</code></a> Release 2025-01-14</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2f70ff490c3474a59ca29dc2661ec5e4f13ec2f7"><code>2f70ff4</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/434159f5c088eb7b484fea9179d6c5e6302e4610"><code>434159f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a80fac9aee1efaac1d16f4005ccc93fed22c62c3"><code>a80fac9</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/44e779abca3ec6dfdc9415ced6af0aa684149bd1"><code>44e779a</code></a> fix several waiter issues (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2953">#2953</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.32.8...v1.33.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.51 to 1.17.54
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fa8a15e1e023fd02edb11264d3e926d65f18f87f"><code>fa8a15e</code></a> Release 2025-01-17</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e08f26d7616d64b6b61c8ebf6b85464e2c801932"><code>e08f26d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/177bcda0d62b5dba1c9e361de38dd3705ca748b3"><code>177bcda</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a5b6a2f1ee1078d88aeecf4f4f5947de77c79994"><code>a5b6a2f</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/74b33baa9e818d9b297f4a9753e0426ed23f510c"><code>74b33ba</code></a> pull retry loop forward to cover everything from resolving auth scheme onward...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7ea97e576564775a287646a2b992f7b31a48b785"><code>7ea97e5</code></a> Add integration test case for flex checksum (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2967">#2967</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c578cddaeb62b06d2c610dac6ef9038649363029"><code>c578cdd</code></a> Release 2025-01-16</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/45afab32d943d9281cd207adc2668b3c8d2774d6"><code>45afab3</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/30703fedc699008a606d283d96aa267e3dd74b37"><code>30703fe</code></a> Update partitions file</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/de0c8ef9ca33f33a769059444d1ad72893eb145b"><code>de0c8ef</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.51...credentials/v1.17.54">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.45.3 to 1.45.6
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d125de3792b20980da07dd1424afa90285d50c90"><code>d125de3</code></a> Release 2024-11-18</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fec51f3fff1b0f1a2cc913f4f7874977a3d47d0d"><code>fec51f3</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fba59970453e163b476913a363a126412641b5fb"><code>fba5997</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0b8e5c842fc161b15aa5e13c8a0caf1357b5e9c7"><code>0b8e5c8</code></a> Bump smithy-go dependency (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2902">#2902</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/50ba45ce162cd2458c65f5da799fd907ad826561"><code>50ba45c</code></a> Release 2024-11-15.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/444bdffccd6dce18f60ae626b74c087641c3d052"><code>444bdff</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/55ab381b20235964ab0c670a29d096821e158e90"><code>55ab381</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/94c083768b80bbd372c0e9feb45f02511442b896"><code>94c0837</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2398a7903ca141692f98da65b8537a4a53e9707e"><code>2398a79</code></a> Remove elastictranscoder service's integration test (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2901">#2901</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/93e0f294f6c692e39adf1b8ec2c8681ba9ee5f01"><code>93e0f29</code></a> Release 2024-11-15</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/rds/v1.45.3...service/polly/v1.45.6">compare view</a></li>
</ul>
</details>
<br />


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
- `@dependabot ignore <dependency name> major version` will close this group update PR and stop Dependabot creating any more for the specific dependency's major version (unless you unignore this specific dependency's major version or upgrade to it yourself)
- `@dependabot ignore <dependency name> minor version` will close this group update PR and stop Dependabot creating any more for the specific dependency's minor version (unless you unignore this specific dependency's minor version or upgrade to it yourself)
- `@dependabot ignore <dependency name>` will close this group update PR and stop Dependabot creating any more for the specific dependency (unless you unignore this specific dependency or upgrade to it yourself)
- `@dependabot unignore <dependency name>` will remove all of the ignore conditions of the specified dependency
- `@dependabot unignore <dependency name> <ignore condition>` will remove the ignore condition of the specified dependency and ignore conditions


</details>

---

## Discussion

### Comment by @jlsherrill on January 20, 2025 at 03:12 PM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Approved on January 20, 2025 at 05:10 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/947*
