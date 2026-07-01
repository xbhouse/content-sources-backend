---
type: pull_request
number: 1256
title: "Build chore(deps): update node.js to v20.19.5"
state: merged
author: red-hat-konflux
created: 2025-10-27T04:15:37Z
updated: 2025-10-31T12:17:21Z
closed: 2025-10-31T09:45:11Z
merged: 2025-10-31T09:45:11Z
base_branch: main
head_branch: konflux/mintmaker/main/node-20.x
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1256
---

# Pull Request #1256: Build chore(deps): update node.js to v20.19.5

**Author**: @red-hat-konflux
**Created**: October 27, 2025 at 04:15 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/node-20.x`

## Description

This PR contains the following updates:

| Package | Update | Change |
|---|---|---|
| [node](https://nodejs.org) ([source](https://redirect.github.com/nodejs/node)) | minor | `v20.17.0` -> `20.19.5` |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>nodejs/node (node)</summary>

### [`v20.19.5`](https://redirect.github.com/nodejs/node/releases/tag/v20.19.5): 2025-09-03, Version 20.19.5 &#x27;Iron&#x27; (LTS), @&#8203;marco-ippolito

[Compare Source](https://redirect.github.com/nodejs/node/compare/v20.19.4...v20.19.5)

##### Notable Changes

- \[[`f5b293ad48`](https://redirect.github.com/nodejs/node/commit/f5b293ad48)] - **doc**: add JonasBa to collaborators (Jonas Badalic) [#&#8203;58355](https://redirect.github.com/nodejs/node/pull/58355)
- \[[`4e6ae787c6`](https://redirect.github.com/nodejs/node/commit/4e6ae787c6)] - **doc**: add puskin to collaborators (Giovanni Bucci) [#&#8203;58308](https://redirect.github.com/nodejs/node/pull/58308)
- \[[`d06db658fc`](https://redirect.github.com/nodejs/node/commit/d06db658fc)] - **doc**: add Filip Skokan to TSC (Rafael Gonzaga) [#&#8203;58499](https://redirect.github.com/nodejs/node/pull/58499)
- \[[`3c6206cac9`](https://redirect.github.com/nodejs/node/commit/3c6206cac9)] - **doc**: add [@&#8203;geeksilva97](https://redirect.github.com/geeksilva97) to collaborators (Edy Silva) [#&#8203;57241](https://redirect.github.com/nodejs/node/pull/57241)

##### Commits

- \[[`ea20403467`](https://redirect.github.com/nodejs/node/commit/ea20403467)] - **build**: fix uvwasi pkgname (Antoine du Hamel) [#&#8203;58270](https://redirect.github.com/nodejs/node/pull/58270)
- \[[`c647aa4b30`](https://redirect.github.com/nodejs/node/commit/c647aa4b30)] - **build**: fix pointer compression builds (Joyee Cheung) [#&#8203;58171](https://redirect.github.com/nodejs/node/pull/58171)
- \[[`d2c5e609ae`](https://redirect.github.com/nodejs/node/commit/d2c5e609ae)] - **build**: disable v8\_enable\_pointer\_compression\_shared\_cage on non-64bit (Shelley Vohr) [#&#8203;58867](https://redirect.github.com/nodejs/node/pull/58867)
- \[[`84d5c4d244`](https://redirect.github.com/nodejs/node/commit/84d5c4d244)] - **build**: search for libnode.so in multiple places (Jan Staněk) [#&#8203;58213](https://redirect.github.com/nodejs/node/pull/58213)
- \[[`068c439552`](https://redirect.github.com/nodejs/node/commit/068c439552)] - **crypto**: fix SHAKE128/256 breaking change introduced with OpenSSL 3.4 (Filip Skokan) [#&#8203;58942](https://redirect.github.com/nodejs/node/pull/58942)
- \[[`edff105c34`](https://redirect.github.com/nodejs/node/commit/edff105c34)] - **debugger**: fix behavior of plain object exec in debugger repl (Dario Piotrowicz) [#&#8203;57498](https://redirect.github.com/nodejs/node/pull/57498)
- \[[`0473e35b7f`](https://redirect.github.com/nodejs/node/commit/0473e35b7f)] - **deps**: update zlib to 1.3.1-470d3a2 (Node.js GitHub Bot) [#&#8203;58628](https://redirect.github.com/nodejs/node/pull/58628)
- \[[`1218dbbea5`](https://redirect.github.com/nodejs/node/commit/1218dbbea5)] - **deps**: update zlib to 1.3.0.1-motley-780819f (Node.js GitHub Bot) [#&#8203;57768](https://redirect.github.com/nodejs/node/pull/57768)
- \[[`0e3cd9ec00`](https://redirect.github.com/nodejs/node/commit/0e3cd9ec00)] - **deps**: update zlib to 1.3.0.1-motley-788cb3c (Node.js GitHub Bot) [#&#8203;56655](https://redirect.github.com/nodejs/node/pull/56655)
- \[[`a194dd9bd4`](https://redirect.github.com/nodejs/node/commit/a194dd9bd4)] - **deps**: update archs files for openssl-3.0.16 (Node.js GitHub Bot) [#&#8203;57335](https://redirect.github.com/nodejs/node/pull/57335)
- \[[`cc9b79ca70`](https://redirect.github.com/nodejs/node/commit/cc9b79ca70)] - **deps**: upgrade openssl sources to quictls/openssl-3.0.16 (Node.js GitHub Bot) [#&#8203;57335](https://redirect.github.com/nodejs/node/pull/57335)
- \[[`82c46d5358`](https://redirect.github.com/nodejs/node/commit/82c46d5358)] - **deps**: update cjs-module-lexer to 2.1.0 (Node.js GitHub Bot) [#&#8203;57180](https://redirect.github.com/nodejs/node/pull/57180)
- \[[`43e3f9b26b`](https://redirect.github.com/nodejs/node/commit/43e3f9b26b)] - **deps**: update cjs-module-lexer to 2.0.0 (Michael Dawson) [#&#8203;56855](https://redirect.github.com/nodejs/node/pull/56855)
- \[[`91282ff16b`](https://redirect.github.com/nodejs/node/commit/91282ff16b)] - **deps**: update corepack to 0.33.0 (Node.js GitHub Bot) [#&#8203;58566](https://redirect.github.com/nodejs/node/pull/58566)
- \[[`b76bca6f38`](https://redirect.github.com/nodejs/node/commit/b76bca6f38)] - **deps**: update acorn to 8.15.0 (Node.js GitHub Bot) [#&#8203;58711](https://redirect.github.com/nodejs/node/pull/58711)
- \[[`ae11481011`](https://redirect.github.com/nodejs/node/commit/ae11481011)] - **deps**: update acorn to 8.14.1 (Node.js GitHub Bot) [#&#8203;57382](https://redirect.github.com/nodejs/node/pull/57382)
- \[[`142d701201`](https://redirect.github.com/nodejs/node/commit/142d701201)] - **deps**: update minimatch to 10.0.3 (Node.js GitHub Bot) [#&#8203;58712](https://redirect.github.com/nodejs/node/pull/58712)
- \[[`fee082d684`](https://redirect.github.com/nodejs/node/commit/fee082d684)] - **deps**: update llhttp to 9.3.0 (Fedor Indutny) [#&#8203;58144](https://redirect.github.com/nodejs/node/pull/58144)
- \[[`c06f6f3f05`](https://redirect.github.com/nodejs/node/commit/c06f6f3f05)] - **dns**: remove redundant code using common variable (Deokjin Kim) [#&#8203;57386](https://redirect.github.com/nodejs/node/pull/57386)
- \[[`cded8e7e77`](https://redirect.github.com/nodejs/node/commit/cded8e7e77)] - **dns**: fix parse memory leaky (theanarkh) [#&#8203;58973](https://redirect.github.com/nodejs/node/pull/58973)
- \[[`182ae67233`](https://redirect.github.com/nodejs/node/commit/182ae67233)] - **dns**: fix dns query cache implementation (Ethan Arrowood) [#&#8203;58404](https://redirect.github.com/nodejs/node/pull/58404)
- \[[`621b66a297`](https://redirect.github.com/nodejs/node/commit/621b66a297)] - **doc**: add review guidelines for collaborator nominations (Antoine du Hamel) [#&#8203;57449](https://redirect.github.com/nodejs/node/pull/57449)
- \[[`b1009b5b72`](https://redirect.github.com/nodejs/node/commit/b1009b5b72)] - **doc**: explicit mention arbitrary code execution as a vuln (Rafael Gonzaga) [#&#8203;57426](https://redirect.github.com/nodejs/node/pull/57426)
- \[[`f5b293ad48`](https://redirect.github.com/nodejs/node/commit/f5b293ad48)] - **doc**: add JonasBa to collaborators (Jonas Badalic) [#&#8203;58355](https://redirect.github.com/nodejs/node/pull/58355)
- \[[`4e6ae787c6`](https://redirect.github.com/nodejs/node/commit/4e6ae787c6)] - **doc**: add puskin to collaborators (Giovanni Bucci) [#&#8203;58308](https://redirect.github.com/nodejs/node/pull/58308)
- \[[`530473f479`](https://redirect.github.com/nodejs/node/commit/530473f479)] - **doc**: add ovflowd back to core collaborators (Claudio W.) [#&#8203;58911](https://redirect.github.com/nodejs/node/pull/58911)
- \[[`38e8bbc131`](https://redirect.github.com/nodejs/node/commit/38e8bbc131)] - **doc**: add info on how project manages social media (Michael Dawson) [#&#8203;57318](https://redirect.github.com/nodejs/node/pull/57318)
- \[[`d06bb4dcc2`](https://redirect.github.com/nodejs/node/commit/d06bb4dcc2)] - **doc**: ping nodejs/tsc for each security pull request (Rafael Gonzaga) [#&#8203;57309](https://redirect.github.com/nodejs/node/pull/57309)
- \[[`d06db658fc`](https://redirect.github.com/nodejs/node/commit/d06db658fc)] - **doc**: add Filip Skokan to TSC (Rafael Gonzaga) [#&#8203;58499](https://redirect.github.com/nodejs/node/pull/58499)
- \[[`8c3bc156ed`](https://redirect.github.com/nodejs/node/commit/8c3bc156ed)] - **doc**: clarify `path.isAbsolute` is not path traversal mitigation (Eric Fortis) [#&#8203;57073](https://redirect.github.com/nodejs/node/pull/57073)
- \[[`e688410bda`](https://redirect.github.com/nodejs/node/commit/e688410bda)] - **doc**: fix rendering of DEP0174 description (David Sanders) [#&#8203;56835](https://redirect.github.com/nodejs/node/pull/56835)
- \[[`e6a0c6a0fa`](https://redirect.github.com/nodejs/node/commit/e6a0c6a0fa)] - **doc**: add missing assert return types (Colin Ihrig) [#&#8203;57219](https://redirect.github.com/nodejs/node/pull/57219)
- \[[`026b3cab6a`](https://redirect.github.com/nodejs/node/commit/026b3cab6a)] - **doc**: add 1ilsang to triage team (1ilsang) [#&#8203;57183](https://redirect.github.com/nodejs/node/pull/57183)
- \[[`3c6206cac9`](https://redirect.github.com/nodejs/node/commit/3c6206cac9)] - **doc**: add [@&#8203;geeksilva97](https://redirect.github.com/geeksilva97) to collaborators (Edy Silva) [#&#8203;57241](https://redirect.github.com/nodejs/node/pull/57241)
- \[[`ef3a4675c7`](https://redirect.github.com/nodejs/node/commit/ef3a4675c7)] - **doc**: fix web.libera.chat link in pull-requests.md (Samuel Bronson) [#&#8203;57076](https://redirect.github.com/nodejs/node/pull/57076)
- \[[`1db42b76f7`](https://redirect.github.com/nodejs/node/commit/1db42b76f7)] - **doc**: remove buffered flag from performance hooks examples (Pavel Romanov) [#&#8203;52607](https://redirect.github.com/nodejs/node/pull/52607)
- \[[`b73a1356ce`](https://redirect.github.com/nodejs/node/commit/b73a1356ce)] - **doc**: add `module namespace object` links (Dario Piotrowicz) [#&#8203;57093](https://redirect.github.com/nodejs/node/pull/57093)
- \[[`09368db20f`](https://redirect.github.com/nodejs/node/commit/09368db20f)] - **doc**: disambiguate pseudo-code statement (Dario Piotrowicz) [#&#8203;57092](https://redirect.github.com/nodejs/node/pull/57092)
- \[[`2c3dc569a1`](https://redirect.github.com/nodejs/node/commit/2c3dc569a1)] - **doc**: fix wrong articles used to address modules (Dario Piotrowicz) [#&#8203;57090](https://redirect.github.com/nodejs/node/pull/57090)
- \[[`cd8259cb4e`](https://redirect.github.com/nodejs/node/commit/cd8259cb4e)] - **doc**: `modules.md`: fix `distance` definition (Alexander “weej” Jones) [#&#8203;57046](https://redirect.github.com/nodejs/node/pull/57046)
- \[[`7b0ea9ab2d`](https://redirect.github.com/nodejs/node/commit/7b0ea9ab2d)] - **doc**: fix wrong verb form (Dario Piotrowicz) [#&#8203;57091](https://redirect.github.com/nodejs/node/pull/57091)
- \[[`14fcfc242b`](https://redirect.github.com/nodejs/node/commit/14fcfc242b)] - **doc**: add a note about `require('../common')` in testing documentation (Aditi) [#&#8203;56953](https://redirect.github.com/nodejs/node/pull/56953)
- \[[`bc7d18b6ea`](https://redirect.github.com/nodejs/node/commit/bc7d18b6ea)] - **doc**: recommend writing tests in new files and including comments (Joyee Cheung) [#&#8203;57028](https://redirect.github.com/nodejs/node/pull/57028)
- \[[`acd4d7f269`](https://redirect.github.com/nodejs/node/commit/acd4d7f269)] - **doc**: improve documentation on argument validation (Aditi) [#&#8203;56954](https://redirect.github.com/nodejs/node/pull/56954)
- \[[`4cd6b3ca73`](https://redirect.github.com/nodejs/node/commit/4cd6b3ca73)] - **doc**: buffer: fix typo on `Buffer.copyBytesFrom(` `offset` option (tpoisseau) [#&#8203;57015](https://redirect.github.com/nodejs/node/pull/57015)
- \[[`01220607f2`](https://redirect.github.com/nodejs/node/commit/01220607f2)] - **doc**: update cleanup to trust on vuln db automation (Rafael Gonzaga) [#&#8203;57004](https://redirect.github.com/nodejs/node/pull/57004)
- \[[`77a0505a32`](https://redirect.github.com/nodejs/node/commit/77a0505a32)] - **doc**: update post sec release process (Rafael Gonzaga) [#&#8203;56907](https://redirect.github.com/nodejs/node/pull/56907)
- \[[`77dbcfce5f`](https://redirect.github.com/nodejs/node/commit/77dbcfce5f)] - **doc**: add section about using npx with permission model (Rafael Gonzaga) [#&#8203;56539](https://redirect.github.com/nodejs/node/pull/56539)
- \[[`73e51407b7`](https://redirect.github.com/nodejs/node/commit/73e51407b7)] - **doc**: remove RedYetiDev from triagers team (Aviv Keller) [#&#8203;55947](https://redirect.github.com/nodejs/node/pull/55947)
- \[[`9a36cbb792`](https://redirect.github.com/nodejs/node/commit/9a36cbb792)] - **doc**: fix relative path mention in --allow-fs (Rafael Gonzaga) [#&#8203;55791](https://redirect.github.com/nodejs/node/pull/55791)
- \[[`04d9c5baeb`](https://redirect.github.com/nodejs/node/commit/04d9c5baeb)] - **doc**: add scroll margin to links (Roman Reiss) [#&#8203;58982](https://redirect.github.com/nodejs/node/pull/58982)
- \[[`959a67f6ff`](https://redirect.github.com/nodejs/node/commit/959a67f6ff)] - **doc**: make Stability labels not sticky in Stability index (Livia Medeiros) [#&#8203;58291](https://redirect.github.com/nodejs/node/pull/58291)
- \[[`8757a5532f`](https://redirect.github.com/nodejs/node/commit/8757a5532f)] - **doc**: update release key for aduh95 (Antoine du Hamel) [#&#8203;58877](https://redirect.github.com/nodejs/node/pull/58877)
- \[[`6fa0626327`](https://redirect.github.com/nodejs/node/commit/6fa0626327)] - **doc,src,test**: fix typos (Noritaka Kobayashi) [#&#8203;58477](https://redirect.github.com/nodejs/node/pull/58477)
- \[[`9991788e4a`](https://redirect.github.com/nodejs/node/commit/9991788e4a)] - **http**: coerce content-length to number (Marco Ippolito) [#&#8203;57458](https://redirect.github.com/nodejs/node/pull/57458)
- \[[`ff5cf8a428`](https://redirect.github.com/nodejs/node/commit/ff5cf8a428)] - **http2**: fix check for `frame->hd.type` (hanguanqiang) [#&#8203;57644](https://redirect.github.com/nodejs/node/pull/57644)
- \[[`2f333b6c51`](https://redirect.github.com/nodejs/node/commit/2f333b6c51)] - **lib**: optimize `prepareStackTrace` on builtin frames (Chengzhong Wu) [#&#8203;56299](https://redirect.github.com/nodejs/node/pull/56299)
- \[[`cdf985071f`](https://redirect.github.com/nodejs/node/commit/cdf985071f)] - **lib**: suppress source map lookup exceptions (Chengzhong Wu) [#&#8203;56299](https://redirect.github.com/nodejs/node/pull/56299)
- \[[`faa08b14ed`](https://redirect.github.com/nodejs/node/commit/faa08b14ed)] - **lib**: fixup incorrect argument order in assertEncoding (James M Snell) [#&#8203;57177](https://redirect.github.com/nodejs/node/pull/57177)
- \[[`a683cd1232`](https://redirect.github.com/nodejs/node/commit/a683cd1232)] - **meta**: add IlyasShabi to collaborators (Ilyas Shabi) [#&#8203;58916](https://redirect.github.com/nodejs/node/pull/58916)
- \[[`b145bb28aa`](https://redirect.github.com/nodejs/node/commit/b145bb28aa)] - **meta**: bump codecov/codecov-action from 5.4.2 to 5.4.3 (dependabot\[bot]) [#&#8203;58551](https://redirect.github.com/nodejs/node/pull/58551)
- \[[`2c59789001`](https://redirect.github.com/nodejs/node/commit/2c59789001)] - **meta**: bump ossf/scorecard-action from 2.4.1 to 2.4.2 (dependabot\[bot]) [#&#8203;58550](https://redirect.github.com/nodejs/node/pull/58550)
- \[[`4095337e96`](https://redirect.github.com/nodejs/node/commit/4095337e96)] - **meta**: bump rtCamp/action-slack-notify from 2.3.2 to 2.3.3 (dependabot\[bot]) [#&#8203;58108](https://redirect.github.com/nodejs/node/pull/58108)
- \[[`631fed8e39`](https://redirect.github.com/nodejs/node/commit/631fed8e39)] - **meta**: move one or more collaborators to emeritus (Node.js GitHub Bot) [#&#8203;58456](https://redirect.github.com/nodejs/node/pull/58456)
- \[[`7d2f7180b6`](https://redirect.github.com/nodejs/node/commit/7d2f7180b6)] - **meta**: bump codecov/codecov-action from 5.4.0 to 5.4.2 (dependabot\[bot]) [#&#8203;58110](https://redirect.github.com/nodejs/node/pull/58110)
- \[[`1558551ea5`](https://redirect.github.com/nodejs/node/commit/1558551ea5)] - **meta**: bump actions/download-artifact from 4.2.1 to 4.3.0 (dependabot\[bot]) [#&#8203;58106](https://redirect.github.com/nodejs/node/pull/58106)
- \[[`e1f12fe737`](https://redirect.github.com/nodejs/node/commit/e1f12fe737)] - **meta**: ignore mailmap changes in linux ci (Jonas Badalic) [#&#8203;58356](https://redirect.github.com/nodejs/node/pull/58356)
- \[[`1b78eb1313`](https://redirect.github.com/nodejs/node/commit/1b78eb1313)] - **meta**: bump actions/setup-node from 4.3.0 to 4.4.0 (dependabot\[bot]) [#&#8203;58111](https://redirect.github.com/nodejs/node/pull/58111)
- \[[`2b8449c39a`](https://redirect.github.com/nodejs/node/commit/2b8449c39a)] - **meta**: bump actions/setup-python from 5.5.0 to 5.6.0 (dependabot\[bot]) [#&#8203;58107](https://redirect.github.com/nodejs/node/pull/58107)
- \[[`833b70bbc5`](https://redirect.github.com/nodejs/node/commit/833b70bbc5)] - **meta**: allow penetration testing on live system with prior authorization (Matteo Collina) [#&#8203;57966](https://redirect.github.com/nodejs/node/pull/57966)
- \[[`c6a88561f5`](https://redirect.github.com/nodejs/node/commit/c6a88561f5)] - **meta**: bump actions/setup-python from 5.4.0 to 5.5.0 (dependabot\[bot]) [#&#8203;57718](https://redirect.github.com/nodejs/node/pull/57718)
- \[[`9046ef4fb3`](https://redirect.github.com/nodejs/node/commit/9046ef4fb3)] - **meta**: bump peter-evans/create-pull-request from 7.0.7 to 7.0.8 (dependabot\[bot]) [#&#8203;57717](https://redirect.github.com/nodejs/node/pull/57717)
- \[[`46388a4e2a`](https://redirect.github.com/nodejs/node/commit/46388a4e2a)] - **meta**: bump actions/cache from 4.2.2 to 4.2.3 (dependabot\[bot]) [#&#8203;57715](https://redirect.github.com/nodejs/node/pull/57715)
- \[[`d3970685bd`](https://redirect.github.com/nodejs/node/commit/d3970685bd)] - **meta**: bump actions/setup-node from 4.2.0 to 4.3.0 (dependabot\[bot]) [#&#8203;57714](https://redirect.github.com/nodejs/node/pull/57714)
- \[[`47004ef37f`](https://redirect.github.com/nodejs/node/commit/47004ef37f)] - **meta**: bump actions/upload-artifact from 4.6.1 to 4.6.2 (dependabot\[bot]) [#&#8203;57713](https://redirect.github.com/nodejs/node/pull/57713)
- \[[`4abe83ec03`](https://redirect.github.com/nodejs/node/commit/4abe83ec03)] - **meta**: add some clarification to the nomination process (James M Snell) [#&#8203;57503](https://redirect.github.com/nodejs/node/pull/57503)
- \[[`45e9b88363`](https://redirect.github.com/nodejs/node/commit/45e9b88363)] - **meta**: remove collaborator self-nomination (Rich Trott) [#&#8203;57537](https://redirect.github.com/nodejs/node/pull/57537)
- \[[`d10949b7d8`](https://redirect.github.com/nodejs/node/commit/d10949b7d8)] - **meta**: edit collaborator nomination process (Antoine du Hamel) [#&#8203;57483](https://redirect.github.com/nodejs/node/pull/57483)
- \[[`704562fb7a`](https://redirect.github.com/nodejs/node/commit/704562fb7a)] - **meta**: move ovflowd to emeritus (Claudio W.) [#&#8203;57443](https://redirect.github.com/nodejs/node/pull/57443)
- \[[`3f981b8537`](https://redirect.github.com/nodejs/node/commit/3f981b8537)] - **meta**: bump codecov/codecov-action from 5.3.1 to 5.4.0 (dependabot\[bot]) [#&#8203;57257](https://redirect.github.com/nodejs/node/pull/57257)
- \[[`7e1ff7b332`](https://redirect.github.com/nodejs/node/commit/7e1ff7b332)] - **meta**: bump ossf/scorecard-action from 2.4.0 to 2.4.1 (dependabot\[bot]) [#&#8203;57253](https://redirect.github.com/nodejs/node/pull/57253)
- \[[`8d4ec412b9`](https://redirect.github.com/nodejs/node/commit/8d4ec412b9)] - **meta**: move RaisinTen back to collaborators, triagers and SEA champion (Darshan Sen) [#&#8203;57292](https://redirect.github.com/nodejs/node/pull/57292)
- \[[`cc2abb5d17`](https://redirect.github.com/nodejs/node/commit/cc2abb5d17)] - **meta**: bump peter-evans/create-pull-request from 7.0.6 to 7.0.7 (dependabot\[bot]) [#&#8203;57259](https://redirect.github.com/nodejs/node/pull/57259)
- \[[`4fad2b8758`](https://redirect.github.com/nodejs/node/commit/4fad2b8758)] - **meta**: bump actions/cache from 4.2.0 to 4.2.2 (dependabot\[bot]) [#&#8203;57256](https://redirect.github.com/nodejs/node/pull/57256)
- \[[`5f5bb8b986`](https://redirect.github.com/nodejs/node/commit/5f5bb8b986)] - **meta**: bump actions/upload-artifact from 4.6.0 to 4.6.1 (dependabot\[bot]) [#&#8203;57255](https://redirect.github.com/nodejs/node/pull/57255)
- \[[`e949359a56`](https://redirect.github.com/nodejs/node/commit/e949359a56)] - **meta**: bump `actions/setup-python` from 5.3.0 to 5.4.0 (dependabot\[bot]) [#&#8203;56867](https://redirect.github.com/nodejs/node/pull/56867)
- \[[`d3c5ad7510`](https://redirect.github.com/nodejs/node/commit/d3c5ad7510)] - **meta**: bump `peter-evans/create-pull-request` from 7.0.5 to 7.0.6 (dependabot\[bot]) [#&#8203;56866](https://redirect.github.com/nodejs/node/pull/56866)
- \[[`56decfe2d1`](https://redirect.github.com/nodejs/node/commit/56decfe2d1)] - **meta**: bump `codecov/codecov-action` from 5.0.7 to 5.3.1 (dependabot\[bot]) [#&#8203;56864](https://redirect.github.com/nodejs/node/pull/56864)
- \[[`52e518444d`](https://redirect.github.com/nodejs/node/commit/52e518444d)] - **meta**: bump `actions/cache` from 4.1.2 to 4.2.0 (dependabot\[bot]) [#&#8203;56862](https://redirect.github.com/nodejs/node/pull/56862)
- \[[`9cac93d9c3`](https://redirect.github.com/nodejs/node/commit/9cac93d9c3)] - **meta**: bump `actions/stale` from 9.0.0 to 9.1.0 (dependabot\[bot]) [#&#8203;56860](https://redirect.github.com/nodejs/node/pull/56860)
- \[[`ecf4252f7c`](https://redirect.github.com/nodejs/node/commit/ecf4252f7c)] - **meta**: update last name for jkrems (Jan Martin) [#&#8203;57006](https://redirect.github.com/nodejs/node/pull/57006)
- \[[`e8beaaaedf`](https://redirect.github.com/nodejs/node/commit/e8beaaaedf)] - **meta**: bump `actions/upload-artifact` from 4.4.3 to 4.6.0 (dependabot\[bot]) [#&#8203;56861](https://redirect.github.com/nodejs/node/pull/56861)
- \[[`5462c257f8`](https://redirect.github.com/nodejs/node/commit/5462c257f8)] - **meta**: bump `actions/setup-node` from 4.1.0 to 4.2.0 (dependabot\[bot]) [#&#8203;56868](https://redirect.github.com/nodejs/node/pull/56868)
- \[[`89c37891a0`](https://redirect.github.com/nodejs/node/commit/89c37891a0)] - **meta**: move one or more collaborators to emeritus (Node.js GitHub Bot) [#&#8203;56889](https://redirect.github.com/nodejs/node/pull/56889)
- \[[`2a0175c291`](https://redirect.github.com/nodejs/node/commit/2a0175c291)] - **meta**: add [@&#8203;nodejs/url](https://redirect.github.com/nodejs/url) as codeowner (Chengzhong Wu) [#&#8203;56783](https://redirect.github.com/nodejs/node/pull/56783)
- \[[`c12aae1e78`](https://redirect.github.com/nodejs/node/commit/c12aae1e78)] - **meta**: bump github/codeql-action from 3.28.18 to 3.29.2 (dependabot\[bot]) [#&#8203;58922](https://redirect.github.com/nodejs/node/pull/58922)
- \[[`4ef09990f1`](https://redirect.github.com/nodejs/node/commit/4ef09990f1)] - **meta**: bump github/codeql-action from 3.28.16 to 3.28.18 (dependabot\[bot]) [#&#8203;58552](https://redirect.github.com/nodejs/node/pull/58552)
- \[[`889654eb2c`](https://redirect.github.com/nodejs/node/commit/889654eb2c)] - **meta**: bump github/codeql-action from 3.28.11 to 3.28.16 (dependabot\[bot]) [#&#8203;58112](https://redirect.github.com/nodejs/node/pull/58112)
- \[[`091e5c1bb9`](https://redirect.github.com/nodejs/node/commit/091e5c1bb9)] - **meta**: bump github/codeql-action from 3.28.10 to 3.28.13 (dependabot\[bot]) [#&#8203;57716](https://redirect.github.com/nodejs/node/pull/57716)
- \[[`01415153de`](https://redirect.github.com/nodejs/node/commit/01415153de)] - **meta**: bump github/codeql-action from 3.28.8 to 3.28.10 (dependabot\[bot]) [#&#8203;57254](https://redirect.github.com/nodejs/node/pull/57254)
- \[[`72ea8aac34`](https://redirect.github.com/nodejs/node/commit/72ea8aac34)] - **meta**: bump `github/codeql-action` from 3.27.5 to 3.28.8 (dependabot\[bot]) [#&#8203;56859](https://redirect.github.com/nodejs/node/pull/56859)
- \[[`99a271e588`](https://redirect.github.com/nodejs/node/commit/99a271e588)] - **meta**: bump step-security/harden-runner from 2.12.0 to 2.12.2 (dependabot\[bot]) [#&#8203;58923](https://redirect.github.com/nodejs/node/pull/58923)
- \[[`b4c4c02490`](https://redirect.github.com/nodejs/node/commit/b4c4c02490)] - **meta**: bump step-security/harden-runner from 2.11.0 to 2.12.0 (dependabot\[bot]) [#&#8203;58109](https://redirect.github.com/nodejs/node/pull/58109)
- \[[`5361bb9157`](https://redirect.github.com/nodejs/node/commit/5361bb9157)] - **meta**: bump step-security/harden-runner from 2.10.4 to 2.11.0 (dependabot\[bot]) [#&#8203;57258](https://redirect.github.com/nodejs/node/pull/57258)
- \[[`28e33acf30`](https://redirect.github.com/nodejs/node/commit/28e33acf30)] - **meta**: bump `step-security/harden-runner` from 2.10.2 to 2.10.4 (dependabot\[bot]) [#&#8203;56863](https://redirect.github.com/nodejs/node/pull/56863)
- \[[`fad773cede`](https://redirect.github.com/nodejs/node/commit/fad773cede)] - **module**: throw error when re-runing errored module jobs (Joyee Cheung) [#&#8203;58957](https://redirect.github.com/nodejs/node/pull/58957)
- \[[`2531185423`](https://redirect.github.com/nodejs/node/commit/2531185423)] - **module**: allow cycles in require() in the CJS handling in ESM loader (Joyee Cheung) [#&#8203;58598](https://redirect.github.com/nodejs/node/pull/58598)
- \[[`ed43b69689`](https://redirect.github.com/nodejs/node/commit/ed43b69689)] - **module**: clarify cjs global-like error on ModuleJobSync (Carlos Espa) [#&#8203;56491](https://redirect.github.com/nodejs/node/pull/56491)
- \[[`6e02db1b12`](https://redirect.github.com/nodejs/node/commit/6e02db1b12)] - **module**: handle instantiated async module jobs in require(esm) (Joyee Cheung) [#&#8203;58067](https://redirect.github.com/nodejs/node/pull/58067)
- \[[`badba50d30`](https://redirect.github.com/nodejs/node/commit/badba50d30)] - **module**: fix incorrect formatting in require(esm) cycle error message (haykam821) [#&#8203;57453](https://redirect.github.com/nodejs/node/pull/57453)
- \[[`939ecf8906`](https://redirect.github.com/nodejs/node/commit/939ecf8906)] - **module**: handle cached linked async jobs in require(esm) (Joyee Cheung) [#&#8203;57187](https://redirect.github.com/nodejs/node/pull/57187)
- \[[`ba7f8a0353`](https://redirect.github.com/nodejs/node/commit/ba7f8a0353)] - **module**: improve error message from asynchronicity in require(esm) (Joyee Cheung) [#&#8203;57126](https://redirect.github.com/nodejs/node/pull/57126)
- \[[`c1e7fa2586`](https://redirect.github.com/nodejs/node/commit/c1e7fa2586)] - **module**: handle .mjs in .js handler in CommonJS (Joyee Cheung) [#&#8203;55590](https://redirect.github.com/nodejs/node/pull/55590)
- \[[`41f3dfd21b`](https://redirect.github.com/nodejs/node/commit/41f3dfd21b)] - **module**: fix require.resolve() crash on non-string paths (Aditi) [#&#8203;56942](https://redirect.github.com/nodejs/node/pull/56942)
- \[[`043dcdd628`](https://redirect.github.com/nodejs/node/commit/043dcdd628)] - **os**: fix GetInterfaceAddresses memory lieaky (theanarkh) [#&#8203;58940](https://redirect.github.com/nodejs/node/pull/58940)
- \[[`9b74e9bfd9`](https://redirect.github.com/nodejs/node/commit/9b74e9bfd9)] - **permission**: ignore internalModuleStat on module loading (Rafael Gonzaga) [#&#8203;55797](https://redirect.github.com/nodejs/node/pull/55797)
- \[[`611a147b45`](https://redirect.github.com/nodejs/node/commit/611a147b45)] - **readline**: fix unresolved promise on abortion (Daniel Venable) [#&#8203;54030](https://redirect.github.com/nodejs/node/pull/54030)
- \[[`f891ae3421`](https://redirect.github.com/nodejs/node/commit/f891ae3421)] - **repl**: avoid deprecated `require.extensions` in tab completion (baki gul) [#&#8203;58653](https://redirect.github.com/nodejs/node/pull/58653)
- \[[`7ba44290bf`](https://redirect.github.com/nodejs/node/commit/7ba44290bf)] - **repl**: fix tab completion not working with computer string properties (Dario Piotrowicz) [#&#8203;58709](https://redirect.github.com/nodejs/node/pull/58709)
- \[[`eb842048b2`](https://redirect.github.com/nodejs/node/commit/eb842048b2)] - **src**: do not format single string argument for THROW\_ERR\_\* (Joyee Cheung) [#&#8203;57126](https://redirect.github.com/nodejs/node/pull/57126)
- \[[`4f004937ec`](https://redirect.github.com/nodejs/node/commit/4f004937ec)] - **src**: fixup errorhandling more in various places (James M Snell) [#&#8203;57852](https://redirect.github.com/nodejs/node/pull/57852)
- \[[`5daa7fe2e2`](https://redirect.github.com/nodejs/node/commit/5daa7fe2e2)] - **src**: fix module buffer allocation (X-BW) [#&#8203;57738](https://redirect.github.com/nodejs/node/pull/57738)
- \[[`586b1be11b`](https://redirect.github.com/nodejs/node/commit/586b1be11b)] - **src**: fix build when using shared simdutf (Antoine du Hamel) [#&#8203;58407](https://redirect.github.com/nodejs/node/pull/58407)
- \[[`563e61f012`](https://redirect.github.com/nodejs/node/commit/563e61f012)] - **src**: fix possible dereference of null pointer (Eusgor) [#&#8203;58459](https://redirect.github.com/nodejs/node/pull/58459)
- \[[`cbec07ea0b`](https://redirect.github.com/nodejs/node/commit/cbec07ea0b)] - **src**: fix FIPS init error handling (Tobias Nießen) [#&#8203;58379](https://redirect.github.com/nodejs/node/pull/58379)
- \[[`80fb80e71b`](https://redirect.github.com/nodejs/node/commit/80fb80e71b)] - **src**: fix -Wunreachable-code in src/node\_api.cc (Shelley Vohr) [#&#8203;58901](https://redirect.github.com/nodejs/node/pull/58901)
- \[[`5e97719860`](https://redirect.github.com/nodejs/node/commit/5e97719860)] - **test**: skip test-http-imports on macos (Marco Ippolito) [#&#8203;59745](https://redirect.github.com/nodejs/node/pull/59745)
- \[[`69c43bdfcc`](https://redirect.github.com/nodejs/node/commit/69c43bdfcc)] - **test**: fix internet/test-dns (Michaël Zasso) [#&#8203;59660](https://redirect.github.com/nodejs/node/pull/59660)
- \[[`6fd58e0338`](https://redirect.github.com/nodejs/node/commit/6fd58e0338)] - **tools**: update coverage GitHub Actions to fixed version (Rich Trott) [#&#8203;59512](https://redirect.github.com/nodejs/node/pull/59512)
- \[[`eb7bbce73e`](https://redirect.github.com/nodejs/node/commit/eb7bbce73e)] - **tools**: disable failing coverage jobs (Antoine du Hamel) [#&#8203;58770](https://redirect.github.com/nodejs/node/pull/58770)
- \[[`65b1669936`](https://redirect.github.com/nodejs/node/commit/65b1669936)] - **util**: fix formatting of objects with built-in Symbol.toPrimitive (Shima Ryuhei) [#&#8203;57832](https://redirect.github.com/nodejs/node/pull/57832)
- \[[`8a29f13bec`](https://redirect.github.com/nodejs/node/commit/8a29f13bec)] - **util**: fix parseEnv incorrectly splitting multiple ‘=‘ in value (HEESEUNG) [#&#8203;57421](https://redirect.github.com/nodejs/node/pull/57421)
- \[[`077d5020c4`](https://redirect.github.com/nodejs/node/commit/077d5020c4)] - **v8**: fix missing callback in heap utils destroy (Ruben Bridgewater) [#&#8203;58846](https://redirect.github.com/nodejs/node/pull/58846)
- \[[`34ae9f8b18`](https://redirect.github.com/nodejs/node/commit/34ae9f8b18)] - **vm**: import call should return a promise in the current context (Chengzhong Wu) [#&#8203;58309](https://redirect.github.com/nodejs/node/pull/58309)
- \[[`0dd3a8d6d1`](https://redirect.github.com/nodejs/node/commit/0dd3a8d6d1)] - **win,build**: fix MSVS v17.14 compilation issue (StefanStojanovic) [#&#8203;58902](https://redirect.github.com/nodejs/node/pull/58902)
- \[[`1b83a2bd2d`](https://redirect.github.com/nodejs/node/commit/1b83a2bd2d)] - **zlib**: remove mentions of unexposed Z\_TREES constant (Jimmy Leung) [#&#8203;58371](https://redirect.github.com/nodejs/node/pull/58371)
- \[[`9dc9604502`](https://redirect.github.com/nodejs/node/commit/9dc9604502)] - **zlib**: fix pointer alignment (jhofstee) [#&#8203;57727](https://redirect.github.com/nodejs/node/pull/57727)

### [`v20.19.4`](https://redirect.github.com/nodejs/node/releases/tag/v20.19.4): 2025-07-15, Version 20.19.4 &#x27;Iron&#x27; (LTS), @&#8203;RafaelGSS

[Compare Source](https://redirect.github.com/nodejs/node/compare/v20.19.3...v20.19.4)

This is a security release.

##### Notable Changes

- (CVE-2025-27210) Windows Device Names (CON, PRN, AUX) Bypass Path Traversal Protection in path.normalize()

##### Commits

- \[[`db7b93fcef`](https://redirect.github.com/nodejs/node/commit/db7b93fcef)] - **(CVE-2025-27210)** **lib**: handle all windows reserved driver name (RafaelGSS) [nodejs-private/node-private#721](https://redirect.github.com/nodejs-private/node-private/pull/721)

### [`v20.19.3`](https://redirect.github.com/nodejs/node/releases/tag/v20.19.3): 2025-06-23, Version 20.19.3 &#x27;Iron&#x27; (LTS), @&#8203;marco-ippolito

[Compare Source](https://redirect.github.com/nodejs/node/compare/v20.19.2...v20.19.3)

##### Notable Changes

- \[[`c535a3c483`](https://redirect.github.com/nodejs/node/commit/c535a3c483)] - **crypto**: graduate WebCryptoAPI [`Ed25519`](https://redirect.github.com/nodejs/node/commit/Ed25519) and X25519 algorithms as stable (Filip Skokan) [#&#8203;56142](https://redirect.github.com/nodejs/node/pull/56142)
- \[[`af1dc63815`](https://redirect.github.com/nodejs/node/commit/af1dc63815)] - **crypto**: update root certificates to NSS 3.108 (Node.js GitHub Bot) [#&#8203;57381](https://redirect.github.com/nodejs/node/pull/57381)
- \[[`01d63a4ddf`](https://redirect.github.com/nodejs/node/commit/01d63a4ddf)] - **deps**: update timezone to 2025b (Node.js GitHub Bot) [#&#8203;57857](https://redirect.github.com/nodejs/node/pull/57857)
- \[[`b6daa344eb`](https://redirect.github.com/nodejs/node/commit/b6daa344eb)] - **doc**: add dario-piotrowicz to collaborators (Dario Piotrowicz) [#&#8203;58102](https://redirect.github.com/nodejs/node/pull/58102)

##### Commits

- \[[`fc1fa7a357`](https://redirect.github.com/nodejs/node/commit/fc1fa7a357)] - **build**: use FILE\_OFFSET\_BITS=64 esp. on 32-bit arch (RafaelGSS) [#&#8203;58090](https://redirect.github.com/nodejs/node/pull/58090)
- \[[`79e0812181`](https://redirect.github.com/nodejs/node/commit/79e0812181)] - **build**: use glob for dependencies of out/Makefile (Richard Lau) [#&#8203;55789](https://redirect.github.com/nodejs/node/pull/55789)
- \[[`f56e62851a`](https://redirect.github.com/nodejs/node/commit/f56e62851a)] - **crypto**: allow length=0 for HKDF and PBKDF2 in SubtleCrypto.deriveBits (Filip Skokan) [#&#8203;55866](https://redirect.github.com/nodejs/node/pull/55866)
- \[[`c535a3c483`](https://redirect.github.com/nodejs/node/commit/c535a3c483)] - **crypto**: graduate WebCryptoAPI [`Ed25519`](https://redirect.github.com/nodejs/node/commit/Ed25519) and X25519 algorithms as stable (Filip Skokan) [#&#8203;56142](https://redirect.github.com/nodejs/node/pull/56142)
- \[[`39925de8b1`](https://redirect.github.com/nodejs/node/commit/39925de8b1)] - **crypto**: allow non-multiple of 8 in SubtleCrypto.deriveBits (Filip Skokan) [#&#8203;55296](https://redirect.github.com/nodejs/node/pull/55296)
- \[[`af1dc63815`](https://redirect.github.com/nodejs/node/commit/af1dc63815)] - **crypto**: update root certificates to NSS 3.108 (Node.js GitHub Bot) [#&#8203;57381](https://redirect.github.com/nodejs/node/pull/57381)
- \[[`d09008add3`](https://redirect.github.com/nodejs/node/commit/d09008add3)] - **deps**: V8: cherry-pick [`1a3ecc2`](https://redirect.github.com/nodejs/node/commit/1a3ecc2483b2) (Michaël Zasso) [#&#8203;58342](https://redirect.github.com/nodejs/node/pull/58342)
- \[[`fd56652425`](https://redirect.github.com/nodejs/node/commit/fd56652425)] - **deps**: V8: cherry-pick [`182d9c0`](https://redirect.github.com/nodejs/node/commit/182d9c05e78b) (Andrey Kosyakov) [#&#8203;58342](https://redirect.github.com/nodejs/node/pull/58342)
- \[[`447481e829`](https://redirect.github.com/nodejs/node/commit/447481e829)] - **deps**: V8: cherry-pick third\_party/zlib@[`646b7f5`](https://redirect.github.com/nodejs/node/commit/646b7f569718) (Hans Wennborg) [#&#8203;58342](https://redirect.github.com/nodejs/node/pull/58342)
- \[[`eb447168df`](https://redirect.github.com/nodejs/node/commit/eb447168df)] - **deps**: update simdutf to 6.4.2 (Node.js GitHub Bot) [#&#8203;57855](https://redirect.github.com/nodejs/node/pull/57855)
- \[[`01d63a4ddf`](https://redirect.github.com/nodejs/node/commit/01d63a4ddf)] - **deps**: update timezone to 2025b (Node.js GitHub Bot) [#&#8203;57857](https://redirect.github.com/nodejs/node/pull/57857)
- \[[`10fb49f2a9`](https://redirect.github.com/nodejs/node/commit/10fb49f2a9)] - **deps**: update icu to 77.1 (Node.js GitHub Bot) [#&#8203;57455](https://redirect.github.com/nodejs/node/pull/57455)
- \[[`f1dc7d0205`](https://redirect.github.com/nodejs/node/commit/f1dc7d0205)] - **deps**: update corepack to 0.32.0 (Node.js GitHub Bot) [#&#8203;57265](https://redirect.github.com/nodejs/node/pull/57265)
- \[[`7a2e64bb8a`](https://redirect.github.com/nodejs/node/commit/7a2e64bb8a)] - **deps**: update simdutf to 6.4.0 (Node.js GitHub Bot) [#&#8203;56764](https://redirect.github.com/nodejs/node/pull/56764)
- \[[`e80669be0d`](https://redirect.github.com/nodejs/node/commit/e80669be0d)] - **doc**: mention reports should align with Node.js CoC (Rafael Gonzaga) [#&#8203;57607](https://redirect.github.com/nodejs/node/pull/57607)
- \[[`7b2c0bc92e`](https://redirect.github.com/nodejs/node/commit/7b2c0bc92e)] - **doc**: add gurgunday as triager (Gürgün Dayıoğlu) [#&#8203;57594](https://redirect.github.com/nodejs/node/pull/57594)
- \[[`791e4879de`](https://redirect.github.com/nodejs/node/commit/791e4879de)] - **doc**: document REPL custom eval arguments (Dario Piotrowicz) [#&#8203;57690](https://redirect.github.com/nodejs/node/pull/57690)
- \[[`2917f09876`](https://redirect.github.com/nodejs/node/commit/2917f09876)] - **doc**: improved fetch docs (Alessandro Miliucci) [#&#8203;57296](https://redirect.github.com/nodejs/node/pull/57296)
- \[[`d940b15843`](https://redirect.github.com/nodejs/node/commit/d940b15843)] - **doc**: clarify `unhandledRejection` events behaviors in process doc (Dario Piotrowicz) [#&#8203;57654](https://redirect.github.com/nodejs/node/pull/57654)
- \[[`71c664fab7`](https://redirect.github.com/nodejs/node/commit/71c664fab7)] - **doc**: update position type to integer | null in fs (Yukihiro Hasegawa) [#&#8203;57745](https://redirect.github.com/nodejs/node/pull/57745)
- \[[`0c0fbfa9c6`](https://redirect.github.com/nodejs/node/commit/0c0fbfa9c6)] - **doc**: add missing v0.x changelog entries (Antoine du Hamel) [#&#8203;57779](https://redirect.github.com/nodejs/node/pull/57779)
- \[[`e99462c9fc`](https://redirect.github.com/nodejs/node/commit/e99462c9fc)] - **doc**: correct deprecation type of `assert.CallTracker` (René) [#&#8203;57997](https://redirect.github.com/nodejs/node/pull/57997)
- \[[`c7e92696ef`](https://redirect.github.com/nodejs/node/commit/c7e92696ef)] - **doc**: add returns for https.get (Eng Zer Jun) [#&#8203;58025](https://redirect.github.com/nodejs/node/pull/58025)
- \[[`ccc42b69ce`](https://redirect.github.com/nodejs/node/commit/ccc42b69ce)] - **doc**: fix env variable name in `util.styleText` (Antoine du Hamel) [#&#8203;58072](https://redirect.github.com/nodejs/node/pull/58072)
- \[[`b6daa344eb`](https://redirect.github.com/nodejs/node/commit/b6daa344eb)] - **doc**: add dario-piotrowicz to collaborators (Dario Piotrowicz) [#&#8203;58102](https://redirect.github.com/nodejs/node/pull/58102)
- \[[`e5d6a3df16`](https://redirect.github.com/nodejs/node/commit/e5d6a3df16)] - **doc**: fix `AsyncLocalStorage` example response changes after node v18 (Naor Tedgi (Abu Emma)) [#&#8203;57969](https://redirect.github.com/nodejs/node/pull/57969)
- \[[`f006411998`](https://redirect.github.com/nodejs/node/commit/f006411998)] - **doc**: fix typo of file `zlib.md` (yusheng chen) [#&#8203;58093](https://redirect.github.com/nodejs/node/pull/58093)
- \[[`5193735df4`](https://redirect.github.com/nodejs/node/commit/5193735df4)] - **doc**: add missing options.signal to readlinePromises.createInterface() (Jimmy Leung) [#&#8203;55456](https://redirect.github.com/nodejs/node/pull/55456)
- \[[`fd44af730f`](https://redirect.github.com/nodejs/node/commit/fd44af730f)] - **doc**: fix misaligned options in vm.compileFunction() (Jimmy Leung) [#&#8203;58145](https://redirect.github.com/nodejs/node/pull/58145)
- \[[`0fdcc0ddcd`](https://redirect.github.com/nodejs/node/commit/0fdcc0ddcd)] - **doc**: add ambassaor message (Brian Muenzenmeyer) [#&#8203;57600](https://redirect.github.com/nodejs/node/pull/57600)
- \[[`5ca9616bd3`](https://redirect.github.com/nodejs/node/commit/5ca9616bd3)] - **doc**: increase z-index of header element (Dario Piotrowicz) [#&#8203;57851](https://redirect.github.com/nodejs/node/pull/57851)
- \[[`81342d10f0`](https://redirect.github.com/nodejs/node/commit/81342d10f0)] - **doc**: fix deprecation type for `DEP0148` (Livia Medeiros) [#&#8203;57785](https://redirect.github.com/nodejs/node/pull/57785)
- \[[`776becfe01`](https://redirect.github.com/nodejs/node/commit/776becfe01)] - **doc**: remove mention of `--require` not supporting ES modules (Huáng Jùnliàng) [#&#8203;57620](https://redirect.github.com/nodejs/node/pull/57620)
- \[[`3140a8f133`](https://redirect.github.com/nodejs/node/commit/3140a8f133)] - **doc**: add missing `deprecated` badges in `fs.md` (Yukihiro Hasegawa) [#&#8203;57384](https://redirect.github.com/nodejs/node/pull/57384)
- \[[`441ce24ae3`](https://redirect.github.com/nodejs/node/commit/441ce24ae3)] - **doc**: deprecate passing invalid types in `fs.existsSync` (Carlos Espa) [#&#8203;55892](https://redirect.github.com/nodejs/node/pull/55892)
- \[[`0556f54544`](https://redirect.github.com/nodejs/node/commit/0556f54544)] - **http**: correctly translate HTTP method (Paolo Insogna) [#&#8203;52701](https://redirect.github.com/nodejs/node/pull/52701)
- \[[`c2c6d2b035`](https://redirect.github.com/nodejs/node/commit/c2c6d2b035)] - **http**: be more generational GC friendly (ywave620) [#&#8203;56767](https://redirect.github.com/nodejs/node/pull/56767)
- \[[`cdf3fa241c`](https://redirect.github.com/nodejs/node/commit/cdf3fa241c)] - **http2**: skip writeHead if stream is closed (Shima Ryuhei) [#&#8203;57686](https://redirect.github.com/nodejs/node/pull/57686)
- \[[`bbd5aec785`](https://redirect.github.com/nodejs/node/commit/bbd5aec785)] - **http2**: fix graceful session close (Kushagra Pandey) [#&#8203;57808](https://redirect.github.com/nodejs/node/pull/57808)
- \[[`b427ae4f34`](https://redirect.github.com/nodejs/node/commit/b427ae4f34)] - **meta**: remove `build-windows.yml` (Aviv Keller) [#&#8203;54662](https://redirect.github.com/nodejs/node/pull/54662)
- \[[`49e624f554`](https://redirect.github.com/nodejs/node/commit/49e624f554)] - **os**: fix netmask format check condition in getCIDR function (Wiyeong Seo) [#&#8203;57324](https://redirect.github.com/nodejs/node/pull/57324)
- \[[`d582954434`](https://redirect.github.com/nodejs/node/commit/d582954434)] - **src**: remove unused variable in crypto\_x509.cc (Michaël Zasso) [#&#8203;57754](https://redirect.github.com/nodejs/node/pull/57754)
- \[[`234a505e96`](https://redirect.github.com/nodejs/node/commit/234a505e96)] - **src**: allow embedder customization of OOMErrorHandler (Shelley Vohr) [#&#8203;57325](https://redirect.github.com/nodejs/node/pull/57325)
- \[[`c0252cd380`](https://redirect.github.com/nodejs/node/commit/c0252cd380)] - **src**: fix -Wunreachable-code-return in node\_sea (Shelley Vohr) [#&#8203;57664](https://redirect.github.com/nodejs/node/pull/57664)
- \[[`fcd1622fc1`](https://redirect.github.com/nodejs/node/commit/fcd1622fc1)] - **src**: fix kill signal 0 on Windows (Stefan Stojanovic) [#&#8203;57695](https://redirect.github.com/nodejs/node/pull/57695)
- \[[`850192b06b`](https://redirect.github.com/nodejs/node/commit/850192b06b)] - **test**: skip broken sea on rhel8 (Marco Ippolito) [#&#8203;58761](https://redirect.github.com/nodejs/node/pull/58761)
- \[[`3cf7cfb695`](https://redirect.github.com/nodejs/node/commit/3cf7cfb695)] - **test**: update WPT for WebCryptoAPI to [`edd42c0`](https://redirect.github.com/nodejs/node/commit/edd42c005c) (Node.js GitHub Bot) [#&#8203;57365](https://redirect.github.com/nodejs/node/pull/57365)
- \[[`f57765bdcf`](https://redirect.github.com/nodejs/node/commit/f57765bdcf)] - **test**: mark test-without-async-context-frame flaky on windows (James M Snell) [#&#8203;56753](https://redirect.github.com/nodejs/node/pull/56753)
- \[[`275ea8e7ef`](https://redirect.github.com/nodejs/node/commit/275ea8e7ef)] - **test**: force GC in test-file-write-stream4 (Luigi Pinca) [#&#8203;57930](https://redirect.github.com/nodejs/node/pull/57930)
- \[[`da6a13c338`](https://redirect.github.com/nodejs/node/commit/da6a13c338)] - **test**: deflake test-http2-options-max-headers-block-length (Luigi Pinca) [#&#8203;57959](https://redirect.github.com/nodejs/node/pull/57959)
- \[[`56fce6691e`](https://redirect.github.com/nodejs/node/commit/56fce6691e)] - **test**: prevent extraneous HOSTNAME substitution in test-runner-output (René) [#&#8203;58076](https://redirect.github.com/nodejs/node/pull/58076)
- \[[`c9c0be5596`](https://redirect.github.com/nodejs/node/commit/c9c0be5596)] - **test**: update expected error message for macOS (Antoine du Hamel) [#&#8203;57742](https://redirect.github.com/nodejs/node/pull/57742)
- \[[`3cbf5f93d2`](https://redirect.github.com/nodejs/node/commit/3cbf5f93d2)] - **test**: fix missing edge case in test-blob-slice-with-large-size (Joyee Cheung) [#&#8203;58414](https://redirect.github.com/nodejs/node/pull/58414)
- \[[`bffd4ec379`](https://redirect.github.com/nodejs/node/commit/bffd4ec379)] - **test**: skip in test-buffer-tostring-rangeerror on allocation failure (Joyee Cheung) [#&#8203;58415](https://redirect.github.com/nodejs/node/pull/58415)
- \[[`8237346fb7`](https://redirect.github.com/nodejs/node/commit/8237346fb7)] - **test,crypto**: update WebCryptoAPI WPT (Filip Skokan) [#&#8203;54593](https://redirect.github.com/nodejs/node/pull/54593)
- \[[`b90c4ab937`](https://redirect.github.com/nodejs/node/commit/b90c4ab937)] - **tools**: remove unused `osx-pkg-postinstall.sh` (Antoine du Hamel) [#&#8203;57667](https://redirect.github.com/nodejs/node/pull/57667)
- \[[`414013dcfb`](https://redirect.github.com/nodejs/node/commit/414013dcfb)] - **tools**: edit create-release-proposal workflow to handle pr body length (Elves Vieira) [#&#8203;57841](https://redirect.github.com/nodejs/node/pull/57841)
- \[[`7c449ed6b3`](https://redirect.github.com/nodejs/node/commit/7c449ed6b3)] - **tools**: fix tarball testing directory (Marco Ippolito) [#&#8203;57994](https://redirect.github.com/nodejs/node/pull/57994)
- \[[`d164dc2d38`](https://redirect.github.com/nodejs/node/commit/d164dc2d38)] - **tools**: update sccache version to v0.10.0 (Marco Ippolito) [#&#8203;57994](https://redirect.github.com/nodejs/node/pull/57994)
- \[[`debd3c2cc0`](https://redirect.github.com/nodejs/node/commit/debd3c2cc0)] - **tools**: disable failing test envs in `test-linux` CI (Antoine du Hamel) [#&#8203;58351](https://redirect.github.com/nodejs/node/pull/58351)
- \[[`152112505a`](https://redirect.github.com/nodejs/node/commit/152112505a)] - **typings**: fix `ImportModuleDynamicallyCallback` return type (Chengzhong Wu) [#&#8203;57160](https://redirect.github.com/nodejs/node/pull/57160)
- \[[`363bf744ab`](https://redirect.github.com/nodejs/node/commit/363bf744ab)] - **worker**: flush stdout and stderr on exit (Matteo Collina) [#&#8203;56428](https://redirect.github.com/nodejs/node/pull/56428)

### [`v20.19.2`](https://redirect.github.com/nodejs/node/releases/tag/v20.19.2): 2025-05-14, Version 20.19.2 &#x27;Iron&#x27; (LTS), @&#8203;RafaelGSS

[Compare Source](https://redirect.github.com/nodejs/node/compare/v20.19.1...v20.19.2)

This is a security release.

##### Notable Changes

- (CVE-2025-23166) fix error handling on async crypto operation
- (CVE-2025-23167) (SEMVER-MAJOR) update llhttp to 9.2.0
- (CVE-2025-23165) add missing call to uv\_fs\_req\_cleanup

##### Commits

- \[[`eb25047b1b`](https://redirect.github.com/nodejs/node/commit/eb25047b1b)] - **deps**: update llhttp to 9.2.0 (Node.js GitHub Bot) [#&#8203;51719](https://redirect.github.com/nodejs/node/pull/51719)
- \[[`12dcd8db08`](https://redirect.github.com/nodejs/node/commit/12dcd8db08)] - **deps**: update llhttp to 9.1.3 (Node.js GitHub Bot) [#&#8203;50080](https://redirect.github.com/nodejs/node/pull/50080)
- \[[`190e45a291`](https://redirect.github.com/nodejs/node/commit/190e45a291)] - **(SEMVER-MAJOR)** **(CVE-2025-23167)** **deps**: update llhttp to 9.1.2 (Paolo Insogna) [#&#8203;48981](https://redirect.github.com/nodejs/node/pull/48981)
- \[[`fc68c44e6a`](https://redirect.github.com/nodejs/node/commit/fc68c44e6a)] - **fs**: added test for missing call to uv\_fs\_req\_cleanup (Justin Nietzel) [#&#8203;57811](https://redirect.github.com/nodejs/node/pull/57811)
- \[[`9e13bf0a81`](https://redirect.github.com/nodejs/node/commit/9e13bf0a81)] - **(CVE-2025-23165)** **fs**: add missing call to uv\_fs\_req\_cleanup (Justin Nietzel) [#&#8203;57811](https://redirect.github.com/nodejs/node/pull/57811)
- \[[`bd0aa5d44c`](https://redirect.github.com/nodejs/node/commit/bd0aa5d44c)] - **(CVE-2024-27982)** **http**: do not allow OBS fold in headers by default (Paolo Insogna) [nodejs-private/node-private#556](https://redirect.github.com/nodejs-private/node-private/pull/556)
- \[[`6c57465920`](https://redirect.github.com/nodejs/node/commit/6c57465920)] - **(CVE-2025-23166)** **src**: fix error handling on async crypto operations (RafaelGSS) [nodejs-private/node-private#710](https://redirect.github.com/nodejs-private/node-private/pull/710)

### [`v20.19.1`](https://redirect.github.com/nodejs/node/releases/tag/v20.19.1): 2025-04-22, Version 20.19.1 &#x27;Iron&#x27; (LTS), @&#8203;UlisesGascon prepared by @&#8203;RafaelGSS

[Compare Source](https://redirect.github.com/nodejs/node/compare/v20.19.0...v20.19.1)

##### Notable Changes

- \[[`d5e73ce0f8`](https://redirect.github.com/nodejs/node/commit/d5e73ce0f8)] - **deps**: update undici to 6.21.2 (Matteo Collina) [#&#8203;57442](https://redirect.github.com/nodejs/node/pull/57442)
- \[[`e4a6323ab2`](https://redirect.github.com/nodejs/node/commit/e4a6323ab2)] - **deps**: update c-ares to v1.34.5 (Node.js GitHub Bot) [#&#8203;57792](https://redirect.github.com/nodejs/node/pull/57792)

##### Commits

- \[[`d5e73ce0f8`](https://redirect.github.com/nodejs/node/commit/d5e73ce0f8)] - **deps**: update undici to 6.21.2 (Matteo Collina) [#&#8203;57442](https://redirect.github.com/nodejs/node/pull/57442)
- \[[`e4a6323ab2`](https://redirect.github.com/nodejs/node/commit/e4a6323ab2)] - **deps**: update c-ares to v1.34.5 (Node.js GitHub Bot) [#&#8203;57792](https://redirect.github.com/nodejs/node/pull/57792)
- \[[`b2b9eb36af`](https://redirect.github.com/nodejs/node/commit/b2b9eb36af)] - **dns**: restore dns query cache ttl (Ethan Arrowood) [#&#8203;57640](https://redirect.github.com/nodejs/node/pull/57640)
- \[[`07a99a5c0b`](https://redirect.github.com/nodejs/node/commit/07a99a5c0b)] - **doc**: correct status of require(esm) warning in v20 changelog (Joyee Cheung) [#&#8203;57529](https://redirect.github.com/nodejs/node/pull/57529)
- \[[`d45517ccbf`](https://redirect.github.com/nodejs/node/commit/d45517ccbf)] - **meta**: bump Mozilla-Actions/sccache-action from 0.0.8 to 0.0.9 (dependabot\[bot]) [#&#8203;57720](https://redirect.github.com/nodejs/node/pull/57720)
- \[[`fa93bb2633`](https://redirect.github.com/nodejs/node/commit/fa93bb2633)] - **test**: update parallel/test-tls-dhe for OpenSSL 3.5 (Richard Lau) [#&#8203;57477](https://redirect.github.com/nodejs/node/pull/57477)
- \[[`29c032403c`](https://redirect.github.com/nodejs/node/commit/29c032403c)] - **tools**: update sccache to support GH cache changes (Michaël Zasso) [#&#8203;57573](https://redirect.github.com/nodejs/node/pull/57573)

### [`v20.19.0`](https://redirect.github.com/nodejs/node/releases/tag/v20.19.0): 2025-03-13, Version 20.19.0 &#x27;Iron&#x27; (LTS), @&#8203;marco-ippolito

[Compare Source](https://redirect.github.com/nodejs/node/compare/v20.18.3...v20.19.0)

##### Notable Changes

##### require(esm) is now enabled by default

Support for loading native ES modules using require() had been available on v20.x under the command line flag --experimental-require-module, and available by default on v22.x and v23.x. In this release, it is now no longer behind a flag on v20.x.

This feature has been tested on v23.x and v22.x, and we are looking for user feedback from v20.x to make more final tweaks before fully stabilizing it.
It now no longer emits a warning unless `--trace-require-module` is explicitly used.
If there happens to be any regressions caused by this feature, users can report it to the Node.js issue tracker. Meanwhile this feature can also be disabled using `--no-experimental-require-module` as a workaround.

With this feature enabled, Node.js will no longer throw `ERR_REQUIRE_ESM` if `require()` is used to load a ES module. It can, however, throw `ERR_REQUIRE_ASYNC_MODULE` if the ES module being loaded or its dependencies contain top-level `await`. When the ES module is loaded successfully by `require()`, the returned object will either be a ES module namespace object similar to what's returned by `import()`, or what gets exported as `"module.exports"` in the ES module.

Users can check `process.features.require_module` to see whether `require(esm)` is enabled in the current Node.js instance. For packages, the `"module-sync"` exports condition can be used as a way to detect `require(esm)` support in the current Node.js instance and allow both `require()` and `import` to load the same native ES module. See [the documentation](https://nodejs.org/docs/latest/api/modules.html#loading-ecmascript-modules-using-require) for more details about this feature.

Contributed by Joyee Cheung in [#&#8203;55085](https://redirect.github.com/nodejs/node/pull/55085)

##### Module syntax detection is now enabled by default

Module syntax detection (the `--experimental-detect-module` flag) is now
enabled by default. Use `--no-experimental-detect-module` to disable it if
needed.

Syntax detection attempts to run ambiguous files as CommonJS, and if the module
fails to parse as CommonJS due to ES module syntax, Node.js tries again and runs
the file as an ES module.
Ambiguous files are those with a `.js` or no extension, where the nearest parent
`package.json` has no `"type"` field (either `"type": "module"` or
`"type": "commonjs"`).
Syntax detection should have no performance impact on CommonJS modules, but it
incurs a slight performance penalty for ES modules; add `"type": "module"` to
the nearest parent `package.json` file to eliminate the performance cost.
A use case unlocked by this feature is the ability to use ES module syntax in
extensionless scripts with no nearby `package.json`.

Thanks to Geoffrey Booth for making this work on [#&#8203;53619](https://redirect.github.com/nodejs/node/pull/53619).

##### Other Notable Changes

- \[[`285bb4ee14`](https://redirect.github.com/nodejs/node/commit/285bb4ee14)] - **crypto**: update root certificates to NSS 3.107 (Node.js GitHub Bot) [#&#8203;56566](https://redirect.github.com/nodejs/node/pull/56566)
- \[[`73b5c16684`](https://redirect.github.com/nodejs/node/commit/73b5c16684)] - **(SEMVER-MINOR)** **worker**: add postMessageToThread (Paolo Insogna) [#&#8203;53682](https://redirect.github.com/nodejs/node/pull/53682)
- \[[`de313b2336`](https://redirect.github.com/nodejs/node/commit/de313b2336)] - **(SEMVER-MINOR)** **module**: only emit require(esm) warning under --trace-require-module (Joyee Cheung) [#&#8203;56194](https://redirect.github.com/nodejs/node/pull/56194)
- \[[`4fba01911d`](https://redirect.github.com/nodejs/node/commit/4fba01911d)] - **(SEMVER-MINOR)** **process**: add process.features.require\_module (Joyee Cheung) [#&#8203;55241](https://redirect.github.com/nodejs/node/pull/55241)
- \[[`df8a045afe`](https://redirect.github.com/nodejs/node/commit/df8a045afe)] - **(SEMVER-MINOR)** **module**: implement the "module-sync" exports condition (Joyee Cheung) [#&#8203;54648](https://redirect.github.com/nodejs/node/pull/54648)
- \[[`f9dc1eaef5`](https://redirect.github.com/nodejs/node/commit/f9dc1eaef5)] - **(SEMVER-MINOR)** **module**: add \_\_esModule to require()'d ESM (Joyee Cheung) [#&#8203;52166](https://redirect.github.com/nodejs/node/pull/52166)

<details>
<summary>
Commits
</summary>

- \[[`d84be843e3`](https://redirect.github.com/nodejs/node/commit/d84be843e3)] - **benchmark**: add validateStream to styleText bench (Rafael Gonzaga) [#&#8203;56556](https://redirect.github.com/nodejs/node/pull/56556)
- \[[`d8eaf5b9b8`](https://redirect.github.com/nodejs/node/commit/d8eaf5b9b8)] - **build**: fix compatibility with V8's `depot_tools` (Richard Lau) [#&#8203;57330](https://redirect.github.com/nodejs/node/pull/57330)
- \[[`1ee4bf9690`](https://redirect.github.com/nodejs/node/commit/1ee4bf9690)] - **build**: test macos-13 on GitHub actions (Michaël Zasso) [#&#8203;56307](https://redirect.github.com/nodejs/node/pull/56307)
- \[[`1cc8d69882`](https://redirect.github.com/nodejs/node/commit/1cc8d69882)] - **build**: build v8 with -fvisibility=hidden on macOS (Joyee Cheung) [#&#8203;56275](https://redirect.github.com/nodejs/node/pull/56275)
- \[[`52f1f7e22b`](https://redirect.github.com/nodejs/node/commit/52f1f7e22b)] - **child\_

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

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6W119-->


---

## Discussion

### Comment by @xbhouse on October 30, 2025 at 05:41 PM UTC

/retest

---

## Reviews

### Review by @TenSt - Approved on October 30, 2025 at 11:24 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1256*
