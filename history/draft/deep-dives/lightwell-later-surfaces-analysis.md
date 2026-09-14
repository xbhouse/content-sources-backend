# After week one — notifications, Lens, Beacon, JFrog, Network API

Period: lightwell  
Contrast: Week one (7/1–7/8) was the Network catalog UI plus the insights-chrome `/lightwell` shell. These PRs add side capabilities on the same Pulp-backed repos.

Insights partner-repos (HMS-107xx, backend#1596) are not in this list.

## Notifications

| PR | Author | What |
|---|---|---|
| backend#1601 | @swadeley | Notification settings DB (HMS-11040) |
| backend#1609 | @rverdile | OSV advisory sync (HMS-11042) |
| backend#1611 / #1625 | @rverdile | Admin test send; new notification format |
| frontend#1149 | @katarinazaprazna | Notification preferences modal |
| backend#1633 | @rverdile | Send Lightwell notifications (HMS-11043) |
| frontend#1160 | @rverdile | Severity selection (LWLP-27) |
| backend#1651 / frontend#1169 | @rverdile | Feature flag (LWLP-543) |
| backend#1676 / #1678 | @rverdile | Per-org dedupe; predisclosure repo |
| frontend#1194 | @katarinazaprazna | Toast notifications (LWLP-820) |

## Lens (coverage)

| PR | Author | What |
|---|---|---|
| backend#1639 | @xbhouse | Coverage report tables and models (LWLP-13) |
| backend#1644 | @xbhouse | Upload SBOM/manifest to S3 on POST (LWLP-589) |
| backend#1662 | @xbhouse | Task that runs coverage analysis (LWLP-18) |
| backend#1680 | @xbhouse | Package match status on API (LWLP-20) |
| backend#1649 / #1677 / frontend#1189 | @xbhouse | Lens/Beacon feature flags |
| frontend#1187 / #1192 / #1206 / #1216 | @xbhouse | Upload size, analysis progress, uploader, report route |
| frontend#1166 / #1173 | @katarinazaprazna | Coverage analyzer page; file size + charts (LWLP-16/17) |
| frontend#1183 / #1188 / #1186 | @katarinazaprazna | Lens polish; covered-package list |
| frontend#1191 / #1221 | @katarinazaprazna | SPUR (components, brand palette) |
| frontend#1193 / #1200 / #1203 | @katarinazaprazna | Manifest format / validation / size limits |
| backend#1645 / #1655 | @katarinazaprazna | CSV + requirements.txt parsers; package matcher (LWLP-15/14) |
| backend#1667 / #1673 | @dominikvagner | SBOM and pom.xml coverage parsers (LWLP-586/587) |
| backend#1711 / #1715 / #1725 | @dominikvagner | Purl fallback/parsing; unlock Lens for internal employees |
| frontend#1175 | @arburka | Coverage Analyzer redesign |
| frontend#1163 | @mattnolting | Nav shell (shared with Beacon) |

## Beacon (vulnerabilities)

| PR | Author | What |
|---|---|---|
| frontend#1163 | @mattnolting | Nav shell + Beacon stub (2026-08-17) |
| frontend#1168 / #1170 / #1172 | @marusak | Beacon UI (LWLP-530) |
| backend#1640 | @TenSt | Clearinghouse / Beacon API layer (LWLP-36) |
| backend#1641 / #1705 | @dominikvagner | Vulnerability ingestion job; vulns seed (LWLP-40) |
| backend#1652 / #1654 | @TenSt | Gate Beacon; flag defaults |
| frontend#1184 | @ochosi | Beacon PDF (LWLP-755) |
| frontend#1199 / #1214 | @marusak | Beacon SPUR updates |

## JFrog integration

See [pr-1664-analysis.md](pr-1664-analysis.md). #1664, #1695. Pulp remains SoT for versions/builds.

## Network API, phase 1

| PR | Author | What |
|---|---|---|
| backend#1643 | @etsien | LWLP-5: advisories API, cross-repo package search, response shape |
| backend#1666 | @etsien | LWLP-5: Lightwell Network API (named first phase) |
