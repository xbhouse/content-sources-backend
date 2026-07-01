---
type: pull_request
number: 1291
title: "Build: update dependency pretty-format to v30"
state: merged
author: red-hat-konflux
created: 2025-11-10T04:16:41Z
updated: 2025-11-10T16:27:46Z
closed: 2025-11-10T14:34:00Z
merged: 2025-11-10T14:34:00Z
base_branch: main
head_branch: konflux/mintmaker/main/major-jest-monorepo
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1291
---

# Pull Request #1291: Build: update dependency pretty-format to v30

**Author**: @red-hat-konflux
**Created**: November 10, 2025 at 04:16 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/major-jest-monorepo`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [pretty-format](https://redirect.github.com/jestjs/jest) ([source](https://redirect.github.com/jestjs/jest/tree/HEAD/packages/pretty-format)) | [`^29.7.0` -> `^30.0.0`](https://renovatebot.com/diffs/npm/pretty-format/29.7.0/30.2.0) | [![age](https://developer.mend.io/api/mc/badges/age/npm/pretty-format/30.2.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/npm/pretty-format/29.7.0/30.2.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

### Release Notes

<details>
<summary>jestjs/jest (pretty-format)</summary>

### [`v30.2.0`](https://redirect.github.com/jestjs/jest/blob/HEAD/CHANGELOG.md#3020)

[Compare Source](https://redirect.github.com/jestjs/jest/compare/v30.0.5...v30.2.0)

##### Chore & Maintenance

- `[*]` Update example repo for testing React Native projects ([#&#8203;15832](https://redirect.github.com/jestjs/jest/pull/15832))
- `[*]` Update `jest-watch-typeahead` to v3 ([#&#8203;15830](https://redirect.github.com/jestjs/jest/pull/15830))

### [`v30.0.5`](https://redirect.github.com/jestjs/jest/blob/HEAD/CHANGELOG.md#3005)

[Compare Source](https://redirect.github.com/jestjs/jest/compare/v30.0.2...v30.0.5)

##### Features

- `[jest-config]` Allow `testMatch` to take a string value
- `[jest-worker]` Let `workerIdleMemoryLimit` accept 0 to always restart worker child processes

##### Fixes

- `[expect]` Fix `bigint` error ([#&#8203;15702](https://redirect.github.com/jestjs/jest/pull/15702))

### [`v30.0.2`](https://redirect.github.com/jestjs/jest/blob/HEAD/CHANGELOG.md#3002)

[Compare Source](https://redirect.github.com/jestjs/jest/compare/v30.0.1...v30.0.2)

##### Fixes

- `[jest-matcher-utils]` Make 'deepCyclicCopyObject' safer by setting descriptors to a null-prototype object ([#&#8203;15689](https://redirect.github.com/jestjs/jest/pull/15689))
- `[jest-util]` Make garbage collection protection property writable ([#&#8203;15689](https://redirect.github.com/jestjs/jest/pull/15689))

### [`v30.0.1`](https://redirect.github.com/jestjs/jest/blob/HEAD/CHANGELOG.md#3001)

[Compare Source](https://redirect.github.com/jestjs/jest/compare/v30.0.0...v30.0.1)

##### Features

- `[jest-resolver]` Implement the `defaultAsyncResolver` ([#&#8203;15679](https://redirect.github.com/jestjs/jest/pull/15679))

##### Fixes

- `[jest-resolver]` Resolve builtin modules correctly ([#&#8203;15683](https://redirect.github.com/jestjs/jest/pull/15683))
- `[jest-environment-node, jest-util]` Avoid setting globals cleanup protection symbol when feature is off ([#&#8203;15684](https://redirect.github.com/jestjs/jest/pull/15684))

##### Chore & Maintenance

- `[*]` Remove and deprecate `jest-repl` package ([#&#8203;15673](https://redirect.github.com/jestjs/jest/pull/15673))
- `[jest-resolver]` Replace custom `isBuiltinModule` with node's `isBuiltin` ([#&#8203;15685](https://redirect.github.com/jestjs/jest/pull/15685))

### [`v30.0.0`](https://redirect.github.com/jestjs/jest/blob/HEAD/CHANGELOG.md#3000)

[Compare Source](https://redirect.github.com/jestjs/jest/compare/v29.7.0...v30.0.0)

##### Features

- `[*]` Renamed `globalsCleanupMode` to `globalsCleanup` and `--waitNextEventLoopTurnForUnhandledRejectionEvents` to `--waitForUnhandledRejections`
- `[expect]` Add `ArrayOf` asymmetric matcher for validating array elements. ([#&#8203;15567](https://redirect.github.com/jestjs/jest/pull/15567))
- `[babel-jest]` Add option `excludeJestPreset` to allow opting out of `babel-preset-jest` ([#&#8203;15164](https://redirect.github.com/jestjs/jest/pull/15164))
- `[expect]` Revert [#&#8203;15038](https://redirect.github.com/jestjs/jest/pull/15038) to fix `expect(fn).toHaveBeenCalledWith(expect.objectContaining(...))` when there are multiple calls ([#&#8203;15508](https://redirect.github.com/jestjs/jest/pull/15508))
- `[jest-circus, jest-cli, jest-config]` Add `waitNextEventLoopTurnForUnhandledRejectionEvents` flag to minimise performance impact of correct detection of unhandled promise rejections introduced in [#&#8203;14315](https://redirect.github.com/jestjs/jest/pull/14315) ([#&#8203;14681](https://redirect.github.com/jestjs/jest/pull/14681))
- `[jest-circus]` Add a `waitBeforeRetry` option to `jest.retryTimes` ([#&#8203;14738](https://redirect.github.com/jestjs/jest/pull/14738))
- `[jest-circus]` Add a `retryImmediately` option to `jest.retryTimes` ([#&#8203;14696](https://redirect.github.com/jestjs/jest/pull/14696))
- `[jest-circus, jest-jasmine2]` Allow `setupFilesAfterEnv` to export an async function ([#&#8203;10962](https://redirect.github.com/jestjs/jest/issues/10962))
- `[jest-circus, jest-test-result]` Add `startedAt` timestamp in `TestCaseResultObject` within `onTestCaseResult` ([#&#8203;15145](https://redirect.github.com/jestjs/jest/pull/15145))
- `[jest-cli]` Export `buildArgv` ([#&#8203;15310](https://redirect.github.com/facebook/jest/pull/15310))
- `[jest-config]` \[**BREAKING**] Add `mts` and `cts` to default `moduleFileExtensions` config ([#&#8203;14369](https://redirect.github.com/facebook/jest/pull/14369))
- `[jest-config]` \[**BREAKING**] Update `testMatch` and `testRegex` default option for supporting `mjs`, `cjs`, `mts`, and `cts` ([#&#8203;14584](https://redirect.github.com/jestjs/jest/pull/14584))
- `[jest-config]` Loads config file from provided path in `package.json` ([#&#8203;14044](https://redirect.github.com/facebook/jest/pull/14044))
- `[jest-config]` Allow loading `jest.config.cts` files ([#&#8203;14070](https://redirect.github.com/facebook/jest/pull/14070))
- `[jest-config]` Show `rootDir` in error message when a `preset` fails to load ([#&#8203;15194](https://redirect.github.com/jestjs/jest/pull/15194))
- `[jest-config]` Support loading TS config files using `esbuild-register` via docblock loader ([#&#8203;15190](https://redirect.github.com/jestjs/jest/pull/15190))
- `[jest-config]` Allow passing TS config loader options via docblock comment ([#&#8203;15234](https://redirect.github.com/jestjs/jest/pull/15234))
- `[jest-config]` If Node is running with type stripping enabled, do not require a TS loader ([#&#8203;15480](https://redirect.github.com/jestjs/jest/pull/15480))
- `[@jest/core]` Group together open handles with the same stack trace ([#&#8203;13417](https://redirect.github.com/jestjs/jest/pull/13417), & [#&#8203;14789](https://redirect.github.com/jestjs/jest/pull/14789))
- `[@jest/core]` Add `perfStats` to surface test setup overhead ([#&#8203;14622](https://redirect.github.com/jestjs/jest/pull/14622))
- `[@jest/core]` \[**BREAKING**] Changed `--filter` to accept an object with shape `{ filtered: Array<string> }` to match [documentation](https://jestjs.io/docs/cli#--filterfile) ([#&#8203;13319](https://redirect.github.com/jestjs/jest/pull/13319))
- `[@jest/core]` Support `--outputFile` option for [`--listTests`](https://jestjs.io/docs/cli#--listtests) ([#&#8203;14980](https://redirect.github.com/jestjs/jest/pull/14980))
- `[@jest/core]` Stringify Errors properly with `--json` flag ([#&#8203;15329](https://redirect.github.com/jestjs/jest/pull/15329))
- `[@jest/core, @&#8203;jest/test-sequencer]` \[**BREAKING**] Exposes `globalConfig` & `contexts` to `TestSequencer` ([#&#8203;14535](https://redirect.github.com/jestjs/jest/pull/14535), & [#&#8203;14543](https://redirect.github.com/jestjs/jest/pull/14543))
- `[jest-each]` Introduce `%$` option to add number of the test to its title ([#&#8203;14710](https://redirect.github.com/jestjs/jest/pull/14710))
- `[@jest/environment]` \[**BREAKING**] Remove deprecated `jest.genMockFromModule()` ([#&#8203;15042](https://redirect.github.com/jestjs/jest/pull/15042))
- `[@jest/environment]` \[**BREAKING**] Remove unnecessary defensive code ([#&#8203;15045](https://redirect.github.com/jestjs/jest/pull/15045))
- `[jest-environment-jsdom]` \[**BREAKING**] Upgrade JSDOM to v22 ([#&#8203;13825](https://redirect.github.com/jestjs/jest/pull/13825))
- `[@jest/environment-jsdom-abstract]` Introduce new package which abstracts over the `jsdom` environment, allowing usage of custom versions of JSDOM ([#&#8203;14717](https://redirect.github.com/jestjs/jest/pull/14717))
- `[jest-environment-node]` Update jest environment with dispose symbols `Symbol` ([#&#8203;14888](https://redirect.github.com/jestjs/jest/pull/14888) & [#&#8203;14909](https://redirect.github.com/jestjs/jest/pull/14909))
- `[expect, @&#8203;jest/expect]` \[**BREAKING**] Add type inference for function parameters in `CalledWith` assertions ([#&#8203;15129](https://redirect.github.com/facebook/jest/pull/15129))
- `[@jest/expect-utils]` Properly compare all types of `TypedArray`s ([#&#8203;15178](https://redirect.github.com/facebook/jest/pull/15178))
- `[@jest/fake-timers]` \[**BREAKING**] Upgrade `@sinonjs/fake-timers` to v13 ([#&#8203;14544](https://redirect.github.com/jestjs/jest/pull/14544) & [#&#8203;15470](https://redirect.github.com/jestjs/jest/pull/15470))
- `[@jest/fake-timers]` Exposing new modern timers function `advanceTimersToFrame()` which advances all timers by the needed milliseconds to execute callbacks currently scheduled with `requestAnimationFrame` ([#&#8203;14598](https://redirect.github.com/jestjs/jest/pull/14598))
- `[jest-matcher-utils]` Add `SERIALIZABLE_PROPERTIES` to allow custom serialization of objects ([#&#8203;14893](https://redirect.github.com/jestjs/jest/pull/14893))
- `[jest-mock]` Add support for the Explicit Resource Management proposal to use the `using` keyword with `jest.spyOn(object, methodName)` ([#&#8203;14895](https://redirect.github.com/jestjs/jest/pull/14895))
- `[jest-reporters]` Add support for [DEC mode 2026](https://gist.github.com/christianparpart/d8a62cc1ab659194337d73e399004036) ([#&#8203;15008](https://redirect.github.com/jestjs/jest/pull/15008))
- `[jest-resolver]` Support `file://` URLs as paths ([#&#8203;15154](https://redirect.github.com/jestjs/jest/pull/15154))
- `[jest-resolve,jest-runtime,jest-resolve-dependencies]` Pass the conditions when resolving stub modules ([#&#8203;15489](https://redirect.github.com/jestjs/jest/pull/15489))
- `[jest-runtime]` Exposing new modern timers function `jest.advanceTimersToFrame()` from `@jest/fake-timers` ([#&#8203;14598](https://redirect.github.com/jestjs/jest/pull/14598))
- `[jest-runtime]` Support `import.meta.filename` and `import.meta.dirname` (available from [Node 20.11](https://nodejs.org/en/blog/release/v20.11.0)) ([#&#8203;14854](https://redirect.github.com/jestjs/jest/pull/14854))
- `[jest-runtime]` Support `import.meta.resolve` ([#&#8203;14930](https://redirect.github.com/jestjs/jest/pull/14930))
- `[jest-runtime]` \[**BREAKING**] Make it mandatory to pass `globalConfig` to the `Runtime` constructor ([#&#8203;15044](https://redirect.github.com/jestjs/jest/pull/15044))
- `[jest-runtime]` Add `unstable_unmockModule` ([#&#8203;15080](https://redirect.github.com/jestjs/jest/pull/15080))
- `[jest-runtime]` Add `onGenerateMock` transformer callback for auto generated callbacks ([#&#8203;15433](https://redirect.github.com/jestjs/jest/pull/15433) & [#&#8203;15482](https://redirect.github.com/jestjs/jest/pull/15482))
- `[jest-runtime]` \[**BREAKING**] Use `vm.compileFunction` over `vm.Script` ([#&#8203;15461](https://redirect.github.com/jestjs/jest/pull/15461))
- `[@jest/schemas]` Upgrade `@sinclair/typebox` to v0.34 ([#&#8203;15450](https://redirect.github.com/jestjs/jest/pull/15450))
- `[@jest/types]` `test.each()`: Accept a readonly (`as const`) table properly ([#&#8203;14565](https://redirect.github.com/jestjs/jest/pull/14565))
- `[@jest/types]` Improve argument type inference passed to `test` and `describe` callback functions from `each` tables ([#&#8203;14920](https://redirect.github.com/jestjs/jest/pull/14920))
- `[jest-snapshot]` \[**BREAKING**] Add support for [Error causes](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Error/cause) in snapshots ([#&#8203;13965](https://redirect.github.com/facebook/jest/pull/13965))
- `[jest-snapshot]` Support Prettier 3 ([#&#8203;14566](https://redirect.github.com/facebook/jest/pull/14566))
- `[@jest/util-snapshot]` Extract utils used by tooling from `jest-snapshot` into its own package ([#&#8203;15095](https://redirect.github.com/facebook/jest/pull/15095))
- `[pretty-format]` \[**BREAKING**] Do not render empty string children (`''`) in React plugin ([#&#8203;14470](https://redirect.github.com/facebook/jest/pull/14470))

##### Fixes

- `[expect]` Show `AggregateError` to display ([#&#8203;15346](https://redirect.github.com/facebook/jest/pull/15346))
- `[*]` Replace `exit` with `exit-x` ([#&#8203;15399](https://redirect.github.com/jestjs/jest/pull/15399))
- `[babel-plugin-jest-hoist]` Use `denylist` instead of the deprecated `blacklist` for Babel 8 support ([#&#8203;14109](https://redirect.github.com/jestjs/jest/pull/14109))
- `[babel-plugin-jest-hoist]` Do not rely on buggy Babel behaviour ([#&#8203;15415](https://redirect.github.com/jestjs/jest/pull/15415))
- `[expect]` Check error instance type for `toThrow/toThrowError` ([#&#8203;14576](https://redirect.github.com/jestjs/jest/pull/14576))
- `[expect]` Improve diff for failing `expect.objectContaining` ([#&#8203;15038](https://redirect.github.com/jestjs/jest/pull/15038))
- `[expect]` Use `Array.isArray` to check if an array is an `Array` ([#&#8203;15101](https://redirect.github.com/jestjs/jest/pull/15101))
- `[expect]` Fix Error `cause` assertion errors ([#&#8203;15339](https://redirect.github.com/jestjs/jest/pull/15339))
- `[jest-changed-files]` Print underlying errors when VCS commands fail ([#&#8203;15052](https://redirect.github.com/jestjs/jest/pull/15052))
- `[jest-changed-files]` Abort `sl root` call if output resembles a steam locomotive ([#&#8203;15053](https://redirect.github.com/jestjs/jest/pull/15053))
- `[jest-circus]` \[**BREAKING**] Prevent false test failures caused by promise rejections handled asynchronously ([#&#8203;14315](https://redirect.github.com/jestjs/jest/pull/14315))
- `[jest-circus]` Replace recursive `makeTestResults` implementation with iterative one ([#&#8203;14760](https://redirect.github.com/jestjs/jest/pull/14760))
- `[jest-circus]` Omit `expect.hasAssertions()` errors if a test already has errors ([#&#8203;14866](https://redirect.github.com/jestjs/jest/pull/14866))
- `[jest-circus, jest-expect, jest-snapshot]` Pass `test.failing` tests when containing failing snapshot matchers ([#&#8203;14313](https://redirect.github.com/jestjs/jest/pull/14313))
- `[jest-circus]` Concurrent tests now emit jest circus events at the correct point and in the expected order. ([#&#8203;15381](https://redirect.github.com/jestjs/jest/pull/15381))
- `[jest-cli]` \[**BREAKING**] Validate CLI flags that require arguments receives them ([#&#8203;14783](https://redirect.github.com/jestjs/jest/pull/14783))
- `[jest-config]` Make sure to respect `runInBand` option ([#&#8203;14578](https://redirect.github.com/jestjs/jest/pull/14578))
- `[jest-config]` Support `testTimeout` in project config ([#&#8203;14697](https://redirect.github.com/jestjs/jest/pull/14697))
- `[jest-config]` Support `coverageReporters` in project config ([#&#8203;14697](https://redirect.github.com/jestjs/jest/pull/14830))
- `[jest-config]` Allow `reporters` in project config ([#&#8203;14768](https://redirect.github.com/jestjs/jest/pull/14768))
- `[jest-config]` Allow Node16/NodeNext/Bundler `moduleResolution` in project's tsconfig ([#&#8203;14739](https://redirect.github.com/jestjs/jest/pull/14739))
- `[@jest/create-cache-key-function]` Correct the return type of `createCacheKey` ([#&#8203;15159](https://redirect.github.com/jestjs/jest/pull/15159))
- `[jest-each]` Allow `$keypath` templates with `null` or `undefined` values ([#&#8203;14831](https://redirect.github.com/jestjs/jest/pull/14831))
- `[@jest/expect-utils]` Fix comparison of `DataView` ([#&#8203;14408](https://redirect.github.com/jestjs/jest/pull/14408))
- `[@jest/expect-utils]` \[**BREAKING**] exclude non-enumerable in object matching ([#&#8203;14670](https://redirect.github.com/jestjs/jest/pull/14670))
- `[@jest/expect-utils]` Fix comparison of `URL` ([#&#8203;14672](https://redirect.github.com/jestjs/jest/pull/14672))
- `[@jest/expect-utils]` Check `Symbol` properties in equality ([#&#8203;14688](https://redirect.github.com/jestjs/jest/pull/14688))
- `[@jest/expect-utils]` Catch circular references within arrays when matching objects ([#&#8203;14894](https://redirect.github.com/jestjs/jest/pull/14894))
- `[@jest/expect-utils]` Fix not addressing to Sets and Maps as objects without keys ([#&#8203;14873](https://redirect.github.com/jestjs/jest/pull/14873))
- `[jest-haste-map]` Fix errors or clobbering with multiple `hasteImplModulePath`s ([#&#8203;15522](https://redirect.github.com/jestjs/jest/pull/15522))
- `[jest-leak-detector]` Make leak-detector more aggressive when running GC ([#&#8203;14526](https://redirect.github.com/jestjs/jest/pull/14526))
- `[jest-runtime]` Properly handle re-exported native modules in ESM via CJS ([#&#8203;14589](https://redirect.github.com/jestjs/jest/pull/14589))
- `[jest-runtime]` Refactor `_importCoreModel` so required core module is consistent if modified while loading ([#&#8203;15077](https://redirect.github.com/jestjs/jest/issues/15077))
- `[jest-schemas, jest-types]` \[**BREAKING**] Fix type of `testFailureExitCode` config option([#&#8203;15232](https://redirect.github.com/jestjs/jest/pull/15232))
- `[jest-util]` Make sure `isInteractive` works in a browser ([#&#8203;14552](https://redirect.github.com/jestjs/jest/pull/14552))
- `[pretty-format]` \[**BREAKING**] Print `ArrayBuffer` and `DataView` correctly ([#&#8203;14290](https://redirect.github.com/jestjs/jest/pull/14290))
- `[pretty-format]` Fixed a bug where "anonymous custom elements" were not being printed as expected. ([#&#8203;15138](https://redirect.github.com/jestjs/jest/pull/15138))
- `[jest-cli]` When specifying paths on the command line, only match against the relative paths of the test files ([#&#8203;12519](https://redirect.github.com/jestjs/jest/pull/12519))
  - \[**BREAKING**] Changes `testPathPattern` configuration option to `testPathPatterns`, which now takes a list of patterns instead of the regex.
  - \[**BREAKING**] `--testPathPattern` is now `--testPathPatterns`
  - \[**BREAKING**] Specifying `testPathPatterns` when programmatically calling `watch` must be specified as `new TestPathPatterns(patterns)`, where `TestPathPatterns` can be imported from `@jest/pattern`
- `[jest-reporters, jest-runner]` Unhandled errors without stack get correctly logged to console ([#&#8203;14619](https://redirect.github.com/jestjs/jest/pull/14619))
- `[jest-util]` Always load `mjs` files with `import` ([#&#8203;15447](https://redirect.github.com/jestjs/jest/pull/15447))
- `[jest-worker]` Properly handle a circular reference error when worker tries to send an assertion fails where either the expected or actual value is circular ([#&#8203;15191](https://redirect.github.com/jestjs/jest/pull/15191))
- `[jest-worker]` Properly handle a BigInt when worker tries to send an assertion fails where either the expected or actual value is BigInt ([#&#8203;15191](https://redirect.github.com/jestjs/jest/pull/15191))
- `[expect]` Resolve issue where `ObjectContaining` matched non-object values. [#&#8203;15463](https://redirect.github.com/jestjs/jest/pull/15463).
  - Adds a `conditional/check` to ensure the argument passed to `expect` is an object.
  - Add unit tests for new `ObjectContaining` behavior.
  - Remove `invalid/wrong` test case assertions for `ObjectContaining`.
- `[jest-worker]` Addresses incorrect state on exit ([#&#8203;15610](https://redirect.github.com/jestjs/jest/pull/15610))

##### Performance

- `[*]` \[**BREAKING**] Bundle all of Jest's modules into `index.js` ([#&#8203;12348](https://redirect.github.com/jestjs/jest/pull/12348), [#&#8203;14550](https://redirect.github.com/jestjs/jest/pull/14550) & [#&#8203;14661](https://redirect.github.com/jestjs/jest/pull/14661))
- `[jest-haste-map]` Only spawn one process to check for `watchman` installation ([#&#8203;14826](https://redirect.github.com/jestjs/jest/pull/14826))
- `[jest-runner]` Better cleanup `source-map-support` after test to resolve (minor) memory leak ([#&#8203;15233](https://redirect.github.com/jestjs/jest/pull/15233))
- `[jest-resolver]` Migrate `resolve` and `resolve.exports` to `unrs-resolver` ([#&#8203;15619](https://redirect.github.com/jestjs/jest/pull/15619))
- `[jest-circus, jest-environment-node, jest-repl, jest-runner, jest-util]` Cleanup global variables on environment teardown to reduce memory leaks ([#&#8203;15215](https://redirect.github.com/jestjs/jest/pull/15215) & [#&#8203;15636](https://redirect.github.com/jestjs/jest/pull/15636) & [#&#8203;15643](https://redirect.github.com/jestjs/jest/pull/15643))

##### Chore & Maintenance

- `[jest-environment-jsdom, jest-environment-jsdom-abstract]` Increased version of jsdom to `^26.0.0` ([#&#8203;15325](https://redirect.github.com/jestjs/jest/issues/15325)[CVE-2024-37890](https://nvd.nist.gov/vuln/detail/CVE-2024-37890))
- `[*]` Increase version of `micromatch` to `^4.0.7` ([#&#8203;15082](https://redirect.github.com/jestjs/jest/pull/15082))
- `[*]` \[**BREAKING**] Drop support for Node.js versions 14, 16, 19, 21 and 23 ([#&#8203;14460](https://redirect.github.com/jestjs/jest/pull/14460), [#&#8203;15118](https://redirect.github.com/jestjs/jest/pull/15118), [#&#8203;15623](https://redirect.github.com/jestjs/jest/pull/15623), [#&#8203;15640](https://redirect.github.com/jestjs/jest/pull/15640))
- `[*]` \[**BREAKING**] Drop support for `typescript@4.3`, minimum version is now `5.4` ([#&#8203;14542](https://redirect.github.com/jestjs/jest/pull/14542), [#&#8203;15621](https://redirect.github.com/jestjs/jest/pull/15621))
- `[*]` Depend on exact versions of monorepo dependencies instead of `^` range ([#&#8203;14553](https://redirect.github.com/jestjs/jest/pull/14553))
- `[*]` \[**BREAKING**] Add ESM wrapper for all of Jest's modules ([#&#8203;14661](https://redirect.github.com/jestjs/jest/pull/14661))
- `[*]` \[**BREAKING**] Upgrade to `glob@10` ([#&#8203;14509](https://redirect.github.com/jestjs/jest/pull/14509))
- `[*]` Use `TypeError` over `Error` where appropriate ([#&#8203;14799](https://redirect.github.com/jestjs/jest/pull/14799))
- `[docs]` Fix typos in `CHANGELOG.md` and `packages/jest-validate/README.md` ([#&#8203;14640](https://redirect.github.com/jestjs/jest/pull/14640))
- `[docs]` Don't use alias matchers in docs ([#&#8203;14631](https://redirect.github.com/jestjs/jest/pull/14631))
- `[babel-jest, babel-preset-jest]` \[**BREAKING**] Increase peer dependency of `@babel/core` to `^7.11` ([#&#8203;14109](https://redirect.github.com/jestjs/jest/pull/14109))
- `[babel-jest, @&#8203;jest/transform]` Update `babel-plugin-istanbul` to v6 ([#&#8203;15156](https://redirect.github.com/jestjs/jest/pull/15156))
- `[babel-plugin-jest-hoist]` Move unnecessary `dependencies` to `devDependencies` ([#&#8203;15010](https://redirect.github.com/jestjs/jest/pull/15010))
- `[expect]` \[**BREAKING**] Remove `.toBeCalled()`, `.toBeCalledTimes()`, `.toBeCalledWith()`, `.lastCalledWith()`, `.nthCalledWith()`, `.toReturn()`, `.toReturnTimes()`, `.toReturnWith()`, `.lastReturnedWith()`, `.nthReturnedWith()` and `.toThrowError()` matcher aliases ([#&#8203;14632](https://redirect.github.com/jestjs/jest/pull/14632))
- `[jest-cli, jest-config, @&#8203;jest/types]` \[**BREAKING**] Remove deprecated `--init` argument ([#&#8203;14490](https://redirect.github.com/jestjs/jest/pull/14490))
- `[jest-config, @&#8203;jest/core, jest-util]` Upgrade `ci-info` ([#&#8203;14655](https://redirect.github.com/jestjs/jest/pull/14655))
- `[jest-mock]` \[**BREAKING**] Remove `MockFunctionMetadataType`, `MockFunctionMetadata` and `SpyInstance` types ([#&#8203;14621](https://redirect.github.com/jestjs/jest/pull/14621))
- `[@jest/reporters]` Upgrade `istanbul-lib-source-maps` ([#&#8203;14924](https://redirect.github.com/jestjs/jest/pull/14924))
- `[jest-schemas]` Upgrade `@sinclair/typebox` ([#&#8203;14775](https://redirect.github.com/jestjs/jest/pull/14775))
- `[jest-transform]` Upgrade `write-file-atomic` ([#&#8203;14274](https://redirect.github.com/jestjs/jest/pull/14274))
- `[jest-util]` Upgrade `picomatch` to v4 ([#&#8203;14653](https://redirect.github.com/jestjs/jest/pull/14653) & [#&#8203;14885](https://redirect.github.com/jestjs/jest/pull/14885))
- `[docs] Append to NODE_OPTIONS, not overwrite ([#&#8203;14730](https://redirect.github.com/jestjs/jest/pull/14730))`
- `[docs]` Updated `.toHaveBeenCalled()` documentation to correctly reflect its functionality ([#&#8203;14842](https://redirect.github.com/jestjs/jest/pull/14842))
- `[docs]` Link NestJS documentation on testing with Jest ([#&#8203;14940](https://redirect.github.com/jestjs/jest/pull/14940))
- `[docs]` `Revised documentation for .toHaveBeenCalled()` to accurately depict its functionality. ([#&#8203;14853](https://redirect.github.com/jestjs/jest/pull/14853))
- `[docs]` Removed ExpressJS reference link from documentation due to dead link ([#&#8203;15270](https://redirect.github.com/jestjs/jest/pull/15270))
- `[docs]` Correct broken links in docs ([#&#8203;15359](https://redirect.github.com/jestjs/jest/pull/15359))

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

### Review by @TenSt - Approved on November 10, 2025 at 02:33 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1291*
