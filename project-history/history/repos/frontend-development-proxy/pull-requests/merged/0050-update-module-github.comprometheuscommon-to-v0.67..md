---
type: pull_request
number: 50
title: "Update module github.com/prometheus/common to v0.67.2"
state: merged
author: red-hat-konflux
created: 2025-11-07T12:18:40Z
updated: 2025-11-11T13:11:53Z
closed: 2025-11-11T13:11:53Z
merged: 2025-11-11T13:11:53Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-prometheus-common-0.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/50
---

# Pull Request #50: Update module github.com/prometheus/common to v0.67.2

**Author**: @red-hat-konflux
**Created**: November 07, 2025 at 12:18 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-prometheus-common-0.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/prometheus/common](https://redirect.github.com/prometheus/common) | `v0.48.0` -> `v0.67.2` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fprometheus%2fcommon/v0.67.2?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fprometheus%2fcommon/v0.48.0/v0.67.2?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>prometheus/common (github.com/prometheus/common)</summary>

### [`v0.67.2`](https://redirect.github.com/prometheus/common/blob/HEAD/CHANGELOG.md#v0672--2025-10-28)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.67.1...v0.67.2)

### [`v0.67.1`](https://redirect.github.com/prometheus/common/blob/HEAD/CHANGELOG.md#v0671--2025-10-07)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.67.0...v0.67.1)

### [`v0.67.0`](https://redirect.github.com/prometheus/common/blob/HEAD/CHANGELOG.md#v0670--2025-10-07)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.66.1...v0.67.0)

### [`v0.66.1`](https://redirect.github.com/prometheus/common/blob/HEAD/CHANGELOG.md#v0661--2025-09-05)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.66.0...v0.66.1)

This release has no functional changes, it just drops the dependencies `github.com/grafana/regexp` and `go.uber.org/atomic` and replaces `gopkg.in/yaml.v2` with `go.yaml.in/yaml/v2` (a drop-in replacement).

##### What's Changed

- Revert "Use github.com/grafana/regexp instead of regexp" by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;835](https://redirect.github.com/prometheus/common/pull/835)
- Move to supported version of yaml parser by [@&#8203;dims](https://redirect.github.com/dims) in [#&#8203;834](https://redirect.github.com/prometheus/common/pull/834)
- Revert "Use go.uber.org/atomic instead of sync/atomic ([#&#8203;825](https://redirect.github.com/prometheus/common/issues/825))" by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;838](https://redirect.github.com/prometheus/common/pull/838)

**Full Changelog**: <https://github.com/prometheus/common/compare/v1.20.99...v0.66.1>

### [`v0.66.0`](https://redirect.github.com/prometheus/common/blob/HEAD/CHANGELOG.md#v0660--2025-09-02)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.65.0...v0.66.0)

##### ⚠️ Breaking Changes ⚠️

- A default-constructed TextParser will be invalid. It must have a valid `scheme` set, so users should use the NewTextParser function to create a valid TextParser. Otherwise parsing will panic with "Invalid name validation scheme requested: unset".

##### What's Changed

- model: add constants for type and unit labels. by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [#&#8203;801](https://redirect.github.com/prometheus/common/pull/801)

- model.ValidationScheme: Support encoding as YAML by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;799](https://redirect.github.com/prometheus/common/pull/799)

- fix(promslog): always print time.Duration values as go duration strings by [@&#8203;tjhop](https://redirect.github.com/tjhop) in [#&#8203;798](https://redirect.github.com/prometheus/common/pull/798)

- Add `ValidationScheme` methods `IsValidMetricName` and `IsValidLabelName` by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;806](https://redirect.github.com/prometheus/common/pull/806)

- Fix delimited proto not escaped correctly by [@&#8203;thampiotr](https://redirect.github.com/thampiotr) in [#&#8203;809](https://redirect.github.com/prometheus/common/pull/809)

- Decoder: Remove use of global name validation and add validation by [@&#8203;ywwg](https://redirect.github.com/ywwg) in [#&#8203;808](https://redirect.github.com/prometheus/common/pull/808)

- ValidationScheme implements pflag.Value and json.Marshaler/Unmarshaler interfaces by [@&#8203;juliusmh](https://redirect.github.com/juliusmh) in [#&#8203;807](https://redirect.github.com/prometheus/common/pull/807)

- expfmt: Add NewTextParser function by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;816](https://redirect.github.com/prometheus/common/pull/816)

- Enable the godot linter by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;821](https://redirect.github.com/prometheus/common/pull/821)

- Enable usestdlibvars linter by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;820](https://redirect.github.com/prometheus/common/pull/820)

- Enable unconvert linter by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;819](https://redirect.github.com/prometheus/common/pull/819)

- Enable the fatcontext linter by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;822](https://redirect.github.com/prometheus/common/pull/822)

- Enable gocritic linter by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;818](https://redirect.github.com/prometheus/common/pull/818)

- Use go.uber.org/atomic instead of sync/atomic by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;825](https://redirect.github.com/prometheus/common/pull/825)

- Enable revive rule unused-parameter by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;824](https://redirect.github.com/prometheus/common/pull/824)

- Enable revive rules by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;823](https://redirect.github.com/prometheus/common/pull/823)

- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;802](https://redirect.github.com/prometheus/common/pull/802)

- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;803](https://redirect.github.com/prometheus/common/pull/803)

- Sync .golangci.yml with prometheus/prometheus by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;817](https://redirect.github.com/prometheus/common/pull/817)

- ci: update upload-actions by [@&#8203;ywwg](https://redirect.github.com/ywwg) in [#&#8203;814](https://redirect.github.com/prometheus/common/pull/814)

- docs: fix typo in expfmt.Negotiate by [@&#8203;wmcram](https://redirect.github.com/wmcram) in [#&#8203;813](https://redirect.github.com/prometheus/common/pull/813)

- build(deps): bump golang.org/x/net from 0.40.0 to 0.41.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;800](https://redirect.github.com/prometheus/common/pull/800)

- build(deps): bump golang.org/x/net from 0.41.0 to 0.42.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;810](https://redirect.github.com/prometheus/common/pull/810)

- build(deps): bump github.com/stretchr/testify from 1.10.0 to 1.11.1 in /assets by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;826](https://redirect.github.com/prometheus/common/pull/826)

- build(deps): bump google.golang.org/protobuf from 1.36.6 to 1.36.8 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;830](https://redirect.github.com/prometheus/common/pull/830)

- build(deps): bump golang.org/x/net from 0.42.0 to 0.43.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;829](https://redirect.github.com/prometheus/common/pull/829)

- build(deps): bump github.com/stretchr/testify from 1.10.0 to 1.11.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;827](https://redirect.github.com/prometheus/common/pull/827)

##### New Contributors

- [@&#8203;aknuds1](https://redirect.github.com/aknuds1) made their first contribution in [#&#8203;799](https://redirect.github.com/prometheus/common/pull/799)
- [@&#8203;thampiotr](https://redirect.github.com/thampiotr) made their first contribution in [#&#8203;809](https://redirect.github.com/prometheus/common/pull/809)
- [@&#8203;wmcram](https://redirect.github.com/wmcram) made their first contribution in [#&#8203;813](https://redirect.github.com/prometheus/common/pull/813)
- [@&#8203;juliusmh](https://redirect.github.com/juliusmh) made their first contribution in [#&#8203;807](https://redirect.github.com/prometheus/common/pull/807)

### [`v0.65.0`](https://redirect.github.com/prometheus/common/releases/tag/v0.65.0)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.64.0...v0.65.0)

#### What's Changed

- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;789](https://redirect.github.com/prometheus/common/pull/789)
- Remove otlptranslator package by [@&#8203;ArthurSens](https://redirect.github.com/ArthurSens) in [#&#8203;791](https://redirect.github.com/prometheus/common/pull/791)
- feat(promslog): add Level() method to get slog.Level by [@&#8203;tjhop](https://redirect.github.com/tjhop) in [#&#8203;795](https://redirect.github.com/prometheus/common/pull/795)
- feat: Support negative duration in new function ParseDurationAllowNegative by [@&#8203;iamhalje](https://redirect.github.com/iamhalje) in [#&#8203;793](https://redirect.github.com/prometheus/common/pull/793)

#### New Contributors

- [@&#8203;iamhalje](https://redirect.github.com/iamhalje) made their first contribution in [#&#8203;793](https://redirect.github.com/prometheus/common/pull/793)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.64.0...v0.65.0>

### [`v0.64.0`](https://redirect.github.com/prometheus/common/releases/tag/v0.64.0)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.63.0...v0.64.0)

#### What's Changed

- Add deprecation notice to otlptranslator by [@&#8203;ArthurSens](https://redirect.github.com/ArthurSens) in [#&#8203;773](https://redirect.github.com/prometheus/common/pull/773)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;774](https://redirect.github.com/prometheus/common/pull/774)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;775](https://redirect.github.com/prometheus/common/pull/775)
- Update Go by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;770](https://redirect.github.com/prometheus/common/pull/770)
- chore: Upgrade golangci-lint to v2 by [@&#8203;kakkoyun](https://redirect.github.com/kakkoyun) in [#&#8203;779](https://redirect.github.com/prometheus/common/pull/779)
- build(deps): bump golang.org/x/net from 0.37.0 to 0.38.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;777](https://redirect.github.com/prometheus/common/pull/777)
- build(deps): bump google.golang.org/protobuf from 1.36.5 to 1.36.6 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;776](https://redirect.github.com/prometheus/common/pull/776)
- promslog: Use the default timezone (again) by [@&#8203;beorn7](https://redirect.github.com/beorn7) in [#&#8203;739](https://redirect.github.com/prometheus/common/pull/739)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;787](https://redirect.github.com/prometheus/common/pull/787)
- build(deps): bump github.com/prometheus/client\_model from 0.6.1 to 0.6.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;784](https://redirect.github.com/prometheus/common/pull/784)
- build(deps): bump golang.org/x/oauth2 from 0.28.0 to 0.29.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;785](https://redirect.github.com/prometheus/common/pull/785)
- build(deps): bump golang.org/x/net from 0.38.0 to 0.39.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;786](https://redirect.github.com/prometheus/common/pull/786)
- refactor(promslog): make `NewNopLogger()` wrapper around `New()` by [@&#8203;tjhop](https://redirect.github.com/tjhop) in [#&#8203;783](https://redirect.github.com/prometheus/common/pull/783)
- build(deps): bump golang.org/x/oauth2 from 0.29.0 to 0.30.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;788](https://redirect.github.com/prometheus/common/pull/788)

#### New Contributors

- [@&#8203;kakkoyun](https://redirect.github.com/kakkoyun) made their first contribution in [#&#8203;779](https://redirect.github.com/prometheus/common/pull/779)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.63.0...v0.64.0>

### [`v0.63.0`](https://redirect.github.com/prometheus/common/releases/tag/v0.63.0)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.62.0...v0.63.0)

#### What's Changed

- Making the map a public variable for promtheus-operator by [@&#8203;dongjiang1989](https://redirect.github.com/dongjiang1989) in [#&#8203;741](https://redirect.github.com/prometheus/common/pull/741)
- setup ossf scorecard and codeql workflows by [@&#8203;mmorel-35](https://redirect.github.com/mmorel-35) in [#&#8203;564](https://redirect.github.com/prometheus/common/pull/564)
- feat(promslog): implement reserved keys, rename duplicates by [@&#8203;tjhop](https://redirect.github.com/tjhop) in [#&#8203;746](https://redirect.github.com/prometheus/common/pull/746)
- Bump golang.org/x/oauth2 from 0.24.0 to 0.25.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;750](https://redirect.github.com/prometheus/common/pull/750)
- Bump golang.org/x/net from 0.33.0 to 0.34.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;749](https://redirect.github.com/prometheus/common/pull/749)
- Bump google.golang.org/protobuf from 1.36.1 to 1.36.3 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;751](https://redirect.github.com/prometheus/common/pull/751)
- promslog: Make AllowedLevel concurrency safe. by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [#&#8203;754](https://redirect.github.com/prometheus/common/pull/754)
- Fix typo 'the an' by [@&#8203;petern48](https://redirect.github.com/petern48) in [#&#8203;752](https://redirect.github.com/prometheus/common/pull/752)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;757](https://redirect.github.com/prometheus/common/pull/757)
- build(deps): bump google.golang.org/protobuf from 1.36.3 to 1.36.4 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;756](https://redirect.github.com/prometheus/common/pull/756)
- build(deps): bump google.golang.org/protobuf from 1.36.4 to 1.36.5 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;761](https://redirect.github.com/prometheus/common/pull/761)
- build(deps): bump github.com/google/go-cmp from 0.6.0 to 0.7.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;763](https://redirect.github.com/prometheus/common/pull/763)
- build(deps): bump golang.org/x/net from 0.34.0 to 0.35.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;762](https://redirect.github.com/prometheus/common/pull/762)
- model: Clarify the purpose of model.NameValidationScheme by [@&#8203;ywwg](https://redirect.github.com/ywwg) in [#&#8203;765](https://redirect.github.com/prometheus/common/pull/765)
- Fix spelling mistake in godoc by [@&#8203;grobinson-grafana](https://redirect.github.com/grobinson-grafana) in [#&#8203;766](https://redirect.github.com/prometheus/common/pull/766)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;767](https://redirect.github.com/prometheus/common/pull/767)
- otlptranslator: Add dependency free package that translates OTLP data into Prometheus metric/label names by [@&#8203;ArthurSens](https://redirect.github.com/ArthurSens) in [#&#8203;768](https://redirect.github.com/prometheus/common/pull/768)

#### New Contributors

- [@&#8203;dongjiang1989](https://redirect.github.com/dongjiang1989) made their first contribution in [#&#8203;741](https://redirect.github.com/prometheus/common/pull/741)
- [@&#8203;petern48](https://redirect.github.com/petern48) made their first contribution in [#&#8203;752](https://redirect.github.com/prometheus/common/pull/752)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.62.0...v0.63.0>

### [`v0.62.0`](https://redirect.github.com/prometheus/common/releases/tag/v0.62.0)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.61.0...v0.62.0)

⚠️ This releases switches internal global to `UTF8Validation` from `LegacyValidation`. This is a breaking change, relaxing the validation. We don't intend to add more schemas and we have to have a global for unmarshalling interfaces, thus the change was made ⚠️

#### What's Changed

- Change default validation scheme to UTF8Validation by [@&#8203;ywwg](https://redirect.github.com/ywwg) in [#&#8203;724](https://redirect.github.com/prometheus/common/pull/724)
- Remove deprecated promlog package by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;738](https://redirect.github.com/prometheus/common/pull/738)
- Remove deprecated sigv4 module by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;737](https://redirect.github.com/prometheus/common/pull/737)
- update links to openmetrics to reference the v1.0.0 release by [@&#8203;dashpole](https://redirect.github.com/dashpole) in [#&#8203;740](https://redirect.github.com/prometheus/common/pull/740)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;742](https://redirect.github.com/prometheus/common/pull/742)
- Bump google.golang.org/protobuf from 1.35.2 to 1.36.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;744](https://redirect.github.com/prometheus/common/pull/744)
- Bump golang.org/x/net from 0.32.0 to 0.33.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;743](https://redirect.github.com/prometheus/common/pull/743)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;747](https://redirect.github.com/prometheus/common/pull/747)
- http\_config: Allow customizing TLS config and settings. by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [#&#8203;748](https://redirect.github.com/prometheus/common/pull/748)

#### New Contributors

- [@&#8203;dashpole](https://redirect.github.com/dashpole) made their first contribution in [#&#8203;740](https://redirect.github.com/prometheus/common/pull/740)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.61.0...v0.62.0>

### [`v0.61.0`](https://redirect.github.com/prometheus/common/releases/tag/v0.61.0)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.60.1...v0.61.0)

#### What's Changed

- Mark sigv4 deprecated by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;715](https://redirect.github.com/prometheus/common/pull/715)
- Provide a way to get UserAgent by [@&#8203;mmorel-35](https://redirect.github.com/mmorel-35) in [#&#8203;716](https://redirect.github.com/prometheus/common/pull/716)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;714](https://redirect.github.com/prometheus/common/pull/714)
- Bump golang.org/x/net from 0.29.0 to 0.30.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;712](https://redirect.github.com/prometheus/common/pull/712)
- chore: enable perfsprint linter by [@&#8203;mmorel-35](https://redirect.github.com/mmorel-35) in [#&#8203;717](https://redirect.github.com/prometheus/common/pull/717)
- chore: use testify instead of testing.Fatal by [@&#8203;mmorel-35](https://redirect.github.com/mmorel-35) in [#&#8203;718](https://redirect.github.com/prometheus/common/pull/718)
- Bump google.golang.org/protobuf from 1.34.2 to 1.35.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;711](https://redirect.github.com/prometheus/common/pull/711)
- setup dependabot for `github.com/prometheus/common/assets` by [@&#8203;mmorel-35](https://redirect.github.com/mmorel-35) in [#&#8203;719](https://redirect.github.com/prometheus/common/pull/719)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;721](https://redirect.github.com/prometheus/common/pull/721)
- Mark promlog deprecated by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;720](https://redirect.github.com/prometheus/common/pull/720)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;722](https://redirect.github.com/prometheus/common/pull/722)
- Allow custom user-agent definition by [@&#8203;mmorel-35](https://redirect.github.com/mmorel-35) in [#&#8203;725](https://redirect.github.com/prometheus/common/pull/725)
- fix: values escaping bugs by [@&#8203;ywwg](https://redirect.github.com/ywwg) in [#&#8203;727](https://redirect.github.com/prometheus/common/pull/727)
- fix(promslog): always use UTC for time by [@&#8203;tjhop](https://redirect.github.com/tjhop) in [#&#8203;735](https://redirect.github.com/prometheus/common/pull/735)
- Bump github.com/stretchr/testify from 1.9.0 to 1.10.0 in /assets by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;729](https://redirect.github.com/prometheus/common/pull/729)
- Bump golang.org/x/oauth2 from 0.23.0 to 0.24.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;730](https://redirect.github.com/prometheus/common/pull/730)
- promslog: always lowercase log level from CLI by [@&#8203;jkroepke](https://redirect.github.com/jkroepke) in [#&#8203;728](https://redirect.github.com/prometheus/common/pull/728)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;726](https://redirect.github.com/prometheus/common/pull/726)
- Bump golang.org/x/net from 0.30.0 to 0.32.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;736](https://redirect.github.com/prometheus/common/pull/736)
- Bump github.com/stretchr/testify from 1.9.0 to 1.10.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;731](https://redirect.github.com/prometheus/common/pull/731)
- Bump google.golang.org/protobuf from 1.35.1 to 1.35.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;732](https://redirect.github.com/prometheus/common/pull/732)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.60.1...v0.61.0>

### [`v0.60.1`](https://redirect.github.com/prometheus/common/releases/tag/v0.60.1)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.60.0...v0.60.1)

#### What's Changed

- promslog: Only log basename, not full path by [@&#8203;roidelapluie](https://redirect.github.com/roidelapluie) in [#&#8203;705](https://redirect.github.com/prometheus/common/pull/705)
- Reload certificates even when no CA is used by [@&#8203;roidelapluie](https://redirect.github.com/roidelapluie) in [#&#8203;707](https://redirect.github.com/prometheus/common/pull/707)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;701](https://redirect.github.com/prometheus/common/pull/701)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.60.0...v0.60.1>

### [`v0.60.0`](https://redirect.github.com/prometheus/common/releases/tag/v0.60.0)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.59.1...v0.60.0)

#### What's Changed

- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;692](https://redirect.github.com/prometheus/common/pull/692)
- slog: expose io.Writer by [@&#8203;jkroepke](https://redirect.github.com/jkroepke) in [#&#8203;694](https://redirect.github.com/prometheus/common/pull/694)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;695](https://redirect.github.com/prometheus/common/pull/695)
- promslog: use UTC timestamps for go-kit log style by [@&#8203;dswarbrick](https://redirect.github.com/dswarbrick) in [#&#8203;696](https://redirect.github.com/prometheus/common/pull/696)
- feat: add `promslog.NewNopLogger()` convenience func by [@&#8203;tjhop](https://redirect.github.com/tjhop) in [#&#8203;697](https://redirect.github.com/prometheus/common/pull/697)
- Bump golang.org/x/net from 0.28.0 to 0.29.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;699](https://redirect.github.com/prometheus/common/pull/699)
- Bump golang.org/x/oauth2 from 0.22.0 to 0.23.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;698](https://redirect.github.com/prometheus/common/pull/698)
- Update supported Go versions by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;700](https://redirect.github.com/prometheus/common/pull/700)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.59.1...v0.60.0>

### [`v0.59.1`](https://redirect.github.com/prometheus/common/releases/tag/v0.59.1)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.59.0...v0.59.1)

#### What's Changed

- fix(utf8): Fix multiple metric name inside braces validation by [@&#8203;fedetorres93](https://redirect.github.com/fedetorres93) in [#&#8203;691](https://redirect.github.com/prometheus/common/pull/691)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.59.0...v0.59.1>

### [`v0.59.0`](https://redirect.github.com/prometheus/common/releases/tag/v0.59.0)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.58.0...v0.59.0)

#### What's Changed

- expfmt: Add WithEscapingScheme to help construct Formats by [@&#8203;ywwg](https://redirect.github.com/ywwg) in [#&#8203;688](https://redirect.github.com/prometheus/common/pull/688)
- Change the default escape method to UnderscoreEscaping by [@&#8203;ywwg](https://redirect.github.com/ywwg) in [#&#8203;690](https://redirect.github.com/prometheus/common/pull/690)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.58.0...v0.59.0>

### [`v0.58.0`](https://redirect.github.com/prometheus/common/releases/tag/v0.58.0)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.57.0...v0.58.0)

#### What's Changed

- docs: mention new promslog package in package list in README by [@&#8203;tjhop](https://redirect.github.com/tjhop) in [#&#8203;683](https://redirect.github.com/prometheus/common/pull/683)
- Bump golang.org/x/oauth2 from 0.21.0 to 0.22.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;684](https://redirect.github.com/prometheus/common/pull/684)
- Bump golang.org/x/net from 0.27.0 to 0.28.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;685](https://redirect.github.com/prometheus/common/pull/685)
- Remove secret file existence check in Validate for headers by [@&#8203;roidelapluie](https://redirect.github.com/roidelapluie) in [#&#8203;687](https://redirect.github.com/prometheus/common/pull/687)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.57.0...v0.58.0>

### [`v0.57.0`](https://redirect.github.com/prometheus/common/releases/tag/v0.57.0)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.56.0...v0.57.0)

#### What's Changed

- feat: new promslog and promslog/flag packages to wrap log/slog by [@&#8203;tjhop](https://redirect.github.com/tjhop) in [#&#8203;677](https://redirect.github.com/prometheus/common/pull/677)

#### New Contributors

- [@&#8203;tjhop](https://redirect.github.com/tjhop) made their first contribution in [#&#8203;677](https://redirect.github.com/prometheus/common/pull/677)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.56.0...v0.57.0>

### [`v0.56.0`](https://redirect.github.com/prometheus/common/releases/tag/v0.56.0)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.55.0...v0.56.0)

#### What's Changed

- Don't always fetch a OAuth2 token, if the secret from a file didn't change by [@&#8203;multani](https://redirect.github.com/multani) in [#&#8203;647](https://redirect.github.com/prometheus/common/pull/647)
- remove dependency to github.com/prometheus/client\_golang by [@&#8203;ilius](https://redirect.github.com/ilius) in [#&#8203;662](https://redirect.github.com/prometheus/common/pull/662)
- Bump github.com/aws/aws-sdk-go from 1.54.7 to 1.54.11 in /sigv4 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;661](https://redirect.github.com/prometheus/common/pull/661)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;664](https://redirect.github.com/prometheus/common/pull/664)
- Revert [#&#8203;576](https://redirect.github.com/prometheus/common/issues/576) and add deprecation notice by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;665](https://redirect.github.com/prometheus/common/pull/665)
- Bump golang.org/x/net from 0.26.0 to 0.27.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;667](https://redirect.github.com/prometheus/common/pull/667)
- use basic string in IsValidLegacyMetricName by [@&#8203;ywwg](https://redirect.github.com/ywwg) in [#&#8203;668](https://redirect.github.com/prometheus/common/pull/668)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;672](https://redirect.github.com/prometheus/common/pull/672)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;674](https://redirect.github.com/prometheus/common/pull/674)
- Bump github.com/aws/aws-sdk-go from 1.54.19 to 1.55.5 in /sigv4 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;671](https://redirect.github.com/prometheus/common/pull/671)
- sigv4: support nil body by [@&#8203;roidelapluie](https://redirect.github.com/roidelapluie) in [#&#8203;673](https://redirect.github.com/prometheus/common/pull/673)
- Fix overflows of untyped int constants on 32-bit by [@&#8203;dswarbrick](https://redirect.github.com/dswarbrick) in [#&#8203;675](https://redirect.github.com/prometheus/common/pull/675)
- Update client\_golang by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;676](https://redirect.github.com/prometheus/common/pull/676)
- Update golangci lint by [@&#8203;roidelapluie](https://redirect.github.com/roidelapluie) in [#&#8203;679](https://redirect.github.com/prometheus/common/pull/679)
- expfmt: Add UTF-8 syntax support in text\_parse.go by [@&#8203;fedetorres93](https://redirect.github.com/fedetorres93) in [#&#8203;670](https://redirect.github.com/prometheus/common/pull/670)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;681](https://redirect.github.com/prometheus/common/pull/681)
- fix(utf8): provide a method for explicitly checking label names for legacy validity by [@&#8203;ywwg](https://redirect.github.com/ywwg) in [#&#8203;682](https://redirect.github.com/prometheus/common/pull/682)

#### New Contributors

- [@&#8203;multani](https://redirect.github.com/multani) made their first contribution in [#&#8203;647](https://redirect.github.com/prometheus/common/pull/647)
- [@&#8203;ilius](https://redirect.github.com/ilius) made their first contribution in [#&#8203;662](https://redirect.github.com/prometheus/common/pull/662)
- [@&#8203;dswarbrick](https://redirect.github.com/dswarbrick) made their first contribution in [#&#8203;675](https://redirect.github.com/prometheus/common/pull/675)
- [@&#8203;fedetorres93](https://redirect.github.com/fedetorres93) made their first contribution in [#&#8203;670](https://redirect.github.com/prometheus/common/pull/670)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.55.0...v0.56.0>

### [`v0.55.0`](https://redirect.github.com/prometheus/common/releases/tag/v0.55.0)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.54.0...v0.55.0)

#### What's Changed

- Move goautoneg to external dependency by [@&#8203;mikelolasagasti](https://redirect.github.com/mikelolasagasti) in [#&#8203;625](https://redirect.github.com/prometheus/common/pull/625)
- Expose secret as SecretReader and InlineSecret from config package by [@&#8203;pracucci](https://redirect.github.com/pracucci) in [#&#8203;650](https://redirect.github.com/prometheus/common/pull/650)
- Fix HTTPClientConfig JSON marshalling by [@&#8203;pracucci](https://redirect.github.com/pracucci) in [#&#8203;651](https://redirect.github.com/prometheus/common/pull/651)
- Expose secret as FileSecret from config package by [@&#8203;alanprot](https://redirect.github.com/alanprot) in [#&#8203;653](https://redirect.github.com/prometheus/common/pull/653)
- Set http\_headers to be omit empty by [@&#8203;yeya24](https://redirect.github.com/yeya24) in [#&#8203;655](https://redirect.github.com/prometheus/common/pull/655)
- chore: add HumanizeTimestamp; make ConvertToFloat exportable by [@&#8203;freak12techno](https://redirect.github.com/freak12techno) in [#&#8203;654](https://redirect.github.com/prometheus/common/pull/654)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;660](https://redirect.github.com/prometheus/common/pull/660)
- Add SigV4 FIPS STS  endpoint config by [@&#8203;rajagopalanand](https://redirect.github.com/rajagopalanand) in [#&#8203;649](https://redirect.github.com/prometheus/common/pull/649)

#### New Contributors

- [@&#8203;gotjosh](https://redirect.github.com/gotjosh) made their first contribution in [#&#8203;644](https://redirect.github.com/prometheus/common/pull/644)
- [@&#8203;mikelolasagasti](https://redirect.github.com/mikelolasagasti) made their first contribution in [#&#8203;625](https://redirect.github.com/prometheus/common/pull/625)
- [@&#8203;alanprot](https://redirect.github.com/alanprot) made their first contribution in [#&#8203;653](https://redirect.github.com/prometheus/common/pull/653)
- [@&#8203;yeya24](https://redirect.github.com/yeya24) made their first contribution in [#&#8203;655](https://redirect.github.com/prometheus/common/pull/655)
- [@&#8203;rajagopalanand](https://redirect.github.com/rajagopalanand) made their first contribution in [#&#8203;649](https://redirect.github.com/prometheus/common/pull/649)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.54.0...v0.55.0>

### [`v0.54.0`](https://redirect.github.com/prometheus/common/releases/tag/v0.54.0)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.53.0...v0.54.0)

#### What's Changed

- Bump golang.org/x/net from 0.22.0 to 0.23.0 in /sigv4 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;624](https://redirect.github.com/prometheus/common/pull/624)
- Bump golang.org/x/net from 0.22.0 to 0.23.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;623](https://redirect.github.com/prometheus/common/pull/623)
- Add HTTP headers support to common HTTP client. by [@&#8203;roidelapluie](https://redirect.github.com/roidelapluie) in [#&#8203;416](https://redirect.github.com/prometheus/common/pull/416)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;633](https://redirect.github.com/prometheus/common/pull/633)
- Bump github.com/aws/aws-sdk-go from 1.51.11 to 1.51.32 in /sigv4 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;632](https://redirect.github.com/prometheus/common/pull/632)
- Bump golang.org/x/oauth2 from 0.18.0 to 0.19.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;628](https://redirect.github.com/prometheus/common/pull/628)
- Bump golang.org/x/net from 0.23.0 to 0.24.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;630](https://redirect.github.com/prometheus/common/pull/630)
- Bump github.com/prometheus/client\_model from 0.6.0 to 0.6.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;631](https://redirect.github.com/prometheus/common/pull/631)
- Bump google.golang.org/protobuf from 1.33.0 to 1.34.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;629](https://redirect.github.com/prometheus/common/pull/629)
- Use common interface to fetch secrets in HTTP client config by [@&#8203;TheSpiritXIII](https://redirect.github.com/TheSpiritXIII) in [#&#8203;538](https://redirect.github.com/prometheus/common/pull/538)
- Add support for secret refs via a secret manager by [@&#8203;TheSpiritXIII](https://redirect.github.com/TheSpiritXIII) in [#&#8203;572](https://redirect.github.com/prometheus/common/pull/572)
- oauth2RoundTripper: Avoid race condition and readability changes. by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [#&#8203;634](https://redirect.github.com/prometheus/common/pull/634)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;636](https://redirect.github.com/prometheus/common/pull/636)
- Bump github.com/aws/aws-sdk-go from 1.51.32 to 1.53.14 in /sigv4 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;638](https://redirect.github.com/prometheus/common/pull/638)
- Bump github.com/prometheus/client\_golang from 1.19.0 to 1.19.1 in /sigv4 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;639](https://redirect.github.com/prometheus/common/pull/639)
- feat: add time template helpers by [@&#8203;freak12techno](https://redirect.github.com/freak12techno) in [#&#8203;627](https://redirect.github.com/prometheus/common/pull/627)

#### New Contributors

- [@&#8203;bwplotka](https://redirect.github.com/bwplotka) made their first contribution in [#&#8203;634](https://redirect.github.com/prometheus/common/pull/634)
- [@&#8203;freak12techno](https://redirect.github.com/freak12techno) made their first contribution in [#&#8203;627](https://redirect.github.com/prometheus/common/pull/627)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.53.0...v0.54.0>

### [`v0.53.0`](https://redirect.github.com/prometheus/common/releases/tag/v0.53.0)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.52.3...v0.53.0)

#### What's Changed

- Add StatusAt method for Alert struct by [@&#8203;grobinson-grafana](https://redirect.github.com/grobinson-grafana) in [#&#8203;618](https://redirect.github.com/prometheus/common/pull/618)
- config: allow exposing real secret value through marshal by [@&#8203;GiedriusS](https://redirect.github.com/GiedriusS) in [#&#8203;487](https://redirect.github.com/prometheus/common/pull/487)
- Fix up config test by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;621](https://redirect.github.com/prometheus/common/pull/621)
- LabelSet.String: restore faster sort call by [@&#8203;bboreham](https://redirect.github.com/bboreham) in [#&#8203;619](https://redirect.github.com/prometheus/common/pull/619)
- LabelSet: add unit test for String method by [@&#8203;bboreham](https://redirect.github.com/bboreham) in [#&#8203;620](https://redirect.github.com/prometheus/common/pull/620)

#### New Contributors

- [@&#8203;grobinson-grafana](https://redirect.github.com/grobinson-grafana) made their first contribution in [#&#8203;618](https://redirect.github.com/prometheus/common/pull/618)
- [@&#8203;GiedriusS](https://redirect.github.com/GiedriusS) made their first contribution in [#&#8203;487](https://redirect.github.com/prometheus/common/pull/487)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.52.3...v0.53.0>

### [`v0.52.3`](https://redirect.github.com/prometheus/common/releases/tag/v0.52.3)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.52.2...v0.52.3)

#### What's Changed

- Support go 1.20 by [@&#8203;ArthurSens](https://redirect.github.com/ArthurSens) in [#&#8203;617](https://redirect.github.com/prometheus/common/pull/617)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.52.2...v0.52.3>

### [`v0.52.2`](https://redirect.github.com/prometheus/common/releases/tag/v0.52.2)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.51.1...v0.52.2)

#### What's Changed

- Drop support for Go older than 1.18 by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;612](https://redirect.github.com/prometheus/common/pull/612)
- fix(protobuf): Correctly decode multi-messages streams by [@&#8203;srebhan](https://redirect.github.com/srebhan) in [#&#8203;616](https://redirect.github.com/prometheus/common/pull/616)
- Bump github.com/aws/aws-sdk-go from 1.50.31 to 1.51.11 in /sigv4 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;615](https://redirect.github.com/prometheus/common/pull/615)

#### New Contributors

- [@&#8203;srebhan](https://redirect.github.com/srebhan) made their first contribution in [#&#8203;616](https://redirect.github.com/prometheus/common/pull/616)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.51.1...v0.52.2>

### [`v0.51.1`](https://redirect.github.com/prometheus/common/releases/tag/v0.51.1)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.51.0...v0.51.1)

#### What's Changed

- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;606](https://redirect.github.com/prometheus/common/pull/606)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;609](https://redirect.github.com/prometheus/common/pull/609)
- Retract v0.50.0 by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;610](https://redirect.github.com/prometheus/common/pull/610)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.51.0...v0.51.1>

### [`v0.51.0`](https://redirect.github.com/prometheus/common/releases/tag/v0.51.0)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.50.0...v0.51.0)

#### What's Changed

- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;604](https://redirect.github.com/prometheus/common/pull/604)
- expfmt: Add a way to generate different OpenMetrics Formats by [@&#8203;ywwg](https://redirect.github.com/ywwg) in [#&#8203;596](https://redirect.github.com/prometheus/common/pull/596)
- Fix string slice definition for FormatFlagOptions. by [@&#8203;gizmoguy](https://redirect.github.com/gizmoguy) in [#&#8203;607](https://redirect.github.com/prometheus/common/pull/607)
- Correct logic in sample naming for counters, add new test by [@&#8203;vesari](https://redirect.github.com/vesari) in [#&#8203;608](https://redirect.github.com/prometheus/common/pull/608)

#### New Contributors

- [@&#8203;gizmoguy](https://redirect.github.com/gizmoguy) made their first contribution in [#&#8203;607](https://redirect.github.com/prometheus/common/pull/607)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.50.0...v0.51.0>

### [`v0.50.0`](https://redirect.github.com/prometheus/common/releases/tag/v0.50.0)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.49.0...v0.50.0)

#### What's Changed

- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;594](https://redirect.github.com/prometheus/common/pull/594)
- Bump github.com/stretchr/testify from 1.8.4 to 1.9.0 in /sigv4 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;593](https://redirect.github.com/prometheus/common/pull/593)
- Bump github.com/aws/aws-sdk-go from 1.50.27 to 1.50.29 in /sigv4 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;592](https://redirect.github.com/prometheus/common/pull/592)
- Bump github.com/aws/aws-sdk-go from 1.50.29 to 1.50.31 in /sigv4 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;595](https://redirect.github.com/prometheus/common/pull/595)
- Remove unused 'Host' member from HTTPClientConfig by [@&#8203;bboreham](https://redirect.github.com/bboreham) in [#&#8203;597](https://redirect.github.com/prometheus/common/pull/597)
- Add OpenMetrics unit support by [@&#8203;vesari](https://redirect.github.com/vesari) in [#&#8203;544](https://redirect.github.com/prometheus/common/pull/544)
- Remove deprecated version function by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;591](https://redirect.github.com/prometheus/common/pull/591)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;599](https://redirect.github.com/prometheus/common/pull/599)
- Bump golang.org/x/oauth2 from 0.17.0 to 0.18.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;600](https://redirect.github.com/prometheus/common/pull/600)
- Bump google.golang.org/protobuf from 1.32.0 to 1.33.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;601](https://redirect.github.com/prometheus/common/pull/601)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.49.0...v0.50.0>

### [`v0.49.0`](https://redirect.github.com/prometheus/common/releases/tag/v0.49.0)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.48.0...v0.49.0)

#### What's Changed

- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;574](https://redirect.github.com/prometheus/common/pull/574)
- Bump github.com/aws/aws-sdk-go from 1.49.13 to 1.50.8 in /sigv4 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;571](https://redirect.github.com/prometheus/common/pull/571)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;581](https://redirect.github.com/prometheus/common/pull/581)
- Update Go by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;588](https://redirect.github.com/prometheus/common/pull/588)
- Deprecate version.NewCollector by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;579](https://redirect.github.com/prometheus/common/pull/579)
- Bump github.com/aws/aws-sdk-go from 1.50.8 to 1.50.27 in /sigv4 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;587](https://redirect.github.com/prometheus/common/pull/587)
- Avoid off-spec openmetrics exposition when exemplars have empty labels by [@&#8203;orls](https://redirect.github.com/orls) in [#&#8203;569](https://redirect.github.com/prometheus/common/pull/569)
- Bump golang.org/x/oauth2 from 0.16.0 to 0.17.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;585](https://redirect.github.com/prometheus/common/pull/585)
- Write created lines when negotiating OpenMetrics by [@&#8203;ArthurSens](https://redirect.github.com/ArthurSens) in [#&#8203;504](https://redirect.github.com/prometheus/common/pull/504)
- Upgrade client\_model to v.0.6.0 by [@&#8203;vesari](https://redirect.github.com/vesari) in [#&#8203;589](https://redirect.github.com/prometheus/common/pull/589)
- http\_config: Add host by [@&#8203;jkroepke](https://redirect.github.com/jkroepke) in [#&#8203;549](https://redirect.github.com/prometheus/common/pull/549)
- LabelSet: Fix alphabetical sorting for prometheus LabelSet by [@&#8203;wasim-nihal](https://redirect.github.com/wasim-nihal) in [#&#8203;575](https://redirect.github.com/prometheus/common/pull/575)
- labelset: optimise String() function by [@&#8203;bboreham](https://redirect.github.com/bboreham) in [#&#8203;590](https://redirect.github.com/prometheus/common/pull/590)

#### New Contributors

- [@&#8203;orls](https://redirect.github.com/orls) made their first contribution in [#&#8203;569](https://redirect.github.com/prometheus/common/pull/569)
- [@&#8203;vesari](https://redirect.github.com/vesari) made their first contribution in [#&#8203;589](https://redirect.github.com/prometheus/common/pull/589)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.48.0...v0.49.0>

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

## Discussion

### Comment by @red-hat-konflux on November 07, 2025 at 12:18 PM UTC

### ℹ Artifact update notice

##### File name: rh_identity_transform/go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 2 additional dependencies were updated


Details:


| **Package**                           | **Change**             |
| :------------------------------------ | :--------------------- |
| `github.com/prometheus/client_golang` | `v1.19.1` -> `v1.20.4` |
| `github.com/prometheus/procfs`        | `v0.12.0` -> `v0.15.1` |

---

## Reviews

### Review by @Hyperkid123 - Approved on November 11, 2025 at 01:11 PM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/50*
