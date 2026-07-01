# Production Timeline

**Period**: 2020-05-01 to 2022-05-31

**Description**: Patchman production hardening and platform integration


**Total Commits**: 1826


## 2020-05 (46 commits)

- **2020-05-05**: Handle 'updated' events _by @Michal Hornicky_ `[patchman-engine]`

- **2020-05-05**: Change default sort key for advisories _by @Michal Hornicky_ `[patchman-engine]`

- **2020-05-05**: upgraded database from postgres 10 to 12 _by @Josef Hak_ `[patchman-engine]`

- **2020-05-05**: Fix tests for new advisories sort _by @Michal Hornicky_ `[patchman-engine]`

- **2020-05-06**: fix(filter): Remove Unknown filter _by @Jiri Dostal_ `[patchman-ui]`

- **2020-05-06**: chore(release): 0.8.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-05-06**: feat(select-all): Support select all on the SystemAdvisories _by @Jiri Dostal_ `[patchman-ui]`

- **2020-05-06**: feat(select-all): Add select all for AffectedSystems _by @Jiri Dostal_ `[patchman-ui]`

- **2020-05-06**: fix(lint): Fix lint _by @Jiri Dostal_ `[patchman-ui]`

- **2020-05-06**: chore(release): 0.9.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-05-11**: Add metric for different database processes _by @Michal Hornicky_ `[patchman-engine]`

- **2020-05-11**: Fix listener crashing when receiving updated messages with nonexistant systems _by @Michal Hornicky_ `[patchman-engine]`

- **2020-05-11**: Cause build _by @Michal Hornicky_ `[patchman-engine]`

- **2020-05-12**: feat(select-all): add indeterminate state for checkboxes _by @Jiri Dostal_ `[patchman-ui]`

- **2020-05-12**: chore(release): 0.10.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-05-12**: added optional KAFKA_READER_MIN_BYTES and KAFKA_READER_MAX_BYTES vars _by @Josef Hak_ `[patchman-engine]`

- **2020-05-13**: Make 'no rows modified' metric a warning _by @Michal Hornicky_ `[patchman-engine]`

- **2020-05-13**: Add a metric to report when calling other services _by @Michal Hornicky_ `[patchman-engine]`

- **2020-05-18**: added option to return all list items setting limit to -1 _by @Josef Hak_ `[patchman-engine]`

- **2020-05-18**: added test for limit -1 value _by @Josef Hak_ `[patchman-engine]`

- **2020-05-18**: updated API docs with limit -1 value _by @Josef Hak_ `[patchman-engine]`

- **2020-05-18**: updated code style due to golint update _by @Josef Hak_ `[patchman-engine]`

- **2020-05-20**: added support to run listener in host OS _by @Josef Hak_ `[patchman-engine]`

- **2020-05-20**: ensured listener does not fall on invalid 'upload' message _by @Josef Hak_ `[patchman-engine]`

- **2020-05-20**: added test for invalid upload message handling _by @Josef Hak_ `[patchman-engine]`

- **2020-05-20**: Add basic support for export endpoints _by @Michal Hornicky_ `[patchman-engine]`

- **2020-05-20**: Add systems export _by @Michal Hornicky_ `[patchman-engine]`

- **2020-05-20**: Add default sort _by @Michal Hornicky_ `[patchman-engine]`

- **2020-05-20**: Fix docs _by @Michal Hornicky_ `[patchman-engine]`

- **2020-05-22**: fix(datetime): Remove momentjs and use Date() instead _by @Jiri Dostal_ `[patchman-ui]`

- **2020-05-22**: fix(tests): Fix test for filter _by @Jiri Dostal_ `[patchman-ui]`

- **2020-05-22**: chore(release): 0.10.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-05-25**: Add filter to export endpoints _by @Michal Hornicky_ `[patchman-engine]`

- **2020-05-25**: Add search _by @Michal Hornicky_ `[patchman-engine]`

- **2020-05-27**: added "/patch" prefix to the manifest _by @Josef Hak_ `[patchman-engine]`

- **2020-05-27**: added gzip support to platform mock _by @Josef Hak_ `[patchman-engine]`

- **2020-05-27**: added response content encoding log info _by @Josef Hak_ `[patchman-engine]`

- **2020-05-27**: added ENABLE_VMAAS_CALL_COMPRESSION option to evaluator _by @Josef Hak_ `[patchman-engine]`

- **2020-05-28**: feat(export): add export to advisories page _by @Jiri Dostal_ `[patchman-ui]`

- **2020-05-28**: feat(export): Add export to systems page _by @Jiri Dostal_ `[patchman-ui]`

- **2020-05-28**: chore(release): 0.11.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-05-28**: fix(manifest): Fix manifest format _by @Jiri Dostal_ `[patchman-ui]`

- **2020-05-28**: chore(release): 0.11.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-05-28**: fix(pagination): Fix bottom pagination in tables: _by @Jiri Dostal_ `[patchman-ui]`

- **2020-05-28**: chore(release): 0.11.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-05-29**: fix(manifest): fix manifest format _by @Jiri Dostal_ `[patchman-ui]`


## 2020-06 (82 commits)

- **2020-06-01**: Remove enabled from API _by @Michal Hornicky_ `[patchman-engine]`

- **2020-06-01**: Generate docs _by @Michal Hornicky_ `[patchman-engine]`

- **2020-06-01**: chore(release): 0.11.3 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-06-01**: fix(select): fix select page number on advisories table _by @Jiri Dostal_ `[patchman-ui]`

- **2020-06-01**: fix(select-all): Fix select page number for affected systems _by @Jiri Dostal_ `[patchman-ui]`

- **2020-06-01**: chore(release): 0.11.4 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-06-02**: fix(empty-state): Fix incorrect empty state _by @Jiri Dostal_ `[patchman-ui]`

- **2020-06-02**: chore(release): 0.11.5 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-06-02**: fix(empty-state): Polish display of empty states _by @Jiri Dostal_ `[patchman-ui]`

- **2020-06-02**: chore(release): 0.11.6 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-06-04**: Update dependencies to versions with fixed tree shaking. _by @Martin Maroši_ `[patchman-ui]`

- **2020-06-04**: Added transform imports config for PF. _by @Martin Maroši_ `[patchman-ui]`

- **2020-06-04**: Added node env to eslint config _by @Martin Maroši_ `[patchman-ui]`

- **2020-06-04**: Use absolute import paths to CJS modules. _by @Martin Maroši_ `[patchman-ui]`

- **2020-06-04**: Removed async component and use React.lazy instead. _by @Martin Maroši_ `[patchman-ui]`

- **2020-06-04**: Allow multiple vendors chunks. _by @Martin Maroši_ `[patchman-ui]`

- **2020-06-04**: Fixed wrong imports of notifications modules. _by @Martin Maroši_ `[patchman-ui]`

- **2020-06-04**: Add PackageData in form of JSONB _by @Michal Hornicky_ `[patchman-engine]`

- **2020-06-04**: Get tests compiling _by @Michal Hornicky_ `[patchman-engine]`

- **2020-06-04**: Create basic endpoint for package data per system, fix doc generation _by @Michal Hornicky_ `[patchman-engine]`

- **2020-06-05**: Add experimental package-systems endpoint _by @Michal Hornicky_ `[patchman-engine]`

- **2020-06-05**: feat(Build tool): Use rollup for lib build _by @Karel Hala_ `[patchman-ui]`

- **2020-06-08**: Use newer version and treeshake lodash _by @Karel Hala_ `[patchman-ui]`

- **2020-06-08**: Update snapshots to match new FEC _by @Karel Hala_ `[patchman-ui]`

- **2020-06-08**: chore(release): 0.12.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-06-09**: feat(filtering): introduce persistent filtering _by @Jiri Dostal_ `[patchman-ui]`

- **2020-06-09**: updated go deps _by @Josef Hak_ `[patchman-engine]`

- **2020-06-09**: Sync package data for advisories _by @Michal Hornicky_ `[patchman-engine]`

- **2020-06-09**: added support to run manager component on local OS _by @Josef Hak_ `[patchman-engine]`

- **2020-06-09**: changed account number in dev/scripts/env.sh to "1" _by @Josef Hak_ `[patchman-engine]`

- **2020-06-09**: fixed dev/scripts/*.sh example scripts _by @Josef Hak_ `[patchman-engine]`

- **2020-06-11**: chore(lint): fix lint _by @Jiri Dostal_ `[patchman-ui]`

- **2020-06-11**: chore(release): 0.13.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-06-13**: Fix tests _by @Michal Hornicky_ `[patchman-engine]`

- **2020-06-13**: Add tests for package endpoints _by @Michal Hornicky_ `[patchman-engine]`

- **2020-06-15**: fix(manifest): Don't include dev deps in the manifest _by @Jiri Dostal_ `[patchman-ui]`

- **2020-06-15**: chore(release): 0.13.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-06-16**: Feedback, RPM versioning fixes _by @Michal Hornicky_ `[patchman-engine]`

- **2020-06-16**: fixed TestRunServer test _by @Josef Hak_ `[patchman-engine]`

- **2020-06-16**: added "cve_list" field to advisory_metadata _by @Josef Hak_ `[patchman-engine]`

- **2020-06-16**: Revert removing port from evaluator _by @Michal Hornicky_ `[patchman-engine]`

- **2020-06-16**: added cve_list syncing to vmaas-sync _by @Josef Hak_ `[patchman-engine]`

- **2020-06-16**: Skip invalid NEVRAs in package profiles _by @Michal Hornicky_ `[patchman-engine]`

- **2020-06-17**: Send updates to remediations after evaluation _by @Michal Hornicky_ `[patchman-engine]`

- **2020-06-17**: Add paging & filters to package systems endpoint _by @Michal Hornicky_ `[patchman-engine]`

- **2020-06-17**: feat(sort): sort by applicable_advisories _by @Jiri Dostal_ `[patchman-ui]`

- **2020-06-17**: chore(release): 0.14.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-06-17**: updated advisory detail with CVEs list _by @Josef Hak_ `[patchman-engine]`

- **2020-06-17**: add Jenkins file for smoke tests _by @Martin Kourim_ `[patchman-engine]`

- **2020-06-17**: used postgres.Jsonb to represent advisory_metadata.cve_list _by @Josef Hak_ `[patchman-engine]`

- **2020-06-17**: added advisory CVEs info into the advisories search _by @Josef Hak_ `[patchman-engine]`

- **2020-06-17**: fixed advisories syncing test _by @Josef Hak_ `[patchman-engine]`

- **2020-06-18**: added advisory packages field into the advisory detail _by @Josef Hak_ `[patchman-engine]`

- **2020-06-18**: updated advisory detail test _by @Josef Hak_ `[patchman-engine]`

- **2020-06-18**: updated advisory detail endpoint docs _by @Josef Hak_ `[patchman-engine]`

- **2020-06-18**: feat(nav): remove tabs and use only left navigation _by @Jiri Dostal_ `[patchman-ui]`

- **2020-06-18**: chore(release): 0.15.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-06-19**: Add test for remediations publishing _by @Michal Hornicky_ `[patchman-engine]`

- **2020-06-22**: Add test for missign package data _by @Michal Hornicky_ `[patchman-engine]`

- **2020-06-22**: Add tags table _by @Michal Hornicky_ `[patchman-engine]`

- **2020-06-22**: updated advisory detail CVEs and Packages fields types _by @Josef Hak_ `[patchman-engine]`

- **2020-06-22**: changed error messages to warn for unsupported export types _by @Josef Hak_ `[patchman-engine]`

- **2020-06-23**: added tests for unsupported export type for systems and advisories _by @Josef Hak_ `[patchman-engine]`

- **2020-06-24**: feat(pf4): upgrade PF4 _by @Jiri Dostal_ `[patchman-ui]`

- **2020-06-24**: chore(tests): Update tests after PF4 upgrade _by @Jiri Dostal_ `[patchman-ui]`

- **2020-06-24**: chore(release): 0.16.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-06-24**: Reduce number of tables, add method for filtering _by @Michal Hornicky_ `[patchman-engine]`

- **2020-06-24**: Better error loggging in remediations _by @Michal Hornicky_ `[patchman-engine]`

- **2020-06-24**: Add app name prefix _by @Michal Hornicky_ `[patchman-engine]`

- **2020-06-24**: Fix topic lookup _by @Michal Hornicky_ `[patchman-engine]`

- **2020-06-24**: Merge pull request #280 from semtexzv/fix-remediations-topic [#280] _by @Michal Hornický_ `[patchman-engine]`

- **2020-06-25**: fix(remediation): fix remediation color disabled _by @Jiri Dostal_ `[patchman-ui]`

- **2020-06-25**: chore(release): 0.16.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-06-25**: fix(styles): fix PF4 styles _by @Jiri Dostal_ `[patchman-ui]`

- **2020-06-25**: chore(release): 0.16.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-06-25**: fix(routing): refresh redirects to landing page _by @Jiri Dostal_ `[patchman-ui]`

- **2020-06-25**: chore(release): 0.16.3 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-06-29**: Implement efficient querying for tags add support for tags on advisory systems endpoint _by @Michal Hornicky_ `[patchman-engine]`

- **2020-06-29**: Split store when building to specific file so we have even better treeshaking _by @Karel Hala_ `[patchman-ui]`

- **2020-06-30**: Database changes for consitent package data _by @Michal Hornicky_ `[patchman-engine]`

- **2020-06-30**: feat(system-packages): Introduce tabs to select between advisories and packages _by @Jiri Dostal_ `[patchman-ui]`

- **2020-06-30**: Pass React to inventory so it will properly load _by @Karel Hala_ `[patchman-ui]`


## 2020-07 (41 commits)

- **2020-07-02**: fix(tables): Rename advisories table to TableView _by @Jiri Dostal_ `[patchman-ui]`

- **2020-07-03**: Feedback _by @Michal Hornicky_ `[patchman-engine]`

- **2020-07-05**: feat(inventory): pass react-redux to inventory _by @Karel Hala_ `[patchman-ui]`

- **2020-07-06**: Implement package syncing in vmaas_sync _by @Michal Hornicky_ `[patchman-engine]`

- **2020-07-07**: Pkgdata API _by @Michal Hornicky_ `[patchman-engine]`

- **2020-07-07**: Implement package data evaluation _by @Michal Hornicky_ `[patchman-engine]`

- **2020-07-09**: chore(release): 0.17.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-07-16**: Add package counts to systems _by @Michal Hornicky_ `[patchman-engine]`

- **2020-07-16**: Migrate from host-egress to events topic _by @Michal Hornicky_ `[patchman-engine]`

- **2020-07-16**: Implement basic package fucntionality _by @Jiri Dostal_ `[patchman-ui]`

- **2020-07-21**: Rebase related fixes _by @Michal Hornicky_ `[patchman-engine]`

- **2020-07-21**: Rename package version to EVRA _by @Michal Hornicky_ `[patchman-engine]`

- **2020-07-21**: Cyndi integration _by @Michal Hornicky_ `[patchman-engine]`

- **2020-07-22**: fix(deps): Fix securitiy issue in lodash by updating _by @Jiri Dostal_ `[patchman-ui]`

- **2020-07-22**: chore(release): 0.17.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-07-22**: Fixing review issues _by @Michal Hornicky_ `[patchman-engine]`

- **2020-07-22**: fix(style): Update style of tabs _by @Jiri Dostal_ `[patchman-ui]`

- **2020-07-23**: feat(advisoryDetail): Add information about CVEs _by @Jiri Dostal_ `[patchman-ui]`

- **2020-07-23**: chore(release): 0.18.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-07-23**: Dedup strings _by @Michal Hornicky_ `[patchman-engine]`

- **2020-07-24**: Merge pull request #290 from semtexzv/cyndi [#290] _by @Michal Hornický_ `[patchman-engine]`

- **2020-07-24**: added testing "inventory.hosts" view to test_data.sql _by @Josef Hak_ `[patchman-engine]`

- **2020-07-27**: Fix datatypes _by @Michal Hornicky_ `[patchman-engine]`

- **2020-07-27**: fix(revert): Revert "feat(advisoryDetail): Add information about CVEs" _by @Jiri Dostal_ `[patchman-ui]`

- **2020-07-27**: fix(manifest): Change manifest prefixes _by @Jiri Dostal_ `[patchman-ui]`

- **2020-07-27**: chore(release): 0.18.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-07-28**: Fix syncing errors _by @Michal Hornicky_ `[patchman-engine]`

- **2020-07-28**: Perf improvements for syncing packages _by @Michal Hornicky_ `[patchman-engine]`

- **2020-07-28**: fix(breadcrumbs): fix breadcrumbs styling _by @Jiri Dostal_ `[patchman-ui]`

- **2020-07-28**: chore(tests): update tests _by @Jiri Dostal_ `[patchman-ui]`

- **2020-07-28**: Fix bad merge, Rebase on top of master _by @Michal Hornicky_ `[patchman-engine]`

- **2020-07-28**: chore(release): 0.18.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-07-28**: Merge pull request #286 from RedHatInsights/pkgdata [#286] _by @Michal Hornický_ `[patchman-engine]`

- **2020-07-28**: Fix multiple installed packages handling _by @Michal Hornicky_ `[patchman-engine]`

- **2020-07-28**: Make package downloads size limited _by @Michal Hornicky_ `[patchman-engine]`

- **2020-07-28**: Re-parse vmaas response NEVRAs to preserve common format _by @Michal Hornicky_ `[patchman-engine]`

- **2020-07-29**: Merge pull request #297 from semtexzv/pkg-download-limit [#297] _by @Michal Hornický_ `[patchman-engine]`

- **2020-07-29**: Fix parsing of NEVRAs _by @Michal Hornicky_ `[patchman-engine]`

- **2020-07-29**: Fix system deletion _by @Michal Hornicky_ `[patchman-engine]`

- **2020-07-30**: Perf improvements, prometheus metrics, regenerated docs _by @Michal Hornicky_ `[patchman-engine]`

- **2020-07-31**: Add pagination, sort, per page etc _by @Jiri Dostal_ `[patchman-ui]`


## 2020-08 (57 commits)

- **2020-08-05**: Fix tests _by @Michal Hornicky_ `[patchman-engine]`

- **2020-08-05**: Pass pf react core to inventory to properly render inventory _by @Karel Hala_ `[patchman-ui]`

- **2020-08-10**: chore(UnitTests): for Reducers/InventoryEntitiesReducer + SystemDetailStore _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-08-10**: Update README.md _by @Muslimjon_ `[patchman-ui]`

- **2020-08-10**: Update README.md _by @Muslimjon_ `[patchman-ui]`

- **2020-08-11**: chore(UnitTests): PresentationalComponents/InfoBox + RegisterPage + Snippets _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-08-12**: Project cyndi tags implementation _by @Michal Hornicky_ `[patchman-engine]`

- **2020-08-12**: added testing "cyndi" data to "inventory.hosts_v1_0" table _by @Josef Hak_ `[patchman-engine]`

- **2020-08-12**: Merge pull request #293 from Josca/cyndi-test-data [#293] _by @Michal Hornický_ `[patchman-engine]`

- **2020-08-13**: chore(unit tests): for api.js and DataMappers.js` _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-08-13**: fix(api) res.json() is not a fucntion error fix _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-08-13**: updated manifest prefix _by @Josef Hak_ `[patchman-engine]`

- **2020-08-13**: fix(packages): add search and fix select all _by @Jiri Dostal_ `[patchman-ui]`

- **2020-08-13**: updated manifest generator to omit kernel rpm _by @Josef Hak_ `[patchman-engine]`

- **2020-08-17**: fix (docs): openapi specs corrected _by @Ales Kasparek_ `[patchman-engine]`

- **2020-08-17**: fix (fix): search attribute added to system entrypoints _by @Ales Kasparek_ `[patchman-engine]`

- **2020-08-17**: added openapi.json file checker, added to test pipeline _by @Josef Hak_ `[patchman-engine]`

- **2020-08-17**: Merge pull request #302 from Josca/manifest [#302] _by @Michal Hornický_ `[patchman-engine]`

- **2020-08-17**: Fix ISE when searching for packages _by @Michal Hornicky_ `[patchman-engine]`

- **2020-08-17**: Add status endpoint _by @Michal Hornicky_ `[patchman-engine]`

- **2020-08-18**: feat(packages): add filtering and refactor usage on other talbes _by @Jiri Dostal_ `[patchman-ui]`

- **2020-08-18**: fix(packages): Update placeholder for search _by @Jiri Dostal_ `[patchman-ui]`

- **2020-08-18**: chore(tests): update tests _by @Jiri Dostal_ `[patchman-ui]`

- **2020-08-18**: chore(files): add missing files _by @Jiri Dostal_ `[patchman-ui]`

- **2020-08-18**: chore(release): 0.19.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-08-19**: changed error log to warn when listener processes unknonw msg type _by @Josef Hak_ `[patchman-engine]`

- **2020-08-19**: fixed TestPackagesSearch _by @Josef Hak_ `[patchman-engine]`

- **2020-08-19**: added remediations test _by @Josef Hak_ `[patchman-engine]`

- **2020-08-19**: added updatable package info into remediations state publishing _by @Josef Hak_ `[patchman-engine]`

- **2020-08-20**: chore(package): Update remediation components _by @Jiri Dostal_ `[patchman-ui]`

- **2020-08-20**: fix(Icons): applicable advisories icons according to mockups _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-08-20**: Add check for updating package counts _by @Michal Hornicky_ `[patchman-engine]`

- **2020-08-20**: fix(tableFooter): footer misalignment fixed _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-08-21**: fix(AffectedSystems): header has space underneath _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-08-21**: chore(release): 0.19.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-08-21**: fix(AnsibleIcon) color changed to white _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-08-24**: Basic dynamic implementation of packages endpoint _by @Michal Hornicky_ `[patchman-engine]`

- **2020-08-24**: fix(ChromeNav) navigation is fixed on page change _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-08-25**: chore(update deps): chrome package update _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-08-25**: Use newer version of utils package to include all remediations dependencies _by @Karel Hala_ `[patchman-ui]`

- **2020-08-25**: Merge pull request #246 from mkholjuraev/update [#246] _by @Muslimjon_ `[patchman-ui]`

- **2020-08-26**: Merge pull request #247 from karelhala/remediations-utils-imports [#247] _by @Muslimjon_ `[patchman-ui]`

- **2020-08-26**: Update flags, add flag for project cyndi provided tags _by @Michal Hornicky_ `[patchman-engine]`

- **2020-08-26**: chore(deps-dev): bump @redhat-cloud-services/frontend-components-config _by @dependabot-preview[bot]_ `[patchman-ui]`

- **2020-08-26**: added support for custom rbac roles _by @Josef Hak_ `[patchman-engine]`

- **2020-08-27**: fix(select): disallow selection of non-updatable packages _by @Jiri Dostal_ `[patchman-ui]`

- **2020-08-27**: added more testing packages _by @Josef Hak_ `[patchman-engine]`

- **2020-08-27**: fix(notifications): notifications has cancal button _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-08-27**: chore(deps): bump @patternfly/react-core from 4.32.1 to 4.47.0 _by @dependabot-preview[bot]_ `[patchman-ui]`

- **2020-08-28**: chore(audit): audit packages _by @Jiri Dostal_ `[patchman-ui]`

- **2020-08-28**: chore(release): 0.19.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-08-28**: Merge pull request #252 from RedHatInsights/select-update-only [#252] _by @Muslimjon_ `[patchman-ui]`

- **2020-08-28**: chore(release): 0.19.3 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-08-31**: feat(SystemPackages):  empty state is added _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-08-31**: feat(packages-count): Count of packages on system lists _by @Jiri Dostal_ `[patchman-ui]`

- **2020-08-31**: chore(release): 0.20.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-08-31**: fix(LastSeen): sorting bug fix by displaying last_upload instead of updated _by @Muslimjon Kholjuraev_ `[patchman-ui]`


## 2020-09 (79 commits)

- **2020-09-01**: feat(packages-remediation): Allow to remediate packages _by @Jiri Dostal_ `[patchman-ui]`

- **2020-09-01**: chore(test): update tests _by @Jiri Dostal_ `[patchman-ui]`

- **2020-09-01**: Merge pull request #257 from RedHatInsights/audit-packages [#257] _by @Muslimjon_ `[patchman-ui]`

- **2020-09-01**: Merge pull request #261 from RedHatInsights/count-packages [#261] _by @Muslimjon_ `[patchman-ui]`

- **2020-09-01**: chore(release): 0.21.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-09-01**: Merge pull request #268 from RedHatInsights/rem-packages [#268] _by @Muslimjon_ `[patchman-ui]`

- **2020-09-01**: chore(release): 0.22.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-09-02**: chore(deps): bump @redhat-cloud-services/frontend-components-utilities _by @dependabot-preview[bot]_ `[patchman-ui]`

- **2020-09-02**: chore(unit tests) AdvisoryHeader, AdvisoryType, TableFooter, TableView, PublishDateFilter _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-09-03**: Merge pull request #251 from RedHatInsights/dependabot/npm_and_yarn/redhat-cloud-services/frontend-components-config-2.1.6 [#251] _by @Muslimjon_ `[patchman-ui]`

- **2020-09-03**: Merge pull request #254 from RedHatInsights/dependabot/npm_and_yarn/patternfly/react-core-4.47.0 [#254] _by @Muslimjon_ `[patchman-ui]`

- **2020-09-03**: chore(deps): bump @patternfly/react-table from 4.12.1 to 4.16.7 _by @dependabot-preview[bot]_ `[patchman-ui]`

- **2020-09-03**: Merge pull request #255 from RedHatInsights/dependabot/npm_and_yarn/patternfly/react-table-4.16.7 [#255] _by @Muslimjon_ `[patchman-ui]`

- **2020-09-03**: chore(deps): bump @patternfly/react-icons from 4.3.5 to 4.7.4 _by @dependabot-preview[bot]_ `[patchman-ui]`

- **2020-09-03**: Merge pull request #256 from RedHatInsights/dependabot/npm_and_yarn/patternfly/react-icons-4.7.4 [#256] _by @Muslimjon_ `[patchman-ui]`

- **2020-09-03**: fix(SystemPackages): Empty state added for empty table in Systems packages page _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-09-03**: chore(release): 0.22.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-09-03**: Add missing routes _by @Michal Hornicky_ `[patchman-engine]`

- **2020-09-04**: chore(Unit Tests): WithLoader.js _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-09-07**: Link between advisories and packages _by @Michal Hornicky_ `[patchman-engine]`

- **2020-09-07**: added testing system_packages items _by @Josef Hak_ `[patchman-engine]`

- **2020-09-07**: added filter[updatable] bool query param to system_packages _by @Josef Hak_ `[patchman-engine]`

- **2020-09-08**: feat (vmaas_sync): Cyndi metrics _by @Ales Kasparek_ `[patchman-engine]`

- **2020-09-08**: Implement dynamic advisory account data calculation _by @Michal Hornicky_ `[patchman-engine]`

- **2020-09-08**: Implement tag filters on additional systems endpoints _by @Michal Hornicky_ `[patchman-engine]`

- **2020-09-08**: added updates due to lint _by @Josef Hak_ `[patchman-engine]`

- **2020-09-08**: Update page sizes _by @Michal Hornicky_ `[patchman-engine]`

- **2020-09-08**: chore(unit test) SmartComponents/ Advisory.js + AdvisoryDetail.js + AffectedSystems.js _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-09-08**: fix(RemediationIcon) color consistancy fix _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-09-08**: feat(filter): Filter by upgradable on packages _by @Jiri Dostal_ `[patchman-ui]`

- **2020-09-08**: feat(packages): Add updatable column to package view _by @Jiri Dostal_ `[patchman-ui]`

- **2020-09-08**: chore(tests): update snapshots _by @Jiri Dostal_ `[patchman-ui]`

- **2020-09-08**: chore(release): 0.23.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-09-09**: detect kafka stuck from kafka error log _by @Josef Hak_ `[patchman-engine]`

- **2020-09-09**: updated kafka-go lib _by @Josef Hak_ `[patchman-engine]`

- **2020-09-09**: Change kafka implementation _by @Michal Hornicky_ `[patchman-engine]`

- **2020-09-09**: Fix tests _by @Michal Hornicky_ `[patchman-engine]`

- **2020-09-10**: chore(release): 0.24.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-09-10**: Fix listener _by @Michal Hornicky_ `[patchman-engine]`

- **2020-09-10**: Fix counting unknown packages _by @Michal Hornicky_ `[patchman-engine]`

- **2020-09-11**: Fix performance issues with loading packages _by @Michal Hornicky_ `[patchman-engine]`

- **2020-09-14**: Update table config for package data tables _by @Michal Hornicky_ `[patchman-engine]`

- **2020-09-14**: Skip unneccesary updates in package evaluation _by @Michal Hornicky_ `[patchman-engine]`

- **2020-09-14**: chore(Remediation) unused Remediation.js removed _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-09-14**: fix(env) depricated isinsights.chrome.navigation([]); is deleted _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-09-15**: Don't log system ids in vmaas sync _by @Michal Hornicky_ `[patchman-engine]`

- **2020-09-15**: Add support for nested query maps _by @Michal Hornicky_ `[patchman-engine]`

- **2020-09-15**: chore(chrome): Disable global filter _by @Jiri Dostal_ `[patchman-ui]`

- **2020-09-15**: fix(AffectedSystems) remediation error: System must be of type string fix _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-09-15**: chore(release): 0.24.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-09-16**: added metric counter for reporter type _by @Josef Hak_ `[patchman-engine]`

- **2020-09-16**: added system_platform.reporter column _by @Josef Hak_ `[patchman-engine]`

- **2020-09-16**: added reporter info storing in listener, updated tests _by @Josef Hak_ `[patchman-engine]`

- **2020-09-17**: chore(unit test): SmartComponents/ SystemDetail.js + InventoryPage.js + SystemAdvisories.js _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-09-21**: feat (manager): endopiont to return single package _by @Ales Kasparek_ `[patchman-engine]`

- **2020-09-21**: feat(GlobalFilter): global filtering is enabled _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-09-22**: use latest_only option for vmaas api _by @Michael Mraka_ `[patchman-engine]`

- **2020-09-22**: Fix tag parsing regex _by @Michal Hornicky_ `[patchman-engine]`

- **2020-09-23**: added database size and table sizes metrics _by @Josef Hak_ `[patchman-engine]`

- **2020-09-23**: added database processes metric _by @Josef Hak_ `[patchman-engine]`

- **2020-09-24**: Mark systems table as full view _by @Karel Hala_ `[patchman-ui]`

- **2020-09-24**: added grafana dashboard configmap _by @Josef Hak_ `[patchman-engine]`

- **2020-09-25**: feat(Packages View) pakcages view page is added _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-09-25**: Revert "Fix tests" _by @Michal Hornicky_ `[patchman-engine]`

- **2020-09-27**: fix(SystemDetails): infinite re-render is fixed _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-09-28**: Add summary to package data _by @Michal Hornicky_ `[patchman-engine]`

- **2020-09-28**: Add summary to package data _by @Michal Hornicky_ `[patchman-engine]`

- **2020-09-28**: Add summary to package data _by @Michal Hornicky_ `[patchman-engine]`

- **2020-09-28**: Docs fixes _by @Michal Hornicky_ `[patchman-engine]`

- **2020-09-29**: Add rh_account_id to system_package _by @Michal Hornicky_ `[patchman-engine]`

- **2020-09-29**: Migration to partition the system_package table _by @Michal Hornicky_ `[patchman-engine]`

- **2020-09-29**: chore(release): 0.24.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-09-29**: removed out-of-date TODO comment _by @Josef Hak_ `[patchman-engine]`

- **2020-09-29**: added upload-evaluation delay metric _by @Josef Hak_ `[patchman-engine]`

- **2020-09-30**: fixed docker-compose.yml kafka setup _by @Josef Hak_ `[patchman-engine]`

- **2020-09-30**: added metrics help strings _by @Josef Hak_ `[patchman-engine]`

- **2020-09-30**: removed duplicit metric from vmaas_sync _by @Josef Hak_ `[patchman-engine]`

- **2020-09-30**: Revert "added support for custom rbac roles" _by @Josef Hak_ `[patchman-engine]`

- **2020-09-30**: Fix crash in auth middleware _by @Michal Hornicky_ `[patchman-engine]`


## 2020-10 (127 commits)

- **2020-10-01**: Fix missed usages of account names instead of IDs _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-01**: chore(release): 0.25.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-10-01**: Fixed issue with evaluation _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-01**: added ENABLE_ADVISORY_ANALYSIS and ENABLE_PACKAGE_ANALYSIS evaluator options _by @Josef Hak_ `[patchman-engine]`

- **2020-10-01**: added database-size, database-processes, reporter metrics to grafana _by @Josef Hak_ `[patchman-engine]`

- **2020-10-01**: added ENABLE_RECALC_MESSAGES_SEND option to vmaas_sync _by @Josef Hak_ `[patchman-engine]`

- **2020-10-01**: added ENABLE_BYPASS option to evaluator _by @Josef Hak_ `[patchman-engine]`

- **2020-10-01**: fixed wrong delete duration measurement in listener _by @Josef Hak_ `[patchman-engine]`

- **2020-10-01**: added ENABLE_BYPASS option to listener _by @Josef Hak_ `[patchman-engine]`

- **2020-10-02**: fix(packages): reflect most recent mockups on system-packages _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-02**: chore(tests): update tests _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-02**: chore(release): 0.25.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-10-05**: fix(packages): Fix not updating entity _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-05**: chore(release): 0.25.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-10-05**: Fix openapi spec for tags _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-06**: chore(packages): align headers with mocks _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-06**: fix(packages): fix loading of packages _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-06**: fix(titles): add dynamic titles _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-06**: Better check to skip writing package updates, reduce loaded columns _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-06**: Remove unused tags code _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-07**: Update delete_system to take partitioning into account: _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-07**: Skip evaluation if message is older than last_evaluation _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-07**: fix(packages): make table compact _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-07**: fix(packages): Trigger global filter _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-07**: grafana update _by @Josef Hak_ `[patchman-engine]`

- **2020-10-08**: fix(APP): fix SAP error in App _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-08**: chore(release): 0.25.3 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-10-08**: added advisory name to vmaas-sync error log _by @Josef Hak_ `[patchman-engine]`

- **2020-10-08**: added some debuging logs to vmaas-sync sync process _by @Josef Hak_ `[patchman-engine]`

- **2020-10-08**: feat(package-detail): implement package detail page _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-08**: feat(packages): link package details _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-08**: set version of generated API doc _by @Michael Mraka_ `[patchman-engine]`

- **2020-10-09**: fix(packages): Add padding to the tabs _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-09**: chore(release): 0.25.4 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-10-09**: Wrap mutex locking + unlocking in inline function _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-09**: Start to work on turnpike integration & admin interface _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-10**: chore(deps): bump @redhat-cloud-services/frontend-components-notifications _by @dependabot-preview[bot]_ `[patchman-ui]`

- **2020-10-12**: Optimize packages query _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-12**: Add docs for workload filters _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-12**: Fix vmaas_sync crashing on ws errors _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-12**: added collapsible rows to grafana _by @Josef Hak_ `[patchman-engine]`

- **2020-10-12**: chore(tests): Update tests _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-12**: chore(release): 0.26.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-10-13**: fix grafana config _by @Josef Hak_ `[patchman-engine]`

- **2020-10-13**: added search to packages endpoint _by @Josef Hak_ `[patchman-engine]`

- **2020-10-13**: added filter support to export system handler _by @Josef Hak_ `[patchman-engine]`

- **2020-10-13**: Update package data on sync _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-13**: change inventory_id type to match inventory.hosts.id _by @Michael Mraka_ `[patchman-engine]`

- **2020-10-13**: moved systems export endpoint to separate file _by @Josef Hak_ `[patchman-engine]`

- **2020-10-13**: refactor(affectedSystems): refactor affected systems _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-13**: updated systems api docs _by @Josef Hak_ `[patchman-engine]`

- **2020-10-13**: updated schema with partitoned system_platform and system_advisories _by @Michael Mraka_ `[patchman-engine]`

- **2020-10-13**: split advisories - advisories_export endpoints _by @Josef Hak_ `[patchman-engine]`

- **2020-10-13**: chore(release): 0.26.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-10-13**: extended tags testing _by @Josef Hak_ `[patchman-engine]`

- **2020-10-14**: fix(globalFilter): remove Patch scope _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-14**: chore(release): 0.26.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-10-14**: Add package_latest_cache view _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-14**: removed [enabled] filter from systems endpoint _by @Josef Hak_ `[patchman-engine]`

- **2020-10-14**: Use package_latest_cache for better packages endpoint _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-14**: Add PR review changes _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-14**: Add PR review changes _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-14**: updated grafana - cyndi chart, upload-eval delay _by @Josef Hak_ `[patchman-engine]`

- **2020-10-15**: Remove duplicate params _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-15**: Fix missing non-updatable packages _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-15**: feat(systemPackages): add systems view per package _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-15**: chore(release): 0.27.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-10-15**: fix(SAP): fix SAP filtering, don't encode url _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-15**: chore(release): 0.27.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-10-15**: Remove the need for the packages endpoint to visit whole partition _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-19**: updated systems_export handler and tags test _by @Josef Hak_ `[patchman-engine]`

- **2020-10-19**: fix(database): Fix ERROR message _by @AlesKas_ `[patchman-engine]`

- **2020-10-19**: fixed wrong offset handling returning 400 - bad request _by @Josef Hak_ `[patchman-engine]`

- **2020-10-19**: added tests to ensure invalid offset handling _by @Josef Hak_ `[patchman-engine]`

- **2020-10-19**: updated dev test data according to partitioning _by @Michael Mraka_ `[patchman-engine]`

- **2020-10-20**: fix(SAP): fix offset when setting global filter _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-20**: Fix invalid tags parsing: _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-20**: chore(release): 0.27.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-10-20**: fix(SAP): enable SID filtering _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-20**: added ENABLE_STALE_SYSTEM_EVALUATION to evaluator _by @Josef Hak_ `[patchman-engine]`

- **2020-10-20**: chore(release): 0.27.3 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-10-20**: fix(packageDetail): affected systems by a package are now clickable _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-20**: Return tags in links response _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-20**: chore(packages): enable package only in beta _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-20**: chore(tests): fix tests _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-20**: chore(release): 0.27.4 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-10-20**: fix(packageDetail): display latest version when same as installed _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-20**: fix(api): reduce size of initial "/systems" response _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-20**: chore(release): 0.27.5 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-10-21**: Store latest evra in system_package index _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-21**: Re-partition system_package _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-21**: fix(SAP): don't trigger update when there is no change _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-21**: chore(release): 0.27.6 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-10-21**: fixing tests according to partitioned schema _by @Michael Mraka_ `[patchman-engine]`

- **2020-10-21**: fix(advisorySystems): deselect of items fixed _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-21**: fix(selectAll): select should take tags into account _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-21**: chore(tests): fix tests _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-21**: chore(release): 0.27.7 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-10-22**: fix(tableView): fix selectAll select page number _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-22**: chore(release): 0.27.8 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-10-22**: Package systems, preload package ids _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-22**: fix(exports): fix exports of components _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-22**: fix(exports): Add new exports to rollup _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-22**: Run unit tests on github workflows _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-22**: chore(release): 0.27.9 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-10-23**: Remove debug enable function from patch component mount _by @Karel Hala_ `[patchman-ui]`

- **2020-10-23**: Delete old system_package table _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-23**: fix(sorting): fix sorting of packageSystems table _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-23**: Trying to get QA to build _by @Michal Hornicky_ `[patchman-engine]`

- **2020-10-23**: fix sap_sids filter length check _by @Josef Hak_ `[patchman-engine]`

- **2020-10-23**: updated tests _by @Josef Hak_ `[patchman-engine]`

- **2020-10-23**: fix(SAP): Fix multiple SID filter result _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-23**: chore(release): 0.27.10 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-10-26**: fix(packages): Fix "isBeta" variable in tests _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-26**: chore(release): 0.27.11 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-10-26**: added metric and log message for unknown reporter _by @Josef Hak_ `[patchman-engine]`

- **2020-10-26**: fix(SID): check for SID existence before checking length _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-26**: chore(release): 0.27.12 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-10-26**: added metric to measure delay between two evaluations _by @Josef Hak_ `[patchman-engine]`

- **2020-10-27**: fix(persistence): Do not persist store when there was an error _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-27**: chore(release): 0.27.13 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-10-27**: fix(url): Fix URL parsing with global filter on _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-27**: chore(release): 0.27.14 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-10-27**: feat(Patch): Patch UI 1.0.0 _by @Jiri Dostal_ `[patchman-ui]`

- **2020-10-27**: chore(release): 1.0.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-10-27**: fixed SystemDeleteHandler using account name instead of account id _by @Josef Hak_ `[patchman-engine]`

- **2020-10-27**: used account names different from account ids _by @Josef Hak_ `[patchman-engine]`


## 2020-11 (41 commits)

- **2020-11-02**: fix system_detail fields: display_name, packages, stale _by @Josef Hak_ `[patchman-engine]`

- **2020-11-04**: return 404 with error message when invalid tag provided _by @Josef Hak_ `[patchman-engine]`

- **2020-11-04**: updated to new patchman-client _by @Michael Mraka_ `[patchman-engine]`

- **2020-11-04**: use updated vmaas client module _by @Michael Mraka_ `[patchman-engine]`

- **2020-11-04**: chore(deps-dev): bump @redhat-cloud-services/frontend-components-config _by @dependabot-preview[bot]_ `[patchman-ui]`

- **2020-11-06**: added project version into metrics _by @Michael Mraka_ `[patchman-engine]`

- **2020-11-06**: add option to show current project version _by @Michael Mraka_ `[patchman-engine]`

- **2020-11-06**: added warning to not edit version manually _by @Michael Mraka_ `[patchman-engine]`

- **2020-11-06**: increased testing db POSTGRESQL_MAX_CONNECTIONS to 250 _by @Josef Hak_ `[patchman-engine]`

- **2020-11-09**: added db-data volume to docker-compose example _by @Josef Hak_ `[patchman-engine]`

- **2020-11-10**: fix(SID): SAP should work together with SID _by @Jiri Dostal_ `[patchman-ui]`

- **2020-11-10**: chore(release): 1.0.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-11-10**: chore(release): 1.0.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-11-10**: fix volume permissions in podman _by @Michael Mraka_ `[patchman-engine]`

- **2020-11-11**: chore(optimization): onSelect function is optimized _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-11-11**: chore(tests): updates snapshots _by @Jiri Dostal_ `[patchman-ui]`

- **2020-11-11**: chore(tests): fix test fails _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-11-11**: chore(optimization) some shared presentational components _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-11-12**: fix(ConditionalFilter): radio typw filter empty value is fixed _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-11-12**: added forgotten metric registration _by @Josef Hak_ `[patchman-engine]`

- **2020-11-13**: chore(release): 1.0.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-11-13**: generate_docs.sh now has mandatory options _by @Michael Mraka_ `[patchman-engine]`

- **2020-11-13**: added EXCLUDED_REPORTERS option to the listener _by @Josef Hak_ `[patchman-engine]`

- **2020-11-14**: task(Inventroy) gate inventory calls before fetch _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-11-16**: chore(unit tests): for System.js, SystemListAssets.js SystemPackages.js _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-11-18**: feat(Errors) handle 403 error _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-11-19**: Add search into query parameters _by @Michal Hornicky_ `[patchman-engine]`

- **2020-11-20**: fixed skipped reporter metric increment _by @Josef Hak_ `[patchman-engine]`

- **2020-11-20**: fixed setting up testing account _by @Josef Hak_ `[patchman-engine]`

- **2020-11-23**: chore(tests): fix tests _by @Jiri Dostal_ `[patchman-ui]`

- **2020-11-24**: feat(OUIA) enable OUIA for Patch _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-11-24**: added reporter id to testing db data _by @Josef Hak_ `[patchman-engine]`

- **2020-11-25**: chore(InventoryAccess): change logic of gating Invetory request _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-11-27**: fix(NoSystemData): fix misleading wording _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-11-30**: chore(deps): bump @redhat-cloud-services/frontend-components-remediations _by @dependabot-preview[bot]_ `[patchman-ui]`

- **2020-11-30**: log slow statements in test/ci/qa _by @Michael Mraka_ `[patchman-engine]`

- **2020-11-30**: chore(deps): bump @redhat-cloud-services/frontend-components _by @dependabot-preview[bot]_ `[patchman-ui]`

- **2020-11-30**: chore(deps): bump @patternfly/react-icons from 4.7.4 to 4.7.18 _by @dependabot-preview[bot]_ `[patchman-ui]`

- **2020-11-30**: added logs and metric for marking systems "stale" _by @Josef Hak_ `[patchman-engine]`

- **2020-11-30**: chore(deps): bump @patternfly/react-table from 4.16.7 to 4.19.24 _by @dependabot-preview[bot]_ `[patchman-ui]`

- **2020-11-30**: mark non-inventory systems culled _by @Josef Hak_ `[patchman-engine]`


## 2020-12 (43 commits)

- **2020-12-01**: Publish admin api from vmaas_sync _by @Michal Hornicky_ `[patchman-engine]`

- **2020-12-01**: Packages export handler _by @Michal Hornicky_ `[patchman-engine]`

- **2020-12-01**: feat(internationalization): install dependencies and adopt intl _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-12-01**: chore(tests): update tests _by @Jiri Dostal_ `[patchman-ui]`

- **2020-12-01**: chore(release): 1.0.3 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-12-01**: fix(globalFilter) error handling is fixed _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-12-02**: feat(internationalization): autogenerate translations _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-12-02**: chore(select): Revert "chore(tests): fix test fails" _by @Jiri Dostal_ `[patchman-ui]`

- **2020-12-02**: fix(select): fix selecting of advisories _by @Jiri Dostal_ `[patchman-ui]`

- **2020-12-02**: updated schema _by @Josef Hak_ `[patchman-engine]`

- **2020-12-03**: chore(release): 1.0.4 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-12-03**: Add endpoint for vuln errata linking _by @Michal Hornicky_ `[patchman-engine]`

- **2020-12-03**: Update docs _by @Michal Hornicky_ `[patchman-engine]`

- **2020-12-04**: Add missing params to system_packages _by @Michal Hornicky_ `[patchman-engine]`

- **2020-12-04**: Add search to ListMeta _by @Michal Hornicky_ `[patchman-engine]`

- **2020-12-04**: Add missing paging params _by @Michal Hornicky_ `[patchman-engine]`

- **2020-12-07**: chore(timeout): remove timeout from API calls _by @Jiri Dostal_ `[patchman-ui]`

- **2020-12-08**: chore(release): 1.1.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-12-08**: added PR template _by @Josef Hak_ `[patchman-engine]`

- **2020-12-08**: added namespace to Grafana board _by @Josef Hak_ `[patchman-engine]`

- **2020-12-09**: fix(treeshaking): Improve treeshaking of imported patch detail _by @Karel Hala_ `[patchman-ui]`

- **2020-12-10**: chore(release): 1.1.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-12-10**: chore(App.js) refactor to functional component _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-12-10**: added inventory.hosts join to system_platform-based queries in endpoints _by @Josef Hak_ `[patchman-engine]`

- **2020-12-11**: New spec check and improve unittest workflow _by @Nikhil Dhandre_ `[patchman-engine]`

- **2020-12-11**: fixed "delete_culled_systems", and added "deleted_limit" argument _by @Josef Hak_ `[patchman-engine]`

- **2020-12-11**: bump 1.4.2 _by @Josef Hak_ `[patchman-engine]`

- **2020-12-11**: script to generate random test data _by @Michael Mraka_ `[patchman-engine]`

- **2020-12-15**: feat(internationalization): translate presentational components _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-12-15**: chore(release): 1.2.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-12-15**: added optional SCHEMA_MIGRATION=int db_admin cfg variable _by @Josef Hak_ `[patchman-engine]`

- **2020-12-15**: feat(SystemDetails) make empty system advisories and packages tables consistant _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-12-16**: fix(nav): fix novaigation updates _by @Jiri Dostal_ `[patchman-ui]`

- **2020-12-17**: feat(internationalization): translate smart components _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-12-17**: added optional env var RESET_SCHEMA to database-admin _by @Josef Hak_ `[patchman-engine]`

- **2020-12-17**: feat(Packages): add export feature _by @Muslimjon Kholjuraev_ `[patchman-ui]`

- **2020-12-17**: chore(release): 1.2.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-12-18**: chore(release): 1.3.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-12-18**: chore(tests): update snapshots _by @Jiri Dostal_ `[patchman-ui]`

- **2020-12-21**: chore(release): 1.3.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2020-12-21**: chore(lint): update linter and enforce sort-keys on Messages _by @Jiri Dostal_ `[patchman-ui]`

- **2020-12-21**: chore(lint): follow lint rules for Messages _by @Jiri Dostal_ `[patchman-ui]`

- **2020-12-22**: added ENABLE_SYNC_ON_START and ENABLE_RECALC_ON_START to vmaas-sync _by @Josef Hak_ `[patchman-engine]`


## 2021-01 (95 commits)

- **2021-01-04**: Log warnings client-side _by @Michal Hornicky_ `[patchman-engine]`

- **2021-01-04**: Support for checking tag presence without value _by @Michal Hornicky_ `[patchman-engine]`

- **2021-01-04**: removed unused DB_INITIALIZED var from update.sh script _by @Josef Hak_ `[patchman-engine]`

- **2021-01-04**: bump 1.4.3 _by @Josef Hak_ `[patchman-engine]`

- **2021-01-05**: system_platform migration - new table and data _by @Michael Mraka_ `[patchman-engine]`

- **2021-01-05**: system_advisories migration - new table and data _by @Michael Mraka_ `[patchman-engine]`

- **2021-01-05**: system_repo migration - new column and data _by @Michael Mraka_ `[patchman-engine]`

- **2021-01-05**: table migration - table switch _by @Michael Mraka_ `[patchman-engine]`

- **2021-01-06**: clear prcedures in sql code _by @Michael Mraka_ `[patchman-engine]`

- **2021-01-06**: removed unpartitoned tables _by @Michael Mraka_ `[patchman-engine]`

- **2021-01-06**: chore(packages): Update PF _by @Jiri Dostal_ `[patchman-ui]`

- **2021-01-06**: chore(tests): update snapshots _by @Jiri Dostal_ `[patchman-ui]`

- **2021-01-06**: fix TestRunServer for podman-compose _by @Josef Hak_ `[patchman-engine]`

- **2021-01-07**: chore(deps): bump @redhat-cloud-services/frontend-components-remediations _by @dependabot-preview[bot]_ `[patchman-ui]`

- **2021-01-07**: feat(Packages):  sort by updatable packages by default _by @mkholjur_ `[patchman-ui]`

- **2021-01-07**: fix(breadcrumb): Fix breadcrumb for consistency _by @Jiri Dostal_ `[patchman-ui]`

- **2021-01-07**: chore(snapshots): update snapshots _by @Jiri Dostal_ `[patchman-ui]`

- **2021-01-07**: chore(release): 1.3.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-01-07**: used sql function "delete_culled_systems(integer)", added DELETE_CULLED_SYSTEMS_LIMIT _by @Josef Hak_ `[patchman-engine]`

- **2021-01-07**: chore(app): hot reload is enabled _by @mkholjur_ `[patchman-ui]`

- **2021-01-08**: bump 1.5.0 _by @Josef Hak_ `[patchman-engine]`

- **2021-01-08**: Add reverse advisories-systems view _by @Michal Hornicky_ `[patchman-engine]`

- **2021-01-08**: added version info into the Patch grafana _by @Michael Mraka_ `[patchman-engine]`

- **2021-01-08**: chore(release): 1.4.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-01-11**: chore(deps): bump @redhat-cloud-services/frontend-components-notifications _by @dependabot-preview[bot]_ `[patchman-ui]`

- **2021-01-11**: chore(deps): bump @redhat-cloud-services/frontend-components _by @dependabot-preview[bot]_ `[patchman-ui]`

- **2021-01-11**: chore(deps): bump @redhat-cloud-services/frontend-components-utilities _by @dependabot-preview[bot]_ `[patchman-ui]`

- **2021-01-11**: fixed grafana versionbox on prod _by @Michael Mraka_ `[patchman-engine]`

- **2021-01-11**: chore(release): 1.5.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-01-11**: chore(dependabot): extend the update checking period _by @mkholjur_ `[patchman-ui]`

- **2021-01-12**: added vmaas-sync options ENABLE_CULLED_SYSTEM_DELETE,ENABLE_SYSTEM_STALING _by @Josef Hak_ `[patchman-engine]`

- **2021-01-12**: Add bulk delete _by @Michal Hornicky_ `[patchman-engine]`

- **2021-01-12**: Manually delete yupana systems _by @Michal Hornicky_ `[patchman-engine]`

- **2021-01-12**: bump 1.5.1 _by @Josef Hak_ `[patchman-engine]`

- **2021-01-12**: feat(GenericError): use frontend components generic error component _by @mkholjur_ `[patchman-ui]`

- **2021-01-12**: Reduce bulk delete size in migration 46 _by @Michal Hornicky_ `[patchman-engine]`

- **2021-01-12**: fix(store): remove getState from reducer _by @Jiri Dostal_ `[patchman-ui]`

- **2021-01-13**: chore(test): resolve test issues _by @mkholjur_ `[patchman-ui]`

- **2021-01-13**: chore(release): 1.5.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-01-14**: Remove yupana systems deletion _by @Michal Hornicky_ `[patchman-engine]`

- **2021-01-14**: feat(Advisories): display cve info _by @mkholjur_ `[patchman-ui]`

- **2021-01-15**: System_packages: speed up packages endpoint _by @Michal Hornicky_ `[patchman-engine]`

- **2021-01-15**: updated doc to 1.6.0 release _by @Michael Mraka_ `[patchman-engine]`

- **2021-01-15**: chore(release): 1.6.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-01-15**: chore(release): 1.7.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-01-15**: added "gpg-pubkey" skip to evaluator _by @Josef Hak_ `[patchman-engine]`

- **2021-01-15**: added os-version testing data to test_data.sql _by @Josef Hak_ `[patchman-engine]`

- **2021-01-15**: Added OSName, OSMajor and OSMinor fields to system(s) endpoint _by @Josef Hak_ `[patchman-engine]`

- **2021-01-18**: Update dependencies _by @Martin Marosi_ `[patchman-ui]`

- **2021-01-18**: Adjust application entry files for chrome 2.0 _by @Martin Marosi_ `[patchman-ui]`

- **2021-01-18**: Use chrome 2.0 compatible version of inventory. _by @Martin Marosi_ `[patchman-ui]`

- **2021-01-18**: Fix AdvisoryDetail tests _by @Martin Marosi_ `[patchman-ui]`

- **2021-01-18**: bug(Inventory): hide tabs in Inventory app when there is not systems registered _by @mkholjur_ `[patchman-ui]`

- **2021-01-19**: Add missing joins using rh_account_id _by @Michal Hornicky_ `[patchman-engine]`

- **2021-01-19**: Empty commit _by @Michal Hornicky_ `[patchman-engine]`

- **2021-01-19**: added rh_account_id select to listener _by @Josef Hak_ `[patchman-engine]`

- **2021-01-20**: Fix incorrect onLoad assignment. _by @Martin Marosi_ `[patchman-ui]`

- **2021-01-20**: Query cleanup _by @Michal Hornicky_ `[patchman-engine]`

- **2021-01-20**: added rh_account_id to system_advisories update _by @Josef Hak_ `[patchman-engine]`

- **2021-01-20**: include AccountID into recalc messages _by @Michael Mraka_ `[patchman-engine]`

- **2021-01-20**: bug(pages): page breaking bugs resolved _by @mkholjur_ `[patchman-ui]`

- **2021-01-21**: Bulk delete fix _by @Michal Hornicky_ `[patchman-engine]`

- **2021-01-21**: Add inventory_id index to system_platform _by @Michal Hornicky_ `[patchman-engine]`

- **2021-01-21**: feat(tables): use compact tables for consistency _by @Jiri Dostal_ `[patchman-ui]`

- **2021-01-21**: chore(tests): fix tests _by @Jiri Dostal_ `[patchman-ui]`

- **2021-01-22**: added system_advisories_sum and system_platform_sum _by @Michael Mraka_ `[patchman-engine]`

- **2021-01-22**: app-interface breaks if there are scripts in dasboards _by @Michael Mraka_ `[patchman-engine]`

- **2021-01-22**: Fix routes in docs _by @Michal Hornicky_ `[patchman-engine]`

- **2021-01-22**: added limit to delete_system(inventory_id uuid) sql function _by @Josef Hak_ `[patchman-engine]`

- **2021-01-22**: Update system_platform opt_out trigger _by @Michal Hornicky_ `[patchman-engine]`

- **2021-01-22**: bug(Notifications): 3x the same notifications due to multiple backend call _by @mkholjur_ `[patchman-ui]`

- **2021-01-22**: bug(Notifications): 3x the same notifications due to multiple backend call _by @mkholjur_ `[patchman-ui]`

- **2021-01-25**: chore(release): 1.8.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-01-25**: bug(StatusFilter): clear filters on unmount, solve 2x upgradable label _by @mkholjur_ `[patchman-ui]`

- **2021-01-25**: fix(imports): update imports for TableVariant _by @Jiri Dostal_ `[patchman-ui]`

- **2021-01-25**: chore(release): 1.8.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-01-25**: added CVE count to advisories _by @Josef Hak_ `[patchman-engine]`

- **2021-01-26**: improve cyndi tag metrics query _by @Michael Mraka_ `[patchman-engine]`

- **2021-01-26**: added RBAC */read/write/ perms support _by @Josef Hak_ `[patchman-engine]`

- **2021-01-27**: Fix system detail missing join _by @Michal Hornicky_ `[patchman-engine]`

- **2021-01-27**: bump version _by @Josef Hak_ `[patchman-engine]`

- **2021-01-28**: removed not used APPLY_INVENTORY_HOSTS var in manager _by @Josef Hak_ `[patchman-engine]`

- **2021-01-28**: fix(select): fix selection of packages with duplicate names _by @Jiri Dostal_ `[patchman-ui]`

- **2021-01-28**: chore(tests): update tests _by @Jiri Dostal_ `[patchman-ui]`

- **2021-01-28**: fix(packages): ID should be unique for packages _by @Jiri Dostal_ `[patchman-ui]`

- **2021-01-28**: chore(tests): update tests _by @Jiri Dostal_ `[patchman-ui]`

- **2021-01-28**: feat(Advisories) display cves count, impact, enable sort _by @mkholjur_ `[patchman-ui]`

- **2021-01-28**: fix(notifications): Display notification when remediation is created _by @Jiri Dostal_ `[patchman-ui]`

- **2021-01-28**: fix(SystemDetail): non-existing entity causes UI stuck _by @mkholjur_ `[patchman-ui]`

- **2021-01-28**: removed "package" join from /packages endpoint _by @Josef Hak_ `[patchman-engine]`

- **2021-01-28**: chore(release): 1.8.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-01-29**: SPM-631: performance improvements _by @Michael Mraka_ `[patchman-engine]`

- **2021-01-29**: chore(release): 1.8.3 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-01-29**: bump 1.6.3 _by @Josef Hak_ `[patchman-engine]`

- **2021-01-29**: updated README.md - how to run single test _by @Josef Hak_ `[patchman-engine]`


## 2021-02 (84 commits)

- **2021-02-01**: bump version 1.7.0 _by @Josef Hak_ `[patchman-engine]`

- **2021-02-01**: feat(OS): Add OS column to systems pages _by @Jiri Dostal_ `[patchman-ui]`

- **2021-02-01**: fix(sort): Make sorting of multiple values work out of the box _by @Jiri Dostal_ `[patchman-ui]`

- **2021-02-01**: fix(OS): display No data if column is empty _by @Jiri Dostal_ `[patchman-ui]`

- **2021-02-01**: chore(release): 1.8.4 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-02-01**: chore(release): 1.9.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-02-02**: added system resource to rbac perms _by @Josef Hak_ `[patchman-engine]`

- **2021-02-02**: Fix last migration _by @mhornick_ `[patchman-engine]`

- **2021-02-02**: fix(styles): fix font-size of inventory tables _by @Jiri Dostal_ `[patchman-ui]`

- **2021-02-02**: updated deps _by @Josef Hak_ `[patchman-engine]`

- **2021-02-02**: chore(package): update packages _by @Jiri Dostal_ `[patchman-ui]`

- **2021-02-02**: chore(package): update to components v3 _by @Jiri Dostal_ `[patchman-ui]`

- **2021-02-02**: Fix query bugs _by @mhornick_ `[patchman-engine]`

- **2021-02-02**: chore(release): 1.9.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-02-03**: Cleanu p tag handling _by @mhornick_ `[patchman-engine]`

- **2021-02-04**: added new charts to grafana _by @Josef Hak_ `[patchman-engine]`

- **2021-02-05**: Remove data migration in 59, leave just schema _by @mhornick_ `[patchman-engine]`

- **2021-02-05**: Revert "Remove data migration in 59, leave just schema" _by @mhornick_ `[patchman-engine]`

- **2021-02-05**: enabled swagger in dev mode _by @Josef Hak_ `[patchman-engine]`

- **2021-02-08**: updated kafka-go dep _by @Josef Hak_ `[patchman-engine]`

- **2021-02-08**: fixed testing upload inventory_id to be valid "uuid" _by @Josef Hak_ `[patchman-engine]`

- **2021-02-08**: fixed waiting for finished migration in "wait-for-services.sh" _by @Josef Hak_ `[patchman-engine]`

- **2021-02-09**: fix(bulkSelect): page selecting after selecting all items _by @mkholjur_ `[patchman-ui]`

- **2021-02-09**: optimize re-calc using rh_account_id sort in vmaas_sync _by @Josef Hak_ `[patchman-engine]`

- **2021-02-09**: added missing "when_patched IS NULL" where filter to manager _by @Josef Hak_ `[patchman-engine]`

- **2021-02-09**: updated tests to verify "patched" are excluded _by @Josef Hak_ `[patchman-engine]`

- **2021-02-10**: chore(packages): Enable global packages in prod _by @Jiri Dostal_ `[patchman-ui]`

- **2021-02-11**: chore(release): 1.10.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-02-11**: feat(Filters): reset filters introduced _by @mkholjur_ `[patchman-ui]`

- **2021-02-11**: bug(App.js): change useEffect dependency to location.pathname _by @mkholjur_ `[patchman-ui]`

- **2021-02-11**: SPM-736: store thirdparty system flag _by @Michael Mraka_ `[patchman-engine]`

- **2021-02-11**: add optional DISABLE_CACHE_COUNTS var to manager _by @Josef Hak_ `[patchman-engine]`

- **2021-02-11**: remove patched advisories during the evaluation _by @Josef Hak_ `[patchman-engine]`

- **2021-02-12**: feat(Packages): change remediation identifier to patch-package _by @mkholjur_ `[patchman-ui]`

- **2021-02-12**: skip gpg-pubkey in packages load method _by @Josef Hak_ `[patchman-engine]`

- **2021-02-12**: Fix store already initialized error. _by @Martin Marosi_ `[patchman-ui]`

- **2021-02-12**: SPM-736: evaluate thirdparty repos _by @Michael Mraka_ `[patchman-engine]`

- **2021-02-12**: bump version _by @Josef Hak_ `[patchman-engine]`

- **2021-02-12**: SPM-736: test for thirdparty flag _by @Michael Mraka_ `[patchman-engine]`

- **2021-02-15**: fix(select): after filering _by @mkholjur_ `[patchman-ui]`

- **2021-02-15**: fixed system_advisories counts in system_platform _by @Josef Hak_ `[patchman-engine]`

- **2021-02-15**: extended evaluation test _by @Josef Hak_ `[patchman-engine]`

- **2021-02-16**: added AssertWait testing method and used in tests _by @Josef Hak_ `[patchman-engine]`

- **2021-02-17**: bump version _by @Josef Hak_ `[patchman-engine]`

- **2021-02-17**: feat(advisories): Add checkboxes on advisory list page _by @Jiri Dostal_ `[patchman-ui]`

- **2021-02-17**: add Haberdasher logging command wrapper support _by @Josef Hak_ `[patchman-engine]`

- **2021-02-17**: chore(dependabot): change update check to monthly _by @mkholjur_ `[patchman-ui]`

- **2021-02-18**: feat(remediation): enable bulk remediations for advisory page _by @Jiri Dostal_ `[patchman-ui]`

- **2021-02-18**: SPM-736: let evaluator run tests _by @Michael Mraka_ `[patchman-engine]`

- **2021-02-18**: increased AssertWait timeout in some tests _by @Josef Hak_ `[patchman-engine]`

- **2021-02-18**: feat(remediation): Show animation when remediation is loading _by @Jiri Dostal_ `[patchman-ui]`

- **2021-02-18**: Remove insights route as it can cause page break _by @Karel Hala_ `[patchman-ui]`

- **2021-02-19**: chore(CVEs): fix the url to /cves endpoint _by @mkholjur_ `[patchman-ui]`

- **2021-02-19**: added kafka encryption support (ENABLE_KAFKA_SSL, KAFKA_SSL_CERT) _by @Josef Hak_ `[patchman-engine]`

- **2021-02-19**: chore(release): 1.11.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-02-19**: fix(bulkSelect): Revert "fix(bulkSelect): page selecting after selecting all items" _by @Jiri Dostal_ `[patchman-ui]`

- **2021-02-19**: chore(release): 1.11.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-02-19**: SPM-736: expose thirdparty flag in API _by @Michael Mraka_ `[patchman-engine]`

- **2021-02-19**: added KAFKA_SSL_SKIP_VERIFY to kafka ssl support _by @Josef Hak_ `[patchman-engine]`

- **2021-02-19**: set up kafka encryption in docker-compose env and tests _by @Josef Hak_ `[patchman-engine]`

- **2021-02-22**: use python3 (unversioned alias can't be available) _by @Michael Mraka_ `[patchman-engine]`

- **2021-02-22**: disable zookeeper admin port _by @Michael Mraka_ `[patchman-engine]`

- **2021-02-22**: extended system_package testing data _by @Josef Hak_ `[patchman-engine]`

- **2021-02-22**: fixed system_packages "Updatable" attribute query _by @Josef Hak_ `[patchman-engine]`

- **2021-02-23**: bump version 1.7.5 _by @Josef Hak_ `[patchman-engine]`

- **2021-02-23**: Add RHSMVersion to systems endpoint _by @AlesKas_ `[patchman-engine]`

- **2021-02-23**: SPM-736: thirdparty API flag test _by @Michael Mraka_ `[patchman-engine]`

- **2021-02-23**: added SonarQube script for static code analysis _by @Josef Hak_ `[patchman-engine]`

- **2021-02-23**: feat(remediation): introduce remediation on system list page _by @Jiri Dostal_ `[patchman-ui]`

- **2021-02-23**: feat(InventoryComp): SPM-754 use custom getEntities _by @mkholjur_ `[patchman-ui]`

- **2021-02-23**: fix(CVEs): inconsistencies in cves info modal _by @mkholjur_ `[patchman-ui]`

- **2021-02-24**: Add link to filters documentation in apidoc _by @mhornick_ `[patchman-engine]`

- **2021-02-24**: Add message for bulk evaluation in recalc _by @mhornick_ `[patchman-engine]`

- **2021-02-24**: Remove opt_out _by @mhornick_ `[patchman-engine]`

- **2021-02-24**: Add turnpike auth check _by @mhornick_ `[patchman-engine]`

- **2021-02-24**: remove first_upload info from system_platform, only .go files _by @mkholjur_ `[patchman-engine]`

- **2021-02-24**: added SonarQube script for static code analysis _by @Josef Hak_ `[patchman-ui]`

- **2021-02-24**: updated sonar config project path _by @Josef Hak_ `[patchman-engine]`

- **2021-02-25**: chore(Routes): remove the appended slash in URL _by @mkholjur_ `[patchman-ui]`

- **2021-02-25**: chore(release): 1.11.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-02-25**: chore(release): 1.11.3 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-02-25**: chore(deps): bump @redhat-cloud-services/frontend-components-translations _by @dependabot-preview[bot]_ `[patchman-ui]`

- **2021-02-25**: chore(deps): bump @redhat-cloud-services/frontend-components-remediations _by @dependabot-preview[bot]_ `[patchman-ui]`

- **2021-02-25**: chore(deps): bump @patternfly/react-table from 4.19.45 to 4.23.2 _by @dependabot-preview[bot]_ `[patchman-ui]`


## 2021-03 (68 commits)

- **2021-03-02**: split evaluate message into batches _by @Michael Mraka_ `[patchman-engine]`

- **2021-03-03**: feat(SystemDetail): add third party info _by @mkholjur_ `[patchman-ui]`

- **2021-03-05**: chore(release): 1.12.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-03-08**: SPM-768: use rhsm release if available _by @Michael Mraka_ `[patchman-engine]`

- **2021-03-08**: fix(SearchFilter): wordings are made consistent with platform _by @mkholjur_ `[patchman-ui]`

- **2021-03-09**: Move html webpack to dev dependencies _by @Martin Marosi_ `[patchman-ui]`

- **2021-03-10**: Additional tests for RHSM _by @AlesKas_ `[patchman-engine]`

- **2021-03-10**: SPM-768: require new inventory client _by @Michael Mraka_ `[patchman-engine]`

- **2021-03-10**: bump version 1.7.6 _by @Josef Hak_ `[patchman-engine]`

- **2021-03-11**: SPM-768: inventory attributes have been renamed _by @Michael Mraka_ `[patchman-engine]`

- **2021-03-11**: chore(release): 1.12.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-03-11**: ficx(CVES): search filter is made case insensitive and wording fixed _by @mkholjur_ `[patchman-ui]`

- **2021-03-11**: SPM-768: make gofmt test happy _by @Michael Mraka_ `[patchman-engine]`

- **2021-03-11**: feat(Systems): display OS version lock info _by @mkholjur_ `[patchman-ui]`

- **2021-03-11**: add codecov analysis to github workflow, removed Travis _by @Josef Hak_ `[patchman-engine]`

- **2021-03-12**: chore(release): 1.13.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-03-12**: excluded description field from search in system_packages.go _by @Josef Hak_ `[patchman-engine]`

- **2021-03-12**: updated status link in README, replaced Travis with GitHub actions _by @Josef Hak_ `[patchman-engine]`

- **2021-03-12**: added better pkg.summary and pkg.descritption to test_data.sql _by @Josef Hak_ `[patchman-engine]`

- **2021-03-12**: fixed typo in method name _by @Josef Hak_ `[patchman-engine]`

- **2021-03-12**: updated go.sum with missing item _by @Josef Hak_ `[patchman-engine]`

- **2021-03-16**: bump version 1.7.7 _by @Josef Hak_ `[patchman-engine]`

- **2021-03-16**: updated data generation according to new schema _by @Michael Mraka_ `[patchman-engine]`

- **2021-03-16**: added progress messages into data generation script _by @Michael Mraka_ `[patchman-engine]`

- **2021-03-16**: query system_platform only once _by @Michael Mraka_ `[patchman-engine]`

- **2021-03-16**: make progress pct parameterizable _by @Michael Mraka_ `[patchman-engine]`

- **2021-03-17**: Rename RHSM to be consistent with inventory_schema _by @AlesKas_ `[patchman-engine]`

- **2021-03-17**: chore(CVEs): change button to link _by @mkholjur_ `[patchman-ui]`

- **2021-03-17**: SPM-736: add thirdparty flag to repos _by @Michael Mraka_ `[patchman-engine]`

- **2021-03-17**: fix(OSLock): change rhsm_version object key name to rhsm _by @mkholjur_ `[patchman-ui]`

- **2021-03-17**: SPM-736: mark all repos from vmaas as not thirdparty _by @Michael Mraka_ `[patchman-engine]`

- **2021-03-17**: chore(release): 1.13.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-03-17**: fix(SonarCube): issues are resolved _by @mkholjur_ `[patchman-ui]`

- **2021-03-18**: feat(remediation): add remediation loading state to systems _by @Jiri Dostal_ `[patchman-ui]`

- **2021-03-18**: fix(remediation): create a separate component for remediation button _by @Jiri Dostal_ `[patchman-ui]`

- **2021-03-18**: chore(tests): update tests _by @Jiri Dostal_ `[patchman-ui]`

- **2021-03-18**: SPM-792 some of sonarCube maintability issues are fixed _by @mkholjur_ `[patchman-ui]`

- **2021-03-18**: fix(remediation): add missing file _by @Jiri Dostal_ `[patchman-ui]`

- **2021-03-18**: removed unused "first_reported" column from system_platform table _by @Josef Hak_ `[patchman-engine]`

- **2021-03-18**: updated golangci-lint config _by @Josef Hak_ `[patchman-engine]`

- **2021-03-19**: chore(release): 1.13.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-03-19**: chore(code): Remove not needed console log _by @Jiri Dostal_ `[patchman-ui]`

- **2021-03-22**: SPM-736: mark system with thirdparty repo as thirdparty system _by @Michael Mraka_ `[patchman-engine]`

- **2021-03-22**: fix(OSLock): add undefined case _by @mkholjur_ `[patchman-ui]`

- **2021-03-22**: SPM-736: test for thirdparty systems _by @Michael Mraka_ `[patchman-engine]`

- **2021-03-22**: bump version 1.7.8 _by @Josef Hak_ `[patchman-engine]`

- **2021-03-23**: SPM-782: add additional info to systems endpoint _by @mkholjur_ `[patchman-engine]`

- **2021-03-25**: feat(TableHeaders): enable sticky headers _by @mkholjur_ `[patchman-ui]`

- **2021-03-25**: chore(tests): update tests _by @Jiri Dostal_ `[patchman-ui]`

- **2021-03-25**: chore(release): 1.14.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-03-25**: feat(git commits): SPM-769 use commitizer plugin to enable conventional commit names _by @mkholjur_ `[patchman-ui]`

- **2021-03-26**: SPM-736: updated TestSync with Repo.ThirdParty updating _by @Josef Hak_ `[patchman-engine]`

- **2021-03-26**: chore(release): 1.14.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-03-29**: bump version 1.7.9 _by @Josef Hak_ `[patchman-engine]`

- **2021-03-29**: added test data for repo and system_repo tables _by @Michael Mraka_ `[patchman-engine]`

- **2021-03-29**: removed unused variables _by @Michael Mraka_ `[patchman-engine]`

- **2021-03-29**: update test data - first_reported has been removed _by @Michael Mraka_ `[patchman-engine]`

- **2021-03-30**: Update all FEC components on 21-03-30 _by @Karel Hala_ `[patchman-ui]`

- **2021-03-30**: chore(tests): update test and disable Inventory tests temporarily _by @Jiri Dostal_ `[patchman-ui]`

- **2021-03-31**: Remove html-webpack-plugin from dependencies _by @Martin Marosi_ `[patchman-ui]`

- **2021-03-31**: fix(CveModal): fix url to Vulnerability app _by @mkholjur_ `[patchman-ui]`

- **2021-03-31**: Wrap inventory detail with wrapper _by @Karel Hala_ `[patchman-ui]`

- **2021-03-31**: SPM-812 fix wait-for-services criteria _by @Josef Hak_ `[patchman-engine]`

- **2021-03-31**: Update version of FEC to include inventory fallback _by @Karel Hala_ `[patchman-ui]`

- **2021-03-31**: SPM-824 added ENABLE_TURNPIKE_AUTH=false to vmaas_sync _by @Josef Hak_ `[patchman-engine]`

- **2021-03-31**: SPM-736 repo-sync added to manual sync handler in vmaas-sync _by @Josef Hak_ `[patchman-engine]`

- **2021-03-31**: chore(tests): fix tests _by @mkholjur_ `[patchman-ui]`

- **2021-03-31**: SPM-727 improved testing AssertWait function _by @Josef Hak_ `[patchman-engine]`


## 2021-04 (125 commits)

- **2021-04-01**: bump version 1.8.0 _by @Josef Hak_ `[patchman-engine]`

- **2021-04-01**: fix(Systems): SPM-825 sorting systems table fixed _by @mkholjur_ `[patchman-ui]`

- **2021-04-01**: SPM-736: gorm can't convert int to bool _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-01**: chore(release): 1.14.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-04-01**: SPM-830 split Evaluate method to multiple ones _by @Josef Hak_ `[patchman-engine]`

- **2021-04-06**: Use no content response when account is not found _by @mhornick_ `[patchman-engine]`

- **2021-04-06**: SPM-736: reflect vmaas thirdparty flag _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-06**: SPM-736: need updated vmaas module _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-06**: chore(release): 1.14.3 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-04-06**: chore(release): 1.15.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-04-06**: chore(release): 1.16.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-04-06**: fix(Packages): rename packages column _by @mkholjur_ `[patchman-ui]`

- **2021-04-06**: SPM-736: update tests to new vmaas module _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-07**: SPM-736: update code to new vmaas module _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-07**: SPM-830 Added lazy package save to evaluation _by @Josef Hak_ `[patchman-engine]`

- **2021-04-07**: bump version 1.8.1 _by @Josef Hak_ `[patchman-engine]`

- **2021-04-07**: remove unused modules _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-07**: fix(env): remove html-webpack-plugin and rebuild _by @mkholjur_ `[patchman-ui]`

- **2021-04-07**: chore(release): 1.16.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-04-08**: keep git repo tidy from temporary files _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-08**: chore(release): 1.16.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-04-08**: Added optional DB_DEBUG=false option to debug database transactions _by @Josef Hak_ `[patchman-engine]`

- **2021-04-08**: added other evaluator envs to testing conf evaluator_commons.env _by @Josef Hak_ `[patchman-engine]`

- **2021-04-08**: added "package_name" and "package" into "test_generate_data.sql" _by @Josef Hak_ `[patchman-engine]`

- **2021-04-09**: SPM-830 fix nil pointer in evaluation _by @Josef Hak_ `[patchman-engine]`

- **2021-04-09**: chore(configuration): Add deployment configuration to run against /beta env _by @Katerina Patticha_ `[patchman-ui]`

- **2021-04-09**: SPM-830 Fill lazy package summary and description _by @Josef Hak_ `[patchman-engine]`

- **2021-04-09**: chore(release): 1.17.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-04-09**: SPM-830 optimized packages metadata loading _by @Josef Hak_ `[patchman-engine]`

- **2021-04-09**: SPM-830 added packages summary order by evra _by @Josef Hak_ `[patchman-engine]`

- **2021-04-12**: fix(Clear filters): onDelete function error fixed _by @mkholjur_ `[patchman-ui]`

- **2021-04-12**: Sorting: make nulls sort as smaller than actual values _by @mhornick_ `[patchman-engine]`

- **2021-04-12**: SPM-830 Fixed "installed" value counting _by @Josef Hak_ `[patchman-engine]`

- **2021-04-12**: aligned "ThirdParty" field in vmaas_sync/repo_based.go _by @Josef Hak_ `[patchman-engine]`

- **2021-04-12**: fix(headers) SPM-845 add table headers to empty table _by @mkholjur_ `[patchman-ui]`

- **2021-04-12**: fix(Advisories): SPM-844 wqremove unwanted checkbox _by @mkholjur_ `[patchman-ui]`

- **2021-04-12**: SPM-830 fix save installed system packages _by @Josef Hak_ `[patchman-engine]`

- **2021-04-12**: SPM-830 added tests for package analysis _by @Josef Hak_ `[patchman-engine]`

- **2021-04-12**: SPM-830 removed duplicit PtrString _by @Josef Hak_ `[patchman-engine]`

- **2021-04-12**: make "alias" for vmaa.Ptr*() functions in utils _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-13**: chore(release): 1.17.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-04-13**: bump version 1.9.0 _by @Josef Hak_ `[patchman-engine]`

- **2021-04-13**: chore(release): 1.17.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-04-13**: SPM-846 added packages metrics to vmaas_sync _by @Josef Hak_ `[patchman-engine]`

- **2021-04-13**: bump version 1.9.1 _by @Josef Hak_ `[patchman-engine]`

- **2021-04-14**: SPM-846 updated grafana config _by @Josef Hak_ `[patchman-engine]`

- **2021-04-14**: SPM-846 added note about grafana update script to the README.md _by @Josef Hak_ `[patchman-engine]`

- **2021-04-14**: fix(remediation): SPM-847 fix double mounted remediation modal _by @Jiri Dostal_ `[patchman-ui]`

- **2021-04-14**: chore(release): 1.17.3 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-04-14**: fixed dereference bug _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-14**: fix(bulk select): make bulk select consistence across tables _by @mkholjur_ `[patchman-ui]`

- **2021-04-14**: Update vmaas client to v0.11.0 _by @mhornick_ `[patchman-engine]`

- **2021-04-14**: fix TestSystemsCounts _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-14**: bump version 1.9.2 _by @Josef Hak_ `[patchman-engine]`

- **2021-04-15**: update rbac client to v0.11.0 _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-15**: update code to new rbac client _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-15**: fix(EmmptyTables): SPM-851 disable sorting and export on empty tables _by @mkholjur_ `[patchman-ui]`

- **2021-04-15**: update inventory client to v0.11.0 _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-15**: chore(release): 1.17.4 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-04-15**: update code to new inventory client _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-15**: update tests to new inventory client _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-16**: fix(Export button): reorder the button _by @mkholjur_ `[patchman-ui]`

- **2021-04-16**: fix(Filters): make filter placeholders lower case _by @mkholjur_ `[patchman-ui]`

- **2021-04-16**: SPM-728 Remove randomly failing test "TestRunReaders" _by @Josef Hak_ `[patchman-engine]`

- **2021-04-16**: SPM-853 Update vmaas client to v0.12.0 _by @Josef Hak_ `[patchman-engine]`

- **2021-04-19**: SPM-868 fixed package-name-cache concurrent access _by @Josef Hak_ `[patchman-engine]`

- **2021-04-19**: chore(release): 1.17.5 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-04-19**: fix(Bulk Select): SPM-861 deselect after page selection _by @mkholjur_ `[patchman-ui]`

- **2021-04-19**: bump version 1.9.3 _by @Josef Hak_ `[patchman-engine]`

- **2021-04-19**: chore(release): 1.17.6 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-04-19**: fix(remediation): SPM-863 make remediation button smaller _by @mkholjur_ `[patchman-ui]`

- **2021-04-19**: fix(Advisories): SPM-837 make type column one line _by @mkholjur_ `[patchman-ui]`

- **2021-04-20**: SPM-853 Refactored "vmaas_sync/advisory_sync.go" process _by @Josef Hak_ `[patchman-engine]`

- **2021-04-20**: chore(release): 1.17.7 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-04-20**: SPM-866: make update_data unique _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-20**: chore(release): 1.17.8 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-04-20**: fix(bulk select): always select all when checkbox is not checked _by @mkholjur_ `[patchman-ui]`

- **2021-04-21**: chore(release): 1.17.9 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-04-21**: SPM-853 Refactored "vmaas_sync/pkg_sync.go" _by @Josef Hak_ `[patchman-engine]`

- **2021-04-21**: fix DeleteNewlyAddedAdvisories and DeleteNewlyAddedPackages _by @Josef Hak_ `[patchman-engine]`

- **2021-04-21**: update postgresql-12 image to centos8/rhel8 _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-21**: SPM-853 Update vmaas client to v0.13.0 _by @Josef Hak_ `[patchman-engine]`

- **2021-04-21**: bump version 1.9.4 _by @Josef Hak_ `[patchman-engine]`

- **2021-04-22**: fix(cves): Fix link to Vulnerability to respect beta route _by @Jiri Dostal_ `[patchman-ui]`

- **2021-04-22**: chore(Advisories): SPM-837 better implementation _by @mkholjur_ `[patchman-ui]`

- **2021-04-22**: chore(release): 1.17.10 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-04-22**: fix(select): SPM-876 Select only visible advisories _by @Jiri Dostal_ `[patchman-ui]`

- **2021-04-22**: chore(AdvisoriesIcons): spacing is refactored _by @mkholjur_ `[patchman-ui]`

- **2021-04-22**: bump version 1.9.5 _by @Josef Hak_ `[patchman-engine]`

- **2021-04-22**: use stable version of openapi generator _by @Josef Hak_ `[patchman-engine]`

- **2021-04-22**: chore(release): 1.17.11 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-04-22**: Remove description from advisory search _by @Josef Hak_ `[patchman-engine]`

- **2021-04-22**: chore(package): update remediations to the latest version _by @Jiri Dostal_ `[patchman-ui]`

- **2021-04-22**: use openapi-generator-cli:v5.0.1 in PR checks _by @Josef Hak_ `[patchman-engine]`

- **2021-04-22**: SPM-853 use vmaas /pkgtree endpoint to sync packages _by @Josef Hak_ `[patchman-engine]`

- **2021-04-23**: empty commit to rebuild image _by @Josef Hak_ `[patchman-engine]`

- **2021-04-23**: SPM-878: group listener messages by account _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-23**: SPM-937: Add support for pruning updates (storing only latest) _by @mhornick_ `[patchman-engine]`

- **2021-04-23**: SPM-853 updated vmaas_sync tests _by @Josef Hak_ `[patchman-engine]`

- **2021-04-26**: SPM-853 added "modifiedSince" (ENABLE_MODIFIED_SINCE_SYNC) to vmaas_sync _by @Josef Hak_ `[patchman-engine]`

- **2021-04-26**: SPM-878: move sendMessages() to a shared code _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-26**: SPM-878: sendMessages() already split systems into batches _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-26**: SPM-853 Added progress info logs to sync process _by @Josef Hak_ `[patchman-engine]`

- **2021-04-27**: SPM-878: flush eval events after timeout _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-28**: Log parsing errors in listener _by @mhornick_ `[patchman-engine]`

- **2021-04-28**: SPM-853 Fix package detail request with empty advisory_id _by @Josef Hak_ `[patchman-engine]`

- **2021-04-28**: Fix message processing, add more lenient datetime parser _by @mhornick_ `[patchman-engine]`

- **2021-04-28**: simplified Packages query, removed sub-selects _by @Josef Hak_ `[patchman-engine]`

- **2021-04-28**: SPM-853 added sync duration info logs _by @Josef Hak_ `[patchman-engine]`

- **2021-04-28**: simplified database data testing functions _by @Josef Hak_ `[patchman-engine]`

- **2021-04-28**: SPM-885 replaced deprecated "scopelint" linter with "exportloopref" _by @Josef Hak_ `[patchman-engine]`

- **2021-04-29**: bump version 1.10.0 _by @Josef Hak_ `[patchman-engine]`

- **2021-04-29**: SPM-886 Revert "update inventory client to v0.11.0" _by @Josef Hak_ `[patchman-engine]`

- **2021-04-29**: SPM-886 Validate go.mod with "go mod vendor" _by @Josef Hak_ `[patchman-engine]`

- **2021-04-29**: SPM-885 Revert "Fix message processing, add more lenient datetime parser" _by @Josef Hak_ `[patchman-engine]`

- **2021-04-29**: GORMv2 update _by @AlesKas_ `[patchman-engine]`

- **2021-04-29**: chore(bulkSelect): refactor bulkSelec config into a separate hook _by @mkholjur_ `[patchman-ui]`

- **2021-04-29**: updated to vmaas API client v0.14.0 _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-29**: SPM-568 Set default database log mode "silent" _by @Josef Hak_ `[patchman-engine]`

- **2021-04-29**: Upgrade to GitHub-native Dependabot _by @dependabot-preview[bot]_ `[patchman-ui]`

- **2021-04-30**: ignore coverage.txt _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-30**: updated to rbac API client v0.14.0 _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-30**: updated to inventory API client v0.14.0 _by @Michael Mraka_ `[patchman-engine]`

- **2021-04-30**: SPM-568 Fix inserting empty repo list in listener _by @Josef Hak_ `[patchman-engine]`

- **2021-04-30**: SPM-568 Fix systems filters _by @Josef Hak_ `[patchman-engine]`


## 2021-05 (85 commits)

- **2021-05-03**: feat(Systems): disable selecting on systems with 0 installed, updatable packages and advisories _by @mkholjur_ `[patchman-ui]`

- **2021-05-03**: chore(SystemDetails): refactor third party info _by @mkholjur_ `[patchman-ui]`

- **2021-05-03**: bump version 1.11.0 _by @Josef Hak_ `[patchman-engine]`

- **2021-05-03**: SPM-866: create package cache to speed up system evaluation _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-03**: SPM-866: create package cache to speed up system evaluation _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-03**: SPM-891 Enabled "third_party" in evaluation vmaas request. _by @Josef Hak_ `[patchman-engine]`

- **2021-05-04**: SPM-878: make timeout configurable from env _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-04**: SPM-866: use package cache in evaluation _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-04**: SPM-866: use package cache in evaluation _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-04**: SPM-887 use retry for vmaas-sync vmaas calls _by @Josef Hak_ `[patchman-engine]`

- **2021-05-04**: Disable defaultColumns in inventory table _by @rvsia_ `[patchman-ui]`

- **2021-05-04**: SPM-887 fix HTTPCallRetry function _by @Josef Hak_ `[patchman-engine]`

- **2021-05-05**: fix(localStorage): SPM-879 use latest store in localStorage _by @mkholjur_ `[patchman-ui]`

- **2021-05-05**: fix test data nevras _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-05**: SPM-866: removed unused package_name_cache _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-05**: SPM-866: removed unused package_name_cache _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-05**: chore(release): 1.17.12 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-05-06**: Fix detection of unknown account _by @mhornick_ `[patchman-engine]`

- **2021-05-06**: SPM-887 validated go.mod _by @Josef Hak_ `[patchman-engine]`

- **2021-05-06**: SPM-887 updated sync logs _by @Josef Hak_ `[patchman-engine]`

- **2021-05-06**: chore(release): 1.18.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-05-06**: feat(system detail): use federated modules to share system detail _by @Karel Hala_ `[patchman-ui]`

- **2021-05-06**: SPM-891 Use "third_party" in vmaas call if needed only _by @Josef Hak_ `[patchman-engine]`

- **2021-05-06**: improve github PR template _by @Josef Hak_ `[patchman-engine]`

- **2021-05-06**: fix(UserStatuses): use different approach to check user account status _by @mkholjur_ `[patchman-ui]`

- **2021-05-06**: bump version 1.11.1 _by @Josef Hak_ `[patchman-engine]`

- **2021-05-06**: Update spandx config and add readme about inventory _by @Karel Hala_ `[patchman-ui]`

- **2021-05-06**: chore(release): 1.18.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-05-06**: chore(release): 1.18.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-05-06**: chore(GeneralComponent): some bug fix and refactoring work _by @mkholjur_ `[patchman-ui]`

- **2021-05-07**: SPM-866: add debuging output to PackageCache _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-07**: SPM-866: add debuging output to PackageCache _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-07**: add timezone to local env _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-07**: SPM-866: fixed 0 epoch nevra in test _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-07**: SPM-866: fixed 0 epoch nevra in test _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-07**: Filter packages without description in package detail endpoint _by @mhornick_ `[patchman-engine]`

- **2021-05-07**: Use optimistic updates _by @mhornick_ `[patchman-engine]`

- **2021-05-07**: fix(SystemPackages): Emptystate was incorrect set _by @mkholjur_ `[patchman-ui]`

- **2021-05-07**: chore(release): 1.18.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-05-07**: chore(release): 1.18.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-05-07**: fix(Advisories): blank page is fixed temporarly _by @mkholjur_ `[patchman-ui]`

- **2021-05-11**: fix(Systems): disabled systems are not selected on page select _by @mkholjur_ `[patchman-ui]`

- **2021-05-11**: chore(release): 1.19.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-05-12**: chore(release): 1.19.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-05-13**: Sync 3rd party advisories _by @mhornick_ `[patchman-engine]`

- **2021-05-13**: Relax schema constraints for EPEL advisories _by @mhornick_ `[patchman-engine]`

- **2021-05-14**: don't try to remove system uuid_* functions _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-14**: SPM-866: reuse utils.SinceStr() for time reporting _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-14**: SPM-866: reuse utils.SinceStr() for time reporting _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-14**: bump version 1.12.0 _by @Josef Hak_ `[patchman-engine]`

- **2021-05-14**: SPM-866: print memory size in human readable format _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-14**: SPM-866: print memory size in human readable format _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-14**: SPM-910: build image just once _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-17**: SPM-674 Add "app-common-go" clowder dep  and go mod tidy _by @Josef Hak_ `[patchman-engine]`

- **2021-05-18**: SPM-674 Added Clowder support (db-admin, manager) _by @Josef Hak_ `[patchman-engine]`

- **2021-05-18**: SPM-932: add insights_id to /systems endpoint _by @mkholjur_ `[patchman-engine]`

- **2021-05-19**: SPM-827 Isolate "segmentio/kafka-go" usage to mqueue only _by @Josef Hak_ `[patchman-engine]`

- **2021-05-19**: SPM-827 Use "confluent-kafka-go", update go.mod, go.sum _by @Josef Hak_ `[patchman-engine]`

- **2021-05-19**: SPM-827 Add support for confluentic Kafka go lib - still using tested kafka-go lib by default - enable use confluent-kafka-go using KAFKA_CLIENT_LIB=confluent-kafka-go _by @Josef Hak_ `[patchman-engine]`

- **2021-05-19**: SPM-910: platform-python is always available _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-20**: chore(release): 1.19.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-05-20**: SPM-910: reorder dockerfile to improve layer caching _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-20**: fix(os_version): Os version sorting is disabled _by @mkholjur_ `[patchman-ui]`

- **2021-05-20**: SPM-910: split Dockerfile into build/devel and runtime image _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-20**: chore(release): 1.19.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-05-20**: fix(SelectedRows): fixlter out previously deselected rows _by @mkholjur_ `[patchman-ui]`

- **2021-05-20**: SPM-910: small shell improvements _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-20**: SPM-910: small rpm_list.sh improvement _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-20**: SPM-910: make rhel8 dockerfile consistent with centos _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-20**: SPM-910: reuse image between db and db_volumefix _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-21**: SPM-910: merge test and development dockerfiles _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-24**: SPM-910 Revert dockerfiles updates due to failing osd3 build _by @Josef Hak_ `[patchman-engine]`

- **2021-05-24**: SPM-910: OSD3 doesn't like absolute file paths _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-24**: Revert "SPM-910 Revert dockerfiles updates due to failing osd3 build" _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-24**: chore(release): 1.19.3 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-05-24**: SPM-910: fix workdir and permission in runtime container _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-24**: SPM-910: show current user id _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-24**: SPM-910: move GORUN into extra file _by @Michael Mraka_ `[patchman-engine]`

- **2021-05-25**: Swap RBAC and account check order _by @mhornick_ `[patchman-engine]`

- **2021-05-25**: chore(Error): SPM-933 improve error handling _by @mkholjur_ `[patchman-ui]`

- **2021-05-25**: SPM-922: Periodically refresh package cached counts _by @mhornick_ `[patchman-engine]`

- **2021-05-26**: SPM-674: update utils.RunServer to accept int port (not string) _by @Josef Hak_ `[patchman-engine]`

- **2021-05-27**: chore(deps): bump browserslist from 4.16.3 to 4.16.6 _by @dependabot[bot]_ `[patchman-ui]`

- **2021-05-27**: SPM-674 Support Kafka dependent components _by @Josef Hak_ `[patchman-engine]`

- **2021-05-31**: removed useless "nolint: lll" comments _by @Josef Hak_ `[patchman-engine]`


## 2021-06 (76 commits)

- **2021-06-01**: chore(deps-dev): bump @redhat-cloud-services/frontend-components-config _by @dependabot[bot]_ `[patchman-ui]`

- **2021-06-01**: chore(release): 1.20.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-06-02**: Integrate handle refresh into getEntities. _by @Martin Marosi_ `[patchman-ui]`

- **2021-06-02**: fix(Inventory): fix clear filters issue after custom inventory fetch _by @mkholjur_ `[patchman-ui]`

- **2021-06-02**: SPM-955: collecting smoke test with CJI _by @Nikhil Dhandre_ `[patchman-engine]`

- **2021-06-02**: chore(release): 1.20.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-06-02**: chore(deps): bump @redhat-cloud-services/frontend-components _by @dependabot[bot]_ `[patchman-ui]`

- **2021-06-02**: SPM-674: Added Cloudwatch config to Clowder _by @Josef Hak_ `[patchman-engine]`

- **2021-06-02**: SPM-674: Used METRICS_PORT and METRICS_PATH _by @Josef Hak_ `[patchman-engine]`

- **2021-06-02**: shrink runtime image +100MB by removing dnf cache _by @Michael Mraka_ `[patchman-engine]`

- **2021-06-03**: fixed return code when composes differ _by @Michael Mraka_ `[patchman-engine]`

- **2021-06-03**: new clowder fixes 'open : no such file or directory' _by @Michael Mraka_ `[patchman-engine]`

- **2021-06-03**: SPM-827: Used kafka sasl auth with kafka-go lib (KAFKA_{USERNAME, PASSWORD}) _by @Josef Hak_ `[patchman-engine]`

- **2021-06-03**: SPM-911: update kafka-go lib 0.4.9 -> 0.4.16 _by @Josef Hak_ `[patchman-engine]`

- **2021-06-03**: use getEntities with custom engine _by @mkholjur_ `[patchman-ui]`

- **2021-06-04**: SPM-827: Add kafka sasl env vars to PrintClowderParams _by @Josef Hak_ `[patchman-engine]`

- **2021-06-04**: run local kafka instance directly from git _by @Michael Mraka_ `[patchman-engine]`

- **2021-06-04**: Detail export endpoints _by @mhornick_ `[patchman-engine]`

- **2021-06-07**: SPM-964: fix Kafka certificate issue adding var to test.env _by @Josef Hak_ `[patchman-engine]`

- **2021-06-08**: SPM-964: generate cert with SAN extension _by @Michael Mraka_ `[patchman-engine]`

- **2021-06-08**: bump version 1.12.1 _by @Josef Hak_ `[patchman-engine]`

- **2021-06-08**: fix(OS): SPM-925 column sorting is fixed _by @mkholjur_ `[patchman-ui]`

- **2021-06-08**: Use python3 symlink in ./scripts/generate_docs.sh _by @Josef Hak_ `[patchman-engine]`

- **2021-06-08**: chore(deps): bump @patternfly/react-icons from 4.9.9 to 4.10.7 _by @dependabot[bot]_ `[patchman-ui]`

- **2021-06-08**: chore(deps): bump @redhat-cloud-services/frontend-components-translations _by @dependabot[bot]_ `[patchman-ui]`

- **2021-06-08**: chore(deps): bump @redhat-cloud-services/frontend-components-notifications _by @dependabot[bot]_ `[patchman-ui]`

- **2021-06-08**: chore(deps): bump @patternfly/react-core from 4.106.2 to 4.121.1 _by @dependabot[bot]_ `[patchman-ui]`

- **2021-06-09**: chore(release): 1.20.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-06-09**: chore(deps): bump @redhat-cloud-services/frontend-components-remediations _by @dependabot[bot]_ `[patchman-ui]`

- **2021-06-09**: chore(OS): better sort value calculation _by @mkholjur_ `[patchman-ui]`

- **2021-06-09**: chore(deps): bump ws from 6.2.1 to 6.2.2 _by @dependabot[bot]_ `[patchman-ui]`

- **2021-06-10**: chore(notifications): removed unnecessary notifications and some cleanup _by @mkholjur_ `[patchman-ui]`

- **2021-06-11**: fix(Inventory): SPM:971 sort value needs correct initialization _by @mkholjur_ `[patchman-ui]`

- **2021-06-14**: add openapi.json to runtime image _by @Michael Mraka_ `[patchman-engine]`

- **2021-06-15**: chore(release): 1.20.3 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-06-15**: fix(systems): SPM-957 Fix setting per-page for getEntities _by @Jiri Dostal_ `[patchman-ui]`

- **2021-06-15**: chore(release): 1.20.4 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-06-16**: feat(notification): Refactor exports and add notifications _by @Jiri Dostal_ `[patchman-ui]`

- **2021-06-16**: Update platform-go-middlewares _by @mhornick_ `[patchman-engine]`

- **2021-06-17**: SPM-984: use postgresql sslrootcert if available _by @Michael Mraka_ `[patchman-engine]`

- **2021-06-18**: fixing lint issue _by @Michael Mraka_ `[patchman-engine]`

- **2021-06-18**: bump version 1.12.2 _by @Josef Hak_ `[patchman-engine]`

- **2021-06-18**: SPM-806: Add "reboot_required" column to advisory_metadata table _by @Josef Hak_ `[patchman-engine]`

- **2021-06-18**: SPM-806: Added RebootRequired field to manager endpoints _by @Josef Hak_ `[patchman-engine]`

- **2021-06-18**: SPM-989: Added "system_packages_export" and test _by @Josef Hak_ `[patchman-engine]`

- **2021-06-18**: moved "nolint: gocritic" lint comment from handlers head _by @Josef Hak_ `[patchman-engine]`

- **2021-06-18**: SPM-989: Added "package_systems_export" and test _by @Josef Hak_ `[patchman-engine]`

- **2021-06-18**: simplified sql query in system_packages.go _by @Josef Hak_ `[patchman-engine]`

- **2021-06-21**: SPM-984: honor sslmode during postgresql check _by @Michael Mraka_ `[patchman-engine]`

- **2021-06-21**: fix(Global Filter): SPM-969 global filters fixed on inventory _by @mkholjur_ `[patchman-ui]`

- **2021-06-21**: bump version 1.12.3 _by @Josef Hak_ `[patchman-engine]`

- **2021-06-21**: SPM-831: Added advisory caches refresh job to vmaas-sync _by @Josef Hak_ `[patchman-engine]`

- **2021-06-21**: SPM-831: Added test for advisory caches refresh job _by @Josef Hak_ `[patchman-engine]`

- **2021-06-22**: feat(export): SPM-916 enable export on detail pages _by @mkholjur_ `[patchman-ui]`

- **2021-06-22**: chore(release): 1.21.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-06-22**: SPM-989: Added "display_name" to "package_systems_(export)" endpoints _by @Josef Hak_ `[patchman-engine]`

- **2021-06-23**: fix(Headers): isLoading is set correctly _by @mkholjur_ `[patchman-ui]`

- **2021-06-23**: bump version 1.12.4 _by @Josef Hak_ `[patchman-engine]`

- **2021-06-24**: feat(PackageSystems): use getEntities _by @mkholjur_ `[patchman-ui]`

- **2021-06-24**: chore(release): 1.21.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-06-24**: SPM-798: updated .golangci.yml - replaced 'golint' with 'revive' _by @Josef Hak_ `[patchman-engine]`

- **2021-06-24**: SPM-798: updated code due to new lint usage _by @Josef Hak_ `[patchman-engine]`

- **2021-06-24**: SPM-896: Simplified packages_test.go tests _by @Josef Hak_ `[patchman-engine]`

- **2021-06-24**: SPM-896: Ensured /packages empty response [] (not "null") _by @Josef Hak_ `[patchman-engine]`

- **2021-06-25**: feat(Inventory): SPM-998 persist parameters _by @mkholjur_ `[patchman-ui]`

- **2021-06-25**: fix(Export button): fix misalignment _by @mkholjur_ `[patchman-ui]`

- **2021-06-28**: chore(release): 1.21.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-06-29**: chore(release): 1.22.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-06-29**: feat(PackageSystems): SPM-960 enable checkbox _by @mkholjur_ `[patchman-ui]`

- **2021-06-29**: SPM-915: Added "package_versions" endpoint. _by @Josef Hak_ `[patchman-engine]`

- **2021-06-30**: chore(release): 1.23.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-06-30**: SPM-1005: Added "Latency uptime" metric to grafana board _by @Josef Hak_ `[patchman-engine]`

- **2021-06-30**: SPM-984: Include DB_SSLROOTCERT var to printed clowder params _by @Josef Hak_ `[patchman-engine]`

- **2021-06-30**: SPM-989: Remove paging from export endpoint docs _by @Michal Hornicky_ `[patchman-engine]`

- **2021-06-30**: SPM-989: Remove unnecessary attributes from export endpoints _by @mhornick_ `[patchman-engine]`

- **2021-06-30**: feat(PackageDetail): SPM-961 enable remediation _by @mkholjur_ `[patchman-ui]`


## 2021-07 (67 commits)

- **2021-07-01**: chore(release): 1.24.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-07-01**: feat(Version filter): SPM-914 version filter is finished _by @mkholjur_ `[patchman-ui]`

- **2021-07-01**: feat(remediation): Add provider for system-package remediation _by @Jiri Dostal_ `[patchman-ui]`

- **2021-07-01**: fix(select): Disable checkbox on up-to-date systems _by @Jiri Dostal_ `[patchman-ui]`

- **2021-07-02**: chore(lint): fix lint _by @Jiri Dostal_ `[patchman-ui]`

- **2021-07-07**: fix(PackageDetail): SPM-1015 hot fix for status filter _by @mkholjur_ `[patchman-ui]`

- **2021-07-07**: SPM-674: remove failing ingress and puptoo from Clowder deps _by @Josef Hak_ `[patchman-engine]`

- **2021-07-07**: chore(release): 1.24.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-07-08**: SPM-995: setup postgresql ssl verify-full in local, CI and QA _by @Michael Mraka_ `[patchman-engine]`

- **2021-07-09**: user verify-full everywhere _by @Michael Mraka_ `[patchman-engine]`

- **2021-07-09**: chore(OS): SPM-1008 reword unknown OS cell _by @mkholjur_ `[patchman-ui]`

- **2021-07-09**: SPM-674 respect vmaas resource settings in ephemeral _by @Patrik Segedy_ `[patchman-engine]`

- **2021-07-09**: SPM-984: Added nil pointer check for DB_SSLROOTCERT _by @Josef Hak_ `[patchman-engine]`

- **2021-07-12**: SPM-915: Removed tags and filters support from package_versions _by @Josef Hak_ `[patchman-engine]`

- **2021-07-12**: bump version 1.12.5 _by @Josef Hak_ `[patchman-engine]`

- **2021-07-12**: chore(ExportMessage): SPM-1014 export success message has file type _by @mkholjur_ `[patchman-ui]`

- **2021-07-12**: SPM-866: load cache after probes init _by @Josef Hak_ `[patchman-engine]`

- **2021-07-12**: Revert "SPM-866: ..." _by @Michael Mraka_ `[patchman-engine]`

- **2021-07-13**: chore(Packages): rename search label to Package _by @mkholjur_ `[patchman-ui]`

- **2021-07-14**: fix(Inventory): SPM-1019 unwanted hosts are removed _by @mkholjur_ `[patchman-ui]`

- **2021-07-15**: fix(select): fix unselection of items on package-systems _by @Jiri Dostal_ `[patchman-ui]`

- **2021-07-15**: chore(tests): update snapshots _by @Jiri Dostal_ `[patchman-ui]`

- **2021-07-15**: bump version 1.12.6 _by @Josef Hak_ `[patchman-engine]`

- **2021-07-15**: SPM-1022: Enable to print Clowder params (SHOW_CLOWDER_VARS=true) _by @Josef Hak_ `[patchman-engine]`

- **2021-07-19**: chore(release): 1.24.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-07-19**: chore(release): 1.25.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-07-19**: SPM-991: Fix latest_package_cache view _by @AlesKas_ `[patchman-engine]`

- **2021-07-20**: SPM-1005: Grafana update, added metric "Latency uptime (30d)" _by @Josef Hak_ `[patchman-engine]`

- **2021-07-20**: SPM-1034: print some package stat during debug _by @Michael Mraka_ `[patchman-engine]`

- **2021-07-20**: SPM-1034: use NEVRA with epoch for package cache even if epoch == 0 _by @Michael Mraka_ `[patchman-engine]`

- **2021-07-20**: SPM-1034: Use epoch 0 for package tests _by @Michael Mraka_ `[patchman-engine]`

- **2021-07-20**: SPM-1034: use LOG_DEBUG values if already set _by @Michael Mraka_ `[patchman-engine]`

- **2021-07-21**: feat(NoAccess): SPM-968 use shared not-connected component _by @mkholjur_ `[patchman-ui]`

- **2021-07-21**: chore(PersistantParams): make all pages consistant to apply url params on load _by @mkholjur_ `[patchman-ui]`

- **2021-07-22**: chore(InventoryReducer): SPM-1038 separate reducer for every page _by @mkholjur_ `[patchman-ui]`

- **2021-07-22**: chore(deps): bump urijs from 1.19.6 to 1.19.7 _by @dependabot[bot]_ `[patchman-ui]`

- **2021-07-22**: chore(deps-dev): bump postcss from 8.2.9 to 8.2.10 _by @dependabot[bot]_ `[patchman-ui]`

- **2021-07-22**: chore(deps): bump hosted-git-info from 2.8.8 to 2.8.9 _by @dependabot[bot]_ `[patchman-ui]`

- **2021-07-22**: chore(deps): bump @redhat-cloud-services/frontend-components-remediations _by @dependabot[bot]_ `[patchman-ui]`

- **2021-07-22**: chore(release): 1.26.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-07-22**: fix(PackageSystems): updatable key is added into reducer _by @mkholjur_ `[patchman-ui]`

- **2021-07-22**: chore(deps): bump @patternfly/react-icons from 4.10.7 to 4.11.0 _by @dependabot[bot]_ `[patchman-ui]`

- **2021-07-22**: bump version 1.12.7 _by @Josef Hak_ `[patchman-engine]`

- **2021-07-22**: chore(release): 1.26.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-07-22**: fix(select): Don't ignore query params in bulk select _by @Jiri Dostal_ `[patchman-ui]`

- **2021-07-23**: chore(release): 1.26.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-07-26**: SPM-987: Add check after advisories incremental sync _by @AlesKas_ `[patchman-engine]`

- **2021-07-26**: feat(StatusFilter): SPM-1004 change to checkbox _by @mkholjur_ `[patchman-ui]`

- **2021-07-26**: SPM-1009: Unify filtering on export endpoints _by @mhornick_ `[patchman-engine]`

- **2021-07-26**: chore(package): Fix package-lock architecture _by @Jiri Dostal_ `[patchman-ui]`

- **2021-07-26**: fix(CVE-2021-33623) SPM-970 is remediated _by @mkholjur_ `[patchman-ui]`

- **2021-07-27**: chore(lint): fix lint _by @Jiri Dostal_ `[patchman-ui]`

- **2021-07-27**: chore(release): 1.27.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-07-27**: SPM-1048: Map EPEL repos based on os version _by @mhornick_ `[patchman-engine]`

- **2021-07-27**: fix(Packages): SPM-1040 remove status filter with X button _by @mkholjur_ `[patchman-ui]`

- **2021-07-27**: fix(GlobalFilter): SPM-1041 tags applied on load _by @mkholjur_ `[patchman-ui]`

- **2021-07-28**: chore(release): 1.27.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-07-28**: SPM-1047: Add unspecified advisory type, add remapping fron newpackage to enhancement _by @mhornick_ `[patchman-engine]`

- **2021-07-28**: chore(Export): file format is added to pending state notication _by @mkholjur_ `[patchman-ui]`

- **2021-07-28**: fix(Inventory): problems after rebase are solved _by @mkholjur_ `[patchman-ui]`

- **2021-07-29**: chore(release): 1.27.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-07-29**: Bump github.com/gin-gonic/gin from 1.6.3 to 1.7.0 _by @dependabot[bot]_ `[patchman-engine]`

- **2021-07-29**: SPM-1058: Update gin to 1.7.2, update its deps _by @Josef Hak_ `[patchman-engine]`

- **2021-07-29**: fix(Clear Filters): SPM-1020 redundant API call is removed _by @mkholjur_ `[patchman-ui]`

- **2021-07-30**: feat(webpack): SPM-1062 Switch to webpack proxy and enable auto reload _by @Jiri Dostal_ `[patchman-ui]`

- **2021-07-30**: SPM-1061: Manager returns 500 (unknonw system) _by @AlesKas_ `[patchman-engine]`

- **2021-07-30**: SPM-1055: Use spkg.name_id to utilize an system_package index _by @Josef Hak_ `[patchman-engine]`


## 2021-08 (99 commits)

- **2021-08-02**: chore(deps-dev): bump @redhat-cloud-services/frontend-components-config _by @dependabot[bot]_ `[patchman-ui]`

- **2021-08-02**: chore(deps-dev): bump @redhat-cloud-services/frontend-components-config _by @dependabot[bot]_ `[patchman-ui]`

- **2021-08-02**: chore(release): 1.28.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-08-02**: SPM-1055: Use packageName argument in package_systems endpoint _by @Josef Hak_ `[patchman-engine]`

- **2021-08-02**: chore(release): 1.28.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-08-02**: bump version 1.12.8 _by @Josef Hak_ `[patchman-engine]`

- **2021-08-02**: SPM-1055: Remove useless tags counts metric _by @Josef Hak_ `[patchman-engine]`

- **2021-08-02**: SPM-1055: Improve /packages sql query _by @Josef Hak_ `[patchman-engine]`

- **2021-08-03**: bump version 1.12.9 _by @Josef Hak_ `[patchman-engine]`

- **2021-08-03**: SPM-674: Use 'ingress' and 'puptoo' in Clowder config _by @Josef Hak_ `[patchman-engine]`

- **2021-08-03**: SPM-674: Fixed clowdapp db credentials for evaluators and listener _by @Josef Hak_ `[patchman-engine]`

- **2021-08-04**: bump version 1.13.0 _by @Josef Hak_ `[patchman-engine]`

- **2021-08-04**: SPM-1055: improve /packages query _by @Michael Mraka_ `[patchman-engine]`

- **2021-08-05**: fix(SystemDetail) do not clear parameters on umount _by @mkholjur_ `[patchman-ui]`

- **2021-08-05**: fix(VersionFilter): make sure installed_evra type is handled correctly _by @mkholjur_ `[patchman-ui]`

- **2021-08-05**: fix(revert): System details do not clear parameters on umount _by @mkholjur_ `[patchman-ui]`

- **2021-08-05**: chore(release): 1.28.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-08-05**: SPM-1076: hint for sql debugging _by @Michael Mraka_ `[patchman-engine]`

- **2021-08-05**: chore(release): 1.28.3 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-08-05**: bump version 1.14.0 _by @Josef Hak_ `[patchman-engine]`

- **2021-08-05**: SPM-1076: run only one query for all 8 system metrics _by @Michael Mraka_ `[patchman-engine]`

- **2021-08-05**: fix(PackageSystems): SPM-1072 missing export button is added _by @mkholjur_ `[patchman-ui]`

- **2021-08-05**: fix(Reducers): SPM-1073 clear details page reducers _by @mkholjur_ `[patchman-ui]`

- **2021-08-05**: chore(release): 1.28.4 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-08-06**: SPM-1089: Fix setting of thid_party flag for repos _by @mhornick_ `[patchman-engine]`

- **2021-08-06**: SPM-1076: updated test for getSystemCounts() metrics _by @Michael Mraka_ `[patchman-engine]`

- **2021-08-09**: SPM-1049: Advisories sync mismatch metric _by @AlesKas_ `[patchman-engine]`

- **2021-08-09**: SPM-1083: Fix package_latest_cache containing no summary _by @mhornick_ `[patchman-engine]`

- **2021-08-09**: fix(Notifications): do not fire multiple notifications _by @mkholjur_ `[patchman-ui]`

- **2021-08-09**: SPM-1081: get all cyndi metrics in one query _by @Michael Mraka_ `[patchman-engine]`

- **2021-08-09**: fix(advisoryDetail): SPM-1080 don't show link to errata pages for non-RH content _by @Jiri Dostal_ `[patchman-ui]`

- **2021-08-09**: SPM-1081: get all advisory type counts at once _by @Michael Mraka_ `[patchman-engine]`

- **2021-08-09**: chore(tests): update tests _by @Jiri Dostal_ `[patchman-ui]`

- **2021-08-09**: SPM-1086: Enable Cyndi and iqePlugin in clowder config _by @Josef Hak_ `[patchman-engine]`

- **2021-08-10**: chore(release): 1.28.5 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-08-10**: feat(remediation): SPM-1050 redesign remediation button _by @Jiri Dostal_ `[patchman-ui]`

- **2021-08-10**: bump version 1.14.1 _by @Josef Hak_ `[patchman-engine]`

- **2021-08-10**: chore(EmptyStates): SPM-1033 empty states wording made more informative _by @mkholjur_ `[patchman-ui]`

- **2021-08-10**: SPM-1091: correct and update package evaluation logs _by @Josef Hak_ `[patchman-engine]`

- **2021-08-10**: chore(tests): update tests _by @Jiri Dostal_ `[patchman-ui]`

- **2021-08-10**: SPM-1091: Fix db system-package items loading _by @Josef Hak_ `[patchman-engine]`

- **2021-08-10**: SPM-1091: Fix advisory type index in evaluate.go _by @Josef Hak_ `[patchman-engine]`

- **2021-08-10**: bump version 1.14.2 _by @Josef Hak_ `[patchman-engine]`

- **2021-08-10**: fix(Synopsis) SPM-1078 synopsis is truncated _by @mkholjur_ `[patchman-ui]`

- **2021-08-11**: SPM-1091: Fix empty inventory ID in success evaluation log _by @Josef Hak_ `[patchman-engine]`

- **2021-08-11**: SPM-1091: make advisory counters more error prone _by @Michael Mraka_ `[patchman-engine]`

- **2021-08-11**: SPM-1091: commit new packages immediatelly _by @Michael Mraka_ `[patchman-engine]`

- **2021-08-11**: bump version 1.15.0 _by @Josef Hak_ `[patchman-engine]`

- **2021-08-11**: bump version 1.15.1 _by @Josef Hak_ `[patchman-engine]`

- **2021-08-11**: Map epel repos for vmaas_json _by @mhornick_ `[patchman-engine]`

- **2021-08-11**: bump version 1.15.2 _by @Josef Hak_ `[patchman-engine]`

- **2021-08-11**: feat(NoAccess): SPM-968 use shared not-connected component _by @mkholjur_ `[patchman-ui]`

- **2021-08-12**: chore(release): 1.29.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-08-12**: SPM-1075: Added secrets and params to Clowder config _by @Josef Hak_ `[patchman-engine]`

- **2021-08-12**: chore(webpack): fix insights proxy config _by @Jiri Dostal_ `[patchman-ui]`

- **2021-08-12**: SPM-947: enhance ListCommon to support different subtotals _by @Michael Mraka_ `[patchman-engine]`

- **2021-08-16**: SPM-947: update ListOpts to use generic TotalFunc _by @Michael Mraka_ `[patchman-engine]`

- **2021-08-16**: SPM-947: add new items into API doc _by @Michael Mraka_ `[patchman-engine]`

- **2021-08-16**: chore(release): 1.30.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-08-16**: SPM-947: move sorting after summing _by @Michael Mraka_ `[patchman-engine]`

- **2021-08-17**: fix(advisory): SPM-1080 Don't show link on Advisories for EPEL _by @Jiri Dostal_ `[patchman-ui]`

- **2021-08-17**: fix(spacing): SPM-1113 Fix inconsistent spacing on toolbars/tables _by @Jiri Dostal_ `[patchman-ui]`

- **2021-08-17**: chore(tests): update tests _by @Jiri Dostal_ `[patchman-ui]`

- **2021-08-17**: chore(lint): fix lint _by @Jiri Dostal_ `[patchman-ui]`

- **2021-08-17**: fix(remediation): SPM-1114 Fix remediation of package-systems _by @Jiri Dostal_ `[patchman-ui]`

- **2021-08-17**: chore(release): 1.30.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-08-17**: SPM-947: save and set Select statements back _by @Josef Hak_ `[patchman-engine]`

- **2021-08-17**: SPM-947: test subtotals on system page _by @Michael Mraka_ `[patchman-engine]`

- **2021-08-18**: Bump version to 1.15.3 _by @mhornick_ `[patchman-engine]`

- **2021-08-18**: fix(Advisories): SPM-1100 make advisory description font consistent _by @mkholjur_ `[patchman-ui]`

- **2021-08-18**: SPM-1111: add column for other advisories to systems _by @Michael Mraka_ `[patchman-engine]`

- **2021-08-18**: SPM-1111: reflect /systems and /advisory API changes _by @Michael Mraka_ `[patchman-engine]`

- **2021-08-18**: SPM-1111: test newly added other_counts _by @Michael Mraka_ `[patchman-engine]`

- **2021-08-18**: fix(Filter Chips): SPM-1095 make search filter chip text consistent with search filter title _by @mkholjur_ `[patchman-ui]`

- **2021-08-18**: fix(VersionFilter): SPM-1098 reword to Filter by version _by @mkholjur_ `[patchman-ui]`

- **2021-08-18**: SPM-1111: update advisory tests to reflect added UNSPEC-10 advisory _by @Michael Mraka_ `[patchman-engine]`

- **2021-08-18**: fix(FilterChips): SPM-1097 remove filter chip group _by @mkholjur_ `[patchman-ui]`

- **2021-08-19**: chore(release): 1.30.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-08-19**: chore(release): 1.30.3 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-08-20**: feat(StatusReport): SPM-1043 create System status report _by @mkholjur_ `[patchman-ui]`

- **2021-08-23**: chore(Detail pages): regression is resolved _by @mkholjur_ `[patchman-ui]`

- **2021-08-23**: bump version 1.15.4 _by @Josef Hak_ `[patchman-engine]`

- **2021-08-23**: fix(Systems): SPM-1120 OS column is placed correctly _by @mkholjur_ `[patchman-ui]`

- **2021-08-24**: SPM-1117: Force downloading of all repos _by @mhornick_ `[patchman-engine]`

- **2021-08-24**: chore(lint): fix lint _by @Jiri Dostal_ `[patchman-ui]`

- **2021-08-24**: chore(release): 1.31.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-08-24**: chore(release): 1.31.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-08-25**: SPM-1059: Return 404 for invalid package name _by @AlesKas_ `[patchman-engine]`

- **2021-08-25**: fix(SystemtatusBar): regressions are resolved _by @mkholjur_ `[patchman-ui]`

- **2021-08-25**: chore(release): 1.31.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-08-25**: feat(Other type): SPM-1133 "Other" type is added _by @mkholjur_ `[patchman-ui]`

- **2021-08-25**: SPM-1125: Updated cyndi.appName (patchman -> patch) in clowdapp.yaml _by @Josef Hak_ `[patchman-engine]`

- **2021-08-25**: SPM-1125: Use podman in build_deploy.sh _by @Josef Hak_ `[patchman-engine]`

- **2021-08-26**: chore(release): 1.32.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-08-26**: SPM-1125: Fix certificate exporting from Clowder config _by @Josef Hak_ `[patchman-engine]`

- **2021-08-26**: bump version 1.15.5 _by @Josef Hak_ `[patchman-engine]`

- **2021-08-30**: SPM-1017: Switch to INFO log level _by @AlesKas_ `[patchman-engine]`

- **2021-08-31**: SPM-1087 Add vmass call log _by @Josef Hak_ `[patchman-engine]`

- **2021-08-31**: SPM-1125: Fix log level params in clowdapp.yml _by @Josef Hak_ `[patchman-engine]`


## 2021-09 (55 commits)

- **2021-09-01**: SPM-1125: Update grafana config _by @Josef Hak_ `[patchman-engine]`

- **2021-09-01**: SPM-1125: Update grafana config - single values _by @Josef Hak_ `[patchman-engine]`

- **2021-09-01**: SPM-1125: Fixed grafana config regex patterns _by @Josef Hak_ `[patchman-engine]`

- **2021-09-01**: chore(deps): bump @patternfly/react-icons from 4.11.0 to 4.11.14 _by @dependabot[bot]_ `[patchman-ui]`

- **2021-09-02**: SPM-1146: Update vmaas address reading in Clowder _by @Josef Hak_ `[patchman-engine]`

- **2021-09-02**: bump version 1.15.6 _by @Josef Hak_ `[patchman-engine]`

- **2021-09-02**: SPM-1125: Add pushing to :qa and :latest quay images _by @Josef Hak_ `[patchman-engine]`

- **2021-09-07**: SPM-1149: remove unused manifest file _by @Michael Mraka_ `[patchman-engine]`

- **2021-09-07**: SPM-1149: remove unused scripts for manifest generation _by @Michael Mraka_ `[patchman-engine]`

- **2021-09-07**: SPM-1149: remove manifest upload from clowder _by @Michael Mraka_ `[patchman-engine]`

- **2021-09-07**: SPM-1148: Add VMAAS_ADDRESS to clowdapp.yaml _by @Josef Hak_ `[patchman-engine]`

- **2021-09-08**: SPM-1148: Add VMAAS_WS_ADDRESS to clowdapp.yaml _by @Josef Hak_ `[patchman-engine]`

- **2021-09-08**: SPM-1149: removed unused secret _by @Michael Mraka_ `[patchman-engine]`

- **2021-09-08**: SPM-1132: Return 400 when using incorrect filter _by @AlesKas_ `[patchman-engine]`

- **2021-09-08**: refactor: use http status codes _by @AlesKas_ `[patchman-engine]`

- **2021-09-09**: SPM-1154: Fix nil pointer in fixEpelRepos, listener component _by @Josef Hak_ `[patchman-engine]`

- **2021-09-09**: bump version 1.15.7 _by @Josef Hak_ `[patchman-engine]`

- **2021-09-09**: SPM-1150: return 404 for unknown package name _by @AlesKas_ `[patchman-engine]`

- **2021-09-09**: SPM-1151: return 404 for invalid package name _by @AlesKas_ `[patchman-engine]`

- **2021-09-09**: chore(manifest): remove manifest generator _by @Jiri Dostal_ `[patchman-ui]`

- **2021-09-09**: SPM-806: Updated vmaas client in go.mod and go.sum _by @Josef Hak_ `[patchman-engine]`

- **2021-09-09**: SPM-806: Update vmaas client usage, vmaas methods names changed _by @Josef Hak_ `[patchman-engine]`

- **2021-09-09**: SPM-806: Added "reboot_required" field storing to advisory_sync.go _by @Josef Hak_ `[patchman-engine]`

- **2021-09-09**: SPM-806: Updated advisories sync tests _by @Josef Hak_ `[patchman-engine]`

- **2021-09-13**: bump version 1.16.0 _by @Josef Hak_ `[patchman-engine]`

- **2021-09-13**: SPM-1159: added liveness, readiness probes to database-admin _by @Josef Hak_ `[patchman-engine]`

- **2021-09-14**: chore(deps): SPM-1141 CVE-2021-3749 is remediated _by @mkholjur_ `[patchman-ui]`

- **2021-09-14**: SPM-806: Added "reboot_required" info to advisory detail API _by @Josef Hak_ `[patchman-engine]`

- **2021-09-14**: chore(deps): SPM-1093 SPM-1094 vulnerabilites are fixed _by @mkholjur_ `[patchman-ui]`

- **2021-09-14**: chore(deps): bump @redhat-cloud-services/frontend-components-remediations _by @dependabot[bot]_ `[patchman-ui]`

- **2021-09-14**: chore(deps): bump @patternfly/react-core from 4.147.0 to 4.152.4 _by @dependabot[bot]_ `[patchman-ui]`

- **2021-09-15**: feat(AdvisoriesReport): SPM-1167 add most impactfull advisory info _by @mkholjur_ `[patchman-ui]`

- **2021-09-15**: bump version 1.16.1 _by @Josef Hak_ `[patchman-engine]`

- **2021-09-15**: key rotation _by @Ryan Long_ `[patchman-ui]`

- **2021-09-16**: feat(Reboot info): SPM-1053 reboot required info is added _by @mkholjur_ `[patchman-ui]`

- **2021-09-16**: chore(deps): bump @redhat-cloud-services/frontend-components-translations _by @dependabot[bot]_ `[patchman-ui]`

- **2021-09-16**: chore(deps): bump @redhat-cloud-services/frontend-components-notifications _by @dependabot[bot]_ `[patchman-ui]`

- **2021-09-16**: SPM-1160: Disable hard cyndi db password update, adding UPDATE_CYNDI_PASSWD _by @Josef Hak_ `[patchman-engine]`

- **2021-09-16**: SPM-1165: Added LRU cache for advisory detail endpoint _by @Josef Hak_ `[patchman-engine]`

- **2021-09-16**: SPM-1165: Added advisory detail cache preloading _by @Josef Hak_ `[patchman-engine]`

- **2021-09-17**: SPM-1165: Add advisory detail cache vars to manager deploy config _by @Josef Hak_ `[patchman-engine]`

- **2021-09-17**: SPM-1169: Added 'semantic release' github action _by @Josef Hak_ `[patchman-engine]`

- **2021-09-20**: v1.16.2 _by @semantic-release_ `[patchman-engine]`

- **2021-09-20**: SPM-1168: switch package cache to standard LRU _by @Michael Mraka_ `[patchman-engine]`

- **2021-09-21**: chore(Readme): SPM-1162 add some more info to readme file _by @mkholjur_ `[patchman-ui]`

- **2021-09-21**: chore(AppEntry): remove obsolete code _by @mkholjur_ `[patchman-ui]`

- **2021-09-22**: SPM-1168: update package cache tests _by @Michael Mraka_ `[patchman-engine]`

- **2021-09-23**: SPM-1168: rpm-devel is needed to build gorpm _by @Michael Mraka_ `[patchman-engine]`

- **2021-09-23**: SPM-1168: temporarily use a copy of gorpm with fixed build issue _by @Michael Mraka_ `[patchman-engine]`

- **2021-09-24**: chore(release): 1.33.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-09-29**: chore(release): 1.34.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-09-29**: SPM-1132: Export endpoint bad filter response _by @AlesKas_ `[patchman-engine]`

- **2021-09-30**: fix(headers): SPM-1189 headers are ignored for export endpoints _by @Jiri Dostal_ `[patchman-ui]`

- **2021-09-30**: chore(tests): update tests _by @Jiri Dostal_ `[patchman-ui]`

- **2021-09-30**: chore(release): 1.34.1 _by @semantic-release-bot_ `[patchman-ui]`


## 2021-10 (74 commits)

- **2021-10-01**: Rotate deploy key once more _by @Karel Hala_ `[patchman-ui]`

- **2021-10-01**: New deploy token _by @Karel Hala_ `[patchman-ui]`

- **2021-10-01**: v1.16.3 _by @semantic-release_ `[patchman-engine]`

- **2021-10-01**: SPM-1185: Add "preference" column to "advisory_type" table. _by @Josef Hak_ `[patchman-engine]`

- **2021-10-01**: SPM-1185: Updated advisories endpoints to use advisory.advisory_type_name _by @Josef Hak_ `[patchman-engine]`

- **2021-10-01**: SPM-1185: Updated API docs _by @Josef Hak_ `[patchman-engine]`

- **2021-10-04**: chore(tests): increase coverage _by @Muslimjon_ `[patchman-ui]`

- **2021-10-04**: SPM-1185: Implement the order_query reflection attribute _by @mhornick_ `[patchman-engine]`

- **2021-10-04**: SPM-1168: add package cache variables to clowder _by @Michael Mraka_ `[patchman-engine]`

- **2021-10-04**: v1.16.4 _by @semantic-release_ `[patchman-engine]`

- **2021-10-04**: chore(tests): increase filters test coverage _by @mkholjur_ `[patchman-ui]`

- **2021-10-05**: v1.17.0 _by @semantic-release_ `[patchman-engine]`

- **2021-10-06**: fix(sorting): fix use advisory_type_name for sorting _by @Jiri Dostal_ `[patchman-ui]`

- **2021-10-06**: feat(Tags): SPM-1178 add tags column, tag filter _by @mkholjur_ `[patchman-ui]`

- **2021-10-06**: SPM-1195: consolidate Dockerfiles _by @Michael Mraka_ `[patchman-engine]`

- **2021-10-06**: chore(tests): update tests _by @Jiri Dostal_ `[patchman-ui]`

- **2021-10-06**: fix(env): fix environment setup _by @Jiri Dostal_ `[patchman-ui]`

- **2021-10-06**: chore(deps): bump @patternfly/react-table from 4.29.0 to 4.30.3 _by @dependabot[bot]_ `[patchman-ui]`

- **2021-10-06**: chore(deps): bump @patternfly/react-core from 4.152.4 to 4.157.3 _by @dependabot[bot]_ `[patchman-ui]`

- **2021-10-06**: fix(advisory_type): SPM-1199 fix advisory type in frontend _by @Jiri Dostal_ `[patchman-ui]`

- **2021-10-06**: SPM-1195: avoid failure when creating cyndi_* roles _by @Michael Mraka_ `[patchman-engine]`

- **2021-10-06**: SPM-1195: replace external psql with go code _by @Michael Mraka_ `[patchman-engine]`

- **2021-10-07**: v1.17.1 _by @semantic-release_ `[patchman-engine]`

- **2021-10-07**: feat(OSfilter): SPM-1193 add os filter _by @mkholjur_ `[patchman-ui]`

- **2021-10-07**: chore(tests): fix failing tests _by @mkholjur_ `[patchman-ui]`

- **2021-10-07**: chore(snapshots): update snapshots _by @Jiri Dostal_ `[patchman-ui]`

- **2021-10-07**: SPM-1190: Fixed grafana memory and CPU resources metrics _by @Josef Hak_ `[patchman-engine]`

- **2021-10-07**: SPM-1193: add os filters into api doc _by @mkholjur_ `[patchman-engine]`

- **2021-10-07**: v1.17.2 _by @semantic-release_ `[patchman-engine]`

- **2021-10-07**: fix(CVE-2021-33623) SPM-970 is remediated _by @mkholjur_ `[patchman-ui]`

- **2021-10-07**: SPM-1172: moved devel database files to dev/database _by @Michael Mraka_ `[patchman-engine]`

- **2021-10-07**: v1.17.3 _by @semantic-release_ `[patchman-engine]`

- **2021-10-07**: chore(tests): fix failing tests _by @mkholjur_ `[patchman-ui]`

- **2021-10-07**: chore(release): 1.34.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-10-07**: v1.17.4 _by @semantic-release_ `[patchman-engine]`

- **2021-10-07**: SPM-1201: Add "os" field to system item for proper filtering _by @Josef Hak_ `[patchman-engine]`

- **2021-10-08**: v1.17.5 _by @semantic-release_ `[patchman-engine]`

- **2021-10-08**: SPM-1195: minimize runimg _by @Michael Mraka_ `[patchman-engine]`

- **2021-10-08**: v1.17.6 _by @semantic-release_ `[patchman-engine]`

- **2021-10-08**: v1.17.7 _by @semantic-release_ `[patchman-engine]`

- **2021-10-13**: chore(deps): bump @patternfly/react-icons from 4.11.14 to 4.11.17 _by @dependabot[bot]_ `[patchman-ui]`

- **2021-10-13**: chore(release): 1.34.3 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-10-14**: chore(release): 1.35.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-10-14**: SPM-1205: Fix panic in filter processing _by @Josef Hak_ `[patchman-engine]`

- **2021-10-14**: SPM-1205: Updated filters tests, added tests helper methods _by @Josef Hak_ `[patchman-engine]`

- **2021-10-14**: chore(OsFilter): SPM-925 fix filter bugs _by @mkholjur_ `[patchman-ui]`

- **2021-10-15**: SPM-1191: Add "max attempts" param to kafka clients _by @Josef Hak_ `[patchman-engine]`

- **2021-10-15**: chore(osFilter): fix bugs _by @mkholjur_ `[patchman-ui]`

- **2021-10-17**: chor(SPUR): SPM-1207 truncate long descriptions _by @mkholjur_ `[patchman-ui]`

- **2021-10-17**: chore(SPUR): SPM-1208 rename status to patch status _by @mkholjur_ `[patchman-ui]`

- **2021-10-17**: chore(SPUR): SPM-1206 rename Clear filters to Reser filters _by @mkholjur_ `[patchman-ui]`

- **2021-10-18**: v1.17.8 _by @semantic-release_ `[patchman-engine]`

- **2021-10-18**: v1.17.9 _by @semantic-release_ `[patchman-engine]`

- **2021-10-18**: ignore .venv _by @Michael Mraka_ `[patchman-engine]`

- **2021-10-18**: SPM-1212: Added "advisory_type_name" to advisory detail endpoint _by @Josef Hak_ `[patchman-engine]`

- **2021-10-19**: v1.17.10 _by @semantic-release_ `[patchman-engine]`

- **2021-10-19**: v1.17.11 _by @semantic-release_ `[patchman-engine]`

- **2021-10-19**: fix(Inventory) SPM-1209 fix remediation spacing _by @mkholjur_ `[patchman-ui]`

- **2021-10-19**: SPM-1195: mount certs into rumtime image on localhost _by @Michael Mraka_ `[patchman-engine]`

- **2021-10-19**: SPM-1191: Fix nil pointer in evaluator _by @Josef Hak_ `[patchman-engine]`

- **2021-10-19**: SPM-1195: mount test_data.sql into prod image _by @Michael Mraka_ `[patchman-engine]`

- **2021-10-19**: v1.17.12 _by @semantic-release_ `[patchman-engine]`

- **2021-10-19**: SPM-1195: fix compose check _by @Michael Mraka_ `[patchman-engine]`

- **2021-10-19**: fix(PackageHeader): SPM-1210 fix title bottom space _by @mkholjur_ `[patchman-ui]`

- **2021-10-20**: feat(AdvisoryType): SPM-1197 add advisory type name _by @mkholjur_ `[patchman-ui]`

- **2021-10-20**: chore(release): 1.35.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-10-20**: v1.17.13 _by @semantic-release_ `[patchman-engine]`

- **2021-10-29**: SPM-1213: Ignore unused mapping struct fields in gorm _by @Josef Hak_ `[patchman-engine]`

- **2021-10-29**: SPM-1213: Add "tags" attribute to /systems response _by @Josef Hak_ `[patchman-engine]`

- **2021-10-29**: SPM-1213: Updated /systems test with "tags" info _by @Josef Hak_ `[patchman-engine]`

- **2021-10-29**: SPM-1213: Update API spec _by @Josef Hak_ `[patchman-engine]`

- **2021-10-29**: SPM-1213: Removed redundant SystemInlineItem, used SystemDBLookup struct _by @Josef Hak_ `[patchman-engine]`

- **2021-10-29**: SPM-1213: Added tags parsing to /advisory_systems endpoint and test. _by @Josef Hak_ `[patchman-engine]`

- **2021-10-29**: chore(InventoryComponent): SPM-1229 fix activeFiltersConfig _by @mkholjur_ `[patchman-ui]`


## 2021-11 (53 commits)

- **2021-11-01**: SPM-1213: Added tags parsing to /systems_export endpoint and test. _by @Josef Hak_ `[patchman-engine]`

- **2021-11-01**: SPM-1213: Added tags parsing to /advisory_systems endpoint and test. _by @Josef Hak_ `[patchman-engine]`

- **2021-11-01**: v1.17.14 _by @semantic-release_ `[patchman-engine]`

- **2021-11-01**: chore(release): 1.36.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-11-01**: SPM-1166: use new vmaas client with errata releases _by @Michael Mraka_ `[patchman-engine]`

- **2021-11-02**: SPM-1166: add errata release_versions into db _by @Michael Mraka_ `[patchman-engine]`

- **2021-11-02**: SPM-1166: store errata release_version into db _by @Michael Mraka_ `[patchman-engine]`

- **2021-11-02**: SPM-1166: include release_versions into /advisory API _by @Michael Mraka_ `[patchman-engine]`

- **2021-11-02**: SPM-1166: test for release_version sync _by @Michael Mraka_ `[patchman-engine]`

- **2021-11-02**: SPM-1166: test release_versions in manager _by @Michael Mraka_ `[patchman-engine]`

- **2021-11-03**: feat(Reboot): SPM-1228 add reboot req into system detail _by @mkholjur_ `[patchman-ui]`

- **2021-11-04**: SPM-1234: Add vmaas call retry vars _by @Josef Hak_ `[patchman-engine]`

- **2021-11-04**: SPM-1234: Add vmaas retry config vars to clowdapp.yaml _by @Josef Hak_ `[patchman-engine]`

- **2021-11-04**: v1.17.15 _by @semantic-release_ `[patchman-engine]`

- **2021-11-04**: v1.18.0 _by @semantic-release_ `[patchman-engine]`

- **2021-11-05**: SPM-889: Updated vmaas client lib, to use /pkglist endpoint _by @Josef Hak_ `[patchman-engine]`

- **2021-11-05**: SPM-889: Add package sync using vmaas /pkglist endpoint _by @Josef Hak_ `[patchman-engine]`

- **2021-11-05**: SPM-889: Added vmaas /pkglist mock handler to testing platform service _by @Josef Hak_ `[patchman-engine]`

- **2021-11-05**: SPM-889: Added test for package sync using /pkglist _by @Josef Hak_ `[patchman-engine]`

- **2021-11-05**: SPM-889: Add ENABLE_PKG_TREE_SYNC env var to clowdapp.yaml _by @Josef Hak_ `[patchman-engine]`

- **2021-11-08**: SPM-1166: add release_versions into /advisories and /system_advisories _by @Michael Mraka_ `[patchman-engine]`

- **2021-11-08**: chore(release): 1.37.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-11-08**: SPM-1166: handle missing release_versions as empty string _by @Michael Mraka_ `[patchman-engine]`

- **2021-11-08**: v1.18.1 _by @semantic-release_ `[patchman-engine]`

- **2021-11-09**: SPM-889: Added packages sync check _by @Josef Hak_ `[patchman-engine]`

- **2021-11-09**: SPM-889: Add ENABLE_PACKAGES_COUNT_CHECK vmaas-sync var _by @Josef Hak_ `[patchman-engine]`

- **2021-11-09**: SPM-889: Updated advisories sync log _by @Josef Hak_ `[patchman-engine]`

- **2021-11-10**: SPM-1166: test data for release_versions in /advisories _by @Michael Mraka_ `[patchman-engine]`

- **2021-11-10**: SPM-1166: update test for release_versions in /advisories and /system_advisories _by @Michael Mraka_ `[patchman-engine]`

- **2021-11-11**: v1.18.2 _by @semantic-release_ `[patchman-engine]`

- **2021-11-11**: chore(deps): fix dependency tree _by @mkholjur_ `[patchman-ui]`

- **2021-11-11**: fix(deps): fix deps for running proxy _by @Jiri Dostal_ `[patchman-ui]`

- **2021-11-11**: chore(tests): fix failing tests _by @mkholjur_ `[patchman-ui]`

- **2021-11-12**: chore(release): 1.37.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-11-15**: SPM-1213: Simplified package_systems_test.go _by @Josef Hak_ `[patchman-engine]`

- **2021-11-15**: SPM-1213: Added tags info to package systems endpoint _by @Josef Hak_ `[patchman-engine]`

- **2021-11-15**: SPM-1213: Updated package systems test to check tags info _by @Josef Hak_ `[patchman-engine]`

- **2021-11-15**: SPM-1213: Added tags info to package systems export json _by @Josef Hak_ `[patchman-engine]`

- **2021-11-15**: SPM-1213: Added tags check to package systems export json _by @Josef Hak_ `[patchman-engine]`

- **2021-11-15**: SPM-1213: Updated API spec with package systems tags info _by @Josef Hak_ `[patchman-engine]`

- **2021-11-15**: SPM-1213: Added "tags" field to *systems csv export _by @Josef Hak_ `[patchman-engine]`

- **2021-11-15**: SPM-1213: Added tags check to *systems tests _by @Josef Hak_ `[patchman-engine]`

- **2021-11-15**: v1.18.3 _by @semantic-release_ `[patchman-engine]`

- **2021-11-15**: v1.18.4 _by @semantic-release_ `[patchman-engine]`

- **2021-11-16**: v1.18.5 _by @semantic-release_ `[patchman-engine]`

- **2021-11-16**: SPM-889: Fixed page number logging in vmaas sync _by @Josef Hak_ `[patchman-engine]`

- **2021-11-16**: SPM-889: Upgrade testify library _by @Josef Hak_ `[patchman-engine]`

- **2021-11-16**: v1.18.6 _by @semantic-release_ `[patchman-engine]`

- **2021-11-22**: SPM-1132: Return 400 when filter is not supported _by @AlesKas_ `[patchman-engine]`

- **2021-11-22**: v1.18.7 _by @semantic-release_ `[patchman-engine]`

- **2021-11-24**: feat(PackageSystems): SPM-1178 add tag filter and column _by @mkholjur_ `[patchman-ui]`

- **2021-11-25**: chore(Tags): fix spacing issues, multiple os column, incorrect number of tags in modal _by @mkholjur_ `[patchman-ui]`

- **2021-11-29**: chore(release): 1.38.0 _by @semantic-release-bot_ `[patchman-ui]`


## 2021-12 (46 commits)

- **2021-12-01**: fix(Last Seen): fix invalid date in last seen column _by @mkholjur_ `[patchman-ui]`

- **2021-12-01**: SPM-1259: Add new partitioned db table "baseline" _by @Josef Hak_ `[patchman-engine]`

- **2021-12-01**: SPM-1259: Updated tests checking db tables count _by @Josef Hak_ `[patchman-engine]`

- **2021-12-02**: SPM-1249: replace centos with ubi8 images _by @Michael Mraka_ `[patchman-engine]`

- **2021-12-02**: chore(release): 1.38.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-12-02**: SPM-1249: update dev database to match oficial postgres image _by @Michael Mraka_ `[patchman-engine]`

- **2021-12-03**: SPM-1249: fix db test _by @Michael Mraka_ `[patchman-engine]`

- **2021-12-03**: SPM-1259: Added baseline info to /*systems endpoints _by @Josef Hak_ `[patchman-engine]`

- **2021-12-03**: SPM-1259: Updated /*systems tests _by @Josef Hak_ `[patchman-engine]`

- **2021-12-03**: SPM-1259: Updated /*systems API doc _by @Josef Hak_ `[patchman-engine]`

- **2021-12-03**: SPM-1259: Added baseline info to system detail endpoint _by @Josef Hak_ `[patchman-engine]`

- **2021-12-03**: SPM-1259: Updated system detail test with baseline info check _by @Josef Hak_ `[patchman-engine]`

- **2021-12-06**: v1.18.8 _by @semantic-release_ `[patchman-engine]`

- **2021-12-06**: v1.18.9 _by @semantic-release_ `[patchman-engine]`

- **2021-12-06**: SPM-1279: Fix db start command in docker-compose files _by @Josef Hak_ `[patchman-engine]`

- **2021-12-06**: SPM-1279: Added users access control before and after migration _by @Josef Hak_ `[patchman-engine]`

- **2021-12-07**: v1.18.10 _by @semantic-release_ `[patchman-engine]`

- **2021-12-07**: rewording comment to be clear _by @Michael Mraka_ `[patchman-engine]`

- **2021-12-07**: SPM-1278: Grafana update - replace histogram charts _by @Josef Hak_ `[patchman-engine]`

- **2021-12-07**: v1.18.11 _by @semantic-release_ `[patchman-engine]`

- **2021-12-08**: fix(modules): unique module names introduced, should fix caching issues _by @Jiri Dostal_ `[patchman-ui]`

- **2021-12-08**: chore(release): 1.38.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2021-12-09**: SPM-1260: remove unnecessary argument from function _by @Michael Mraka_ `[patchman-engine]`

- **2021-12-10**: SPM-1259: Temporarily hide baseline changes in API _by @Josef Hak_ `[patchman-engine]`

- **2021-12-10**: SPM-1259: Update tests after hiding baseline API changes _by @Josef Hak_ `[patchman-engine]`

- **2021-12-10**: SPM-1259: Update API docs after hiding baselines info _by @Josef Hak_ `[patchman-engine]`

- **2021-12-10**: SPM-1261: endpoint base code updates for baselines _by @mkholjur_ `[patchman-engine]`

- **2021-12-10**: SPM-1261: endpoint base code update tests for baselines _by @mkholjur_ `[patchman-engine]`

- **2021-12-10**: SPM-1261: endpoint for baselines docs _by @mkholjur_ `[patchman-engine]`

- **2021-12-10**: SPM-1261: endpoint base code updates for baseline systems _by @mkholjur_ `[patchman-engine]`

- **2021-12-10**: v1.18.12 _by @semantic-release_ `[patchman-engine]`

- **2021-12-10**: v1.18.13 _by @semantic-release_ `[patchman-engine]`

- **2021-12-14**: SPM-1292: sonar-project properties _by @Nikhil Dhandre_ `[patchman-engine]`

- **2021-12-14**: SPM-1260: test for baseline _by @Michael Mraka_ `[patchman-engine]`

- **2021-12-14**: SPM-1261: new endpoints for reading baseline systems _by @mkholjur_ `[patchman-engine]`

- **2021-12-14**: SPM-1261: endpoint for baseline systems docs _by @mkholjur_ `[patchman-engine]`

- **2021-12-16**: v1.18.14 _by @semantic-release_ `[patchman-engine]`

- **2021-12-16**: v1.18.15 _by @semantic-release_ `[patchman-engine]`

- **2021-12-16**: SPM-1290: fix tags dropdown alignment _by @mkholjur_ `[patchman-ui]`

- **2021-12-20**: SPM-1261: grant all permission on system_platform to manager _by @mkholjur_ `[patchman-engine]`

- **2021-12-21**: SPM-1261: Added delete_baseline endpoint _by @mkholjur_ `[patchman-engine]`

- **2021-12-21**: SPM-1261: Added delete_baseline tests _by @mkholjur_ `[patchman-engine]`

- **2021-12-21**: SPM-1261: Added delete_baseline to API docs _by @mkholjur_ `[patchman-engine]`

- **2021-12-21**: SPM-1261: base code updates to create a baseline _by @mkholjur_ `[patchman-engine]`

- **2021-12-21**: SPM-1261: base code update tests to create a baseline _by @mkholjur_ `[patchman-engine]`

- **2021-12-22**: fix(CVE): SPM-1247 fix CVE-2021-3918 _by @mkholjur_ `[patchman-ui]`


## 2022-01 (76 commits)

- **2022-01-04**: SPM-1260: filter baseline advisories from vmaas response _by @Michael Mraka_ `[patchman-engine]`

- **2022-01-05**: SPM-1260: tests for baseline evaluation _by @Michael Mraka_ `[patchman-engine]`

- **2022-01-06**: chore(release): 1.38.3 _by @semantic-release-bot_ `[patchman-ui]`

- **2022-01-06**: chore(deps): bump @patternfly/react-icons from 4.19.8 to 4.32.1 _by @dependabot[bot]_ `[patchman-ui]`

- **2022-01-07**: SPM-1260: propagate ENABLE_BASELINE_EVAL env via clowder _by @Michael Mraka_ `[patchman-engine]`

- **2022-01-07**: SPM-1259: Use non-standard pg port 5433, not to conflict with vmaas _by @Josef Hak_ `[patchman-engine]`

- **2022-01-07**: SPM-1259: Add more info to "PostgreSQL is unavailable" message (wait-for-services.sh) _by @Josef Hak_ `[patchman-engine]`

- **2022-01-07**: SPM-1259: Fixed typo "ParseReponseBody > ParseResponseBody" _by @Josef Hak_ `[patchman-engine]`

- **2022-01-10**: v1.18.16 _by @semantic-release_ `[patchman-engine]`

- **2022-01-10**: SPM-1312: fix sonarqube hardcoded IP warning _by @Michael Mraka_ `[patchman-engine]`

- **2022-01-10**: SPM-1312: fix sonarqube WHERE warnings _by @Michael Mraka_ `[patchman-engine]`

- **2022-01-10**: SPM-1261: base code updates to update a baseline _by @mkholjur_ `[patchman-engine]`

- **2022-01-10**: SPM-1261: tests to update a baseline _by @mkholjur_ `[patchman-engine]`

- **2022-01-10**: SPM-1261: api doc changes to update a baseline _by @mkholjur_ `[patchman-engine]`

- **2022-01-10**: SPM-1261: Added missing inventoryIDs check _by @Josef Hak_ `[patchman-engine]`

- **2022-01-10**: SPM-1261: Updated test TestUpdateBaselineInvalidSystem _by @Josef Hak_ `[patchman-engine]`

- **2022-01-10**: SPM-1261: Do not update baseline values on null params _by @Josef Hak_ `[patchman-engine]`

- **2022-01-10**: SPM-1261: Test baseline update with null values _by @Josef Hak_ `[patchman-engine]`

- **2022-01-11**: SPM-1261: Remove INSERT persmission from manager (sql schema) _by @Josef Hak_ `[patchman-engine]`

- **2022-01-11**: v1.18.17 _by @semantic-release_ `[patchman-engine]`

- **2022-01-11**: SPM-1261: Reused database.BaselineConfig for baseline update _by @Josef Hak_ `[patchman-engine]`

- **2022-01-11**: SPM-1261: Use RFC 3339 timestamps in baseline update tests _by @Josef Hak_ `[patchman-engine]`

- **2022-01-11**: SPM-1289: Added "unknown" and "unspecified" testing advisory data. _by @Josef Hak_ `[patchman-engine]`

- **2022-01-11**: SPM-1289: Updated tests due to updated testing sql data _by @Josef Hak_ `[patchman-engine]`

- **2022-01-11**: SPM-1289: Add support for filter[advisory_type_name]=other _by @Josef Hak_ `[patchman-engine]`

- **2022-01-11**: SPM-1289: Updated tests for filter[advisory_type_name]=other _by @Josef Hak_ `[patchman-engine]`

- **2022-01-11**: SPM-1289: Added filter[advisory_type_name]=other to API spec _by @Josef Hak_ `[patchman-engine]`

- **2022-01-11**: SPM-1289: Add "other" advisories count to metrics _by @Josef Hak_ `[patchman-engine]`

- **2022-01-11**: SPM-1289: Update advisories count metrics test _by @Josef Hak_ `[patchman-engine]`

- **2022-01-12**: v1.18.18 _by @semantic-release_ `[patchman-engine]`

- **2022-01-12**: chore(deps): bump @redhat-cloud-services/frontend-components _by @dependabot[bot]_ `[patchman-ui]`

- **2022-01-13**: v1.18.19 _by @semantic-release_ `[patchman-engine]`

- **2022-01-13**: SPM-1261: Updated baseline create endpoint _by @Josef Hak_ `[patchman-engine]`

- **2022-01-13**: SPM-1261: Extended baseline_create_test.go _by @Josef Hak_ `[patchman-engine]`

- **2022-01-13**: SPM-1261: Added baseline create to API docs _by @Josef Hak_ `[patchman-engine]`

- **2022-01-13**: SPM-1282: Added VMAAS_CALL_USE_OPTIMISTIC_UPDATES to evaluator _by @Josef Hak_ `[patchman-engine]`

- **2022-01-13**: SPM-1282: Added VMAAS_CALL_USE_OPTIMISTIC_UPDATES to clowdapp.yaml _by @Josef Hak_ `[patchman-engine]`

- **2022-01-13**: v1.18.20 _by @semantic-release_ `[patchman-engine]`

- **2022-01-14**: SPM-1287: make system status card respect only tags, os, systemPlatform _by @mkholjur_ `[patchman-ui]`

- **2022-01-14**: v1.18.21 _by @semantic-release_ `[patchman-engine]`

- **2022-01-14**: SPM-1261: Add baseline_id validation to baseline_update endpoint _by @Josef Hak_ `[patchman-engine]`

- **2022-01-14**: SPM-1261: Updated rbac methods (w: PUT, DELETE; r: GET, POST) _by @Josef Hak_ `[patchman-engine]`

- **2022-01-17**: v1.18.22 _by @semantic-release_ `[patchman-engine]`

- **2022-01-17**: SPM-1238: make os version consistent with vulnerability _by @Michael Mraka_ `[patchman-engine]`

- **2022-01-18**: v1.18.23 _by @semantic-release_ `[patchman-engine]`

- **2022-01-19**: SPM-1289: Added log to show "other advisory types" _by @Josef Hak_ `[patchman-engine]`

- **2022-01-19**: SPM-1289: Fix, added filter operator transformation for "other" value _by @Josef Hak_ `[patchman-engine]`

- **2022-01-19**: SPM-1289: Updated tests for filter[advisory_type_name]=other _by @Josef Hak_ `[patchman-engine]`

- **2022-01-20**: v1.18.24 _by @semantic-release_ `[patchman-engine]`

- **2022-01-21**: SPM-1330: Added manual inventory mapping _by @Josef Hak_ `[patchman-engine]`

- **2022-01-21**: SPM-1330: Remove generated inventory go client _by @Josef Hak_ `[patchman-engine]`

- **2022-01-21**: SPM-1330: Reduced inventory.SystemProfile, used attributes only _by @Josef Hak_ `[patchman-engine]`

- **2022-01-21**: SPM-1330: Added getters to inventory.SystemProfile _by @Josef Hak_ `[patchman-engine]`

- **2022-01-21**: SPM-1330: Removed useless attributes from listener.Host struct _by @Josef Hak_ `[patchman-engine]`

- **2022-01-21**: SPM-1330: Added EXCLUDED_HOST_TYPES feature to listener _by @Josef Hak_ `[patchman-engine]`

- **2022-01-21**: SPM-1330: Added test for host type excluding _by @Josef Hak_ `[patchman-engine]`

- **2022-01-21**: SPM-1330: Added EXCLUDED_HOST_TYPES var to clowdapp.yaml _by @Josef Hak_ `[patchman-engine]`

- **2022-01-21**: SPM-1331: Mark db systems with host_type=edge as stale (db migr.) _by @Josef Hak_ `[patchman-engine]`

- **2022-01-24**: v1.18.25 _by @semantic-release_ `[patchman-engine]`

- **2022-01-24**: SPM-1334: Added "feature flag" ENABLE_BASELINES_API=true _by @Josef Hak_ `[patchman-engine]`

- **2022-01-24**: SPM-1334: Added ENABLE_BASELINES_API to clowdapp.yaml _by @Josef Hak_ `[patchman-engine]`

- **2022-01-24**: SPM-1334: Added tests for openapi.json filtering _by @Josef Hak_ `[patchman-engine]`

- **2022-01-24**: v1.18.26 _by @semantic-release_ `[patchman-engine]`

- **2022-01-24**: v1.18.27 _by @semantic-release_ `[patchman-engine]`

- **2022-01-25**: v1.18.28 _by @semantic-release_ `[patchman-engine]`

- **2022-01-27**: SPM-1349: Fix evaluator test - added package cache loading _by @Josef Hak_ `[patchman-engine]`

- **2022-01-31**: SPM-1321: make advisory cards responsive _by @mkholjur_ `[patchman-ui]`

- **2022-01-31**: SPM-1349: Removed useless pkg tree option in vmaas_sync _by @Josef Hak_ `[patchman-engine]`

- **2022-01-31**: SPM-1349: Created simplified vmaas client, used generic api.Client _by @Josef Hak_ `[patchman-engine]`

- **2022-01-31**: SPM-1349: Used simplified vmaas client in evaluator _by @Josef Hak_ `[patchman-engine]`

- **2022-01-31**: SPM-1349: Fixed lint - added "nolint: bodyclose" in vmaas_sync _by @Josef Hak_ `[patchman-engine]`

- **2022-01-31**: SPM-1349: Used simplified vmaas client in vmaas_sync/advisory_sync.go _by @Josef Hak_ `[patchman-engine]`

- **2022-01-31**: SPM-1349: Used simplified vmaas client in vmaas_sync/package_sync.go _by @Josef Hak_ `[patchman-engine]`

- **2022-01-31**: SPM-1349: Used simplified vmaas client in vmaas_sync/repo_based.go _by @Josef Hak_ `[patchman-engine]`

- **2022-01-31**: SPM-1349: Used simplified vmaas client in vmaas_sync/vmaas_sync.go _by @Josef Hak_ `[patchman-engine]`

- **2022-01-31**: SPM-1349: Used CentOS 8 Stream in Dockerfile (Centos 8 EOL) _by @Josef Hak_ `[patchman-engine]`


## 2022-02 (78 commits)

- **2022-02-01**: SPM-1261: Fixed baseline_update apidoc (POST -> PUT) _by @Josef Hak_ `[patchman-engine]`

- **2022-02-01**: v1.18.29 _by @semantic-release_ `[patchman-engine]`

- **2022-02-01**: v1.18.30 _by @semantic-release_ `[patchman-engine]`

- **2022-02-01**: chore(deps): bump @patternfly/react-icons from 4.32.1 to 4.43.15 _by @dependabot[bot]_ `[patchman-ui]`

- **2022-02-01**: chore(deps): bump @redhat-cloud-services/frontend-components-translations _by @dependabot[bot]_ `[patchman-ui]`

- **2022-02-01**: SPM-1349: Added "method" param to api client (POST, GET, PUT...) _by @Josef Hak_ `[patchman-engine]`

- **2022-02-01**: SPM-1349: Simplified rbac client, removed generated client usage _by @Josef Hak_ `[patchman-engine]`

- **2022-02-02**: v1.18.31 _by @semantic-release_ `[patchman-engine]`

- **2022-02-02**: SPM-1349: Fixed TestRun, added wg.Wait() _by @Josef Hak_ `[patchman-engine]`

- **2022-02-02**: v1.18.32 _by @semantic-release_ `[patchman-engine]`

- **2022-02-02**: SPM-1268: Added Floorist (DataHub) config to deploy/clowdapp.yaml _by @Josef Hak_ `[patchman-engine]`

- **2022-02-03**: v1.18.33 _by @semantic-release_ `[patchman-engine]`

- **2022-02-03**: SPM-1349: Used simplified vmaas client in listener (structs only) _by @Josef Hak_ `[patchman-engine]`

- **2022-02-03**: SPM-1349: Redefined Ptr functions not to use gen. vmaas client _by @Josef Hak_ `[patchman-engine]`

- **2022-02-03**: SPM-1349: Removed vmaas client usage from platform/vmaas.go _by @Josef Hak_ `[patchman-engine]`

- **2022-02-03**: SPM-1349: Removed not used vmaas pkgtree mock (platform/vmaas.go) _by @Josef Hak_ `[patchman-engine]`

- **2022-02-03**: SPM-1349: Removed patchman-clients/rbac usage from platform/rbac.go _by @Josef Hak_ `[patchman-engine]`

- **2022-02-03**: SPM-1361: Added "synced" boolean table columns _by @Josef Hak_ `[patchman-engine]`

- **2022-02-03**: SPM-1361: Updated advisory_sync and tests _by @Josef Hak_ `[patchman-engine]`

- **2022-02-03**: SPM-1361: Updated package_sync, package eval and tests _by @Josef Hak_ `[patchman-engine]`

- **2022-02-03**: SPM-1361: Fixed permission issue in listener test _by @Josef Hak_ `[patchman-engine]`

- **2022-02-04**: v1.18.34 _by @semantic-release_ `[patchman-engine]`

- **2022-02-04**: SPM-1268: Fixed floorist sql query in deploy/clowdapp.yaml _by @Josef Hak_ `[patchman-engine]`

- **2022-02-04**: v1.18.35 _by @semantic-release_ `[patchman-engine]`

- **2022-02-04**: SPM-1261: Updated baselines sql test data _by @Josef Hak_ `[patchman-engine]`

- **2022-02-07**: SPM-1261: Fixed baselines listing query, added tags support _by @Josef Hak_ `[patchman-engine]`

- **2022-02-07**: SPM-1261: Added tags support to baseline systems endpoint _by @Josef Hak_ `[patchman-engine]`

- **2022-02-07**: SPM-1261: Added "tags" filter to baselines API doc _by @Josef Hak_ `[patchman-engine]`

- **2022-02-07**: SPM-1261: Added baseline detail endpoint _by @Josef Hak_ `[patchman-engine]`

- **2022-02-07**: SPM-1261: Added baseline detail endpoint test _by @Josef Hak_ `[patchman-engine]`

- **2022-02-07**: SPM-1261: Added baseline detail endpoint to API doc _by @Josef Hak_ `[patchman-engine]`

- **2022-02-07**: SPM-1261: Clarified SQL condition in baseline_delete _by @Josef Hak_ `[patchman-engine]`

- **2022-02-07**: SPM-1261: Added "account" condition into the baseline_update _by @Josef Hak_ `[patchman-engine]`

- **2022-02-08**: v1.18.36 _by @semantic-release_ `[patchman-engine]`

- **2022-02-08**: v1.18.37 _by @semantic-release_ `[patchman-engine]`

- **2022-02-08**: SPM-1261: Updated structs comments (swag annotations) _by @Josef Hak_ `[patchman-engine]`

- **2022-02-08**: SPM-1261: Updated (re-generated) API spec _by @Josef Hak_ `[patchman-engine]`

- **2022-02-09**: v1.18.38 _by @semantic-release_ `[patchman-engine]`

- **2022-02-09**: SPM-1261: Update "system_platform.unchanged_since" on baseline update _by @Josef Hak_ `[patchman-engine]`

- **2022-02-10**: v1.18.39 _by @semantic-release_ `[patchman-engine]`

- **2022-02-10**: SPM-1261: Added Kafka writer to manager _by @Josef Hak_ `[patchman-engine]`

- **2022-02-10**: SPM-1261: Used Kafka writer in manager baseline endpoints _by @Josef Hak_ `[patchman-engine]`

- **2022-02-10**: SPM-1261: Added manager Kafka writer tests _by @Josef Hak_ `[patchman-engine]`

- **2022-02-10**: SPM-1261: Added Kafka params to manager section in clowdapp.yaml _by @Josef Hak_ `[patchman-engine]`

- **2022-02-10**: SPM-1261: Fixed testing setup for docker-compose.yml _by @Josef Hak_ `[patchman-engine]`

- **2022-02-11**: chore(deps-dev): bump node-sass from 6.0.1 to 7.0.0 _by @dependabot[bot]_ `[patchman-ui]`

- **2022-02-11**: SPM-1370: Rename sonarqube project-key _by @Nikhil Dhandre_ `[patchman-engine]`

- **2022-02-11**: v1.18.40 _by @semantic-release_ `[patchman-engine]`

- **2022-02-11**: SPM-1372: fix links in the documentation _by @Michael Mraka_ `[patchman-engine]`

- **2022-02-11**: SPM-1261: Use inventory IDs sending in kafka.go instead of baseline ID _by @Josef Hak_ `[patchman-engine]`

- **2022-02-11**: v1.18.41 _by @semantic-release_ `[patchman-engine]`

- **2022-02-11**: feat(SPM-1308): assign system to a patch set _by @mkholjur_ `[patchman-ui]`

- **2022-02-11**: feat(SPM-1306): create patch set wizard _by @mkholjur_ `[patchman-ui]`

- **2022-02-14**: sonarqube config _by @Nikhil Dhandre_ `[patchman-ui]`

- **2022-02-15**: SPM-1261: Change env var: ENABLE_EVALUATION_REQUESTS -> ENABLE_BASELINE_CHANGE_EVAL _by @Josef Hak_ `[patchman-engine]`

- **2022-02-15**: SPM-1261: Correct var name in manager/kafka/kafka.go _by @Josef Hak_ `[patchman-engine]`

- **2022-02-15**: chore(deps): bump follow-redirects from 1.14.7 to 1.14.8 _by @dependabot[bot]_ `[patchman-ui]`

- **2022-02-15**: SPM-1361: delete unused package and advisory data _by @Michael Mraka_ `[patchman-engine]`

- **2022-02-15**: v1.18.42 _by @semantic-release_ `[patchman-engine]`

- **2022-02-17**: fix(SystemsStatusReport): SPM-1358 fixed responsibility _by @Alexandr Čelakovský_ `[patchman-ui]`

- **2022-02-17**: fix(AdvisoriesStatusReport): SPM-1358 fixed responsiblity _by @Alexandr Čelakovský_ `[patchman-ui]`

- **2022-02-17**: SPM-1349: Removed unused patchman-clients/{rbac,vmaas} deps _by @Josef Hak_ `[patchman-engine]`

- **2022-02-17**: v1.18.43 _by @semantic-release_ `[patchman-engine]`

- **2022-02-18**: SPM-1315: Add tag and SAP filters into metadata _by @AlesKas_ `[patchman-engine]`

- **2022-02-21**: feat(SPM-1305): create patch set table _by @mkholjur_ `[patchman-ui]`

- **2022-02-21**: chore(deps-dev): bump @redhat-cloud-services/frontend-components-config _by @dependabot[bot]_ `[patchman-ui]`

- **2022-02-21**: v1.18.44 _by @semantic-release_ `[patchman-engine]`

- **2022-02-22**: chore(deps): install sass-loader _by @mkholjur_ `[patchman-ui]`

- **2022-02-22**: chore(release): 1.39.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2022-02-22**: chore(PatchSetWizard): rename patchSet var to rows _by @mkholjur_ `[patchman-ui]`

- **2022-02-22**: SPM-1382: fix data path in postgresql container _by @Michael Mraka_ `[patchman-engine]`

- **2022-02-22**: feat(RHIF-22): enable AAP & MSSQL fitlers _by @mkholjur_ `[patchman-ui]`

- **2022-02-23**: fix(InventoryDetail): Changed Alert to inline _by @Alexandr Čelakovský_ `[patchman-ui]`

- **2022-02-24**: v1.18.45 _by @semantic-release_ `[patchman-engine]`

- **2022-02-24**: SPM-1361: add system_package data generation _by @Michael Mraka_ `[patchman-engine]`

- **2022-02-24**: SPM-1261: Fix baseline API responses (create, delete, update) _by @Josef Hak_ `[patchman-engine]`

- **2022-02-24**: SPM-1261: Updated API docs - baseline endpoints responses _by @Josef Hak_ `[patchman-engine]`

- **2022-02-25**: v1.18.46 _by @semantic-release_ `[patchman-engine]`


## 2022-03 (68 commits)

- **2022-03-01**: feat(SPM-1380): integrate patch set table with patch set actions _by @mkholjur_ `[patchman-ui]`

- **2022-03-01**: SPM-1349: Added deps backup repo link to README.md _by @Josef Hak_ `[patchman-engine]`

- **2022-03-01**: fix(Travis): update node version _by @mkholjur_ `[patchman-ui]`

- **2022-03-01**: chore(release): 1.40.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2022-03-01**: chore(Travis): fix errors _by @mkholjur_ `[patchman-ui]`

- **2022-03-01**: v1.18.47 _by @semantic-release_ `[patchman-engine]`

- **2022-03-02**: fix(PatchSetWizard): separate request to update and create set _by @mkholjur_ `[patchman-ui]`

- **2022-03-03**: SPM-1371: Add subtotals metadata to advisories _by @michalslomczynski_ `[patchman-engine]`

- **2022-03-03**: v1.18.48 _by @semantic-release_ `[patchman-engine]`

- **2022-03-03**: Revert "chore(Travis): fix errors" _by @mkholjur_ `[patchman-ui]`

- **2022-03-03**: fix(tests): failing tests are resolved with some cleanup _by @mkholjur_ `[patchman-ui]`

- **2022-03-04**: SPM-1361: add index for package_id foreign key _by @Michael Mraka_ `[patchman-engine]`

- **2022-03-07**: SPM-1391: Add AAP and MSSQL filters _by @michalslomczynski_ `[patchman-engine]`

- **2022-03-07**: SPM-1400: Fix GetBaselineConfig _by @Josef Hak_ `[patchman-engine]`

- **2022-03-07**: SPM-1400: Updated GetBaselineConfig test _by @Josef Hak_ `[patchman-engine]`

- **2022-03-07**: v1.18.49 _by @semantic-release_ `[patchman-engine]`

- **2022-03-07**: v1.18.50 _by @semantic-release_ `[patchman-engine]`

- **2022-03-08**: fix(SPM-1381): Pagination dropdown disabled when table empty (#740) [#740] _by @Alexandr Čelakovský_ `[patchman-ui]`

- **2022-03-08**: chore(release): 1.40.1 [#740] _by @semantic-release-bot_ `[patchman-ui]`

- **2022-03-09**: chore(tests): PR review improvements _by @mkholjur_ `[patchman-ui]`

- **2022-03-09**: SPM-1398: Simplify swagger doc generator _by @Josef Hak_ `[patchman-engine]`

- **2022-03-10**: chore(release): 1.40.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2022-03-10**: SPM-1391: Fix global filters - handle not_nil values _by @michalslomczynski_ `[patchman-engine]`

- **2022-03-10**: fix(Remediation): apply new loading approach _by @mkholjur_ `[patchman-ui]`

- **2022-03-10**: v1.18.51 _by @semantic-release_ `[patchman-engine]`

- **2022-03-10**: v1.18.52 _by @semantic-release_ `[patchman-engine]`

- **2022-03-11**: feat(PatchSet): put patch set work under feature flag _by @mkholjur_ `[patchman-ui]`

- **2022-03-11**: SPM-1404: Fix multiple SAP IDs _by @michalslomczynski_ `[patchman-engine]`

- **2022-03-14**: SPM-1403: Return baseline details on baseline-update _by @michalslomczynski_ `[patchman-engine]`

- **2022-03-14**: SPM-1320: Update non-200 response API types _by @michalslomczynski_ `[patchman-engine]`

- **2022-03-14**: v1.18.53 _by @semantic-release_ `[patchman-engine]`

- **2022-03-14**: v1.18.54 _by @semantic-release_ `[patchman-engine]`

- **2022-03-14**: chore(release): 1.40.3 _by @semantic-release-bot_ `[patchman-ui]`

- **2022-03-15**: v1.18.55 _by @semantic-release_ `[patchman-engine]`

- **2022-03-15**: chore(tests): update snapshots _by @mkholjur_ `[patchman-ui]`

- **2022-03-15**: chore(Travis): add custom release scripts _by @mkholjur_ `[patchman-ui]`

- **2022-03-15**: SPM-1398: Extend nevra parse error log _by @Josef Hak_ `[patchman-engine]`

- **2022-03-16**: SPM-1361: tests for clean_unused_data _by @Michael Mraka_ `[patchman-engine]`

- **2022-03-16**: chore(PatchSet): enable different behaviour on Systems and patch-set tables _by @mkholjur_ `[patchman-ui]`

- **2022-03-16**: SPM-1361: mark initial test packages as synced _by @Michael Mraka_ `[patchman-engine]`

- **2022-03-17**: SPM-1412: update gorm.io/driver/postgres to fix broken deps _by @Michael Mraka_ `[patchman-engine]`

- **2022-03-17**: v1.18.56 _by @semantic-release_ `[patchman-engine]`

- **2022-03-17**: SPM-1414: Revert "Temporarily hide baseline changes in API" _by @Josef Hak_ `[patchman-engine]`

- **2022-03-17**: SPM-1414: Revert "Update tests after hiding baseline API changes" _by @Josef Hak_ `[patchman-engine]`

- **2022-03-17**: SPM-1414: Revert "Update API docs after hiding baselines info" _by @Josef Hak_ `[patchman-engine]`

- **2022-03-17**: fix(PatchSet): to_date field is moved into config object _by @mkholjur_ `[patchman-ui]`

- **2022-03-17**: chore(PatchWizard): now patch-set table reloads after successfull request _by @mkholjur_ `[patchman-ui]`

- **2022-03-18**: v1.18.57 _by @semantic-release_ `[patchman-engine]`

- **2022-03-18**: v1.18.58 _by @semantic-release_ `[patchman-engine]`

- **2022-03-18**: chore(PatchSet): pre-select already assigned systems in review systems step _by @mkholjur_ `[patchman-ui]`

- **2022-03-18**: Update README with latest build config (#753) [#753] _by @Samuel Olekšák_ `[patchman-ui]`

- **2022-03-18**: chore(PatchSet): configuration step to behaves different for systems page, patch-sets page _by @mkholjur_ `[patchman-ui]`

- **2022-03-18**: chore(release): 1.40.4 _by @semantic-release-bot_ `[patchman-ui]`

- **2022-03-21**: fix(PatchSetWizard): fetchPatchSetsAction is called only assigning a system _by @mkholjur_ `[patchman-ui]`

- **2022-03-21**: chore(PatchSet): enable loading effect on patch-set selector _by @mkholjur_ `[patchman-ui]`

- **2022-03-22**: SPM-1361: enable new env variables in clowder _by @Michael Mraka_ `[patchman-engine]`

- **2022-03-23**: v1.18.59 _by @semantic-release_ `[patchman-engine]`

- **2022-03-23**: SPM-1325: Update listener to read updateinfo from kafka _by @michalslomczynski_ `[patchman-engine]`

- **2022-03-23**: chore(GlobalFilters): update string concatenation _by @mkholjur_ `[patchman-ui]`

- **2022-03-24**: v1.18.60 _by @semantic-release_ `[patchman-engine]`

- **2022-03-25**: SPM-1326: Handle updateinfo in evaluator _by @michalslomczynski_ `[patchman-engine]`

- **2022-03-25**: SPM-1326: function to compare NEVRAs _by @Michael Mraka_ `[patchman-engine]`

- **2022-03-25**: chore(PatchSet): loading state enabled for configuration fields _by @mkholjur_ `[patchman-ui]`

- **2022-03-25**: SPM-1326: use NEVRA compare function _by @Michael Mraka_ `[patchman-engine]`

- **2022-03-28**: SPM-1326: Add yum evaluator test _by @michalslomczynski_ `[patchman-engine]`

- **2022-03-29**: chore(PatchSetWizards): write test cases to cover InputFields _by @mkholjur_ `[patchman-ui]`

- **2022-03-30**: chore(deps): bump minimist from 1.2.5 to 1.2.6 _by @dependabot[bot]_ `[patchman-ui]`

- **2022-03-31**: chore(PatchSetWizard): fix bulk systems selection and some clean-up _by @mkholjur_ `[patchman-ui]`


## 2022-04 (75 commits)

- **2022-04-01**: SPM-1416: Add baseline description field _by @michalslomczynski_ `[patchman-engine]`

- **2022-04-01**: chore(release): 1.41.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2022-04-01**: SPM-1416: make error handling reusable for different pg errors _by @Michael Mraka_ `[patchman-engine]`

- **2022-04-01**: SPM-1392: Fix nil pointer dereference on tags _by @michalslomczynski_ `[patchman-engine]`

- **2022-04-01**: SPM-1416: fixed pointer logic in optional/mandatory attributes _by @Michael Mraka_ `[patchman-engine]`

- **2022-04-01**: SPM-1411: Fix infinite loop in nestedQuery _by @michalslomczynski_ `[patchman-engine]`

- **2022-04-04**: v1.18.61 _by @semantic-release_ `[patchman-engine]`

- **2022-04-04**: SPM-1392: key should not start with "/" if namespace is empty _by @Michael Mraka_ `[patchman-engine]`

- **2022-04-04**: v1.18.62 _by @semantic-release_ `[patchman-engine]`

- **2022-04-04**: v1.18.63 _by @semantic-release_ `[patchman-engine]`

- **2022-04-06**: fix(Systems): SPM-1394 fetch filtered data for selection when search param exists _by @mkholjur_ `[patchman-ui]`

- **2022-04-06**: SPM-1416: make sure there are no duplicate before creating unique constraint _by @Michael Mraka_ `[patchman-engine]`

- **2022-04-06**: v1.18.64 _by @semantic-release_ `[patchman-engine]`

- **2022-04-06**: chore(tests): Systems stable bulk select test cases with some cleanup _by @mkholjur_ `[patchman-ui]`

- **2022-04-06**: v1.18.65 _by @semantic-release_ `[patchman-engine]`

- **2022-04-07**: chore(SIDs): change SIDs filter from comma separated values to single valued filters _by @mkholjur_ `[patchman-ui]`

- **2022-04-07**: SPM-1326: simplify mergeUpdates() _by @Michael Mraka_ `[patchman-engine]`

- **2022-04-07**: SPM-1326: improved updates merge tests _by @Michael Mraka_ `[patchman-engine]`

- **2022-04-08**: SPM-1416: Fix baseline update condition _by @michalslomczynski_ `[patchman-engine]`

- **2022-04-08**: v1.18.66 _by @semantic-release_ `[patchman-engine]`

- **2022-04-08**: chore(PatchSet): preselect and unassign systems, correct time format _by @mkholjur_ `[patchman-ui]`

- **2022-04-08**: chore(PatchSet): patch set configuration result step update _by @mkholjur_ `[patchman-ui]`

- **2022-04-08**: chore(deps-dev): bump @redhat-cloud-services/frontend-components-config _by @dependabot[bot]_ `[patchman-ui]`

- **2022-04-08**: feat(Systems): patch-set column added _by @mkholjur_ `[patchman-ui]`

- **2022-04-08**: chore(release): 1.42.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2022-04-08**: chore(release): 1.43.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2022-04-11**: SPM-1326: fix yum update evaluation logic _by @Michael Mraka_ `[patchman-engine]`

- **2022-04-11**: feat(PatchSet): assign multiple systems, existing patch-set selection filterable, paginated _by @mkholjur_ `[patchman-ui]`

- **2022-04-11**: chore(release): 1.44.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2022-04-12**: fix(ReviewSystems): persist system selection across page and step change _by @mkholjur_ `[patchman-ui]`

- **2022-04-12**: fix(PatchSetTable): make columns sortable, remove kebab menu on empty table _by @mkholjur_ `[patchman-ui]`

- **2022-04-12**: chore(release): 1.44.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2022-04-12**: fix(ConfigurationFields): validate toDate filed with regex _by @mkholjur_ `[patchman-ui]`

- **2022-04-12**: SPM-1416: Remove omitempty from response _by @michalslomczynski_ `[patchman-engine]`

- **2022-04-12**: chore(release): 1.45.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2022-04-12**: v1.18.67 _by @semantic-release_ `[patchman-engine]`

- **2022-04-14**: chore(release): 1.45.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2022-04-18**: Initial commit _by @Justin Sherrill_ `[content-sources-backend]`

- **2022-04-18**: Initial Structure _by @Justin Sherrill_ `[content-sources-backend]`

- **2022-04-18**: initial app with deployment, dockerfile, and ping api _by @Justin Sherrill_ `[content-sources-backend]`

- **2022-04-18**: correcting quay location _by @Justin Sherrill_ `[content-sources-backend]`

- **2022-04-18**: Fixing api path _by @Justin Sherrill_ `[content-sources-backend]`

- **2022-04-19**: feat(SPM-1311): assign systems to a patch-set in System details page _by @mkholjur_ `[patchman-ui]`

- **2022-04-20**: chore(release): 1.46.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2022-04-20**: SPM-1452: added patch-set/baseline filter for systems _by @Michael Mraka_ `[patchman-engine]`

- **2022-04-20**: SPM-1452: just reformat comment _by @Michael Mraka_ `[patchman-engine]`

- **2022-04-20**: chore: Bump frontend-components-config (#777) [#777] _by @Samuel Olekšák_ `[patchman-ui]`

- **2022-04-20**: chore(deps): bump @redhat-cloud-services/frontend-components _by @dependabot[bot]_ `[patchman-ui]`

- **2022-04-21**: chore(Tests): fix failing tests, update snapshots _by @mkholjur_ `[patchman-ui]`

- **2022-04-21**: SPM-1455: make local image names cosistent with quay.io _by @Michael Mraka_ `[patchman-engine]`

- **2022-04-21**: SPM-1455: minimize prod runtime to ubi8-micro _by @Michael Mraka_ `[patchman-engine]`

- **2022-04-21**: SPM-1455: clean unused haberdasher _by @Michael Mraka_ `[patchman-engine]`

- **2022-04-21**: feat(ReviewSystems): add patch-set column _by @mkholjur_ `[patchman-ui]`

- **2022-04-21**: fix(SystemDetail): show tags and baseline info _by @mkholjur_ `[patchman-ui]`

- **2022-04-21**: chore(ReviewStep): status, patch status filters added. Defaults filters enabled _by @mkholjur_ `[patchman-ui]`

- **2022-04-22**: v1.18.68 _by @semantic-release_ `[patchman-engine]`

- **2022-04-22**: SPM-1433: Add endpoint to remove patch sets from system list _by @michalslomczynski_ `[patchman-engine]`

- **2022-04-22**: chore(PatchSet): remove flags for Patch set table and patch set column in Systems table _by @mkholjur_ `[patchman-ui]`

- **2022-04-22**: SPM-1455: fix dev script url _by @Michael Mraka_ `[patchman-engine]`

- **2022-04-22**: chore(deps): update frontend-components-config package _by @mkholjur_ `[patchman-ui]`

- **2022-04-22**: SPM-1455: remove redundant messages which mess clowder param output _by @Michael Mraka_ `[patchman-engine]`

- **2022-04-25**: SPM-1455: remove redundant websocket test from vmaas_sync _by @Michael Mraka_ `[patchman-engine]`

- **2022-04-26**: SPM-1442: Sort baselines list by name ascending _by @michalslomczynski_ `[patchman-engine]`

- **2022-04-26**: chore(release): 1.47.0 _by @semantic-release-bot_ `[patchman-ui]`

- **2022-04-26**: chore(InventoryDetail) remove feature flag, show patch set only when it is present _by @mkholjur_ `[patchman-ui]`

- **2022-04-26**: chore(InventoryDetail): fix tests and reword wizard title _by @mkholjur_ `[patchman-ui]`

- **2022-04-26**: chore(ReviewSystems): fix filters breaking the page _by @mkholjur_ `[patchman-ui]`

- **2022-04-26**: v1.18.69 _by @semantic-release_ `[patchman-engine]`

- **2022-04-26**: SPM-1433: Return only 200 status code as response _by @michalslomczynski_ `[patchman-engine]`

- **2022-04-27**: chore(PatchSet): fix date field dropdown _by @mkholjur_ `[patchman-ui]`

- **2022-04-27**: v1.18.70 _by @semantic-release_ `[patchman-engine]`

- **2022-04-27**: SPM-1460: Check baseline IDs during update _by @michalslomczynski_ `[patchman-engine]`

- **2022-04-27**: SPM-1460: Add test for update with system IDs of another baseline _by @michalslomczynski_ `[patchman-engine]`

- **2022-04-27**: SPM-1460: Improve baselines DB query performance _by @michalslomczynski_ `[patchman-engine]`

- **2022-04-28**: v1.18.71 _by @semantic-release_ `[patchman-engine]`


## 2022-05 (86 commits)

- **2022-05-02**: v1.18.72 _by @semantic-release_ `[patchman-engine]`

- **2022-05-03**: SPM-1469: temporarily force go-toolset-1.17 from centos stream _by @Michael Mraka_ `[patchman-engine]`

- **2022-05-03**: v1.18.73 _by @semantic-release_ `[patchman-engine]`

- **2022-05-03**: SPM-1469: allow build on aarch64 _by @Patrik Segedy_ `[patchman-engine]`

- **2022-05-03**: Update iqe plugin name in deployment.yaml _by @Mike Shriver_ `[content-sources-backend]`

- **2022-05-04**: SPM-1469: pin golangci-lint to v1.44.2 _by @Patrik Segedy_ `[patchman-engine]`

- **2022-05-05**: v1.18.74 _by @semantic-release_ `[patchman-engine]`

- **2022-05-05**: SPM-1459: create rh_account on first request _by @Patrik Segedy_ `[patchman-engine]`

- **2022-05-09**: SPM-1479: use db clowder config directly _by @Michael Mraka_ `[patchman-engine]`

- **2022-05-09**: SPM-1479: use API clowder config directly _by @Michael Mraka_ `[patchman-engine]`

- **2022-05-09**: SPM-1479: use kafka clowder config directly _by @Michael Mraka_ `[patchman-engine]`

- **2022-05-09**: SPM-1479: use service clowder config _by @Michael Mraka_ `[patchman-engine]`

- **2022-05-09**: SPM-1479: use cloudwatch clowder config _by @Michael Mraka_ `[patchman-engine]`

- **2022-05-09**: SPM-1479: use internal cfg for database _by @Michael Mraka_ `[patchman-engine]`

- **2022-05-09**: Initial _by @Andrew Dewar_ `[yummy]`

- **2022-05-09**: Update Readme _by @Andrew Dewar_ `[yummy]`

- **2022-05-10**: SPM-1479: use internal cfg for API _by @Michael Mraka_ `[patchman-engine]`

- **2022-05-10**: SPM-1479: use internal cfg for kafka _by @Michael Mraka_ `[patchman-engine]`

- **2022-05-10**: SPM-1479: use internal cfg for services _by @Michael Mraka_ `[patchman-engine]`

- **2022-05-10**: SPM-1479: use internal cloudwatch config _by @Michael Mraka_ `[patchman-engine]`

- **2022-05-10**: CONTENT-33: initial api work _by @Justin Sherrill_ `[content-sources-backend]`

- **2022-05-10**: SPM-1475: export/system baseline filter _by @Patrik Segedy_ `[patchman-engine]`

- **2022-05-10**: SPM:1495: remove confluent-kafka-go _by @Patrik Segedy_ `[patchman-engine]`

- **2022-05-10**: fix(SPM-1488): remove systems from patch sets _by @mkholjur_ `[patchman-ui]`

- **2022-05-10**: updating Readme and adding arch doc _by @Justin Sherrill_ `[content-sources-backend]`

- **2022-05-10**: Update ElapsedWithMemory with alloc instead of totalAlloc _by @Andrew Dewar_ `[yummy]`

- **2022-05-10**: Separate Extract functions and filter for "rpm" package types only _by @Andrew Dewar_ `[yummy]`

- **2022-05-10**: Bump go version to match fedora 35 _by @Andrew Dewar_ `[yummy]`

- **2022-05-11**: Revert "SPM-1469: temporarily force go-toolset-1.17 from centos stream" _by @Michael Mraka_ `[patchman-engine]`

- **2022-05-11**: Initial _by @Andrew Dewar_ `[yummy]`

- **2022-05-11**: Add Functionality and readme.md _by @Andrew Dewar_ `[yummy]`

- **2022-05-11**: Updated a few names _by @Andrew Dewar_ `[yummy]`

- **2022-05-11**: Updated filename from Elapsed to Time _by @Andrew Dewar_ `[yummy]`

- **2022-05-12**: fix(1494): fix selected systems across steps _by @mkholjur_ `[patchman-ui]`

- **2022-05-12**: fix(SPM-1492): system details page refreshs after patch-set change _by @mkholjur_ `[patchman-ui]`

- **2022-05-12**: fix(SPM-1490, SPM-1493): wizard titles made dynamic _by @mkholjur_ `[patchman-ui]`

- **2022-05-12**: fix(SPM-1489): rename assign action title _by @mkholjur_ `[patchman-ui]`

- **2022-05-12**: CONTENT-39: Repository Configurations model and database _by @Ryan Verdile_ `[content-sources-backend]`

- **2022-05-12**: fix(SPM-1458): change patch set page title _by @mkholjur_ `[patchman-ui]`

- **2022-05-12**: Update names, remove time package _by @Andrew Dewar_ `[yummy]`

- **2022-05-12**: Update package with tests, rename many things _by @Andrew Dewar_ `[yummy]`

- **2022-05-13**: SPM-1505: temporarily switch back to workaround cert issue in stage _by @Michael Mraka_ `[patchman-engine]`

- **2022-05-13**: v1.18.75 _by @semantic-release_ `[patchman-engine]`

- **2022-05-16**: SPM-1505: use RHEL 8.5 images but go 1.17 _by @Michael Mraka_ `[patchman-engine]`

- **2022-05-16**: v1.18.76 _by @semantic-release_ `[patchman-engine]`

- **2022-05-16**: chore(release): 1.47.1 _by @semantic-release-bot_ `[patchman-ui]`

- **2022-05-17**: v1.18.77 _by @semantic-release_ `[patchman-engine]`

- **2022-05-17**: chore(release): 1.47.2 _by @semantic-release-bot_ `[patchman-ui]`

- **2022-05-17**: SPM-1505: remove unused x509ignoreCN=0 _by @Michael Mraka_ `[patchman-engine]`

- **2022-05-17**: fix(SPM-1487): let users create patch sets with zero systems assigned _by @mkholjur_ `[patchman-ui]`

- **2022-05-17**: v1.18.78 _by @semantic-release_ `[patchman-engine]`

- **2022-05-17**: chore(release): 1.47.3 _by @semantic-release-bot_ `[patchman-ui]`

- **2022-05-17**: HMSCONTENT-45 _by @Andrew Dewar_ `[content-sources-backend]`

- **2022-05-17**: Merge pull request #1 from content-services/ReviewBranch [#1] _by @Andrew_ `[yummy]`

- **2022-05-17**: CONTENT-45: Add tests/vet/lint/format to github actions _by @Andrew Dewar_ `[yummy]`

- **2022-05-17**: Merge pull request #2 from content-services/CONTENT-45 [#2] _by @Andrew_ `[yummy]`

- **2022-05-17**: Merge pull request #5 from content-services/HMSCONTENT-45 [#5] _by @Andrew_ `[content-sources-backend]`

- **2022-05-17**: Add temporary build & push script _by @Justin Sherrill_ `[content-sources-backend]`

- **2022-05-17**: build and push container on merge _by @Justin Sherrill_ `[content-sources-backend]`

- **2022-05-19**: SPM-1484: Add base notification _by @michalslomczynski_ `[patchman-engine]`

- **2022-05-19**: SPM-1484: Add notification config _by @michalslomczynski_ `[patchman-engine]`

- **2022-05-19**: SPM-1484: Add advisory notification to evaluator _by @michalslomczynski_ `[patchman-engine]`

- **2022-05-19**: SPM-1484: Add advisory notification test _by @michalslomczynski_ `[patchman-engine]`

- **2022-05-19**: v1.18.79 _by @semantic-release_ `[patchman-engine]`

- **2022-05-19**: Merge pull request #4 from rverdile/my-changes [#4] _by @rverdile_ `[content-sources-backend]`

- **2022-05-23**: SPM-1501: add payload tracker msgs with org_id _by @Patrik Segedy_ `[patchman-engine]`

- **2022-05-23**: CONTENT-58: Update DB config to use clowder _by @Ryan Verdile_ `[content-sources-backend]`

- **2022-05-23**: Refs CONTENT-58: add database to deployment _by @Ryan Verdile_ `[content-sources-backend]`

- **2022-05-24**: trim hash for image pushes _by @Justin Sherrill_ `[content-sources-backend]`

- **2022-05-24**: Update arch doc to use puml _by @Justin Sherrill_ `[content-sources-backend]`

- **2022-05-25**: SPM-1324: let's log whole message when tracing error _by @Michael Mraka_ `[patchman-engine]`

- **2022-05-25**: push only one image at a time _by @Justin Sherrill_ `[content-sources-backend]`

- **2022-05-26**: chore(release): 1.47.4 _by @semantic-release-bot_ `[patchman-ui]`

- **2022-05-26**: SPM-1538: add root ca certs into runtime image _by @Michael Mraka_ `[patchman-engine]`

- **2022-05-26**: Adding missing message _by @Gabi Podolnikova_ `[patchman-ui]`

- **2022-05-26**: v1.18.80 _by @semantic-release_ `[patchman-engine]`

- **2022-05-26**: chore(Messages): fix invalid message IDs _by @mkholjur_ `[patchman-ui]`

- **2022-05-26**: CONTENT-54: Support delete of repositories _by @Justin Sherrill_ `[content-sources-backend]`

- **2022-05-26**: Changes message _by @Gabi Podolnikova_ `[patchman-ui]`

- **2022-05-26**: SPM-1499,SPM-1506: store org_id to db from identity header _by @Patrik Segedy_ `[patchman-engine]`

- **2022-05-26**: CONTENT-60 - expose api over /v1/ _by @Justin Sherrill_ `[content-sources-backend]`

- **2022-05-26**: CONTENT-50: Add zerolog logging (#9) [#9] _by @Andrew_ `[content-sources-backend]`

- **2022-05-26**: CONTENT-40: Create content source _by @Ryan Verdile_ `[content-sources-backend]`

- **2022-05-27**: v1.18.81 _by @semantic-release_ `[patchman-engine]`

- **2022-05-27**: v1.18.82 _by @semantic-release_ `[patchman-engine]`

- **2022-05-27**: CONTENT-59: integrate viper for config loading _by @Justin Sherrill_ `[content-sources-backend]`
