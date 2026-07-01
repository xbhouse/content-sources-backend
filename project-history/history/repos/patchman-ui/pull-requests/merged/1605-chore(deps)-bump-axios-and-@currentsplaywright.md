---
type: pull_request
number: 1605
title: "chore(deps): bump axios and @currents/playwright"
state: merged
author: dependabot
created: 2026-05-08T00:54:21Z
updated: 2026-05-08T14:24:40Z
closed: 2026-05-08T14:24:31Z
merged: 2026-05-08T14:24:31Z
base_branch: master
head_branch: dependabot/npm_and_yarn/multi-8aafb98a29
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1605
---

# Pull Request #1605: chore(deps): bump axios and @currents/playwright

**Author**: @dependabot
**Created**: May 08, 2026 at 12:54 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/multi-8aafb98a29`

## Description

Bumps [axios](https://github.com/axios/axios) and [@currents/playwright](https://github.com/currents-dev/currents-playwright-changelog). These dependencies needed to be updated together.
Updates `axios` from 1.15.0 to 1.16.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/axios/axios/releases">axios's releases</a>.</em></p>
<blockquote>
<h2>v1.16.0 — May 2, 2026</h2>
<p>This release adds support for the QUERY HTTP method and a new <code>ECONNREFUSED</code> error constant, lands a substantial wave of HTTP, fetch, and XHR adapter bug fixes around redirects, aborts, headers, and timeouts, and welcomes 23 new contributors.</p>
<h2>⚠️ Notable Changes</h2>
<p>A handful of fixes in this release are either security-adjacent or change observable behaviour. Please review before upgrading:</p>
<ul>
<li><strong>Fetch adapter now enforces <code>maxBodyLength</code> and <code>maxContentLength</code>.</strong> These limits were silently ignored on the fetch adapter prior to 1.16.0 — anyone relying on them as a safety net (DoS protection, accidental large uploads) had no protection. (<strong><a href="https://redirect.github.com/axios/axios/issues/10795">#10795</a></strong>)</li>
<li><strong>Proxy requests now preserve user-supplied <code>Host</code> headers.</strong> Previously, the proxy path could overwrite a custom <code>Host</code>. Virtual-host-style routing through a proxy will now behave correctly. (<strong><a href="https://redirect.github.com/axios/axios/issues/10822">#10822</a></strong>)</li>
<li><strong>Basic auth credentials embedded in URLs are now URL-decoded.</strong> If you have percent-encoded credentials in a URL (e.g. <code>https://user:p%40ss@host</code>), the decoded value is what now goes on the wire. (<strong><a href="https://redirect.github.com/axios/axios/issues/10825">#10825</a></strong>)</li>
<li><strong><code>parseProtocol</code> now strictly requires a colon in the protocol separator.</strong> Strings that loosely parsed as protocols before may no longer match. (<strong><a href="https://redirect.github.com/axios/axios/issues/10729">#10729</a></strong>)</li>
<li><strong>Deprecated <code>unescape()</code> replaced with modern UTF-8 encoding.</strong> Non-ASCII URL handling is now spec-correct; consumers depending on legacy <code>unescape()</code> quirks may see different output bytes. (<strong><a href="https://redirect.github.com/axios/axios/issues/7378">#7378</a></strong>)</li>
<li><strong><code>transformRequest</code> input typing change was reverted.</strong> The typing change introduced in <a href="https://redirect.github.com/axios/axios/issues/10745">#10745</a> was reverted in <a href="https://redirect.github.com/axios/axios/issues/10810">#10810</a> after follow-up review — net behavior is unchanged from 1.15.2. (<strong><a href="https://redirect.github.com/axios/axios/issues/10745">#10745</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10810">#10810</a></strong>)</li>
</ul>
<h2>🚀 New Features</h2>
<ul>
<li><strong>QUERY HTTP Method:</strong> Added support for the QUERY HTTP method across adapters and type definitions. (<strong><a href="https://redirect.github.com/axios/axios/issues/10802">#10802</a></strong>)</li>
<li><strong>ECONNREFUSED Error Constant:</strong> Exposed <code>ECONNREFUSED</code> as a constant on <code>AxiosError</code> so callers can match connection-refused failures without comparing string literals (closes <a href="https://redirect.github.com/axios/axios/issues/6485">#6485</a>). (<strong><a href="https://redirect.github.com/axios/axios/issues/10680">#10680</a></strong>)</li>
<li><strong>Encode Helper Export:</strong> Exported the internal <code>encode</code> helper from <code>buildURL</code> so userland param serializers can reuse the same encoding logic that axios uses internally. (<strong><a href="https://redirect.github.com/axios/axios/issues/6897">#6897</a></strong>)</li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li><strong>HTTP Adapter — Redirects &amp; Headers:</strong> Cleared stale headers when a redirect targets a no-proxy host, fixed the redirect listener chain so listeners no longer stack across hops, restored the missing <code>requestDetails</code> argument on <code>beforeRedirect</code>, preserved user-supplied <code>Host</code> headers when forwarding through a proxy, and properly URL-decoded basic auth credentials. (<strong><a href="https://redirect.github.com/axios/axios/issues/10794">#10794</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10800">#10800</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/6241">#6241</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10822">#10822</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10825">#10825</a></strong>)</li>
<li><strong>HTTP Adapter — Streams &amp; Timeouts:</strong> Preserved the partial response object on <code>AxiosError</code> when a stream is aborted after headers arrive, honoured the <code>timeout</code> option during the connect phase when redirects are disabled, and resolved an unsettled-promise hang when an aborted request was combined with compression and <code>maxRedirects: 0</code>. (<strong><a href="https://redirect.github.com/axios/axios/issues/10708">#10708</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10819">#10819</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/7149">#7149</a></strong>)</li>
<li><strong>Fetch Adapter:</strong> Enforced <code>maxBodyLength</code> / <code>maxContentLength</code> in the fetch adapter, set the <code>User-Agent</code> header to match the HTTP adapter, preserved the original abort reason instead of replacing it with a generic error, and deferred global access so importing the module no longer throws a <code>TypeError</code> in restricted environments. (<strong><a href="https://redirect.github.com/axios/axios/issues/10795">#10795</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10772">#10772</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10806">#10806</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/7260">#7260</a></strong>)</li>
<li><strong>XHR Adapter:</strong> Unsubscribed the <code>cancelToken</code> and <code>AbortSignal</code> listeners on the error, timeout, and abort code paths to prevent leaked subscriptions. (<strong><a href="https://redirect.github.com/axios/axios/issues/10787">#10787</a></strong>)</li>
<li><strong>Error Handling:</strong> Attached the parsed response to <code>AxiosError</code> when <code>JSON.parse</code> fails inside <code>dispatchRequest</code>, prevented <code>settle</code> from emitting <code>undefined</code> error codes, and tightened the <code>parseProtocol</code> regex to require a colon in the protocol separator. (<strong><a href="https://redirect.github.com/axios/axios/issues/10724">#10724</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/7276">#7276</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10729">#10729</a></strong>)</li>
<li><strong>Types &amp; Exports:</strong> Aligned the CommonJS <code>CancelToken</code> typings with the ESM build, fixed a compiler error caused by <code>RawAxiosHeaders</code>, and re-exported <code>create</code> from the package index. (<strong><a href="https://redirect.github.com/axios/axios/issues/7414">#7414</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/6389">#6389</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/6460">#6460</a></strong>)</li>
<li><strong>UTF-8 Encoding:</strong> Replaced the deprecated <code>unescape()</code> call with a modern UTF-8 encoding implementation. (<strong><a href="https://redirect.github.com/axios/axios/issues/7378">#7378</a></strong>)</li>
<li><strong>Misc Cleanup:</strong> Resolved a batch of small inconsistencies and gadget-level issues across the codebase. (<strong><a href="https://redirect.github.com/axios/axios/issues/10833">#10833</a></strong>)</li>
</ul>
<h2>🔧 Maintenance &amp; Chores</h2>
<ul>
<li><strong>Refactor — ES6 Modernisation:</strong> Modernised the <code>utils</code> module and XHR adapter to use ES6 features, and tidied the multipart boundary error message. (<strong><a href="https://redirect.github.com/axios/axios/issues/10588">#10588</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/7419">#7419</a></strong>)</li>
<li><strong>Tests:</strong> Hardened the HTTP test server lifecycle to fix flaky <code>FormData</code> EPIPE failures, fixed Win32 platform support for the pipe tests, and corrected an incorrect test assumption. (<strong><a href="https://redirect.github.com/axios/axios/issues/10820">#10820</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10791">#10791</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10796">#10796</a></strong>)</li>
<li><strong>Docs:</strong> Documented <code>paramsSerializer.encode</code> for strict RFC 3986 query encoding, updated the <code>parseReviver</code> TypeScript definitions and configuration docs for ES2023, added timeout guidance to the README's first async example, and expanded notes around the recent type changes. (<strong><a href="https://redirect.github.com/axios/axios/issues/10821">#10821</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10782">#10782</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10759">#10759</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10804">#10804</a></strong>)</li>
<li><strong>Reverted:</strong> Reverted the <code>transformRequest</code> input typing change from <a href="https://redirect.github.com/axios/axios/issues/10745">#10745</a> after follow-up review. (<strong><a href="https://redirect.github.com/axios/axios/issues/10745">#10745</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10810">#10810</a></strong>)</li>
<li><strong>Dependencies:</strong> Bumped <code>actions/setup-node</code>, the <code>github-actions</code> group, and <code>postcss</code> (in <code>/docs</code>) to their latest versions. (<strong><a href="https://redirect.github.com/axios/axios/issues/10785">#10785</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10813">#10813</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10814">#10814</a></strong>)</li>
<li><strong>Release:</strong> Updated changelog and packages, and prepared the 1.16.0 release. (<strong><a href="https://redirect.github.com/axios/axios/issues/10790">#10790</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10834">#10834</a></strong>)</li>
</ul>
<h2>🌟 New Contributors</h2>
<p>We are thrilled to welcome our new contributors. Thank you for helping improve axios:</p>
<ul>
<li><strong><a href="https://github.com/singhankit001"><code>@​singhankit001</code></a></strong> (<strong><a href="https://redirect.github.com/axios/axios/issues/10588">#10588</a></strong>)</li>
<li><strong><a href="https://github.com/cuiweixie"><code>@​cuiweixie</code></a></strong> (<strong><a href="https://redirect.github.com/axios/axios/issues/7419">#7419</a></strong>)</li>
<li><strong><a href="https://github.com/iruizsalinas"><code>@​iruizsalinas</code></a></strong> (<strong><a href="https://redirect.github.com/axios/axios/issues/10787">#10787</a></strong>)</li>
<li><strong><a href="https://github.com/MarcosNocetti"><code>@​MarcosNocetti</code></a></strong> (<strong><a href="https://redirect.github.com/axios/axios/issues/10680">#10680</a></strong>)</li>
<li><strong><a href="https://github.com/deepview-autofix"><code>@​deepview-autofix</code></a></strong> (<strong><a href="https://redirect.github.com/axios/axios/issues/10729">#10729</a></strong>)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/axios/axios/blob/v1.x/CHANGELOG.md">axios's changelog</a>.</em></p>
<blockquote>
<h2>v1.16.0 — May 2, 2026</h2>
<p>This release adds support for the QUERY HTTP method and a new <code>ECONNREFUSED</code> error constant, lands a substantial wave of HTTP, fetch, and XHR adapter bug fixes around redirects, aborts, headers, and timeouts, and welcomes 23 new contributors.</p>
<h2>⚠️ Notable Changes</h2>
<p>A handful of fixes in this release are either security-adjacent or change observable behaviour. Please review before upgrading:</p>
<ul>
<li><strong>Fetch adapter now enforces <code>maxBodyLength</code> and <code>maxContentLength</code>.</strong> These limits were silently ignored on the fetch adapter prior to 1.16.0 — anyone relying on them as a safety net (DoS protection, accidental large uploads) had no protection. (<strong><a href="https://redirect.github.com/axios/axios/issues/10795">#10795</a></strong>)</li>
<li><strong>Proxy requests now preserve user-supplied <code>Host</code> headers.</strong> Previously, the proxy path could overwrite a custom <code>Host</code>. Virtual-host-style routing through a proxy will now behave correctly. (<strong><a href="https://redirect.github.com/axios/axios/issues/10822">#10822</a></strong>)</li>
<li><strong>Basic auth credentials embedded in URLs are now URL-decoded.</strong> If you have percent-encoded credentials in a URL (e.g. <code>https://user:p%40ss@host</code>), the decoded value is what now goes on the wire. (<strong><a href="https://redirect.github.com/axios/axios/issues/10825">#10825</a></strong>)</li>
<li><strong><code>parseProtocol</code> now strictly requires a colon in the protocol separator.</strong> Strings that loosely parsed as protocols before may no longer match. (<strong><a href="https://redirect.github.com/axios/axios/issues/10729">#10729</a></strong>)</li>
<li><strong>Deprecated <code>unescape()</code> replaced with modern UTF-8 encoding.</strong> Non-ASCII URL handling is now spec-correct; consumers depending on legacy <code>unescape()</code> quirks may see different output bytes. (<strong><a href="https://redirect.github.com/axios/axios/issues/7378">#7378</a></strong>)</li>
<li><strong><code>transformRequest</code> input typing change was reverted.</strong> The typing change introduced in <a href="https://redirect.github.com/axios/axios/issues/10745">#10745</a> was reverted in <a href="https://redirect.github.com/axios/axios/issues/10810">#10810</a> after follow-up review — net behavior is unchanged from 1.15.2. (<strong><a href="https://redirect.github.com/axios/axios/issues/10745">#10745</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10810">#10810</a></strong>)</li>
</ul>
<h2>🚀 New Features</h2>
<ul>
<li><strong>QUERY HTTP Method:</strong> Added support for the QUERY HTTP method across adapters and type definitions. (<strong><a href="https://redirect.github.com/axios/axios/issues/10802">#10802</a></strong>)</li>
<li><strong>ECONNREFUSED Error Constant:</strong> Exposed <code>ECONNREFUSED</code> as a constant on <code>AxiosError</code> so callers can match connection-refused failures without comparing string literals (closes <a href="https://redirect.github.com/axios/axios/issues/6485">#6485</a>). (<strong><a href="https://redirect.github.com/axios/axios/issues/10680">#10680</a></strong>)</li>
<li><strong>Encode Helper Export:</strong> Exported the internal <code>encode</code> helper from <code>buildURL</code> so userland param serializers can reuse the same encoding logic that axios uses internally. (<strong><a href="https://redirect.github.com/axios/axios/issues/6897">#6897</a></strong>)</li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li><strong>HTTP Adapter — Redirects &amp; Headers:</strong> Cleared stale headers when a redirect targets a no-proxy host, fixed the redirect listener chain so listeners no longer stack across hops, restored the missing <code>requestDetails</code> argument on <code>beforeRedirect</code>, preserved user-supplied <code>Host</code> headers when forwarding through a proxy, and properly URL-decoded basic auth credentials. (<strong><a href="https://redirect.github.com/axios/axios/issues/10794">#10794</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10800">#10800</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/6241">#6241</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10822">#10822</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10825">#10825</a></strong>)</li>
<li><strong>HTTP Adapter — Streams &amp; Timeouts:</strong> Preserved the partial response object on <code>AxiosError</code> when a stream is aborted after headers arrive, honoured the <code>timeout</code> option during the connect phase when redirects are disabled, and resolved an unsettled-promise hang when an aborted request was combined with compression and <code>maxRedirects: 0</code>. (<strong><a href="https://redirect.github.com/axios/axios/issues/10708">#10708</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10819">#10819</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/7149">#7149</a></strong>)</li>
<li><strong>Fetch Adapter:</strong> Enforced <code>maxBodyLength</code> / <code>maxContentLength</code> in the fetch adapter, set the <code>User-Agent</code> header to match the HTTP adapter, preserved the original abort reason instead of replacing it with a generic error, and deferred global access so importing the module no longer throws a <code>TypeError</code> in restricted environments. (<strong><a href="https://redirect.github.com/axios/axios/issues/10795">#10795</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10772">#10772</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10806">#10806</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/7260">#7260</a></strong>)</li>
<li><strong>XHR Adapter:</strong> Unsubscribed the <code>cancelToken</code> and <code>AbortSignal</code> listeners on the error, timeout, and abort code paths to prevent leaked subscriptions. (<strong><a href="https://redirect.github.com/axios/axios/issues/10787">#10787</a></strong>)</li>
<li><strong>Error Handling:</strong> Attached the parsed response to <code>AxiosError</code> when <code>JSON.parse</code> fails inside <code>dispatchRequest</code>, prevented <code>settle</code> from emitting <code>undefined</code> error codes, and tightened the <code>parseProtocol</code> regex to require a colon in the protocol separator. (<strong><a href="https://redirect.github.com/axios/axios/issues/10724">#10724</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/7276">#7276</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10729">#10729</a></strong>)</li>
<li><strong>Types &amp; Exports:</strong> Aligned the CommonJS <code>CancelToken</code> typings with the ESM build, fixed a compiler error caused by <code>RawAxiosHeaders</code>, and re-exported <code>create</code> from the package index. (<strong><a href="https://redirect.github.com/axios/axios/issues/7414">#7414</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/6389">#6389</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/6460">#6460</a></strong>)</li>
<li><strong>UTF-8 Encoding:</strong> Replaced the deprecated <code>unescape()</code> call with a modern UTF-8 encoding implementation. (<strong><a href="https://redirect.github.com/axios/axios/issues/7378">#7378</a></strong>)</li>
<li><strong>Misc Cleanup:</strong> Resolved a batch of small inconsistencies and gadget-level issues across the codebase. (<strong><a href="https://redirect.github.com/axios/axios/issues/10833">#10833</a></strong>)</li>
</ul>
<h2>🔧 Maintenance &amp; Chores</h2>
<ul>
<li><strong>Refactor — ES6 Modernisation:</strong> Modernised the <code>utils</code> module and XHR adapter to use ES6 features, and tidied the multipart boundary error message. (<strong><a href="https://redirect.github.com/axios/axios/issues/10588">#10588</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/7419">#7419</a></strong>)</li>
<li><strong>Tests:</strong> Hardened the HTTP test server lifecycle to fix flaky <code>FormData</code> EPIPE failures, fixed Win32 platform support for the pipe tests, and corrected an incorrect test assumption. (<strong><a href="https://redirect.github.com/axios/axios/issues/10820">#10820</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10791">#10791</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10796">#10796</a></strong>)</li>
<li><strong>Docs:</strong> Documented <code>paramsSerializer.encode</code> for strict RFC 3986 query encoding, updated the <code>parseReviver</code> TypeScript definitions and configuration docs for ES2023, added timeout guidance to the README's first async example, and expanded notes around the recent type changes. (<strong><a href="https://redirect.github.com/axios/axios/issues/10821">#10821</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10782">#10782</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10759">#10759</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10804">#10804</a></strong>)</li>
<li><strong>Reverted:</strong> Reverted the <code>transformRequest</code> input typing change from <a href="https://redirect.github.com/axios/axios/issues/10745">#10745</a> after follow-up review. (<strong><a href="https://redirect.github.com/axios/axios/issues/10745">#10745</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10810">#10810</a></strong>)</li>
<li><strong>Dependencies:</strong> Bumped <code>actions/setup-node</code>, the <code>github-actions</code> group, and <code>postcss</code> (in <code>/docs</code>) to their latest versions. (<strong><a href="https://redirect.github.com/axios/axios/issues/10785">#10785</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10813">#10813</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10814">#10814</a></strong>)</li>
<li><strong>Release:</strong> Updated changelog and packages, and prepared the 1.16.0 release. (<strong><a href="https://redirect.github.com/axios/axios/issues/10790">#10790</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10834">#10834</a></strong>)</li>
</ul>
<h2>🌟 New Contributors</h2>
<p>We are thrilled to welcome our new contributors. Thank you for helping improve axios:</p>
<ul>
<li><strong><a href="https://github.com/singhankit001"><code>@​singhankit001</code></a></strong> (<strong><a href="https://redirect.github.com/axios/axios/issues/10588">#10588</a></strong>)</li>
<li><strong><a href="https://github.com/cuiweixie"><code>@​cuiweixie</code></a></strong> (<strong><a href="https://redirect.github.com/axios/axios/issues/7419">#7419</a></strong>)</li>
<li><strong><a href="https://github.com/iruizsalinas"><code>@​iruizsalinas</code></a></strong> (<strong><a href="https://redirect.github.com/axios/axios/issues/10787">#10787</a></strong>)</li>
<li><strong><a href="https://github.com/MarcosNocetti"><code>@​MarcosNocetti</code></a></strong> (<strong><a href="https://redirect.github.com/axios/axios/issues/10680">#10680</a></strong>)</li>
<li><strong><a href="https://github.com/deepview-autofix"><code>@​deepview-autofix</code></a></strong> (<strong><a href="https://redirect.github.com/axios/axios/issues/10729">#10729</a></strong>)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/axios/axios/commit/df53d7dd99b202fb194217abd127ae6a630e70dc"><code>df53d7d</code></a> chore(release): prepare release 1.16.0 (<a href="https://redirect.github.com/axios/axios/issues/10834">#10834</a>)</li>
<li><a href="https://github.com/axios/axios/commit/9d92bcd32639d1eea5b89f03ae45f248d3bb058e"><code>9d92bcd</code></a> fix: gadgets and smaller issues (<a href="https://redirect.github.com/axios/axios/issues/10833">#10833</a>)</li>
<li><a href="https://github.com/axios/axios/commit/5107ee69aee527b19eabaf80000ca65752135435"><code>5107ee6</code></a> fix: prevent undefined error codes in settle (<a href="https://redirect.github.com/axios/axios/issues/7276">#7276</a>)</li>
<li><a href="https://github.com/axios/axios/commit/e57349992f230b6b13e80613eb84302560aa5ba8"><code>e573499</code></a> fix(fetch): defer global access in fetch adapter (<a href="https://redirect.github.com/axios/axios/issues/7260">#7260</a>)</li>
<li><a href="https://github.com/axios/axios/commit/ad68e1a484b50086af427f767bbd7d6e3aab7ac3"><code>ad68e1a</code></a> fix(http): honor timeout during connect without redirects (<a href="https://redirect.github.com/axios/axios/issues/10819">#10819</a>)</li>
<li><a href="https://github.com/axios/axios/commit/2a51828213128691d2e37502b5eb2cf4965a737d"><code>2a51828</code></a> fix(http): decode URL basic auth credentials (<a href="https://redirect.github.com/axios/axios/issues/10825">#10825</a>)</li>
<li><a href="https://github.com/axios/axios/commit/0e8b6bbb542131bae9940618d84d5286255d4db1"><code>0e8b6bb</code></a> fix(http): preserve user-supplied Host header when forwarding through a proxy...</li>
<li><a href="https://github.com/axios/axios/commit/79f39e1d041dca87173226d0255f90eaf252564b"><code>79f39e1</code></a> docs: document paramsSerializer.encode for strict RFC 3986 query encoding (<a href="https://redirect.github.com/axios/axios/issues/1">#1</a>...</li>
<li><a href="https://github.com/axios/axios/commit/0fe3a5fc14829535e1d517c662d448e86c33438e"><code>0fe3a5f</code></a> [Docs/Types] Update <code>parseReviver</code> TypeScript definitions for ES2023 and add ...</li>
<li><a href="https://github.com/axios/axios/commit/cd6737fd84bdb7caf2a319d3579573a49f9d238d"><code>cd6737f</code></a> chore: matches the sibling responseStream.on(aborted) handler and added tests...</li>
<li>Additional commits viewable in <a href="https://github.com/axios/axios/compare/v1.15.0...v1.16.0">compare view</a></li>
</ul>
</details>
<br />

Updates `@currents/playwright` from 1.23.0 to 1.23.3
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/currents-dev/currents-playwright-changelog/blob/main/CHANGELOG.md">@​currents/playwright's changelog</a>.</em></p>
<blockquote>
<h2><a href="https://github.com/currents-dev/currents-playwright/compare/v1.23.2...v1.23.3">1.23.3</a> (2026-05-06)</h2>
<h3>Bug Fixes</h3>
<ul>
<li>bump axios from 1.15.0 to 1.16.0 (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/809">#809</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/09430754181ff6356a68b89680d6e2ad37d3d947">0943075</a>)</li>
<li>bump follow-redirects from 1.15.11 to 1.16.0 (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/801">#801</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/429940704e90fd3313a9a8de1989c8894eccb9b9">4299407</a>)</li>
<li>bump postcss from 8.5.9 to 8.5.10 (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/800">#800</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/cfd9556b674f2d459cc55a5196294ae2506839f1">cfd9556</a>)</li>
</ul>
<h2><a href="https://github.com/currents-dev/currents-playwright/compare/v1.23.1...v1.23.2">1.23.2</a> (2026-04-17)</h2>
<h3>Bug Fixes</h3>
<ul>
<li>retain trace for quarantine in retain-on-failures trace mode (<a href="https://github.com/currents-dev/currents-playwright/commit/1f271c188296fc5f6a8e3631c8c8ec035bd60ef3">1f271c1</a>)</li>
</ul>
<h2><a href="https://github.com/currents-dev/currents-playwright/compare/v1.23.0...v1.23.1">1.23.1</a> (2026-04-16)</h2>
<h3>Bug Fixes</h3>
<ul>
<li>project dependencies missing in discovery when filtered by last-failed [CSR-4073] (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/785">#785</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/480bb590088479d56ab3489c438f006bfe768535">480bb59</a>), closes <a href="https://redirect.github.com/currents-dev/currents-playwright/issues/788">#788</a></li>
<li>run only id-only preOnlyTestFilters in discovery scanner [CSR-4094] (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/790">#790</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/288a240357745cb43748b7ac32a749eab4a11e0e">288a240</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li>See full diff in <a href="https://github.com/currents-dev/currents-playwright-changelog/commits">compare view</a></li>
</ul>
</details>
<br />


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
You can disable automated security fix PRs for this repo from the [Security Alerts page](https://github.com/RedHatInsights/patchman-ui/network/alerts).

</details>

---

## Discussion

### Comment by @codecov-commenter on May 08, 2026 at 01:01 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1605?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 76.34%. Comparing base ([`89b84d9`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/89b84d97090a083b96744a7e6a251edabb9c45c9?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`e2072d3`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/e2072d39afe1d1642b596388c17d943d142f27a9?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1605   +/-   ##
=======================================
  Coverage   76.34%   76.34%           
=======================================
  Files         103      103           
  Lines        3187     3187           
  Branches      693      692    -1     
=======================================
  Hits         2433     2433           
  Misses        675      675           
  Partials       79       79           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1605?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @swadeley - Approved on May 08, 2026 at 09:28 AM UTC

ACK

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1605*
