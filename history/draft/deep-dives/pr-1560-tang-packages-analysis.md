# backend#1560 + tang#25/#26 — Lightwell package APIs

Authors: @jlsherrill (backend#1560, tang#25), @TenSt (tang#26, Python)  
Merged: 2026-07-03  
URLs: https://github.com/content-services/content-sources-backend/pull/1560 · https://github.com/content-services/tang/pull/25 · https://github.com/content-services/tang/pull/26  
Period: lightwell

## Problem

Insights snapshot UIs list RPMs (and errata) via tang. Lightwell needed the same for Maven (GAV + builds) and Python (distributions), without a new search service. RPM-shaped APIs (`name`, `version`, `arch`) do not fit Maven `groupId:artifactId`.

## Approach

tang#25 added list support for Maven packages and builds on the latest repository version. @TenSt approved with a note on in-memory vs SQL pagination; RPM tests broke on an old Pulp API and the one-line fix landed in tang#26 (Python listing). backend#1560 exposed those tang methods as content-type-agnostic HTTP APIs: “these two apis are agnostic to what content type it is (although for something like python, the group field may be blank).” Merge waited on a tang release so CI could compile.

Follow-ups the same week: Maven details (#1564), Python details (tang#30), versions list (#32, #1564/#1576), prefix search (tang#28/#29). Frontend package list/detail pages landed in parallel (@xbhouse, @rverdile).

## Design decisions

- Reuse tang, do not add a Lightwell-only query service. Same Pulp DB, new content types.
- Agnostic API, optional group. Avoids `MavenPackage` vs `PythonPackage` route explosion on the backend. Clients that need GAV pass group; Python leaves it empty.
- Latest Pulp repo version only (v1). tang#25’s description: “the latest repo version of a repository.” Pulp is the source of truth for versions and builds; there is no Artifactory/JFrog query on this path.
- tang is the current Lightwell query layer, not the destination. Catalog search/order APIs are being added in Pulp: [pulp_maven#451](https://github.com/pulp/pulp_maven/pull/451) and [pulp_python#1359](https://github.com/pulp/pulp_python/pull/1359) (@TenSt). Neither has merged; the endpoints and indexes are already structured there.

## Impact

This is how Lightwell becomes more than a repo table: packages, builds, versions, then OSV advisories (#1609). The pattern is the 2024 errata split (tang#6 / backend#608) applied to non-RPM types.

## Quote

> “The idea being that these two apis are agnostic to what content type it is. (although for something like python, the group field may be blank).” — @jlsherrill, #1560
