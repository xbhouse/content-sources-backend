---
type: pull_request
number: 1876
title: "Update module github.com/quic-go/quic-go to v0.55.0"
state: merged
author: red-hat-konflux
created: 2025-10-06T08:16:34Z
updated: 2025-10-08T16:16:20Z
closed: 2025-10-06T08:22:08Z
merged: 2025-10-06T08:22:08Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-quic-go-quic-go-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1876
---

# Pull Request #1876: Update module github.com/quic-go/quic-go to v0.55.0

**Author**: @red-hat-konflux
**Created**: October 06, 2025 at 08:16 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-quic-go-quic-go-0.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/quic-go/quic-go](https://redirect.github.com/quic-go/quic-go) | `v0.54.1` -> `v0.55.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fquic-go%2fquic-go/v0.55.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fquic-go%2fquic-go/v0.54.1/v0.55.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>quic-go/quic-go (github.com/quic-go/quic-go)</summary>

### [`v0.55.0`](https://redirect.github.com/quic-go/quic-go/releases/tag/v0.55.0)

[Compare Source](https://redirect.github.com/quic-go/quic-go/compare/v0.54.1...v0.55.0)

This release contains a number of improvements and fixes, and it updates the supported Go versions to 1.24 and 1.25.

#### Optimizations

When sending packets on a QUIC connection, RFC 9002 requires us to save the timestamp for every packet sent. In [#&#8203;5344](https://redirect.github.com/quic-go/quic-go/issues/5344), we implemented a memory-optimized drop-in replacement for `time.Time`, which reduces the memory required from 24 to 8 bytes, and vastly speeds up timer calculations (which happen very frequently).

#### New Features

- Basic connection statistics are now exposed via `Conn.ConnectionStats`, thanks to [@&#8203;MarcoPolo](https://redirect.github.com/MarcoPolo)
- On some links, packet reordering can lead to spurious detections of packet loss when using the loss detection logic specified in RFC 9002. [#&#8203;5355](https://redirect.github.com/quic-go/quic-go/issues/5355) adds logic detect when packet loss is detected spuriously.

#### Notable Fixes

- http3: don't allow usage of closed `Transport`: [#&#8203;5324](https://redirect.github.com/quic-go/quic-go/issues/5324), thanks to [@&#8203;Glonee](https://redirect.github.com/Glonee)
- http3: fix race in concurrent `Transport.Roundtrip` calls: [#&#8203;5323](https://redirect.github.com/quic-go/quic-go/issues/5323), thanks to [@&#8203;Glonee](https://redirect.github.com/Glonee)
- improve and fix connection timer logic: [#&#8203;5339](https://redirect.github.com/quic-go/quic-go/issues/5339), thanks to [@&#8203;sukunrt](https://redirect.github.com/sukunrt) for a very comprehensive code review

#### Behind the Scenes

We have started transitioning tests to make use of the new `synctest` package that was added in Go 1.25 (and was available as a `GOEXPERIMENT` in Go 1.24): [#&#8203;5291](https://redirect.github.com/quic-go/quic-go/issues/5291), [#&#8203;5296](https://redirect.github.com/quic-go/quic-go/issues/5296), [#&#8203;5298](https://redirect.github.com/quic-go/quic-go/issues/5298), [#&#8203;5299](https://redirect.github.com/quic-go/quic-go/issues/5299), [#&#8203;5302](https://redirect.github.com/quic-go/quic-go/issues/5302), [#&#8203;5304](https://redirect.github.com/quic-go/quic-go/issues/5304), [#&#8203;5305](https://redirect.github.com/quic-go/quic-go/issues/5305), [#&#8203;5306](https://redirect.github.com/quic-go/quic-go/issues/5306), [#&#8203;5317](https://redirect.github.com/quic-go/quic-go/issues/5317). This is a lot of work, but it makes the test execution both faster and more reliable.

#### Changelog

- wire: implement parsing and writing of the ACK\_FREQUENCY frame by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5264](https://redirect.github.com/quic-go/quic-go/pull/5264)
- wire: implement parsing and writing of the IMMEDIATE\_ACK frame by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5265](https://redirect.github.com/quic-go/quic-go/pull/5265)
- fuzzing: fix timeout in frame parser by [@&#8203;jannis-seemann](https://redirect.github.com/jannis-seemann) in [#&#8203;5268](https://redirect.github.com/quic-go/quic-go/pull/5268)
- wire: add support for the min\_ack\_delay transport parameter by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5266](https://redirect.github.com/quic-go/quic-go/pull/5266)
- fix missing log statement for STREAM, DATAGRAM and ACK by [@&#8203;jannis-seemann](https://redirect.github.com/jannis-seemann) in [#&#8203;5273](https://redirect.github.com/quic-go/quic-go/pull/5273)
- qlog: add support for ACK\_FREQUENCY and IMMEDIATE\_ACK frames by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5276](https://redirect.github.com/quic-go/quic-go/pull/5276)
- ackhandler: remove unused time from receivedPacketHandler.ReceivedPacket by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5277](https://redirect.github.com/quic-go/quic-go/pull/5277)
- quicvarint: extend benchmark to use quicvarint.Reader by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5278](https://redirect.github.com/quic-go/quic-go/pull/5278)
- quicvarint: tolerate empty reads of the underlying io.Reader by [@&#8203;bemasc](https://redirect.github.com/bemasc) in [#&#8203;5275](https://redirect.github.com/quic-go/quic-go/pull/5275)
- http3: fix documentation for Server.ServeListener by [@&#8203;WeidiDeng](https://redirect.github.com/WeidiDeng) in [#&#8203;5282](https://redirect.github.com/quic-go/quic-go/pull/5282)
- expose connection stats via Conn.ConnectionStats by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5281](https://redirect.github.com/quic-go/quic-go/pull/5281)
- ackhandler: generalize check for missing packets below threshold by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5260](https://redirect.github.com/quic-go/quic-go/pull/5260)
- update to Go 1.25, drop Go 1.23, use go tool for gomock by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5283](https://redirect.github.com/quic-go/quic-go/pull/5283)
- replace `interface{}` with `any` by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5290](https://redirect.github.com/quic-go/quic-go/pull/5290)
- use testing.B.Loop in all benchmark tests by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5285](https://redirect.github.com/quic-go/quic-go/pull/5285)
- build(deps): bump actions/checkout from 4 to 5 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;5293](https://redirect.github.com/quic-go/quic-go/pull/5293)
- use synctest to make receive stream tests fully deterministc by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5291](https://redirect.github.com/quic-go/quic-go/pull/5291)
- use synctest to make streams map tests fully deterministic by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5296](https://redirect.github.com/quic-go/quic-go/pull/5296)
- ci: cache the Go build cache for cross-compilation workflow by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5297](https://redirect.github.com/quic-go/quic-go/pull/5297)
- ci: fix cache save and restore logic for cross compile workflow by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5300](https://redirect.github.com/quic-go/quic-go/pull/5300)
- restore previously deleted TestStreamsMapConcurrent test by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5301](https://redirect.github.com/quic-go/quic-go/pull/5301)
- use synctest to make the send queue tests fully deterministic by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5302](https://redirect.github.com/quic-go/quic-go/pull/5302)
- use synctest to make the send stream tests fully deterministic by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5298](https://redirect.github.com/quic-go/quic-go/pull/5298)
- ci: use `go mod tidy -diff` to check for tidied `go.mod` by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5303](https://redirect.github.com/quic-go/quic-go/pull/5303)
- use synctest to make the datagram queue tests fully deterministic by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5305](https://redirect.github.com/quic-go/quic-go/pull/5305)
- utils: use synctest to make the timer tests fully deterministic by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5306](https://redirect.github.com/quic-go/quic-go/pull/5306)
- ackhandler: fix resetting of packet.isPathProbePacket by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5310](https://redirect.github.com/quic-go/quic-go/pull/5310)
- ackhandler: use an iterator to process received packet ranges by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5309](https://redirect.github.com/quic-go/quic-go/pull/5309)
- ackhandler: use a typed mock for the ECNHandler by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5311](https://redirect.github.com/quic-go/quic-go/pull/5311)
- ackhandler: immediately clear ackedPacket slice after processing ACK by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5313](https://redirect.github.com/quic-go/quic-go/pull/5313)
- ci: improve cache key generation for the cross compilation job by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5315](https://redirect.github.com/quic-go/quic-go/pull/5315)
- ci: fix cache paths in cross compile workflow by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5318](https://redirect.github.com/quic-go/quic-go/pull/5318)
- ackhandler: avoid storing packet number in packet struct by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5312](https://redirect.github.com/quic-go/quic-go/pull/5312)
- ackhandler: store skipped packet numbers separately by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5314](https://redirect.github.com/quic-go/quic-go/pull/5314)
- ackhandler: account for skipped packets in packet threshold calculation by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5316](https://redirect.github.com/quic-go/quic-go/pull/5316)
- ackhandler: store the last four skipped packets by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5322](https://redirect.github.com/quic-go/quic-go/pull/5322)
- http3: fix data race in Transport by [@&#8203;Glonee](https://redirect.github.com/Glonee) in [#&#8203;5323](https://redirect.github.com/quic-go/quic-go/pull/5323)
- qlog: add a benchmark for the ConnectionTracer by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5328](https://redirect.github.com/quic-go/quic-go/pull/5328)
- qlog: merge event category and name by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5329](https://redirect.github.com/quic-go/quic-go/pull/5329)
- http3: don't allow usage of closed Transport by [@&#8203;Glonee](https://redirect.github.com/Glonee) in [#&#8203;5324](https://redirect.github.com/quic-go/quic-go/pull/5324)
- build(deps): bump actions/setup-go from 5 to 6 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;5330](https://redirect.github.com/quic-go/quic-go/pull/5330)
- fix: return stream frames to pool on error paths by [@&#8203;lidel](https://redirect.github.com/lidel) in [#&#8203;5327](https://redirect.github.com/quic-go/quic-go/pull/5327)
- ackhandler: add a benchmark for sending and acknowledging packets by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5333](https://redirect.github.com/quic-go/quic-go/pull/5333)
- implement a memory-optimized time.Time replacement by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5334](https://redirect.github.com/quic-go/quic-go/pull/5334)
- add a benchmark test for data transfers by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5335](https://redirect.github.com/quic-go/quic-go/pull/5335)
- improve connection timer logic by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5339](https://redirect.github.com/quic-go/quic-go/pull/5339)
- use synctest to make the connection tests fully deterministic by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5317](https://redirect.github.com/quic-go/quic-go/pull/5317)
- drop initial packets when the handshake is confirmed by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5354](https://redirect.github.com/quic-go/quic-go/pull/5354)
- protocol: optimize ConnectionID.String by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5351](https://redirect.github.com/quic-go/quic-go/pull/5351)
- fix missing tracing of restored transport parameters by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5349](https://redirect.github.com/quic-go/quic-go/pull/5349)
- ackhandler: track lost packets and detect spurious losses by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5355](https://redirect.github.com/quic-go/quic-go/pull/5355)

#### New Contributors

- [@&#8203;bemasc](https://redirect.github.com/bemasc) made their first contribution in [#&#8203;5275](https://redirect.github.com/quic-go/quic-go/pull/5275)
- [@&#8203;lidel](https://redirect.github.com/lidel) made their first contribution in [#&#8203;5327](https://redirect.github.com/quic-go/quic-go/pull/5327)

**Full Changelog**: <https://github.com/quic-go/quic-go/compare/v0.54.0...v0.55.0>

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

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @jira-linking on October 06, 2025 at 08:16 AM UTC

Commits missing Jira IDs:
df7a05ca6e230c1bc84eb868c4180cb229e51384


### Comment by @codecov-commenter on October 06, 2025 at 08:21 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1876?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 54.44%. Comparing base ([`43cc7e8`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/43cc7e82ecf5394f27affff68d56d601271406f7?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`df7a05c`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/df7a05ca6e230c1bc84eb868c4180cb229e51384?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 3 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1876   +/-   ##
=======================================
  Coverage   54.44%   54.44%           
=======================================
  Files         138      138           
  Lines       10735    10735           
=======================================
  Hits         5845     5845           
  Misses       4350     4350           
  Partials      540      540           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1876/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1876/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `54.44% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1876?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1876*
