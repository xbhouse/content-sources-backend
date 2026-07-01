---
type: pull_request
number: 95
title: "chore(deps): update module github.com/quic-go/quic-go to v0.58.0 - autoclosed"
state: closed
author: red-hat-konflux
created: 2025-12-08T00:26:59Z
updated: 2026-01-07T08:42:41Z
closed: 2026-01-07T08:42:40Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-quic-go-quic-go-0.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/95
---

# Pull Request #95: chore(deps): update module github.com/quic-go/quic-go to v0.58.0 - autoclosed

**Author**: @red-hat-konflux
**Created**: December 08, 2025 at 12:26 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-quic-go-quic-go-0.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/quic-go/quic-go](https://redirect.github.com/quic-go/quic-go) | `v0.56.0` -> `v0.58.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fquic-go%2fquic-go/v0.58.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fquic-go%2fquic-go/v0.56.0/v0.58.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

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

### [`v0.57.1`](https://redirect.github.com/quic-go/quic-go/releases/tag/v0.57.1)

[Compare Source](https://redirect.github.com/quic-go/quic-go/compare/v0.57.0...v0.57.1)

This release resolves a panic during the server handshake when using the upcoming Go 1.26 toolchain, specifically occurring with TLS session tickets disabled ([#&#8203;5462](https://redirect.github.com/quic-go/quic-go/issues/5462)). This issue does not impact builds on Go 1.25 or earlier versions.

### [`v0.57.0`](https://redirect.github.com/quic-go/quic-go/releases/tag/v0.57.0)

[Compare Source](https://redirect.github.com/quic-go/quic-go/compare/v0.56.0...v0.57.0)

This release contains a fix for CVE-2025-64702 by reworking the HTTP/3 header processing logic:

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

📅 **Schedule**: Branch creation - At any time (no schedule defined), Automerge - At any time (no schedule defined).

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

---
### Documentation

Find out how to configure dependency updates in [MintMaker documentation](https://konflux-ci.dev/docs/mintmaker/user/) or see all available configuration options in [Renovate documentation](https://docs.renovatebot.com/configuration-options/).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6W119-->


---

## Discussion

### Comment by @red-hat-konflux on December 21, 2025 at 12:38 PM UTC

### ℹ Artifact update notice

##### File name: rh_identity_transform/go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 1 additional dependency was updated


Details:


| **Package**                | **Change**           |
| :------------------------- | :------------------- |
| `github.com/quic-go/qpack` | `v0.5.1` -> `v0.6.0` |

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/95*
