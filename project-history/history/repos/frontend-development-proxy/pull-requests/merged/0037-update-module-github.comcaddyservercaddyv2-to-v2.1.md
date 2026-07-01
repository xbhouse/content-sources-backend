---
type: pull_request
number: 37
title: "Update module github.com/caddyserver/caddy/v2 to v2.10.2"
state: merged
author: red-hat-konflux
created: 2025-10-24T20:15:53Z
updated: 2025-11-24T11:36:35Z
closed: 2025-11-24T11:35:16Z
merged: 2025-11-24T11:35:16Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-caddyserver-caddy-v2-2.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/37
---

# Pull Request #37: Update module github.com/caddyserver/caddy/v2 to v2.10.2

**Author**: @red-hat-konflux
**Created**: October 24, 2025 at 08:15 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-caddyserver-caddy-v2-2.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/caddyserver/caddy/v2](https://redirect.github.com/caddyserver/caddy) | `v2.9.1` -> `v2.10.2` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fcaddyserver%2fcaddy%2fv2/v2.10.2?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fcaddyserver%2fcaddy%2fv2/v2.9.1/v2.10.2?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>caddyserver/caddy (github.com/caddyserver/caddy/v2)</summary>

### [`v2.10.2`](https://redirect.github.com/caddyserver/caddy/releases/tag/v2.10.2)

[Compare Source](https://redirect.github.com/caddyserver/caddy/compare/v2.10.1...v2.10.2)

This is a hotfix release to fix a couple critical issues from v2.10.1

#### What's Changed

- http: Make logger first, before TLS provisioning by [@&#8203;francislavoie](https://redirect.github.com/francislavoie) in [#&#8203;7198](https://redirect.github.com/caddyserver/caddy/pull/7198)
- httpcaddyfile: Fix `acme_dns` regression by [@&#8203;francislavoie](https://redirect.github.com/francislavoie) in [#&#8203;7199](https://redirect.github.com/caddyserver/caddy/pull/7199)
- caddyfile: Fix importing nested tokens for `{block}` by [@&#8203;BeeJay28](https://redirect.github.com/BeeJay28) in [#&#8203;7189](https://redirect.github.com/caddyserver/caddy/pull/7189)

#### Changelog

- [`551f793`](https://redirect.github.com/caddyserver/caddy/commit/551f793700fe1550845c824470b623fd1aa03d36) caddyfile: Fix importing nested tokens for `{block}` ([#&#8203;7189](https://redirect.github.com/caddyserver/caddy/issues/7189))
- [`16fe83c`](https://redirect.github.com/caddyserver/caddy/commit/16fe83c7afe2152b0bb53ae35078a28f87e6dcf2) http: Make logger first, before TLS provisioning ([#&#8203;7198](https://redirect.github.com/caddyserver/caddy/issues/7198))
- [`4564261`](https://redirect.github.com/caddyserver/caddy/commit/4564261d8350f8010b7e001e646e260e9bba5746) httpcaddyfile: Fix `acme_dns` regression ([#&#8203;7199](https://redirect.github.com/caddyserver/caddy/issues/7199))

#### New Contributors

- [@&#8203;BeeJay28](https://redirect.github.com/BeeJay28) made their first contribution in [#&#8203;7189](https://redirect.github.com/caddyserver/caddy/pull/7189)

**Full Changelog**: <https://github.com/caddyserver/caddy/compare/v2.10.1...v2.10.2>

### [`v2.10.1`](https://redirect.github.com/caddyserver/caddy/releases/tag/v2.10.1)

[Compare Source](https://redirect.github.com/caddyserver/caddy/compare/v2.10.0...v2.10.1)

This is probably our biggest patch release ever -- not that lots of things were broken, but there's lots of refinement happening thanks to broader adoption and contributions from many more people. Just look at the New Contributors below!

Anyway, this release does contain some bug fixes and dependency upgrades which we hope will serve you well. Let us know if there's any issues! And thank you to all who contributed, especially our reliable maintainer team!

This version of Caddy requires [Go v1.25.0 or newer](https://golang.org/dl/).

#### What's Changed

- update quic-go to v0.51.0 by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;6972](https://redirect.github.com/caddyserver/caddy/pull/6972)
- forwardproxy: reference correct field name in LoadModule by [@&#8203;mohammed90](https://redirect.github.com/mohammed90) in [#&#8203;6978](https://redirect.github.com/caddyserver/caddy/pull/6978)
- fix: Remove `nil` arg from `zapslog.NewHandler` call by [@&#8203;IndraGunawan](https://redirect.github.com/IndraGunawan) in [#&#8203;6984](https://redirect.github.com/caddyserver/caddy/pull/6984)
- fileserver: Add support for .avif image format by [@&#8203;steffenbusch](https://redirect.github.com/steffenbusch) in [#&#8203;6988](https://redirect.github.com/caddyserver/caddy/pull/6988)
- reverseproxy: use DialTLSContext for TLS if servername has placeholder by [@&#8203;WeidiDeng](https://redirect.github.com/WeidiDeng) in [#&#8203;6955](https://redirect.github.com/caddyserver/caddy/pull/6955)
- admin: Make sure that any admin routers are provisioned when local/re… by [@&#8203;Compy](https://redirect.github.com/Compy) in [#&#8203;6997](https://redirect.github.com/caddyserver/caddy/pull/6997)
- log: default logger should respect `{in,ex}clude` by [@&#8203;mohammed90](https://redirect.github.com/mohammed90) in [#&#8203;6995](https://redirect.github.com/caddyserver/caddy/pull/6995)
- Move local admin server replacement logic below data structure initia… by [@&#8203;Compy](https://redirect.github.com/Compy) in [#&#8203;7004](https://redirect.github.com/caddyserver/caddy/pull/7004)
- acme\_server: fix policy parsing in caddyfile by [@&#8203;mohammed90](https://redirect.github.com/mohammed90) in [#&#8203;7006](https://redirect.github.com/caddyserver/caddy/pull/7006)
- implement Unwrap for interceptedResponseHandler by [@&#8203;WeidiDeng](https://redirect.github.com/WeidiDeng) in [#&#8203;7016](https://redirect.github.com/caddyserver/caddy/pull/7016)
- fileserver: map invalid path errors to fs.ErrInvalid, and return 400 … by [@&#8203;Compy](https://redirect.github.com/Compy) in [#&#8203;7017](https://redirect.github.com/caddyserver/caddy/pull/7017)
- caddyhttp: fix route sort by comparing paths without wildcard if they don't shar… by [@&#8203;WeidiDeng](https://redirect.github.com/WeidiDeng) in [#&#8203;7015](https://redirect.github.com/caddyserver/caddy/pull/7015)
- refactor: use maps.Copy for cleaner map handling by [@&#8203;eveneast](https://redirect.github.com/eveneast) in [#&#8203;7009](https://redirect.github.com/caddyserver/caddy/pull/7009)
- refactor: use slices.Contains to simplify code by [@&#8203;tongjicoder](https://redirect.github.com/tongjicoder) in [#&#8203;7039](https://redirect.github.com/caddyserver/caddy/pull/7039)
- chore: upgrade .golangci.yml and workflow to v2 by [@&#8203;mohammed90](https://redirect.github.com/mohammed90) in [#&#8203;6924](https://redirect.github.com/caddyserver/caddy/pull/6924)
- build(deps): bump golangci/golangci-lint-action from 6 to 8 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;7044](https://redirect.github.com/caddyserver/caddy/pull/7044)
- fix: crash - null check on event origin by [@&#8203;suxatcode](https://redirect.github.com/suxatcode) in [#&#8203;7047](https://redirect.github.com/caddyserver/caddy/pull/7047)
- fix: prevent error handler from overriding sub handler matchers by [@&#8203;Hellio404](https://redirect.github.com/Hellio404) in [#&#8203;6999](https://redirect.github.com/caddyserver/caddy/pull/6999)
- client\_auth: wire up leaf verifier Caddyfile by [@&#8203;mohammed90](https://redirect.github.com/mohammed90) in [#&#8203;6772](https://redirect.github.com/caddyserver/caddy/pull/6772)
- caddyfile: reject blocks in log\_skip directive by [@&#8203;IwatsukaYura](https://redirect.github.com/IwatsukaYura) in [#&#8203;7056](https://redirect.github.com/caddyserver/caddy/pull/7056)
- build(deps): bump github.com/cloudflare/circl from 1.6.0 to 1.6.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;7058](https://redirect.github.com/caddyserver/caddy/pull/7058)
- cmd: fix `Commands` function not returning all registered commands by [@&#8203;hslatman](https://redirect.github.com/hslatman) in [#&#8203;7059](https://redirect.github.com/caddyserver/caddy/pull/7059)
- ci: add dep review, OSSF scorecard actions by [@&#8203;mohammed90](https://redirect.github.com/mohammed90) in [#&#8203;7063](https://redirect.github.com/caddyserver/caddy/pull/7063)
- ci: add `{base,head}-ref` to  dep review check by [@&#8203;mohammed90](https://redirect.github.com/mohammed90) in [#&#8203;7064](https://redirect.github.com/caddyserver/caddy/pull/7064)
- core: clean up new config if it failed to run by [@&#8203;WeidiDeng](https://redirect.github.com/WeidiDeng) in [#&#8203;7068](https://redirect.github.com/caddyserver/caddy/pull/7068)
- chore: apply security best practices for CI by [@&#8203;mohammed90](https://redirect.github.com/mohammed90) in [#&#8203;7066](https://redirect.github.com/caddyserver/caddy/pull/7066)
- refactor: use the built-in max/min to simplify the code by [@&#8203;xiaoxiangirl](https://redirect.github.com/xiaoxiangirl) in [#&#8203;7081](https://redirect.github.com/caddyserver/caddy/pull/7081)
- \[ADD] sort buttons in grid mode by [@&#8203;filipRatajczak](https://redirect.github.com/filipRatajczak) in [#&#8203;7089](https://redirect.github.com/caddyserver/caddy/pull/7089)
- update quic-go to v0.53.0 by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;7094](https://redirect.github.com/caddyserver/caddy/pull/7094)
- refactor: replace HasPrefix+TrimPrefix with CutPrefix by [@&#8203;gopherorg](https://redirect.github.com/gopherorg) in [#&#8203;7095](https://redirect.github.com/caddyserver/caddy/pull/7095)
- docs: fix some minor issues in the comments by [@&#8203;mountdisk](https://redirect.github.com/mountdisk) in [#&#8203;7101](https://redirect.github.com/caddyserver/caddy/pull/7101)
- httpcaddyfile: Validates TLS DNS challenge options by [@&#8203;francislavoie](https://redirect.github.com/francislavoie) in [#&#8203;7099](https://redirect.github.com/caddyserver/caddy/pull/7099)
- chore: fix struct name in comment by [@&#8203;bytetigers](https://redirect.github.com/bytetigers) in [#&#8203;7114](https://redirect.github.com/caddyserver/caddy/pull/7114)
- reverse proxy: validate versions in http transport by [@&#8203;WeidiDeng](https://redirect.github.com/WeidiDeng) in [#&#8203;7112](https://redirect.github.com/caddyserver/caddy/pull/7112)
- chore: fix function in comment by [@&#8203;bytesingsong](https://redirect.github.com/bytesingsong) in [#&#8203;7121](https://redirect.github.com/caddyserver/caddy/pull/7121)
- Fix: Support placeholders in header replacement search patterns by [@&#8203;zongzewu23](https://redirect.github.com/zongzewu23) in [#&#8203;7117](https://redirect.github.com/caddyserver/caddy/pull/7117)
- fileserver: specify license for embedded JavaScript by [@&#8203;infertux](https://redirect.github.com/infertux) in [#&#8203;7127](https://redirect.github.com/caddyserver/caddy/pull/7127)
- fix dead link by [@&#8203;eeemmmmmm](https://redirect.github.com/eeemmmmmm) in [#&#8203;7136](https://redirect.github.com/caddyserver/caddy/pull/7136)
- update quic-go to v0.54.0 by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;7138](https://redirect.github.com/caddyserver/caddy/pull/7138)
- chore: fix minor issue in comment by [@&#8203;pingshuijie](https://redirect.github.com/pingshuijie) in [#&#8203;7140](https://redirect.github.com/caddyserver/caddy/pull/7140)
- refactor: use slices.Equal to simplify code by [@&#8203;minxinyi](https://redirect.github.com/minxinyi) in [#&#8203;7141](https://redirect.github.com/caddyserver/caddy/pull/7141)
- ci: reduce dependabot spam by [@&#8203;mohammed90](https://redirect.github.com/mohammed90) in [#&#8203;7078](https://redirect.github.com/caddyserver/caddy/pull/7078)
- fix(provisioning): `Context.App` or `Context.AppIfConfigured` will return `(val, nil)` even if the app failed to provision or validate the first time by [@&#8203;alexandre-daubois](https://redirect.github.com/alexandre-daubois) in [#&#8203;7070](https://redirect.github.com/caddyserver/caddy/pull/7070)
- build(deps): bump the actions-deps group with 6 updates by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;7142](https://redirect.github.com/caddyserver/caddy/pull/7142)
- Use KeepAliveConfig to pass keepalive\_interval to listener's accepted sockets by [@&#8203;joshuamcbeth](https://redirect.github.com/joshuamcbeth) in [#&#8203;7151](https://redirect.github.com/caddyserver/caddy/pull/7151)
- build(deps): bump the all-updates group across 1 directory with 17 updates by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;7155](https://redirect.github.com/caddyserver/caddy/pull/7155)
- cmd: Allow `caddy adapt` to read from stdin by [@&#8203;bosdhill](https://redirect.github.com/bosdhill) in [#&#8203;7163](https://redirect.github.com/caddyserver/caddy/pull/7163)
- feat: add bcrypt cost parameter to hash-password by [@&#8203;GreyXor](https://redirect.github.com/GreyXor) in [#&#8203;7149](https://redirect.github.com/caddyserver/caddy/pull/7149)
- fix typo in bcrypt cost flag name by [@&#8203;GreyXor](https://redirect.github.com/GreyXor) in [#&#8203;7168](https://redirect.github.com/caddyserver/caddy/pull/7168)
- chore: fix inconsistent function name in comment by [@&#8203;youzichuan](https://redirect.github.com/youzichuan) in [#&#8203;7174](https://redirect.github.com/caddyserver/caddy/pull/7174)
- caddytls: fix regression in external certificate manager support by [@&#8203;quagsirus](https://redirect.github.com/quagsirus) in [#&#8203;7179](https://redirect.github.com/caddyserver/caddy/pull/7179)
- http: free up quic listener when stopping by [@&#8203;WeidiDeng](https://redirect.github.com/WeidiDeng) in [#&#8203;7177](https://redirect.github.com/caddyserver/caddy/pull/7177)
- chore: bump Go to v1.25 by [@&#8203;dunglas](https://redirect.github.com/dunglas) in [#&#8203;7184](https://redirect.github.com/caddyserver/caddy/pull/7184)
- caddyhttp: refactor to use reflect.TypeFor by [@&#8203;cuiweixie](https://redirect.github.com/cuiweixie) in [#&#8203;7187](https://redirect.github.com/caddyserver/caddy/pull/7187)
- refactor: use a more modern writing style to simplify code by [@&#8203;joemicky](https://redirect.github.com/joemicky) in [#&#8203;7182](https://redirect.github.com/caddyserver/caddy/pull/7182)
- http: disable keepalive when KeepAliveInterval is negative by [@&#8203;WeidiDeng](https://redirect.github.com/WeidiDeng) in [#&#8203;7158](https://redirect.github.com/caddyserver/caddy/pull/7158)
- http: clean up listeners if some of the listeners fail to bind by [@&#8203;WeidiDeng](https://redirect.github.com/WeidiDeng) in [#&#8203;7176](https://redirect.github.com/caddyserver/caddy/pull/7176)
- reverse\_proxy: use the new KeepAliveConfig to set probe interval by [@&#8203;WeidiDeng](https://redirect.github.com/WeidiDeng) in [#&#8203;7157](https://redirect.github.com/caddyserver/caddy/pull/7157)
- ci: set proper build tags in golangci and minor cleanup by [@&#8203;dunglas](https://redirect.github.com/dunglas) in [#&#8203;7183](https://redirect.github.com/caddyserver/caddy/pull/7183)
- doc: Add a few lines about Etag file content by [@&#8203;Pizmovc](https://redirect.github.com/Pizmovc) in [#&#8203;7173](https://redirect.github.com/caddyserver/caddy/pull/7173)
- file\_server: set Range header for precompressed static files to force Content Length header to appear by [@&#8203;WeidiDeng](https://redirect.github.com/WeidiDeng) in [#&#8203;7042](https://redirect.github.com/caddyserver/caddy/pull/7042)
- caddyhttp: use the new http.Protocols to handle h1, h2 and h2c requests by [@&#8203;WeidiDeng](https://redirect.github.com/WeidiDeng) in [#&#8203;6961](https://redirect.github.com/caddyserver/caddy/pull/6961)

#### Changelog

- [`44d078b`](https://redirect.github.com/caddyserver/caddy/commit/44d078b6705c7abcabb2a60f501568ff7f5a57a1) acme\_server: fix policy parsing in caddyfile ([#&#8203;7006](https://redirect.github.com/caddyserver/caddy/issues/7006))
- [`320c572`](https://redirect.github.com/caddyserver/caddy/commit/320c57291dbe06e00e0759bdb5cbbf0d622e5968) admin: Make sure that any admin routers are provisioned when local/re… ([#&#8203;6997](https://redirect.github.com/caddyserver/caddy/issues/6997))
- [`49dac61`](https://redirect.github.com/caddyserver/caddy/commit/49dac61b078a704b3e98566134c108d6def19450) bcrypt: add cost parameter to hash-password ([#&#8203;7149](https://redirect.github.com/caddyserver/caddy/issues/7149))
- [`4bfc3b9`](https://redirect.github.com/caddyserver/caddy/commit/4bfc3b95b5f88a1042a5103d8ad3fac3f42bf129) bcrypt: wrong cost flag name ([#&#8203;7168](https://redirect.github.com/caddyserver/caddy/issues/7168))
- [`4b01d77`](https://redirect.github.com/caddyserver/caddy/commit/4b01d77b81a9ebda046637026cf57671a3cc5859) build(deps): bump github.com/cloudflare/circl from 1.6.0 to 1.6.1 ([#&#8203;7058](https://redirect.github.com/caddyserver/caddy/issues/7058))
- [`45c9341`](https://redirect.github.com/caddyserver/caddy/commit/45c9341deb9818638ab389b98e7b4c74dc9f6afc) build(deps): bump golangci/golangci-lint-action from 6 to 8 ([#&#8203;7044](https://redirect.github.com/caddyserver/caddy/issues/7044))
- [`5bc2afb`](https://redirect.github.com/caddyserver/caddy/commit/5bc2afbbb6ee5bc1537b521fd507506fd86ae43f) build(deps): bump the actions-deps group with 6 updates ([#&#8203;7142](https://redirect.github.com/caddyserver/caddy/issues/7142))
- [`007f406`](https://redirect.github.com/caddyserver/caddy/commit/007f4066f6abf55f42195155117c53b230490a73) build(deps): bump the all-updates group across 1 directory with 17 updates ([#&#8203;7155](https://redirect.github.com/caddyserver/caddy/issues/7155))
- [`8524386`](https://redirect.github.com/caddyserver/caddy/commit/8524386737e7dcaf6ab2378e5bc9456f82b91cd1) caddyhttp: Compare paths w/o wildcard if prefixes differ ([#&#8203;7015](https://redirect.github.com/caddyserver/caddy/issues/7015))
- [`7590c9c`](https://redirect.github.com/caddyserver/caddy/commit/7590c9ca1ba6096408574a7dd3b9dc42d12cb948) caddyhttp: Free up quic listener when stopping ([#&#8203;7177](https://redirect.github.com/caddyserver/caddy/issues/7177))
- [`b15ed9b`](https://redirect.github.com/caddyserver/caddy/commit/b15ed9b0844dd7b73977f4c6dacfb3348579ce4a) caddyhttp: refactor to use reflect.TypeFor ([#&#8203;7187](https://redirect.github.com/caddyserver/caddy/issues/7187))
- [`14a63a2`](https://redirect.github.com/caddyserver/caddy/commit/14a63a26b9a673857fc37cba37aedc35a10ce634) caddyhttp: use the new http.Protocols to handle h1, h2 and h2c requests ([#&#8203;6961](https://redirect.github.com/caddyserver/caddy/issues/6961))
- [`731e6c2`](https://redirect.github.com/caddyserver/caddy/commit/731e6c24821630f7ac117ff74188e79a06a75d1e) caddytls: Improve ECH error logging (close [#&#8203;7152](https://redirect.github.com/caddyserver/caddy/issues/7152))
- [`105eee6`](https://redirect.github.com/caddyserver/caddy/commit/105eee671c384459de666889be953857a7175afa) caddytls: Set local\_ip, not remote\_ip ([#&#8203;6952](https://redirect.github.com/caddyserver/caddy/issues/6952))
- [`b898873`](https://redirect.github.com/caddyserver/caddy/commit/b898873b90b5cb804400888ec2f2994f6a2dd270) caddytls: fix regression in external certificate manager support ([#&#8203;7179](https://redirect.github.com/caddyserver/caddy/issues/7179))
- [`1481c04`](https://redirect.github.com/caddyserver/caddy/commit/1481c0411aa3ce3a53c206e18ee9ce4223cc156d) caddytls: wire up client\_auth leaf verifier Caddyfile ([#&#8203;6772](https://redirect.github.com/caddyserver/caddy/issues/6772))
- [`19ff47a`](https://redirect.github.com/caddyserver/caddy/commit/19ff47a63b9ff2ae424790b02548d9ba4afc56ba) cmd: Allow `caddy adapt` to read from stdin ([#&#8203;7163](https://redirect.github.com/caddyserver/caddy/issues/7163))
- [`e633d01`](https://redirect.github.com/caddyserver/caddy/commit/e633d013f64d057f462786ccf4a430cd27817d4d) cmd: fix `Commands` function not returning all registered commands ([#&#8203;7059](https://redirect.github.com/caddyserver/caddy/issues/7059))
- [`7099892`](https://redirect.github.com/caddyserver/caddy/commit/7099892958fbee17f5e087c768f4db1940303fa8) core: Check for nil event origin ([#&#8203;7047](https://redirect.github.com/caddyserver/caddy/issues/7047))
- [`3d0b4fa`](https://redirect.github.com/caddyserver/caddy/commit/3d0b4fac5a583d615fe411a4af9a24a7bcc3bee3) core: Clean up new config if it failed to run ([#&#8203;7068](https://redirect.github.com/caddyserver/caddy/issues/7068))
- [`051e73a`](https://redirect.github.com/caddyserver/caddy/commit/051e73aefca4cc3d930e8b637d848deb5e100126) core: Replace admin server later in provisionContext ([#&#8203;7004](https://redirect.github.com/caddyserver/caddy/issues/7004))
- [`fe41ff3`](https://redirect.github.com/caddyserver/caddy/commit/fe41ff3c5bbee0aaa44110c110fd6b4fbf052048) core: Save app provisioning errors with context ([#&#8203;7070](https://redirect.github.com/caddyserver/caddy/issues/7070))
- [`e4447c4`](https://redirect.github.com/caddyserver/caddy/commit/e4447c4ba783482c35b096cfbe0e0ffa0403b450) core: Use KeepAliveConfig to pass keepalive\_interval to listener's accepted sockets ([#&#8203;7151](https://redirect.github.com/caddyserver/caddy/issues/7151))
- [`b9710c6`](https://redirect.github.com/caddyserver/caddy/commit/b9710c6af4f764b463a8e0c080783f2b7fb15ce0) fileserver: Add a few doc lines about Etag file content ([#&#8203;7173](https://redirect.github.com/caddyserver/caddy/issues/7173))
- [`3b4d966`](https://redirect.github.com/caddyserver/caddy/commit/3b4d966fba069e9895980dbbf05f84bf818968ba) fileserver: Add sort buttons in grid mode ([#&#8203;7089](https://redirect.github.com/caddyserver/caddy/issues/7089))
- [`54d03ce`](https://redirect.github.com/caddyserver/caddy/commit/54d03ced48a8ed2264ae9248c81f00a1c2648d82) fileserver: Add support for .avif image format ([#&#8203;6988](https://redirect.github.com/caddyserver/caddy/issues/6988))
- [`790f3e0`](https://redirect.github.com/caddyserver/caddy/commit/790f3e0885cae4f6362f35f8ed4f725824d46089) fileserver: denote license for embedded JavaScript for LibreJS ([#&#8203;7127](https://redirect.github.com/caddyserver/caddy/issues/7127))
- [`94147ca`](https://redirect.github.com/caddyserver/caddy/commit/94147caf31f7e56a919432accc2779a22b2ed1a0) fileserver: map invalid path errors to fs.ErrInvalid, and return 400 for any invalid path errors. (close [#&#8203;7008](https://redirect.github.com/caddyserver/caddy/issues/7008)) ([#&#8203;7017](https://redirect.github.com/caddyserver/caddy/issues/7017))
- [`67debd0`](https://redirect.github.com/caddyserver/caddy/commit/67debd0e11c3c880bb6fc7a92325989576b0a7fa) fileserver: set Range header for precompressed static files to force Content Length header to appear ([#&#8203;7042](https://redirect.github.com/caddyserver/caddy/issues/7042))
- [`89ed5f4`](https://redirect.github.com/caddyserver/caddy/commit/89ed5f44de61fcb0c3b7ce93bfefbb8e775d1964) fix: Remove nil arg from zapslog.NewHandler call ([#&#8203;6984](https://redirect.github.com/caddyserver/caddy/issues/6984))
- [`3723e89`](https://redirect.github.com/caddyserver/caddy/commit/3723e895854a1f9e528ec278acb0ed8ef1c63f81) go.mod: Upgrade CertMagic to v0.24.0
- [`3bd4135`](https://redirect.github.com/caddyserver/caddy/commit/3bd413546bf8f7b5e19fcb2104a21600c55cc146) go.mod: Upgrade dependencies
- [`a6d488a`](https://redirect.github.com/caddyserver/caddy/commit/a6d488a15beb01369384e74d0e0159da11272bc3) go.mod: update quic-go to v0.51.0 ([#&#8203;6972](https://redirect.github.com/caddyserver/caddy/issues/6972))
- [`11c6dae`](https://redirect.github.com/caddyserver/caddy/commit/11c6daecd7663bd625ec823f585a1c502468b1e7) go.mod: update quic-go to v0.53.0 ([#&#8203;7094](https://redirect.github.com/caddyserver/caddy/issues/7094))
- [`bbf1dfc`](https://redirect.github.com/caddyserver/caddy/commit/bbf1dfcea283c4b002f2e4c970d2def38fa2046f) headers: Support placeholders in replacement search patterns ([#&#8203;7117](https://redirect.github.com/caddyserver/caddy/issues/7117))
- [`f11c780`](https://redirect.github.com/caddyserver/caddy/commit/f11c780fdc2e4e5298a64ef88d110dd392060a36) http: clean up listeners if some of the listeners fail to bind ([#&#8203;7176](https://redirect.github.com/caddyserver/caddy/issues/7176))
- [`fdf6108`](https://redirect.github.com/caddyserver/caddy/commit/fdf610850b5e5dcb518eb2ad475817d6990b8a8d) http: disable keepalive when KeepAliveInterval is negative ([#&#8203;7158](https://redirect.github.com/caddyserver/caddy/issues/7158))
- [`5b727bd`](https://redirect.github.com/caddyserver/caddy/commit/5b727bde2992e7cb9987208453db42ae6e1c6e1e) httpcaddyfile: Allow naked acme\_dns if dns is set (fix [#&#8203;7091](https://redirect.github.com/caddyserver/caddy/issues/7091))
- [`0badb07`](https://redirect.github.com/caddyserver/caddy/commit/0badb071efc38bb9cc055076f0a48d1725fe8cc8) httpcaddyfile: Fix generated config related to ACME global options
- [`092913a`](https://redirect.github.com/caddyserver/caddy/commit/092913a7a568a5eb4b28c06e12c1274bd5eb1140) httpcaddyfile: Prevent error handler from overriding sub-handler matchers ([#&#8203;6999](https://redirect.github.com/caddyserver/caddy/issues/6999))
- [`77dd12c`](https://redirect.github.com/caddyserver/caddy/commit/77dd12cc785990c5c5da947b4e883029ab8bd552) httpcaddyfile: Validates TLS DNS challenge options ([#&#8203;7099](https://redirect.github.com/caddyserver/caddy/issues/7099))
- [`0f209f6`](https://redirect.github.com/caddyserver/caddy/commit/0f209f62eb1f33e67070ada7fa6f4a7899b8e99d) httpcaddyfile: reject blocks in log\_skip directive ([#&#8203;7056](https://redirect.github.com/caddyserver/caddy/issues/7056))
- [`716d72e`](https://redirect.github.com/caddyserver/caddy/commit/716d72e47538cc4f7bab43b1d973e0f8aa0a9fba) intercept: implement Unwrap for interceptedResponseHandler ([#&#8203;7016](https://redirect.github.com/caddyserver/caddy/issues/7016))
- [`9f71483`](https://redirect.github.com/caddyserver/caddy/commit/9f7148392adb72a6121bf99070efaa1a90ffe901) log: default logger should respect `{in,ex}clude` ([#&#8203;6995](https://redirect.github.com/caddyserver/caddy/issues/6995))
- [`33c88bd`](https://redirect.github.com/caddyserver/caddy/commit/33c88bd2bb543a726274cdf52899edb0639cf5f6) refactor: replace HasPrefix+TrimPrefix with CutPrefix ([#&#8203;7095](https://redirect.github.com/caddyserver/caddy/issues/7095))
- [`ab3b2d6`](https://redirect.github.com/caddyserver/caddy/commit/ab3b2d64ba9bf7dadd4440a84894ebab0ee6d5ea) refactor: use slices.Equal to simplify code ([#&#8203;7141](https://redirect.github.com/caddyserver/caddy/issues/7141))
- [`1c596e3`](https://redirect.github.com/caddyserver/caddy/commit/1c596e3c5a1cd0b52febb1506ffe471918bd1128) reverse\_proxy: use the new KeepAliveConfig to set probe interval ([#&#8203;7157](https://redirect.github.com/caddyserver/caddy/issues/7157))
- [`aa3d20b`](https://redirect.github.com/caddyserver/caddy/commit/aa3d20be3ee451af9465470a28937690104e9422) reverseproxy: Use DialTLSContext if ServerName has placeholder ([#&#8203;6955](https://redirect.github.com/caddyserver/caddy/issues/6955))
- [`737936c`](https://redirect.github.com/caddyserver/caddy/commit/737936c06be001a40c2d743d17d1e3df148408f0) reverseproxy: reference correct field name in LoadModule ([#&#8203;6978](https://redirect.github.com/caddyserver/caddy/issues/6978))
- [`1209b5c`](https://redirect.github.com/caddyserver/caddy/commit/1209b5c5669fc5d0a708931f138bfa7a5d4c5ebc) reverseproxy: validate versions in http transport ([#&#8203;7112](https://redirect.github.com/caddyserver/caddy/issues/7112))

#### New Contributors

- [@&#8203;IndraGunawan](https://redirect.github.com/IndraGunawan) made their first contribution in [#&#8203;6984](https://redirect.github.com/caddyserver/caddy/pull/6984)
- [@&#8203;Compy](https://redirect.github.com/Compy) made their first contribution in [#&#8203;6997](https://redirect.github.com/caddyserver/caddy/pull/6997)
- [@&#8203;eveneast](https://redirect.github.com/eveneast) made their first contribution in [#&#8203;7009](https://redirect.github.com/caddyserver/caddy/pull/7009)
- [@&#8203;tongjicoder](https://redirect.github.com/tongjicoder) made their first contribution in [#&#8203;7039](https://redirect.github.com/caddyserver/caddy/pull/7039)
- [@&#8203;suxatcode](https://redirect.github.com/suxatcode) made their first contribution in [#&#8203;7047](https://redirect.github.com/caddyserver/caddy/pull/7047)
- [@&#8203;Hellio404](https://redirect.github.com/Hellio404) made their first contribution in [#&#8203;6999](https://redirect.github.com/caddyserver/caddy/pull/6999)
- [@&#8203;IwatsukaYura](https://redirect.github.com/IwatsukaYura) made their first contribution in [#&#8203;7056](https://redirect.github.com/caddyserver/caddy/pull/7056)
- [@&#8203;xiaoxiangirl](https://redirect.github.com/xiaoxiangirl) made their first contribution in [#&#8203;7081](https://redirect.github.com/caddyserver/caddy/pull/7081)
- [@&#8203;filipRatajczak](https://redirect.github.com/filipRatajczak) made their first contribution in [#&#8203;7089](https://redirect.github.com/caddyserver/caddy/pull/7089)
- [@&#8203;gopherorg](https://redirect.github.com/gopherorg) made their first contribution in [#&#8203;7095](https://redirect.github.com/caddyserver/caddy/pull/7095)
- [@&#8203;mountdisk](https://redirect.github.com/mountdisk) made their first contribution in [#&#8203;7101](https://redirect.github.com/caddyserver/caddy/pull/7101)
- [@&#8203;bytetigers](https://redirect.github.com/bytetigers) made their first contribution in [#&#8203;7114](https://redirect.github.com/caddyserver/caddy/pull/7114)
- [@&#8203;bytesingsong](https://redirect.github.com/bytesingsong) made their first contribution in [#&#8203;7121](https://redirect.github.com/caddyserver/caddy/pull/7121)
- [@&#8203;zongzewu23](https://redirect.github.com/zongzewu23) made their first contribution in [#&#8203;7117](https://redirect.github.com/caddyserver/caddy/pull/7117)
- [@&#8203;infertux](https://redirect.github.com/infertux) made their first contribution in [#&#8203;7127](https://redirect.github.com/caddyserver/caddy/pull/7127)
- [@&#8203;eeemmmmmm](https://redirect.github.com/eeemmmmmm) made their first contribution in [#&#8203;7136](https://redirect.github.com/caddyserver/caddy/pull/7136)
- [@&#8203;pingshuijie](https://redirect.github.com/pingshuijie) made their first contribution in [#&#8203;7140](https://redirect.github.com/caddyserver/caddy/pull/7140)
- [@&#8203;minxinyi](https://redirect.github.com/minxinyi) made their first contribution in [#&#8203;7141](https://redirect.github.com/caddyserver/caddy/pull/7141)
- [@&#8203;alexandre-daubois](https://redirect.github.com/alexandre-daubois) made their first contribution in [#&#8203;7070](https://redirect.github.com/caddyserver/caddy/pull/7070)
- [@&#8203;joshuamcbeth](https://redirect.github.com/joshuamcbeth) made their first contribution in [#&#8203;7151](https://redirect.github.com/caddyserver/caddy/pull/7151)
- [@&#8203;bosdhill](https://redirect.github.com/bosdhill) made their first contribution in [#&#8203;7163](https://redirect.github.com/caddyserver/caddy/pull/7163)
- [@&#8203;GreyXor](https://redirect.github.com/GreyXor) made their first contribution in [#&#8203;7149](https://redirect.github.com/caddyserver/caddy/pull/7149)
- [@&#8203;youzichuan](https://redirect.github.com/youzichuan) made their first contribution in [#&#8203;7174](https://redirect.github.com/caddyserver/caddy/pull/7174)
- [@&#8203;quagsirus](https://redirect.github.com/quagsirus) made their first contribution in [#&#8203;7179](https://redirect.github.com/caddyserver/caddy/pull/7179)
- [@&#8203;cuiweixie](https://redirect.github.com/cuiweixie) made their first contribution in [#&#8203;7187](https://redirect.github.com/caddyserver/caddy/pull/7187)
- [@&#8203;joemicky](https://redirect.github.com/joemicky) made their first contribution in [#&#8203;7182](https://redirect.github.com/caddyserver/caddy/pull/7182)
- [@&#8203;Pizmovc](https://redirect.github.com/Pizmovc) made their first contribution in [#&#8203;7173](https://redirect.github.com/caddyserver/caddy/pull/7173)

**Full Changelog**: <https://github.com/caddyserver/caddy/compare/v2.10.0...v2.10.1>

### [`v2.10.0`](https://redirect.github.com/caddyserver/caddy/releases/tag/v2.10.0)

[Compare Source](https://redirect.github.com/caddyserver/caddy/compare/v2.9.1...v2.10.0)

Caddy 2.10 is here! Aside from bug fixes, this release features:

- **Encrypted ClientHello (ECH):** This new technology encrypts the last plaintext portion of a TLS connection: the ClientHello, which includes the domain name being connected to. The [draft spec](https://www.ietf.org/archive/id/draft-ietf-tls-esni-24.html) for ECH is almost finalized, so we can now support this privacy feature for TLS. This is a powerful but nuanced capability; we highly recommend reading [the ECH documentation](https://caddyserver.com/docs/automatic-https#encrypted-clienthello-ech) on our website.
- **Post-quantum (PQC) key exchange:** Caddy now supports the standardized `x25519mlkem768` cryptographic group by default.
- **ACME profiles:** ACME profiles are an experimental draft that allow you to choose properties of your certificates with more flexibility than traditional CSR methods. For example, [Let's Encrypt will issue 6-day certificates](https://letsencrypt.org/2025/01/16/6-day-and-ip-certs/) under a certain profile. Caddy may eventually use that profile by default.
- **Via header:** The reverse proxy now sets a Via header instead of a duplicate Server header.
- **Global DNS provider:** You can now specify a default "global" DNS module to use instead of having to configure it locally in every part of your config that requires a DNS provider (for example, ACME DNS challenges, and ECH). This is the `dns` global option in the Caddyfile, or in JSON config, it's the `dns` parameter in the `tls` app configuration.
- **Wildcards used by default:** Previously, Caddy would obtain individual certificates for every domain in your config literally; now wildcards, if present, will be utilized for subdomains, rather than obtaining individual certificates. This change was motivated by the novel possibility for subdomain privacy afforded by ECH. It can be overridden with `tls force_automate` in the Caddyfile. The experimental `auto_https prefer_wildcard` option has been removed.
- **libdns 1.0 APIs:** Many of you use [DNS provider modules](https://redirect.github.com/caddy-dns) to solve ACME DNS challenges or to enable dynamic DNS. They implement interfaces defined by [libdns](https://redirect.github.com/libdns/libdns) to get, set, append, and delete DNS records. After 5 years of production experience, including lessons learned with ECH, libdns APIs have been updated and 1.0 beta has been tagged. DNS provider packages will need to update their code to be compatible, which will help ensure stability and well-defined semantics for the future. Several packages have already updated or are in the process of updating (cloudflare, rfc2136, and desec to name a few).
- **Global `dns` config:** Now that several components of Caddy configuration may affect DNS records (ACME challenges, ECH publication, etc.), there is a new `dns` global option that can be used to specify your DNS provider config in a single place. This prevents repetition of credentials for servers where all the domains are managed by a single DNS provider.

**Thank you to the many contributors who have helped to make this possible!** :tada: :partying\_face: :champagne:

:warning: While have traditionally supported the last 2 minor Go versions to accommodate some distribution / package manager policies, we now only support the latest minor Go version. The privacy and security benefits added in new Go versions (such as post-quantum cryptography) are worth making available to everyone as soon as possible, rather than holding back the entire user base or maintaining multiple code compilation configurations.

#### Encrypted ClientHello (ECH) details

(This is a brief overview. We recommend reading [the full documentation](https://caddyserver.com/docs/automatic-https#encrypted-clienthello-ech).)

Typically, server names (domain names, or "SNI") are sent in the plaintext ClientHello when establishing TLS connections. With ECH, the true server name is encrypted (and wrapped) by an "outer" ClientHello which has a generic SNI of your choosing. With many sites on the same server sharing the same outer SNI, both clients and the server have more privacy related to domain names.

Caddy implements fully automated ECH, meaning that it generates (and [soon](https://redirect.github.com/golang/go/issues/71920), rotates), publishes, and serves ECH configurations simply by specifying a DNS provider, and the outer/public domain name to use.

**Fully automated ECH requires a DNS module built into your Caddy binary.** In order for a client, such as a browser, to know it can use ECH, and what parameters to use, the server's ECH configuration must be published. This config includes the public name, cryptographic parameters, and a public key for encrypting the inner ClientHello. By convention, browsers read the standardized HTTPS-type DNS record containing a `ech` SvcParamKey. Caddy sets this DNS record for all domains being protected, but it needs that DNS provider module plugged in and configured in order to do this. If you are already using the DNS ACME challenge, you should already have a DNS provider plugged in. If you prefer to build Caddy from source with [a DNS module](https://redirect.github.com/caddy-dns), it's easy with [xcaddy](https://redirect.github.com/caddyserver/xcaddy), for example: `$ xcaddy build --with github.com/caddy-dns/cloudflare`

The minimum config needed to enable ClientHello is also the *recommended* config, as it maximizes privacy benefits in most situations. You just need the `ech` global option and a DNS provider specified. Here's an example using Cloudflare as the nameserver:

**Caddyfile:**

```caddy
{
	debug  # not required; recommended while testing
	dns cloudflare {env.CLOUDFLARE_API_KEY}
	ech ech.example.net
}

example.com {
	respond "Hello there!"
}
```

This protects all your sites (`example.com` in this case) behind the public name of `ech.example.net`. (As another example, Cloudflare uses `cloudflare-ech.com` for all the sites it serves. We recommend choosing a single public domain and use it to protect all your sites.)

**The outer/public name you choose should point to your server.** Caddy will obtain a certificate for this name in order to facilitate safe, reliable connections for clients when needed. Without a certificate, clients may be forced to connect insecurely, or fail to connect at all, in some cases, which not only leaves them vulnerable, but also risks exposing the names of your server's sites.

Caddy then uses the specified DNS provider to publish the ECH config(s) for your various site names. It creates (or augments) HTTPS-type records for the domains of your sites (not your ECH public name). Note that DNS provider modules are independently-maintained, and may not have been tested for compatibility with HTTPS-type records. Please contact your module's maintainers if you experience issues.

If you have more advanced configuration needs, you can use the JSON configuration (more details coming soon; for now, see [#&#8203;6862](https://redirect.github.com/caddyserver/caddy/issues/6862) or look at the source code; or use `caddy adapt` to convert a Caddyfile to JSON).

##### Testing and verifying Encrypted ClientHello

First make sure Caddy runs successfully with ECH enabled (and a DNS module) in the config. You should see logs that it is generating an ECH config and publishing it to your domain name(s).

You will need to use a client that supports ECH. Some custom builds of `curl` do, and Firefox and modern Chrome-based browsers do as well, but you need to enable DNS-over-HTTPS or DNS-over-TLS first (since, obviously, querying DNS in plaintext for a protected domain name will expose the domain and defeat the purpose of ECH).

If reusing an existing domain name, clear your DNS cache. Firefox has a way of doing this for its cache at `about:networking#dns`.

Once you have a suitable client, use [Wireshark](https://www.wireshark.org/) to capture network packets as you load your site. You should see *only* the outer/public name as SNI (ServerName Indicator) values in the packet capture. If at any time you see the true site name, ECH is not working properly -- it could be a client or server issue. Before filing a bug, please try to pinpoint it as a server issue first. But definitely report server bugs! Thank you!

(Note that ECH is not automatically published for CNAME'd domains, and the domain must already have a record in the zone.)

#### Commits

##### Beta 1:

- [`96c5c55`](https://redirect.github.com/caddyserver/caddy/commit/96c5c554c1241430ac9ddea6f4b68948adcc961b) admin: fix index validation for PUT requests ([#&#8203;6824](https://redirect.github.com/caddyserver/caddy/issues/6824))
- [`3644ee3`](https://redirect.github.com/caddyserver/caddy/commit/3644ee31cae8e20493d7ccd0c55b0a9c21f20693) build(deps): bump github.com/cloudflare/circl from 1.3.3 to 1.3.7 ([#&#8203;6876](https://redirect.github.com/caddyserver/caddy/issues/6876))
- [`eacd772`](https://redirect.github.com/caddyserver/caddy/commit/eacd7720e99f51b6d2dd340849897c0ff812b8c8) build(deps): bump github.com/go-jose/go-jose/v3 from 3.0.3 to 3.0.4 ([#&#8203;6871](https://redirect.github.com/caddyserver/caddy/issues/6871))
- [`9996d6a`](https://redirect.github.com/caddyserver/caddy/commit/9996d6a70ba76a94dfc90548f25fbac0ce9da497) build(deps): bump github.com/golang/glog from 1.2.2 to 1.2.4 ([#&#8203;6814](https://redirect.github.com/caddyserver/caddy/issues/6814))
- [`1115158`](https://redirect.github.com/caddyserver/caddy/commit/11151586165946453275b66ef33794d41a5e047b) caddyhttp: ResponseRecorder sets stream regardless of 1xx
- [`8861eae`](https://redirect.github.com/caddyserver/caddy/commit/8861eae22350d9e8f94653db951faf85a50a82da) caddytest: Support configuration defaults override ([#&#8203;6850](https://redirect.github.com/caddyserver/caddy/issues/6850))
- [`d7764df`](https://redirect.github.com/caddyserver/caddy/commit/d7764dfdbbee04d2f63aa1b05150737dfddc0bcf) caddytls: Encrypted ClientHello (ECH) ([#&#8203;6862](https://redirect.github.com/caddyserver/caddy/issues/6862))
- [`a807fe0`](https://redirect.github.com/caddyserver/caddy/commit/a807fe065959baa8ee2ad95156183c0850c2b584) caddytls: Enhance ECH documentation
- [`bc3d497`](https://redirect.github.com/caddyserver/caddy/commit/bc3d497739444a5ce550696b7b0da36e6e3bc777) caddytls: Fix broken refactor
- [`7b8f350`](https://redirect.github.com/caddyserver/caddy/commit/7b8f3505e33139de0d542566478e98b361bb84bf) caddytls: Fix sni\_regexp matcher to obtain layer4 contexts ([#&#8203;6804](https://redirect.github.com/caddyserver/caddy/issues/6804))
- [`2c4295e`](https://redirect.github.com/caddyserver/caddy/commit/2c4295ee48f494bc8dda5fa09b37612d520c8b3b) caddytls: Initial support for ACME profiles
- [`d7872c3`](https://redirect.github.com/caddyserver/caddy/commit/d7872c3bfa673ce9584d00f01a725b93fa7bedf1) caddytls: Refactor sni matcher ([#&#8203;6812](https://redirect.github.com/caddyserver/caddy/issues/6812))
- [`172136a`](https://redirect.github.com/caddyserver/caddy/commit/172136a0a0f6aa47be4eab3727fa2482d7af6617) caddytls: Support post-quantum key exchange mechanism X25519MLKEM768
- [`066d770`](https://redirect.github.com/caddyserver/caddy/commit/066d770409917b409d0bdc14cb5ba33d3e4cb33e) cmd: automatically set GOMEMLIMIT ([#&#8203;6809](https://redirect.github.com/caddyserver/caddy/issues/6809))
- [`1f35a8a`](https://redirect.github.com/caddyserver/caddy/commit/1f35a8a4029a338e89998acafa95e1e931a46a27) fastcgi: improve parsePHPFastCGI docs ([#&#8203;6779](https://redirect.github.com/caddyserver/caddy/issues/6779))
- [`22563a7`](https://redirect.github.com/caddyserver/caddy/commit/22563a70eb7b590fcb698680a3ec6d76c0968748) file\_server: use the UTC timezone for modified time ([#&#8203;6830](https://redirect.github.com/caddyserver/caddy/issues/6830))
- [`cfc3af6`](https://redirect.github.com/caddyserver/caddy/commit/cfc3af67492eba22686fd13a2b2201c66cd737f3) fix: update broken link to Ardan Labs ([#&#8203;6800](https://redirect.github.com/caddyserver/caddy/issues/6800))
- [`99073ea`](https://redirect.github.com/caddyserver/caddy/commit/99073eaa33af62bff51c31305e3437c57d936284) go.mod: Upgrade CertMagic to v0.21.7
- [`1641e76`](https://redirect.github.com/caddyserver/caddy/commit/1641e76fd742408c85363e4826451ba9ef22bc99) go.mod: Upgrade dependencies
- [`0d7c639`](https://redirect.github.com/caddyserver/caddy/commit/0d7c63920daecec510202c42816c883fd2dbe047) go.mod: remove glog dependency ([#&#8203;6838](https://redirect.github.com/caddyserver/caddy/issues/6838))
- [`932dac1`](https://redirect.github.com/caddyserver/caddy/commit/932dac157a3c4693b80576477498bb86208b9b30) logging: Always set fields func; fix [#&#8203;6829](https://redirect.github.com/caddyserver/caddy/issues/6829)
- [`9e0e5a4`](https://redirect.github.com/caddyserver/caddy/commit/9e0e5a4b4c2babda81c58f28fe61adfa91d04524) logging: Fix crash if logging error is not HandlerError ([#&#8203;6777](https://redirect.github.com/caddyserver/caddy/issues/6777))
- [`904a0fa`](https://redirect.github.com/caddyserver/caddy/commit/904a0fa368b7eacac3c7156ce4a1f6ced8f61f34) reverse\_proxy: re-add healthy upstreams metric ([#&#8203;6806](https://redirect.github.com/caddyserver/caddy/issues/6806))
- [`e7da3b2`](https://redirect.github.com/caddyserver/caddy/commit/e7da3b267bcec986aaca960dd22ef834d3b9d4a6) reverseproxy: Via header ([#&#8203;6275](https://redirect.github.com/caddyserver/caddy/issues/6275))
- [`9283770`](https://redirect.github.com/caddyserver/caddy/commit/9283770f68f570f47ca20aa9c6f9de8cc50063ba) reverseproxy: ignore duplicate collector registration error ([#&#8203;6820](https://redirect.github.com/caddyserver/caddy/issues/6820))

##### Beta 2:

- [`f4432a3`](https://redirect.github.com/caddyserver/caddy/commit/f4432a306ac59feee1fc45c8efefad3619e37629) caddyfile: add error handling for unrecognized subdirective/options in various modules ([#&#8203;6884](https://redirect.github.com/caddyserver/caddy/issues/6884))
- [`84364ff`](https://redirect.github.com/caddyserver/caddy/commit/84364ffcd06e35a93c9bb08ed80617bde72d4f74) caddypki: Remove lifetime check at Caddyfile parse (fix [#&#8203;6878](https://redirect.github.com/caddyserver/caddy/issues/6878))
- [`adbe7f8`](https://redirect.github.com/caddyserver/caddy/commit/adbe7f87e6bda96a1dddd94ecedefe3219a5304d) caddytls: Only make DNS solver if not already set (fix [#&#8203;6880](https://redirect.github.com/caddyserver/caddy/issues/6880))
- [`d57ab21`](https://redirect.github.com/caddyserver/caddy/commit/d57ab215a2f198a465ea33abe4588bb5696e7abd) caddytls: Pointer receiver (fix [#&#8203;6885](https://redirect.github.com/caddyserver/caddy/issues/6885))
- [`4ebcfed`](https://redirect.github.com/caddyserver/caddy/commit/4ebcfed9c942c59f473f12f8108e1d0fa92e0855) caddytls: Reorder provisioning steps (fix [#&#8203;6877](https://redirect.github.com/caddyserver/caddy/issues/6877))
- [`a686f7c`](https://redirect.github.com/caddyserver/caddy/commit/a686f7c346fe011ad153a3bd4ac3e31e6758bcce) cmd: Only set memory/CPU limits on run (fix [#&#8203;6879](https://redirect.github.com/caddyserver/caddy/issues/6879))
- [`1987620`](https://redirect.github.com/caddyserver/caddy/commit/19876208c79a476a46beec2430e554d4161ab426) cmd: Promote undo maxProcs func to caller
- [`220cd1c`](https://redirect.github.com/caddyserver/caddy/commit/220cd1c2bcecc07bcf6a0141069538c1b1109907) reverseproxy: more comments about buffering and add new tests ([#&#8203;6778](https://redirect.github.com/caddyserver/caddy/issues/6778))

##### Beta 3:

- [`b3e692e`](https://redirect.github.com/caddyserver/caddy/commit/b3e692ed09f8ba15b741621c4b16d8bfee38f8a1) caddyfile: Fix formatting for backquote wrapped braces ([#&#8203;6903](https://redirect.github.com/caddyserver/caddy/issues/6903))
- [`55c89cc`](https://redirect.github.com/caddyserver/caddy/commit/55c89ccf2a39dcfd7286fcaed54787821ff9a1aa) caddytls: Convert AP subjects to punycode
- [`1f8dab5`](https://redirect.github.com/caddyserver/caddy/commit/1f8dab572ca9681464fdadc65bfb5f250fc496c3) caddytls: Don't publish ECH configs if other records don't exist
- [`782a3c7`](https://redirect.github.com/caddyserver/caddy/commit/782a3c7ac60c82311fe9fb8889dd843dfe26c0bc) caddytls: Don't publish HTTPS record for CNAME'd domain (fix [#&#8203;6922](https://redirect.github.com/caddyserver/caddy/issues/6922))
- [`49f9af9`](https://redirect.github.com/caddyserver/caddy/commit/49f9af9a4ab2a28fa5c445630017f5284a5afa48) caddytls: Fix TrustedCACerts backwards compatibility ([#&#8203;6889](https://redirect.github.com/caddyserver/caddy/issues/6889))
- [`e276994`](https://redirect.github.com/caddyserver/caddy/commit/e276994174983dbb190d4bb9acaab157ef14373b) caddytls: Initialize permission module earlier (fix [#&#8203;6901](https://redirect.github.com/caddyserver/caddy/issues/6901))
- [`39262f8`](https://redirect.github.com/caddyserver/caddy/commit/39262f86632401ae4915600b042ef5a28141d3d5) caddytls: Minor fixes for ECH
- [`1735730`](https://redirect.github.com/caddyserver/caddy/commit/173573035c7484bb4aad4498a90bf5a1cf1bb5be) core: add modular `network_proxy` support ([#&#8203;6399](https://redirect.github.com/caddyserver/caddy/issues/6399))
- [`86c620f`](https://redirect.github.com/caddyserver/caddy/commit/86c620fb4e7bfad5888832c491147af53fd5390a) go.mod: Minor dependency upgrades
- [`af2d33a`](https://redirect.github.com/caddyserver/caddy/commit/af2d33afbb52389cda139a6a0fd8a9d65f558676) headers: Allow nil HeaderOps (fix [#&#8203;6893](https://redirect.github.com/caddyserver/caddy/issues/6893))
- [`dccf3d8`](https://redirect.github.com/caddyserver/caddy/commit/dccf3d8982d1b428e840d43f71fa5c3becf6ea8f) requestbody: Add set option to replace request body ([#&#8203;5795](https://redirect.github.com/caddyserver/caddy/issues/5795))
- [`2ac09fd`](https://redirect.github.com/caddyserver/caddy/commit/2ac09fdb2046957597e17096adf6335a6d589a2f) requestbody: Fix ContentLength calculation after body replacement ([#&#8203;6896](https://redirect.github.com/caddyserver/caddy/issues/6896))

##### v2.10.0:

- [`f297bc0`](https://redirect.github.com/caddyserver/caddy/commit/f297bc0a04dcab6c2585b47f3672d045c4f6b54b) admin: Remove host checking for UDS (close [#&#8203;6832](https://redirect.github.com/caddyserver/caddy/issues/6832))
- [`0b2802f`](https://redirect.github.com/caddyserver/caddy/commit/0b2802faa47faa378181a3de5b0d1dcc769a715d) build(deps): bump golang.org/x/net from 0.37.0 to 0.38.0 ([#&#8203;6960](https://redirect.github.com/caddyserver/caddy/issues/6960))
- [`5be77d0`](https://redirect.github.com/caddyserver/caddy/commit/5be77d07ab730e6035ec7a47fb0fe161785af35c) caddyauth: Set authentication provider error in placeholder ([#&#8203;6932](https://redirect.github.com/caddyserver/caddy/issues/6932))
- [`b06a949`](https://redirect.github.com/caddyserver/caddy/commit/b06a9496d130cb06466156d53138a9691342e5a2) caddyhttp: Document side effect of HTTP/3 early data (close [#&#8203;6936](https://redirect.github.com/caddyserver/caddy/issues/6936))
- [`35c8c2d`](https://redirect.github.com/caddyserver/caddy/commit/35c8c2d92d26208642cea0d1549c77a00124e154) caddytls: Add remote\_ip to HTTP cert manager (close [#&#8203;6952](https://redirect.github.com/caddyserver/caddy/issues/6952))
- [`fb22a26`](https://redirect.github.com/caddyserver/caddy/commit/fb22a26b1a08a2fa3b2526d1852467904ee140f6) caddytls: Allow missing ECH meta file
- [`1bfa111`](https://redirect.github.com/caddyserver/caddy/commit/1bfa111552eff8b30bc1a5f76516426f29c66a88) caddytls: Prefer managed wildcard certs over individual subdomain certs ([#&#8203;6959](https://redirect.github.com/caddyserver/caddy/issues/6959))
- [`ea77a9a`](https://redirect.github.com/caddyserver/caddy/commit/ea77a9ab67d8c04f513adaf0a1c648c738e25922) caddytls: Temporarily treat "" and "@&#8203;" as equivalent for DNS publication
- [`5a6b2f8`](https://redirect.github.com/caddyserver/caddy/commit/5a6b2f8d1d4633622b551357f3cc9d27ec669d02) events: Refactor; move Event into core, so core can emit events ([#&#8203;6930](https://redirect.github.com/caddyserver/caddy/issues/6930))
- [`137711a`](https://redirect.github.com/caddyserver/caddy/commit/137711ae3e2d9aa48d7c48dba5ca176af628f073) go.mod: Upgrade acmez and certmagic
- [`9becf61`](https://redirect.github.com/caddyserver/caddy/commit/9becf61a9f5bafb88a15823ce80c1325d3a30a4f) go.mod: Upgrade to libdns 1.0 beta APIs (requires upgraded DNS providers)
- [`6c38ae7`](https://redirect.github.com/caddyserver/caddy/commit/6c38ae7381b3338b173c59706673d11783091dee) reverseproxy: Add valid Upstream to DialInfo in active health checks ([#&#8203;6949](https://redirect.github.com/caddyserver/caddy/issues/6949))

#### What's Changed

- docs: improve parsePHPFastCGI docs by [@&#8203;dunglas](https://redirect.github.com/dunglas) in [#&#8203;6779](https://redirect.github.com/caddyserver/caddy/pull/6779)
- Fixes crash if logging error is not HandlerError by [@&#8203;kkroo](https://redirect.github.com/kkroo) in [#&#8203;6777](https://redirect.github.com/caddyserver/caddy/pull/6777)
- chore: update quic-go to v0.49.0 by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;6803](https://redirect.github.com/caddyserver/caddy/pull/6803)
- chore: don't use deprecated `archives.format_overrides.format` by [@&#8203;mohammed90](https://redirect.github.com/mohammed90) in [#&#8203;6807](https://redirect.github.com/caddyserver/caddy/pull/6807)
- caddytls: Fix sni\_regexp matcher to obtain layer4 contexts by [@&#8203;vnxme](https://redirect.github.com/vnxme) in [#&#8203;6804](https://redirect.github.com/caddyserver/caddy/pull/6804)
- feat: automatically set GOMEMLIMIT by [@&#8203;dunglas](https://redirect.github.com/dunglas) in [#&#8203;6809](https://redirect.github.com/caddyserver/caddy/pull/6809)
- caddytls: Refactor sni matcher by [@&#8203;vnxme](https://redirect.github.com/vnxme) in [#&#8203;6812](https://redirect.github.com/caddyserver/caddy/pull/6812)
- reverse\_proxy: re-add healthy upstreams metric by [@&#8203;mohammed90](https://redirect.github.com/mohammed90) in [#&#8203;6806](https://redirect.github.com/caddyserver/caddy/pull/6806)
- fix: update broken link to Ardan Labs by [@&#8203;sbruens](https://redirect.github.com/sbruens) in [#&#8203;6800](https://redirect.github.com/caddyserver/caddy/pull/6800)
- build(deps): bump github.com/golang/glog from 1.2.2 to 1.2.4 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;6814](https://redirect.github.com/caddyserver/caddy/pull/6814)
- reverseproxy: ignore duplicate collector registration error by [@&#8203;mohammed90](https://redirect.github.com/mohammed90) in [#&#8203;6820](https://redirect.github.com/caddyserver/caddy/pull/6820)
- fix: fix index validation for PUT requests by [@&#8203;debug-ing](https://redirect.github.com/debug-ing) in [#&#8203;6824](https://redirect.github.com/caddyserver/caddy/pull/6824)
- file\_server: use the UTC timezone for modified time by [@&#8203;WeidiDeng](https://redirect.github.com/WeidiDeng) in [#&#8203;6830](https://redirect.github.com/caddyserver/caddy/pull/6830)
- feat/tests: tests for error handling & metrics in admin endpoints by [@&#8203;gdhameeja](https://redirect.github.com/gdhameeja) in [#&#8203;6805](https://redirect.github.com/caddyserver/caddy/pull/6805)
- chore: upgrade Go version to 1.24 by [@&#8203;mohammed90](https://redirect.github.com/mohammed90) in [#&#8203;6839](https://redirect.github.com/caddyserver/caddy/pull/6839)
- remove glog dependency by [@&#8203;Ns2Kracy](https://redirect.github.com/Ns2Kracy) in [#&#8203;6838](https://redirect.github.com/caddyserver/caddy/pull/6838)
- update quic-go to v0.50.0 by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;6854](https://redirect.github.com/caddyserver/caddy/pull/6854)
- Support Caddy Test Configuration Defaults Override. by [@&#8203;baruchyahalom](https://redirect.github.com/baruchyahalom) in [#&#8203;6850](https://redirect.github.com/caddyserver/caddy/pull/6850)
- chore: upgrade cobra by [@&#8203;mohammed90](https://redirect.github.com/mohammed90) in [#&#8203;6868](https://redirect.github.com/caddyserver/caddy/pull/6868)
- build(deps): bump github.com/go-jose/go-jose/v3 from 3.0.3 to 3.0.4 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;6871](https://redirect.github.com/caddyserver/caddy/pull/6871)
- caddytls: Encrypted ClientHello (ECH) by [@&#8203;mholt](https://redirect.github.com/mholt) in [#&#8203;6862](https://redirect.github.com/caddyserver/caddy/pull/6862)
- build(deps): bump github.com/cloudflare/circl from 1.3.3 to 1.3.7 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;6876](https://redirect.github.com/caddyserver/caddy/pull/6876)
- docs: replaced the name and twitter link by [@&#8203;sashaphmn](https://redirect.github.com/sashaphmn) in [#&#8203;6874](https://redirect.github.com/caddyserver/caddy/pull/6874)
- ci: allow using th

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

### Comment by @red-hat-konflux on October 24, 2025 at 08:15 PM UTC

### ℹ Artifact update notice

##### File name: rh_identity_transform/go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 7 additional dependencies were updated
- The `go` directive was updated for compatibility reasons


Details:


| **Package**                           | **Change**              |
| :------------------------------------ | :---------------------- |
| `go`                                  | `1.24.0` -> `1.25`      |
| `github.com/caddyserver/certmagic`    | `v0.21.6` -> `v0.24.0`  |
| `github.com/klauspost/compress`       | `v1.17.11` -> `v1.18.0` |
| `github.com/libdns/libdns`            | `v0.2.3` -> `v1.1.0`    |
| `github.com/prometheus/client_golang` | `v1.20.4` -> `v1.23.0`  |
| `github.com/prometheus/procfs`        | `v0.15.1` -> `v0.16.1`  |
| `github.com/smallstep/certificates`   | `v0.26.1` -> `v0.28.4`  |
| `go.etcd.io/bbolt`                    | `v1.3.9` -> `v1.3.10`   |

---

## Reviews

### Review by @Hyperkid123 - Approved on November 24, 2025 at 11:35 AM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/37*
