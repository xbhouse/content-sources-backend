---
type: pull_request
number: 40
title: "Update module github.com/go-sql-driver/mysql to v1.9.3"
state: merged
author: red-hat-konflux
created: 2025-11-05T20:18:34Z
updated: 2025-11-11T12:48:11Z
closed: 2025-11-11T12:48:11Z
merged: 2025-11-11T12:48:11Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-go-sql-driver-mysql-1.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/40
---

# Pull Request #40: Update module github.com/go-sql-driver/mysql to v1.9.3

**Author**: @red-hat-konflux
**Created**: November 05, 2025 at 08:18 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-go-sql-driver-mysql-1.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/go-sql-driver/mysql](https://redirect.github.com/go-sql-driver/mysql) | `v1.7.1` -> `v1.9.3` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fgo-sql-driver%2fmysql/v1.9.3?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fgo-sql-driver%2fmysql/v1.7.1/v1.9.3?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>go-sql-driver/mysql (github.com/go-sql-driver/mysql)</summary>

### [`v1.9.3`](https://redirect.github.com/go-sql-driver/mysql/releases/tag/v1.9.3)

[Compare Source](https://redirect.github.com/go-sql-driver/mysql/compare/v1.9.2...v1.9.3)

#### What's Changed

- \[1.9] test stability improvement. by [@&#8203;methane](https://redirect.github.com/methane) in [#&#8203;1699](https://redirect.github.com/go-sql-driver/mysql/pull/1699)
- \[1.9] Transaction Commit/Rollback returns conn's cached error by [@&#8203;methane](https://redirect.github.com/methane) in [#&#8203;1702](https://redirect.github.com/go-sql-driver/mysql/pull/1702)
- backport benchmark\_test by [@&#8203;methane](https://redirect.github.com/methane) in [#&#8203;1706](https://redirect.github.com/go-sql-driver/mysql/pull/1706)
- \[1.9] optimize readPacket ([#&#8203;1705](https://redirect.github.com/go-sql-driver/mysql/issues/1705)) by [@&#8203;methane](https://redirect.github.com/methane) in [#&#8203;1707](https://redirect.github.com/go-sql-driver/mysql/pull/1707)
- \[1.9] fix PING on compressed connections by [@&#8203;methane](https://redirect.github.com/methane) in [#&#8203;1723](https://redirect.github.com/go-sql-driver/mysql/pull/1723)
- release v1.9.3 by [@&#8203;methane](https://redirect.github.com/methane) in [#&#8203;1725](https://redirect.github.com/go-sql-driver/mysql/pull/1725)

**Full Changelog**: <https://github.com/go-sql-driver/mysql/compare/v1.9.2...v1.9.3>

### [`v1.9.2`](https://redirect.github.com/go-sql-driver/mysql/blob/HEAD/CHANGELOG.md#v192-2025-04-07)

[Compare Source](https://redirect.github.com/go-sql-driver/mysql/compare/v1.9.1...v1.9.2)

v1.9.2 is a re-release of v1.9.1 due to a release process issue; no changes were made to the content.

### [`v1.9.1`](https://redirect.github.com/go-sql-driver/mysql/blob/HEAD/CHANGELOG.md#v191-2025-03-21)

[Compare Source](https://redirect.github.com/go-sql-driver/mysql/compare/v1.9.0...v1.9.1)

##### Major Changes

- Add Charset() option. ([#&#8203;1679](https://redirect.github.com/go-sql-driver/mysql/issues/1679))

##### Bugfixes

- go.mod: fix go version format ([#&#8203;1682](https://redirect.github.com/go-sql-driver/mysql/issues/1682))
- Fix FormatDSN missing ConnectionAttributes ([#&#8203;1619](https://redirect.github.com/go-sql-driver/mysql/issues/1619))

### [`v1.9.0`](https://redirect.github.com/go-sql-driver/mysql/blob/HEAD/CHANGELOG.md#v190-2025-02-18)

[Compare Source](https://redirect.github.com/go-sql-driver/mysql/compare/v1.8.1...v1.9.0)

##### Major Changes

- Implement zlib compression. ([#&#8203;1487](https://redirect.github.com/go-sql-driver/mysql/issues/1487))
- Supported Go version is updated to Go 1.21+. ([#&#8203;1639](https://redirect.github.com/go-sql-driver/mysql/issues/1639))
- Add support for VECTOR type introduced in MySQL 9.0. ([#&#8203;1609](https://redirect.github.com/go-sql-driver/mysql/issues/1609))
- Config object can have custom dial function. ([#&#8203;1527](https://redirect.github.com/go-sql-driver/mysql/issues/1527))

##### Bugfixes

- Fix auth errors when username/password are too long. ([#&#8203;1625](https://redirect.github.com/go-sql-driver/mysql/issues/1625))
- Check if MySQL supports CLIENT\_CONNECT\_ATTRS before sending client attributes. ([#&#8203;1640](https://redirect.github.com/go-sql-driver/mysql/issues/1640))
- Fix auth switch request handling. ([#&#8203;1666](https://redirect.github.com/go-sql-driver/mysql/issues/1666))

##### Other changes

- Add "filename:line" prefix to log in go-mysql. Custom loggers now show it. ([#&#8203;1589](https://redirect.github.com/go-sql-driver/mysql/issues/1589))
- Improve error handling. It reduces the "busy buffer" errors. ([#&#8203;1595](https://redirect.github.com/go-sql-driver/mysql/issues/1595), [#&#8203;1601](https://redirect.github.com/go-sql-driver/mysql/issues/1601), [#&#8203;1641](https://redirect.github.com/go-sql-driver/mysql/issues/1641))
- Use `strconv.Atoi` to parse max\_allowed\_packet. ([#&#8203;1661](https://redirect.github.com/go-sql-driver/mysql/issues/1661))
- `rejectReadOnly` option now handles ER\_READ\_ONLY\_MODE (1290) error too. ([#&#8203;1660](https://redirect.github.com/go-sql-driver/mysql/issues/1660))

### [`v1.8.1`](https://redirect.github.com/go-sql-driver/mysql/releases/tag/v1.8.1)

[Compare Source](https://redirect.github.com/go-sql-driver/mysql/compare/v1.8.0...v1.8.1)

#### What's Changed

Bugfixes:

- fix race condition when context is canceled in [#&#8203;1562](https://redirect.github.com/go-sql-driver/mysql/pull/1562) and [#&#8203;1570](https://redirect.github.com/go-sql-driver/mysql/pull/1570)

**Full Changelog**: <https://github.com/go-sql-driver/mysql/compare/v1.8.0...v1.8.1>

### [`v1.8.0`](https://redirect.github.com/go-sql-driver/mysql/releases/tag/v1.8.0)

[Compare Source](https://redirect.github.com/go-sql-driver/mysql/compare/v1.7.1...v1.8.0)

#### What's Changed

#### Major changes

- Use `SET NAMES charset COLLATE collation`. by [@&#8203;methane](https://redirect.github.com/methane) in [#&#8203;1437](https://redirect.github.com/go-sql-driver/mysql/pull/1437)

  - Older go-mysql-driver used `collation_id` in the handshake packet. But it caused collation mismatch in some situation.
  - If you don't specify charset nor collation, go-mysql-driver sends `SET NAMES utf8mb4` for new connection. This uses server's default collation for utf8mb4.
  - If you specify charset, go-mysql-driver sends `SET NAMES <charset>`. This uses the server's default collation for `<charset>`.
  - If you specify collation and/or charset, go-mysql-driver sends `SET NAMES charset COLLATE collation`.

- PathEscape dbname in DSN. by [@&#8203;methane](https://redirect.github.com/methane) in [#&#8203;1432](https://redirect.github.com/go-sql-driver/mysql/pull/1432)

  - This is backward incompatible in rare case. Check your DSN.

- Drop Go 1.13-17 support by [@&#8203;methane](https://redirect.github.com/methane) in [#&#8203;1420](https://redirect.github.com/go-sql-driver/mysql/pull/1420)

  - Use Go 1.18+

- Parse numbers on text protocol too by [@&#8203;methane](https://redirect.github.com/methane) in [#&#8203;1452](https://redirect.github.com/go-sql-driver/mysql/pull/1452)

  - When text protocol is used, go-mysql-driver passed bare `[]byte` to database/sql for avoid unnecessary allocation and conversion.
  - If user specified `*any` to `Scan()`, database/sql passed the `[]byte` into the target variabe.
  - This confused users because most user doesn't know when text/binary protocol used.
  - go-mysql-driver 1.8 converts integer/float values into int64/double even in text protocol. This doesn't increase allocation compared to `[]byte` and conversion cost is negilible.

- New options start using the Functional Option Pattern to avoid increasing technical debt in the Config object. Future version may introduce Functional Option for existing options, but not for now.
  - Make TimeTruncate functional option by [@&#8203;methane](https://redirect.github.com/methane) in [#&#8203;1552](https://redirect.github.com/go-sql-driver/mysql/pull/1552)
  - Add BeforeConnect callback to configuration object by [@&#8203;ItalyPaleAle](https://redirect.github.com/ItalyPaleAle) in [#&#8203;1469](https://redirect.github.com/go-sql-driver/mysql/pull/1469)

##### Other changes

- Adding DeregisterDialContext to prevent memory leaks with dialers we don't need anymore by [@&#8203;jypelle](https://redirect.github.com/jypelle) in [#&#8203;1422](https://redirect.github.com/go-sql-driver/mysql/pull/1422)

- Make logger configurable per connection by [@&#8203;frozenbonito](https://redirect.github.com/frozenbonito) in [#&#8203;1408](https://redirect.github.com/go-sql-driver/mysql/pull/1408)

- Fix ColumnType.DatabaseTypeName for mediumint unsigned by [@&#8203;evanelias](https://redirect.github.com/evanelias) in [#&#8203;1428](https://redirect.github.com/go-sql-driver/mysql/pull/1428)

- Add connection attributes by [@&#8203;Daemonxiao](https://redirect.github.com/Daemonxiao) in [#&#8203;1389](https://redirect.github.com/go-sql-driver/mysql/pull/1389)

- Stop `ColumnTypeScanType()` from returning `sql.RawBytes` by [@&#8203;methane](https://redirect.github.com/methane) in [#&#8203;1424](https://redirect.github.com/go-sql-driver/mysql/pull/1424)

- Exec() now provides access to status of multiple statements. by [@&#8203;mherr-google](https://redirect.github.com/mherr-google) in [#&#8203;1309](https://redirect.github.com/go-sql-driver/mysql/pull/1309)

- Allow to change (or disable) the default driver name for registration by [@&#8203;dolmen](https://redirect.github.com/dolmen) in [#&#8203;1499](https://redirect.github.com/go-sql-driver/mysql/pull/1499)

- Add default connection attribute '\_server\_host' by [@&#8203;oblitorum](https://redirect.github.com/oblitorum) in [#&#8203;1506](https://redirect.github.com/go-sql-driver/mysql/pull/1506)

- QueryUnescape DSN ConnectionAttribute value by [@&#8203;zhangyangyu](https://redirect.github.com/zhangyangyu) in [#&#8203;1470](https://redirect.github.com/go-sql-driver/mysql/pull/1470)

- Add client\_ed25519 authentication by [@&#8203;Gusted](https://redirect.github.com/Gusted) in [#&#8203;1518](https://redirect.github.com/go-sql-driver/mysql/pull/1518)

- Reduced allocation on connection.go by [@&#8203;EPuncker](https://redirect.github.com/EPuncker) in [#&#8203;1421](https://redirect.github.com/go-sql-driver/mysql/pull/1421)

- Avoid panic in TestRowsColumnTypes by [@&#8203;wayyoungboy](https://redirect.github.com/wayyoungboy) in [#&#8203;1426](https://redirect.github.com/go-sql-driver/mysql/pull/1426)

- Add benchmark to receive massive rows. by [@&#8203;methane](https://redirect.github.com/methane) in [#&#8203;1415](https://redirect.github.com/go-sql-driver/mysql/pull/1415)

- README: Update multistatement by [@&#8203;methane](https://redirect.github.com/methane) in [#&#8203;1431](https://redirect.github.com/go-sql-driver/mysql/pull/1431)

- all: replace ioutil pkg to new package by [@&#8203;uji](https://redirect.github.com/uji) in [#&#8203;1438](https://redirect.github.com/go-sql-driver/mysql/pull/1438)

- chore: code optimization by [@&#8203;testwill](https://redirect.github.com/testwill) in [#&#8203;1439](https://redirect.github.com/go-sql-driver/mysql/pull/1439)

- Reduce map lookup in ColumnTypeDatabaseTypeName. by [@&#8203;methane](https://redirect.github.com/methane) in [#&#8203;1436](https://redirect.github.com/go-sql-driver/mysql/pull/1436)

- doc: add link to NewConnector from FormatDSN by [@&#8203;dolmen](https://redirect.github.com/dolmen) in [#&#8203;1442](https://redirect.github.com/go-sql-driver/mysql/pull/1442)

- Add fuzz test for ParseDSN / FormatDSN roundtrip by [@&#8203;dolmen](https://redirect.github.com/dolmen) in [#&#8203;1444](https://redirect.github.com/go-sql-driver/mysql/pull/1444)

- TestDSNReformat: add more roundtrip checks by [@&#8203;dolmen](https://redirect.github.com/dolmen) in [#&#8203;1443](https://redirect.github.com/go-sql-driver/mysql/pull/1443)

- tcp: handle errors returned by SetKeepAlive by [@&#8203;achille-roussel](https://redirect.github.com/achille-roussel) in [#&#8203;1448](https://redirect.github.com/go-sql-driver/mysql/pull/1448)

- use staticcheck by [@&#8203;methane](https://redirect.github.com/methane) in [#&#8203;1449](https://redirect.github.com/go-sql-driver/mysql/pull/1449)

- Add Daemonxiao to AUTHORS by [@&#8203;Daemonxiao](https://redirect.github.com/Daemonxiao) in [#&#8203;1459](https://redirect.github.com/go-sql-driver/mysql/pull/1459)

- Update link about `LOAD DATA LOCAL` in README.md by [@&#8203;i7a7467](https://redirect.github.com/i7a7467) in [#&#8203;1468](https://redirect.github.com/go-sql-driver/mysql/pull/1468)

- Update README.md by [@&#8203;Netzer7](https://redirect.github.com/Netzer7) in [#&#8203;1464](https://redirect.github.com/go-sql-driver/mysql/pull/1464)

- add Go 1.21 and MySQL 8.1 to the build matrix by [@&#8203;shogo82148](https://redirect.github.com/shogo82148) in [#&#8203;1472](https://redirect.github.com/go-sql-driver/mysql/pull/1472)

- Improve DSN docstsrings by [@&#8203;golddranks](https://redirect.github.com/golddranks) in [#&#8203;1475](https://redirect.github.com/go-sql-driver/mysql/pull/1475)

- Fix [#&#8203;1478](https://redirect.github.com/go-sql-driver/mysql/issues/1478) remove length check by [@&#8203;ShenFeng312](https://redirect.github.com/ShenFeng312) in [#&#8203;1481](https://redirect.github.com/go-sql-driver/mysql/pull/1481)

- README: fix markup error by [@&#8203;methane](https://redirect.github.com/methane) in [#&#8203;1480](https://redirect.github.com/go-sql-driver/mysql/pull/1480)

- Close connection on ErrPktSync and ErrPktSyncMul by [@&#8203;owbone](https://redirect.github.com/owbone) in [#&#8203;1473](https://redirect.github.com/go-sql-driver/mysql/pull/1473)

- Spelling, grammar, and link fixes by [@&#8203;scop](https://redirect.github.com/scop) in [#&#8203;1485](https://redirect.github.com/go-sql-driver/mysql/pull/1485)

- Make use of strings.Cut by [@&#8203;scop](https://redirect.github.com/scop) in [#&#8203;1486](https://redirect.github.com/go-sql-driver/mysql/pull/1486)

- move stale connection check to ResetSession() by [@&#8203;methane](https://redirect.github.com/methane) in [#&#8203;1496](https://redirect.github.com/go-sql-driver/mysql/pull/1496)

- fix race condition of TestConcurrent by [@&#8203;shogo82148](https://redirect.github.com/shogo82148) in [#&#8203;1490](https://redirect.github.com/go-sql-driver/mysql/pull/1490)

- mark fail, mustExec and mustQuery as test helpers by [@&#8203;shogo82148](https://redirect.github.com/shogo82148) in [#&#8203;1488](https://redirect.github.com/go-sql-driver/mysql/pull/1488)

- Remove obsolete fuzz.go [#&#8203;1445](https://redirect.github.com/go-sql-driver/mysql/issues/1445) by [@&#8203;dolmen](https://redirect.github.com/dolmen) in [#&#8203;1498](https://redirect.github.com/go-sql-driver/mysql/pull/1498)

- testing: expose testing.TB in DBTest instead of full \*testing.T by [@&#8203;dolmen](https://redirect.github.com/dolmen) in [#&#8203;1500](https://redirect.github.com/go-sql-driver/mysql/pull/1500)

- symbol removed from installation command by [@&#8203;panvalkar1994](https://redirect.github.com/panvalkar1994) in [#&#8203;1510](https://redirect.github.com/go-sql-driver/mysql/pull/1510)

- fix issue 1361 by [@&#8203;keeplearning20221](https://redirect.github.com/keeplearning20221) in [#&#8203;1462](https://redirect.github.com/go-sql-driver/mysql/pull/1462)

- fix fragile test by [@&#8203;methane](https://redirect.github.com/methane) in [#&#8203;1522](https://redirect.github.com/go-sql-driver/mysql/pull/1522)

- Fix sql.RawBytes corruption issue by [@&#8203;shogo82148](https://redirect.github.com/shogo82148) in [#&#8203;1523](https://redirect.github.com/go-sql-driver/mysql/pull/1523)

- fix for enum and set field type to column type identifying by [@&#8203;jennifersp](https://redirect.github.com/jennifersp) in [#&#8203;1520](https://redirect.github.com/go-sql-driver/mysql/pull/1520)

- Parallelize test by [@&#8203;shogo82148](https://redirect.github.com/shogo82148) in [#&#8203;1525](https://redirect.github.com/go-sql-driver/mysql/pull/1525)

- Fix unsigned int overflow by [@&#8203;shiyuhang0](https://redirect.github.com/shiyuhang0) in [#&#8203;1530](https://redirect.github.com/go-sql-driver/mysql/pull/1530)

- Introduce `timeTruncate` parameter for `time.Time` arguments by [@&#8203;PauliusLozys](https://redirect.github.com/PauliusLozys) in [#&#8203;1541](https://redirect.github.com/go-sql-driver/mysql/pull/1541)

- add TiDB support in README.md by [@&#8203;crazycs520](https://redirect.github.com/crazycs520) in [#&#8203;1333](https://redirect.github.com/go-sql-driver/mysql/pull/1333)

- Update workflows by [@&#8203;methane](https://redirect.github.com/methane) in [#&#8203;1547](https://redirect.github.com/go-sql-driver/mysql/pull/1547)

#### New Contributors

- [@&#8203;EPuncker](https://redirect.github.com/EPuncker) made their first contribution in [#&#8203;1421](https://redirect.github.com/go-sql-driver/mysql/pull/1421)
- [@&#8203;jypelle](https://redirect.github.com/jypelle) made their first contribution in [#&#8203;1422](https://redirect.github.com/go-sql-driver/mysql/pull/1422)
- [@&#8203;frozenbonito](https://redirect.github.com/frozenbonito) made their first contribution in [#&#8203;1408](https://redirect.github.com/go-sql-driver/mysql/pull/1408)
- [@&#8203;wayyoungboy](https://redirect.github.com/wayyoungboy) made their first contribution in [#&#8203;1426](https://redirect.github.com/go-sql-driver/mysql/pull/1426)
- [@&#8203;evanelias](https://redirect.github.com/evanelias) made their first contribution in [#&#8203;1428](https://redirect.github.com/go-sql-driver/mysql/pull/1428)
- [@&#8203;Daemonxiao](https://redirect.github.com/Daemonxiao) made their first contribution in [#&#8203;1389](https://redirect.github.com/go-sql-driver/mysql/pull/1389)
- [@&#8203;uji](https://redirect.github.com/uji) made their first contribution in [#&#8203;1438](https://redirect.github.com/go-sql-driver/mysql/pull/1438)
- [@&#8203;testwill](https://redirect.github.com/testwill) made their first contribution in [#&#8203;1439](https://redirect.github.com/go-sql-driver/mysql/pull/1439)
- [@&#8203;i7a7467](https://redirect.github.com/i7a7467) made their first contribution in [#&#8203;1468](https://redirect.github.com/go-sql-driver/mysql/pull/1468)
- [@&#8203;Netzer7](https://redirect.github.com/Netzer7) made their first contribution in [#&#8203;1464](https://redirect.github.com/go-sql-driver/mysql/pull/1464)
- [@&#8203;golddranks](https://redirect.github.com/golddranks) made their first contribution in [#&#8203;1475](https://redirect.github.com/go-sql-driver/mysql/pull/1475)
- [@&#8203;ShenFeng312](https://redirect.github.com/ShenFeng312) made their first contribution in [#&#8203;1481](https://redirect.github.com/go-sql-driver/mysql/pull/1481)
- [@&#8203;owbone](https://redirect.github.com/owbone) made their first contribution in [#&#8203;1473](https://redirect.github.com/go-sql-driver/mysql/pull/1473)
- [@&#8203;scop](https://redirect.github.com/scop) made their first contribution in [#&#8203;1485](https://redirect.github.com/go-sql-driver/mysql/pull/1485)
- [@&#8203;panvalkar1994](https://redirect.github.com/panvalkar1994) made their first contribution in [#&#8203;1510](https://redirect.github.com/go-sql-driver/mysql/pull/1510)
- [@&#8203;zhangyangyu](https://redirect.github.com/zhangyangyu) made their first contribution in [#&#8203;1470](https://redirect.github.com/go-sql-driver/mysql/pull/1470)
- [@&#8203;keeplearning20221](https://redirect.github.com/keeplearning20221) made their first contribution in [#&#8203;1462](https://redirect.github.com/go-sql-driver/mysql/pull/1462)
- [@&#8203;oblitorum](https://redirect.github.com/oblitorum) made their first contribution in [#&#8203;1506](https://redirect.github.com/go-sql-driver/mysql/pull/1506)
- [@&#8203;Gusted](https://redirect.github.com/Gusted) made their first contribution in [#&#8203;1518](https://redirect.github.com/go-sql-driver/mysql/pull/1518)
- [@&#8203;jennifersp](https://redirect.github.com/jennifersp) made their first contribution in [#&#8203;1520](https://redirect.github.com/go-sql-driver/mysql/pull/1520)
- [@&#8203;shiyuhang0](https://redirect.github.com/shiyuhang0) made their first contribution in [#&#8203;1530](https://redirect.github.com/go-sql-driver/mysql/pull/1530)
- [@&#8203;PauliusLozys](https://redirect.github.com/PauliusLozys) made their first contribution in [#&#8203;1541](https://redirect.github.com/go-sql-driver/mysql/pull/1541)
- [@&#8203;crazycs520](https://redirect.github.com/crazycs520) made their first contribution in [#&#8203;1333](https://redirect.github.com/go-sql-driver/mysql/pull/1333)
- [@&#8203;ItalyPaleAle](https://redirect.github.com/ItalyPaleAle) made their first contribution in [#&#8203;1469](https://redirect.github.com/go-sql-driver/mysql/pull/1469)

**Full Changelog**: <https://github.com/go-sql-driver/mysql/compare/v1.7.1...v1.8.0>

</details>

---

### Configuration

📅 **Schedule**: Branch creation - At any time (no schedule defined), Automerge - At any time (no schedule defined).

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

---
### Documentation

Find out how to configure dependency updates in [MintMaker documentation](https://konflux-ci.dev/docs/mintmaker/user/) or see all available configuration options in [Renovate documentation](https://docs.renovatebot.com/configuration-options/).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6W119-->


---

## Reviews

### Review by @Hyperkid123 - Approved on November 11, 2025 at 12:48 PM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/40*
