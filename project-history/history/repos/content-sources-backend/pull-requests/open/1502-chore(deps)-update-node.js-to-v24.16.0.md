---
type: pull_request
number: 1502
title: "chore(deps): update node.js to v24.16.0"
state: open
author: red-hat-konflux
created: 2026-05-25T05:35:20Z
updated: 2026-06-15T17:30:37Z
base_branch: main
head_branch: konflux/mintmaker/main/node-24.x
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1502
---

# Pull Request #1502: chore(deps): update node.js to v24.16.0

**Author**: @red-hat-konflux
**Created**: May 25, 2026 at 05:35 AM UTC
**Status**: Open
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/node-24.x`

## Description

This PR contains the following updates:

| Package | Update | Change |
|---|---|---|
| [node](https://nodejs.org) ([source](https://redirect.github.com/nodejs/node)) | minor | `24.15.0` → `24.16.0` |

---

### Release Notes

<details>
<summary>nodejs/node (node)</summary>

### [`v24.16.0`](https://redirect.github.com/nodejs/node/releases/tag/v24.16.0): 2026-05-21, Version 24.16.0 &#x27;Krypton&#x27; (LTS), @&#8203;aduh95

[Compare Source](https://redirect.github.com/nodejs/node/compare/v24.15.0...v24.16.0)

##### Notable Changes

- \[[`b267f6bca3`](https://redirect.github.com/nodejs/node/commit/b267f6bca3)] - **(SEMVER-MINOR)** **crypto**: implement `randomUUIDv7()` (nabeel378) [#&#8203;62553](https://redirect.github.com/nodejs/node/pull/62553)
- \[[`ec2451b9cd`](https://redirect.github.com/nodejs/node/commit/ec2451b9cd)] - **(SEMVER-MINOR)** **debugger**: add edit-free runtime expression probes to `node inspect` (Joyee Cheung) [#&#8203;62713](https://redirect.github.com/nodejs/node/pull/62713)
- \[[`9705f628d9`](https://redirect.github.com/nodejs/node/commit/9705f628d9)] - **(SEMVER-MINOR)** **fs**: add signal option to `fs.stat()` (Mert Can Altin) [#&#8203;57775](https://redirect.github.com/nodejs/node/pull/57775)
- \[[`40ccfdecf9`](https://redirect.github.com/nodejs/node/commit/40ccfdecf9)] - **(SEMVER-MINOR)** **fs**: expose `frsize` field in `statfs` (Jinho Jang) [#&#8203;62277](https://redirect.github.com/nodejs/node/pull/62277)
- \[[`d7188af5c9`](https://redirect.github.com/nodejs/node/commit/d7188af5c9)] - **(SEMVER-MINOR)** **http**: harden `ClientRequest` options merge (Matteo Collina) [#&#8203;63082](https://redirect.github.com/nodejs/node/pull/63082)
- \[[`aa1d8a9afc`](https://redirect.github.com/nodejs/node/commit/aa1d8a9afc)] - **(SEMVER-MINOR)** **http**: add `req.signal` to `IncomingMessage` (Akshat) [#&#8203;62541](https://redirect.github.com/nodejs/node/pull/62541)
- \[[`6f37f7e240`](https://redirect.github.com/nodejs/node/commit/6f37f7e240)] - **(SEMVER-MINOR)** **stream**: propagate destruction in `duplexPair` (Ahmed Elhor) [#&#8203;61098](https://redirect.github.com/nodejs/node/pull/61098)
- \[[`d14029be7f`](https://redirect.github.com/nodejs/node/commit/d14029be7f)] - **(SEMVER-MINOR)** **test\_runner**: support test order randomization (Pietro Marchini) [#&#8203;61747](https://redirect.github.com/nodejs/node/pull/61747)
- \[[`d142c584cd`](https://redirect.github.com/nodejs/node/commit/d142c584cd)] - **(SEMVER-MINOR)** **test\_runner**: align mock timeout api (sangwook) [#&#8203;62820](https://redirect.github.com/nodejs/node/pull/62820)
- \[[`01a9552585`](https://redirect.github.com/nodejs/node/commit/01a9552585)] - **(SEMVER-MINOR)** **test\_runner**: add mock-timers support for `AbortSignal.timeout` (DeveloperViraj) [#&#8203;60751](https://redirect.github.com/nodejs/node/pull/60751)
- \[[`00705a459a`](https://redirect.github.com/nodejs/node/commit/00705a459a)] - **(SEMVER-MINOR)** **util**: colorize text with hex colors (Guilherme Araújo) [#&#8203;61556](https://redirect.github.com/nodejs/node/pull/61556)

##### Commits

- \[[`dd72df060d`](https://redirect.github.com/nodejs/node/commit/dd72df060d)] - **assert,util**: fix stale nested cycle memo entries (Ruben Bridgewater) [#&#8203;62509](https://redirect.github.com/nodejs/node/pull/62509)
- \[[`add94f4bc3`](https://redirect.github.com/nodejs/node/commit/add94f4bc3)] - **build**: track PDL files as inputs in inspector GN build (Robo) [#&#8203;62888](https://redirect.github.com/nodejs/node/pull/62888)
- \[[`1b1eb9e334`](https://redirect.github.com/nodejs/node/commit/1b1eb9e334)] - **build**: remove redundant -fuse-linker-plugin from GCC LTO flags (Daniel Lando) [#&#8203;62667](https://redirect.github.com/nodejs/node/pull/62667)
- \[[`8752b604ec`](https://redirect.github.com/nodejs/node/commit/8752b604ec)] - **crypto**: deduplicate and canonicalize CryptoKey usages (Filip Skokan) [#&#8203;62902](https://redirect.github.com/nodejs/node/pull/62902)
- \[[`341947e7fd`](https://redirect.github.com/nodejs/node/commit/341947e7fd)] - **crypto**: reject unintended raw key format string input (Filip Skokan) [#&#8203;62974](https://redirect.github.com/nodejs/node/pull/62974)
- \[[`28a78747fc`](https://redirect.github.com/nodejs/node/commit/28a78747fc)] - **crypto**: remove Argon2 KDF derivation from its job setup (Filip Skokan) [#&#8203;62863](https://redirect.github.com/nodejs/node/pull/62863)
- \[[`16e8c2b54d`](https://redirect.github.com/nodejs/node/commit/16e8c2b54d)] - **crypto**: fix unsigned conversion of 4-byte RSA publicExponent (DeepView Autofix) [#&#8203;62839](https://redirect.github.com/nodejs/node/pull/62839)
- \[[`eeae754a87`](https://redirect.github.com/nodejs/node/commit/eeae754a87)] - **crypto**: reject inherited key type names (Jonathan Lopes) [#&#8203;62875](https://redirect.github.com/nodejs/node/pull/62875)
- \[[`9dd5540325`](https://redirect.github.com/nodejs/node/commit/9dd5540325)] - **crypto**: add memory tracking for secureContext openssl objects (Mert Can Altin) [#&#8203;59051](https://redirect.github.com/nodejs/node/pull/59051)
- \[[`b267f6bca3`](https://redirect.github.com/nodejs/node/commit/b267f6bca3)] - **(SEMVER-MINOR)** **crypto**: implement randomUUIDv7() (nabeel378) [#&#8203;62553](https://redirect.github.com/nodejs/node/pull/62553)
- \[[`7597d204c1`](https://redirect.github.com/nodejs/node/commit/7597d204c1)] - **crypto**: add support for [`Ed25519`](https://redirect.github.com/nodejs/node/commit/Ed25519) context parameter (Filip Skokan) [#&#8203;62474](https://redirect.github.com/nodejs/node/pull/62474)
- \[[`4bf85845da`](https://redirect.github.com/nodejs/node/commit/4bf85845da)] - **debugger**: move ProbeInspectorSession and helpers to separate files (Joyee Cheung) [#&#8203;63013](https://redirect.github.com/nodejs/node/pull/63013)
- \[[`ec2451b9cd`](https://redirect.github.com/nodejs/node/commit/ec2451b9cd)] - **(SEMVER-MINOR)** **debugger**: add edit-free runtime expression probes to `node inspect` (Joyee Cheung) [#&#8203;62713](https://redirect.github.com/nodejs/node/pull/62713)
- \[[`83e98f77b7`](https://redirect.github.com/nodejs/node/commit/83e98f77b7)] - **deps**: update corepack to 0.35.0 (Node.js GitHub Bot) [#&#8203;63375](https://redirect.github.com/nodejs/node/pull/63375)
- \[[`ec8c6b939a`](https://redirect.github.com/nodejs/node/commit/ec8c6b939a)] - **deps**: V8: cherry-pick [`657d8de`](https://redirect.github.com/nodejs/node/commit/657d8de27427) (Guy Bedford) [#&#8203;62784](https://redirect.github.com/nodejs/node/pull/62784)
- \[[`722c0c3274`](https://redirect.github.com/nodejs/node/commit/722c0c3274)] - **deps**: update nghttp3 to 1.14.0 (Node.js GitHub Bot) [#&#8203;61187](https://redirect.github.com/nodejs/node/pull/61187)
- \[[`5304db93d3`](https://redirect.github.com/nodejs/node/commit/5304db93d3)] - **deps**: update nghttp3 to 1.13.1 (Node.js GitHub Bot) [#&#8203;60046](https://redirect.github.com/nodejs/node/pull/60046)
- \[[`e073b3811d`](https://redirect.github.com/nodejs/node/commit/e073b3811d)] - **deps**: update nghttp3 to 1.11.0 (James M Snell) [#&#8203;59249](https://redirect.github.com/nodejs/node/pull/59249)
- \[[`1d00313fb2`](https://redirect.github.com/nodejs/node/commit/1d00313fb2)] - **deps**: update ngtcp2 to 1.14.0 (James M Snell) [#&#8203;59249](https://redirect.github.com/nodejs/node/pull/59249)
- \[[`8b3a4fc18f`](https://redirect.github.com/nodejs/node/commit/8b3a4fc18f)] - **deps**: update amaro to 1.1.9 (Node.js GitHub Bot) [#&#8203;63090](https://redirect.github.com/nodejs/node/pull/63090)
- \[[`62fe0cfcd1`](https://redirect.github.com/nodejs/node/commit/62fe0cfcd1)] - **deps**: update llhttp to 9.4.1 (Node.js GitHub Bot) [#&#8203;63045](https://redirect.github.com/nodejs/node/pull/63045)
- \[[`137e09c8e9`](https://redirect.github.com/nodejs/node/commit/137e09c8e9)] - **deps**: update corepack to 0.34.7 (Node.js GitHub Bot) [#&#8203;62810](https://redirect.github.com/nodejs/node/pull/62810)
- \[[`14a4cb8fbc`](https://redirect.github.com/nodejs/node/commit/14a4cb8fbc)] - **deps**: update timezone to 2026b (Node.js GitHub Bot) [#&#8203;62962](https://redirect.github.com/nodejs/node/pull/62962)
- \[[`3e1036583a`](https://redirect.github.com/nodejs/node/commit/3e1036583a)] - **deps**: upgrade npm to 11.13.0 (npm team) [#&#8203;62898](https://redirect.github.com/nodejs/node/pull/62898)
- \[[`01dfe5961c`](https://redirect.github.com/nodejs/node/commit/01dfe5961c)] - **deps**: cherry-pick [libuv/libuv@`439a54b`](https://redirect.github.com/libuv/libuv/commit/439a54b) (skooch) [#&#8203;62881](https://redirect.github.com/nodejs/node/pull/62881)
- \[[`6cd368b10c`](https://redirect.github.com/nodejs/node/commit/6cd368b10c)] - **deps**: update sqlite to 3.53.0 (Node.js GitHub Bot) [#&#8203;62699](https://redirect.github.com/nodejs/node/pull/62699)
- \[[`f218a4f553`](https://redirect.github.com/nodejs/node/commit/f218a4f553)] - **deps**: update nbytes to 0.1.4 (Node.js GitHub Bot) [#&#8203;62698](https://redirect.github.com/nodejs/node/pull/62698)
- \[[`b47688524a`](https://redirect.github.com/nodejs/node/commit/b47688524a)] - **deps**: update archs files for openssl-3.5.6 (Node.js GitHub Bot) [#&#8203;62629](https://redirect.github.com/nodejs/node/pull/62629)
- \[[`d202e2d343`](https://redirect.github.com/nodejs/node/commit/d202e2d343)] - **deps**: upgrade openssl sources to openssl-3.5.6 (Node.js GitHub Bot) [#&#8203;62629](https://redirect.github.com/nodejs/node/pull/62629)
- \[[`2faba66341`](https://redirect.github.com/nodejs/node/commit/2faba66341)] - **deps**: update minimatch to 10.2.5 (Node.js GitHub Bot) [#&#8203;62594](https://redirect.github.com/nodejs/node/pull/62594)
- \[[`fa46c90c5d`](https://redirect.github.com/nodejs/node/commit/fa46c90c5d)] - **deps**: update googletest to [`d72f9c8`](https://redirect.github.com/nodejs/node/commit/d72f9c8aea6817cdf1ca0ac10887f328de7f3da2) (Node.js GitHub Bot) [#&#8203;62593](https://redirect.github.com/nodejs/node/pull/62593)
- \[[`099ded5713`](https://redirect.github.com/nodejs/node/commit/099ded5713)] - **deps**: update simdjson to 4.6.1 (Node.js GitHub Bot) [#&#8203;62592](https://redirect.github.com/nodejs/node/pull/62592)
- \[[`7ce95afe96`](https://redirect.github.com/nodejs/node/commit/7ce95afe96)] - **deps**: libuv: cherry-pick [`aabb765`](https://redirect.github.com/nodejs/node/commit/aabb7651de) (Santiago Gimeno) [#&#8203;62561](https://redirect.github.com/nodejs/node/pull/62561)
- \[[`57ef845623`](https://redirect.github.com/nodejs/node/commit/57ef845623)] - **deps**: update icu to 78.3 (Node.js GitHub Bot) [#&#8203;62324](https://redirect.github.com/nodejs/node/pull/62324)
- \[[`493ac40e12`](https://redirect.github.com/nodejs/node/commit/493ac40e12)] - **deps**: update libuv to 1.52.1 (Node.js GitHub Bot) [#&#8203;61829](https://redirect.github.com/nodejs/node/pull/61829)
- \[[`b39508b368`](https://redirect.github.com/nodejs/node/commit/b39508b368)] - **deps**: update undici to 7.25.0 (Node.js GitHub Bot) [#&#8203;63011](https://redirect.github.com/nodejs/node/pull/63011)
- \[[`cb67a925e9`](https://redirect.github.com/nodejs/node/commit/cb67a925e9)] - **deps**: use npm undici\@&#8203;seven tag in `update-undici.sh` (Matteo Collina) [#&#8203;62739](https://redirect.github.com/nodejs/node/pull/62739)
- \[[`aa1e0bc28b`](https://redirect.github.com/nodejs/node/commit/aa1e0bc28b)] - **doc**: fix typos and inconsistencies in crypto.md and webcrypto.md (Filip Skokan) [#&#8203;62828](https://redirect.github.com/nodejs/node/pull/62828)
- \[[`f2a1735ed9`](https://redirect.github.com/nodejs/node/commit/f2a1735ed9)] - **doc**: fix duplicate word "to to" in util.styleText (Daijiro Wachi) [#&#8203;62917](https://redirect.github.com/nodejs/node/pull/62917)
- \[[`b6378e215c`](https://redirect.github.com/nodejs/node/commit/b6378e215c)] - **doc**: fix node-config-schema (Сковорода Никита Андреевич) [#&#8203;61596](https://redirect.github.com/nodejs/node/pull/61596)
- \[[`233894a9ce`](https://redirect.github.com/nodejs/node/commit/233894a9ce)] - **doc**: fix the TypeScript Execute (tsx) project link (David Thornton) [#&#8203;63093](https://redirect.github.com/nodejs/node/pull/63093)
- \[[`5d97919f8f`](https://redirect.github.com/nodejs/node/commit/5d97919f8f)] - **doc**: correct diagnostics\_channel built-in channel names (Bryan English) [#&#8203;62995](https://redirect.github.com/nodejs/node/pull/62995)
- \[[`2a9ccc927e`](https://redirect.github.com/nodejs/node/commit/2a9ccc927e)] - **doc**: use mjs/cjs blocks for callbackify null reason example (Daijiro Wachi) [#&#8203;62884](https://redirect.github.com/nodejs/node/pull/62884)
- \[[`ef413b5358`](https://redirect.github.com/nodejs/node/commit/ef413b5358)] - **doc**: fix typo in test.md (Rich Trott) [#&#8203;62960](https://redirect.github.com/nodejs/node/pull/62960)
- \[[`76f21c5070`](https://redirect.github.com/nodejs/node/commit/76f21c5070)] - **doc**: correct typo in PR contribution instructions (Mike McCready) [#&#8203;62738](https://redirect.github.com/nodejs/node/pull/62738)
- \[[`ca02af1f7d`](https://redirect.github.com/nodejs/node/commit/ca02af1f7d)] - **doc**: fix duplicate word "of of" in postMessageToThread (Daijiro Wachi) [#&#8203;62917](https://redirect.github.com/nodejs/node/pull/62917)
- \[[`46c99ed526`](https://redirect.github.com/nodejs/node/commit/46c99ed526)] - **doc**: fix duplicate word "for for" in compile cache (Daijiro Wachi) [#&#8203;62917](https://redirect.github.com/nodejs/node/pull/62917)
- \[[`1a60851734`](https://redirect.github.com/nodejs/node/commit/1a60851734)] - **doc**: fix typo in dns.lookup options description (Daijiro Wachi) [#&#8203;62882](https://redirect.github.com/nodejs/node/pull/62882)
- \[[`169b5ea2ed`](https://redirect.github.com/nodejs/node/commit/169b5ea2ed)] - **doc**: fix Argon2 parameter bounds (Tobias Nießen) [#&#8203;62868](https://redirect.github.com/nodejs/node/pull/62868)
- \[[`9a3a190f4e`](https://redirect.github.com/nodejs/node/commit/9a3a190f4e)] - **doc**: clarify diffieHellman.generateKeys recomputes same key (Kit Dallege) [#&#8203;62205](https://redirect.github.com/nodejs/node/pull/62205)
- \[[`0fba9e87d6`](https://redirect.github.com/nodejs/node/commit/0fba9e87d6)] - **doc**: remove Ayase-252 and meixg from triagger team (Antoine du Hamel) [#&#8203;62841](https://redirect.github.com/nodejs/node/pull/62841)
- \[[`9c700f3446`](https://redirect.github.com/nodejs/node/commit/9c700f3446)] - **doc**: clarify dns.lookup() callback signature when all is true (eungi) [#&#8203;62800](https://redirect.github.com/nodejs/node/pull/62800)
- \[[`6b7280bc17`](https://redirect.github.com/nodejs/node/commit/6b7280bc17)] - **doc**: add experimental modules lifetime policy (Paolo Insogna) [#&#8203;62753](https://redirect.github.com/nodejs/node/pull/62753)
- \[[`ce47ea31c9`](https://redirect.github.com/nodejs/node/commit/ce47ea31c9)] - **doc**: clarify process.\_debugProcess() in Permission Model (Fahad Khan) [#&#8203;62537](https://redirect.github.com/nodejs/node/pull/62537)
- \[[`ba01633757`](https://redirect.github.com/nodejs/node/commit/ba01633757)] - **doc**: fix typo in devcontainer guide (Rohan Santhosh Kumar) [#&#8203;62687](https://redirect.github.com/nodejs/node/pull/62687)
- \[[`70b4d5839b`](https://redirect.github.com/nodejs/node/commit/70b4d5839b)] - **doc**: clarify Backport-PR-URL metadata added automatically (Mike McCready) [#&#8203;62668](https://redirect.github.com/nodejs/node/pull/62668)
- \[[`8126d1c3eb`](https://redirect.github.com/nodejs/node/commit/8126d1c3eb)] - **doc**: update WPT test runner README.md (Filip Skokan) [#&#8203;62680](https://redirect.github.com/nodejs/node/pull/62680)
- \[[`978afea4b5`](https://redirect.github.com/nodejs/node/commit/978afea4b5)] - **doc**: fix spelling in release announcement guidance (Rohan Santhosh Kumar) [#&#8203;62663](https://redirect.github.com/nodejs/node/pull/62663)
- \[[`1684ab8ff8`](https://redirect.github.com/nodejs/node/commit/1684ab8ff8)] - **doc**: note non-monotonic clock in crypto.randomUUIDv7 (nabeel378) [#&#8203;62600](https://redirect.github.com/nodejs/node/pull/62600)
- \[[`86d4f07930`](https://redirect.github.com/nodejs/node/commit/86d4f07930)] - **doc**: update bug bounty program (Rafael Gonzaga) [#&#8203;62590](https://redirect.github.com/nodejs/node/pull/62590)
- \[[`736ed8a08f`](https://redirect.github.com/nodejs/node/commit/736ed8a08f)] - **doc**: document TransformStream transformer.cancel option (Tom Pereira) [#&#8203;62566](https://redirect.github.com/nodejs/node/pull/62566)
- \[[`938af9be01`](https://redirect.github.com/nodejs/node/commit/938af9be01)] - **doc**: mention test runner retry attemp is zero based (Moshe Atlow) [#&#8203;62504](https://redirect.github.com/nodejs/node/pull/62504)
- \[[`94433e450f`](https://redirect.github.com/nodejs/node/commit/94433e450f)] - **doc,src,test**: fix dead inspector help URL (semimikoh) [#&#8203;62745](https://redirect.github.com/nodejs/node/pull/62745)
- \[[`ddf1f01659`](https://redirect.github.com/nodejs/node/commit/ddf1f01659)] - **esm**: add `ERR_REQUIRE_ESM_RACE_CONDITION` (Antoine du Hamel) [#&#8203;62462](https://redirect.github.com/nodejs/node/pull/62462)
- \[[`4a506acd16`](https://redirect.github.com/nodejs/node/commit/4a506acd16)] - **fs**: add followSymlinks option to glob (Matteo Collina) [#&#8203;62695](https://redirect.github.com/nodejs/node/pull/62695)
- \[[`f4ea495f9b`](https://redirect.github.com/nodejs/node/commit/f4ea495f9b)] - **fs**: restore fs patchability in ESM loader (Joyee Cheung) [#&#8203;62835](https://redirect.github.com/nodejs/node/pull/62835)
- \[[`63c111cd60`](https://redirect.github.com/nodejs/node/commit/63c111cd60)] - **fs**: validate position argument before length === 0 early return (Edy Silva) [#&#8203;62674](https://redirect.github.com/nodejs/node/pull/62674)
- \[[`9705f628d9`](https://redirect.github.com/nodejs/node/commit/9705f628d9)] - **(SEMVER-MINOR)** **fs**: add signal option to fs.stat() (Mert Can Altin) [#&#8203;57775](https://redirect.github.com/nodejs/node/pull/57775)
- \[[`40ccfdecf9`](https://redirect.github.com/nodejs/node/commit/40ccfdecf9)] - **(SEMVER-MINOR)** **fs**: expose frsize field in statfs (Jinho Jang) [#&#8203;62277](https://redirect.github.com/nodejs/node/pull/62277)
- \[[`717476a24e`](https://redirect.github.com/nodejs/node/commit/717476a24e)] - **http**: emit 'drain' on OutgoingMessage only after buffers drain (Robert Nagy) [#&#8203;62936](https://redirect.github.com/nodejs/node/pull/62936)
- \[[`d7188af5c9`](https://redirect.github.com/nodejs/node/commit/d7188af5c9)] - **(SEMVER-MINOR)** **http**: harden ClientRequest options merge (Matteo Collina) [#&#8203;63082](https://redirect.github.com/nodejs/node/pull/63082)
- \[[`64f15c274a`](https://redirect.github.com/nodejs/node/commit/64f15c274a)] - **http**: fix leaked error listener on sync HTTP req create + destroy (Tim Perry) [#&#8203;62872](https://redirect.github.com/nodejs/node/pull/62872)
- \[[`5c4798d799`](https://redirect.github.com/nodejs/node/commit/5c4798d799)] - **http**: fix no\_proxy leading-dot suffix matching (Daijiro Wachi) [#&#8203;62333](https://redirect.github.com/nodejs/node/pull/62333)
- \[[`9f3bc70ae5`](https://redirect.github.com/nodejs/node/commit/9f3bc70ae5)] - **http**: cleanup pipeline queue (Robert Nagy) [#&#8203;62534](https://redirect.github.com/nodejs/node/pull/62534)
- \[[`aa1d8a9afc`](https://redirect.github.com/nodejs/node/commit/aa1d8a9afc)] - **(SEMVER-MINOR)** **http**: add req.signal to IncomingMessage (Akshat) [#&#8203;62541](https://redirect.github.com/nodejs/node/pull/62541)
- \[[`900dc758ff`](https://redirect.github.com/nodejs/node/commit/900dc758ff)] - **http2**: expose writable stream state on compat response (T) [#&#8203;63003](https://redirect.github.com/nodejs/node/pull/63003)
- \[[`b3bfe35912`](https://redirect.github.com/nodejs/node/commit/b3bfe35912)] - **inspector**: coerce key and value to string in webstorage events (Ali Hassan) [#&#8203;62616](https://redirect.github.com/nodejs/node/pull/62616)
- \[[`3dc3fb6ad8`](https://redirect.github.com/nodejs/node/commit/3dc3fb6ad8)] - **inspector**: return errors when CDP protocol event emission fails (Ryuhei Shima) [#&#8203;62162](https://redirect.github.com/nodejs/node/pull/62162)
- \[[`4f3f21bd7c`](https://redirect.github.com/nodejs/node/commit/4f3f21bd7c)] - **inspector**: auto collect webstorage data (Ryuhei Shima) [#&#8203;62145](https://redirect.github.com/nodejs/node/pull/62145)
- \[[`36cc04189d`](https://redirect.github.com/nodejs/node/commit/36cc04189d)] - **inspector**: initial support storage inspection (Ryuhei Shima) [#&#8203;61139](https://redirect.github.com/nodejs/node/pull/61139)
- \[[`1718bc3b9b`](https://redirect.github.com/nodejs/node/commit/1718bc3b9b)] - **inspector**: fix absolute URLs in network http (bugyaluwang) [#&#8203;62955](https://redirect.github.com/nodejs/node/pull/62955)
- \[[`97e32c7a74`](https://redirect.github.com/nodejs/node/commit/97e32c7a74)] - **lib**: avoid quadratic shift() in startup snapshot callback (Daijiro Wachi) [#&#8203;62914](https://redirect.github.com/nodejs/node/pull/62914)
- \[[`25d2e999de`](https://redirect.github.com/nodejs/node/commit/25d2e999de)] - **lib**: harden kKeyOps lookup with null prototype (Filip Skokan) [#&#8203;62877](https://redirect.github.com/nodejs/node/pull/62877)
- \[[`37d3913c8f`](https://redirect.github.com/nodejs/node/commit/37d3913c8f)] - **lib**: short-circuit WebIDL BufferSource SAB check (Filip Skokan) [#&#8203;62833](https://redirect.github.com/nodejs/node/pull/62833)
- \[[`430c69d25f`](https://redirect.github.com/nodejs/node/commit/430c69d25f)] - **lib**: use js-only implementation of `isDataView()` (René) [#&#8203;62780](https://redirect.github.com/nodejs/node/pull/62780)
- \[[`3ba0add6a0`](https://redirect.github.com/nodejs/node/commit/3ba0add6a0)] - **lib**: fix lint in internal/webstreams/util.js (Filip Skokan) [#&#8203;62806](https://redirect.github.com/nodejs/node/pull/62806)
- \[[`9b95c41398`](https://redirect.github.com/nodejs/node/commit/9b95c41398)] - **lib**: fix sequence argument handling in Blob constructor (Ms2ger) [#&#8203;62179](https://redirect.github.com/nodejs/node/pull/62179)
- \[[`314dacdbee`](https://redirect.github.com/nodejs/node/commit/314dacdbee)] - **lib**: improve Web Cryptography key validation ordering (Filip Skokan) [#&#8203;62749](https://redirect.github.com/nodejs/node/pull/62749)
- \[[`3d18162430`](https://redirect.github.com/nodejs/node/commit/3d18162430)] - **lib**: reject SharedArrayBuffer in web APIs per spec (Ali Hassan) [#&#8203;62632](https://redirect.github.com/nodejs/node/pull/62632)
- \[[`ada3ce879d`](https://redirect.github.com/nodejs/node/commit/ada3ce879d)] - **lib**: defer AbortSignal.any() following (sangwook) [#&#8203;62367](https://redirect.github.com/nodejs/node/pull/62367)
- \[[`b2981ec7eb`](https://redirect.github.com/nodejs/node/commit/b2981ec7eb)] - **meta**: bump actions/download-artifact from 8.0.0 to 8.0.1 (dependabot\[bot]) [#&#8203;62549](https://redirect.github.com/nodejs/node/pull/62549)
- \[[`7cd20667b5`](https://redirect.github.com/nodejs/node/commit/7cd20667b5)] - **meta**: bump github/codeql-action from 4.35.1 to 4.35.3 (dependabot\[bot]) [#&#8203;63074](https://redirect.github.com/nodejs/node/pull/63074)
- \[[`91a07cfe9f`](https://redirect.github.com/nodejs/node/commit/91a07cfe9f)] - **meta**: bump Mozilla-Actions/sccache-action from 0.0.9 to 0.0.10 (dependabot\[bot]) [#&#8203;63073](https://redirect.github.com/nodejs/node/pull/63073)
- \[[`09e17fe47c`](https://redirect.github.com/nodejs/node/commit/09e17fe47c)] - **meta**: add automation policy (Chengzhong Wu) [#&#8203;62871](https://redirect.github.com/nodejs/node/pull/62871)
- \[[`59e7fb7986`](https://redirect.github.com/nodejs/node/commit/59e7fb7986)] - **meta**: move VoltrexKeyva to emeritus (Matteo Collina) [#&#8203;62895](https://redirect.github.com/nodejs/node/pull/62895)
- \[[`1e2915cfa6`](https://redirect.github.com/nodejs/node/commit/1e2915cfa6)] - **meta**: bump peter-evans/create-pull-request from 8.1.0 to 8.1.1 (dependabot\[bot]) [#&#8203;62845](https://redirect.github.com/nodejs/node/pull/62845)
- \[[`0253c6e2be`](https://redirect.github.com/nodejs/node/commit/0253c6e2be)] - **meta**: bump step-security/harden-runner from 2.16.1 to 2.19.0 (dependabot\[bot]) [#&#8203;62844](https://redirect.github.com/nodejs/node/pull/62844)
- \[[`f503675b86`](https://redirect.github.com/nodejs/node/commit/f503675b86)] - **meta**: bump actions/setup-node from 6.3.0 to 6.4.0 (dependabot\[bot]) [#&#8203;62842](https://redirect.github.com/nodejs/node/pull/62842)
- \[[`5e14e4d26e`](https://redirect.github.com/nodejs/node/commit/5e14e4d26e)] - **meta**: broaden stale bot (Aviv Keller) [#&#8203;62658](https://redirect.github.com/nodejs/node/pull/62658)
- \[[`795db76f87`](https://redirect.github.com/nodejs/node/commit/795db76f87)] - **meta**: pass release version to release worker (flakey5) [#&#8203;62777](https://redirect.github.com/nodejs/node/pull/62777)
- \[[`ef384fe39f`](https://redirect.github.com/nodejs/node/commit/ef384fe39f)] - **meta**: add QUIC to CODEOWNERS (Tim Perry) [#&#8203;62652](https://redirect.github.com/nodejs/node/pull/62652)
- \[[`67e0ac568d`](https://redirect.github.com/nodejs/node/commit/67e0ac568d)] - **meta**: move Michael to emeritus (Michael Dawson) [#&#8203;62536](https://redirect.github.com/nodejs/node/pull/62536)
- \[[`5dad616393`](https://redirect.github.com/nodejs/node/commit/5dad616393)] - **meta**: populate apt list for slim runner in update-openssl workflow (René) [#&#8203;62628](https://redirect.github.com/nodejs/node/pull/62628)
- \[[`a869d25d8a`](https://redirect.github.com/nodejs/node/commit/a869d25d8a)] - **meta**: bump step-security/harden-runner from 2.15.0 to 2.16.1 (dependabot\[bot]) [#&#8203;62550](https://redirect.github.com/nodejs/node/pull/62550)
- \[[`769efc0403`](https://redirect.github.com/nodejs/node/commit/769efc0403)] - **meta**: bump actions/setup-node from 6.2.0 to 6.3.0 (dependabot\[bot]) [#&#8203;62548](https://redirect.github.com/nodejs/node/pull/62548)
- \[[`73fcc2b055`](https://redirect.github.com/nodejs/node/commit/73fcc2b055)] - **meta**: bump github/codeql-action from 4.32.4 to 4.35.1 (dependabot\[bot]) [#&#8203;62547](https://redirect.github.com/nodejs/node/pull/62547)
- \[[`6c001246fe`](https://redirect.github.com/nodejs/node/commit/6c001246fe)] - **meta**: bump codecov/codecov-action from 5.5.2 to 6.0.0 (dependabot\[bot]) [#&#8203;62545](https://redirect.github.com/nodejs/node/pull/62545)
- \[[`5ee40d6a03`](https://redirect.github.com/nodejs/node/commit/5ee40d6a03)] - **meta**: bump actions/cache from 5.0.3 to 5.0.4 (dependabot\[bot]) [#&#8203;62543](https://redirect.github.com/nodejs/node/pull/62543)
- \[[`ca16ad8a05`](https://redirect.github.com/nodejs/node/commit/ca16ad8a05)] - **meta**: require DCO signoff in commit message guidelines (James M Snell) [#&#8203;62510](https://redirect.github.com/nodejs/node/pull/62510)
- \[[`db9497fc41`](https://redirect.github.com/nodejs/node/commit/db9497fc41)] - **meta**: expand memory leak DoS criteria to all DoS (Joyee Cheung) [#&#8203;62505](https://redirect.github.com/nodejs/node/pull/62505)
- \[[`13b7d08b8d`](https://redirect.github.com/nodejs/node/commit/13b7d08b8d)] - **module**: remove duplicated checks from `_resolveFilename` (Antoine du Hamel) [#&#8203;62729](https://redirect.github.com/nodejs/node/pull/62729)
- \[[`6b53efb53a`](https://redirect.github.com/nodejs/node/commit/6b53efb53a)] - **module,win**: fix long subpath import (Stefan Stojanovic) [#&#8203;62101](https://redirect.github.com/nodejs/node/pull/62101)
- \[[`841dfbf6fc`](https://redirect.github.com/nodejs/node/commit/841dfbf6fc)] - **node-api**: update libuv ABI stability note (Chengzhong Wu) [#&#8203;62789](https://redirect.github.com/nodejs/node/pull/62789)
- \[[`01090f2aa1`](https://redirect.github.com/nodejs/node/commit/01090f2aa1)] - **node-api**: add napi\_create\_external\_sharedarraybuffer (Ben Noordhuis) [#&#8203;62623](https://redirect.github.com/nodejs/node/pull/62623)
- \[[`87443b4355`](https://redirect.github.com/nodejs/node/commit/87443b4355)] - **node-api**: execute tsfn finalizer after queue drains when aborted (Kevin Eady) [#&#8203;61956](https://redirect.github.com/nodejs/node/pull/61956)
- \[[`e95570c054`](https://redirect.github.com/nodejs/node/commit/e95570c054)] - **process**: handle rejections only when needed (Gürgün Dayıoğlu) [#&#8203;62919](https://redirect.github.com/nodejs/node/pull/62919)
- \[[`37d49f3219`](https://redirect.github.com/nodejs/node/commit/37d49f3219)] - **process**: optimize asyncHandledRejections by using FixedQueue (Gürgün Dayıoğlu) [#&#8203;60854](https://redirect.github.com/nodejs/node/pull/60854)
- \[[`f697c55e38`](https://redirect.github.com/nodejs/node/commit/f697c55e38)] - **quic**: add QuicEndpoint.listening & QuicStream.destroy() and tests (Tim Perry) [#&#8203;62648](https://redirect.github.com/nodejs/node/pull/62648)
- \[[`c128942b69`](https://redirect.github.com/nodejs/node/commit/c128942b69)] - **quic**: fixup token verification to handle zero expiration (James M Snell) [#&#8203;62620](https://redirect.github.com/nodejs/node/pull/62620)
- \[[`abb881ec92`](https://redirect.github.com/nodejs/node/commit/abb881ec92)] - **quic**: support multiple ALPN negotiation (James M Snell) [#&#8203;62620](https://redirect.github.com/nodejs/node/pull/62620)
- \[[`476926c2ad`](https://redirect.github.com/nodejs/node/commit/476926c2ad)] - **quic**: apply multiple TLS context improvements and SNI support (James M Snell) [#&#8203;62620](https://redirect.github.com/nodejs/node/pull/62620)
- \[[`76d9c24b95`](https://redirect.github.com/nodejs/node/commit/76d9c24b95)] - **quic**: implement rapidhash for hashing improvements (James M Snell) [#&#8203;62620](https://redirect.github.com/nodejs/node/pull/62620)
- \[[`08726cd43d`](https://redirect.github.com/nodejs/node/commit/08726cd43d)] - **quic**: move quic behind compile time flag (Matteo Collina) [#&#8203;61444](https://redirect.github.com/nodejs/node/pull/61444)
- \[[`ea4f19aaa7`](https://redirect.github.com/nodejs/node/commit/ea4f19aaa7)] - **quic**: use arena allocation for packets (James M Snell) [#&#8203;62589](https://redirect.github.com/nodejs/node/pull/62589)
- \[[`21e9239e2a`](https://redirect.github.com/nodejs/node/commit/21e9239e2a)] - **quic**: fixup linting/formatting issues (James M Snell) [#&#8203;62387](https://redirect.github.com/nodejs/node/pull/62387)
- \[[`edeed4303b`](https://redirect.github.com/nodejs/node/commit/edeed4303b)] - **quic**: update http3 impl details (James M Snell) [#&#8203;62387](https://redirect.github.com/nodejs/node/pull/62387)
- \[[`7f3a85e6aa`](https://redirect.github.com/nodejs/node/commit/7f3a85e6aa)] - **quic**: fix a handful of bugs and missing functionality (James M Snell) [#&#8203;62387](https://redirect.github.com/nodejs/node/pull/62387)
- \[[`45c1ebddf8`](https://redirect.github.com/nodejs/node/commit/45c1ebddf8)] - **quic**: copy options.certs buffer instead of detaching (Chengzhong Wu) [#&#8203;61403](https://redirect.github.com/nodejs/node/pull/61403)
- \[[`a31a8ee680`](https://redirect.github.com/nodejs/node/commit/a31a8ee680)] - **quic**: reduce boilerplate and other minor cleanups (James M Snell) [#&#8203;59342](https://redirect.github.com/nodejs/node/pull/59342)
- \[[`3be70ff43a`](https://redirect.github.com/nodejs/node/commit/3be70ff43a)] - **quic**: multiple fixups and updates (James M Snell) [#&#8203;59342](https://redirect.github.com/nodejs/node/pull/59342)
- \[[`b91a93444c`](https://redirect.github.com/nodejs/node/commit/b91a93444c)] - **quic**: update more of the quic to the new compile guard (James M Snell) [#&#8203;59342](https://redirect.github.com/nodejs/node/pull/59342)
- \[[`ca0080c164`](https://redirect.github.com/nodejs/node/commit/ca0080c164)] - **quic**: few additional small comment edits in cid.h (James M Snell) [#&#8203;59342](https://redirect.github.com/nodejs/node/pull/59342)
- \[[`6553202d83`](https://redirect.github.com/nodejs/node/commit/6553202d83)] - **quic**: fixup NO\_ERROR macro conflict on windows (James M Snell) [#&#8203;59381](https://redirect.github.com/nodejs/node/pull/59381)
- \[[`6df1508ac2`](https://redirect.github.com/nodejs/node/commit/6df1508ac2)] - **quic**: fixup windows coverage compile error (James M Snell) [#&#8203;59381](https://redirect.github.com/nodejs/node/pull/59381)
- \[[`b2b0bf8b04`](https://redirect.github.com/nodejs/node/commit/b2b0bf8b04)] - **quic**: update the guard to check openssl version (James M Snell) [#&#8203;59249](https://redirect.github.com/nodejs/node/pull/59249)
- \[[`5556b154bd`](https://redirect.github.com/nodejs/node/commit/5556b154bd)] - **quic**: start re-enabling quic with openssl 3.5 (James M Snell) [#&#8203;59249](https://redirect.github.com/nodejs/node/pull/59249)
- \[[`2ca42c8263`](https://redirect.github.com/nodejs/node/commit/2ca42c8263)] - **repl**: keep reference count for `process.on('newListener')` (Anna Henningsen) [#&#8203;61895](https://redirect.github.com/nodejs/node/pull/61895)
- \[[`2f37f9177f`](https://redirect.github.com/nodejs/node/commit/2f37f9177f)] - **sqlite**: use OneByte for ASCII text and internalize col names (Ali Hassan) [#&#8203;61954](https://redirect.github.com/nodejs/node/pull/61954)
- \[[`3c96ae1b2f`](https://redirect.github.com/nodejs/node/commit/3c96ae1b2f)] - **sqlite**: add serialize() and deserialize() (Ali Hassan) [#&#8203;62579](https://redirect.github.com/nodejs/node/pull/62579)
- \[[`be4d2f3a4c`](https://redirect.github.com/nodejs/node/commit/be4d2f3a4c)] - **sqlite**: enable Percentile extension (Jurj Andrei George) [#&#8203;61295](https://redirect.github.com/nodejs/node/pull/61295)
- \[[`dafed453b2`](https://redirect.github.com/nodejs/node/commit/dafed453b2)] - **src**: clean up experimental flag variables (Antoine du Hamel) [#&#8203;62759](https://redirect.github.com/nodejs/node/pull/62759)
- \[[`dca1e6aeea`](https://redirect.github.com/nodejs/node/commit/dca1e6aeea)] - **src**: expose help texts into node-config-schema.json (Pietro Marchini) [#&#8203;58680](https://redirect.github.com/nodejs/node/pull/58680)
- \[[`28c4f44eb1`](https://redirect.github.com/nodejs/node/commit/28c4f44eb1)] - **src**: add permission support to config file (Marco Ippolito) [#&#8203;60746](https://redirect.github.com/nodejs/node/pull/60746)
- \[[`f49175b220`](https://redirect.github.com/nodejs/node/commit/f49175b220)] - **src**: fix small compile warning in quic/streams.cc (James M Snell) [#&#8203;60118](https://redirect.github.com/nodejs/node/pull/60118)
- \[[`c9d4a446d8`](https://redirect.github.com/nodejs/node/commit/c9d4a446d8)] - **src**: cleanup quic TransportParams class (James M Snell) [#&#8203;59884](https://redirect.github.com/nodejs/node/pull/59884)
- \[[`99bb02fd9e`](https://redirect.github.com/nodejs/node/commit/99bb02fd9e)] - **src**: swap dotenv and config file parsing order (Marco Ippolito) [#&#8203;63035](https://redirect.github.com/nodejs/node/pull/63035)
- \[[`ecb4d49b7b`](https://redirect.github.com/nodejs/node/commit/ecb4d49b7b)] - **src**: add missing \<cstdlib> for abort() declaration (Charles Kerr) [#&#8203;63001](https://redirect.github.com/nodejs/node/pull/63001)
- \[[`b6219b6362`](https://redirect.github.com/nodejs/node/commit/b6219b6362)] - **src**: fix crash in GetErrorSource() for invalid using syntax (semimikoh) [#&#8203;62770](https://redirect.github.com/nodejs/node/pull/62770)
- \[[`b5ca5ad4c5`](https://redirect.github.com/nodejs/node/commit/b5ca5ad4c5)] - **src**: simplify `TCPWrap::Connect` signature (Anna Henningsen) [#&#8203;62929](https://redirect.github.com/nodejs/node/pull/62929)
- \[[`ef7ffce7cf`](https://redirect.github.com/nodejs/node/commit/ef7ffce7cf)] - **src**: use DCHECK in AsyncWrap::MakeCallback instead emiting a warning (Gerhard Stöbich) [#&#8203;62795](https://redirect.github.com/nodejs/node/pull/62795)
- \[[`cd9890a5ab`](https://redirect.github.com/nodejs/node/commit/cd9890a5ab)] - **src**: fix MaybeStackBuffer char\_traits deprecation warning (om-ghante) [#&#8203;62507](https://redirect.github.com/nodejs/node/pull/62507)
- \[[`c70ff44aee`](https://redirect.github.com/nodejs/node/commit/c70ff44aee)] - **src**: use context-free V8 message column getters (René) [#&#8203;62778](https://redirect.github.com/nodejs/node/pull/62778)
- \[[`06c405f1d7`](https://redirect.github.com/nodejs/node/commit/06c405f1d7)] - **src**: coerce `spawnSync` args to string once (Antoine du Hamel) [#&#8203;62633](https://redirect.github.com/nodejs/node/pull/62633)
- \[[`6151999ad6`](https://redirect.github.com/nodejs/node/commit/6151999ad6)] - **src**: use stack allocation for small string encoding (Ali Hassan) [#&#8203;62431](https://redirect.github.com/nodejs/node/pull/62431)
- \[[`a71a4ac7a3`](https://redirect.github.com/nodejs/node/commit/a71a4ac7a3)] - **src**: add contextify interceptor debug logs (Chengzhong Wu) [#&#8203;62460](https://redirect.github.com/nodejs/node/pull/62460)
- \[[`ad9a2909c2`](https://redirect.github.com/nodejs/node/commit/ad9a2909c2)] - **src**: workaround AIX libc++ std::filesystem bug (Richard Lau) [#&#8203;62788](https://redirect.github.com/nodejs/node/pull/62788)
- \[[`7792f1ae47`](https://redirect.github.com/nodejs/node/commit/7792f1ae47)] - **stream**: copyedit `webstreams/adapter.js` (Antoine du Hamel) [#&#8203;63034](https://redirect.github.com/nodejs/node/pull/63034)
- \[[`1397d8ce5c`](https://redirect.github.com/nodejs/node/commit/1397d8ce5c)] - **stream**: remove duplicated utility (Antoine du Hamel) [#&#8203;63031](https://redirect.github.com/nodejs/node/pull/63031)
- \[[`ff86b1d64f`](https://redirect.github.com/nodejs/node/commit/ff86b1d64f)] - **stream**: simplify `setPromiseHandled` utility (Antoine du Hamel) [#&#8203;63032](https://redirect.github.com/nodejs/node/pull/63032)
- \[[`24a078149a`](https://redirect.github.com/nodejs/node/commit/24a078149a)] - **stream**: validate ReadableStream.from iterator objects (Daeyeon Jeong) [#&#8203;62911](https://redirect.github.com/nodejs/node/pull/62911)
- \[[`cfb1fa9680`](https://redirect.github.com/nodejs/node/commit/cfb1fa9680)] - **stream**: reject duplicate nested transferables (Daeyeon Jeong) [#&#8203;62831](https://redirect.github.com/nodejs/node/pull/62831)
- \[[`d0c913758a`](https://redirect.github.com/nodejs/node/commit/d0c913758a)] - **stream**: ensuring cross-destruction in \_duplexify to prevent leaks (Daijiro Wachi) [#&#8203;62824](https://redirect.github.com/nodejs/node/pull/62824)
- \[[`978f5c15d7`](https://redirect.github.com/nodejs/node/commit/978f5c15d7)] - **stream**: simplify `readableStreamFromIterable` (Antoine du Hamel) [#&#8203;62651](https://redirect.github.com/nodejs/node/pull/62651)
- \[[`3527646ba5`](https://redirect.github.com/nodejs/node/commit/3527646ba5)] - **stream**: fix nested compose error propagation (Matteo Collina) [#&#8203;62556](https://redirect.github.com/nodejs/node/pull/62556)
- \[[`dfb9edef4f`](https://redirect.github.com/nodejs/node/commit/dfb9edef4f)] - **stream**: allow shared array buffer sources in writable webstream adapter (René) [#&#8203;62163](https://redirect.github.com/nodejs/node/pull/62163)
- \[[`f00cdab627`](https://redirect.github.com/nodejs/node/commit/f00cdab627)] - **stream**: simplify `createPromiseCallback` (Antoine du Hamel) [#&#8203;62650](https://redirect.github.com/nodejs/node/pull/62650)
- \[[`3ed783535f`](https://redirect.github.com/nodejs/node/commit/3ed783535f)] - **stream**: fix writev unhandled rejection in fromWeb (sangwook) [#&#8203;62297](https://redirect.github.com/nodejs/node/pull/62297)
- \[[`29b196694c`](https://redirect.github.com/nodejs/node/commit/29b196694c)] - **stream**: noop pause/resume on destroyed streams (Robert Nagy) [#&#8203;62557](https://redirect.github.com/nodejs/node/pull/62557)
- \[[`d73dbb9fc8`](https://redirect.github.com/nodejs/node/commit/d73dbb9fc8)] - **stream**: refactor duplexify to be less suceptible to prototype pollution (Antoine du Hamel) [#&#8203;62559](https://redirect.github.com/nodejs/node/pull/62559)
- \[[`6f37f7e240`](https://redirect.github.com/nodejs/node/commit/6f37f7e240)] - **(SEMVER-MINOR)** **stream**: propagate destruction in duplexPair (Ahmed Elhor) [#&#8203;61098](https://redirect.github.com/nodejs/node/pull/61098)
- \[[`b8816580e9`](https://redirect.github.com/nodejs/node/commit/b8816580e9)] - **test**: generate `localstorage.db` in a temp dir (Chengzhong Wu) [#&#8203;62660](https://redirect.github.com/nodejs/node/pull/62660)
- \[[`31a863fd29`](https://redirect.github.com/nodejs/node/commit/31a863fd29)] - **test**: update WPT for url to [`258f285`](https://redirect.github.com/nodejs/node/commit/258f285de0) (Node.js GitHub Bot) [#&#8203;63087](https://redirect.github.com/nodejs/node/pull/63087)
- \[[`d0d19bd8e3`](https://redirect.github.com/nodejs/node/commit/d0d19bd8e3)] - **test**: update WPT for streams to [`f8f26a3`](https://redirect.github.com/nodejs/node/commit/f8f26a372f) (Node.js GitHub Bot) [#&#8203;62864](https://redirect.github.com/nodejs/node/pull/62864)
- \[[`f50ac5bc78`](https://redirect.github.com/nodejs/node/commit/f50ac5bc78)] - **test**: improve config-file permission test coverage (Rafael Gonzaga) [#&#8203;60929](https://redirect.github.com/nodejs/node/pull/60929)
- \[[`a0f90000f4`](https://redirect.github.com/nodejs/node/commit/a0f90000f4)] - **test**: export isRiscv64 from common module (Jamie Magee) [#&#8203;62609](https://redirect.github.com/nodejs/node/pull/62609)
- \[[`da4dd8646f`](https://redirect.github.com/nodejs/node/commit/da4dd8646f)] - **test**: normalize known inspector crash as completion (Joyee Cheung) [#&#8203;62851](https://redirect.github.com/nodejs/node/pull/62851)
- \[[`b7fdd94a4c`](https://redirect.github.com/nodejs/node/commit/b7fdd94a4c)] - **test**: account for RFC 7919 FFDHE negotiation in OpenSSL 4.0 (Filip Skokan) [#&#8203;62805](https://redirect.github.com/nodejs/node/pull/62805)
- \[[`375a993aaf`](https://redirect.github.com/nodejs/node/commit/375a993aaf)] - **test**: skip tls-deprecated secp256k1 on OpenSSL 4.0 (Filip Skokan) [#&#8203;62805](https://redirect.github.com/nodejs/node/pull/62805)
- \[[`698d8287d1`](https://redirect.github.com/nodejs/node/commit/698d8287d1)] - **test**: use an always invalid cipher and cover OpenSSL 4.0 behaviours (Filip Skokan) [#&#8203;62805](https://redirect.github.com/nodejs/node/pull/62805)
- \[[`036bc6f300`](https://redirect.github.com/nodejs/node/commit/036bc6f300)] - **test**: use valid DER OCSP responses (Filip Skokan) [#&#8203;62805](https://redirect.github.com/nodejs/node/pull/62805)
- \[[`3aa9938da8`](https://redirect.github.com/nodejs/node/commit/3aa9938da8)] - **test**: skip test-tls-error-stack when engines are unsupported (Filip Skokan) [#&#8203;62805](https://redirect.github.com/nodejs/node/pull/62805)
- \[[`947f1ae246`](https://redirect.github.com/nodejs/node/commit/947f1ae246)] - **test**: accept renamed OpenSSL 4.0 error code and reason (Filip Skokan) [#&#8203;62805](https://redirect.github.com/nodejs/node/pull/62805)
- \[[`afdd355622`](https://redirect.github.com/nodejs/node/commit/afdd355622)] - **test**: update test/addons/openssl-binding for OpenSSL 4.0 (Filip Skokan) [#&#8203;62805](https://redirect.github.com/nodejs/node/pull/62805)
- \[[`8637524a99`](https://redirect.github.com/nodejs/node/commit/8637524a99)] - **test**: mark test-snapshot-reproducible flaky (Filip Skokan) [#&#8203;62808](https://redirect.github.com/nodejs/node/pull/62808)
- \[[`c22d34134b`](https://redirect.github.com/nodejs/node/commit/c22d34134b)] - **test**: check contextify contextual store behavior in strict mode (René) [#&#8203;62571](https://redirect.github.com/nodejs/node/pull/62571)
- \[[`0b4e0d3c94`](https://redirect.github.com/nodejs/node/commit/0b4e0d3c94)] - **test**: update tls junk data error expectations (Filip Skokan) [#&#8203;62629](https://redirect.github.com/nodejs/node/pull/62629)
- \[[`85d83c2cdb`](https://redirect.github.com/nodejs/node/commit/85d83c2cdb)] - **test**: ensure WPT report is in out/wpt (Filip Skokan) [#&#8203;62637](https://redirect.github.com/nodejs/node/pull/62637)
- \[[`9e21711c60`](https://redirect.github.com/nodejs/node/commit/9e21711c60)] - **test**: improve WPT runner summary (Filip Skokan) [#&#8203;62636](https://redirect.github.com/nodejs/node/pull/62636)
- \[[`e04e2c9ac1`](https://redirect.github.com/nodejs/node/commit/e04e2c9ac1)] - **test**: skip url WPT subtests instead of modifying test script (Filip Skokan) [#&#8203;62635](https://redirect.github.com/nodejs/node/pull/62635)
- \[[`7b1211f88c`](https://redirect.github.com/nodejs/node/commit/7b1211f88c)] - **test**: capture negative utimes mtime at call time (Yuya Inoue) [#&#8203;62490](https://redirect.github.com/nodejs/node/pull/62490)
- \[[`f1a6e9fcc7`](https://redirect.github.com/nodejs/node/commit/f1a6e9fcc7)] - **test**: allow skipping individual WPT subtests (Filip Skokan) [#&#8203;62517](https://redirect.github.com/nodejs/node/pull/62517)
- \[[`23f927542e`](https://redirect.github.com/nodejs/node/commit/23f927542e)] - **test**: use on-disk fixture for test-npm-install (Joyee Cheung) [#&#8203;62584](https://redirect.github.com/nodejs/node/pull/62584)
- \[[`4739c45879`](https://redirect.github.com/nodejs/node/commit/4739c45879)] - **test**: update WPT for url to [`7a3645b`](https://redirect.github.com/nodejs/node/commit/7a3645b79a) (Node.js GitHub Bot) [#&#8203;62591](https://redirect.github.com/nodejs/node/pull/62591)
- \[[`f68189b839`](https://redirect.github.com/nodejs/node/commit/f68189b839)] - **test\_runner**: add `testId` to test events (Moshe Atlow) [#&#8203;62772](https://redirect.github.com/nodejs/node/pull/62772)
- \[[`5c2770446e`](https://redirect.github.com/nodejs/node/commit/5c2770446e)] - **test\_runner**: publish to TracingChannel for OTel instrumentation (Moshe Atlow) [#&#8203;62502](https://redirect.github.com/nodejs/node/pull/62502)
- \[[`d14029be7f`](https://redirect.github.com/nodejs/node/commit/d14029be7f)] - **(SEMVER-MINOR)** **test\_runner**: support test order randomization (Pietro Marchini) [#&#8203;61747](https://redirect.github.com/nodejs/node/pull/61747)
- \[[`3f74a58979`](https://redirect.github.com/nodejs/node/commit/3f74a58979)] - **test\_runner**: update node-config-schema (Pietro Marchini) [#&#8203;58680](https://redirect.github.com/nodejs/node/pull/58680)
- \[[`60c83f6199`](https://redirect.github.com/nodejs/node/commit/60c83f6199)] - **test\_runner**: fix failing suite hooks when marked with `todo` (Moshe Atlow) [#&#8203;63097](https://redirect.github.com/nodejs/node/pull/63097)
- \[[`d142c584cd`](https://redirect.github.com/nodejs/node/commit/d142c584cd)] - **(SEMVER-MINOR)** **test\_runner**: align mock timeout api (sangwook) [#&#8203;62820](https://redirect.github.com/nodejs/node/pull/62820)
- \[[`3e72065ed6`](https://redirect.github.com/nodejs/node/commit/3e72065ed6)] - **test\_runner**: fix suite rerun edge case (Moshe Atlow) [#&#8203;62860](https://redirect.github.com/nodejs/node/pull/62860)
- \[[`01a9552585`](https://redirect.github.com/nodejs/node/commit/01a9552585)] - **(SEMVER-MINOR)** **test\_runner**: add mock-timers support for AbortSignal.timeout (DeveloperViraj) [#&#8203;60751](https://redirect.github.com/nodejs/node/pull/60751)
- \[[`dd43efffa6`](https://redirect.github.com/nodejs/node/commit/dd43efffa6)] - **test\_runner**: add passed, attempt, and diagnostic to SuiteContext (Moshe Atlow) [#&#8203;62504](https://redirect.github.com/nodejs/node/pull/62504)
- \[[`a12dc445cc`](https://redirect.github.com/nodejs/node/commit/a12dc445cc)] - **tools**: add a check for clean git tree after tests (Antoine du Hamel) [#&#8203;62661](https://redirect.github.com/nodejs/node/pull/62661)
- \[[`5b49178375`](https://redirect.github.com/nodejs/node/commit/5b49178375)] - **tools**: use LTS Node.js in notify-on-push workflow (Nenad Spasenic) [#&#8203;63084](https://redirect.github.com/nodejs/node/pull/63084)
- \[[`5a93bde5bb`](https://redirect.github.com/nodejs/node/commit/5a93bde5bb)] - **tools**: update gr2m/create-or-update-pull-request-action to v1.10.1 (Mike McCready) [#&#8203;63065](https://redirect.github.com/nodejs/node/pull/63065)
- \[[`b133019d19`](https://redirect.github.com/nodejs/node/commit/b133019d19)] - **tools**: simplify `update-undici.sh` (Antoine du Hamel) [#&#8203;63044](https://redirect.github.com/nodejs/node/pull/63044)
- \[[`04d3538074`](https://redirect.github.com/nodejs/node/commit/04d3538074)] - **tools**: do not run `test-linux` on unrelated tools changes (Antoine du Hamel) [#&#8203;63037](https://redirect.github.com/nodejs/node/pull/63037)
- \[[`4d396ac4a5`](https://redirect.github.com/nodejs/node/commit/4d396ac4a5)] - **tools**: bump the eslint group in /tools/eslint with 4 updates (dependabot\[bot]) [#&#8203;62848](https://redirect.github.com/nodejs/node/pull/62848)
- \[[`9354bf40e7`](https://redirect.github.com/nodejs/node/commit/9354bf40e7)] - **tools**: update gyp-next to 0.22.1 (Node.js GitHub Bot) [#&#8203;62961](https://redirect.github.com/nodejs/node/pull/62961)
- \[[`c23db1ca85`](https://redirect.github.com/nodejs/node/commit/c23db1ca85)] - **tools**: fix commit linter for semver-major release proposals (Antoine du Hamel) [#&#8203;62993](https://redirect.github.com/nodejs/node/pull/62993)
- \[[`6e097ee3f1`](https://redirect.github.com/nodejs/node/commit/6e097ee3f1)] - **tools**: consolidate and simplify .editorconfig deps section (Daijiro Wachi) [#&#8203;62887](https://redirect.github.com/nodejs/node/pull/62887)
- \[[`a47ea6d6ea`](https://redirect.github.com/nodejs/node/commit/a47ea6d6ea)] - **tools**: set bot as author of tools-deps-update PRs (Antoine du Hamel) [#&#8203;62856](https://redirect.github.com/nodejs/node/pull/62856)
- \[[`00e86f0471`](https://redirect.github.com/nodejs/node/commit/00e86f0471)] - **tools**: bump brace-expansion from 5.0.4 to 5.0.5 in /tools/eslint (dependabot\[bot]) [#&#8203;62458](https://redirect.github.com/nodejs/node/pull/62458)
- \[[`cd7e262e75`](https://redirect.github.com/nodejs/node/commit/cd7e262e75)] - **tools**: bump brace-expansion in /tools/clang-format (dependabot\[bot]) [#&#8203;62467](https://redirect.github.com/nodejs/node/pull/62467)
- \[[`bfc1319bc8`](https://redirect.github.com/nodejs/node/commit/bfc1319bc8)] - **tools**: exclude [@&#8203;node-core/doc-kit](https://redirect.github.com/node-core/doc-kit) from dependabot cooldown (Levi Zim) [#&#8203;62775](https://redirect.github.com/nodejs/node/pull/62775)
- \[[`a932fbd10b`](https://redirect.github.com/nodejs/node/commit/a932fbd10b)] - **tools**: re-enable undici WPTs in daily wpt.fyi job (Filip Skokan) [#&#8203;62677](https://redirect.github.com/nodejs/node/pull/62677)
- \[[`f7bd9e3055`](https://redirect.github.com/nodejs/node/commit/f7bd9e3055)] - **tools**: update gyp-next to 0.22.0 (Node.js GitHub Bot) [#&#8203;62697](https://redirect.github.com/nodejs/node/pull/62697)
- \[[`c400d46d87`](https://redirect.github.com/nodejs/node/commit/c400d46d87)] - **tools**: improve backport review script (Antoine du Hamel) [#&#8203;62573](https://redirect.github.com/nodejs/node/pull/62573)
- \[[`be23b75814`](https://redirect.github.com/nodejs/node/commit/be23b75814)] - **tools**: improve output for unexpected passes in WTP tests (Antoine du Hamel) [#&#8203;62587](https://redirect.github.com/nodejs/node/pull/62587)
- \[[`609c013ece`](https://redirect.github.com/nodejs/node/commit/609c013ece)] - **tools**: revert OpenSSL update workflow to ubuntu-latest (Richard Lau) [#&#8203;62627](https://redirect.github.com/nodejs/node/pull/62627)
- \[[`81bac1ebfd`](https://redirect.github.com/nodejs/node/commit/81bac1ebfd)] - **tools**: bump the eslint group in /tools/eslint with 2 updates (dependabot\[bot]) [#&#8203;62552](https://redirect.github.com/nodejs/node/pull/62552)
- \[[`1fee26522d`](https://redirect.github.com/nodejs/node/commit/1fee26522d)] - **tools**: allow triagers to queue a PR for CI until it's reviewed (Antoine du Hamel) [#&#8203;62524](https://redirect.github.com/nodejs/node/pull/62524)
- \[[`332088f929`](https://redirect.github.com/nodejs/node/commit/332088f929)] - **tools**: do not run `commit-lint` on release proposals (Antoine du Hamel) [#&#8203;62523](https://redirect.github.com/nodejs/node/pull/62523)
- \[[`9a25fc8a4d`](https://redirect.github.com/nodejs/node/commit/9a25fc8a4d)] - **url**: process crash via malformed UNC hostname in pathToFileURL() (Nicola Del Gobbo) [#&#8203;62574](https://redirect.github.com/nodejs/node/pull/62574)
- \[[`7bd08ff60a`](https://redirect.github.com/nodejs/node/commit/7bd08ff60a)] - **url**: optimize URLSearchParams set/delete duplicate handling (Gürgün Dayıoğlu) [#&#8203;62266](https://redirect.github.com/nodejs/node/pull/62266)
- \[[`2d636388fa`](https://redirect.github.com/nodejs/node/commit/2d636388fa)] - **url**: align default argument handling for URLPattern with webidl (Filip Skokan) [#&#8203;62719](https://redirect.github.com/nodejs/node/pull/62719)
- \[[`00705a459a`](https://redirect.github.com/nodejs/node/commit/00705a459a)] - **(SEMVER-MINOR)** **util**: colorize text with hex colors (Guilherme Araújo) [#&#8203;61556](https://redirect.github.com/nodejs/node/pull/61556)
- \[[`0e2adb3e45`](https://redirect.github.com/nodejs/node/commit/0e2adb3e45)] - **watch**: track worker entry files in watch mode (SudhansuBandha) [#&#8203;62368](https://redirect.github.com/nodejs/node/pull/62368)
- \[[`c58fe38211`](https://redirect.github.com/nodejs/node/commit/c58fe38211)] - **watch**: fix --env-file-if-exists crashing on linux if the file is missing (Efe) [#&#8203;61870](https://redirect.github.com/nodejs/node/pull/61870)

</details>

---

### Configuration

📅 **Schedule**: Branch creation - "on Monday after 3am and before 10am" (UTC), Automerge - At any time (no schedule defined).

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0Mi45OS4wLXJwbSIsInVwZGF0ZWRJblZlciI6IjQyLjk5LjAtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6W119-->


---

## Discussion

### Comment by @swadeley on June 08, 2026 at 05:47 PM UTC

Hi @dominikvagner, are you happy for browsers to be installed serially? See my commit just now.

### Comment by @red-hat-konflux on June 15, 2026 at 05:30 PM UTC

### Edited/Blocked Notification

Renovate will not automatically rebase this PR, because it does not recognize the last commit author and assumes somebody else may have edited the PR.

You can manually request rebase by checking the rebase/retry box above.

 ⚠️ **Warning**: custom changes will be lost.

---

## Reviews

### Review by @swadeley - Dismissed on May 26, 2026 at 07:22 AM UTC

build fails

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1502*
