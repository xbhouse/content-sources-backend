# Key Contributors

Analysis of git commit history (9,474 commits across 9 repos), PR authorship (6,796 PRs), and review/approval activity from archived GitHub data. Bot accounts (dependabot, red-hat-konflux, semantic-release) are excluded from rankings.

**Data sources**: local git logs, `history/repos/*/pull-requests/`, `history/draft/data/important_prs.json`, `theme_analysis.json`

---

## Summary Rankings

### By commit volume (humans only)

| Rank | GitHub | Real name | Commits | Primary repos |
|------|--------|-----------|---------|---------------|
| 1 | @MichaelMraka | Michael Mraka | 914 | patchman-engine |
| 2 | @josef-hak | Josef Hak | 763 | patchman-engine |
| 3 | @jlsherrill | Justin Sherrill | ~632* | content-sources-backend, zest, caliri |
| 4 | @xbhouse | Bryttanie House | 604‡ | content-sources-backend, content-sources-frontend, yummy |
| 5 | @Andrewgdewar | Andrew Dewar | 602 | zest, content-sources-frontend, content-sources-backend |
| 6 | @psegedy | Patrik Segedy | 557 | patchman-engine |
| 7 | @jiridostal | Jiri Dostal | 351 | patchman-ui |
| 8 | @mhornick / @semtexzv | Michal Hornicky | ~349+† | patchman-engine |
| 9 | @mkholjuraev | Mukhammadkhol Kholiujurakhmad | 305 | patchman-ui |
| 10 | @swadeley | Stephen Wadeley | 268 | content-sources-frontend, content-sources-backend |
| 11 | @rverdile | Ryan Verdile | 259 | content-sources-backend, tang |
| 12 | @leSamo | Samuel Olekšák | 179 | patchman-ui |
| 13 | @dominikvagner | Dominik Vagner | 158 | content-sources-frontend, frontend-development-proxy |
| 14 | @Dugowitch | Jakub Dugovič | 140 | patchman-engine |
| 15 | @TenSt | Stepan Makymchuk | 133 | patchman-engine |

\* Justin Sherrill uses two emails (`jlsherrill@gmail.com` + `jsherril@redhat.com`); combined total shown.  
† Early patchman-engine work committed as @semtexzv (same person — Michal Hornicky).  
‡ Bryttanie House uses `bhouse@redhat.com` and `28575816+xbhouse@users.noreply.github.com`; combined total shown.

### By PR authorship (humans only)

| Rank | GitHub | PRs authored | Top repos |
|------|--------|-------------|-----------|
| 1 | @jlsherrill | 619 | content-sources-backend (467) |
| 2 | @josef-hak | 487 | patchman-engine (486) |
| 3 | @mkholjuraev | 363 | patchman-ui (350) |
| 4 | @psegedy | 343 | patchman-engine |
| 5 | @MichaelMraka | 336 | patchman-engine |
| 6 | @semtexzv | 303 | patchman-engine |
| 7 | @swadeley | 270 | content-sources-frontend (182) |
| 8 | @rverdile | 260 | content-sources-backend (217) |
| 9 | @Andrewgdewar | 218 | content-sources-frontend (131) |
| 10 | @xbhouse | 205 | content-sources-backend (116) |

### By review participation

| Rank | GitHub | Review comments |
|------|--------|----------------|
| 1 | @rverdile | 1,430 |
| 2 | @jlsherrill | 1,376 |
| 3 | @swadeley | 1,174 |
| 4 | @MichaelMraka | 861 |
| 5 | @xbhouse | 838 |
| 6 | @Andrewgdewar | 814 |
| 7 | @josef-hak | 589 |
| 8 | @semtexzv | 413 |
| 9 | @TenSt | 411 |
| 10 | @dominikvagner | 411 |

### By PR approvals (maintainer signal)

| Rank | GitHub | Approvals |
|------|--------|-----------|
| 1 | @MichaelMraka | 609 |
| 2 | @swadeley | 542 |
| 3 | @jlsherrill | 446 |
| 4 | @rverdile | 371 |
| 5 | @xbhouse | 346 |
| 6 | @TenSt | 309 |
| 7 | @Andrewgdewar | 271 |
| 8 | @josef-hak | 239 |
| 9 | @psegedy | 184 |
| 10 | @semtexzv | 164 |

---

## Contributor Profiles

### @MichaelMraka — Michael Mraka
**Most active**: 2022–2024 (peak Jul 2022 – May 2024)  
**Primary areas**: patchman-engine infrastructure — PostgreSQL migrations, Kafka configuration, OpenShift/Clowder deployments, developer setup, OpenAPI tooling, semantic-release integration  
**Notable impact**: Longest-tenured patchman-engine maintainer after the founding era; 609 PR approvals; 225 scored "important" PRs. Drove SSL/TLS hardening for Postgres and Kafka, Floorist config fixes, and ongoing schema migration reliability.  
**Email**: michael.mraka@redhat.com

### @josef-hak — Josef Hak
**Most active**: 2019–2022 (peak Jan–Mar 2020)  
**Primary areas**: patchman-engine core — Kafka clients, database performance, NEVRA parsing, benchmarking, advisory evaluation, Floorist/DataHub introduction (PR #898)  
**Notable impact**: 763 commits, 487 PRs, 183 important PRs. Built much of the early Go prototype and production engine. Activity tapered after Mar 2022 but reviews continued through 2024.  
**Email**: jhak@redhat.com

### @jlsherrill — Justin Sherrill
**Most active**: 2022–2024 (peak Feb 2023 – Jan 2024)  
**Primary areas**: content-sources ecosystem founder — backend architecture, introspection, OpenAPI spec, zest/caliri bindings, deployment/CI  
**Notable impact**: Founded content-sources-backend, caliri, and contributed to zest. 619 PRs (most authored in workspace), 1,376 reviews, 446 approvals. Authored introspection command (PR #55), early API design, and ongoing architectural oversight across content-services repos.  
**Email**: jlsherrill@gmail.com, jsherril@redhat.com

### @Andrewgdewar — Andrew Dewar
**Most active**: 2023–2025 (peak Mar 2023 – Mar 2025)  
**Primary areas**: zest (Pulp bindings), content-sources-frontend, snapshot/template features, Playwright CI  
**Notable impact**: Founded zest and yummy; bootstrapped content-sources-frontend (PR #1). 602 commits heavily concentrated in zest binding automation. Led snapshot endpoint (PR #458) and Playwright backend POC (PR #956).  
**Email**: andrewgdewar@gmail.com

### @psegedy — Patrik Segedy
**Most active**: 2021–2023 (peak Aug 2022 – Aug 2023)  
**Primary areas**: patchman-engine — org_id migration, Kafka listeners, account/cache management, evaluation pipeline  
**Notable impact**: 557 commits, 343 PRs, 202 important PRs. Authored org_id migration (PR #960) and many post-migration identity fixes. Primary engine contributor during the growth→maturity transition.  
**Email**: psegedy@redhat.com

### @jiridostal — Jiri Dostal
**Most active**: 2019–2020 (peak Jan–Oct 2020)  
**Primary areas**: patchman-ui — React/Redux frontend, webpack build, landing page, component library, early test infrastructure  
**Notable impact**: Founded patchman-ui (initial commit Nov 2019). 351 commits almost entirely in patchman-ui. Set the frontend patterns (Redux store, Header, deployment keys) that persisted for years.  
**Email**: jdostal@redhat.com, jiridostal@hotmail.com

### @semtexzv / @mhornick — Michal Hornicky
**Most active**: 2019–2021 (peak Jan–Mar 2020)  
**Primary areas**: patchman-engine genesis — Go prototype, database schema, Kafka, OpenShift templates, table partitioning, system table migrations  
**Notable impact**: Co-founded patchman-engine (Oct 2019 commits as semtexzv). 303 PRs as @semtexzv plus 349 commits as mhornick@redhat.com. Built base application structure from Go prototype, partition migration for system_package, Kafka implementation changes.  
**Email**: semtexzv@gmail.com, mhornick@redhat.com

### @mkholjuraev — Mukhammadkhol Kholiujurakhmad
**Most active**: 2022 (peak Mar–Jun 2022)  
**Primary areas**: patchman-ui — patch templates wizard, feature flags, Enzyme→RTL migration, internationalization  
**Notable impact**: 305 commits, 363 PRs, 249 important PRs (highest human score). Led patch template UI build-out (rename patch set → patch template, wizard, notifications). Drove RTL migration (PR #1164).  
**Email**: mkholjur@redhat.com

### @rverdile — Ryan Verdile
**Most active**: 2024–2026 (peak Sep 2024 – Apr 2026)  
**Primary areas**: content-sources-backend — database/dao layer, content templates, introspection, tang service  
**Notable impact**: Most prolific reviewer in workspace (1,430 review comments, 371 approvals). Founded tang (PR #1). Authored repository DB model (PR #4), content templates CRUD (PR #510). Bridge between backend architecture and code review culture.  
**Email**: ryanverdile@gmail.com

### @swadeley — Stephen Wadeley
**Most active**: 2025–2026 (peak Oct 2025 – Mar 2026)  
**Primary areas**: content-sources-frontend, Playwright testing, template Kafka events, QE workflows  
**Notable impact**: 270 PRs, 542 approvals (2nd highest). Primary frontend maintainer in the maturity period. Authored SendTemplateUpdateEvents job (PR #1537). Heavy review load across backend and frontend.  
**Email**: swadeley@redhat.com

### @leSamo — Samuel Olekšák
**Most active**: 2023–2025 (peak Feb 2023 – Jan 2025)  
**Primary areas**: patchman-ui — template detail pages, template list filters, template wizard UX  
**Notable impact**: 179 commits, 100 PRs, 80 important PRs — all patchman-ui. Built template detail page, assignment flows, and creator filters during the patch-templates feature expansion.  
**Email**: soleksak@redhat.com

### @dominikvagner — Dominik Vagner
**Most active**: 2025–2026 (peak Mar 2025 – Jan 2026)  
**Primary areas**: content-sources-frontend, frontend-development-proxy, Playwright, Konflux CI  
**Notable impact**: Founded frontend-development-proxy (May 2025). 158 commits across frontend repos. Initialized Playwright in patchman-ui (PR #1445). Konflux pipeline adoption (PR #1 on proxy repo).  
**Email**: dvagner@redhat.com

### @xbhouse — Bryttanie House
**Most active**: Oct 2023 – present (peak Mar–Jul 2025)  
**Primary areas**: content-sources backend/frontend, yummy introspection (comps, groups, environments), snapshots, patchman cross-integration  
**Notable impact**: 604 commits, 205 PRs, 838 reviews, 346 approvals. Joined mid-expansion as a full-stack contributor across API, UI, and yummy parsing. Authored groups/environments list API for repository introspection metadata (#468), template advisory endpoints (#1536), and patchman-ui/engine integration work during convergence.  
**Email**: bhouse@redhat.com, 28575816+xbhouse@users.noreply.github.com

### @Dugowitch — Jakub Dugovič
**Most active**: 2024–2026 (peak Feb 2024 – Mar 2026)  
**Primary areas**: patchman-engine — Cyndi/inventory sync, system_platform table split, database migrations  
**Notable impact**: 140 commits, 72 PRs. Authored system_platform split for Cyndi changes (PR #1988) — one of the most significant recent schema migrations.  
**Email**: jdugovic@redhat.com

### @TenSt — Stepan Makymchuk
**Most active**: 2025–2026 (peak Oct 2025 – Jun 2026)  
**Primary areas**: patchman-engine — DB migration infrastructure, ClowdApp jobs, deployment reliability  
**Notable impact**: 133 commits in ~9 months. Authored move db-migration to separate Job (PR #2236). 309 approvals — emerging maintainer for engine deployment architecture.  
**Email**: smaksymc@redhat.com

---

## Special Roles

### Project Founders (first commits per repo)

| Repo | Founder | Date | First commit |
|------|---------|------|-------------|
| patchman-engine | @jiridostal (Jiri Dostal) | Oct 2019 | Initial commit |
| patchman-engine | @semtexzv (Michal Hornicky) | Oct 2019 | Dockerfiles, Python scaffold, DB schema |
| patchman-ui | @jiridostal (Jiri Dostal) | Nov 2019 | Initial commit, build setup |
| content-sources-backend | @jlsherrill (Justin Sherrill) | Apr 2022 | Initial commit, ping API |
| content-sources-frontend | @Andrewgdewar (Andrew Dewar) | May 2022 | Initial commit, React boilerplate |
| yummy | @Andrewgdewar (Andrew Dewar) | May 2022 | Initial introspection library |
| zest | @Andrewgdewar (Andrew Dewar) | Mar 2023 | Init — Pulp bindings generator |
| tang | @rverdile (Ryan Verdile) | Nov 2023 | Initial commit, RPM search service |
| caliri | @jlsherrill (Justin Sherrill) | Feb 2024 | Initial commit — Candlepin bindings |
| frontend-development-proxy | @dominikvagner (Dominik Vagner) | May 2025 | Initial commit, HCC proxy |

**Patchman co-founders** (engine): @jiridostal, @semtexzv, @Josef Hak (@josef-hak), with early Go work from @josef-hak and @semtexzv in Nov 2019.

### Maintainers (frequent approvers, cross-repo)

These contributors consistently approve PRs across multiple repositories, indicating maintainer responsibility:

| GitHub | Approvals | Scope |
|--------|-----------|-------|
| @MichaelMraka | 609 | patchman-engine (primary gatekeeper) |
| @swadeley | 542 | content-sources-frontend/backend, patchman-ui |
| @jlsherrill | 446 | All content-services repos |
| @rverdile | 371 | content-sources-backend (primary), tang |
| @xbhouse | 346 | content-sources backend + frontend, yummy; growing patchman-ui/engine |
| @TenSt | 309 | patchman-engine deployment/migrations |
| @Andrewgdewar | 271 | zest, content-sources, tang |
| @josef-hak | 239 | patchman-engine (historical) |
| @psegedy | 184 | patchman-engine |
| @dominikvagner | 137 | frontend repos, proxy |

### Domain Experts (by repository / theme)

| Domain | Expert(s) | Evidence |
|--------|-----------|----------|
| **Patch evaluation engine** | @josef-hak, @MichaelMraka, @psegedy | 763 + 914 + 557 commits in patchman-engine |
| **Patch UI / templates** | @mkholjuraev, @leSamo, @jiridostal | 305 + 179 + 351 commits in patchman-ui; template wizard era 2022–2023 |
| **Content Sources API** | @jlsherrill, @xbhouse, @rverdile | Backend commits; snapshots, introspection, templates |
| **Pulp bindings (zest)** | @Andrewgdewar, @jlsherrill | 351 + 49 zest commits; automated binding releases |
| **Snapshot errata (tang)** | @rverdile, @xbhouse | tang founder + 6 commits (errata sorting, CVE listing) |
| **Introspection (yummy)** | @Andrewgdewar, @xbhouse, @jlsherrill | yummy founder + 7 PRs/18 commits (comps/groups) + PR #142 extraction |
| **Candlepin bindings (caliri)** | @jlsherrill | caliri founder; 15 commits |
| **Playwright / E2E testing** | @swadeley, @dominikvagner, @Andrewgdewar | Backend POC (#956), patchman-ui init (#1445), frontend CI |
| **DB migrations / Cyndi** | @Dugowitch, @TenSt, @MichaelMraka | system_platform split, migration Job, ongoing schema work |
| **Frontend dev tooling** | @dominikvagner | frontend-development-proxy founder; Konflux pipelines |

Theme analysis (`theme_analysis.json`) shows **playwright** (944 commits in maturity), **snapshots** (646), **pulp** (458), and **content templates** (226) as dominant maturity themes — contributors above map to those themes by repo affiliation rather than commit-message keyword matching (keywords rarely appear in commit subjects).

---

## Additional Notable Contributors

| GitHub | Real name | Role | Activity |
|--------|-----------|------|----------|
| @avisiedo | Alejandro Viciedo | Kafka introspection (PR #113), backend APIs | 31 PRs, 263 reviews |
| @mshriver | Mike Shriver | Early content-sources deployment, ClowdApp, CI | 42 PRs, 64 approvals |
| @marusak | Vendy | patchman-engine evaluation, reviews | 36 PRs, 239 reviews |
| @mayurilahane | Mayuri Lahane | content-sources-frontend features | 63 PRs |
| @katarinazaprazna | Katarina Zaprazna | Frontend testing, PatternFly migrations | 46 PRs, 179 reviews |
| @Hyperkid123 | — | Reviews across patchman-ui | 77 approvals |

---

## Collaboration Patterns

1. **Two product lines, shared leadership**: Patchman (engine + UI) and Content Sources (backend + frontend + libraries) evolved as parallel tracks with @jlsherrill bridging both via content-templates integration into patchman-ui.

2. **Review-heavy culture**: Top engineers spend as much time reviewing as authoring — @rverdile (1,430 reviews vs 260 PRs), @jlsherrill (1,376 vs 619), @xbhouse (838 vs 205). Knowledge transfer happens through PR discussion.

3. **Generational handoff on patchman-engine**: Founding era (@semtexzv, @josef-hak, 2019–2022) → @psegedy/@MichaelMraka (2021–2024) → @Dugowitch/@TenSt (2024–2026) for schema and deployment work.

4. **Content-services burst**: Apr 2022 founding → rapid team growth through 2024 templates/snapshots work → Playwright/Konflux modernization 2025–2026.

5. **Binding generators as shared infrastructure**: @Andrewgdewar and @jlsherrill maintain zest/caliri with automated PR streams consumed by the whole backend team.

---

## Validation Notes

- **Username ↔ name mapping** derived from git author fields and PR `author:` frontmatter. Confidence is **high** for Red Hat email addresses; **medium** for GitHub noreply handles matched by convention.
- **@semtexzv = Michal Hornicky** confirmed via git author string `Michal Hornicky <semtexzv@gmail.com>`.
- **@mkholjuraev** display name varies (`mkholjur` in commits); GitHub handle used consistently in PRs.
- **@xbhouse = Bryttanie House** confirmed via `bhouse@redhat.com` and PR author field; earlier analysis under-counted by matching GitHub handle only (142 → 604 commits).
- **@swadeley = Stephen Wadeley** — first name Stephen, not Sam; git author field uses `swadeley`.
- Bot activity (dependabot: 939 PRs, konflux: 997 PRs) excluded from human rankings but noted as significant automation contributors.
- Review counts include comment-only reviews, not just approvals.

---

*Generated: Phase 2, Prompt 2 · Next: validate attribution with Prompt 5 (`contributor_mapping.json`) during Phase 3*
