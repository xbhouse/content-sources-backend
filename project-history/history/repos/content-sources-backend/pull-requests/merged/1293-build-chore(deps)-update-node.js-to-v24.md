---
type: pull_request
number: 1293
title: "Build chore(deps): update node.js to v24"
state: merged
author: red-hat-konflux
created: 2025-11-10T08:19:31Z
updated: 2025-11-10T20:32:00Z
closed: 2025-11-10T18:02:52Z
merged: 2025-11-10T18:02:52Z
base_branch: main
head_branch: konflux/mintmaker/main/node-24.x
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1293
---

# Pull Request #1293: Build chore(deps): update node.js to v24

**Author**: @red-hat-konflux
**Created**: November 10, 2025 at 08:19 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/node-24.x`

## Description

This PR contains the following updates:

| Package | Type | Update | Change | Age | Confidence |
|---|---|---|---|---|---|
| [node](https://nodejs.org) ([source](https://redirect.github.com/nodejs/node)) |  | major | `20.19.5` -> `24.11.0` | [![age](https://developer.mend.io/api/mc/badges/age/node-version/node/v24.11.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/node-version/node/v20.19.5/v24.11.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |
| [@types/node](https://redirect.github.com/DefinitelyTyped/DefinitelyTyped/tree/master/types/node) ([source](https://redirect.github.com/DefinitelyTyped/DefinitelyTyped/tree/HEAD/types/node)) | devDependencies | major | [`^22.10.2` -> `^24.0.0`](https://renovatebot.com/diffs/npm/@types%2fnode/22.19.0/24.10.0) | [![age](https://developer.mend.io/api/mc/badges/age/npm/@types%2fnode/24.10.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/npm/@types%2fnode/22.19.0/24.10.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

### Release Notes

<details>
<summary>nodejs/node (node)</summary>

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
- \[[`18c79d9e1c`](https://redirect.github.com/nodejs/node/commit/18c79d9e1c)] - **(SEMVER-MINOR)** **sqlite**: create authorization api (Guilherme Araújo) [#&#8203;59928](https://redirect.github.com/nodejs/node/pull/59928)

##### Commits

- \[[`e8cff3d51e`](https://redirect.github.com/nodejs/node/commit/e8cff3d51e)] - **benchmark**: remove unused variable from util/priority-queue (Bruno Rodrigues) [#&#8203;59872](https://redirect.github.com/nodejs/node/pull/59872)
- \[[`03294252ab`](https://redirect.github.com/nodejs/node/commit/03294252ab)] - **benchmark**: update count to n in permission startup (Bruno Rodrigues) [#&#8203;59872](https://redirect.github.com/nodejs/node/pull/59872)
- \[[`3c8a609d9b`](https://redirect.github.com/nodejs/node/commit/3c8a609d9b)] - **benchmark**: update num to n in dgram offset-length (Bruno Rodrigues) [#&#8203;59872](https://redirect.github.com/nodejs/node/pull/59872)
- \[[`7b2032b13e`](https://redirect.github.com/nodejs/node/commit/7b2032b13e)] - **benchmark**: adjust dgram offset-length len values (Bruno Rodrigues) [#&#8203;59708](https://redirect.github.com/nodejs/node/pull/59708)
- \[[`552d887aee`](https://redirect.github.com/nodejs/node/commit/552d887aee)] - **benchmark**: update num to n in dgram offset-length (Bruno Rodrigues) [#&#8203;59708](https://redirect.github.com/nodejs/node/pull/59708)
- \[[`31bb476895`](https://redirect.github.com/nodejs/node/commit/31bb476895)] - **(SEMVER-MINOR)** **console**: allow per-stream `inspectOptions` option (Anna Henningsen) [#&#8203;60082](https://redirect.github.com/nodejs/node/pull/60082)
- \[[`0bf022d4c0`](https://redirect.github.com/nodejs/node/commit/0bf022d4c0)] - **console,util**: improve array inspection performance (Ruben Bridgewater) [#&#8203;60037](https://redirect.github.com/nodejs/node/pull/60037)
- \[[`04d568e591`](https://redirect.github.com/nodejs/node/commit/04d568e591)] - **deps**: V8: cherry-pick [`f93055f`](https://redirect.github.com/nodejs/node/commit/f93055fbd5aa) (Olivier Flückiger) [#&#8203;60105](https://redirect.github.com/nodejs/node/pull/60105)
- \[[`621058b3bf`](https://redirect.github.com/nodejs/node/commit/621058b3bf)] - **deps**: update archs files for openssl-3.5.4 (Node.js GitHub Bot) [#&#8203;60101](https://redirect.github.com/nodejs/node/pull/60101)
- \[[`81b3009fe6`](https://redirect.github.com/nodejs/node/commit/81b3009fe6)] - **deps**: upgrade openssl sources to openssl-3.5.4 (Node.js GitHub Bot) [#&#8203;60101](https://redirect.github.com/nodejs/node/pull/60101)
- \[[`dc44c9f349`](https://redirect.github.com/nodejs/node/commit/dc44c9f349)] - **deps**: upgrade npm to 11.6.1 (npm team) [#&#8203;60012](https://redirect.github.com/nodejs/node/pull/60012)
- \[[`ec0f137198`](https://redirect.github.com/nodejs/node/commit/ec0f137198)] - **deps**: update ada to 3.3.0 (Node.js GitHub Bot) [#&#8203;60045](https://redirect.github.com/nodejs/node/pull/60045)
- \[[`f490f91874`](https://redirect.github.com/nodejs/node/commit/f490f91874)] - **deps**: update amaro to 1.1.4 (pmarchini) [#&#8203;60044](https://redirect.github.com/nodejs/node/pull/60044)
- \[[`de7a7cd0d7`](https://redirect.github.com/nodejs/node/commit/de7a7cd0d7)] - **deps**: update ada to 3.2.9 (Node.js GitHub Bot) [#&#8203;59987](https://redirect.github.com/nodejs/node/pull/59987)
- \[[`a533e5b5db`](https://redirect.github.com/nodejs/node/commit/a533e5b5db)] - **doc**: add automated migration info to deprecations (Augustin Mauroy) [#&#8203;60022](https://redirect.github.com/nodejs/node/pull/60022)
- \[[`7fb8fe4875`](https://redirect.github.com/nodejs/node/commit/7fb8fe4875)] - **doc**: fix typo on child\_process.md (Angelo Gazzola) [#&#8203;60114](https://redirect.github.com/nodejs/node/pull/60114)
- \[[`24c1ef9846`](https://redirect.github.com/nodejs/node/commit/24c1ef9846)] - **doc**: remove optional title prefixes (Aviv Keller) [#&#8203;60087](https://redirect.github.com/nodejs/node/pull/60087)
- \[[`08b9eb8e19`](https://redirect.github.com/nodejs/node/commit/08b9eb8e19)] - **doc**: mark `.env` files support as stable (Santeri Hiltunen) [#&#8203;59925](https://redirect.github.com/nodejs/node/pull/59925)
- \[[`66d90b8063`](https://redirect.github.com/nodejs/node/commit/66d90b8063)] - **doc**: mention reverse proxy and include simple example (Steven) [#&#8203;59736](https://redirect.github.com/nodejs/node/pull/59736)
- \[[`14aa1119cb`](https://redirect.github.com/nodejs/node/commit/14aa1119cb)] - **doc**: provide alternative to `url.parse()` using WHATWG URL (Steven) [#&#8203;59736](https://redirect.github.com/nodejs/node/pull/59736)
- \[[`f9412324f6`](https://redirect.github.com/nodejs/node/commit/f9412324f6)] - **doc**: fix typo of built-in module specifier in worker\_threads (Deokjin Kim) [#&#8203;59992](https://redirect.github.com/nodejs/node/pull/59992)
- \[[`64e738a342`](https://redirect.github.com/nodejs/node/commit/64e738a342)] - **doc,crypto**: reorder ML-KEM in the asymmetric key types table (Filip Skokan) [#&#8203;60067](https://redirect.github.com/nodejs/node/pull/60067)
- \[[`1b25008b41`](https://redirect.github.com/nodejs/node/commit/1b25008b41)] - **http**: improve writeEarlyHints by avoiding for-of loop (Haram Jeong) [#&#8203;59958](https://redirect.github.com/nodejs/node/pull/59958)
- \[[`35f9b6b28f`](https://redirect.github.com/nodejs/node/commit/35f9b6b28f)] - **inspector**: improve batch diagnostic channel subscriptions (Chengzhong Wu) [#&#8203;60009](https://redirect.github.com/nodejs/node/pull/60009)
- \[[`3b92be2fb8`](https://redirect.github.com/nodejs/node/commit/3b92be2fb8)] - **(SEMVER-MINOR)** **lib**: remove util.getCallSite (Rafael Gonzaga) [#&#8203;59980](https://redirect.github.com/nodejs/node/pull/59980)
- \[[`c495e1fe57`](https://redirect.github.com/nodejs/node/commit/c495e1fe57)] - **lib**: optimize priority queue (Gürgün Dayıoğlu) [#&#8203;60039](https://redirect.github.com/nodejs/node/pull/60039)
- \[[`6be31fb9f3`](https://redirect.github.com/nodejs/node/commit/6be31fb9f3)] - **lib**: implement passive listener behavior per spec (BCD1me) [#&#8203;59995](https://redirect.github.com/nodejs/node/pull/59995)
- \[[`c5e4aa763b`](https://redirect.github.com/nodejs/node/commit/c5e4aa763b)] - **meta**: bump actions/setup-python from 5.6.0 to 6.0.0 (dependabot\[bot]) [#&#8203;60090](https://redirect.github.com/nodejs/node/pull/60090)
- \[[`50fa1f4a76`](https://redirect.github.com/nodejs/node/commit/50fa1f4a76)] - **meta**: bump ossf/scorecard-action from 2.4.2 to 2.4.3 (dependabot\[bot]) [#&#8203;60096](https://redirect.github.com/nodejs/node/pull/60096)
- \[[`def4ce976c`](https://redirect.github.com/nodejs/node/commit/def4ce976c)] - **meta**: bump actions/cache from 4.2.4 to 4.3.0 (dependabot\[bot]) [#&#8203;60095](https://redirect.github.com/nodejs/node/pull/60095)
- \[[`24b5abc0e9`](https://redirect.github.com/nodejs/node/commit/24b5abc0e9)] - **meta**: bump step-security/harden-runner from 2.12.2 to 2.13.1 (dependabot\[bot]) [#&#8203;60094](https://redirect.github.com/nodejs/node/pull/60094)
- \[[`8ccf2b0b34`](https://redirect.github.com/nodejs/node/commit/8ccf2b0b34)] - **meta**: bump actions/setup-node from 4.4.0 to 5.0.0 (dependabot\[bot]) [#&#8203;60093](https://redirect.github.com/nodejs/node/pull/60093)
- \[[`78580147ef`](https://redirect.github.com/nodejs/node/commit/78580147ef)] - **meta**: bump actions/stale from 9.1.0 to 10.0.0 (dependabot\[bot]) [#&#8203;60092](https://redirect.github.com/nodejs/node/pull/60092)
- \[[`705686b5c4`](https://redirect.github.com/nodejs/node/commit/705686b5c4)] - **meta**: bump codecov/codecov-action from 5.5.0 to 5.5.1 (dependabot\[bot]) [#&#8203;60091](https://redirect.github.com/nodejs/node/pull/60091)
- \[[`423a6bc744`](https://redirect.github.com/nodejs/node/commit/423a6bc744)] - **meta**: bump github/codeql-action from 3.30.0 to 3.30.5 (dependabot\[bot]) [#&#8203;60089](https://redirect.github.com/nodejs/node/pull/60089)
- \[[`9d9bd0fb4f`](https://redirect.github.com/nodejs/node/commit/9d9bd0fb4f)] - **meta**: move Michael to emeritus (Michael Dawson) [#&#8203;60070](https://redirect.github.com/nodejs/node/pull/60070)
- \[[`dbeee55824`](https://redirect.github.com/nodejs/node/commit/dbeee55824)] - **module**: use sync cjs when importing cts (Marco Ippolito) [#&#8203;60072](https://redirect.github.com/nodejs/node/pull/60072)
- \[[`a722f677ac`](https://redirect.github.com/nodejs/node/commit/a722f677ac)] - **perf\_hooks**: fix histogram fast call signatures (Renegade334) [#&#8203;59600](https://redirect.github.com/nodejs/node/pull/59600)
- \[[`b3295b8353`](https://redirect.github.com/nodejs/node/commit/b3295b8353)] - **process**: fix wrong asyncContext under unhandled-rejections=strict (Shima Ryuhei) [#&#8203;60103](https://redirect.github.com/nodejs/node/pull/60103)
- \[[`cff4a7608a`](https://redirect.github.com/nodejs/node/commit/cff4a7608a)] - **process**: fix default `env` for `process.execve` (Richard Lau) [#&#8203;60029](https://redirect.github.com/nodejs/node/pull/60029)
- \[[`cd034e927f`](https://redirect.github.com/nodejs/node/commit/cd034e927f)] - **process**: fix hrtime fast call signatures (Renegade334) [#&#8203;59600](https://redirect.github.com/nodejs/node/pull/59600)
- \[[`18c79d9e1c`](https://redirect.github.com/nodejs/node/commit/18c79d9e1c)] - **(SEMVER-MINOR)** **sqlite**: create authorization api (Guilherme Araújo) [#&#8203;59928](https://redirect.github.com/nodejs/node/pull/59928)
- \[[`d949222043`](https://redirect.github.com/nodejs/node/commit/d949222043)] - **sqlite**: replace `ToLocalChecked` and improve filter error handling (Edy Silva) [#&#8203;60028](https://redirect.github.com/nodejs/node/pull/60028)
- \[[`6417dc879e`](https://redirect.github.com/nodejs/node/commit/6417dc879e)] - **src**: bring permissions macros in line with general C/C++ standards (Anna Henningsen) [#&#8203;60053](https://redirect.github.com/nodejs/node/pull/60053)
- \[[`e273c2020c`](https://redirect.github.com/nodejs/node/commit/e273c2020c)] - **src**: update contextify to use DictionaryTemplate (James M Snell) [#&#8203;60059](https://redirect.github.com/nodejs/node/pull/60059)
- \[[`5f9ff60664`](https://redirect.github.com/nodejs/node/commit/5f9ff60664)] - **src**: remove `AnalyzeTemporaryDtors` option from .clang-tidy (iknoom) [#&#8203;60008](https://redirect.github.com/nodejs/node/pull/60008)
- \[[`9db54adccc`](https://redirect.github.com/nodejs/node/commit/9db54adccc)] - **src**: update cares\_wrap to use DictionaryTemplates (James M Snell) [#&#8203;60033](https://redirect.github.com/nodejs/node/pull/60033)
- \[[`fc0ceb7b82`](https://redirect.github.com/nodejs/node/commit/fc0ceb7b82)] - **src**: correct the error handling in StatementExecutionHelper (James M Snell) [#&#8203;60040](https://redirect.github.com/nodejs/node/pull/60040)
- \[[`3e8fdc1d8d`](https://redirect.github.com/nodejs/node/commit/3e8fdc1d8d)] - **src**: remove unused variables from report (Moonki Choi) [#&#8203;60047](https://redirect.github.com/nodejs/node/pull/60047)
- \[[`d744324d8e`](https://redirect.github.com/nodejs/node/commit/d744324d8e)] - **src**: avoid unnecessary string allocations in SPrintF impl (Anna Henningsen) [#&#8203;60052](https://redirect.github.com/nodejs/node/pull/60052)
- \[[`de65a5c719`](https://redirect.github.com/nodejs/node/commit/de65a5c719)] - **src**: make ToLower/ToUpper input args more flexible (Anna Henningsen) [#&#8203;60052](https://redirect.github.com/nodejs/node/pull/60052)
- \[[`354026df5a`](https://redirect.github.com/nodejs/node/commit/354026df5a)] - **src**: allow `std::string_view` arguments to `SPrintF()` and friends (Anna Henningsen) [#&#8203;60058](https://redirect.github.com/nodejs/node/pull/60058)
- \[[`42f7d7cb20`](https://redirect.github.com/nodejs/node/commit/42f7d7cb20)] - **src**: remove unnecessary `std::string` error messages (Anna Henningsen) [#&#8203;60057](https://redirect.github.com/nodejs/node/pull/60057)
- \[[`30c2c0fedd`](https://redirect.github.com/nodejs/node/commit/30c2c0fedd)] - **src**: remove unnecessary shadowed functions on Utf8Value & BufferValue (Anna Henningsen) [#&#8203;60056](https://redirect.github.com/nodejs/node/pull/60056)
- \[[`eb99eec09b`](https://redirect.github.com/nodejs/node/commit/eb99eec09b)] - **src**: avoid unnecessary string -> `char*` -> string round trips (Anna Henningsen) [#&#8203;60055](https://redirect.github.com/nodejs/node/pull/60055)
- \[[`c1f1dbdce2`](https://redirect.github.com/nodejs/node/commit/c1f1dbdce2)] - **src**: remove useless dereferencing in `THROW_...` (Anna Henningsen) [#&#8203;60054](https://redirect.github.com/nodejs/node/pull/60054)
- \[[`ea0f5e575d`](https://redirect.github.com/nodejs/node/commit/ea0f5e575d)] - **src**: fill `options_args`, `options_env` after vectors are finalized (iknoom) [#&#8203;59945](https://redirect.github.com/nodejs/node/pull/59945)
- \[[`415fff217a`](https://redirect.github.com/nodejs/node/commit/415fff217a)] - **src**: use RAII for uv\_process\_options\_t (iknoom) [#&#8203;59945](https://redirect.github.com/nodejs/node/pull/59945)
- \[[`982b03ecbd`](https://redirect.github.com/nodejs/node/commit/982b03ecbd)] - **test**: mark `test-runner-run-watch` flaky on macOS (Richard Lau) [#&#8203;60115](https://redirect.github.com/nodejs/node/pull/60115)
- \[[`831a0d3d28`](https://redirect.github.com/nodejs/node/commit/831a0d3d28)] - **test**: ensure that the message event is fired (Luigi Pinca) [#&#8203;59952](https://redirect.github.com/nodejs/node/pull/59952)
- \[[`5538cfc1e8`](https://redirect.github.com/nodejs/node/commit/5538cfc1e8)] - **test**: replace diagnostics\_channel stackframe in output snapshots (Chengzhong Wu) [#&#8203;60024](https://redirect.github.com/nodejs/node/pull/60024)
- \[[`77ec400d90`](https://redirect.github.com/nodejs/node/commit/77ec400d90)] - **test**: mark test-web-locks skip on IBM i (SRAVANI GUNDEPALLI) [#&#8203;59996](https://redirect.github.com/nodejs/node/pull/59996)
- \[[`1aaadb9e31`](https://redirect.github.com/nodejs/node/commit/1aaadb9e31)] - **test**: ensure message event fires in worker message port test (Jarred Sumner) [#&#8203;59885](https://redirect.github.com/nodejs/node/pull/59885)
- \[[`1d5cc5e57a`](https://redirect.github.com/nodejs/node/commit/1d5cc5e57a)] - **test**: mark sea tests flaky on macOS x64 (Richard Lau) [#&#8203;60068](https://redirect.github.com/nodejs/node/pull/60068)
- \[[`c412b1855d`](https://redirect.github.com/nodejs/node/commit/c412b1855d)] - **test**: expand tls-check-server-identity coverage (Diango Gavidia) [#&#8203;60002](https://redirect.github.com/nodejs/node/pull/60002)
- \[[`ad87975029`](https://redirect.github.com/nodejs/node/commit/ad87975029)] - **test**: fix typo of test-benchmark-readline.js (Deokjin Kim) [#&#8203;59993](https://redirect.github.com/nodejs/node/pull/59993)
- \[[`bad4b9b878`](https://redirect.github.com/nodejs/node/commit/bad4b9b878)] - **test**: add new `startNewREPLSever` testing utility (Dario Piotrowicz) [#&#8203;59964](https://redirect.github.com/nodejs/node/pull/59964)
- \[[`ef90b0f456`](https://redirect.github.com/nodejs/node/commit/ef90b0f456)] - **test**: verify tracing channel doesn't swallow unhandledRejection (Gerhard Stöbich) [#&#8203;59974](https://redirect.github.com/nodejs/node/pull/59974)
- \[[`d7285459fe`](https://redirect.github.com/nodejs/node/commit/d7285459fe)] - **timers**: fix binding fast call signatures (Renegade334) [#&#8203;59600](https://redirect.github.com/nodejs/node/pull/59600)
- \[[`6529ae9b0c`](https://redirect.github.com/nodejs/node/commit/6529ae9b0c)] - **tools**: add message on auto-fixing js lint issues in gh workflow (Dario Piotrowicz) [#&#8203;59128](https://redirect.github.com/nodejs/node/pull/59128)
- \[[`1ca116a6ea`](https://redirect.github.com/nodejs/node/commit/1ca116a6ea)] - **tools**: verify signatures when updating nghttp\* (Antoine du Hamel) [#&#8203;60113](https://redirect.github.com/nodejs/node/pull/60113)
- \[[`20d10a2398`](https://redirect.github.com/nodejs/node/commit/20d10a2398)] - **tools**: use dependabot cooldown and move tools/doc (Rafael Gonzaga) [#&#8203;59978](https://redirect.github.com/nodejs/node/pull/59978)
- \[[`275c07064c`](https://redirect.github.com/nodejs/node/commit/275c07064c)] - **typings**: update 'types' binding (René) [#&#8203;59692](https://redirect.github.com/nodejs/node/pull/59692)
- \[[`8c21c4b286`](https://redirect.github.com/nodejs/node/commit/8c21c4b286)] - **wasi**: fix WasiFunction fast call signature (Renegade334) [#&#8203;59600](https://redirect.github.com/nodejs/node/pull/59600)
- \[[`b865074641`](https://redirect.github.com/nodejs/node/commit/b865074641)] - **win,tools**: add description to signature (Martin Costello) [#&#8203;59877](https://redirect.github.com/nodejs/node/pull/59877)

### [`v24.9.0`](https://redirect.github.com/nodejs/node/releases/tag/v24.9.0): 2025-09-25, Version 24.9.0 (Current), @&#8203;targos

[Compare Source](https://redirect.github.com/nodejs/node/compare/v24.8.0...v24.9.0)

##### Notable Changes

- \[[`9b043a9096`](https://redirect.github.com/nodejs/node/commit/9b043a9096)] - **(SEMVER-MINOR)** **http**: add shouldUpgradeCallback to let servers control HTTP upgrades (Tim Perry) [#&#8203;59824](https://redirect.github.com/nodejs/node/pull/59824)
- \[[`a6456ab90a`](https://redirect.github.com/nodejs/node/commit/a6456ab90a)] - **(SEMVER-MINOR)** **sqlite**: cleanup ERM support and export Session class (James M Snell) [#&#8203;58378](https://redirect.github.com/nodejs/node/pull/58378)
- \[[`5563361d22`](https://redirect.github.com/nodejs/node/commit/5563361d22)] - **(SEMVER-MINOR)** **sqlite**: add tagged template (0hm☘️) [#&#8203;58748](https://redirect.github.com/nodejs/node/pull/58748)
- \[[`04013ee933`](https://redirect.github.com/nodejs/node/commit/04013ee933)] - **(SEMVER-MINOR)** **worker**: add heap profile API (theanarkh) [#&#8203;59846](https://redirect.github.com/nodejs/node/pull/59846)

##### Commits

- \[[`cbec4fd6de`](https://redirect.github.com/nodejs/node/commit/cbec4fd6de)] - **benchmark**: calibrate config dgram multi-buffer (Bruno Rodrigues) [#&#8203;59696](https://redirect.github.com/nodejs/node/pull/59696)
- \[[`9a4bbdc3c5`](https://redirect.github.com/nodejs/node/commit/9a4bbdc3c5)] - **benchmark**: calibrate config cluster/echo.js (Nam Yooseong) [#&#8203;59836](https://redirect.github.com/nodejs/node/pull/59836)
- \[[`0b284d86e8`](https://redirect.github.com/nodejs/node/commit/0b284d86e8)] - **build**: add the missing macro definitions for OpenHarmony (hqzing) [#&#8203;59804](https://redirect.github.com/nodejs/node/pull/59804)
- \[[`43e6e54d66`](https://redirect.github.com/nodejs/node/commit/43e6e54d66)] - **build**: do not include custom ESLint rules testing in tarball (Antoine du Hamel) [#&#8203;59809](https://redirect.github.com/nodejs/node/pull/59809)
- \[[`039ac19154`](https://redirect.github.com/nodejs/node/commit/039ac19154)] - **crypto**: expose signatureAlgorithm on X509Certificate (Patrick Costa) [#&#8203;59235](https://redirect.github.com/nodejs/node/pull/59235)
- \[[`647c332704`](https://redirect.github.com/nodejs/node/commit/647c332704)] - **crypto**: use `return await` when returning Promises from async functions (Renegade334) [#&#8203;59841](https://redirect.github.com/nodejs/node/pull/59841)
- \[[`8ed4587cf0`](https://redirect.github.com/nodejs/node/commit/8ed4587cf0)] - **crypto**: use async functions for non-stub Promise-returning functions (Renegade334) [#&#8203;59841](https://redirect.github.com/nodejs/node/pull/59841)
- \[[`bb051c56ef`](https://redirect.github.com/nodejs/node/commit/bb051c56ef)] - **crypto**: avoid calls to `promise.catch()` (Renegade334) [#&#8203;59841](https://redirect.github.com/nodejs/node/pull/59841)
- \[[`05e560dd25`](https://redirect.github.com/nodejs/node/commit/05e560dd25)] - **deps**: update googletest to [`50b8600`](https://redirect.github.com/nodejs/node/commit/50b8600) (Node.js GitHub Bot) [#&#8203;59955](https://redirect.github.com/nodejs/node/pull/59955)
- \[[`fa40d3a785`](https://redirect.github.com/nodejs/node/commit/fa40d3a785)] - **deps**: update archs files for openssl-3.5.3 (Node.js GitHub Bot) [#&#8203;59901](https://redirect.github.com/nodejs/node/pull/59901)
- \[[`8c85570d18`](https://redirect.github.com/nodejs/node/commit/8c85570d18)] - **deps**: upgrade openssl sources to openssl-3.5.3 (Node.js GitHub Bot) [#&#8203;59901](https://redirect.github.com/nodejs/node/pull/59901)
- \[[`b71125664e`](https://redirect.github.com/nodejs/node/commit/b71125664e)] - **deps**: update undici to 7.16.0 (Node.js GitHub Bot) [#&#8203;59830](https://redirect.github.com/nodejs/node/pull/59830)
- \[[`dea5dd7077`](https://redirect.github.com/nodejs/node/commit/dea5dd7077)] - **dgram**: restore buffer optimization in fixBufferList (Yoo) [#&#8203;59934](https://redirect.github.com/nodejs/node/pull/59934)
- \[[`b0c1e67532`](https://redirect.github.com/nodejs/node/commit/b0c1e67532)] - **diagnostics\_channel**: fix race condition with diagnostics\_channel and GC (Ugaitz Urien) [#&#8203;59910](https://redirect.github.com/nodejs/node/pull/59910)
- \[[`0b37b594c3`](https://redirect.github.com/nodejs/node/commit/0b37b594c3)] - **doc**: use "WebAssembly" instead of "Web Assembly" (Tobias Nießen) [#&#8203;59954](https://redirect.github.com/nodejs/node/pull/59954)
- \[[`1e723f9c6b`](https://redirect.github.com/nodejs/node/commit/1e723f9c6b)] - **doc**: fix typo in section on microtask order (Tobias Nießen) [#&#8203;59932](https://redirect.github.com/nodejs/node/pull/59932)
- \[[`a28962a85c`](https://redirect.github.com/nodejs/node/commit/a28962a85c)] - **doc**: update V8 fast API guidance (René) [#&#8203;58999](https://redirect.github.com/nodejs/node/pull/58999)
- \[[`bd767c5d1b`](https://redirect.github.com/nodejs/node/commit/bd767c5d1b)] - **doc**: add security escalation policy (Ulises Gascón) [#&#8203;59806](https://redirect.github.com/nodejs/node/pull/59806)
- \[[`9df91e59e1`](https://redirect.github.com/nodejs/node/commit/9df91e59e1)] - **doc**: type improvement of file `http.md` (yusheng chen) [#&#8203;58189](https://redirect.github.com/nodejs/node/pull/58189)
- \[[`e4f571680b`](https://redirect.github.com/nodejs/node/commit/e4f571680b)] - **doc**: deprecate closing `fs.Dir` on garbage collection (Livia Medeiros) [#&#8203;59839](https://redirect.github.com/nodejs/node/pull/59839)
- \[[`e9cb986fa5`](https://redirect.github.com/nodejs/node/commit/e9cb986fa5)] - **doc**: rephrase dynamic import() description (Nam Yooseong) [#&#8203;59224](https://redirect.github.com/nodejs/node/pull/59224)
- \[[`026d4e33f7`](https://redirect.github.com/nodejs/node/commit/026d4e33f7)] - **doc,crypto**: update subtle.generateKey and subtle.importKey (Filip Skokan) [#&#8203;59851](https://redirect.github.com/nodejs/node/pull/59851)
- \[[`2b2591db52`](https://redirect.github.com/nodejs/node/commit/2b2591db52)] - **esm**: make hasAsyncGraph non-enumerable (Joyee Cheung) [#&#8203;59905](https://redirect.github.com/nodejs/node/pull/59905)
- \[[`993f05d323`](https://redirect.github.com/nodejs/node/commit/993f05d323)] - **fs,win**: do not add a second trailing slash in readdir (Gerhard Stöbich) [#&#8203;59847](https://redirect.github.com/nodejs/node/pull/59847)
- \[[`7aec53b607`](https://redirect.github.com/nodejs/node/commit/7aec53b607)] - **(SEMVER-MINOR)** **http**: add shouldUpgradeCallback to let servers control HTTP upgrades (Tim Perry) [#&#8203;59824](https://redirect.github.com/nodejs/node/pull/59824)
- \[[`83ae6102e7`](https://redirect.github.com/nodejs/node/commit/83ae6102e7)] - **http**: optimize checkIsHttpToken for short strings (방진혁) [#&#8203;59832](https://redirect.github.com/nodejs/node/pull/59832)
- \[[`6695067636`](https://redirect.github.com/nodejs/node/commit/6695067636)] - **http,https**: handle IPv6 with proxies (Joyee Cheung) [#&#8203;59894](https://redirect.github.com/nodejs/node/pull/59894)
- \[[`c5d910a0a9`](https://redirect.github.com/nodejs/node/commit/c5d910a0a9)] - **http2**: fix allowHttp1+Upgrade, broken by shouldUpgradeCallback (Tim Perry) [#&#8203;59924](https://redirect.github.com/nodejs/node/pull/59924)
- \[[`acada1fb82`](https://redirect.github.com/nodejs/node/commit/acada1fb82)] - **inspector**: ensure adequate memory allocation for `Binary::toBase64` (René) [#&#8203;59870](https://redirect.github.com/nodejs/node/pull/59870)
- \[[`396cc8ec65`](https://redirect.github.com/nodejs/node/commit/396cc8ec65)] - **lib**: update inspect output format for subclasses (Miguel Marcondes Filho) [#&#8203;59687](https://redirect.github.com/nodejs/node/pull/59687)
- \[[`fed1dac8de`](https://redirect.github.com/nodejs/node/commit/fed1dac8de)] - **lib**: update isDeepStrictEqual to support options (Miguel Marcondes Filho) [#&#8203;59762](https://redirect.github.com/nodejs/node/pull/59762)
- \[[`d785929fd7`](https://redirect.github.com/nodejs/node/commit/d785929fd7)] - **lib**: add source map support for assert messages (Chengzhong Wu) [#&#8203;59751](https://redirect.github.com/nodejs/node/pull/59751)
- \[[`ff13d1d61e`](https://redirect.github.com/nodejs/node/commit/ff13d1d61e)] - **lib,src**: cache ModuleWrap.hasAsyncGraph (Chengzhong Wu) [#&#8203;59703](https://redirect.github.com/nodejs/node/pull/59703)
- \[[`b200cd8470`](https://redirect.github.com/nodejs/node/commit/b200cd8470)] - **lib,src**: refactor assert to load error source from memory (Chengzhong Wu) [#&#8203;59751](https://redirect.github.com/nodejs/node/pull/59751)
- \[[`e94c57301b`](https://redirect.github.com/nodejs/node/commit/e94c57301b)] - **meta**: add .npmrc with ignore-scripts=true (Joyee Cheung) [#&#8203;59914](https://redirect.github.com/nodejs/node/pull/59914)
- \[[`728472a57b`](https://redirect.github.com/nodejs/node/commit/728472a57b)] - **module**: only put directly require-d ESM into require.cache (Joyee Cheung) [#&#8203;59874](https://redirect.github.com/nodejs/node/pull/59874)
- \[[`be48760b93`](https://redirect.github.com/nodejs/node/commit/be48760b93)] - **node-api**: added SharedArrayBuffer api (Mert Can Altin) [#&#8203;59071](https://redirect.github.com/nodejs/node/pull/59071)
- \[[`f006a14522`](https://redirect.github.com/nodejs/node/commit/f006a14522)] - **node-api**: make napi\_delete\_reference use node\_api\_basic\_env (Jeetu Suthar) [#&#8203;59684](https://redirect.github.com/nodejs/node/pull/59684)
- \[[`0f46c1c3b0`](https://redirect.github.com/nodejs/node/commit/0f46c1c3b0)] - **repl**: fix cpu overhead pasting big strings to the REPL (Ruben Bridgewater) [#&#8203;59857](https://redirect.github.com/nodejs/node/pull/59857)
- \[[`3eeb7b47ea`](https://redirect.github.com/nodejs/node/commit/3eeb7b47ea)] - **sqlite**: fix crash session extension callbacks with workers (Bart Louwers) [#&#8203;59848](https://redirect.github.com/nodejs/node/pull/59848)
- \[[`0fe53375ec`](https://redirect.github.com/nodejs/node/commit/0fe53375ec)] - **(SEMVER-MINOR)** **sqlite**: cleanup ERM support and export Session class (James M Snell) [#&#8203;58378](https://redirect.github.com/nodejs/node/pull/58378)
- \[[`9a3e58a007`](https://redirect.github.com/nodejs/node/commit/9a3e58a007)] - **(SEMVER-MINOR)** **sqlite**: add tagged template (0hm☘️) [#&#8203;58748](https://redirect.github.com/nodejs/node/pull/58748)
- \[[`f14ed5ab7b`](https://redirect.github.com/nodejs/node/commit/f14ed5ab7b)] - **src**: simplify watchdog instantiations via `std::optional` (Anna Henningsen) [#&#8203;59960](https://redirect.github.com/nodejs/node/pull/59960)
- \[[`e330f03f84`](https://redirect.github.com/nodejs/node/commit/e330f03f84)] - **src**: update crypto objects to use DictionaryTemplate (James M Snell) [#&#8203;59942](https://redirect.github.com/nodejs/node/pull/59942)
- \[[`69b5607cf4`](https://redirect.github.com/nodejs/node/commit/69b5607cf4)] - **src**: simplify is\_callable by making it a concept (Tobias Nießen) [#&#8203;58169](https://redirect.github.com/nodejs/node/pull/58169)
- \[[`86150f3401`](https://redirect.github.com/nodejs/node/commit/86150f3401)] - **src**: rename private fields to follow naming convention (Moonki Choi) [#&#8203;59923](https://redirect.github.com/nodejs/node/pull/59923)
- \[[`d17f299539`](https://redirect.github.com/nodejs/node/commit/d17f299539)] - **src**: use DictionaryTemplate more in URLPattern (James M Snell) [#&#8203;59892](https://redirect.github.com/nodejs/node/pull/59892)
- \[[`ac784912ac`](https://redirect.github.com/nodejs/node/commit/ac784912ac)] - **src**: reduce the nearest parent package JSON cache size (Michael Smith) [#&#8203;59888](https://redirect.github.com/nodejs/node/pull/59888)
- \[[`abecdcb536`](https://redirect.github.com/nodejs/node/commit/abecdcb536)] - **src**: replace FIXED\_ONE\_BYTE\_STRING with Environment-cached strings (Moonki Choi) [#&#8203;59891](https://redirect.github.com/nodejs/node/pull/59891)
- \[[`2bb152500b`](https://redirect.github.com/nodejs/node/commit/2bb152500b)] - **src**: create strings in `FIXED_ONE_BYTE_STRING` as internalized (Anna Henningsen) [#&#8203;59826](https://redirect.github.com/nodejs/node/pull/59826)
- \[[`03116a7cd8`](https://redirect.github.com/nodejs/node/commit/03116a7cd8)] - **src**: remove `std::array` overload of `FIXED_ONE_BYTE_STRING` (Anna Henningsen) [#&#8203;59826](https://redirect.github.com/nodejs/node/pull/59826)
- \[[`8a5325d6e3`](https://redirect.github.com/nodejs/node/commit/8a5325d6e3)] - **src**: ensure `v8::Eternal` is empty before setting it (Anna Henningsen) [#&#8203;59825](https://redirect.github.com/nodejs/node/pull/59825)
- \[[`f0c20ccd81`](https://redirect.github.com/nodejs/node/commit/f0c20ccd81)] - **src**: remove unnecessary `Environment::GetCurrent()` calls (Moonki Choi) [#&#8203;59814](https://redirect.github.com/nodejs/node/pull/59814)
- \[[`213188e491`](https://redirect.github.com/nodejs/node/commit/213188e491)] - **stream**: use new AsyncResource instead of bind (Matteo Collina) [#&#8203;59867](https://redirect.github.com/nodejs/node/pull/59867)
- \[[`ce8435b003`](https://redirect.github.com/nodejs/node/commit/ce8435b003)] - **test**: testcase demonstrating issue 59541 (Eric Rannaud) [#&#8203;59801](https://redirect.github.com/nodejs/node/pull/59801)
- \[[`8f32746142`](https://redirect.github.com/nodejs/node/commit/8f32746142)] - **test**: guard write to proxy client if proxy connection is ended (Joyee Cheung) [#&#8203;59742](https://redirect.github.com/nodejs/node/pull/59742)
- \[[`6790093fcb`](https://redirect.github.com/nodejs/node/commit/6790093fcb)] - **tls**: load bundled and extra certificates off-thread (Joyee Cheung) [#&#8203;59856](https://redirect.github.com/nodejs/node/pull/59856)
- \[[`f5d3f919d8`](https://redirect.github.com/nodejs/node/commit/f5d3f919d8)] - **tls**: only do off-thread certificate loading on loading tls (Joyee Cheung) [#&#8203;59856](https://redirect.github.com/nodejs/node/pull/59856)
- \[[`87bbaa23a0`](https://redirect.github.com/nodejs/node/commit/87bbaa23a0)] - **tools**: fix `tools/make-v8.sh` for clang (Richard Lau) [#&#8203;59893](https://redirect.github.com/nodejs/node/pull/59893)
- \[[`0d23fd525b`](https://redirect.github.com/nodejs/node/commit/0d23fd525b)] - **tools**: skip test-internet workflow for draft PRs (Michaël Zasso) [#&#8203;59817](https://redirect.github.com/nodejs/node/pull/59817)
- \[[`e17c73731a`](https://redirect.github.com/nodejs/node/commit/e17c73731a)] - **tools**: copyedit `build-tarball.yml` (Antoine du Hamel) [#&#8203;59808](https://redirect.github.com/nodejs/node/pull/59808)
- \[[`97c4e1bac9`](https://redirect.github.com/nodejs/node/commit/97c4e1bac9)] - **typings**: remove unused imports (Nam Yooseong) [#&#8203;59880](https://redirect.github.com/nodejs/node/pull/59880)
- \[[`8b29bbca76`](https://redirect.github.com/nodejs/node/commit/8b29bbca76)] - **url**: replaced slice with at (Mikhail) [#&#8203;59181](https://redirect.github.com/nodejs/node/pull/59181)
- \[[`6458867a6b`](https://redirect.github.com/nodejs/node/commit/6458867a6b)] - **url**: add type checking to urlToHttpOptions() (simon-id) [#&#8203;59753](https://redirect.github.com/nodejs/node/pull/59753)
- \[[`3c62b3886f`](https://redirect.github.com/nodejs/node/commit/3c62b3886f)] - **util**: inspect objects with throwing Symbol.toStringTag (Ruben Bridgewater) [#&#8203;59860](https://redirect.github.com/nodejs/node/pull/59860)
- \[[`6133a82875`](https://redirect.github.com/nodejs/node/commit/6133a82875)] - **util**: fix debuglog.enabled not being present with callback logger (Ruben Bridgewater) [#&#8203;59858](https://redirect.github.com/nodejs/node/pull/59858)
- \[[`9347ddddf4`](https://redirect.github.com/nodejs/node/commit/9347ddddf4)] - **vm**: explain how to share promises between contexts w/ afterEvaluate (Eric Rannaud) [#&#8203;59801](https://redirect.github.com/nodejs/node/pull/59801)
- \[[`44ce971619`](https://redirect.github.com/nodejs/node/commit/44ce971619)] - **vm**: "afterEvaluate", evaluate() return a promise from the outer context (Eric Rannaud) [#&#8203;59801](https://redirect.github.com/nodejs/node/pull/59801)
- \[[`6e586a1409`](https://redirect.github.com/nodejs/node/commit/6e586a1409)] - **vm**: expose hasTopLevelAwait on SourceTextModule (Chengzhong Wu) [#&#8203;59865](https://redirect.github.com/nodejs/node/pull/59865)
- \[[`49747a58a3`](https://redirect.github.com/nodejs/node/commit/49747a58a3)] - **(SEMVER-MINOR)** **worker**: add heap profile API (theanarkh) [#&#8203;59846](https://redirect.github.com/nodejs/node/pull/59846)
- \[[`b970c0bbc2`](https://redirect.github.com/nodejs/node/commit/b970c0bbc2)] - **zlib**: reduce code duplication (jhofstee) [#&#8203;57810](https://redirect.github.com/nodejs/node/pull/57810)
- \[[`9782ca2b1b`](https://redirect.github.com/nodejs/node/commit/9782ca2b1b)] - **zlib**: implement fast path for crc32 (Gürgün Dayıoğlu) [#&#8203;59813](https://redirect.github.com/nodejs/node/pull/59813)

### [`v24.8.0`](https://redirect.github.com/nodejs/node/releases/tag/v24.8.0): 2025-09-10, Version 24.8.0 (Current), @&#8203;targos

[Compare Source](https://redirect.github.com/nodejs/node/compare/v24.7.0...v24.8.0)

##### Notable Changes

##### HTTP/2 Network Inspection Support in Node.js

Node.js now supports inspection of HTTP/2 network calls in Chrome DevTools for Node.js.

##### Usage

Write a `test.js` script that makes HTTP/2 requests.

```js
const http2 = require('node:http2');

const client = http2.connect('https://nghttp2.org');

const req = client.request([
  ':path', '/',
  ':method', 'GET',
]);
```

Run it with these options:

```bash
node --inspect-wait --experimental-network-inspection test.js
```

Open `about:inspect` on Google Chrome and click on `Open dedicated DevTools for Node`.
The `Network` tab will let you track your HTTP/2 calls.

Contributed by Darshan Sen in [#&#8203;59611](https://redirect.github.com/nodejs/node/pull/59611).

##### Other Notable Changes

- \[[`7a8e2c251d`](https://redirect.github.com/nodejs/node/commit/7a8e2c251d)] - **(SEMVER-MINOR)** **crypto**: support Ed448 and ML-DSA context parameter in node:crypto (Filip Skokan) [#&#8203;59570](https://redirect.github.com/nodejs/node/pull/59570)
- \[[`4b631be0b0`](https://redirect.github.com/nodejs/node/commit/4b631be0b0)] - **(SEMVER-MINOR)** **crypto**: support Ed448 and ML-DSA context parameter in Web Cryptography (Filip Skokan) [#&#8203;59570](https://redirect.github.com/nodejs/node/pull/59570)
- \[[`3e4b1e732c`](https://redirect.github.com/nodejs/node/commit/3e4b1e732c)] - **(SEMVER-MINOR)** **crypto**: add KMAC Web Cryptography algorithms (Filip Skokan) [#&#8203;59647](https://redirect.github.com/nodejs/node/pull/59647)
- \[[`b1d28785b2`](https://redirect.github.com/nodejs/node/commit/b1d28785b2)] - **(SEMVER-MINOR)** **crypto**: add Argon2 Web Cryptography algorithms (Filip Skokan) [#&#8203;59544](https://redirect.github.com/nodejs/node/pull/59544)
- \[[`430691d1af`](https://redirect.github.com/nodejs/node/commit/430691d1af)] - **(SEMVER-MINOR)** **crypto**: support SLH-DSA KeyObject, sign, and verify (Filip Skokan) [#&#8203;59537](https://redirect.github.com/nodejs/node/pull/59537)
- \[[`d6d05ba397`](https://redirect.github.com/nodejs/node/commit/d6d05ba397)] - **(SEMVER-MINOR)** **worker**: add cpu profile APIs for worker (theanarkh) [#&#8203;59428](https://redirect.github.com/nodejs/node/pull/59428)

##### Commits

- \[[`d913872369`](https://redirect.github.com/nodejs/node/commit/d913872369)] - **assert**: cap input size in myersDiff to avoid Int32Array overflow (Haram Jeong) [#&#8203;59578](https://redirect.github.com/nodejs/node/pull/59578)
- \[[`7bbbcf6666`](https://redirect.github.com/nodejs/node/commit/7bbbcf6666)] - **benchmark**: sqlite prevent create both tables on prepare selects (Bruno Rodrigues) [#&#8203;59709](https://redirect.github.com/nodejs/node/pull/59709)
- \[[`44d7b92271`](https://redirect.github.com/nodejs/node/commit/44d7b92271)] - **benchmark**: calibrate config array-vs-concat (Rafael Gonzaga) [#&#8203;59587](https://redirect.github.com/nodejs/node/pull/59587)
- \[[`7f347fc551`](https://redirect.github.com/nodejs/node/commit/7f347fc551)] - **build**: fix getting OpenSSL version on Windows (Michaël Zasso) [#&#8203;59609](https://redirect.github.com/nodejs/node/pull/59609)
- \[[`4a317150d5`](https://redirect.github.com/nodejs/node/commit/4a317150d5)] - **build**: fix 'implicit-function-declaration' on OpenHarmony platform (hqzing) [#&#8203;59547](https://redirect.github.com/nodejs/node/pull/59547)
- \[[`bda32af587`](https://redirect.github.com/nodejs/node/commit/bda32af587)] - **build**: use `windows-2025` runner (Michaël Zasso) [#&#8203;59673](https://redirect.github.com/nodejs/node/pull/59673)
- \[[`a4a8ed8f6e`](https://redirect.github.com/nodejs/node/commit/a4a8ed8f6e)] - **build**: compile bundled uvwasi conditionally (Carlo Cabrera) [#&#8203;59622](https://redirect.github.com/nodejs/node/pull/59622)
- \[[`d944a87761`](https://redirect.github.com/nodejs/node/commit/d944a87761)] - **crypto**: refactor subtle methods to use synchronous import (Filip Skokan) [#&#8203;59771](https://redirect.github.com/nodejs/node/pull/59771)
- \[[`7a8e2c251d`](https://redirect.github.com/nodejs/node/commit/7a8e2c251d)] - **(SEMVER-MINOR)** **crypto**: support Ed448 and ML-DSA context parameter in node:crypto (Filip Skokan) [#&#8203;59570](https://redirect.github.com/nodejs/node/pull/59570)
- \[[`4b631be0b0`](https://redirect.github.com/nodejs/node/commit/4b631be0b0)] - **(SEMVER-MINOR)** **crypto**: support Ed448 and ML-DSA context parameter in Web Cryptography (Filip Skokan) [#&#8203;59570](https://redirect.github.com/nodejs/node/pull/59570)
- \[[`3e4b1e732c`](https://redirect.github.com/nodejs/node/commit/3e4b1e732c)] - **(SEMVER-MINOR)** **crypto**: add KMAC Web Cryptography algorithms (Filip Skokan) [#&#8203;59647](https://redirect.github.com/nodejs/node/pull/59647)
- \[[`b1d28785b2`](https://redirect.github.com/nodejs/node/commit/b1d28785b2)] - **(SEMVER-MINOR)** **crypto**: add Argon2 Web Cryptography algorithms (Filip Skokan) [#&#8203;59544](https://redirect.github.com/nodejs/node/pull/59544)
- \[[`430691d1af`](https://redirect.github.com/nodejs/node/commit/430691d1af)] - **(SEMVER-MINOR)** **crypto**: support SLH-DSA KeyObject, sign, and verify (Filip Skokan) [#&#8203;59537](https://redirect.github.com/nodejs/node/pull/59537)
- \[[`0d1e53d935`](https://redirect.github.com/nodejs/node/commit/0d1e53d935)] - **deps**: update uvwasi to 0.0.23 (Node.js GitHub Bot) [#&#8203;59791](https://redirect.github.com/nodejs/node/pull/59791)
- \[[`68732cf426`](https://redirect.github.com/nodejs/node/commit/68732cf426)] - **deps**: update histogram to 0.11.9 (Node.js GitHub Bot) [#&#8203;59689](https://redirect.github.com/nodejs/node/pull/59689)
- \[[`f12c1ad961`](https://redirect.github.com/nodejs/node/commit/f12c1ad961)] - **deps**: update googletest to [`eb2d85e`](https://redirect.github.com/nodejs/node/commit/eb2d85e) (Node.js GitHub Bot) [#&#8203;59335](https://redirect.github.com/nodejs/node/pull/59335)
- \[[`45af6966ae`](https://redirect.github.com/nodejs/node/commit/45af6966ae)] - **deps**: upgrade npm to 11.6.0 (npm team) [#&#8203;59750](https://redirect.github.com/nodejs/node/pull/59750)
- \[[`57617244a4`](https://redirect.github.com/nodejs/node/commit/57617244a4)] - **deps**: V8: cherry-pick [`6b1b9bc`](https://redirect.github.com/nodejs/node/commit/6b1b9bca2a8) (Xiao-Tao) [#&#8203;59283](https://redirect.github.com/nodejs/node/pull/59283)
- \[[`2e6225a747`](https://redirect.github.com/nodejs/node/commit/2e6225a747)] - **deps**: update amaro to 1.1.2 (Node.js GitHub Bot) [#&#8203;59616](https://redirect.github.com/nodejs/node/pull/59616)
- \[[`1f7f6dfae6`](https://redirect.github.com/nodejs/node/commit/1f7f6dfae6)] - **diagnostics\_channel**: revoke DEP0163 (René) [#&#8203;59758](https://redirect.github.com/nodejs/node/pull/59758)
- \[[`8671a6cdb3`](https://redirect.github.com/nodejs/node/commit/8671a6cdb3)] - **doc**: stabilize --disable-sigusr1 (Rafael Gonzaga) [#&#8203;59707](https://redirect.github.com/nodejs/node/pull/59707)
- \[[`583b1b255d`](https://redirect.github.com/nodejs/node/commit/583b1b255d)] - **doc**: update OpenSSL default security level to 2 (Jeetu Suthar) [#&#8203;59723](https://redirect.github.com/nodejs/node/pull/59723)
- \[[`9b5eb6eb50`](https://redirect.github.com/nodejs/node/commit/9b5eb6eb50)] - **doc**: fix missing links in the `errors` page (Nam Yooseong) [#&#8203;59427](https://redirect.github.com/nodejs/node/pull/59427)
- \[[`e7bf712c57`](https://redirect.github.com/nodejs/node/commit/e7bf712c57)] - **doc**: update "Type stripping in dependencies" section (Josh Kelley) [#&#8203;59652](https://redirect.github.com/nodejs/node/pull/59652)
- \[[`96db47f91e`](https://redirect.github.com/nodejs/node/commit/96db47f91e)] - **doc**: add Miles Guicent as triager (Miles Guicent) [#&#8203;59562](https://redirect.github.com/nodejs/node/pull/59562)
- \[[`87f829bd0c`](https://redirect.github.com/nodejs/node/commit/87f829bd0c)] - **doc**: mark `path.matchesGlob` as stable (Aviv Keller) [#&#8203;59572](https://redirect.github.com/nodejs/node/pull/59572)
- \[[`062b2f705e`](https://redirect.github.com/nodejs/node/commit/062b2f705e)] - **doc**: improve documentation for raw headers in HTTP/2 APIs (Tim Perry) [#&#8203;59633](https://redirect.github.com/nodejs/node/pull/59633)
- \[[`6ab9306370`](https://redirect.github.com/nodejs/node/commit/6ab9306370)] - **doc**: update install\_tools.bat free disk space (Stefan Stojanovic) [#&#8203;59579](https://redirect.github.com/nodejs/node/pull/59579)
- \[[`c8d6b60da6`](https://redirect.github.com/nodejs/node/commit/c8d6b60da6)] - **doc**: fix quic session instance typo (jakecastelli) [#&#8203;59642](https://redirect.github.com/nodejs/node/pull/59642)
- \[[`61d0a2d1ba`](https://redirect.github.com/nodejs/node/commit/61d0a2d1ba)] - **doc**: fix filehandle.read typo (Ruy Adorno) [#&#8203;59635](https://redirect.github.com/nodejs/node/pull/59635)
- \[[`3276bfa0d0`](https://redirect.github.com/nodejs/node/commit/3276bfa0d0)] - **doc**: update migration recomendations for `util.is**()` deprecations (Augustin Mauroy) [#&#8203;59269](https://redirect.github.com/nodejs/node/pull/59269)
- \[[`11de6c7ebb`](https://redirect.github.com/nodejs/node/commit/11de6c7ebb)] - **doc**: fix missing link to the Error documentation in the `http` page (Alexander Makarenko) [#&#8203;59080](https://redirect.github.com/nodejs/node/pull/59080)
- \[[`f5b6829bba`](https://redirect.github.com/nodejs/node/commit/f5b6829bba)] - **doc,crypto**: add description to the KEM and supports() methods (Filip Skokan) [#&#8203;59644](https://redirect.github.com/nodejs/node/pull/59644)
- \[[`5bfdc7ee74`](https://redirect.github.com/nodejs/node/commit/5bfdc7ee74)] - **doc,crypto**: cleanup unlinked and self method references webcrypto.md (Filip Skokan) [#&#8203;59608](https://redirect.github.com/nodejs/node/pull/59608)
- \[[`010458d061`](https://redirect.github.com/nodejs/node/commit/010458d061)] - **esm**: populate separate cache for require(esm) in imported CJS (Joyee Cheung) [#&#8203;59679](https://redirect.github.com/nodejs/node/pull/59679)
- \[[`dbe6e63baf`](https://redirect.github.com/nodejs/node/commit/dbe6e63baf)] - **esm**: fix missed renaming in ModuleJob.runSync (Joyee Cheung) [#&#8203;59724](https://redirect.github.com/nodejs/node/pull/59724)
- \[[`8eb0d9d834`](https://redirect.github.com/nodejs/node/commit/8eb0d9d834)] - **fs**: fix wrong order of file names in cpSync error message (Nicholas Paun) [#&#8203;59775](https://redirect.github.com/nodejs/node/pull/59775)
- \[[`e69be5611f`](https://redirect.github.com/nodejs/node/commit/e69be5611f)] - **fs**: fix dereference: false on cpSync (Nicholas Paun) [#&#8203;59681](https://redirect.github.com/nodejs/node/pull/59681)
- \[[`2865d2ac20`](https://redirect.github.com/nodejs/node/commit/2865d2ac20)] - **http**: unbreak keepAliveTimeoutBuffer (Robert Nagy) [#&#8203;59784](https://redirect.github.com/nodejs/node/pull/59784)
- \[[`ade1175475`](https://redirect.github.com/nodejs/node/commit/ade1175475)] - **http**: use cached '1.1' http version string (Robert Nagy) [#&#8203;59717](https://redirect.github.com/nodejs/node/pull/59717)
- \[[`74a09482de`](https://redirect.github.com/nodejs/node/commit/74a09482de)] - **inspector**: undici as shared-library should pass tests (Aras Abbasi) [#&#8203;59837](https://redirect.github.com/nodejs/node/pull/59837)
- \[[`772f8f415a`](https://redirect.github.com/nodejs/node/commit/772f8f415a)] - **inspector**: add http2 tracking support (Darshan Sen) [#&#8203;59611](https://redirect.github.com/nodejs/node/pull/59611)
- \[[`3d225572d7`](https://redirect.github.com/nodejs/node/commit/3d225572d7)] - ***Revert*** "**lib**: optimize writable stream buffer clearing" (Yoo) [#&#8203;59743](https://redirect.github.com/nodejs/node/pull/59743)
- \[[`4fd213ce73`](https://redirect.github.com/nodejs/node/commit/4fd213ce73)] - **lib**: fix isReadable and isWritable return type value (Gabriel Quaresma) [#&#8203;59089](https://redirect.github.com/nodejs/node/pull/59089)
- \[[`39befddb87`](https://redirect.github.com/nodejs/node/commit/39befddb87)] - **lib**: prefer TypedArrayPrototype primordials (Filip Skokan) [#&#8203;59766](https://redirect.github.com/nodejs/node/pull/59766)
- \[[`0748160d2e`](https://redirect.github.com/nodejs/node/commit/0748160d2e)] - **lib**: fix DOMException subclass support (Chengzhong Wu) [#&#8203;59680](https://redirect.github.com/nodejs/node/pull/59680)
- \[[`1a93df808c`](https://redirect.github.com/nodejs/node/commit/1a93df808c)] - **lib**: revert to using default derived class constructors (René) [#&#8203;59650](https://redirect.github.com/nodejs/node/pull/59650)
- \[[`bb0755df37`](https://redirect.github.com/nodejs/node/commit/bb0755df37)] - **meta**: bump `codecov/codecov-action` (dependabot\[bot]) [#&#8203;59726](https://redirect.github.com/nodejs/node/pull/59726)
- \[[`45d148d9be`](https://redirect.github.com/nodejs/node/commit/45d148d9be)] - **meta**: bump actions/download-artifact from 4.3.0 to 5.0.0 (dependabot\[bot]) [#&#8203;59729](https://redirect.github.com/nodejs/node/pull/59729)
- \[[`01b66b122e`](https://redirect.github.com/nodejs/node/commit/01b66b122e)] - **meta**: bump github/codeql-action from 3.29.2 to 3.30.0 (dependabot\[bot]) [#&#8203;59728](https://redirect.github.com/nodejs/node/pull/59728)
- \[[`34f7ab5502`](https://redirect.github.com/nodejs/node/commit/34f7ab5502)] - **meta**: bump actions/cache from 4.2.3 to 4.2.4 (dependabot\[bot]) [#&#8203;59727](https://redirect.github.com/nodejs/node/pull/59727)
- \[[`5806ea02af`](https://redirect.github.com/nodejs/node/commit/5806ea02af)] - **meta**: bump actions/checkout from 4.2.2 to 5.0.0 (dependabot\[bot]) [#&#8203;59725](https://redirect.github.com/nodejs/node/pull/59725)
- \[[`f667215583`](https://redirect.github.com/nodejs/node/commit/f667215583)] - **path**: refactor path joining logic for clarity and performance (Lee Jiho) [#&#8203;59781](https://redirect.github.com/nodejs/node/pull/59781)
- \[[`0340fe92a6`](https://redirect.github.com/nodejs/node/commit/0340fe92a6)] - **repl**: do not cause side effects in tab completion (Anna Henningsen) [#&#8203;59774](https://redirect.github.com/nodejs/node/pull/59774)
- \[[`a414c1eb51`](https://redirect.github.com/nodejs/node/commit/a414c1eb51)] - **repl**: fix REPL completion under unary expressions (Kingsword) [#&#8203;59744](https://redirect.github.com/nodejs/node/pull/59744)
- \[[`c206f8dd87`](https://redirect.github.com/nodejs/node/commit/c206f8dd87)] - **repl**: add isValidParentheses check before wrap input (Xuguang Mei) [#&#8203;59607](https://redirect.github.com/nodejs/node/pull/59607)
- \[[`0bf9775ee2`](https://redirect.github.com/nodejs/node/commit/0bf9775ee2)] - **sea**: implement sea.getAssetKeys() (Joyee Cheung) [#&#8203;59661](https://redirect.github.com/nodejs/node/pull/59661)
- \[[`bf26b478d8`](https://redirect.github.com/nodejs/node/commit/bf26b478d8)] - **sea**: allow using inspector command line flags with SEA (Joyee Cheung) [#&#8203;59568](https://redirect.github.com/nodejs/node/pull/59568)
- \[[`92128a8fe2`](https://redirect.github.com/nodejs/node/commit/92128a8fe2)] - **src**: use DictionaryTemplate for node\_url\_pattern (James M Snell) [#&#8203;59802](https://redirect.github.com/nodejs/node/pull/59802)
- \[[`bcb29fb84f`](https://redirect.github.com/nodejs/node/commit/bcb29fb84f)] - **src**: correctly report memory changes to V8 (Yaksh Bariya) [#&#8203;59623](https://redirect.github.com/nodejs/node/pull/59623)
- \[[`44c24657d3`](https://redirect.github.com/nodejs/node/commit/44c24657d3)] - **src**: fixup node\_messaging error handling (James M Snell) [#&#8203;59792](https://redirect.github.com/nodejs/node/pull/59792)
- \[[`2cd6a3b7ec`](https://redirect.github.com/nodejs/node/commit/2cd6a3b7ec)] - **src**: track async resources via pointers to stack-allocated handles (Anna Henningsen) [#&#8203;59704](https://redirect.github.com/nodejs/node/pull/59704)
- \[[`34d752586f`](https://redirect.github.com/nodejs/node/commit/34d752586f)] - **src**: fix build on NetBSD (Thomas Klausner) [#&#8203;59718](https://redirect.github.com/nodejs/node/pull/59718)
- \[[`15fa779ac5`](https://redirect.github.com/nodejs/node/commit/15fa779ac5)] - **src**: fix race on process exit and off thread CA loading (Chengzhong Wu) [#&#8203;59632](https://redirect.github.com/nodejs/node/pull/59632)
- \[[`15cbd3966a`](https://redirect.github.com/nodejs/node/commit/15cbd3966a)] - **src**: separate module.hasAsyncGraph and module.hasTopLevelAwait (Joyee Cheung) [#&#8203;59675](https://redirect.github.com/nodejs/node/pull/59675)
- \[[`88d1ca8990`](https://redirect.github.com/nodejs/node/commit/88d1ca8990)] - **src**: use non-deprecated Get/SetPrototype methods (Michaël Zasso) [#&#8203;59671](https://redirect.github.com/nodejs/node/pull/59671)
- \[[`56ac9a2d46`](https://redirect.github.com/nodejs/node/commit/56ac9a2d46)] - **src**: migrate WriteOneByte to WriteOneByteV2 (Chengzhong Wu) [#&#8203;59634](https://redirect.github.com/nodejs/node/pull/59634)
- \[[`3d88aa9f2f`](https://redirect.github.com/nodejs/node/commit/3d88aa9f2f)] - **src**: remove duplicate code (theanarkh) [#&#8203;59649](https://redirect.github.com/nodejs/node/pull/59649)
- \[[`0718a70b2a`](https://redirect.github.com/nodejs/node/commit/0718a70b2a)] - **src**: add name for more threads (theanarkh) [#&#8203;59601](https://redirect.github.com/nodejs/node/pull/59601)
- \[[`0379a8b254`](https://redirect.github.com/nodejs/node/commit/0379a8b254)] - **src**: remove JSONParser (Joyee Cheung) [#&#8203;59619](https://redirect.github.com/nodejs/node/pull/59619)
- \[[`90d0a1b2e9`](https://redirect.github.com/nodejs/node/commit/90d0a1b2e9)] - **src,sqlite**: refactor value conversion (Edy Silva) [#&#8203;59659](https://redirect.github.com/nodejs/node/pull/59659)
- \[[`5e025c7ca7`](https://redirect.github.com/nodejs/node/commit/5e025c7ca7)] - **stream**: replace manual function validation with validateFunction (방진혁) [#&#8203;59529](https://redirect.github.com/nodejs/node/pull/59529)
- \[[`155a999bed`](https://redirect.github.com/nodejs/node/commit/155a999bed)] - **test**: skip tests failing when run under root (Livia Medeiros) [#&#8203;59779](https://redirect.github.com/nodejs/node/pull/59779)
- \[[`6313706c69`](https://redirect.github.com/nodejs/node/commit/6313706c69)] - **test**: update WPT for urlpattern to [`cff1ac1`](https://redirect.github.com/nodejs/node/commit/cff1ac1123) (Node.js GitHub Bot) [#&#8203;59602](https://redirect.github.com/nodejs/node/pull/59602)
- \[[`41245ad4c7`](https://redirect.github.com/nodejs/node/commit/41245ad4c7)] - **test**: skip more sea tests on Linux ppc64le (Richard Lau) [#&#8203;59755](https://redirect.github.com/nodejs/node/pull/59755)
- \[[`df63d37ec4`](https://redirect.github.com/nodejs/node/commit/df63d37ec4)] - **test**: fix internet/test-dns (Michaël Zasso) [#&#8203;59660](https://redirect.github.com/nodejs/node/pull/59660)
- \[[`1f6c335e82`](https://redirect.github.com/nodejs/node/commit/1f6c335e82)] - **test**: mark test-inspector-network-fetch as flaky again (Joyee Cheung) [#&#8203;59640](https://redirect.github.com/nodejs/node/pull/59640)
- \[[`1798683df1`](https://redirect.github.com/nodejs/node/commit/1798683df1)] - **test**: skip test-fs-cp\* tests that are constantly failing on Windows (Joyee Cheung) [#&#8203;59637](https://redirect.github.com/nodejs/node/pull/59637)
- \[[`4c48ec09e5`](https://redirect.github.com/nodejs/node/commit/4c48ec09e5)] - **test**: deflake test-http-keep-alive-empty-line (Luigi Pinca) [#&#8203;59595](https://redirect.github.com/nodejs/node/pull/59595)
- \[[`dcdb259e85`](https://redirect.github.com/nodejs/node/commit/dcdb259e85)] - **test\_runner**: fix todo inheritance (Moshe Atlow) [#&#8203;59721](https://redirect.github.com/nodejs/node/pull/59721)
- \[[`24177973a2`](https://redirect.github.com/nodejs/node/commit/24177973a2)] - **test\_runner**: set mock timer's interval undefined (hotpineapple) [#&#8203;59479](https://redirec

</details>

---

### Configuration

📅 **Schedule**: Branch creation - "on Monday after 3am and before 10am" (UTC), Automerge - At any time (no schedule defined).

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about these updates again.

---

 - [x] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

---
### Documentation

Find out how to configure dependency updates in [MintMaker documentation](https://konflux-ci.dev/docs/mintmaker/user/) or see all available configuration options in [Renovate documentation](https://docs.renovatebot.com/configuration-options/).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6W119-->


---

## Reviews

### Review by @swadeley - Approved on November 10, 2025 at 06:02 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1293*
