---
type: pull_request
number: 1401
title: "Build: Update Node.js to v24.14.0"
state: merged
author: red-hat-konflux
created: 2026-03-02T05:18:49Z
updated: 2026-03-02T08:30:55Z
closed: 2026-03-02T08:30:54Z
merged: 2026-03-02T08:30:54Z
base_branch: main
head_branch: konflux/mintmaker/main/node-24.x
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1401
---

# Pull Request #1401: Build: Update Node.js to v24.14.0

**Author**: @red-hat-konflux
**Created**: March 02, 2026 at 05:18 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/node-24.x`

## Description

This PR contains the following updates:

| Package | Update | Change |
|---|---|---|
| [node](https://nodejs.org) ([source](https://redirect.github.com/nodejs/node)) | minor | `24.13.1` -> `24.14.0` |

---

### Release Notes

<details>
<summary>nodejs/node (node)</summary>

### [`v24.14.0`](https://redirect.github.com/nodejs/node/releases/tag/v24.14.0): 2026-02-24, Version 24.14.0 &#x27;Krypton&#x27; (LTS), @&#8203;ruyadorno prepared by @&#8203;aduh95

[Compare Source](https://redirect.github.com/nodejs/node/compare/v24.13.1...v24.14.0)

##### Notable Changes

- \[[`8b6d31d379`](https://redirect.github.com/nodejs/node/commit/8b6d31d379)] - **(SEMVER-MINOR)** **async\_hooks**: add `trackPromises` option to `createHook()` (Joyee Cheung) [#&#8203;61415](https://redirect.github.com/nodejs/node/pull/61415)
- \[[`68da144b4e`](https://redirect.github.com/nodejs/node/commit/68da144b4e)] - **build,deps**: replace cjs-module-lexer with merve (Yagiz Nizipli) [#&#8203;61456](https://redirect.github.com/nodejs/node/pull/61456)
- \[[`f3a24c76e4`](https://redirect.github.com/nodejs/node/commit/f3a24c76e4)] - **(SEMVER-MINOR)** **deps**: add LIEF as a dependency (Joyee Cheung) [#&#8203;61167](https://redirect.github.com/nodejs/node/pull/61167)
- \[[`1948861d23`](https://redirect.github.com/nodejs/node/commit/1948861d23)] - **(SEMVER-MINOR)** **events**: repurpose `events.listenerCount()` to accept EventTargets (René) [#&#8203;60214](https://redirect.github.com/nodejs/node/pull/60214)
- \[[`d6f7c8d06f`](https://redirect.github.com/nodejs/node/commit/d6f7c8d06f)] - **(SEMVER-MINOR)** **fs**: add `ignore` option to `fs.watch` (Matteo Collina) [#&#8203;61433](https://redirect.github.com/nodejs/node/pull/61433)
- \[[`cb54b3ca6e`](https://redirect.github.com/nodejs/node/commit/cb54b3ca6e)] - **(SEMVER-MINOR)** **http**: add `http.setGlobalProxyFromEnv()` (Joyee Cheung) [#&#8203;60953](https://redirect.github.com/nodejs/node/pull/60953)
- \[[`35b1759d06`](https://redirect.github.com/nodejs/node/commit/35b1759d06)] - **(SEMVER-MINOR)** **module**: allow subpath imports that start with `#/` (Jan Martin) [#&#8203;60864](https://redirect.github.com/nodejs/node/pull/60864)
- \[[`2d72ea66f2`](https://redirect.github.com/nodejs/node/commit/2d72ea66f2)] - **(SEMVER-MINOR)** **process**: preserve `AsyncLocalStorage` in `queueMicrotask` only when needed (Gürgün Dayıoğlu) [#&#8203;60913](https://redirect.github.com/nodejs/node/pull/60913)
- \[[`6f4a4f6c8e`](https://redirect.github.com/nodejs/node/commit/6f4a4f6c8e)] - **(SEMVER-MINOR)** **sea**: split sea binary manipulation code (Joyee Cheung) [#&#8203;61167](https://redirect.github.com/nodejs/node/pull/61167)
- \[[`c0ceb9b065`](https://redirect.github.com/nodejs/node/commit/c0ceb9b065)] - **(SEMVER-MINOR)** **sqlite**: enable defensive mode by default (Bart Louwers) [#&#8203;61266](https://redirect.github.com/nodejs/node/pull/61266)
- \[[`33d8e8303b`](https://redirect.github.com/nodejs/node/commit/33d8e8303b)] - **(SEMVER-MINOR)** **sqlite**: add sqlite prepare options args (Guilherme Araújo) [#&#8203;61311](https://redirect.github.com/nodejs/node/pull/61311)
- \[[`563ab699eb`](https://redirect.github.com/nodejs/node/commit/563ab699eb)] - **(SEMVER-MINOR)** **src**: add initial support for ESM in embedder API (Joyee Cheung) [#&#8203;61548](https://redirect.github.com/nodejs/node/pull/61548)
- \[[`4c80031000`](https://redirect.github.com/nodejs/node/commit/4c80031000)] - **(SEMVER-MINOR)** **stream**: add `bytes()` method to `node:stream/consumers` (wantaek) [#&#8203;60426](https://redirect.github.com/nodejs/node/pull/60426)
- \[[`f5233df4ff`](https://redirect.github.com/nodejs/node/commit/f5233df4ff)] - **(SEMVER-MINOR)** **stream**: do not pass `readable.compose()` output via `Readable.from()` (René) [#&#8203;60907](https://redirect.github.com/nodejs/node/pull/60907)
- \[[`345a40fda3`](https://redirect.github.com/nodejs/node/commit/345a40fda3)] - **(SEMVER-MINOR)** **test**: use fixture directories for sea tests (Joyee Cheung) [#&#8203;61167](https://redirect.github.com/nodejs/node/pull/61167)
- \[[`972f82411d`](https://redirect.github.com/nodejs/node/commit/972f82411d)] - **(SEMVER-MINOR)** **test\_runner**: add `env` option to `run` function (Ethan Arrowood) [#&#8203;61367](https://redirect.github.com/nodejs/node/pull/61367)
- \[[`d77f98c4b6`](https://redirect.github.com/nodejs/node/commit/d77f98c4b6)] - **(SEMVER-MINOR)** **test\_runner**: support expecting a test-case to fail (Jacob Smith) [#&#8203;60669](https://redirect.github.com/nodejs/node/pull/60669)
- \[[`8e900af6ba`](https://redirect.github.com/nodejs/node/commit/8e900af6ba)] - **(SEMVER-MINOR)** **util**: add `convertProcessSignalToExitCode` utility (Erick Wendel) [#&#8203;60963](https://redirect.github.com/nodejs/node/pull/60963)

##### Commits

- \[[`180778fb9a`](https://redirect.github.com/nodejs/node/commit/180778fb9a)] - **assert**: fix loose deepEqual arrays with undefined and null failing (Ruben Bridgewater) [#&#8203;61587](https://redirect.github.com/nodejs/node/pull/61587)
- \[[`8b6d31d379`](https://redirect.github.com/nodejs/node/commit/8b6d31d379)] - **(SEMVER-MINOR)** **async\_hooks**: add trackPromises option to createHook() (Joyee Cheung) [#&#8203;61415](https://redirect.github.com/nodejs/node/pull/61415)
- \[[`83bcd38d35`](https://redirect.github.com/nodejs/node/commit/83bcd38d35)] - **benchmark**: add streaming TextDecoder benchmark (Сковорода Никита Андреевич) [#&#8203;61549](https://redirect.github.com/nodejs/node/pull/61549)
- \[[`4c105844c5`](https://redirect.github.com/nodejs/node/commit/4c105844c5)] - **build**: add support for Visual Studio 2026 (Michaël Zasso) [#&#8203;60727](https://redirect.github.com/nodejs/node/pull/60727)
- \[[`1f84fd91d9`](https://redirect.github.com/nodejs/node/commit/1f84fd91d9)] - **build**: skip sscache action on non-main branches (Joyee Cheung) [#&#8203;61790](https://redirect.github.com/nodejs/node/pull/61790)
- \[[`30601b680f`](https://redirect.github.com/nodejs/node/commit/30601b680f)] - **build**: add `--shared-nbytes` configure flag (Antoine du Hamel) [#&#8203;61341](https://redirect.github.com/nodejs/node/pull/61341)
- \[[`c6253eda49`](https://redirect.github.com/nodejs/node/commit/c6253eda49)] - **build**: add `--shared-hdr-histogram` configure flag (Antoine du Hamel) [#&#8203;61280](https://redirect.github.com/nodejs/node/pull/61280)
- \[[`584c189037`](https://redirect.github.com/nodejs/node/commit/584c189037)] - **build**: add `--shared-gtest` configure flag (Antoine du Hamel) [#&#8203;61279](https://redirect.github.com/nodejs/node/pull/61279)
- \[[`5998987881`](https://redirect.github.com/nodejs/node/commit/5998987881)] - **build**: aix: deoptimize implementation-visitor.cc with --shared (Stewart X Addison) [#&#8203;61550](https://redirect.github.com/nodejs/node/pull/61550)
- \[[`68da144b4e`](https://redirect.github.com/nodejs/node/commit/68da144b4e)] - **build,deps**: replace cjs-module-lexer with merve (Yagiz Nizipli) [#&#8203;61456](https://redirect.github.com/nodejs/node/pull/61456)
- \[[`6a4511bafb`](https://redirect.github.com/nodejs/node/commit/6a4511bafb)] - **build,win**: fix vs2022 compilation (Stefan Stojanovic) [#&#8203;61530](https://redirect.github.com/nodejs/node/pull/61530)
- \[[`2d6735db8a`](https://redirect.github.com/nodejs/node/commit/2d6735db8a)] - **deps**: upgrade npm to 11.9.0 (npm team) [#&#8203;61685](https://redirect.github.com/nodejs/node/pull/61685)
- \[[`699e2f8f81`](https://redirect.github.com/nodejs/node/commit/699e2f8f81)] - **deps**: update amaro to 1.1.7 (Node.js GitHub Bot) [#&#8203;61730](https://redirect.github.com/nodejs/node/pull/61730)
- \[[`7be76316d6`](https://redirect.github.com/nodejs/node/commit/7be76316d6)] - **deps**: update minimatch to 10.1.2 (Node.js GitHub Bot) [#&#8203;61732](https://redirect.github.com/nodejs/node/pull/61732)
- \[[`97e5a65013`](https://redirect.github.com/nodejs/node/commit/97e5a65013)] - **deps**: update undici to 7.21.0 (Node.js GitHub Bot) [#&#8203;61683](https://redirect.github.com/nodejs/node/pull/61683)
- \[[`74e4710ee7`](https://redirect.github.com/nodejs/node/commit/74e4710ee7)] - **deps**: update googletest to [`56efe39`](https://redirect.github.com/nodejs/node/commit/56efe3983185e3f37e43415d1afa97e3860f187f) (Node.js GitHub Bot) [#&#8203;61605](https://redirect.github.com/nodejs/node/pull/61605)
- \[[`b5113e2a2a`](https://redirect.github.com/nodejs/node/commit/b5113e2a2a)] - **deps**: update amaro to 1.1.6 (Node.js GitHub Bot) [#&#8203;61603](https://redirect.github.com/nodejs/node/pull/61603)
- \[[`f3a24c76e4`](https://redirect.github.com/nodejs/node/commit/f3a24c76e4)] - **(SEMVER-MINOR)** **deps**: add LIEF as a dependency (Joyee Cheung) [#&#8203;61167](https://redirect.github.com/nodejs/node/pull/61167)
- \[[`c370c3dc06`](https://redirect.github.com/nodejs/node/commit/c370c3dc06)] - **(SEMVER-MINOR)** **deps**: add tools and scripts to pull LIEF as a dependency (Joyee Cheung) [#&#8203;61167](https://redirect.github.com/nodejs/node/pull/61167)
- \[[`e54975e17d`](https://redirect.github.com/nodejs/node/commit/e54975e17d)] - **deps**: V8: cherry-pick [highway@`dcc0ca1`](https://redirect.github.com/highway/node/commit/dcc0ca1cd42) (Richard Lau) [#&#8203;61008](https://redirect.github.com/nodejs/node/pull/61008)
- \[[`625b90b76b`](https://redirect.github.com/nodejs/node/commit/625b90b76b)] - **deps**: update undici to 7.19.2 (Node.js GitHub Bot) [#&#8203;61566](https://redirect.github.com/nodejs/node/pull/61566)
- \[[`05e9a9fb5e`](https://redirect.github.com/nodejs/node/commit/05e9a9fb5e)] - **deps**: update undici to 7.19.1 (Node.js GitHub Bot) [#&#8203;61514](https://redirect.github.com/nodejs/node/pull/61514)
- \[[`3d41643e38`](https://redirect.github.com/nodejs/node/commit/3d41643e38)] - **deps**: update undici to 7.19.0 (Node.js GitHub Bot) [#&#8203;61470](https://redirect.github.com/nodejs/node/pull/61470)
- \[[`17b363a66c`](https://redirect.github.com/nodejs/node/commit/17b363a66c)] - **dns**: fix Windows SRV ECONNREFUSED by adjusting c-ares fallback detection (notvivek12) [#&#8203;61453](https://redirect.github.com/nodejs/node/pull/61453)
- \[[`33d0a8c22d`](https://redirect.github.com/nodejs/node/commit/33d0a8c22d)] - **doc**: clarify EventEmitter error handling in threat model (Matteo Collina) [#&#8203;61701](https://redirect.github.com/nodejs/node/pull/61701)
- \[[`5b8e72cf85`](https://redirect.github.com/nodejs/node/commit/5b8e72cf85)] - **doc**: mention default option for test runner env (Steven) [#&#8203;61659](https://redirect.github.com/nodejs/node/pull/61659)
- \[[`f44e67fac2`](https://redirect.github.com/nodejs/node/commit/f44e67fac2)] - **doc**: fix --inspect security warning section (Tim Perry) [#&#8203;61675](https://redirect.github.com/nodejs/node/pull/61675)
- \[[`a0e09c9043`](https://redirect.github.com/nodejs/node/commit/a0e09c9043)] - **doc**: document `url.format(urlString)` as deprecated under DEP0169 (René) [#&#8203;61644](https://redirect.github.com/nodejs/node/pull/61644)
- \[[`5e719248fe`](https://redirect.github.com/nodejs/node/commit/5e719248fe)] - **doc**: deprecation add more codemod (Augustin Mauroy) [#&#8203;61642](https://redirect.github.com/nodejs/node/pull/61642)
- \[[`8f5a3e5df4`](https://redirect.github.com/nodejs/node/commit/8f5a3e5df4)] - **doc**: fix grammatical error in README.md (ayj8201) [#&#8203;61653](https://redirect.github.com/nodejs/node/pull/61653)
- \[[`d52b535163`](https://redirect.github.com/nodejs/node/commit/d52b535163)] - **doc**: correct tools README Boxstarter link (Mike McCready) [#&#8203;61638](https://redirect.github.com/nodejs/node/pull/61638)
- \[[`4889dc4f59`](https://redirect.github.com/nodejs/node/commit/4889dc4f59)] - **doc**: update `server.dropMaxConnection` link (YuSheng Chen) [#&#8203;61584](https://redirect.github.com/nodejs/node/pull/61584)
- \[[`8e48e72f2a`](https://redirect.github.com/nodejs/node/commit/8e48e72f2a)] - **doc**: clean up writing-and-running-benchmarks.md (Hardanish Singh) [#&#8203;61345](https://redirect.github.com/nodejs/node/pull/61345)
- \[[`1948861d23`](https://redirect.github.com/nodejs/node/commit/1948861d23)] - **(SEMVER-MINOR)** **events**: repurpose `events.listenerCount()` to accept EventTargets (René) [#&#8203;60214](https://redirect.github.com/nodejs/node/pull/60214)
- \[[`d6f7c8d06f`](https://redirect.github.com/nodejs/node/commit/d6f7c8d06f)] - **(SEMVER-MINOR)** **fs**: add ignore option to fs.watch (Matteo Collina) [#&#8203;61433](https://redirect.github.com/nodejs/node/pull/61433)
- \[[`2d7e5f9581`](https://redirect.github.com/nodejs/node/commit/2d7e5f9581)] - **http**: implement slab allocation for HTTP header parsing (Mert Can Altin) [#&#8203;61375](https://redirect.github.com/nodejs/node/pull/61375)
- \[[`cb54b3ca6e`](https://redirect.github.com/nodejs/node/commit/cb54b3ca6e)] - **(SEMVER-MINOR)** **http**: add http.setGlobalProxyFromEnv() (Joyee Cheung) [#&#8203;60953](https://redirect.github.com/nodejs/node/pull/60953)
- \[[`6df8be48ce`](https://redirect.github.com/nodejs/node/commit/6df8be48ce)] - **lib**: use utf8 fast path for streaming TextDecoder (Сковорода Никита Андреевич) [#&#8203;61549](https://redirect.github.com/nodejs/node/pull/61549)
- \[[`830fff0aca`](https://redirect.github.com/nodejs/node/commit/830fff0aca)] - **lib**: recycle queues (Robert Nagy) [#&#8203;61461](https://redirect.github.com/nodejs/node/pull/61461)
- \[[`069874bdbd`](https://redirect.github.com/nodejs/node/commit/069874bdbd)] - **lib**: use StringPrototypeStartsWith from primordials in locks (Taejin Kim) [#&#8203;61492](https://redirect.github.com/nodejs/node/pull/61492)
- \[[`7824c7589e`](https://redirect.github.com/nodejs/node/commit/7824c7589e)] - **lib**: unify ICU and no-ICU TextDecoder (Сковорода Никита Андреевич) [#&#8203;61409](https://redirect.github.com/nodejs/node/pull/61409)
- \[[`f81430702a`](https://redirect.github.com/nodejs/node/commit/f81430702a)] - **lib**: prefer `call()` over `apply()` if argument list is not array (Livia Medeiros) [#&#8203;60796](https://redirect.github.com/nodejs/node/pull/60796)
- \[[`a723f72e1e`](https://redirect.github.com/nodejs/node/commit/a723f72e1e)] - **lib**: add support for readable byte streams to .toWeb() (Hans Klunder) [#&#8203;58664](https://redirect.github.com/nodejs/node/pull/58664)
- \[[`b78d814b3d`](https://redirect.github.com/nodejs/node/commit/b78d814b3d)] - **meta**: persist sccache daemon until end of build workflows (René) [#&#8203;61639](https://redirect.github.com/nodejs/node/pull/61639)
- \[[`40a872a4b9`](https://redirect.github.com/nodejs/node/commit/40a872a4b9)] - **meta**: bump github/codeql-action from 4.31.9 to 4.32.0 (dependabot\[bot]) [#&#8203;61622](https://redirect.github.com/nodejs/node/pull/61622)
- \[[`0637bdb3be`](https://redirect.github.com/nodejs/node/commit/0637bdb3be)] - **meta**: bump step-security/harden-runner from 2.14.0 to 2.14.1 (dependabot\[bot]) [#&#8203;61621](https://redirect.github.com/nodejs/node/pull/61621)
- \[[`e8d9bd9fc5`](https://redirect.github.com/nodejs/node/commit/e8d9bd9fc5)] - **meta**: bump actions/setup-python from 6.1.0 to 6.2.0 (dependabot\[bot]) [#&#8203;61627](https://redirect.github.com/nodejs/node/pull/61627)
- \[[`c517df2b65`](https://redirect.github.com/nodejs/node/commit/c517df2b65)] - **meta**: bump actions/setup-node from 6.1.0 to 6.2.0 (dependabot\[bot]) [#&#8203;61625](https://redirect.github.com/nodejs/node/pull/61625)
- \[[`9a64f2f25d`](https://redirect.github.com/nodejs/node/commit/9a64f2f25d)] - **meta**: bump actions/cache from 5.0.1 to 5.0.3 (dependabot\[bot]) [#&#8203;61624](https://redirect.github.com/nodejs/node/pull/61624)
- \[[`0e5922e95e`](https://redirect.github.com/nodejs/node/commit/0e5922e95e)] - **meta**: bump peter-evans/create-pull-request from 8.0.0 to 8.1.0 (dependabot\[bot]) [#&#8203;61623](https://redirect.github.com/nodejs/node/pull/61623)
- \[[`5da7b51091`](https://redirect.github.com/nodejs/node/commit/5da7b51091)] - **meta**: bump actions/stale from 10.1.0 to 10.1.1 (dependabot\[bot]) [#&#8203;61620](https://redirect.github.com/nodejs/node/pull/61620)
- \[[`c085c8a43f`](https://redirect.github.com/nodejs/node/commit/c085c8a43f)] - **meta**: bump actions/checkout from 6.0.1 to 6.0.2 (dependabot\[bot]) [#&#8203;61619](https://redirect.github.com/nodejs/node/pull/61619)
- \[[`ce2acf0275`](https://redirect.github.com/nodejs/node/commit/ce2acf0275)] - **meta**: bump actions/download-artifact from 6.0.0 to 7.0.0 (dependabot\[bot]) [#&#8203;61242](https://redirect.github.com/nodejs/node/pull/61242)
- \[[`629f0eaac5`](https://redirect.github.com/nodejs/node/commit/629f0eaac5)] - **meta**: bump actions/checkout from 6.0.0 to 6.0.1 (dependabot\[bot]) [#&#8203;61239](https://redirect.github.com/nodejs/node/pull/61239)
- \[[`cd80d369c9`](https://redirect.github.com/nodejs/node/commit/cd80d369c9)] - **meta**: bump actions/upload-artifact from 5.0.0 to 6.0.0 (dependabot\[bot]) [#&#8203;61238](https://redirect.github.com/nodejs/node/pull/61238)
- \[[`8c75e4e1fa`](https://redirect.github.com/nodejs/node/commit/8c75e4e1fa)] - **meta**: bump actions/checkout from 5.0.1 to 6.0.0 (dependabot\[bot]) [#&#8203;60925](https://redirect.github.com/nodejs/node/pull/60925)
- \[[`5a9e9f4127`](https://redirect.github.com/nodejs/node/commit/5a9e9f4127)] - **meta**: bump actions/checkout from 5.0.0 to 5.0.1 (dependabot\[bot]) [#&#8203;60767](https://redirect.github.com/nodejs/node/pull/60767)
- \[[`1519251dd1`](https://redirect.github.com/nodejs/node/commit/1519251dd1)] - **module**: do not invoke resolve hooks twice for imported cjs (Joyee Cheung) [#&#8203;61529](https://redirect.github.com/nodejs/node/pull/61529)
- \[[`8d7190b3fe`](https://redirect.github.com/nodejs/node/commit/8d7190b3fe)] - **module**: do not wrap module.\_load when tracing is not enabled (Joyee Cheung) [#&#8203;61479](https://redirect.github.com/nodejs/node/pull/61479)
- \[[`35b1759d06`](https://redirect.github.com/nodejs/node/commit/35b1759d06)] - **(SEMVER-MINOR)** **module**: allow subpath imports that start with `#/` (Jan Martin) [#&#8203;60864](https://redirect.github.com/nodejs/node/pull/60864)
- \[[`7a83b38921`](https://redirect.github.com/nodejs/node/commit/7a83b38921)] - **net**: defer synchronous destroy calls in internalConnect (RajeshKumar11) [#&#8203;61658](https://redirect.github.com/nodejs/node/pull/61658)
- \[[`16bab79421`](https://redirect.github.com/nodejs/node/commit/16bab79421)] - **process**: do not truncate long strings in `--print` (Mohamed Akram) [#&#8203;61497](https://redirect.github.com/nodejs/node/pull/61497)
- \[[`2d72ea66f2`](https://redirect.github.com/nodejs/node/commit/2d72ea66f2)] - **(SEMVER-MINOR)** **process**: preserve AsyncLocalStorage in queueMicrotask only when needed (Gürgün Dayıoğlu) [#&#8203;60913](https://redirect.github.com/nodejs/node/pull/60913)
- \[[`9cc1c4604f`](https://redirect.github.com/nodejs/node/commit/9cc1c4604f)] - **repl**: fix getters triggering side effects during completion (Dario Piotrowicz) [#&#8203;61043](https://redirect.github.com/nodejs/node/pull/61043)
- \[[`93703306a1`](https://redirect.github.com/nodejs/node/commit/93703306a1)] - **repl**: tab completion targets `<class>` instead of `new <class>` (Đỗ Trọng Hải) [#&#8203;60319](https://redirect.github.com/nodejs/node/pull/60319)
- \[[`6f4a4f6c8e`](https://redirect.github.com/nodejs/node/commit/6f4a4f6c8e)] - **(SEMVER-MINOR)** **sea**: split sea binary manipulation code (Joyee Cheung) [#&#8203;61167](https://redirect.github.com/nodejs/node/pull/61167)
- \[[`46a2dad4db`](https://redirect.github.com/nodejs/node/commit/46a2dad4db)] - **sqlite**: avoid extra copy for large text binds (Ali Hassan) [#&#8203;61580](https://redirect.github.com/nodejs/node/pull/61580)
- \[[`f91a377f7e`](https://redirect.github.com/nodejs/node/commit/f91a377f7e)] - **sqlite**: use DictionaryTemplate for run() result (Mert Can Altin) [#&#8203;61432](https://redirect.github.com/nodejs/node/pull/61432)
- \[[`0e7571ae3e`](https://redirect.github.com/nodejs/node/commit/0e7571ae3e)] - **sqlite**: change approach to fix segfault SQLTagStore (Bart Louwers) [#&#8203;60462](https://redirect.github.com/nodejs/node/pull/60462)
- \[[`8e8f70524a`](https://redirect.github.com/nodejs/node/commit/8e8f70524a)] - **sqlite**: reserve vectors space (Guilherme Araújo) [#&#8203;61540](https://redirect.github.com/nodejs/node/pull/61540)
- \[[`c0ceb9b065`](https://redirect.github.com/nodejs/node/commit/c0ceb9b065)] - **(SEMVER-MINOR)** **sqlite**: enable defensive mode by default (Bart Louwers) [#&#8203;61266](https://redirect.github.com/nodejs/node/pull/61266)
- \[[`33d8e8303b`](https://redirect.github.com/nodejs/node/commit/33d8e8303b)] - **(SEMVER-MINOR)** **sqlite**: add sqlite prepare options args (Guilherme Araújo) [#&#8203;61311](https://redirect.github.com/nodejs/node/pull/61311)
- \[[`f0d8f37002`](https://redirect.github.com/nodejs/node/commit/f0d8f37002)] - **src**: elide heap allocation in structured clone implementation (Anna Henningsen) [#&#8203;61703](https://redirect.github.com/nodejs/node/pull/61703)
- \[[`db478c4336`](https://redirect.github.com/nodejs/node/commit/db478c4336)] - **src**: use simdutf for one-byte string UTF-8 write in stringBytes (Mert Can Altin) [#&#8203;61696](https://redirect.github.com/nodejs/node/pull/61696)
- \[[`563ab699eb`](https://redirect.github.com/nodejs/node/commit/563ab699eb)] - **(SEMVER-MINOR)** **src**: add initial support for ESM in embedder API (Joyee Cheung) [#&#8203;61548](https://redirect.github.com/nodejs/node/pull/61548)
- \[[`da13186a15`](https://redirect.github.com/nodejs/node/commit/da13186a15)] - **src**: throw RangeError on failed ArrayBuffer BackingStore allocation (Chengzhong Wu) [#&#8203;61480](https://redirect.github.com/nodejs/node/pull/61480)
- \[[`4c80031000`](https://redirect.github.com/nodejs/node/commit/4c80031000)] - **(SEMVER-MINOR)** **stream**: add bytes() method to stream/consumers (wantaek) [#&#8203;60426](https://redirect.github.com/nodejs/node/pull/60426)
- \[[`f5233df4ff`](https://redirect.github.com/nodejs/node/commit/f5233df4ff)] - **(SEMVER-MINOR)** **stream**: do not pass `readable.compose()` output via `Readable.from()` (René) [#&#8203;60907](https://redirect.github.com/nodejs/node/pull/60907)
- \[[`ad04a469c8`](https://redirect.github.com/nodejs/node/commit/ad04a469c8)] - **test**: restraint version replacement pattern in snapshots (Chengzhong Wu) [#&#8203;61748](https://redirect.github.com/nodejs/node/pull/61748)
- \[[`2d3b4a8d65`](https://redirect.github.com/nodejs/node/commit/2d3b4a8d65)] - **test**: print stack immediately avoiding GC interleaving (Chengzhong Wu) [#&#8203;61699](https://redirect.github.com/nodejs/node/pull/61699)
- \[[`38f43a6d4e`](https://redirect.github.com/nodejs/node/commit/38f43a6d4e)] - **test**: fix case-insensitive path matching on Windows (Matteo Collina) [#&#8203;61682](https://redirect.github.com/nodejs/node/pull/61682)
- \[[`06513f5ff2`](https://redirect.github.com/nodejs/node/commit/06513f5ff2)] - **test**: fix flaky test-performance-eventloopdelay (Matteo Collina) [#&#8203;61629](https://redirect.github.com/nodejs/node/pull/61629)
- \[[`9d79c66c61`](https://redirect.github.com/nodejs/node/commit/9d79c66c61)] - **test**: remove duplicate wpt tests (Filip Skokan) [#&#8203;61617](https://redirect.github.com/nodejs/node/pull/61617)
- \[[`eac9f4f401`](https://redirect.github.com/nodejs/node/commit/eac9f4f401)] - **test**: fix race condition in watch mode tests (Matteo Collina) [#&#8203;61615](https://redirect.github.com/nodejs/node/pull/61615)
- \[[`ecf5947575`](https://redirect.github.com/nodejs/node/commit/ecf5947575)] - **test**: update WPT for url to [`e3c46fd`](https://redirect.github.com/nodejs/node/commit/e3c46fdf55) (Node.js GitHub Bot) [#&#8203;61602](https://redirect.github.com/nodejs/node/pull/61602)
- \[[`356ff5fece`](https://redirect.github.com/nodejs/node/commit/356ff5fece)] - **test**: use the skipIfNoWatch() utility function (Luigi Pinca) [#&#8203;61531](https://redirect.github.com/nodejs/node/pull/61531)
- \[[`4b2187aea2`](https://redirect.github.com/nodejs/node/commit/4b2187aea2)] - **test**: unify assertSnapshot common patterns (Chengzhong Wu) [#&#8203;61590](https://redirect.github.com/nodejs/node/pull/61590)
- \[[`8c25489d63`](https://redirect.github.com/nodejs/node/commit/8c25489d63)] - **test**: split test-fs-watch-ignore-\* (Luigi Pinca) [#&#8203;61494](https://redirect.github.com/nodejs/node/pull/61494)
- \[[`43b8a2b7e7`](https://redirect.github.com/nodejs/node/commit/43b8a2b7e7)] - **test**: add some validation for JSON doc output (Antoine du Hamel) [#&#8203;61413](https://redirect.github.com/nodejs/node/pull/61413)
- \[[`345a40fda3`](https://redirect.github.com/nodejs/node/commit/345a40fda3)] - **(SEMVER-MINOR)** **test**: use fixture directories for sea tests (Joyee Cheung) [#&#8203;61167](https://redirect.github.com/nodejs/node/pull/61167)
- \[[`24cf6b8326`](https://redirect.github.com/nodejs/node/commit/24cf6b8326)] - **test**: reveal wpt evaluation errors in status files (Chengzhong Wu) [#&#8203;61358](https://redirect.github.com/nodejs/node/pull/61358)
- \[[`d4034dfb62`](https://redirect.github.com/nodejs/node/commit/d4034dfb62)] - **test**: forbid use of named imports for fixtures (Antoine du Hamel) [#&#8203;61228](https://redirect.github.com/nodejs/node/pull/61228)
- \[[`4f871ee897`](https://redirect.github.com/nodejs/node/commit/4f871ee897)] - **test**: enforce better never-settling-promise detection (Antoine du Hamel) [#&#8203;60976](https://redirect.github.com/nodejs/node/pull/60976)
- \[[`8e9adedf02`](https://redirect.github.com/nodejs/node/commit/8e9adedf02)] - **test**: ensure assertions are reached on all tests (Antoine du Hamel) [#&#8203;60845](https://redirect.github.com/nodejs/node/pull/60845)
- \[[`273832802e`](https://redirect.github.com/nodejs/node/commit/273832802e)] - **test**: ensure assertions are reached on more tests (Antoine du Hamel) [#&#8203;60763](https://redirect.github.com/nodejs/node/pull/60763)
- \[[`e06adcb52f`](https://redirect.github.com/nodejs/node/commit/e06adcb52f)] - **test**: ensure assertions are reached on more tests (Antoine du Hamel) [#&#8203;60760](https://redirect.github.com/nodejs/node/pull/60760)
- \[[`aeed0ccc02`](https://redirect.github.com/nodejs/node/commit/aeed0ccc02)] - **test**: use `RegExp.escape` to improve test reliability (Antoine du Hamel) [#&#8203;60803](https://redirect.github.com/nodejs/node/pull/60803)
- \[[`74bcd0adab`](https://redirect.github.com/nodejs/node/commit/74bcd0adab)] - **test**: ensure assertions are reached on more tests (Antoine du Hamel) [#&#8203;60728](https://redirect.github.com/nodejs/node/pull/60728)
- \[[`407807b08e`](https://redirect.github.com/nodejs/node/commit/407807b08e)] - **test**: skip tests not passing without `NODE_OPTIONS` support (Antoine du Hamel) [#&#8203;60912](https://redirect.github.com/nodejs/node/pull/60912)
- \[[`a9e70cefb0`](https://redirect.github.com/nodejs/node/commit/a9e70cefb0)] - **test**: ensure assertions are reached on more tests (Antoine du Hamel) [#&#8203;60634](https://redirect.github.com/nodejs/node/pull/60634)
- \[[`21b23cd0d0`](https://redirect.github.com/nodejs/node/commit/21b23cd0d0)] - **test\_runner**: fix test enqueue when test file has syntax error (Edy Silva) [#&#8203;61573](https://redirect.github.com/nodejs/node/pull/61573)
- \[[`6a4de694b2`](https://redirect.github.com/nodejs/node/commit/6a4de694b2)] - **test\_runner**: fix passing `expectFailure` (Moshe Atlow) [#&#8203;61568](https://redirect.github.com/nodejs/node/pull/61568)
- \[[`6640de2b0f`](https://redirect.github.com/nodejs/node/commit/6640de2b0f)] - **test\_runner**: differentiate todo and failure styles (Moshe Atlow) [#&#8203;61564](https://redirect.github.com/nodejs/node/pull/61564)
- \[[`972f82411d`](https://redirect.github.com/nodejs/node/commit/972f82411d)] - **(SEMVER-MINOR)** **test\_runner**: add env option to run function (Ethan Arrowood) [#&#8203;61367](https://redirect.github.com/nodejs/node/pull/61367)
- \[[`d77f98c4b6`](https://redirect.github.com/nodejs/node/commit/d77f98c4b6)] - **(SEMVER-MINOR)** **test\_runner**: support expecting a test-case to fail (Jacob Smith) [#&#8203;60669](https://redirect.github.com/nodejs/node/pull/60669)
- \[[`f98986cbb9`](https://redirect.github.com/nodejs/node/commit/f98986cbb9)] - **tools**: switch to ARM runners on GHA jobs (Antoine du Hamel) [#&#8203;61903](https://redirect.github.com/nodejs/node/pull/61903)
- \[[`034589dd93`](https://redirect.github.com/nodejs/node/commit/034589dd93)] - **tools**: avoid building twice in coverage jobs (Antoine du Hamel) [#&#8203;61899](https://redirect.github.com/nodejs/node/pull/61899)
- \[[`e50e2f00bb`](https://redirect.github.com/nodejs/node/commit/e50e2f00bb)] - **tools**: use ubuntu-slim runner in GHA (Antoine du Hamel) [#&#8203;61759](https://redirect.github.com/nodejs/node/pull/61759)
- \[[`f658f48ccb`](https://redirect.github.com/nodejs/node/commit/f658f48ccb)] - **tools**: use ubuntu-slim runner in GHA (Antoine du Hamel) [#&#8203;61734](https://redirect.github.com/nodejs/node/pull/61734)
- \[[`65c77d74ff`](https://redirect.github.com/nodejs/node/commit/65c77d74ff)] - **tools**: use ubuntu-latest runner in `notify-on-push` workflow (Antoine du Hamel) [#&#8203;61742](https://redirect.github.com/nodejs/node/pull/61742)
- \[[`605905556a`](https://redirect.github.com/nodejs/node/commit/605905556a)] - **tools**: enforce removal of `lts-watch-*` labels on release proposals (Antoine du Hamel) [#&#8203;61672](https://redirect.github.com/nodejs/node/pull/61672)
- \[[`f0f98d4c03`](https://redirect.github.com/nodejs/node/commit/f0f98d4c03)] - **tools**: use ubuntu-slim runner in meta GitHub Actions (Tierney Cyren) [#&#8203;61663](https://redirect.github.com/nodejs/node/pull/61663)
- \[[`ab63ddf354`](https://redirect.github.com/nodejs/node/commit/ab63ddf354)] - **tools**: add LIEF to license builder (Chengzhong Wu) [#&#8203;61523](https://redirect.github.com/nodejs/node/pull/61523)
- \[[`8a0f6192c9`](https://redirect.github.com/nodejs/node/commit/8a0f6192c9)] - **tools**: enforce trailing commas in `test/es-module` (Antoine du Hamel) [#&#8203;60891](https://redirect.github.com/nodejs/node/pull/60891)
- \[[`4afbbcf39e`](https://redirect.github.com/nodejs/node/commit/4afbbcf39e)] - **tools**: enforce trailing commas in `test/sequential` (Antoine du Hamel) [#&#8203;60892](https://redirect.github.com/nodejs/node/pull/60892)
- \[[`4c1abf752c`](https://redirect.github.com/nodejs/node/commit/4c1abf752c)] - **tools,win**: upgrade install additional tools to Visual Studio 2026 (Mike McCready) [#&#8203;61562](https://redirect.github.com/nodejs/node/pull/61562)
- \[[`8e900af6ba`](https://redirect.github.com/nodejs/node/commit/8e900af6ba)] - **(SEMVER-MINOR)** **util**: add convertProcessSignalToExitCode utility (Erick Wendel) [#&#8203;60963](https://redirect.github.com/nodejs/node/pull/60963)

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0Mi4yNi41LXJwbSIsInVwZGF0ZWRJblZlciI6IjQyLjI2LjUtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6W119-->


---

## Reviews

### Review by @swadeley - Approved on March 02, 2026 at 08:30 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1401*
