---
type: pull_request
number: 2011
title: "Update module github.com/quic-go/quic-go to v0.59.0"
state: merged
author: red-hat-konflux
created: 2026-01-12T08:39:18Z
updated: 2026-01-12T12:40:04Z
closed: 2026-01-12T08:49:37Z
merged: 2026-01-12T08:49:37Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-quic-go-quic-go-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2011
---

# Pull Request #2011: Update module github.com/quic-go/quic-go to v0.59.0

**Author**: @red-hat-konflux
**Created**: January 12, 2026 at 08:39 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-quic-go-quic-go-0.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/quic-go/quic-go](https://redirect.github.com/quic-go/quic-go) | `v0.58.0` -> `v0.59.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fquic-go%2fquic-go/v0.59.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fquic-go%2fquic-go/v0.58.0/v0.59.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>quic-go/quic-go (github.com/quic-go/quic-go)</summary>

### [`v0.59.0`](https://redirect.github.com/quic-go/quic-go/releases/tag/v0.59.0)

[Compare Source](https://redirect.github.com/quic-go/quic-go/compare/v0.58.1...v0.59.0)

This release adds a couple of new features:

- Adds an API to peek stream data on `ReceiveStream` and `Stream`: [#&#8203;5501](https://redirect.github.com/quic-go/quic-go/issues/5501)
- Adds an API to peek the next varint on a stream: [#&#8203;5502](https://redirect.github.com/quic-go/quic-go/issues/5502)
- Reworks the API exposed by the HTTP/3 package for WebTransport: [#&#8203;5509](https://redirect.github.com/quic-go/quic-go/issues/5509), [#&#8203;5512](https://redirect.github.com/quic-go/quic-go/issues/5512). Regular HTTP/3 use cases should not be affected by these changes.
- Adds support for HTTP request trailers (trailers sent by the client): [#&#8203;5507](https://redirect.github.com/quic-go/quic-go/issues/5507)

#### Breaking Changes

- Removes the deprecated `ClientHelloInfo`: [#&#8203;5497](https://redirect.github.com/quic-go/quic-go/issues/5497)
- Removes the deprecated `ConnectionTracingID` and `ConnectionTracingKey`: [#&#8203;5521](https://redirect.github.com/quic-go/quic-go/issues/5521)
- http3: the qlogger is now closed after all streams have been handled: [#&#8203;5524](https://redirect.github.com/quic-go/quic-go/issues/5524)
- The `ConnectionState` now reports both the local and the remote status of the QUIC Datagram and Reliable Stream Reset extensions: [#&#8203;5533](https://redirect.github.com/quic-go/quic-go/issues/5533)

#### Other Notable Fixes

- Fixes an infinite loop of PING-only packets caused by a bug in the PTO queueing logic: [#&#8203;5538](https://redirect.github.com/quic-go/quic-go/issues/5538) and [#&#8203;5539](https://redirect.github.com/quic-go/quic-go/issues/5539)
- http3: Fixes a race condition between new request streams and GOAWAY: [#&#8203;5522](https://redirect.github.com/quic-go/quic-go/issues/5522)
- qlog: Fixes a race condition between `RecordEvent` and `Close`: [#&#8203;5523](https://redirect.github.com/quic-go/quic-go/issues/5523)

#### Changelog

- remove deprecated ClientHelloInfo by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5497](https://redirect.github.com/quic-go/quic-go/pull/5497)
- expose SetReliableBoundary method on the Stream by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5498](https://redirect.github.com/quic-go/quic-go/pull/5498)
- simplify ReceiveStream.Read implementation by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5500](https://redirect.github.com/quic-go/quic-go/pull/5500)
- ci: use codecov-action instead of codecov/test-results-action by [@&#8203;Copilot](https://redirect.github.com/Copilot) in [#&#8203;5504](https://redirect.github.com/quic-go/quic-go/pull/5504)
- http3: add support for request trailers by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5507](https://redirect.github.com/quic-go/quic-go/pull/5507)
- http3: remove stream hijacking API by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5508](https://redirect.github.com/quic-go/quic-go/pull/5508)
- http3: remove Hijacker, add Settingser interface by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5509](https://redirect.github.com/quic-go/quic-go/pull/5509)
- implement a stream peeking API by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5501](https://redirect.github.com/quic-go/quic-go/pull/5501)
- ci: fix Codecov upload by untracking integrationtests before removal by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5511](https://redirect.github.com/quic-go/quic-go/pull/5511)
- quicvarint: implement function to peek the next varint by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5502](https://redirect.github.com/quic-go/quic-go/pull/5502)
- http3: implement new API to allow fine-grained control of connections by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5512](https://redirect.github.com/quic-go/quic-go/pull/5512)
- remove deprecated ConnectionTracingID and ConnectionTracingKey by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5521](https://redirect.github.com/quic-go/quic-go/pull/5521)
- http3: fix race between new streams and GOAWAY by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5522](https://redirect.github.com/quic-go/quic-go/pull/5522)
- qlogwriter: fix race between RecordEvent and Close by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5523](https://redirect.github.com/quic-go/quic-go/pull/5523)
- polish the security policy by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5526](https://redirect.github.com/quic-go/quic-go/pull/5526)
- http3: close qlogger after all streams have been handled by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5524](https://redirect.github.com/quic-go/quic-go/pull/5524)
- fix flaky TestHTTP3Qlog by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5532](https://redirect.github.com/quic-go/quic-go/pull/5532)
- expose local and remote settings in ConnectionState by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5533](https://redirect.github.com/quic-go/quic-go/pull/5533)
- ackhandler: fix qlogging of outstanding packet count by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5538](https://redirect.github.com/quic-go/quic-go/pull/5538)
- ackhandler: fix counting of packets queued for PTO probing by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5539](https://redirect.github.com/quic-go/quic-go/pull/5539)

**Full Changelog**: <https://github.com/quic-go/quic-go/compare/v0.58.0...v0.59.0>

### [`v0.58.1`](https://redirect.github.com/quic-go/quic-go/releases/tag/v0.58.1)

[Compare Source](https://redirect.github.com/quic-go/quic-go/compare/v0.58.0...v0.58.1)

This patch release backports fixes for a bug in the PTO queueing logic that could lead to an infinite loop of PING packets.

#### Bug Fixes

- ackhandler: fix qlogging of outstanding packet count ([#&#8203;5538](https://redirect.github.com/quic-go/quic-go/issues/5538))
- ackhandler: fix counting of packets queued for PTO probing ([#&#8203;5539](https://redirect.github.com/quic-go/quic-go/issues/5539))

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

### Comment by @jira-linking on January 12, 2026 at 08:39 AM UTC

Commits missing Jira IDs:
a000a85f67888a1a4fae96ac95016ba03763453b


### Comment by @codecov-commenter on January 12, 2026 at 08:45 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2011?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.01%. Comparing base ([`62f4a7f`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/62f4a7f52de3e53edb10206b8b3927d13120177b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`a000a85`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/a000a85f67888a1a4fae96ac95016ba03763453b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2011   +/-   ##
=======================================
  Coverage   59.01%   59.01%           
=======================================
  Files         131      131           
  Lines        8493     8493           
=======================================
  Hits         5012     5012           
  Misses       2947     2947           
  Partials      534      534           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2011/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2011/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.01% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2011?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2011*
