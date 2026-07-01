---
type: pull_request
number: 1615
title: "chore(deps): update slackapi/slack-github-action action to v3"
state: merged
author: red-hat-konflux
created: 2026-05-08T21:33:37Z
updated: 2026-05-10T21:33:12Z
closed: 2026-05-10T18:13:50Z
merged: 2026-05-10T18:13:50Z
base_branch: master
head_branch: konflux/mintmaker/master/slackapi-slack-github-action-3.x
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1615
---

# Pull Request #1615: chore(deps): update slackapi/slack-github-action action to v3

**Author**: @red-hat-konflux
**Created**: May 08, 2026 at 09:33 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/slackapi-slack-github-action-3.x`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [slackapi/slack-github-action](https://redirect.github.com/slackapi/slack-github-action) | action | major | `v2.1.1` → `v3.0.3` |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>slackapi/slack-github-action (slackapi/slack-github-action)</summary>

### [`v3.0.3`](https://redirect.github.com/slackapi/slack-github-action/releases/tag/v3.0.3): Slack GitHub Action v3.0.3

[Compare Source](https://redirect.github.com/slackapi/slack-github-action/compare/v3.0.2...v3.0.3)

##### Patch Changes

- [`66834e4`](https://redirect.github.com/slackapi/slack-github-action/commit/66834e4): feat: add instrumentation to address error rates

### [`v3.0.2`](https://redirect.github.com/slackapi/slack-github-action/releases/tag/v3.0.2): Slack GitHub Action v3.0.2

[Compare Source](https://redirect.github.com/slackapi/slack-github-action/compare/v3.0.1...v3.0.2)

##### Patch Changes

- [`79529d7`](https://redirect.github.com/slackapi/slack-github-action/commit/79529d7): fix: resolve url.parse deprecation warning for webhook techniques

### [`v3.0.1`](https://redirect.github.com/slackapi/slack-github-action/releases/tag/v3.0.1): Slack GitHub Action v3.0.1

[Compare Source](https://redirect.github.com/slackapi/slack-github-action/compare/v3.0.0...v3.0.1)

#### What's Changed

Alongside the breaking changes of [`@v3.0.0`](https://redirect.github.com/slackapi/slack-github-action/releases/tag/v3.0.0) and a [new technique](https://docs.slack.dev/tools/slack-github-action/sending-techniques/running-slack-cli-commands/) to run Slack CLI commands, we tried the wrong name to publish to the GitHub Marketplace 🐙  This action is now noted as [**The Slack GitHub Action**](https://redirect.github.com/marketplace/actions/the-slack-github-action) in listings 🎶 ✨

##### :art: Maintenance

- chore: use a unique title for marketplace in [#&#8203;576](https://redirect.github.com/slackapi/slack-github-action/pull/576) - Thanks [@&#8203;zimeg](https://redirect.github.com/zimeg)!
- chore(release): tag version 3.0.1 in [#&#8203;577](https://redirect.github.com/slackapi/slack-github-action/pull/577) - Thanks [@&#8203;zimeg](https://redirect.github.com/zimeg)!

**Full Changelog**: <https://github.com/slackapi/slack-github-action/compare/v3.0.0...v3.0.1>

### [`v3.0.0`](https://redirect.github.com/slackapi/slack-github-action/releases/tag/v3.0.0): Slack GitHub Action v3.0.0

[Compare Source](https://redirect.github.com/slackapi/slack-github-action/compare/v2.1.1...v3.0.0)

> The `@v3.0.0` release had a hiccup on publish and we recommend using [**@&#8203;v3.0.1**](https://redirect.github.com/slackapi/slack-github-action/releases/tag/v3.0.1) or a more recent version when updating! Oops!

🎽 **Running Slack CLI commands and the active Node runtime, both included in this release** 👟 ✨

##### ⚠️ Breaking change: Node.js 24 the runtime

This major version updates the GitHub Actions required runtime to [**Node.js 24**.](https://nodejs.org/en/about/previous-releases) Most [GitHub-hosted runners](https://redirect.github.com/actions/runner-images?tab=readme-ov-file#software-and-image-support) already include this, but self-hosted runners may need to be updated ahead of [planned deprecations of Node 20 on GitHub Actions runners](https://github.blog/changelog/2025-09-19-deprecation-of-node-20-on-github-actions-runners/).

##### 📺 Enhancement: Run Slack CLI commands

This release introduces a new technique for running [Slack CLI](https://docs.slack.dev/tools/slack-cli) commands directly in GitHub Actions workflows. Use this to install the latest version (or a specific one) of the CLI and execute commands like `deploy` for merges to main, `manifest validate` with tests, and other [commands](https://docs.slack.dev/tools/slack-cli/reference/commands/slack).

Gather a token using the following CLI command to store with repo secrets, then get started with an example below:

```
$ slack auth token
```

##### 🧪 Validate an app manifest on pull requests

Check that your app manifest is valid before merging changes:

🔗 <https://docs.slack.dev/tools/slack-github-action/sending-techniques/running-slack-cli-commands/validate-a-manifest>

```yaml
- name: Validate the manifest
  uses: slackapi/slack-github-action/cli@v3.0.0
  with:
    command: "manifest validate --app ${{ vars.SLACK_APP_ID }}"
    token: ${{ secrets.SLACK_SERVICE_TOKEN }}
```

##### 🚀 Deploy your app on push to main

Automate deployments whenever changes land on your main branch:

🔗 <https://docs.slack.dev/tools/slack-github-action/sending-techniques/running-slack-cli-commands/deploy-an-app>

```yaml
- name: Deploy the app
  uses: slackapi/slack-github-action/cli@v3.0.0
  with:
    command: "deploy --app ${{ vars.SLACK_APP_ID }} --force"
    token: ${{ secrets.SLACK_SERVICE_TOKEN }}
```

Any Slack CLI command can be passed through the `command` option without the "slack" prefix 🍀

The `token` input accepts a [service token](https://docs.slack.dev/authentication/tokens/#service) for authentication. You can gather this token by running [`slack auth token`](https://docs.slack.dev/tools/slack-cli/reference/commands/slack_auth_token) with the Slack CLI and storing the value as a repository secret.

The latest Slack CLI version is used by default, but a specific one can be set with the `version` input.

***

🏆 Huge thanks to [@&#8203;ewanek1](https://redirect.github.com/ewanek1) for explorations and prototypes toward the scripted CLI technique!

For full documentation on the CLI technique, check out the [docs](https://docs.slack.dev/tools/slack-github-action/sending-techniques/running-slack-cli-commands/) and explore the related pages 📚

#### What's Changed

##### :space\_invader: Enhancements

- feat: support slack cli commands with composite action inputs in [#&#8203;560](https://redirect.github.com/slackapi/slack-github-action/pull/560) - Thanks [@&#8203;zimeg](https://redirect.github.com/zimeg)!
- build!: update node runtime to version 24 in [#&#8203;567](https://redirect.github.com/slackapi/slack-github-action/pull/567) - Thanks [@&#8203;desrosj](https://redirect.github.com/desrosj)!

##### :books: Documentation

- docs: updates links to point to `docs.slack.dev/tools*` paths in [#&#8203;485](https://redirect.github.com/slackapi/slack-github-action/pull/485) - Thanks [@&#8203;lukegalbraithrussell](https://redirect.github.com/lukegalbraithrussell)!
- docs: fix typos and misspellings in descriptions in [#&#8203;530](https://redirect.github.com/slackapi/slack-github-action/pull/530) - Thanks [@&#8203;szepeviktor](https://redirect.github.com/szepeviktor)!
- docs: move pull request requirements instructions into a comment in [#&#8203;551](https://redirect.github.com/slackapi/slack-github-action/pull/551) - Thanks [@&#8203;zimeg](https://redirect.github.com/zimeg)!
- docs: redirect links to the current developer documentation pages in [#&#8203;532](https://redirect.github.com/slackapi/slack-github-action/pull/532) - Thanks [@&#8203;zimeg](https://redirect.github.com/zimeg)!

##### :toolbox: Maintenance

- chore: update steps taken to release a new latest tag and version in [#&#8203;439](https://redirect.github.com/slackapi/slack-github-action/pull/439) - Thanks [@&#8203;zimeg](https://redirect.github.com/zimeg)!
- chore: Add .github/CODEOWNERS file in [#&#8203;509](https://redirect.github.com/slackapi/slack-github-action/pull/509) - Thanks [@&#8203;mwbrooks](https://redirect.github.com/mwbrooks)!
- build: match the node types package and node version in [#&#8203;531](https://redirect.github.com/slackapi/slack-github-action/pull/531) - Thanks [@&#8203;zimeg](https://redirect.github.com/zimeg)!
- test: use node test runner to assert expected cases in [#&#8203;538](https://redirect.github.com/slackapi/slack-github-action/pull/538) - Thanks [@&#8203;zimeg](https://redirect.github.com/zimeg)!
- test: switch junit test reporter for lcov results to upload in [#&#8203;539](https://redirect.github.com/slackapi/slack-github-action/pull/539) - Thanks [@&#8203;zimeg](https://redirect.github.com/zimeg)!
- build: update [@&#8203;actions](https://redirect.github.com/actions) dependencies to versions with esm support in [#&#8203;547](https://redirect.github.com/slackapi/slack-github-action/pull/547) - Thanks [@&#8203;zimeg](https://redirect.github.com/zimeg)!
- ci(deps): auto-approve / auto-merge dependencies from dependabot in [#&#8203;548](https://redirect.github.com/slackapi/slack-github-action/pull/548) - Thanks [@&#8203;mwbrooks](https://redirect.github.com/mwbrooks)!
- build: ignore dist when linting and formating in [#&#8203;550](https://redirect.github.com/slackapi/slack-github-action/pull/550) - Thanks [@&#8203;zimeg](https://redirect.github.com/zimeg)!
- ci: run tests without permission checks for origin pull requests in [#&#8203;553](https://redirect.github.com/slackapi/slack-github-action/pull/553) - Thanks [@&#8203;zimeg](https://redirect.github.com/zimeg)!
- chore(release): tag version 3.0.0 in [#&#8203;575](https://redirect.github.com/slackapi/slack-github-action/pull/575) - Thanks [@&#8203;zimeg](https://redirect.github.com/zimeg)!

##### :gift: Dependencies

- build(deps): bump [@&#8203;actions/core](https://redirect.github.com/actions/core) from 1.11.1 to 2.0.1 in [#&#8203;526](https://redirect.github.com/slackapi/slack-github-action/pull/526) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump [@&#8203;actions/github](https://redirect.github.com/actions/github) from 6.0.1 to 7.0.0 in [#&#8203;537](https://redirect.github.com/slackapi/slack-github-action/pull/537) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump [@&#8203;slack/logger](https://redirect.github.com/slack/logger) from 4.0.0 to 4.0.1 in [#&#8203;573](https://redirect.github.com/slackapi/slack-github-action/pull/573) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump [@&#8203;slack/web-api](https://redirect.github.com/slack/web-api) from 7.9.3 to 7.10.0 in [#&#8203;497](https://redirect.github.com/slackapi/slack-github-action/pull/497) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump [@&#8203;slack/web-api](https://redirect.github.com/slack/web-api) from 7.10.0 to 7.12.0 in [#&#8203;506](https://redirect.github.com/slackapi/slack-github-action/pull/506) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump [@&#8203;slack/web-api](https://redirect.github.com/slack/web-api) from 7.12.0 to 7.13.0 in [#&#8203;514](https://redirect.github.com/slackapi/slack-github-action/pull/514) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump [@&#8203;slack/web-api](https://redirect.github.com/slack/web-api) from 7.13.0 to 7.14.1 in [#&#8203;564](https://redirect.github.com/slackapi/slack-github-action/pull/564) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump [@&#8203;slack/web-api](https://redirect.github.com/slack/web-api) from 7.14.1 to 7.15.0 in [#&#8203;574](https://redirect.github.com/slackapi/slack-github-action/pull/574) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump actions/checkout from 4.2.2 to 5.0.0 in [#&#8203;495](https://redirect.github.com/slackapi/slack-github-action/pull/495) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump actions/checkout from 5.0.0 to 6.0.0 in [#&#8203;519](https://redirect.github.com/slackapi/slack-github-action/pull/519) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump actions/checkout from 6.0.0 to 6.0.1 in [#&#8203;522](https://redirect.github.com/slackapi/slack-github-action/pull/522) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump actions/checkout from 6.0.1 to 6.0.2 in [#&#8203;540](https://redirect.github.com/slackapi/slack-github-action/pull/540) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump actions/setup-node from 4.4.0 to 5.0.0 in [#&#8203;494](https://redirect.github.com/slackapi/slack-github-action/pull/494) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump actions/setup-node from 5.0.0 to 6.0.0 in [#&#8203;503](https://redirect.github.com/slackapi/slack-github-action/pull/503) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump actions/setup-node from 6.0.0 to 6.1.0 in [#&#8203;523](https://redirect.github.com/slackapi/slack-github-action/pull/523) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump actions/setup-node from 6.1.0 to 6.2.0 in [#&#8203;541](https://redirect.github.com/slackapi/slack-github-action/pull/541) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump actions/setup-node from 6.2.0 to 6.3.0 in [#&#8203;569](https://redirect.github.com/slackapi/slack-github-action/pull/569) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump axios from 1.10.0 to 1.11.0 in [#&#8203;478](https://redirect.github.com/slackapi/slack-github-action/pull/478) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump axios from 1.11.0 to 1.12.2 in [#&#8203;493](https://redirect.github.com/slackapi/slack-github-action/pull/493) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump axios from 1.12.2 to 1.13.2 in [#&#8203;515](https://redirect.github.com/slackapi/slack-github-action/pull/515) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump axios from 1.13.2 to 1.13.4 in [#&#8203;543](https://redirect.github.com/slackapi/slack-github-action/pull/543) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump axios from 1.13.4 to 1.13.5 in [#&#8203;558](https://redirect.github.com/slackapi/slack-github-action/pull/558) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump axios from 1.13.5 to 1.13.6 in [#&#8203;565](https://redirect.github.com/slackapi/slack-github-action/pull/565) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump codecov/codecov-action from 5.4.3 to 5.5.1 in [#&#8203;496](https://redirect.github.com/slackapi/slack-github-action/pull/496) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump codecov/codecov-action from 5.5.1 to 5.5.2 in [#&#8203;525](https://redirect.github.com/slackapi/slack-github-action/pull/525) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump form-data from 4.0.0 to 4.0.4 in [#&#8203;477](https://redirect.github.com/slackapi/slack-github-action/pull/477) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump glob from 10.4.5 to 10.5.0 in [#&#8203;512](https://redirect.github.com/slackapi/slack-github-action/pull/512) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump https-proxy-agent from 7.0.6 to 8.0.0 in [#&#8203;572](https://redirect.github.com/slackapi/slack-github-action/pull/572) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps): bump js-yaml from 4.1.0 to 4.1.1 in [#&#8203;510](https://redirect.github.com/slackapi/slack-github-action/pull/510) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump [@&#8203;biomejs/biome](https://redirect.github.com/biomejs/biome) from 2.0.6 to 2.1.3 in [#&#8203;482](https://redirect.github.com/slackapi/slack-github-action/pull/482) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump [@&#8203;biomejs/biome](https://redirect.github.com/biomejs/biome) from 2.1.3 to 2.2.4 in [#&#8203;499](https://redirect.github.com/slackapi/slack-github-action/pull/499) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump [@&#8203;biomejs/biome](https://redirect.github.com/biomejs/biome) from 2.2.4 to 2.3.2 in [#&#8203;507](https://redirect.github.com/slackapi/slack-github-action/pull/507) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump [@&#8203;biomejs/biome](https://redirect.github.com/biomejs/biome) from 2.3.10 to 2.3.11 in [#&#8203;534](https://redirect.github.com/slackapi/slack-github-action/pull/534) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump [@&#8203;biomejs/biome](https://redirect.github.com/biomejs/biome) from 2.3.11 to 2.3.13 in [#&#8203;545](https://redirect.github.com/slackapi/slack-github-action/pull/545) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump [@&#8203;biomejs/biome](https://redirect.github.com/biomejs/biome) from 2.3.13 to 2.4.4 in [#&#8203;563](https://redirect.github.com/slackapi/slack-github-action/pull/563) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump [@&#8203;biomejs/biome](https://redirect.github.com/biomejs/biome) from 2.3.3 to 2.3.8 in [#&#8203;518](https://redirect.github.com/slackapi/slack-github-action/pull/518) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump [@&#8203;biomejs/biome](https://redirect.github.com/biomejs/biome) from 2.3.8 to 2.3.10 in [#&#8203;527](https://redirect.github.com/slackapi/slack-github-action/pull/527) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump [@&#8203;biomejs/biome](https://redirect.github.com/biomejs/biome) from 2.4.4 to 2.4.6 in [#&#8203;570](https://redirect.github.com/slackapi/slack-github-action/pull/570) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump [@&#8203;types/node](https://redirect.github.com/types/node) from 24.0.8 to 24.1.0 in [#&#8203;483](https://redirect.github.com/slackapi/slack-github-action/pull/483) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump [@&#8203;types/node](https://redirect.github.com/types/node) from 24.1.0 to 24.6.1 in [#&#8203;498](https://redirect.github.com/slackapi/slack-github-action/pull/498) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump [@&#8203;types/node](https://redirect.github.com/types/node) from 24.6.1 to 24.9.2 in [#&#8203;504](https://redirect.github.com/slackapi/slack-github-action/pull/504) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump [@&#8203;types/node](https://redirect.github.com/types/node) from 24.10.0 to 24.10.1 in [#&#8203;517](https://redirect.github.com/slackapi/slack-github-action/pull/517) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump [@&#8203;types/node](https://redirect.github.com/types/node) from 20.19.27 to 20.19.28 in [#&#8203;535](https://redirect.github.com/slackapi/slack-github-action/pull/535) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump [@&#8203;types/node](https://redirect.github.com/types/node) from 20.19.28 to 20.19.30 in [#&#8203;546](https://redirect.github.com/slackapi/slack-github-action/pull/546) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump [@&#8203;types/node](https://redirect.github.com/types/node) from 20.19.30 to 20.19.35 in [#&#8203;562](https://redirect.github.com/slackapi/slack-github-action/pull/562) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump [@&#8203;types/sinon](https://redirect.github.com/types/sinon) from 17.0.4 to 21.0.0 in [#&#8203;516](https://redirect.github.com/slackapi/slack-github-action/pull/516) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump [@&#8203;vercel/ncc](https://redirect.github.com/vercel/ncc) from 0.38.3 to 0.38.4 in [#&#8203;528](https://redirect.github.com/slackapi/slack-github-action/pull/528) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump chai from 5.1.2 to 6.2.0 in [#&#8203;508](https://redirect.github.com/slackapi/slack-github-action/pull/508) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump mocha from 11.7.1 to 11.7.3 in [#&#8203;500](https://redirect.github.com/slackapi/slack-github-action/pull/500) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump mocha from 11.7.3 to 11.7.4 in [#&#8203;505](https://redirect.github.com/slackapi/slack-github-action/pull/505) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump sinon from 21.0.0 to 21.0.1 in [#&#8203;524](https://redirect.github.com/slackapi/slack-github-action/pull/524) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump sinon from 21.0.1 to 21.0.2 in [#&#8203;571](https://redirect.github.com/slackapi/slack-github-action/pull/571) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump typescript from 5.8.3 to 5.9.2 in [#&#8203;481](https://redirect.github.com/slackapi/slack-github-action/pull/481) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!
- build(deps-dev): bump typescript from 5.9.2 to 5.9.3 in [#&#8203;501](https://redirect.github.com/slackapi/slack-github-action/pull/501) - Thanks [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot]!

#### :wave: New Contributors

- [@&#8203;szepeviktor](https://redirect.github.com/szepeviktor) made their first contribution in [#&#8203;530](https://redirect.github.com/slackapi/slack-github-action/pull/530)!

**Full Changelog**: <https://github.com/slackapi/slack-github-action/compare/v2.1.1...v3.0.0>

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0Mi45OS4wLXJwbSIsInVwZGF0ZWRJblZlciI6IjQyLjk5LjAtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @codecov-commenter on May 08, 2026 at 09:36 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1615?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 76.34%. Comparing base ([`0e669c9`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/0e669c99596688190b2b913341652c1c75a73571?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`a979218`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/a979218049c4f30b7f6ccf1f6c0f88bc5d6b18b4?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1615   +/-   ##
=======================================
  Coverage   76.34%   76.34%           
=======================================
  Files         103      103           
  Lines        3187     3187           
  Branches      698      698           
=======================================
  Hits         2433     2433           
  Misses        675      675           
  Partials       79       79           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1615?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @github-actions - Approved on May 08, 2026 at 09:33 PM UTC

This PR from Konflux has been automatically approved.

### Review by @swadeley - Approved on May 10, 2026 at 06:07 PM UTC

ACK

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1615*
