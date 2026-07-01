---
type: pull_request
number: 1670
title: "chore(deps): update node.js to v24.18.0"
state: merged
author: red-hat-konflux
created: 2026-06-28T00:22:40Z
updated: 2026-06-28T13:45:01Z
closed: 2026-06-28T13:45:01Z
merged: 2026-06-28T13:45:01Z
base_branch: master
head_branch: konflux/mintmaker/master/node-24.x
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1670
---

# Pull Request #1670: chore(deps): update node.js to v24.18.0

**Author**: @red-hat-konflux
**Created**: June 28, 2026 at 12:22 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/node-24.x`

## Description

This PR contains the following updates:

| Package | Update | Change |
|---|---|---|
| [node](https://nodejs.org) ([source](https://redirect.github.com/nodejs/node)) | minor | `24.17.0` → `24.18.0` |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>nodejs/node (node)</summary>

### [`v24.18.0`](https://redirect.github.com/nodejs/node/releases/tag/v24.18.0): 2026-06-23, Version 24.18.0 'Krypton' (LTS), @&#8203;richardlau prepared by @&#8203;sxa

[Compare Source](https://redirect.github.com/nodejs/node/compare/v24.17.0...v24.18.0)

##### Notable Changes

- \[[`e07e7a31e1`](https://redirect.github.com/nodejs/node/commit/e07e7a31e1)] - **crypto**: update root certificates to NSS 3.123.1 (Node.js GitHub Bot) [#&#8203;63527](https://redirect.github.com/nodejs/node/pull/63527)
- \[[`44c8ebcbd6`](https://redirect.github.com/nodejs/node/commit/44c8ebcbd6)] - **http**: avoid stream listeners on idle agent sockets (Matteo Collina) [#&#8203;64004](https://redirect.github.com/nodejs/node/pull/64004)
- \[[`d3ef4122ee`](https://redirect.github.com/nodejs/node/commit/d3ef4122ee)] - **(SEMVER-MINOR)** **buffer**: increase Buffer.poolSize default to 64 KiB (Matteo Collina) [#&#8203;63597](https://redirect.github.com/nodejs/node/pull/63597)
- \[[`bb2857b85a`](https://redirect.github.com/nodejs/node/commit/bb2857b85a)] - **(SEMVER-MINOR)** **crypto**: align key argument names in docs and error messages (Filip Skokan) [#&#8203;62527](https://redirect.github.com/nodejs/node/pull/62527)
- \[[`b9d5e87880`](https://redirect.github.com/nodejs/node/commit/b9d5e87880)] - **(SEMVER-MINOR)** **crypto**: accept key data in crypto.diffieHellman() and cleanup DH jobs (Filip Skokan) [#&#8203;62527](https://redirect.github.com/nodejs/node/pull/62527)
- \[[`ccd756d61e`](https://redirect.github.com/nodejs/node/commit/ccd756d61e)] - **(SEMVER-MINOR)** **crypto**: add TurboSHAKE and KangarooTwelve Web Cryptography algorithms (Filip Skokan) [#&#8203;62183](https://redirect.github.com/nodejs/node/pull/62183)
- \[[`4c9251fc09`](https://redirect.github.com/nodejs/node/commit/4c9251fc09)] - **(SEMVER-MINOR)** **http**: add writeInformation to send arbitrary 1xx status codes (Tim Perry) [#&#8203;63155](https://redirect.github.com/nodejs/node/pull/63155)
- \[[`8c989ec4a3`](https://redirect.github.com/nodejs/node/commit/8c989ec4a3)] - **(SEMVER-MINOR)** **inspector**: expose precise coverage start to JS runtime (sangwook) [#&#8203;63079](https://redirect.github.com/nodejs/node/pull/63079)
- \[[`3f54c8ba32`](https://redirect.github.com/nodejs/node/commit/3f54c8ba32)] - ***Revert*** "**stream**: noop pause/resume on destroyed streams" (Stewart X Addison) [#&#8203;63834](https://redirect.github.com/nodejs/node/pull/63834)

##### Commits

- \[[`d3ef4122ee`](https://redirect.github.com/nodejs/node/commit/d3ef4122ee)] - **(SEMVER-MINOR)** **buffer**: increase Buffer.poolSize default to 64 KiB (Matteo Collina) [#&#8203;63597](https://redirect.github.com/nodejs/node/pull/63597)
- \[[`9ff36e40f0`](https://redirect.github.com/nodejs/node/commit/9ff36e40f0)] - **build**: add --enable-all-experimentals build flag (Paolo Insogna) [#&#8203;62755](https://redirect.github.com/nodejs/node/pull/62755)
- \[[`7c22ee23aa`](https://redirect.github.com/nodejs/node/commit/7c22ee23aa)] - **build**: def `NODE_USE_NODE_CODE_CACHE` only used in node\_mksnapshot (Chengzhong Wu) [#&#8203;63588](https://redirect.github.com/nodejs/node/pull/63588)
- \[[`2551abdb4a`](https://redirect.github.com/nodejs/node/commit/2551abdb4a)] - **build,win**: enable x64 PGO (Stefan Stojanovic) [#&#8203;62761](https://redirect.github.com/nodejs/node/pull/62761)
- \[[`e8a55ce9b1`](https://redirect.github.com/nodejs/node/commit/e8a55ce9b1)] - **crypto**: strengthen argument CHECKs in TurboSHAKE (Tobias Nießen) [#&#8203;62763](https://redirect.github.com/nodejs/node/pull/62763)
- \[[`ae61cd68f3`](https://redirect.github.com/nodejs/node/commit/ae61cd68f3)] - **crypto**: harden WebCrypto against prototype pollution (Filip Skokan) [#&#8203;63363](https://redirect.github.com/nodejs/node/pull/63363)
- \[[`3d05a1d396`](https://redirect.github.com/nodejs/node/commit/3d05a1d396)] - **crypto**: pass CryptoKey handles to KDF jobs (Filip Skokan) [#&#8203;63363](https://redirect.github.com/nodejs/node/pull/63363)
- \[[`f9d10a3f6b`](https://redirect.github.com/nodejs/node/commit/f9d10a3f6b)] - **crypto**: remove async from WebCrypto methods (Filip Skokan) [#&#8203;63363](https://redirect.github.com/nodejs/node/pull/63363)
- \[[`e431d93e9e`](https://redirect.github.com/nodejs/node/commit/e431d93e9e)] - **crypto**: add WebCrypto CryptoJob mode (Filip Skokan) [#&#8203;63363](https://redirect.github.com/nodejs/node/pull/63363)
- \[[`56e2505e48`](https://redirect.github.com/nodejs/node/commit/56e2505e48)] - **crypto**: wire ML-DSA and ML-KEM for use when using BoringSSL (Filip Skokan) [#&#8203;63255](https://redirect.github.com/nodejs/node/pull/63255)
- \[[`3bac77f2a8`](https://redirect.github.com/nodejs/node/commit/3bac77f2a8)] - **crypto**: wire ChaCha20-Poly1305 in Web Cryptography when using BoringSSL (Filip Skokan) [#&#8203;63255](https://redirect.github.com/nodejs/node/pull/63255)
- \[[`1bff901b09`](https://redirect.github.com/nodejs/node/commit/1bff901b09)] - **crypto**: wire AES-KW in Web Cryptography when using BoringSSL (Filip Skokan) [#&#8203;63255](https://redirect.github.com/nodejs/node/pull/63255)
- \[[`4433fca3df`](https://redirect.github.com/nodejs/node/commit/4433fca3df)] - **crypto**: harden CryptoKey algorithm slots (Filip Skokan) [#&#8203;63111](https://redirect.github.com/nodejs/node/pull/63111)
- \[[`b5cf01217a`](https://redirect.github.com/nodejs/node/commit/b5cf01217a)] - **crypto**: harden KeyObject internal slots (Filip Skokan) [#&#8203;63111](https://redirect.github.com/nodejs/node/pull/63111)
- \[[`ce84aef37d`](https://redirect.github.com/nodejs/node/commit/ce84aef37d)] - **crypto**: add guards and adjust tests for BoringSSL (Filip Skokan) [#&#8203;62883](https://redirect.github.com/nodejs/node/pull/62883)
- \[[`26781689b0`](https://redirect.github.com/nodejs/node/commit/26781689b0)] - **crypto**: reject duplicate ML-KEM JWK key\_ops (Filip Skokan) [#&#8203;62905](https://redirect.github.com/nodejs/node/pull/62905)
- \[[`aeea8f4970`](https://redirect.github.com/nodejs/node/commit/aeea8f4970)] - **crypto**: add JWK support for ML-KEM and SLH-DSA key types (Filip Skokan) [#&#8203;62706](https://redirect.github.com/nodejs/node/pull/62706)
- \[[`407cf91656`](https://redirect.github.com/nodejs/node/commit/407cf91656)] - **crypto**: guard against size\_t overflow on experimental 32-bit arch (Filip Skokan) [#&#8203;62626](https://redirect.github.com/nodejs/node/pull/62626)
- \[[`bb2857b85a`](https://redirect.github.com/nodejs/node/commit/bb2857b85a)] - **(SEMVER-MINOR)** **crypto**: align key argument names in docs and error messages (Filip Skokan) [#&#8203;62527](https://redirect.github.com/nodejs/node/pull/62527)
- \[[`b9d5e87880`](https://redirect.github.com/nodejs/node/commit/b9d5e87880)] - **(SEMVER-MINOR)** **crypto**: accept key data in crypto.diffieHellman() and cleanup DH jobs (Filip Skokan) [#&#8203;62527](https://redirect.github.com/nodejs/node/pull/62527)
- \[[`b46d52b283`](https://redirect.github.com/nodejs/node/commit/b46d52b283)] - **crypto**: unify asymmetric key import through KeyObjectHandle::Init (Filip Skokan) [#&#8203;62499](https://redirect.github.com/nodejs/node/pull/62499)
- \[[`ccd756d61e`](https://redirect.github.com/nodejs/node/commit/ccd756d61e)] - **(SEMVER-MINOR)** **crypto**: add TurboSHAKE and KangarooTwelve Web Cryptography algorithms (Filip Skokan) [#&#8203;62183](https://redirect.github.com/nodejs/node/pull/62183)
- \[[`e07e7a31e1`](https://redirect.github.com/nodejs/node/commit/e07e7a31e1)] - **crypto**: update root certificates to NSS 3.123.1 (Node.js GitHub Bot) [#&#8203;63527](https://redirect.github.com/nodejs/node/pull/63527)
- \[[`61826df455`](https://redirect.github.com/nodejs/node/commit/61826df455)] - **crypto**: coerce -0 keylen to +0 in pbkdf2 and scrypt (Jordan Harband) [#&#8203;63531](https://redirect.github.com/nodejs/node/pull/63531)
- \[[`16d2fd3c07`](https://redirect.github.com/nodejs/node/commit/16d2fd3c07)] - **crypto**: align verifyOneShot accepted types (Anshika Jain) [#&#8203;63280](https://redirect.github.com/nodejs/node/pull/63280)
- \[[`3b8330deda`](https://redirect.github.com/nodejs/node/commit/3b8330deda)] - **crypto**: improve system certificate enumeration logic on macOS (Robo) [#&#8203;62576](https://redirect.github.com/nodejs/node/pull/62576)
- \[[`141de35399`](https://redirect.github.com/nodejs/node/commit/141de35399)] - **debugger**: add --help to `node inspect` and improve docs (Joyee Cheung) [#&#8203;63201](https://redirect.github.com/nodejs/node/pull/63201)
- \[[`b76bfcd4fa`](https://redirect.github.com/nodejs/node/commit/b76bfcd4fa)] - **deps**: upgrade npm to 11.16.0 (npm team) [#&#8203;63602](https://redirect.github.com/nodejs/node/pull/63602)
- \[[`4ec142314c`](https://redirect.github.com/nodejs/node/commit/4ec142314c)] - **deps**: SQLite: cherry-pick [`b869ed6`](https://redirect.github.com/nodejs/node/commit/b869ed6b067d623cb1383549f2a18aa35508385d) (Junsu Han) [#&#8203;63525](https://redirect.github.com/nodejs/node/pull/63525)
- \[[`19e8ce1c36`](https://redirect.github.com/nodejs/node/commit/19e8ce1c36)] - **deps**: upgrade npm to 11.15.0 (npm team) [#&#8203;63463](https://redirect.github.com/nodejs/node/pull/63463)
- \[[`8a264260e2`](https://redirect.github.com/nodejs/node/commit/8a264260e2)] - **deps**: update sqlite to 3.53.1 (Node.js GitHub Bot) [#&#8203;63217](https://redirect.github.com/nodejs/node/pull/63217)
- \[[`50c8ff3f94`](https://redirect.github.com/nodejs/node/commit/50c8ff3f94)] - **deps**: update simdjson to 4.6.4 (Node.js GitHub Bot) [#&#8203;62811](https://redirect.github.com/nodejs/node/pull/62811)
- \[[`6e56f01c4b`](https://redirect.github.com/nodejs/node/commit/6e56f01c4b)] - **deps**: V8: cherry-pick [`435a2cd`](https://redirect.github.com/nodejs/node/commit/435a2cdf664c) (Matthias Liedtke) [#&#8203;63136](https://redirect.github.com/nodejs/node/pull/63136)
- \[[`3ba813b242`](https://redirect.github.com/nodejs/node/commit/3ba813b242)] - **deps**: cherry-pick [libuv/libuv@`a43e543`](https://redirect.github.com/libuv/libuv/commit/a43e543) (Ali Hassan) [#&#8203;63222](https://redirect.github.com/nodejs/node/pull/63222)
- \[[`2390e3a5ac`](https://redirect.github.com/nodejs/node/commit/2390e3a5ac)] - **doc**: remove duplicated sentences in large-pull-requests.md (Joyee Cheung) [#&#8203;63650](https://redirect.github.com/nodejs/node/pull/63650)
- \[[`52a1c18374`](https://redirect.github.com/nodejs/node/commit/52a1c18374)] - **doc**: update `git node land` instructions for security releases (Antoine du Hamel) [#&#8203;63586](https://redirect.github.com/nodejs/node/pull/63586)
- \[[`3e6b4da037`](https://redirect.github.com/nodejs/node/commit/3e6b4da037)] - **doc**: drop --experimental from --permission (Rafael Gonzaga) [#&#8203;63583](https://redirect.github.com/nodejs/node/pull/63583)
- \[[`84d05163b9`](https://redirect.github.com/nodejs/node/commit/84d05163b9)] - **doc**: explicitly ask for reproducible in JS (Rafael Gonzaga) [#&#8203;63479](https://redirect.github.com/nodejs/node/pull/63479)
- \[[`7da2a4450e`](https://redirect.github.com/nodejs/node/commit/7da2a4450e)] - **doc**: fix URL postMessage example in worker\_threads (Kit Dallege) [#&#8203;62203](https://redirect.github.com/nodejs/node/pull/62203)
- \[[`3d79bd8b29`](https://redirect.github.com/nodejs/node/commit/3d79bd8b29)] - **doc**: clarify `filter` option of `sqlite.database.applyChangeset` (Antoine du Hamel) [#&#8203;63515](https://redirect.github.com/nodejs/node/pull/63515)
- \[[`4f4174aace`](https://redirect.github.com/nodejs/node/commit/4f4174aace)] - **doc**: fix double spaces in ERR\_TLS\_INVALID\_PROTOCOL\_METHOD (Daijiro Wachi) [#&#8203;63511](https://redirect.github.com/nodejs/node/pull/63511)
- \[[`388323ca4b`](https://redirect.github.com/nodejs/node/commit/388323ca4b)] - **doc**: fix double space in modules.md (Daijiro Wachi) [#&#8203;63512](https://redirect.github.com/nodejs/node/pull/63512)
- \[[`5258ccc058`](https://redirect.github.com/nodejs/node/commit/5258ccc058)] - **doc**: fix "options" to "option" in tls.createServer (Daijiro Wachi) [#&#8203;63453](https://redirect.github.com/nodejs/node/pull/63453)
- \[[`43e83e6507`](https://redirect.github.com/nodejs/node/commit/43e83e6507)] - **doc**: fix typo in deprecations (Daijiro Wachi) [#&#8203;63434](https://redirect.github.com/nodejs/node/pull/63434)
- \[[`f05a61d54c`](https://redirect.github.com/nodejs/node/commit/f05a61d54c)] - **doc**: remove unsupported template type from v8.md (René) [#&#8203;63410](https://redirect.github.com/nodejs/node/pull/63410)
- \[[`c39d5fc820`](https://redirect.github.com/nodejs/node/commit/c39d5fc820)] - **doc**: fix article usage before vowel-sound acronyms (joao-oliveira-softtor) [#&#8203;62696](https://redirect.github.com/nodejs/node/pull/62696)
- \[[`398261f911`](https://redirect.github.com/nodejs/node/commit/398261f911)] - **doc**: remove the bi-monthly contributor spotlight section (Claudio Wunder) [#&#8203;62734](https://redirect.github.com/nodejs/node/pull/62734)
- \[[`fd9e14c405`](https://redirect.github.com/nodejs/node/commit/fd9e14c405)] - **doc**: update http2's `push` and `trailers` events with `rawHeaders` param (YuSheng Chen) [#&#8203;63259](https://redirect.github.com/nodejs/node/pull/63259)
- \[[`b943ce6933`](https://redirect.github.com/nodejs/node/commit/b943ce6933)] - **doc**: remove inactive members from Triagers list (Antoine du Hamel) [#&#8203;63329](https://redirect.github.com/nodejs/node/pull/63329)
- \[[`4b9cdfc022`](https://redirect.github.com/nodejs/node/commit/4b9cdfc022)] - **doc**: reference correct function in Module docs (Robin Malfait) [#&#8203;63247](https://redirect.github.com/nodejs/node/pull/63247)
- \[[`bed84b6df2`](https://redirect.github.com/nodejs/node/commit/bed84b6df2)] - **doc**: replace Visual Studio 2022 Evergreen version reference with 17.14 (Mike McCready) [#&#8203;63211](https://redirect.github.com/nodejs/node/pull/63211)
- \[[`32ea70569b`](https://redirect.github.com/nodejs/node/commit/32ea70569b)] - **doc**: recommend explicitly Tier 1 or 2 for production applications (Mike McCready) [#&#8203;63187](https://redirect.github.com/nodejs/node/pull/63187)
- \[[`4627bcfd82`](https://redirect.github.com/nodejs/node/commit/4627bcfd82)] - **doc**: run license-builder (github-actions\[bot]) [#&#8203;63232](https://redirect.github.com/nodejs/node/pull/63232)
- \[[`28eba71845`](https://redirect.github.com/nodejs/node/commit/28eba71845)] - **doc**: add large pull requests contributing guide (Matteo Collina) [#&#8203;62829](https://redirect.github.com/nodejs/node/pull/62829)
- \[[`2648efd438`](https://redirect.github.com/nodejs/node/commit/2648efd438)] - **doc**: remove unnecessary `<!-- eslint-` magic comments (Antoine du Hamel) [#&#8203;63200](https://redirect.github.com/nodejs/node/pull/63200)
- \[[`a95fc1f8fc`](https://redirect.github.com/nodejs/node/commit/a95fc1f8fc)] - **doc**: clarify SEA platform support excludes darwin-x64 (MJSHANG) [#&#8203;63181](https://redirect.github.com/nodejs/node/pull/63181)
- \[[`aaef29e2e1`](https://redirect.github.com/nodejs/node/commit/aaef29e2e1)] - **doc**: update release steps when post-release fails (Rafael Gonzaga) [#&#8203;63131](https://redirect.github.com/nodejs/node/pull/63131)
- \[[`7d81419cf2`](https://redirect.github.com/nodejs/node/commit/7d81419cf2)] - **doc**: add Hmac.digest() documentation-only deprecation (DEP0206) (Anshika Jain) [#&#8203;63121](https://redirect.github.com/nodejs/node/pull/63121)
- \[[`ececd80d81`](https://redirect.github.com/nodejs/node/commit/ececd80d81)] - **doc**: document the latest-vX.x schema (Marco Ippolito) [#&#8203;63033](https://redirect.github.com/nodejs/node/pull/63033)
- \[[`27c1c1d842`](https://redirect.github.com/nodejs/node/commit/27c1c1d842)] - **doc**: remove list of versions in `BUILDING.md` (Antoine du Hamel) [#&#8203;63113](https://redirect.github.com/nodejs/node/pull/63113)
- \[[`e369886a65`](https://redirect.github.com/nodejs/node/commit/e369886a65)] - **doc,sqlite**: document entryPoint argument for loadExtension (Edy Silva) [#&#8203;63152](https://redirect.github.com/nodejs/node/pull/63152)
- \[[`e4e5137cbd`](https://redirect.github.com/nodejs/node/commit/e4e5137cbd)] - **errors**: handle V8 warnings in DisallowJavascriptExecutionScope (Divyanshu Sharma) [#&#8203;63491](https://redirect.github.com/nodejs/node/pull/63491)
- \[[`6d1f6048d2`](https://redirect.github.com/nodejs/node/commit/6d1f6048d2)] - **fs**: make `Date` properties on `Stats` enumerable (LiviaMedeiros) [#&#8203;63328](https://redirect.github.com/nodejs/node/pull/63328)
- \[[`44c8ebcbd6`](https://redirect.github.com/nodejs/node/commit/44c8ebcbd6)] - **http**: avoid stream listeners on idle agent sockets (Matteo Collina) [#&#8203;64004](https://redirect.github.com/nodejs/node/pull/64004)
- \[[`4c9251fc09`](https://redirect.github.com/nodejs/node/commit/4c9251fc09)] - **(SEMVER-MINOR)** **http**: add writeInformation to send arbitrary 1xx status codes (Tim Perry) [#&#8203;63155](https://redirect.github.com/nodejs/node/pull/63155)
- \[[`39f61fb06c`](https://redirect.github.com/nodejs/node/commit/39f61fb06c)] - **http2**: emit session close before stream close (Matteo Collina) [#&#8203;63414](https://redirect.github.com/nodejs/node/pull/63414)
- \[[`8a8f2127d1`](https://redirect.github.com/nodejs/node/commit/8a8f2127d1)] - **http2**: validate non-link headers in writeEarlyHints (Matteo Collina) [#&#8203;62017](https://redirect.github.com/nodejs/node/pull/62017)
- \[[`8c989ec4a3`](https://redirect.github.com/nodejs/node/commit/8c989ec4a3)] - **(SEMVER-MINOR)** **inspector**: expose precise coverage start to JS runtime (sangwook) [#&#8203;63079](https://redirect.github.com/nodejs/node/pull/63079)
- \[[`c05f38229b`](https://redirect.github.com/nodejs/node/commit/c05f38229b)] - **lib**: cleanup stateless diffiehellman key handling (Filip Skokan) [#&#8203;62645](https://redirect.github.com/nodejs/node/pull/62645)
- \[[`1c16b45d35`](https://redirect.github.com/nodejs/node/commit/1c16b45d35)] - **lib**: refactor internal webidl converters (Filip Skokan) [#&#8203;62979](https://redirect.github.com/nodejs/node/pull/62979)
- \[[`02f35d6dce`](https://redirect.github.com/nodejs/node/commit/02f35d6dce)] - **lib**: define `kEnumerableProperty` atomically (Antoine du Hamel) [#&#8203;63609](https://redirect.github.com/nodejs/node/pull/63609)
- \[[`12c51547ba`](https://redirect.github.com/nodejs/node/commit/12c51547ba)] - **lib**: fix typos in esm loader comments (RonGamzu) [#&#8203;63465](https://redirect.github.com/nodejs/node/pull/63465)
- \[[`9b03b84262`](https://redirect.github.com/nodejs/node/commit/9b03b84262)] - **lib**: fix typo idenity => identity (Daijiro Wachi) [#&#8203;63112](https://redirect.github.com/nodejs/node/pull/63112)
- \[[`a84e6b0567`](https://redirect.github.com/nodejs/node/commit/a84e6b0567)] - **lib**: fixes validator message (Daijiro Wachi) [#&#8203;62823](https://redirect.github.com/nodejs/node/pull/62823)
- \[[`11734166a8`](https://redirect.github.com/nodejs/node/commit/11734166a8)] - **lib**: narrow ReadableStreamBYOBRequest.view return type to Uint8Array (RoomWithOutRoof) [#&#8203;63017](https://redirect.github.com/nodejs/node/pull/63017)
- \[[`7cead61d21`](https://redirect.github.com/nodejs/node/commit/7cead61d21)] - **meta**: flip mcollina emails in .mailmap (Matteo Collina) [#&#8203;63621](https://redirect.github.com/nodejs/node/pull/63621)
- \[[`a08cfcfd35`](https://redirect.github.com/nodejs/node/commit/a08cfcfd35)] - **meta**: label "source maps" PRs (Chengzhong Wu) [#&#8203;63591](https://redirect.github.com/nodejs/node/pull/63591)
- \[[`d56e8d2512`](https://redirect.github.com/nodejs/node/commit/d56e8d2512)] - **meta**: add `vfs` subsystem label (René) [#&#8203;62331](https://redirect.github.com/nodejs/node/pull/62331)
- \[[`6201cfe488`](https://redirect.github.com/nodejs/node/commit/6201cfe488)] - **meta**: skip scheduled workflows on forks (Jamie Magee) [#&#8203;63565](https://redirect.github.com/nodejs/node/pull/63565)
- \[[`f095e2bd31`](https://redirect.github.com/nodejs/node/commit/f095e2bd31)] - **meta**: add additional gitignore entries (James M Snell) [#&#8203;63267](https://redirect.github.com/nodejs/node/pull/63267)
- \[[`1ea52c444c`](https://redirect.github.com/nodejs/node/commit/1ea52c444c)] - **meta**: move one or more collaborators to emeritus (Node.js GitHub Bot) [#&#8203;63402](https://redirect.github.com/nodejs/node/pull/63402)
- \[[`b1b2327611`](https://redirect.github.com/nodejs/node/commit/b1b2327611)] - **meta**: move one or more collaborators to emeritus (Node.js GitHub Bot) [#&#8203;63235](https://redirect.github.com/nodejs/node/pull/63235)
- \[[`7d88e130a9`](https://redirect.github.com/nodejs/node/commit/7d88e130a9)] - **meta**: ignore AI assistants files (Matteo Collina) [#&#8203;62612](https://redirect.github.com/nodejs/node/pull/62612)
- \[[`a53b51df38`](https://redirect.github.com/nodejs/node/commit/a53b51df38)] - **module**: load ESM helpers eagerly in the snapshot (Joyee Cheung) [#&#8203;63550](https://redirect.github.com/nodejs/node/pull/63550)
- \[[`69df688fff`](https://redirect.github.com/nodejs/node/commit/69df688fff)] - **module**: fix sync hook short-circuit in require() in imported CJS (Joyee Cheung) [#&#8203;62920](https://redirect.github.com/nodejs/node/pull/62920)
- \[[`75d9a4ed47`](https://redirect.github.com/nodejs/node/commit/75d9a4ed47)] - **node-api**: support SharedArrayBuffer in napi\_create\_typedarray (Yilong Li) [#&#8203;62710](https://redirect.github.com/nodejs/node/pull/62710)
- \[[`c20aa4c47b`](https://redirect.github.com/nodejs/node/commit/c20aa4c47b)] - **quic**: add reusePort option to QuicEndpoint (James M Snell) [#&#8203;63267](https://redirect.github.com/nodejs/node/pull/63267)
- \[[`26a30d8a7f`](https://redirect.github.com/nodejs/node/commit/26a30d8a7f)] - **quic**: implement rate limiting for version nego and immediate close (James M Snell) [#&#8203;63267](https://redirect.github.com/nodejs/node/pull/63267)
- \[[`0b534b5770`](https://redirect.github.com/nodejs/node/commit/0b534b5770)] - **quic**: fixup linting issue after other changes (James M Snell) [#&#8203;63267](https://redirect.github.com/nodejs/node/pull/63267)
- \[[`4b367cbe09`](https://redirect.github.com/nodejs/node/commit/4b367cbe09)] - **quic**: remove unused binding variable in session.cc (James M Snell) [#&#8203;63177](https://redirect.github.com/nodejs/node/pull/63177)
- \[[`2574bef5a6`](https://redirect.github.com/nodejs/node/commit/2574bef5a6)] - **repl**: fix dedup comparing normalized line against raw history (Daijiro Wachi) [#&#8203;62886](https://redirect.github.com/nodejs/node/pull/62886)
- \[[`30e71c7e49`](https://redirect.github.com/nodejs/node/commit/30e71c7e49)] - **sqlite**: keep source database alive during backup (Matteo Collina) [#&#8203;62673](https://redirect.github.com/nodejs/node/pull/62673)
- \[[`677ca7e76c`](https://redirect.github.com/nodejs/node/commit/677ca7e76c)] - **src**: simplify OpenSSL feature gates (Filip Skokan) [#&#8203;63255](https://redirect.github.com/nodejs/node/pull/63255)
- \[[`c863c75c39`](https://redirect.github.com/nodejs/node/commit/c863c75c39)] - **src**: add BoringSSL EVP enumeration fallback (Filip Skokan) [#&#8203;63206](https://redirect.github.com/nodejs/node/pull/63206)
- \[[`f6b2466921`](https://redirect.github.com/nodejs/node/commit/f6b2466921)] - **src**: decouple KeyObject and CryptoKey and move CryptoKey to src (Filip Skokan) [#&#8203;62924](https://redirect.github.com/nodejs/node/pull/62924)
- \[[`92d4f07dd2`](https://redirect.github.com/nodejs/node/commit/92d4f07dd2)] - **src**: remove license headers for new node\_profiling files (Chengzhong Wu) [#&#8203;63066](https://redirect.github.com/nodejs/node/pull/63066)
- \[[`8ac5d771c8`](https://redirect.github.com/nodejs/node/commit/8ac5d771c8)] - **src**: split profiling helpers from util (Ilyas Shabi) [#&#8203;63008](https://redirect.github.com/nodejs/node/pull/63008)
- \[[`85d1639495`](https://redirect.github.com/nodejs/node/commit/85d1639495)] - **src**: remove TOCTOU race condition when encoding SAB-backed `Buffer`s (Antoine du Hamel) [#&#8203;63517](https://redirect.github.com/nodejs/node/pull/63517)
- \[[`9473c5f05c`](https://redirect.github.com/nodejs/node/commit/9473c5f05c)] - **src**: skip duplicate UTF-8 validation in TextDecoder fatal path (Mert Can Altin) [#&#8203;63231](https://redirect.github.com/nodejs/node/pull/63231)
- \[[`f35c91ee68`](https://redirect.github.com/nodejs/node/commit/f35c91ee68)] - **src**: improve token return value check (James M Snell) [#&#8203;63483](https://redirect.github.com/nodejs/node/pull/63483)
- \[[`26f677c1c5`](https://redirect.github.com/nodejs/node/commit/26f677c1c5)] - **src**: expose `node::RegisterContext` to make a node managed context (Chengzhong Wu) [#&#8203;62322](https://redirect.github.com/nodejs/node/pull/62322)
- \[[`275cf909b6`](https://redirect.github.com/nodejs/node/commit/275cf909b6)] - **src,sqlite**: only pass `xFilter` when user provided a callback (Antoine du Hamel) [#&#8203;63516](https://redirect.github.com/nodejs/node/pull/63516)
- \[[`287e02303f`](https://redirect.github.com/nodejs/node/commit/287e02303f)] - **src,sqlite**: remove dead code (Edy Silva) [#&#8203;63204](https://redirect.github.com/nodejs/node/pull/63204)
- \[[`58fa2ee189`](https://redirect.github.com/nodejs/node/commit/58fa2ee189)] - **stream**: switch to internal `sleep` binding (Antoine du Hamel) [#&#8203;63611](https://redirect.github.com/nodejs/node/pull/63611)
- \[[`f954ab3f1a`](https://redirect.github.com/nodejs/node/commit/f954ab3f1a)] - **stream**: use data listener for compose forwarding (Trivikram Kamat) [#&#8203;63593](https://redirect.github.com/nodejs/node/pull/63593)
- \[[`dc57173003`](https://redirect.github.com/nodejs/node/commit/dc57173003)] - **stream**: fix Writable.toWeb() hang on synchronous drain (sangwook) [#&#8203;61197](https://redirect.github.com/nodejs/node/pull/61197)
- \[[`3f54c8ba32`](https://redirect.github.com/nodejs/node/commit/3f54c8ba32)] - ***Revert*** "**stream**: noop pause/resume on destroyed streams" (Stewart X Addison) [#&#8203;63834](https://redirect.github.com/nodejs/node/pull/63834)
- \[[`cee279c5d6`](https://redirect.github.com/nodejs/node/commit/cee279c5d6)] - **stream**: remove unnecessary check (Antoine du Hamel) [#&#8203;63030](https://redirect.github.com/nodejs/node/pull/63030)
- \[[`61b20f60a3`](https://redirect.github.com/nodejs/node/commit/61b20f60a3)] - **test**: update tls/crypto behaviour expectations when using BoringSSL (Filip Skokan) [#&#8203;63161](https://redirect.github.com/nodejs/node/pull/63161)
- \[[`a835363808`](https://redirect.github.com/nodejs/node/commit/a835363808)] - **test**: update WPT for WebCryptoAPI to [`97bbc72`](https://redirect.github.com/nodejs/node/commit/97bbc7247a) (Node.js GitHub Bot) [#&#8203;63417](https://redirect.github.com/nodejs/node/pull/63417)
- \[[`a00297480b`](https://redirect.github.com/nodejs/node/commit/a00297480b)] - **test**: update WPT resources, interfaces and WebCryptoAPI (Node.js GitHub Bot) [#&#8203;62389](https://redirect.github.com/nodejs/node/pull/62389)
- \[[`5a95a2b055`](https://redirect.github.com/nodejs/node/commit/5a95a2b055)] - **test**: shorten path in net pipe connect errors (Matteo Collina) [#&#8203;63405](https://redirect.github.com/nodejs/node/pull/63405)
- \[[`5e8ff22d8f`](https://redirect.github.com/nodejs/node/commit/5e8ff22d8f)] - **test**: remove test-node-output-v8-warning (Joyee Cheung) [#&#8203;63469](https://redirect.github.com/nodejs/node/pull/63469)
- \[[`ee15380950`](https://redirect.github.com/nodejs/node/commit/ee15380950)] - **test**: update test426-fixtures to [`9b9e225`](https://redirect.github.com/nodejs/node/commit/9b9e225b5a63139e9a95cdd1bf874a8f0b9d131) (Node.js GitHub Bot) [#&#8203;63373](https://redirect.github.com/nodejs/node/pull/63373)
- \[[`9e063d9bea`](https://redirect.github.com/nodejs/node/commit/9e063d9bea)] - **test**: update WPT for url to [`e4a4672`](https://redirect.github.com/nodejs/node/commit/e4a4672e9e) (Node.js GitHub Bot) [#&#8203;63372](https://redirect.github.com/nodejs/node/pull/63372)
- \[[`503bee4b43`](https://redirect.github.com/nodejs/node/commit/503bee4b43)] - **test**: deflake async-hooks statwatcher test (Trivikram Kamat) [#&#8203;63396](https://redirect.github.com/nodejs/node/pull/63396)
- \[[`cccc7c32d8`](https://redirect.github.com/nodejs/node/commit/cccc7c32d8)] - **test**: avoid test\_runner watch restart in spec snapshot (Trivikram Kamat) [#&#8203;63392](https://redirect.github.com/nodejs/node/pull/63392)
- \[[`c89489258c`](https://redirect.github.com/nodejs/node/commit/c89489258c)] - **test**: reduce watch mode restart flakiness (Trivikram Kamat) [#&#8203;63390](https://redirect.github.com/nodejs/node/pull/63390)
- \[[`e4d5e2578e`](https://redirect.github.com/nodejs/node/commit/e4d5e2578e)] - **test**: isolate rerun-failures state file under tmpdir (Chemi Atlow) [#&#8203;63449](https://redirect.github.com/nodejs/node/pull/63449)
- \[[`362644a9ba`](https://redirect.github.com/nodejs/node/commit/362644a9ba)] - **test**: wait for ok before initial break after restart (Yuya Inoue) [#&#8203;62807](https://redirect.github.com/nodejs/node/pull/62807)
- \[[`c4058d0e05`](https://redirect.github.com/nodejs/node/commit/c4058d0e05)] - **test**: disable Maglev in near-heap-limit worker test (Trivikram Kamat) [#&#8203;63398](https://redirect.github.com/nodejs/node/pull/63398)
- \[[`214da630a7`](https://redirect.github.com/nodejs/node/commit/214da630a7)] - **test**: deflake connection refused proxy tests (Trivikram Kamat) [#&#8203;63395](https://redirect.github.com/nodejs/node/pull/63395)
- \[[`1d61a29876`](https://redirect.github.com/nodejs/node/commit/1d61a29876)] - **test**: avoid repeated writes in watch helper (Trivikram Kamat) [#&#8203;63386](https://redirect.github.com/nodejs/node/pull/63386)
- \[[`2004e25387`](https://redirect.github.com/nodejs/node/commit/2004e25387)] - **test**: deflake watch mode worker test (Trivikram Kamat) [#&#8203;63384](https://redirect.github.com/nodejs/node/pull/63384)
- \[[`d691cccfc1`](https://redirect.github.com/nodejs/node/commit/d691cccfc1)] - **test**: relax test-memory-usage arrayBuffers check (inoway46) [#&#8203;63244](https://redirect.github.com/nodejs/node/pull/63244)
- \[[`0ff6bf853c`](https://redirect.github.com/nodejs/node/commit/0ff6bf853c)] - **test**: reduce flakiness of `different-registry-per-thread` (Antoine du Hamel) [#&#8203;63244](https://redirect.github.com/nodejs/node/pull/63244)
- \[[`d9f4e8e503`](https://redirect.github.com/nodejs/node/commit/d9f4e8e503)] - **test**: fix flaky test-watch-mode-inspect timeout (Matteo Collina) [#&#8203;63361](https://redirect.github.com/nodejs/node/pull/63361)
- \[[`6d7cd50328`](https://redirect.github.com/nodejs/node/commit/6d7cd50328)] - **test**: relax min assertion in test-performance-eventloopdelay (Marco) [#&#8203;63100](https://redirect.github.com/nodejs/node/pull/63100)
- \[[`9dafe1d2d8`](https://redirect.github.com/nodejs/node/commit/9dafe1d2d8)] - **test**: avoid flaky restart sync in debugger exceptions test (Yuya Inoue) [#&#8203;62055](https://redirect.github.com/nodejs/node/pull/62055)
- \[[`989b2de973`](https://redirect.github.com/nodejs/node/commit/989b2de973)] - **test**: avoid initial-break wait in restart-message (inoway46) [#&#8203;62060](https://redirect.github.com/nodejs/node/pull/62060)
- \[[`a072a25ee7`](https://redirect.github.com/nodejs/node/commit/a072a25ee7)] - **test**: move FFI tests to `NATIVE_SUITES` (Antoine du Hamel) [#&#8203;63165](https://redirect.github.com/nodejs/node/pull/63165)
- \[[`64efbfd878`](https://redirect.github.com/nodejs/node/commit/64efbfd878)] - **test**: use ERM to destroy sqlite database handles after tests (René) [#&#8203;63076](https://redirect.github.com/nodejs/node/pull/63076)
- \[[`7dee66cd94`](https://redirect.github.com/nodejs/node/commit/7dee66cd94)] - **test\_runner**: dont buffer unordered events in process isolation mode (Moshe Atlow) [#&#8203;63432](https://redirect.github.com/nodejs/node/pull/63432)
- \[[`d257eec1e3`](https://redirect.github.com/nodejs/node/commit/d257eec1e3)] - **test\_runner**: fix --test-rerun-failures swallowing failures on retry (Chemi Atlow) [#&#8203;63431](https://redirect.github.com/nodejs/node/pull/63431)
- \[[`288c320e2f`](https://redirect.github.com/nodejs/node/commit/288c320e2f)] - **test\_runner**: show replayed-from-attempt hint in spec reporter (Moshe Atlow) [#&#8203;63429](https://redirect.github.com/nodejs/node/pull/63429)
- \[[`904bdf5bb4`](https://redirect.github.com/nodejs/node/commit/904bdf5bb4)] - **test\_runner**: preserve run duration when using test-rerun (Moshe Atlow) [#&#8203;63429](https://redirect.github.com/nodejs/node/pull/63429)
- \[[`df183d7bfa`](https://redirect.github.com/nodejs/node/commit/df183d7bfa)] - **test\_runner**: avoid hanging on incomplete v8 frames (Ali Hassan) [#&#8203;62704](https://redirect.github.com/nodejs/node/pull/62704)
- \[[`ec86c69726`](https://redirect.github.com/nodejs/node/commit/ec86c69726)] - **test\_runner**: fix diagnostics channel context tracking (Moshe Atlow) [#&#8203;63283](https://redirect.github.com/nodejs/node/pull/63283)
- \[[`94e5f63b83`](https://redirect.github.com/nodejs/node/commit/94e5f63b83)] - **tls**: add unsupported renegotiation error (Filip Skokan) [#&#8203;63161](https://redirect.github.com/nodejs/node/pull/63161)
- \[[`06d308fb61`](https://redirect.github.com/nodejs/node/commit/06d308fb61)] - **tools**: prevent lib code from reading KeyObject and CryptoKey accessors (Filip Skokan) [#&#8203;63111](https://redirect.github.com/nodejs/node/pull/63111)
- \[[`2e4a0d0c91`](https://redirect.github.com/nodejs/node/commit/2e4a0d0c91)] - **tools**: bump brace-expansion from 5.0.5 to 5.0.6 in /tools/eslint (dependabot\[bot]) [#&#8203;63415](https://redirect.github.com/nodejs/node/pull/63415)
- \[[`4c9666b366`](https://redirect.github.com/nodejs/node/commit/4c9666b366)] - **tools**: skip commit-lint on backport pull requests (Marco) [#&#8203;63378](https://redirect.github.com/nodejs/node/pull/63378)
- \[[`67d0c490a8`](https://redirect.github.com/nodejs/node/commit/67d0c490a8)] - **tools**: fix skip of `test-internet` on forks (Antoine du Hamel) [#&#8203;63492](https://redirect.github.com/nodejs/node/pull/63492)
- \[[`02f73c7cac`](https://redirect.github.com/nodejs/node/commit/02f73c7cac)] - **tools**: bump the eslint group in /tools/eslint with 4 updates (dependabot\[bot]) [#&#8203;63075](https://redirect.github.com/nodejs/node/pull/63075)
- \[[`5d016d3241`](https://redirect.github.com/nodejs/node/commit/5d016d3241)] - **tools**: update gyp-next to 0.22.2 (Node.js GitHub Bot) [#&#8203;63374](https://redirect.github.com/nodejs/node/pull/63374)
- \[[`55af0f0edb`](https://redirect.github.com/nodejs/node/commit/55af0f0edb)] - **tools**: fix test426 updater (Antoine du Hamel) [#&#8203;63271](https://redirect.github.com/nodejs/node/pull/63271)
- \[[`d8475e167a`](https://redirect.github.com/nodejs/node/commit/d8475e167a)] - **tools**: use different branch for tool updates on staging branches (Antoine du Hamel) [#&#8203;63110](https://redirect.github.com/nodejs/node/pull/63110)
- \[[`c605df9e50`](https://redirect.github.com/nodejs/node/commit/c605df9e50)] - **util**: remove unused functions (Antoine du Hamel) [#&#8203;63612](https://redirect.github.com/nodejs/node/pull/63612)
- \[[`fe4540ebdb`](https://redirect.github.com/nodejs/node/commit/fe4540ebdb)] - **util**: create hex style cache and fast path (Guilherme Araújo) [#&#8203;62999](https://redirect.github.com/nodejs/node/pull/62999)

</details>

---

### Configuration

📅 **Schedule**: (UTC)

- Branch creation
  - At any time (no schedule defined)
- Automerge
  - At any time (no schedule defined)

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0My4yMTAuMC1ycG0iLCJ1cGRhdGVkSW5WZXIiOiI0My4yMTAuMC1ycG0iLCJ0YXJnZXRCcmFuY2giOiJtYXN0ZXIiLCJsYWJlbHMiOltdfQ==-->


---

## Discussion

### Comment by @codecov-commenter on June 28, 2026 at 12:25 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1670?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 77.27%. Comparing base ([`49d0fb1`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/49d0fb1e3781c62e9029b9593a6cc546021cf53b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`1c76fa8`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/1c76fa8275d18db11fc995ac6835299c32cf9c5e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 13 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1670   +/-   ##
=======================================
  Coverage   77.27%   77.27%           
=======================================
  Files         103      103           
  Lines        3287     3287           
  Branches      740      734    -6     
=======================================
  Hits         2540     2540           
  Misses        668      668           
  Partials       79       79           
```
</details>

[:umbrella: View full report in Codecov by Harness](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1670?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @github-actions - Approved on June 28, 2026 at 12:22 AM UTC

This PR from Konflux has been automatically approved.

### Review by @TenSt - Approved on June 28, 2026 at 01:44 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1670*
