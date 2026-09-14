# Key Contributors

People who shaped Pulp integration, the Insights expansion years, or Lightwell. Bot accounts (Dependabot, Konflux, Snyk, RHTAP) excluded. Identities merged across work/personal emails. Confirm mappings in Phase 3 Prompt 5 before publishing.

## Pulp integration

| GitHub | Name | Role in this story |
|---|---|---|
| @jlsherrill | Justin Sherrill | Founder. Snapshot task (#248), integrate snapshots (#279), Pulp domains (#331), nightly sync (#387), RH snapshotting, zest/tang contributor |
| @rverdile | Ryan Verdile | Tasking system (#253), snapshot list API (#308), Pulp client refactor (#438), tang founder (2023-11-30) |
| @Andrewgdewar | Andrew Dewar | zest repo (2023-03-09) and binding automation; snapshot list/trigger APIs (#342, #458); early snapshot UI |
| @avisiedo | Alejandro Visiedo | Pre-Pulp: RPM search, Kafka introspect (#113), RBAC (#149). Background only |
| @dpang314 | David Pang | Admin/task APIs (#316) that operate the snapshot workers |
| @xbhouse | Bryttanie House | Snapshot info on the content list (#155); later template/snapshot APIs in expansion |

## Lightwell repos

| GitHub | Name | Role in this story |
|---|---|---|
| @xbhouse | Bryttanie House | Week-1 Network UI (#1059/#1061/#1066/#1069). Lens backend: coverage tables (#1639), S3 upload of manifests (#1644), analysis task (#1662), match status (#1680) |
| @rverdile | Ryan Verdile | Import (#1559), fetch, Maven details, Java/Python list UI (#1065/#1068), notifications send (#1633), OSV (#1609) |
| @arburka | | Connect snippets (#1088), week-1 polish, later Lens coverage UI (#1175) |
| @marusak | Matej Marusak | Week-1 chrome (#1063); later Beacon UI (#1168, SPUR #1199/#1214) |
| @ochosi | | Week-1 chrome/RBAC; Beacon PDF (#1184) |
| @jlsherrill | Justin Sherrill | Feature filter (#1557), count cache (#1558), package API (#1560), tang Maven list |
| @TenSt | Stepan Maksymchuk | Python/tang package APIs; Beacon clearinghouse API (#1640); Pulp catalog APIs in flight ([pulp_maven#451](https://github.com/pulp/pulp_maven/pull/451), [pulp_python#1359](https://github.com/pulp/pulp_python/pull/1359)) |
| @katarinazaprazna | Katarína Zápražná | Lens: coverage analyzer UI (#1166/#1183/#1186), backend parsers/matcher (#1645/#1655); notifications (#1149); chrome footer (#3679) |
| @dominikvagner | Dominik Vagner | Beacon ingestion (#1641); Lens SBOM/pom/purl (#1667/#1673/#1715), unlock Lens (#1725); chrome breadcrumbs (#3676) |
| @mattnolting | | Nav shell + Lens/Beacon stub (#1163) |
| @etsien | | JFrog advisory bridge (#1664, #1695), Network API phase 1 (#1643/#1666) |
| @swadeley | Steve Wadeley | Notification settings DB (#1601) |
| @florkbr | Bryan Florkiewicz | insights-chrome `/lightwell` scalprum route (#3582) |
| @karelhala | Karel Hala | Lightwell logo, clickable wordmark, logout (#3638/#3645/#3648) |
| @johnelliott626 | John Elliott | Insights templates/repos: columns/detail (#1132), upload-repo URLs (#1134), never-introspected status (#1657) |
| @Starle21 | Lenka Staronova | Insights UI (RepositoriesCard); npm openapi (#1586). |

## Origins

| GitHub | Name | Why they appear |
|---|---|---|
| @jlsherrill | Justin Sherrill | Backend first commits (2022-04-18) |
| @Andrewgdewar | Andrew Dewar | Frontend first commit (2022-05-31) |
| @rverdile | Ryan Verdile | Repository configuration model (#4), early DAO |
| @swadeley | Stephen Wadeley | QE |
| @mshriver | Mike Shriver | Clowder/IQE jobs — platform background |

## Expansion (Insights product)

| GitHub | Name | Role in this story |
|---|---|---|
| @jlsherrill | Justin Sherrill | Snapshot/Candlepin/template APIs continue from 2023 |
| @rverdile | Ryan Verdile | Templates (#510), snapshot APIs |
| @xbhouse | Bryttanie House | Template/snapshot APIs |
| @Andrewgdewar | Andrew Dewar | Frontend and zest through this era |
| @dominikvagner | Dominik Vagner | Templates UX, retention, systems |
| @katarinazaprazna | Katarína Zápražná | Template assignment, Candlepin client |
| @Starle21 | Lenka Staronova | Repositories and templates UI; resumable uploads; table refactor; backend openapi/deletion tasks. Git author `lstarono` |
| @marusak | Matej Marusak | Templates/repos UI through this era |
| @Dugowitch | Jakub Dugovič | Templates UI (wizard copy, assign modal) |
| @swadeley | Stephen Wadeley | QE |
| @mayurilahane | Mayuri Lahane | QE |
| @TenSt | Stepan Maksymchuk | Pulp/search work into 2026 |

## Identity notes (unverified)

- Justin Sherrill: `jlsherrill` — `jlsherrill@gmail.com` + `jsherril@redhat.com`
- Andrew Dewar: `Andrewgdewar` / `DewardianDev` / `adewar`
- Ryan Verdile: `rverdile`
- Bryttanie House: `xbhouse` / `bhouse@redhat.com`
- Stepan: `TenSt` / `smaksymc@redhat.com`
- Matej Marusak: `marusak` / `mmarusak@redhat.com`
- Katarína Zápražná: `katarinazaprazna` / `kzaprazn@redhat.com`
- Dominik Vagner: `dominikvagner`
- Bryan Florkiewicz: `florkbr` / `bflorkie@redhat.com`
- Karel Hala: `karelhala` / `khala@redhat.com`
- Lenka Staronova: GitHub `Starle21`; git author `lstarono` / `lstarono@redhat.com`. PRs are @Starle21 (branches often `lstarono/…` or `starle/…`)
- John Elliott: `johnelliott626` / `jelliott@redhat.com`
- Jakub Dugovič: `Dugowitch`
- Mayuri Lahane: `mayurilahane` / git `mlahane`

Use `@usernames` only in the published history.
