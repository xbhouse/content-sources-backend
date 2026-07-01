---
type: pull_request
number: 1421
title: "Build: Bump the go group with 4 updates"
state: merged
author: dependabot
created: 2026-03-23T04:33:31Z
updated: 2026-03-23T08:28:27Z
closed: 2026-03-23T08:28:24Z
merged: 2026-03-23T08:28:24Z
base_branch: main
head_branch: dependabot/go_modules/go-b2afe243a9
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1421
---

# Pull Request #1421: Build: Bump the go group with 4 updates

**Author**: @dependabot
**Created**: March 23, 2026 at 04:33 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-b2afe243a9`

## Description

Bumps the go group with 4 updates: [github.com/ProtonMail/go-crypto](https://github.com/ProtonMail/go-crypto), [github.com/lib/pq](https://github.com/lib/pq), [github.com/jackc/pgx/v5](https://github.com/jackc/pgx) and [github.com/project-kessel/kessel-sdk-go](https://github.com/project-kessel/kessel-sdk-go).

Updates `github.com/ProtonMail/go-crypto` from 1.4.0 to 1.4.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/ProtonMail/go-crypto/releases">github.com/ProtonMail/go-crypto's releases</a>.</em></p>
<blockquote>
<h2>Release v1.4.1</h2>
<h2>What's Changed</h2>
<ul>
<li>Properly handle ECC keys with invalid points in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/304">ProtonMail/go-crypto#304</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/ProtonMail/go-crypto/compare/v1.4.0...v1.4.1">https://github.com/ProtonMail/go-crypto/compare/v1.4.0...v1.4.1</a></p>
<h2>Release v1.4.1-proton</h2>
<p>This release is v1.4.1 with support for the following non-standardized features:</p>
<ul>
<li>Presistent symmetric keys experimental + latest draft <a href="https://www.ietf.org/archive/id/draft-ietf-openpgp-persistent-symmetric-keys-00.html">draft-ietf-openpgp-persistent-symmetric-keys-00</a></li>
<li>Automatic forwarding <a href="https://www.ietf.org/archive/id/draft-wussler-openpgp-forwarding-00.html">draft-wussler-openpgp-forwarding-00</a></li>
<li>Post-quantum algorithms <a href="https://datatracker.ietf.org/doc/draft-ietf-openpgp-pqc/">draft-ietf-openpgp-pqc</a> (Updated to draft-ietf-openpgp-pqc-09)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/2e73b118bb72881b92b292f85cb2d057c3d7bef0"><code>2e73b11</code></a> Properly handle ECC keys with invalid points (<a href="https://redirect.github.com/ProtonMail/go-crypto/issues/304">#304</a>)</li>
<li>See full diff in <a href="https://github.com/ProtonMail/go-crypto/compare/v1.4.0...v1.4.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/lib/pq` from 1.11.2 to 1.12.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/lib/pq/releases">github.com/lib/pq's releases</a>.</em></p>
<blockquote>
<h2>v1.12.0</h2>
<ul>
<li>
<p>The next release may change the default sslmode from <code>require</code> to <code>prefer</code>. See <a href="https://redirect.github.com/lib/pq/issues/1271">#1271</a> for details.</p>
</li>
<li>
<p><code>CopyIn()</code> and <code>CopyInToSchema()</code> have been marked as deprecated. These are simple query builders and not needed for <code>COPY [..] FROM STDIN</code> support (which is <em>not</em> deprecated). (<a href="https://redirect.github.com/lib/pq/issues/1279">#1279</a>)</p>
<pre><code>// Old
tx.Prepare(CopyIn(&quot;temp&quot;, &quot;num&quot;, &quot;text&quot;, &quot;blob&quot;, &quot;nothing&quot;))
<p>// Replacement
tx.Prepare(<code>copy temp (num, text, blob, nothing) from stdin</code>)
</code></pre></p>
</li>
</ul>
<h3>Features</h3>
<ul>
<li>
<p>Support protocol 3.2, and the <code>min_protocol_version</code> and <code>max_protocol_version</code> DSN parameters (<a href="https://redirect.github.com/lib/pq/issues/1258">#1258</a>).</p>
</li>
<li>
<p>Support <code>sslmode=prefer</code> and <code>sslmode=allow</code> (<a href="https://redirect.github.com/lib/pq/issues/1270">#1270</a>).</p>
</li>
<li>
<p>Support <code>ssl_min_protocol_version</code> and <code>ssl_max_protocol_version</code> (<a href="https://redirect.github.com/lib/pq/issues/1277">#1277</a>).</p>
</li>
<li>
<p>Support connection service file to load connection details (<a href="https://redirect.github.com/lib/pq/issues/1285">#1285</a>).</p>
</li>
<li>
<p>Support <code>sslrootcert=system</code> and use <code>~/.postgresql/root.crt</code> as the default value of sslrootcert (<a href="https://redirect.github.com/lib/pq/issues/1280">#1280</a>, <a href="https://redirect.github.com/lib/pq/issues/1281">#1281</a>).</p>
</li>
<li>
<p>Add a new <code>pqerror</code> package with PostgreSQL error codes (<a href="https://redirect.github.com/lib/pq/issues/1275">#1275</a>).</p>
<p>For example, to test if an error is a UNIQUE constraint violation:</p>
<pre><code>if pqErr, ok := errors.AsType[*pq.Error](https://github.com/lib/pq/blob/HEAD/err); ok &amp;&amp; pqErr.Code == pqerror.UniqueViolation {
    log.Fatalf(&quot;email %q already exsts&quot;, email)
}
</code></pre>
<p>To make this a bit more convenient, it also adds a <code>pq.As()</code> function:</p>
<pre><code>pqErr := pq.As(err, pqerror.UniqueViolation)
if pqErr != nil {
    log.Fatalf(&quot;email %q already exsts&quot;, email)
}
</code></pre>
</li>
</ul>
<h3>Fixes</h3>
<ul>
<li>
<p>Fix SSL key permission check to allow modes stricter than 0600/0640 (<a href="https://redirect.github.com/lib/pq/issues/1265">#1265</a>).</p>
</li>
<li>
<p>Fix Hstore to work with binary parameters (<a href="https://redirect.github.com/lib/pq/issues/1278">#1278</a>).</p>
</li>
<li>
<p>Clearer error when starting a new query while pq is still processing another query (<a href="https://redirect.github.com/lib/pq/issues/1272">#1272</a>).</p>
</li>
<li>
<p>Send intermediate CAs with client certificates, so they can be signed by an intermediate CA (<a href="https://redirect.github.com/lib/pq/issues/1267">#1267</a>).</p>
</li>
<li>
<p>Use <code>time.UTC</code> for UTC aliases such as <code>Etc/UTC</code> (<a href="https://redirect.github.com/lib/pq/issues/1283">#1283</a>).</p>
</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/lib/pq/blob/master/CHANGELOG.md">github.com/lib/pq's changelog</a>.</em></p>
<blockquote>
<h2>v1.12.0 (2026-03-18)</h2>
<ul>
<li>
<p>The next release may change the default sslmode from <code>require</code> to <code>prefer</code>.
See <a href="https://redirect.github.com/lib/pq/issues/1271">#1271</a> for details.</p>
</li>
<li>
<p><code>CopyIn()</code> and <code>CopyInToSchema()</code> have been marked as deprecated. These are
simple query builders and not needed for <code>COPY [..] FROM STDIN</code> support (which
is <em>not</em> deprecated). (<a href="https://redirect.github.com/lib/pq/issues/1279">#1279</a>)</p>
<pre><code>// Old
tx.Prepare(CopyIn(&quot;temp&quot;, &quot;num&quot;, &quot;text&quot;, &quot;blob&quot;, &quot;nothing&quot;))
<p>// Replacement
tx.Prepare(<code>copy temp (num, text, blob, nothing) from stdin</code>)
</code></pre></p>
</li>
</ul>
<h3>Features</h3>
<ul>
<li>
<p>Support protocol 3.2, and the <code>min_protocol_version</code> and
<code>max_protocol_version</code> DSN parameters (<a href="https://redirect.github.com/lib/pq/issues/1258">#1258</a>).</p>
</li>
<li>
<p>Support <code>sslmode=prefer</code> and <code>sslmode=allow</code> (<a href="https://redirect.github.com/lib/pq/issues/1270">#1270</a>).</p>
</li>
<li>
<p>Support <code>ssl_min_protocol_version</code> and <code>ssl_max_protocol_version</code> (<a href="https://redirect.github.com/lib/pq/issues/1277">#1277</a>).</p>
</li>
<li>
<p>Support connection service file to load connection details (<a href="https://redirect.github.com/lib/pq/issues/1285">#1285</a>).</p>
</li>
<li>
<p>Support <code>sslrootcert=system</code> and use <code>~/.postgresql/root.crt</code> as the default
value of sslrootcert (<a href="https://redirect.github.com/lib/pq/issues/1280">#1280</a>, <a href="https://redirect.github.com/lib/pq/issues/1281">#1281</a>).</p>
</li>
<li>
<p>Add a new <code>pqerror</code> package with PostgreSQL error codes (<a href="https://redirect.github.com/lib/pq/issues/1275">#1275</a>).</p>
<p>For example, to test if an error is a UNIQUE constraint violation:</p>
<pre><code>if pqErr, ok := errors.AsType[*pq.Error](https://github.com/lib/pq/blob/master/err); ok &amp;&amp; pqErr.Code == pqerror.UniqueViolation {
    log.Fatalf(&quot;email %q already exsts&quot;, email)
}
</code></pre>
<p>To make this a bit more convenient, it also adds a <code>pq.As()</code> function:</p>
<pre><code>pqErr := pq.As(err, pqerror.UniqueViolation)
if pqErr != nil {
    log.Fatalf(&quot;email %q already exsts&quot;, email)
}
</code></pre>
</li>
</ul>
<h3>Fixes</h3>
<ul>
<li>
<p>Fix SSL key permission check to allow modes stricter than <a href="https://redirect.github.com/0600/0640/issues/0600">0600/06400600</a> (<a href="https://redirect.github.com/lib/pq/issues/1265">#1265</a>).</p>
</li>
<li>
<p>Fix Hstore to work with binary parameters (<a href="https://redirect.github.com/lib/pq/issues/1278">#1278</a>).</p>
</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/lib/pq/commit/42ab0ff7c4b1b585fa6cd216a12a126715b5af06"><code>42ab0ff</code></a> Change default sslmode from &quot;require&quot; to &quot;prefer&quot;</li>
<li><a href="https://github.com/lib/pq/commit/6d40f13cfefe5ae4b489404b45114ba03224947a"><code>6d40f13</code></a> Release v1.12.0</li>
<li><a href="https://github.com/lib/pq/commit/386fc0e764f82951e9706dd8e618fc416fb5fb2c"><code>386fc0e</code></a> Document NULL behaviour with COPY</li>
<li><a href="https://github.com/lib/pq/commit/a62682e9083edee5ffe2d3904e80c30aaaa0877f"><code>a62682e</code></a> Better staticcheck cache 2</li>
<li><a href="https://github.com/lib/pq/commit/87ee06c600347478cfef07ad31b5ea617008aee7"><code>87ee06c</code></a> Better staticcheck cache</li>
<li><a href="https://github.com/lib/pq/commit/09624587cd3d4b69cd4a1a3b403b544ec0ddb2f7"><code>0962458</code></a> Rewrite tests to use pqerror, pq.As()</li>
<li><a href="https://github.com/lib/pq/commit/0d209810a155b9b8a173ab6d5da332f8c6cbc2c9"><code>0d20981</code></a> Don't move pq.Error to pqerror.Error</li>
<li><a href="https://github.com/lib/pq/commit/433213859e656dfe7dbba17b2838839e03a43ca2"><code>4332138</code></a> Add pqerror package</li>
<li><a href="https://github.com/lib/pq/commit/620d6d50817f9c4dcbfcbfeadb12ec552b636bf3"><code>620d6d5</code></a> Make tests run faster</li>
<li><a href="https://github.com/lib/pq/commit/dc8ff5d519b00bfa67626e68ea1ff3fdde0cc1b2"><code>dc8ff5d</code></a> Implement connection service file</li>
<li>Additional commits viewable in <a href="https://github.com/lib/pq/compare/v1.11.2...v1.12.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/jackc/pgx/v5` from 5.8.0 to 5.9.1
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jackc/pgx/blob/master/CHANGELOG.md">github.com/jackc/pgx/v5's changelog</a>.</em></p>
<blockquote>
<h1>5.9.1 (March 22, 2026)</h1>
<ul>
<li>Fix: batch result format corruption when using cached prepared statements (reported by Dirkjan Bussink)</li>
</ul>
<h1>5.9.0 (March 21, 2026)</h1>
<p>This release includes a number of new features such as SCRAM-SHA-256-PLUS support, OAuth authentication support, and
PostgreSQL protocol 3.2 support.</p>
<p>It significantly reduces the amount of network traffic when using prepared statements (which are used automatically by
default) by avoiding unnecessary Describe Portal messages. This also reduces local memory usage.</p>
<p>It also includes multiple fixes for potential DoS due to panic or OOM if connected to a malicious server that sends
deliberately malformed messages.</p>
<ul>
<li>Require Go 1.25+</li>
<li>Add SCRAM-SHA-256-PLUS support (Adam Brightwell)</li>
<li>Add OAuth authentication support for PostgreSQL 18 (David Schneider)</li>
<li>Add PostgreSQL protocol 3.2 support (Dirkjan Bussink)</li>
<li>Add tsvector type support (Adam Brightwell)</li>
<li>Skip Describe Portal for cached prepared statements reducing network round trips</li>
<li>Make LoadTypes query easier to support on &quot;postgres-like&quot; servers (Jelte Fennema-Nio)</li>
<li>Default empty user to current OS user matching libpq behavior (ShivangSrivastava)</li>
<li>Optimize LRU statement cache with custom linked list and node pooling (Mathias Bogaert)</li>
<li>Optimize date scanning by replacing regex with manual parsing (Mathias Bogaert)</li>
<li>Optimize pgio append/set functions with direct byte shifts (Mathias Bogaert)</li>
<li>Make RowsAffected faster (Abhishek Chanda)</li>
<li>Fix: Pipeline.Close panic when server sends multiple FATAL errors (Varun Chawla)</li>
<li>Fix: ContextWatcher goroutine leak (Hank Donnay)</li>
<li>Fix: stdlib discard connections with open transactions in ResetSession (Jeremy Schneider)</li>
<li>Fix: pipelineBatchResults.Exec silently swallowing lastRows error</li>
<li>Fix: ColumnTypeLength using BPCharArrayOID instead of BPCharOID</li>
<li>Fix: TSVector text encoding returning nil for valid empty tsvector</li>
<li>Fix: wrong error messages for Int2 and Int4 underflow</li>
<li>Fix: Numeric nil Int pointer dereference with Valid: true</li>
<li>Fix: reversed strings.ContainsAny arguments in Numeric.ScanScientific</li>
<li>Fix: message length parsing on 32-bit platforms</li>
<li>Fix: FunctionCallResponse.Decode mishandling of signed result size</li>
<li>Fix: returning wrong error in configTLS when DecryptPEMBlock fails (Maxim Motyshen)</li>
<li>Fix: misleading ParseConfig error when default_query_exec_mode is invalid (Skarm)</li>
<li>Fix: missed Unwatch in Pipeline error paths</li>
<li>Clarify too many failed acquire attempts error message</li>
<li>Better error wrapping with context and SQL statement (Aneesh Makala)</li>
<li>Enable govet and ineffassign linters (Federico Guerinoni)</li>
<li>Guard against various malformed binary messages (arrays, hstore, multirange, protocol messages)</li>
<li>Fix various godoc comments (ferhat elmas)</li>
<li>Fix typos in comments (Oleksandr Redko)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jackc/pgx/commit/4e4eaedb47b7b3cfba0a1b0a9e6a3f015764f046"><code>4e4eaed</code></a> Release v5.9.1</li>
<li><a href="https://github.com/jackc/pgx/commit/62731882651a90348febb43b2119b5f8bd9272de"><code>6273188</code></a> Fix batch result format corruption when using cached prepared statements</li>
<li><a href="https://github.com/jackc/pgx/commit/f7b90c2f1ac099f00e67d6d4d0fee6deb330bc94"><code>f7b90c2</code></a> Merge pull request <a href="https://redirect.github.com/jackc/pgx/issues/2524">#2524</a> from dbussink/pipeline-result-format-reuse</li>
<li><a href="https://github.com/jackc/pgx/commit/3ce6d75be4baa8d1e4b4f880da5f9ad68ab14e7f"><code>3ce6d75</code></a> Add failing test: batch scan corrupted in cache_statement mode</li>
<li><a href="https://github.com/jackc/pgx/commit/b4d8e62b6616d0c09c5021500363de0c56e01631"><code>b4d8e62</code></a> Release v5.9.0</li>
<li><a href="https://github.com/jackc/pgx/commit/c227cd4f76fa2b1a47c0156621e05c076f4cf5c9"><code>c227cd4</code></a> Bump minimum Go version from 1.24 to 1.25</li>
<li><a href="https://github.com/jackc/pgx/commit/f492c14836d7d442e8103b09f2c0c74a80c56347"><code>f492c14</code></a> Use reflect.TypeFor instead of reflect.TypeOf for static types</li>
<li><a href="https://github.com/jackc/pgx/commit/ad8fb08d3f1a36c0e475c9f80dc9bb19d075d8e2"><code>ad8fb08</code></a> Use sync.WaitGroup.Go to simplify goroutine spawning</li>
<li><a href="https://github.com/jackc/pgx/commit/303377376df43ba3d1a99728eaa9f9a6bcaab767"><code>3033773</code></a> Remove go1.26 build tag from synctest test</li>
<li><a href="https://github.com/jackc/pgx/commit/83ffb3c2220737cf11c7dd88c80be9166753102f"><code>83ffb3c</code></a> Validate multirange element count against source length before allocating</li>
<li>Additional commits viewable in <a href="https://github.com/jackc/pgx/compare/v5.8.0...v5.9.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/project-kessel/kessel-sdk-go` from 1.5.0 to 1.6.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/project-kessel/kessel-sdk-go/releases">github.com/project-kessel/kessel-sdk-go's releases</a>.</em></p>
<blockquote>
<h2>v1.6.0</h2>
<h2>What's Changed</h2>
<ul>
<li>Bump buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go from 1.36.11-20251209175733-2a1774d88802.1 to 1.36.11-20260209202127-80ab13bee0bf.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/pull/34">project-kessel/kessel-sdk-go#34</a></li>
<li>Bump github.com/zitadel/oidc/v3 from 3.44.0 to 3.45.5 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/pull/38">project-kessel/kessel-sdk-go#38</a></li>
<li>Bump google.golang.org/grpc from 1.73.0 to 1.79.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/pull/37">project-kessel/kessel-sdk-go#37</a></li>
<li>Update generated protobuf files by <a href="https://github.com/github-actions"><code>@​github-actions</code></a>[bot] in <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/pull/40">project-kessel/kessel-sdk-go#40</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/project-kessel/kessel-sdk-go/compare/v1.5.0...v1.6.0">https://github.com/project-kessel/kessel-sdk-go/compare/v1.5.0...v1.6.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/3632a37ee2c9892dc565f8676ca71cfeb2f7bf82"><code>3632a37</code></a> Merge pull request <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/issues/40">#40</a> from project-kessel/buf-generate-update</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/d845078d2acbd8cb00c785cf3e7f232fe1457c11"><code>d845078</code></a> Regenerate protobuf files from buf</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/cdf999695394942cde5a11c280b42f824406224f"><code>cdf9996</code></a> Merge pull request <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/issues/37">#37</a> from project-kessel/dependabot/go_modules/google.golan...</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/a6a934f23877dc7e38c31eb26c3bccf2e74dcf2f"><code>a6a934f</code></a> Go 1.23 EOL -&gt; go 1.24 at minimum</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/159ee50c379247da89239bfd63be1a7f872e91b9"><code>159ee50</code></a> Bump google.golang.org/grpc from 1.73.0 to 1.79.1</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/fb4f95e2540075df46d64f634bd97ef7dfb5d4f4"><code>fb4f95e</code></a> Merge pull request <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/issues/38">#38</a> from project-kessel/dependabot/go_modules/github.com/z...</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/a8689410a48b13d88c37663e3d9e3a1b7cb23ca9"><code>a868941</code></a> Bump github.com/zitadel/oidc/v3 from 3.44.0 to 3.45.5</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/876b33fedfebb6efa1a91795c5d43370147f55b3"><code>876b33f</code></a> Merge pull request <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/issues/34">#34</a> from project-kessel/dependabot/go_modules/buf.build/ge...</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/5348e644216c0879082ca73a7999d29eb89dffc5"><code>5348e64</code></a> Bump buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go</li>
<li>See full diff in <a href="https://github.com/project-kessel/kessel-sdk-go/compare/v1.5.0...v1.6.0">compare view</a></li>
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

### Review by @swadeley - Approved on March 23, 2026 at 08:28 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1421*
