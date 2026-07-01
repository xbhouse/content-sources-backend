---
type: pull_request
number: 1338
title: "Build: chore(deps): update node.js to v24.12.0"
state: merged
author: red-hat-konflux
created: 2025-12-15T04:38:24Z
updated: 2025-12-17T12:21:59Z
closed: 2025-12-17T12:21:57Z
merged: 2025-12-17T12:21:57Z
base_branch: main
head_branch: konflux/mintmaker/main/node-24.x
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1338
---

# Pull Request #1338: Build: chore(deps): update node.js to v24.12.0

**Author**: @red-hat-konflux
**Created**: December 15, 2025 at 04:38 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/node-24.x`

## Description

This PR contains the following updates:

| Package | Update | Change |
|---|---|---|
| [node](https://nodejs.org) ([source](https://redirect.github.com/nodejs/node)) | minor | `24.11.1` -> `24.12.0` |

---

### Release Notes

<details>
<summary>nodejs/node (node)</summary>

### [`v24.12.0`](https://redirect.github.com/nodejs/node/releases/tag/v24.12.0): 2025-12-10, Version 24.12.0 &#x27;Krypton&#x27; (LTS), @&#8203;targos

[Compare Source](https://redirect.github.com/nodejs/node/compare/v24.11.1...v24.12.0)

##### Notable Changes

- \[[`1a00b5f68a`](https://redirect.github.com/nodejs/node/commit/1a00b5f68a)] - **(SEMVER-MINOR)** **http**: add optimizeEmptyRequests server option (Rafael Gonzaga) [#&#8203;59778](https://redirect.github.com/nodejs/node/pull/59778)
- \[[`ff5754077d`](https://redirect.github.com/nodejs/node/commit/ff5754077d)] - **(SEMVER-MINOR)** **lib**: add options to util.deprecate (Rafael Gonzaga) [#&#8203;59982](https://redirect.github.com/nodejs/node/pull/59982)
- \[[`8987159234`](https://redirect.github.com/nodejs/node/commit/8987159234)] - **(SEMVER-MINOR)** **module**: mark type stripping as stable (Marco Ippolito) [#&#8203;60600](https://redirect.github.com/nodejs/node/pull/60600)
- \[[`92c484ebf4`](https://redirect.github.com/nodejs/node/commit/92c484ebf4)] - **(SEMVER-MINOR)** **node-api**: add napi\_create\_object\_with\_properties (Miguel Marcondes Filho) [#&#8203;59953](https://redirect.github.com/nodejs/node/pull/59953)
- \[[`b11bc5984e`](https://redirect.github.com/nodejs/node/commit/b11bc5984e)] - **(SEMVER-MINOR)** **sqlite**: allow setting defensive flag (Bart Louwers) [#&#8203;60217](https://redirect.github.com/nodejs/node/pull/60217)
- \[[`e7da5b4b7d`](https://redirect.github.com/nodejs/node/commit/e7da5b4b7d)] - **(SEMVER-MINOR)** **src**: add watch config namespace (Marco Ippolito) [#&#8203;60178](https://redirect.github.com/nodejs/node/pull/60178)
- \[[`a7f7d10c06`](https://redirect.github.com/nodejs/node/commit/a7f7d10c06)] - **(SEMVER-MINOR)** **src**: add an option to make compile cache portable (Aditi) [#&#8203;58797](https://redirect.github.com/nodejs/node/pull/58797)
- \[[`92ea669240`](https://redirect.github.com/nodejs/node/commit/92ea669240)] - **(SEMVER-MINOR)** **src,permission**: add --allow-inspector ability (Rafael Gonzaga) [#&#8203;59711](https://redirect.github.com/nodejs/node/pull/59711)
- \[[`05d7509bd2`](https://redirect.github.com/nodejs/node/commit/05d7509bd2)] - **(SEMVER-MINOR)** **v8**: add cpu profile (theanarkh) [#&#8203;59807](https://redirect.github.com/nodejs/node/pull/59807)

##### Commits

- \[[`e4a23a35ac`](https://redirect.github.com/nodejs/node/commit/e4a23a35ac)] - **benchmark**: focus on import.meta intialization in import-meta benchmark (Joyee Cheung) [#&#8203;60603](https://redirect.github.com/nodejs/node/pull/60603)
- \[[`b6114ae5c9`](https://redirect.github.com/nodejs/node/commit/b6114ae5c9)] - **benchmark**: add per-suite setup option (Joyee Cheung) [#&#8203;60574](https://redirect.github.com/nodejs/node/pull/60574)
- \[[`ac8e90af7c`](https://redirect.github.com/nodejs/node/commit/ac8e90af7c)] - **buffer**: speed up concat via TypedArray#set (Gürgün Dayıoğlu) [#&#8203;60399](https://redirect.github.com/nodejs/node/pull/60399)
- \[[`acbc8ca13e`](https://redirect.github.com/nodejs/node/commit/acbc8ca13e)] - **build**: upgrade Python linter ruff, add rules ASYNC,PERF (Christian Clauss) [#&#8203;59984](https://redirect.github.com/nodejs/node/pull/59984)
- \[[`f97a609a07`](https://redirect.github.com/nodejs/node/commit/f97a609a07)] - **console**: optimize single-string logging (Gürgün Dayıoğlu) [#&#8203;60422](https://redirect.github.com/nodejs/node/pull/60422)
- \[[`6cd9bdc580`](https://redirect.github.com/nodejs/node/commit/6cd9bdc580)] - **crypto**: ensure documented RSA-PSS saltLength default is used (Filip Skokan) [#&#8203;60662](https://redirect.github.com/nodejs/node/pull/60662)
- \[[`0fafe24d9b`](https://redirect.github.com/nodejs/node/commit/0fafe24d9b)] - **crypto**: fix argument validation in crypto.timingSafeEqual fast path (Joyee Cheung) [#&#8203;60538](https://redirect.github.com/nodejs/node/pull/60538)
- \[[`54421e0419`](https://redirect.github.com/nodejs/node/commit/54421e0419)] - **debugger**: fix event listener leak in the run command (Joyee Cheung) [#&#8203;60464](https://redirect.github.com/nodejs/node/pull/60464)
- \[[`c361a628b4`](https://redirect.github.com/nodejs/node/commit/c361a628b4)] - **deps**: V8: cherry-pick [`72b0e27`](https://redirect.github.com/nodejs/node/commit/72b0e27bd936) (pthier) [#&#8203;60732](https://redirect.github.com/nodejs/node/pull/60732)
- \[[`c70f4588dd`](https://redirect.github.com/nodejs/node/commit/c70f4588dd)] - **deps**: V8: cherry-pick [`6bb32bd`](https://redirect.github.com/nodejs/node/commit/6bb32bd2c194) (Erik Corry) [#&#8203;60732](https://redirect.github.com/nodejs/node/pull/60732)
- \[[`881fe784c5`](https://redirect.github.com/nodejs/node/commit/881fe784c5)] - **deps**: V8: cherry-pick [`0dd2318`](https://redirect.github.com/nodejs/node/commit/0dd2318b5237) (Erik Corry) [#&#8203;60732](https://redirect.github.com/nodejs/node/pull/60732)
- \[[`457c33efcc`](https://redirect.github.com/nodejs/node/commit/457c33efcc)] - **deps**: V8: cherry-pick [`df20105`](https://redirect.github.com/nodejs/node/commit/df20105ccf36) (Erik Corry) [#&#8203;60732](https://redirect.github.com/nodejs/node/pull/60732)
- \[[`0bf45a829c`](https://redirect.github.com/nodejs/node/commit/0bf45a829c)] - **deps**: V8: backport [`e5dbbba`](https://redirect.github.com/nodejs/node/commit/e5dbbbadcbff) (Darshan Sen) [#&#8203;60524](https://redirect.github.com/nodejs/node/pull/60524)
- \[[`4993bdc476`](https://redirect.github.com/nodejs/node/commit/4993bdc476)] - **deps**: V8: cherry-pick [`5ba9200`](https://redirect.github.com/nodejs/node/commit/5ba9200cd046) (Juan José Arboleda) [#&#8203;60620](https://redirect.github.com/nodejs/node/pull/60620)
- \[[`1e9abe0078`](https://redirect.github.com/nodejs/node/commit/1e9abe0078)] - **deps**: update corepack to 0.34.5 (Node.js GitHub Bot) [#&#8203;60842](https://redirect.github.com/nodejs/node/pull/60842)
- \[[`3f704ed08f`](https://redirect.github.com/nodejs/node/commit/3f704ed08f)] - **deps**: update corepack to 0.34.4 (Node.js GitHub Bot) [#&#8203;60643](https://redirect.github.com/nodejs/node/pull/60643)
- \[[`04e360fdb1`](https://redirect.github.com/nodejs/node/commit/04e360fdb1)] - **deps**: V8: cherry-pick [`06bf293`](https://redirect.github.com/nodejs/node/commit/06bf293610ef), [`146962d`](https://redirect.github.com/nodejs/node/commit/146962dda8d2) and [`e0fb10b`](https://redirect.github.com/nodejs/node/commit/e0fb10b5148c) (Michaël Zasso) [#&#8203;60713](https://redirect.github.com/nodejs/node/pull/60713)
- \[[`fcbd8dbbde`](https://redirect.github.com/nodejs/node/commit/fcbd8dbbde)] - **deps**: patch V8 to 13.6.233.17 (Michaël Zasso) [#&#8203;60712](https://redirect.github.com/nodejs/node/pull/60712)
- \[[`28e9433f39`](https://redirect.github.com/nodejs/node/commit/28e9433f39)] - **deps**: V8: cherry-pick [`8735658`](https://redirect.github.com/nodejs/node/commit/87356585659b) (Joyee Cheung) [#&#8203;60069](https://redirect.github.com/nodejs/node/pull/60069)
- \[[`3cac85b243`](https://redirect.github.com/nodejs/node/commit/3cac85b243)] - **deps**: V8: backport [`2e4c5cf`](https://redirect.github.com/nodejs/node/commit/2e4c5cf9b112) (Michaël Zasso) [#&#8203;60654](https://redirect.github.com/nodejs/node/pull/60654)
- \[[`1daece1970`](https://redirect.github.com/nodejs/node/commit/1daece1970)] - **deps**: call OPENSSL\_free after ANS1\_STRING\_to\_UTF8 (Rafael Gonzaga) [#&#8203;60609](https://redirect.github.com/nodejs/node/pull/60609)
- \[[`5f55a9c9ea`](https://redirect.github.com/nodejs/node/commit/5f55a9c9ea)] - **deps**: nghttp2: revert [`7784fa9`](https://redirect.github.com/nodejs/node/commit/7784fa979d0b) (Antoine du Hamel) [#&#8203;59790](https://redirect.github.com/nodejs/node/pull/59790)
- \[[`1d9e7c1f4d`](https://redirect.github.com/nodejs/node/commit/1d9e7c1f4d)] - **deps**: update nghttp2 to 1.67.1 (nodejs-github-bot) [#&#8203;59790](https://redirect.github.com/nodejs/node/pull/59790)
- \[[`3140415068`](https://redirect.github.com/nodejs/node/commit/3140415068)] - **deps**: update simdjson to 4.1.0 (Node.js GitHub Bot) [#&#8203;60542](https://redirect.github.com/nodejs/node/pull/60542)
- \[[`d911f9f1b8`](https://redirect.github.com/nodejs/node/commit/d911f9f1b8)] - **deps**: update amaro to 1.1.5 (Node.js GitHub Bot) [#&#8203;60541](https://redirect.github.com/nodejs/node/pull/60541)
- \[[`daaaf04a32`](https://redirect.github.com/nodejs/node/commit/daaaf04a32)] - **deps**: V8: cherry-pick [`2abc613`](https://redirect.github.com/nodejs/node/commit/2abc61361dd4) (Richard Lau) [#&#8203;60177](https://redirect.github.com/nodejs/node/pull/60177)
- \[[`b4f63ee5f8`](https://redirect.github.com/nodejs/node/commit/b4f63ee5f8)] - **doc**: update Collaborators list to reflect hybrist handle change (Antoine du Hamel) [#&#8203;60650](https://redirect.github.com/nodejs/node/pull/60650)
- \[[`effcf7a8ab`](https://redirect.github.com/nodejs/node/commit/effcf7a8ab)] - **doc**: fix link in `--env-file=file` section (N. Bighetti) [#&#8203;60563](https://redirect.github.com/nodejs/node/pull/60563)
- \[[`7011736703`](https://redirect.github.com/nodejs/node/commit/7011736703)] - **doc**: fix linter issues (Antoine du Hamel) [#&#8203;60636](https://redirect.github.com/nodejs/node/pull/60636)
- \[[`5cc79d8945`](https://redirect.github.com/nodejs/node/commit/5cc79d8945)] - **doc**: add missing history entry for `sqlite.md` (Antoine du Hamel) [#&#8203;60607](https://redirect.github.com/nodejs/node/pull/60607)
- \[[`bbc649057c`](https://redirect.github.com/nodejs/node/commit/bbc649057c)] - **doc**: correct values/references for buffer.kMaxLength (René) [#&#8203;60305](https://redirect.github.com/nodejs/node/pull/60305)
- \[[`ea7ecb517b`](https://redirect.github.com/nodejs/node/commit/ea7ecb517b)] - **doc**: recommend events.once to manage 'close' event (Dan Fabulich) [#&#8203;60017](https://redirect.github.com/nodejs/node/pull/60017)
- \[[`58bff04cc2`](https://redirect.github.com/nodejs/node/commit/58bff04cc2)] - **doc**: highlight module loading difference between import and require (Ajay A) [#&#8203;59815](https://redirect.github.com/nodejs/node/pull/59815)
- \[[`bbcbff9b4d`](https://redirect.github.com/nodejs/node/commit/bbcbff9b4d)] - **doc**: add CJS code snippets in `sqlite.md` (Allon Murienik) [#&#8203;60395](https://redirect.github.com/nodejs/node/pull/60395)
- \[[`f8af33d5a7`](https://redirect.github.com/nodejs/node/commit/f8af33d5a7)] - **doc**: fix typo in `process.unref` documentation (우혁) [#&#8203;59698](https://redirect.github.com/nodejs/node/pull/59698)
- \[[`df105dc351`](https://redirect.github.com/nodejs/node/commit/df105dc351)] - **doc**: add some entries to `glossary.md` (Mohataseem Khan) [#&#8203;59277](https://redirect.github.com/nodejs/node/pull/59277)
- \[[`4955cb2b5b`](https://redirect.github.com/nodejs/node/commit/4955cb2b5b)] - **doc**: improve agent.createConnection docs for http and https agents (JaeHo Jang) [#&#8203;58205](https://redirect.github.com/nodejs/node/pull/58205)
- \[[`6283bb5cc9`](https://redirect.github.com/nodejs/node/commit/6283bb5cc9)] - **doc**: fix pseudo code in modules.md (chirsz) [#&#8203;57677](https://redirect.github.com/nodejs/node/pull/57677)
- \[[`d5059ea537`](https://redirect.github.com/nodejs/node/commit/d5059ea537)] - **doc**: add missing variable in code snippet (Koushil Mankali) [#&#8203;55478](https://redirect.github.com/nodejs/node/pull/55478)
- \[[`900de373ae`](https://redirect.github.com/nodejs/node/commit/900de373ae)] - **doc**: add missing word in `single-executable-applications.md` (Konstantin Tsabolov) [#&#8203;53864](https://redirect.github.com/nodejs/node/pull/53864)
- \[[`5735044c8b`](https://redirect.github.com/nodejs/node/commit/5735044c8b)] - **doc**: fix typo in http.md (Michael Solomon) [#&#8203;59354](https://redirect.github.com/nodejs/node/pull/59354)
- \[[`2dee6df831`](https://redirect.github.com/nodejs/node/commit/2dee6df831)] - **doc**: update devcontainer.json and add documentation (Joyee Cheung) [#&#8203;60472](https://redirect.github.com/nodejs/node/pull/60472)
- \[[`8f2d98d7d2`](https://redirect.github.com/nodejs/node/commit/8f2d98d7d2)] - **doc**: add haramj as triager (Haram Jeong) [#&#8203;60348](https://redirect.github.com/nodejs/node/pull/60348)
- \[[`bbd7fdfff4`](https://redirect.github.com/nodejs/node/commit/bbd7fdfff4)] - **doc**: clarify require(esm) description (dynst) [#&#8203;60520](https://redirect.github.com/nodejs/node/pull/60520)
- \[[`33ad11a764`](https://redirect.github.com/nodejs/node/commit/33ad11a764)] - **doc**: instantiate resolver object (Donghoon Nam) [#&#8203;60476](https://redirect.github.com/nodejs/node/pull/60476)
- \[[`81a61274f3`](https://redirect.github.com/nodejs/node/commit/81a61274f3)] - **doc**: correct module loading descriptions (Joyee Cheung) [#&#8203;60346](https://redirect.github.com/nodejs/node/pull/60346)
- \[[`77911185fe`](https://redirect.github.com/nodejs/node/commit/77911185fe)] - **doc**: clarify --use-system-ca support status (Joyee Cheung) [#&#8203;60340](https://redirect.github.com/nodejs/node/pull/60340)
- \[[`185f6e95d9`](https://redirect.github.com/nodejs/node/commit/185f6e95d9)] - **doc,crypto**: link keygen to supported types (Filip Skokan) [#&#8203;60585](https://redirect.github.com/nodejs/node/pull/60585)
- \[[`772d6c6608`](https://redirect.github.com/nodejs/node/commit/772d6c6608)] - **doc,src,lib**: clarify experimental status of Web Storage support (Antoine du Hamel) [#&#8203;60708](https://redirect.github.com/nodejs/node/pull/60708)
- \[[`ad98e11ac2`](https://redirect.github.com/nodejs/node/commit/ad98e11ac2)] - **esm**: use sync loading/resolving on non-loader-hook thread (Joyee Cheung) [#&#8203;60380](https://redirect.github.com/nodejs/node/pull/60380)
- \[[`1a00b5f68a`](https://redirect.github.com/nodejs/node/commit/1a00b5f68a)] - **(SEMVER-MINOR)** **http**: add optimizeEmptyRequests server option (Rafael Gonzaga) [#&#8203;59778](https://redirect.github.com/nodejs/node/pull/59778)
- \[[`5703ce68bc`](https://redirect.github.com/nodejs/node/commit/5703ce68bc)] - **http**: replace startsWith with strict equality (btea) [#&#8203;59394](https://redirect.github.com/nodejs/node/pull/59394)
- \[[`2b696ffad8`](https://redirect.github.com/nodejs/node/commit/2b696ffad8)] - **http2**: add diagnostics channels for client stream request body (Darshan Sen) [#&#8203;60480](https://redirect.github.com/nodejs/node/pull/60480)
- \[[`dbdf4cb5a5`](https://redirect.github.com/nodejs/node/commit/dbdf4cb5a5)] - **inspector**: inspect HTTP response body (Chengzhong Wu) [#&#8203;60572](https://redirect.github.com/nodejs/node/pull/60572)
- \[[`9dc9a7d33d`](https://redirect.github.com/nodejs/node/commit/9dc9a7d33d)] - **inspector**: support inspecting HTTP/2 request and response bodies (Darshan Sen) [#&#8203;60483](https://redirect.github.com/nodejs/node/pull/60483)
- \[[`89fa2befe4`](https://redirect.github.com/nodejs/node/commit/89fa2befe4)] - **inspector**: fix crash when receiving non json message (Shima Ryuhei) [#&#8203;60388](https://redirect.github.com/nodejs/node/pull/60388)
- \[[`ff5754077d`](https://redirect.github.com/nodejs/node/commit/ff5754077d)] - **(SEMVER-MINOR)** **lib**: add options to util.deprecate (Rafael Gonzaga) [#&#8203;59982](https://redirect.github.com/nodejs/node/pull/59982)
- \[[`33baaf42c8`](https://redirect.github.com/nodejs/node/commit/33baaf42c8)] - **lib**: replace global SharedArrayBuffer constructor with bound method (Renegade334) [#&#8203;60497](https://redirect.github.com/nodejs/node/pull/60497)
- \[[`b047586a08`](https://redirect.github.com/nodejs/node/commit/b047586a08)] - **meta**: bump actions/download-artifact from 5.0.0 to 6.0.0 (dependabot\[bot]) [#&#8203;60532](https://redirect.github.com/nodejs/node/pull/60532)
- \[[`64192176d7`](https://redirect.github.com/nodejs/node/commit/64192176d7)] - **meta**: bump actions/upload-artifact from 4.6.2 to 5.0.0 (dependabot\[bot]) [#&#8203;60531](https://redirect.github.com/nodejs/node/pull/60531)
- \[[`af6d4a6b9b`](https://redirect.github.com/nodejs/node/commit/af6d4a6b9b)] - **meta**: bump github/codeql-action from 3.30.5 to 4.31.2 (dependabot\[bot]) [#&#8203;60533](https://redirect.github.com/nodejs/node/pull/60533)
- \[[`c17276fd24`](https://redirect.github.com/nodejs/node/commit/c17276fd24)] - **meta**: bump actions/setup-node from 5.0.0 to 6.0.0 (dependabot\[bot]) [#&#8203;60529](https://redirect.github.com/nodejs/node/pull/60529)
- \[[`6e8b52a7dc`](https://redirect.github.com/nodejs/node/commit/6e8b52a7dc)] - **meta**: bump actions/stale from 10.0.0 to 10.1.0 (dependabot\[bot]) [#&#8203;60528](https://redirect.github.com/nodejs/node/pull/60528)
- \[[`a12658595b`](https://redirect.github.com/nodejs/node/commit/a12658595b)] - **meta**: call `create-release-post.yml` post release (Aviv Keller) [#&#8203;60366](https://redirect.github.com/nodejs/node/pull/60366)
- \[[`8987159234`](https://redirect.github.com/nodejs/node/commit/8987159234)] - **(SEMVER-MINOR)** **module**: mark type stripping as stable (Marco Ippolito) [#&#8203;60600](https://redirect.github.com/nodejs/node/pull/60600)
- \[[`36da413663`](https://redirect.github.com/nodejs/node/commit/36da413663)] - **module**: fix directory option in the enableCompileCache() API (Joyee Cheung) [#&#8203;59931](https://redirect.github.com/nodejs/node/pull/59931)
- \[[`92c484ebf4`](https://redirect.github.com/nodejs/node/commit/92c484ebf4)] - **(SEMVER-MINOR)** **node-api**: add napi\_create\_object\_with\_properties (Miguel Marcondes Filho) [#&#8203;59953](https://redirect.github.com/nodejs/node/pull/59953)
- \[[`545162b0d4`](https://redirect.github.com/nodejs/node/commit/545162b0d4)] - **node-api**: use local files for instanceof test (Vladimir Morozov) [#&#8203;60190](https://redirect.github.com/nodejs/node/pull/60190)
- \[[`526c011d89`](https://redirect.github.com/nodejs/node/commit/526c011d89)] - **perf\_hooks**: fix stack overflow error (Antoine du Hamel) [#&#8203;60084](https://redirect.github.com/nodejs/node/pull/60084)
- \[[`1de0476939`](https://redirect.github.com/nodejs/node/commit/1de0476939)] - **perf\_hooks**: move non-standard performance properties to perf\_hooks (Chengzhong Wu) [#&#8203;60370](https://redirect.github.com/nodejs/node/pull/60370)
- \[[`07ec1239ef`](https://redirect.github.com/nodejs/node/commit/07ec1239ef)] - **repl**: fix pasting after moving the cursor to the left (Ruben Bridgewater) [#&#8203;60470](https://redirect.github.com/nodejs/node/pull/60470)
- \[[`b11bc5984e`](https://redirect.github.com/nodejs/node/commit/b11bc5984e)] - **(SEMVER-MINOR)** **sqlite**: allow setting defensive flag (Bart Louwers) [#&#8203;60217](https://redirect.github.com/nodejs/node/pull/60217)
- \[[`273c9661fd`](https://redirect.github.com/nodejs/node/commit/273c9661fd)] - **sqlite,doc**: fix StatementSync section (Edy Silva) [#&#8203;60474](https://redirect.github.com/nodejs/node/pull/60474)
- \[[`d92ec21a4c`](https://redirect.github.com/nodejs/node/commit/d92ec21a4c)] - **src**: use CP\_UTF8 for wide file names on win32 (Fedor Indutny) [#&#8203;60575](https://redirect.github.com/nodejs/node/pull/60575)
- \[[`baef0468ed`](https://redirect.github.com/nodejs/node/commit/baef0468ed)] - **src**: move Node-API version detection to where it is used (Anna Henningsen) [#&#8203;60512](https://redirect.github.com/nodejs/node/pull/60512)
- \[[`e7da5b4b7d`](https://redirect.github.com/nodejs/node/commit/e7da5b4b7d)] - **(SEMVER-MINOR)** **src**: add watch config namespace (Marco Ippolito) [#&#8203;60178](https://redirect.github.com/nodejs/node/pull/60178)
- \[[`a7f7d10c06`](https://redirect.github.com/nodejs/node/commit/a7f7d10c06)] - **(SEMVER-MINOR)** **src**: add an option to make compile cache portable (Aditi) [#&#8203;58797](https://redirect.github.com/nodejs/node/pull/58797)
- \[[`566add0b19`](https://redirect.github.com/nodejs/node/commit/566add0b19)] - **src**: avoid C strings in more C++ exception throws (Anna Henningsen) [#&#8203;60592](https://redirect.github.com/nodejs/node/pull/60592)
- \[[`9b796347c1`](https://redirect.github.com/nodejs/node/commit/9b796347c1)] - **src**: add internal binding for constructing SharedArrayBuffers (Renegade334) [#&#8203;60497](https://redirect.github.com/nodejs/node/pull/60497)
- \[[`3b01cbb411`](https://redirect.github.com/nodejs/node/commit/3b01cbb411)] - **src**: move `napi_addon_register_func` to `node_api_types.h` (Anna Henningsen) [#&#8203;60512](https://redirect.github.com/nodejs/node/pull/60512)
- \[[`02fb7f4ecb`](https://redirect.github.com/nodejs/node/commit/02fb7f4ecb)] - **src**: remove unconditional NAPI\_EXPERIMENTAL in node.h (Chengzhong Wu) [#&#8203;60345](https://redirect.github.com/nodejs/node/pull/60345)
- \[[`bd09ae24e4`](https://redirect.github.com/nodejs/node/commit/bd09ae24e4)] - **src**: clean up generic counter implementation (Anna Henningsen) [#&#8203;60447](https://redirect.github.com/nodejs/node/pull/60447)
- \[[`cd6bf51dbd`](https://redirect.github.com/nodejs/node/commit/cd6bf51dbd)] - **src**: add enum handle for ToStringHelper + formatting (Burkov Egor) [#&#8203;56829](https://redirect.github.com/nodejs/node/pull/56829)
- \[[`92ea669240`](https://redirect.github.com/nodejs/node/commit/92ea669240)] - **(SEMVER-MINOR)** **src,permission**: add --allow-inspector ability (Rafael Gonzaga) [#&#8203;59711](https://redirect.github.com/nodejs/node/pull/59711)
- \[[`ac3dbe48f7`](https://redirect.github.com/nodejs/node/commit/ac3dbe48f7)] - **stream**: don't try to read more if reading (Robert Nagy) [#&#8203;60454](https://redirect.github.com/nodejs/node/pull/60454)
- \[[`790288a93b`](https://redirect.github.com/nodejs/node/commit/790288a93b)] - **test**: ensure assertions are reachable in `test/internet` (Antoine du Hamel) [#&#8203;60513](https://redirect.github.com/nodejs/node/pull/60513)
- \[[`0a85132989`](https://redirect.github.com/nodejs/node/commit/0a85132989)] - **test**: fix status when compiled without inspector (Antoine du Hamel) [#&#8203;60289](https://redirect.github.com/nodejs/node/pull/60289)
- \[[`2f57673172`](https://redirect.github.com/nodejs/node/commit/2f57673172)] - **test**: deflake test-perf-hooks-timerify-histogram-sync (Joyee Cheung) [#&#8203;60639](https://redirect.github.com/nodejs/node/pull/60639)
- \[[`09726269de`](https://redirect.github.com/nodejs/node/commit/09726269de)] - **test**: apply a delay to `watch-mode-kill-signal` tests (Joyee Cheung) [#&#8203;60610](https://redirect.github.com/nodejs/node/pull/60610)
- \[[`45537b9562`](https://redirect.github.com/nodejs/node/commit/45537b9562)] - **test**: async iife in repl (Tony Gorez) [#&#8203;44878](https://redirect.github.com/nodejs/node/pull/44878)
- \[[`4ca81f101d`](https://redirect.github.com/nodejs/node/commit/4ca81f101d)] - **test**: parallelize sea tests when there's enough disk space (Joyee Cheung) [#&#8203;60604](https://redirect.github.com/nodejs/node/pull/60604)
- \[[`ea71e96191`](https://redirect.github.com/nodejs/node/commit/ea71e96191)] - **test**: only show overridden env in child process failures (Joyee Cheung) [#&#8203;60556](https://redirect.github.com/nodejs/node/pull/60556)
- \[[`06b2e348c7`](https://redirect.github.com/nodejs/node/commit/06b2e348c7)] - **test**: ensure assertions are reached on more tests (Antoine du Hamel) [#&#8203;60498](https://redirect.github.com/nodejs/node/pull/60498)
- \[[`de9c8cb670`](https://redirect.github.com/nodejs/node/commit/de9c8cb670)] - **test**: ensure assertions are reachable in `test/es-module` (Antoine du Hamel) [#&#8203;60501](https://redirect.github.com/nodejs/node/pull/60501)
- \[[`75bc40fced`](https://redirect.github.com/nodejs/node/commit/75bc40fced)] - **test**: ensure assertions are reached on more tests (Antoine du Hamel) [#&#8203;60485](https://redirect.github.com/nodejs/node/pull/60485)
- \[[`1a6084cfd3`](https://redirect.github.com/nodejs/node/commit/1a6084cfd3)] - **test**: ensure assertions are reached on more tests (Antoine du Hamel) [#&#8203;60500](https://redirect.github.com/nodejs/node/pull/60500)
- \[[`2c651c90cf`](https://redirect.github.com/nodejs/node/commit/2c651c90cf)] - **test**: split test-perf-hooks-timerify (Joyee Cheung) [#&#8203;60568](https://redirect.github.com/nodejs/node/pull/60568)
- \[[`6e8b5f7345`](https://redirect.github.com/nodejs/node/commit/6e8b5f7345)] - **test**: add more logs to test-esm-loader-hooks-inspect-wait (Joyee Cheung) [#&#8203;60466](https://redirect.github.com/nodejs/node/pull/60466)
- \[[`9dea7ffa30`](https://redirect.github.com/nodejs/node/commit/9dea7ffa30)] - **test**: mark stringbytes-external-exceed-max tests as flaky on AIX (Joyee Cheung) [#&#8203;60565](https://redirect.github.com/nodejs/node/pull/60565)
- \[[`0b3c3b710a`](https://redirect.github.com/nodejs/node/commit/0b3c3b710a)] - **test**: split test-esm-wasm.js (Joyee Cheung) [#&#8203;60491](https://redirect.github.com/nodejs/node/pull/60491)
- \[[`a15b795b34`](https://redirect.github.com/nodejs/node/commit/a15b795b34)] - **test**: correct conditional secure heap flags test (Shelley Vohr) [#&#8203;60385](https://redirect.github.com/nodejs/node/pull/60385)
- \[[`38b77b3a44`](https://redirect.github.com/nodejs/node/commit/38b77b3a44)] - **test**: fix flaky test-watch-mode-kill-signal-\* (Joyee Cheung) [#&#8203;60443](https://redirect.github.com/nodejs/node/pull/60443)
- \[[`e8d7598057`](https://redirect.github.com/nodejs/node/commit/e8d7598057)] - **test**: capture stack trace in debugger timeout errors (Joyee Cheung) [#&#8203;60457](https://redirect.github.com/nodejs/node/pull/60457)
- \[[`674befeb81`](https://redirect.github.com/nodejs/node/commit/674befeb81)] - **test**: ensure assertions are reachable in `test/sequential` (Antoine du Hamel) [#&#8203;60412](https://redirect.github.com/nodejs/node/pull/60412)
- \[[`952c08a735`](https://redirect.github.com/nodejs/node/commit/952c08a735)] - **test**: ensure assertions are reachable in more folders (Antoine du Hamel) [#&#8203;60411](https://redirect.github.com/nodejs/node/pull/60411)
- \[[`bbca57584b`](https://redirect.github.com/nodejs/node/commit/bbca57584b)] - **test**: split test-runner-watch-mode (Joyee Cheung) [#&#8203;60391](https://redirect.github.com/nodejs/node/pull/60391)
- \[[`e78e0cf6e7`](https://redirect.github.com/nodejs/node/commit/e78e0cf6e7)] - **test**: move test-runner-watch-mode helper into common (Joyee Cheung) [#&#8203;60391](https://redirect.github.com/nodejs/node/pull/60391)
- \[[`84576ef021`](https://redirect.github.com/nodejs/node/commit/84576ef021)] - **test**: ensure assertions are reachable in `test/addons` (Antoine du Hamel) [#&#8203;60142](https://redirect.github.com/nodejs/node/pull/60142)
- \[[`1659078c11`](https://redirect.github.com/nodejs/node/commit/1659078c11)] - **test**: ignore EPIPE errors in https proxy invalid URL test (Joyee Cheung) [#&#8203;60269](https://redirect.github.com/nodejs/node/pull/60269)
- \[[`79ffee80ec`](https://redirect.github.com/nodejs/node/commit/79ffee80ec)] - **test**: ensure assertions are reachable in `test/client-proxy` (Antoine du Hamel) [#&#8203;60175](https://redirect.github.com/nodejs/node/pull/60175)
- \[[`e5a812243a`](https://redirect.github.com/nodejs/node/commit/e5a812243a)] - **test**: ensure assertions are reachable in `test/async-hooks` (Antoine du Hamel) [#&#8203;60150](https://redirect.github.com/nodejs/node/pull/60150)
- \[[`e924fd72e3`](https://redirect.github.com/nodejs/node/commit/e924fd72e3)] - **test,crypto**: handle a few more BoringSSL tests (Shelley Vohr) [#&#8203;59030](https://redirect.github.com/nodejs/node/pull/59030)
- \[[`a55ac11611`](https://redirect.github.com/nodejs/node/commit/a55ac11611)] - **test,crypto**: update x448 and ed448 expectation when on boringssl (Shelley Vohr) [#&#8203;60387](https://redirect.github.com/nodejs/node/pull/60387)
- \[[`55d5e9ec73`](https://redirect.github.com/nodejs/node/commit/55d5e9ec73)] - **tls**: fix leak on invalid protocol method (Shelley Vohr) [#&#8203;60427](https://redirect.github.com/nodejs/node/pull/60427)
- \[[`5763c96e7c`](https://redirect.github.com/nodejs/node/commit/5763c96e7c)] - **tools**: replace invalid expression in dependabot config (Riddhi) [#&#8203;60649](https://redirect.github.com/nodejs/node/pull/60649)
- \[[`b6e21b47d7`](https://redirect.github.com/nodejs/node/commit/b6e21b47d7)] - **tools**: skip unaffected GHA jobs for changes in `test/internet` (Antoine du Hamel) [#&#8203;60517](https://redirect.github.com/nodejs/node/pull/60517)
- \[[`999664c76d`](https://redirect.github.com/nodejs/node/commit/999664c76d)] - **tools**: do not use short hashes for deps versioning to avoid collision (Antoine du Hamel) [#&#8203;60407](https://redirect.github.com/nodejs/node/pull/60407)
- \[[`ada856d0fb`](https://redirect.github.com/nodejs/node/commit/ada856d0fb)] - **tools**: only add test reporter args when node:test is used (Joyee Cheung) [#&#8203;60551](https://redirect.github.com/nodejs/node/pull/60551)
- \[[`1812c56bb3`](https://redirect.github.com/nodejs/node/commit/1812c56bb3)] - **tools**: fix update-icu script (Michaël Zasso) [#&#8203;60521](https://redirect.github.com/nodejs/node/pull/60521)
- \[[`747040438a`](https://redirect.github.com/nodejs/node/commit/747040438a)] - **tools**: fix linter for semver-major release proposals (Antoine du Hamel) [#&#8203;60481](https://redirect.github.com/nodejs/node/pull/60481)
- \[[`f170551e40`](https://redirect.github.com/nodejs/node/commit/f170551e40)] - **tools**: fix failing release-proposal linter for LTS transitions (Antoine du Hamel) [#&#8203;60465](https://redirect.github.com/nodejs/node/pull/60465)
- \[[`2db4ea0ce4`](https://redirect.github.com/nodejs/node/commit/2db4ea0ce4)] - **tools**: remove undici from daily wpt.fyi job (Filip Skokan) [#&#8203;60444](https://redirect.github.com/nodejs/node/pull/60444)
- \[[`2a85aa4e7b`](https://redirect.github.com/nodejs/node/commit/2a85aa4e7b)] - **tools**: add lint rule to ensure assertions are reached (Antoine du Hamel) [#&#8203;60125](https://redirect.github.com/nodejs/node/pull/60125)
- \[[`48299ef5fb`](https://redirect.github.com/nodejs/node/commit/48299ef5fb)] - **tools,doc**: update JavaScript primitive types to match MDN Web Docs (JustApple) [#&#8203;60581](https://redirect.github.com/nodejs/node/pull/60581)
- \[[`7ec04cf936`](https://redirect.github.com/nodejs/node/commit/7ec04cf936)] - **util**: fix stylize of special properties in inspect (Ge Gao) [#&#8203;60479](https://redirect.github.com/nodejs/node/pull/60479)
- \[[`05d7509bd2`](https://redirect.github.com/nodejs/node/commit/05d7509bd2)] - **(SEMVER-MINOR)** **v8**: add cpu profile (theanarkh) [#&#8203;59807](https://redirect.github.com/nodejs/node/pull/59807)
- \[[`884fe884a1`](https://redirect.github.com/nodejs/node/commit/884fe884a1)] - **vm**: hint module identifier in instantiate errors (Chengzhong Wu) [#&#8203;60199](https://redirect.github.com/nodejs/node/pull/60199)
- \[[`a2caf19f70`](https://redirect.github.com/nodejs/node/commit/a2caf19f70)] - **watch**: fix interaction with multiple env files (Marco Ippolito) [#&#8203;60605](https://redirect.github.com/nodejs/node/pull/60605)

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6W119-->


---

## Reviews

### Review by @TenSt - Approved on December 17, 2025 at 12:21 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1338*
