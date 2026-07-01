# Expansion (June 2022 – December 2023)

**2,284 commits · 6 repositories**

[← Project History](../PROJECT_HISTORY.md) · [← Production](production.md) · [Templates Era →](templates-era.md)

---

## Opening

The expansion period is when the workspace stopped being "Patchman plus some experiments" and became a **program**: two product lines sharing platform conventions but solving different problems. Content Sources — custom repository introspection, snapshotting, and (eventually) template composition — went from @jlsherrill's April 2022 backend commit to a six-repo ecosystem with Pulp integration, generated API clients, and a dedicated snapshot query service. Patchman did not pause; it absorbed org_id fallout, continued template UI work, and executed the largest schema migration since 2021.

Commit activity jumped 133% in June 2022 versus the prior six-month average. By December 2023 the workspace had content-sources-backend, content-sources-frontend, yummy, zest, and tang — plus patchman-engine and patchman-ui still contributing nearly 60% of commits. The period ends with snapshot infrastructure complete, introspection event-driven, and system_package2 migration landing — the prerequisites for the **Templates Era** that begins January 2024.

---

## Major Themes

### Content Sources founding: api/dao/handler from the first PR

@jlsherrill created content-sources-backend in April 2022; by June the frontend existed (@Andrewgdewar), repository CRUD was live, and @Ryan Verdile was establishing code layout conventions. The architectural pattern emerged immediately: PostgreSQL via GORM, OpenAPI-documented handlers, dao layer separation, org-scoped access via `x-rh-identity` headers through platform-go-middlewares.

PR #4 (May 2022, merged during expansion activity) set the database foundation. PR #19 (June 2022) established the create endpoint with org-scoped unique URL constraints. These were not boilerplate — they encoded decisions about transaction-wrapped migrations, force-rollback on failure, and identity integration that every subsequent feature inherited.

@jlsherrill became the highest-volume PR author in the workspace (619 PRs lifetime), and @Ryan Verdile emerged as the primary reviewer (1,430 review comments lifetime) — collaboration patterns that would define content-sources development for four years.

### Introspection: the core product capability

Content Sources' value proposition is knowing what packages exist in a customer's custom YUM/DNF repositories. PR #55 (@jlsherrill, July 2022) implemented the introspection CLI — downloading repodata, parsing packages, storing results in PostgreSQL. @jlsherrill self-reviewed with characteristic honesty: "A lot of these methods would benefit from a comment." The PR merged with known todos; the team iterated rather than blocked.

PR #113 (@avisiedo, October 2022) made introspection **event-driven**: a `kafka-introspect` command consuming Kafka messages to trigger introspection when repositories were created or updated. Schema-generated Go types, kafkacat test tooling, validation across three ephemeral Kafka pool configurations. @avisiedo's ephemeral test discovered 13,106 packages on first run — proof the pipeline worked at realistic scale.

@jlsherrill pushed for producer-side integration too: Kafka messages should fire when repos were created via API, not just when consumed. Contract-first Kafka schemas (.gen.go regeneration via `make build`) established the same discipline as OpenAPI on the REST side.

### yummy extraction: shared libraries begin

PR #142 (@rverdile, December 2022) extracted parsing logic from `dao/external_resources.go` into standalone `github.com/content-services/yummy`. @jlsherrill's review captured the architectural intent: "Overall, a big +1, i don't have any major comments."

This was the first of content-services' micro-libraries — yummy for repo parsing, zest for Pulp bindings, tang for snapshot queries, caliri for Candlepin bindings. The pattern: extract when reuse or independent testing justifies the module boundary; coordinate backend + library PRs with go.mod replace for local dev.

@xbhouse joined the program in October 2023 and would become yummy's primary maintainer — a contributor who bridged backend, frontend, and library work.

### Pulp ecosystem: zest, snapshots, tang

March 2023 saw zest created — automated Go client generation from Pulp's OpenAPI spec, eliminating hand-maintained HTTP clients. @jlsherrill drove binding automation; Konflux CI pipelines followed in later PRs. Every Pulp interaction in content-sources-backend would flow through zest-generated bindings.

PR #458 (@Andrewgdewar, December 2023) added `POST /repositories/{uuid}/snapshot/` — point-in-time repository state captured in Pulp for template composition. @jlsherrill's review: "Two small changes, and then this is good." The endpoint was manual-trigger-first; automated snapshot scheduling came later.

tang (PR #1, January 2024 — just after period boundary but rooted in November 2023 work) offloaded complex RPM/errata queries from the monolithic backend into a dedicated microservice with Pulp podman-compose for local dev. The expansion period built the **storage layer**; tang built the **query layer**.

PR #486 (snapshot-for-date API, December 2023) enabled templates to resolve "which snapshot existed on date X" — merged at period end, enabling the Templates Era's reproducible content views.

### Patchman at scale: system_package2 migration

While content-sources grew, patchman-engine faced its own scale crisis. The `system_package` table stored package data as JSON blobs per system — simple early on, catastrophic at 100k+ systems per account. RHINENG-5394 initiated migration to normalized `system_package2` with per-account batching.

PR #1353 (@psegedy, December 2023) implemented per-account migration: iterate accounts, migrate rows in batches, run VACUUM in goroutines to avoid API timeouts. Part of a multi-PR sequence (#1316 POC through #1356 drop old table) authored primarily by @psegedy and @MichaelMraka. @MichaelMraka approved; merged same day.

The parallel with January 2021's system table split is clear: **normalize when JSON convenience becomes query pain**. The difference is migration sophistication — per-account batching, async vacuum, POC validation through admin API before production migration.

### Patch templates UI and inventory integration

patchman-ui continued patch-set/template work (@mkholjuraev) — systems page cleanup, patch set assignment UX, rename from patch set to patch template (PR #832 area, June 2022). PR #1092/#1114 (2023) integrated Host Inventory groups into the patch systems view — tying compliance to inventory organization structure.

### zest and tang: generated clients and query offload

zest (March 2023) automated Go bindings from Pulp's OpenAPI spec. Before zest, Pulp HTTP clients were hand-maintained — every Pulp API upgrade risked drift between client and server. @jlsherrill drove generator templates, UBI-based builds, and Konflux CI (PR #27 area). The zest pattern: regenerate bindings on Pulp version bump, release versioned module, backend bumps go.mod.

tang emerged from the snapshot query problem: once snapshots lived in Pulp, answering "what RPMs are in this snapshot version?" required complex Pulp API traversal ill-suited to content-sources-backend's request path. tang PR #1 (@rverdile, January 2024) introduced `Tangy` interface with `RpmRepositoryVersionPackageSearch()` — dedicated microservice with podman-compose for local dev.

The tradeoff is service count: backend + yummy + zest + tang + caliri (starting February 2024) means five content-services components where expansion started with one backend. Each extraction was justified — independent testing, versioned release, specialized deployment — but operational surface area grew accordingly.

### Frontend parallel track

content-sources-frontend (@Andrewgdewar, @swadeley, @dominikvagner) shipped repository list/delete UI, pagination, ClowdApp deployment, module federation patterns, and snapshot testing hooks wired to backend PR #458. The frontend was not a thin CRUD shell — snapshot workflows, repository configuration forms, and admin task views required coordinated backend/frontend PRs across repos.

@xbhouse's October 2023 arrival accelerated introspection work — groups and environments metadata from comps parsing, yummy maintenance, and backend/frontend API depth.

---

### PR #4 — Repository configuration database (CONTENT-39)

@rverdile implemented GORM models, `dbmigrate` CLI with force-rollback on failure, migration file templates auto-wrapping transactions, and TestMain setup. [Deep dive](../deep-dives/content-sources-backend-pr-4-analysis.md)

Force rollback on failed migrations was non-negotiable — partial schema states are harder to debug than clean rollbacks. Migration templates ensuring transaction syntax in every new file prevented a class of "forgot to wrap in BEGIN/COMMIT" mistakes.

Without this PR: no repositories, no introspection, no snapshots, no templates. Everything in content-sources-backend depends on this foundation.

The week-long review with @jlsherrill was architectural tutoring as much as code review — migration ergonomics, clowder integration todos, TestMain patterns. @rverdile would go on to author 260 PRs and become the workspace's most active reviewer; PR #4 established his backend voice: careful schema work, explicit rollback semantics, documentation for local dev via Makefile env vars.

### PR #19 — Create content source endpoint (CONTENT-40)

@rverdile's June 2022 create endpoint established the **api/dao/handler** layering that became convention for 1,500+ subsequent backend PRs: API structs for request/response binding, dao layer for DB operations, handlers wiring HTTP to dao. Platform identity via `x-rh-identity` through platform-go-middlewares matched other Insights services. Org-scoped URL uniqueness — duplicate URLs rejected within a tenant, allowed across tenants — encoded multi-tenancy in the dao layer, not just middleware.

Merged during content-sources' "go live" month alongside frontend bootstrap — the first customer-facing API proved the PR #4 schema supported real requests.

### PR #55 — Repository introspection command (HMSCONTENT-49)

@jlsherrill's introspection CLI was the heart of the product — downloading and parsing external YUM/DNF repositories. SSL certificate configuration, public repo support, revision tracking. CLI-first delivery validated parsing logic before API automation.

The tradeoff: merged with known todos (cert config via config files, introspect-all repos, more tests). @avisiedo added tests in review. This incremental delivery style — ship the core path, iterate on edge cases — characterized content-sources' first two years.

### PR #113 — Kafka-driven introspection (HMSCONTENT-151)

@avisiedo converted introspection from CLI-only to event-driven. [Deep dive](../deep-dives/content-sources-backend-pr-113-analysis.md)

@jlsherrill on the request ID header: "We should probably not make it required." @avisiedo converted TODO comments to Jira cards before merge — review hygiene that kept PR scope bounded.

> @avisiedo: "IntrospectionUrl returned 13106 new packages"

That quote captures the moment introspection went from "code that parses repodata" to "system that handles real repositories at scale."

### PR #142 — Yummy library extraction (HMSCONTENT-200)

@rverdile moved parsing to yummy; backend imported via Go module. Established the micro-library architecture: shared code lives in versioned modules, backend coordinates releases, local replace enables parallel development.

Impact rippled through 2023–2024: comps parsing, groups support, compressed comps handling — all in yummy, consumed by backend, testable independently.

### PR #458 — Repository snapshot endpoint (HMS-1973)

@Andrewgdewar's snapshot endpoint introduced Pulp-backed content storage — the second pillar of content-sources after introspection. Manual trigger first; Pulp infrastructure required; November-to-December merge cycle with @jlsherrill and @swadeley review.

The frontend temporarily wired "Introspect Now" to the snapshot endpoint for testing — pragmatic integration before dedicated snapshot UI shipped in content-sources-frontend.

### PR #1353 — system_package per-account migration (RHINENG-5394)

@psegedy's per-account migration was one of patchman-engine's largest schema changes. Per-account batching avoided single long transactions; async vacuum preserved API responsiveness during migration operations.

Merged the same week as Cyndi filter changes — parallel inventory pipeline work showing patchman-engine's dual mandate: scale existing compliance data while adapting to platform inventory evolution.

---

## Impact

Expansion doubled the workspace's scope. Content Sources went from idea to introspection + snapshots + Pulp integration in 18 months. Patchman proved it could survive schema migrations at scale without API breaks. The shared library pattern (yummy, zest) and event-driven architecture (Kafka introspection) became content-services' DNA.

Tradeoffs acknowledged: Pulp as mandatory infrastructure (operational complexity), hand-off of snapshot queries to tang (service proliferation), patch-set templates in patchman-ui while content-sources built the replacement API (intentional parallel development), and org_id migration tail work bleeding into expansion's first weeks.

December 2023 is the natural period boundary: snapshot endpoint merged, system_package2 migration landing, tang maturing, @xbhouse active on groups/environments introspection APIs (#461, #468). January 2024 opens with content templates CRUD (PR #510) — the Templates Era's defining feature, built on every expansion-period foundation.

---

[← Project History](../PROJECT_HISTORY.md) · [← Production](production.md) · [Templates Era →](templates-era.md)

*Sources: [expansion timeline](../timeline/expansion.md) · [deep dives](../deep-dives/) (#4, #19, #55, #113, #142, #458, #1353) · [architectural milestones](../curated/architectural_milestones.md) (milestones 7, 16–23, 27–28)*
