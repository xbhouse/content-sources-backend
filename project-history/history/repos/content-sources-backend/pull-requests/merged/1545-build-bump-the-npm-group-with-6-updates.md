---
type: pull_request
number: 1545
title: "Build: Bump the npm group with 6 updates"
state: merged
author: dependabot
created: 2026-06-22T04:36:29Z
updated: 2026-06-22T07:27:58Z
closed: 2026-06-22T07:27:56Z
merged: 2026-06-22T07:27:56Z
base_branch: main
head_branch: dependabot/npm_and_yarn/npm-a7ef11845a
labels: ["dependencies", "javascript"]
url: https://github.com/content-services/content-sources-backend/pull/1545
---

# Pull Request #1545: Build: Bump the npm group with 6 updates

**Author**: @dependabot
**Created**: June 22, 2026 at 04:36 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `javascript`
**Base**: `main` ← **Head**: `dependabot/npm_and_yarn/npm-a7ef11845a`

## Description

Bumps the npm group with 6 updates:

| Package | From | To |
| --- | --- | --- |
| [@currents/playwright](https://github.com/currents-dev/currents-playwright-changelog) | `2.3.0` | `2.3.1` |
| [@playwright/test](https://github.com/microsoft/playwright) | `1.60.0` | `1.61.0` |
| [@types/node](https://github.com/DefinitelyTyped/DefinitelyTyped/tree/HEAD/types/node) | `25.9.2` | `25.9.3` |
| [@typescript-eslint/eslint-plugin](https://github.com/typescript-eslint/typescript-eslint/tree/HEAD/packages/eslint-plugin) | `8.61.0` | `8.61.1` |
| [@typescript-eslint/parser](https://github.com/typescript-eslint/typescript-eslint/tree/HEAD/packages/parser) | `8.61.0` | `8.61.1` |
| [eslint](https://github.com/eslint/eslint) | `10.4.1` | `10.5.0` |

Updates `@currents/playwright` from 2.3.0 to 2.3.1
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/currents-dev/currents-playwright-changelog/blob/main/CHANGELOG.md">@​currents/playwright's changelog</a>.</em></p>
<blockquote>
<h2><a href="https://github.com/currents-dev/currents-playwright/compare/v2.3.0...v2.3.1">2.3.1</a> (2026-06-16)</h2>
<h3>Bug Fixes</h3>
<ul>
<li><code>@​currents/playwright</code> compatibility with Playwright 1.61.0 [ENG-680] (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/881">#881</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/87d233b839bd357c425be559cdfbd43350dbc2b7">87d233b</a>)</li>
<li>bump <code>@​babel/code-frame</code> from 7.29.0 to 7.29.7 (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/875">#875</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/75767776743a5be817cfaf80fe9fac1f11f8011c">7576777</a>)</li>
<li>ensure pwc-p run rejects the -G flag (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/882">#882</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/172e9744557e26431d51df7507f78486c8bd0f59">172e974</a>)</li>
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

Updates `@playwright/test` from 1.60.0 to 1.61.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/microsoft/playwright/releases">@​playwright/test's releases</a>.</em></p>
<blockquote>
<h2>v1.61.0</h2>
<h2>🔑 WebAuthn passkeys</h2>
<p>New <a href="https://playwright.dev/docs/api/class-credentials">Credentials</a> virtual authenticator, available via <a href="https://playwright.dev/docs/api/class-browsercontext#browser-context-credentials">browserContext.credentials</a>, lets tests register passkeys and answer <code>navigator.credentials.create()</code> / <code>navigator.credentials.get()</code> ceremonies in the page — no real hardware key required, works in all browsers:</p>
<pre lang="js"><code>const context = await browser.newContext();
<p>// Seed a passkey your backend provisioned for a test user.<br />
await context.credentials.create('example.com', {<br />
id: credentialId,<br />
userHandle,<br />
privateKey,<br />
publicKey,<br />
});<br />
await context.credentials.install();</p>
<p>const page = await context.newPage();<br />
await page.goto('<a href="https://example.com/login">https://example.com/login</a>');<br />
// The page's navigator.credentials.get() is answered with the seeded passkey.<br />
</code></pre></p>
<p>You can also let the app register a passkey once in a setup test, read it back with <a href="https://playwright.dev/docs/api/class-credentials#credentials-get">credentials.get()</a>, and seed it into later tests — see <a href="https://playwright.dev/docs/api/class-credentials">Credentials</a> for details.</p>
<h2>🗃️ Web Storage</h2>
<p>New <a href="https://playwright.dev/docs/api/class-webstorage">WebStorage</a> API, available via <a href="https://playwright.dev/docs/api/class-page#page-local-storage">page.localStorage</a> and <a href="https://playwright.dev/docs/api/class-page#page-session-storage">page.sessionStorage</a>, reads and writes the page's storage for the current origin:</p>
<pre lang="js"><code>await page.localStorage.setItem('token', 'abc');
const token = await page.localStorage.getItem('token');
const items = await page.sessionStorage.items();
</code></pre>
<h2>New APIs</h2>
<h3>Network</h3>
<ul>
<li><a href="https://playwright.dev/docs/api/class-apiresponse#api-response-security-details">apiResponse.securityDetails()</a> and <a href="https://playwright.dev/docs/api/class-apiresponse#api-response-server-addr">apiResponse.serverAddr()</a> mirror the browser-side <a href="https://playwright.dev/docs/api/class-response#response-security-details">response.securityDetails()</a> and <a href="https://playwright.dev/docs/api/class-response#response-server-addr">response.serverAddr()</a>.</li>
</ul>
<h3>Browser and Screencast</h3>
<ul>
<li>New option <code>artifactsDir</code> in <a href="https://playwright.dev/docs/api/class-browsertype#browser-type-connect-over-cdp">browserType.connectOverCDP()</a> controls where artifacts such as traces and downloads are stored when attached to an existing browser.</li>
<li>New option <code>cursor</code> in <a href="https://playwright.dev/docs/api/class-screencast#screencast-show-actions">screencast.showActions()</a> controls the cursor decoration rendered for pointer actions.</li>
<li>The <code>onFrame</code> callback in <a href="https://playwright.dev/docs/api/class-screencast#screencast-start">screencast.start()</a> now receives a <code>timestamp</code> of when the frame was presented by the browser.</li>
</ul>
<h3>Test runner</h3>
<ul>
<li>The <a href="https://playwright.dev/docs/api/class-testoptions#test-options-video">testOptions.video</a> option now supports the same set of modes as <code>trace</code>: new <code>'on-all-retries'</code>, <code>'retain-on-first-failure'</code> and <code>'retain-on-failure-and-retries'</code> values. See the <a href="https://playwright.dev/docs/test-use-options#video-modes">video modes table</a> for which runs are recorded and kept in each mode.</li>
<li>Supported <code>expect.soft.poll(...)</code>.</li>
<li>New <a href="https://playwright.dev/docs/api/class-fullconfig#full-config-argv">fullConfig.argv</a> — a snapshot of <code>process.argv</code> from the runner process, handy for reading custom arguments passed after the <code>--</code> separator.</li>
<li>New <a href="https://playwright.dev/docs/api/class-fullconfig#full-config-fail-on-flaky-tests">fullConfig.failOnFlakyTests</a> mirrors the config option, so reporters can explain why a flaky run failed.</li>
<li><a href="https://playwright.dev/docs/api/class-testinfo#test-info-errors">testInfo.errors</a> now lists each sub-error of an <code>AggregateError</code> as a separate entry.</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/microsoft/playwright/commit/1cc5a90cfa3eaa430b1a991963100f95126caa47"><code>1cc5a90</code></a> cherry-pick(<a href="https://redirect.github.com/microsoft/playwright/issues/41295">#41295</a>): chore: PLAYWRIGHT_TRACING_NO_WEBSOCKET_FRAMES and PLAYWR...</li>
<li><a href="https://github.com/microsoft/playwright/commit/a6772bdede34028cbbd417a3b3d778801899e870"><code>a6772bd</code></a> cherry-pick(<a href="https://redirect.github.com/microsoft/playwright/issues/41280">#41280</a>): Revert &quot;fix(trace-viewer): add keyboard navigation to `N...</li>
<li><a href="https://github.com/microsoft/playwright/commit/8133dcf97d52818d36022ed37797a616ff6cb934"><code>8133dcf</code></a> cherry-pick(<a href="https://redirect.github.com/microsoft/playwright/issues/41283">#41283</a>): docs: add Ubuntu 26.04 and Node.js 26.x to system requir...</li>
<li><a href="https://github.com/microsoft/playwright/commit/812432e070afec9e44d22e95915f975965b7d5b7"><code>812432e</code></a> chore: mark v1.61.0 (<a href="https://redirect.github.com/microsoft/playwright/issues/41277">#41277</a>)</li>
<li><a href="https://github.com/microsoft/playwright/commit/ac05145c8d9eb1303c8f3bfd4d860b6d1ca261ae"><code>ac05145</code></a> fix(fetch): report serverAddr and securityDetails for reused sockets (<a href="https://redirect.github.com/microsoft/playwright/issues/41267">#41267</a>)</li>
<li><a href="https://github.com/microsoft/playwright/commit/056efc9f5c0a870d0944e53a835d6283a77f200f"><code>056efc9</code></a> fix(trace-viewer): add keyboard navigation to <code>NetworkFilters</code> component (<a href="https://redirect.github.com/microsoft/playwright/issues/41">#41</a>...</li>
<li><a href="https://github.com/microsoft/playwright/commit/41f7b9a0db0d1ada12ff0d9244393eea8f81b796"><code>41f7b9a</code></a> chore: fixes uncovered by the .NET 1.61 roll (<a href="https://redirect.github.com/microsoft/playwright/issues/41266">#41266</a>)</li>
<li><a href="https://github.com/microsoft/playwright/commit/ba507783ae48724a1882f6423d8e8ec208bf366a"><code>ba50778</code></a> fix(mcp): assign caps as array for legacy --vision flag (<a href="https://redirect.github.com/microsoft/playwright/issues/41253">#41253</a>)</li>
<li><a href="https://github.com/microsoft/playwright/commit/b8ee5ae27fd068e3744852209dfcb5c1a142909f"><code>b8ee5ae</code></a> docs: release notes for v1.61 (<a href="https://redirect.github.com/microsoft/playwright/issues/41261">#41261</a>)</li>
<li><a href="https://github.com/microsoft/playwright/commit/49c1f694c9bc06c9d1f6966afe8b6dfd4f388b3e"><code>49c1f69</code></a> fix(trace viewer): load trace from a local file (<a href="https://redirect.github.com/microsoft/playwright/issues/41263">#41263</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/microsoft/playwright/compare/v1.60.0...v1.61.0">compare view</a></li>
</ul>
</details>
<br />

Updates `@types/node` from 25.9.2 to 25.9.3
<details>
<summary>Commits</summary>
<ul>
<li>See full diff in <a href="https://github.com/DefinitelyTyped/DefinitelyTyped/commits/HEAD/types/node">compare view</a></li>
</ul>
</details>
<br />

Updates `@typescript-eslint/eslint-plugin` from 8.61.0 to 8.61.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/typescript-eslint/typescript-eslint/releases">@​typescript-eslint/eslint-plugin's releases</a>.</em></p>
<blockquote>
<h2>v8.61.1</h2>
<h2>8.61.1 (2026-06-15)</h2>
<h3>🩹 Fixes</h3>
<ul>
<li><strong>eslint-plugin:</strong> [consistent-indexed-object-style] do not remove comments when fixing (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12396">#12396</a>, <a href="https://redirect.github.com/typescript-eslint/typescript-eslint/issues/10577">#10577</a>)</li>
<li><strong>eslint-plugin:</strong> [no-unnecessary-type-assertion] avoid false positive for template literal expressions (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12281">#12281</a>)</li>
<li><strong>eslint-plugin:</strong> [no-unnecessary-type-assertion] wrap object literal in parens when removing TSTypeAssertion in arrow body (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12394">#12394</a>, <a href="https://redirect.github.com/typescript-eslint/typescript-eslint/issues/12393">#12393</a>)</li>
<li><strong>eslint-plugin:</strong> [no-unnecessary-boolean-literal-compare] fix precedence bug in autofix (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12413">#12413</a>)</li>
<li><strong>eslint-plugin:</strong> [no-unnecessary-template-expression] respect ECMAScript line terminators (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12388">#12388</a>)</li>
</ul>
<h3>❤️ Thank You</h3>
<ul>
<li>Anas <a href="https://github.com/anasm266"><code>@​anasm266</code></a></li>
<li>Deftera <a href="https://github.com/Deftera186"><code>@​Deftera186</code></a></li>
<li>Kirk Waiblinger <a href="https://github.com/kirkwaiblinger"><code>@​kirkwaiblinger</code></a></li>
<li>lumir</li>
<li>Sarath Francis <a href="https://github.com/sarathfrancis90"><code>@​sarathfrancis90</code></a></li>
</ul>
<p>See <a href="https://github.com/typescript-eslint/typescript-eslint/releases/tag/v8.61.1">GitHub Releases</a> for more information.</p>
<p>You can read about our <a href="https://typescript-eslint.io/users/versioning">versioning strategy</a> and <a href="https://typescript-eslint.io/users/releases">releases</a> on our website.</p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/typescript-eslint/typescript-eslint/blob/main/packages/eslint-plugin/CHANGELOG.md">@​typescript-eslint/eslint-plugin's changelog</a>.</em></p>
<blockquote>
<h2>8.61.1 (2026-06-15)</h2>
<h3>🩹 Fixes</h3>
<ul>
<li><strong>eslint-plugin:</strong> [no-unnecessary-template-expression] respect ECMAScript line terminators (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12388">#12388</a>)</li>
<li><strong>eslint-plugin:</strong> [no-unnecessary-boolean-literal-compare] fix precedence bug in autofix (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12413">#12413</a>)</li>
<li><strong>eslint-plugin:</strong> [no-unnecessary-type-assertion] wrap object literal in parens when removing TSTypeAssertion in arrow body (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12394">#12394</a>, <a href="https://redirect.github.com/typescript-eslint/typescript-eslint/issues/12393">#12393</a>)</li>
<li><strong>eslint-plugin:</strong> [no-unnecessary-type-assertion] avoid false positive for template literal expressions (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12281">#12281</a>)</li>
<li><strong>eslint-plugin:</strong> [consistent-indexed-object-style] do not remove comments when fixing (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12396">#12396</a>, <a href="https://redirect.github.com/typescript-eslint/typescript-eslint/issues/10577">#10577</a>)</li>
</ul>
<h3>❤️ Thank You</h3>
<ul>
<li>Anas <a href="https://github.com/anasm266"><code>@​anasm266</code></a></li>
<li>Deftera <a href="https://github.com/Deftera186"><code>@​Deftera186</code></a></li>
<li>Kirk Waiblinger <a href="https://github.com/kirkwaiblinger"><code>@​kirkwaiblinger</code></a></li>
<li>lumir</li>
<li>Sarath Francis <a href="https://github.com/sarathfrancis90"><code>@​sarathfrancis90</code></a></li>
</ul>
<p>See <a href="https://github.com/typescript-eslint/typescript-eslint/releases/tag/v8.61.1">GitHub Releases</a> for more information.</p>
<p>You can read about our <a href="https://typescript-eslint.io/users/versioning">versioning strategy</a> and <a href="https://typescript-eslint.io/users/releases">releases</a> on our website.</p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/aaad7187b529c4b6ff8088ffd7e948c69c2763b6"><code>aaad718</code></a> chore(release): publish 8.61.1</li>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/0cc8f3559591221400bed1e8bf8debce5edca4c4"><code>0cc8f35</code></a> fix(eslint-plugin): [no-unnecessary-template-expression] respect ECMAScript l...</li>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/6f269e21e3d65388886ee17c5e568d4c7e55bb24"><code>6f269e2</code></a> fix(eslint-plugin): [no-unnecessary-boolean-literal-compare] fix precedence b...</li>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/1b5d5430831e0baacca248027bf840260d4e597b"><code>1b5d543</code></a> fix(eslint-plugin): [no-unnecessary-type-assertion] wrap object literal in pa...</li>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/565e6668b2c69b02a74d8ca27c41633f20f4a699"><code>565e666</code></a> fix(eslint-plugin): [no-unnecessary-type-assertion] avoid false positive for ...</li>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/204eabc2949c59802b4d7057d9e16874fa19db6f"><code>204eabc</code></a> fix(eslint-plugin): [consistent-indexed-object-style] do not remove comments ...</li>
<li>See full diff in <a href="https://github.com/typescript-eslint/typescript-eslint/commits/v8.61.1/packages/eslint-plugin">compare view</a></li>
</ul>
</details>
<br />

Updates `@typescript-eslint/parser` from 8.61.0 to 8.61.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/typescript-eslint/typescript-eslint/releases">@​typescript-eslint/parser's releases</a>.</em></p>
<blockquote>
<h2>v8.61.1</h2>
<h2>8.61.1 (2026-06-15)</h2>
<h3>🩹 Fixes</h3>
<ul>
<li><strong>eslint-plugin:</strong> [consistent-indexed-object-style] do not remove comments when fixing (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12396">#12396</a>, <a href="https://redirect.github.com/typescript-eslint/typescript-eslint/issues/10577">#10577</a>)</li>
<li><strong>eslint-plugin:</strong> [no-unnecessary-type-assertion] avoid false positive for template literal expressions (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12281">#12281</a>)</li>
<li><strong>eslint-plugin:</strong> [no-unnecessary-type-assertion] wrap object literal in parens when removing TSTypeAssertion in arrow body (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12394">#12394</a>, <a href="https://redirect.github.com/typescript-eslint/typescript-eslint/issues/12393">#12393</a>)</li>
<li><strong>eslint-plugin:</strong> [no-unnecessary-boolean-literal-compare] fix precedence bug in autofix (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12413">#12413</a>)</li>
<li><strong>eslint-plugin:</strong> [no-unnecessary-template-expression] respect ECMAScript line terminators (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12388">#12388</a>)</li>
</ul>
<h3>❤️ Thank You</h3>
<ul>
<li>Anas <a href="https://github.com/anasm266"><code>@​anasm266</code></a></li>
<li>Deftera <a href="https://github.com/Deftera186"><code>@​Deftera186</code></a></li>
<li>Kirk Waiblinger <a href="https://github.com/kirkwaiblinger"><code>@​kirkwaiblinger</code></a></li>
<li>lumir</li>
<li>Sarath Francis <a href="https://github.com/sarathfrancis90"><code>@​sarathfrancis90</code></a></li>
</ul>
<p>See <a href="https://github.com/typescript-eslint/typescript-eslint/releases/tag/v8.61.1">GitHub Releases</a> for more information.</p>
<p>You can read about our <a href="https://typescript-eslint.io/users/versioning">versioning strategy</a> and <a href="https://typescript-eslint.io/users/releases">releases</a> on our website.</p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/typescript-eslint/typescript-eslint/blob/main/packages/parser/CHANGELOG.md">@​typescript-eslint/parser's changelog</a>.</em></p>
<blockquote>
<h2>8.61.1 (2026-06-15)</h2>
<p>This was a version bump only for parser to align it with other projects, there were no code changes.</p>
<p>See <a href="https://github.com/typescript-eslint/typescript-eslint/releases/tag/v8.61.1">GitHub Releases</a> for more information.</p>
<p>You can read about our <a href="https://typescript-eslint.io/users/versioning">versioning strategy</a> and <a href="https://typescript-eslint.io/users/releases">releases</a> on our website.</p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/aaad7187b529c4b6ff8088ffd7e948c69c2763b6"><code>aaad718</code></a> chore(release): publish 8.61.1</li>
<li>See full diff in <a href="https://github.com/typescript-eslint/typescript-eslint/commits/v8.61.1/packages/parser">compare view</a></li>
</ul>
</details>
<br />

Updates `eslint` from 10.4.1 to 10.5.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/eslint/eslint/releases">eslint's releases</a>.</em></p>
<blockquote>
<h2>v10.5.0</h2>
<h2>Features</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/5ca8c5278edea1fd84d3ba83d8ea3f52fb3831ad"><code>5ca8c52</code></a> feat: correct stack tracking in max-nested-callbacks (<a href="https://redirect.github.com/eslint/eslint/issues/20973">#20973</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/b5657837604fa5e8cf1278074782025cadd34b6c"><code>b565783</code></a> feat: report no-with violations at the with keyword (<a href="https://redirect.github.com/eslint/eslint/issues/20971">#20971</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/2ce032fbc72a1a80c024c084a4f382fb6dece684"><code>2ce032f</code></a> feat: report max-lines-per-function violations at function head (<a href="https://redirect.github.com/eslint/eslint/issues/20966">#20966</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/732cb3e09d5b8b809b5f461d118a5d9fdcd6427f"><code>732cb3e</code></a> feat: report max-nested-callbacks violations at function head (<a href="https://redirect.github.com/eslint/eslint/issues/20967">#20967</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/f9c138a0ba7d8e37aed39aef4a3ff1cae8c669f7"><code>f9c138a</code></a> feat: report max-depth violations on keywords (<a href="https://redirect.github.com/eslint/eslint/issues/20943">#20943</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/bdb496cc0d54b6d0a023aef9abd5f040ccff2101"><code>bdb496c</code></a> feat: correct max-depth handling for else-if chains (<a href="https://redirect.github.com/eslint/eslint/issues/20944">#20944</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/c29687354a7f96093f57f7d73eecb866ad5e2953"><code>c296873</code></a> feat: update error loc in <code>max-statements</code> to function header (<a href="https://redirect.github.com/eslint/eslint/issues/20907">#20907</a>) (Taejin Kim)</li>
</ul>
<h2>Documentation</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/8ae1b5b856dc031cd6c701d89a4df7da4772cd56"><code>8ae1b5b</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/ca7eb90127dcad917188bb1342623f02a272e781"><code>ca7eb90</code></a> docs: update Node.js prerequisites to include ICU support (<a href="https://redirect.github.com/eslint/eslint/issues/20962">#20962</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/f99b47a6799be25321552402a49303bb06a43fe4"><code>f99b47a</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/acf03d4eed31d259c7dc62af5b9640629784f7cc"><code>acf03d4</code></a> docs: clarify precedence of parserOptions over languageOptions (<a href="https://redirect.github.com/eslint/eslint/issues/20926">#20926</a>) (sethamus)</li>
</ul>
<h2>Chores</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/b18bf58c5ac748415ffffdff2d96980fbd6a57e8"><code>b18bf58</code></a> chore: update ecosystem plugins (<a href="https://redirect.github.com/eslint/eslint/issues/20959">#20959</a>) (ESLint Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/c2d1444df77cb42e5a0b89ab70496879d180a54d"><code>c2d1444</code></a> refactor: replace areAllSegmentsUnreachable with !isAnySegmentReachable (<a href="https://redirect.github.com/eslint/eslint/issues/20951">#20951</a>) (Taejin Kim)</li>
<li><a href="https://github.com/eslint/eslint/commit/243b8c56014bbbe63771185b0731d8dd4d1316e9"><code>243b8c5</code></a> chore: enhance config-rule to support oneOf, anyOf, and nested schemas (<a href="https://redirect.github.com/eslint/eslint/issues/20788">#20788</a>) (kuldeep kumar)</li>
<li><a href="https://github.com/eslint/eslint/commit/217b2a91f46137c5ffd693965e71306c4c15ea6b"><code>217b2a9</code></a> test: add unit tests for ParserService (<a href="https://redirect.github.com/eslint/eslint/issues/20949">#20949</a>) (Taejin Kim)</li>
<li><a href="https://github.com/eslint/eslint/commit/72003e781d76bd4ee0d98a6601730d0b829070f9"><code>72003e7</code></a> test: add location information to error messages in <code>max-statements</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20945">#20945</a>) (lumir)</li>
<li><a href="https://github.com/eslint/eslint/commit/7797c266977b0bc4971aa79721813d480de72cd1"><code>7797c26</code></a> refactor: deduplicate isAnySegmentReachable across rules (<a href="https://redirect.github.com/eslint/eslint/issues/20890">#20890</a>) (Taejin Kim)</li>
<li><a href="https://github.com/eslint/eslint/commit/67c46fa6e4f34e88cc6bc82f8a0dcc917c65d257"><code>67c46fa</code></a> chore: update ecosystem plugins (<a href="https://redirect.github.com/eslint/eslint/issues/20938">#20938</a>) (ESLint Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/95d8c7a99f991abd8ab618d0ee2cbd4f58effc29"><code>95d8c7a</code></a> chore: update dependency <code>@​eslint/json</code> to v2 (<a href="https://redirect.github.com/eslint/eslint/issues/20934">#20934</a>) (renovate[bot])</li>
<li><a href="https://github.com/eslint/eslint/commit/cf9e496205142cd4971b9f98aed85866d1010b9c"><code>cf9e496</code></a> chore: update <code>@​arethetypeswrong/cli</code> to 0.18.3 (<a href="https://redirect.github.com/eslint/eslint/issues/20933">#20933</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/fb6d3960cacc51fc12383fa5ded2382adbf90c1c"><code>fb6d396</code></a> test: run type tests with TypeScript 7 (<a href="https://redirect.github.com/eslint/eslint/issues/20868">#20868</a>) (sethamus)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/de3b672a267e32607db04176ce4775664acb3145"><code>de3b672</code></a> 10.5.0</li>
<li><a href="https://github.com/eslint/eslint/commit/362a5185134290db696d39f97c9da609ded54040"><code>362a518</code></a> Build: changelog update for 10.5.0</li>
<li><a href="https://github.com/eslint/eslint/commit/5ca8c5278edea1fd84d3ba83d8ea3f52fb3831ad"><code>5ca8c52</code></a> feat: correct stack tracking in max-nested-callbacks (<a href="https://redirect.github.com/eslint/eslint/issues/20973">#20973</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/b5657837604fa5e8cf1278074782025cadd34b6c"><code>b565783</code></a> feat: report no-with violations at the with keyword (<a href="https://redirect.github.com/eslint/eslint/issues/20971">#20971</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/2ce032fbc72a1a80c024c084a4f382fb6dece684"><code>2ce032f</code></a> feat: report max-lines-per-function violations at function head (<a href="https://redirect.github.com/eslint/eslint/issues/20966">#20966</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/732cb3e09d5b8b809b5f461d118a5d9fdcd6427f"><code>732cb3e</code></a> feat: report max-nested-callbacks violations at function head (<a href="https://redirect.github.com/eslint/eslint/issues/20967">#20967</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/f9c138a0ba7d8e37aed39aef4a3ff1cae8c669f7"><code>f9c138a</code></a> feat: report max-depth violations on keywords (<a href="https://redirect.github.com/eslint/eslint/issues/20943">#20943</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/8ae1b5b856dc031cd6c701d89a4df7da4772cd56"><code>8ae1b5b</code></a> docs: Update README</li>
<li><a href="https://github.com/eslint/eslint/commit/ca7eb90127dcad917188bb1342623f02a272e781"><code>ca7eb90</code></a> docs: update Node.js prerequisites to include ICU support (<a href="https://redirect.github.com/eslint/eslint/issues/20962">#20962</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/b18bf58c5ac748415ffffdff2d96980fbd6a57e8"><code>b18bf58</code></a> chore: update ecosystem plugins (<a href="https://redirect.github.com/eslint/eslint/issues/20959">#20959</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/eslint/eslint/compare/v10.4.1...v10.5.0">compare view</a></li>
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
- `@dependabot ignore <dependency name> major version` will close this group update PR and stop Dependabot creating any more for the specific dependency's major version (unless you unignore this specific dependency's major version or upgrade to it yourself)
- `@dependabot ignore <dependency name> minor version` will close this group update PR and stop Dependabot creating any more for the specific dependency's minor version (unless you unignore this specific dependency's minor version or upgrade to it yourself)
- `@dependabot ignore <dependency name>` will close this group update PR and stop Dependabot creating any more for the specific dependency (unless you unignore this specific dependency or upgrade to it yourself)
- `@dependabot unignore <dependency name>` will remove all of the ignore conditions of the specified dependency
- `@dependabot unignore <dependency name> <ignore condition>` will remove the ignore condition of the specified dependency and ignore conditions


</details>

---

## Reviews

### Review by @swadeley - Approved on June 22, 2026 at 07:27 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1545*
