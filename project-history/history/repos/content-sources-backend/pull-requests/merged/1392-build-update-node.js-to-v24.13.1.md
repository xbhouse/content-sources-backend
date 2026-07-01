---
type: pull_request
number: 1392
title: "Build: Update Node.js to v24.13.1"
state: merged
author: red-hat-konflux
created: 2026-02-16T04:52:33Z
updated: 2026-02-16T05:30:54Z
closed: 2026-02-16T05:30:53Z
merged: 2026-02-16T05:30:53Z
base_branch: main
head_branch: konflux/mintmaker/main/node-24.x
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1392
---

# Pull Request #1392: Build: Update Node.js to v24.13.1

**Author**: @red-hat-konflux
**Created**: February 16, 2026 at 04:52 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/node-24.x`

## Description

> **Note:** This PR body was truncated due to platform limits.

This PR contains the following updates:

| Package | Update | Change |
|---|---|---|
| [node](https://nodejs.org) ([source](https://redirect.github.com/nodejs/node)) | patch | `24.13.0` -> `24.13.1` |

---

### Release Notes

<details>
<summary>nodejs/node (node)</summary>

### [`v24.13.1`](https://redirect.github.com/nodejs/node/releases/tag/v24.13.1): 2026-02-10, Version 24.13.1 &#x27;Krypton&#x27; (LTS), @&#8203;aduh95

[Compare Source](https://redirect.github.com/nodejs/node/compare/v24.13.0...v24.13.1)

##### Notable Changes

- \[[`1f64d6841e`](https://redirect.github.com/nodejs/node/commit/1f64d6841e)] - **build**: add support for Python 3.14 (Christian Clauss) [#&#8203;59983](https://redirect.github.com/nodejs/node/pull/59983)
- \[[`30e500fc09`](https://redirect.github.com/nodejs/node/commit/30e500fc09)] - **cli**: mark `--heapsnapshot-near-heap-limit` as stable (Joyee Cheung) [#&#8203;60956](https://redirect.github.com/nodejs/node/pull/60956)
- \[[`bc0a55f086`](https://redirect.github.com/nodejs/node/commit/bc0a55f086)] - **crypto**: update root certificates to NSS 3.119 (Node.js GitHub Bot) [#&#8203;61419](https://redirect.github.com/nodejs/node/pull/61419)
- \[[`8a67c00bf5`](https://redirect.github.com/nodejs/node/commit/8a67c00bf5)] - **doc**: mark `--build-snapshot` and `--build-snapshot-config` as stable (Joyee Cheung) [#&#8203;60954](https://redirect.github.com/nodejs/node/pull/60954)
- \[[`3999c2a910`](https://redirect.github.com/nodejs/node/commit/3999c2a910)] - **meta**: add avivkeller to collaborators (Aviv Keller) [#&#8203;61115](https://redirect.github.com/nodejs/node/pull/61115)
- \[[`fa542fbae6`](https://redirect.github.com/nodejs/node/commit/fa542fbae6)] - **meta**: add gurgunday to collaborators (Gürgün Dayıoğlu) [#&#8203;61094](https://redirect.github.com/nodejs/node/pull/61094)
- \[[`ff11eda2f2`](https://redirect.github.com/nodejs/node/commit/ff11eda2f2)] - **meta**: add Renegade334 to collaborators (Renegade334) [#&#8203;60714](https://redirect.github.com/nodejs/node/pull/60714)
- \[[`2e387fb969`](https://redirect.github.com/nodejs/node/commit/2e387fb969)] - **url**: update ada to v3.4.2 and support unicode 17 (Yagiz Nizipli) [#&#8203;61593](https://redirect.github.com/nodejs/node/pull/61593)
- \[[`bb206782d4`](https://redirect.github.com/nodejs/node/commit/bb206782d4)] - **v8**: mark `v8.queryObjects()` as stable (Joyee Cheung) [#&#8203;60957](https://redirect.github.com/nodejs/node/pull/60957)

##### Commits

- \[[`a73279c60d`](https://redirect.github.com/nodejs/node/commit/a73279c60d)] - **assert**: use a set instead of an array for faster lookup (Ruben Bridgewater) [#&#8203;61076](https://redirect.github.com/nodejs/node/pull/61076)
- \[[`6a61bcd73c`](https://redirect.github.com/nodejs/node/commit/6a61bcd73c)] - **assert,util**: fix deep comparison for sets and maps with mixed types (Ruben Bridgewater) [#&#8203;61388](https://redirect.github.com/nodejs/node/pull/61388)
- \[[`cf0eabcd42`](https://redirect.github.com/nodejs/node/commit/cf0eabcd42)] - **assert,util**: improve deep comparison performance (Ruben Bridgewater) [#&#8203;61076](https://redirect.github.com/nodejs/node/pull/61076)
- \[[`ff3b9ac183`](https://redirect.github.com/nodejs/node/commit/ff3b9ac183)] - **benchmark**: add SQLite benchmarks (Guilherme Araújo) [#&#8203;61401](https://redirect.github.com/nodejs/node/pull/61401)
- \[[`e1f7d68c94`](https://redirect.github.com/nodejs/node/commit/e1f7d68c94)] - **benchmark**: use boolean options in benchmark tests (SeokhunEom) [#&#8203;60129](https://redirect.github.com/nodejs/node/pull/60129)
- \[[`91127c91cd`](https://redirect.github.com/nodejs/node/commit/91127c91cd)] - **benchmark**: allow boolean option values (SeokhunEom) [#&#8203;60129](https://redirect.github.com/nodejs/node/pull/60129)
- \[[`170fda55f6`](https://redirect.github.com/nodejs/node/commit/170fda55f6)] - **benchmark**: add microbench on isInsideNodeModules (Chengzhong Wu) [#&#8203;60991](https://redirect.github.com/nodejs/node/pull/60991)
- \[[`3976381b41`](https://redirect.github.com/nodejs/node/commit/3976381b41)] - **benchmark**: fix incorrect base64 input in byteLength benchmark (semimikoh) [#&#8203;60841](https://redirect.github.com/nodejs/node/pull/60841)
- \[[`c702fccd76`](https://redirect.github.com/nodejs/node/commit/c702fccd76)] - **benchmark**: use typescript for import cjs benchmark (Joyee Cheung) [#&#8203;60663](https://redirect.github.com/nodejs/node/pull/60663)
- \[[`92c517c62d`](https://redirect.github.com/nodejs/node/commit/92c517c62d)] - **buffer**: make methods work on Uint8Array instances (Neal Beeken) [#&#8203;56578](https://redirect.github.com/nodejs/node/pull/56578)
- \[[`be95382edb`](https://redirect.github.com/nodejs/node/commit/be95382edb)] - **buffer**: let Buffer.of use heap (Сковорода Никита Андреевич) [#&#8203;60503](https://redirect.github.com/nodejs/node/pull/60503)
- \[[`1f64d6841e`](https://redirect.github.com/nodejs/node/commit/1f64d6841e)] - **build**: test on Python 3.14 (Christian Clauss) [#&#8203;59983](https://redirect.github.com/nodejs/node/pull/59983)
- \[[`ea4687981b`](https://redirect.github.com/nodejs/node/commit/ea4687981b)] - **build**: update android-patches/trap-handler.h.patch (Mo Luo) [#&#8203;60369](https://redirect.github.com/nodejs/node/pull/60369)
- \[[`b3a7a8c780`](https://redirect.github.com/nodejs/node/commit/b3a7a8c780)] - **build**: update devcontainer.json to use paired nix env (Joyee Cheung) [#&#8203;61414](https://redirect.github.com/nodejs/node/pull/61414)
- \[[`7168d0b5e3`](https://redirect.github.com/nodejs/node/commit/7168d0b5e3)] - **build**: add embedtest into native suite (Joyee Cheung) [#&#8203;61357](https://redirect.github.com/nodejs/node/pull/61357)
- \[[`e00755a977`](https://redirect.github.com/nodejs/node/commit/e00755a977)] - **build**: fix misplaced comma in ldflags (hqzing) [#&#8203;61294](https://redirect.github.com/nodejs/node/pull/61294)
- \[[`72fcc3ee9d`](https://redirect.github.com/nodejs/node/commit/72fcc3ee9d)] - **build**: fix crate vendor file checksums on windows (Chengzhong Wu) [#&#8203;61329](https://redirect.github.com/nodejs/node/pull/61329)
- \[[`76a73d68fd`](https://redirect.github.com/nodejs/node/commit/76a73d68fd)] - **build**: expose libplatform symbols in shared libnode (Joyee Cheung) [#&#8203;61144](https://redirect.github.com/nodejs/node/pull/61144)
- \[[`ef8d26ce5c`](https://redirect.github.com/nodejs/node/commit/ef8d26ce5c)] - **build**: fix inconsistent quoting in `Makefile` (Antoine du Hamel) [#&#8203;60511](https://redirect.github.com/nodejs/node/pull/60511)
- \[[`2d23968783`](https://redirect.github.com/nodejs/node/commit/2d23968783)] - **build**: remove temporal updater (Chengzhong Wu) [#&#8203;61151](https://redirect.github.com/nodejs/node/pull/61151)
- \[[`4c2655f1c2`](https://redirect.github.com/nodejs/node/commit/4c2655f1c2)] - **build**: update test-wpt-report to use NODE instead of OUT\_NODE (Filip Skokan) [#&#8203;61024](https://redirect.github.com/nodejs/node/pull/61024)
- \[[`eaea6821fc`](https://redirect.github.com/nodejs/node/commit/eaea6821fc)] - **build**: skip build-ci on actions with a separate test step (Chengzhong Wu) [#&#8203;61073](https://redirect.github.com/nodejs/node/pull/61073)
- \[[`dfd4e12037`](https://redirect.github.com/nodejs/node/commit/dfd4e12037)] - **build**: run embedtest with node\_g when BUILDTYPE=Debug (Chengzhong Wu) [#&#8203;60850](https://redirect.github.com/nodejs/node/pull/60850)
- \[[`775c77234b`](https://redirect.github.com/nodejs/node/commit/775c77234b)] - **build,tools**: fix addon build deadlock on errors (Vladimir Morozov) [#&#8203;61321](https://redirect.github.com/nodejs/node/pull/61321)
- \[[`5deafc10fa`](https://redirect.github.com/nodejs/node/commit/5deafc10fa)] - **build,win**: improve logs when ClangCL is missing (Mike McCready) [#&#8203;61438](https://redirect.github.com/nodejs/node/pull/61438)
- \[[`e2481c5c6e`](https://redirect.github.com/nodejs/node/commit/e2481c5c6e)] - **build,win**: update WinGet configurations to Python 3.14 (Mike McCready) [#&#8203;61431](https://redirect.github.com/nodejs/node/pull/61431)
- \[[`d2586b7e4c`](https://redirect.github.com/nodejs/node/commit/d2586b7e4c)] - **child\_process**: treat ipc length header as unsigned uint32 (Ryuhei Shima) [#&#8203;61344](https://redirect.github.com/nodejs/node/pull/61344)
- \[[`30e500fc09`](https://redirect.github.com/nodejs/node/commit/30e500fc09)] - **cli**: mark --heapsnapshot-near-heap-limit as stable (Joyee Cheung) [#&#8203;60956](https://redirect.github.com/nodejs/node/pull/60956)
- \[[`2c7da15612`](https://redirect.github.com/nodejs/node/commit/2c7da15612)] - **cluster**: fix port reuse between cluster (Ryuhei Shima) [#&#8203;60141](https://redirect.github.com/nodejs/node/pull/60141)
- \[[`bc0a55f086`](https://redirect.github.com/nodejs/node/commit/bc0a55f086)] - **crypto**: update root certificates to NSS 3.119 (Node.js GitHub Bot) [#&#8203;61419](https://redirect.github.com/nodejs/node/pull/61419)
- \[[`2d5f20e9c3`](https://redirect.github.com/nodejs/node/commit/2d5f20e9c3)] - **crypto**: update root certificates to NSS 3.117 (Node.js GitHub Bot) [#&#8203;60741](https://redirect.github.com/nodejs/node/pull/60741)
- \[[`fba95be188`](https://redirect.github.com/nodejs/node/commit/fba95be188)] - **deps**: update archs files for openssl-3.5.5 (Node.js GitHub Bot) [#&#8203;61547](https://redirect.github.com/nodejs/node/pull/61547)
- \[[`08697289e0`](https://redirect.github.com/nodejs/node/commit/08697289e0)] - **deps**: upgrade openssl sources to openssl-3.5.5 (Node.js GitHub Bot) [#&#8203;61547](https://redirect.github.com/nodejs/node/pull/61547)
- \[[`403c50c04d`](https://redirect.github.com/nodejs/node/commit/403c50c04d)] - **deps**: update corepack to 0.34.6 (Node.js GitHub Bot) [#&#8203;61510](https://redirect.github.com/nodejs/node/pull/61510)
- \[[`3b24691aeb`](https://redirect.github.com/nodejs/node/commit/3b24691aeb)] - **deps**: upgrade npm to 11.8.0 (npm team) [#&#8203;61466](https://redirect.github.com/nodejs/node/pull/61466)
- \[[`2bba7efdc4`](https://redirect.github.com/nodejs/node/commit/2bba7efdc4)] - **deps**: update googletest to [`8508785`](https://redirect.github.com/nodejs/node/commit/85087857ad10bd407cd6ed2f52f7ea9752db621f) (Node.js GitHub Bot) [#&#8203;61417](https://redirect.github.com/nodejs/node/pull/61417)
- \[[`8f8c6f6162`](https://redirect.github.com/nodejs/node/commit/8f8c6f6162)] - **deps**: update sqlite to 3.51.2 (Node.js GitHub Bot) [#&#8203;61339](https://redirect.github.com/nodejs/node/pull/61339)
- \[[`c46009053c`](https://redirect.github.com/nodejs/node/commit/c46009053c)] - **deps**: update icu to 78.2 (Node.js GitHub Bot) [#&#8203;60523](https://redirect.github.com/nodejs/node/pull/60523)
- \[[`b46b8dd91b`](https://redirect.github.com/nodejs/node/commit/b46b8dd91b)] - **deps**: update ada to v3.4.0 (Yagiz Nizipli) [#&#8203;61315](https://redirect.github.com/nodejs/node/pull/61315)
- \[[`88c6b17e18`](https://redirect.github.com/nodejs/node/commit/88c6b17e18)] - **deps**: update zlib to 1.3.1-e00f703 (Node.js GitHub Bot) [#&#8203;61135](https://redirect.github.com/nodejs/node/pull/61135)
- \[[`0030c05ba9`](https://redirect.github.com/nodejs/node/commit/0030c05ba9)] - **deps**: update cjs-module-lexer to 2.2.0 (Node.js GitHub Bot) [#&#8203;61271](https://redirect.github.com/nodejs/node/pull/61271)
- \[[`77437cff89`](https://redirect.github.com/nodejs/node/commit/77437cff89)] - **deps**: update nbytes to 0.1.2 (Node.js GitHub Bot) [#&#8203;61270](https://redirect.github.com/nodejs/node/pull/61270)
- \[[`fb0f05a937`](https://redirect.github.com/nodejs/node/commit/fb0f05a937)] - **deps**: update timezone to 2025c (Node.js GitHub Bot) [#&#8203;61138](https://redirect.github.com/nodejs/node/pull/61138)
- \[[`b426a47c05`](https://redirect.github.com/nodejs/node/commit/b426a47c05)] - **deps**: nghttp2: revert [`7784fa9`](https://redirect.github.com/nodejs/node/commit/7784fa979d0b) (Antoine du Hamel) [#&#8203;61136](https://redirect.github.com/nodejs/node/pull/61136)
- \[[`c07a38f700`](https://redirect.github.com/nodejs/node/commit/c07a38f700)] - **deps**: update nghttp2 to 1.68.0 (nodejs-github-bot) [#&#8203;61136](https://redirect.github.com/nodejs/node/pull/61136)
- \[[`c2ddc9a18b`](https://redirect.github.com/nodejs/node/commit/c2ddc9a18b)] - **deps**: update simdjson to 4.2.4 (Node.js GitHub Bot) [#&#8203;61056](https://redirect.github.com/nodejs/node/pull/61056)
- \[[`f38cd6da8e`](https://redirect.github.com/nodejs/node/commit/f38cd6da8e)] - **deps**: update googletest to [`065127f`](https://redirect.github.com/nodejs/node/commit/065127f1e4b46c5f14fc73cf8d323c221f9dc68e) (Node.js GitHub Bot) [#&#8203;61055](https://redirect.github.com/nodejs/node/pull/61055)
- \[[`a9a6a4cdb2`](https://redirect.github.com/nodejs/node/commit/a9a6a4cdb2)] - **deps**: brotli: cherry-pick [`e230f47`](https://redirect.github.com/nodejs/node/commit/e230f474b87) (liujiahui) [#&#8203;61003](https://redirect.github.com/nodejs/node/pull/61003)
- \[[`5a40023aae`](https://redirect.github.com/nodejs/node/commit/5a40023aae)] - **deps**: upgrade npm to 11.7.0 (npm team) [#&#8203;61011](https://redirect.github.com/nodejs/node/pull/61011)
- \[[`4121e7a413`](https://redirect.github.com/nodejs/node/commit/4121e7a413)] - **deps**: update sqlite to 3.51.1 (Node.js GitHub Bot) [#&#8203;60899](https://redirect.github.com/nodejs/node/pull/60899)
- \[[`e8a09fc896`](https://redirect.github.com/nodejs/node/commit/e8a09fc896)] - **deps**: update zlib to 1.3.1-63d7e16 (Node.js GitHub Bot) [#&#8203;60898](https://redirect.github.com/nodejs/node/pull/60898)
- \[[`8df5862ee5`](https://redirect.github.com/nodejs/node/commit/8df5862ee5)] - **deps**: upgrade npm to 11.6.4 (npm team) [#&#8203;60853](https://redirect.github.com/nodejs/node/pull/60853)
- \[[`6c1c8cbdcc`](https://redirect.github.com/nodejs/node/commit/6c1c8cbdcc)] - **deps**: update sqlite to 3.51.0 (Node.js GitHub Bot) [#&#8203;60614](https://redirect.github.com/nodejs/node/pull/60614)
- \[[`2d1efc7c1b`](https://redirect.github.com/nodejs/node/commit/2d1efc7c1b)] - **deps**: upgrade npm to 11.6.3 (npm team) [#&#8203;60785](https://redirect.github.com/nodejs/node/pull/60785)
- \[[`3a2de1c23b`](https://redirect.github.com/nodejs/node/commit/3a2de1c23b)] - **deps**: update brotli to 1.2.0 (Node.js GitHub Bot) [#&#8203;60540](https://redirect.github.com/nodejs/node/pull/60540)
- \[[`58c5d40bd1`](https://redirect.github.com/nodejs/node/commit/58c5d40bd1)] - **deps**: update simdjson to 4.2.2 (Node.js GitHub Bot) [#&#8203;60740](https://redirect.github.com/nodejs/node/pull/60740)
- \[[`e6b607ef50`](https://redirect.github.com/nodejs/node/commit/e6b607ef50)] - **deps**: update googletest to [`1b96fa1`](https://redirect.github.com/nodejs/node/commit/1b96fa13f549387b7549cc89e1a785cf143a1a50) (Node.js GitHub Bot) [#&#8203;60739](https://redirect.github.com/nodejs/node/pull/60739)
- \[[`650c9e0305`](https://redirect.github.com/nodejs/node/commit/650c9e0305)] - **deps**: update minimatch to 10.1.1 (Node.js GitHub Bot) [#&#8203;60543](https://redirect.github.com/nodejs/node/pull/60543)
- \[[`ef1951d5d5`](https://redirect.github.com/nodejs/node/commit/ef1951d5d5)] - **deps**: update inspector\_protocol to [`1b1bcbb`](https://redirect.github.com/nodejs/node/commit/1b1bcbbe060e8c8cd8704f00f78978c50991) (Node.js GitHub Bot) [#&#8203;60705](https://redirect.github.com/nodejs/node/pull/60705)
- \[[`eb068305dd`](https://redirect.github.com/nodejs/node/commit/eb068305dd)] - **deps**: update cjs-module-lexer to 2.1.1 (Node.js GitHub Bot) [#&#8203;60646](https://redirect.github.com/nodejs/node/pull/60646)
- \[[`ee1d99131c`](https://redirect.github.com/nodejs/node/commit/ee1d99131c)] - **deps**: update simdjson to 4.2.1 (Node.js GitHub Bot) [#&#8203;60644](https://redirect.github.com/nodejs/node/pull/60644)
- \[[`23582967b7`](https://redirect.github.com/nodejs/node/commit/23582967b7)] - **deps**: V8: cherry-pick [`1441665`](https://redirect.github.com/nodejs/node/commit/1441665e0d87) (Domagoj Stolfa) [#&#8203;60989](https://redirect.github.com/nodejs/node/pull/60989)
- \[[`155eaedff2`](https://redirect.github.com/nodejs/node/commit/155eaedff2)] - **deps**: V8: cherry-pick [`394a805`](https://redirect.github.com/nodejs/node/commit/394a8053b59e) (Lu Yahan) [#&#8203;60962](https://redirect.github.com/nodejs/node/pull/60962)
- \[[`c95a4a0f43`](https://redirect.github.com/nodejs/node/commit/c95a4a0f43)] - **deps**: V8: backport [`bbaae8e`](https://redirect.github.com/nodejs/node/commit/bbaae8e36164) (Lu Yahan) [#&#8203;60962](https://redirect.github.com/nodejs/node/pull/60962)
- \[[`6f123f186d`](https://redirect.github.com/nodejs/node/commit/6f123f186d)] - **doc**: move Security-Team from TSC to SECURITY (Rafael Gonzaga) [#&#8203;61495](https://redirect.github.com/nodejs/node/pull/61495)
- \[[`2e3337d15b`](https://redirect.github.com/nodejs/node/commit/2e3337d15b)] - **doc**: added `requestOCSP` option to `tls.connect` (ikeyan) [#&#8203;61064](https://redirect.github.com/nodejs/node/pull/61064)
- \[[`f505f81577`](https://redirect.github.com/nodejs/node/commit/f505f81577)] - **doc**: restore [@&#8203;ChALkeR](https://redirect.github.com/ChALkeR) to collaborators (Сковорода Никита Андреевич) [#&#8203;61553](https://redirect.github.com/nodejs/node/pull/61553)
- \[[`12fb95d0c9`](https://redirect.github.com/nodejs/node/commit/12fb95d0c9)] - **doc**: update IBM/Red Hat volunteers with dedicated project time (Beth Griggs) [#&#8203;61588](https://redirect.github.com/nodejs/node/pull/61588)
- \[[`283ab61ed2`](https://redirect.github.com/nodejs/node/commit/283ab61ed2)] - **doc**: align Buffer.concat documentation with behavior (Gürgün Dayıoğlu) [#&#8203;60405](https://redirect.github.com/nodejs/node/pull/60405)
- \[[`fc9c906d5f`](https://redirect.github.com/nodejs/node/commit/fc9c906d5f)] - **doc**: remove `v` prefix for version references (Mike McCready) [#&#8203;61488](https://redirect.github.com/nodejs/node/pull/61488)
- \[[`4a88ed09e8`](https://redirect.github.com/nodejs/node/commit/4a88ed09e8)] - **doc**: mention constructor comparison in assert.deepStrictEqual (Hamza Kargin) [#&#8203;60253](https://redirect.github.com/nodejs/node/pull/60253)
- \[[`9b29d56491`](https://redirect.github.com/nodejs/node/commit/9b29d56491)] - **doc**: add CVE delay mention (Rafael Gonzaga) [#&#8203;61465](https://redirect.github.com/nodejs/node/pull/61465)
- \[[`4815e4ac52`](https://redirect.github.com/nodejs/node/commit/4815e4ac52)] - **doc**: update previous version links in BUILDING (Mike McCready) [#&#8203;61457](https://redirect.github.com/nodejs/node/pull/61457)
- \[[`8a43244e6c`](https://redirect.github.com/nodejs/node/commit/8a43244e6c)] - **doc**: include OpenJSF handle for security stewards (Rafael Gonzaga) [#&#8203;61454](https://redirect.github.com/nodejs/node/pull/61454)
- \[[`89a7f184a1`](https://redirect.github.com/nodejs/node/commit/89a7f184a1)] - **doc**: clarify process.argv\[1] behavior for -e/--eval (Jeevankumar S) [#&#8203;61366](https://redirect.github.com/nodejs/node/pull/61366)
- \[[`b4041aba1c`](https://redirect.github.com/nodejs/node/commit/b4041aba1c)] - **doc**: remove Windows Dev Home instructions from BUILDING (Mike McCready) [#&#8203;61434](https://redirect.github.com/nodejs/node/pull/61434)
- \[[`fa7830bac0`](https://redirect.github.com/nodejs/node/commit/fa7830bac0)] - **doc**: clarify TypedArray properties on Buffer (Roman Reiss) [#&#8203;61355](https://redirect.github.com/nodejs/node/pull/61355)
- \[[`45663c8956`](https://redirect.github.com/nodejs/node/commit/45663c8956)] - **doc**: update Python 3.14 manual install instructions (Windows) (Mike McCready) [#&#8203;61428](https://redirect.github.com/nodejs/node/pull/61428)
- \[[`0248357f26`](https://redirect.github.com/nodejs/node/commit/0248357f26)] - **doc**: note resume build should not be done on node-test-commit (Stewart X Addison) [#&#8203;61373](https://redirect.github.com/nodejs/node/pull/61373)
- \[[`b254bab513`](https://redirect.github.com/nodejs/node/commit/b254bab513)] - **doc**: refine WebAssembly error documentation (sangwook) [#&#8203;61382](https://redirect.github.com/nodejs/node/pull/61382)
- \[[`8aca37c6ef`](https://redirect.github.com/nodejs/node/commit/8aca37c6ef)] - **doc**: add deprecation history for url.parse (Eng Zer Jun) [#&#8203;61389](https://redirect.github.com/nodejs/node/pull/61389)
- \[[`8047ac3aac`](https://redirect.github.com/nodejs/node/commit/8047ac3aac)] - **doc**: add marco and rafael in last sec release (Marco Ippolito) [#&#8203;61383](https://redirect.github.com/nodejs/node/pull/61383)
- \[[`61190bf4b4`](https://redirect.github.com/nodejs/node/commit/61190bf4b4)] - **doc**: packages: example of private import switch to internal (coderaiser) [#&#8203;61343](https://redirect.github.com/nodejs/node/pull/61343)
- \[[`346311c42f`](https://redirect.github.com/nodejs/node/commit/346311c42f)] - **doc**: add esm and cjs examples to node:v8 (Alfredo González) [#&#8203;61328](https://redirect.github.com/nodejs/node/pull/61328)
- \[[`c07c80717c`](https://redirect.github.com/nodejs/node/commit/c07c80717c)] - **doc**: added 'secure' event to tls.TLSSocket (ikeyan) [#&#8203;61066](https://redirect.github.com/nodejs/node/pull/61066)
- \[[`9f68d30f11`](https://redirect.github.com/nodejs/node/commit/9f68d30f11)] - **doc**: restore [@&#8203;watilde](https://redirect.github.com/watilde) to collaborators (Daijiro Wachi) [#&#8203;61350](https://redirect.github.com/nodejs/node/pull/61350)
- \[[`a3b08ddb51`](https://redirect.github.com/nodejs/node/commit/a3b08ddb51)] - **doc**: run license-builder (github-actions\[bot]) [#&#8203;61348](https://redirect.github.com/nodejs/node/pull/61348)
- \[[`4990812dd9`](https://redirect.github.com/nodejs/node/commit/4990812dd9)] - **doc**: document ALPNCallback option for TLSSocket constructor (ikeyan) [#&#8203;61331](https://redirect.github.com/nodejs/node/pull/61331)
- \[[`89e9d19693`](https://redirect.github.com/nodejs/node/commit/89e9d19693)] - **doc**: update MDN links (Livia Medeiros) [#&#8203;61062](https://redirect.github.com/nodejs/node/pull/61062)
- \[[`dcffa88fec`](https://redirect.github.com/nodejs/node/commit/dcffa88fec)] - **doc**: correct description of `error.stack` accessor behavior (René) [#&#8203;61090](https://redirect.github.com/nodejs/node/pull/61090)
- \[[`31476cd4d1`](https://redirect.github.com/nodejs/node/commit/31476cd4d1)] - **doc**: add documentation for process.traceProcessWarnings (Alireza Ebrahimkhani) [#&#8203;53641](https://redirect.github.com/nodejs/node/pull/53641)
- \[[`99c783b9ec`](https://redirect.github.com/nodejs/node/commit/99c783b9ec)] - **doc**: add sqlite session disposal method (René) [#&#8203;61273](https://redirect.github.com/nodejs/node/pull/61273)
- \[[`c7764bed35`](https://redirect.github.com/nodejs/node/commit/c7764bed35)] - **doc**: fix filename typo (Hardanish Singh) [#&#8203;61297](https://redirect.github.com/nodejs/node/pull/61297)
- \[[`0f16bca9d8`](https://redirect.github.com/nodejs/node/commit/0f16bca9d8)] - **doc**: fix typos and grammar in `BUILDING.md` & `onboarding.md` (Hardanish Singh) [#&#8203;61267](https://redirect.github.com/nodejs/node/pull/61267)
- \[[`4b691b562d`](https://redirect.github.com/nodejs/node/commit/4b691b562d)] - **doc**: mention --newVersion release script (Rafael Gonzaga) [#&#8203;61255](https://redirect.github.com/nodejs/node/pull/61255)
- \[[`32e56ab71f`](https://redirect.github.com/nodejs/node/commit/32e56ab71f)] - **doc**: correct typo in api contributing doc (Mike McCready) [#&#8203;61260](https://redirect.github.com/nodejs/node/pull/61260)
- \[[`9ebf1ffbeb`](https://redirect.github.com/nodejs/node/commit/9ebf1ffbeb)] - **doc**: add PR-URL requirement for security backports (Rafael Gonzaga) [#&#8203;61256](https://redirect.github.com/nodejs/node/pull/61256)
- \[[`940f83d95d`](https://redirect.github.com/nodejs/node/commit/940f83d95d)] - **doc**: add reusePort error behavior to net module (mag123c) [#&#8203;61250](https://redirect.github.com/nodejs/node/pull/61250)
- \[[`8881859ee0`](https://redirect.github.com/nodejs/node/commit/8881859ee0)] - **doc**: note corepack package removal in distribution doc (Mike McCready) [#&#8203;61207](https://redirect.github.com/nodejs/node/pull/61207)
- \[[`03a1540cd1`](https://redirect.github.com/nodejs/node/commit/03a1540cd1)] - **doc**: fix tls.connect() timeout documentation (Azad Gupta) [#&#8203;61079](https://redirect.github.com/nodejs/node/pull/61079)
- \[[`816ce7530d`](https://redirect.github.com/nodejs/node/commit/816ce7530d)] - **doc**: missing `passed`, `error` and `passed` properties on `TestContext` (Xavier Stouder) [#&#8203;61185](https://redirect.github.com/nodejs/node/pull/61185)
- \[[`d825c8858a`](https://redirect.github.com/nodejs/node/commit/d825c8858a)] - **doc**: clarify threat model for application-level API exposure (Rafael Gonzaga) [#&#8203;61184](https://redirect.github.com/nodejs/node/pull/61184)
- \[[`a3dd30d0e0`](https://redirect.github.com/nodejs/node/commit/a3dd30d0e0)] - **doc**: correct options for net.Socket class and socket.connect (Xavier Stouder) [#&#8203;61179](https://redirect.github.com/nodejs/node/pull/61179)
- \[[`c3e776becd`](https://redirect.github.com/nodejs/node/commit/c3e776becd)] - **doc**: document error event on readline InterfaceConstructor (Xavier Stouder) [#&#8203;61170](https://redirect.github.com/nodejs/node/pull/61170)
- \[[`05a6372d30`](https://redirect.github.com/nodejs/node/commit/05a6372d30)] - **doc**: add a smooth scrolling effect to the sidebar (btea) [#&#8203;59007](https://redirect.github.com/nodejs/node/pull/59007)
- \[[`76a7eb09ef`](https://redirect.github.com/nodejs/node/commit/76a7eb09ef)] - **doc**: fix test settime docs (Efe) [#&#8203;61117](https://redirect.github.com/nodejs/node/pull/61117)
- \[[`bcbbde6ccc`](https://redirect.github.com/nodejs/node/commit/bcbbde6ccc)] - **doc**: correct invalid collaborator profile (JJ) [#&#8203;61091](https://redirect.github.com/nodejs/node/pull/61091)
- \[[`084741d09d`](https://redirect.github.com/nodejs/node/commit/084741d09d)] - **doc**: add a tip about developer mode on Windows (Joyee Cheung) [#&#8203;61112](https://redirect.github.com/nodejs/node/pull/61112)
- \[[`ed4de371d3`](https://redirect.github.com/nodejs/node/commit/ed4de371d3)] - **doc**: exclude compile-time flag features from security policy (Matteo Collina) [#&#8203;61109](https://redirect.github.com/nodejs/node/pull/61109)
- \[[`3999c2a910`](https://redirect.github.com/nodejs/node/commit/3999c2a910)] - **doc**: add [@&#8203;avivkeller](https://redirect.github.com/avivkeller) to collaborators (Aviv Keller) [#&#8203;61115](https://redirect.github.com/nodejs/node/pull/61115)
- \[[`f3ec066f1a`](https://redirect.github.com/nodejs/node/commit/f3ec066f1a)] - **doc**: warn about short GCM tags visibly (Tobias Nießen) [#&#8203;61082](https://redirect.github.com/nodejs/node/pull/61082)
- \[[`fa542fbae6`](https://redirect.github.com/nodejs/node/commit/fa542fbae6)] - **doc**: add gurgunday to collaborators (Gürgün Dayıoğlu) [#&#8203;61094](https://redirect.github.com/nodejs/node/pull/61094)
- \[[`49f36722dc`](https://redirect.github.com/nodejs/node/commit/49f36722dc)] - **doc**: mark sync module hooks as release candidate (Joyee Cheung) [#&#8203;60960](https://redirect.github.com/nodejs/node/pull/60960)
- \[[`a0adc6afd2`](https://redirect.github.com/nodejs/node/commit/a0adc6afd2)] - **doc**: reorganize docs of module customization hooks (Joyee Cheung) [#&#8203;60960](https://redirect.github.com/nodejs/node/pull/60960)
- \[[`a4097ca048`](https://redirect.github.com/nodejs/node/commit/a4097ca048)] - **doc**: mark crypto.hash as stable (Joyee Cheung) [#&#8203;60994](https://redirect.github.com/nodejs/node/pull/60994)
- \[[`8a67c00bf5`](https://redirect.github.com/nodejs/node/commit/8a67c00bf5)] - **doc**: mark --build-snapshot and --build-snapshot-config as stable (Joyee Cheung) [#&#8203;60954](https://redirect.github.com/nodejs/node/pull/60954)
- \[[`0c83169c35`](https://redirect.github.com/nodejs/node/commit/0c83169c35)] - **doc**: add File modes cross-references in fs methods (Mohit Raj Saxena) [#&#8203;60286](https://redirect.github.com/nodejs/node/pull/60286)
- \[[`dae815262a`](https://redirect.github.com/nodejs/node/commit/dae815262a)] - **doc**: add missing `zstd` to mjs example of zlib (Deokjin Kim) [#&#8203;60915](https://redirect.github.com/nodejs/node/pull/60915)
- \[[`28b284880e`](https://redirect.github.com/nodejs/node/commit/28b284880e)] - **doc**: clarify fileURLToPath security considerations (Rafael Gonzaga) [#&#8203;60887](https://redirect.github.com/nodejs/node/pull/60887)
- \[[`6c440af39b`](https://redirect.github.com/nodejs/node/commit/6c440af39b)] - **doc**: show the use of string expressions in the SQLTagStore example (schliepa) [#&#8203;60873](https://redirect.github.com/nodejs/node/pull/60873)
- \[[`4c5b62209c`](https://redirect.github.com/nodejs/node/commit/4c5b62209c)] - **doc**: replace column with columnNumber in example of `util.getCallSites` (Deokjin Kim) [#&#8203;60881](https://redirect.github.com/nodejs/node/pull/60881)
- \[[`8875c9148e`](https://redirect.github.com/nodejs/node/commit/8875c9148e)] - **doc**: correct spelling in BUILDING.md (Rich Trott) [#&#8203;60875](https://redirect.github.com/nodejs/node/pull/60875)
- \[[`d6cb762426`](https://redirect.github.com/nodejs/node/commit/d6cb762426)] - **doc**: update debuglog examples to use 'foo-bar' instead of 'foo' (xiaoyao) [#&#8203;60867](https://redirect.github.com/nodejs/node/pull/60867)
- \[[`9eae518796`](https://redirect.github.com/nodejs/node/commit/9eae518796)] - **doc**: correct 'event handle' to 'event handler' in Utf8Stream drop event (Riddhi) [#&#8203;60692](https://redirect.github.com/nodejs/node/pull/60692)
- \[[`c3c3ed27c1`](https://redirect.github.com/nodejs/node/commit/c3c3ed27c1)] - **doc**: fix typos in changelogs (Rich Trott) [#&#8203;60855](https://redirect.github.com/nodejs/node/pull/60855)
- \[[`1b975e3017`](https://redirect.github.com/nodejs/node/commit/1b975e3017)] - **doc**: mark module.register as active development (Chengzhong Wu) [#&#8203;60849](https://redirect.github.com/nodejs/node/pull/60849)
- \[[`6a6fc0c851`](https://redirect.github.com/nodejs/node/commit/6a6fc0c851)] - **doc**: add fullName property to SuiteContext (PaulyBearCoding) [#&#8203;60762](https://redirect.github.com/nodejs/node/pull/60762)
- \[[`8347d734e6`](https://redirect.github.com/nodejs/node/commit/8347d734e6)] - **doc**: add additional codemods for deprecation (Augustin Mauroy) [#&#8203;60811](https://redirect.github.com/nodejs/node/pull/60811)
- \[[`7cc87037c3`](https://redirect.github.com/nodejs/node/commit/7cc87037c3)] - **doc**: keep sidebar module visible when navigating docs (Botato) [#&#8203;60410](https://redirect.github.com/nodejs/node/pull/60410)
- \[[`1c6618f643`](https://redirect.github.com/nodejs/node/commit/1c6618f643)] - **doc**: correct concurrency wording in test() documentation (Azad Gupta) [#&#8203;60773](https://redirect.github.com/nodejs/node/pull/60773)
- \[[`488208004e`](https://redirect.github.com/nodejs/node/commit/488208004e)] - **doc**: clarify that CQ only picks up PRs targeting `main` (René) [#&#8203;60731](https://redirect.github.com/nodejs/node/pull/60731)
- \[[`34517940c2`](https://redirect.github.com/nodejs/node/commit/34517940c2)] - **doc**: clarify license section and add contributor note (KaleruMadhu) [#&#8203;60590](https://redirect.github.com/nodejs/node/pull/60590)
- \[[`f080721df4`](https://redirect.github.com/nodejs/node/commit/f080721df4)] - **doc**: correct and expand documentation for SQLTagStore (René) [#&#8203;60200](https://redirect.github.com/nodejs/node/pull/60200)
- \[[`be3d26709d`](https://redirect.github.com/nodejs/node/commit/be3d26709d)] - **doc**: correct tls ALPNProtocols types (René) [#&#8203;60143](https://redirect.github.com/nodejs/node/pull/60143)
- \[[`ef82c53131`](https://redirect.github.com/nodejs/node/commit/ef82c53131)] - **doc**: remove mention of SMS 2FA (Antoine du Hamel) [#&#8203;60707](https://redirect.github.com/nodejs/node/pull/60707)
- \[[`11b190f63e`](https://redirect.github.com/nodejs/node/commit/11b190f63e)] - **doc**: add info about renamed flag in `cli.md` (Antoine du Hamel) [#&#8203;60690](https://redirect.github.com/nodejs/node/pull/60690)
- \[[`59db9bc654`](https://redirect.github.com/nodejs/node/commit/59db9bc654)] - **doc**: fix incorrect slh-dsa oids in crypto.md (Artsiom Malakhau) [#&#8203;60681](https://redirect.github.com/nodejs/node/pull/60681)
- \[[`ad52750cf6`](https://redirect.github.com/nodejs/node/commit/ad52750cf6)] - **doc**: `domain.add()` does not accept timer objects (René) [#&#8203;60675](https://redirect.github.com/nodejs/node/pull/60675)
- \[[`2592d94e29`](https://redirect.github.com/nodejs/node/commit/2592d94e29)] - **doc**: fix v24 changelog after security release (Marco Ippolito) [#&#8203;61371](https://redirect.github.com/nodejs/node/pull/61371)
- \[[`e0f4ad0af0`](https://redirect.github.com/nodejs/node/commit/e0f4ad0af0)] - **doc,test**: add documentation and test on how to use addons in SEA (Joyee Cheung) [#&#8203;59582](https://redirect.github.com/nodejs/node/pull/59582)
- \[[`13af640d94`](https://redirect.github.com/nodejs/node/commit/13af640d94)] - **esm**: ensure watch mode restarts after syntax errors (Xavier Stouder) [#&#8203;61232](https://redirect.github.com/nodejs/node/pull/61232)
- \[[`31afe95d15`](https://redirect.github.com/nodejs/node/commit/31afe95d15)] - **esm**: avoid throw when module specifier is not url (Craig Macomber (Microsoft)) [#&#8203;61000](https://redirect.github.com/nodejs/node/pull/61000)
- \[[`311a04cf2d`](https://redirect.github.com/nodejs/node/commit/311a04cf2d)] - **esm**: improve error messages for ambiguous module syntax (mag123c) [#&#8203;60376](https://redirect.github.com/nodejs/node/pull/60376)
- \[[`cacef92937`](https://redirect.github.com/nodejs/node/commit/cacef92937)] - **events**: remove redundant todo (Gürgün Dayıoğlu) [#&#8203;60595](https://redirect.github.com/nodejs/node/pull/60595)
- \[[`42e1f72561`](https://redirect.github.com/nodejs/node/commit/42e1f72561)] - **events**: remove eventtarget custom inspect branding (Efe) [#&#8203;61128](https://redirect.github.com/nodejs/node/pull/61128)
- \[[`fd8b61369b`](https://redirect.github.com/nodejs/node/commit/fd8b61369b)] - **fs**: remove duplicate getValidatedPath calls (Mert Can Altin) [#&#8203;61359](https://redirect.github.com/nodejs/node/pull/61359)
- \[[`9bb9fc7f2c`](https://redirect.github.com/nodejs/node/commit/9bb9fc7f2c)] - **fs**: fix errorOnExist behavior for directory copy in fs.cp (Nicholas Paun) [#&#8203;60946](https://redirect.github.com/nodejs/node/pull/60946)
- \[[`55a3c70780`](https://redirect.github.com/nodejs/node/commit/55a3c70780)] - **fs**: fix ENOTDIR in globSync when file is treated as dir (sangwook) [#&#8203;61259](https://redirect.github.com/nodejs/node/pull/61259)
- \[[`073a145095`](https://redirect.github.com/nodejs/node/commit/073a145095)] - **fs**: remove duplicate fd validation in sync functions (Mert Can Altin) [#&#8203;61361](https://redirect.github.com/nodejs/node/pull/61361)
- \[[`b132ecdf60`](https://redirect.github.com/nodejs/node/commit/b132ecdf60)] - **fs**: validate statfs path (Efe) [#&#8203;61230](https://redirect.github.com/nodejs/node/pull/61230)
- \[[`0ed0a30f74`](https://redirect.github.com/nodejs/node/commit/0ed0a30f74)] - **fs**: fix rmSync to handle non-ASCII characters (Yeaseen) [#&#8203;61108](https://redirect.github.com/nodejs/node/pull/61108)
- \[[`99632b1a3b`](https://redirect.github.com/nodejs/node/commit/99632b1a3b)] - **fs**: remove broken symlinks in rmSync (sangwook) [#&#8203;61040](https://redirect.github.com/nodejs/node/pull/61040)
- \[[`9cb6757a67`](https://redirect.github.com/nodejs/node/commit/9cb6757a67)] - **fs**: detect dot files when using globstar (Robin van Wijngaarden) [#&#8203;61012](https://redirect.github.com/nodejs/node/pull/61012)
- \[[`e22aad19e0`](https://redirect.github.com/nodejs/node/commit/e22aad19e0)] - **gyp**: aix: change gcc version detection so CXX="ccache g++" works (Stewart X Addison) [#&#8203;61464](https://redirect.github.com/nodejs/node/pull/61464)
- \[[`59d94ba7e7`](https://redirect.github.com/nodejs/node/commit/59d94ba7e7)] - **http**: fix rawHeaders exceeding maxHeadersCount limit (Max Harari) [#&#8203;61285](https://redirect.github.com/nodejs/node/pull/61285)
- \[[`ae6a1fd40a`](https://redirect.github.com/nodejs/node/commit/ae6a1fd40a)] - **http,https**: fix double ERR\_PROXY\_TUNNEL emission (Shima Ryuhei) [#&#8203;60699](https://redirect.github.com/nodejs/node/pull/60699)
- \[[`53bfbaa4b1`](https://redirect.github.com/nodejs/node/commit/53bfbaa4b1)] - **http2**: validate initialWindowSize per HTTP/2 spec (Matteo Collina) [#&#8203;61402](https://redirect.github.com/nodejs/node/pull/61402)
- \[[`14b421b677`](https://redirect.github.com/nodejs/node/commit/14b421b677)] - **http2,zlib**: prefer `call()` over `apply()` if argument list is not array (Livia Medeiros) [#&#8203;60834](https://redirect.github.com/nodejs/node/pull/60834)
- \[[`32b03d0604`](https://redirect.github.com/nodejs/node/commit/32b03d0604)] - **(CVE-2025-59465)** **lib**: add TLSSocket default error handler (RafaelGSS) [nodejs-private/node-private#750](https://redirect.github.com/nodejs-private/node-private/pull/750)
- \[[`4ef7a6c77e`](https://redirect.github.com/nodejs/node/commit/4ef7a6c77e)] - **lib**: backport `_tls_common` and `_tls_wrap` refactors (Dario Piotrowicz) [#&#8203;57643](https://redirect.github.com/nodejs/node/pull/57643)
- \[[`820e0a5cfa`](https://redirect.github.com/nodejs/node/commit/820e0a5cfa)] - **lib**: fix typo in `util.js` comment (Taejin Kim) [#&#8203;61365](https://redirect.github.com/nodejs/node/pull/61365)
- \[[`8de391e1cb`](https://redirect.github.com/nodejs/node/commit/8de391e1cb)] - **lib**: fix TypeScript support check in jitless mode (sangwook) [#&#8203;61382](https://redirect.github.com/nodejs/node/pull/61382)
- \[[`f22f622b3e`](https://redirect.github.com/nodejs/node/commit/f22f622b3e)] - **lib**: add lint rules for reflective function calls (Michaël Zasso) [#&#8203;60825](https://redirect.github.com/nodejs/node/pull/60825)
- \[[`603f0bf8e1`](https://redirect.github.com/nodejs/node/commit/603f0bf8e1)] - **lib**: implement all 1-byte encodings in js (Сковорода Никита Андреевич) [#&#8203;61093](https://redirect.github.com/nodejs/node/pull/61093)
- \[[`1c0a1aa5ef`](https://redirect.github.com/nodejs/node/commit/1c0a1aa5ef)] - **lib**: gbk decoder is gb18030 decoder per spec (Сковорода Никита Андреевич) [#&#8203;61099](https://redirect.github.com/nodejs/node/pull/61099)
- \[[`2cf963df73`](https://redirect.github.com/nodejs/node/commit/2cf963df73)] - **lib**: enforce use of `URLParse` (Antoine du Hamel) [#&#8203;61016](https://redirect.github.com/nodejs/node/pull/61016)
- \[[`bb90630470`](https://redirect.github.com/nodejs/node/commit/bb90630470)] - **lib**: use `FastBuffer` for empty buffer allocation (Gürgün Dayıoğlu) [#&#8203;60558](https://redirect.github.com/nodejs/node/pull/60558)
- \[[`10893a6f13`](https://redirect.github.com/nodejs/node/commit/10893a6f13)] - **lib**: refactor JWK import PQC support check (Filip Skokan) [#&#8203;60586](https://redirect.github.com/nodejs/node/pull/60586)
- \[[`d43806291f`](https://redirect.github.com/nodejs/node/commit/d43806291f)] - **lib,src**: isInsideNodeModules should test on the first non-internal frame (Chengzhong Wu) [#&#8203;60991](https://redirect.github.com/nodejs/node/pull/60991)
- \[[`0bb8f5fe03`](https://redirect.github.com/nodejs/node/commit/0bb8f5fe03)] - **lib,src,test**: fix tests without SQLite (Antoine du Hamel) [#&#8203;60906](https://redirect.github.com/nodejs/node/pull/60906)
- \[[`f3fe0e7fc2`](https://redirect.github.com/nodejs/node/commit/f3fe0e7fc2)] - **lib,test**: enforce use of `assert.fail` via a lint rule (Antoine du Hamel) [#&#8203;61004](https://redirect.github.com/nodejs/node/pull/61004)
- \[[`8b783d46ef`](https://redirect.github.com/nodejs/node/commit/8b783d46ef)] - **meta**: do not fast-track npm updates (Antoine du Hamel) [#&#8203;61475](https://redirect.github.com/nodejs/node/pull/61475)
- \[[`de4a11b50e`](https://redirect.github.com/nodejs/node/commit/de4a11b50e)] - **meta**: fix typos in issue template config (Daijiro Wachi) [#&#8203;61399](https://redirect.github.com/nodejs/node/pull/61399)
- \[[`97b1492783`](https://redirect.github.com/nodejs/node/commit/97b1492783)] - **meta**: label v8 module PRs (René) [#&#8203;61325](https://redirect.github.com/nodejs/node/pull/61325)
- \[[`9bf899b743`](https://redirect.github.com/nodejs/node/commit/9bf899b743)] - **meta**: bump step-security/harden-runner from 2.13.2 to 2.14.0 (dependabot\[bot]) [#&#8203;61245](https://redirect.github.com/nodejs/node/pull/61245)
- \[[`4df7134324`](https://redirect.github.com/nodejs/node/commit/4df7134324)] - **meta**: bump actions/setup-node from 6.0.0 to 6.1.0 (dependabot\[bot]) [#&#8203;61244](https://redirect.github.com/nodejs/node/pull/61244)
- \[[`ff98f610d8`](https://redirect.github.com/nodejs/node/commit/ff98f610d8)] - **meta**: bump actions/cache from 4.3.0 to 5.0.1 (dependabot\[bot]) [#&#8203;61243](https://redirect.github.com/nodejs/node/pull/61243)
- \[[`86950a41ab`](https://redirect.github.com/nodejs/node/commit/86950a41ab)] - **meta**: bump github/codeql-action from 4.31.6 to 4.31.9 (dependabot\[bot]) [#&#8203;61241](https://redirect.github.com/nodejs/node/pull/61241)
- \[[`96901b4828`](https://redirect.github.com/nodejs/node/commit/96901b4828)] - **meta**: bump codecov/codecov-action from 5.5.1 to 5.5.2 (dependabot\[bot]) [#&#8203;61240](https://redirect.github.com/nodejs/node/pull/61240)
- \[[`c90fc7c0d3`](https://redirect.github.com/nodejs/node/commit/c90fc7c0d3)] - **meta**: bump peter-evans/create-pull-request from 7.0.9 to 8.0.0 (dependabot\[bot]) [#&#8203;61237](https://redirect.github.com/nodejs/node/pull/61237)
- \[[`f130d4b6de`](https://redirect.github.com/nodejs/node/commit/f130d4b6de)] - **meta**: move lukekarrys to emeritus (Node.js GitHub Bot) [#&#8203;60985](https://redirect.github.com/nodejs/node/pull/60985)
- \[[`416f34ccfc`](https://redirect.github.com/nodejs/node/commit/416f34ccfc)] - **meta**: bump actions/setup-python from 6.0.0 to 6.1.0 (dependabot\[bot]) [#&#8203;60927](https://redirect.github.com/nodejs/node/pull/60927)
- \[[`2239939305`](https://redirect.github.com/nodejs/node/commit/2239939305)] - **meta**: bump github/codeql-action from 4.31.3 to 4.31.6 (dependabot\[bot]) [#&#8203;60926](https://redirect.github.com/nodejs/node/pull/60926)
- \[[`7f146b6a97`](https://redirect.github.com/nodejs/node/commit/7f146b6a97)] - **meta**: bump peter-evans/create-pull-request from 7.0.8 to 7.0.9 (dependabot\[bot]) [#&#8203;60924](https://redirect.github.com/nodejs/node/pull/60924)
- \[[`d9020f0089`](https://redirect.github.com/nodejs/node/commit/d9020f0089)] - **meta**: bump github/codeql-action from 4.31.2 to 4.31.3 (dependabot\[bot]) [#&#8203;60770](https://redirect.github.com/nodejs/node/pull/60770)
- \[[`4bba259d3b`](https://redirect.github.com/nodejs/node/commit/4bba259d3b)] - **meta**: bump step-security/harden-runner from 2.13.1 to 2.13.2 (dependabot\[bot]) [#&#8203;60769](https://redirect.github.com/nodejs/node/pull/60769)
- \[[`ff11eda2f2`](https://redirect.github.com/nodejs/node/commit/ff11eda2f2)] - **meta**: add Renegade334 to collaborators (Renegade334) [#&#8203;60714](https://redirect.github.com/nodejs/node/pull/60714)
- \[[`e3b5593c0f`](https://redirect.github.com/nodejs/node/commit/e3b5593c0f)] - **module**: fix sync resolve hooks for require with node: prefixes (Joyee Cheung) [#&#8203;61088](https://redirect.github.com/nodejs/node/pull/61088)
- \[[`edec5be805`](https://redirect.github.com/nodejs/node/commit/edec5be805)] - **module**: preserve URL in the parent created by createRequire() (Joyee Cheung) [#&#8203;60974](https://redirect.github.com/nodejs/node/pull/60974)
- \[[`5cc3596eb4`](https://redirect.github.com/nodejs/node/commit/5cc3596eb4)] - **node-api**: fix node\_api\_create\_object\_with\_properties name (Vladimir Morozov) [#&#8203;61319](https://redirect.github.com/nodejs/node/pull/61319)
- \[[`179162fe42`](https://redirect.github.com/nodejs/node/commit/179162fe42)] - **node-api**: use Node-API in comments (Vladimir Morozov) [#&#8203;61320](https://redirect.github.com/nodejs/node/pull/61320)
- \[[`b3fe457a89`](https://redirect.github.com/nodejs/node/commit/b3fe457a89)] - **node-api**: add napi\_set\_prototype (siaeyy) [#&#8203;60711](https://redirect.github.com/nodejs/node/pull/60711)
- \[[`1e13e84f16`](https://redirect.github.com/nodejs/node/commit/1e13e84f16)] - **node-api**: fix data race and use-after-free in napi\_threadsafe\_function (Mika Fischer) [#&#8203;55877](https://redirect.github.com/nodejs/node/pull/55877)
- \[[`36ce6d636d`](https://redirect.github.com/nodejs/node/commit/36ce6d636d)] - **node-api**: add support for Float16Array (Ilyas Shabi) [#&#8203;58879](https://redirect.github.com/nodejs/node/pull/58879)
- \[[`95e6659e2b`](https://redirect.github.com/nodejs/node/commit/95e6659e2b)] - **node-api**: support SharedArrayBuffer in napi\_create\_dataview (Kevin Eady) [#&#8203;60473](https://redirect.github.com/nodejs/node/pull/60473)
- \[[`54f58e2fb2`](https://redirect.github.com/nodejs/node/commit/54f58e2fb2)] - **os**: freeze signals constant (Xavier Stouder) [#&#8203;61038](https://redirect.github.com/nodejs/node/pull/61038)
- \[[`31489310f8`](https://redirect.github.com/nodejs/node/commit/31489310f8)] - **process**: improve process.cwd() error message (TseIan) [#&#8203;61164](https://redirect.github.com/nodejs/node/pull/61164)
- \[[`f7450a90ed`](https://redirect.github.com/nodejs/node/commit/f7450a90ed)] - **repl**: move completion logic to internal module (Dario Piotrowicz) [#&#8203;59889](https://redirect.github.com/nodejs/node/pull/59889)
- \[[`27117625df`](https://redirect.github.com/nodejs/node/commit/27117625df)] - **sqlite**: add some tests (Guilherme Araújo) [#&#8203;61410](https://redirect.github.com/nodejs/node/pull/61410)
- \[[`d56066ce8c`](https://redirect.github.com/nodejs/node/commit/d56066ce8c)] - **sqlite**: improve error messages for tag store (Pramit Sharma) [#&#8203;61096](https://redirect.github.com/nodejs/node/pull/61096)
- \[[`9d993be6c1`](https://redirect.github.com/nodejs/node/commit/9d993be6c1)] - **sqlite**: make `SQLTagStore.prototype.size` a getter (René) [#&#8203;60246](https://redirect.github.com/nodejs/node/pull/60246)
- \[[`ceaa200d16`](https://redirect.github.com/nodejs/node/commit/ceaa200d16)] - **src**: improve StringBytes::Encode perf on UTF8 (Сковорода Никита Андреевич) [#&#8203;61131](https://redirect.github.com/nodejs/node/pull/61131)
- \[[`034a5f2346`](https://redirect.github.com/nodejs/node/commit/034a5f2346)] - **src**: add missing override specifier to Clean() (Tobias Nießen) [#&#8203;61429](https://redirect.github.com/nodejs/node/pull/61429)
- \[[`977f46cc20`](https://redirect.github.com/nodejs/node/commit/977f46cc20)] - **src**: cache context lookup in vectored io loops (Mert Can Altin) [#&#8203;61387](https://redirect.github.com/nodejs/node/pull/61387)
- \[[`bb9e4e0784`](https://redirect.github.com/nodejs/node/commit/bb9e4e0784)] - **src**: cache missing package.json files in the C++ package config cache (Michael Smith) [#&#8203;60425](https://redirect.github.com/nodejs/node/pull/60425)
- \[[`c1aa9f49cd`](https://redirect.github.com/nodejs/node/commit/c1aa9f49cd)] - **src**: use starts\_with instead of rfind/find (Tobias Nießen) [#&#8203;61426](https://redirect.github.com/nodejs/node/pull/61426)
- \[[`d3676d0a82`](https://redirect.github.com/nodejs/node/commit/d3676d0a82)] - **src**: use C++ nullptr in sqlite (Tobias Nießen) [#&#8203;61416](https://redirect.github.com/nodejs/node/pull/61416)
- \[[`001be8aa7c`](https://redirect.github.com/nodejs/node/commit/001be8aa7c)] - **src**: use C++ nullptr in webstorage (Tobias Nießen) [#&#8203;61407](https://redirect.github.com/nodejs/node/pull/61407)
- \[[`4f832b1e3d`](https://redirect.github.com/nodejs/node/commit/4f832b1e3d)] - **src**: fix pointer alignment (jhofstee) [#&#8203;61336](https://redirect.github.com/nodejs/node/pull/61336)
- \[[`a0a8c96fd1`](https://redirect.github.com/nodejs/node/commit/a0a8c96fd1)] - **src**: dump snapshot source with node:generate\_default\_snapshot\_source (Joyee Cheung) [#&#8203;61101](https://redirect.github.com/nodejs/node/pull/61101)
- \[[`b6d3caeda8`](https://redirect.github.com/nodejs/node/commit/b6d3caeda8)] - **src**: improve StringBytes::Encode perf on ASCII (Сковорода Никита Андреевич) [#&#8203;61119](https://redirect.github.com/nodejs/node/pull/61119)
- \[[`9c80e5ac87`](https://redirect.github.com/nodejs/node/commit/9c80e5ac87)] - **src**: add HandleScope to edge loop in heap\_utils (Mert Can Altin) [#&#8203;60885](https://redirect.github.com/nodejs/node/pull/60885)
- \[[`09ccd94312`](https://redirect.github.com/nodejs/node/commit/09ccd94312)] - **src**: remove redundant CHECK (Tobias Nießen) [#&#8203;61130](https://redirect.github.com/nodejs/node/pull/61130)
- \[[`6008354b8a`](https://redirect.github.com/nodejs/node/commit/6008354b8a)] - **src**: remove unused private field in `SQLTagStore` (Michaël Zasso) [#&#8203;61027](https://redirect.github.com/nodejs/node/pull/61027)
- \[[`7484a34a7d`](https://redirect.github.com/nodejs/node/commit/7484a34a7d)] - **src**: implement Windows-1252 encoding support and update related tests (Mert Can Altin) [#&#8203;60893](https://redirect.github.com/nodejs/node/pull/60893)
- \[[`47851db855`](https://redirect.github.com/nodejs/node/commit/47851db855)] - **src**: fix off-thread cert loading in bundled cert mode (Joyee Cheung) [#&#8203;60764](https://redirect.github.com/nodejs/node/pull/60764)
- \[[`4702a8696b`](https://redirect.github.com/nodejs/node/commit/4702a8696b)] - **src**: handle DER decoding errors from system certificates (Joyee Cheung) [#&#8203;60787](https://redirect.github.com/nodejs/node/pull/60787)
- \[[`19a4926965`](https://redirect.github.com/nodejs/node/commit/19a4926965)] - **src**: use static\_cast instead of C-style cast (Michaël Zasso) [#&#8203;60868](https://redirect.github.com/nodejs/node/pull/60868)
- \[[`6529334dec`](https://redirect.github.com/nodejs/node/commit/6529334dec)] - **src**: add test flag to config file (Marco Ippolito) [#&#8203;60798](https://redirect.github.com/nodejs/node/pull/60798)
- \[[`d153b30773`](https://redirect.github.com/nodejs/node/commit/d153b30773)] - **src**: split inspector protocol domains files (Chengzhong Wu) [#&#8203;60754](https://redirect.github.com/nodejs/node/pull/60754)
- \[[`7191b847c6`](https://redirect.github.com/nodejs/node/commit/7191b847c6)] - **src,permission**: fix permission.has on empty param (Rafael Gonzaga) [#&#8203;60674](https://redirect.github.com/nodejs/node/pull/60674)
- \[[`a188b954bb`](https://redirect.github.com/nodejs/node/commit/a188b954bb)] - **src,permission**: add debug log on is\_tree\_granted (Rafael Gonzaga) [#&#8203;60668](https://redirect.github.com/nodejs/node/pull/60668)
- \[[`b483b5a8ea`](https://redirect.github.com/nodejs/node/commit/b483b5a8ea)] - **stream**: export namespace object from internal end-of-stream module (René) [#&#8203;61455](https://redirect.github.com/nodejs/node/pull/61455)
- \[[`0472104536`](https://redirect.github.com/nodejs/node/commit/0472104536)] - **stream**: fix isErrored/isWritable for WritableStreams (René) [#&#8203;60905](https://redirect.github.com/nodejs/node/pull/60905)
- \[[`dd13f1046f`](https://redirect.github.com/nodejs/node/commit/dd13f1046f)] - **test**: skip --build-sea tests on platforms where SEA is flaky (Joyee Cheung) [#&#8203;61504](https://redirect.github.com/nodejs/node/pull/61504)
- \[[`6c18bf26f4`](https://redirect.github.com/nodejs/node/commit/6c18bf26f4)] - **test**: update WPT for url to [`81a2aed`](https://redirect.github.com/nodejs/node/commit/81a2aed262) (Node.js GitHub Bot) [#&#8203;61509](https://redirect.github.com/nodejs/node/pull/61509)
- \[[`f511c24d6b`](https://redirect.github.com/nodejs/node/commit/f511c24d6b)] - **test**: fix flaky debugger test (Ryuhei Shima) [#&#8203;58324](https://redirect.github.com/nodejs/node/pull/58324)
- \[[`41710ba953`](https://redirect.github.com/nodejs/node/commit/41710ba953)] - **test**: ensure removeListener event fires for once() listeners (sangwook) [#&#8203;60137](https://redirect.github.com/nodejs/node/pull/60137)
- \[[`0035f3fa0f`](https://redirect.github.com/nodejs/node/commit/0035f3fa0f)] - **test**: delay writing the files only on macOS (Luigi Pinca) [#&#8203;61532](https://redirect.github.com/nodejs/node/pull/61532)
- \[[`99c29eb261`](https://redirect.github.com/nodejs/node/commit/99c29eb261)] - **test**: add implicit test for fs dispose handling with using (Ilyas Shabi) [#&#8203;61140](https://redirect.github.com/nodejs/node/pull/61140)
- \[[`e349d34c8a`](https://redirect.github.com/nodejs/node/commit/e349d34c8a)] - **test**: check new WebCryptoAPI enum values (Filip Skokan) [#&#8203;61406](https://redirect.github.com/nodejs/node/pull/61406)
- \[[`e75617d25f`](https://redirect.github.com/nodejs/node/commit/e75617d25f)] - **test**: split test-esm-loader-hooks (Joyee Cheung) [#&#8203;61374](https://redirect.github.com/nodejs/node/pull/61374)
- \[[`42110af62a`](https://redirect.github.com/nodejs/node/commit/42110af62a)] - **test**: aix: mark test-emit-on-destroyed as flaky (Stewart X Addison) [#&#8203;61381](https://redirect.github.com/nodejs/node/pull/61381)
- \[[`180fdbf188`](https://redirect.github.com/nodejs/node/commit/180fdbf188)] - **test**: update url web-platform tests (Yagiz Nizipli) [#&#8203;61315](https://redirect.github.com/nodejs/node/pull/61315)
- \[[`4bac4ecd9d`](https://redirect.github.com/nodejs/node/commit/4bac4ecd9d)] - **test**: ensure assertions are reached on more tests (Antoine du Hamel) [#&#8203;60761](https://redirect.github.com/nodejs/node/pull/60761)
- \[[`39ca74e57e`](https://redirect.github.com/nodejs/node/commit/39ca74e57e)] - **test**: ensure assertions are reached on more tests (Antoine du Hamel) [#&#8203;60759](https://redirect.github.com/nodejs/node/pull/60759)
- \[[`7327b04875`](https://redirect.github.com/nodejs/node/commit/7327b04875)] - **test**: ensure assertions are reached on more tests (Antoine du Hamel) [#&#8203;60726](https://redirect.github.com/nodejs/node/pull/60726)
- \[[`fd6601c710`](https://redirect.github.com/nodejs/node/commit/fd6601c710)] - **test**: asserts that import.meta.resolve invokes sync loader hooks (Chengzhong Wu) [#&#8203;61158](https://redirect.github.com/nodejs/node/pull/61158)
- \[[`da4d4d4fde`](https://redirect.github.com/nodejs/node/commit/da4d4d4fde)] - **test**: check util.parseArgs argv parsing with actual process execution (René) [#&#8203;61089](https://redirect.github.com/nodejs/node/pull/61089)
- \[[`368b32d410`](https://redirect.github.com/nodejs/node/commit/368b32d410)] - **test**: update WPT for urlpattern to [`a2e15ad`](https://redirect.github.com/nodejs/node/commit/a2e15ad405) (Node.js GitHub Bot) [#&#8203;61134](https://redirect.github.com/nodejs/node/pull/61134)
- \[[`e880062de8`](https://redirect.github.com/nodejs/node/commit/e880062de8)] - **test**: make buffer sizes 32bit-aware in test-internal-util-construct-sab (René) [#&#8203;61026](https://redirect.github.com/nodejs/node/pull/61026)
- \[[`f2706e1166`](https://redirect.github.com/nodejs/node/commit/f2706e1166)] - **test**: remove unneccessary repl magic\_mode tests (Dario Piotrowicz) [#&#8203;61053](https://redirect.github.com/nodejs/node/pull/61053)
- \[[`327dd25f86`](https://redirect.github.com/nodejs/node/commit/327dd25f86)] - **test**: skip sea tests on riscv64 (Stewart X Addison) [#&#8203;61111](https://redirect.github.com/nodejs/node/pull/61111)
- \[[`6da34027e2`](https://redirect.github.com/nodejs/node/commit/6da34027e2)] - **test**: simplify `test-cli-node-options-docs` (Antoine du Hamel) [#&#8203;61006](https://redirect.github.com/nodejs/node/pull/61006)
- \[[`74df70d1da`](https://redirect.github.com/nodejs/node/commit/74df70d1da)] - **test**: mark stringbytes-external-max flaky on AIX (Stewart X Addison) [#&#8203;60995](https://redirect.github.com/nodejs/node/pull/60995)
- \[[`5513338446`](https://redirect.github.com/nodejs/node/commit/5513338446)] - **test**: update test426 fixtures (Rich Trott) [#&#8203;60982](https://redirect.github.com/nodejs/node/pull/60982)
- \[[`9f594f53a7`](https://redirect.github.com/nodejs/node/commit/9f594f53a7)] - **test**: update WPT for urlpattern to [`aed1f3d`](https://redirect.github.com/nodejs/node/commit/aed1f3d244) (Node.js GitHub Bot) [#&#8203;60642](https://redirect.github.com/nodejs/node/pull/60642)
- \[[`18e3b91bf1`](https://redirect.github.com/nodejs/node/commit/18e3b91bf1)] - **test**: deflake `test-repl-paste-big-data` (Livia Medeiros) [#&#8203;60975]

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

### Review by @swadeley - Approved on February 16, 2026 at 05:06 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1392*
