---
type: pull_request
number: 1090
title: "Build: Bump the go group across 1 directory with 3 updates"
state: merged
author: dependabot
created: 2025-04-25T14:35:47Z
updated: 2025-04-29T13:12:02Z
closed: 2025-04-29T13:11:55Z
merged: 2025-04-29T13:11:55Z
base_branch: main
head_branch: dependabot/go_modules/go-3382638607
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1090
---

# Pull Request #1090: Build: Bump the go group across 1 directory with 3 updates

**Author**: @dependabot
**Created**: April 25, 2025 at 02:35 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-3382638607`

## Description

Bumps the go group with 3 updates in the / directory: [github.com/golang-migrate/migrate/v4](https://github.com/golang-migrate/migrate), [gorm.io/gorm](https://github.com/go-gorm/gorm) and [github.com/content-services/zest/release/v2025](https://github.com/content-services/zest).

Updates `github.com/golang-migrate/migrate/v4` from 4.18.2 to 4.18.3
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/golang-migrate/migrate/releases">github.com/golang-migrate/migrate/v4's releases</a>.</em></p>
<blockquote>
<h2>v4.18.3</h2>
<h2>Changelog</h2>
<ul>
<li>a4d0a1b Bump github.com/golang-jwt/jwt/v4 from 4.5.1 to 4.5.2</li>
<li>f37ef79 Bump golang.org/x/crypto from 0.31.0 to 0.35.0</li>
<li>5b97c92 Bump golang.org/x/net from 0.33.0 to 0.38.0</li>
<li>e6d84f6 Drop support for Go 1.22 and add support for Go 1.24</li>
<li>fccd197 Mention CLI install instructions in main README</li>
<li>34c2b4a Remove redundant build tags</li>
<li>a868033 Update FAQ.md - typo</li>
<li>7269490 Update golangci-lint version used in GitHub Actions</li>
<li>c5137c4 Update migrate -help output for the readme file</li>
<li>033835a Update to dktest v0.4.5</li>
<li>8b09191 fix: typo limited not limitted</li>
<li>60d73be refactor: replace github.com/pkg/errors with stdlib</li>
<li>36d17ba tests: fix various tests (<a href="https://redirect.github.com/golang-migrate/migrate/issues/1209">#1209</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/golang-migrate/migrate/commit/9023d66a0b649dbd6c99950903ebc7fd90d77ded"><code>9023d66</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1244">#1244</a> from alexandear-org/chore-redundant-build-tags</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/10494908afcb1f055c3ae007bfe1fb3c0100b7ee"><code>1049490</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1179">#1179</a> from lunfel/master</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/033835ac1bda6473673e3839c4383578c45258d9"><code>033835a</code></a> Update to dktest v0.4.5</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/329152ed6239455b8068774023ebec78132c3dc5"><code>329152e</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1196">#1196</a> from Rambatino/patch-1</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/34c2b4ad9c10984c40dfbe9348cdefdc8cf6b871"><code>34c2b4a</code></a> Remove redundant build tags</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/a3b7633246f2e0352cb3e5e907f0f6836cf16025"><code>a3b7633</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1259">#1259</a> from golang-migrate/dependabot/go_modules/golang.org...</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/5b97c921edce9c1a4b9aa0f642a23491fe84a72e"><code>5b97c92</code></a> Bump golang.org/x/net from 0.33.0 to 0.38.0</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/3c3ce91f29e79c86b65624231a1cfd78f4197847"><code>3c3ce91</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1258">#1258</a> from golang-migrate/dependabot/go_modules/golang.org...</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/f37ef795dbfb9dc56e217c1fee437ca1d463c304"><code>f37ef79</code></a> Bump golang.org/x/crypto from 0.31.0 to 0.35.0</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/1af841d787e67871a5af6eceb2a79099f8b9455d"><code>1af841d</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1260">#1260</a> from dhui/update_go</li>
<li>Additional commits viewable in <a href="https://github.com/golang-migrate/migrate/compare/v4.18.2...v4.18.3">compare view</a></li>
</ul>
</details>
<br />

Updates `gorm.io/gorm` from 1.25.12 to 1.26.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/go-gorm/gorm/releases">gorm.io/gorm's releases</a>.</em></p>
<blockquote>
<h2>Release v1.26.0</h2>
<h2>Changes</h2>
<ul>
<li>Preparestmt use LRU Map instead default map <a href="https://github.com/Yaccc"><code>@​Yaccc</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7435">#7435</a>)</li>
<li>use golangci replace reviewdog <a href="https://github.com/jinzhu"><code>@​jinzhu</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7426">#7426</a>)</li>
<li>test: mssql ci <a href="https://github.com/a631807682"><code>@​a631807682</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7388">#7388</a>)</li>
<li>fix deprecated reflect.PtrTo reflect.PointerTo usage <a href="https://github.com/Aman-Shitta"><code>@​Aman-Shitta</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7366">#7366</a>)</li>
<li>Fix concurrent map writes <a href="https://redirect.github.com/go-gorm/gorm/issues/7297">#7297</a> <a href="https://github.com/Ponywka"><code>@​Ponywka</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7298">#7298</a>)</li>
<li>chore: update copyright year <a href="https://github.com/maxktz"><code>@​maxktz</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7332">#7332</a>)</li>
<li>Enhance db.Scan with ParamsFilter - Issue 7336 - Suggestion <a href="https://github.com/evyaffe"><code>@​evyaffe</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7337">#7337</a>)</li>
<li>Fixed Empty Returning Clause Merge Edge Case <a href="https://github.com/aviyam181199"><code>@​aviyam181199</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7339">#7339</a>)</li>
<li>feat:Capitalize the priority field of IndexOption（<a href="https://redirect.github.com/go-gorm/gorm/issues/7331%EF%BC%89">go-gorm/gorm#7331</a> <a href="https://github.com/nowindexman"><code>@​nowindexman</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7342">#7342</a>)</li>
<li>fix: deterministic index ordering when migrating <a href="https://github.com/bamo"><code>@​bamo</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7208">#7208</a>)</li>
<li>use map look-up for indexes <a href="https://github.com/abbyssoul"><code>@​abbyssoul</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7242">#7242</a>)</li>
<li><a href="https://redirect.github.com/go-gorm/gorm/issues/6372">#6372</a> Fixed nullable constraint bug for columns during auto migration <a href="https://github.com/wookie0"><code>@​wookie0</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7269">#7269</a>)</li>
<li>Create CODE_OF_CONDUCT.md <a href="https://github.com/omidfth"><code>@​omidfth</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7240">#7240</a>)</li>
<li>Enhance Release Process: Implement Automated Release Management and Notes Generation <a href="https://github.com/YidiDev"><code>@​YidiDev</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7224">#7224</a>)</li>
<li>refactor: improve logging for unimplemented ErrorTranslator in TranlateError config enabled <a href="https://github.com/Invidam"><code>@​Invidam</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7225">#7225</a>)</li>
<li>Add GitHub Actions workflow to automate release creation on tagged pushes <a href="https://github.com/YidiDev"><code>@​YidiDev</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7209">#7209</a>)</li>
<li>Use official SQL Server docker image for tests <a href="https://github.com/omkar-foss"><code>@​omkar-foss</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7205">#7205</a>)</li>
<li>feat: Modernize Docker Compose File <a href="https://github.com/isso-719"><code>@​isso-719</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7086">#7086</a>)</li>
<li>Generate unique savepoint names for nested transactions <a href="https://github.com/phroggyy"><code>@​phroggyy</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7174">#7174</a>)</li>
<li>fix: AfterQuery using safer right trim while clearing from clause's j… <a href="https://github.com/bhowmik-abhijeet"><code>@​bhowmik-abhijeet</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7153">#7153</a>)</li>
<li>fix memory leaks in PrepareStatementDB <a href="https://github.com/ivila"><code>@​ivila</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7142">#7142</a>)</li>
<li>ci: Add PostgreSQL 14 and 15 to GitHub Actions matrix <a href="https://github.com/enomotodev"><code>@​enomotodev</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7081">#7081</a>)</li>
<li>feat: add MapColumns method <a href="https://github.com/molon"><code>@​molon</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/6901">#6901</a>)</li>
<li>add DB level propagation for the Unscoped flag <a href="https://github.com/sprataa"><code>@​sprataa</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7007">#7007</a>)</li>
<li>fix(scan): update Scan function to reset structs to zero values for each scan <a href="https://github.com/Waldeedle"><code>@​Waldeedle</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7061">#7061</a>)</li>
<li>fix: use reflect.Append when preloading nested associations instead of making a slice with fixed size <a href="https://github.com/emilienkofman"><code>@​emilienkofman</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7014">#7014</a>)</li>
<li>Fix association replace non-addressable panic <a href="https://github.com/SergeiSadov"><code>@​SergeiSadov</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7012">#7012</a>)</li>
<li>fix: <code>unsupported data</code> on nested joins with preloads <a href="https://github.com/N-Schaef"><code>@​N-Schaef</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/6957">#6957</a>)</li>
<li>fix: AfterQuery should clear FROM Clause's Joins rather than the Stat… <a href="https://github.com/liov"><code>@​liov</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7027">#7027</a>)</li>
<li>feat: chainable order support clause.OrderBy <a href="https://github.com/supergem3000"><code>@​supergem3000</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7054">#7054</a>)</li>
<li>fix: strings.Title -&gt; cases.Title bcs strings.Title library is deprecated <a href="https://github.com/ryuji-cre8ive"><code>@​ryuji-cre8ive</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/6999">#6999</a>)</li>
<li>fix: typo <a href="https://github.com/hakusai22"><code>@​hakusai22</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7003">#7003</a>)</li>
<li>Fix handling of unknown column types <a href="https://github.com/looi"><code>@​looi</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/6540">#6540</a>)</li>
<li>Fix panic bug in migrator due to lack of nil check for stmt.Schema <a href="https://github.com/pixelmaxQm"><code>@​pixelmaxQm</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/6932">#6932</a>)</li>
<li>Add new error for &quot;Violation Check Constraint&quot; <a href="https://github.com/anilsenay"><code>@​anilsenay</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/6992">#6992</a>)</li>
<li>fix: not clause with or condition <a href="https://github.com/a631807682"><code>@​a631807682</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/6984">#6984</a>)</li>
<li>perf: merge nested preload query when using join <a href="https://github.com/a631807682"><code>@​a631807682</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/6990">#6990</a>)</li>
<li>Faster utils.FileWithLineNum <a href="https://github.com/kkocdko"><code>@​kkocdko</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/6981">#6981</a>)</li>
<li>Added comment describing Unscoped() method <a href="https://github.com/AntonyChR"><code>@​AntonyChR</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/6969">#6969</a>)</li>
<li>fix: duplicated preload <a href="https://github.com/yetone"><code>@​yetone</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/6948">#6948</a>)</li>
<li>feat: prepare_stmt support ping <a href="https://github.com/philhuan"><code>@​philhuan</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/6924">#6924</a>)</li>
<li>fix: remove <code>callback</code> from <code>callbacks</code> if <code>Remove()</code> called <a href="https://github.com/snackmgmg"><code>@​snackmgmg</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/6916">#6916</a>)</li>
<li>fix: 'type XXXX int' will print wrong sql to terminal <a href="https://github.com/wangzeping722"><code>@​wangzeping722</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/6917">#6917</a>)</li>
<li>chore: optimize <code>regEnLetterAndMidline</code> regular <a href="https://github.com/demoManito"><code>@​demoManito</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/6908">#6908</a>)</li>
<li>fix(create): fix insert column order <a href="https://github.com/archever"><code>@​archever</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/6855">#6855</a>)</li>
<li>chore: optimize conversion of nanoseconds to milliseconds <a href="https://github.com/demoManito"><code>@​demoManito</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/6907">#6907</a>)</li>
<li>fix(scan): array element is set to a zero value <a href="https://github.com/demoManito"><code>@​demoManito</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/6890">#6890</a>)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/go-gorm/gorm/commit/a827495be126f896c0f744043410642d70bbac1b"><code>a827495</code></a> Preparestmt use LRU Map instead default map (<a href="https://redirect.github.com/go-gorm/gorm/issues/7435">#7435</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/489a56329318ee9316bc8a73f035320df6855e53"><code>489a563</code></a> only check new issues for golangci linter</li>
<li><a href="https://github.com/go-gorm/gorm/commit/42bd4f603c318fbdad5eff132679107b1ea34b07"><code>42bd4f6</code></a> use golangci replace reviewdog (<a href="https://redirect.github.com/go-gorm/gorm/issues/7426">#7426</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/a9d27293de2267a36fa6c9f8892977d3159cf8ea"><code>a9d2729</code></a> test: mssql ci</li>
<li><a href="https://github.com/go-gorm/gorm/commit/3876ffe4bb34bb8d5ceede0ded9694d34089bdef"><code>3876ffe</code></a> test: mssql ci</li>
<li><a href="https://github.com/go-gorm/gorm/commit/ee3b549d7dbdaad6610d80726093ada0c0ca90c1"><code>ee3b549</code></a> Update tests script</li>
<li><a href="https://github.com/go-gorm/gorm/commit/9f273777f58a247f7ae4176a014f6d59ac9fff8c"><code>9f27377</code></a> fix deprecated reflect.PtrTo reflect.PointerTo usage (<a href="https://redirect.github.com/go-gorm/gorm/issues/7366">#7366</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/9ca84b3dde0400ac69cc6cbf8dd8ec531f9638bd"><code>9ca84b3</code></a> fix concurrent map writes (<a href="https://redirect.github.com/go-gorm/gorm/issues/7298">#7298</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/86b1d229113ee803401645e7033d6c59245b197a"><code>86b1d22</code></a> chore: update copyright year (<a href="https://redirect.github.com/go-gorm/gorm/issues/7332">#7332</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/fed49230cbc0f84804cd5cc23d101e5ab1e068fa"><code>fed4923</code></a> Enhance db.Scan with ParamsFilter - Issue 7336 - Suggestion (<a href="https://redirect.github.com/go-gorm/gorm/issues/7337">#7337</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/go-gorm/gorm/compare/v1.25.12...v1.26.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2025` from 2025.4.1745343172 to 2025.4.1745523042
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/784e87e2a165b368fa4e4b1851d821e22fff8077"><code>784e87e</code></a> Update pulp bindings to d6a984592e569b42a26e536b94e27a48e9e38e1b7a8d53bd49838...</li>
<li><a href="https://github.com/content-services/zest/commit/ea8b09a8a5661dd224f7c558246074aacc0c1307"><code>ea8b09a</code></a> Update pulp bindings to b59d5a8be2d6a9584b32d696a45276e2b89bb3f27b8e3dae95843...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2025.4.1745343172...release/v2025.4.1745523042">compare view</a></li>
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
- `@dependabot merge` will merge this PR after your CI passes on it
- `@dependabot squash and merge` will squash and merge this PR after your CI passes on it
- `@dependabot cancel merge` will cancel a previously requested merge and block automerging
- `@dependabot reopen` will reopen this PR if it is closed
- `@dependabot close` will close this PR and stop Dependabot recreating it. You can achieve the same result by closing it manually
- `@dependabot show <dependency name> ignore conditions` will show all of the ignore conditions of the specified dependency
- `@dependabot ignore <dependency name> major version` will close this group update PR and stop Dependabot creating any more for the specific dependency's major version (unless you unignore this specific dependency's major version or upgrade to it yourself)
- `@dependabot ignore <dependency name> minor version` will close this group update PR and stop Dependabot creating any more for the specific dependency's minor version (unless you unignore this specific dependency's minor version or upgrade to it yourself)
- `@dependabot ignore <dependency name>` will close this group update PR and stop Dependabot creating any more for the specific dependency (unless you unignore this specific dependency or upgrade to it yourself)
- `@dependabot unignore <dependency name>` will remove all of the ignore conditions of the specified dependency
- `@dependabot unignore <dependency name> <ignore condition>` will remove the ignore condition of the specified dependency and ignore conditions


</details>

---

## Reviews

### Review by @rverdile - Approved on April 29, 2025 at 01:07 PM UTC

### Review by @rverdile - Approved on April 29, 2025 at 01:11 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1090*
