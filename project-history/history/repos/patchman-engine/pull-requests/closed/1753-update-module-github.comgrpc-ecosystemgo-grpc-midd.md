---
type: pull_request
number: 1753
title: "Update module github.com/grpc-ecosystem/go-grpc-middleware to v2"
state: closed
author: red-hat-konflux
created: 2025-07-27T13:50:01Z
updated: 2025-07-28T10:34:37Z
closed: 2025-07-28T08:24:13Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-grpc-ecosystem-go-grpc-middleware-2.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1753
---

# Pull Request #1753: Update module github.com/grpc-ecosystem/go-grpc-middleware to v2

**Author**: @red-hat-konflux
**Created**: July 27, 2025 at 01:50 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-grpc-ecosystem-go-grpc-middleware-2.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/grpc-ecosystem/go-grpc-middleware](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware) | `v1.4.0` -> `v2.3.2` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fgrpc-ecosystem%2fgo-grpc-middleware/v2.3.2?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fgrpc-ecosystem%2fgo-grpc-middleware/v1.4.0/v2.3.2?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

### Release Notes

<details>
<summary>grpc-ecosystem/go-grpc-middleware (github.com/grpc-ecosystem/go-grpc-middleware)</summary>

### [`v2.3.2`](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/releases/tag/v2.3.2)

[Compare Source](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/compare/v2.3.1...v2.3.2)

#### What's Changed

- chore: fix some typos by [@&#8203;dockercui](https://redirect.github.com/dockercui) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/755](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/755)
- update module for protovalidate by [@&#8203;calmh](https://redirect.github.com/calmh) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/757](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/757)

#### New Contributors

- [@&#8203;dockercui](https://redirect.github.com/dockercui) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/755](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/755)
- [@&#8203;calmh](https://redirect.github.com/calmh) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/757](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/757)

**Full Changelog**: https://github.com/grpc-ecosystem/go-grpc-middleware/compare/v2.3.1...v2.3.2

### [`v2.3.1`](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/releases/tag/v2.3.1)

[Compare Source](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/compare/v2.3.0...v2.3.1)

#### What's Changed

- Return uint64 from exponentBase2 function by [@&#8203;fesiqueira](https://redirect.github.com/fesiqueira) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/753](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/753)
- add build tag to disable tracing by [@&#8203;rashmi-tondare](https://redirect.github.com/rashmi-tondare) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/754](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/754)

#### New Contributors

- [@&#8203;fesiqueira](https://redirect.github.com/fesiqueira) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/753](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/753)
- [@&#8203;rashmi-tondare](https://redirect.github.com/rashmi-tondare) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/754](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/754)

**Full Changelog**: https://github.com/grpc-ecosystem/go-grpc-middleware/compare/v2.3.0...v2.3.1

### [`v2.3.0`](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/releases/tag/v2.3.0)

[Compare Source](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/compare/v2.2.0...v2.3.0)

#### What's Changed

- logging: add AddFields by [@&#8203;kindermoumoute](https://redirect.github.com/kindermoumoute) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/739](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/739)
- logging: store the propagated context in the reporter by [@&#8203;kindermoumoute](https://redirect.github.com/kindermoumoute) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/740](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/740)
- Add skip\_healthchecks logging example by [@&#8203;bkane-msft](https://redirect.github.com/bkane-msft) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/742](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/742)
- grpc\_retry backoff overflow by [@&#8203;JacobSMoller](https://redirect.github.com/JacobSMoller) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/747](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/747)
- defer cancel() leaks memory by [@&#8203;JacobSMoller](https://redirect.github.com/JacobSMoller) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/748](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/748)
- protovalidate: support new protovalidate-go Validator interface by [@&#8203;zchee](https://redirect.github.com/zchee) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/746](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/746)

#### New Contributors

- [@&#8203;bkane-msft](https://redirect.github.com/bkane-msft) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/742](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/742)
- [@&#8203;JacobSMoller](https://redirect.github.com/JacobSMoller) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/747](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/747)
- [@&#8203;zchee](https://redirect.github.com/zchee) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/746](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/746)

**Full Changelog**: https://github.com/grpc-ecosystem/go-grpc-middleware/compare/v2.2.0...v2.3.0

### [`v2.2.0`](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/releases/tag/v2.2.0)

[Compare Source](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/compare/v2.1.0...v2.2.0)

#### What's Changed

-   Call retry callback on retry by [@&#8203;fredr](https://redirect.github.com/fredr) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/700](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/700)
-   interceptors: Update logging interceptor Reporter to re-extract fields from context before logging by [@&#8203;chancez](https://redirect.github.com/chancez) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/702](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/702)
-   logging: Document correct WithFieldsFromContext/WithFieldsFromContextAndCallMeta usage by [@&#8203;chancez](https://redirect.github.com/chancez) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/703](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/703)
-   Include error details in protovalidate responses by [@&#8203;akshayjshah](https://redirect.github.com/akshayjshah) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/714](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/714)
-   protovalidate: avoid pointer comparisons by [@&#8203;akshayjshah](https://redirect.github.com/akshayjshah) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/715](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/715)
-   Support for namespace in grpc prometheus counter and histogram metrics by [@&#8203;hyungi](https://redirect.github.com/hyungi) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/718](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/718)
-   Protovalidate interceptor cleanup, Go version bump by [@&#8203;ash2k](https://redirect.github.com/ash2k) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/721](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/721)
-   Use ValueFromIncomingContext() to reduce allocations and copying by [@&#8203;ash2k](https://redirect.github.com/ash2k) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/723](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/723)
-   Update examples to the latest otelgrpc API by [@&#8203;nmittler](https://redirect.github.com/nmittler) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/729](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/729)
-   Fix grpc middleware interceptor not PostCall-ing when a streaming RPC with non-streaming server finishes successfully. by [@&#8203;alexandrupitis1](https://redirect.github.com/alexandrupitis1) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/725](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/725)
-   x-retry-attempt to StreamClientInterceptor by [@&#8203;Boklazhenko](https://redirect.github.com/Boklazhenko) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/733](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/733)
-   logging: add WithErrorFields by [@&#8203;kindermoumoute](https://redirect.github.com/kindermoumoute) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/734](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/734)
-   example: use slog instead of go-kit by [@&#8203;kindermoumoute](https://redirect.github.com/kindermoumoute) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/735](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/735)

#### New Contributors

-   [@&#8203;fredr](https://redirect.github.com/fredr) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/700](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/700)
-   [@&#8203;marefr](https://redirect.github.com/marefr) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/706](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/706)
-   [@&#8203;akshayjshah](https://redirect.github.com/akshayjshah) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/714](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/714)
-   [@&#8203;hyungi](https://redirect.github.com/hyungi) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/718](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/718)
-   [@&#8203;nmittler](https://redirect.github.com/nmittler) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/729](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/729)
-   [@&#8203;alexandrupitis1](https://redirect.github.com/alexandrupitis1) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/725](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/725)
-   [@&#8203;Boklazhenko](https://redirect.github.com/Boklazhenko) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/733](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/733)
-   [@&#8203;kindermoumoute](https://redirect.github.com/kindermoumoute) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/734](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/734)

**Full Changelog**: https://github.com/grpc-ecosystem/go-grpc-middleware/compare/v2.1.0...v2.2.0

### [`v2.1.0`](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/releases/tag/v2.1.0)

[Compare Source](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/compare/v2.0.1...v2.1.0)

#### What's Changed

-   Support for subsystem in grpc prometheus counter and histogram metrics by [@&#8203;rohsaini](https://redirect.github.com/rohsaini) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/643](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/643)
-   doc: update client interceptors chaining example with grpc functions by [@&#8203;dethi](https://redirect.github.com/dethi) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/669](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/669)
-   adds fields from durationFieldFunc to request/response log entries by [@&#8203;vroldanbet](https://redirect.github.com/vroldanbet) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/670](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/670)
-   add doc for disabling log opts by [@&#8203;coleenquadros](https://redirect.github.com/coleenquadros) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/680](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/680)
-   Middleware for determining the real ip of the client by [@&#8203;MadsRC](https://redirect.github.com/MadsRC) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/682](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/682)
-   protovalidate: add option to ignore certain message types by [@&#8203;igor-tsiglyar](https://redirect.github.com/igor-tsiglyar) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/684](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/684)
-   Update README.md by [@&#8203;zeroboo](https://redirect.github.com/zeroboo) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/688](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/688)
-   Fix `InitializeMetrics` signature to allow use with `xds.GRPCServer` by [@&#8203;bozaro](https://redirect.github.com/bozaro) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/689](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/689)
-   Support retriable func condition by [@&#8203;tamayika](https://redirect.github.com/tamayika) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/687](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/687)
-   Extend realip parsing of GRPC peer address to handle IPv6 by [@&#8203;surik](https://redirect.github.com/surik) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/692](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/692)
-   Fix logging Example : log only first field by [@&#8203;arckadious](https://redirect.github.com/arckadious) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/694](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/694)
-   Extent realip interceptors with ip selection based on proxy count and list by [@&#8203;surik](https://redirect.github.com/surik) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/695](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/695)
-   Fix for vulnerability CVE-2023-44487 by [@&#8203;vkaushik](https://redirect.github.com/vkaushik) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/696](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/696)

#### New Contributors

-   [@&#8203;rohsaini](https://redirect.github.com/rohsaini) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/643](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/643)
-   [@&#8203;dethi](https://redirect.github.com/dethi) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/669](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/669)
-   [@&#8203;vroldanbet](https://redirect.github.com/vroldanbet) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/670](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/670)
-   [@&#8203;MadsRC](https://redirect.github.com/MadsRC) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/682](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/682)
-   [@&#8203;igor-tsiglyar](https://redirect.github.com/igor-tsiglyar) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/684](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/684)
-   [@&#8203;zeroboo](https://redirect.github.com/zeroboo) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/688](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/688)
-   [@&#8203;bozaro](https://redirect.github.com/bozaro) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/689](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/689)
-   [@&#8203;tamayika](https://redirect.github.com/tamayika) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/687](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/687)
-   [@&#8203;surik](https://redirect.github.com/surik) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/692](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/692)
-   [@&#8203;arckadious](https://redirect.github.com/arckadious) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/694](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/694)
-   [@&#8203;vkaushik](https://redirect.github.com/vkaushik) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/696](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/696)

**Full Changelog**: https://github.com/grpc-ecosystem/go-grpc-middleware/compare/v2.0.1...v2.1.0

### [`v2.0.1`](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/releases/tag/v2.0.1)

[Compare Source](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/compare/v2.0.0...v2.0.1)

#### What's Changed

-   Fix outdated 'make proto' command by [@&#8203;takp](https://redirect.github.com/takp) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/623](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/623)
-   Fix linting errors by [@&#8203;takp](https://redirect.github.com/takp) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/624](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/624)
-   Logging: Add missing variadic operator for fields by [@&#8203;olivierlemasle](https://redirect.github.com/olivierlemasle) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/629](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/629)
-   feat: Support extracting fields from CallMeta by [@&#8203;fsaintjacques](https://redirect.github.com/fsaintjacques) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/628](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/628)
-   Fix "make test" and "make lint" by [@&#8203;olivierlemasle](https://redirect.github.com/olivierlemasle) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/627](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/627)
-   Do not set timeout for stream initialization by [@&#8203;DavyJohnes](https://redirect.github.com/DavyJohnes) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/645](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/645)
-   Add logging option to disable fields in log entry by [@&#8203;coleenquadros](https://redirect.github.com/coleenquadros) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/631](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/631)
-   Update logging adapter docs by [@&#8203;aboryslawski](https://redirect.github.com/aboryslawski) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/647](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/647)

#### New Contributors

-   [@&#8203;takp](https://redirect.github.com/takp) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/623](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/623)
-   [@&#8203;olivierlemasle](https://redirect.github.com/olivierlemasle) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/629](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/629)
-   [@&#8203;fsaintjacques](https://redirect.github.com/fsaintjacques) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/628](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/628)
-   [@&#8203;DavyJohnes](https://redirect.github.com/DavyJohnes) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/645](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/645)
-   [@&#8203;coleenquadros](https://redirect.github.com/coleenquadros) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/631](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/631)
-   [@&#8203;aboryslawski](https://redirect.github.com/aboryslawski) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/647](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/647)

**Full Changelog**: https://github.com/grpc-ecosystem/go-grpc-middleware/compare/v2.0.0...v2.0.1

### [`v2.0.0`](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/releases/tag/v2.0.0)

[Compare Source](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/compare/v1.4.0...v2.0.0)

This is the first stable release of the new v2 release branch 🎉

Many of the interceptors have been rewritten from scratch and the project has been upgraded to use the Go Protobuf v2 API.

See the project README for details and migration guide. Thanks to all contributors who made this possible! 💪🏽

#### What's Changed

-   Initial change for v2. by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/276](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/276)
-   Updated README with note that it's under development. by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/278](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/278)
-   Fix typo in field extractor (splices -> slices) ([#&#8203;287](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/issues/287)) by [@&#8203;bvwells](https://redirect.github.com/bvwells) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/289](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/289)
-   Moved imports to v2; Moved to Go 1.14.2 by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/290](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/290)
-   Formatted code; Added goimports to Makefile, Renamed pb_testproto to testpb. by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/291](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/291)
-   Fixed providers go modules, examples and consistency. by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/292](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/292)
-   added example for AuthFuncOverride v2 branch by [@&#8203;tegk](https://redirect.github.com/tegk) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/294](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/294)
-   Added some description of the Makefile in the contributing.md by [@&#8203;yashrsharma44](https://redirect.github.com/yashrsharma44) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/298](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/298)
-   v2: Add support for the zerolog logging provider by [@&#8203;irridia](https://redirect.github.com/irridia) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/299](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/299)
-   proto: fix gogoproto import by [@&#8203;johanbrandhorst](https://redirect.github.com/johanbrandhorst) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/302](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/302)
-   Retry dial and connection errors for grpc stream. by [@&#8203;kartlee](https://redirect.github.com/kartlee) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/308](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/308)
-   Moved to GH actions; Added lint; Added issue/PR templates. by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/296](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/296)
-   inline localhost certificate into go file by [@&#8203;bmon](https://redirect.github.com/bmon) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/318](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/318)
-   Update streaming interceptor example by [@&#8203;G07cha](https://redirect.github.com/G07cha) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/322](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/322)
-   Do not stop retrying based on earlier good message from the stream by [@&#8203;kartlee](https://redirect.github.com/kartlee) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/323](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/323)
-   test certs - cherry-pick PR325 on v2 by [@&#8203;dmitris](https://redirect.github.com/dmitris) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/331](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/331)
-   add all make target, reword instructions by [@&#8203;dmitris](https://redirect.github.com/dmitris) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/335](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/335)
-   remove 1.12.x from build config for consistency with master by [@&#8203;dmitris](https://redirect.github.com/dmitris) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/337](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/337)
-   \[v2] Fix the special case for jaeger format traceid extraction by [@&#8203;nvx](https://redirect.github.com/nvx) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/340](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/340)
-   \[v2] Fix ctxtags TagBasedRequestFieldExtractor extracting from fields in a oneof by [@&#8203;nvx](https://redirect.github.com/nvx) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/339](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/339)
-   Request Logging by [@&#8203;yashrsharma44](https://redirect.github.com/yashrsharma44) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/311](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/311)
-   Bug fix for data race by [@&#8203;yashrsharma44](https://redirect.github.com/yashrsharma44) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/354](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/354)
-   make ratelimit interface context aware by [@&#8203;xinxiao](https://redirect.github.com/xinxiao) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/367](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/367)
-   Add error param to the decider method of logging middleware by [@&#8203;yashrsharma44](https://redirect.github.com/yashrsharma44) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/372](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/372)
-   \[v2] Add skip interceptor by [@&#8203;XSAM](https://redirect.github.com/XSAM) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/364](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/364)
-   Chain middleware by [@&#8203;drewwells](https://redirect.github.com/drewwells) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/385](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/385)
-   Update travis ci badget to Github actions badge. by [@&#8203;drewwells](https://redirect.github.com/drewwells) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/384](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/384)
-   Upgraded proto related deps: grpc and protobuf; removed gogo from core. by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/321](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/321)
-   improve v2 rate-limiter by [@&#8203;MalloZup](https://redirect.github.com/MalloZup) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/380](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/380)
-   Moved to buf; Added buf lint; Fixed ping service to match standards; … by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/383](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/383)
-   Add timer interface for OpenMetrics(Prometheus) Provider by [@&#8203;yashrsharma44](https://redirect.github.com/yashrsharma44) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/387](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/387)
-   \[Rate-limit provider]: Add token bucket implementation of rate-limiter by [@&#8203;MalloZup](https://redirect.github.com/MalloZup) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/386](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/386)
-   Add OpenMetrics(Prometheus) in the provider module by [@&#8203;yashrsharma44](https://redirect.github.com/yashrsharma44) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/379](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/379)
-   v2: Client unary interceptor timeout on v2 branch by [@&#8203;instabledesign](https://redirect.github.com/instabledesign) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/330](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/330)
-   add onRetryCallback callback function by [@&#8203;shamil](https://redirect.github.com/shamil) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/405](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/405)
-   v2: validator support for protoc-gen-validate 0.6.0 by [@&#8203;danielhochman](https://redirect.github.com/danielhochman) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/418](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/418)
-   v2: Refactor metrics interceptor and fix tests by [@&#8203;ash2k](https://redirect.github.com/ash2k) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/413](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/413)
-   Support customization of timestamp format (v2 branch) by [@&#8203;stanhu](https://redirect.github.com/stanhu) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/399](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/399)
-   Fixed misleading comments in the interceptor file by [@&#8203;iamrajiv](https://redirect.github.com/iamrajiv) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/424](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/424)
-   v2: Switch from github.com/go-kit/kit to github.com/go-kit/log interfaces by [@&#8203;liggitt](https://redirect.github.com/liggitt) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/427](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/427)
-   v2: Add support for the phuslog logging provider by [@&#8203;ogimenezb](https://redirect.github.com/ogimenezb) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/425](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/425)
-   v2:providers/zap: fix caller annotation by [@&#8203;jkawamoto](https://redirect.github.com/jkawamoto) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/432](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/432)
-   Added Dependabot by [@&#8203;iamrajiv](https://redirect.github.com/iamrajiv) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/376](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/376)
-   Added a Copyright check in the Makefile by [@&#8203;yashrsharma44](https://redirect.github.com/yashrsharma44) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/420](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/420)
-   Cleanup v2 with some updates by [@&#8203;yashrsharma44](https://redirect.github.com/yashrsharma44) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/419](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/419)
-   openmetrics: forward server context by [@&#8203;amenzhinsky](https://redirect.github.com/amenzhinsky) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/434](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/434)
-   recovery: change the default behavior by [@&#8203;amenzhinsky](https://redirect.github.com/amenzhinsky) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/439](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/439)
-   Add all-validator support by [@&#8203;leventeliu](https://redirect.github.com/leventeliu) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/443](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/443)
-   Remove backoffutils and added the files to retry package by [@&#8203;yashrsharma44](https://redirect.github.com/yashrsharma44) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/390](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/390)
-   v2:interceptors/logging: allow to separate request response payload logging by [@&#8203;michaljemala](https://redirect.github.com/michaljemala) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/442](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/442)
-   Removed tags; Simplified interceptor code; Added logging fields editability. by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/394](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/394)
-   InterceptorTestSuite client connection optimize by [@&#8203;HUSTtoKTH](https://redirect.github.com/HUSTtoKTH) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/455](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/455)
-   Remove opentracing from go.mod by [@&#8203;jpkrohling](https://redirect.github.com/jpkrohling) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/477](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/477)
-   Replace two old Go versions with two new ones by [@&#8203;jpkrohling](https://redirect.github.com/jpkrohling) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/478](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/478)
-   Move util/metautils to root-level package metadata, fixes [#&#8203;392](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/issues/392) by [@&#8203;rahulkhairwar](https://redirect.github.com/rahulkhairwar) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/474](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/474)
-   Remove data race from zerolog provider by [@&#8203;ecordell](https://redirect.github.com/ecordell) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/487](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/487)
-   Update provider/kit by [@&#8203;metalmatze](https://redirect.github.com/metalmatze) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/490](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/490)
-   Refactor tracing interceptor by [@&#8203;XSAM](https://redirect.github.com/XSAM) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/450](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/450)
-   Add opentracing provider by [@&#8203;XSAM](https://redirect.github.com/XSAM) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/492](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/492)
-   update phuslog to fix typo by [@&#8203;ogimenezb](https://redirect.github.com/ogimenezb) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/499](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/499)
-   Added logr as logging Library provider by [@&#8203;mcdoker18](https://redirect.github.com/mcdoker18) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/510](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/510)
-   ✨ new feat: selector middleware by [@&#8203;aimuz](https://redirect.github.com/aimuz) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/511](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/511)
-   Add 1.18.x and 1.19.x unit tests by [@&#8203;aimuz](https://redirect.github.com/aimuz) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/513](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/513)
-   change the doc.go to the latest format by [@&#8203;aimuz](https://redirect.github.com/aimuz) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/512](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/512)
-   fix provider examples by [@&#8203;forsaken628](https://redirect.github.com/forsaken628) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/529](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/529)
-   docs: add `logging.InjectFields` usage description by [@&#8203;aimuz](https://redirect.github.com/aimuz) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/541](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/541)
-   Bump golang.org/x/net from 0.0.0-20201021035429-f5854403a974 to 0.7.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/537](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/537)
-   v2: All for v2: Exemplars, Cleanup, Docs, Lint, Proto upgrades and more by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/543](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/543)
-   \[interceptors/validator] feat: add error logging in validator by [@&#8203;rohanraj7316](https://redirect.github.com/rohanraj7316) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/544](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/544)
-   fix auto-generated docs by [@&#8203;peczenyj](https://redirect.github.com/peczenyj) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/548](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/548)
-   fix vulnerability GO-2022-0603 by [@&#8203;peczenyj](https://redirect.github.com/peczenyj) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/549](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/549)
-   add support to trace on grpc_logrus.DefaultMessageProducer by [@&#8203;peczenyj](https://redirect.github.com/peczenyj) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/547](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/547)
-   Simplified logging middleware; Fields are now "any" type; Moved logging providers to examples only. by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/552](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/552)
-   Removed deciders; Cleaned up validators. by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/554](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/554)
-   Adjustments to README and consistency of callback options. by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/555](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/555)
-   Merge v2 into main (with -X theirs) by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/556](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/556)
-   Bump golang.org/x/net from 0.5.0 to 0.7.0 in /providers/prometheus by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/561](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/561)
-   Fix overwritten logger in zerolog example by [@&#8203;longshine](https://redirect.github.com/longshine) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/574](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/574)
-   Changed for for mapping fields, different var for logger in zap example by [@&#8203;MichalFikejs](https://redirect.github.com/MichalFikejs) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/581](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/581)
-   cleanup: no cap definition required by [@&#8203;aimuz](https://redirect.github.com/aimuz) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/582](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/582)
-   providers/prometheus: Add WithHistogramOpts for native histograms by [@&#8203;metalmatze](https://redirect.github.com/metalmatze) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/584](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/584)
-   fix: Refactor logger initialization in example_test.go by [@&#8203;aimuz](https://redirect.github.com/aimuz) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/580](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/580)
-   Minor code cleanups by [@&#8203;ash2k](https://redirect.github.com/ash2k) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/586](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/586)
-   fix prometheus interceptors not converting context errors to gRPC codes by [@&#8203;vtermanis](https://redirect.github.com/vtermanis) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/571](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/571)
-   Update README.md by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/600](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/600)
-   Update PULL_REQUEST_TEMPLATE.md by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/601](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/601)
-   Add Client rate limit interceptors 520 by [@&#8203;rahulkhairwar](https://redirect.github.com/rahulkhairwar) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/599](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/599)
-   Use default go errors package instead of github.com/pkg/errors by [@&#8203;rifkyazizf](https://redirect.github.com/rifkyazizf) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/608](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/608)
-   fix bug : should drain channel of timer after stop by [@&#8203;ikenchina](https://redirect.github.com/ikenchina) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/612](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/612)
-   feat: add interceptor for bufbuild/protovalidate by [@&#8203;gvencadze](https://redirect.github.com/gvencadze) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/614](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/614)
-   Enhancement: Introduce Option Interface for Future Interceptor Customization by [@&#8203;elliotmjackson](https://redirect.github.com/elliotmjackson) in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/615](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/615)

#### New Contributors

-   [@&#8203;irridia](https://redirect.github.com/irridia) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/299](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/299)
-   [@&#8203;xinxiao](https://redirect.github.com/xinxiao) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/367](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/367)
-   [@&#8203;XSAM](https://redirect.github.com/XSAM) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/364](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/364)
-   [@&#8203;MalloZup](https://redirect.github.com/MalloZup) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/380](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/380)
-   [@&#8203;instabledesign](https://redirect.github.com/instabledesign) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/330](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/330)
-   [@&#8203;shamil](https://redirect.github.com/shamil) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/405](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/405)
-   [@&#8203;ash2k](https://redirect.github.com/ash2k) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/413](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/413)
-   [@&#8203;ogimenezb](https://redirect.github.com/ogimenezb) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/425](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/425)
-   [@&#8203;amenzhinsky](https://redirect.github.com/amenzhinsky) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/434](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/434)
-   [@&#8203;leventeliu](https://redirect.github.com/leventeliu) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/443](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/443)
-   [@&#8203;michaljemala](https://redirect.github.com/michaljemala) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/442](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/442)
-   [@&#8203;HUSTtoKTH](https://redirect.github.com/HUSTtoKTH) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/455](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/455)
-   [@&#8203;jpkrohling](https://redirect.github.com/jpkrohling) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/477](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/477)
-   [@&#8203;rahulkhairwar](https://redirect.github.com/rahulkhairwar) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/474](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/474)
-   [@&#8203;ecordell](https://redirect.github.com/ecordell) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/487](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/487)
-   [@&#8203;metalmatze](https://redirect.github.com/metalmatze) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/490](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/490)
-   [@&#8203;mcdoker18](https://redirect.github.com/mcdoker18) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/510](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/510)
-   [@&#8203;aimuz](https://redirect.github.com/aimuz) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/511](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/511)
-   [@&#8203;forsaken628](https://redirect.github.com/forsaken628) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/529](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/529)
-   [@&#8203;dependabot](https://redirect.github.com/dependabot) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/537](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/537)
-   [@&#8203;rohanraj7316](https://redirect.github.com/rohanraj7316) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/544](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/544)
-   [@&#8203;peczenyj](https://redirect.github.com/peczenyj) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/548](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/548)
-   [@&#8203;longshine](https://redirect.github.com/longshine) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/574](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/574)
-   [@&#8203;MichalFikejs](https://redirect.github.com/MichalFikejs) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/581](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/581)
-   [@&#8203;vtermanis](https://redirect.github.com/vtermanis) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/571](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/571)
-   [@&#8203;rifkyazizf](https://redirect.github.com/rifkyazizf) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/608](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/608)
-   [@&#8203;ikenchina](https://redirect.github.com/ikenchina) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/612](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/612)
-   [@&#8203;gvencadze](https://redirect.github.com/gvencadze) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/614](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/614)
-   [@&#8203;elliotmjackson](https://redirect.github.com/elliotmjackson) made their first contribution in [https://github.com/grpc-ecosystem/go-grpc-middleware/pull/615](https://redirect.github.com/grpc-ecosystem/go-grpc-middleware/pull/615)

**Full Changelog**: https://github.com/grpc-ecosystem/go-grpc-middleware/compare/v1.4.0...v2.0.0

</details>

---

### Configuration

📅 **Schedule**: Branch creation - "after 5am on sunday" in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Enabled.

♻ **Rebasing**: Whenever PR is behind base branch, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS43LjAtcnBtIiwidXBkYXRlZEluVmVyIjoiNDEuNy4wLXJwbSIsInRhcmdldEJyYW5jaCI6Im1hc3RlciIsImxhYmVscyI6W119-->


---

## Discussion

### Comment by @jira-linking on July 27, 2025 at 01:50 PM UTC

Commits missing Jira IDs:
2104fc5d63b4576654142491d7d3a6f0fd0093a2


### Comment by @red-hat-konflux on July 28, 2025 at 10:34 AM UTC

### Renovate Ignore Notification

Because you closed this PR without merging, Renovate will ignore this update. You will not get PRs for *any* future `2.x` releases. But if you manually upgrade to `2.x` then Renovate will re-enable `minor` and `patch` updates automatically.

If you accidentally closed this PR, or if you changed your mind: rename this PR to get a fresh replacement PR.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1753*
