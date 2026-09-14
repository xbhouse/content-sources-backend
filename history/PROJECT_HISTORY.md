# Content Sources — project history

This history follows two decisions: how content-sources adopted Pulp, and how Lightwell repositories were added on that stack. Insights content-sources is the RPM product; Lightwell is a second product on the same backend for Java (Maven) and Python, with JavaScript next. Origins and the 2024–2026 product years are included as context, not as a feature chronicle.

Repos: [content-sources-backend](https://github.com/content-services/content-sources-backend), [content-sources-frontend](https://github.com/content-services/content-sources-frontend), [zest](https://github.com/content-services/zest), [tang](https://github.com/content-services/tang), [insights-chrome](https://github.com/RedHatInsights/insights-chrome) (Lightwell shell only, 2026-07-01 onward).

If you are joining to work on Lightwell: [Lightwell today](#lightwell-today) (path, surfaces, in flight), then [origins](draft/narrative/origins.md), [pulp-integration](draft/narrative/pulp-integration.md), and [lightwell](draft/narrative/lightwell.md). Expansion is Insights product context. Skip the raw timelines until you need a date.

## Executive summary

Content-sources is the Insights service for RPM repositories and the packages in them. It started in April 2022 as RPM-only: CRUD plus introspection — parse `repomd`, show package counts, do not host a copy. That shipped on Clowder with a console UI and RBAC. It is still the RPM product. Java (Maven) and Python arrived in 2026 as Lightwell, a separate product on the same backend. JavaScript is next for Lightwell.

In 2023 the team put Pulp behind those rows. zest generated the Go client. A snapshot task mirrored remote → sync → publication → distribution. A Postgres queue replaced Kafka for work that must resume. Pulp domains isolated orgs. Nightly sync and a shared RH catalog (`org_id=-1`) made snapshots something you could pin to. tang moved heavy “what is in this repository version?” queries out of the API process.

2024–mid-2026 used that foundation: content templates, Candlepin environments, upload repos, errata, one breaking change to `snapshots/for_date`. Those years are the largest commit bucket. They are a short chapter here.

In July 2026 Lightwell Network appeared: imported repos, not introspected URLs; Java (Maven) and Python, listed via tang today. JavaScript is the next type. Pulp is always the source of truth for versions and builds. Package search is moving from tang onto Pulp catalog APIs ([pulp_maven#451](https://github.com/pulp/pulp_maven/pull/451), [pulp_python#1359](https://github.com/pulp/pulp_python/pull/1359)); those PRs are open, the structure is there. The entire Network UI (insights-chrome `/lightwell` shell, repo list, package list/detail, connect snippets) shipped in the first week (2026-07-01 through 2026-07-08), then was polished. Later: notifications, Lens, Beacon, a JFrog advisory bridge, and a first-phase Network API. Insights partner-repos (foreign-org upload sharing) is a different product that happened to ship on the same calendar. Org-id middleware and the Insights notifications bus were the first things that did not fit.

The through-line: RPM metadata → RPM mirrors in Pulp (Insights content-sources) → a second product on the same machinery for Java and Python (JavaScript next). Artifactory/Nexus appear only as client connect snippets. JFrog is told about advisories; it does not store Lightwell package versions.

## Timeline

| Period | Dates | Theme | Major milestones |
|---|---|---|---|
| [Origins](draft/narrative/origins.md) | 2022-04 – 2023-02 | RPM metadata service | Repo model (#4), `/v1/` (#18), introspect (#55), Kafka introspect (#113), RBAC (#149) |
| [Pulp integration](draft/narrative/pulp-integration.md) | 2023-03 – 2023-12 | Mirror + isolate + read | zest; snapshot task (#248); tasking (#253/#279); domains (#331); nightly + RH (#387/#407); tang; client refactor (#438) |
| [Expansion](draft/narrative/content-sources-expansion.md) | 2024-01 – 2026-06 | Use the mirror | Templates (#510); Candlepin (#570); errata; uploads (#732); `for_date` break (#935) |
| [Lightwell](draft/narrative/lightwell.md) | 2026-07 – | Java and Python | Week-1 chrome + Network UI; Maven/Python packages (#1560 / tang#25); JavaScript next; later notifications, Lens/Beacon, JFrog (#1664), Network API (#1666) |

## Period summaries

### Origins (2022-04 – 2023-02)

Full narrative: [origins.md](draft/narrative/origins.md)

- Backend: @jlsherrill (2022-04-18). Frontend: @Andrewgdewar (2022-05-31). Repo model: @rverdile.
- Only package type: RPM.
- Introspection and RPM search made the UI useful.
- Insights-native (Clowder, OpenAPI, stage deploys) before it stored a single RPM.
- “Valid” meant “we parsed metadata.”

### Pulp integration (2023-03 – 2023-12)

Full narrative: [pulp-integration.md](draft/narrative/pulp-integration.md)

- Architectural year. Spine: @jlsherrill’s #248 / #279 / #331 / #387 and @rverdile’s #253 / #438, plus tang.
- Resume and domains were known-incomplete in v1; reviews said so.
- By December the product was Pulp plus tasks plus tang.

### Expansion (2024-01 – 2026-06)

Full narrative: [content-sources-expansion.md](draft/narrative/content-sources-expansion.md)

- Templates are pointers at snapshots.
- Candlepin is how a pointer becomes a RHEL environment.
- Uploads and errata reuse the same pipeline and sidecar.
- Largest team-volume chapter: templates/repos UI (@Starle21, @dominikvagner, @katarinazaprazna, @marusak, @Dugowitch) and QE (@swadeley, @mayurilahane), not only the APIs.

### Lightwell (2026-07 – )

Full narrative: [lightwell.md](draft/narrative/lightwell.md)

- Same backend and Pulp machinery as Insights content-sources. Lightwell itself is Java (Maven) and Python. JavaScript is next.
- Import-first repos, content-type-agnostic package APIs.
- Week one: insights-chrome shell and the Network UI.
- Later launches: notifications, Lens, Beacon, JFrog, Network API.
- tang is the current package query path, not the planned one.
- Did not reuse introspect-as-source-of-truth, Insights partner-repos, or the Insights notification firehose.
- Pulp stayed the package store.
- Current path, surfaces, and in-flight work: [Lightwell today](#lightwell-today).

## Architectural journey

1–8 are the stack still in use. 9–10 are Lightwell tenancy and the JFrog bus.

1. Postgres row (`repository_configurations`) is the customer-visible object throughout.
2. yummy introspection fills status and RPM lists. Still used; no longer the only truth.
3. zest is the written Pulp API. Do not hand-wrap Pulp.
4. Snapshot task copies a URL into Pulp. Distribution path is what users (and templates) consume.
5. Postgres tasking is the worker. Kafka inspection is deleted (#390). Payload holds Pulp task hrefs so resume works.
6. Pulp domains prefix paths per org. Created lazily. Pre-domain data needed repair.
7. tang reads Pulp’s DB for packages: RPMs (and errata, module streams) for Insights content-sources; Maven and Python for Lightwell. Catalog APIs in pulp_maven and pulp_python are meant to replace the Lightwell tang path.
8. Templates + Candlepin pin and entitle. Optional for understanding Pulp; required for understanding why nightly sync mattered.
9. Lightwell import adds rows that are not the caller’s org. List APIs and middleware had to change. Versions still come from Pulp; tang is the query layer until the plugin catalog APIs land.
10. JFrog bridge publishes advisory events to a dedicated Kafka topic. It is not a second package repository and not Insights partner-repos.

## Lightwell today

### Do not confuse

- Insights content-sources is RPM. Lightwell is Java (Maven) and Python on the same backend.
- A snapshot is a Pulp mirror of a URL. A Lightwell repo is imported. Introspection is not Lightwell’s source of truth.
- zest talks to Pulp. tang reads Pulp’s DB (for now). Pulp stores versions and builds.
- Artifactory/Nexus in the UI are connect snippets. JFrog is an outbound advisory topic. Neither holds Lightwell versions.
- Insights partner-repos (HMS-107xx) are foreign-org upload sharing. Not Lightwell. JFrog “partner” is not that feature.
- Templates and Candlepin are Insights entitlement. Lightwell does not use that path.
- Tickets: LWLP is Lightwell. HMS is Insights content-sources (including partner-repos).

### Request path

insights-chrome `/lightwell` → content-sources-frontend module → content-sources-backend → zest → Pulp (write and store). Package list/detail today: backend → tang → Pulp’s DB. Planned: backend → Pulp catalog APIs (see In flight).

List and fetch cannot assume every row is the caller’s org. Import writes foreign-org rows; Insights `enforce_consistent_org_id` returned 500 on the first import (#1559).

Where to change what:

- [insights-chrome](https://github.com/RedHatInsights/insights-chrome) — `/lightwell` shell only (route, theme, header, nav). Not Network, Lens, or Beacon pages.
- [content-sources-frontend](https://github.com/content-services/content-sources-frontend) — the Lightwell console (catalog, connect snippets, Lens, Beacon, notification prefs).
- [content-sources-backend](https://github.com/content-services/content-sources-backend) — APIs, import, tasks, Lens analysis, Beacon ingestion, notifications send, JFrog producer, Network API.
- [zest](https://github.com/content-services/zest) — generated Pulp Go client. Do not hand-edit bindings or wrap Pulp. Change Pulp’s OpenAPI and regenerate.
- [tang](https://github.com/content-services/tang) — current package query layer (Maven/Python for Lightwell). Shrinks when the catalog APIs land.
- Pulp — source of truth for versions, builds, and distribution URLs. Not in this org.

### Surfaces

| Surface | Job | Reads | Writes |
|---|---|---|---|
| Network | Catalog: repos, packages, versions, connect | backend list/fetch/packages → tang → Pulp | Import seeds rows (`external-repos import`). The UI does not write package versions. Connect tabs only copy Pulp URLs into Maven/Gradle/pip/Artifactory/Nexus clients. |
| Lens | Coverage: uploaded SBOM/manifest vs the catalog | Catalog via the same package path; report + match status | Upload to S3 on POST; analysis task; report tables |
| Beacon | Vulnerability tracker | Clearinghouse API | Ingestion job (and seed). PDF is export, not a store. |
| Notifications | Advisory mail/toasts and prefs | OSV sync; settings in Postgres | Send path; preference rows. Not the Insights notifications firehose. |
| JFrog | Outbound advisory bridge | Advisory events | Kafka `platform.lightwell.advisory-created`. Does not store packages. Not Insights partner-repos. |
| Network API | HTTP API, phase 1 | Advisories, cross-repo package search | Does not replace the Network console |

### In flight

Do not re-decide these from scratch.

- Package read path: move list/search/order off tang onto Pulp plugin catalog APIs. Shape is in open PRs from @TenSt: [pulp_maven#451](https://github.com/pulp/pulp_maven/pull/451), [pulp_python#1359](https://github.com/pulp/pulp_python/pull/1359). Neither has landed. Until they do, content-sources still goes tang → Pulp DB.
- JavaScript is the next Lightwell content type. Same pattern as Maven and Python (Pulp store, content-type-agnostic package APIs).
- Network API is still phase 1 (backend#1643 / #1666). It does not replace the console.
- JFrog stays a dedicated topic. The Insights notifications bus was tried and rejected (volume and per-org fan-out).
- Artifactory/Nexus stay connect snippets. They are not a second package repository.

## Community and contributors

Bots (Dependabot, Konflux, zest binding bumps) are excluded from timelines. They are most of the raw commit volume after 2024.

### Origins

- @jlsherrill — backend first commits (2022-04-18)
- @Andrewgdewar — frontend first commit (2022-05-31)
- @rverdile — repository-configuration model, early DAO
- @avisiedo — RPM search, Kafka introspect, RBAC
- @swadeley, @mshriver — QE and Clowder/IQE jobs

### Pulp integration

- @jlsherrill — snapshot task, integrate, domains, nightly, RH snapshotting
- @rverdile — tasking, snapshot list, Pulp client refactor, tang
- @Andrewgdewar — zest, early snapshot UI
- @dpang314 — admin/task APIs
- @swadeley — IQE for snapshots and domains

### Expansion

The Insights product years (templates, Candlepin, uploads, errata, snapshot lifecycle). Most of the team’s volume is here, not only the API work in the period summary.

- @jlsherrill, @rverdile, @xbhouse, @Andrewgdewar — snapshot, template, and Candlepin APIs and UI
- @dominikvagner, @katarinazaprazna — templates (wizard, assignment, retention, systems)
- @Starle21 (Lenka Staronova) — repositories and templates UI (styling, tables, uploads, remove-RPMs); backend upload/openapi and deletion-task work. Git author is `lstarono`
- @marusak, @Dugowitch (Jakub Dugovič) — templates and repositories UI
- @swadeley, @mayurilahane (Mayuri Lahane) — QE
- @TenSt — Pulp/search work that continues into Lightwell

### Lightwell

- Week-1 Network UI: @xbhouse, @rverdile, @arburka, @marusak, @ochosi
- insights-chrome shell: @florkbr, @marusak, @ochosi; later @karelhala, @dominikvagner, @katarinazaprazna
- Pulp/tang packages: @jlsherrill, @TenSt
- Notifications: @rverdile, @katarinazaprazna, @swadeley
- Lens: @xbhouse (tables, S3 upload, analysis task), @katarinazaprazna (analyzer UI and parsers), @dominikvagner (SBOM/pom/purl)
- Beacon: @dominikvagner (ingestion), @mattnolting, @marusak, @TenSt
- JFrog and Network API: @etsien, with @jlsherrill on the bus

Insights templates/repos on this calendar (not Lightwell Network): @johnelliott626 (John Elliott) — repositories and templates UI (columns, detail spacing, upload-repo URLs), never-introspected status, yummy bump. @Starle21 continues (RepositoriesCard, npm openapi). Partner-repos on the same dates is the content-sources product, not Lightwell.

## Where to read next

- Start: [Lightwell today](#lightwell-today), then [origins](draft/narrative/origins.md), [pulp-integration](draft/narrative/pulp-integration.md), [lightwell](draft/narrative/lightwell.md)
- Insights context: [expansion](draft/narrative/content-sources-expansion.md)
- Later, for a specific PR: [milestones](draft/curated/architectural_milestones.md), [deep-dives](draft/deep-dives/)
- How the periods were cut: [period_structure.md](draft/curated/period_structure.md)

