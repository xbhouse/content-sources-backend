# Architectural Milestones

Scope for this history: how content-sources adopted Pulp, and how Lightwell repos were added. Origins and later product work (templates, Candlepin, uploads) are included only as context.

Repo shorthand: backend = content-sources-backend, frontend = content-sources-frontend.

The auto-scored PR list is still noisy (Dependabot, Konflux). This list is curated.

---

## Origins — background (2022-04 to 2023-02)

These made Pulp integration possible. Do not deep-dive unless a later Pulp PR requires them.

### 1. backend#4 — CONTENT-39: Repository Configurations model
- Why: Core `repository_configurations` row that snapshots later attach to.
- What changed: Schema and DAO for custom repos (name, URL, arch, version).
- Breaking: First schema.

### 2. backend#18 — CONTENT-60: expose API over /v1/
- Why: Stable Insights path (`/api/content-sources/v1/`).
- What changed: Versioned routing.

### 3. backend#55 — Fixes 49: add introspection command
- Why: Pre-Pulp way of learning repo contents (yummy / repomd). Snapshots later *mirror* content; introspection still reports status and package counts.
- What changed: Introspect command, later nightly cron and Kafka-triggered introspect (#113).

### 4. backend#149 — Fixes 470: Add RBAC integration
- Why: Authorization model before zest existed (Feb 2023).
- What changed: RBAC on routes; Redis cache followed in #264 (May, during Pulp work).

---

## Pulp integration — main act 1 (2023-03 to 2023-12)

### 5. zest (2023-03-09) — generated Pulp Go bindings
- Why: Stopped hand-writing Pulp HTTP. Backend consumes zest; bindings regenerate from Pulp OpenAPI.
- What changed: New repo; frequent Pulp-version bumps.
- Breaking: Backend now tracks zest/Pulp versions.

### 6. backend#248 — Fixes 1468: add repo snapshot task
- Why: The turn from “introspect metadata” to “mirror content in Pulp.” Everything after (templates, RH repos, uploads, Lightwell counts) assumes this pipeline.
- What changed: Snapshot task (remote → repo → sync → publication → distribution), snapshot DAO, DAO registry.
- Breaking: Feature-flagged; Pulp required. Recovery from mid-sync interruption was explicitly incomplete.

### 7. backend#253 — Fixes 1447: add postgres tasking system
- Why: Durable workers (heartbeat, requeue, cancel) instead of cron/Kafka one-shots. Snapshots needed this to be production-safe.
- What changed: Task table and enqueue API; nightly introspect moved onto it (#299). Kafka inspection removed later (#390).

### 8. backend#279 — Fixes 1471: Integrate snapshot task into application
- Why: Snapshots became a user-facing repo action, not a manual command.
- What changed: Wired into repo create/update and tasking; frontend toggle (#130).
- Breaking: Snapshot-enabled repos create Pulp remotes and take nightly sync (#387).

### 9. backend#308 — Fixes 1501: add API to list snapshots
- Why: Snapshots became a queryable resource.
- What changed: List API; listing detail (#342); snapshot task ID on repo response (#371).

### 10. backend#331 — Fixes 1939: pulp domain integration
- Why: Multi-tenant isolation in Pulp (per-org domains) instead of a shared namespace.
- What changed: Domain create/use on the Pulp client; S3 ACL / bucket hostname follow-ups (#355, #395).
- Breaking: Existing Pulp objects had to become domain-aware; stage repair jobs (#478–#492) cleaned leftover snapshots.

### 11. backend#340 — Fixes 2244,2280: simplify snap model and soft-delete repo config
- Why: Lifecycle: soft-delete repos, delete snapshots with the config (#314).
- What changed: Schema and delete semantics.

### 12. backend#387 — Fixes 1963: nightly sync
- Why: Snapshot freshness became automatic, not only on-demand.
- What changed: Nightly snapshot/sync task for snapshot-enabled repos.
- Follow-on: #407 added Red Hat repos to that nightly path (briefly reverted as #421, then restored).

### 13. backend#390 — Fixes 2264: remove kafka based inspection
- Why: Tasking fully replaced the 2022 Kafka introspect consumer.
- What changed: Kafka inspection path deleted.

### 14. backend#438 — Fixes 2896: refactor pulp client in dao
- Why: Pulp access centralized in the DAO layer — the client shape Lightwell and later features still use.
- What changed: Pulp client refactor.

### 15. tang (2023-11-30) — Pulp content search service
- Why: Split heavy Pulp/Postgres queries (RPMs, later errata, module streams, Maven/Python/npm) out of the API process.
- What changed: New Go service; backend calls tang for snapshot content.
- Breaking: New runtime dependency (Clowder RDS CA wiring in 2024, #629).

### 16. backend#458 — Fixes 1973: Add repository snapshot endpoint
- Why: Explicit “snapshot this repo now” API, completing the 2023 surface.
- What changed: Trigger endpoint alongside list (#308).

---

## Content Sources expansion — background (2024-01 to 2026-06)

Do not treat these as first-class architecture for this history. They show how the Pulp foundation was *used*. One paragraph each is enough in the narrative.

### 17. backend#510 — templates CRUD (Jan 2024)
Templates pin repos to a date/arch/version. They only work because snapshots exist. Candlepin environments (#570, #639) attach on top.

### 18. backend#608 / tang#6 — snapshot errata (Apr 2024)
tang grew `RpmRepositoryVersionErrataList`; backend exposed `listSnapshotErrata`. Same search-sidecar pattern Lightwell later uses for Maven/Python.

### 19. backend#732 — snapshot on create for upload repos (Jul 2024)
Customer-uploaded RPMs entered the same Pulp snapshot pipeline as URL repos.

### 20. backend#743 / #836 — `use_latest` and direct template↔snapshot association (2024)
Templates stopped resolving snapshots only at read time; they store the association. Retention (#901) then deletes unreferenced snapshots.

### 21. backend#935 — breaking `snapshots/for_date` date format (Jan 2025)
Clients could no longer send `YYYY-MM-DD`. Mention as the one documented snapshot API break.

### 22. backend#943 — feature content guards for RH repos (Jan 2025)
Entitlement-gated Red Hat content on Pulp distributions. Relevant as “Pulp content guards,” not as a Lightwell precursor.

---

## Lightwell — main act 2 (2026-07 to present)

### 23. backend#1557 / #1558 — HMS-10941 feature filter and HMS-10937 cached counts
- Why: Lightwell repos appear on the same repository list API, gated by feature, with cached Pulp/tang package/build counts instead of live hits on every list.
- What changed: List filter; count cache. Same day as the frontend app module.

### 24. frontend#1059–#1094 + insights-chrome#3582–#3597 — Lightwell Network UI (2026-07-01 through 2026-07-08)
- Why: The entire Network console shipped in the first week: chrome `/lightwell` shell, repo list, package list, Maven/Python detail, version switching, connect/copy snippets.
- What changed: chrome#3582 scalprum route; #3585 felt theme; #3589 hidden Insights nav; frontend#1059 app module; #1061 repo list; #1066 package list; #1068/#1069/#1075 detail and routing; #1088 connect popover (clients pointed at Pulp URLs).

### 25. backend#1559 — HMS-10936/10938: import for Lightwell repositories
- Why: Lightwell content is *imported* (already in Pulp), not customer-URL-introspected. Pulp is the source of truth for versions and builds.
- What changed: Import path for Lightwell repos.
- Breaking: Feature-gated product surface.

### 26. backend#1560 + tang#25 / #26 — Lightwell package APIs (Maven, Python)
- Why: Package model expanded beyond RPM: Maven GAV, Python distributions (npm followed). tang currently reads the latest Pulp repository version; Lightwell is moving that query onto Pulp catalog APIs ([pulp_maven#451](https://github.com/pulp/pulp_maven/pull/451), [pulp_python#1359](https://github.com/pulp/pulp_python/pull/1359), not landed).
- What changed: tang list/detail; backend package + versions endpoints; frontend package list/detail.
- Breaking: New routes; some Maven queries allow empty name/version.

### 27. backend#1562 / #1565 / frontend#1065 — Python (and later npm) on the Lightwell list
- Why: Repo-type expansion on the shared list, not a separate microservice or a second store.
- What changed: Python package listing; npm support in tang/zest shortly after.

### 28. backend#1609 + #1633 / frontend#1149 — OSV + Lightwell notifications
- Why: Advisories for non-RPM ecosystems, then notify subscribers (settings, prefs modal, send path, flags, dedupe).
- What changed: OSV sync (#1609); notification settings (#1601); prefs UI (#1149); send (#1633 HMS-11043).

### 29. Lens and Beacon (mid-August) — @katarinazaprazna / @dominikvagner plus UI/API
- Why: Coverage (Lens) and vulnerability tracker (Beacon) on the Network shell. Not Insights partner-repos.
- What changed: Coverage store/pipeline backend#1639/#1644/#1662 (@xbhouse, including S3 upload); analyzer UI + parsers frontend#1166 / backend#1645/#1655 (@katarinazaprazna); SBOM/pom/purl/ingestion backend#1667/#1673/#1641 (@dominikvagner); Beacon UI #1168 (@marusak); clearinghouse API #1640 (@TenSt); nav stub #1163 (@mattnolting).

### 30. backend#1664 / #1695 — LWLP-626: Lightwell JFrog bridge
- Why: Tell JFrog that an advisory exists. Not a Lightwell package source — Pulp remains SoT for versions/builds. Not the Insights partner-repo feature.
- What changed: Dedicated `platform.lightwell.advisory-created` topic; flag-gated bridge; #1695 follow-up.
- Breaking: New integration dependency (advisory path only).

### 31. backend#1643 / #1666 — LWLP-5: Lightwell Network API (phase 1)
- Why: Advisories API, cross-repo package search, and a Network API surface — first phase, after the week-one console.
- What changed: #1643 phase-1 shape; #1666 named Network API PR.

---

## Deep-dive priority (Phase 3)

Pulp (do these):

1. zest introduction (repo, not a single PR)
2. backend#248 snapshot task
3. backend#253 / #279 tasking + snapshot integration
4. backend#331 Pulp domains
5. backend#387 nightly sync (+ #407 RH repos)
6. tang introduction
7. backend#438 Pulp client refactor

Lightwell (do these):

8. backend#1559 Lightwell import
9. backend#1560 + tang#25/#26 package APIs
10. frontend week-1 Network UI (#1059–#1094) + insights-chrome `/lightwell` shell
11. backend#1664 JFrog advisory bridge (not a package store; not partner-repos)

Optional one-pagers (context only): #510 templates, #570 Candlepin, #935 `for_date` break.
