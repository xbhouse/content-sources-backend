---
type: pull_request
number: 1423
title: "Build: Bump github.com/buger/jsonparser from 1.1.1 to 1.1.2"
state: merged
author: dependabot
created: 2026-03-23T08:27:44Z
updated: 2026-03-23T09:06:51Z
closed: 2026-03-23T09:06:49Z
merged: 2026-03-23T09:06:49Z
base_branch: main
head_branch: dependabot/go_modules/github.com/buger/jsonparser-1.1.2
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1423
---

# Pull Request #1423: Build: Bump github.com/buger/jsonparser from 1.1.1 to 1.1.2

**Author**: @dependabot
**Created**: March 23, 2026 at 08:27 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/github.com/buger/jsonparser-1.1.2`

## Description

Bumps [github.com/buger/jsonparser](https://github.com/buger/jsonparser) from 1.1.1 to 1.1.2.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/buger/jsonparser/releases">github.com/buger/jsonparser's releases</a>.</em></p>
<blockquote>
<h2>v1.1.2</h2>
<h2>What's Changed</h2>
<ul>
<li>Updated travis to build for 1.13 to 1.15 by <a href="https://github.com/janreggie"><code>@​janreggie</code></a> in <a href="https://redirect.github.com/buger/jsonparser/pull/225">buger/jsonparser#225</a></li>
<li>
<ul>
<li>eliminate 2 allocations in EachKey() by <a href="https://github.com/Villenny"><code>@​Villenny</code></a> in <a href="https://redirect.github.com/buger/jsonparser/pull/223">buger/jsonparser#223</a></li>
</ul>
</li>
<li>fix issue <a href="https://redirect.github.com/buger/jsonparser/issues/150">#150</a> (in deleting case) by <a href="https://github.com/daria-kay"><code>@​daria-kay</code></a> in <a href="https://redirect.github.com/buger/jsonparser/pull/226">buger/jsonparser#226</a></li>
<li>fixing the oss-fuzz issue by <a href="https://github.com/daria-kay"><code>@​daria-kay</code></a> in <a href="https://redirect.github.com/buger/jsonparser/pull/227">buger/jsonparser#227</a></li>
<li>Fix parseInt overflow check false negative by <a href="https://github.com/carsonip"><code>@​carsonip</code></a> in <a href="https://redirect.github.com/buger/jsonparser/pull/231">buger/jsonparser#231</a></li>
<li>Added bespoke error for null cases by <a href="https://github.com/jonomacd"><code>@​jonomacd</code></a> in <a href="https://redirect.github.com/buger/jsonparser/pull/228">buger/jsonparser#228</a></li>
<li>Fuzzing: Add CIFuzz by <a href="https://github.com/AdamKorcz"><code>@​AdamKorcz</code></a> in <a href="https://redirect.github.com/buger/jsonparser/pull/239">buger/jsonparser#239</a></li>
<li>Added latest versions of go to tests by <a href="https://github.com/moredure"><code>@​moredure</code></a> in <a href="https://redirect.github.com/buger/jsonparser/pull/244">buger/jsonparser#244</a></li>
<li>fix EachKey pIdxFlags allocation by <a href="https://github.com/unxcepted"><code>@​unxcepted</code></a> in <a href="https://redirect.github.com/buger/jsonparser/pull/241">buger/jsonparser#241</a></li>
<li>fix: prevent panic on negative slice index in Delete with malformed JSON (GO-2026-4514) by <a href="https://github.com/dbarrosop"><code>@​dbarrosop</code></a> in <a href="https://redirect.github.com/buger/jsonparser/pull/276">buger/jsonparser#276</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/janreggie"><code>@​janreggie</code></a> made their first contribution in <a href="https://redirect.github.com/buger/jsonparser/pull/225">buger/jsonparser#225</a></li>
<li><a href="https://github.com/Villenny"><code>@​Villenny</code></a> made their first contribution in <a href="https://redirect.github.com/buger/jsonparser/pull/223">buger/jsonparser#223</a></li>
<li><a href="https://github.com/daria-kay"><code>@​daria-kay</code></a> made their first contribution in <a href="https://redirect.github.com/buger/jsonparser/pull/226">buger/jsonparser#226</a></li>
<li><a href="https://github.com/carsonip"><code>@​carsonip</code></a> made their first contribution in <a href="https://redirect.github.com/buger/jsonparser/pull/231">buger/jsonparser#231</a></li>
<li><a href="https://github.com/jonomacd"><code>@​jonomacd</code></a> made their first contribution in <a href="https://redirect.github.com/buger/jsonparser/pull/228">buger/jsonparser#228</a></li>
<li><a href="https://github.com/moredure"><code>@​moredure</code></a> made their first contribution in <a href="https://redirect.github.com/buger/jsonparser/pull/244">buger/jsonparser#244</a></li>
<li><a href="https://github.com/unxcepted"><code>@​unxcepted</code></a> made their first contribution in <a href="https://redirect.github.com/buger/jsonparser/pull/241">buger/jsonparser#241</a></li>
<li><a href="https://github.com/dbarrosop"><code>@​dbarrosop</code></a> made their first contribution in <a href="https://redirect.github.com/buger/jsonparser/pull/276">buger/jsonparser#276</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/buger/jsonparser/compare/v1.1.1...v1.1.2">https://github.com/buger/jsonparser/compare/v1.1.1...v1.1.2</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/buger/jsonparser/commit/a69e7e01cd4ad67bdfd3ac2c080b9212af16f4b0"><code>a69e7e0</code></a> Merge pull request <a href="https://redirect.github.com/buger/jsonparser/issues/276">#276</a> from dbarrosop/master</li>
<li><a href="https://github.com/buger/jsonparser/commit/d3eacc0bab779d6cf98221f5268828fff287876e"><code>d3eacc0</code></a> fix: prevent panic on negative slice index in Delete with malformed JSON (GO-...</li>
<li><a href="https://github.com/buger/jsonparser/commit/61b32cfdfa0f5d368ef7c7daef28ce12d538740f"><code>61b32cf</code></a> Merge pull request <a href="https://redirect.github.com/buger/jsonparser/issues/241">#241</a> from unxcepted/master</li>
<li><a href="https://github.com/buger/jsonparser/commit/2181e8398f18397c9cacbaea9889314bb585e868"><code>2181e83</code></a> Merge pull request <a href="https://redirect.github.com/buger/jsonparser/issues/244">#244</a> from ScaleChamp/patch-2</li>
<li><a href="https://github.com/buger/jsonparser/commit/1510b5194182fc2fb898f28cdbceb42fd7258bfa"><code>1510b51</code></a> Added latest versions of go to tests</li>
<li><a href="https://github.com/buger/jsonparser/commit/6fc2e488ed3cc4f1f1debec3b0c70715bd7be6fd"><code>6fc2e48</code></a> fix: eachkey allocation</li>
<li><a href="https://github.com/buger/jsonparser/commit/a6f867eb7787e4ec54536b77b5d628ddf5c4f73d"><code>a6f867e</code></a> Merge pull request <a href="https://redirect.github.com/buger/jsonparser/issues/239">#239</a> from AdamKorcz/cifuzz1</li>
<li><a href="https://github.com/buger/jsonparser/commit/cbc01fdbbe131706e89eeaaf0cd917760d8d3949"><code>cbc01fd</code></a> Fuzzing: Add CIFuzz</li>
<li><a href="https://github.com/buger/jsonparser/commit/dc92d6932a1272b4d8f485f798a88c3a75106256"><code>dc92d69</code></a> Merge pull request <a href="https://redirect.github.com/buger/jsonparser/issues/228">#228</a> from jonomacd/null-handling</li>
<li><a href="https://github.com/buger/jsonparser/commit/2d9d6343e8621ddc18c70749663f74bc584c0de4"><code>2d9d634</code></a> Merge pull request <a href="https://redirect.github.com/buger/jsonparser/issues/231">#231</a> from carsonip/fix-parseint-overflow-check</li>
<li>Additional commits viewable in <a href="https://github.com/buger/jsonparser/compare/v1.1.1...v1.1.2">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=github.com/buger/jsonparser&package-manager=go_modules&previous-version=1.1.1&new-version=1.1.2)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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
- `@dependabot show <dependency name> ignore conditions` will show all of the ignore conditions of the specified dependency
- `@dependabot ignore this major version` will close this PR and stop Dependabot creating any more for this major version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this minor version` will close this PR and stop Dependabot creating any more for this minor version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this dependency` will close this PR and stop Dependabot creating any more for this dependency (unless you reopen the PR or upgrade to it yourself)
You can disable automated security fix PRs for this repo from the [Security Alerts page](https://github.com/content-services/content-sources-backend/network/alerts).

</details>

---

## Discussion

### Comment by @swadeley on March 23, 2026 at 08:40 AM UTC

/retest

---

## Reviews

### Review by @swadeley - Approved on March 23, 2026 at 08:34 AM UTC

ACK

### Review by @swadeley - Approved on March 23, 2026 at 08:35 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1423*
