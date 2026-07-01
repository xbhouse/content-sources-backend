---
type: pull_request
number: 1522
title: "Build: Bump @playwright/test from 1.57.0 to 1.60.0"
state: merged
author: dependabot
created: 2026-06-08T04:36:24Z
updated: 2026-06-08T11:08:18Z
closed: 2026-06-08T11:08:16Z
merged: 2026-06-08T11:08:16Z
base_branch: main
head_branch: dependabot/npm_and_yarn/playwright/test-1.60.0
labels: ["dependencies", "javascript"]
url: https://github.com/content-services/content-sources-backend/pull/1522
---

# Pull Request #1522: Build: Bump @playwright/test from 1.57.0 to 1.60.0

**Author**: @dependabot
**Created**: June 08, 2026 at 04:36 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `javascript`
**Base**: `main` ← **Head**: `dependabot/npm_and_yarn/playwright/test-1.60.0`

## Description

Bumps [@playwright/test](https://github.com/microsoft/playwright) from 1.57.0 to 1.60.0.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/microsoft/playwright/releases">@​playwright/test's releases</a>.</em></p>
<blockquote>
<h2>v1.60.0</h2>
<h2>🌐 HAR recording on Tracing</h2>
<p><a href="https://playwright.dev/docs/api/class-tracing#tracing-start-har">tracing.startHar()</a> / <a href="https://playwright.dev/docs/api/class-tracing#tracing-stop-har">tracing.stopHar()</a> expose HAR recording as a first-class tracing API, with the same <code>content</code>, <code>mode</code> and <code>urlFilter</code> options as <code>recordHar</code>. The returned <a href="https://playwright.dev/docs/api/class-disposable">Disposable</a> makes it easy to scope a recording with <code>await using</code>:</p>
<pre lang="js"><code>await using har = await context.tracing.startHar('trace.har');
const page = await context.newPage();
await page.goto('https://playwright.dev');
// HAR is finalized when `har` goes out of scope.
</code></pre>
<h2>🪝 Drop API</h2>
<p>New <a href="https://playwright.dev/docs/api/class-locator#locator-drop">locator.drop()</a> simulates an external drag-and-drop of files or clipboard-like data onto an element. Playwright dispatches <code>dragenter</code>, <code>dragover</code>, and <code>drop</code> with a synthetic [DataTransfer] in the page context — works cross-browser and is great for testing upload zones:</p>
<pre lang="js"><code>await page.locator('#dropzone').drop({
  files: { name: 'note.txt', mimeType: 'text/plain', buffer: Buffer.from('hello') },
});
<p>await page.locator('#dropzone').drop({
data: {
'text/plain': 'hello world',
'text/uri-list': '<a href="https://example.com">https://example.com</a>',
},
});
</code></pre></p>
<h2>🎯 Aria snapshots</h2>
<ul>
<li><a href="https://playwright.dev/docs/api/class-pageassertions#page-assertions-to-match-aria-snapshot">expect(page).toMatchAriaSnapshot()</a> now works on a <a href="https://playwright.dev/docs/api/class-page">Page</a>, in addition to a <a href="https://playwright.dev/docs/api/class-locator">Locator</a> — equivalent to asserting against <code>page.locator('body')</code>.</li>
<li>New <code>boxes</code> option on <a href="https://playwright.dev/docs/api/class-locator#locator-aria-snapshot">locator.ariaSnapshot()</a> / <a href="https://playwright.dev/docs/api/class-page#page-aria-snapshot">page.ariaSnapshot()</a> appends each element's bounding box as <code>[box=x,y,width,height]</code>, useful for AI consumption.</li>
</ul>
<h2>🛑 test.abort()</h2>
<p>New <a href="https://playwright.dev/docs/api/class-test#test-abort">test.abort()</a> aborts the currently running test from a fixture, hook, or route handler with an optional message. Use it when you have detected an unrecoverable misuse and want to fail the test right away:</p>
<pre lang="js"><code>test('does not publish to the shared page', async ({ page }) =&gt; {
  await page.route('**/publish', route =&gt; {
    test.abort('Tests must not publish to the shared page. Use the `clone` option.');
    return route.abort();
  });
  // ...
});
</code></pre>
<h2>New APIs</h2>
<h3>Browser, Context and Page</h3>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/microsoft/playwright/commit/87bb9ddbd78f329df18c2b24847bc9409240cd07"><code>87bb9dd</code></a> cherry-pick(<a href="https://redirect.github.com/microsoft/playwright/issues/40747">#40747</a>): fix(yauzl): vendor yauzl with destroy-lifecycle fix</li>
<li><a href="https://github.com/microsoft/playwright/commit/9a9c51cb7d1b39fab51ca288e59f8ca38fd19910"><code>9a9c51c</code></a> cherry-pick(<a href="https://redirect.github.com/microsoft/playwright/issues/40733">#40733</a>): chore(electron): revert <a href="https://redirect.github.com/microsoft/playwright/issues/40184">#40184</a> (move Electron API to a s...</li>
<li><a href="https://github.com/microsoft/playwright/commit/4b3b628663031bcaaeca907e337892263524634d"><code>4b3b628</code></a> cherry-pick(<a href="https://redirect.github.com/microsoft/playwright/issues/40736">#40736</a>): Revert &quot;feat(electron): add timeout option to electronAp...</li>
<li><a href="https://github.com/microsoft/playwright/commit/f869f96bbe6607cc3b88b4ca96fd82f17b301b50"><code>f869f96</code></a> chore: bump version to v1.60.0 (<a href="https://redirect.github.com/microsoft/playwright/issues/40714">#40714</a>)</li>
<li><a href="https://github.com/microsoft/playwright/commit/7eb6918afadfb0dd5c7e94ca9ffbddd84d8fbb39"><code>7eb6918</code></a> cherry-pick(<a href="https://redirect.github.com/microsoft/playwright/issues/40710">#40710</a>): docs: release notes v1.60</li>
<li><a href="https://github.com/microsoft/playwright/commit/118d2aa6076d82840decca15d96b48611b08e392"><code>118d2aa</code></a> cherry-pick(<a href="https://redirect.github.com/microsoft/playwright/issues/40693">#40693</a>): chore(python): formdata path type</li>
<li><a href="https://github.com/microsoft/playwright/commit/54012f5dcc586da2e5d6cccd75f13ca367b94579"><code>54012f5</code></a> chore(deps): bump ip-address and express-rate-limit (<a href="https://redirect.github.com/microsoft/playwright/issues/40680">#40680</a>)</li>
<li><a href="https://github.com/microsoft/playwright/commit/9fa531da5677a3807d6e1dccd22c5137339a44f7"><code>9fa531d</code></a> fix(screencast): unblock frame ack when an async client disconnects (<a href="https://redirect.github.com/microsoft/playwright/issues/40674">#40674</a>)</li>
<li><a href="https://github.com/microsoft/playwright/commit/3649db560ff943e724185784d34f7db131a11961"><code>3649db5</code></a> chore(mcp): bump default extension protocol to v2 (<a href="https://redirect.github.com/microsoft/playwright/issues/40678">#40678</a>)</li>
<li><a href="https://github.com/microsoft/playwright/commit/bb6c00957f47ba04caad7fca75d426309a2d32d4"><code>bb6c009</code></a> chore(extension): mark 0.2.1 (<a href="https://redirect.github.com/microsoft/playwright/issues/40679">#40679</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/microsoft/playwright/compare/v1.57.0...v1.60.0">compare view</a></li>
</ul>
</details>
<br />


---

## Discussion

### Comment by @swadeley on June 08, 2026 at 08:12 AM UTC

build fails

---

## Reviews

### Review by @swadeley - Approved on June 08, 2026 at 11:08 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1522*
