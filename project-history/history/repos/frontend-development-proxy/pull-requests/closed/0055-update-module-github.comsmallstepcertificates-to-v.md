---
type: pull_request
number: 55
title: "Update module github.com/smallstep/certificates to v0.28.4 - autoclosed"
state: closed
author: red-hat-konflux
created: 2025-11-07T20:18:26Z
updated: 2025-11-24T12:51:13Z
closed: 2025-11-24T12:51:11Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-smallstep-certificates-0.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/55
---

# Pull Request #55: Update module github.com/smallstep/certificates to v0.28.4 - autoclosed

**Author**: @red-hat-konflux
**Created**: November 07, 2025 at 08:18 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-smallstep-certificates-0.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/smallstep/certificates](https://redirect.github.com/smallstep/certificates) | `v0.26.1` -> `v0.28.4` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fsmallstep%2fcertificates/v0.28.4?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fsmallstep%2fcertificates/v0.26.1/v0.28.4?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>smallstep/certificates (github.com/smallstep/certificates)</summary>

### [`v0.28.4`](https://redirect.github.com/smallstep/certificates/releases/tag/v0.28.4): Step CA v0.28.4 (25-07-14)

[Compare Source](https://redirect.github.com/smallstep/certificates/compare/v0.28.3...v0.28.4)

#### Official Release Artifacts

##### Linux

- 📦 [step-ca\_linux\_0.28.4\_amd64.tar.gz](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.4/step-ca_linux_0.28.4_amd64.tar.gz)
- 📦 [step-ca\_0.28.4-1\_amd64.deb](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.4/step-ca_0.28.4-1_amd64.deb)
- 📦 [step-ca-0.28.4-1.x86\_64.rpm](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.4/step-ca-0.28.4-1.x86_64.rpm)
- 📦 [step-ca\_0.28.4-1\_arm64.deb](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.4/step-ca_0.28.4-1_arm64.deb)
- 📦 [step-ca-0.28.4-1.aarch64.rpm](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.4/step-ca-0.28.4-1.aarch64.rpm)

##### OSX Darwin

- 📦 [step-ca\_darwin\_0.28.4\_amd64.tar.gz](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.4/step-ca_darwin_0.28.4_amd64.tar.gz)
- 📦 [step-ca\_darwin\_0.28.4\_arm64.tar.gz](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.4/step-ca_darwin_0.28.4_arm64.tar.gz)

##### Windows

- 📦 [step-ca\_windows\_0.28.4\_amd64.zip](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.4/step-ca_windows_0.28.4_amd64.zip)

For more builds across platforms and architectures, see the `Assets` section below.
And for packaged versions (Docker, k8s, Homebrew), see our [installation docs](https://smallstep.com/docs/step-ca/installation).

Don't see the artifact you need? Open an issue [here](https://redirect.github.com/smallstep/certificates/issues/new/choose).

#### Signatures and Checksums

`step-ca` uses [sigstore/cosign](https://redirect.github.com/sigstore/cosign) for signing and verifying release artifacts.

Below is an example using `cosign` to verify a release artifact:

```
cosign verify-blob \
  --certificate step-ca_darwin_0.28.4_amd64.tar.gz.pem \
  --signature step-ca_darwin_0.28.4_amd64.tar.gz.sig \
  --certificate-identity-regexp "https://github\.com/smallstep/workflows/.*" \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com \
  step-ca_darwin_0.28.4_amd64.tar.gz
```

The `checksums.txt` file (in the `Assets` section below) contains a checksum for every artifact in the release.

#### Changelog

- [`2c61c44`](https://redirect.github.com/smallstep/certificates/commit/2c61c44176a89885ea69dd341dca16fb2875d868) Update changelog ([#&#8203;2332](https://redirect.github.com/smallstep/certificates/issues/2332))
- [`c86cf07`](https://redirect.github.com/smallstep/certificates/commit/c86cf07be9c5909a08e631ca0490662f734c3505) Merge pull request [#&#8203;2331](https://redirect.github.com/smallstep/certificates/issues/2331) from smallstep/mariano/fix-tests
- [`831d005`](https://redirect.github.com/smallstep/certificates/commit/831d005df8f245ba2cc98028524488f6d0a7442c) Fix gcp unit tests
- [`bc09e46`](https://redirect.github.com/smallstep/certificates/commit/bc09e46c3c8263b1d10cd5afcf50da34a1c97b82) Merge pull request [#&#8203;2133](https://redirect.github.com/smallstep/certificates/issues/2133) from ericnorris/feat-gcp-enable-organization-checking
- [`0d9f051`](https://redirect.github.com/smallstep/certificates/commit/0d9f0513cfd5506398f972dd8c40a5f3973be769) Merge branch 'master' into feat-gcp-enable-organization-checking
- [`197d0d3`](https://redirect.github.com/smallstep/certificates/commit/197d0d3508d5f424865682a4df61a05850331487) Changelog updates ([#&#8203;2330](https://redirect.github.com/smallstep/certificates/issues/2330))
- [`2932225`](https://redirect.github.com/smallstep/certificates/commit/293222505539eed2e3dff9078070f1368dd55d99) Merge pull request [#&#8203;2329](https://redirect.github.com/smallstep/certificates/issues/2329) from smallstep/dependabot/go\_modules/google.golang.org/api-0.240.0
- [`b1dd5a6`](https://redirect.github.com/smallstep/certificates/commit/b1dd5a6ebd2384468940d37b63fdb74dc59a2a82) Bump google.golang.org/api from 0.239.0 to 0.240.0
- [`8b9bd89`](https://redirect.github.com/smallstep/certificates/commit/8b9bd89bccd755a5c4600e7192bd9fcdd85a6bdb) refactor: make projectIDs authoritative
- [`a3db8de`](https://redirect.github.com/smallstep/certificates/commit/a3db8de6614ea02de4a853a243a571434a232f43) feat(gcp): enable organization checking

#### Thanks!

Those were the changes on v0.28.4!

Come join us on [Discord](https://discord.gg/X2RKGwEbV9) to ask questions, chat about PKI, or get a sneak peek at the freshest PKI memes.

### [`v0.28.3`](https://redirect.github.com/smallstep/certificates/releases/tag/v0.28.3): Step CA v0.28.3 (25-03-18)

[Compare Source](https://redirect.github.com/smallstep/certificates/compare/v0.28.2...v0.28.3)

#### Official Release Artifacts

##### Linux

- 📦 [step-ca\_linux\_0.28.3\_amd64.tar.gz](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.3/step-ca_linux_0.28.3_amd64.tar.gz)
- 📦 [step-ca\_0.28.3-1\_amd64.deb](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.3/step-ca_0.28.3-1_amd64.deb)
- 📦 [step-ca-0.28.3-1.x86\_64.rpm](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.3/step-ca-0.28.3-1.x86_64.rpm)
- 📦 [step-ca\_0.28.3-1\_arm64.deb](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.3/step-ca_0.28.3-1_arm64.deb)
- 📦 [step-ca-0.28.3-1.aarch64.rpm](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.3/step-ca-0.28.3-1.aarch64.rpm)

##### OSX Darwin

- 📦 [step-ca\_darwin\_0.28.3\_amd64.tar.gz](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.3/step-ca_darwin_0.28.3_amd64.tar.gz)
- 📦 [step-ca\_darwin\_0.28.3\_arm64.tar.gz](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.3/step-ca_darwin_0.28.3_arm64.tar.gz)

##### Windows

- 📦 [step-ca\_windows\_0.28.3\_amd64.zip](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.3/step-ca_windows_0.28.3_amd64.zip)

For more builds across platforms and architectures, see the `Assets` section below.
And for packaged versions (Docker, k8s, Homebrew), see our [installation docs](https://smallstep.com/docs/step-ca/installation).

Don't see the artifact you need? Open an issue [here](https://redirect.github.com/smallstep/certificates/issues/new/choose).

#### Signatures and Checksums

`step-ca` uses [sigstore/cosign](https://redirect.github.com/sigstore/cosign) for signing and verifying release artifacts.

Below is an example using `cosign` to verify a release artifact:

```
cosign verify-blob \
  --certificate step-ca_darwin_0.28.3_amd64.tar.gz.sig.pem \
  --signature step-ca_darwin_0.28.3_amd64.tar.gz.sig \
  --certificate-identity-regexp "https://github\.com/smallstep/workflows/.*" \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com \
  step-ca_darwin_0.28.3_amd64.tar.gz
```

The `checksums.txt` file (in the `Assets` section below) contains a checksum for every artifact in the release.

#### Changelog

- [`0cf1c56`](https://redirect.github.com/smallstep/certificates/commit/0cf1c5688708ec4a910c007d7f151c617b722268) empty commit
- [`78745ba`](https://redirect.github.com/smallstep/certificates/commit/78745ba9ff05d489f4bb95789f163217818adf26) empty commit ([#&#8203;2216](https://redirect.github.com/smallstep/certificates/issues/2216))

#### Thanks!

Those were the changes on v0.28.3!

Come join us on [Discord](https://discord.gg/X2RKGwEbV9) to ask questions, chat about PKI, or get a sneak peek at the freshest PKI memes.

### [`v0.28.2`](https://redirect.github.com/smallstep/certificates/releases/tag/v0.28.2): Step CA v0.28.2 (25-02-21)

[Compare Source](https://redirect.github.com/smallstep/certificates/compare/v0.28.1...v0.28.2)

#### Official Release Artifacts

##### Linux

- 📦 [step-ca\_linux\_0.28.2\_amd64.tar.gz](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.2/step-ca_linux_0.28.2_amd64.tar.gz)
- 📦 [step-ca\_0.28.2-1\_amd64.deb](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.2/step-ca_0.28.2-1_amd64.deb)
- 📦 [step-ca-0.28.2-1.x86\_64.rpm](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.2/step-ca-0.28.2-1.x86_64.rpm)
- 📦 [step-ca\_0.28.2-1\_arm64.deb](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.2/step-ca_0.28.2-1_arm64.deb)
- 📦 [step-ca-0.28.2-1.aarch64.rpm](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.2/step-ca-0.28.2-1.aarch64.rpm)

##### OSX Darwin

- 📦 [step-ca\_darwin\_0.28.2\_amd64.tar.gz](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.2/step-ca_darwin_0.28.2_amd64.tar.gz)
- 📦 [step-ca\_darwin\_0.28.2\_arm64.tar.gz](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.2/step-ca_darwin_0.28.2_arm64.tar.gz)

##### Windows

- 📦 [step-ca\_windows\_0.28.2\_amd64.zip](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.2/step-ca_windows_0.28.2_amd64.zip)

For more builds across platforms and architectures, see the `Assets` section below.
And for packaged versions (Docker, k8s, Homebrew), see our [installation docs](https://smallstep.com/docs/step-ca/installation).

Don't see the artifact you need? Open an issue [here](https://redirect.github.com/smallstep/certificates/issues/new/choose).

#### Signatures and Checksums

`step-ca` uses [sigstore/cosign](https://redirect.github.com/sigstore/cosign) for signing and verifying release artifacts.

Below is an example using `cosign` to verify a release artifact:

```
cosign verify-blob \
  --certificate step-ca_darwin_0.28.2_amd64.tar.gz.sig.pem \
  --signature step-ca_darwin_0.28.2_amd64.tar.gz.sig \
  --certificate-identity-regexp "https://github\.com/smallstep/workflows/.*" \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com \
  step-ca_darwin_0.28.2_amd64.tar.gz
```

The `checksums.txt` file (in the `Assets` section below) contains a checksum for every artifact in the release.

#### Changelog

- [`d21e756`](https://redirect.github.com/smallstep/certificates/commit/d21e756af0d310af38afe36b41875e9003d4873c) Add changelog for v0.28.2 ([#&#8203;2178](https://redirect.github.com/smallstep/certificates/issues/2178))
- [`d4960a1`](https://redirect.github.com/smallstep/certificates/commit/d4960a194b626a8ee114a6277034e14dfde87c19) Merge pull request [#&#8203;2179](https://redirect.github.com/smallstep/certificates/issues/2179) from smallstep/carl/linked-ca-init
- [`3e4393e`](https://redirect.github.com/smallstep/certificates/commit/3e4393e1275920460b0cb49e9caf47ae88a60656) Update copy for linked CA init
- [`b22e186`](https://redirect.github.com/smallstep/certificates/commit/b22e186ae944440f03386e8c45bd9db229c4dc7d) Merge pull request [#&#8203;2171](https://redirect.github.com/smallstep/certificates/issues/2171) from smallstep/dependabot/go\_modules/github.com/smallstep/pkcs7-0.2.1
- [`8debf5b`](https://redirect.github.com/smallstep/certificates/commit/8debf5bc971749f109f216956808310fa42a55c0) Bump github.com/smallstep/pkcs7 from 0.1.1 to 0.2.1
- [`8610bf6`](https://redirect.github.com/smallstep/certificates/commit/8610bf631e00b65c549c8fa7089ac51a2c61a82d) Merge pull request [#&#8203;2175](https://redirect.github.com/smallstep/certificates/issues/2175) from smallstep/dependabot/go\_modules/google.golang.org/api-0.221.0
- [`f59221b`](https://redirect.github.com/smallstep/certificates/commit/f59221b9eb876f4f8be673c7a40f52a14ce52b80) Bump google.golang.org/api from 0.220.0 to 0.221.0
- [`3d94255`](https://redirect.github.com/smallstep/certificates/commit/3d942551413ab036f3a9d066d3913b63dbc9f88e) Merge pull request [#&#8203;2172](https://redirect.github.com/smallstep/certificates/issues/2172) from smallstep/dependabot/go\_modules/github.com/hashicorp/vault/api/auth/aws-0.9.0
- [`0ce1ecd`](https://redirect.github.com/smallstep/certificates/commit/0ce1ecd1a5f36f7a3852497cba5057ba784a693b) Bump github.com/hashicorp/vault/api/auth/aws from 0.8.0 to 0.9.0
- [`497cb20`](https://redirect.github.com/smallstep/certificates/commit/497cb20a193bbfd7bf2e1c622fdc0d315d47992d) Merge pull request [#&#8203;2174](https://redirect.github.com/smallstep/certificates/issues/2174) from smallstep/dependabot/go\_modules/github.com/hashicorp/vault/api/auth/approle-0.9.0
- [`9e717d7`](https://redirect.github.com/smallstep/certificates/commit/9e717d77488978770d4872fe3f7bc757214e9228) Bump github.com/hashicorp/vault/api/auth/approle from 0.8.0 to 0.9.0
- [`3db57f7`](https://redirect.github.com/smallstep/certificates/commit/3db57f724abfb4b9d4a08ee5ee75c721af4f5b58) Merge pull request [#&#8203;2173](https://redirect.github.com/smallstep/certificates/issues/2173) from smallstep/dependabot/go\_modules/github.com/hashicorp/vault/api/auth/kubernetes-0.9.0
- [`ed2c5ca`](https://redirect.github.com/smallstep/certificates/commit/ed2c5ca81e6d560fb9262592e63fde203d4ea6a2) Bump github.com/hashicorp/vault/api/auth/kubernetes from 0.8.0 to 0.9.0
- [`e38dec9`](https://redirect.github.com/smallstep/certificates/commit/e38dec9d098dc32f2c593b08145467933817a452) Merge pull request [#&#8203;2167](https://redirect.github.com/smallstep/certificates/issues/2167) from smallstep/herman/upgrade-linter
- [`27944b4`](https://redirect.github.com/smallstep/certificates/commit/27944b4eaee2c5cc32852ade074a1affa1e91f8d) Fix linter issues
- [`86c04f0`](https://redirect.github.com/smallstep/certificates/commit/86c04f0ce830c0a9d2d3c23c8bb531a77857cae7) Merge pull request [#&#8203;2166](https://redirect.github.com/smallstep/certificates/issues/2166) from smallstep/dependabot/go\_modules/google.golang.org/api-0.220.0
- [`e236934`](https://redirect.github.com/smallstep/certificates/commit/e236934007849c9e61dfec0a4d1c447e6b67f4d4) Bump to Go v1.22.12
- [`dca1c6b`](https://redirect.github.com/smallstep/certificates/commit/dca1c6b1e09e4d515b0864a205399bec8f0410ab) Bump google.golang.org/api from 0.219.0 to 0.220.0
- [`726e0dc`](https://redirect.github.com/smallstep/certificates/commit/726e0dcf8ad9f1045fc414d114ffef3af0bdf68a) Merge pull request [#&#8203;2164](https://redirect.github.com/smallstep/certificates/issues/2164) from smallstep/dependabot/go\_modules/github.com/go-chi/chi/v5-5.2.1
- [`d48c5a1`](https://redirect.github.com/smallstep/certificates/commit/d48c5a10787fa90e927154e6896596cd3b011e6f) Merge pull request [#&#8203;2165](https://redirect.github.com/smallstep/certificates/issues/2165) from smallstep/dependabot/go\_modules/google.golang.org/protobuf-1.36.5
- [`e0ffc47`](https://redirect.github.com/smallstep/certificates/commit/e0ffc4706b7c2196c5f02f826b045709efd791e7) Merge pull request [#&#8203;2163](https://redirect.github.com/smallstep/certificates/issues/2163) from smallstep/dependabot/go\_modules/golang.org/x/crypto-0.33.0
- [`b3f7349`](https://redirect.github.com/smallstep/certificates/commit/b3f73499ac1b394666ececc2e3d40e9cb69d249d) Bump google.golang.org/protobuf from 1.36.4 to 1.36.5
- [`791e2af`](https://redirect.github.com/smallstep/certificates/commit/791e2af9123ba3e777e373b9e3f508e06e095064) Bump github.com/go-chi/chi/v5 from 5.2.0 to 5.2.1
- [`52a4a1b`](https://redirect.github.com/smallstep/certificates/commit/52a4a1b7d7417fad1e7f1e4140a3514f411c1b68) Bump golang.org/x/crypto from 0.32.0 to 0.33.0
- [`3159a74`](https://redirect.github.com/smallstep/certificates/commit/3159a74c5ecfc8ae8a49e6066c7a4be3a7018711) Merge pull request [#&#8203;2153](https://redirect.github.com/smallstep/certificates/issues/2153) from smallstep/dependabot/go\_modules/google.golang.org/api-0.219.0
- [`1f78066`](https://redirect.github.com/smallstep/certificates/commit/1f78066fce8ef8552d41fa24ca362b8c83dc1e72) Merge pull request [#&#8203;2154](https://redirect.github.com/smallstep/certificates/issues/2154) from smallstep/dependabot/go\_modules/go.step.sm/crypto-0.57.1
- [`14b6f61`](https://redirect.github.com/smallstep/certificates/commit/14b6f618b008a088342948ea581340b2244a85f2) Bump go.step.sm/crypto from 0.57.0 to 0.57.1
- [`9dabb19`](https://redirect.github.com/smallstep/certificates/commit/9dabb19da02989a81ff65de0ae63c27496a3f0f8) Bump google.golang.org/api from 0.218.0 to 0.219.0
- [`cf1ef2a`](https://redirect.github.com/smallstep/certificates/commit/cf1ef2a77dfab39389423bfa453cd637e8cd1e0c) Merge pull request [#&#8203;2151](https://redirect.github.com/smallstep/certificates/issues/2151) from smallstep/dependabot/go\_modules/github.com/golang/glog-1.2.4
- [`f879d10`](https://redirect.github.com/smallstep/certificates/commit/f879d10ff7f5edff772dee64cb56f9daab073ccf) Bump github.com/golang/glog from 1.2.3 to 1.2.4
- [`4ff06cf`](https://redirect.github.com/smallstep/certificates/commit/4ff06cf1fc01d4def9bde7fde6d4c6cecba4fc28) Merge pull request [#&#8203;2149](https://redirect.github.com/smallstep/certificates/issues/2149) from smallstep/dependabot/go\_modules/google.golang.org/api-0.218.0
- [`b4cddb4`](https://redirect.github.com/smallstep/certificates/commit/b4cddb424b040affc2c8d21f119c4e35c33c443a) Bump google.golang.org/api from 0.217.0 to 0.218.0
- [`1df134e`](https://redirect.github.com/smallstep/certificates/commit/1df134ed8492b826262011a886894c31ef8779b8) Merge pull request [#&#8203;2147](https://redirect.github.com/smallstep/certificates/issues/2147) from smallstep/dependabot/go\_modules/google.golang.org/protobuf-1.36.4
- [`c70a072`](https://redirect.github.com/smallstep/certificates/commit/c70a072ddf8c2253a2e711b6198efbe555369d7f) Bump google.golang.org/protobuf from 1.36.3 to 1.36.4
- [`58dca0f`](https://redirect.github.com/smallstep/certificates/commit/58dca0f8221d948dda410a7414fd5d60ae6f1142) Merge pull request [#&#8203;2148](https://redirect.github.com/smallstep/certificates/issues/2148) from smallstep/dependabot/go\_modules/google.golang.org/grpc-1.70.0
- [`b0c0344`](https://redirect.github.com/smallstep/certificates/commit/b0c034479ce977fc85eeefb61b3cd0a9a81de65a) Bump google.golang.org/grpc from 1.69.4 to 1.70.0
- [`93cd8fc`](https://redirect.github.com/smallstep/certificates/commit/93cd8fc98153e9b66e98385728a298516966affc) Merge pull request [#&#8203;2141](https://redirect.github.com/smallstep/certificates/issues/2141) from smallstep/dependabot/go\_modules/go.step.sm/crypto-0.57.0
- [`6b1a82d`](https://redirect.github.com/smallstep/certificates/commit/6b1a82d169a6e366debdf217ddb58a11aa4dd057) Merge pull request [#&#8203;2144](https://redirect.github.com/smallstep/certificates/issues/2144) from smallstep/dependabot/go\_modules/github.com/newrelic/go-agent/v3-3.36.0
- [`775071a`](https://redirect.github.com/smallstep/certificates/commit/775071a54f3a35a91273f6fa569666a61d7696b5) Merge pull request [#&#8203;2142](https://redirect.github.com/smallstep/certificates/issues/2142) from smallstep/dependabot/go\_modules/google.golang.org/api-0.217.0
- [`54ea865`](https://redirect.github.com/smallstep/certificates/commit/54ea865e58c29402424389d4a32ce924100e6764) Bump google.golang.org/api from 0.216.0 to 0.217.0
- [`565cab2`](https://redirect.github.com/smallstep/certificates/commit/565cab261682268cbf0a7d848686b7970757bb4c) Merge pull request [#&#8203;2143](https://redirect.github.com/smallstep/certificates/issues/2143) from smallstep/dependabot/go\_modules/google.golang.org/protobuf-1.36.3
- [`78979c6`](https://redirect.github.com/smallstep/certificates/commit/78979c6221dc5d428e7a546baa1b5407a67af6d3) Bump github.com/newrelic/go-agent/v3 from 3.35.1 to 3.36.0
- [`9d90fb8`](https://redirect.github.com/smallstep/certificates/commit/9d90fb80065b29c4a654817017337727253c6d49) Bump google.golang.org/protobuf from 1.36.2 to 1.36.3
- [`a291fee`](https://redirect.github.com/smallstep/certificates/commit/a291fee64763280a797df7aa30240219057c5c51) Bump go.step.sm/crypto from 0.56.0 to 0.57.0
- [`a729f1f`](https://redirect.github.com/smallstep/certificates/commit/a729f1fc1ea755bbddb6d708cffb46fdaebfe7ba) Merge pull request [#&#8203;2140](https://redirect.github.com/smallstep/certificates/issues/2140) from smallstep/herman/fix-download-urls
- [`d1cfae0`](https://redirect.github.com/smallstep/certificates/commit/d1cfae029242b081e5493f46ba8af0ae3ada362f) Fix download URLs that were pointing to `cli`
- [`41820c8`](https://redirect.github.com/smallstep/certificates/commit/41820c8f76f5fff3cbd08c88df7e84de29e463d2) Merge pull request [#&#8203;2138](https://redirect.github.com/smallstep/certificates/issues/2138) from smallstep/dependabot/go\_modules/google.golang.org/grpc-1.69.4
- [`e76ae73`](https://redirect.github.com/smallstep/certificates/commit/e76ae73c1e72aedb66dc54643221208c2ebbde1b) Bump google.golang.org/grpc from 1.69.2 to 1.69.4
- [`c5a0249`](https://redirect.github.com/smallstep/certificates/commit/c5a024992f328d8ad6e2c3712f645bf08558c013) Merge pull request [#&#8203;2137](https://redirect.github.com/smallstep/certificates/issues/2137) from smallstep/dependabot/go\_modules/google.golang.org/protobuf-1.36.2
- [`97d1692`](https://redirect.github.com/smallstep/certificates/commit/97d1692e1d2312205f90d8b3e3305ceb0192b122) Merge pull request [#&#8203;2136](https://redirect.github.com/smallstep/certificates/issues/2136) from smallstep/dependabot/go\_modules/golang.org/x/net-0.34.0
- [`a7f7a15`](https://redirect.github.com/smallstep/certificates/commit/a7f7a1549f5effb47446108e4dc152be26088db0) Bump golang.org/x/net from 0.33.0 to 0.34.0
- [`7a180e4`](https://redirect.github.com/smallstep/certificates/commit/7a180e4477121b16d30ec2495ad9d98880d2e6b3) Bump google.golang.org/protobuf from 1.36.1 to 1.36.2
- [`b1195f2`](https://redirect.github.com/smallstep/certificates/commit/b1195f216d818fcfd1997341d155fd3046e5aff0) Merge pull request [#&#8203;2135](https://redirect.github.com/smallstep/certificates/issues/2135) from smallstep/dependabot/github\_actions/softprops/action-gh-release-2.2.1
- [`57b0378`](https://redirect.github.com/smallstep/certificates/commit/57b03786ee1c0913bc2a8434d287455347ee056b) Merge pull request [#&#8203;2139](https://redirect.github.com/smallstep/certificates/issues/2139) from smallstep/dependabot/go\_modules/google.golang.org/api-0.216.0
- [`6283fe6`](https://redirect.github.com/smallstep/certificates/commit/6283fe6acc760c867da08aa92e98e85764195294) Bump google.golang.org/api from 0.214.0 to 0.216.0
- [`9f50dd9`](https://redirect.github.com/smallstep/certificates/commit/9f50dd9030897f20afcacbf36f61feff216691c1) Bump softprops/action-gh-release from 2.2.0 to 2.2.1
- [`47c08d5`](https://redirect.github.com/smallstep/certificates/commit/47c08d5cf04671bf38685b78742b1d36b9427fe6) Merge pull request [#&#8203;2131](https://redirect.github.com/smallstep/certificates/issues/2131) from smallstep/herman/unflake-request-id-integration-test
- [`b3fb927`](https://redirect.github.com/smallstep/certificates/commit/b3fb92756fe66cdd6b19267c4718b8eb59cf9850) Add comment to `requireCAServerToBeAvailable`
- [`c2a6b5b`](https://redirect.github.com/smallstep/certificates/commit/c2a6b5b5880356657bde0fe1b029567f04036f76) Wait till a connection to the CA host/port can be made
- [`cda0eec`](https://redirect.github.com/smallstep/certificates/commit/cda0eec3a596fbc7c72b48ee4b576e79ecc7625a) Merge pull request [#&#8203;2129](https://redirect.github.com/smallstep/certificates/issues/2129) from smallstep/dependabot/go\_modules/cloud.google.com/go/security-1.18.3
- [`6796e4f`](https://redirect.github.com/smallstep/certificates/commit/6796e4fa156635200f535eda04d79eaeff3b3914) Bump cloud.google.com/go/security from 1.18.2 to 1.18.3
- [`98ddef5`](https://redirect.github.com/smallstep/certificates/commit/98ddef57295692fb05af0c79550419d0b831cae4) Merge pull request [#&#8203;2127](https://redirect.github.com/smallstep/certificates/issues/2127) from smallstep/dependabot/go\_modules/github.com/coreos/go-oidc/v3-3.12.0
- [`1d965aa`](https://redirect.github.com/smallstep/certificates/commit/1d965aad16834f44ca3c39acc66f257c6116d30b) Merge pull request [#&#8203;2130](https://redirect.github.com/smallstep/certificates/issues/2130) from smallstep/dependabot/go\_modules/golang.org/x/crypto-0.32.0
- [`f7cd3ba`](https://redirect.github.com/smallstep/certificates/commit/f7cd3bad14335ee230a601bcdc5d58fc5e34d0e0) Bump github.com/coreos/go-oidc/v3 from 3.11.0 to 3.12.0
- [`2dd29b1`](https://redirect.github.com/smallstep/certificates/commit/2dd29b19b13b437e72ac859de3218de6be7bc437) Merge pull request [#&#8203;2128](https://redirect.github.com/smallstep/certificates/issues/2128) from smallstep/dependabot/go\_modules/cloud.google.com/go/longrunning-0.6.4
- [`1ecf556`](https://redirect.github.com/smallstep/certificates/commit/1ecf556ff1ff159873084e2b2c051badbafb5d49) Bump golang.org/x/crypto from 0.31.0 to 0.32.0
- [`164f9a9`](https://redirect.github.com/smallstep/certificates/commit/164f9a9b9cfc35b7220f7b5f48031b0c15390141) Bump cloud.google.com/go/longrunning from 0.6.3 to 0.6.4
- [`e57a3ca`](https://redirect.github.com/smallstep/certificates/commit/e57a3ca3807e3abbdb567ce0dcf04be2376ba402) Merge pull request [#&#8203;2126](https://redirect.github.com/smallstep/certificates/issues/2126) from smallstep/herman/upgrade-linkedca-v0.23.0
- [`f473632`](https://redirect.github.com/smallstep/certificates/commit/f4736325fa4737877baa9112a78aa7cb047ec5f0) Use `github.com/smallstep/linkedca` @&#8203; `v0.23.0`
- [`3f9dc06`](https://redirect.github.com/smallstep/certificates/commit/3f9dc06cf5907c2b8b394a9142221841e60f15ff) Merge pull request [#&#8203;2124](https://redirect.github.com/smallstep/certificates/issues/2124) from smallstep/mariano/payloadformat
- [`ba9e082`](https://redirect.github.com/smallstep/certificates/commit/ba9e08243ab2c40083ee763fbb463344a5084d65) Add attestation format to challenge
- [`143e484`](https://redirect.github.com/smallstep/certificates/commit/143e4846e0c14bf44149a95d057ebe12e04d600f) Merge pull request [#&#8203;2123](https://redirect.github.com/smallstep/certificates/issues/2123) from smallstep/dependabot/go\_modules/google.golang.org/protobuf-1.36.1
- [`7847e67`](https://redirect.github.com/smallstep/certificates/commit/7847e67ea3fd4448e100742e7f9fe3b6090c3924) Bump google.golang.org/protobuf from 1.36.0 to 1.36.1
- [`f812cf2`](https://redirect.github.com/smallstep/certificates/commit/f812cf26c4b5c9f633ea579e1888d6cfa6272203) Allow storing the attestation payload ([#&#8203;2114](https://redirect.github.com/smallstep/certificates/issues/2114))
- [`232f464`](https://redirect.github.com/smallstep/certificates/commit/232f46454cf3aaf4c3286898492ba080bcec7aaa) Merge pull request [#&#8203;2117](https://redirect.github.com/smallstep/certificates/issues/2117) from smallstep/dependabot/go\_modules/google.golang.org/grpc-1.69.2
- [`c917e97`](https://redirect.github.com/smallstep/certificates/commit/c917e97931d988e898fa6d714db6ec2bfb48c0be) Bump google.golang.org/grpc from 1.69.0 to 1.69.2
- [`2dce25c`](https://redirect.github.com/smallstep/certificates/commit/2dce25cac75258d0c98ef96ec0653f1c1311a070) Merge pull request [#&#8203;2118](https://redirect.github.com/smallstep/certificates/issues/2118) from smallstep/dependabot/go\_modules/google.golang.org/api-0.214.0
- [`c70a905`](https://redirect.github.com/smallstep/certificates/commit/c70a9052e962b27953debdf61a0c39b8be5e8186) Bump google.golang.org/api from 0.212.0 to 0.214.0
- [`3bf1bba`](https://redirect.github.com/smallstep/certificates/commit/3bf1bba33610c3b3d2834cb88393f669a3e9899a) Merge pull request [#&#8203;2116](https://redirect.github.com/smallstep/certificates/issues/2116) from smallstep/dependabot/go\_modules/github.com/google/go-tpm-0.9.3
- [`7398b43`](https://redirect.github.com/smallstep/certificates/commit/7398b4375c3c5503ea8bbaeb9ce6bb6dffeb69e4) Merge pull request [#&#8203;2120](https://redirect.github.com/smallstep/certificates/issues/2120) from smallstep/dependabot/go\_modules/github.com/googleapis/gax-go/v2-2.14.1
- [`45fe6ce`](https://redirect.github.com/smallstep/certificates/commit/45fe6ce32273c12e6cb88fc4af041782c81a157e) Merge pull request [#&#8203;2119](https://redirect.github.com/smallstep/certificates/issues/2119) from smallstep/dependabot/go\_modules/github.com/smallstep/pkcs7-0.1.1
- [`db0377c`](https://redirect.github.com/smallstep/certificates/commit/db0377c8db10beec5917e7863a4fdefbeebb0cef) Bump github.com/googleapis/gax-go/v2 from 2.14.0 to 2.14.1
- [`19bcc33`](https://redirect.github.com/smallstep/certificates/commit/19bcc3367fc07e015ab9b165b1bea9ae3e62c531) Bump github.com/smallstep/pkcs7
- [`ab422b2`](https://redirect.github.com/smallstep/certificates/commit/ab422b212be377118cd9ca6341b3352b27934d03) Bump github.com/google/go-tpm from 0.9.2 to 0.9.3
- [`cba7add`](https://redirect.github.com/smallstep/certificates/commit/cba7add37e4d1e8a951efb490b3dee23fb158720) Merge pull request [#&#8203;2113](https://redirect.github.com/smallstep/certificates/issues/2113) from smallstep/mariano/fix-2101
- [`ac5b36f`](https://redirect.github.com/smallstep/certificates/commit/ac5b36fee8d3d19de04ec0d7c2f29f84b3229c0e) Upgrade golang.org/x/net
- [`4af059b`](https://redirect.github.com/smallstep/certificates/commit/4af059b6b3c5c592c6ffdf3bb11edc690707353f) Upgrade crypto with support for imported keys on YibiKey KMS
- [`6938c78`](https://redirect.github.com/smallstep/certificates/commit/6938c78effaa8015f156e00ea9529ab8ff273b1d) Merge pull request [#&#8203;2107](https://redirect.github.com/smallstep/certificates/issues/2107) from smallstep/dependabot/go\_modules/github.com/go-chi/chi/v5-5.2.0
- [`97c5e69`](https://redirect.github.com/smallstep/certificates/commit/97c5e690420b723bd391e3f5cd882040788bcc87) Merge pull request [#&#8203;2106](https://redirect.github.com/smallstep/certificates/issues/2106) from smallstep/dependabot/go\_modules/google.golang.org/grpc-1.69.0
- [`2a528c9`](https://redirect.github.com/smallstep/certificates/commit/2a528c9de109d0921becba45baf469e949b4cc85) Bump github.com/go-chi/chi/v5 from 5.1.0 to 5.2.0
- [`197bcd0`](https://redirect.github.com/smallstep/certificates/commit/197bcd0647b44fd721cdfb6901067f26bbf4bdcf) Bump google.golang.org/grpc from 1.68.1 to 1.69.0
- [`ea2016b`](https://redirect.github.com/smallstep/certificates/commit/ea2016bd3f52d2b018c1458bc890581d41e7076f) Merge pull request [#&#8203;2110](https://redirect.github.com/smallstep/certificates/issues/2110) from smallstep/dependabot/go\_modules/google.golang.org/api-0.211.0
- [`a085c22`](https://redirect.github.com/smallstep/certificates/commit/a085c229f6a2811de621734f9c66a7b890c5c16a) Bump google.golang.org/api from 0.210.0 to 0.211.0
- [`7dbc429`](https://redirect.github.com/smallstep/certificates/commit/7dbc429279dcae1bfac7a78403cd32f19f2fb1cc) Merge pull request [#&#8203;2109](https://redirect.github.com/smallstep/certificates/issues/2109) from smallstep/dependabot/go\_modules/golang.org/x/crypto-0.31.0
- [`91faf6a`](https://redirect.github.com/smallstep/certificates/commit/91faf6af82b32cf05e629908e7e25bc786e9e65f) Merge pull request [#&#8203;2108](https://redirect.github.com/smallstep/certificates/issues/2108) from smallstep/dependabot/go\_modules/golang.org/x/net-0.32.0
- [`b071e0c`](https://redirect.github.com/smallstep/certificates/commit/b071e0c8ad912fc4b210c98e0a4f0e485c8f729b) Merge pull request [#&#8203;2111](https://redirect.github.com/smallstep/certificates/issues/2111) from smallstep/dependabot/github\_actions/softprops/action-gh-release-2.2.0
- [`1bc4ede`](https://redirect.github.com/smallstep/certificates/commit/1bc4eded3b2d0e02313cff55ccbaa0a1843b1fbe) Bump softprops/action-gh-release from 2.1.0 to 2.2.0
- [`9192369`](https://redirect.github.com/smallstep/certificates/commit/9192369a12dfc87e90a3eba659fc1ddb72d56359) Bump golang.org/x/crypto from 0.30.0 to 0.31.0
- [`397e17b`](https://redirect.github.com/smallstep/certificates/commit/397e17b3754833891548af7908d864c19596eaf1) Bump golang.org/x/net from 0.31.0 to 0.32.0
- [`351287d`](https://redirect.github.com/smallstep/certificates/commit/351287d23b2b5bac7c514c96942fff323efc4dcd) Merge pull request [#&#8203;2104](https://redirect.github.com/smallstep/certificates/issues/2104) from smallstep/mariano/wraptransport
- [`98087fe`](https://redirect.github.com/smallstep/certificates/commit/98087fec306ad68ffc6efbf7ba26a42815552879) Move wrap transport initialization to constructor
- [`809c702`](https://redirect.github.com/smallstep/certificates/commit/809c7023c95fe99ba8fa89d1a2494c109b7730d7) Transport wrappers ([#&#8203;2103](https://redirect.github.com/smallstep/certificates/issues/2103))
- [`c986962`](https://redirect.github.com/smallstep/certificates/commit/c9869621549de59d075b895e7c73b5c65592623e) internal/httptransport: initial implementation of the package ([#&#8203;2098](https://redirect.github.com/smallstep/certificates/issues/2098))
- [`51e253b`](https://redirect.github.com/smallstep/certificates/commit/51e253be7198f3669d3703ed7b570704202b199b) Merge pull request [#&#8203;2097](https://redirect.github.com/smallstep/certificates/issues/2097) from smallstep/dependabot/go\_modules/github.com/slackhq/nebula-1.9.5
- [`4c7aa8a`](https://redirect.github.com/smallstep/certificates/commit/4c7aa8a6233514ec1cce677c44619f66bda0c9d0) Add test case for Nebula certificate errors
- [`9000271`](https://redirect.github.com/smallstep/certificates/commit/9000271ce0c6c21b0c06c412b7fe8f44a5480a63) Fix new return value from `nebula.NewCAPoolFromBytes`
- [`f902049`](https://redirect.github.com/smallstep/certificates/commit/f90204930e41ed45f248910745eb7b0911017250) Merge pull request [#&#8203;2094](https://redirect.github.com/smallstep/certificates/issues/2094) from smallstep/dependabot/go\_modules/google.golang.org/grpc-1.68.1
- [`7a7714f`](https://redirect.github.com/smallstep/certificates/commit/7a7714fb370c85854f72b8bf9e879fa9d3408b71) Merge pull request [#&#8203;2096](https://redirect.github.com/smallstep/certificates/issues/2096) from smallstep/dependabot/go\_modules/golang.org/x/crypto-0.30.0
- [`30dd9a6`](https://redirect.github.com/smallstep/certificates/commit/30dd9a6c132e82deeae574171d36dec13e101ece) Bump google.golang.org/grpc from 1.67.1 to 1.68.1
- [`7df5f54`](https://redirect.github.com/smallstep/certificates/commit/7df5f546dc8fd4b29e5b587eaf5aa6c835951e34) Bump golang.org/x/crypto from 0.29.0 to 0.30.0
- [`dbeedf3`](https://redirect.github.com/smallstep/certificates/commit/dbeedf342a883fd7932d37c72a313d48fc2e0e51) Merge pull request [#&#8203;2095](https://redirect.github.com/smallstep/certificates/issues/2095) from smallstep/dependabot/go\_modules/go.step.sm/crypto-0.55.0
- [`9f68609`](https://redirect.github.com/smallstep/certificates/commit/9f68609a85d27a7e6d2829bcff91ec61a3bbbec8) Merge pull request [#&#8203;2093](https://redirect.github.com/smallstep/certificates/issues/2093) from smallstep/dependabot/go\_modules/google.golang.org/api-0.210.0
- [`43400d6`](https://redirect.github.com/smallstep/certificates/commit/43400d6ffb1e4a17f25c889ffbdcd4a256bfa234) Bump github.com/slackhq/nebula from 1.9.4 to 1.9.5
- [`eb1b7e0`](https://redirect.github.com/smallstep/certificates/commit/eb1b7e090e0a4a280e2da35dd430a3c6808b22c2) Bump go.step.sm/crypto from 0.54.2 to 0.55.0
- [`abfcc59`](https://redirect.github.com/smallstep/certificates/commit/abfcc59ae8306b1878266e7e4746e09ef19bf6b6) Bump google.golang.org/api from 0.209.0 to 0.210.0
- [`6b8da39`](https://redirect.github.com/smallstep/certificates/commit/6b8da39183ddc61b44dbebdacd646b6af103ebea) Merge pull request [#&#8203;2082](https://redirect.github.com/smallstep/certificates/issues/2082) from smallstep/dependabot/go\_modules/cloud.google.com/go/longrunning-0.6.3
- [`8cffe10`](https://redirect.github.com/smallstep/certificates/commit/8cffe10cecbc238a7b370b8af546a8add28f51bc) Merge pull request [#&#8203;2084](https://redirect.github.com/smallstep/certificates/issues/2084) from smallstep/dependabot/go\_modules/google.golang.org/api-0.209.0
- [`09b917b`](https://redirect.github.com/smallstep/certificates/commit/09b917badd8fa116fbf3f6e01aa112079946453f) Merge pull request [#&#8203;2083](https://redirect.github.com/smallstep/certificates/issues/2083) from smallstep/dependabot/go\_modules/github.com/stretchr/testify-1.10.0
- [`e2e3f36`](https://redirect.github.com/smallstep/certificates/commit/e2e3f36d7798d14ada84d1568f92b2b6ea053758) Bump google.golang.org/api from 0.206.0 to 0.209.0
- [`29e1811`](https://redirect.github.com/smallstep/certificates/commit/29e1811c479864cc5bb4a966cd1cc8293d1a1e3c) Bump github.com/stretchr/testify from 1.9.0 to 1.10.0
- [`e93d542`](https://redirect.github.com/smallstep/certificates/commit/e93d542145b5028268747ec797c7811fc38df165) Bump cloud.google.com/go/longrunning from 0.6.2 to 0.6.3
- [`bc1accb`](https://redirect.github.com/smallstep/certificates/commit/bc1accbce132863721d1e1e99ec4ca2ca59c9c9b) add go.uber.org/mock to tools.go, generate mocks using go run ([#&#8203;2080](https://redirect.github.com/smallstep/certificates/issues/2080))

#### Thanks!

Those were the changes on v0.28.2!

Come join us on [Discord](https://discord.gg/X2RKGwEbV9) to ask questions, chat about PKI, or get a sneak peek at the freshest PKI memes.

### [`v0.28.1`](https://redirect.github.com/smallstep/certificates/releases/tag/v0.28.1): Step CA v0.28.1 (24-11-20)

[Compare Source](https://redirect.github.com/smallstep/certificates/compare/v0.28.0...v0.28.1)

#### Official Release Artifacts

##### Linux

- 📦 [step-ca\_linux\_0.28.1\_amd64.tar.gz](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.1/step-ca_linux_0.28.1_amd64.tar.gz)
- 📦 [step-ca\_0.28.1-1\_amd64.deb](https://dl.smallstep.com/gh-release/cli/gh-release-header/v0.28.1/step-ca_0.28.1-1_amd64.deb)
- 📦 [step-ca-0.28.1-1.x86\_64.rpm](https://dl.smallstep.com/gh-release/cli/gh-release-header/v0.28.1/step-ca-0.28.1-1.x86_64.rpm)
- 📦 [step-ca\_0.28.1-1\_arm64.deb](https://dl.smallstep.com/gh-release/cli/gh-release-header/v0.28.1/step-ca_0.28.1-1_arm64.deb)
- 📦 [step-ca-0.28.1-1.aarch64.rpm](https://dl.smallstep.com/gh-release/cli/gh-release-header/v0.28.1/step-ca-0.28.1-1.aarch64.rpm)

##### OSX Darwin

- 📦 [step-ca\_darwin\_0.28.1\_amd64.tar.gz](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.1/step-ca_darwin_0.28.1_amd64.tar.gz)
- 📦 [step-ca\_darwin\_0.28.1\_arm64.tar.gz](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.1/step-ca_darwin_0.28.1_arm64.tar.gz)

##### Windows

- 📦 [step-ca\_windows\_0.28.1\_amd64.zip](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.1/step-ca_windows_0.28.1_amd64.zip)

For more builds across platforms and architectures, see the `Assets` section below.
And for packaged versions (Docker, k8s, Homebrew), see our [installation docs](https://smallstep.com/docs/step-ca/installation).

Don't see the artifact you need? Open an issue [here](https://redirect.github.com/smallstep/certificates/issues/new/choose).

#### Signatures and Checksums

`step-ca` uses [sigstore/cosign](https://redirect.github.com/sigstore/cosign) for signing and verifying release artifacts.

Below is an example using `cosign` to verify a release artifact:

```
cosign verify-blob \
  --certificate step-ca_darwin_0.28.1_amd64.tar.gz.sig.pem \
  --signature step-ca_darwin_0.28.1_amd64.tar.gz.sig \
  --certificate-identity-regexp "https://github\.com/smallstep/workflows/.*" \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com \
  step-ca_darwin_0.28.1_amd64.tar.gz
```

The `checksums.txt` file (in the `Assets` section below) contains a checksum for every artifact in the release.

#### Changelog

- [`d327203`](https://redirect.github.com/smallstep/certificates/commit/d327203c1c2a1900bd019a0d9b38bd280fbd5c24) Add a few PRs to the 0.28.1 changelog ([#&#8203;2079](https://redirect.github.com/smallstep/certificates/issues/2079))
- [`f5b45d8`](https://redirect.github.com/smallstep/certificates/commit/f5b45d8ae18446ea0a884b6666937a06a6047183) Changelog for 0.28.1 ([#&#8203;2078](https://redirect.github.com/smallstep/certificates/issues/2078))
- [`f88a136`](https://redirect.github.com/smallstep/certificates/commit/f88a136bc23e62d32b7449de87cefda2a0f8dd00) Merge pull request [#&#8203;2076](https://redirect.github.com/smallstep/certificates/issues/2076) from smallstep/jdoss/repos
- [`354af7f`](https://redirect.github.com/smallstep/certificates/commit/354af7f914dff87d16fbbe8a4040d6653f9234d2) Merge branch 'master' into jdoss/repos
- [`c0d41d7`](https://redirect.github.com/smallstep/certificates/commit/c0d41d70ac4727a4cfbe35dcbbfdb6e82866bde0) Add p11-kit package.
- [`cd57b50`](https://redirect.github.com/smallstep/certificates/commit/cd57b50db2842cb19f36be7ae495ecee44af2632) Add in PKCS11 packages.
- [`7c9e3ff`](https://redirect.github.com/smallstep/certificates/commit/7c9e3ff1181b3e845d68a992a114b55454908dff) Merge pull request [#&#8203;2077](https://redirect.github.com/smallstep/certificates/issues/2077) from smallstep/jdoss/HSM\_container\_packages
- [`ae54b5a`](https://redirect.github.com/smallstep/certificates/commit/ae54b5a152549cc2c78395022fd0707ad00091ca) Add p11-kit package.
- [`e63b649`](https://redirect.github.com/smallstep/certificates/commit/e63b64904dc2659214e784bfb753cd968414471f) Make script +x.
- [`0851563`](https://redirect.github.com/smallstep/certificates/commit/08515639bd80022af5c7681ed909a908f260f1a6) Add in PKCS11 packages.
- [`330b28f`](https://redirect.github.com/smallstep/certificates/commit/330b28f83f19c5caa6e5d3c50eececf43632ea24) Merge branch 'master' into jdoss/repos
- [`b1c3d18`](https://redirect.github.com/smallstep/certificates/commit/b1c3d18d31bf5b59da403666534b79a6a31eb86f) Merge pull request [#&#8203;2072](https://redirect.github.com/smallstep/certificates/issues/2072) from smallstep/dependabot/go\_modules/google.golang.org/api-0.206.0
- [`e0196e4`](https://redirect.github.com/smallstep/certificates/commit/e0196e4860d72b8e2ea708f548f3fd425203a378) Bump google.golang.org/api from 0.205.0 to 0.206.0
- [`61af4c3`](https://redirect.github.com/smallstep/certificates/commit/61af4c3d79e7ebd5225882c26da550283b69936e) Merge pull request [#&#8203;2074](https://redirect.github.com/smallstep/certificates/issues/2074) from smallstep/dependabot/go\_modules/google.golang.org/protobuf-1.35.2
- [`2847fbd`](https://redirect.github.com/smallstep/certificates/commit/2847fbd3f362636169c3c1b3ed5174f9e1efdc17) Merge pull request [#&#8203;2073](https://redirect.github.com/smallstep/certificates/issues/2073) from smallstep/dependabot/go\_modules/go.step.sm/crypto-0.54.2
- [`b9296c1`](https://redirect.github.com/smallstep/certificates/commit/b9296c1b6fcd45613d444f01b492bcd89933f764) Merge pull request [#&#8203;2075](https://redirect.github.com/smallstep/certificates/issues/2075) from smallstep/dependabot/go\_modules/github.com/googleapis/gax-go/v2-2.14.0
- [`2b4894a`](https://redirect.github.com/smallstep/certificates/commit/2b4894a87136757bd01d1cf2be3c591f086b5572) Configure GitHub Actions to publish RPMs and Debs to packages.smallstep.com.
- [`a0048b4`](https://redirect.github.com/smallstep/certificates/commit/a0048b414bbeb95b8a66e61128c18855ca867ab2) Bump github.com/googleapis/gax-go/v2 from 2.13.0 to 2.14.0
- [`e9a1a7b`](https://redirect.github.com/smallstep/certificates/commit/e9a1a7ba158a84f6d6dbae4f099fe966036e62cb) Bump google.golang.org/protobuf from 1.35.1 to 1.35.2
- [`35bc30b`](https://redirect.github.com/smallstep/certificates/commit/35bc30b1f5535a59a4693c946c00d76cbfd52ed9) Bump go.step.sm/crypto from 0.54.0 to 0.54.2
- [`707f537`](https://redirect.github.com/smallstep/certificates/commit/707f53722f23363cc5dc873434d87085d91ade7a) Merge pull request [#&#8203;2071](https://redirect.github.com/smallstep/certificates/issues/2071) from smallstep/dependabot/github\_actions/softprops/action-gh-release-2.1.0
- [`fc007df`](https://redirect.github.com/smallstep/certificates/commit/fc007dfd039e6821a736d03b7170203e0d0e7e64) Bump softprops/action-gh-release from 2.0.9 to 2.1.0
- [`7eccbc5`](https://redirect.github.com/smallstep/certificates/commit/7eccbc51febc2a8f64142d3c8b35595cd3a22cba) Merge pull request [#&#8203;2069](https://redirect.github.com/smallstep/certificates/issues/2069) from smallstep/mariano/urn
- [`1a2e647`](https://redirect.github.com/smallstep/certificates/commit/1a2e64724ad944701a740005f3ed26ec2e6b2ecb) Change URN for acme errors
- [`73f97e2`](https://redirect.github.com/smallstep/certificates/commit/73f97e244c02bd0a01b3c84f574063bca58b75b7) Ignore non-constant format string linting error ([#&#8203;2068](https://redirect.github.com/smallstep/certificates/issues/2068))
- [`fdaa665`](https://redirect.github.com/smallstep/certificates/commit/fdaa6657e1674a7231949af8de3aebb9dc12f8c2) Merge pull request [#&#8203;2066](https://redirect.github.com/smallstep/certificates/issues/2066) from smallstep/mariano/webhook-errors
- [`05295d9`](https://redirect.github.com/smallstep/certificates/commit/05295d9c6a8b0baaad0a351fb685594191c08db5) Propagate human errors from webhooks
- [`2f7690c`](https://redirect.github.com/smallstep/certificates/commit/2f7690c2ae1e91460dd4590cca0055e1dce1afad) Merge pull request [#&#8203;2065](https://redirect.github.com/smallstep/certificates/issues/2065) from smallstep/mariano/challenge-webhook
- [`ff37bf1`](https://redirect.github.com/smallstep/certificates/commit/ff37bf181190ef5143ef9588c22c0408a1c99b54) Add unit tests for scepchallenge webhooks
- [`f2663dd`](https://redirect.github.com/smallstep/certificates/commit/f2663dd9d9bd3bc13296d7fd3fc30f4d3a0ae9fb) Add data support on SCEPCHALLENGE webhooks
- [`295892e`](https://redirect.github.com/smallstep/certificates/commit/295892ecba96a058a7229e99f7c782f888635550) Merge pull request [#&#8203;2060](https://redirect.github.com/smallstep/certificates/issues/2060) from smallstep/dependabot/go\_modules/google.golang.org/api-0.205.0
- [`15e9041`](https://redirect.github.com/smallstep/certificates/commit/15e904109a6c604d0b8c0e79d9f2dab7afc9bfa4) Bump google.golang.org/api from 0.204.0 to 0.205.0
- [`8465db8`](https://redirect.github.com/smallstep/certificates/commit/8465db8b308f76cd896de23c7cf66efc7ff1eb2e) Merge pull request [#&#8203;2062](https://redirect.github.com/smallstep/certificates/issues/2062) from smallstep/dependabot/go\_modules/golang.org/x/net-0.31.0
- [`8313e47`](https://redirect.github.com/smallstep/certificates/commit/8313e4706180c326d022c410bdbf86d2f8ea0b06) Bump golang.org/x/net from 0.30.0 to 0.31.0
- [`c4c0d92`](https://redirect.github.com/smallstep/certificates/commit/c4c0d923d6748cc4d221b05a8d4a034e29a55d10) Merge pull request [#&#8203;2061](https://redirect.github.com/smallstep/certificates/issues/2061) from smallstep/dependabot/go\_modules/golang.org/x/crypto-0.29.0
- [`6dda3c1`](https://redirect.github.com/smallstep/certificates/commit/6dda3c1c4d836c7dade7b7d318f003f2f837cd6d) Bump golang.org/x/crypto from 0.28.0 to 0.29.0
- [`e82c228`](https://redirect.github.com/smallstep/certificates/commit/e82c228a530968e553b475637a6ee789a8f0fce9) Merge pull request [#&#8203;2057](https://redirect.github.com/smallstep/certificates/issues/2057) from smallstep/dependabot/github\_actions/softprops/action-gh-release-2.0.9
- [`59deaa5`](https://redirect.github.com/smallstep/certificates/commit/59deaa5180735fb4e63bbe3220846723f97cb002) Merge pull request [#&#8203;2058](https://redirect.github.com/smallstep/certificates/issues/2058) from smallstep/dependabot/go\_modules/google.golang.org/api-0.204.0
- [`1ab27b1`](https://redirect.github.com/smallstep/certificates/commit/1ab27b13e5208ec387d0634949588ea2b08cb9b7) Bump google.golang.org/api from 0.203.0 to 0.204.0
- [`0a0dfc0`](https://redirect.github.com/smallstep/certificates/commit/0a0dfc008eef0f9747276a896f58bd9d788106d4) Bump softprops/action-gh-release from 2.0.8 to 2.0.9

#### Thanks!

Those were the changes on v0.28.1!

Come join us on [Discord](https://discord.gg/X2RKGwEbV9) to ask questions, chat about PKI, or get a sneak peek at the freshest PKI memes.

### [`v0.28.0`](https://redirect.github.com/smallstep/certificates/releases/tag/v0.28.0): Step CA v0.28.0 (24-10-30)

[Compare Source](https://redirect.github.com/smallstep/certificates/compare/v0.27.5...v0.28.0)

#### Official Release Artifacts

##### Linux

- 📦 [step-ca\_linux\_0.28.0\_amd64.tar.gz](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.0/step-ca_linux_0.28.0_amd64.tar.gz)
- 📦 [step-ca\_0.28.0\_amd64.deb](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.0/step-ca_0.28.0_amd64.deb)

##### OSX Darwin

- 📦 [step-ca\_darwin\_0.28.0\_amd64.tar.gz](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.0/step-ca_darwin_0.28.0_amd64.tar.gz)
- 📦 [step-ca\_darwin\_0.28.0\_arm64.tar.gz](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.0/step-ca_darwin_0.28.0_arm64.tar.gz)

##### Windows

- 📦 [step-ca\_windows\_0.28.0\_amd64.zip](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.28.0/step-ca_windows_0.28.0_amd64.zip)

For more builds across platforms and architectures, see the `Assets` section below.
And for packaged versions (Docker, k8s, Homebrew), see our [installation docs](https://smallstep.com/docs/step-ca/installation).

Don't see the artifact you need? Open an issue [here](https://redirect.github.com/smallstep/certificates/issues/new/choose).

#### Signatures and Checksums

`step-ca` uses [sigstore/cosign](https://redirect.github.com/sigstore/cosign) for signing and verifying release artifacts.

Below is an example using `cosign` to verify a release artifact:

```
cosign verify-blob \
  --certificate step-ca_darwin_0.28.0_amd64.tar.gz.sig.pem \
  --signature step-ca_darwin_0.28.0_amd64.tar.gz.sig \
  --certificate-identity-regexp "https://github\.com/smallstep/workflows/.*" \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com \
  step-ca_darwin_0.28.0_amd64.tar.gz
```

The `checksums.txt` file (in the `Assets` section below) contains a checksum for every artifact in the release.

#### Changelog

- [`df13aae`](https://redirect.github.com/smallstep/certificates/commit/df13aaef0b9cb5f0d6e9c011695214dbafff7aca) Update changelog for v0.28.0 ([#&#8203;2051](https://redirect.github.com/smallstep/certificates/issues/2051))
- [`77667e7`](https://redirect.github.com/smallstep/certificates/commit/77667e7f43fba96a720b07870856d75fa3503dca) Merge pull request [#&#8203;2049](https://redirect.github.com/smallstep/certificates/issues/2049) from smallstep/dependabot/go\_modules/cloud.google.com/go/security-1.18.2
- [`5147c60`](https://redirect.github.com/smallstep/certificates/commit/5147c6024049529b128cca9e9bc23f520b5c6331) Bump cloud.google.com/go/security from 1.18.1 to 1.18.2
- [`efd324d`](https://redirect.github.com/smallstep/certificates/commit/efd324d6d62c58d848f303f46091ce8c7c95fca2) Merge pull request [#&#8203;2048](https://redirect.github.com/smallstep/certificates/issues/2048) from smallstep/dependabot/go\_modules/cloud.google.com/go/longrunning-0.6.2
- [`a51f670`](https://redirect.github.com/smallstep/certificates/commit/a51f670f77179b9d69c9506e69caa0d45204fcaf) Merge pull request [#&#8203;2047](https://redirect.github.com/smallstep/certificates/issues/2047) from smallstep/dependabot/go\_modules/github.com/newrelic/go-agent/v3-3.35.1
- [`9943bf1`](https://redirect.github.com/smallstep/certificates/commit/9943bf1984b249bfac5820d5f7507a92c3b3b367) Merge pull request [#&#8203;2050](https://redirect.github.com/smallstep/certificates/issues/2050) from smallstep/dependabot/go\_modules/google.golang.org/api-0.203.0
- [`bb8605c`](https://redirect.github.com/smallstep/certificates/commit/bb8605c079ca4aa5cc459ee3b3641a7735d93322) Add DisableSSHCAUser and DisableSSHCAHost options to linkedca GCP provisioner ([#&#8203;2045](https://redirect.github.com/smallstep/certificates/issues/2045))
- [`95a6cad`](https://redirect.github.com/smallstep/certificates/commit/95a6cad404fe5065bd74f3218da248a4b2840c84) Bump google.golang.org/api from 0.201.0 to 0.203.0
- [`99baf67`](https://redirect.github.com/smallstep/certificates/commit/99baf6744aa02732fa3c8a2a5e3f95370ea0e5f0) Bump cloud.google.com/go/longrunning from 0.6.1 to 0.6.2
- [`bfe436b`](https://redirect.github.com/smallstep/certificates/commit/bfe436b14560519d4c863a39dc713a0742414235) Bump github.com/newrelic/go-agent/v3 from 3.35.0 to 3.35.1
- [`34ba7a2`](https://redirect.github.com/smallstep/certificates/commit/34ba7a2f3eff1ae7da44dd3784eec4c15682f6ac) Merge pull request [#&#8203;2046](https://redirect.github.com/smallstep/certificates/issues/2046) from smallstep/herman/refactor-cli-utils-import
- [`b45b73f`](https://redirect.github.com/smallstep/certificates/commit/b45b73f4cc9746b7c91252dccc6bf374002b3351) Use `github.com/smallstep/cli-utils` instead of `go.step.sm/cli-utils`
- [`88443dd`](https://redirect.github.com/smallstep/certificates/commit/88443ddab988c5f1f80c4bd768342391326c7d07) Use dnsNamesSubsetValidator for IID provisioners ([#&#8203;2044](https://redirect.github.com/smallstep/certificates/issues/2044))
- [`946bba2`](https://redirect.github.com/smallstep/certificates/commit/946bba29bb1f02fde5e2c2f31cc63550aebb8b7e) Merge pull request [#&#8203;2040](https://redirect.github.com/smallstep/certificates/issues/2040) from smallstep/dependabot/go\_modules/github.com/prometheus/client\_golang-1.20.5
- [`61fb5f4`](https://redirect.github.com/smallstep/certificates/commit/61fb5f4071c10b189056130b9c0341a3eb51b0ec) Merge pull request [#&#8203;2039](https://redirect.github.com/smallstep/certificates/issues/2039) from smallstep/dependabot/go\_modules/google.golang.org/api-0.201.0
- [`5afe6b1`](https://redirect.github.com/smallstep/certificates/commit/5afe6b1e549ebfbdc823206e3c82361d84f1da94) Bump github.com/prometheus/client\_golang from 1.20.4 to 1.20.5
- [`6ce8f7d`](https://redirect.github.com/smallstep/certificates/commit/6ce8f7d606d98d0d7f13ace1a2818fde18a2a82d) Bump google.golang.org/api from 0.200.0 to 0.201.0

#### Thanks!

Those were the changes on v0.28.0!

Come join us on [Discord](https://discord.gg/X2RKGwEbV9) to ask questions, chat about PKI, or get a sneak peek at the freshest PKI memes.

### [`v0.27.5`](https://redirect.github.com/smallstep/certificates/releases/tag/v0.27.5): Step CA v0.27.5 (24-10-17)

[Compare Source](https://redirect.github.com/smallstep/certificates/compare/v0.27.4...v0.27.5)

#### Official Release Artifacts

##### Linux

- 📦 [step-ca\_linux\_0.27.5\_amd64.tar.gz](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.27.5/step-ca_linux_0.27.5_amd64.tar.gz)
- 📦 [step-ca\_0.27.5\_amd64.deb](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.27.5/step-ca_0.27.5_amd64.deb)

##### OSX Darwin

- 📦 [step-ca\_darwin\_0.27.5\_amd64.tar.gz](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.27.5/step-ca_darwin_0.27.5_amd64.tar.gz)
- 📦 [step-ca\_darwin\_0.27.5\_arm64.tar.gz](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.27.5/step-ca_darwin_0.27.5_arm64.tar.gz)

##### Windows

- 📦 [step-ca\_windows\_0.27.5\_amd64.zip](https://dl.smallstep.com/gh-release/certificates/gh-release-header/v0.27.5/step-ca_windows_0.27.5_amd64.zip)

For more builds across platforms and architectures, see the `Assets` section below.
And for packaged versions (Docker, k8s, Homebrew), see our [installation docs](https://smallstep.com/docs/step-ca/installation).

Don't see the artifact you need? Open an issue [here](https://redirect.github.com/smallstep/certificates/issues/new/choose).

#### Signatures and Checksums

`step-ca` uses [sigstore/cosign](https://redirect.github.com/sigstore/cosign) for signing and verifying release artifacts.

Below is an example using `cosign` to verify a release artifact:

```
cosign verify-blob \
  --certificate step-ca_darwin_0.27.5_amd64.tar.gz.sig.pem \
  --signature step-ca_darwin_0.27.5_amd64.tar.gz.sig \
  --certificate-identity-regexp "https://github\.com/smallstep/workflows/.*" \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com \
  step-ca_darwin_0.27.5_amd64.tar.gz
```

The `checksums.txt` file (in the `Assets` section below) contains a checksum for every artifact in the release.

#### Changelog

- [`

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

### Comment by @red-hat-konflux on November 07, 2025 at 08:18 PM UTC

### ℹ Artifact update notice

##### File name: rh_identity_transform/go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 3 additional dependencies were updated


Details:


| **Package**                           | **Change**              |
| :------------------------------------ | :---------------------- |
| `github.com/klauspost/compress`       | `v1.17.11` -> `v1.18.0` |
| `github.com/prometheus/client_golang` | `v1.20.4` -> `v1.22.0`  |
| `go.etcd.io/bbolt`                    | `v1.3.9` -> `v1.3.10`   |

---

## Reviews

### Review by @Hyperkid123 - Approved on November 11, 2025 at 02:13 PM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/55*
