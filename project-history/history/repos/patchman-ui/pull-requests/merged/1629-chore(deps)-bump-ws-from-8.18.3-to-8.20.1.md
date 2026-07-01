---
type: pull_request
number: 1629
title: "chore(deps): bump ws from 8.18.3 to 8.20.1"
state: merged
author: dependabot
created: 2026-05-20T11:30:17Z
updated: 2026-05-21T07:23:13Z
closed: 2026-05-21T07:23:04Z
merged: 2026-05-21T07:23:04Z
base_branch: master
head_branch: dependabot/npm_and_yarn/ws-8.20.1
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1629
---

# Pull Request #1629: chore(deps): bump ws from 8.18.3 to 8.20.1

**Author**: @dependabot
**Created**: May 20, 2026 at 11:30 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/ws-8.20.1`

## Description

Bumps [ws](https://github.com/websockets/ws) from 8.18.3 to 8.20.1.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/websockets/ws/releases">ws's releases</a>.</em></p>
<blockquote>
<h2>8.20.1</h2>
<h1>Bug fixes</h1>
<ul>
<li>Fixed an uninitialized memory disclosure issue in <code>websocket.close()</code>
(c0327ec1).</li>
</ul>
<p>Providing a <code>TypedArray</code> (e.g. <code>Float32Array</code>) as the <code>reason</code> argument for
<code>websocket.close()</code>, rather than the supported string or <code>Buffer</code> types, caused
uninitialized memory to be disclosed to the remote peer.</p>
<pre lang="js"><code>import { deepStrictEqual } from 'node:assert';
import { WebSocket, WebSocketServer } from 'ws';
<p>const wss = new WebSocketServer(
{ port: 0, skipUTF8Validation: true },
function () {
const { port } = wss.address();
const ws = new WebSocket(<code>ws://localhost:${port}</code>, {
skipUTF8Validation: true
});</p>
<pre><code>ws.on('close', function (code, reason) {
  deepStrictEqual(reason, Buffer.alloc(80));
});
</code></pre>
<p>}
);</p>
<p>wss.on('connection', function (ws) {
ws.close(1000, new Float32Array(20));
});
</code></pre></p>
<p>The issue was privately reported by <a href="https://github.com/ChALkeR">Nikita Skovoroda</a>.</p>
<h2>8.20.0</h2>
<h1>Features</h1>
<ul>
<li>Added exports for the <code>PerMessageDeflate</code> class and utilities for the
<code>Sec-WebSocket-Extensions</code> and <code>Sec-WebSocket-Protocol</code> headers (d3503c1f).</li>
</ul>
<h2>8.19.0</h2>
<h1>Features</h1>
<ul>
<li>Added the <code>closeTimeout</code> option (<a href="https://redirect.github.com/websockets/ws/issues/2308">#2308</a>).</li>
</ul>
<h1>Bug fixes</h1>
<ul>
<li>Handled a forthcoming breaking change in Node.js core (19984854).</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/websockets/ws/commit/5d9b316230ea931532a6671cc450f18c11edd02f"><code>5d9b316</code></a> [dist] 8.20.1</li>
<li><a href="https://github.com/websockets/ws/commit/c0327ec15a54d701eb6ccefaa8bef328cfc03086"><code>c0327ec</code></a> [security] Fix uninitialized memory disclosure in <code>websocket.close()</code></li>
<li><a href="https://github.com/websockets/ws/commit/ce2a3d62437995a47e6056d485a33d21b6a8f867"><code>ce2a3d6</code></a> [ci] Test on node 26</li>
<li><a href="https://github.com/websockets/ws/commit/58e45b872bb0f35a3edd553c27e105300a4f5bd0"><code>58e45b8</code></a> [ci] Do not test on node 25</li>
<li><a href="https://github.com/websockets/ws/commit/5f26c245231a4b018479a9269e8c3da4773fe42f"><code>5f26c24</code></a> [ci] Run the lint step on node 24</li>
<li><a href="https://github.com/websockets/ws/commit/843925544e2f4cffe445e0179947f56d6c5b608f"><code>8439255</code></a> [dist] 8.20.0</li>
<li><a href="https://github.com/websockets/ws/commit/d3503c1fd36a310985108f62b343bae18346ab67"><code>d3503c1</code></a> [minor] Export the <code>PerMessageDeflate</code> class and header utils</li>
<li><a href="https://github.com/websockets/ws/commit/3ee5349a0b1580f6e1f347b59ec3371011bd8481"><code>3ee5349</code></a> [api] Convert the <code>isServer</code> and <code>maxPayload</code> parameters to options</li>
<li><a href="https://github.com/websockets/ws/commit/91707b470ebd803aaa3fd1e896217740f39267d4"><code>91707b4</code></a> [doc] Add missing space</li>
<li><a href="https://github.com/websockets/ws/commit/8b553192268810a83253e2a4a39ac16768e75bb3"><code>8b55319</code></a> [pkg] Update eslint to version 10.0.1</li>
<li>Additional commits viewable in <a href="https://github.com/websockets/ws/compare/8.18.3...8.20.1">compare view</a></li>
</ul>
</details>
<br />


---

## Discussion

### Comment by @codecov-commenter on May 20, 2026 at 11:33 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1629?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 77.58%. Comparing base ([`ee372fc`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/ee372fca1965107dc1b77ce9ba34fa860897ca63?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`d226fae`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/d226fae167a8aecb55c3339a4591f49495be4c79?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1629   +/-   ##
=======================================
  Coverage   77.58%   77.58%           
=======================================
  Files         103      103           
  Lines        3266     3266           
  Branches      729      729           
=======================================
  Hits         2534     2534           
  Misses        655      655           
  Partials       77       77           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1629?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @swadeley - Approved on May 21, 2026 at 07:22 AM UTC

ACK

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1629*
