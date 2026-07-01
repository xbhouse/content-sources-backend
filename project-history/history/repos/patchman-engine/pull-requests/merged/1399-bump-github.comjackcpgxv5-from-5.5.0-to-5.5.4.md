---
type: pull_request
number: 1399
title: "Bump github.com/jackc/pgx/v5 from 5.5.0 to 5.5.4"
state: merged
author: dependabot
created: 2024-03-14T22:17:39Z
updated: 2024-03-15T12:21:24Z
closed: 2024-03-15T12:21:16Z
merged: 2024-03-15T12:21:16Z
base_branch: master
head_branch: dependabot/go_modules/github.com/jackc/pgx/v5-5.5.4
labels: ["dependencies"]
url: https://github.com/RedHatInsights/patchman-engine/pull/1399
---

# Pull Request #1399: Bump github.com/jackc/pgx/v5 from 5.5.0 to 5.5.4

**Author**: @dependabot
**Created**: March 14, 2024 at 10:17 PM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `master` ← **Head**: `dependabot/go_modules/github.com/jackc/pgx/v5-5.5.4`

## Description

Bumps [github.com/jackc/pgx/v5](https://github.com/jackc/pgx) from 5.5.0 to 5.5.4.
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jackc/pgx/blob/master/CHANGELOG.md">github.com/jackc/pgx/v5's changelog</a>.</em></p>
<blockquote>
<h1>5.5.4 (March 4, 2024)</h1>
<p>Fix CVE-2024-27304</p>
<p>SQL injection can occur if an attacker can cause a single query or bind message to exceed 4 GB in size. An integer
overflow in the calculated message size can cause the one large message to be sent as multiple messages under the
attacker's control.</p>
<p>Thanks to Paul Gerste for reporting this issue.</p>
<ul>
<li>Fix behavior of CollectRows to return empty slice if Rows are empty (Felix)</li>
<li>Fix simple protocol encoding of json.RawMessage</li>
<li>Fix *Pipeline.getResults should close pipeline on error</li>
<li>Fix panic in TryFindUnderlyingTypeScanPlan (David Kurman)</li>
<li>Fix deallocation of invalidated cached statements in a transaction</li>
<li>Handle invalid sslkey file</li>
<li>Fix scan float4 into sql.Scanner</li>
<li>Fix pgtype.Bits not making copy of data from read buffer. This would cause the data to be corrupted by future reads.</li>
</ul>
<h1>5.5.3 (February 3, 2024)</h1>
<ul>
<li>Fix: prepared statement already exists</li>
<li>Improve CopyFrom auto-conversion of text-ish values</li>
<li>Add ltree type support (Florent Viel)</li>
<li>Make some properties of Batch and QueuedQuery public (Pavlo Golub)</li>
<li>Add AppendRows function (Edoardo Spadolini)</li>
<li>Optimize convert UUID [16]byte to string (Kirill Malikov)</li>
<li>Fix: LargeObject Read and Write of more than ~1GB at a time (Mitar)</li>
</ul>
<h1>5.5.2 (January 13, 2024)</h1>
<ul>
<li>Allow NamedArgs to start with underscore</li>
<li>pgproto3: Maximum message body length support (jeremy.spriet)</li>
<li>Upgrade golang.org/x/crypto to v0.17.0</li>
<li>Add snake_case support to RowToStructByName (Tikhon Fedulov)</li>
<li>Fix: update description cache after exec prepare (James Hartig)</li>
<li>Fix: pipeline checks if it is closed (James Hartig and Ryan Fowler)</li>
<li>Fix: normalize timeout / context errors during TLS startup (Samuel Stauffer)</li>
<li>Add OnPgError for easier centralized error handling (James Hartig)</li>
</ul>
<h1>5.5.1 (December 9, 2023)</h1>
<ul>
<li>Add CopyFromFunc helper function. (robford)</li>
<li>Add PgConn.Deallocate method that uses PostgreSQL protocol Close message.</li>
<li>pgx uses new PgConn.Deallocate method. This allows deallocating statements to work in a failed transaction. This fixes a case where the prepared statement map could become invalid.</li>
<li>Fix: Prefer driver.Valuer over json.Marshaler for json fields. (Jacopo)</li>
<li>Fix: simple protocol SQL sanitizer previously panicked if an invalid $0 placeholder was used. This now returns an error instead. (maksymnevajdev)</li>
<li>Add pgtype.Numeric.ScanScientific (Eshton Robateau)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jackc/pgx/commit/da6f2c98f2664b215b40b1606551fdfcc7f3ea5c"><code>da6f2c9</code></a> Update changelog</li>
<li><a href="https://github.com/jackc/pgx/commit/c543134753a0c5d22881c12404025724cb05ffd8"><code>c543134</code></a> SQL sanitizer wraps arguments in parentheses</li>
<li><a href="https://github.com/jackc/pgx/commit/20344dfae83e672eff73a03324398543a1d87f43"><code>20344df</code></a> Check for overflow on uint16 sizes in pgproto3</li>
<li><a href="https://github.com/jackc/pgx/commit/adbb38f298c76e283ffc7c7a3f571036fea47fd4"><code>adbb38f</code></a> Do not allow protocol messages larger than ~1GB</li>
<li><a href="https://github.com/jackc/pgx/commit/c1b0a01ca75ac9eb3a7dbc1396f583ab5dbf9557"><code>c1b0a01</code></a> Fix behavior of CollectRows to return empty slice if Rows are empty</li>
<li><a href="https://github.com/jackc/pgx/commit/88dfc22ae4aa031783cda90841d5358edd85ff2c"><code>88dfc22</code></a> Fix simple protocol encoding of json.RawMessage</li>
<li><a href="https://github.com/jackc/pgx/commit/2e84dccaf57b4fa803149884f2149dfa9e923593"><code>2e84dcc</code></a> *Pipeline.getResults should close pipeline on error</li>
<li><a href="https://github.com/jackc/pgx/commit/d149d3fe5c50d1d98bd6265d3c928519ba4b3f4b"><code>d149d3f</code></a> Fix panic in TryFindUnderlyingTypeScanPlan</li>
<li><a href="https://github.com/jackc/pgx/commit/046f497efb4e92caa9575a0e9c351e4906af14c6"><code>046f497</code></a> deallocateInvalidatedCachedStatements now runs in transactions</li>
<li><a href="https://github.com/jackc/pgx/commit/8896bd697781ed4aee392daa90b90cde142319fe"><code>8896bd6</code></a> Handle invalid sslkey file</li>
<li>Additional commits viewable in <a href="https://github.com/jackc/pgx/compare/v5.5.0...v5.5.4">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=github.com/jackc/pgx/v5&package-manager=go_modules&previous-version=5.5.0&new-version=5.5.4)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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
You can disable automated security fix PRs for this repo from the [Security Alerts page](https://github.com/RedHatInsights/patchman-engine/network/alerts).

</details>

---

## Discussion

### Comment by @jira-linking on March 14, 2024 at 10:17 PM UTC

Commits missing Jira IDs:
b708d8c0ee78327fc2012ab3c4dff9b762b05de5


### Comment by @codecov-commenter on March 14, 2024 at 10:22 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1399?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 63.71%. Comparing base [(`04e5cd4`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/04e5cd4377ec6c5f524e207cb0f970e10c512835?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`b708d8c`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1399?dropdown=coverage&src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 4 commits behind head on master.


<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1399      +/-   ##
==========================================
+ Coverage   63.67%   63.71%   +0.04%     
==========================================
  Files         107      107              
  Lines        6521     6521              
==========================================
+ Hits         4152     4155       +3     
+ Misses       1880     1877       -3     
  Partials      489      489              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1399/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1399/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `63.71% <ø> (+0.04%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1399?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1399*
