---
type: pull_request
number: 1477
title: "Build: Bump go.opentelemetry.io/otel from 1.40.0 to 1.41.0"
state: merged
author: dependabot
created: 2026-04-24T20:53:35Z
updated: 2026-04-27T07:30:11Z
closed: 2026-04-27T07:30:09Z
merged: 2026-04-27T07:30:09Z
base_branch: main
head_branch: dependabot/go_modules/go.opentelemetry.io/otel-1.41.0
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1477
---

# Pull Request #1477: Build: Bump go.opentelemetry.io/otel from 1.40.0 to 1.41.0

**Author**: @dependabot
**Created**: April 24, 2026 at 08:53 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go.opentelemetry.io/otel-1.41.0`

## Description

Bumps [go.opentelemetry.io/otel](https://github.com/open-telemetry/opentelemetry-go) from 1.40.0 to 1.41.0.
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/open-telemetry/opentelemetry-go/blob/main/CHANGELOG.md">go.opentelemetry.io/otel's changelog</a>.</em></p>
<blockquote>
<h2>[1.41.0/0.63.0/0.17.0/0.0.15] 2026-03-02</h2>
<p>This release is the last to support [Go 1.24].
The next release will require at least [Go 1.25].</p>
<h3>Added</h3>
<ul>
<li>Support testing of [Go 1.26]. (<a href="https://redirect.github.com/open-telemetry/opentelemetry-go/issues/7902">#7902</a>)</li>
</ul>
<h3>Fixed</h3>
<ul>
<li>Update <code>Baggage</code> in <code>go.opentelemetry.io/otel/propagation</code> and <code>Parse</code> and <code>New</code> in <code>go.opentelemetry.io/otel/baggage</code> to comply with W3C Baggage specification limits.
<code>New</code> and <code>Parse</code> now return partial baggage along with an error when limits are exceeded.
Errors from baggage extraction are reported to the global error handler. (<a href="https://redirect.github.com/open-telemetry/opentelemetry-go/issues/7880">#7880</a>)</li>
<li>Return an error when the endpoint is configured as insecure and with TLS configuration in <code>go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp</code>. (<a href="https://redirect.github.com/open-telemetry/opentelemetry-go/issues/7914">#7914</a>)</li>
<li>Return an error when the endpoint is configured as insecure and with TLS configuration in <code>go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp</code>. (<a href="https://redirect.github.com/open-telemetry/opentelemetry-go/issues/7914">#7914</a>)</li>
<li>Return an error when the endpoint is configured as insecure and with TLS configuration in <code>go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp</code>. (<a href="https://redirect.github.com/open-telemetry/opentelemetry-go/issues/7914">#7914</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/open-telemetry/opentelemetry-go/commit/4575a9774dd9415ffc858dd34955493b0031065a"><code>4575a97</code></a> Release 1.41.0/0.63.0/0.17.0/0.0.15 (<a href="https://redirect.github.com/open-telemetry/opentelemetry-go/issues/7977">#7977</a>)</li>
<li><a href="https://github.com/open-telemetry/opentelemetry-go/commit/66fc10d9dff9653c65bcca111b965137d06f09aa"><code>66fc10d</code></a> fix: add error handling for insecure HTTP endpoints with TLS client configura...</li>
<li><a href="https://github.com/open-telemetry/opentelemetry-go/commit/76e6eec88f186f06a0708b5620324d2b002d9a97"><code>76e6eec</code></a> chore(deps): update github/codeql-action action to v4.32.5 (<a href="https://redirect.github.com/open-telemetry/opentelemetry-go/issues/7980">#7980</a>)</li>
<li><a href="https://github.com/open-telemetry/opentelemetry-go/commit/0d50f9008c8c93fe49a7caa45c88c30370479d27"><code>0d50f90</code></a> Revert &quot;Generate semconv/v1.40.0&quot; (<a href="https://redirect.github.com/open-telemetry/opentelemetry-go/issues/7978">#7978</a>)</li>
<li><a href="https://github.com/open-telemetry/opentelemetry-go/commit/c38a4a57c320b6098ca5c92f0a85201034780b1f"><code>c38a4a5</code></a> Generate semconv/v1.40.0 (<a href="https://redirect.github.com/open-telemetry/opentelemetry-go/issues/7929">#7929</a>)</li>
<li><a href="https://github.com/open-telemetry/opentelemetry-go/commit/0f1a22484ec52d6beb1efdb0fa1b63a31e7405af"><code>0f1a224</code></a> chore(deps): update module github.com/securego/gosec/v2 to v2.23.0 (<a href="https://redirect.github.com/open-telemetry/opentelemetry-go/issues/7899">#7899</a>)</li>
<li><a href="https://github.com/open-telemetry/opentelemetry-go/commit/c79ebf43eb1cff6dd76a33bb1549f2c082dab604"><code>c79ebf4</code></a> chore(deps): update module github.com/daixiang0/gci to v0.14.0 (<a href="https://redirect.github.com/open-telemetry/opentelemetry-go/issues/7973">#7973</a>)</li>
<li><a href="https://github.com/open-telemetry/opentelemetry-go/commit/f75815746541d0d4ac84e1c5955bdcd1a2df2d7d"><code>f758157</code></a> chore(deps): update module github.com/sonatard/noctx to v0.5.0 (<a href="https://redirect.github.com/open-telemetry/opentelemetry-go/issues/7968">#7968</a>)</li>
<li><a href="https://github.com/open-telemetry/opentelemetry-go/commit/92a11645724515630187def073ae39f1b6cb3c69"><code>92a1164</code></a> fix(deps): update github.com/opentracing-contrib/go-grpc/test digest to d566b...</li>
<li><a href="https://github.com/open-telemetry/opentelemetry-go/commit/3cd7c27e840ea3114115459db2e299a27fffaff8"><code>3cd7c27</code></a> chore(deps): update module github.com/protonmail/go-crypto to v1.4.0 (<a href="https://redirect.github.com/open-telemetry/opentelemetry-go/issues/7969">#7969</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/open-telemetry/opentelemetry-go/compare/v1.40.0...v1.41.0">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=go.opentelemetry.io/otel&package-manager=go_modules&previous-version=1.40.0&new-version=1.41.0)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

Dependabot will resolve any conflicts with this PR as long as you don't alter it yourself. You can also trigger a rebase manually by commenting `@dependabot rebase`.

[//]: # (dependabot-automerge-start)
[//]: # (dependabot-automerge-end)

---

<details>
<summary>Dependabot commands and options</summary>
<br />

You can trigger Dependabot actions by commenting on this PR:
- `@dependabot rebase` will rebase this PR
- `@dependabot recreate` will recreate this PR, overwriting any edits that have been made to it
- `@dependabot show <dependency name> ignore conditions` will show all of the ignore conditions of the specified dependency
- `@dependabot ignore this major version` will close this PR and stop Dependabot creating any more for this major version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this minor version` will close this PR and stop Dependabot creating any more for this minor version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this dependency` will close this PR and stop Dependabot creating any more for this dependency (unless you reopen the PR or upgrade to it yourself)
You can disable automated security fix PRs for this repo from the [Security Alerts page](https://github.com/content-services/content-sources-backend/network/alerts).

</details>

---

## Reviews

### Review by @swadeley - Approved on April 27, 2026 at 07:29 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1477*
