---
type: pull_request
number: 42
title: "Update module github.com/google/cel-go to v0.26.1"
state: merged
author: red-hat-konflux
created: 2025-11-06T00:18:15Z
updated: 2025-11-11T14:20:00Z
closed: 2025-11-11T14:19:57Z
merged: 2025-11-11T14:19:57Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-google-cel-go-0.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/42
---

# Pull Request #42: Update module github.com/google/cel-go to v0.26.1

**Author**: @red-hat-konflux
**Created**: November 06, 2025 at 12:18 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-google-cel-go-0.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/google/cel-go](https://redirect.github.com/google/cel-go) | `v0.21.0` -> `v0.26.1` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fgoogle%2fcel-go/v0.26.1?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fgoogle%2fcel-go/v0.21.0/v0.26.1?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>google/cel-go (github.com/google/cel-go)</summary>

### [`v0.26.1`](https://redirect.github.com/google/cel-go/releases/tag/v0.26.1)

[Compare Source](https://redirect.github.com/google/cel-go/compare/v0.26.0...v0.26.1)

#### What's Changed

- Comprehension nesting limit typo, allow nesting limit validator to accept doubles as limits by [@&#8203;l46kok](https://redirect.github.com/l46kok) in [#&#8203;1196](https://redirect.github.com/google/cel-go/pull/1196)
- Minor compatibility fixes for google3-import. by [@&#8203;jnthntatum](https://redirect.github.com/jnthntatum) in [#&#8203;1198](https://redirect.github.com/google/cel-go/pull/1198)
- Bump the npm\_and\_yarn group across 1 directory with 2 updates by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;1197](https://redirect.github.com/google/cel-go/pull/1197)
- Init function bindings on environment init by [@&#8203;beldmian](https://redirect.github.com/beldmian) in [#&#8203;1199](https://redirect.github.com/google/cel-go/pull/1199)
- Support variable descriptions in the AI prompt template by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1205](https://redirect.github.com/google/cel-go/pull/1205)
- Add support for nested element type by [@&#8203;MisLink](https://redirect.github.com/MisLink) in [#&#8203;1190](https://redirect.github.com/google/cel-go/pull/1190)
- Support unwrapping unknown implementations of `proto.Message` by [@&#8203;srikrsna](https://redirect.github.com/srikrsna) in [#&#8203;1207](https://redirect.github.com/google/cel-go/pull/1207)

#### New Contributors

- [@&#8203;beldmian](https://redirect.github.com/beldmian) made their first contribution in [#&#8203;1199](https://redirect.github.com/google/cel-go/pull/1199)
- [@&#8203;MisLink](https://redirect.github.com/MisLink) made their first contribution in [#&#8203;1190](https://redirect.github.com/google/cel-go/pull/1190)
- [@&#8203;srikrsna](https://redirect.github.com/srikrsna) made their first contribution in [#&#8203;1207](https://redirect.github.com/google/cel-go/pull/1207)

**Full Changelog**: <https://github.com/google/cel-go/compare/v0.25.1...v0.26.1>

### [`v0.26.0`](https://redirect.github.com/google/cel-go/releases/tag/v0.26.0)

[Compare Source](https://redirect.github.com/google/cel-go/compare/v0.25.1...v0.26.0)

##### New Features ✨

- **Add support for global constants to constant folder** by [@&#8203;zeitgeist87](https://redirect.github.com/zeitgeist87) in [#&#8203;1180](https://redirect.github.com/google/cel-go/issues/1180)
- **Adding CEL Regex Extensions** by [@&#8203;maskri17](https://redirect.github.com/maskri17) in [#&#8203;1187](https://redirect.github.com/google/cel-go/issues/1187)
- **Sqrt func** by [@&#8203;haribalan](https://redirect.github.com/haribalan) in [#&#8203;1166](https://redirect.github.com/google/cel-go/issues/1166)
- **Add bazel rule to trigger cel tests and return policy metadata while creating CEL programs** by [@&#8203;aakash070](https://redirect.github.com/aakash070) in [#&#8203;1176](https://redirect.github.com/google/cel-go/issues/1176)
- **Create an util method to convert rpc status to eval status** by [@&#8203;ChinmayMadeshi](https://redirect.github.com/ChinmayMadeshi) in [#&#8203;1178](https://redirect.github.com/google/cel-go/issues/1178)
- **Add test runner option to configure a custom test suite parser** by [@&#8203;aakash070](https://redirect.github.com/aakash070) in [#&#8203;1189](https://redirect.github.com/google/cel-go/issues/1189)
- **Cost tracking for list operations** by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1192](https://redirect.github.com/google/cel-go/issues/1192)

##### Bug Fixes 🐛

- **Fix lastIndexOf behavior against an empty string in strings extension** by [@&#8203;l46kok](https://redirect.github.com/l46kok) in [#&#8203;1173](https://redirect.github.com/google/cel-go/issues/1173)
- **Fix container setting for cel test all types example in online REPL** by [@&#8203;l46kok](https://redirect.github.com/l46kok) in [#&#8203;1182](https://redirect.github.com/google/cel-go/issues/1182)
- **fix(checker): Correct Sprintf argument count** by [@&#8203;cuishuang](https://redirect.github.com/cuishuang) in [#&#8203;1185](https://redirect.github.com/google/cel-go/issues/1185)

##### Test Updates 🧪

- **fix test runner test cases** by [@&#8203;aakash070](https://redirect.github.com/aakash070) in [#&#8203;1170](https://redirect.github.com/google/cel-go/issues/1170)
- **Update test runner to avoid using flags when not necessary** by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1174](https://redirect.github.com/google/cel-go/issues/1174)
- **add support for handling unknown expression ids in test result** by [@&#8203;aakash070](https://redirect.github.com/aakash070) in [#&#8203;1183](https://redirect.github.com/google/cel-go/issues/1183)
- **Test case for aliasing container imports** by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1193](https://redirect.github.com/google/cel-go/issues/1193)

##### Documentation 📚

- **Add documentation for YAML quirks in celpolicy** by [@&#8203;jnthntatum](https://redirect.github.com/jnthntatum) in [#&#8203;1186](https://redirect.github.com/google/cel-go/issues/1186)

##### Dependency Updates ⬆️

- **Bump the npm\_and\_yarn group across 1 directory with 2 updates** by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;1188](https://redirect.github.com/google/cel-go/issues/1188)

### [`v0.25.1`](https://redirect.github.com/google/cel-go/compare/v0.25.0...v0.25.1)

[Compare Source](https://redirect.github.com/google/cel-go/compare/v0.25.0...v0.25.1)

### [`v0.25.0`](https://redirect.github.com/google/cel-go/releases/tag/v0.25.0)

[Compare Source](https://redirect.github.com/google/cel-go/compare/v0.24.1...v0.25.0)

#### Features & Enhancements

This release introduces features for richer configuration-based CEL, AI prompt generation from config files, additional documentation, and 3x performance when evaluating traced / state-tracking expressions. This release also introduces a unit test runner framwork.

[#&#8203;1141](https://redirect.github.com/google/cel-go/pull/1141): Expose extension option factory as a public method

[#&#8203;1143](https://redirect.github.com/google/cel-go/pull/1143): Add a new compiler tool which can be used to compile CEL expressions and policies using serialized environment

[#&#8203;1151](https://redirect.github.com/google/cel-go/pull/1151): Lightweight observable evaluation

[#&#8203;1155](https://redirect.github.com/google/cel-go/pull/1155): Utilities for formatting and parsing documentation strings

[#&#8203;1156](https://redirect.github.com/google/cel-go/pull/1156): Support for documentation and example strings in CEL environments

[#&#8203;1158](https://redirect.github.com/google/cel-go/pull/1158): Re-export interpreter.AttributePattern in package cel.

[#&#8203;1159](https://redirect.github.com/google/cel-go/pull/1159): Document the standard library macros and functions

[#&#8203;1160](https://redirect.github.com/google/cel-go/pull/1160): Prompt generation for AI-assisted authoring based on a CEL environment

[#&#8203;1117](https://redirect.github.com/google/cel-go/pull/1117): Add LateFunctionBinding declaration and fix constant folding

[#&#8203;1163](https://redirect.github.com/google/cel-go/pull/1163): Initialize stateful observers prior to evaluation

[#&#8203;1164](https://redirect.github.com/google/cel-go/pull/1164): Unparse Expr values to strings

[#&#8203;1149](https://redirect.github.com/google/cel-go/pull/1149): Add test runner library

[#&#8203;1167](https://redirect.github.com/google/cel-go/pull/1167): REPL: Add an extension option for two var comprehensions

#### Fixes

Several fixes were implemented, including updating strings.format to better adhere to the specification, correcting constant folding logic alongside the late binding feature, removing a non-functional check in test code, and adding argument count validation for optFieldSelect.

[#&#8203;1133](https://redirect.github.com/google/cel-go/pull/1133): Update strings.format to adhere to the specification

[#&#8203;1117](https://redirect.github.com/google/cel-go/pull/1117): Add LateFunctionBinding declaration and fix constant folding

[#&#8203;1161](https://redirect.github.com/google/cel-go/pull/1161): Remove non-functional optional check in test-only selection

[#&#8203;1168](https://redirect.github.com/google/cel-go/pull/1168): Check arg count when validating optFieldSelect

#### Refactoring & Internal Improvements

General refactoring was performed across the codebase. Coverage and comments for Activation methods were improved. The test runner library was refactored to create options from flags and improve code structure.

[#&#8203;1145](https://redirect.github.com/google/cel-go/pull/1145): Refactoring changes

[#&#8203;1150](https://redirect.github.com/google/cel-go/pull/1150): Additional comments and coverage for Activation methods

[#&#8203;1165](https://redirect.github.com/google/cel-go/pull/1165): Refactoring changes to create a test runner option from passed flags, correct indentation and add package level comment for test

##### Documentation

Documentation was enhanced, including updates to the NativeTypes documentation regarding the cel tag, adding documentation for the optional library, and documenting the standard library functions/macros as part of the documentation string feature.

[#&#8203;1148](https://redirect.github.com/google/cel-go/pull/1148): Update NativeTypes doc to reflect how to enable cel tag

[#&#8203;1155](https://redirect.github.com/google/cel-go/pull/1155): Utilities for formatting and parsing documentation strings

[#&#8203;1156](https://redirect.github.com/google/cel-go/pull/1156): Support for documentation and example strings in CEL environments

[#&#8203;1159](https://redirect.github.com/google/cel-go/pull/1159): Document the standard library macros and functions

[#&#8203;1162](https://redirect.github.com/google/cel-go/pull/1162): Document optional library and increase docs coverage

##### Build System

Configuration fixes were made for Bzlmod compatibility.

[#&#8203;1146](https://redirect.github.com/google/cel-go/pull/1146): Bzlmod configuration fixes

##### Type System

Type formatting was updated to correctly handle type parameters.

[#&#8203;1154](https://redirect.github.com/google/cel-go/pull/1154): Update type formatting for type params

### [`v0.24.1`](https://redirect.github.com/google/cel-go/releases/tag/v0.24.1)

[Compare Source](https://redirect.github.com/google/cel-go/compare/v0.24.0...v0.24.1)

#### Fixes

- Separate unnest optimization from composer to capture type info \[[#&#8203;1138](https://redirect.github.com/google/cel-go/issues/1138)]

**Full Changelog**: <https://github.com/google/cel-go/compare/v0.24.0...v0.24.1>

### [`v0.24.0`](https://redirect.github.com/google/cel-go/releases/tag/v0.24.0)

[Compare Source](https://redirect.github.com/google/cel-go/compare/v0.23.2...v0.24.0)

Support for subsetting CEL standard library and serialization of CEL environments to YAML.

CEL is an official Google product \[[#&#8203;1122](https://redirect.github.com/google/cel-go/issues/1122)]

#### Features

- Helper methods for subsetting function overloads \[[#&#8203;1120](https://redirect.github.com/google/cel-go/issues/1120)]
- Introduce cel package aliases for Activation \[[#&#8203;1123](https://redirect.github.com/google/cel-go/issues/1123)]
- Canonical environment description and stdlib subsetting \[[#&#8203;1125](https://redirect.github.com/google/cel-go/issues/1125)]
- Support for cel.Env conversion to YAML-serializable config \[[#&#8203;1128](https://redirect.github.com/google/cel-go/issues/1128)]
- Option to configure CEL via env.Config object \[[#&#8203;1129](https://redirect.github.com/google/cel-go/issues/1129)]
- Support for feature flags and validators in env.Config \[[#&#8203;1132](https://redirect.github.com/google/cel-go/issues/1132)]
- Add k8s custom policy tag handler for test \[[#&#8203;1121](https://redirect.github.com/google/cel-go/issues/1121)]

#### Fixes

- ContextEval support for Unknowns \[[#&#8203;1126](https://redirect.github.com/google/cel-go/issues/1126)]
- Fix godoc formatting for Lists and OptionalTypes functions \[[#&#8203;1127](https://redirect.github.com/google/cel-go/issues/1127)]
- Default enable DefaultUTCTimeZone \[[#&#8203;1130](https://redirect.github.com/google/cel-go/issues/1130)]
- Support for splitting nested branching operators within policies \[[#&#8203;1136](https://redirect.github.com/google/cel-go/issues/1136)]

#### New Contributors

- [@&#8203;chaewonkong](https://redirect.github.com/chaewonkong) made their first contribution in [#&#8203;1127](https://redirect.github.com/google/cel-go/issues/1127)

**Full Changelog**: <https://github.com/google/cel-go/compare/v0.23.2...v0.24.0>

### [`v0.23.2`](https://redirect.github.com/google/cel-go/releases/tag/v0.23.2)

[Compare Source](https://redirect.github.com/google/cel-go/compare/v0.23.1...v0.23.2)

Corrects one remaining issue for cost computations from the v0.23.0 releases

#### Fixes

- Bump github.com/golang/glog from 1.0.0 to 1.2.4 in /codelab in the go\_modules group across 1 directory by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1115](https://redirect.github.com/google/cel-go/pull/1115)
- Minor update on cost order by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1119](https://redirect.github.com/google/cel-go/pull/1119)

**Full Changelog**: <https://github.com/google/cel-go/compare/v0.23.1...v0.23.2>

### [`v0.23.1`](https://redirect.github.com/google/cel-go/releases/tag/v0.23.1)

[Compare Source](https://redirect.github.com/google/cel-go/compare/v0.23.0...v0.23.1)

Minor release to address cost tracking and size estimation \[[#&#8203;1113](https://redirect.github.com/google/cel-go/issues/1113)]

**Full Changelog**: <https://github.com/google/cel-go/compare/v0.23.0...v0.23.1>

### [`v0.23.0`](https://redirect.github.com/google/cel-go/releases/tag/v0.23.0)

[Compare Source](https://redirect.github.com/google/cel-go/compare/v0.22.1...v0.23.0)

#### Features

- First and last element in list support \[[#&#8203;1067](https://redirect.github.com/google/cel-go/issues/1067)]
- Add support for typed conformance tests. \[[#&#8203;1089](https://redirect.github.com/google/cel-go/issues/1089)]
- Add syntax for escaped field selectors. \[[#&#8203;1002](https://redirect.github.com/google/cel-go/issues/1002)]
- Add optional.unwrap() / .unwrapOpt() function \[[#&#8203;1103](https://redirect.github.com/google/cel-go/issues/1103)]
- Cost tracking for two-variable comprehensions and bindings \[[#&#8203;1104](https://redirect.github.com/google/cel-go/issues/1104)]

#### Fixes

PR [#&#8203;1099](https://redirect.github.com/google/cel-go/issues/1099) enables a change in the internal variable name used for comprehension result accumulation. This change may break some tests which inspect the AST contents in text form; however, will not break any existing uses of CEL during parse, check, or evaluation.

- Improve policy compiler error message for incompatible outputs. \[[#&#8203;1082](https://redirect.github.com/google/cel-go/issues/1082)]
- Fix partial evaluation with the comprehension folder objects \[[#&#8203;1084](https://redirect.github.com/google/cel-go/issues/1084)]
- Introduce versioning options to all extensions \[[#&#8203;1075](https://redirect.github.com/google/cel-go/issues/1075)]
- Fix a crash in mismatched output check for nested rules \[[#&#8203;1086](https://redirect.github.com/google/cel-go/issues/1086)]
- improve debug output to properly quote byte strings \[[#&#8203;1088](https://redirect.github.com/google/cel-go/issues/1088)]
- Fix two-variable comprehension pruning \[[#&#8203;1083](https://redirect.github.com/google/cel-go/issues/1083)]
- Replace checks for valid UTF-8 in strings with go-maintained calls \[[#&#8203;1094](https://redirect.github.com/google/cel-go/issues/1094)]
- Policy nested rule fix \[[#&#8203;1092](https://redirect.github.com/google/cel-go/issues/1092)]
- Address non-const format string lint findings \[[#&#8203;1096](https://redirect.github.com/google/cel-go/issues/1096)]
- Fix typos in ext/README.md \[[#&#8203;1098](https://redirect.github.com/google/cel-go/issues/1098)]
- Add option to use inaccessible accumulator var \[[#&#8203;1097](https://redirect.github.com/google/cel-go/issues/1097)]
- Add test cases for `string.format` covering various edge cases \[[#&#8203;1101](https://redirect.github.com/google/cel-go/issues/1101)]
- Add base\_config and partial\_config files under restricted\_destination testdata \[[#&#8203;1106](https://redirect.github.com/google/cel-go/issues/1106)]
- Default enable using hidden accumulator name \[[#&#8203;1099](https://redirect.github.com/google/cel-go/issues/1099)]
- Update PruneAst to support constants of optional type \[[#&#8203;1109](https://redirect.github.com/google/cel-go/issues/1109)]

#### New Contributors

- [@&#8203;bigkevmcd](https://redirect.github.com/bigkevmcd) made their first contribution in [#&#8203;1067](https://redirect.github.com/google/cel-go/pull/1067)
- [@&#8203;hudlow](https://redirect.github.com/hudlow) made their first contribution in [#&#8203;1088](https://redirect.github.com/google/cel-go/pull/1088)
- [@&#8203;dbuduev](https://redirect.github.com/dbuduev) made their first contribution in [#&#8203;1098](https://redirect.github.com/google/cel-go/pull/1098)
- [@&#8203;zeitgeist87](https://redirect.github.com/zeitgeist87) made their first contribution in [#&#8203;1101](https://redirect.github.com/google/cel-go/pull/1101)

**Full Changelog**: <https://github.com/google/cel-go/compare/v0.22.1...v0.23.0>

### [`v0.22.1`](https://redirect.github.com/google/cel-go/releases/tag/v0.22.1)

[Compare Source](https://redirect.github.com/google/cel-go/compare/v0.22.0...v0.22.1)

#### Fixes

- Additional hardening on legacy macros \[[#&#8203;1064](https://redirect.github.com/google/cel-go/pull/1064)]
- Additional nil-safety checks with corresponding test updates \[[#&#8203;1073](https://redirect.github.com/google/cel-go/pull/1073)]
- Add two-variable comprehension support to cel-policy \[[#&#8203;1074](https://redirect.github.com/google/cel-go/pull/1074)]
- Fix optional test to short-circuit \[[#&#8203;1076](https://redirect.github.com/google/cel-go/pull/1076)]
- Fix nil-type when two-var comprehension has a dyn range \[[#&#8203;1077](https://redirect.github.com/google/cel-go/pull/1077)]

#### New Contributors

- [@&#8203;aakash070](https://redirect.github.com/aakash070) made their first contribution in [#&#8203;1069](https://redirect.github.com/google/cel-go/pull/1069)

**Full Changelog**: <https://github.com/google/cel-go/compare/v0.22.0...v0.22.1>

### [`v0.22.0`](https://redirect.github.com/google/cel-go/releases/tag/v0.22.0)

[Compare Source](https://redirect.github.com/google/cel-go/compare/v0.21.0...v0.22.0)

#### What's Changed

##### Core CEL

- Add list flatten() by [@&#8203;fabriziosestito](https://redirect.github.com/fabriziosestito) in [#&#8203;980](https://redirect.github.com/google/cel-go/pull/980)
- Add sets and lists extensions to REPL by [@&#8203;l46kok](https://redirect.github.com/l46kok) in [#&#8203;1005](https://redirect.github.com/google/cel-go/pull/1005)
- CEL Spec and Proto Updates by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1011](https://redirect.github.com/google/cel-go/pull/1011)
- Implement native conformance test runner by [@&#8203;jcking](https://redirect.github.com/jcking) in [#&#8203;1001](https://redirect.github.com/google/cel-go/pull/1001)
- Add support for no padding in base64 decode by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1017](https://redirect.github.com/google/cel-go/pull/1017)
- Allow configurable tag name overrides in native types by [@&#8203;matthchr](https://redirect.github.com/matthchr) in [#&#8203;1009](https://redirect.github.com/google/cel-go/pull/1009)
- Introduce helper trait types for lists and maps by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1026](https://redirect.github.com/google/cel-go/pull/1026)
- ext: add list.sort() by [@&#8203;cezar-guimaraes](https://redirect.github.com/cezar-guimaraes) in [#&#8203;1021](https://redirect.github.com/google/cel-go/pull/1021)
- Foldable Maps and Lists by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;995](https://redirect.github.com/google/cel-go/pull/995)
- Update REPL examples by [@&#8203;jnthntatum](https://redirect.github.com/jnthntatum) in [#&#8203;1028](https://redirect.github.com/google/cel-go/pull/1028)
- Interop foldable maps and lists with map mutation helper by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1029](https://redirect.github.com/google/cel-go/pull/1029)
- Update the Go AST representation to handle a second iteration variable by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1031](https://redirect.github.com/google/cel-go/pull/1031)
- Runtime support for two-variable comprehensions by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1032](https://redirect.github.com/google/cel-go/pull/1032)
- Two-variable comprehension support by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1034](https://redirect.github.com/google/cel-go/pull/1034)
- Update tag-based parsing to use lambda for additional customization by [@&#8203;matthchr](https://redirect.github.com/matthchr) in [#&#8203;1039](https://redirect.github.com/google/cel-go/pull/1039)
- Fix string reverse member overload name by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1045](https://redirect.github.com/google/cel-go/pull/1045)
- Remove unused server directory by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1041](https://redirect.github.com/google/cel-go/pull/1041)
- Add functions to the lists extension. by [@&#8203;seirl](https://redirect.github.com/seirl) in [#&#8203;1037](https://redirect.github.com/google/cel-go/pull/1037)
- Rename strings version test by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1047](https://redirect.github.com/google/cel-go/pull/1047)
- Fix doc strings and version support on math extensions by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1046](https://redirect.github.com/google/cel-go/pull/1046)
- Upgrade cel-go to support bazel-mod by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1049](https://redirect.github.com/google/cel-go/pull/1049)
- Fix out of range error for non-negative string indexing offsets by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1052](https://redirect.github.com/google/cel-go/pull/1052)
- Expand visibility of  parser/gen package by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1057](https://redirect.github.com/google/cel-go/pull/1057)
- Ensure variables in comprehensions don't collide by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1062](https://redirect.github.com/google/cel-go/pull/1062)

##### Policy

- Support for typename import aliases by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;993](https://redirect.github.com/google/cel-go/pull/993)
- Error on condition-only match blocks by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1003](https://redirect.github.com/google/cel-go/pull/1003)
- Update context proto resolution to use type provider by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1013](https://redirect.github.com/google/cel-go/pull/1013)
- Support for context proto declarations by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1006](https://redirect.github.com/google/cel-go/pull/1006)
- CEL Policy Readme by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1025](https://redirect.github.com/google/cel-go/pull/1025)
- Support for cel.[@&#8203;block](https://redirect.github.com/block) during policy composition by [@&#8203;TristonianJones](https://redirect.github.com/TristonianJones) in [#&#8203;1056](https://redirect.github.com/google/cel-go/pull/1056)

#### New Contributors

- [@&#8203;fabriziosestito](https://redirect.github.com/fabriziosestito) made their first contribution in [#&#8203;980](https://redirect.github.com/google/cel-go/pull/980)
- [@&#8203;matthchr](https://redirect.github.com/matthchr) made their first contribution in [#&#8203;1009](https://redirect.github.com/google/cel-go/pull/1009)
- [@&#8203;cezar-guimaraes](https://redirect.github.com/cezar-guimaraes) made their first contribution in [#&#8203;1021](https://redirect.github.com/google/cel-go/pull/1021)
- [@&#8203;sreeram-venkitesh](https://redirect.github.com/sreeram-venkitesh) made their first contribution in [#&#8203;1035](https://redirect.github.com/google/cel-go/pull/1035)

**Full Changelog**: <https://github.com/google/cel-go/compare/v0.21.0...v0.22.0>

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6W119-->


---

## Reviews

### Review by @Hyperkid123 - Approved on November 11, 2025 at 02:19 PM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/42*
