---
type: pull_request
number: 1473
title: "Build chore(deps): update node.js to v24.15.0"
state: merged
author: red-hat-konflux
created: 2026-04-20T05:28:25Z
updated: 2026-04-20T07:28:25Z
closed: 2026-04-20T07:28:23Z
merged: 2026-04-20T07:28:23Z
base_branch: main
head_branch: konflux/mintmaker/main/node-24.x
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1473
---

# Pull Request #1473: Build chore(deps): update node.js to v24.15.0

**Author**: @red-hat-konflux
**Created**: April 20, 2026 at 05:28 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/node-24.x`

## Description

> ℹ️ **Note**
> 
> This PR body was truncated due to platform limits.

This PR contains the following updates:

| Package | Update | Change |
|---|---|---|
| [node](https://nodejs.org) ([source](https://redirect.github.com/nodejs/node)) | minor | `24.14.1` → `24.15.0` |

---

### Release Notes

<details>
<summary>nodejs/node (node)</summary>

### [`v24.15.0`](https://redirect.github.com/nodejs/node/releases/tag/v24.15.0): 2026-04-15, Version 24.15.0 &#x27;Krypton&#x27; (LTS), @&#8203;aduh95

[Compare Source](https://redirect.github.com/nodejs/node/compare/v24.14.1...v24.15.0)

##### Notable Changes

- \[[`3d87ecacbc`](https://redirect.github.com/nodejs/node/commit/3d87ecacbc)] - **(SEMVER-MINOR)** **cli**: add --max-heap-size option (tannal) [#&#8203;58708](https://redirect.github.com/nodejs/node/pull/58708)
- \[[`83c38672f7`](https://redirect.github.com/nodejs/node/commit/83c38672f7)] - **cli**: add --require-module/--no-require-module (Joyee Cheung) [#&#8203;60959](https://redirect.github.com/nodejs/node/pull/60959)
- \[[`54ef940e01`](https://redirect.github.com/nodejs/node/commit/54ef940e01)] - **(SEMVER-MINOR)** **crypto**: add raw key formats support to the KeyObject APIs (Filip Skokan) [#&#8203;62240](https://redirect.github.com/nodejs/node/pull/62240)
- \[[`f4a3edc47a`](https://redirect.github.com/nodejs/node/commit/f4a3edc47a)] - **(SEMVER-MINOR)** **fs**: add `throwIfNoEntry` option for fs.stat and fs.promises.stat (Juan José) [#&#8203;61178](https://redirect.github.com/nodejs/node/pull/61178)
- \[[`5cdcba17cc`](https://redirect.github.com/nodejs/node/commit/5cdcba17cc)] - **(SEMVER-MINOR)** **http2**: add http1Options for HTTP/1 fallback configuration (Amol Yadav) [#&#8203;61713](https://redirect.github.com/nodejs/node/pull/61713)
- \[[`8b6be3fe14`](https://redirect.github.com/nodejs/node/commit/8b6be3fe14)] - **module**: mark require(esm) as stable (Joyee Cheung) [#&#8203;60959](https://redirect.github.com/nodejs/node/pull/60959)
- \[[`68fbc0c6cc`](https://redirect.github.com/nodejs/node/commit/68fbc0c6cc)] - **module**: mark module compile cache as stable (Joyee Cheung) [#&#8203;60971](https://redirect.github.com/nodejs/node/pull/60971)
- \[[`c851e76f8c`](https://redirect.github.com/nodejs/node/commit/c851e76f8c)] - **(SEMVER-MINOR)** **net**: add `setTOS` and `getTOS` to `Socket` (Amol Yadav) [#&#8203;61503](https://redirect.github.com/nodejs/node/pull/61503)
- \[[`6ac4304c87`](https://redirect.github.com/nodejs/node/commit/6ac4304c87)] - **(SEMVER-MINOR)** **sqlite**: add limits property to DatabaseSync (Mert Can Altin) [#&#8203;61298](https://redirect.github.com/nodejs/node/pull/61298)
- \[[`aaf9af1672`](https://redirect.github.com/nodejs/node/commit/aaf9af1672)] - **sqlite**: mark as release candidate (Matteo Collina) [#&#8203;61262](https://redirect.github.com/nodejs/node/pull/61262)
- \[[`eb77a7a297`](https://redirect.github.com/nodejs/node/commit/eb77a7a297)] - **(SEMVER-MINOR)** **src**: add C++ support for diagnostics channels (RafaelGSS) [#&#8203;61869](https://redirect.github.com/nodejs/node/pull/61869)
- \[[`6834ca13bb`](https://redirect.github.com/nodejs/node/commit/6834ca13bb)] - **(SEMVER-MINOR)** **stream**: rename `Duplex.toWeb()` type option to `readableType` (René) [#&#8203;61632](https://redirect.github.com/nodejs/node/pull/61632)
- \[[`f5f21d36a6`](https://redirect.github.com/nodejs/node/commit/f5f21d36a6)] - **test\_runner**: add exports option for module mocks (sangwook) [#&#8203;61727](https://redirect.github.com/nodejs/node/pull/61727)
- \[[`1f2025fd1e`](https://redirect.github.com/nodejs/node/commit/1f2025fd1e)] - **(SEMVER-MINOR)** **test\_runner**: expose worker ID for concurrent test execution (Ali Hassan) [#&#8203;61394](https://redirect.github.com/nodejs/node/pull/61394)
- \[[`1ca20fc33d`](https://redirect.github.com/nodejs/node/commit/1ca20fc33d)] - **(SEMVER-MINOR)** **test\_runner**: show interrupted test on SIGINT (Matteo Collina) [#&#8203;61676](https://redirect.github.com/nodejs/node/pull/61676)

##### Commits

- \[[`148373cea1`](https://redirect.github.com/nodejs/node/commit/148373cea1)] - **assert,util**: improve comparison performance (Ruben Bridgewater) [#&#8203;61176](https://redirect.github.com/nodejs/node/pull/61176)
- \[[`e5558b0859`](https://redirect.github.com/nodejs/node/commit/e5558b0859)] - **assert,util**: fix deep comparing invalid dates skipping properties (Ruben Bridgewater) [#&#8203;61076](https://redirect.github.com/nodejs/node/pull/61076)
- \[[`83cffd92b5`](https://redirect.github.com/nodejs/node/commit/83cffd92b5)] - **async\_hooks**: enabledHooksExist shall return if hooks are enabled (Gerhard Stöbich) [#&#8203;61054](https://redirect.github.com/nodejs/node/pull/61054)
- \[[`2c9436b43d`](https://redirect.github.com/nodejs/node/commit/2c9436b43d)] - **benchmark**: fix destructuring in dgram/single-buffer (Ali Hassan) [#&#8203;62084](https://redirect.github.com/nodejs/node/pull/62084)
- \[[`837acd7382`](https://redirect.github.com/nodejs/node/commit/837acd7382)] - **benchmark**: add startup benchmark for ESM entrypoint (Joyee Cheung) [#&#8203;61769](https://redirect.github.com/nodejs/node/pull/61769)
- \[[`a6ced7d272`](https://redirect.github.com/nodejs/node/commit/a6ced7d272)] - **buffer**: improve performance of multiple Buffer operations (Ali Hassan) [#&#8203;61871](https://redirect.github.com/nodejs/node/pull/61871)
- \[[`a82003bf8b`](https://redirect.github.com/nodejs/node/commit/a82003bf8b)] - **buffer**: optimize buffer.concat performance (Mert Can Altin) [#&#8203;61721](https://redirect.github.com/nodejs/node/pull/61721)
- \[[`83dfd0be1d`](https://redirect.github.com/nodejs/node/commit/83dfd0be1d)] - **buffer**: disallow ArrayBuffer transfer on pooled buffer (Chengzhong Wu) [#&#8203;61372](https://redirect.github.com/nodejs/node/pull/61372)
- \[[`ed2d0cb1bf`](https://redirect.github.com/nodejs/node/commit/ed2d0cb1bf)] - **build**: support empty libname flags in `configure.py` (Antoine du Hamel) [#&#8203;62477](https://redirect.github.com/nodejs/node/pull/62477)
- \[[`09f7920267`](https://redirect.github.com/nodejs/node/commit/09f7920267)] - **build**: fix timezone-update path references (Chengzhong Wu) [#&#8203;62280](https://redirect.github.com/nodejs/node/pull/62280)
- \[[`af46b15b91`](https://redirect.github.com/nodejs/node/commit/af46b15b91)] - **build**: use path-ignore in GHA coverage-windows.yml (Chengzhong Wu) [#&#8203;61811](https://redirect.github.com/nodejs/node/pull/61811)
- \[[`2cf77eadd1`](https://redirect.github.com/nodejs/node/commit/2cf77eadd1)] - **build**: generate\_config\_gypi.py generates valid JSON (Shelley Vohr) [#&#8203;61791](https://redirect.github.com/nodejs/node/pull/61791)
- \[[`e0220f0c35`](https://redirect.github.com/nodejs/node/commit/e0220f0c35)] - **build**: build with v8 gdbjit support on supported platform (Joyee Cheung) [#&#8203;61010](https://redirect.github.com/nodejs/node/pull/61010)
- \[[`5505511dcb`](https://redirect.github.com/nodejs/node/commit/5505511dcb)] - **build**: enable -DV8\_ENABLE\_CHECKS flag (Ryuhei Shima) [#&#8203;61327](https://redirect.github.com/nodejs/node/pull/61327)
- \[[`5f8ecf3940`](https://redirect.github.com/nodejs/node/commit/5f8ecf3940)] - **build**: add --debug-symbols to build with -g without enabling DCHECKs (Joyee Cheung) [#&#8203;61100](https://redirect.github.com/nodejs/node/pull/61100)
- \[[`ab18c0867b`](https://redirect.github.com/nodejs/node/commit/ab18c0867b)] - **build**: fix --node-builtin-modules-path (Filip Skokan) [#&#8203;62115](https://redirect.github.com/nodejs/node/pull/62115)
- \[[`bfa60d5782`](https://redirect.github.com/nodejs/node/commit/bfa60d5782)] - **build**: fix GN for new merve dep (Shelley Vohr) [#&#8203;61984](https://redirect.github.com/nodejs/node/pull/61984)
- \[[`0d1975fe3a`](https://redirect.github.com/nodejs/node/commit/0d1975fe3a)] - **build,win**: add WinGet Visual Studio 2022 Build Tools Edition config (Mike McCready) [#&#8203;61652](https://redirect.github.com/nodejs/node/pull/61652)
- \[[`10b2bb5fa6`](https://redirect.github.com/nodejs/node/commit/10b2bb5fa6)] - **child\_process**: add tracing channel for spawn (Marco) [#&#8203;61836](https://redirect.github.com/nodejs/node/pull/61836)
- \[[`3d87ecacbc`](https://redirect.github.com/nodejs/node/commit/3d87ecacbc)] - **(SEMVER-MINOR)** **cli**: add --max-heap-size option (tannal) [#&#8203;58708](https://redirect.github.com/nodejs/node/pull/58708)
- \[[`83c38672f7`](https://redirect.github.com/nodejs/node/commit/83c38672f7)] - **cli**: add --require-module/--no-require-module (Joyee Cheung) [#&#8203;60959](https://redirect.github.com/nodejs/node/pull/60959)
- \[[`9d37233824`](https://redirect.github.com/nodejs/node/commit/9d37233824)] - **crypto**: update root certificates to NSS 3.121 (Node.js GitHub Bot) [#&#8203;62485](https://redirect.github.com/nodejs/node/pull/62485)
- \[[`b0cbfe38a4`](https://redirect.github.com/nodejs/node/commit/b0cbfe38a4)] - **crypto**: add crypto::GetSSLCtx API for addon access to OpenSSL contexts (Tim Perry) [#&#8203;62254](https://redirect.github.com/nodejs/node/pull/62254)
- \[[`dc034a4ac9`](https://redirect.github.com/nodejs/node/commit/dc034a4ac9)] - **crypto**: reject ML-KEM/ML-DSA [PKCS#8](https://redirect.github.com/PKCS/node/issues/8) import without seed in SubtleCrypto (Filip Skokan) [#&#8203;62218](https://redirect.github.com/nodejs/node/pull/62218)
- \[[`8aa6e706df`](https://redirect.github.com/nodejs/node/commit/8aa6e706df)] - **crypto**: refactor WebCrypto AEAD algorithms auth tag handling (Filip Skokan) [#&#8203;62169](https://redirect.github.com/nodejs/node/pull/62169)
- \[[`20cb932bcf`](https://redirect.github.com/nodejs/node/commit/20cb932bcf)] - **crypto**: read algorithm name property only once in normalizeAlgorithm (Filip Skokan) [#&#8203;62170](https://redirect.github.com/nodejs/node/pull/62170)
- \[[`e2934162b4`](https://redirect.github.com/nodejs/node/commit/e2934162b4)] - **crypto**: add missing AES dictionaries (Filip Skokan) [#&#8203;62099](https://redirect.github.com/nodejs/node/pull/62099)
- \[[`8b8db52f65`](https://redirect.github.com/nodejs/node/commit/8b8db52f65)] - **crypto**: fix importKey required argument count check (Filip Skokan) [#&#8203;62099](https://redirect.github.com/nodejs/node/pull/62099)
- \[[`bd5458db29`](https://redirect.github.com/nodejs/node/commit/bd5458db29)] - **crypto**: fix missing nullptr check on RSA\_new() (ndossche) [#&#8203;61888](https://redirect.github.com/nodejs/node/pull/61888)
- \[[`7302c7ed22`](https://redirect.github.com/nodejs/node/commit/7302c7ed22)] - **crypto**: fix handling of null BUF\_MEM\* in ToV8Value() (Nora Dossche) [#&#8203;61885](https://redirect.github.com/nodejs/node/pull/61885)
- \[[`8d0c22ea20`](https://redirect.github.com/nodejs/node/commit/8d0c22ea20)] - **crypto**: fix potential null pointer dereference when BIO\_meth\_new() fails (Nora Dossche) [#&#8203;61788](https://redirect.github.com/nodejs/node/pull/61788)
- \[[`72aad8b40f`](https://redirect.github.com/nodejs/node/commit/72aad8b40f)] - **crypto**: always return certificate serial numbers as uppercase (Anna Henningsen) [#&#8203;61752](https://redirect.github.com/nodejs/node/pull/61752)
- \[[`2395fc0f4d`](https://redirect.github.com/nodejs/node/commit/2395fc0f4d)] - **crypto**: rename CShakeParams and KmacParams length to outputLength (Filip Skokan) [#&#8203;61875](https://redirect.github.com/nodejs/node/pull/61875)
- \[[`541be3aaf2`](https://redirect.github.com/nodejs/node/commit/541be3aaf2)] - **crypto**: recognize raw formats in keygen (Filip Skokan) [#&#8203;62480](https://redirect.github.com/nodejs/node/pull/62480)
- \[[`54ef940e01`](https://redirect.github.com/nodejs/node/commit/54ef940e01)] - **(SEMVER-MINOR)** **crypto**: add raw key formats support to the KeyObject APIs (Filip Skokan) [#&#8203;62240](https://redirect.github.com/nodejs/node/pull/62240)
- \[[`bef1949823`](https://redirect.github.com/nodejs/node/commit/bef1949823)] - **deps**: V8: cherry-pick [`33e7739`](https://redirect.github.com/nodejs/node/commit/33e7739c134d) (Thibaud Michaud) [#&#8203;62567](https://redirect.github.com/nodejs/node/pull/62567)
- \[[`2e1a565a55`](https://redirect.github.com/nodejs/node/commit/2e1a565a55)] - **deps**: update ada to 3.4.4 (Node.js GitHub Bot) [#&#8203;62414](https://redirect.github.com/nodejs/node/pull/62414)
- \[[`d0418bad10`](https://redirect.github.com/nodejs/node/commit/d0418bad10)] - **deps**: update timezone to 2026a (Node.js GitHub Bot) [#&#8203;62164](https://redirect.github.com/nodejs/node/pull/62164)
- \[[`53aad66415`](https://redirect.github.com/nodejs/node/commit/53aad66415)] - **deps**: update googletest to [`2461743`](https://redirect.github.com/nodejs/node/commit/2461743991f9aa53e9a3625eafcbacd81a3c74cd) (Node.js GitHub Bot) [#&#8203;62484](https://redirect.github.com/nodejs/node/pull/62484)
- \[[`90fab71a84`](https://redirect.github.com/nodejs/node/commit/90fab71a84)] - **deps**: update simdjson to 4.5.0 (Node.js GitHub Bot) [#&#8203;62382](https://redirect.github.com/nodejs/node/pull/62382)
- \[[`a416ddf6d9`](https://redirect.github.com/nodejs/node/commit/a416ddf6d9)] - **deps**: V8: cherry-pick [`cf1bce4`](https://redirect.github.com/nodejs/node/commit/cf1bce40a5ef) (Richard Lau) [#&#8203;62449](https://redirect.github.com/nodejs/node/pull/62449)
- \[[`4d9123e57d`](https://redirect.github.com/nodejs/node/commit/4d9123e57d)] - **deps**: upgrade npm to 11.12.1 (npm team) [#&#8203;62448](https://redirect.github.com/nodejs/node/pull/62448)
- \[[`952d715028`](https://redirect.github.com/nodejs/node/commit/952d715028)] - **deps**: update sqlite to 3.51.3 (Node.js GitHub Bot) [#&#8203;62256](https://redirect.github.com/nodejs/node/pull/62256)
- \[[`f3fd7ed426`](https://redirect.github.com/nodejs/node/commit/f3fd7ed426)] - **deps**: update googletest to [`73a63ea`](https://redirect.github.com/nodejs/node/commit/73a63ea05dc8ca29ec1d2c1d66481dd0de1950f1) (Node.js GitHub Bot) [#&#8203;61927](https://redirect.github.com/nodejs/node/pull/61927)
- \[[`71a2f82d7c`](https://redirect.github.com/nodejs/node/commit/71a2f82d7c)] - **deps**: upgrade npm to 11.11.1 (npm team) [#&#8203;62216](https://redirect.github.com/nodejs/node/pull/62216)
- \[[`84f60c26f7`](https://redirect.github.com/nodejs/node/commit/84f60c26f7)] - **deps**: update amaro to 1.1.8 (Node.js GitHub Bot) [#&#8203;62151](https://redirect.github.com/nodejs/node/pull/62151)
- \[[`43159d0e5f`](https://redirect.github.com/nodejs/node/commit/43159d0e5f)] - **deps**: update sqlite to 3.52.0 (Node.js GitHub Bot) [#&#8203;62150](https://redirect.github.com/nodejs/node/pull/62150)
- \[[`b887657b38`](https://redirect.github.com/nodejs/node/commit/b887657b38)] - **deps**: V8: cherry-pick [`aa0b288`](https://redirect.github.com/nodejs/node/commit/aa0b288f87cc) (Richard Lau) [#&#8203;62136](https://redirect.github.com/nodejs/node/pull/62136)
- \[[`7ab885b323`](https://redirect.github.com/nodejs/node/commit/7ab885b323)] - **deps**: update ada to 3.4.3 (Node.js GitHub Bot) [#&#8203;62049](https://redirect.github.com/nodejs/node/pull/62049)
- \[[`671ddec2b9`](https://redirect.github.com/nodejs/node/commit/671ddec2b9)] - **deps**: update minimatch to 10.2.4 (Node.js GitHub Bot) [#&#8203;62016](https://redirect.github.com/nodejs/node/pull/62016)
- \[[`290fe37d4d`](https://redirect.github.com/nodejs/node/commit/290fe37d4d)] - **deps**: update simdjson to 4.3.1 (Node.js GitHub Bot) [#&#8203;61930](https://redirect.github.com/nodejs/node/pull/61930)
- \[[`a13bee76b5`](https://redirect.github.com/nodejs/node/commit/a13bee76b5)] - **deps**: update acorn-walk to 8.3.5 (Node.js GitHub Bot) [#&#8203;61928](https://redirect.github.com/nodejs/node/pull/61928)
- \[[`f0e40b35b9`](https://redirect.github.com/nodejs/node/commit/f0e40b35b9)] - **deps**: update acorn to 8.16.0 (Node.js GitHub Bot) [#&#8203;61925](https://redirect.github.com/nodejs/node/pull/61925)
- \[[`463dfa023a`](https://redirect.github.com/nodejs/node/commit/463dfa023a)] - **deps**: update minimatch to 10.2.2 (Node.js GitHub Bot) [#&#8203;61830](https://redirect.github.com/nodejs/node/pull/61830)
- \[[`4b2e4bb108`](https://redirect.github.com/nodejs/node/commit/4b2e4bb108)] - **deps**: update nbytes to 0.1.3 (Node.js GitHub Bot) [#&#8203;61879](https://redirect.github.com/nodejs/node/pull/61879)
- \[[`5626cb83d0`](https://redirect.github.com/nodejs/node/commit/5626cb83d0)] - **deps**: remove stale OpenSSL arch configs (René) [#&#8203;61834](https://redirect.github.com/nodejs/node/pull/61834)
- \[[`52668874fd`](https://redirect.github.com/nodejs/node/commit/52668874fd)] - **deps**: update llhttp to 9.3.1 (Node.js GitHub Bot) [#&#8203;61827](https://redirect.github.com/nodejs/node/pull/61827)
- \[[`b3387b07b1`](https://redirect.github.com/nodejs/node/commit/b3387b07b1)] - **deps**: update googletest to [`5a9c3f9`](https://redirect.github.com/nodejs/node/commit/5a9c3f9e8d9b90bbbe8feb32902146cb8f7c1757) (Node.js GitHub Bot) [#&#8203;61731](https://redirect.github.com/nodejs/node/pull/61731)
- \[[`196268cb4c`](https://redirect.github.com/nodejs/node/commit/196268cb4c)] - **deps**: V8: cherry-pick [`c5ff7c4`](https://redirect.github.com/nodejs/node/commit/c5ff7c4d6cde) (Chengzhong Wu) [#&#8203;61372](https://redirect.github.com/nodejs/node/pull/61372)
- \[[`36869b52de`](https://redirect.github.com/nodejs/node/commit/36869b52de)] - **deps**: update merve to 1.2.2 (Node.js GitHub Bot) [#&#8203;62213](https://redirect.github.com/nodejs/node/pull/62213)
- \[[`3cbac055de`](https://redirect.github.com/nodejs/node/commit/3cbac055de)] - **deps**: update merve to 1.2.0 (Node.js GitHub Bot) [#&#8203;62149](https://redirect.github.com/nodejs/node/pull/62149)
- \[[`7757cc3495`](https://redirect.github.com/nodejs/node/commit/7757cc3495)] - **deps**: V8: backport [`6a0a25a`](https://redirect.github.com/nodejs/node/commit/6a0a25abaed3) (Vivian Wang) [#&#8203;61670](https://redirect.github.com/nodejs/node/pull/61670)
- \[[`359797c2fb`](https://redirect.github.com/nodejs/node/commit/359797c2fb)] - **deps,src**: prepare for cpplint update (Michaël Zasso) [#&#8203;60901](https://redirect.github.com/nodejs/node/pull/60901)
- \[[`ace802e59b`](https://redirect.github.com/nodejs/node/commit/ace802e59b)] - **diagnostics\_channel**: add diagnostics channels for web locks (Ilyas Shabi) [#&#8203;62123](https://redirect.github.com/nodejs/node/pull/62123)
- \[[`a072411b03`](https://redirect.github.com/nodejs/node/commit/a072411b03)] - **doc**: remove spawn with shell example from bat/cmd section (Kit Dallege) [#&#8203;62243](https://redirect.github.com/nodejs/node/pull/62243)
- \[[`0b152449af`](https://redirect.github.com/nodejs/node/commit/0b152449af)] - **doc**: fix typo in --disable-wasm-trap-handler description (Dmytro Semchuk) [#&#8203;61820](https://redirect.github.com/nodejs/node/pull/61820)
- \[[`73ea387ad7`](https://redirect.github.com/nodejs/node/commit/73ea387ad7)] - **doc**: remove obsolete Boxstarter automated install (Mike McCready) [#&#8203;61785](https://redirect.github.com/nodejs/node/pull/61785)
- \[[`7f234add8e`](https://redirect.github.com/nodejs/node/commit/7f234add8e)] - **doc**: deprecate `module.register()` (DEP0205) (Geoffrey Booth) [#&#8203;62395](https://redirect.github.com/nodejs/node/pull/62395)
- \[[`12fc3c6a30`](https://redirect.github.com/nodejs/node/commit/12fc3c6a30)] - **doc**: clarify that features cannot be both experimental and deprecated (Antoine du Hamel) [#&#8203;62456](https://redirect.github.com/nodejs/node/pull/62456)
- \[[`1ecc5962a2`](https://redirect.github.com/nodejs/node/commit/1ecc5962a2)] - **doc**: fix 'transfered' typo in quic.md (lilianakatrina684-a11y) [#&#8203;62492](https://redirect.github.com/nodejs/node/pull/62492)
- \[[`56741a1303`](https://redirect.github.com/nodejs/node/commit/56741a1303)] - **doc**: move sqlite type conversion section to correct level (René) [#&#8203;62482](https://redirect.github.com/nodejs/node/pull/62482)
- \[[`12b04d17d5`](https://redirect.github.com/nodejs/node/commit/12b04d17d5)] - **doc**: add Rafael to last security release steward (Rafael Gonzaga) [#&#8203;62423](https://redirect.github.com/nodejs/node/pull/62423)
- \[[`c4567e4a8d`](https://redirect.github.com/nodejs/node/commit/c4567e4a8d)] - **doc**: fix overstated Date header requirement in response.sendDate (Kit Dallege) [#&#8203;62206](https://redirect.github.com/nodejs/node/pull/62206)
- \[[`384a41047f`](https://redirect.github.com/nodejs/node/commit/384a41047f)] - **doc**: enhance clarification about the main field (Mowafak Almahaini) [#&#8203;62302](https://redirect.github.com/nodejs/node/pull/62302)
- \[[`93d19b1a1c`](https://redirect.github.com/nodejs/node/commit/93d19b1a1c)] - **doc**: minor typo fix (Jeff Matson) [#&#8203;62358](https://redirect.github.com/nodejs/node/pull/62358)
- \[[`3db35d2c59`](https://redirect.github.com/nodejs/node/commit/3db35d2c59)] - **doc**: add path to vulnerabilities.json mention (Rafael Gonzaga) [#&#8203;62355](https://redirect.github.com/nodejs/node/pull/62355)
- \[[`57b105c9d5`](https://redirect.github.com/nodejs/node/commit/57b105c9d5)] - **doc**: deprecate CryptoKey use in node:crypto (Filip Skokan) [#&#8203;62321](https://redirect.github.com/nodejs/node/pull/62321)
- \[[`490168c993`](https://redirect.github.com/nodejs/node/commit/490168c993)] - **doc**: fix small environment\_variables typo (chris) [#&#8203;62279](https://redirect.github.com/nodejs/node/pull/62279)
- \[[`0291be584b`](https://redirect.github.com/nodejs/node/commit/0291be584b)] - **doc**: test and test-only targets do not run linter (Xavier Stouder) [#&#8203;62120](https://redirect.github.com/nodejs/node/pull/62120)
- \[[`ba0a82a1e1`](https://redirect.github.com/nodejs/node/commit/ba0a82a1e1)] - **doc**: clarify fs.ReadStream and fs.WriteStream are not constructable (Kit Dallege) [#&#8203;62208](https://redirect.github.com/nodejs/node/pull/62208)
- \[[`125bdbf504`](https://redirect.github.com/nodejs/node/commit/125bdbf504)] - **doc**: clarify that any truthy value of `shell` is part of DEP0190 (Antoine du Hamel) [#&#8203;62249](https://redirect.github.com/nodejs/node/pull/62249)
- \[[`a141ad0aeb`](https://redirect.github.com/nodejs/node/commit/a141ad0aeb)] - **doc**: remove outdated Chrome 66 and ndb references from debugger (Kit Dallege) [#&#8203;62202](https://redirect.github.com/nodejs/node/pull/62202)
- \[[`44bde8e573`](https://redirect.github.com/nodejs/node/commit/44bde8e573)] - **doc**: add note (and caveat) for `mock.module` about customization hooks (Jacob Smith) [#&#8203;62075](https://redirect.github.com/nodejs/node/pull/62075)
- \[[`8c46a1ca1a`](https://redirect.github.com/nodejs/node/commit/8c46a1ca1a)] - **doc**: copyedit `addons.md` (Antoine du Hamel) [#&#8203;62071](https://redirect.github.com/nodejs/node/pull/62071)
- \[[`7f989f02f7`](https://redirect.github.com/nodejs/node/commit/7f989f02f7)] - **doc**: correct `util.convertProcessSignalToExitCode` validation behavior (René) [#&#8203;62134](https://redirect.github.com/nodejs/node/pull/62134)
- \[[`a4466ebdac`](https://redirect.github.com/nodejs/node/commit/a4466ebdac)] - **doc**: add efekrskl as triager (Efe) [#&#8203;61876](https://redirect.github.com/nodejs/node/pull/61876)
- \[[`db516eca3a`](https://redirect.github.com/nodejs/node/commit/db516eca3a)] - **doc**: fix markdown for `expectFailure` values (Jacob Smith) [#&#8203;62100](https://redirect.github.com/nodejs/node/pull/62100)
- \[[`ad97045125`](https://redirect.github.com/nodejs/node/commit/ad97045125)] - **doc**: include url.resolve() in DEP0169 application deprecation (Mike McCready) [#&#8203;62002](https://redirect.github.com/nodejs/node/pull/62002)
- \[[`309f37ba42`](https://redirect.github.com/nodejs/node/commit/309f37ba42)] - **doc**: expand SECURITY.md with non-vulnerability examples (Rafael Gonzaga) [#&#8203;61972](https://redirect.github.com/nodejs/node/pull/61972)
- \[[`dbb3551b7b`](https://redirect.github.com/nodejs/node/commit/dbb3551b7b)] - **doc**: separate in-types and out-types in SQLite conversion docs (René) [#&#8203;62034](https://redirect.github.com/nodejs/node/pull/62034)
- \[[`191c433db8`](https://redirect.github.com/nodejs/node/commit/191c433db8)] - **doc**: fix small logic error in DETECT\_MODULE\_SYNTAX (René) [#&#8203;62025](https://redirect.github.com/nodejs/node/pull/62025)
- \[[`8511b1c784`](https://redirect.github.com/nodejs/node/commit/8511b1c784)] - **doc**: fix module.stripTypeScriptTypes indentation (René) [#&#8203;61992](https://redirect.github.com/nodejs/node/pull/61992)
- \[[`dd1139f52c`](https://redirect.github.com/nodejs/node/commit/dd1139f52c)] - **doc**: update DEP0040 (punycode) to application type deprecation (Mike McCready) [#&#8203;61916](https://redirect.github.com/nodejs/node/pull/61916)
- \[[`54009e9c62`](https://redirect.github.com/nodejs/node/commit/54009e9c62)] - **doc**: explicitly mention Slack handle (Rafael Gonzaga) [#&#8203;61986](https://redirect.github.com/nodejs/node/pull/61986)
- \[[`78fa1a1a49`](https://redirect.github.com/nodejs/node/commit/78fa1a1a49)] - **doc**: support toolchain Visual Studio 2022 & 2026 + Windows 11 SDK (Mike McCready) [#&#8203;61864](https://redirect.github.com/nodejs/node/pull/61864)
- \[[`d8204d3cdb`](https://redirect.github.com/nodejs/node/commit/d8204d3cdb)] - **doc**: rename invalid `function` parameter (René) [#&#8203;61942](https://redirect.github.com/nodejs/node/pull/61942)
- \[[`a5a14482fb`](https://redirect.github.com/nodejs/node/commit/a5a14482fb)] - **doc**: clarify status of feature request issues (Antoine du Hamel) [#&#8203;61505](https://redirect.github.com/nodejs/node/pull/61505)
- \[[`bd0688feb6`](https://redirect.github.com/nodejs/node/commit/bd0688feb6)] - **doc**: add esm and cjs examples to node:vm (Alfredo González) [#&#8203;61498](https://redirect.github.com/nodejs/node/pull/61498)
- \[[`240b512f9f`](https://redirect.github.com/nodejs/node/commit/240b512f9f)] - **doc**: clarify build environment is trusted in threat model (Matteo Collina) [#&#8203;61865](https://redirect.github.com/nodejs/node/pull/61865)
- \[[`5dd48e3456`](https://redirect.github.com/nodejs/node/commit/5dd48e3456)] - **doc**: remove incorrect mention of `module` in `typescript.md` (Rob Palmer) [#&#8203;61839](https://redirect.github.com/nodejs/node/pull/61839)
- \[[`9502c22055`](https://redirect.github.com/nodejs/node/commit/9502c22055)] - **doc**: simplify addAbortListener example (Chemi Atlow) [#&#8203;61842](https://redirect.github.com/nodejs/node/pull/61842)
- \[[`6fec397828`](https://redirect.github.com/nodejs/node/commit/6fec397828)] - **doc**: clean up globals.md (René) [#&#8203;61822](https://redirect.github.com/nodejs/node/pull/61822)
- \[[`a810f5ccef`](https://redirect.github.com/nodejs/node/commit/a810f5ccef)] - **doc**: clarify async caveats for `events.once()` (René) [#&#8203;61572](https://redirect.github.com/nodejs/node/pull/61572)
- \[[`2bf990bb1a`](https://redirect.github.com/nodejs/node/commit/2bf990bb1a)] - **doc**: update Juan's security steward info (Juan José) [#&#8203;61754](https://redirect.github.com/nodejs/node/pull/61754)
- \[[`0312db948d`](https://redirect.github.com/nodejs/node/commit/0312db948d)] - **doc**: fix methods being documented as properties in `process.md` (Antoine du Hamel) [#&#8203;61765](https://redirect.github.com/nodejs/node/pull/61765)
- \[[`e558b26e7f`](https://redirect.github.com/nodejs/node/commit/e558b26e7f)] - **doc**: add riscv64 info into platform list (Lu Yahan) [#&#8203;42251](https://redirect.github.com/nodejs/node/pull/42251)
- \[[`49254e3dc0`](https://redirect.github.com/nodejs/node/commit/49254e3dc0)] - **doc**: fix dropdown menu being obscured at <600px due to stacking context (Jeff) [#&#8203;61735](https://redirect.github.com/nodejs/node/pull/61735)
- \[[`4ff01b5c10`](https://redirect.github.com/nodejs/node/commit/4ff01b5c10)] - **doc**: fix spacing in process message event (Aviv Keller) [#&#8203;61756](https://redirect.github.com/nodejs/node/pull/61756)
- \[[`94097a79d6`](https://redirect.github.com/nodejs/node/commit/94097a79d6)] - **doc**: move describe/it aliases section before expectFailure (Luca Raveri) [#&#8203;61567](https://redirect.github.com/nodejs/node/pull/61567)
- \[[`b7cd31acbe`](https://redirect.github.com/nodejs/node/commit/b7cd31acbe)] - **doc**: fix broken links of net.md (YuSheng Chen) [#&#8203;61673](https://redirect.github.com/nodejs/node/pull/61673)
- \[[`ae5e353fe2`](https://redirect.github.com/nodejs/node/commit/ae5e353fe2)] - **doc**: clean up Windows code snippet in `child_process.md` (reillylm) [#&#8203;61422](https://redirect.github.com/nodejs/node/pull/61422)
- \[[`ea9beb6a3c`](https://redirect.github.com/nodejs/node/commit/ea9beb6a3c)] - **doc**: update to Visual Studio 2026 manual install (Mike McCready) [#&#8203;61655](https://redirect.github.com/nodejs/node/pull/61655)
- \[[`42057c84e2`](https://redirect.github.com/nodejs/node/commit/42057c84e2)] - **doc,module**: add missing doc for syncHooks.deregister() (Joyee Cheung) [#&#8203;61959](https://redirect.github.com/nodejs/node/pull/61959)
- \[[`a035bd5235`](https://redirect.github.com/nodejs/node/commit/a035bd5235)] - **doc,test**: clarify --eval syntax for leading '-' scripts (kovan) [#&#8203;62244](https://redirect.github.com/nodejs/node/pull/62244)
- \[[`deb0b78460`](https://redirect.github.com/nodejs/node/commit/deb0b78460)] - **esm**: fix typo in worker loader hook comment (jakecastelli) [#&#8203;62475](https://redirect.github.com/nodejs/node/pull/62475)
- \[[`b93bf7dbfc`](https://redirect.github.com/nodejs/node/commit/b93bf7dbfc)] - **esm**: fix source phase identity bug in loadCache eviction (Guy Bedford) [#&#8203;62415](https://redirect.github.com/nodejs/node/pull/62415)
- \[[`679d18b57f`](https://redirect.github.com/nodejs/node/commit/679d18b57f)] - **esm**: fix path normalization in `finalizeResolution` (Antoine du Hamel) [#&#8203;62080](https://redirect.github.com/nodejs/node/pull/62080)
- \[[`171e9fc268`](https://redirect.github.com/nodejs/node/commit/171e9fc268)] - **esm**: update outdated FIXME comment in translators.js (Karan Mangtani) [#&#8203;61715](https://redirect.github.com/nodejs/node/pull/61715)
- \[[`cc19728228`](https://redirect.github.com/nodejs/node/commit/cc19728228)] - **events**: avoid cloning listeners array on every emit (Gürgün Dayıoğlu) [#&#8203;62261](https://redirect.github.com/nodejs/node/pull/62261)
- \[[`458c92be52`](https://redirect.github.com/nodejs/node/commit/458c92be52)] - **events**: don't call resume after close (Сковорода Никита Андреевич) [#&#8203;60548](https://redirect.github.com/nodejs/node/pull/60548)
- \[[`4691f3e7fb`](https://redirect.github.com/nodejs/node/commit/4691f3e7fb)] - **fs**: fix cpSync to handle non-ASCII characters (Stefan Stojanovic) [#&#8203;61950](https://redirect.github.com/nodejs/node/pull/61950)
- \[[`f4a3edc47a`](https://redirect.github.com/nodejs/node/commit/f4a3edc47a)] - **(SEMVER-MINOR)** **fs**: add `throwIfNoEntry` option for fs.stat and fs.promises.stat (Juan José) [#&#8203;61178](https://redirect.github.com/nodejs/node/pull/61178)
- \[[`58e4d50cd0`](https://redirect.github.com/nodejs/node/commit/58e4d50cd0)] - **http**: fix use-after-free when freeParser is called during llhttp\_execute (Gerhard Stöbich) [#&#8203;62095](https://redirect.github.com/nodejs/node/pull/62095)
- \[[`0a4ad85ab0`](https://redirect.github.com/nodejs/node/commit/0a4ad85ab0)] - **http**: validate ClientRequest path on set (Matteo Collina) [#&#8203;62030](https://redirect.github.com/nodejs/node/pull/62030)
- \[[`f8178ac3e6`](https://redirect.github.com/nodejs/node/commit/f8178ac3e6)] - **http**: validate headers in writeEarlyHints (Richard Clarke) [#&#8203;61897](https://redirect.github.com/nodejs/node/pull/61897)
- \[[`899884d0ed`](https://redirect.github.com/nodejs/node/commit/899884d0ed)] - **http**: remove redundant keepAliveTimeoutBuffer assignment (Efe) [#&#8203;61743](https://redirect.github.com/nodejs/node/pull/61743)
- \[[`08d2e40694`](https://redirect.github.com/nodejs/node/commit/08d2e40694)] - **http**: attach error handler to socket synchronously in onSocket (RajeshKumar11) [#&#8203;61770](https://redirect.github.com/nodejs/node/pull/61770)
- \[[`1c2064c1f8`](https://redirect.github.com/nodejs/node/commit/1c2064c1f8)] - **http**: fix keep-alive socket reuse race in requestOnFinish (Martin Slota) [#&#8203;61710](https://redirect.github.com/nodejs/node/pull/61710)
- \[[`38e9c66e0f`](https://redirect.github.com/nodejs/node/commit/38e9c66e0f)] - **http2**: add strictSingleValueFields option to relax header validation (Tim Perry) [#&#8203;59917](https://redirect.github.com/nodejs/node/pull/59917)
- \[[`5cdcba17cc`](https://redirect.github.com/nodejs/node/commit/5cdcba17cc)] - **(SEMVER-MINOR)** **http2**: add http1Options for HTTP/1 fallback configuration (Amol Yadav) [#&#8203;61713](https://redirect.github.com/nodejs/node/pull/61713)
- \[[`687c0acd00`](https://redirect.github.com/nodejs/node/commit/687c0acd00)] - **http2**: fix FileHandle leak in respondWithFile (sangwook) [#&#8203;61707](https://redirect.github.com/nodejs/node/pull/61707)
- \[[`0c8f802ec2`](https://redirect.github.com/nodejs/node/commit/0c8f802ec2)] - **inspector**: add Target.getTargets and extract TargetManager (Kohei) [#&#8203;62487](https://redirect.github.com/nodejs/node/pull/62487)
- \[[`7de8a303c1`](https://redirect.github.com/nodejs/node/commit/7de8a303c1)] - **inspector**: unwrap internal/debugger/inspect imports (René) [#&#8203;61974](https://redirect.github.com/nodejs/node/pull/61974)
- \[[`59ac10a4fd`](https://redirect.github.com/nodejs/node/commit/59ac10a4fd)] - **lib**: make SubtleCrypto.supports enumerable (Filip Skokan) [#&#8203;62307](https://redirect.github.com/nodejs/node/pull/62307)
- \[[`9dc102ba90`](https://redirect.github.com/nodejs/node/commit/9dc102ba90)] - **lib**: prefer primordials in SubtleCrypto (Filip Skokan) [#&#8203;62226](https://redirect.github.com/nodejs/node/pull/62226)
- \[[`78a9aa8f32`](https://redirect.github.com/nodejs/node/commit/78a9aa8f32)] - **lib**: fix source map url parse in dynamic imports (Chengzhong Wu) [#&#8203;61990](https://redirect.github.com/nodejs/node/pull/61990)
- \[[`16b8cc6643`](https://redirect.github.com/nodejs/node/commit/16b8cc6643)] - **lib**: improve argument handling in Blob constructor (Ms2ger) [#&#8203;61980](https://redirect.github.com/nodejs/node/pull/61980)
- \[[`a03b5d39b8`](https://redirect.github.com/nodejs/node/commit/a03b5d39b8)] - **lib**: reduce cycles in esm loader and load it in snapshot (Joyee Cheung) [#&#8203;61769](https://redirect.github.com/nodejs/node/pull/61769)
- \[[`1017bf5f86`](https://redirect.github.com/nodejs/node/commit/1017bf5f86)] - **lib**: remove top-level getOptionValue() calls in lib/internal/modules (Joyee Cheung) [#&#8203;61769](https://redirect.github.com/nodejs/node/pull/61769)
- \[[`d79984b41b`](https://redirect.github.com/nodejs/node/commit/d79984b41b)] - **lib**: optimize styleText when validateStream is false (Rafael Gonzaga) [#&#8203;61792](https://redirect.github.com/nodejs/node/pull/61792)
- \[[`6462b89d10`](https://redirect.github.com/nodejs/node/commit/6462b89d10)] - **meta**: bump actions/download-artifact from 7.0.0 to 8.0.0 (dependabot\[bot]) [#&#8203;62063](https://redirect.github.com/nodejs/node/pull/62063)
- \[[`5bb89916ea`](https://redirect.github.com/nodejs/node/commit/5bb89916ea)] - **meta**: bump actions/upload-artifact from 6.0.0 to 7.0.0 (dependabot\[bot]) [#&#8203;62062](https://redirect.github.com/nodejs/node/pull/62062)
- \[[`b067d74d94`](https://redirect.github.com/nodejs/node/commit/b067d74d94)] - **meta**: bump step-security/harden-runner from 2.14.2 to 2.15.0 (dependabot\[bot]) [#&#8203;62064](https://redirect.github.com/nodejs/node/pull/62064)
- \[[`830e5cd125`](https://redirect.github.com/nodejs/node/commit/830e5cd125)] - **meta**: bump github/codeql-action from 4.32.0 to 4.32.4 (dependabot\[bot]) [#&#8203;61911](https://redirect.github.com/nodejs/node/pull/61911)
- \[[`16c839a3dd`](https://redirect.github.com/nodejs/node/commit/16c839a3dd)] - **meta**: bump step-security/harden-runner from 2.14.1 to 2.14.2 (dependabot\[bot]) [#&#8203;61909](https://redirect.github.com/nodejs/node/pull/61909)
- \[[`498abf661e`](https://redirect.github.com/nodejs/node/commit/498abf661e)] - **meta**: bump actions/stale from 10.1.1 to 10.2.0 (dependabot\[bot]) [#&#8203;61908](https://redirect.github.com/nodejs/node/pull/61908)
- \[[`78ac17f426`](https://redirect.github.com/nodejs/node/commit/78ac17f426)] - **module**: fix coverage of mocked CJS modules imported from ESM (Marco) [#&#8203;62133](https://redirect.github.com/nodejs/node/pull/62133)
- \[[`46cfad4138`](https://redirect.github.com/nodejs/node/commit/46cfad4138)] - **module**: run require.resolve through module.registerHooks() (Joyee Cheung) [#&#8203;62028](https://redirect.github.com/nodejs/node/pull/62028)
- \[[`8b6be3fe14`](https://redirect.github.com/nodejs/node/commit/8b6be3fe14)] - **module**: mark require(esm) as stable (Joyee Cheung) [#&#8203;60959](https://redirect.github.com/nodejs/node/pull/60959)
- \[[`68fbc0c6cc`](https://redirect.github.com/nodejs/node/commit/68fbc0c6cc)] - **module**: mark module compile cache as stable (Joyee Cheung) [#&#8203;60971](https://redirect.github.com/nodejs/node/pull/60971)
- \[[`c851e76f8c`](https://redirect.github.com/nodejs/node/commit/c851e76f8c)] - **(SEMVER-MINOR)** **net**: add `setTOS` and `getTOS` to `Socket` (Amol Yadav) [#&#8203;61503](https://redirect.github.com/nodejs/node/pull/61503)
- \[[`4c206ecb31`](https://redirect.github.com/nodejs/node/commit/4c206ecb31)] - **quic**: remove CryptoKey support from session keys option (Filip Skokan) [#&#8203;62335](https://redirect.github.com/nodejs/node/pull/62335)
- \[[`2f9c085cf5`](https://redirect.github.com/nodejs/node/commit/2f9c085cf5)] - **sqlite**: handle stmt invalidation (Guilherme Araújo) [#&#8203;61877](https://redirect.github.com/nodejs/node/pull/61877)
- \[[`6ac4304c87`](https://redirect.github.com/nodejs/node/commit/6ac4304c87)] - **(SEMVER-MINOR)** **sqlite**: add limits property to DatabaseSync (Mert Can Altin) [#&#8203;61298](https://redirect.github.com/nodejs/node/pull/61298)
- \[[`aaf9af1672`](https://redirect.github.com/nodejs/node/commit/aaf9af1672)] - **sqlite**: mark as release candidate (Matteo Collina) [#&#8203;61262](https://redirect.github.com/nodejs/node/pull/61262)
- \[[`7d67e5d693`](https://redirect.github.com/nodejs/node/commit/7d67e5d693)] - **src**: convert context\_frame field in AsyncWrap to internal field (Anna Henningsen) [#&#8203;62103](https://redirect.github.com/nodejs/node/pull/62103)
- \[[`d8ea1aaa8a`](https://redirect.github.com/nodejs/node/commit/d8ea1aaa8a)] - **src**: make AsyncWrap subclass internal field counts explicit (Anna Henningsen) [#&#8203;62103](https://redirect.github.com/nodejs/node/pull/62103)
- \[[`1dbf3bedbe`](https://redirect.github.com/nodejs/node/commit/1dbf3bedbe)] - **src**: improve EC JWK import performance (Filip Skokan) [#&#8203;62396](https://redirect.github.com/nodejs/node/pull/62396)
- \[[`cd84af747b`](https://redirect.github.com/nodejs/node/commit/cd84af747b)] - **src**: handle null backing store in ArrayBufferViewContents::Read (Mert Can Altin) [#&#8203;62343](https://redirect.github.com/nodejs/node/pull/62343)
- \[[`4f553cdc01`](https://redirect.github.com/nodejs/node/commit/4f553cdc01)] - **src**: enable compilation/linking with OpenSSL 4.0 (Filip Skokan) [#&#8203;62410](https://redirect.github.com/nodejs/node/pull/62410)
- \[[`70f8057258`](https://redirect.github.com/nodejs/node/commit/70f8057258)] - **src**: use stack allocation in indexOf latin1 path (Mert Can Altin) [#&#8203;62268](https://redirect.github.com/nodejs/node/pull/62268)
- \[[`d788467b6a`](https://redirect.github.com/nodejs/node/commit/d788467b6a)] - **src**: expose async context frame debugging helper to JS (Anna Henningsen) [#&#8203;62103](https://redirect.github.com/nodejs/node/pull/62103)
- \[[`4213f893ec`](https://redirect.github.com/nodejs/node/commit/4213f893ec)] - **src**: release context frame in AsyncWrap::EmitDestroy (Gerhard Stöbich) [#&#8203;61995](https://redirect.github.com/nodejs/node/pull/61995)
- \[[`79fb8cbcf5`](https://redirect.github.com/nodejs/node/commit/79fb8cbcf5)] - **src**: use validate\_ascii\_with\_errors instead of validate\_ascii (Сковорода Никита Андреевич) [#&#8203;61122](https://redirect.github.com/nodejs/node/pull/61122)
- \[[`2df328d59e`](https://redirect.github.com/nodejs/node/commit/2df328d59e)] - **src**: fix flags argument offset in JSUdpWrap (Weixie Cui) [#&#8203;61948](https://redirect.github.com/nodejs/node/pull/61948)
- \[[`eb77a7a297`](https://redirect.github.com/nodejs/node/commit/eb77a7a297)] - **(SEMVER-MINOR)** **src**: add C++ support for diagnostics channels (RafaelGSS) [#&#8203;61869](https://redirect.github.com/nodejs/node/pull/61869)
- \[[`6cda3d30c0`](https://redirect.github.com/nodejs/node/commit/6cda3d30c0)] - **src**: remove unnecessary `c_str()` conversions in diagnostic messages (Anna Henningsen) [#&#8203;61786](https://redirect.github.com/nodejs/node/pull/61786)
- \[[`26c6045363`](https://redirect.github.com/nodejs/node/commit/26c6045363)] - **src**: use bool literals in TraceEnvVarOptions (Tobias Nießen) [#&#8203;61425](https://redirect.github.com/nodejs/node/pull/61425)
- \[[`3c8f700fd7`](https://redirect.github.com/nodejs/node/commit/3c8f700fd7)] - **src**: track allocations made by zstd streams (Anna Henningsen) [#&#8203;61717](https://redirect.github.com/nodejs/node/pull/61717)
- \[[`94dbb36d4d`](https://redirect.github.com/nodejs/node/commit/94dbb36d4d)] - **src**: do not store compression methods on Brotli classes (Anna Henningsen) [#&#8203;61717](https://redirect.github.com/nodejs/node/pull/61717)
- \[[`bef661f182`](https://redirect.github.com/nodejs/node/commit/bef661f182)] - **src**: extract zlib allocation tracking into its own class (Anna Henningsen) [#&#8203;61717](https://redirect.github.com/nodejs/node/pull/61717)
- \[[`e8079a8297`](https://redirect.github.com/nodejs/node/commit/e8079a8297)] - **src**: release memory for zstd contexts in `Close()` (Anna Henningsen) [#&#8203;61717](https://redirect.github.com/nodejs/node/pull/61717)
- \[[`6e1197a3cc`](https://redirect.github.com/nodejs/node/commit/6e1197a3cc)] - **src**: add more checks and clarify docs for external references (Joyee Cheung) [#&#8203;61719](https://redirect.github.com/nodejs/node/pull/61719)
- \[[`c28a22c4be`](https://redirect.github.com/nodejs/node/commit/c28a22c4be)] - **src**: fix cjs\_lexer external reference registration (Joyee Cheung) [#&#8203;61718](https://redirect.github.com/nodejs/node/pull/61718)
- \[[`9e2c5fd7c9`](https://redirect.github.com/nodejs/node/commit/9e2c5fd7c9)] - **src**: simply uint32 to string as it must not fail (Chengzhong Wu) [#&#8203;60846](https://redirect.github.com/nodejs/node/pull/60846)
- \[[`df435d32b8`](https://redirect.github.com/nodejs/node/commit/df435d32b8)] - **src**: build v8 tick processor as built-in source text modules (Joyee Cheung) [#&#8203;60518](https://redirect.github.com/nodejs/node/pull/60518)
- \[[`2cb3573735`](https://redirect.github.com/nodejs/node/commit/2cb3573735)] - **src,sqlite**: fix filterFunc dangling reference (Edy Silva) [#&#8203;62281](https://redirect.github.com/nodejs/node/pull/62281)
- \[[`c44f53b544`](https://redirect.github.com/nodejs/node/commit/c44f53b544)] - **stream**: preserve error over AbortError in pipeline (Marco) [#&#8203;62113](https://redirect.github.com/nodejs/node/pull/62113)
- \[[`dc541370b4`](https://redirect.github.com/nodejs/node/commit/dc541370b4)] - **stream**: replace bind with arrow function for onwrite callback (Ali Hassan) [#&#8203;62087](https://redirect.github.com/nodejs/node/pull/62087)
- \[[`f6cdfbfaa7`](https://redirect.github.com/nodejs/node/commit/f6cdfbfaa7)] - **stream**: optimize webstreams pipeTo (Mattias Buelens) [#&#8203;62079](https://redirect.github.com/nodejs/node/pull/62079)
- \[[`fcf2a9f788`](https://redirect.github.com/nodejs/node/commit/fcf2a9f788)] - **stream**: fix brotli error handling in web compression streams (Filip Skokan) [#&#8203;62107](https://redirect.github.com/nodejs/node/pull/62107)
- \[[`cdec579c6b`](https://redirect.github.com/nodejs/node/commit/cdec579c6b)] - **stream**: improve Web Compression spec compliance (Filip Skokan) [#&#8203;62107](https://redirect.github.com/nodejs/node/pull/62107)
- \[[`dbe5898379`](https://redirect.github.com/nodejs/node/commit/dbe5898379)] - **stream**: fix UTF-8 character corruption in fast-utf8-stream (Matteo Collina) [#&#8203;61745](https://redirect.github.com/nodejs/node/pull/61745)
- \[[`531e62cd74`](https://redirect.github.com/nodejs/node/commit/531e62cd74)] - **stream**: fix TransformStream race on cancel with pending write (Marco) [#&#8203;62040](https://redirect.github.com/nodejs/node/pull/62040)
- \[[`a3751f2249`](https://redirect.github.com/nodejs/node/commit/a3751f2249)] - **stream**: accept ArrayBuffer in CompressionStream and DecompressionStream (조수민) [#&#8203;61913](https://redirect.github.com/nodejs/node/pull/61913)
- \[[`65aa8f68d0`](https://redirect.github.com/nodejs/node/commit/65aa8f68d0)] - **stream**: fix pipeTo to defer writes per WHATWG spec (Matteo Collina) [#&#8203;61800](https://redirect.github.com/nodejs/node/pull/61800)
- \[[`15f32b4935`](https://redirect.github.com/nodejs/node/commit/15f32b4935)] - **stream**: fix decoded fromList chunk boundary check (Thomas Watson) [#&#8203;61884](https://redirect.github.com/nodejs/node/pull/61884)
- \[[`569767e52e`](https://redirect.github.com/nodejs/node/commit/569767e52e)] - **stream**: add fast paths for webstreams read and pipeTo (Matteo Collina) [#&#8203;61807](https://redirect.github.com/nodejs/node/pull/61807)
- \[[`6834ca13bb`](https://redirect.github.com/nodejs/node/commit/6834ca13bb)] - **(SEMVER-MINOR)** **stream**: rename `Duplex.toWeb()` type option to `readableType` (René) [#&#8203;61632](https://redirect.github.com/nodejs/node/pull/61632)
- \[[`5ed5474437`](https://redirect.github.com/nodejs/node/commit/5ed5474437)] - **test**: update WPT for WebCryptoAPI to [`2cb332d`](https://redirect.github.com/nodejs/node/commit/2cb332d710) (Node.js GitHub Bot) [#&#8203;62483](https://redirect.github.com/nodejs/node/pull/62483)
- \[[`3c9c0f8577`](https://redirect.github.com/nodejs/node/commit/3c9c0f8577)] - **test**: fix test-buffer-zero-fill-cli to be effective (Сковорода Никита Андреевич) [#&#8203;60623](https://redirect.github.com/nodejs/node/pull/60623)
- \[[`19a52a1abe`](https://redirect.github.com/nodejs/node/commit/19a52a1abe)] - **test**: update WPT for url to [`fc3e651`](https://redirect.github.com/nodejs/node/commit/fc3e651593) (Node.js GitHub Bot) [#&#8203;62379](https://redirect.github.com/nodejs/node/pull/62379)
- \[[`111ba9bd5b`](https://redirect.github.com/nodejs/node/commit/111ba9bd5b)] - **test**: wait for reattach before initial break on restart (Yuya Inoue) [#&#8203;62471](https://redirect.github.com/nodejs/node/pull/62471)
- \[[`0897c6cc08`](https://redirect.github.com/nodejs/node/commit/0897c6cc08)] - **test**: disable flaky WPT Blob test on AIX (James M Snell) [#&#8203;62470](https://redirect.github.com/nodejs/node/pull/62470)
- \[[`1c3d93bfab`](https://redirect.github.com/nodejs/node/commit/1c3d93bfab)] - **test**: avoid flaky run wait in debugger restart test (Yuya Inoue) [#&#8203;62112](https://redirect.github.com/nodejs/node/pull/62112)
- \[[`83416a640a`](https://redirect.github.com/nodejs/node/commit/83416a640a)] - **test**: skip test-cluster-dgram-reuse on AIX 7.3 (Stewart X Addison) [#&#8203;62238](https://redirect.github.com/nodejs/node/pull/62238)
- \[[`af8d0922dd`](https://redirect.github.com/nodejs/node/commit/af8d0922dd)] - **test**: add WebCrypto Promise.prototype.then pollution regression tests (Filip Skokan) [#&#8203;62226](https://redirect.github.com/nodejs/node/pull/62226)
- \[[`fc9a60ec74`](https://redirect.github.com/nodejs/node/commit/fc9a60ec74)] - **test**: update WPT for WebCryptoAPI to [`6a1c545`](https://redirect.github.com/nodejs/node/commit/6a1c545d77) (Node.js GitHub Bot) [#&#8203;62187](https://redirect.github.com/nodejs/node/pull/62187)
- \[[`12ba2d74fe`](https://redirect.github.com/nodejs/node/commit/12ba2d74fe)] - **test**: update WPT for url to [`c928b19`](https://redirect.github.com/nodejs/node/commit/c928b19ab0) (Node.js GitHub Bot) [#&#8203;62148](https://redirect.github.com/nodejs/node/pull/62148)
- \[[`4e15e5b647`](https://redirect.github.com/nodejs/node/commit/4e15e5b647)] - **test**: update WPT for WebCryptoAPI to [`c9e9558`](https://redirect.github.com/nodejs/node/commit/c9e955840a) (Node.js GitHub Bot) [#&#8203;62147](https://redirect.github.com/nodejs/node/pull/62147)
- \[[`dc66a05558`](https://redirect.github.com/nodejs/node/commit/dc66a05558)] - **test**: improve WPT report runner (Filip Skokan) [#&#8203;62107](https://redirect.github.com/nodejs/node/pull/62107)
- \[[`9536e5621b`](https://redirect.github.com/nodejs/node/commit/9536e5621b)] - **test**: update WPT compression to [`ae05f5c`](https://redirect.github.com/nodejs/node/commit/ae05f5cb53) (Filip Skokan) [#&#8203;62107](https://redirect.github.com/nodejs/node/pull/62107)
- \[[`fb1c0bda0a`](https://redirect.github.com/nodejs/node/commit/fb1c0bda0a)] - **test**: update WPT for WebCryptoAPI to [`42e4732`](https://redirect.github.com/nodejs/node/commit/42e47329fd) (Node.js GitHub Bot) [#&#8203;62048](https://redirect.github.com/nodejs/node/pull/62048)
- \[[`d886f27485`](https://redirect.github.com/nodejs/node/commit/d886f27485)] - **test**: fix skipping behavior for `test-runner-run-files-undefined` (Antoine du Hamel) [#&#8203;62026](https://redirect.github.com/nodejs/node/pull/62026)
- \[[`f79df03e0b`](https://redirect.github.com/nodejs/node/commit/f79df03e0b)] - **test**: remove unnecessary `process.exit` calls from test files (Antoine du Hamel) [#&#8203;62020](https://redirect.github.com/nodejs/node/pull/62020)
- \[[`1319295467`](https://redirect.github.com/nodejs/node/commit/1319295467)] - **test**: skip `test-url` on `--shared-ada` builds (Antoine du Hamel) [#&#8203;62019](https://redirect.github.com/nodejs/node/pull/62019)
- \[[`2ea06727c6`](https://redirect.github.com/nodejs/node/commit/2ea06727c6)] - **test**: skip strace test with shared openssl (Richard Lau) [#&#8203;61987](https://redirect.github.com/nodejs/node/pull/61987)
- \[[`c0680d5df7`](https://redirect.github.com/nodejs/node/commit/c0680d5df7)] - **test**: avoid flaky debugger restart waits (Yuya Inoue) [#&#8203;61773](https://redirect.github.com/nodejs/node/pull/61773)
- \[[`22b748ef72`](https://redirect.github.com/nodejs/node/commit/22b748ef72)] - **test**: fix typos in test files (Daijiro Wachi) [#&#8203;61408](https://redirect.github.com/nodejs/node/pull/61408)
- \[[`a20bf9a84d`](https://redirect.github.com/nodejs/node/commit/a20bf9a84d)] - **test**: allow filtering async internal frames in assertSnapshot (Joyee Cheung) [#&#8203;61769](https://redirect.github.com/nodejs/node/pull/61769)
- \[[`ec2913f036`](https://redirect.github.com/nodejs/node/commit/ec2913f036)] - **test**: unify assertSnapshot stacktrace transform (Chengzhong Wu) [#&#8203;61665](https://redirect.github.com/nodejs/node/pull/61665)
- \[[`460f41233d`](https://redirect.github.com/nodejs/node/commit/460f41233d)] - **test**: check stability block position in API markdown (René) [#&#8203;58590](https://redirect.github.com/nodejs/node/pull/58590)
- \[[`9ad02065d5`](https://redirect.github.com/nodejs/node/commit/9ad02065d5)] - **test**: adapt buffer test for v8 sandbox (Shelley Vohr) [#&#8203;61772](https://redirect.github.com/nodejs/node/pull/61772)
- \[[`5cf001736e`](https://redirect.github.com/nodejs/node/commit/5cf001736e)] - **test**: update FileAPI tests from WPT (Ms2ger) [#&#8203;61750](https://redirect.github.com/nodejs/node/pull/61750)
- \[[`84c7a23223`](https://redirect.github.com/nodejs/node/commit/84c7a23223)] - **test**: update WPT for WebCryptoAPI to [`7cbe7e8`](https://redirect.github.com/nodejs/node/commit/7cbe7e8ed9) (Node.js GitHub Bot) [#&#8203;61729](https://redirect.github.com/nodejs/node/pull/61729)
- \[[`276a32fd10`](https://redirect.github.com/nodejs/node/commit/276a32fd10)] - **test**: update WPT for url to [`efb889e`](https://redirect.github.com/nodejs/node/commit/efb889eb4c) (Node.js GitHub Bot) [#&#8203;61728](https://redirect.github.com/nodejs/node/pull/61728)
- \[[`f5f21d36a6`](https://redirect.github.com/nodejs/node/commit/f5f21d36a6)] - **test\_runner**: add exports option for module mocks (sangwook) [#&#8203;61727](https://redirect.github.com/nodejs/node/pull/61727)
- \[[`bfc8a12977`](https://redirect.github.com/nodejs/node/commit/bfc8a12977)] - **test\_runner**: make it compatible with fake timers (Matteo Collina) [#&#8203;59272](https://redirect.github.com/nodejs/node/pull/59272)
- \[[`e0cde40e1d`](https://redirect.github.com/nodejs/node/commit/e0cde40e1d)] - **test\_runner**: set non-zero exit code when suite errors occur (Edy Silva) [#&#8203;62282](https://redirect.github.com/nodejs/node/pull/62282)
- \[[`d74efd6834`](https://redirect.github.com/nodejs/node/commit/d74efd6834)] - **test\_runner**: run afterEach on runtime skip (Igor Shevelenkov) [#&#8203;61525](https://redirect.github.com/nodejs/node/pull/61525)
- \[[`8287ca749e`](https://redirect.github.com/nodejs/node/commit/8287ca749e)] - **test\_runner**: expose expectFailure message (sangwook) [#&#8203;61563](https://redirect.github.com/nodejs/node/pull/61563)
- \[[`1f2025fd1e`](https://redirect.github.com/nodejs/node/commit/1f2025fd1e)] - **(SEMVER-MINOR)** **test\_runner**: expose worker ID for concurrent test execution (Ali Hassan) [#&#8203;61394](https://redirect.github.com/nodejs/node/pull/61394)
- \[[`b1199c7bb4`](https://redirect.github.com/nodejs/node/commit/b1199c7bb4)] - **test\_runner**: replace native methods with primordials (Ayoub Mabrouk) [#&#8203;61219](https://redirect.github.com/nodejs/node/pull/61219)
- \[[`1ca20fc33d`](https://redirect.github.com/nodejs/node/commit/1ca20fc33d)] - **(SEMVER-MINOR)** **test\_runner**: show interrupted test on SIGINT (Matteo Collina) [#&#8203;61676](https://redirect.github.com/nodejs/node/pull/61676)
- \[[`207ba4f89f`](https://redirect.github.com/nodejs/node/commit/207ba4f89f)] - **test\_runner**: fix suite rerun (Moshe Atlow) [#&#8203;61775](https://redirect.github.com/nodejs/node/pull/61775)
- \[[`9927335c11`](https://redirect.github.com/nodejs/node/commit/9927335c11)] - **tls**: forward keepAlive, keepAliveInitialDelay, noDelay to socket (Sergey Zelenov) [#&#8203;62004](https://redirect.github.com/nodejs/node/pull/62004)
- \[[`a1c3c901c0`](https://redirect.github.com/nodejs/node/commit/a1c3c901c0)] - **tools**: bump picomatch from 4.0.3 to 4.0.4 in /tools/eslint (dependabot\[bot]) [#&#8203;62439](https://redirect.github.com/nodejs/node/pull/62439)
- \[[`1c6f5ed7c2`](https://redirect.github.com/nodejs/node/commit/1c6f5ed7c2)] - **tools**: adopt the `--check-for-duplicates` NCU flag (Antoine du Hamel) [#&#8203;62478](https://redirect.github.com/nodejs/node/pull/62478)
- \[[`b53377e8fe`](https://redirect.github.com/nodejs/node/commit/b53377e8fe)] - **tools**: bump flatted from 3.4.1 to 3.4.2 in /tools/eslint (dependabot\[bot]) [#&#8203;62375](https://redirect.github.com/nodejs/node/pull/62375)
- \[[`f102e79b80`](https://redirect.github.com/nodejs/node/commit/f102e79b80)] - **tools**: bump eslint deps (Huáng Jùnliàng) [#&#8203;62356](https://redirect.github.com/nodejs/node/pull/62356)
- \[[`f5d74f8216`](https://redirect.github.com/nodejs/node/commit/f5d74f8216)] - **tools**: add eslint-plugin-regexp (Huáng Jùnliàng) [#&#8203;62093](https://redirect.github.com/nodejs/node/pull/62093)
- \[[`bc5b9a04ad`](https://redirect.github.com/nodejs/node/commit/bc5b9a04ad)] - **tools**: bump flatted from 3.3.3 to 3.4.1 in /tools/eslint (dependabot\[bot]) [#&#8203;62255](https://redirect.github.com/nodejs/node/pull/62255)
- \[[`bad48b9700`](https://redirect.github.com/nodejs/node/commit/bad48b9700)] - **tools**: validate all commits that are pushed to `main` (Antoine du Hamel) [#&#8203;62246](https://redirect.github.com/nodejs/node/pull/62246)
- \[[`795d663ff4`](https://redirect.github.com/nodejs/node/commit/795d663ff4)] - **tools**: keep GN files when updating Merve (Antoine du Hamel) [#&#8203;62167](https://redirect.github.com/nodejs/node/pull/62167)
- \[[`0b6fa913f1`](https://redirect.github.com/nodejs/node/commit/0b6fa913f1)] - **tools**: revert timezone update GHA workflow to ubuntu-latest (Richard Lau) [#&#8203;62140](https://redirect.github.com/nodejs/node/pull/62140)
- \[[`840e098e99`](https://redirect.github.com/nodejs/node/commit/840e098e99)] - **tools**: improve error handling in test426 update script (Rich Trott) [#&#8203;62121](https://redirect.github.com/nodejs/node/pull/62121)
- \[[`bd34e53a8e`](https://redirect.github.com/nodejs/node/commit/bd34e53a8e)] - **tools**: bump the eslint group across 1 directory with 2 updates (dependabot\[bo

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

## Reviews

### Review by @swadeley - Approved on April 20, 2026 at 07:26 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1473*
