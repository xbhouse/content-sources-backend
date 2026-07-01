---
type: pull_request
number: 1449
title: "chore(deps): update node.js to v24"
state: merged
author: red-hat-konflux
created: 2025-12-08T08:38:39Z
updated: 2026-02-10T08:48:23Z
closed: 2026-02-10T07:47:21Z
merged: 2026-02-10T07:47:21Z
base_branch: master
head_branch: konflux/mintmaker/master/node-24.x
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1449
---

# Pull Request #1449: chore(deps): update node.js to v24

**Author**: @red-hat-konflux
**Created**: December 08, 2025 at 08:38 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/node-24.x`

## Description

> **Note:** This PR body was truncated due to platform limits.

This PR contains the following updates:

| Package | Update | Change |
|---|---|---|
| [node](https://nodejs.org) ([source](https://redirect.github.com/nodejs/node)) | major | `22` -> `24` |

---

### Release Notes

<details>
<summary>nodejs/node (node)</summary>

### [`v24.13.0`](https://redirect.github.com/nodejs/node/releases/tag/v24.13.0): 2026-01-13, Version 24.13.0 &#x27;Krypton&#x27; (LTS), @&#8203;marco-ippolito

[Compare Source](https://redirect.github.com/nodejs/node/compare/v24.12.0...v24.13.0)

This is a security release.

##### Notable Changes

lib:

- (CVE-2025-59465) add TLSSocket default error handler (RafaelGSS) [nodejs-private/node-private#797](https://redirect.github.com/nodejs-private/node-private/pull/797)
- (CVE-2025-55132) disable futimes when permission model is enabled (RafaelGSS) [nodejs-private/node-private#748](https://redirect.github.com/nodejs-private/node-private/pull/748)
  lib,permission:
- (CVE-2025-55130) require full read and write to symlink APIs (RafaelGSS) [nodejs-private/node-private#760](https://redirect.github.com/nodejs-private/node-private/pull/760)
  src:
- (CVE-2025-59466) rethrow stack overflow exceptions in async\_hooks (Matteo Collina) [nodejs-private/node-private#773](https://redirect.github.com/nodejs-private/node-private/pull/773)
  src,lib:
- (CVE-2025-55131) refactor unsafe buffer creation to remove zero-fill toggle (Сковорода Никита Андреевич) [nodejs-private/node-private#759](https://redirect.github.com/nodejs-private/node-private/pull/759)
  tls:
- (CVE-2026-21637) route callback exceptions through error handlers (Matteo Collina) [nodejs-private/node-private#796](https://redirect.github.com/nodejs-private/node-private/pull/796)

##### Commits

- \[[`2092785d01`](https://redirect.github.com/nodejs/node/commit/2092785d01)] - **deps**: update c-ares to v1.34.6 (Node.js GitHub Bot) [#&#8203;60997](https://redirect.github.com/nodejs/node/pull/60997)
- \[[`3e58b7f2af`](https://redirect.github.com/nodejs/node/commit/3e58b7f2af)] - **deps**: update undici to 7.18.2 (Node.js GitHub Bot) [#&#8203;61283](https://redirect.github.com/nodejs/node/pull/61283)
- \[[`4ba536a5a6`](https://redirect.github.com/nodejs/node/commit/4ba536a5a6)] - **(CVE-2025-59465)** **lib**: add TLSSocket default error handler (RafaelGSS) [nodejs-private/node-private#797](https://redirect.github.com/nodejs-private/node-private/pull/797)
- \[[`89adaa21fd`](https://redirect.github.com/nodejs/node/commit/89adaa21fd)] - **(CVE-2025-55132)** **lib**: disable futimes when permission model is enabled (RafaelGSS) [nodejs-private/node-private#748](https://redirect.github.com/nodejs-private/node-private/pull/748)
- \[[`7302b4dae1`](https://redirect.github.com/nodejs/node/commit/7302b4dae1)] - **(CVE-2025-55130)** **lib,permission**: require full read and write to symlink APIs (RafaelGSS) [nodejs-private/node-private#760](https://redirect.github.com/nodejs-private/node-private/pull/760)
- \[[`ac030753c4`](https://redirect.github.com/nodejs/node/commit/ac030753c4)] - **(CVE-2025-59466)** **src**: rethrow stack overflow exceptions in async\_hooks (Matteo Collina) [nodejs-private/node-private#773](https://redirect.github.com/nodejs-private/node-private/pull/773)
- \[[`20075692fe`](https://redirect.github.com/nodejs/node/commit/20075692fe)] - **(CVE-2025-55131)** **src,lib**: refactor unsafe buffer creation to remove zero-fill toggle (Сковорода Никита Андреевич) [nodejs-private/node-private#759](https://redirect.github.com/nodejs-private/node-private/pull/759)
- \[[`20591b0618`](https://redirect.github.com/nodejs/node/commit/20591b0618)] - **(CVE-2026-21637)** **tls**: route callback exceptions through error handlers (Matteo Collina) [nodejs-private/node-private#796](https://redirect.github.com/nodejs-private/node-private/pull/796)

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

### [`v24.11.1`](https://redirect.github.com/nodejs/node/releases/tag/v24.11.1): 2025-11-11, Version 24.11.1 &#x27;Krypton&#x27; (LTS), @&#8203;aduh95

[Compare Source](https://redirect.github.com/nodejs/node/compare/v24.11.0...v24.11.1)

##### Notable Changes

The known issue relating to `Buffer.allocUnsafe` incorrectly zero-filling buffers
has now been addressed and now returns uninitialized memory as documented in the
[`Buffer.allocUnsafe`](https://nodejs.org/docs/latest-v24.x/api/buffer.html#static-method-bufferallocunsafesize)
documentation.

##### Commits

- \[[`0a15ccf3f4`](https://redirect.github.com/nodejs/node/commit/0a15ccf3f4)] - **benchmark**: improve cpu.sh for safety and usability (Nam Yooseong) [#&#8203;60162](https://redirect.github.com/nodejs/node/pull/60162)
- \[[`a1c7d1dac9`](https://redirect.github.com/nodejs/node/commit/a1c7d1dac9)] - **benchmark**: add benchmark for leaf source text modules (Joyee Cheung) [#&#8203;60205](https://redirect.github.com/nodejs/node/pull/60205)
- \[[`99e2acf46b`](https://redirect.github.com/nodejs/node/commit/99e2acf46b)] - **benchmark**: add vm.SourceTextModule benchmark (Joyee Cheung) [#&#8203;59396](https://redirect.github.com/nodejs/node/pull/59396)
- \[[`c01c72b407`](https://redirect.github.com/nodejs/node/commit/c01c72b407)] - **benchmark**: use non-deprecated WriteUtf8V2 method (Michaël Zasso) [#&#8203;60173](https://redirect.github.com/nodejs/node/pull/60173)
- \[[`a42dbd138e`](https://redirect.github.com/nodejs/node/commit/a42dbd138e)] - **build**: ibmi follow aix visibility (SRAVANI GUNDEPALLI) [#&#8203;60360](https://redirect.github.com/nodejs/node/pull/60360)
- \[[`5673a54a5d`](https://redirect.github.com/nodejs/node/commit/5673a54a5d)] - **build**: use call command when calling python configure (Jacob Nichols) [#&#8203;60098](https://redirect.github.com/nodejs/node/pull/60098)
- \[[`c67cb727cb`](https://redirect.github.com/nodejs/node/commit/c67cb727cb)] - **build**: build v8 with -fvisibility=hidden -fvisibility-inlines-hidden (Joyee Cheung) [#&#8203;56290](https://redirect.github.com/nodejs/node/pull/56290)
- \[[`b03f7b93b1`](https://redirect.github.com/nodejs/node/commit/b03f7b93b1)] - **build**: remove V8\_COMPRESS\_POINTERS\_IN\_ISOLATE\_CAGE defs (Joyee Cheung) [#&#8203;60296](https://redirect.github.com/nodejs/node/pull/60296)
- \[[`2505568531`](https://redirect.github.com/nodejs/node/commit/2505568531)] - **build, src**: fix include paths for vtune files (Rahul) [#&#8203;59999](https://redirect.github.com/nodejs/node/pull/59999)
- \[[`95330b036f`](https://redirect.github.com/nodejs/node/commit/95330b036f)] - **crypto**: update root certificates to NSS 3.116 (Node.js GitHub Bot) [#&#8203;59956](https://redirect.github.com/nodejs/node/pull/59956)
- \[[`c221d892ef`](https://redirect.github.com/nodejs/node/commit/c221d892ef)] - **deps**: update corepack to 0.34.2 (Node.js GitHub Bot) [#&#8203;60550](https://redirect.github.com/nodejs/node/pull/60550)
- \[[`bc00aa4c77`](https://redirect.github.com/nodejs/node/commit/bc00aa4c77)] - **deps**: update simdjson to 4.0.7 (Node.js GitHub Bot) [#&#8203;59883](https://redirect.github.com/nodejs/node/pull/59883)
- \[[`d03b89ec53`](https://redirect.github.com/nodejs/node/commit/d03b89ec53)] - **deps**: update corepack to 0.34.1 (Node.js GitHub Bot) [#&#8203;60314](https://redirect.github.com/nodejs/node/pull/60314)
- \[[`b7882090de`](https://redirect.github.com/nodejs/node/commit/b7882090de)] - **deps**: update inspector\_protocol to [`af7f5a8`](https://redirect.github.com/nodejs/node/commit/af7f5a8173fdbc29f0835ec94395932e328b) (Node.js GitHub Bot) [#&#8203;60312](https://redirect.github.com/nodejs/node/pull/60312)
- \[[`7007f9dd65`](https://redirect.github.com/nodejs/node/commit/7007f9dd65)] - **deps**: update googletest to [`279f847`](https://redirect.github.com/nodejs/node/commit/279f847) (Node.js GitHub Bot) [#&#8203;60219](https://redirect.github.com/nodejs/node/pull/60219)
- \[[`a56aa9ffa8`](https://redirect.github.com/nodejs/node/commit/a56aa9ffa8)] - **deps**: upgrade npm to 11.6.2 (npm team) [#&#8203;60168](https://redirect.github.com/nodejs/node/pull/60168)
- \[[`0bf8952721`](https://redirect.github.com/nodejs/node/commit/0bf8952721)] - **doc**: mention more codemods in `deprecations.md` (Augustin Mauroy) [#&#8203;60243](https://redirect.github.com/nodejs/node/pull/60243)
- \[[`2473ca77f6`](https://redirect.github.com/nodejs/node/commit/2473ca77f6)] - **doc**: add missing CAA type to dns.resolveAny() & dnsPromises.resolveAny() (Jimmy Leung) [#&#8203;58899](https://redirect.github.com/nodejs/node/pull/58899)
- \[[`39ddd8522e`](https://redirect.github.com/nodejs/node/commit/39ddd8522e)] - **doc**: use `any` for `worker_threads.Worker` 'error' event argument `err` (Jonas Geiler) [#&#8203;60300](https://redirect.github.com/nodejs/node/pull/60300)
- \[[`eaa825fd97`](https://redirect.github.com/nodejs/node/commit/eaa825fd97)] - **doc**: update decorator documentation to reflect actual policy (Muhammad Salman Aziz) [#&#8203;60288](https://redirect.github.com/nodejs/node/pull/60288)
- \[[`a744e42282`](https://redirect.github.com/nodejs/node/commit/a744e42282)] - **doc**: document wildcard supported by tools/test.py (Joyee Cheung) [#&#8203;60265](https://redirect.github.com/nodejs/node/pull/60265)
- \[[`ec0d5beb09`](https://redirect.github.com/nodejs/node/commit/ec0d5beb09)] - **doc**: add --heap-snapshot-on-oom to useful v8 flag (jakecastelli) [#&#8203;60260](https://redirect.github.com/nodejs/node/pull/60260)
- \[[`13da0df12a`](https://redirect.github.com/nodejs/node/commit/13da0df12a)] - **doc**: fix `blob.bytes()` heading level (XTY) [#&#8203;60252](https://redirect.github.com/nodejs/node/pull/60252)
- \[[`8e771632b7`](https://redirect.github.com/nodejs/node/commit/8e771632b7)] - **doc**: fix not working code example in vm docs (Artur Gawlik) [#&#8203;60224](https://redirect.github.com/nodejs/node/pull/60224)
- \[[`70c2080bff`](https://redirect.github.com/nodejs/node/commit/70c2080bff)] - **doc**: improve code snippet alternative of url.parse() using WHATWG URL (Steven) [#&#8203;60209](https://redirect.github.com/nodejs/node/pull/60209)
- \[[`beadcf176e`](https://redirect.github.com/nodejs/node/commit/beadcf176e)] - **doc**: `createSQLTagStore` -> `createTagStore` (Aviv Keller) [#&#8203;60182](https://redirect.github.com/nodejs/node/pull/60182)
- \[[`b0da3b9c6a`](https://redirect.github.com/nodejs/node/commit/b0da3b9c6a)] - **doc**: use markdown when branch-diff major release (Rafael Gonzaga) [#&#8203;60179](https://redirect.github.com/nodejs/node/pull/60179)
- \[[`688115aa6b`](https://redirect.github.com/nodejs/node/commit/688115aa6b)] - **doc**: update teams in collaborator-guide.md and add links (Bart Louwers) [#&#8203;60065](https://redirect.github.com/nodejs/node/pull/60065)
- \[[`923082a064`](https://redirect.github.com/nodejs/node/commit/923082a064)] - **doc**: disambiguate top-level `worker_threads` module exports (René) [#&#8203;59890](https://redirect.github.com/nodejs/node/pull/59890)
- \[[`7be4330870`](https://redirect.github.com/nodejs/node/commit/7be4330870)] - **doc**: add known issue to v24.11.0 release notes (Richard Lau) [#&#8203;60467](https://redirect.github.com/nodejs/node/pull/60467)
- \[[`4d8f62aeaf`](https://redirect.github.com/nodejs/node/commit/4d8f62aeaf)] - **doc, module**: change async customization hooks to experimental (Gerhard Stöbich) [#&#8203;60302](https://redirect.github.com/nodejs/node/pull/60302)
- \[[`d86a118bbd`](https://redirect.github.com/nodejs/node/commit/d86a118bbd)] - **http**: lazy allocate cookies array (Robert Nagy) [#&#8203;59734](https://redirect.github.com/nodejs/node/pull/59734)
- \[[`8c256d4139`](https://redirect.github.com/nodejs/node/commit/8c256d4139)] - **http**: fix http client leaky with double response (theanarkh) [#&#8203;60062](https://redirect.github.com/nodejs/node/pull/60062)
- \[[`265e9d59fa`](https://redirect.github.com/nodejs/node/commit/265e9d59fa)] - **http2**: rename variable to additionalPseudoHeaders (Tobias Nießen) [#&#8203;60208](https://redirect.github.com/nodejs/node/pull/60208)
- \[[`65bec037e2`](https://redirect.github.com/nodejs/node/commit/65bec037e2)] - **http2**: do not crash on mismatched ping buffer length (René) [#&#8203;60135](https://redirect.github.com/nodejs/node/pull/60135)
- \[[`9b83ef53b7`](https://redirect.github.com/nodejs/node/commit/9b83ef53b7)] - **inspector**: add network payload buffer size limits (Chengzhong Wu) [#&#8203;60236](https://redirect.github.com/nodejs/node/pull/60236)
- \[[`03ac05c458`](https://redirect.github.com/nodejs/node/commit/03ac05c458)] - **inspector**: support handshake response for websocket inspection (Shima Ryuhei) [#&#8203;60225](https://redirect.github.com/nodejs/node/pull/60225)
- \[[`aa04f06190`](https://redirect.github.com/nodejs/node/commit/aa04f06190)] - **lib**: fix typo in createBlobReaderStream (SeokHun) [#&#8203;60132](https://redirect.github.com/nodejs/node/pull/60132)
- \[[`5aea1a429e`](https://redirect.github.com/nodejs/node/commit/5aea1a429e)] - **lib**: fix constructor in \_errnoException stack tree (SeokHun) [#&#8203;60156](https://redirect.github.com/nodejs/node/pull/60156)
- \[[`4f7745acc7`](https://redirect.github.com/nodejs/node/commit/4f7745acc7)] - **lib**: fix typo in QuicSessionStats (SeokHun) [#&#8203;60155](https://redirect.github.com/nodejs/node/pull/60155)
- \[[`f8725861ea`](https://redirect.github.com/nodejs/node/commit/f8725861ea)] - **lib**: remove redundant destroyHook checks (Gürgün Dayıoğlu) [#&#8203;60120](https://redirect.github.com/nodejs/node/pull/60120)
- \[[`696c20bf3f`](https://redirect.github.com/nodejs/node/commit/696c20bf3f)] - **meta**: move one or more collaborators to emeritus (Node.js GitHub Bot) [#&#8203;60325](https://redirect.github.com/nodejs/node/pull/60325)
- \[[`90434ff99a`](https://redirect.github.com/nodejs/node/commit/90434ff99a)] - **meta**: loop userland-migrations in deprecations (Chengzhong Wu) [#&#8203;60299](https://redirect.github.com/nodejs/node/pull/60299)
- \[[`ffbc0ae60a`](https://redirect.github.com/nodejs/node/commit/ffbc0ae60a)] - **module**: refactor and clarify async loader hook customizations (Joyee Cheung) [#&#8203;60278](https://redirect.github.com/nodejs/node/pull/60278)
- \[[`6ed6062f7d`](https://redirect.github.com/nodejs/node/commit/6ed6062f7d)] - **module**: handle null source from async loader hooks in sync hooks (Joyee Cheung) [#&#8203;59929](https://redirect.github.com/nodejs/node/pull/59929)
- \[[`a2871baed2`](https://redirect.github.com/nodejs/node/commit/a2871baed2)] - **msi**: fix WiX warnings (Stefan Stojanovic) [#&#8203;60251](https://redirect.github.com/nodejs/node/pull/60251)
- \[[`6199541d67`](https://redirect.github.com/nodejs/node/commit/6199541d67)] - **src**: fix timing of snapshot serialize callback (Joyee Cheung) [#&#8203;60434](https://redirect.github.com/nodejs/node/pull/60434)
- \[[`13b687959a`](https://redirect.github.com/nodejs/node/commit/13b687959a)] - **src**: add COUNT\_GENERIC\_USAGE utility for tests (Joyee Cheung) [#&#8203;60434](https://redirect.github.com/nodejs/node/pull/60434)
- \[[`a587623b4f`](https://redirect.github.com/nodejs/node/commit/a587623b4f)] - **src**: conditionally disable source phase imports by default (Shelley Vohr) [#&#8203;60364](https://redirect.github.com/nodejs/node/pull/60364)
- \[[`e483267995`](https://redirect.github.com/nodejs/node/commit/e483267995)] - **src**: use cached primordials\_string (Sohyeon Kim) [#&#8203;60255](https://redirect.github.com/nodejs/node/pull/60255)
- \[[`4c9a64fbaf`](https://redirect.github.com/nodejs/node/commit/4c9a64fbaf)] - **src**: replace Environment::GetCurrent with args.GetIsolate (Sohyeon Kim) [#&#8203;60256](https://redirect.github.com/nodejs/node/pull/60256)
- \[[`eb8a0493d1`](https://redirect.github.com/nodejs/node/commit/eb8a0493d1)] - **src**: initial enablement of IsolateGroups (James M Snell) [#&#8203;60254](https://redirect.github.com/nodejs/node/pull/60254)
- \[[`463c6450cf`](https://redirect.github.com/nodejs/node/commit/463c6450cf)] - **src**: use `Utf8Value` and `TwoByteValue` instead of V8 helpers (Anna Henningsen) [#&#8203;60244](https://redirect.github.com/nodejs/node/pull/60244)
- \[[`b370e02789`](https://redirect.github.com/nodejs/node/commit/b370e02789)] - **src**: add a default branch for module phase (Chengzhong Wu) [#&#8203;60261](https://redirect.github.com/nodejs/node/pull/60261)
- \[[`4e1c5c5601`](https://redirect.github.com/nodejs/node/commit/4e1c5c5601)] - **src**: make additional cleanups in node locks impl (James M Snell) [#&#8203;60061](https://redirect.github.com/nodejs/node/pull/60061)
- \[[`f00d4c10fc`](https://redirect.github.com/nodejs/node/commit/f00d4c10fc)] - **src**: update locks to use DictionaryTemplate (James M Snell) [#&#8203;60061](https://redirect.github.com/nodejs/node/pull/60061)
- \[[`1c8716e97c`](https://redirect.github.com/nodejs/node/commit/1c8716e97c)] - **test**: increase debugger waitFor timeout on macOS (Chengzhong Wu) [#&#8203;60367](https://redirect.github.com/nodejs/node/pull/60367)
- \[[`17b4f38e9c`](https://redirect.github.com/nodejs/node/commit/17b4f38e9c)] - **test**: put helper in test-runner-output into common (Joyee Cheung) [#&#8203;60330](https://redirect.github.com/nodejs/node/pull/60330)
- \[[`43b9ea8389`](https://redirect.github.com/nodejs/node/commit/43b9ea8389)] - **test**: fix small compile warning in test\_network\_requests\_buffer.cc (xiaocainiao633) [#&#8203;60281](https://redirect.github.com/nodejs/node/pull/60281)
- \[[`38a62980ad`](https://redirect.github.com/nodejs/node/commit/38a62980ad)] - **test**: split test-runner-watch-mode-kill-signal (Joyee Cheung) [#&#8203;60298](https://redirect.github.com/nodejs/node/pull/60298)
- \[[`34e4c8c84f`](https://redirect.github.com/nodejs/node/commit/34e4c8c84f)] - **test**: fix incorrect calculation in test-perf-hooks.js (Joyee Cheung) [#&#8203;60271](https://redirect.github.com/nodejs/node/pull/60271)
- \[[`4481feb17b`](https://redirect.github.com/nodejs/node/commit/4481feb17b)] - **test**: parallelize test-without-async-context-frame correctly (Joyee Cheung) [#&#8203;60273](https://redirect.github.com/nodejs/node/pull/60273)
- \[[`91ea9b06e0`](https://redirect.github.com/nodejs/node/commit/91ea9b06e0)] - **test**: skip sea tests on x64 macOS (Joyee Cheung) [#&#8203;60250](https://redirect.github.com/nodejs/node/pull/60250)
- \[[`cedba09e60`](https://redirect.github.com/nodejs/node/commit/cedba09e60)] - **test**: move sea tests into test/sea (Joyee Cheung) [#&#8203;60250](https://redirect.github.com/nodejs/node/pull/60250)
- \[[`635af55e12`](https://redirect.github.com/nodejs/node/commit/635af55e12)] - ***Revert*** "**test**: ensure message event fires in worker message port test" (Luigi Pinca) [#&#8203;60126](https://redirect.github.com/nodejs/node/pull/60126)
- \[[`68f678028e`](https://redirect.github.com/nodejs/node/commit/68f678028e)] - **test**: skip tests that cause timeouts on IBM i (SRAVANI GUNDEPALLI) [#&#8203;60148](https://redirect.github.com/nodejs/node/pull/60148)
- \[[`cc3a70598c`](https://redirect.github.com/nodejs/node/commit/cc3a70598c)] - **test**: deflake test-fs-promises-watch-iterator (Luigi Pinca) [#&#8203;60060](https://redirect.github.com/nodejs/node/pull/60060)
- \[[`3d784dd766`](https://redirect.github.com/nodejs/node/commit/3d784dd766)] - **test**: prepare junit file attribute normalization (sangwook) [#&#8203;59432](https://redirect.github.com/nodejs/node/pull/59432)
- \[[`84974d97ad`](https://redirect.github.com/nodejs/node/commit/84974d97ad)] - **test**: skip failing test on macOS 15.7+ (Antoine du Hamel) [#&#8203;60419](https://redirect.github.com/nodejs/node/pull/60419)
- \[[`fabf8e4975`](https://redirect.github.com/nodejs/node/commit/fabf8e4975)] - **test,crypto**: fix conditional SHA3-\* skip on BoringSSL (Filip Skokan) [#&#8203;60379](https://redirect.github.com/nodejs/node/pull/60379)
- \[[`8faa494bf2`](https://redirect.github.com/nodejs/node/commit/8faa494bf2)] - **test,crypto**: sha3 algorithms aren't supported with BoringSSL (Shelley Vohr) [#&#8203;60374](https://redirect.github.com/nodejs/node/pull/60374)
- \[[`538a00c0f6`](https://redirect.github.com/nodejs/node/commit/538a00c0f6)] - **test,doc**: skip --max-old-space-size-percentage on 32-bit platforms (Asaf Federman) [#&#8203;60144](https://redirect.github.com/nodejs/node/pull/60144)
- \[[`9ac5dbb694`](https://redirect.github.com/nodejs/node/commit/9ac5dbb694)] - **test\_runner**: use module.registerHooks in module mocks (Joyee Cheung) [#&#8203;60326](https://redirect.github.com/nodejs/node/pull/60326)
- \[[`f6ff6e7166`](https://redirect.github.com/nodejs/node/commit/f6ff6e7166)] - **test\_runner**: fix suite timeout (Moshe Atlow) [#&#8203;59853](https://redirect.github.com/nodejs/node/pull/59853)
- \[[`455bfeb52d`](https://redirect.github.com/nodejs/node/commit/455bfeb52d)] - **test\_runner**: add junit file attribute support (sangwook) [#&#8203;59432](https://redirect.github.com/nodejs/node/pull/59432)
- \[[`223c5e105d`](https://redirect.github.com/nodejs/node/commit/223c5e105d)] - **tools**: update gyp-next to 0.20.5 (Node.js GitHub Bot) [#&#8203;60313](https://redirect.github.com/nodejs/node/pull/60313)
- \[[`2949408fc1`](https://redirect.github.com/nodejs/node/commit/2949408fc1)] - **tools**: limit inspector protocol PR title length (Chengzhong Wu) [#&#8203;60324](https://redirect.github.com/nodejs/node/pull/60324)
- \[[`b36a898650`](https://redirect.github.com/nodejs/node/commit/b36a898650)] - **tools**: fix inspector\_protocol updater (Chengzhong Wu) [#&#8203;60277](https://redirect.github.com/nodejs/node/pull/60277)
- \[[`d60f002b62`](https://redirect.github.com/nodejs/node/commit/d60f002b62)] - **tools**: optimize wildcard execution in tools/test.py (Joyee Cheung) [#&#8203;60266](https://redirect.github.com/nodejs/node/pull/60266)
- \[[`9d4e422419`](https://redirect.github.com/nodejs/node/commit/9d4e422419)] - **tools**: add inspector\_protocol updater (Chengzhong Wu) [#&#8203;60245](https://redirect.github.com/nodejs/node/pull/60245)
- \[[`2f93a9894f`](https://redirect.github.com/nodejs/node/commit/2f93a9894f)] - **tools**: use cooldown property correctly (Rafael Gonzaga) [#&#8203;60134](https://redirect.github.com/nodejs/node/pull/60134)
- \[[`9468ade95d`](https://redirect.github.com/nodejs/node/commit/9468ade95d)] - **typings**: add missing properties and method in Worker (Woohyun Sung) [#&#8203;60257](https://redirect.github.com/nodejs/node/pull/60257)
- \[[`f611ec0a9e`](https://redirect.github.com/nodejs/node/commit/f611ec0a9e)] - **typings**: add missing properties in HTTPParser (Woohyun Sung) [#&#8203;60257](https://redirect.github.com/nodejs/node/pull/60257)
- \[[`301c1347a1`](https://redirect.github.com/nodejs/node/commit/301c1347a1)] - **typings**: delete undefined property in ConfigBinding (Woohyun Sung) [#&#8203;60257](https://redirect.github.com/nodejs/node/pull/60257)
- \[[`80fdb3d39b`](https://redirect.github.com/nodejs/node/commit/80fdb3d39b)] - **typings**: add buffer internalBinding typing (방진혁) [#&#8203;60163](https://redirect.github.com/nodejs/node/pull/60163)
- \[[`8cb3b77039`](https://redirect.github.com/nodejs/node/commit/8cb3b77039)] - **util**: use more defensive code when inspecting error objects (Antoine du Hamel) [#&#8203;60139](https://redirect.github.com/nodejs/node/pull/60139)
- \[[`748d4f6430`](https://redirect.github.com/nodejs/node/commit/748d4f6430)] - **util**: mark special properties when inspecting them (Ruben Bridgewater) [#&#8203;60131](https://redirect.github.com/nodejs/node/pull/60131)
- \[[`6183a759d7`](https://redirect.github.com/nodejs/node/commit/6183a759d7)] - **vm**: make vm.Module.evaluate() conditionally synchronous (Joyee Cheung) [#&#8203;60205](https://redirect.github.com/nodejs/node/pull/60205)
- \[[`4b8506628f`](https://redirect.github.com/nodejs/node/commit/4b8506628f)] - **win**: upgrade Visual Studio workload from 2019 to 2022 (Jiawen Geng) [#&#8203;60318](https://redirect.github.com/nodejs/node/pull/60318)

### [`v24.11.0`](https://redirect.github.com/nodejs/node/releases/tag/v24.11.0): 2025-10-28, Version 24.11.0 &#x27;Krypton&#x27; (LTS), @&#8203;richardlau

[Compare Source](https://redirect.github.com/nodejs/node/compare/v24.10.0...v24.11.0)

##### Notable Changes

This release marks the transition of Node.js 24.x into Long Term Support (LTS)
with the codename 'Krypton'. It will continue to receive updates through to
the end of April 2028.

Other than updating metadata, such as the `process.release` object, to reflect
that the release is LTS, no further changes from Node.js 24.10.0 are included.

##### Known issue

An issue has been identified in the Node.js 24.x line with `Buffer.allocUnsafe`
unintentionally returning zero-filled buffers. This API is
[documented to return uninitialized memory](https://nodejs.org/docs/latest-v24.x/api/buffer.html#static-method-bufferallocunsafesize).
The documented behavior will be restored in the next Node.js 24.x LTS release to bring
it back in line with previous releases. For more information, see
[#&#8203;60423](https://redirect.github.com/nodejs/node/issues/60423).

### [`v24.10.0`](https://redirect.github.com/nodejs/node/releases/tag/v24.10.0): 2025-10-08, Version 24.10.0 (Current), @&#8203;RafaelGSS

[Compare Source](https://redirect.github.com/nodejs/node/compare/v24.9.0...v24.10.0)

##### Notable Changes

- \[[`31bb476895`](https://redirect.github.com/nodejs/node/commit/31bb476895)] - **(SEMVER-MINOR)** **console**: allow per-stream `inspectOptions` option (Anna Henningsen) [#&#8203;60082](https://redirect.github.com/nodejs/node/pull/60082)
- \[[`3b92be2fb8`](https://redirect.github.com/nodejs/node/commit/3b92be2fb8)] - **(SEMVER-MINOR)** **lib**: remove util.getCallSite (Rafael Gonzaga) [#&#8203;59980](https://redirect.github.com/nodejs/node/pull/59980)
- \[[`18c79d9e1c`](https://redirect.github.com/nodejs/node/commit/18c79d9e1c)] - **(SEMVER-MINOR)** **sqlite**: create authorization api (Guilherme Araújo) [#&#8203;59928]

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQyLjI2LjUtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @codecov-commenter on February 09, 2026 at 05:21 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1449?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.53%. Comparing base ([`4437a55`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/4437a559dd63fd5542cfc62723c58a0341755583?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`ec773fc`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/ec773fc0299f96670af2ca0e86485f5389443ec0?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1449   +/-   ##
=======================================
  Coverage   62.53%   62.53%           
=======================================
  Files         127      127           
  Lines        3310     3310           
  Branches      892      892           
=======================================
  Hits         2070     2070           
  Misses       1120     1120           
  Partials      120      120           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1449?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @github-actions - Approved on December 08, 2025 at 08:38 AM UTC

This PR from Konflux has been automatically approved.

### Review by @swadeley - Approved on February 10, 2026 at 07:47 AM UTC

ACK

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1449*
