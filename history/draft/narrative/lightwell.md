# Lightwell (2026-07 to present)

On 2026-07-01 the same backend grew a second product: Lightwell Network. Insights content-sources stayed the RPM product. Lightwell is Java (Maven) and Python only, with JavaScript next. It does not include RPMs. Lightwell repositories are not customer-typed URLs to introspect. They are imported into content-sources from a catalog that already exists. Pulp is always the source of truth for repository content, package versions, and builds. tang is what queries Pulp for Lightwell packages today. The console’s Artifactory/Nexus/pip tabs are connect snippets — how a developer points a client at the Pulp distribution URL — not a second store.

The stack is the 2023 one: zest, tasks, domains, tang. The tenancy and list contracts are not. Nothing Lightwell-related exists in git before this date. Current path, surfaces, and in-flight work: [Lightwell today](../../PROJECT_HISTORY.md#lightwell-today).

The product shipped in two tempos. Week one (2026-07-01 through 2026-07-08) built the Lightwell Network UI — app shell in insights-chrome, repository list, package list, Maven and Python detail, version switching, and connect/copy flows. The same surface was then polished, and Lightwell grew notifications, Lens, Beacon, JFrog integration, and a first-phase Network API.

Insights partner repositories (HMS-107xx, backend#1596, foreign-org upload sharing) are a different product. They share a calendar with Lightwell. They are not Lightwell, and “partner” on the JFrog advisory bridge is not that feature.

## Week one: chrome shell + Network UI (7/1–7/8)

insights-chrome got a `/lightwell` layout on day one: scalprum route (#3582, @florkbr), felt theme (#3585), hidden Insights chrome (#3589, @marusak), simplified header (#3591), no theme/banner flicker (#3592, @ochosi). The Network console is a content-sources frontend module inside that shell.

Frontend#1059 (@xbhouse) added the empty table, app entry, and module on day one. frontend#1061 was the repo list (HMS-10939). frontend#1066 the package list (HMS-10957). frontend#1068 / #1069 / #1075 were Java and Python package detail and version/group routing (@rverdile, @xbhouse). frontend#1088 (@arburka) put auth and tool tabs (Maven, Gradle, pip, Poetry, Artifactory, Nexus) in the connect popover — still documenting how to consume Pulp URLs. @marusak hid in-app chrome that would navigate away from Lightwell (#1063). @ochosi wired stylesheet overrides and RBAC-safe routing (#1074, #1079).

The backend of that week is import and read APIs, not a new content engine. backend#1559 (@rverdile) imported Lightwell repo rows (`external-repos import`, idempotent). List/fetch were deferred on purpose. @jlsherrill hit `enforce_consistent_org_id` (500, org mismatch): Insights middleware assumed every row belonged to the caller. Feature filter on list and cached Pulp/tang counts (backend#1557 HMS-10941 / #1558 HMS-10937) landed 2026-07-01 so the new tables would not live-query on every load. Fetch (#1561), Python as a repo type (#1562) and on the list (#1565 / frontend#1065), and package APIs (#1560 + tang#25/#26) filled the week. Maven/Python details and versions (backend#1564/#1568) closed it.

By 2026-07-08 the Network console was a complete path: Lightwell-branded chrome, see repos, see packages, see a version, copy a Pulp URL into your build tool.

## Pulp-backed packages (same week, read path)

backend#1560 and tang#25/#26 (2026-07-03) added content-type-agnostic list/detail:

> “these two apis are agnostic to what content type it is. (although for something like python, the group field may be blank).” — @jlsherrill

tang reads the latest Pulp repository version. There is no JFrog or Artifactory query for “what versions exist.” JavaScript is intended to use the same pattern next.

That tang path is temporary. Lightwell is moving package list/search/order off tang and onto catalog APIs in the Pulp plugins themselves. The shape is already in open PRs from @TenSt: [pulp_maven#451](https://github.com/pulp/pulp_maven/pull/451) (search on groupId/artifactId, last_updated ordering, newest-first versions, trigram indexes) and [pulp_python#1359](https://github.com/pulp/pulp_python/pull/1359) (HTTP `packages`/`metrics` catalog, name search, PEP 440 newest-first). Neither has landed. Until they do, content-sources still goes tang → Pulp DB.

## After week one

Chrome kept going: branding (#3617), logo and logout (@karelhala, #3638/#3645/#3648), console nav entry (#3616), horizontal subnav (#3656, #3663), theme-aware icons (#3670), same layout at all widths (#3675). @dominikvagner dropped the last breadcrumb segment for Lens/Beacon routes (#3676). @katarinazaprazna put a minimal footer outside the page card (#3679). frontend#1213 (LWLP-750) hid Chrome favorites on Lightwell pages.

Catalog polish continued (spacing, copy buttons, `/lightwell` distribution URLs, Python Artifactory snippet fixes in frontend#1116 — which Pulp path to paste into an Artifactory client, not where packages live).

Launches on that shell:

- Notifications (late July–August). Settings in the DB (backend#1601, @swadeley), preferences modal (frontend#1149, @katarinazaprazna), send path (backend#1633, @rverdile). Format change (#1625), severity (#1160), feature flag (#1651 / frontend#1169), per-org dedupe (#1676), predisclosure repo (#1678 / frontend#1205). OSV advisory sync (#1609) is the data behind the mails.
- Lens. Coverage of an uploaded SBOM/manifest against the Lightwell catalog. @xbhouse: report tables (backend#1639), store the upload in S3 on POST (backend#1644), analysis task (backend#1662), match status on the API (backend#1680), flags (#1649/#1677, frontend#1189), plus upload size / progress / uploader / report route (frontend#1187/#1192/#1206/#1216). @katarinazaprazna: coverage analyzer UI (frontend#1166), charts and polish (#1173/#1183/#1188/#1186), SPUR (#1191/#1221), manifest constraints (#1193/#1200/#1203), CSV/requirements parsers and package matcher (backend#1645/#1655). @dominikvagner: SBOM and pom.xml parsers (#1667/#1673), purl (#1711/#1715), unlock Lens for internal employees (#1725). @arburka redesigned Coverage Analyzer chrome (frontend#1175).
- Beacon. Vulnerability tracker on the same nav. frontend#1163 (@mattnolting) added the nav shell and a tracker stub; frontend#1168 (@marusak) is the Beacon UI, with #1170/#1172 and SPUR updates (#1199/#1214). backend#1640 (@TenSt) is the clearinghouse read API. @dominikvagner’s #1641 is the vulnerability ingestion job; #1705 seeded vulns. PDF export: frontend#1184 (@ochosi).
- JFrog integration. Advisory bridge, not a package store and not Insights partner-repos. backend#1664 (@etsien) first tried the Insights notifications firehose; @jlsherrill stopped that (millions of messages, per-org fan-out). Dedicated topic `platform.lightwell.advisory-created`, flag-gated, with an admin test endpoint. Follow-ups: #1695, #1707, LWLP-951. Pulp remains the version/build store; JFrog is told when an advisory exists.
- Network API, phase 1. backend#1643 (@etsien) started advisories API, cross-repo package search, and response-shape work. backend#1666 is the named Lightwell Network API PR. First phase; it does not replace the Network console.

## What was reused, what was not

Reused: zest, Pulp as content store, domain-scoped client, tang as the current read model (to be replaced by Pulp catalog APIs), import-as-seed, feature flags, insights-chrome as the product shell.

Not reused: customer URL introspection as source of truth, Candlepin, `org_id` equality on every list, the Insights notifications topic, RPM-shaped package fields, the Insights partner-repo sharing model.

@xbhouse, @rverdile, @arburka, @marusak, and @ochosi built the week-one Network UI. @florkbr, @marusak, @ochosi, and chrome platform work made `/lightwell` a first-class layout that week. @jlsherrill and @TenSt owned Pulp/tang package APIs. Notifications: @rverdile, @katarinazaprazna, @swadeley. Lens: @xbhouse (coverage tables, S3 upload, analysis task), @katarinazaprazna (analyzer UI and parsers), @dominikvagner (SBOM/pom/purl). Beacon: @dominikvagner (ingestion), @mattnolting, @marusak, @TenSt. JFrog/Network API: @etsien, with Justin on the bus.

## Impact

Lightwell Network’s catalog shipped in eight days on the 2023 Pulp foundation and a dedicated insights-chrome shell. Notifications, Lens, Beacon, JFrog, and the Network API launched on that same stack. The console never grew a second package database.
