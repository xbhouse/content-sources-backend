---
type: pull_request
number: 52
title: "Update module github.com/quic-go/quic-go to v0.56.0"
state: merged
author: red-hat-konflux
created: 2025-11-07T16:20:00Z
updated: 2025-11-10T15:24:31Z
closed: 2025-11-10T15:24:29Z
merged: 2025-11-10T15:24:28Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-quic-go-quic-go-0.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/52
---

# Pull Request #52: Update module github.com/quic-go/quic-go to v0.56.0

**Author**: @red-hat-konflux
**Created**: November 07, 2025 at 04:20 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-quic-go-quic-go-0.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/quic-go/quic-go](https://redirect.github.com/quic-go/quic-go) | `v0.48.2` -> `v0.56.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fquic-go%2fquic-go/v0.56.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fquic-go%2fquic-go/v0.48.2/v0.56.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

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

### [`v0.54.1`](https://redirect.github.com/quic-go/quic-go/compare/v0.54.0...v0.54.1)

[Compare Source](https://redirect.github.com/quic-go/quic-go/compare/v0.54.0...v0.54.1)

### [`v0.54.0`](https://redirect.github.com/quic-go/quic-go/releases/tag/v0.54.0)

[Compare Source](https://redirect.github.com/quic-go/quic-go/compare/v0.53.0...v0.54.0)

This release adds support for [QUIC Stream Resets with Partial Delivery](https://datatracker.ietf.org/doc/html/draft-ietf-quic-reliable-stream-reset-07), a QUIC extension that allows resetting a stream, while guaranteeing delivery of stream data up to a certain byte offset ([#&#8203;5155](https://redirect.github.com/quic-go/quic-go/issues/5155), [#&#8203;5158](https://redirect.github.com/quic-go/quic-go/issues/5158), [#&#8203;5160](https://redirect.github.com/quic-go/quic-go/issues/5160), [#&#8203;5235](https://redirect.github.com/quic-go/quic-go/issues/5235), [#&#8203;5242](https://redirect.github.com/quic-go/quic-go/issues/5242), [#&#8203;5243](https://redirect.github.com/quic-go/quic-go/issues/5243)). This extension is a requirement of newer versions of [WebTransport over HTTP/3](https://datatracker.ietf.org/doc/draft-ietf-webtrans-http3/).

#### Other Notable Changes

- http3: the package now doesn't depend on any internal quic-go packages: [#&#8203;5256](https://redirect.github.com/quic-go/quic-go/issues/5256)
- wire: return concrete structs (instead of a `wire.Frame`) for common frame types (STREAM, DATAGRAM, ACK), speeding up STREAM frame parsing by \~18%: [#&#8203;5253](https://redirect.github.com/quic-go/quic-go/issues/5253), [#&#8203;5227](https://redirect.github.com/quic-go/quic-go/issues/5227), thanks to [@&#8203;jannis-seemann](https://redirect.github.com/jannis-seemann)

#### Fixes

- fix retransmission logic for path probing packets: [#&#8203;5241](https://redirect.github.com/quic-go/quic-go/issues/5241)
- close the` Transport` when `DialAddr` fails: [#&#8203;5259](https://redirect.github.com/quic-go/quic-go/issues/5259), thanks to [@&#8203;rbqvq](https://redirect.github.com/rbqvq)

#### Changelog

- fix retransmission logic for path probing packets by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5241](https://redirect.github.com/quic-go/quic-go/pull/5241)
- implement receiver side behavior for RESET\_STREAM\_AT by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5235](https://redirect.github.com/quic-go/quic-go/pull/5235)
- implement sender side behavior for RESET\_STREAM\_AT by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5242](https://redirect.github.com/quic-go/quic-go/pull/5242)
- fix flaky TestTransportReplaceWithClosed by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5245](https://redirect.github.com/quic-go/quic-go/pull/5245)
- fix flaky TestDrainServerAcceptQueue by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5247](https://redirect.github.com/quic-go/quic-go/pull/5247)
- fix flaky TestServerReceiveQueue by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5249](https://redirect.github.com/quic-go/quic-go/pull/5249)
- http3: fix flaky TestConnGoAwayFailures by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5252](https://redirect.github.com/quic-go/quic-go/pull/5252)
- add a Config and ConnectionState flag for RESET\_STREAM\_AT by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5243](https://redirect.github.com/quic-go/quic-go/pull/5243)
- fix flaky TestPostQuantumClientHello by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5253](https://redirect.github.com/quic-go/quic-go/pull/5253)
- http3: Remove dependency on quic internal package by [@&#8203;rthellend](https://redirect.github.com/rthellend) in [#&#8203;5256](https://redirect.github.com/quic-go/quic-go/pull/5256)
- close Transport when DialAddr fails by [@&#8203;rbqvq](https://redirect.github.com/rbqvq) in [#&#8203;5259](https://redirect.github.com/quic-go/quic-go/pull/5259)
- wire: improve frame parsing benchmarks by [@&#8203;jannis-seemann](https://redirect.github.com/jannis-seemann) in [#&#8203;5263](https://redirect.github.com/quic-go/quic-go/pull/5263)
- optimize parsing logic for STREAM, DATAGRAM and ACK frames by [@&#8203;jannis-seemann](https://redirect.github.com/jannis-seemann) in [#&#8203;5227](https://redirect.github.com/quic-go/quic-go/pull/5227)

#### New Contributors

- [@&#8203;rbqvq](https://redirect.github.com/rbqvq) made their first contribution in [#&#8203;5259](https://redirect.github.com/quic-go/quic-go/pull/5259)

**Full Changelog**: <https://github.com/quic-go/quic-go/compare/v0.53.0...v0.54.0>

### [`v0.53.0`](https://redirect.github.com/quic-go/quic-go/releases/tag/v0.53.0)

[Compare Source](https://redirect.github.com/quic-go/quic-go/compare/v0.52.0...v0.53.0)

This release introduces a massive overhaul of the quic-go API. See this [blog post](https://seemann.io/posts/2025-06-25---quic-go-api-revamp/) for more details about the motivation. Most users will need to make some changes when upgrading to this version.

- The `Connection` interface was removed in favor of a `Conn` struct ([#&#8203;5195](https://redirect.github.com/quic-go/quic-go/issues/5195)).
- The `ReceiveStream`, `SendStream` and `Stream` interfaces were replaced with structs of the same name ([#&#8203;5149](https://redirect.github.com/quic-go/quic-go/issues/5149), [#&#8203;5172](https://redirect.github.com/quic-go/quic-go/issues/5172), [#&#8203;5173](https://redirect.github.com/quic-go/quic-go/issues/5173), [#&#8203;5214](https://redirect.github.com/quic-go/quic-go/issues/5214)).

In most cases, migrating downstream code should be fairly straightforward. For example, a method that used to accept a `quic.Connection` as a parameter now needs to accept a `*quic.Conn`, and a function handling a `quic.Stream` now needs to handle a `*quic.Stream`. Of course, consumers of quic-go are free to define their own interfaces.

Similarly, on the HTTP/3 layer:

- The `Connection` interface was replaced with a `Conn` struct ([#&#8203;5204](https://redirect.github.com/quic-go/quic-go/issues/5204)).
- The `RequestStream` interface was converted to a struct ([#&#8203;5153](https://redirect.github.com/quic-go/quic-go/issues/5153), [#&#8203;5216](https://redirect.github.com/quic-go/quic-go/issues/5216)).
- The `Stream` interface was converted to a struct ([#&#8203;5154](https://redirect.github.com/quic-go/quic-go/issues/5154)).

We expect that most HTTP/3 users won't need to adjust their code, if they use the package to run an HTTP/3 server and dial HTTP/3 connection. More advanced use cases, such as WebTransport and the various MASQUE protocols, will require updates. We have already released new versions of [webtransport-go](https://redirect.github.com/quic-go/webtransport-go/releases/tag/v0.9.0) and [masque-go](https://redirect.github.com/quic-go/masque-go/releases/tag/v0.3.0) to support these changes.

#### Other Breaking Changes

- http3: the deprecated `SingleDestinationRoundTripper` was removed ([#&#8203;5217](https://redirect.github.com/quic-go/quic-go/issues/5217))

#### Notable Fixes and Improvements

- fix Goroutine leak when receiving a Version Negotiation packets race with dial context cancellation ([#&#8203;5203](https://redirect.github.com/quic-go/quic-go/issues/5203))
- drain the server accept queue when closing the transport ([#&#8203;5237](https://redirect.github.com/quic-go/quic-go/issues/5237)), thanks to [@&#8203;sukunrt](https://redirect.github.com/sukunrt)
- fix a race condition when closing transport ([#&#8203;5220](https://redirect.github.com/quic-go/quic-go/issues/5220)), thanks to [@&#8203;sukunrt](https://redirect.github.com/sukunrt)
- quicvarint: speed up parsing of 1, 2 and 4-byte varints (\~12.5% for 1 and 2 bytes, \~1% for 4 bytes) ([#&#8203;5229](https://redirect.github.com/quic-go/quic-go/issues/5229)), thanks to [@&#8203;jannis-seemann](https://redirect.github.com/jannis-seemann)
- http3: expose `ClientConn.Context`, `CloseWithError` and `Conn`: [#&#8203;5219](https://redirect.github.com/quic-go/quic-go/issues/5219)
- http3: `RequestStream` could be misused in many different ways, that's why we tightened the error checks ([#&#8203;5231](https://redirect.github.com/quic-go/quic-go/issues/5231))

#### Behind The Scenes

We've completed the migration of the entire test suite away from Ginkgo ([#&#8203;3652](https://redirect.github.com/quic-go/quic-go/issues/3652)) and towards standard Go tests ([#&#8203;5084](https://redirect.github.com/quic-go/quic-go/issues/5084), [#&#8203;5150](https://redirect.github.com/quic-go/quic-go/issues/5150), [#&#8203;5151](https://redirect.github.com/quic-go/quic-go/issues/5151), [#&#8203;5193](https://redirect.github.com/quic-go/quic-go/issues/5193), [#&#8203;5194](https://redirect.github.com/quic-go/quic-go/issues/5194), [#&#8203;5196](https://redirect.github.com/quic-go/quic-go/issues/5196), [#&#8203;5198](https://redirect.github.com/quic-go/quic-go/issues/5198)). This was a major undertaking, spanning roughly 9 months and resulting in a complete rewrite of quic-go's test suite (> 40,000 lines of code!). Users will now benefit from a significantly slimmed-down dependency tree when upgrading.

#### Changelog

- http3: migrate the stream tests away from Ginkgo by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5150](https://redirect.github.com/quic-go/quic-go/pull/5150)
- http3: migrate the state tracking stream tests away from Ginkgo by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5151](https://redirect.github.com/quic-go/quic-go/pull/5151)
- implement parsing and writing of RESET\_STREAM\_AT frames by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5155](https://redirect.github.com/quic-go/quic-go/pull/5155)
- wire: add support for the reset\_stream\_at transport parameter by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5158](https://redirect.github.com/quic-go/quic-go/pull/5158)
- qlog: add support for reset\_stream\_at frame and transport parameter by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5160](https://redirect.github.com/quic-go/quic-go/pull/5160)
- http3: convert RequestStream from an interface to a struct by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5153](https://redirect.github.com/quic-go/quic-go/pull/5153)
- http3: convert Stream from an interface to a struct by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5154](https://redirect.github.com/quic-go/quic-go/pull/5154)
- http3: simplify HTTP datagram handling by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5156](https://redirect.github.com/quic-go/quic-go/pull/5156)
- http3: use actual QUIC connection and stream in server tests by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5161](https://redirect.github.com/quic-go/quic-go/pull/5161)
- http3: use actual QUIC connection and stream in client tests by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5162](https://redirect.github.com/quic-go/quic-go/pull/5162)
- http3: use actual QUIC connection and stream in conn tests by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5163](https://redirect.github.com/quic-go/quic-go/pull/5163)
- http3: use actual QUIC connection in transport tests by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5164](https://redirect.github.com/quic-go/quic-go/pull/5164)
- http3: use actual QUIC stream in state tracking stream tests by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5166](https://redirect.github.com/quic-go/quic-go/pull/5166)
- http3: use actual QUIC connection in frames tests by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5165](https://redirect.github.com/quic-go/quic-go/pull/5165)
- http3: fix flaky TestClientStreamHijacking by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5169](https://redirect.github.com/quic-go/quic-go/pull/5169)
- http3: use actual QUIC connection in stream tests by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5170](https://redirect.github.com/quic-go/quic-go/pull/5170)
- mockquic: remove package by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5171](https://redirect.github.com/quic-go/quic-go/pull/5171)
- fix flaky TestDatagramLoss by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5174](https://redirect.github.com/quic-go/quic-go/pull/5174)
- fix flaky TestConnectionPathValidation by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5175](https://redirect.github.com/quic-go/quic-go/pull/5175)
- fix flaky TestHTTPRequestAfterGracefulShutdown by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5178](https://redirect.github.com/quic-go/quic-go/pull/5178)
- fix flaky TestGracefulShutdownPendingStreams by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5179](https://redirect.github.com/quic-go/quic-go/pull/5179)
- fix flaky TestTransportReplaceWithClosed by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5181](https://redirect.github.com/quic-go/quic-go/pull/5181)
- http3: fix flaky TestClientResponseValidation by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5183](https://redirect.github.com/quic-go/quic-go/pull/5183)
- http3: fix flaky TestServerRequestHeaderTooLarge by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5186](https://redirect.github.com/quic-go/quic-go/pull/5186)
- http3: fix flaky TestConnControlStreamFailure by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5188](https://redirect.github.com/quic-go/quic-go/pull/5188)
- avoid triggering macOS dual-stack flakiness in HTTP/3 integration tests by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5187](https://redirect.github.com/quic-go/quic-go/pull/5187)
- convert Stream interface to a struct by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5149](https://redirect.github.com/quic-go/quic-go/pull/5149)
- convert SendStream interface to a struct by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5172](https://redirect.github.com/quic-go/quic-go/pull/5172)
- convert ReceiveStream interface to a struct by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5173](https://redirect.github.com/quic-go/quic-go/pull/5173)
- ackhandler: migrate the ECN tests away from Ginkgo by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5084](https://redirect.github.com/quic-go/quic-go/pull/5084)
- congestion: migrate tests away from Ginkgo by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5193](https://redirect.github.com/quic-go/quic-go/pull/5193)
- ci: stop using Ginkgo test command by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5194](https://redirect.github.com/quic-go/quic-go/pull/5194)
- mocks: simplify mockgen command to generate MockCryptoSetup by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5197](https://redirect.github.com/quic-go/quic-go/pull/5197)
- ci: remove leftover check for Ginkgo imports by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5198](https://redirect.github.com/quic-go/quic-go/pull/5198)
- http3: simplify connection closing in the frame parser by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5196](https://redirect.github.com/quic-go/quic-go/pull/5196)
- fix Goroutine leak on version negotiation race with context cancel by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5203](https://redirect.github.com/quic-go/quic-go/pull/5203)
- simplify stream ID handling in the incoming streams map by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5207](https://redirect.github.com/quic-go/quic-go/pull/5207)
- simplify stream ID handling in the outgoing streams map by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5209](https://redirect.github.com/quic-go/quic-go/pull/5209)
- ci: enable Codecov test analysis by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5210](https://redirect.github.com/quic-go/quic-go/pull/5210)
- remove connection flow controller mock by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5213](https://redirect.github.com/quic-go/quic-go/pull/5213)
- handle stream-related frame in the streams map by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5212](https://redirect.github.com/quic-go/quic-go/pull/5212)
- explictly expose all method on the Stream by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5214](https://redirect.github.com/quic-go/quic-go/pull/5214)
- convert Connection interface to Conn struct by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5195](https://redirect.github.com/quic-go/quic-go/pull/5195)
- http3: convert Connection interface to Conn struct by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5204](https://redirect.github.com/quic-go/quic-go/pull/5204)
- http3: remove deprecated SingleDestinationRoundTripper type by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5217](https://redirect.github.com/quic-go/quic-go/pull/5217)
- rename Conn receiver variable by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5215](https://redirect.github.com/quic-go/quic-go/pull/5215)
- ci: enable the nolintlint linter in golangci-lint by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5221](https://redirect.github.com/quic-go/quic-go/pull/5221)
- ci: enable the usetesting linter in golangci-lint by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5222](https://redirect.github.com/quic-go/quic-go/pull/5222)
- ci: add Go 1.25rc1 to tested Go versions by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5224](https://redirect.github.com/quic-go/quic-go/pull/5224)
- http3: add ClientConn.Context, CloseWithError and Conn by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5219](https://redirect.github.com/quic-go/quic-go/pull/5219)
- http3: explicitly expose all method on the RequestStream by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5216](https://redirect.github.com/quic-go/quic-go/pull/5216)
- http3: remove deprecated RoundTripper by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5230](https://redirect.github.com/quic-go/quic-go/pull/5230)
- http3: avoid reinitilising the frame parser on the stream by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5232](https://redirect.github.com/quic-go/quic-go/pull/5232)
- http3: tighten checks for incorrect use of RequestStream by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5231](https://redirect.github.com/quic-go/quic-go/pull/5231)
- improve documentation for the various error types by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5233](https://redirect.github.com/quic-go/quic-go/pull/5233)
- handshake: store key update interval in an atomic by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5234](https://redirect.github.com/quic-go/quic-go/pull/5234)
- fix deadlock when closing the Transport by [@&#8203;sukunrt](https://redirect.github.com/sukunrt) in [#&#8203;5220](https://redirect.github.com/quic-go/quic-go/pull/5220)
- drain server accept queue when the transport is closed by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5237](https://redirect.github.com/quic-go/quic-go/pull/5237)
- quicvarint: speed up parsing of 1, 2 and 4 byte varints  by [@&#8203;jannis-seemann](https://redirect.github.com/jannis-seemann) in [#&#8203;5229](https://redirect.github.com/quic-go/quic-go/pull/5229)

#### New Contributors

- [@&#8203;jannis-seemann](https://redirect.github.com/jannis-seemann) made their first contribution in [#&#8203;5229](https://redirect.github.com/quic-go/quic-go/pull/5229)

**Full Changelog**: <https://github.com/quic-go/quic-go/compare/v0.52.0...v0.53.0>

### [`v0.52.0`](https://redirect.github.com/quic-go/quic-go/releases/tag/v0.52.0)

[Compare Source](https://redirect.github.com/quic-go/quic-go/compare/v0.51.0...v0.52.0)

This release focus on HTTP/3 graceful shutdown using the GOAWAY mechnism.

On the server side graceful shutdown is initiated by calling the `http3.Server.Shutdown` method:

- A single GOAWAY frame is sent, instructing the client to not issue any new requests ([#&#8203;5114](https://redirect.github.com/quic-go/quic-go/issues/5114)).
- New requests are rejected by resetting the streams using the H3\_REQUEST\_REJECTED reset error code: [#&#8203;5116](https://redirect.github.com/quic-go/quic-go/issues/5116).
- QUIC listeners created by the HTTP/3 server (i.e. when using `http3.Server.ListenAndServe` and `ListenAndServeTLS`) are immediately closed ([#&#8203;5101](https://redirect.github.com/quic-go/quic-go/issues/5101)). QUIC listeners created by the application (i.e. when using `http3.Server.ServeListener`) are left running, it is the application's responsibility to close them, or use them in a new server instance ([#&#8203;5129](https://redirect.github.com/quic-go/quic-go/issues/5129)).
- Note that the during the graceful shutdown period, the server does not close existing connections, as this is racy in the presence of packet reordering.

On the client side, when receiving a GOAWAY frame:

- No new streams will be opened on the new connection. Requests to the same origin will be sent on a freshly established QUIC connection.
- Once all requests have completed / were cancelled, the underlying QUIC connection is closed: [#&#8203;5143](https://redirect.github.com/quic-go/quic-go/issues/5143) and [#&#8203;5145](https://redirect.github.com/quic-go/quic-go/issues/5145).

#### Breaking Changes

- `Transport.ConnContext` now passes the `ClientInfo` to the callback, and allows rejecting handshakes by introducing an `error` return value: [#&#8203;5122](https://redirect.github.com/quic-go/quic-go/issues/5122). This allows applications to build more sophisticated DoS defenses. Thanks to [@&#8203;sukunrt](https://redirect.github.com/sukunrt)!
- Connections accepted from a `Listener` using the `Listen` and `ListenAddr` convenience functions now aren't closed when the listener is closed. This makes the shutdown behavior consistent with listeners created from a `Transport`, and the standard library's `net.Listener`: [#&#8203;5108](https://redirect.github.com/quic-go/quic-go/issues/5108).

#### Other Notable Changes

- The TLS ClientHello is now fragmented into multiple pieces, complicating Deep Packet Inspection of the handshake by middlebox: [#&#8203;5107](https://redirect.github.com/quic-go/quic-go/issues/5107). This behavior can be disabled by setting the `QUIC_GO_DISABLE_CLIENTHELLO_SCRAMBLING` environment variable.
- Kernel control messages for ECN and for picking the correct network interface now work on big-endian platforms such as s390x: [#&#8203;5094](https://redirect.github.com/quic-go/quic-go/issues/5094) and [#&#8203;5105](https://redirect.github.com/quic-go/quic-go/issues/5105). Thanks to [@&#8203;Zxilly](https://redirect.github.com/Zxilly)!
- The RTT estimate is now stored in the resumption token (and not in the TLS session ticket): [#&#8203;5065](https://redirect.github.com/quic-go/quic-go/issues/5065). Thanks to [@&#8203;tanghaowillow](https://redirect.github.com/tanghaowillow)!

#### Fixes

- http3: The Alt-Svc entry is now removed when `http3.Server.Serve` returns: [#&#8203;5093](https://redirect.github.com/quic-go/quic-go/issues/5093).
- http3: `http3.Server.ServeQUICConn` now returns `http.ErrServerClosed` when called on a closed server: [#&#8203;5095](https://redirect.github.com/quic-go/quic-go/issues/5095).
- http3: The datagram receive loop now doesn't prematurely return when receiving a datagram for an unknown stream: [#&#8203;5136](https://redirect.github.com/quic-go/quic-go/issues/5136).
- http3: Requests are now only retried when it can be guaranteed that the server didn't process the request: [#&#8203;5141](https://redirect.github.com/quic-go/quic-go/issues/5141).
- http3: Prevent a stream leak when the server sends too many 1xx responses: [#&#8203;5144](https://redirect.github.com/quic-go/quic-go/issues/5144).

#### Behind The Scenes

As in the last couple of releases, we continued our ongoing effort to migrate away from the Ginkgo test suite (tracking issue [#&#8203;3652](https://redirect.github.com/quic-go/quic-go/issues/3652)), mostly in the HTTP/3 package: [#&#8203;5068](https://redirect.github.com/quic-go/quic-go/issues/5068), [#&#8203;5069](https://redirect.github.com/quic-go/quic-go/issues/5069), [#&#8203;5070](https://redirect.github.com/quic-go/quic-go/issues/5070), [#&#8203;5073](https://redirect.github.com/quic-go/quic-go/issues/5073), [#&#8203;5075](https://redirect.github.com/quic-go/quic-go/issues/5075), [#&#8203;5078](https://redirect.github.com/quic-go/quic-go/issues/5078), [#&#8203;5067](https://redirect.github.com/quic-go/quic-go/issues/5067), [#&#8203;5081](https://redirect.github.com/quic-go/quic-go/issues/5081), [#&#8203;5085](https://redirect.github.com/quic-go/quic-go/issues/5085), [#&#8203;5096](https://redirect.github.com/quic-go/quic-go/issues/5096), [#&#8203;5133](https://redirect.github.com/quic-go/quic-go/issues/5133). There are still \~1400 LOC of Ginkgo tests to clean up, scattered across the code base.

#### Changelog

- use assert.AnError consistently in tests by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5066](https://redirect.github.com/quic-go/quic-go/pull/5066)
- http3: migrate the request writer tests away from Ginkgo by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5069](https://redirect.github.com/quic-go/quic-go/pull/5069)
- http3: migrate the headers tests away from Ginkgo by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5068](https://redirect.github.com/quic-go/quic-go/pull/5068)
- http3: simplify request writer by writing to an io.Writer by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5070](https://redirect.github.com/quic-go/quic-go/pull/5070)
- http3: check response writer for http.ResponseController methods by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5071](https://redirect.github.com/quic-go/quic-go/pull/5071)
- http3: migrate the capsule tests away from Ginkgo by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5073](https://redirect.github.com/quic-go/quic-go/pull/5073)
- http3: migrate the response writer tests away from Ginkgo by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5075](https://redirect.github.com/quic-go/quic-go/pull/5075)
- http3: migrate the datagram tests away from Ginkgo by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5076](https://redirect.github.com/quic-go/quic-go/pull/5076)
- http3: migrate the body tests away from Ginkgo by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5078](https://redirect.github.com/quic-go/quic-go/pull/5078)
- http3: migrate the frames tests away from Ginkgo by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5067](https://redirect.github.com/quic-go/quic-go/pull/5067)
- ackhandler: migrate the packet number generator tests away from Ginkgo by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5081](https://redirect.github.com/quic-go/quic-go/pull/5081)
- http3: use httptest.NewRequest by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5086](https://redirect.github.com/quic-go/quic-go/pull/5086)
- http3: update `Hijacker` and `HTTPStreamer` documentation by [@&#8203;TheoTechnicguy](https://redirect.github.com/TheoTechnicguy) in [#&#8203;5089](https://redirect.github.com/quic-go/quic-go/pull/5089)
- http3: use a slice instead of a map to store active listeners by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5087](https://redirect.github.com/quic-go/quic-go/pull/5087)
- http3: remove Alt-Svc entry when Server.Serve returns by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5093](https://redirect.github.com/quic-go/quic-go/pull/5093)
- http3: return http.ErrServerClosed for ServeQUICConn after Server.Close by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5095](https://redirect.github.com/quic-go/quic-go/pull/5095)
- fix: parse ifindex from packet correctly by [@&#8203;Zxilly](https://redirect.github.com/Zxilly) in [#&#8203;5094](https://redirect.github.com/quic-go/quic-go/pull/5094)
- http3: migrate the server tests away from Ginkgo by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5085](https://redirect.github.com/quic-go/quic-go/pull/5085)
- fix dequeuing logic for tiny CRYPTO frames by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5104](https://redirect.github.com/quic-go/quic-go/pull/5104)
- remove periodic logging functionality from packet handler map by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5110](https://redirect.github.com/quic-go/quic-go/pull/5110)
- delete retired connection IDs after 3 PTOs by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;5109](https://redirect.github.com/quic-go/quic-go/pull/5109)
- simplify tracking of Transports for connection migration by [@&

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

### Comment by @red-hat-konflux on November 08, 2025 at 08:36 AM UTC

### ℹ Artifact update notice

##### File name: rh_identity_transform/go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 1 additional dependency was updated


Details:


| **Package**         | **Change**            |
| :------------------ | :-------------------- |
| `golang.org/x/time` | `v0.7.0` -> `v0.12.0` |

---

## Reviews

### Review by @Hyperkid123 - Approved on November 10, 2025 at 03:24 PM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/52*
