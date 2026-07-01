---
type: pull_request
number: 1080
title: "Build: Bump the go group with 7 updates"
state: merged
author: dependabot
created: 2025-04-14T04:14:55Z
updated: 2025-04-14T17:00:39Z
closed: 2025-04-14T17:00:36Z
merged: 2025-04-14T17:00:35Z
base_branch: main
head_branch: dependabot/go_modules/go-4b13c79f43
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1080
---

# Pull Request #1080: Build: Bump the go group with 7 updates

**Author**: @dependabot
**Created**: April 14, 2025 at 04:14 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-4b13c79f43`

## Description

Bumps the go group with 7 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/ProtonMail/go-crypto](https://github.com/ProtonMail/go-crypto) | `1.1.6` | `1.2.0` |
| [github.com/prometheus/client_golang](https://github.com/prometheus/client_golang) | `1.21.1` | `1.22.0` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.29.13` | `1.29.14` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.66` | `1.17.67` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.47.2` | `1.47.3` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.79.1` | `1.79.2` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.31.1` | `0.32.0` |

Updates `github.com/ProtonMail/go-crypto` from 1.1.6 to 1.2.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/ProtonMail/go-crypto/releases">github.com/ProtonMail/go-crypto's releases</a>.</em></p>
<blockquote>
<h2>v1.2.0</h2>
<h2>What's Changed</h2>
<ul>
<li>Update min go version to 1.22.0 by <a href="https://github.com/lubux"><code>@​lubux</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/278">ProtonMail/go-crypto#278</a></li>
<li>Change the max AEAD chunk size to 4MiB from 64KiB by <a href="https://github.com/lubux"><code>@​lubux</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/280">ProtonMail/go-crypto#280</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/ProtonMail/go-crypto/compare/v1.1.6...v1.2.0">https://github.com/ProtonMail/go-crypto/compare/v1.1.6...v1.2.0</a></p>
<h2>Release v1.2.0-proton</h2>
<h2>What's Changed</h2>
<p>This release is v1.2.0 with support for the following non-standardized features:</p>
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
<li><a href="https://github.com/ProtonMail/go-crypto/commit/70930d7c5799a230834b3d0888ca341f0fd33361"><code>70930d7</code></a> Merge pull request <a href="https://redirect.github.com/ProtonMail/go-crypto/issues/280">#280</a> from ProtonMail/fix/aead-max-chunk-size</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/54e82fd684b78fd7120386cd0476ed10fe447594"><code>54e82fd</code></a> fix: The max AEAD chunk size must be 4MiB not 65KiB</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/61ae167ce5df0735ddcb0c53af42bb599a724fc4"><code>61ae167</code></a> Merge pull request <a href="https://redirect.github.com/ProtonMail/go-crypto/issues/278">#278</a> from ProtonMail/chore/bump-go-version</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/af200f275a9d5e2d9146519ec99d1de49a7c7adc"><code>af200f2</code></a> chore: Update min go version to 1.22.0</li>
<li>See full diff in <a href="https://github.com/ProtonMail/go-crypto/compare/v1.1.6...v1.2.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/prometheus/client_golang` from 1.21.1 to 1.22.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/releases">github.com/prometheus/client_golang's releases</a>.</em></p>
<blockquote>
<h2>v1.22.0 - 2025-04-07</h2>
<p>:warning: This release contains potential breaking change if you use experimental <code>zstd</code> support introduce in <a href="https://redirect.github.com/prometheus/client_golang/issues/1496">#1496</a> :warning:</p>
<p>Experimental support for <code>zstd</code> on scrape was added, controlled by the request <code>Accept-Encoding</code> header.
It was enabled by default since version 1.20, but now you need to add a blank import to enable it.
The decision to make it opt-in by default was originally made because the Go standard library was expected to have default zstd support added soon,
<a href="https://redirect.github.com/golang/go/issues/62513">golang/go#62513</a> however, the work took longer than anticipated and it will be postponed to upcoming major Go versions.</p>
<p>e.g.:</p>
<blockquote>
<pre lang="go"><code>import (
  _ &quot;github.com/prometheus/client_golang/prometheus/promhttp/zstd&quot;
)
</code></pre>
</blockquote>
<ul>
<li>[FEATURE] prometheus: Add new CollectorFunc utility <a href="https://redirect.github.com/prometheus/client_golang/issues/1724">#1724</a></li>
<li>[CHANGE] Minimum required Go version is now 1.22 (we also test client_golang against latest go version - 1.24) <a href="https://redirect.github.com/prometheus/client_golang/issues/1738">#1738</a></li>
<li>[FEATURE] api: <code>WithLookbackDelta</code> and <code>WithStats</code> options have been added to API client. <a href="https://redirect.github.com/prometheus/client_golang/issues/1743">#1743</a></li>
<li>[CHANGE] :warning: promhttp: Isolate zstd support and klauspost/compress library use to promhttp/zstd package. <a href="https://redirect.github.com/prometheus/client_golang/issues/1765">#1765</a></li>
</ul>
<!-- raw HTML omitted -->
<ul>
<li>build(deps): bump golang.org/x/sys from 0.28.0 to 0.29.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1720">prometheus/client_golang#1720</a></li>
<li>build(deps): bump google.golang.org/protobuf from 1.36.1 to 1.36.3 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1719">prometheus/client_golang#1719</a></li>
<li>Update RELEASE.md by <a href="https://github.com/bwplotka"><code>@​bwplotka</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1721">prometheus/client_golang#1721</a></li>
<li>chore(docs): Add links for the upstream PRs by <a href="https://github.com/kakkoyun"><code>@​kakkoyun</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1722">prometheus/client_golang#1722</a></li>
<li>Added tips on releasing client and checking with k8s. by <a href="https://github.com/bwplotka"><code>@​bwplotka</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1723">prometheus/client_golang#1723</a></li>
<li>feat: Add new CollectorFunc utility by <a href="https://github.com/Saumya40-codes"><code>@​Saumya40-codes</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1724">prometheus/client_golang#1724</a></li>
<li>build(deps): bump google.golang.org/protobuf from 1.36.3 to 1.36.4 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1725">prometheus/client_golang#1725</a></li>
<li>build(deps): bump the github-actions group with 5 updates by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1726">prometheus/client_golang#1726</a></li>
<li>Synchronize common files from prometheus/prometheus by <a href="https://github.com/prombot"><code>@​prombot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1727">prometheus/client_golang#1727</a></li>
<li>Synchronize common files from prometheus/prometheus by <a href="https://github.com/prombot"><code>@​prombot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1731">prometheus/client_golang#1731</a></li>
<li>build(deps): bump golang.org/x/sys from 0.29.0 to 0.30.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1739">prometheus/client_golang#1739</a></li>
<li>build(deps): bump google.golang.org/protobuf from 1.36.4 to 1.36.5 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1740">prometheus/client_golang#1740</a></li>
<li>Cleanup dependabot config by <a href="https://github.com/SuperQ"><code>@​SuperQ</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1741">prometheus/client_golang#1741</a></li>
<li>Upgrade Golang version v1.24 by <a href="https://github.com/dongjiang1989"><code>@​dongjiang1989</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1738">prometheus/client_golang#1738</a></li>
<li>build(deps): bump the github-actions group with 2 updates by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1742">prometheus/client_golang#1742</a></li>
<li>Merging 1.21 release back to main. by <a href="https://github.com/bwplotka"><code>@​bwplotka</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1744">prometheus/client_golang#1744</a></li>
<li>Synchronize common files from prometheus/prometheus by <a href="https://github.com/prombot"><code>@​prombot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1745">prometheus/client_golang#1745</a></li>
<li>Add support for undocumented query options for API by <a href="https://github.com/mahendrapaipuri"><code>@​mahendrapaipuri</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1743">prometheus/client_golang#1743</a></li>
<li>exp/api: Add experimental exp module; Add remote API with write client and handler. by <a href="https://github.com/bwplotka"><code>@​bwplotka</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1658">prometheus/client_golang#1658</a></li>
<li>exp/api: Add accepted msg type validation to handler by <a href="https://github.com/saswatamcode"><code>@​saswatamcode</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1750">prometheus/client_golang#1750</a></li>
<li>build(deps): bump the github-actions group with 5 updates by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1751">prometheus/client_golang#1751</a></li>
<li>build(deps): bump github.com/klauspost/compress from 1.17.11 to 1.18.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1752">prometheus/client_golang#1752</a></li>
<li>build(deps): bump github.com/google/go-cmp from 0.6.0 to 0.7.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1753">prometheus/client_golang#1753</a></li>
<li>exp: Reset snappy buf by <a href="https://github.com/saswatamcode"><code>@​saswatamcode</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1756">prometheus/client_golang#1756</a></li>
<li>Merge release 1.21.1 to main. by <a href="https://github.com/bwplotka"><code>@​bwplotka</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1762">prometheus/client_golang#1762</a></li>
<li>exp: Add dependabot config by <a href="https://github.com/saswatamcode"><code>@​saswatamcode</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1754">prometheus/client_golang#1754</a></li>
<li>build(deps): bump peter-evans/create-pull-request from 7.0.7 to 7.0.8 in the github-actions group by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1764">prometheus/client_golang#1764</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/blob/main/CHANGELOG.md">github.com/prometheus/client_golang's changelog</a>.</em></p>
<blockquote>
<h2>1.22.0 / 2025-04-07</h2>
<p>:warning: This release contains potential breaking change if you use experimental <code>zstd</code> support introduce in <a href="https://redirect.github.com/prometheus/client_golang/issues/1496">#1496</a> :warning:</p>
<p>Experimental support for <code>zstd</code> on scrape was added, controlled by the request <code>Accept-Encoding</code> header.
It was enabled by default since version 1.20, but now you need to add a blank import to enable it.
The decision to make it opt-in by default was originally made because the Go standard library was expected to have default zstd support added soon,
<a href="https://redirect.github.com/golang/go/issues/62513">golang/go#62513</a> however, the work took longer than anticipated and it will be postponed to upcoming major Go versions.</p>
<p>e.g.:</p>
<blockquote>
<pre lang="go"><code>import (
  _ &quot;github.com/prometheus/client_golang/prometheus/promhttp/zstd&quot;
)
</code></pre>
</blockquote>
<ul>
<li>[FEATURE] prometheus: Add new CollectorFunc utility <a href="https://redirect.github.com/prometheus/client_golang/issues/1724">#1724</a></li>
<li>[CHANGE] Minimum required Go version is now 1.22 (we also test client_golang against latest go version - 1.24) <a href="https://redirect.github.com/prometheus/client_golang/issues/1738">#1738</a></li>
<li>[FEATURE] api: <code>WithLookbackDelta</code> and <code>WithStats</code> options have been added to API client. <a href="https://redirect.github.com/prometheus/client_golang/issues/1743">#1743</a></li>
<li>[CHANGE] :warning: promhttp: Isolate zstd support and klauspost/compress library use to promhttp/zstd package. <a href="https://redirect.github.com/prometheus/client_golang/issues/1765">#1765</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/prometheus/client_golang/commit/d50be25511d790f4c166d68ce7d046c2977d148b"><code>d50be25</code></a> Cut 1.22.0 (<a href="https://redirect.github.com/prometheus/client_golang/issues/1793">#1793</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/1043db7cb8735186b341bbff6ae45d7ff56018dd"><code>1043db7</code></a> Cut 1.22.0-rc.0 (<a href="https://redirect.github.com/prometheus/client_golang/issues/1768">#1768</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/e575c9c04e40fe0784d4fee7f1f56c1a66ef090d"><code>e575c9c</code></a> promhttp: Isolate zstd support and klauspost/compress library use to promhttp...</li>
<li><a href="https://github.com/prometheus/client_golang/commit/f2276aa7d4b6e6526e0011762d0e4070513cf9cd"><code>f2276aa</code></a> Merge pull request <a href="https://redirect.github.com/prometheus/client_golang/issues/1764">#1764</a> from prometheus/dependabot/github_actions/github-act...</li>
<li><a href="https://github.com/prometheus/client_golang/commit/9df772cc5f399a2946a9158e57fd0aff66daf7d1"><code>9df772c</code></a> build(deps): bump peter-evans/create-pull-request</li>
<li><a href="https://github.com/prometheus/client_golang/commit/a3548c5aa811a0a52d1d723c3dfb3b01930481fb"><code>a3548c5</code></a> Merge pull request <a href="https://redirect.github.com/prometheus/client_golang/issues/1754">#1754</a> from saswatamcode/exp-eh</li>
<li><a href="https://github.com/prometheus/client_golang/commit/60fd2b0490db6e7c74e623651031feae93c7ebc1"><code>60fd2b0</code></a> Remove go.work file for now</li>
<li><a href="https://github.com/prometheus/client_golang/commit/8f9d0de6893dd7a470cdb1cffc5d98d1b5d7b50d"><code>8f9d0de</code></a> exp: Add dependabot config</li>
<li><a href="https://github.com/prometheus/client_golang/commit/c5cf981312d510414c209927e4f9e888d6776b5c"><code>c5cf981</code></a> Merge pull request <a href="https://redirect.github.com/prometheus/client_golang/issues/1762">#1762</a> from prometheus/release-1.21</li>
<li><a href="https://github.com/prometheus/client_golang/commit/e84c30512a658de7fdccce84b434c9a5696ec5af"><code>e84c305</code></a> exp: Reset snappy buf (<a href="https://redirect.github.com/prometheus/client_golang/issues/1756">#1756</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/prometheus/client_golang/compare/v1.21.1...v1.22.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.29.13 to 1.29.14
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2a0c73e76f5f06579a2cb24239ca054d60ced4c2"><code>2a0c73e</code></a> Release 2025-04-10</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b7b05ea6b196e40720c2620dc19fbd57945b11b3"><code>b7b05ea</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2a62cee847726f5a607d1801bebe9e1e77f4f400"><code>2a62cee</code></a> Update partitions file</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6af33138b18ea9a336f87e180fc9fdc286b1117e"><code>6af3313</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/915cbf3d7c2f9bc5dcd9a2a1ab02521f01c2f48c"><code>915cbf3</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ac82667a0ef9b9c97a9559810eb106d62bb1cd1e"><code>ac82667</code></a> Release 2025-04-09</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7db3afa50ca2f10dba50ffcb103f453010a2fdb8"><code>7db3afa</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f56d4c1e532d38245bbb638c73b669c536b3263e"><code>f56d4c1</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bf17208deff95cb9703ad5febb75f59af1da7fa9"><code>bf17208</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f529add9a2cd0d97281fd81f711c620c1e95cfb8"><code>f529add</code></a> Release 2025-04-08</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.29.13...config/v1.29.14">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.66 to 1.17.67
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2a0c73e76f5f06579a2cb24239ca054d60ced4c2"><code>2a0c73e</code></a> Release 2025-04-10</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b7b05ea6b196e40720c2620dc19fbd57945b11b3"><code>b7b05ea</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2a62cee847726f5a607d1801bebe9e1e77f4f400"><code>2a62cee</code></a> Update partitions file</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6af33138b18ea9a336f87e180fc9fdc286b1117e"><code>6af3313</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/915cbf3d7c2f9bc5dcd9a2a1ab02521f01c2f48c"><code>915cbf3</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ac82667a0ef9b9c97a9559810eb106d62bb1cd1e"><code>ac82667</code></a> Release 2025-04-09</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7db3afa50ca2f10dba50ffcb103f453010a2fdb8"><code>7db3afa</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f56d4c1e532d38245bbb638c73b669c536b3263e"><code>f56d4c1</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bf17208deff95cb9703ad5febb75f59af1da7fa9"><code>bf17208</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f529add9a2cd0d97281fd81f711c620c1e95cfb8"><code>f529add</code></a> Release 2025-04-08</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.66...credentials/v1.17.67">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.47.2 to 1.47.3
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8c02c46d4857cbf2c1eb607484e3b94458fe50eb"><code>8c02c46</code></a> Release 2023-12-06</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/58dc75bfb38bf8b8d746e386240a58d857342491"><code>58dc75b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/05f28b1a792aa3c00c08cacffaa896107a69df09"><code>05f28b1</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a5c0a503ef7007945e0016017dc8dad56ed696e6"><code>a5c0a50</code></a> Restore pre-SRA optionalAuth/no-auth behaviors (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2410">#2410</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9805a196767e31076150a0e5ff38e2356a93e840"><code>9805a19</code></a> Release 2023-12-05</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/26cb6b4486d9e3b8a7e9f8b3c023b237310302eb"><code>26cb6b4</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5beea61355437cd3871d3c37435c5adbb30dc431"><code>5beea61</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/baa4447e926a969a69a8faba52c0acf46c827ef3"><code>baa4447</code></a> Smithy upgrade (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2407">#2407</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b8307d199c3c49b8ec6bde395f07d9de0294fe89"><code>b8307d1</code></a> Release 2023-12-04</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/14590dd5a7419101b6b4d321be04c69c72115974"><code>14590dd</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.47.2...service/s3/v1.47.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.79.1 to 1.79.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2a0c73e76f5f06579a2cb24239ca054d60ced4c2"><code>2a0c73e</code></a> Release 2025-04-10</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b7b05ea6b196e40720c2620dc19fbd57945b11b3"><code>b7b05ea</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2a62cee847726f5a607d1801bebe9e1e77f4f400"><code>2a62cee</code></a> Update partitions file</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6af33138b18ea9a336f87e180fc9fdc286b1117e"><code>6af3313</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/915cbf3d7c2f9bc5dcd9a2a1ab02521f01c2f48c"><code>915cbf3</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ac82667a0ef9b9c97a9559810eb106d62bb1cd1e"><code>ac82667</code></a> Release 2025-04-09</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7db3afa50ca2f10dba50ffcb103f453010a2fdb8"><code>7db3afa</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f56d4c1e532d38245bbb638c73b669c536b3263e"><code>f56d4c1</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bf17208deff95cb9703ad5febb75f59af1da7fa9"><code>bf17208</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f529add9a2cd0d97281fd81f711c620c1e95cfb8"><code>f529add</code></a> Release 2025-04-08</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.79.1...service/s3/v1.79.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.31.1 to 0.32.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.32.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.32.0.</p>
<h3>Breaking Changes</h3>
<ul>
<li>Bump the minimum Go version to 1.22. The supported versions are 1.22, 1.23 and 1.24. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/967">#967</a>)</li>
<li>Setting any values on <code>span.Extra</code> has no effect anymore. Use <code>SetData(name string, value interface{})</code> instead. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/864">#864</a>)</li>
</ul>
<h3>Features</h3>
<ul>
<li>Add a <code>MockTransport</code> and <code>MockScope</code>. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/972">#972</a>)</li>
</ul>
<h3>Bug Fixes</h3>
<ul>
<li>Fix writing <code>*http.Request</code> in the Logrus JSONFormatter. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/955">#955</a>)</li>
</ul>
<h3>Misc</h3>
<ul>
<li>Transaction <code>data</code> attributes are now seralized as trace context data attributes, allowing you to query these attributes in the <a href="https://docs.sentry.io/product/explore/traces/">Trace Explorer</a>.</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.32.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.32.0.</p>
<h3>Breaking Changes</h3>
<ul>
<li>Bump the minimum Go version to 1.22. The supported versions are 1.22, 1.23 and 1.24. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/967">#967</a>)</li>
<li>Setting any values on <code>span.Extra</code> has no effect anymore. Use <code>SetData(name string, value interface{})</code> instead. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/864">#864</a>)</li>
</ul>
<h3>Features</h3>
<ul>
<li>Add a <code>MockTransport</code> and <code>MockScope</code>. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/972">#972</a>)</li>
</ul>
<h3>Bug Fixes</h3>
<ul>
<li>Fix writing <code>*http.Request</code> in the Logrus JSONFormatter. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/955">#955</a>)</li>
</ul>
<h3>Misc</h3>
<ul>
<li>Transaction <code>data</code> attributes are now seralized as trace context data attributes, allowing you to query these attributes in the <a href="https://docs.sentry.io/product/explore/traces/">Trace Explorer</a>.</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/530f74d352a3533855e84572c8d0f4e790ae674b"><code>530f74d</code></a> release: 0.32.0</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/e068944300b5225ae481b0bd8f57bfd8e6441bae"><code>e068944</code></a> Prepare 0.32.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/992">#992</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/32c062f0707339f391efc62a05e11e5e3f0b0780"><code>32c062f</code></a> Add <code>transaction.data</code> to <code>contexts.trace.data</code> (<a href="https://redirect.github.com/getsentry/sentry-go/issues/864">#864</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/6019770393710ba5afe949c459289635e73a55b1"><code>6019770</code></a> Fix lint issues (<a href="https://redirect.github.com/getsentry/sentry-go/issues/981">#981</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/cfbfc6b38a42c0254ff00e0f8244d04f785c0f9c"><code>cfbfc6b</code></a> Expose MockTransport, MockScope in root sentry package (<a href="https://redirect.github.com/getsentry/sentry-go/issues/972">#972</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/a4e0ea8a065445c29d25effb3d93e16170bb2931"><code>a4e0ea8</code></a> Remove github token in lint (<a href="https://redirect.github.com/getsentry/sentry-go/issues/982">#982</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/031ec47dc927e3fb737fa818285ccf3f46b47c7f"><code>031ec47</code></a> build(deps): bump golangci/golangci-lint-action from 6.2.0 to 6.5.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/975">#975</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/bdbb6be4a0bae2f225b22b5dca9f1c7af53c06f3"><code>bdbb6be</code></a> build(deps): bump codecov/codecov-action from 5.3.1 to 5.4.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/974">#974</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/3d8d0e14fd1ddff49af13cff94931114dec68c67"><code>3d8d0e1</code></a> build(deps): bump actions/create-github-app-token from 1.11.2 to 1.11.5 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/973">#973</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/e56ac3010f8c48bbb32839ff8a7e751f420c361f"><code>e56ac30</code></a> build(deps): bump codecov/codecov-action from 5.1.2 to 5.3.1 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/962">#962</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/getsentry/sentry-go/compare/v0.31.1...v0.32.0">compare view</a></li>
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

## Reviews

### Review by @rverdile - Approved on April 14, 2025 at 05:00 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1080*
