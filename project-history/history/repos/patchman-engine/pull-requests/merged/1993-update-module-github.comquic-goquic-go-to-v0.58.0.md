---
type: pull_request
number: 1993
title: "Update module github.com/quic-go/quic-go to v0.58.0"
state: merged
author: red-hat-konflux
created: 2025-12-22T04:43:16Z
updated: 2025-12-22T08:40:28Z
closed: 2025-12-22T06:13:24Z
merged: 2025-12-22T06:13:24Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-quic-go-quic-go-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1993
---

# Pull Request #1993: Update module github.com/quic-go/quic-go to v0.58.0

**Author**: @red-hat-konflux
**Created**: December 22, 2025 at 04:43 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-quic-go-quic-go-0.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/quic-go/quic-go](https://redirect.github.com/quic-go/quic-go) | `v0.57.1` -> `v0.58.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fquic-go%2fquic-go/v0.58.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fquic-go%2fquic-go/v0.57.1/v0.58.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>quic-go/quic-go (github.com/quic-go/quic-go)</summary>

### [`v0.58.0`](https://redirect.github.com/quic-go/quic-go/releases/tag/v0.58.0)

[Compare Source](https://redirect.github.com/quic-go/quic-go/compare/v0.57.1...v0.58.0)

This release optimizes the QUIC handshake:

- Multiple incoming packets are now processed before sending an acknowledgment, reducing the total number of packets sent: [#&#8203;5451](https://redirect.github.com/quic-go/quic-go/issues/5451)
- ACK frames are now packed into coalesced packets, reducing the need to send a separate packet just for the ACK in many cases: [#&#8203;5477](https://redirect.github.com/quic-go/quic-go/issues/5477)
- When packets are buffered during the handshake, this now doesn't lead to inflated RTT measurements anymore: [#&#8203;5493](https://redirect.github.com/quic-go/quic-go/issues/5493), [#&#8203;5494](https://redirect.github.com/quic-go/quic-go/issues/5494)

#### Other notable changes

- quic-go now has a new logo: [#&#8203;5484](https://redirect.github.com/quic-go/quic-go/issues/5484)
- ACK frames can now be encoded with up to 64 ranges (previously: 32): [#&#8203;5476](https://redirect.github.com/quic-go/quic-go/issues/5476)
- Serializing ACK frames is now significantly faster: [#&#8203;5476](https://redirect.github.com/quic-go/quic-go/issues/5476)
- Improved batch packet processing logic: [#&#8203;5478](https://redirect.github.com/quic-go/quic-go/issues/5478)
- qlog: added support for the `datagram_id` on `packet_sent`, `packet_received` and `packet_buffered` events, using the CRC32 of the packet

#### Changelog

- qlog: add a DebugEvent by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5456](https://redirect.github.com/quic-go/quic-go/pull/5456)
- qlog the datagram\_id for long header and coalesced packets by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5455](https://redirect.github.com/quic-go/quic-go/pull/5455)
- build(deps): bump actions/checkout from 5 to 6 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;5458](https://redirect.github.com/quic-go/quic-go/pull/5458)
- handshake: pass cipherSuite by value instead of pointer by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5459](https://redirect.github.com/quic-go/quic-go/pull/5459)
- handshake: add a benchmark for token decoding by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5460](https://redirect.github.com/quic-go/quic-go/pull/5460)
- use synctest in the session resumption test by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5424](https://redirect.github.com/quic-go/quic-go/pull/5424)
- handshake: fix QUIC events with session tickets disabled on Go 1.26 by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5462](https://redirect.github.com/quic-go/quic-go/pull/5462)
- fix flaky TestConnectionCongestionControl by [@&#8203;Copilot](https://redirect.github.com/Copilot) in [#&#8203;5467](https://redirect.github.com/quic-go/quic-go/pull/5467)
- simnet: move latency configuration to Router, simplify queueing logic by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5463](https://redirect.github.com/quic-go/quic-go/pull/5463)
- use synctest for the RTT and reordering tests by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5464](https://redirect.github.com/quic-go/quic-go/pull/5464)
- use synctest for 0-RTT integration tests by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5465](https://redirect.github.com/quic-go/quic-go/pull/5465)
- allow processing of multiple packets before handshake completion by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5451](https://redirect.github.com/quic-go/quic-go/pull/5451)
- ackhandler: disentangle ReceivedPacketHandler and SentPacketHandler by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5469](https://redirect.github.com/quic-go/quic-go/pull/5469)
- refactor connection tests to remove ReceivedPacketHandler mock by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5468](https://redirect.github.com/quic-go/quic-go/pull/5468)
- ackhandler: remove ReceivedPacketHandler interface, use struct directly by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5472](https://redirect.github.com/quic-go/quic-go/pull/5472)
- wire: add a function to trunctate an ACK frame by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5476](https://redirect.github.com/quic-go/quic-go/pull/5476)
- allow packing of ACKs in coalesced packets by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5477](https://redirect.github.com/quic-go/quic-go/pull/5477)
- improve batch packet processing logic by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5478](https://redirect.github.com/quic-go/quic-go/pull/5478)
- build(deps): bump actions/cache from 4 to 5 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;5482](https://redirect.github.com/quic-go/quic-go/pull/5482)
- build(deps): bump actions/upload-artifact from 5 to 6 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;5483](https://redirect.github.com/quic-go/quic-go/pull/5483)
- update the logo by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5484](https://redirect.github.com/quic-go/quic-go/pull/5484)
- ci: add Go 1.26rc1 to tested Go versions by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5486](https://redirect.github.com/quic-go/quic-go/pull/5486)
- utils: make TestAddTimestamp work in all time zones by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5492](https://redirect.github.com/quic-go/quic-go/pull/5492)
- ackhandler: record RTT measurements for non-ack-eliciting packets by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5494](https://redirect.github.com/quic-go/quic-go/pull/5494)
- ackhandler: only generate RTT sample for the last ack-eliciting packet by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5493](https://redirect.github.com/quic-go/quic-go/pull/5493)

**Full Changelog**: <https://github.com/quic-go/quic-go/compare/v0.57.0...v0.58.0>

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

### Comment by @jira-linking on December 22, 2025 at 04:43 AM UTC

Commits missing Jira IDs:
5f08fba4c640047b5e52dc79e11656e5f28e005b


### Comment by @codecov-commenter on December 22, 2025 at 04:49 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1993?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.01%. Comparing base ([`e90b3ef`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/e90b3ef44c967d50ee185921e94081b25e639c62?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`5f08fba`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/5f08fba4c640047b5e52dc79e11656e5f28e005b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 6 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1993   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1993/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1993/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.01% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1993?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1993*
