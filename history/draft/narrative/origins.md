# Origins (2022-04 to 2023-02)

Content-sources began as an Insights service that could remember a customer’s custom RPM repositories and tell them what was in those repos. RPM was the only package type. It did not yet mirror content. That distinction — metadata versus a Pulp snapshot — is the whole point of the next chapter. This one is how the service became real enough to hang Pulp on. Java and Python arrive in 2026; they are not part of this period.

@jlsherrill’s first backend commits landed on 2022-04-18: a Go app, Dockerfile, Clowder-shaped deployment, and a ping API. @Andrewgdewar started the frontend on 2022-05-31. @rverdile joined within weeks on the repository-configuration model. By February 2023 the API lived at `/api/content-sources/v1/`, the console UI listed and edited repos, introspection ran on a schedule (and on Kafka), and Insights RBAC gated writes. None of that required Pulp.

## A repository is a row, not a mirror

The core object is still `repository_configurations` (backend#4, @rverdile, May 2022): name, URL, architecture, distribution versions, later GPG key and verification flags. CRUD, pagination, filters, and org-scoping (backend#29) made it an Insights API instead of a script. backend#18 put the public path on `/v1/`. @mshriver moved the OpenAPI spec into the layout the rest of the org expects (backend#24).

What the service knew about a repo came from introspection, not from hosting packages. backend#55 (@jlsherrill, July 2022) added the introspect command: fetch `repomd` via yummy, record package count and status. @avisiedo added RPM list and package search (backend#42, #51). The frontend grew a packages modal and an introspection-status column (@rverdile, frontend#29). Daily introspect-all (backend#76) and “smarter” introspect (backend#104) were about not hammering every URL every time.

This is the model Lightwell later refused: customer-typed URL, periodic metadata scrape, no frozen tree. Snapshots keep introspection for status. They add a second, Pulp-backed truth.

## Insights-shaped from week one

Clowder was not a later migration. backend#10 (CONTENT-58) loaded DB config from Clowder. @mshriver’s `pr_check` / `build_deploy` and frontend ClowdApp templates (August 2022) put the app on the same path as other console services. By November the frontend was pushing to stage-beta and stage-stable (frontend#66, #69).

Auth followed the platform: identity-header middleware (backend#88, @avisiedo), then full RBAC in backend#149 (February 2023) — the last origins milestone. Redis caching of RBAC (#264) arrived in May, during Pulp work; the permission model did not.

Kafka entered as an introspect trigger (backend#113, @avisiedo, October 2022), not as a content bus. That consumer is the thing the 2023 Postgres tasking system replaced. Popular repositories (backend#148, December 2022) added a curated catalog — still URLs to introspect, not content to host.

## What this period did not decide

Nobody here chose Pulp domains, a snapshot resume story, or a search sidecar. zest did not exist. The DAO layer and OpenAPI contract were ready for those things. The product was custom RPM repos in Insights, and it shipped that way.

The people who would build Pulp integration were already in the repo: @jlsherrill on architecture and introspect, @rverdile on the DAO, @Andrewgdewar on the UI that would later grow a snapshot toggle, @avisiedo on search/Kafka/RBAC, @swadeley and @mshriver on QE and deploy.

## Impact

Origins left a production-shaped Insights service that understood references to RPM repositories and nothing else. The next nine months turned those references into Pulp remotes. Until May 2023, “the repo is valid” meant “we parsed its RPM metadata,” not “we have a copy.”
