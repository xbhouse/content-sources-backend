---
type: pull_request
number: 2136
title: "Update module github.com/lib/pq to v1.12.0"
state: merged
author: red-hat-konflux
created: 2026-03-30T01:26:24Z
updated: 2026-03-30T05:13:53Z
closed: 2026-03-30T01:26:42Z
merged: 2026-03-30T01:26:42Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-lib-pq-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2136
---

# Pull Request #2136: Update module github.com/lib/pq to v1.12.0

**Author**: @red-hat-konflux
**Created**: March 30, 2026 at 01:26 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-lib-pq-1.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/lib/pq](https://redirect.github.com/lib/pq) | `v1.11.2` -> `v1.12.0` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2flib%2fpq/v1.12.0?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2flib%2fpq/v1.11.2/v1.12.0?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>lib/pq (github.com/lib/pq)</summary>

### [`v1.12.0`](https://redirect.github.com/lib/pq/blob/HEAD/CHANGELOG.md#v1120-2026-03-18)

[Compare Source](https://redirect.github.com/lib/pq/compare/v1.11.2...v1.12.0)

- The next release may change the default sslmode from `require` to `prefer`.
  See [#&#8203;1271] for details.

- `CopyIn()` and `CopyInToSchema()` have been marked as deprecated. These are
  simple query builders and not needed for `COPY [..] FROM STDIN` support (which
  is *not* deprecated). ([#&#8203;1279])

  ```
  // Old
  tx.Prepare(CopyIn("temp", "num", "text", "blob", "nothing"))

  // Replacement
  tx.Prepare(`copy temp (num, text, blob, nothing) from stdin`)
  ```

##### Features

- Support protocol 3.2, and the `min_protocol_version` and
  `max_protocol_version` DSN parameters ([#&#8203;1258]).

- Support `sslmode=prefer` and `sslmode=allow` ([#&#8203;1270]).

- Support `ssl_min_protocol_version` and `ssl_max_protocol_version` ([#&#8203;1277]).

- Support connection service file to load connection details ([#&#8203;1285]).

- Support `sslrootcert=system` and use `~/.postgresql/root.crt` as the default
  value of sslrootcert ([#&#8203;1280], [#&#8203;1281]).

- Add a new `pqerror` package with PostgreSQL error codes ([#&#8203;1275]).

  For example, to test if an error is a UNIQUE constraint violation:

  ```
  if pqErr, ok := errors.AsType[*pq.Error](err); ok && pqErr.Code == pqerror.UniqueViolation {
      log.Fatalf("email %q already exsts", email)
  }
  ```

  To make this a bit more convenient, it also adds a `pq.As()` function:

  ```
  pqErr := pq.As(err, pqerror.UniqueViolation)
  if pqErr != nil {
      log.Fatalf("email %q already exsts", email)
  }
  ```

##### Fixes

- Fix SSL key permission check to allow modes stricter than [0600/0640#1265](https://redirect.github.com/0600/0640/issues/1265) ([#&#8203;1265]).

- Fix Hstore to work with binary parameters ([#&#8203;1278]).

- Clearer error when starting a new query while pq is still processing another
  query ([#&#8203;1272]).

- Send intermediate CAs with client certificates, so they can be signed by an
  intermediate CA ([#&#8203;1267]).

- Use `time.UTC` for UTC aliases such as `Etc/UTC` ([#&#8203;1282]).

[#&#8203;1258]: https://redirect.github.com/lib/pq/pull/1258

[#&#8203;1265]: https://redirect.github.com/lib/pq/pull/1265

[#&#8203;1267]: https://redirect.github.com/lib/pq/pull/1267

[#&#8203;1270]: https://redirect.github.com/lib/pq/pull/1270

[#&#8203;1271]: https://redirect.github.com/lib/pq/pull/1271

[#&#8203;1272]: https://redirect.github.com/lib/pq/pull/1272

[#&#8203;1275]: https://redirect.github.com/lib/pq/pull/1275

[#&#8203;1277]: https://redirect.github.com/lib/pq/pull/1277

[#&#8203;1278]: https://redirect.github.com/lib/pq/pull/1278

[#&#8203;1279]: https://redirect.github.com/lib/pq/pull/1279

[#&#8203;1280]: https://redirect.github.com/lib/pq/pull/1280

[#&#8203;1281]: https://redirect.github.com/lib/pq/pull/1281

[#&#8203;1282]: https://redirect.github.com/lib/pq/pull/1282

[#&#8203;1283]: https://redirect.github.com/lib/pq/pull/1283

[#&#8203;1285]: https://redirect.github.com/lib/pq/pull/1285

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

## Summary by Sourcery

Build:
- Bump github.com/lib/pq from v1.11.2 to v1.12.0 in go.mod and go.sum.

---

## Discussion

### Comment by @sourcery-ai on March 30, 2026 at 01:26 AM UTC

<!-- Generated by sourcery-ai[bot]: start review_guide -->

<details>
<summary>Reviewer's guide (collapsed on small PRs)</summary>

## Reviewer's Guide

Bumps the PostgreSQL driver dependency github.com/lib/pq from v1.11.2 to v1.12.0 in go.mod (and go.sum), pulling in upstream feature, deprecation, and bugfix changes while leaving application code untouched.

### File-Level Changes

| Change | Details | Files |
| ------ | ------- | ----- |
| Update the PostgreSQL driver dependency github.com/lib/pq from v1.11.2 to v1.12.0. | <ul><li>Change the required version of github.com/lib/pq in go.mod to v1.12.0.</li><li>Refresh go.sum entries to match the new github.com/lib/pq version and its checksums.</li></ul> | `go.mod`<br/>`go.sum` |

</details>

---

<details>
<summary>Tips and commands</summary>

#### Interacting with Sourcery

- **Trigger a new review:** Comment `@sourcery-ai review` on the pull request.
- **Continue discussions:** Reply directly to Sourcery's review comments.
- **Generate a GitHub issue from a review comment:** Ask Sourcery to create an
  issue from a review comment by replying to it. You can also reply to a
  review comment with `@sourcery-ai issue` to create an issue from it.
- **Generate a pull request title:** Write `@sourcery-ai` anywhere in the pull
  request title to generate a title at any time. You can also comment
  `@sourcery-ai title` on the pull request to (re-)generate the title at any time.
- **Generate a pull request summary:** Write `@sourcery-ai summary` anywhere in
  the pull request body to generate a PR summary at any time exactly where you
  want it. You can also comment `@sourcery-ai summary` on the pull request to
  (re-)generate the summary at any time.
- **Generate reviewer's guide:** Comment `@sourcery-ai guide` on the pull
  request to (re-)generate the reviewer's guide at any time.
- **Resolve all Sourcery comments:** Comment `@sourcery-ai resolve` on the
  pull request to resolve all Sourcery comments. Useful if you've already
  addressed all the comments and don't want to see them anymore.
- **Dismiss all Sourcery reviews:** Comment `@sourcery-ai dismiss` on the pull
  request to dismiss all existing Sourcery reviews. Especially useful if you
  want to start fresh with a new review - don't forget to comment
  `@sourcery-ai review` to trigger a new review!

#### Customizing Your Experience

Access your [dashboard](https://app.sourcery.ai) to:
- Enable or disable review features such as the Sourcery-generated pull request
  summary, the reviewer's guide, and others.
- Change the review language.
- Add, remove or edit custom review instructions.
- Adjust other review settings.

#### Getting Help

- [Contact our support team](mailto:support@sourcery.ai) for questions or feedback.
- Visit our [documentation](https://docs.sourcery.ai) for detailed guides and information.
- Keep in touch with the Sourcery team by following us on [X/Twitter](https://x.com/SourceryAI), [LinkedIn](https://www.linkedin.com/company/sourcery-ai/) or [GitHub](https://github.com/sourcery-ai).

</details>

<!-- Generated by sourcery-ai[bot]: end review_guide -->

### Comment by @github-actions on March 30, 2026 at 01:26 AM UTC

<!-- sc-environment-impact-check -->
## SC Environment Impact Assessment

**Overall Impact:** ⚪ **NONE**

No SC Environment-specific impacts detected in this PR.

<details>
<summary>What was checked</summary>

This PR was automatically scanned for:
- Database migrations
- ClowdApp configuration changes
- Kessel integration changes
- AWS service integrations (S3, RDS, ElastiCache)
- Kafka topic changes
- Secrets management changes
- External dependencies
</details>

### Comment by @codecov-commenter on March 30, 2026 at 01:32 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2136?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.18%. Comparing base ([`c4888aa`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/c4888aa8434a265eb6e68e683ee57a42a5193cd9?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`f50e88c`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/f50e88c529ffe9ee6dae8d7da1fa8ba16d3c8546?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 7 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2136   +/-   ##
=======================================
  Coverage   59.18%   59.18%           
=======================================
  Files         134      134           
  Lines        8624     8624           
=======================================
  Hits         5104     5104           
  Misses       2985     2985           
  Partials      535      535           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2136/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2136/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.18% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2136?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @sourcery-ai - Commented on March 30, 2026 at 01:26 AM UTC

Hey - I've left some high level feedback:

- Given the upstream note that the next lib/pq release may change the default `sslmode` from `require` to `prefer`, it would be good to scan existing DSNs and explicitly set `sslmode` where you rely on current defaults to avoid subtle behavioral changes in a future bump.
- Since `CopyIn`/`CopyInToSchema` are now deprecated, check whether this codebase uses them and, if so, plan a follow-up PR to migrate those call sites to raw `COPY ... FROM STDIN` statements as suggested in the changelog.

<details>
<summary>Prompt for AI Agents</summary>

~~~markdown
Please address the comments from this code review:

## Overall Comments
- Given the upstream note that the next lib/pq release may change the default `sslmode` from `require` to `prefer`, it would be good to scan existing DSNs and explicitly set `sslmode` where you rely on current defaults to avoid subtle behavioral changes in a future bump.
- Since `CopyIn`/`CopyInToSchema` are now deprecated, check whether this codebase uses them and, if so, plan a follow-up PR to migrate those call sites to raw `COPY ... FROM STDIN` statements as suggested in the changelog.
~~~

</details>

***

<details>
<summary>Sourcery is free for open source - if you like our reviews please consider sharing them ✨</summary>

- [X](https://twitter.com/intent/tweet?text=I%20just%20got%20an%20instant%20code%20review%20from%20%40SourceryAI%2C%20and%20it%20was%20brilliant%21%20It%27s%20free%20for%20open%20source%20and%20has%20a%20free%20trial%20for%20private%20code.%20Check%20it%20out%20https%3A//sourcery.ai)
- [Mastodon](https://mastodon.social/share?text=I%20just%20got%20an%20instant%20code%20review%20from%20%40SourceryAI%2C%20and%20it%20was%20brilliant%21%20It%27s%20free%20for%20open%20source%20and%20has%20a%20free%20trial%20for%20private%20code.%20Check%20it%20out%20https%3A//sourcery.ai)
- [LinkedIn](https://www.linkedin.com/sharing/share-offsite/?url=https://sourcery.ai)
- [Facebook](https://www.facebook.com/sharer/sharer.php?u=https://sourcery.ai)

</details>

<sub>
Help me be more useful! Please click 👍 or 👎 on each comment and I'll use the feedback to improve your reviews.
</sub>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2136*
