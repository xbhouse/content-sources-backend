---
type: pull_request
number: 2038
title: "Update module github.com/lib/pq to v1.11.1"
state: merged
author: red-hat-konflux
created: 2026-02-02T08:48:37Z
updated: 2026-02-02T12:48:17Z
closed: 2026-02-02T09:03:04Z
merged: 2026-02-02T09:03:04Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-lib-pq-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2038
---

# Pull Request #2038: Update module github.com/lib/pq to v1.11.1

**Author**: @red-hat-konflux
**Created**: February 02, 2026 at 08:48 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-lib-pq-1.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/lib/pq](https://redirect.github.com/lib/pq) | `v1.10.9` -> `v1.11.1` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2flib%2fpq/v1.11.1?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2flib%2fpq/v1.10.9/v1.11.1?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>lib/pq (github.com/lib/pq)</summary>

### [`v1.11.1`](https://redirect.github.com/lib/pq/blob/HEAD/CHANGELOG.md#v1111-2025-01-29)

[Compare Source](https://redirect.github.com/lib/pq/compare/v1.11.0...v1.11.1)

This fixes two regressions present in the v1.11.0 release:

- Fix build on 32bit systems, Windows, and Plan 9 ([#&#8203;1253]).

- Named \[]byte types and pointers to \[]byte (e.g. `*[]byte`, `json.RawMessage`)
  would be treated as an array instead of bytea ([#&#8203;1252]).

[#&#8203;1252]: https://redirect.github.com/lib/pq/pull/1252

[#&#8203;1253]: https://redirect.github.com/lib/pq/pull/1253

### [`v1.11.0`](https://redirect.github.com/lib/pq/blob/HEAD/CHANGELOG.md#v1110-2025-01-28)

[Compare Source](https://redirect.github.com/lib/pq/compare/v1.10.9...v1.11.0)

This version of pq requires Go 1.21 or newer.

pq now supports only maintained PostgreSQL releases, which is PostgreSQL 14 and
newer. Previously PostgreSQL 8.4 and newer were supported.

##### Features

- The `pq.Error.Error()` text  includes the position of the error (if reported
  by PostgreSQL) and SQLSTATE code ([#&#8203;1219], [#&#8203;1224]):

  ```
  pq: column "columndoesntexist" does not exist at column 8 (42703)
  pq: syntax error at or near ")" at position 2:71 (42601)
  ```

- The `pq.Error.ErrorWithDetail()` method prints a more detailed multiline
  message, with the Detail, Hint, and error position (if any) ([#&#8203;1219]):

  ```
  ERROR:   syntax error at or near ")" (42601)
  CONTEXT: line 12, column 1:

       10 |     name           varchar,
       11 |     version        varchar,
       12 | );
            ^
  ```

- Add `Config`, `NewConfig()`, and `NewConnectorConfig()` to supply connection
  details in a more structured way ([#&#8203;1240]).

- Support `hostaddr` and `$PGHOSTADDR` ([#&#8203;1243]).

- Support multiple values in `host`, `port`, and `hostaddr`, which are each
  tried in order, or randomly if `load_balance_hosts=random` is set ([#&#8203;1246]).

- Support `target_session_attrs` connection parameter ([#&#8203;1246]).

- Support [`sslnegotiation`] to use SSL without negotiation ([#&#8203;1180]).

- Allow using a custom `tls.Config`, for example for encrypted keys ([#&#8203;1228]).

- Add `PQGO_DEBUG=1` print the communication with PostgreSQL to stderr, to aid
  in debugging, testing, and bug reports ([#&#8203;1223]).

- Add support for NamedValueChecker interface ([#&#8203;1125], [#&#8203;1238]).

##### Fixes

- Match HOME directory lookup logic with libpq: prefer $HOME over /etc/passwd,
  ignore ENOTDIR errors, and use APPDATA on Windows ([#&#8203;1214]).

- Fix `sslmode=verify-ca` verifying the hostname anyway when connecting to a DNS
  name (rather than IP) ([#&#8203;1226]).

- Correctly detect pre-protocol errors such as the server not being able to fork
  or running out of memory ([#&#8203;1248]).

- Fix build with wasm ([#&#8203;1184]), appengine ([#&#8203;745]), and Plan 9 ([#&#8203;1133]).

- Deprecate and type alias `pq.NullTime` to `sql.NullTime` ([#&#8203;1211]).

- Enforce integer limits of the Postgres wire protocol ([#&#8203;1161]).

- Accept the `passfile` connection parameter to override `PGPASSFILE` ([#&#8203;1129]).

- Fix connecting to socket on Windows systems ([#&#8203;1179]).

- Don't perform a permission check on the .pgpass file on Windows ([#&#8203;595]).

- Warn about incorrect .pgpass permissions ([#&#8203;595]).

- Don't set extra\_float\_digits ([#&#8203;1212]).

- Decode bpchar into a string ([#&#8203;949]).

- Fix panic in Ping() by not requiring CommandComplete or EmptyQueryResponse in
  simpleQuery() ([#&#8203;1234])

- Recognize bit/varbit ([#&#8203;743]) and float types ([#&#8203;1166]) in ColumnTypeScanType().

- Accept `PGGSSLIB` and `PGKRBSRVNAME` environment variables ([#&#8203;1143]).

- Handle ErrorResponse in readReadyForQuery and return proper error ([#&#8203;1136]).

- CopyIn() and CopyInSchema() now work if the list of columns is empty, in which
  case it will copy all columns ([#&#8203;1239]).

- Treat nil \[]byte in query parameters as nil/NULL rather than `""` ([#&#8203;838]).

- Accept multiple authentication methods before checking AuthOk, which improves
  compatibility with PgPool-II ([#&#8203;1188]).

[`sslnegotiation`]: https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNECT-SSLNEGOTIATION

[#&#8203;595]: https://redirect.github.com/lib/pq/pull/595

[#&#8203;745]: https://redirect.github.com/lib/pq/pull/745

[#&#8203;743]: https://redirect.github.com/lib/pq/pull/743

[#&#8203;838]: https://redirect.github.com/lib/pq/pull/838

[#&#8203;949]: https://redirect.github.com/lib/pq/pull/949

[#&#8203;1125]: https://redirect.github.com/lib/pq/pull/1125

[#&#8203;1129]: https://redirect.github.com/lib/pq/pull/1129

[#&#8203;1133]: https://redirect.github.com/lib/pq/pull/1133

[#&#8203;1136]: https://redirect.github.com/lib/pq/pull/1136

[#&#8203;1143]: https://redirect.github.com/lib/pq/pull/1143

[#&#8203;1161]: https://redirect.github.com/lib/pq/pull/1161

[#&#8203;1166]: https://redirect.github.com/lib/pq/pull/1166

[#&#8203;1179]: https://redirect.github.com/lib/pq/pull/1179

[#&#8203;1180]: https://redirect.github.com/lib/pq/pull/1180

[#&#8203;1184]: https://redirect.github.com/lib/pq/pull/1184

[#&#8203;1188]: https://redirect.github.com/lib/pq/pull/1188

[#&#8203;1211]: https://redirect.github.com/lib/pq/pull/1211

[#&#8203;1212]: https://redirect.github.com/lib/pq/pull/1212

[#&#8203;1214]: https://redirect.github.com/lib/pq/pull/1214

[#&#8203;1219]: https://redirect.github.com/lib/pq/pull/1219

[#&#8203;1223]: https://redirect.github.com/lib/pq/pull/1223

[#&#8203;1224]: https://redirect.github.com/lib/pq/pull/1224

[#&#8203;1226]: https://redirect.github.com/lib/pq/pull/1226

[#&#8203;1228]: https://redirect.github.com/lib/pq/pull/1228

[#&#8203;1234]: https://redirect.github.com/lib/pq/pull/1234

[#&#8203;1238]: https://redirect.github.com/lib/pq/pull/1238

[#&#8203;1239]: https://redirect.github.com/lib/pq/pull/1239

[#&#8203;1240]: https://redirect.github.com/lib/pq/pull/1240

[#&#8203;1243]: https://redirect.github.com/lib/pq/pull/1243

[#&#8203;1246]: https://redirect.github.com/lib/pq/pull/1246

[#&#8203;1248]: https://redirect.github.com/lib/pq/pull/1248

</details>

---

### Configuration

📅 **Schedule**: Branch creation - Between 03:00 AM and 10:59 AM, only on Monday ( * 3-10 * * 1 ) in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Enabled.

♻ **Rebasing**: Whenever PR is behind base branch, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

---
### Documentation

Find out how to configure dependency updates in [MintMaker documentation](https://konflux-ci.dev/docs/mintmaker/user/) or see all available configuration options in [Renovate documentation](https://docs.renovatebot.com/configuration-options/).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0Mi4yNi41LXJwbSIsInVwZGF0ZWRJblZlciI6IjQyLjI2LjUtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @jira-linking on February 02, 2026 at 08:48 AM UTC

Commits missing Jira IDs:
6bfc72b850d65575a3b5ab1d65029b5f4ef94888


### Comment by @codecov-commenter on February 02, 2026 at 08:53 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2038?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.25%. Comparing base ([`92592c9`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/92592c95755f9762d74e9f8cc68fac318de73333?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`6bfc72b`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/6bfc72b850d65575a3b5ab1d65029b5f4ef94888?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2038   +/-   ##
=======================================
  Coverage   59.25%   59.25%           
=======================================
  Files         134      134           
  Lines        8615     8615           
=======================================
  Hits         5105     5105           
  Misses       2967     2967           
  Partials      543      543           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2038/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2038/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.25% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2038?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2038*
