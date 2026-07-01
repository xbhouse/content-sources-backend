---
type: pull_request
number: 1302
title: "Build chore(deps): update node.js to v24.11.1"
state: merged
author: red-hat-konflux
created: 2025-11-17T04:38:28Z
updated: 2025-11-17T09:41:11Z
closed: 2025-11-17T09:41:10Z
merged: 2025-11-17T09:41:10Z
base_branch: main
head_branch: konflux/mintmaker/main/node-24.x
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1302
---

# Pull Request #1302: Build chore(deps): update node.js to v24.11.1

**Author**: @red-hat-konflux
**Created**: November 17, 2025 at 04:38 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/node-24.x`

## Description

This PR contains the following updates:

| Package | Update | Change |
|---|---|---|
| [node](https://nodejs.org) ([source](https://redirect.github.com/nodejs/node)) | patch | `24.11.0` -> `24.11.1` |

---

### Release Notes

<details>
<summary>nodejs/node (node)</summary>

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

### Review by @swadeley - Approved on November 17, 2025 at 09:21 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1302*
