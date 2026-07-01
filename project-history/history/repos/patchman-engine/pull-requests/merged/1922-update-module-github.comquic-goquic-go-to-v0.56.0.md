---
type: pull_request
number: 1922
title: "Update module github.com/quic-go/quic-go to v0.56.0"
state: merged
author: red-hat-konflux
created: 2025-11-10T04:17:03Z
updated: 2025-11-10T08:20:26Z
closed: 2025-11-10T04:31:34Z
merged: 2025-11-10T04:31:34Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-quic-go-quic-go-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1922
---

# Pull Request #1922: Update module github.com/quic-go/quic-go to v0.56.0

**Author**: @red-hat-konflux
**Created**: November 10, 2025 at 04:17 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-quic-go-quic-go-0.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/quic-go/quic-go](https://redirect.github.com/quic-go/quic-go) | `v0.55.0` -> `v0.56.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fquic-go%2fquic-go/v0.56.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fquic-go%2fquic-go/v0.55.0/v0.56.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>quic-go/quic-go (github.com/quic-go/quic-go)</summary>

### [`v0.56.0`](https://redirect.github.com/quic-go/quic-go/releases/tag/v0.56.0)

[Compare Source](https://redirect.github.com/quic-go/quic-go/compare/v0.55.0...v0.56.0)

This release introduces qlog support for HTTP/3 ([#&#8203;5367](https://redirect.github.com/quic-go/quic-go/issues/5367), [#&#8203;5372](https://redirect.github.com/quic-go/quic-go/issues/5372), [#&#8203;5374](https://redirect.github.com/quic-go/quic-go/issues/5374), [#&#8203;5375](https://redirect.github.com/quic-go/quic-go/issues/5375), [#&#8203;5376](https://redirect.github.com/quic-go/quic-go/issues/5376), [#&#8203;5381](https://redirect.github.com/quic-go/quic-go/issues/5381), [#&#8203;5383](https://redirect.github.com/quic-go/quic-go/issues/5383)).

For this, we completely changed how connection tracing works. Instead of a general-purpose `logging.ConnectionTracer` (which we removed entirely), we now have a qlog-specific tracer ([#&#8203;5356](https://redirect.github.com/quic-go/quic-go/issues/5356), [#&#8203;5417](https://redirect.github.com/quic-go/quic-go/issues/5417)). quic-go users can now implement their own qlog events.

It also removes the Prometheus-based metrics collection. Please comment on the tracking issue ([#&#8203;5294](https://redirect.github.com/quic-go/quic-go/issues/5294)) if you rely on metrics and are interested in seeing metrics brought back in a future release.

#### Notable Changes

- replaced the unmaintained gojay with a custom, performance-optimized JSON encoder ([#&#8203;5353](https://redirect.github.com/quic-go/quic-go/issues/5353), [#&#8203;5371](https://redirect.github.com/quic-go/quic-go/issues/5371))
- quicvarint: improved panic message for numbers larger than 2^62 ([#&#8203;5410](https://redirect.github.com/quic-go/quic-go/issues/5410))

#### Behind the Scenes

Go 1.25 [introduced](https://go.dev/blog/synctest) support for testing concurrent code using `testing/synctest`. We've been working on transitioning tests to use synctest ([#&#8203;5357](https://redirect.github.com/quic-go/quic-go/issues/5357), [#&#8203;5391](https://redirect.github.com/quic-go/quic-go/issues/5391), [#&#8203;5393](https://redirect.github.com/quic-go/quic-go/issues/5393), [#&#8203;5397](https://redirect.github.com/quic-go/quic-go/issues/5397), [#&#8203;5398](https://redirect.github.com/quic-go/quic-go/issues/5398), [#&#8203;5403](https://redirect.github.com/quic-go/quic-go/issues/5403), [#&#8203;5414](https://redirect.github.com/quic-go/quic-go/issues/5414), [#&#8203;5415](https://redirect.github.com/quic-go/quic-go/issues/5415)), using [@&#8203;MarcoPolo](https://redirect.github.com/MarcoPolo)'s simnet package to simulate a network in memory.

Using synctest makes test execution more reliable (reducing flakiness). The use of a synthetic clock leads to a massive speedup; the execution time of some integration tests was reduced from 20s to less than 1ms. The work will continue for the next release (see tracking issue: [#&#8203;5386](https://redirect.github.com/quic-go/quic-go/issues/5386)).

#### Changelog

- qlog: implement a minimal jsontext-like JSON encoder by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5353](https://redirect.github.com/quic-go/quic-go/pull/5353)
- ci: remove 386 (32 bit x86) by [@&#8203;MarcoPolo](https://redirect.github.com/MarcoPolo) in [#&#8203;5352](https://redirect.github.com/quic-go/quic-go/pull/5352)
- use synctest in more connection tests by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5357](https://redirect.github.com/quic-go/quic-go/pull/5357)
- qlog: split serializiation and event definitions, remove logging abstraction by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5356](https://redirect.github.com/quic-go/quic-go/pull/5356)
- qlogwriter: implement the draft-12 trace header by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5360](https://redirect.github.com/quic-go/quic-go/pull/5360)
- qlogwriter: add support for event\_schemas in the trace header by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5361](https://redirect.github.com/quic-go/quic-go/pull/5361)
- qlogwriter: pass the event time to Event.Encode by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5362](https://redirect.github.com/quic-go/quic-go/pull/5362)
- ackhandler: fix qlogging of alarm timer expiration time by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5363](https://redirect.github.com/quic-go/quic-go/pull/5363)
- qlog: privatize Encode functions of non-Event structs by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5364](https://redirect.github.com/quic-go/quic-go/pull/5364)
- fix qlogging of the short header payload length by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5365](https://redirect.github.com/quic-go/quic-go/pull/5365)
- ci: include OS and Go version in Codecov test report upload by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5370](https://redirect.github.com/quic-go/quic-go/pull/5370)
- http3: add basic server-side qlog support by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5367](https://redirect.github.com/quic-go/quic-go/pull/5367)
- jsontext: add support for encoding null by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5371](https://redirect.github.com/quic-go/quic-go/pull/5371)
- qlog: use PathEndpointInfo in connection\_started by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5368](https://redirect.github.com/quic-go/quic-go/pull/5368)
- http3: fix qlog encoding of frame\_parsed and frame\_created events by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5372](https://redirect.github.com/quic-go/quic-go/pull/5372)
- http3: add basic client-side qlog support by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5374](https://redirect.github.com/quic-go/quic-go/pull/5374)
- readme: update oss-fuzz link by [@&#8203;kriztalz](https://redirect.github.com/kriztalz) in [#&#8203;5377](https://redirect.github.com/quic-go/quic-go/pull/5377)
- http3: qlog sent and received GOAWAY frames by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5376](https://redirect.github.com/quic-go/quic-go/pull/5376)
- http3: qlog sent and received DATAGRAMs by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5375](https://redirect.github.com/quic-go/quic-go/pull/5375)
- http3: move qlogging of frames into the frame parser by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5378](https://redirect.github.com/quic-go/quic-go/pull/5378)
- http3: qlog sent and received SETTINGS frames by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5379](https://redirect.github.com/quic-go/quic-go/pull/5379)
- http3: qlog the frame length and payload length of parsed frames by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5380](https://redirect.github.com/quic-go/quic-go/pull/5380)
- http3: qlog reserved, unsupported and unknown frames by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5381](https://redirect.github.com/quic-go/quic-go/pull/5381)
- http3: add the qlog event schema to trace header by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5383](https://redirect.github.com/quic-go/quic-go/pull/5383)
- use default RTT (100ms) for 0-RTT if no prior estimate by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5388](https://redirect.github.com/quic-go/quic-go/pull/5388)
- congestion: avoid overflows when calculating pacer budget by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5390](https://redirect.github.com/quic-go/quic-go/pull/5390)
- add simnet package to simulate a net.PacketConn in memory by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5385](https://redirect.github.com/quic-go/quic-go/pull/5385)
- use synctest for transport tests by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5391](https://redirect.github.com/quic-go/quic-go/pull/5391)
- use simnet in CONNECTION\_CLOSE retransmission test by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5395](https://redirect.github.com/quic-go/quic-go/pull/5395)
- use synctest for the packet drop test by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5393](https://redirect.github.com/quic-go/quic-go/pull/5393)
- use synctest for the handshake drop test by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5397](https://redirect.github.com/quic-go/quic-go/pull/5397)
- use synctest for the datagram test by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5398](https://redirect.github.com/quic-go/quic-go/pull/5398)
- use synctest for the timeout tests by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5403](https://redirect.github.com/quic-go/quic-go/pull/5403)
- fix flaky TestConnectionUnpackFailureDropped by [@&#8203;Copilot](https://redirect.github.com/Copilot) in [#&#8203;5382](https://redirect.github.com/quic-go/quic-go/pull/5382)
- fix flaky TestServerTransportClose by [@&#8203;Copilot](https://redirect.github.com/Copilot) in [#&#8203;5407](https://redirect.github.com/quic-go/quic-go/pull/5407)
- ci: use gcassert to check that quicvarint.Len is inlined by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5409](https://redirect.github.com/quic-go/quic-go/pull/5409)
- quicvarint: improve panic message for numbers larger 2^62 by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5410](https://redirect.github.com/quic-go/quic-go/pull/5410)
- ci: bump actions/upload-artifact from 4 to 5 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;5411](https://redirect.github.com/quic-go/quic-go/pull/5411)
- ci: update golangci-lint to v2.6.0 by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5412](https://redirect.github.com/quic-go/quic-go/pull/5412)
- qlog: rename owner to initiator by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5416](https://redirect.github.com/quic-go/quic-go/pull/5416)
- use synctest for the stateless reset tests by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5415](https://redirect.github.com/quic-go/quic-go/pull/5415)
- ackhandler: fix qlogging of RTT values by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5418](https://redirect.github.com/quic-go/quic-go/pull/5418)
- qlog: rework the ConnectionClosed event by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5417](https://redirect.github.com/quic-go/quic-go/pull/5417)
- qlog: split the PTO count updates ouf of the MetricsUpdated event by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5421](https://redirect.github.com/quic-go/quic-go/pull/5421)

#### New Contributors

- [@&#8203;kriztalz](https://redirect.github.com/kriztalz) made their first contribution in [#&#8203;5377](https://redirect.github.com/quic-go/quic-go/pull/5377)
- [@&#8203;Copilot](https://redirect.github.com/Copilot) made their first contribution in [#&#8203;5382](https://redirect.github.com/quic-go/quic-go/pull/5382)

**Full Changelog**: <https://github.com/quic-go/quic-go/compare/v0.55.0...v0.56.0>

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

### Comment by @codecov-commenter on November 10, 2025 at 04:22 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1922?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 58.90%. Comparing base ([`adae82b`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/adae82bcb08271fd9297c40c73ee6a2307521190?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`dff6bd6`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/dff6bd6fdb8808eb82d4b5dba5b75c0ee72ac60c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1922   +/-   ##
=======================================
  Coverage   58.90%   58.90%           
=======================================
  Files         131      131           
  Lines        8421     8421           
=======================================
  Hits         4960     4960           
  Misses       2927     2927           
  Partials      534      534           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1922/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1922/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.90% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1922?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1922*
