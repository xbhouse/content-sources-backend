---
type: pull_request
number: 1949
title: "Update module github.com/quic-go/quic-go to v0.57.1"
state: merged
author: red-hat-konflux
created: 2025-11-24T08:44:59Z
updated: 2025-12-01T08:41:49Z
closed: 2025-12-01T04:49:34Z
merged: 2025-12-01T04:49:34Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-quic-go-quic-go-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1949
---

# Pull Request #1949: Update module github.com/quic-go/quic-go to v0.57.1

**Author**: @red-hat-konflux
**Created**: November 24, 2025 at 08:44 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-quic-go-quic-go-0.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/quic-go/quic-go](https://redirect.github.com/quic-go/quic-go) | `v0.56.0` -> `v0.57.1` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fquic-go%2fquic-go/v0.57.1?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fquic-go%2fquic-go/v0.56.0/v0.57.1?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>quic-go/quic-go (github.com/quic-go/quic-go)</summary>

### [`v0.57.1`](https://redirect.github.com/quic-go/quic-go/releases/tag/v0.57.1)

[Compare Source](https://redirect.github.com/quic-go/quic-go/compare/v0.57.0...v0.57.1)

This release resolves a panic during the server handshake when using the upcoming Go 1.26 toolchain, specifically occurring with TLS session tickets disabled ([#&#8203;5462](https://redirect.github.com/quic-go/quic-go/issues/5462)). This issue does not impact builds on Go 1.25 or earlier versions.

### [`v0.57.0`](https://redirect.github.com/quic-go/quic-go/releases/tag/v0.57.0)

[Compare Source](https://redirect.github.com/quic-go/quic-go/compare/v0.56.0...v0.57.0)

This release reworks the HTTP/3 header processing logic:

- Both client and server now send their respective header size constraints using the SETTINGS\_MAX\_FIELD\_SECTION\_SIZE setting: [#&#8203;5431](https://redirect.github.com/quic-go/quic-go/issues/5431)
- For any QPACK-related errors, the correct error code (QPACK\_DECOMPRESSION\_FAILED) is now used: [#&#8203;5439](https://redirect.github.com/quic-go/quic-go/issues/5439)
- QPACK header parsing is now incremental (instead of parsing all headers at once), which is \~5-10% faster and reduces allocations: [#&#8203;5435](https://redirect.github.com/quic-go/quic-go/issues/5435) (and [quic-go/qpack#67](https://redirect.github.com/quic-go/qpack/pull/67))
- The server now sends a 431 status code (Request Header Fields Too Large) when encountering HTTP header fields exceeding the size constraint: [#&#8203;5452](https://redirect.github.com/quic-go/quic-go/issues/5452)

 

#### Breaking Changes

- http3: `Transport.MaxResponseBytes` is now an `int` (before: `int64`): [#&#8203;5433](https://redirect.github.com/quic-go/quic-go/issues/5433)
   

#### Notable Fixes

- qlogwriter: fix storing of event schemas (this prevented qlog event logging from working for HTTP/3): [#&#8203;5430](https://redirect.github.com/quic-go/quic-go/issues/5430)
- http3: errors sending the request are now ignored, instead, the response from the server is read (thereby allowing the client to read the status code, for example): [#&#8203;5432](https://redirect.github.com/quic-go/quic-go/issues/5432)

#### What's Changed

- build(deps): bump golangci/golangci-lint-action from 8 to 9 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;5426](https://redirect.github.com/quic-go/quic-go/pull/5426)
- qlogwriter: fix storing of event schemas by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5430](https://redirect.github.com/quic-go/quic-go/pull/5430)
- http3: send SETTINGS\_MAX\_FIELD\_SECTION\_SIZE in the SETTINGS frame by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5431](https://redirect.github.com/quic-go/quic-go/pull/5431)
- http3: read response after encountering error sending the request by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5432](https://redirect.github.com/quic-go/quic-go/pull/5432)
- http3: make Transport.MaxResponseBytes an int by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5433](https://redirect.github.com/quic-go/quic-go/pull/5433)
- http3: add a benchmark for header parsing by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5435](https://redirect.github.com/quic-go/quic-go/pull/5435)
- update qpack to v0.6.0 by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5434](https://redirect.github.com/quic-go/quic-go/pull/5434)
- http3: use QPACK\_DECOMPRESSION\_FAILED for QPACK errors by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5439](https://redirect.github.com/quic-go/quic-go/pull/5439)
- add documentation for Conn.NextConnection by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5442](https://redirect.github.com/quic-go/quic-go/pull/5442)
- ackhandler: don’t generate an immediate ACK for the first packet by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5447](https://redirect.github.com/quic-go/quic-go/pull/5447)
- don’t arm connection timer for connection ID retirement by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5449](https://redirect.github.com/quic-go/quic-go/pull/5449)
- README: add nodepass to list of projects by [@&#8203;yosebyte](https://redirect.github.com/yosebyte) in [#&#8203;5448](https://redirect.github.com/quic-go/quic-go/pull/5448)
- qlogwriter: use synctest to make tests deterministic by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5454](https://redirect.github.com/quic-go/quic-go/pull/5454)
- http3: limit size of decompressed headers by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5452](https://redirect.github.com/quic-go/quic-go/pull/5452)

#### New Contributors

- [@&#8203;yosebyte](https://redirect.github.com/yosebyte) made their first contribution in [#&#8203;5448](https://redirect.github.com/quic-go/quic-go/pull/5448)

**Full Changelog**: <https://github.com/quic-go/quic-go/compare/v0.56.0...v0.57.0>

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @red-hat-konflux on November 24, 2025 at 08:45 AM UTC

### ℹ Artifact update notice

##### File name: go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 1 additional dependency was updated


Details:


| **Package**                | **Change**           |
| :------------------------- | :------------------- |
| `github.com/quic-go/qpack` | `v0.5.1` -> `v0.6.0` |

### Comment by @codecov-commenter on December 01, 2025 at 04:46 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1949?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 58.83%. Comparing base ([`29b134c`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/29b134cd85fb4993e8bc7ab6e6d006aedb8ff205?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`bc48ae7`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/bc48ae72c5c24a83c7c7e7d9877d99334c8e6f3e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 6 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1949      +/-   ##
==========================================
+ Coverage   58.79%   58.83%   +0.03%     
==========================================
  Files         131      131              
  Lines        8407     8407              
==========================================
+ Hits         4943     4946       +3     
+ Misses       2930     2927       -3     
  Partials      534      534              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1949/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1949/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.83% <ø> (+0.03%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1949?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1949*
