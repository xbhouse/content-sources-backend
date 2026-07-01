---
type: pull_request
number: 54
title: "Update module github.com/slackhq/nebula to v1.9.7"
state: merged
author: red-hat-konflux
created: 2025-11-07T20:17:50Z
updated: 2025-11-11T14:17:11Z
closed: 2025-11-11T14:17:08Z
merged: 2025-11-11T14:17:08Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-slackhq-nebula-1.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/54
---

# Pull Request #54: Update module github.com/slackhq/nebula to v1.9.7

**Author**: @red-hat-konflux
**Created**: November 07, 2025 at 08:17 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-slackhq-nebula-1.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/slackhq/nebula](https://redirect.github.com/slackhq/nebula) | `v1.6.1` -> `v1.9.7` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fslackhq%2fnebula/v1.9.7?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fslackhq%2fnebula/v1.6.1/v1.9.7?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>slackhq/nebula (github.com/slackhq/nebula)</summary>

### [`v1.9.7`](https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.7)

[Compare Source](https://redirect.github.com/slackhq/nebula/compare/v1.9.6...v1.9.7)

##### Security

- Fix an issue where Nebula could incorrectly accept and process a packet from an erroneous source IP when the sender's
  certificate is configured with unsafe\_routes (cert v1/v2) or multiple IPs (cert v2). ([#&#8203;1494](https://redirect.github.com/slackhq/nebula/issues/1494))

##### Changed

- Disable sending `recv_error` messages when a packet is received outside the allowable counter window.  ([#&#8203;1459](https://redirect.github.com/slackhq/nebula/issues/1459))
- Improve error messages and remove some unnecessary fatal conditions in the Windows and generic udp listener.  ([#&#8203;1453](https://redirect.github.com/slackhq/nebula/issues/1453))

### [`v1.9.6`](https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.6)

[Compare Source](https://redirect.github.com/slackhq/nebula/compare/v1.9.5...v1.9.6)

##### Added

- Support dropping inactive tunnels. This is disabled by default in this release but can be enabled with `tunnels.drop_inactive`. See example config for more details. ([#&#8203;1413](https://redirect.github.com/slackhq/nebula/issues/1413))

##### Fixed

- Fix Darwin freeze due to presence of some Network Extensions ([#&#8203;1426](https://redirect.github.com/slackhq/nebula/issues/1426))
- Ensure the same relay tunnel is always used when multiple relay tunnels are present ([#&#8203;1422](https://redirect.github.com/slackhq/nebula/issues/1422))
- Fix Windows freeze due to ICMP error handling ([#&#8203;1412](https://redirect.github.com/slackhq/nebula/issues/1412))
- Fix relay migration panic ([#&#8203;1403](https://redirect.github.com/slackhq/nebula/issues/1403))

### [`v1.9.5`](https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.5)

[Compare Source](https://redirect.github.com/slackhq/nebula/compare/v1.9.4...v1.9.5)

##### Added

- Gracefully ignore v2 certificates. ([#&#8203;1282](https://redirect.github.com/slackhq/nebula/issues/1282))

##### Fixed

- Fix relays that refuse to re-establish after one of the remote tunnel pairs breaks. ([#&#8203;1277](https://redirect.github.com/slackhq/nebula/issues/1277))

### [`v1.9.4`](https://redirect.github.com/slackhq/nebula/blob/HEAD/CHANGELOG.md#100---2019-11-19)

[Compare Source](https://redirect.github.com/slackhq/nebula/compare/v1.9.3...v1.9.4)

##### Added

- Initial public release.

[Unreleased]: https://redirect.github.com/slackhq/nebula/compare/v1.9.4...HEAD

[1.9.4]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.4

[1.9.3]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.3

[1.9.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.2

[1.9.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.1

[1.9.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.0

[1.8.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.2

[1.8.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.1

[1.8.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.0

[1.7.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.2

[1.7.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.1

[1.7.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.0

[1.6.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.6.1

[1.6.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.6.0

[1.5.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.5.2

[1.5.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.5.0

[1.4.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.4.0

[1.3.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.3.0

[1.2.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.2.0

[1.1.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.1.0

[1.0.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.0.0

### [`v1.9.3`](https://redirect.github.com/slackhq/nebula/blob/HEAD/CHANGELOG.md#100---2019-11-19)

[Compare Source](https://redirect.github.com/slackhq/nebula/compare/v1.9.2...v1.9.3)

##### Added

- Initial public release.

[Unreleased]: https://redirect.github.com/slackhq/nebula/compare/v1.9.4...HEAD

[1.9.4]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.4

[1.9.3]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.3

[1.9.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.2

[1.9.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.1

[1.9.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.0

[1.8.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.2

[1.8.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.1

[1.8.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.0

[1.7.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.2

[1.7.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.1

[1.7.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.0

[1.6.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.6.1

[1.6.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.6.0

[1.5.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.5.2

[1.5.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.5.0

[1.4.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.4.0

[1.3.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.3.0

[1.2.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.2.0

[1.1.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.1.0

[1.0.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.0.0

### [`v1.9.2`](https://redirect.github.com/slackhq/nebula/blob/HEAD/CHANGELOG.md#100---2019-11-19)

[Compare Source](https://redirect.github.com/slackhq/nebula/compare/v1.9.1...v1.9.2)

##### Added

- Initial public release.

[Unreleased]: https://redirect.github.com/slackhq/nebula/compare/v1.9.4...HEAD

[1.9.4]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.4

[1.9.3]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.3

[1.9.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.2

[1.9.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.1

[1.9.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.0

[1.8.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.2

[1.8.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.1

[1.8.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.0

[1.7.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.2

[1.7.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.1

[1.7.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.0

[1.6.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.6.1

[1.6.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.6.0

[1.5.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.5.2

[1.5.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.5.0

[1.4.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.4.0

[1.3.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.3.0

[1.2.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.2.0

[1.1.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.1.0

[1.0.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.0.0

### [`v1.9.1`](https://redirect.github.com/slackhq/nebula/blob/HEAD/CHANGELOG.md#100---2019-11-19)

[Compare Source](https://redirect.github.com/slackhq/nebula/compare/v1.9.0...v1.9.1)

##### Added

- Initial public release.

[Unreleased]: https://redirect.github.com/slackhq/nebula/compare/v1.9.4...HEAD

[1.9.4]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.4

[1.9.3]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.3

[1.9.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.2

[1.9.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.1

[1.9.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.0

[1.8.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.2

[1.8.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.1

[1.8.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.0

[1.7.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.2

[1.7.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.1

[1.7.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.0

[1.6.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.6.1

[1.6.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.6.0

[1.5.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.5.2

[1.5.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.5.0

[1.4.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.4.0

[1.3.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.3.0

[1.2.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.2.0

[1.1.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.1.0

[1.0.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.0.0

### [`v1.9.0`](https://redirect.github.com/slackhq/nebula/blob/HEAD/CHANGELOG.md#100---2019-11-19)

[Compare Source](https://redirect.github.com/slackhq/nebula/compare/v1.8.2...v1.9.0)

##### Added

- Initial public release.

[Unreleased]: https://redirect.github.com/slackhq/nebula/compare/v1.9.4...HEAD

[1.9.4]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.4

[1.9.3]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.3

[1.9.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.2

[1.9.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.1

[1.9.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.0

[1.8.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.2

[1.8.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.1

[1.8.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.0

[1.7.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.2

[1.7.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.1

[1.7.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.0

[1.6.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.6.1

[1.6.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.6.0

[1.5.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.5.2

[1.5.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.5.0

[1.4.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.4.0

[1.3.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.3.0

[1.2.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.2.0

[1.1.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.1.0

[1.0.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.0.0

### [`v1.8.2`](https://redirect.github.com/slackhq/nebula/blob/HEAD/CHANGELOG.md#100---2019-11-19)

[Compare Source](https://redirect.github.com/slackhq/nebula/compare/v1.8.1...v1.8.2)

##### Added

- Initial public release.

[Unreleased]: https://redirect.github.com/slackhq/nebula/compare/v1.9.4...HEAD

[1.9.4]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.4

[1.9.3]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.3

[1.9.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.2

[1.9.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.1

[1.9.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.0

[1.8.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.2

[1.8.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.1

[1.8.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.0

[1.7.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.2

[1.7.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.1

[1.7.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.0

[1.6.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.6.1

[1.6.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.6.0

[1.5.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.5.2

[1.5.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.5.0

[1.4.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.4.0

[1.3.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.3.0

[1.2.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.2.0

[1.1.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.1.0

[1.0.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.0.0

### [`v1.8.1`](https://redirect.github.com/slackhq/nebula/blob/HEAD/CHANGELOG.md#100---2019-11-19)

[Compare Source](https://redirect.github.com/slackhq/nebula/compare/v1.8.0...v1.8.1)

##### Added

- Initial public release.

[Unreleased]: https://redirect.github.com/slackhq/nebula/compare/v1.9.4...HEAD

[1.9.4]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.4

[1.9.3]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.3

[1.9.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.2

[1.9.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.1

[1.9.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.0

[1.8.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.2

[1.8.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.1

[1.8.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.0

[1.7.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.2

[1.7.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.1

[1.7.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.0

[1.6.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.6.1

[1.6.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.6.0

[1.5.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.5.2

[1.5.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.5.0

[1.4.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.4.0

[1.3.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.3.0

[1.2.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.2.0

[1.1.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.1.0

[1.0.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.0.0

### [`v1.8.0`](https://redirect.github.com/slackhq/nebula/blob/HEAD/CHANGELOG.md#100---2019-11-19)

[Compare Source](https://redirect.github.com/slackhq/nebula/compare/v1.7.2...v1.8.0)

##### Added

- Initial public release.

[Unreleased]: https://redirect.github.com/slackhq/nebula/compare/v1.9.4...HEAD

[1.9.4]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.4

[1.9.3]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.3

[1.9.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.2

[1.9.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.1

[1.9.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.0

[1.8.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.2

[1.8.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.1

[1.8.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.0

[1.7.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.2

[1.7.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.1

[1.7.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.0

[1.6.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.6.1

[1.6.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.6.0

[1.5.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.5.2

[1.5.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.5.0

[1.4.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.4.0

[1.3.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.3.0

[1.2.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.2.0

[1.1.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.1.0

[1.0.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.0.0

### [`v1.7.2`](https://redirect.github.com/slackhq/nebula/blob/HEAD/CHANGELOG.md#100---2019-11-19)

[Compare Source](https://redirect.github.com/slackhq/nebula/compare/v1.7.1...v1.7.2)

##### Added

- Initial public release.

[Unreleased]: https://redirect.github.com/slackhq/nebula/compare/v1.9.4...HEAD

[1.9.4]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.4

[1.9.3]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.3

[1.9.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.2

[1.9.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.1

[1.9.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.0

[1.8.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.2

[1.8.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.1

[1.8.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.0

[1.7.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.2

[1.7.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.1

[1.7.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.0

[1.6.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.6.1

[1.6.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.6.0

[1.5.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.5.2

[1.5.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.5.0

[1.4.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.4.0

[1.3.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.3.0

[1.2.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.2.0

[1.1.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.1.0

[1.0.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.0.0

### [`v1.7.1`](https://redirect.github.com/slackhq/nebula/blob/HEAD/CHANGELOG.md#100---2019-11-19)

[Compare Source](https://redirect.github.com/slackhq/nebula/compare/v1.7.0...v1.7.1)

##### Added

- Initial public release.

[Unreleased]: https://redirect.github.com/slackhq/nebula/compare/v1.9.4...HEAD

[1.9.4]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.4

[1.9.3]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.3

[1.9.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.2

[1.9.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.1

[1.9.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.0

[1.8.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.2

[1.8.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.1

[1.8.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.0

[1.7.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.2

[1.7.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.1

[1.7.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.0

[1.6.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.6.1

[1.6.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.6.0

[1.5.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.5.2

[1.5.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.5.0

[1.4.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.4.0

[1.3.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.3.0

[1.2.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.2.0

[1.1.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.1.0

[1.0.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.0.0

### [`v1.7.0`](https://redirect.github.com/slackhq/nebula/blob/HEAD/CHANGELOG.md#100---2019-11-19)

[Compare Source](https://redirect.github.com/slackhq/nebula/compare/v1.6.1...v1.7.0)

##### Added

- Initial public release.

[Unreleased]: https://redirect.github.com/slackhq/nebula/compare/v1.9.4...HEAD

[1.9.4]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.4

[1.9.3]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.3

[1.9.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.2

[1.9.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.1

[1.9.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.9.0

[1.8.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.2

[1.8.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.1

[1.8.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.8.0

[1.7.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.2

[1.7.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.1

[1.7.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.7.0

[1.6.1]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.6.1

[1.6.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.6.0

[1.5.2]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.5.2

[1.5.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.5.0

[1.4.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.4.0

[1.3.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.3.0

[1.2.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.2.0

[1.1.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.1.0

[1.0.0]: https://redirect.github.com/slackhq/nebula/releases/tag/v1.0.0

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

### Review by @Hyperkid123 - Approved on November 11, 2025 at 02:17 PM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/54*
