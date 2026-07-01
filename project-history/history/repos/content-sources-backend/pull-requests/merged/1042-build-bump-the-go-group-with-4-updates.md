---
type: pull_request
number: 1042
title: "Build: Bump the go group with 4 updates"
state: merged
author: dependabot
created: 2025-03-24T04:46:55Z
updated: 2025-03-24T12:25:32Z
closed: 2025-03-24T12:25:24Z
merged: 2025-03-24T12:25:24Z
base_branch: main
head_branch: dependabot/go_modules/go-3d53e2d7c4
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1042
---

# Pull Request #1042: Build: Bump the go group with 4 updates

**Author**: @dependabot
**Created**: March 24, 2025 at 04:46 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-3d53e2d7c4`

## Description

Bumps the go group with 4 updates: [github.com/getkin/kin-openapi](https://github.com/getkin/kin-openapi), [github.com/rs/zerolog](https://github.com/rs/zerolog), [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) and [github.com/jackc/pgx/v5](https://github.com/jackc/pgx).

Updates `github.com/getkin/kin-openapi` from 0.130.0 to 0.131.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getkin/kin-openapi/releases">github.com/getkin/kin-openapi's releases</a>.</em></p>
<blockquote>
<h2>v0.131.0</h2>
<h2>What's Changed</h2>
<ul>
<li>openapi3filter: de-register ZipFileBodyDecoder and make a few decoders public by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1059">getkin/kin-openapi#1059</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/getkin/kin-openapi/compare/v0.130.0...v0.131.0">https://github.com/getkin/kin-openapi/compare/v0.130.0...v0.131.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getkin/kin-openapi/commit/67f0b233ffc01332f7d993f79490fbea5f4455f1"><code>67f0b23</code></a> openapi3filter: de-register ZipFileBodyDecoder and make a few decoders public...</li>
<li>See full diff in <a href="https://github.com/getkin/kin-openapi/compare/v0.130.0...v0.131.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/rs/zerolog` from 1.33.0 to 1.34.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/rs/zerolog/commit/db9d1bebd9c0810ec04bc7e4199655f85c5b479b"><code>db9d1be</code></a> Update go versions covered by CI</li>
<li><a href="https://github.com/rs/zerolog/commit/5f4b880a01c1a905549ab70ecad797e19f09d9a7"><code>5f4b880</code></a> Delete _config.yml</li>
<li><a href="https://github.com/rs/zerolog/commit/ffb27080ca298f9827a625d247d9c50fdf28cd8f"><code>ffb2708</code></a> Remove CNAME file</li>
<li><a href="https://github.com/rs/zerolog/commit/cc4dde7383252eb0f0b99067dd1d7a2e61ba6ac2"><code>cc4dde7</code></a> Create CONTRIBUTING.md</li>
<li><a href="https://github.com/rs/zerolog/commit/04ea0f4371b35472e5cda5468daeb3ea5c574ffb"><code>04ea0f4</code></a> Implement Close() for zerolog.FilteredLevelWriter (<a href="https://redirect.github.com/rs/zerolog/issues/715">#715</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/039860087ce1099876fb1f03a069a9feaa9cbefe"><code>0398600</code></a> fix: reset condition in burst sampler (<a href="https://redirect.github.com/rs/zerolog/issues/711">#711</a>) (<a href="https://redirect.github.com/rs/zerolog/issues/712">#712</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/1869fa55bea5c09e93a06368c6a8756780dca5f7"><code>1869fa5</code></a> FormatPartValueByName for flexible custom formatting for ConsoleWriter (<a href="https://redirect.github.com/rs/zerolog/issues/541">#541</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/31e7995c5b60af4ba6ae1aa808a026f8cbf3b8cf"><code>31e7995</code></a> remove unnecessary nil checks (<a href="https://redirect.github.com/rs/zerolog/issues/701">#701</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/582f820cf0db2a9532afbb61570bec4e15ff2b1f"><code>582f820</code></a> Get BasicSampler(0), RandomSampler(0), and BurstSampler(0) to behave the same...</li>
<li><a href="https://github.com/rs/zerolog/commit/6abadab4881e4af6d122332f6aef0365507c248a"><code>6abadab</code></a> Bump github.com/rs/xid from 1.5.0 to 1.6.0 (<a href="https://redirect.github.com/rs/zerolog/issues/684">#684</a>)</li>
<li>See full diff in <a href="https://github.com/rs/zerolog/compare/v1.33.0...v1.34.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.47.0 to 1.47.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f3ea5b8aca81cfc644b822ffc9910f2382e474c2"><code>f3ea5b8</code></a> Release 2023-11-30</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bb443747086264230b376f3a965df21acfce6927"><code>bb44374</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1fd796ca4039fe0f3358736643ed1eba0f59b071"><code>1fd796c</code></a> Update SDK's smithy-go dependency to v1.18.1</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5e061f0f21478d50bb8636c8b73b7c0a32b66816"><code>5e061f0</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e6de9ae7e5c9631913c08d37619e35cf3d9628e0"><code>e6de9ae</code></a> Update API model</li>
<li>See full diff in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.47.0...service/s3/v1.47.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/jackc/pgx/v5` from 5.7.2 to 5.7.3
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jackc/pgx/blob/master/CHANGELOG.md">github.com/jackc/pgx/v5's changelog</a>.</em></p>
<blockquote>
<h1>5.7.3 (March 21, 2025)</h1>
<ul>
<li>Expose EmptyAcquireWaitTime in pgxpool.Stat (vamshiaruru32)</li>
<li>Improve SQL sanitizer performance (ninedraft)</li>
<li>Fix Scan confusion with json(b), sql.Scanner, and automatic dereferencing (moukoublen, felix-roehrich)</li>
<li>Fix Values() for xml type always returning nil instead of []byte</li>
<li>Add ability to send Flush message in pipeline mode (zenkovev)</li>
<li>Fix pgtype.Timestamp's JSON behavior to match PostgreSQL (pconstantinou)</li>
<li>Better error messages when scanning structs (logicbomb)</li>
<li>Fix handling of error on batch write (bonnefoa)</li>
<li>Match libpq's connection fallback behavior more closely (felix-roehrich)</li>
<li>Add MinIdleConns to pgxpool (djahandarie)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jackc/pgx/commit/5c1fbf4806cb7fc61f88e0808fb6bdabbac6df3f"><code>5c1fbf4</code></a> Update changelog for v5.7.3</li>
<li><a href="https://github.com/jackc/pgx/commit/05fe5f8b05285cac23190b1d99543e98924f2dee"><code>05fe5f8</code></a> Explain seemingly redundant rows.Close() in CollectOneRow</li>
<li><a href="https://github.com/jackc/pgx/commit/70c9a147a20c3645cd566b539dd681e03c986e1a"><code>70c9a14</code></a> Merge pull request <a href="https://redirect.github.com/jackc/pgx/issues/2279">#2279</a> from djahandarie/min-idle-conns</li>
<li><a href="https://github.com/jackc/pgx/commit/6603ddfbe446491f948edcb31bed97e835eb3b28"><code>6603ddf</code></a> add MinIdleConns</li>
<li><a href="https://github.com/jackc/pgx/commit/70f7cad2226dc12406b105f8bb5be9c62780aaf7"><code>70f7cad</code></a> Add link to <a href="https://github.com/Arlandaren/pgxWrappy">https://github.com/Arlandaren/pgxWrappy</a></li>
<li><a href="https://github.com/jackc/pgx/commit/6bf1b0b1b90399c72fc9cc0da9f538341490af7c"><code>6bf1b0b</code></a> Add database/sql to overview of scanning</li>
<li><a href="https://github.com/jackc/pgx/commit/14bda65a0cb3a66672ed9290a337a1cb7e103d3e"><code>14bda65</code></a> Correct pgtype docs</li>
<li><a href="https://github.com/jackc/pgx/commit/9e3c4fb40fcd8d5c764d24fe1a651319a0d20122"><code>9e3c4fb</code></a> Merge pull request <a href="https://redirect.github.com/jackc/pgx/issues/2257">#2257</a> from felix-roehrich/fr/change-connect-logic</li>
<li><a href="https://github.com/jackc/pgx/commit/05e72a5ab10edd4ebf166624d2f64476791dbd4f"><code>05e72a5</code></a> make connection logic more forgiving</li>
<li><a href="https://github.com/jackc/pgx/commit/47d631e34be7128997a0aa89b75885cc4ad4c82e"><code>47d631e</code></a> Added missed change to v5.7.2 changelog</li>
<li>Additional commits viewable in <a href="https://github.com/jackc/pgx/compare/v5.7.2...v5.7.3">compare view</a></li>
</ul>
</details>
<br />

<details>
<summary>Most Recent Ignore Conditions Applied to This Pull Request</summary>

| Dependency Name | Ignore Conditions |
| --- | --- |
| github.com/getkin/kin-openapi | [>= 0.128.a, < 0.129] |
</details>


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

### Review by @jlsherrill - Approved on March 24, 2025 at 12:25 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1042*
