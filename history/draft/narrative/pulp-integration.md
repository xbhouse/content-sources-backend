# Pulp integration (2023-03 to 2023-12)

This is the year content-sources stopped being a metadata service and became a Pulp frontend with its own Postgres, API, and workers. zest appeared in March. By December, snapshots ran on a durable task queue, each org had a Pulp domain, nightly sync existed, Red Hat repos shared that pipeline, and tang was started so the API process would not query Pulp’s database itself.

The sequence was deliberate, and the gaps were spoken out loud in review. Snapshot create (#248) could not resume a mid-sync Pulp task. Tasking (#253) existed before snapshots used it. Integration (#279) stored Pulp hrefs so a killed worker could continue. Domains (#331) isolated tenants after the first snapshots had already been taken in a shared namespace — stage later needed repair jobs. That is not a tidy rewrite. It is how the foundation was actually poured.

## zest: stop handwriting Pulp

@Andrewgdewar created [zest](https://github.com/content-services/zest) on 2023-03-09. The job is to generate Go bindings from Pulp’s OpenAPI and publish a module the backend pins. Almost all zest git history is automated `Update pulp bindings` commits; those are excluded from this analysis. The decision that matters is: Pulp’s spec is the source of truth. When the spec was wrong, the team patched generated types (for example @rverdile’s `TaskResponse.Error` → string) rather than wrapping Pulp by hand.

Versioning thrashed in mid-2023 (`v$year`, content checksums as versions) and settled on a normal module path. @jlsherrill and later @TenSt kept the generator alive (UBI image, 404 retries, domain flags, operation-id collisions). Every later snapshot, domain, and Lightwell Pulp call goes through this client.

## The snapshot pipeline, then the queue

backend#248 (@jlsherrill, merged 2023-05-03) added the pipeline: Pulp remote → repository → sync → publication → distribution, plus a snapshot DAO and a DAO registry. The first interface was a command: `go run cmd/external-repos/main.go pulp-create <url> <org>`. Users still did not pass `snapshot=true`. Resume worked after remote or repo creation, not during sync, publish, or distribute.

@rverdile tested that. He Ctrl-C’d `pulp-create`, reran it, got “successfully created,” and found nothing at `/pulp/content`. Justin’s reply set the next two months of work:

> “What we're gonna need to do is save the task ids on the task itself somehow (to be done during integration) and then as part of the task check that the sync task/publish task is completed, and if not retry it.”

backend#253 (@rverdile, 2023-05-16) put a Postgres tasking system under introspect first: queue, client, worker pool, feature flag `new_tasking_system`. Kafka remained the old path. Justin asked for architecture docs and for queue-latency / processing-failure metrics; they were already in the worker. The task table is still the execution model — later template updates, upload conversions, and Lightwell import are rows in it.

backend#279 (@jlsherrill, 2023-06-03) connected the two. Create/bulk-create with `snapshot=true` enqueued a snapshot task whose payload stored Pulp sync, publication, and distribution hrefs. Kill the worker, restart it, the snapshot should finish. QE (@swadeley) could pass the boolean in IQE but could not yet watch completion from the API — you selected `distribution_path` from `snapshots`. Justin disabled Pulp in stage and prod before merge so the flag could land dark.

List APIs and a trigger endpoint followed (#308, #342, #371, #458). Soft-delete and “delete snapshots with the repo” (#340, #314) made lifecycle match Insights tenancy. Kafka inspection died in #390. By September the queue had won.

## Domains: isolation after the fact

backend#331 (@jlsherrill, 2023-08-04) created a Pulp domain per org on first snapshot, split the client into global (domain CRUD) and domain-scoped, and put a domain prefix on `repository_path`. Ephemeral Pulp needed `DOMAIN_ENABLED: "true"`. Compose switched to the team’s Pulp stack; S3/minio config landed in the same change (ACL follow-ups #355, #395).

Bulk-create of a new org raced domain create and returned 500s. Ryan: the domain still appeared and “no actual functionality is disrupted,” but the error was consistent. Justin treated concurrent create as something to swallow. @swadeley showed two orgs, two prefixes (`2d18751c` vs `12d903db`).

The important product fact: orgs that never snapshot never touch Pulp. Domain name is not `org_id`; mapping stays in Postgres. Pre-domain snapshots in stage were wrong enough to need repair jobs (#478–#492) in November. Lightwell import (#1559) still sits on this model — and immediately tripped org-id middleware written for a different tenancy story.

backend#438 (@rverdile, November) replaced a bug-prone “set domain client” helper from #414 with `WithContext(ctx).WithDomain(name)`. DAOs construct a client at injection time with `context.Background()` and rebind per request. Plumbing, but it is the client shape every later content type uses.

## Nightly sync and Red Hat as `org_id=-1`

A create-time snapshot is a day-one copy. backend#387 (2023-09-27) enqueued snapshot tasks at night for every `snapshot=true` repo not snapshotted in about 24 hours. `OPTIONS_ALWAYS_RUN_CRON_TASKS` ignored the window for debugging. Ryan asked whether a failed snapshot should retry in eight hours; v1 queued “older than ~20 hours” and left faster retry for later (HMS-5225 in 2025).

backend#407 (2023-10-10) loaded Red Hat CDN repos from JSON, one GPG key, `origin=red_hat`, `org_id=-1`. They are shared rows, not copied per customer. List hid them unless `origin=red_hat` or `NewRepositoryFiltering` was on — Ryan’s first curl returned both RH and custom because the flag was on. Sync used the CDN cert. #421 briefly reverted RH nightly snapshotting; it came back.

This is why `origin` exists, and why Lightwell did not pretend to be another `red_hat` catalog.

## tang: read path out of the API process

[tang](https://github.com/content-services/tang) started 2023-11-30 (@rverdile). tang#1 (merged 2024-01-03) exported `Tangy` with `RpmRepositoryVersionPackageSearch()` over Pulp hrefs, plus a mock in-tree because Justin did not want every app to invent one. tang is a Go module pointed at Pulp’s database, not an HTTP sidecar — which is why the backend later had to pass Clowder’s RDS CA (#629).

Errata (tang#6 / backend#608) and module streams (tang#14) are expansion-era uses of this split. Lightwell Maven/Python (tang#25/#26) are the same interface with a blank `group` for Python. Snapshots without tang are a distribution URL. With tang they are searchable.

## Who built it

@jlsherrill owned the snapshot task, integration, domains, nightly sync, and RH import. @rverdile owned tasking, snapshot list, the Pulp client refactor, and tang. @Andrewgdewar owned zest and the first snapshot UI (toggle, list modal). @swadeley made IQE tell the truth about `snapshot=true` and domains. @dpang314’s admin task API (#316) is how anyone saw the queue.

## Impact

By the end of 2023, content-sources was “Insights API + Postgres tasks + per-org Pulp + tang reads.” Templates, Candlepin, uploads, and Lightwell do not replace that. They enqueue more task types and ask tang new questions. The honest leftover: first snapshots were not domain-safe, first resume did not cover sync, and first RH list behavior depended on a flag nobody had documented in the testing steps.
