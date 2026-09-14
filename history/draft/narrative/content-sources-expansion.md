# Content Sources expansion (2024-01 to 2026-06)

Two and a half years, 1,468 non-bot commits, and one job for this history: show how the 2023 Pulp foundation was used. Templates, Candlepin, uploads, errata, and a breaking date format are consequences of snapshots. They are not a second Pulp origin story. Konflux, Playwright, and template-wizard UX stay out.

## Templates pin what Pulp already mirrored

backend#510 (@rverdile, 2024-01-15) added create/list/fetch for content templates: a name, repo UUIDs, arch, version — later a date or `use_latest`. The first related API was already in December 2023 (#486, `content_templates`). Delete (#535), update (#547), and the console table (frontend#193) followed immediately. Review was product-polish: OpenAPI `UUID` vs `uuid` (@swadeley), 404 text that said “repository” instead of “template” (@xbhouse).

`use_latest` (#743) and direct template↔snapshot association (#836) stopped resolving “the snapshot for this date” only at read time. Retention (#901) could then delete unreferenced snapshots. That is snapshot lifecycle, triggered by template semantics.

None of it works without #248’s distribution and #387’s nightly refresh. A template is a pointer into Pulp, plus later a Candlepin environment.

## Candlepin: entitlement side effect

backend#570 (@jlsherrill, 2024-02-28) bootstrapped a Candlepin client. Compose creates a devel org and imports a manifest; stage/prod use the user’s org (`devel_org` flips this). #639 created Candlepin content for a template; #670 and #682 fixed environment paths, GPG, and modular hotfixes.

## The same pipeline, more content shapes

Upload repos (#732, July 2024) entered the snapshot pipeline at create — no introspect URL. Templates later accepted them (#872). Errata (tang#6, backend#608, April 2024) made a snapshot inspectable for advisories the way #51 had done for RPM names in 2022, except the query now lived in tang. Module streams (#897 / tang#14) and feature-gated RH content (#943) are Pulp content guards and more tang columns, not a new store.

## Impact

Expansion proved the 2023 bet: if snapshots are durable and isolated, the product can be “pin a system to repo state” and “search what is in the pin.” The cost was API surface (dates, origins, feature flags) and a lot of UI. The next chapter does not add a third content engine. It imports a different catalog onto the same Pulp, tasks, and tang stack. Pulp remains the source of truth for Lightwell versions and builds.
