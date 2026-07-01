---
type: pull_request
number: 1649
title: "Update module gorm.io/gorm to v1.26.0"
state: closed
author: red-hat-konflux
created: 2025-04-27T16:37:41Z
updated: 2025-05-27T13:08:32Z
closed: 2025-05-27T13:08:32Z
base_branch: master
head_branch: konflux/mintmaker/master/gorm.io-gorm-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1649
---

# Pull Request #1649: Update module gorm.io/gorm to v1.26.0

**Author**: @red-hat-konflux
**Created**: April 27, 2025 at 04:37 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/gorm.io-gorm-1.x`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [gorm.io/gorm](https://redirect.github.com/go-gorm/gorm) | require | minor | `v1.25.12` -> `v1.26.0` |

---

### Release Notes

<details>
<summary>go-gorm/gorm (gorm.io/gorm)</summary>

### [`v1.26.0`](https://redirect.github.com/go-gorm/gorm/releases/tag/v1.26.0)

[Compare Source](https://redirect.github.com/go-gorm/gorm/compare/v1.25.12...v1.26.0)

#### Changes

-   Preparestmt use LRU Map instead default map [@&#8203;Yaccc](https://redirect.github.com/Yaccc) ([#&#8203;7435](https://redirect.github.com/go-gorm/gorm/issues/7435))
-   use golangci replace reviewdog [@&#8203;jinzhu](https://redirect.github.com/jinzhu) ([#&#8203;7426](https://redirect.github.com/go-gorm/gorm/issues/7426))
-   test: mssql ci [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;7388](https://redirect.github.com/go-gorm/gorm/issues/7388))
-   fix deprecated reflect.PtrTo reflect.PointerTo usage [@&#8203;Aman-Shitta](https://redirect.github.com/Aman-Shitta) ([#&#8203;7366](https://redirect.github.com/go-gorm/gorm/issues/7366))
-   Fix concurrent map writes [#&#8203;7297](https://redirect.github.com/go-gorm/gorm/issues/7297) [@&#8203;Ponywka](https://redirect.github.com/Ponywka) ([#&#8203;7298](https://redirect.github.com/go-gorm/gorm/issues/7298))
-   chore: update copyright year [@&#8203;maxktz](https://redirect.github.com/maxktz) ([#&#8203;7332](https://redirect.github.com/go-gorm/gorm/issues/7332))
-   Enhance db.Scan with ParamsFilter - Issue 7336 - Suggestion [@&#8203;evyaffe](https://redirect.github.com/evyaffe) ([#&#8203;7337](https://redirect.github.com/go-gorm/gorm/issues/7337))
-   Fixed Empty Returning Clause Merge Edge Case [@&#8203;aviyam181199](https://redirect.github.com/aviyam181199) ([#&#8203;7339](https://redirect.github.com/go-gorm/gorm/issues/7339))
-   feat:Capitalize the priority field of IndexOption（[https://github.com/go-gorm/gorm/issues/7331](https://redirect.github.com/go-gorm/gorm/issues/7331)） [@&#8203;nowindexman](https://redirect.github.com/nowindexman) ([#&#8203;7342](https://redirect.github.com/go-gorm/gorm/issues/7342))
-   fix: deterministic index ordering when migrating [@&#8203;bamo](https://redirect.github.com/bamo) ([#&#8203;7208](https://redirect.github.com/go-gorm/gorm/issues/7208))
-   use map look-up for indexes [@&#8203;abbyssoul](https://redirect.github.com/abbyssoul) ([#&#8203;7242](https://redirect.github.com/go-gorm/gorm/issues/7242))
-   \[[#&#8203;6372](https://redirect.github.com/go-gorm/gorm/issues/6372)] Fixed nullable constraint bug for columns during auto migration [@&#8203;wookie0](https://redirect.github.com/wookie0) ([#&#8203;7269](https://redirect.github.com/go-gorm/gorm/issues/7269))
-   Create CODE_OF_CONDUCT.md [@&#8203;omidfth](https://redirect.github.com/omidfth) ([#&#8203;7240](https://redirect.github.com/go-gorm/gorm/issues/7240))
-   Enhance Release Process: Implement Automated Release Management and Notes Generation [@&#8203;YidiDev](https://redirect.github.com/YidiDev) ([#&#8203;7224](https://redirect.github.com/go-gorm/gorm/issues/7224))
-   refactor: improve logging for unimplemented ErrorTranslator in TranlateError config enabled [@&#8203;Invidam](https://redirect.github.com/Invidam) ([#&#8203;7225](https://redirect.github.com/go-gorm/gorm/issues/7225))
-   Add GitHub Actions workflow to automate release creation on tagged pushes [@&#8203;YidiDev](https://redirect.github.com/YidiDev) ([#&#8203;7209](https://redirect.github.com/go-gorm/gorm/issues/7209))
-   Use official SQL Server docker image for tests [@&#8203;omkar-foss](https://redirect.github.com/omkar-foss) ([#&#8203;7205](https://redirect.github.com/go-gorm/gorm/issues/7205))
-   feat: Modernize Docker Compose File [@&#8203;isso-719](https://redirect.github.com/isso-719) ([#&#8203;7086](https://redirect.github.com/go-gorm/gorm/issues/7086))
-   Generate unique savepoint names for nested transactions [@&#8203;phroggyy](https://redirect.github.com/phroggyy) ([#&#8203;7174](https://redirect.github.com/go-gorm/gorm/issues/7174))
-   fix: AfterQuery using safer right trim while clearing from clause's j… [@&#8203;bhowmik-abhijeet](https://redirect.github.com/bhowmik-abhijeet) ([#&#8203;7153](https://redirect.github.com/go-gorm/gorm/issues/7153))
-   fix memory leaks in PrepareStatementDB [@&#8203;ivila](https://redirect.github.com/ivila) ([#&#8203;7142](https://redirect.github.com/go-gorm/gorm/issues/7142))
-   ci: Add PostgreSQL 14 and 15 to GitHub Actions matrix [@&#8203;enomotodev](https://redirect.github.com/enomotodev) ([#&#8203;7081](https://redirect.github.com/go-gorm/gorm/issues/7081))
-   feat: add MapColumns method [@&#8203;molon](https://redirect.github.com/molon) ([#&#8203;6901](https://redirect.github.com/go-gorm/gorm/issues/6901))
-   add DB level propagation for the Unscoped flag [@&#8203;sprataa](https://redirect.github.com/sprataa) ([#&#8203;7007](https://redirect.github.com/go-gorm/gorm/issues/7007))
-   fix(scan): update Scan function to reset structs to zero values for each scan [@&#8203;Waldeedle](https://redirect.github.com/Waldeedle) ([#&#8203;7061](https://redirect.github.com/go-gorm/gorm/issues/7061))
-   fix: use reflect.Append when preloading nested associations instead of making a slice with fixed size [@&#8203;emilienkofman](https://redirect.github.com/emilienkofman) ([#&#8203;7014](https://redirect.github.com/go-gorm/gorm/issues/7014))
-   Fix association replace non-addressable panic [@&#8203;SergeiSadov](https://redirect.github.com/SergeiSadov) ([#&#8203;7012](https://redirect.github.com/go-gorm/gorm/issues/7012))
-   fix: `unsupported data` on nested joins with preloads [@&#8203;N-Schaef](https://redirect.github.com/N-Schaef) ([#&#8203;6957](https://redirect.github.com/go-gorm/gorm/issues/6957))
-   fix: AfterQuery should clear FROM Clause's Joins rather than the Stat… [@&#8203;liov](https://redirect.github.com/liov) ([#&#8203;7027](https://redirect.github.com/go-gorm/gorm/issues/7027))
-   feat: chainable order support clause.OrderBy [@&#8203;supergem3000](https://redirect.github.com/supergem3000) ([#&#8203;7054](https://redirect.github.com/go-gorm/gorm/issues/7054))
-   fix: strings.Title -> cases.Title bcs strings.Title library is deprecated [@&#8203;ryuji-cre8ive](https://redirect.github.com/ryuji-cre8ive) ([#&#8203;6999](https://redirect.github.com/go-gorm/gorm/issues/6999))
-   fix: typo [@&#8203;hakusai22](https://redirect.github.com/hakusai22) ([#&#8203;7003](https://redirect.github.com/go-gorm/gorm/issues/7003))
-   Fix handling of unknown column types [@&#8203;looi](https://redirect.github.com/looi) ([#&#8203;6540](https://redirect.github.com/go-gorm/gorm/issues/6540))
-   Fix panic bug in migrator due to lack of nil check for stmt.Schema [@&#8203;pixelmaxQm](https://redirect.github.com/pixelmaxQm) ([#&#8203;6932](https://redirect.github.com/go-gorm/gorm/issues/6932))
-   Add new error for "Violation Check Constraint" [@&#8203;anilsenay](https://redirect.github.com/anilsenay) ([#&#8203;6992](https://redirect.github.com/go-gorm/gorm/issues/6992))
-   fix: not clause with or condition [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;6984](https://redirect.github.com/go-gorm/gorm/issues/6984))
-   perf: merge nested preload query when using join [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;6990](https://redirect.github.com/go-gorm/gorm/issues/6990))
-   Faster utils.FileWithLineNum [@&#8203;kkocdko](https://redirect.github.com/kkocdko) ([#&#8203;6981](https://redirect.github.com/go-gorm/gorm/issues/6981))
-   Added comment describing Unscoped() method [@&#8203;AntonyChR](https://redirect.github.com/AntonyChR) ([#&#8203;6969](https://redirect.github.com/go-gorm/gorm/issues/6969))
-   fix: duplicated preload [@&#8203;yetone](https://redirect.github.com/yetone) ([#&#8203;6948](https://redirect.github.com/go-gorm/gorm/issues/6948))
-   feat: prepare_stmt support ping [@&#8203;philhuan](https://redirect.github.com/philhuan) ([#&#8203;6924](https://redirect.github.com/go-gorm/gorm/issues/6924))
-   fix: remove `callback` from `callbacks` if `Remove()` called [@&#8203;snackmgmg](https://redirect.github.com/snackmgmg) ([#&#8203;6916](https://redirect.github.com/go-gorm/gorm/issues/6916))
-   fix: 'type XXXX int' will print wrong sql to terminal [@&#8203;wangzeping722](https://redirect.github.com/wangzeping722) ([#&#8203;6917](https://redirect.github.com/go-gorm/gorm/issues/6917))
-   chore: optimize `regEnLetterAndMidline` regular [@&#8203;demoManito](https://redirect.github.com/demoManito) ([#&#8203;6908](https://redirect.github.com/go-gorm/gorm/issues/6908))
-   fix(create): fix insert column order [@&#8203;archever](https://redirect.github.com/archever) ([#&#8203;6855](https://redirect.github.com/go-gorm/gorm/issues/6855))
-   chore: optimize conversion of nanoseconds to milliseconds [@&#8203;demoManito](https://redirect.github.com/demoManito) ([#&#8203;6907](https://redirect.github.com/go-gorm/gorm/issues/6907))
-   fix(scan): array element is set to a zero value [@&#8203;demoManito](https://redirect.github.com/demoManito) ([#&#8203;6890](https://redirect.github.com/go-gorm/gorm/issues/6890))
-   fix: nested preload with join panic when find [@&#8203;black-06](https://redirect.github.com/black-06) ([#&#8203;6877](https://redirect.github.com/go-gorm/gorm/issues/6877))
-   fix(scan.go): reflect.MakeSlice may pass in the wrong type reflect.Array [@&#8203;demoManito](https://redirect.github.com/demoManito) ([#&#8203;6880](https://redirect.github.com/go-gorm/gorm/issues/6880))
-   test: namer identifier lenght [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;6872](https://redirect.github.com/go-gorm/gorm/issues/6872))
-   fix some typos in tests [@&#8203;hishope](https://redirect.github.com/hishope) ([#&#8203;6879](https://redirect.github.com/go-gorm/gorm/issues/6879))
-   Fix regression in db.Not introduced in v1.25.6. [@&#8203;tsuba3](https://redirect.github.com/tsuba3) ([#&#8203;6844](https://redirect.github.com/go-gorm/gorm/issues/6844))
-   Add unittest test helper function ConvertSliceOfMapToValuesForCreate [@&#8203;naruchet](https://redirect.github.com/naruchet) ([#&#8203;6854](https://redirect.github.com/go-gorm/gorm/issues/6854))
-   CHORE add unittest test function  ConvertMapToValueForCreate [@&#8203;naruchet](https://redirect.github.com/naruchet) ([#&#8203;6846](https://redirect.github.com/go-gorm/gorm/issues/6846))
-   Fix: panic on nullable value with multiple foreign key usage [@&#8203;shtrih](https://redirect.github.com/shtrih) ([#&#8203;6839](https://redirect.github.com/go-gorm/gorm/issues/6839))
-   refactor: part 2 of distinguish between Unique and UniqueIndex [@&#8203;black-06](https://redirect.github.com/black-06) ([#&#8203;6822](https://redirect.github.com/go-gorm/gorm/issues/6822))
-   let limit and offset use bind parameter [@&#8203;jasonchuan](https://redirect.github.com/jasonchuan) ([#&#8203;6806](https://redirect.github.com/go-gorm/gorm/issues/6806))
-   refactor: distinguish between Unique and UniqueIndex [@&#8203;black-06](https://redirect.github.com/black-06) ([#&#8203;6386](https://redirect.github.com/go-gorm/gorm/issues/6386))
-   fix: preload shouldn't overwrite the value of join [@&#8203;black-06](https://redirect.github.com/black-06) ([#&#8203;6771](https://redirect.github.com/go-gorm/gorm/issues/6771))
-   chore(deps): bump actions/cache from 3 to 4 @&#8203;[dependabot\[bot\]](https://redirect.github.com/apps/dependabot) ([#&#8203;6802](https://redirect.github.com/go-gorm/gorm/issues/6802))
-   fix: ignore .gen.go suffix in logger to get the real caller when using gen [@&#8203;0xJacky](https://redirect.github.com/0xJacky) ([#&#8203;6785](https://redirect.github.com/go-gorm/gorm/issues/6785))
-   fix: ExplainSQL using consecutive pairs of escaper in SQL string represents an escaper [@&#8203;iTanken](https://redirect.github.com/iTanken) ([#&#8203;6766](https://redirect.github.com/go-gorm/gorm/issues/6766))
-   fix: join and select mytable.\* not working [@&#8203;StephanoGeorge](https://redirect.github.com/StephanoGeorge) ([#&#8203;6761](https://redirect.github.com/go-gorm/gorm/issues/6761))
-   feature: bring custom type and id column name to polymorphism [@&#8203;alexisvisco](https://redirect.github.com/alexisvisco) ([#&#8203;6716](https://redirect.github.com/go-gorm/gorm/issues/6716))
-   Making locking parameters more intuitive [@&#8203;dogenkigen](https://redirect.github.com/dogenkigen) ([#&#8203;6719](https://redirect.github.com/go-gorm/gorm/issues/6719))
-   refactor: Resolve implicit memory aliasing in for loop [@&#8203;BugKillerPro](https://redirect.github.com/BugKillerPro) ([#&#8203;6730](https://redirect.github.com/go-gorm/gorm/issues/6730))
-   map insert support return increment id [@&#8203;FangSqing](https://redirect.github.com/FangSqing) ([#&#8203;6662](https://redirect.github.com/go-gorm/gorm/issues/6662))
-   docs: fix broken link [@&#8203;kijimaD](https://redirect.github.com/kijimaD) ([#&#8203;6673](https://redirect.github.com/go-gorm/gorm/issues/6673))
-   chore(logger): optimize [@&#8203;flc1125](https://redirect.github.com/flc1125) ([#&#8203;6675](https://redirect.github.com/go-gorm/gorm/issues/6675))
-   feat: add MigrateColumnUnique [@&#8203;black-06](https://redirect.github.com/black-06) ([#&#8203;6640](https://redirect.github.com/go-gorm/gorm/issues/6640))
-   test: fix TestEmbeddedRelations [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;6639](https://redirect.github.com/go-gorm/gorm/issues/6639))
-   tests for sqilte: enable FOREIGN_KEYS inside OpenTestConnection [@&#8203;glebarez](https://redirect.github.com/glebarez) ([#&#8203;6641](https://redirect.github.com/go-gorm/gorm/issues/6641))
-   Add support for returning in sqlserver [@&#8203;FrancoLiberali](https://redirect.github.com/FrancoLiberali) ([#&#8203;6585](https://redirect.github.com/go-gorm/gorm/issues/6585))
-   chore(deps): bump actions/checkout from 3 to 4 @&#8203;[dependabot\[bot\]](https://redirect.github.com/apps/dependabot) ([#&#8203;6586](https://redirect.github.com/go-gorm/gorm/issues/6586))
-   Don't call MethodByName with a variable arg [@&#8203;jquirke](https://redirect.github.com/jquirke) ([#&#8203;6602](https://redirect.github.com/go-gorm/gorm/issues/6602))
-   feat: reuse name [@&#8203;philhuan](https://redirect.github.com/philhuan) ([#&#8203;6626](https://redirect.github.com/go-gorm/gorm/issues/6626))
-   fix: sqlite dialector cannot apply `PRIMARY KEY AUTOINCREMENT` type [@&#8203;samuelncui](https://redirect.github.com/samuelncui) ([#&#8203;6624](https://redirect.github.com/go-gorm/gorm/issues/6624))
-   Fixed error message when dialector fails to initialize [@&#8203;RatajVaver](https://redirect.github.com/RatajVaver) ([#&#8203;6509](https://redirect.github.com/go-gorm/gorm/issues/6509))
-   fix schema GetIdentityFieldValuesMap interface or ptr [@&#8203;uptutu](https://redirect.github.com/uptutu) ([#&#8203;6417](https://redirect.github.com/go-gorm/gorm/issues/6417))
-   add float32 test case for  keep float precision in ExplainSQL [@&#8203;Heliner](https://redirect.github.com/Heliner) ([#&#8203;6530](https://redirect.github.com/go-gorm/gorm/issues/6530))
-   feat: rm GetDBConnWithContext method [@&#8203;qqxhb](https://redirect.github.com/qqxhb) ([#&#8203;6535](https://redirect.github.com/go-gorm/gorm/issues/6535))
-   fix(clause): clause equal empty array [@&#8203;whcao](https://redirect.github.com/whcao) ([#&#8203;6503](https://redirect.github.com/go-gorm/gorm/issues/6503))
-   refactor: Regex description [@&#8203;fayvori](https://redirect.github.com/fayvori) ([#&#8203;6507](https://redirect.github.com/go-gorm/gorm/issues/6507))
-   fix: rectify `SkipHooks` not working with `NewDB` set in Session method [@&#8203;aayushacharya](https://redirect.github.com/aayushacharya) ([#&#8203;6484](https://redirect.github.com/go-gorm/gorm/issues/6484))
-   keep float precision in ExplainSQL [@&#8203;kumakichi](https://redirect.github.com/kumakichi) ([#&#8203;6495](https://redirect.github.com/go-gorm/gorm/issues/6495))
-   test: coverage for tabletype added [@&#8203;saeidee](https://redirect.github.com/saeidee) ([#&#8203;6496](https://redirect.github.com/go-gorm/gorm/issues/6496))
-   test: coverage for foreign key violation err [@&#8203;saeidee](https://redirect.github.com/saeidee) ([#&#8203;6403](https://redirect.github.com/go-gorm/gorm/issues/6403))
-   ci: fix mariadb mysqladmin [@&#8203;saeidee](https://redirect.github.com/saeidee) ([#&#8203;6401](https://redirect.github.com/go-gorm/gorm/issues/6401))
-   test: coverage for duplicated key err [@&#8203;saeidee](https://redirect.github.com/saeidee) ([#&#8203;6389](https://redirect.github.com/go-gorm/gorm/issues/6389))
-   Fix incorrect documentation comment (has many -> has one) [@&#8203;johannes-riecken](https://redirect.github.com/johannes-riecken) ([#&#8203;6382](https://redirect.github.com/go-gorm/gorm/issues/6382))
-   fix: database/sql.Scanner should not retain references [@&#8203;ncruces](https://redirect.github.com/ncruces) ([#&#8203;6380](https://redirect.github.com/go-gorm/gorm/issues/6380))
-   feat: add \*sql.DB connector that uses database context [@&#8203;lzakharov](https://redirect.github.com/lzakharov) ([#&#8203;6366](https://redirect.github.com/go-gorm/gorm/issues/6366))
-   reafactor: add nil detection when sqldb return [@&#8203;KantaHasegawa](https://redirect.github.com/KantaHasegawa) ([#&#8203;6373](https://redirect.github.com/go-gorm/gorm/issues/6373))
-   refactor: remove unnecessary prepared statement allocation [@&#8203;lzakharov](https://redirect.github.com/lzakharov) ([#&#8203;6374](https://redirect.github.com/go-gorm/gorm/issues/6374))
-   fix: avoid panic when open fails [@&#8203;black-06](https://redirect.github.com/black-06) ([#&#8203;6368](https://redirect.github.com/go-gorm/gorm/issues/6368))
-   fix: begin transaction fail, rollback panic [@&#8203;zhouit](https://redirect.github.com/zhouit) ([#&#8203;6365](https://redirect.github.com/go-gorm/gorm/issues/6365))
-   max identifier length changed to 63 [@&#8203;alidevhere](https://redirect.github.com/alidevhere) ([#&#8203;6337](https://redirect.github.com/go-gorm/gorm/issues/6337))
-   fix: save with hook ([#&#8203;6285](https://redirect.github.com/go-gorm/gorm/issues/6285)) [@&#8203;black-06](https://redirect.github.com/black-06) ([#&#8203;6294](https://redirect.github.com/go-gorm/gorm/issues/6294))
-   fix(nested transaction): SavePoint SQL Statement not support in Prepared Statements [@&#8203;wangliuyang520](https://redirect.github.com/wangliuyang520) ([#&#8203;6220](https://redirect.github.com/go-gorm/gorm/issues/6220))
-   refactor: error translator test [@&#8203;saeidee](https://redirect.github.com/saeidee) ([#&#8203;6350](https://redirect.github.com/go-gorm/gorm/issues/6350))
-   Added support of "Violates Foreign Key Constraint" [@&#8203;amirejaz75](https://redirect.github.com/amirejaz75) ([#&#8203;6329](https://redirect.github.com/go-gorm/gorm/issues/6329))
-   feature: rename License to LICENSE ([#&#8203;6331](https://redirect.github.com/go-gorm/gorm/issues/6331)) [@&#8203;Avinaba-Bhattacharjee](https://redirect.github.com/Avinaba-Bhattacharjee) ([#&#8203;6336](https://redirect.github.com/go-gorm/gorm/issues/6336))
-   fix:clickhouse error not capture([#&#8203;6277](https://redirect.github.com/go-gorm/gorm/issues/6277)) [@&#8203;201430098137](https://redirect.github.com/201430098137) ([#&#8203;6321](https://redirect.github.com/go-gorm/gorm/issues/6321))
-   fix: 🐛 embedded struct test failed with custom datatypes [@&#8203;aclich](https://redirect.github.com/aclich) ([#&#8203;6311](https://redirect.github.com/go-gorm/gorm/issues/6311))
-   feat: migrator support table comment [@&#8203;johnmai-dev](https://redirect.github.com/johnmai-dev) ([#&#8203;6225](https://redirect.github.com/go-gorm/gorm/issues/6225))
-   feat: unscoped association ([#&#8203;5899](https://redirect.github.com/go-gorm/gorm/issues/5899)) [@&#8203;black-06](https://redirect.github.com/black-06) ([#&#8203;6246](https://redirect.github.com/go-gorm/gorm/issues/6246))
-   fix: 🐛 some numeric types in embedded pointer type struct cause test failed [@&#8203;hykuan](https://redirect.github.com/hykuan) ([#&#8203;6293](https://redirect.github.com/go-gorm/gorm/issues/6293))
-   fix: use slice Stale sort [@&#8203;Hanwn](https://redirect.github.com/Hanwn) ([#&#8203;6263](https://redirect.github.com/go-gorm/gorm/issues/6263))
-   fix: nested joins alias [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;6265](https://redirect.github.com/go-gorm/gorm/issues/6265))
-   fix typo in comment example of DB.Table() [@&#8203;yikakia](https://redirect.github.com/yikakia) ([#&#8203;6266](https://redirect.github.com/go-gorm/gorm/issues/6266))
-   fix: avoid coroutine leaks when the dialecter initialization fails. [@&#8203;onlyice](https://redirect.github.com/onlyice) ([#&#8203;6249](https://redirect.github.com/go-gorm/gorm/issues/6249))
-   fix: unit test [@&#8203;black-06](https://redirect.github.com/black-06) ([#&#8203;6250](https://redirect.github.com/go-gorm/gorm/issues/6250))
-   feat: support embedded preload [@&#8203;black-06](https://redirect.github.com/black-06) ([#&#8203;6137](https://redirect.github.com/go-gorm/gorm/issues/6137))
-   fix cond in scopes [@&#8203;black-06](https://redirect.github.com/black-06) ([#&#8203;6152](https://redirect.github.com/go-gorm/gorm/issues/6152))
-   fix: many2many association with duplicate belongs to elem [@&#8203;bsmith-auth0](https://redirect.github.com/bsmith-auth0) ([#&#8203;6206](https://redirect.github.com/go-gorm/gorm/issues/6206))
-   refactor(migrator): non-standard codes [@&#8203;demoManito](https://redirect.github.com/demoManito) ([#&#8203;6180](https://redirect.github.com/go-gorm/gorm/issues/6180))
-   chore(deps): bump actions/stale from 7 to 8 @&#8203;[dependabot\[bot\]](https://redirect.github.com/apps/dependabot) ([#&#8203;6190](https://redirect.github.com/go-gorm/gorm/issues/6190))
-   fix: `limit(0).offset(b)` return all data, where b <= 0 [@&#8203;Hanwn](https://redirect.github.com/Hanwn) ([#&#8203;6191](https://redirect.github.com/go-gorm/gorm/issues/6191))
-   fix: embedded should be nil if not exists [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;6219](https://redirect.github.com/go-gorm/gorm/issues/6219))
-   refactor: translatorError flag added for backward compatibility [@&#8203;saeidee](https://redirect.github.com/saeidee) ([#&#8203;6178](https://redirect.github.com/go-gorm/gorm/issues/6178))
-   avoid starting a transaction when performing only one insert operation in CreateInBatches [@&#8203;chenyahui](https://redirect.github.com/chenyahui) ([#&#8203;6174](https://redirect.github.com/go-gorm/gorm/issues/6174))
-   fix: count with group ([#&#8203;6157](https://redirect.github.com/go-gorm/gorm/issues/6157)) [@&#8203;black-06](https://redirect.github.com/black-06) ([#&#8203;6160](https://redirect.github.com/go-gorm/gorm/issues/6160))
-   save should be idempotent [@&#8203;black-06](https://redirect.github.com/black-06) ([#&#8203;6149](https://redirect.github.com/go-gorm/gorm/issues/6149))
-   chore(deps): bump actions/setup-go from 3 to 4 @&#8203;[dependabot\[bot\]](https://redirect.github.com/apps/dependabot) ([#&#8203;6165](https://redirect.github.com/go-gorm/gorm/issues/6165))
-   feat: support nested join [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;6067](https://redirect.github.com/go-gorm/gorm/issues/6067))
-   test: pgsql alter column from smallint or string to boolean [@&#8203;jeffry-luqman](https://redirect.github.com/jeffry-luqman) ([#&#8203;6107](https://redirect.github.com/go-gorm/gorm/issues/6107))
-   fix: diff schema update assign value [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;6096](https://redirect.github.com/go-gorm/gorm/issues/6096))
-   refactor: translate error only when it is not nil [@&#8203;saeidee](https://redirect.github.com/saeidee) ([#&#8203;6133](https://redirect.github.com/go-gorm/gorm/issues/6133))
-   Fix: Composite primary key with auto-increment value returns 0 after insert [@&#8203;truongns](https://redirect.github.com/truongns) ([#&#8203;6127](https://redirect.github.com/go-gorm/gorm/issues/6127))
-   fix: on confilct with default null [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;6129](https://redirect.github.com/go-gorm/gorm/issues/6129))
-   feat: Unique Constraint Violation error translator for different drivers [@&#8203;saeidee](https://redirect.github.com/saeidee) ([#&#8203;6004](https://redirect.github.com/go-gorm/gorm/issues/6004))
-   Refactor: Remove redundant code [@&#8203;xiaoliwang](https://redirect.github.com/xiaoliwang) ([#&#8203;6087](https://redirect.github.com/go-gorm/gorm/issues/6087))
-   Create and drop view [@&#8203;black-06](https://redirect.github.com/black-06) ([#&#8203;6097](https://redirect.github.com/go-gorm/gorm/issues/6097))
-   quotes on docker-compose.yml ports [@&#8203;xiaoliwang](https://redirect.github.com/xiaoliwang) ([#&#8203;6089](https://redirect.github.com/go-gorm/gorm/issues/6089))
-   test: pgsql migrate unique index [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;6028](https://redirect.github.com/go-gorm/gorm/issues/6028))
-   fix: update panic if model is not ptr [@&#8203;black-06](https://redirect.github.com/black-06) ([#&#8203;6037](https://redirect.github.com/go-gorm/gorm/issues/6037))
-   fix: association concurrently appending [@&#8203;black-06](https://redirect.github.com/black-06) ([#&#8203;6044](https://redirect.github.com/go-gorm/gorm/issues/6044))
-   fix: miss join type [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;6056](https://redirect.github.com/go-gorm/gorm/issues/6056))
-   Issue 6054: Unscoped not working with PreLoad on Joins [@&#8203;manstis](https://redirect.github.com/manstis) ([#&#8203;6058](https://redirect.github.com/go-gorm/gorm/issues/6058))
-   feat: add tidb integration test cases [@&#8203;Icemap](https://redirect.github.com/Icemap) ([#&#8203;6014](https://redirect.github.com/go-gorm/gorm/issues/6014))
-   fix:throw model value required error [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;6031](https://redirect.github.com/go-gorm/gorm/issues/6031))
-   fix: ignore nil query [@&#8203;chyroc](https://redirect.github.com/chyroc) ([#&#8203;6021](https://redirect.github.com/go-gorm/gorm/issues/6021))
-   fix: support zeroValue tag on DeletedAt [@&#8203;qiankunli](https://redirect.github.com/qiankunli) ([#&#8203;6011](https://redirect.github.com/go-gorm/gorm/issues/6011))
-   fix(schema): field is only unique when there is one unique index [@&#8203;Nomango](https://redirect.github.com/Nomango) ([#&#8203;5974](https://redirect.github.com/go-gorm/gorm/issues/5974))
-   fix(migrator): Tag default:'null' always causes field migration [#&#8203;5953](https://redirect.github.com/go-gorm/gorm/issues/5953) [@&#8203;Nomango](https://redirect.github.com/Nomango) ([#&#8203;5954](https://redirect.github.com/go-gorm/gorm/issues/5954))
-   fix(migrator): ignore relationships when migrating [#&#8203;5913](https://redirect.github.com/go-gorm/gorm/issues/5913) [@&#8203;Nomango](https://redirect.github.com/Nomango) ([#&#8203;5946](https://redirect.github.com/go-gorm/gorm/issues/5946))
-   chore(deps): bump actions/stale from 6 to 7 @&#8203;[dependabot\[bot\]](https://redirect.github.com/apps/dependabot) ([#&#8203;5945](https://redirect.github.com/go-gorm/gorm/issues/5945))
-   test(MigrateColumn): mock alter column to improve field compare [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5499](https://redirect.github.com/go-gorm/gorm/issues/5499))
-   feat: support inner join [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5583](https://redirect.github.com/go-gorm/gorm/issues/5583))
-   DryRun for migrator [@&#8203;defool](https://redirect.github.com/defool) ([#&#8203;5689](https://redirect.github.com/go-gorm/gorm/issues/5689))
-   fix:Issue migrating field with CURRENT_TIMESTAMP [@&#8203;0fv](https://redirect.github.com/0fv) ([#&#8203;5906](https://redirect.github.com/go-gorm/gorm/issues/5906))
-   Update func comments in chainable_api and FirstOr\_ [@&#8203;naterarmstrong](https://redirect.github.com/naterarmstrong) ([#&#8203;5935](https://redirect.github.com/go-gorm/gorm/issues/5935))
-   Add test case for embedded value selects [@&#8203;emcfarlane](https://redirect.github.com/emcfarlane) ([#&#8203;5901](https://redirect.github.com/go-gorm/gorm/issues/5901))
-   fix: skip append relation field to default db value [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5885](https://redirect.github.com/go-gorm/gorm/issues/5885))
-   clear code syntax [@&#8203;wjw1758548031](https://redirect.github.com/wjw1758548031) ([#&#8203;5889](https://redirect.github.com/go-gorm/gorm/issues/5889))
-   fix(FindInBatches): throw err if pk not exists [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5868](https://redirect.github.com/go-gorm/gorm/issues/5868))
-   fix bug in windows [@&#8203;kvii](https://redirect.github.com/kvii) ([#&#8203;5844](https://redirect.github.com/go-gorm/gorm/issues/5844))
-   cleanup(prepare_stmt.go): unnecessary map delete [@&#8203;demoManito](https://redirect.github.com/demoManito) ([#&#8203;5849](https://redirect.github.com/go-gorm/gorm/issues/5849))
-   doc(README.md): add contributors [@&#8203;demoManito](https://redirect.github.com/demoManito) ([#&#8203;5847](https://redirect.github.com/go-gorm/gorm/issues/5847))
-   fix logger path bug [@&#8203;kvii](https://redirect.github.com/kvii) ([#&#8203;5836](https://redirect.github.com/go-gorm/gorm/issues/5836))
-   test(utils): add utils unit test [@&#8203;demoManito](https://redirect.github.com/demoManito) ([#&#8203;5834](https://redirect.github.com/go-gorm/gorm/issues/5834))
-   feat: golangci add goimports and whitespace [@&#8203;demoManito](https://redirect.github.com/demoManito) ([#&#8203;5835](https://redirect.github.com/go-gorm/gorm/issues/5835))
-   test(clause/joins): add join unit test [@&#8203;demoManito](https://redirect.github.com/demoManito) ([#&#8203;5832](https://redirect.github.com/go-gorm/gorm/issues/5832))
-   fix(Joins): args with select and omit [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5790](https://redirect.github.com/go-gorm/gorm/issues/5790))
-   test: invalid cache plan with prepare stmt [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5778](https://redirect.github.com/go-gorm/gorm/issues/5778))
-   feat(PreparedStmtDB): support reset [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5782](https://redirect.github.com/go-gorm/gorm/issues/5782))
-   fix: association without pks [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5779](https://redirect.github.com/go-gorm/gorm/issues/5779))
-   fix: limit=0 results ([#&#8203;5735](https://redirect.github.com/go-gorm/gorm/issues/5735)) [@&#8203;robhafner](https://redirect.github.com/robhafner) ([#&#8203;5736](https://redirect.github.com/go-gorm/gorm/issues/5736))
-   fix: `func (schema *Schema) guessRelation(...)` primaryFields are overwritten [@&#8203;demoManito](https://redirect.github.com/demoManito) ([#&#8203;5721](https://redirect.github.com/go-gorm/gorm/issues/5721))
-   Fix OnConstraint builder [@&#8203;xwjdsh](https://redirect.github.com/xwjdsh) ([#&#8203;5738](https://redirect.github.com/go-gorm/gorm/issues/5738))
-   fix: prepare deadlock [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5568](https://redirect.github.com/go-gorm/gorm/issues/5568))
-   fix maybe nil panic [@&#8203;demoManito](https://redirect.github.com/demoManito) ([#&#8203;5720](https://redirect.github.com/go-gorm/gorm/issues/5720))
-   chore(deps): bump actions/stale from 5 to 6 @&#8203;[dependabot\[bot\]](https://redirect.github.com/apps/dependabot) ([#&#8203;5717](https://redirect.github.com/go-gorm/gorm/issues/5717))
-   fix: return id which have type string after created [@&#8203;nohattee](https://redirect.github.com/nohattee) ([#&#8203;5477](https://redirect.github.com/go-gorm/gorm/issues/5477))
-   feat: migrator support type aliases [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5627](https://redirect.github.com/go-gorm/gorm/issues/5627))
-   fix: scan array [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5624](https://redirect.github.com/go-gorm/gorm/issues/5624))
-   support scan assign slice cap [@&#8203;demoManito](https://redirect.github.com/demoManito) ([#&#8203;5634](https://redirect.github.com/go-gorm/gorm/issues/5634))
-   fix(migrate): same embedded filed name [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5705](https://redirect.github.com/go-gorm/gorm/issues/5705))
-   fix: update omit [@&#8203;qqxhb](https://redirect.github.com/qqxhb) ([#&#8203;5699](https://redirect.github.com/go-gorm/gorm/issues/5699))
-   AutoMigrate() should always migrate checks, even there is no relationship constraints. [@&#8203;googollee](https://redirect.github.com/googollee) ([#&#8203;5644](https://redirect.github.com/go-gorm/gorm/issues/5644))
-   Rewrite of finisher_api Godocs [@&#8203;bruc3mackenzi3](https://redirect.github.com/bruc3mackenzi3) ([#&#8203;5618](https://redirect.github.com/go-gorm/gorm/issues/5618))
-   simplified regexp [@&#8203;xiaoliwang](https://redirect.github.com/xiaoliwang) ([#&#8203;5677](https://redirect.github.com/go-gorm/gorm/issues/5677))
-   Optimize: code logic db.scanIntoStruct() [@&#8203;demoManito](https://redirect.github.com/demoManito) ([#&#8203;5633](https://redirect.github.com/go-gorm/gorm/issues/5633))
-   test: remove uuid autoincrement tag [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5620](https://redirect.github.com/go-gorm/gorm/issues/5620))
-   Replace `ioutil.Discard` with `io.Discard` [@&#8203;zaneli](https://redirect.github.com/zaneli) ([#&#8203;5603](https://redirect.github.com/go-gorm/gorm/issues/5603))
-   Refactor: redundant type from composite literal [@&#8203;zaneli](https://redirect.github.com/zaneli) ([#&#8203;5604](https://redirect.github.com/go-gorm/gorm/issues/5604))
-   Add Go 1.19 Support [@&#8203;Aoang](https://redirect.github.com/Aoang) ([#&#8203;5608](https://redirect.github.com/go-gorm/gorm/issues/5608))
-   fix: correct grammar [@&#8203;seaworn](https://redirect.github.com/seaworn) ([#&#8203;5600](https://redirect.github.com/go-gorm/gorm/issues/5600))
-   Update Delete Godoc to describe soft delete behaviour [@&#8203;bruc3mackenzi3](https://redirect.github.com/bruc3mackenzi3) ([#&#8203;5554](https://redirect.github.com/go-gorm/gorm/issues/5554))
-   chore: fix gorm tag [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5577](https://redirect.github.com/go-gorm/gorm/issues/5577))
-   fix: empty serilizer err [#&#8203;5524](https://redirect.github.com/go-gorm/gorm/issues/5524) [@&#8203;philhuan](https://redirect.github.com/philhuan) ([#&#8203;5525](https://redirect.github.com/go-gorm/gorm/issues/5525))
-   Fixed some typos in the code comment [@&#8203;Minjerous](https://redirect.github.com/Minjerous) ([#&#8203;5549](https://redirect.github.com/go-gorm/gorm/issues/5549))
-   fix: embedded default value [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5540](https://redirect.github.com/go-gorm/gorm/issues/5540))
-   fix bad logging performance of bulk create ([#&#8203;5520](https://redirect.github.com/go-gorm/gorm/issues/5520)) [@&#8203;zxdvd](https://redirect.github.com/zxdvd) ([#&#8203;5521](https://redirect.github.com/go-gorm/gorm/issues/5521))
-   fix empty QueryClauses in association ([#&#8203;5502](https://redirect.github.com/go-gorm/gorm/issues/5502)) [@&#8203;goxiaoy](https://redirect.github.com/goxiaoy) ([#&#8203;5503](https://redirect.github.com/go-gorm/gorm/issues/5503))
-   Adjust ToStringKey use unpack params, fix  pass \[]any as any in variadic function [@&#8203;alingse](https://redirect.github.com/alingse) ([#&#8203;5500](https://redirect.github.com/go-gorm/gorm/issues/5500))
-   test: pg array type [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5480](https://redirect.github.com/go-gorm/gorm/issues/5480))
-   feat: use callback to handle transaction [@&#8203;sunfuze](https://redirect.github.com/sunfuze) ([#&#8203;5455](https://redirect.github.com/go-gorm/gorm/issues/5455))
-   fix: association many2many duplicate elem [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5473](https://redirect.github.com/go-gorm/gorm/issues/5473))
-   fix(MigrateColumn):declared different type without length [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5465](https://redirect.github.com/go-gorm/gorm/issues/5465))
-   fix:serializer contain field panic [@&#8203;wuweishuo](https://redirect.github.com/wuweishuo) ([#&#8203;5461](https://redirect.github.com/go-gorm/gorm/issues/5461))
-   feat: add method GetIndexes [@&#8203;qqxhb](https://redirect.github.com/qqxhb) ([#&#8203;5436](https://redirect.github.com/go-gorm/gorm/issues/5436))
-   fix: reset null value in slice [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5417](https://redirect.github.com/go-gorm/gorm/issues/5417))
-   enhancement: Avoid calling reflect.New() when passing in slice of values to `Scan()` [@&#8203;Bexanderthebex](https://redirect.github.com/Bexanderthebex) ([#&#8203;5388](https://redirect.github.com/go-gorm/gorm/issues/5388))
-   chore(deps): bump gorm.io/driver/mysql from 1.3.3 to 1.3.4 in /tests @&#8203;[dependabot\[bot\]](https://redirect.github.com/apps/dependabot) ([#&#8203;5385](https://redirect.github.com/go-gorm/gorm/issues/5385))
-   fix: migrate column default value [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5359](https://redirect.github.com/go-gorm/gorm/issues/5359))
-   fixed FirstOrCreate not handled error when table is not exists [@&#8203;ophum](https://redirect.github.com/ophum) ([#&#8203;5367](https://redirect.github.com/go-gorm/gorm/issues/5367))
-   fix: duplicate column scan [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5369](https://redirect.github.com/go-gorm/gorm/issues/5369))
-   test: test for skip prepared when auto migrate [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5350](https://redirect.github.com/go-gorm/gorm/issues/5350))
-   Fixed [#&#8203;5355](https://redirect.github.com/go-gorm/gorm/issues/5355) - Named variables don't work when followed by Windows CRLF line endings [@&#8203;clarkmcc](https://redirect.github.com/clarkmcc) ([#&#8203;5356](https://redirect.github.com/go-gorm/gorm/issues/5356))
-   fix: trx in hooks clone stmt [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5338](https://redirect.github.com/go-gorm/gorm/issues/5338))
-   fix: quote index when creating table [@&#8203;black-06](https://redirect.github.com/black-06) ([#&#8203;5331](https://redirect.github.com/go-gorm/gorm/issues/5331))
-   fix: many2many auto migrate [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5322](https://redirect.github.com/go-gorm/gorm/issues/5322))
-   fix: preload with skip hooks [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5310](https://redirect.github.com/go-gorm/gorm/issues/5310))
-   fix: callbcak sort when using multiple plugin [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5304](https://redirect.github.com/go-gorm/gorm/issues/5304))
-   fix: add judge result of auto_migrate [@&#8203;Heliner](https://redirect.github.com/Heliner) ([#&#8203;5306](https://redirect.github.com/go-gorm/gorm/issues/5306))
-   fix: AutoMigrate with special table name [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5301](https://redirect.github.com/go-gorm/gorm/issues/5301))
-   index: add composite id [@&#8203;photon3108](https://redirect.github.com/photon3108) ([#&#8203;5269](https://redirect.github.com/go-gorm/gorm/issues/5269))
-   test: test for postgrs serial column [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5234](https://redirect.github.com/go-gorm/gorm/issues/5234))
-   check for pointer to pointer value [@&#8203;aelmel](https://redirect.github.com/aelmel) ([#&#8203;5278](https://redirect.github.com/go-gorm/gorm/issues/5278))
-   fix: stmt.Changed zero value filed behavior [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5281](https://redirect.github.com/go-gorm/gorm/issues/5281))
-   fix missing error-check in AutoMigrate [@&#8203;glebarez](https://redirect.github.com/glebarez) ([#&#8203;5283](https://redirect.github.com/go-gorm/gorm/issues/5283))
-   fix: FindInBatches with offset limit [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5255](https://redirect.github.com/go-gorm/gorm/issues/5255))
-   fix spelling mistake [@&#8203;ZhangShenao](https://redirect.github.com/ZhangShenao) ([#&#8203;5256](https://redirect.github.com/go-gorm/gorm/issues/5256))
-   chore(deps): bump actions/setup-go from 2 to 3 @&#8203;[dependabot\[bot\]](https://redirect.github.com/apps/dependabot) ([#&#8203;5243](https://redirect.github.com/go-gorm/gorm/issues/5243))
-   chore(deps): bump actions/stale from 4 to 5 @&#8203;[dependabot\[bot\]](https://redirect.github.com/apps/dependabot) ([#&#8203;5244](https://redirect.github.com/go-gorm/gorm/issues/5244))
-   fix: FirstOrCreate RowsAffected [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5250](https://redirect.github.com/go-gorm/gorm/issues/5250))
-   Fix scanIntoStruct [@&#8203;delmorof](https://redirect.github.com/delmorof) ([#&#8203;5241](https://redirect.github.com/go-gorm/gorm/issues/5241))
-   Set permissions for GitHub actions [@&#8203;naveensrinivasan](https://redirect.github.com/naveensrinivasan) ([#&#8203;5237](https://redirect.github.com/go-gorm/gorm/issues/5237))
-   Offset issue resolved for scanning results back into struct [@&#8203;Joker666](https://redirect.github.com/Joker666) ([#&#8203;5227](https://redirect.github.com/go-gorm/gorm/issues/5227))
-   unify db receiver name [@&#8203;ZhangShenao](https://redirect.github.com/ZhangShenao) ([#&#8203;5215](https://redirect.github.com/go-gorm/gorm/issues/5215))
-   fix: context missing in association [@&#8203;goxiaoy](https://redirect.github.com/goxiaoy) ([#&#8203;5214](https://redirect.github.com/go-gorm/gorm/issues/5214))
-   fix variable shadowing [@&#8203;ZhangShenao](https://redirect.github.com/ZhangShenao) ([#&#8203;5212](https://redirect.github.com/go-gorm/gorm/issues/5212))
-   chore(deps): bump actions/cache from 2 to 3 @&#8203;[dependabot\[bot\]](https://redirect.github.com/apps/dependabot) ([#&#8203;5196](https://redirect.github.com/go-gorm/gorm/issues/5196))
-   fix: throw err if association model miss primary key [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5187](https://redirect.github.com/go-gorm/gorm/issues/5187))
-   style: fix coding typo [@&#8203;ag9920](https://redirect.github.com/ag9920) ([#&#8203;5184](https://redirect.github.com/go-gorm/gorm/issues/5184))
-   style: fix linter check problem for NamingStrategy and onConflictOption [@&#8203;ag9920](https://redirect.github.com/ag9920) ([#&#8203;5174](https://redirect.github.com/go-gorm/gorm/issues/5174))
-   test: fix test case and utils [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5172](https://redirect.github.com/go-gorm/gorm/issues/5172))
-   Use WriteByte for single byte operations [@&#8203;moredure](https://redirect.github.com/moredure) ([#&#8203;5167](https://redirect.github.com/go-gorm/gorm/issues/5167))
-   fix when index name is "type", parseFieldIndexes will set index TYPE is "TYPE" [@&#8203;labulakalia](https://redirect.github.com/labulakalia) ([#&#8203;5155](https://redirect.github.com/go-gorm/gorm/issues/5155))
-   chore(deps): bump actions/checkout from 2 to 3 @&#8203;[dependabot\[bot\]](https://redirect.github.com/apps/dependabot) ([#&#8203;5133](https://redirect.github.com/go-gorm/gorm/issues/5133))
-   ToSQL should enable SkipDefaultTransaction by default [@&#8203;CaoManhDat](https://redirect.github.com/CaoManhDat) ([#&#8203;5125](https://redirect.github.com/go-gorm/gorm/issues/5125))
-   fix: query scanner in single column [@&#8203;a631807682](https://redirect.github.com/a631807682) ([#&#8203;5111](https://redirect.github.com/go-gorm/gorm/issues/5111))
-   feat: support gob serialize [@&#8203;ag9920](https://redirect.github.com/ag9920) ([#&#8203;5108](https://redirect.github.com/go-gorm/gorm/issues/5108))
-   Fix invalid foreign key constraint name with long schema qualified tables [@&#8203;mnussbaum](https://redirect.github.com/mnussbaum) ([#&#8203;5045](https://redirect.github.com/go-gorm/gorm/issues/5045))
-   fix typo in TxCommitter interface comment & improve CheckTruth, chek val empty first [@&#8203;0x2d3c](https://redirect.github.com/0x2d3c) ([#&#8203;5094](https://redirect.github.com/go-gorm/gorm/issues/5094))
-   style: use ReplaceAll instead of Replace [@&#8203;hu-quan-er](https://redirect.github.com/hu-quan-er) ([#&#8203;5095](https://redirect.github.com/go-gorm/gorm/issues/5095))
-   Inherit clone flag (NewDB) on transaction creation [@&#8203;Gilwe](https://redirect.github.com/Gilwe) ([#&#8203;5012](https://redirect.github.com/go-gorm/gorm/issues/5012))
-   Fixed the use of "or" to be " OR ", to account for words that contain… [@&#8203;sammyrnycreal](https://redirect.github.com/sammyrnycreal) ([#&#8203;5074](https://redirect.github.com/go-gorm/gorm/issues/5074))
-   Add Serializer, Update migrator ColumnType interface [@&#8203;jinzhu](https://redirect.github.com/jinzhu) ([#&#8203;5089](https://redirect.github.com/go-gorm/gorm/issues/5089))
-   fix: isPrintable incorrect [@&#8203;li-jin-gou](https://redirect.github.com/li-jin-gou) ([#&#8203;5076](https://redirect.github.com/go-gorm/gorm/issues/5076))
-   fix: replace empty table name result in panic [@&#8203;li-jin-gou](https://redirect.github.com/li-jin-gou) ([#&#8203;5048](https://redirect.github.com/go-gorm/gorm/issues/5048))
-   Added comments to existing methods [@&#8203;Saurabh-Thakre](https://redirect.github.com/Saurabh-Thakre) ([#&#8203;5043](https://redirect.github.com/go-gorm/gorm/issues/5043))
-   preoload not allowd before count [@&#8203;0fv](https://redirect.github.com/0fv) ([#&#8203;5023](https://redirect.github.com/go-gorm/gorm/issues/5023))
-   fix: omit not work when use join  [@&#8203;li-jin-gou](https://redirect.github.com/li-jin-gou) ([#&#8203;5034](https://redirect.github.com/go-gorm/gorm/issues/5034))
-   chore(deps): bump gorm.io/driver/mysql from 1.2.1 to 1.2.3 in /tests @&#8203;[dependabot\[bot\]](https://redirect.github.com/apps/dependabot) ([#&#8203;4987](https://redirect.github.com/go-gorm/gorm/issues/4987))
-   time.Time, \[]byte type add alias support. (rebase master) [@&#8203;piyongcai](https://redirect.github.com/piyongcai) ([#&#8203;4992](https://redirect.github.com/go-gorm/gorm/issues/4992))
-   fix: auto migration column order unpredictable [@&#8203;halfcrazy](https://redirect.github.com/halfcrazy) ([#&#8203;4980](https://redirect.github.com/go-gorm/gorm/issues/4980))
-   improve the error handle in tests_test [@&#8203;liweitingwt](https://redirect.github.com/liweitingwt) ([#&#8203;4964](https://redirect.github.com/go-gorm/gorm/issues/4964))
-   Fix: Where clauses with named arguments may cause generation of unintended queries [@&#8203;emregullu](https://redirect.github.com/emregullu) ([#&#8203;4937](https://redirect.github.com/go-gorm/gorm/issues/4937))
-   modify unscoped judge [@&#8203;liweitingwt](https://redirect.github.com/liweitingwt) ([#&#8203;4929](https://redirect.github.com/go-gorm/gorm/issues/4929))
-   fix type alias AutoMigrate bug（Add Test Case） [@&#8203;piyongcai](https://redirect.github.com/piyongcai) ([#&#8203;4888](https://redirect.github.com/go-gorm/gorm/issues/4888))
-   Use Golangci configuration file [@&#8203;mmorel-35](https://redirect.github.com/mmorel-35) ([#&#8203;4896](https://redirect.github.com/go-gorm/gorm/issues/4896))
-   feat: go code style adjust and optimize code for callbacks package [@&#8203;daheige](https://redirect.github.com/daheige) ([#&#8203;4861](https://redirect.github.com/go-gorm/gorm/issues/4861))
-   Add unused argument [@&#8203;jinzhu](https://redirect.github.com/jinzhu) ([#&#8203;4871](https://redirect.github.com/go-gorm/gorm/issues/4871))
-   Bump gorm.io/driver/mysql from 1.1.3 to 1.2.0 in /tests @&#8203;[dependabot\[bot\]](https://redirect.github.com/apps/dependabot) ([#&#8203;4856](https://redirect.github.com/go-gorm/gorm/issues/4856))
-   Bump github.com/jinzhu/now from 1.1.2 to 1.1.3 in /tests @&#8203;[dependabot\[bot\]](https://redirect.github.com/apps/dependabot) ([#&#8203;4866](https://redirect.github.com/go-gorm/gorm/issues/4866))
-   Bump github.com/jinzhu/now from 1.1.2 to 1.1.3 @&#8203;[dependabot\[bot\]](https://redirect.github.com/apps/dependabot) ([#&#8203;4865](https://redirect.github.com/go-gorm/gorm/issues/4865))
-   feat(migrator,migrator/migrator.go,tests/migrate_test.go) : Get multiple data tables for migrator. [@&#8203;dino-ma](https://redirect.github.com/dino-ma) ([#&#8203;4841](https://redirect.github.com/go-gorm/gorm/issues/4841))
-   Fix self-referential belongs to constraint [@&#8203;mgovilla](https://redirect.github.com/mgovilla) ([#&#8203;4801](https://redirect.github.com/go-gorm/gorm/issues/4801))
-   Refactor if logic [@&#8203;daheige](https://redirect.github.com/daheige) ([#&#8203;4683](https://redirect.github.com/go-gorm/gorm/issues/4683))
-   Bump gorm.io/driver/sqlserver from 1.1.2 to 1.2.0 in /tests @&#8203;[dependabot\[bot\]](https://redirect.github.com/apps/dependabot) ([#&#8203;4820](https://redirect.github.com/go-gorm/gorm/issues/4820))
-   Add ToSQL support to generate SQL string. [@&#8203;huacnlee](https://redirect.github.com/huacnlee) ([#&#8203;4787](https://redirect.github.com/go-gorm/gorm/issues/4787))
-   Refactor ParseWithSchemaTable method and improve test. [@&#8203;huacnlee](https://redirect.github.com/huacnlee) ([#&#8203;4789](https://redirect.github.com/go-gorm/gorm/issues/4789))
-   fix: automigrate error caused by indexes while using dynamic table name [@&#8203;xwjdsh](https://redirect.github.com/xwjdsh) ([#&#8203;4773](https://redirect.github.com/go-gorm/gorm/issues/4773))
-   feat: Convert SQL null values to zero values for model fields which are not pointers. [@&#8203;jimlambrt](https://redirect.github.com/jimlambrt) ([#&#8203;4710](https://redirect.github.com/go-gorm/gorm/issues/4710))
-   feat: adjust PreparedStmtDB unlock location and BuildCondition if logic [@&#8203;daheige](https://redirect.github.com/daheige) ([#&#8203;4681](https://redirect.github.com/go-gorm/gorm/issues/4681))
-   feat: adjust SetupJoinTable func if..else code [@&#8203;daheige](https://redirect.github.com/daheige) ([#&#8203;4680](https://redirect.github.com/go-gorm/gorm/issues/4680))
-   fixed belongs_to & has_one reversed if field same (proper fix) [@&#8203;paraswaykole](https://redirect.github.com/paraswaykole) ([#&#8203;4694](https://redirect.github.com/go-gorm/gorm/issues/4694))
-   Bump gorm.io/driver/postgres from 1.1.1 to 1.1.2 in /tests @&#8203;[dependabot\[bot\]](https://redirect.github.com/apps/dependabot) ([#&#8203;4740](https://redirect.github.com/go-gorm/gorm/issues/4740))
-   Update `tests.yml` [@&#8203;s-takehana](https://redirect.github.com/s-takehana) ([#&#8203;4741](https://redirect.github.com/go-gorm/gorm/issues/4741))
-   fix: QuoteTo not fully support raw mode [@&#8203;tr1v3r](https://redirect.github.com/tr1v3r) ([#&#8203;4735](https://redirect.github.com/go-gorm/gorm/issues/4735))
-   fix (clause/expression): Allow sql stmt terminator [@&#8203;jimlambrt](https://redirect.github.com/jimlambrt) ([#&#8203;4693](https://redirect.github.com/go-gorm/gorm/issues/4693))
-   Bump gorm.io/driver/postgres from 1.1.0 to 1.1.1 in /tests @&#8203;[dependabot\[bot\]](https://redirect.github.com/apps/dependabot) ([#&#8203;4699](https://redirect.github.com/go-gorm/gorm/issues/4699))
-   Bump gorm.io/driver/sqlite from 1.1.4 to 1.1.5 in /tests @&#8203;[dependabot\[bot\]](https://redirect.github.com/apps/dependabot) ([#&#8203;4701](https://redirect.github.com/go-gorm/gorm/issues/4701))
-   Refactor update record [@&#8203;jinzhu](https://redirect.github.com/jinzhu) ([#&#8203;4679](https://redirect.github.com/go-gorm/gorm/issues/4679))
-   Bump github.com/lib/pq from 1.10.2 to 1.10.3 in /tests @&#8203;[dependabot\[bot\]](https://redirect.github.com/apps/dependabot) ([#&#8203;4676](https://redirect.github.com/go-gorm/gorm/issues/4676))
-   Bump gorm.io/gorm from 1.21.13 to 1.21.14 in /tests @&#8203;[dependabot\[bot\]](https://redirect.github.com/apps/dependabot) ([#&#8203;4655](https://redirect.github.com/go-gorm/gorm/issues/4655))
-   Add Go 1.17 [@&#8203;jxlwqq](https://redirect.github.com/jxlwqq) ([#&#8203;4666](https://redirect.github.com/go-gorm/gorm/issues/4666))
-   Bump gorm.io/driver/sqlserver from 1.0.7 to 1.0.8 in /tests @&#8203;[dependabot\[bot\]](https://redirect.github.com/apps/dependabot) ([#&#8203;4631](https://redirect.github.com/go-gorm/gorm/issues/4631))
-   Fix extra 'AND' when len(values) == 0 ON IN.NegationBuild() [@&#8203;secake](https://redirect.github.com/secake) ([#&#8203;4618](https://redirect.github.com/go-gorm/gorm/issues/4618))
-   Bump gorm.io/driver/mysql from 1.1.1 to 1.1.2 in /tests @&#8203;[dependabot\[bot\]](https://redirect.github.com/apps/dependabot) ([#&#8203;4615](https://redirect.github.com/go-gorm/gorm/issues/4615))
-   Bump gorm.io/gorm from 1.21.12 to 1.21.13 in /tests @&#8203;[dependabot\[bot\]](https://redirect.github.com/apps/dependabot) ([#&#8203;4616](https://redirect.github.com/go-gorm/gorm/issues/4616))
-   feat: count accpet `db`.`table` [@&#8203;tr1v3r](https://redirect.github.com/tr1v3r) ([#&#8203;4626](https://redirect.github.com/go-gorm/gorm/issues/4626))
-   feat: QuoteTo accept clause.Expr [@&#8203;tr1v3r](https://redirect.github.com/tr1v3r) ([#&#8203;4621](https://redirect.github.com/go-gorm/gorm/issues/4621))
-   chore(logger): explicitly set config of Default Logger [@&#8203;HurSungYun](https://redirect.github.com/HurSungYun) ([#&#8203;4605](https://redirect.github.com/go-gorm/gorm/issues/4605))
-   Fix create with ignore migration [@&#8203;zkqiang](https://redirect.github.com/zkqiang) ([#&#8203;4571](https://redirect.github.com/go-gorm/gorm/issues/4571))
-   fix: table couln't be reentrant [@&#8203;SmallTianTian](https://redirect.github.com/SmallTianTian) ([#&#8203;4556](https://redirect.github.com/go-gorm/gorm/issues/4556))
-   Update Dependencies  [@&#8203;mmorel-35](https://redirect.github.com/mmorel-35) ([#&#8203;4582](https://redirect.github.com/go-gorm/gorm/issues/4582))
-   Do not emit ORDER BY for empty values [@&#8203;wfscheper](https://redirect.github.com/wfscheper) ([#&#8203;4592](https://redirect.github.com/go-gorm/gorm/issues/4592))
-   fix return value for \*schema.Check [@&#8203;daheige](https://redirect.github.com/daheige) ([#&#8203;4527](https://redirect.github.com/go-gorm/gorm/issues/4527))
-   optimize migrator.go MigrateColumn and ColumnTypes func. [@&#8203;daheige](https://redirect.github.com/daheige) ([#&#8203;4532](https://redirect.github.com/go-gorm/gorm/issues/4532))
-   optimize Parse func for fieldValue.Interface [@&#8203;daheige](https://redirect.github.com/daheige) ([#&#8203;4526](https://redirect.github.com/go-gorm/gorm/issues/4526))
-   update comment for ConvertSliceOfMapToValuesForCreate func [@&#8203;daheige](https://redirect.github.com/daheige) ([#&#8203;4529](https://redirect.github.com/go-gorm/gorm/issues/4529))
-   optimize setupValuerAndSetter func [@&#8203;daheige](https://redirect.github.com/daheige) ([#&#8203;4530](https://redirect.github.com/go-gorm/gorm/issues/4530))
-   adjust Preload fmt.Errorf [@&#8203;daheige](https://redirect.github.com/daheige) ([#&#8203;4531](https://redirect.github.com/go-gorm/gorm/issues/4531))
-   Fix create index with comments in MySQL [@&#8203;s-takehana](https://redirect.github.com/s-takehana) ([#&#8203;4521](https://redirect.github.com/go-gorm/gorm/issues/4521))
-   New Comma Expression [@&#8203;tr1v3r](https://redirect.github.com/tr1v3r) ([#&#8203;4524](https://redirect.github.com/go-gorm/gorm/issues/4524))
-   slightly better callback warning [@&#8203;bdemirpolat](https://redirect.github.com/bdemirpolat) ([#&#8203;4495](https://redirect.github.com/go-gorm/gorm/issues/4495))
-   title foreign schema for many2many to avoid panic [@&#8203;wangyuehong](https://redirect.github.com/wangyuehong) ([#&#8203;4496](https://redirect.github.com/go-gorm/gorm/issues/4496))
-   fix: fix race issue in prepare method [@&#8203;shiyu7](https://redirect.github.com/shiyu7) ([#&#8203;4487](https://redirect.github.com/go-gorm/gorm/issues/4487))
-   Fix Pluck's usage [#&#8203;4473](https://redirect.github.com/go-gorm/gorm/issues/4473) [@&#8203;wuwenchi](https://redirect.github.com/wuwenchi) ([#&#8203;4479](https://redirect.github.com/go-gorm/gorm/issues/4479))
-   Added return names to logger.Interface.Trace [@&#8203;applejag](https://redirect.github.com/applejag) ([#&#8203;4450](https://redirect.github.com/go-gorm/gorm/issues/4450))
-   Use count(\*) instead of count(1) include NULL and non-NULL rows(SQL-92). [@&#8203;tony95271](https://redirect.github.com/tony95271) ([#&#8203;4453](https://redirect.github.com/go-gorm/gorm/issues/4453))
-   Code optimize [@&#8203;daheige](https://redirect.github.com/daheige) ([#&#8203;4415](https://redirect.github.com/go-gorm/gorm/issues/4415))
-   Fix: FirstOrCreate slice out of bounds error when using 'Assign' [@&#8203;liamrfell](https://redirect.github.com/liamrfell) ([#&#8203;4436](https://redirect.github.com/go-gorm/gorm/issues/4436))
-   Support partial indexes for Upsert in Postgres [@&#8203;VitalyShein](https://redirect.github.com/VitalyShein) ([#&#8203;4442](https://redirect.github.com/go-gorm/gorm/issues/4442))
-   Update version in `tests.yml` [@&#8203;s-takehana](https://redirect.github.com/s-takehana) ([#&#8203;4432](https://redirect.github.com/go-gorm/gorm/issues/4432))
-   golint standard [@&#8203;heyanfu](https://redirect.github.com/heyanfu) ([#&#8203;4421](https://redirect.github.com/go-gorm/gorm/issues/4421))
-   Fix typo in associations_test.go [@&#8203;eltociear](https://redirect.github.com/eltociear) ([#&#8203;4407](https://redirect.github.com/go-gorm/gorm/issues/4407))
-   Small grammar fix in error message [@&#8203;Br3nda](https://redirect.github.com/Br3nda) ([#&#8203;4406](https://redirect.github.com/go-gorm/gorm/issues/4406))
-   fixed has_many stopped working if field names are identical [@&#8203;paraswaykole](https://redirect.githu

</details>

---

### Configuration

📅 **Schedule**: Branch creation - "after 5am on sunday" in timezone Europe/Prague, Automerge - "* 3-10 * * 1" in timezone Europe/Prague.

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [x] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOS4xNTguMC1ycG0iLCJ1cGRhdGVkSW5WZXIiOiIzOS4xNTguMC1ycG0iLCJ0YXJnZXRCcmFuY2giOiJtYXN0ZXIiLCJsYWJlbHMiOltdfQ==-->


---

## Discussion

### Comment by @jira-linking on April 27, 2025 at 04:37 PM UTC

Commits missing Jira IDs:
6ad8b70426ed986fe409f4698b9e3577c75f6c79


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1649*
