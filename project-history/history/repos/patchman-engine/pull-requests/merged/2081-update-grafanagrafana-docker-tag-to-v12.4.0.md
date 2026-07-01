---
type: pull_request
number: 2081
title: "Update grafana/grafana Docker tag to v12.4.0"
state: merged
author: red-hat-konflux
created: 2026-03-02T05:22:40Z
updated: 2026-03-02T09:18:24Z
closed: 2026-03-02T05:28:22Z
merged: 2026-03-02T05:28:22Z
base_branch: master
head_branch: konflux/mintmaker/master/grafana-monorepo
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2081
---

# Pull Request #2081: Update grafana/grafana Docker tag to v12.4.0

**Author**: @red-hat-konflux
**Created**: March 02, 2026 at 05:22 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/grafana-monorepo`

## Description

> **Note:** This PR body was truncated due to platform limits.

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [grafana/grafana](https://redirect.github.com/grafana/grafana) | final | minor | `12.3.3` -> `12.4.0` |

---

### Release Notes

<details>
<summary>grafana/grafana (grafana/grafana)</summary>

### [`v12.4.0`](https://redirect.github.com/grafana/grafana/blob/HEAD/CHANGELOG.md#1240-2026-02-24)

[Compare Source](https://redirect.github.com/grafana/grafana/compare/v12.3.4...v12.4.0)

##### Features and enhancements

- **API:** Add missing scope check on dashboards [#&#8203;116885](https://redirect.github.com/grafana/grafana/pull/116885), [@&#8203;Proximyst](https://redirect.github.com/Proximyst)
- **Alerting Enrichment:** Add new RBAC permissions for reading and writing enrichments (Enterprise)
- **Alerting:** Add Alert Rules tabs navigation with feature toggle [#&#8203;116253](https://redirect.github.com/grafana/grafana/pull/116253), [@&#8203;aifraenkel](https://redirect.github.com/aifraenkel)
- **Alerting:** Add Alert activity card to alerting home page [#&#8203;115822](https://redirect.github.com/grafana/grafana/pull/115822), [@&#8203;dhalachliyski](https://redirect.github.com/dhalachliyski)
- **Alerting:** Add Cursor frontmatter to CLAUDE.md for auto-loading [#&#8203;115613](https://redirect.github.com/grafana/grafana/pull/115613), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Add Edit/Export actions to group rows, clickable folders, and square icon for recording rules [#&#8203;117763](https://redirect.github.com/grafana/grafana/pull/117763), [@&#8203;konrad147](https://redirect.github.com/konrad147)
- **Alerting:** Add RBAC for enrichment [#&#8203;113296](https://redirect.github.com/grafana/grafana/pull/113296), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Add RBAC to enrichments (Enterprise)
- **Alerting:** Add UI for imported time intervals [#&#8203;116249](https://redirect.github.com/grafana/grafana/pull/116249), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Add alert labels as tags on annotations ([#&#8203;28610](https://redirect.github.com/grafana/grafana/issues/28610)) [#&#8203;116244](https://redirect.github.com/grafana/grafana/pull/116244), [@&#8203;msvechla](https://redirect.github.com/msvechla)
- **Alerting:** Add alertingSyncNotifiersApiMigration feature flag [#&#8203;117946](https://redirect.github.com/grafana/grafana/pull/117946), [@&#8203;rodrigopk](https://redirect.github.com/rodrigopk)
- **Alerting:** Add compressed periodic save for alert instances [#&#8203;111803](https://redirect.github.com/grafana/grafana/pull/111803), [@&#8203;softho0n](https://redirect.github.com/softho0n)
- **Alerting:** Add counts for firing and pending alert rules [#&#8203;113309](https://redirect.github.com/grafana/grafana/pull/113309), [@&#8203;laurenashleigh](https://redirect.github.com/laurenashleigh)
- **Alerting:** Add empty state to triage page WIP [#&#8203;113390](https://redirect.github.com/grafana/grafana/pull/113390), [@&#8203;laurenashleigh](https://redirect.github.com/laurenashleigh)
- **Alerting:** Add expression type information to webhook valueString [#&#8203;112312](https://redirect.github.com/grafana/grafana/pull/112312), [@&#8203;softho0n](https://redirect.github.com/softho0n)
- **Alerting:** Add feature toggle to disable DMA creation in UI [#&#8203;116830](https://redirect.github.com/grafana/grafana/pull/116830), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Add first CLAUDE.md in the frontend alerting folder [#&#8203;114308](https://redirect.github.com/grafana/grafana/pull/114308), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Add folder\_uid label to the grafana\_alerting\_rule\_group\_rules metric [#&#8203;115129](https://redirect.github.com/grafana/grafana/pull/115129), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Add gh in CLAUDE.md [#&#8203;114992](https://redirect.github.com/grafana/grafana/pull/114992), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Add limits for the size of expanded notification templates [#&#8203;115242](https://redirect.github.com/grafana/grafana/pull/115242), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
- **Alerting:** Add managed folder validation frontend [#&#8203;115203](https://redirect.github.com/grafana/grafana/pull/115203), [@&#8203;laurenashleigh](https://redirect.github.com/laurenashleigh)
- **Alerting:** Add policy selector in the alert rule form [#&#8203;117464](https://redirect.github.com/grafana/grafana/pull/117464), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Add saved searches feature for alert rules page [#&#8203;115001](https://redirect.github.com/grafana/grafana/pull/115001), [@&#8203;dhalachliyski](https://redirect.github.com/dhalachliyski)
- **Alerting:** Add viz wrapper for run queries in enrichment (Enterprise)
- **Alerting:** Alerts page performance improvements [#&#8203;113391](https://redirect.github.com/grafana/grafana/pull/113391), [@&#8203;konrad147](https://redirect.github.com/konrad147)
- **Alerting:** Analyze an alert rule with Grafana Assistant [#&#8203;114420](https://redirect.github.com/grafana/grafana/pull/114420), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Apply pending period to NoData and Error alerts [#&#8203;117024](https://redirect.github.com/grafana/grafana/pull/117024), [@&#8203;santihernandezc](https://redirect.github.com/santihernandezc)
- **Alerting:** Change group filtering to search-based using lightweight BE endpoint [#&#8203;114347](https://redirect.github.com/grafana/grafana/pull/114347), [@&#8203;laurenashleigh](https://redirect.github.com/laurenashleigh)
- **Alerting:** Collate alert\_rule.namespace\_uid column as binary [#&#8203;115152](https://redirect.github.com/grafana/grafana/pull/115152), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Collate rule\_group column as binary [#&#8203;114365](https://redirect.github.com/grafana/grafana/pull/114365), [@&#8203;rwwiv](https://redirect.github.com/rwwiv)
- **Alerting:** Config option to set default datasource in Prometheus rule import [#&#8203;115665](https://redirect.github.com/grafana/grafana/pull/115665), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Correct relative paths in CLAUDE.md Required Reading links [#&#8203;114709](https://redirect.github.com/grafana/grafana/pull/114709), [@&#8203;dhalachliyski](https://redirect.github.com/dhalachliyski)
- **Alerting:** Dedicated permission for Template testing API [#&#8203;115032](https://redirect.github.com/grafana/grafana/pull/115032), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
- **Alerting:** Deprecate OpsGenie integration [#&#8203;117085](https://redirect.github.com/grafana/grafana/pull/117085), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
- **Alerting:** Filter out imported contact points from simplified routing dropdown [#&#8203;116408](https://redirect.github.com/grafana/grafana/pull/116408), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Hide DMA options when no manageAlerts datasources exist [#&#8203;115952](https://redirect.github.com/grafana/grafana/pull/115952), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Import to Grafana Alerting Wizard - first iteration [#&#8203;116924](https://redirect.github.com/grafana/grafana/pull/116924), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Improve ASH Loki query efficiency by including folderUID [#&#8203;113322](https://redirect.github.com/grafana/grafana/pull/113322), [@&#8203;JacobsonMT](https://redirect.github.com/JacobsonMT)
- **Alerting:** Improve instance count display [#&#8203;114997](https://redirect.github.com/grafana/grafana/pull/114997), [@&#8203;laurenashleigh](https://redirect.github.com/laurenashleigh)
- **Alerting:** Make AI Helper in triage to use only assistant (Enterprise)
- **Alerting:** Make default notification configuration use empty receiver [#&#8203;116368](https://redirect.github.com/grafana/grafana/pull/116368), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
- **Alerting:** Make saved search name clickable to apply search [#&#8203;116832](https://redirect.github.com/grafana/grafana/pull/116832), [@&#8203;dhalachliyski](https://redirect.github.com/dhalachliyski)
- **Alerting:** Migrate to K8s style receiver testing API [#&#8203;116847](https://redirect.github.com/grafana/grafana/pull/116847), [@&#8203;rodrigopk](https://redirect.github.com/rodrigopk)
- **Alerting:** Notification configuration tabs [#&#8203;116749](https://redirect.github.com/grafana/grafana/pull/116749), [@&#8203;aifraenkel](https://redirect.github.com/aifraenkel)
- **Alerting:** Prevent routing preview from auto-triggering on mount [#&#8203;113749](https://redirect.github.com/grafana/grafana/pull/113749), [@&#8203;laurenashleigh](https://redirect.github.com/laurenashleigh)
- **Alerting:** Prevent users from saving rules to git-synced folders [#&#8203;114944](https://redirect.github.com/grafana/grafana/pull/114944), [@&#8203;rwwiv](https://redirect.github.com/rwwiv)
- **Alerting:** Protected fields for Contact points [#&#8203;115442](https://redirect.github.com/grafana/grafana/pull/115442), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
- **Alerting:** Receiver testing via app platform APIs [#&#8203;111338](https://redirect.github.com/grafana/grafana/pull/111338), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
- **Alerting:** Refactor error for duplicate names on notificationPolicy creation [#&#8203;117797](https://redirect.github.com/grafana/grafana/pull/117797), [@&#8203;rodrigopk](https://redirect.github.com/rodrigopk)
- **Alerting:** Replace the static radio button list for notification routing with a dropdown [#&#8203;117414](https://redirect.github.com/grafana/grafana/pull/117414), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Resize created\_by and updated\_by columns in alert rules tables [#&#8203;113870](https://redirect.github.com/grafana/grafana/pull/113870), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
- **Alerting:** Restrict import UI to admin users only [#&#8203;117441](https://redirect.github.com/grafana/grafana/pull/117441), [@&#8203;rodrigopk](https://redirect.github.com/rodrigopk)
- **Alerting:** Show alert rule scoping in the UI to enrichments list and form (Enterprise)
- **Alerting:** Single alertmanager contact points versions [#&#8203;116076](https://redirect.github.com/grafana/grafana/pull/116076), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Update GOPS labels API calls to v2alpha1 [#&#8203;116327](https://redirect.github.com/grafana/grafana/pull/116327), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Update RuleGroupConfig definitions with missing fields [#&#8203;115850](https://redirect.github.com/grafana/grafana/pull/115850), [@&#8203;JacobsonMT](https://redirect.github.com/JacobsonMT)
- **Alerting:** Update UI of instance counts on triage page [#&#8203;113660](https://redirect.github.com/grafana/grafana/pull/113660), [@&#8203;laurenashleigh](https://redirect.github.com/laurenashleigh)
- **Alerting:** Update createdBy field when silence is being Recreated [#&#8203;115543](https://redirect.github.com/grafana/grafana/pull/115543), [@&#8203;paulojmdias](https://redirect.github.com/paulojmdias)
- **Alerting:** Update docs for ash AI helper button [#&#8203;114229](https://redirect.github.com/grafana/grafana/pull/114229), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Update import wizard to use policyTreeName as config identifier [#&#8203;117382](https://redirect.github.com/grafana/grafana/pull/117382), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Update logic handling canCreate in integrations version, and handle the new deprecated field in the schema [#&#8203;116672](https://redirect.github.com/grafana/grafana/pull/116672), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Update origin for openAssistant in ash (Enterprise)
- **Alerting:** Update prompt for Analyze rule AI button [#&#8203;115341](https://redirect.github.com/grafana/grafana/pull/115341), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Update prompt for the button 'Analyze rule with assistant' button [#&#8203;114593](https://redirect.github.com/grafana/grafana/pull/114593), [@&#8203;konrad147](https://redirect.github.com/konrad147)
- **Alerting:** Update tooltip message when routing preview is disabled [#&#8203;113962](https://redirect.github.com/grafana/grafana/pull/113962), [@&#8203;laurenashleigh](https://redirect.github.com/laurenashleigh)
- **Alerting:** Update translations (Enterprise)
- **Alerting:** Use assistant inline hook instead of llm for template ai button (Enterprise)
- **Alerting:** Use canUse instead of provenance to filter out time intervals [#&#8203;117036](https://redirect.github.com/grafana/grafana/pull/117036), [@&#8203;rodrigopk](https://redirect.github.com/rodrigopk)
- **Alerting:** Use data source headers when remote writing [#&#8203;114528](https://redirect.github.com/grafana/grafana/pull/114528), [@&#8203;santihernandezc](https://redirect.github.com/santihernandezc)
- **AppChrome:** Add proper menu icon for menu, logo icon becomes home [#&#8203;114713](https://redirect.github.com/grafana/grafana/pull/114713), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **Auditing:** Allow configuring Loki retries and timeout (Enterprise)
- **Auditing:** Track uid endpoints for dashboards, not id (Enterprise)
- **Auth:** Add SSO settings PATCH endpoint [#&#8203;117346](https://redirect.github.com/grafana/grafana/pull/117346), [@&#8203;colin-stuart](https://redirect.github.com/colin-stuart)
- **Auth:** Add support for validating OAuth ID token signatures [#&#8203;116442](https://redirect.github.com/grafana/grafana/pull/116442), [@&#8203;DanCech](https://redirect.github.com/DanCech)
- **Auth:** Promote SCIM to GA [#&#8203;116963](https://redirect.github.com/grafana/grafana/pull/116963), [@&#8203;linoman](https://redirect.github.com/linoman)
- **Authz:** Implement Query operation for Zanzana with folder parent retrieval [#&#8203;113483](https://redirect.github.com/grafana/grafana/pull/113483), [@&#8203;mihai-turdean](https://redirect.github.com/mihai-turdean)
- **Avatar:** Require sign-in, remove queue, respect timeout [#&#8203;116891](https://redirect.github.com/grafana/grafana/pull/116891), [@&#8203;macabu](https://redirect.github.com/macabu)
- **Azure Monitor:** Clear filter options in logs builder when key changes [#&#8203;116329](https://redirect.github.com/grafana/grafana/pull/116329), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **Azure:** Improved column handling in logs query builder [#&#8203;114667](https://redirect.github.com/grafana/grafana/pull/114667), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **Azure:** Include aggregate columns in logs builder [#&#8203;114684](https://redirect.github.com/grafana/grafana/pull/114684), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **CandleStick:** Add timeRangePan [#&#8203;113888](https://redirect.github.com/grafana/grafana/pull/113888), [@&#8203;drew08t](https://redirect.github.com/drew08t)
- **Chore:** API: add query params to the spec [#&#8203;117217](https://redirect.github.com/grafana/grafana/pull/117217), [@&#8203;yudintsevegor](https://redirect.github.com/yudintsevegor)
- **Chore:** Access API: add missing query params (Enterprise)
- **Chore:** Deprecate experimental restore dashboard API [#&#8203;116256](https://redirect.github.com/grafana/grafana/pull/116256), [@&#8203;ryantxu](https://redirect.github.com/ryantxu)
- **Chore:** Deprecate the localeFormatPreference feature toggle [#&#8203;116621](https://redirect.github.com/grafana/grafana/pull/116621), [@&#8203;joshhunt](https://redirect.github.com/joshhunt)
- **Chore:** Improve packaging/docker/run.sh [#&#8203;114012](https://redirect.github.com/grafana/grafana/pull/114012), [@&#8203;dmotte](https://redirect.github.com/dmotte)
- **Chore:** RBAC: Migrate role picker to rtkq [#&#8203;116571](https://redirect.github.com/grafana/grafana/pull/116571), [@&#8203;yudintsevegor](https://redirect.github.com/yudintsevegor)
- **Chore:** Remove Drilldown Investigations [#&#8203;115471](https://redirect.github.com/grafana/grafana/pull/115471), [@&#8203;joey-grafana](https://redirect.github.com/joey-grafana)
- **Chore:** Remove `logRequestsInstrumentedAsUnknown` feature flag [#&#8203;116417](https://redirect.github.com/grafana/grafana/pull/116417), [@&#8203;undef1nd](https://redirect.github.com/undef1nd)
- **Chore:** Remove `pinNavItems` feature toggle [#&#8203;113855](https://redirect.github.com/grafana/grafana/pull/113855), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
- **Chore:** Remove `unifiedHistory` feature toggle and associated code [#&#8203;113857](https://redirect.github.com/grafana/grafana/pull/113857), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
- **Chore:** Remove deprecated language\_provider methods in prometheus package [#&#8203;114361](https://redirect.github.com/grafana/grafana/pull/114361), [@&#8203;itsmylife](https://redirect.github.com/itsmylife)
- **Chore:** Remove experimental feature individualCookiePreferences [#&#8203;116374](https://redirect.github.com/grafana/grafana/pull/116374), [@&#8203;hairyhenderson](https://redirect.github.com/hairyhenderson)
- **Chore:** Remove unused+experimental /dashboards/calculate-diff API support [#&#8203;114151](https://redirect.github.com/grafana/grafana/pull/114151), [@&#8203;ryantxu](https://redirect.github.com/ryantxu)
- **Chore:** Rudderstack upgrade to SDK v3 behind flag [#&#8203;114126](https://redirect.github.com/grafana/grafana/pull/114126), [@&#8203;samsch](https://redirect.github.com/samsch)
- **Chore:** Upgrade Grafana Faro to v2, removing `web_vitals_attribution_enabled` [#&#8203;117516](https://redirect.github.com/grafana/grafana/pull/117516), [@&#8203;tskarhed](https://redirect.github.com/tskarhed)
- **Cleanup:** Remove CSV drag-and-drop snapshot query feature [#&#8203;113645](https://redirect.github.com/grafana/grafana/pull/113645), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **Cloud Monitoring:** Add support for Google Cloud universe\_domain [#&#8203;115931](https://redirect.github.com/grafana/grafana/pull/115931), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **CloudMigrations:** Remove feature toggle and introduce config setting to disable it [#&#8203;114223](https://redirect.github.com/grafana/grafana/pull/114223), [@&#8203;macabu](https://redirect.github.com/macabu)
- **CloudWatch Logs:** Hide internal logs field [#&#8203;114121](https://redirect.github.com/grafana/grafana/pull/114121), [@&#8203;kevinwcyu](https://redirect.github.com/kevinwcyu)
- **CloudWatch Logs:** Limit CloudWatch logs queries to use logGroupIdentifiers only for monitoring accounts [#&#8203;113137](https://redirect.github.com/grafana/grafana/pull/113137), [@&#8203;kevinwcyu](https://redirect.github.com/kevinwcyu)
- **CloudWatch Logs:** Select log groups with the log group selector and $\_\_logGroups macro for OpenSearch Structured Query Language queries [#&#8203;116222](https://redirect.github.com/grafana/grafana/pull/116222), [@&#8203;kevinwcyu](https://redirect.github.com/kevinwcyu)
- **CloudWatch:** Add anomaly command to language support, add documentation for anomaly queries [#&#8203;113311](https://redirect.github.com/grafana/grafana/pull/113311), [@&#8203;idastambuk](https://redirect.github.com/idastambuk)
- **CloudWatch:** Add links to data source docs in the config editor [#&#8203;113795](https://redirect.github.com/grafana/grafana/pull/113795), [@&#8203;kevinwcyu](https://redirect.github.com/kevinwcyu)
- **CloudWatch:** Make match exact toggle false by default [#&#8203;113314](https://redirect.github.com/grafana/grafana/pull/113314), [@&#8203;idastambuk](https://redirect.github.com/idastambuk)
- **Cloudwatch:** Make cloudwatchBatchQueries GA [#&#8203;117448](https://redirect.github.com/grafana/grafana/pull/117448), [@&#8203;iwysiu](https://redirect.github.com/iwysiu)
- **Cloudwatch:** Mark missing default region error downstream [#&#8203;117551](https://redirect.github.com/grafana/grafana/pull/117551), [@&#8203;iwysiu](https://redirect.github.com/iwysiu)
- **Cloudwatch:** Update grafana-aws-sdk to 1.4.2 [#&#8203;115855](https://redirect.github.com/grafana/grafana/pull/115855), [@&#8203;iwysiu](https://redirect.github.com/iwysiu)
- **Config:** Set skip migrations in defaults.ini + override when running frontend service locally [#&#8203;114007](https://redirect.github.com/grafana/grafana/pull/114007), [@&#8203;ashharrison90](https://redirect.github.com/ashharrison90)
- **Correlations:** Remove support for org\_id=0 [#&#8203;116877](https://redirect.github.com/grafana/grafana/pull/116877), [@&#8203;gelicia](https://redirect.github.com/gelicia)
- **Dashboard :** Allow applying variable regex to display text [#&#8203;114426](https://redirect.github.com/grafana/grafana/pull/114426), [@&#8203;kristinademeshchik](https://redirect.github.com/kristinademeshchik)
- **Dashboard Controls:** Add UI for displaying under menu [#&#8203;113517](https://redirect.github.com/grafana/grafana/pull/113517), [@&#8203;leventebalogh](https://redirect.github.com/leventebalogh)
- **Dashboard provisioning:** Add support for v2 schema [#&#8203;113620](https://redirect.github.com/grafana/grafana/pull/113620), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
- **Dashboard:** Do not select element always [#&#8203;116986](https://redirect.github.com/grafana/grafana/pull/116986), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **Dashboard:** Hide sidebar in kiosk mode [#&#8203;115387](https://redirect.github.com/grafana/grafana/pull/115387), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **Dashboard:** Hide sidebar on mobile when in view mode [#&#8203;117369](https://redirect.github.com/grafana/grafana/pull/117369), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **Dashboard:** Hide sidebar when playlist is playing [#&#8203;115414](https://redirect.github.com/grafana/grafana/pull/115414), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **Dashboard:** New experimental time range zoom shortcuts [#&#8203;114190](https://redirect.github.com/grafana/grafana/pull/114190), [@&#8203;jesdavpet](https://redirect.github.com/jesdavpet)
- **Dashboard:** Round x/y/w/h when importing a dashboard with floats [#&#8203;117072](https://redirect.github.com/grafana/grafana/pull/117072), [@&#8203;bfmatei](https://redirect.github.com/bfmatei)
- **Dashboards:** Avoid using internal id from the frontend [#&#8203;117398](https://redirect.github.com/grafana/grafana/pull/117398), [@&#8203;ryantxu](https://redirect.github.com/ryantxu)
- **Dashboards:** Do not show alert rules button for new dashboads [#&#8203;115571](https://redirect.github.com/grafana/grafana/pull/115571), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **Dashboards:** Make clear all of variable dropdown accessible by keyboard navigation [#&#8203;117462](https://redirect.github.com/grafana/grafana/pull/117462), [@&#8203;oscarkilhed](https://redirect.github.com/oscarkilhed)
- **Dashboards:** Per panel filtering for timeseries [#&#8203;114499](https://redirect.github.com/grafana/grafana/pull/114499), [@&#8203;mdvictor](https://redirect.github.com/mdvictor)
- **Dashboards:** Prevent memory leak in CUE validation by reusing context only for 100 validations [#&#8203;114818](https://redirect.github.com/grafana/grafana/pull/114818), [@&#8203;MissingRoberto](https://redirect.github.com/MissingRoberto)
- **Dashboards:** Remove deprecated dashboard id endpoints [#&#8203;117227](https://redirect.github.com/grafana/grafana/pull/117227), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
- **DashboardsAPI:** Deprecate /api/dashboards/home [#&#8203;115333](https://redirect.github.com/grafana/grafana/pull/115333), [@&#8203;ryantxu](https://redirect.github.com/ryantxu)
- **DataSources:** Deprecate api routes with name and internal IDs [#&#8203;116391](https://redirect.github.com/grafana/grafana/pull/116391), [@&#8203;ryantxu](https://redirect.github.com/ryantxu)
- **DataSources:** Update SDKs in support of auth service [#&#8203;112101](https://redirect.github.com/grafana/grafana/pull/112101), [@&#8203;njvrzm](https://redirect.github.com/njvrzm)
- **Datagrid:** Deprecate panel [#&#8203;116071](https://redirect.github.com/grafana/grafana/pull/116071), [@&#8203;natellium](https://redirect.github.com/natellium)
- **Datasources:** Experimental API group names use full plugin IDs [#&#8203;112961](https://redirect.github.com/grafana/grafana/pull/112961), [@&#8203;dafydd-t](https://redirect.github.com/dafydd-t)
- **Datasources:** Support new temp creds AWS datasources in auth service (Enterprise)
- **Dependencies:** Bump Go to v1.25.5 [#&#8203;114749](https://redirect.github.com/grafana/grafana/pull/114749), [@&#8203;macabu](https://redirect.github.com/macabu)
- **Docs:** Add Knowledge Graph trace & profile configuration section [#&#8203;117155](https://redirect.github.com/grafana/grafana/pull/117155), [@&#8203;github-actions\[bot\]](https://redirect.github.com/github-actions\[bot])
- **Docs:** Add a "DO NOT MODIFY" warning to the `public/img/*` source code directory [#&#8203;115502](https://redirect.github.com/grafana/grafana/pull/115502), [@&#8203;jesdavpet](https://redirect.github.com/jesdavpet)
- **Docs:** Clarify section title for repeating rows and tabs [#&#8203;115170](https://redirect.github.com/grafana/grafana/pull/115170), [@&#8203;imatwawana](https://redirect.github.com/imatwawana)
- **Docs:** Cleanup enterprise tag usage [#&#8203;114694](https://redirect.github.com/grafana/grafana/pull/114694), [@&#8203;Hipska](https://redirect.github.com/Hipska)
- **Docs:** Cleanup enterprise tag usage (Enterprise)
- **Dynamic Dashboards:** Add new panel button with drag & drop [#&#8203;116276](https://redirect.github.com/grafana/grafana/pull/116276), [@&#8203;idastambuk](https://redirect.github.com/idastambuk)
- **Dynamic Dashboards:** Disallow adding empty row and tab titles [#&#8203;113941](https://redirect.github.com/grafana/grafana/pull/113941), [@&#8203;idastambuk](https://redirect.github.com/idastambuk)
- **Dynamic Dashboards:** Make outline open by default [#&#8203;114146](https://redirect.github.com/grafana/grafana/pull/114146), [@&#8203;idastambuk](https://redirect.github.com/idastambuk)
- **Dynamic Dashboards:** Show hidden variables greyed out [#&#8203;115723](https://redirect.github.com/grafana/grafana/pull/115723), [@&#8203;idastambuk](https://redirect.github.com/idastambuk)
- **EchoSrv:** Enable auto route tracking for Azure App Insights [#&#8203;113354](https://redirect.github.com/grafana/grafana/pull/113354), [@&#8203;joshhunt](https://redirect.github.com/joshhunt)
- **ElasticSearch:** Update annotation time-range properties [#&#8203;115500](https://redirect.github.com/grafana/grafana/pull/115500), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **Elasticsearch:** Add default query mode config setting [#&#8203;112540](https://redirect.github.com/grafana/grafana/pull/112540), [@&#8203;cauemarcondes](https://redirect.github.com/cauemarcondes)
- **Elasticsearch:** Add support for serverless connections [#&#8203;114855](https://redirect.github.com/grafana/grafana/pull/114855), [@&#8203;cauemarcondes](https://redirect.github.com/cauemarcondes)
- **Elasticsearch:** Clear code editor query when switching query types [#&#8203;116318](https://redirect.github.com/grafana/grafana/pull/116318), [@&#8203;Milad93R](https://redirect.github.com/Milad93R)
- **Elasticsearch:** Handle keyed filters buckets and emit frames [#&#8203;113478](https://redirect.github.com/grafana/grafana/pull/113478), [@&#8203;adamyeats](https://redirect.github.com/adamyeats)
- **Elasticsearch:** Raw query editor for DSL [#&#8203;114066](https://redirect.github.com/grafana/grafana/pull/114066), [@&#8203;bossinc](https://redirect.github.com/bossinc)
- **Explore:** Add keyboard shortcut to run queries ([#&#8203;111675](https://redirect.github.com/grafana/grafana/issues/111675)) [#&#8203;115811](https://redirect.github.com/grafana/grafana/pull/115811), [@&#8203;naimeshpatel5295](https://redirect.github.com/naimeshpatel5295)
- **Explore:** Ensure data source is part of query object in internal data links [#&#8203;112949](https://redirect.github.com/grafana/grafana/pull/112949), [@&#8203;ifrost](https://redirect.github.com/ifrost)
- **Explore:** Remove use of AppChrome navbar [#&#8203;114680](https://redirect.github.com/grafana/grafana/pull/114680), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **Explore:** Reset legend when a new query is run [#&#8203;116323](https://redirect.github.com/grafana/grafana/pull/116323), [@&#8203;ifrost](https://redirect.github.com/ifrost)
- **Explore:** Traces query that will work with either logs drilldown or explore [#&#8203;115837](https://redirect.github.com/grafana/grafana/pull/115837), [@&#8203;gtk-grafana](https://redirect.github.com/gtk-grafana)
- **Explore:** Use new Table component [#&#8203;111463](https://redirect.github.com/grafana/grafana/pull/111463), [@&#8203;SamarthBagga](https://redirect.github.com/SamarthBagga)
- **ExternalPlugins:** Restore backward compatability for util function [#&#8203;113735](https://redirect.github.com/grafana/grafana/pull/113735), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **Feat:** Datasources Auth Service (Enterprise)
- **Feat:** Experimental sandbox mode for community & PPT plugins (Enterprise)
- **Feat:** Experimental sandbox mode for community plugins [#&#8203;115936](https://redirect.github.com/grafana/grafana/pull/115936), [@&#8203;njvrzm](https://redirect.github.com/njvrzm)
- **Feat:** Remove experimental `permissionsFilterRemoveSubquery` feature [#&#8203;116405](https://redirect.github.com/grafana/grafana/pull/116405), [@&#8203;papagian](https://redirect.github.com/papagian)
- **FeatureToggle:** Create experimental `timeRangePan` flag [#&#8203;112988](https://redirect.github.com/grafana/grafana/pull/112988), [@&#8203;jesdavpet](https://redirect.github.com/jesdavpet)
- **FeatureToggle:** Enable time range pan zoom flags by default as generally available [#&#8203;116970](https://redirect.github.com/grafana/grafana/pull/116970), [@&#8203;jesdavpet](https://redirect.github.com/jesdavpet)
- **FieldColor:** Add accessible color palettes [#&#8203;114424](https://redirect.github.com/grafana/grafana/pull/114424), [@&#8203;ashharrison90](https://redirect.github.com/ashharrison90)
- **Folders:** Deprecate `getFolderByUID` method [#&#8203;113173](https://redirect.github.com/grafana/grafana/pull/113173), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
- **Folders:** Improve wording for actions and move/delete [#&#8203;114090](https://redirect.github.com/grafana/grafana/pull/114090), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
- **Folders:** Manage folder owner reference [#&#8203;117426](https://redirect.github.com/grafana/grafana/pull/117426), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
- **Folders:** Send permissions query param with app platform for folder picker [#&#8203;114158](https://redirect.github.com/grafana/grafana/pull/114158), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
- **Folders:** Show owner references on folder details pages [#&#8203;116843](https://redirect.github.com/grafana/grafana/pull/116843), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
- **Gauge:** Delete radialbar plugin to avoid migrations [#&#8203;116722](https://redirect.github.com/grafana/grafana/pull/116722), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **Gauge:** Mark grafana/ui export as deprecated [#&#8203;116436](https://redirect.github.com/grafana/grafana/pull/116436), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **Geomap:** Min/Max Zoom options for XYZ Tile Layer [#&#8203;114947](https://redirect.github.com/grafana/grafana/pull/114947), [@&#8203;WoozyMasta](https://redirect.github.com/WoozyMasta)
- **Geomap:** Variable support in the XYZ Tile layer [#&#8203;116654](https://redirect.github.com/grafana/grafana/pull/116654), [@&#8203;WoozyMasta](https://redirect.github.com/WoozyMasta)
- **Go:** Update to 1.25.6 [#&#8203;116394](https://redirect.github.com/grafana/grafana/pull/116394), [@&#8203;macabu](https://redirect.github.com/macabu)
- **Go:** Update to 1.25.7 [#&#8203;117470](https://redirect.github.com/grafana/grafana/pull/117470), [@&#8203;macabu](https://redirect.github.com/macabu)
- **Grafana Cli:** Add admin flush-rbac-seed-assignment command [#&#8203;116716](https://redirect.github.com/grafana/grafana/pull/116716), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
- **Grafana Monitoring:** Enable native HTTP histograms by default, make classic histograms configurable [#&#8203;116534](https://redirect.github.com/grafana/grafana/pull/116534), [@&#8203;undef1nd](https://redirect.github.com/undef1nd)
- **GrafanaBootData:** Deprecate config.apps [#&#8203;115610](https://redirect.github.com/grafana/grafana/pull/115610), [@&#8203;hugohaggmark](https://redirect.github.com/hugohaggmark)
- **GrafanaBootData:** Deprecate config.panels [#&#8203;116918](https://redirect.github.com/grafana/grafana/pull/116918), [@&#8203;hugohaggmark](https://redirect.github.com/hugohaggmark)
- **Graphite:** Revert naming convention changes [#&#8203;117158](https://redirect.github.com/grafana/grafana/pull/117158), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **Heatmap:** Add timeRangePan [#&#8203;113889](https://redirect.github.com/grafana/grafana/pull/113889), [@&#8203;drew08t](https://redirect.github.com/drew08t)
- **Heatmap:** Support for linear y axis [#&#8203;113337](https://redirect.github.com/grafana/grafana/pull/113337), [@&#8203;leeoniya](https://redirect.github.com/leeoniya)
- **I18n:** Ignore dist folder in packages when extracting translations [#&#8203;116532](https://redirect.github.com/grafana/grafana/pull/116532), [@&#8203;aocenas](https://redirect.github.com/aocenas)
- **IAM:** Optionally make refresh tokens required if use\_refresh\_token is enabled [#&#8203;114174](https://redirect.github.com/grafana/grafana/pull/114174), [@&#8203;cinaglia](https://redirect.github.com/cinaglia)
- **InteractiveTable:** Extend sort options with `disableSortRemove` and `sortDescFirst` [#&#8203;115352](https://redirect.github.com/grafana/grafana/pull/115352), [@&#8203;mikkancso](https://redirect.github.com/mikkancso)
- **InteractiveTable:** Prevent reset to first page after `data` property change unless `autoResetPage` property is specified [#&#8203;117546](https://redirect.github.com/grafana/grafana/pull/117546), [@&#8203;darrenjaneczek](https://redirect.github.com/darrenjaneczek)
- **Library Elements:** Deprecate folderFilter query param; update docs for folderFilterUIDs [#&#8203;116048](https://redirect.github.com/grafana/grafana/pull/116048), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
- **Live:** Add configuration for client\_queue\_max\_size [#&#8203;114225](https://redirect.github.com/grafana/grafana/pull/114225), [@&#8203;itsgareth](https://redirect.github.com/itsgareth)
- **Live:** Use namespace rather than OrgID [#&#8203;117275](https://redirect.github.com/grafana/grafana/pull/117275), [@&#8203;ryantxu](https://redirect.github.com/ryantxu)
- **Log Line Context:** Internally manage displayed fields [#&#8203;116912](https://redirect.github.com/grafana/grafana/pull/116912), [@&#8203;matyax](https://redirect.github.com/matyax)
- **Logs Panel:** Added support for transformations when using infinite scrolling [#&#8203;116528](https://redirect.github.com/grafana/grafana/pull/116528), [@&#8203;matyax](https://redirect.github.com/matyax)
- **Logs Panel:** Added support for unwrapped logs with optional columns for displayed fields [#&#8203;117402](https://redirect.github.com/grafana/grafana/pull/117402), [@&#8203;matyax](https://redirect.github.com/matyax)
- **Logs Panel:** Integrate client-side search with Popover Menu [#&#8203;114653](https://redirect.github.com/grafana/grafana/pull/114653), [@&#8203;colega](https://redirect.github.com/colega)
- **Logs Volume:** Show visible range of logs in Explore [#&#8203;114501](https://redirect.github.com/grafana/grafana/pull/114501), [@&#8203;matyax](https://redirect.github.com/matyax)
- **Logs:** Cell format value on inspect should use Code view for arrays, objects, and JSON strings [#&#8203;115037](https://redirect.github.com/grafana/grafana/pull/115037), [@&#8203;L2D2Grafana](https://redirect.github.com/L2D2Grafana)
- **Logs:** Feature flag logRowsPopoverMenu removed [#&#8203;113583](https://redirect.github.com/grafana/grafana/pull/113583), [@&#8203;matyax](https://redirect.github.com/matyax)
- **Logs:** Feature flag logsInfiniteScrolling removed [#&#8203;113585](https://redirect.github.com/grafana/grafana/pull/113585), [@&#8203;matyax](https://redirect.github.com/matyax)
- **Logs:** Improved flexibility of `hasSupplementaryQuerySupport` [#&#8203;115348](https://redirect.github.com/grafana/grafana/pull/115348), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **Logs:** Persist sort order in the Explore URL [#&#8203;114350](https://redirect.github.com/grafana/grafana/pull/114350), [@&#8203;matyax](https://redirect.github.com/matyax)
- **Loki:** Apply default\_manage\_alerts\_ui\_toggle config [#&#8203;112297](https://redirect.github.com/grafana/grafana/pull/112297), [@&#8203;416e64726579](https://redirect.github.com/416e64726579)
- **MSSQL:** Current-user authentication [#&#8203;113977](https://redirect.github.com/grafana/grafana/pull/113977), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **MetricsDrilldown:** Remove `exploreMetricsRelatedLogs` feature toggle [#&#8203;116090](https://redirect.github.com/grafana/grafana/pull/116090), [@&#8203;NWRichmond](https://redirect.github.com/NWRichmond)
- **MySQL:** Add variable query editor support [#&#8203;116900](https://redirect.github.com/grafana/grafana/pull/116900), [@&#8203;yesoreyeram](https://redirect.github.com/yesoreyeram)
- **NPM:** Dispatch to plugin-tools on e2e-selectors changes [#&#8203;115218](https://redirect.github.com/grafana/grafana/pull/115218), [@&#8203;sunker](https://redirect.github.com/sunker)
- **New Logs Panel:** Enable new visualization by default [#&#8203;113340](https://redirect.github.com/grafana/grafana/pull/113340), [@&#8203;matyax](https://redirect.github.com/matyax)
- **News Panel:** Modify pubDate logic to use updated date as fallback [#&#8203;113329](https://redirect.github.com/grafana/grafana/pull/113329), [@&#8203;swiffer](https://redirect.github.com/swiffer)
- **Node Graph:** Use first numeric field as fallback for main stat [#&#8203;116530](https://redirect.github.com/grafana/grafana/pull/116530), [@&#8203;ifrost](https://redirect.github.com/ifrost)
- **PDFTables:** Dynamically shrink font to try and fit whole table in pdf page width (Enterprise)
- **Page:** Background prop to support canvas background for standard layout pages [#&#8203;111174](https://redirect.github.com/grafana/grafana/pull/111174), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **Panel Menu:** Allow using icons for link extensions [#&#8203;114836](https://redirect.github.com/grafana/grafana/pull/114836), [@&#8203;leventebalogh](https://redirect.github.com/leventebalogh)
- **Panel visualizations:** Focus on search input when changing visualizations [#&#8203;115484](https://redirect.github.com/grafana/grafana/pull/115484), [@&#8203;idastambuk](https://redirect.github.com/idastambuk)
- **PanelChrome:** Enable new panel padding by default [#&#8203;114492](https://redirect.github.com/grafana/grafana/pull/114492), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **PanelChrome:** Feature toggle increased panel header height and padding [#&#8203;112613](https://redirect.github.com/grafana/grafana/pull/112613), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **Playlists:** Graduate to v1 apis [#&#8203;117638](https://redirect.github.com/grafana/grafana/pull/117638), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
- **Plugin Metrics:** Improve metrics on long duration queries within grafana [#&#8203;116371](https://redirect.github.com/grafana/grafana/pull/116371), [@&#8203;sarahzinger](https://redirect.github.com/sarahzinger)
- **PostgreSQL:** Add variable query editor support [#&#8203;115974](https://redirect.github.com/grafana/grafana/pull/115974), [@&#8203;yesoreyeram](https://redirect.github.com/yesoreyeram)
- **PostgreSQL:** Remove feature toggle `postgresDSUsePGX` [#&#8203;113675](https://redirect.github.com/grafana/grafana/pull/113675), [@&#8203;zoltanbedi](https://redirect.github.com/zoltanbedi)
- **Preferences:** Add API validation and update documentation [#&#8203;116045](https://redirect.github.com/grafana/grafana/pull/116045), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
- **Prometheus Dashboards:** Use $\_\_rate\_interval instead of hardcoded value [#&#8203;111899](https://redirect.github.com/grafana/grafana/pull/111899), [@&#8203;attu0](https://redirect.github.com/attu0)
- **Prometheus:** Add variable job and replaced hardcoded values in prometheus 2.0 stats dashboard [#&#8203;115916](https://redirect.github.com/grafana/grafana/pull/115916), [@&#8203;saurabh007007](https://redirect.github.com/saurabh007007)
- **Prometheus:** Hide 'Kick start your query' button for existing queries [#&#8203;113980](https://redirect.github.com/grafana/grafana/pull/113980), [@&#8203;priyansh3006](https://redirect.github.com/priyansh3006)
- **Prometheus:** Introduce failsafe PromQueryFormat unmarshalling [#&#8203;116670](https://redirect.github.com/grafana/grafana/pull/116670), [@&#8203;itsmylife](https://redirect.github.com/itsmylife)
- **Prometheus:** Introduce filtering /series endpoint for prometheus versions that don't support match\[] parameter [#&#8203;116648](https://redirect.github.com/grafana/grafana/pull/116648), [@&#8203;itsmylife](https://redirect.github.com/itsmylife)
- **Prometheus:** Optimize regex pattern for multi-value label matchers [#&#8203;116233](https://redirect.github.com/grafana/grafana/pull/116233), [@&#8203;Krishnachaitanyakc](https://redirect.github.com/Krishnachaitanyakc)
- **Prometheus:** Revert "Prometheus: Make sure "Min Step" has precedence ([#&#8203;115941](https://redirect.github.com/grafana/grafana/issues/115941))" [#&#8203;116959](https://redirect.github.com/grafana/grafana/pull/116959), [@&#8203;ellisda](https://redirect.github.com/ellisda)
- **Provisioning:** Enable editing dashboard via JSON model [#&#8203;115420](https://redirect.github.com/grafana/grafana/pull/115420), [@&#8203;Clarity-89](https://redirect.github.com/Clarity-89)
- **Provisioning:** Integrate GH app connections into the wizard flow [#&#8203;116547](https://redirect.github.com/grafana/grafana/pull/116547), [@&#8203;Clarity-89](https://redirect.github.com/Clarity-89)
- **Pyroscope:** Exemplar support for series queries [#&#8203;113926](https://redirect.github.com/grafana/grafana/pull/113926), [@&#8203;alsoba13](https://redirect.github.com/alsoba13)
- **Query Editor:** Add Query Options footer and sidebar for new query editor [#&#8203;117403](https://redirect.github.com/grafana/grafana/pull/117403), [@&#8203;Develer](https://redirect.github.com/Develer)
- **QueryEditorRows:** Clear hideSeriesFrom override on query edit [#&#8203;114315](https://redirect.github.com/grafana/grafana/pull/114315), [@&#8203;Sergej-Vlasov](https://redirect.github.com/Sergej-Vlasov)
- **Reporting:** Productize reporting retries feature [#&#8203;117378](https://redirect.github.com/grafana/grafana/pull/117378), [@&#8203;macabu](https://redirect.github.com/macabu)
- **Reporting:** Remove newPDFRendering feature flag, stabilising it (Enterprise)
- **Reporting:** Support editing template variables in the form for dashboards v2 (Enterprise)
- **Restore dashboards:** Improve permissions [#&#8203;116266](https://redirect.github.com/grafana/grafana/pull/116266), [@&#8203;Clarity-89](https://redirect.github.com/Clarity-89)
- **SQL Expressions:** Add "NOT" keyword to allow list [#&#8203;116802](https://redirect.github.com/grafana/grafana/pull/116802), [@&#8203;net0pyr](https://redirect.github.com/net0pyr)
- **SQLDataSource:** Use UID rather than internal ID [#&#8203;116461](https://redirect.github.com/grafana/grafana/pull/116461), [@&#8203;ryantxu](https://redirect.github.com/ryantxu)
- **SQLExpressions:** Add new schema inspector panel [#&#8203;113545](https://redirect.github.com/grafana/grafana/pull/113545), [@&#8203;alexjonspencer1](https://redirect.github.com/alexjonspencer1)
- **Scopes:** Scope input UI update [#&#8203;114002](https://redirect.github.com/grafana/grafana/pull/114002), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **Search:** Move experimental panelTitleSearch from searchV2 to unified search [#&#8203;116326](https://redirect.github.com/grafana/grafana/pull/116326), [@&#8203;ryantxu](https://redirect.github.com/ryantxu)
- **SearchAPI:** Return "shared with me" children based on the permission query param [#&#8203;116254](https://redirect.github.com/grafana/grafana/pull/116254), [@&#8203;aocenas](https://redirect.github.com/aocenas)
- **Secrets Keeper:** Add secretsKeeperUI feature flag [#&#8203;117427](https://redirect.github.com/grafana/grafana/pull/117427), [@&#8203;ericrshields](https://redirect.github.com/ericrshields)
- **Secrets Keeper:** UI shell with tab navigation (Enterprise)
- **Security:** Sanitize TraceView html [#&#8203;117853](https://redirect.github.com/grafana/grafana/pull/117853), [@&#8203;github-actions\[bot\]](https://redirect.github.com/github-actions\[bot])
- **Security:** Use dashboard timerange if time selection disabled [#&#8203;117854](https://redirect.github.com/grafana/grafana/pull/117854), [@&#8203;dana-axinte](https://redirect.github.com/dana-axinte)
- **SelectBase:** Use standard portal container [#&#8203;114844](https://redirect.github.com/grafana/grafana/pull/114844), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **Short URL:** Change default expiration to never [#&#8203;115029](https://redirect.github.com/grafana/grafana/pull/115029), [@&#8203;nmarrs](https://redirect.github.com/nmarrs)
- **Sidebar:** A new reusable component for side toolbars and panes [#&#8203;114141](https://redirect.github.com/grafana/grafana/pull/114141), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **Span Details:** Two-column view [#&#8203;112856](https://redirect.github.com/grafana/grafana/pull/112856), [@&#8203;ifrost](https://redirect.github.com/ifrost)
- **Sparkline:** Improve min/max logic to avoid issues for very narrow deltas [#&#8203;115030](https://redirect.github.com/grafana/grafana/pull/115030), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **Sparkline:** Prevent infinite loop when rendering a sparkline with a single value [#&#8203;114203](https://redirect.github.com/grafana/grafana/pull/114203), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **Table:** Add title attribute to make truncated headings legible [#&#8203;115155](https://redirect.github.com/grafana/grafana/pull/115155), [@&#8203;jesdavpet](https://redirect.github.com/jesdavpet)
- **Table:** Clamp Safari exclusions to 26.0 and 26.1 [#&#8203;114454](https://redirect.github.com/grafana/grafana/pull/114454), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **Table:** Clean up filter popover layout and improve filter selection UX [#&#8203;114052](https://redirect.github.com/grafana/grafana/pull/114052), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **Table:** Keyboard accessibility for filter [#&#8203;117354](https://redirect.github.com/grafana/grafana/pull/117354), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **Table:** Remove hardcoded assumption of \_\_nestedFrames field name [#&#8203;115117](https://redirect.github.com/grafana/grafana/pull/115117), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **TeamFolders:** Show team folders in folder picker [#&#8203;117381](https://redirect.github.com/grafana/grafana/pull/117381), [@&#8203;aocenas](https://redirect.github.com/aocenas)
- **Tempo:** Encode header values before adding them to outgoing context [#&#8203;117279](https://redirect.github.com/grafana/grafana/pull/117279), [@&#8203;jcolladokuri](https://redirect.github.com/jcolladokuri)
- **Tempo:** Remove forwarding incoming and team headers for streaming requests [#&#8203;117813](https://redirect.github.com/grafana/grafana/pull/117813), [@&#8203;jcolladokuri](https://redirect.github.com/jcolladokuri)
- **Theme:** Add breakpoint methods for container queries [#&#8203;113619](https://redirect.github.com/grafana/grafana/pull/113619), [@&#8203;MattIPv4](https://redirect.github.com/MattIPv4)
- **TimePicker:** Show new shortcut for zoom out when experimental flag toggled on [#&#8203;114506](https://redirect.github.com/grafana/grafana/pull/114506), [@&#8203;jesdavpet](https://redirect.github.com/jesdavpet)
- **TimeRange:** Additional keyboard shortcut `t =` to complement `t +` for zoom in [#&#8203;115022](https://redirect.github.com/grafana/grafana/pull/115022), [@&#8203;jesdavpet](https://redirect.github.com/jesdavpet)
- **TimeRange:** Avoid x-axis pan jump caused by data loading latency [#&#8203;114496](https://redirect.github.com/grafana/grafana/pull/114496), [@&#8203;jesdavpet](https://redirect.github.com/jesdavpet)
- **TimeSeries:** X-axis (time range) click-and-drag panning in panel [#&#8203;112982](https://redirect.github.com/grafana/grafana/pull/112982), [@&#8203;jesdavpet](https://redirect.github.com/jesdavpet)
- **Timeline:** Add timeRangePan [#&#8203;113890](https://redirect.github.com/grafana/grafana/pull/113890), [@&#8203;drew08t](https://redirect.github.com/drew08t)
- **Timeseries:** Change mouse cursors to indicate active x-axis and y-axis zoom interactions [#&#8203;113465](https://redirect.github.com/grafana/grafana/pull/113465), [@&#8203;jesdavpet](https://redirect.github.com/jesdavpet)
- **Timeseries:** More nuanced editing of linear threshold to avoid crashes [#&#8203;112301](https://redirect.github.com/grafana/grafana/pull/112301), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **Trace View:** Span filters updated to use combobox filters [#&#8203;112287](https://redirect.github.com/grafana/grafana/pull/112287), [@&#8203;adrapereira](https://redirect.github.com/adrapereira)
- **Trace datasources:** Add Victoria Metrics support for "traces to metrics" [#&#8203;114962](https://redirect.github.com/grafana/grafana/pull/114962), [@&#8203;arturminchukov](https://redirect.github.com/arturminchukov)
- **Transformers:** Add smoothing transformer [#&#8203;111077](https://redirect.github.com/grafana/grafana/pull/111077), [@&#8203;vesalaakso-oura](https://redirect.github.com/vesalaakso-oura)
- **UI Extensions:** Add `openInNewTab` property to link extensions [#&#8203;114831](https://redirect.github.com/grafana/grafana/pull/114831), [@&#8203;leventebalogh](https://redirect.github.com/leventebalogh)
- **UI:** Use react-table column header types in InteractiveTable with story and tests [#&#8203;116091](https://redirect.github.com/grafana/grafana/pull/116091), [@&#8203;Alan-eMartin](https://redirect.github.com/Alan-eMartin)
- **Unified:** Run resource data migrations at startup [#&#8203;114857](https://redirect.github.com/grafana/grafana/pull/114857), [@&#8203;RafaelPaulovic](https://redirect.github.com/RafaelPaulovic)
- **Viz:** Update OutsideRangePlugin to support single datapoint [#&#8203;117278](https://redirect.github.com/grafana/grafana/pull/117278), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)

##### Bug fixes

- **Alerting:** Add support for client certificate authentication and TLS options to External Alertmanager [#&#8203;115716](https://redirect.github.com/grafana/grafana/pull/115716), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
- **Alerting:** Bug fix for regex matching in Alerts page [#&#8203;113400](https://redirect.github.com/grafana/grafana/pull/113400), [@&#8203;laurenashleigh](https://redirect.github.com/laurenashleigh)
- **Alerting:** Fix 'Rule group does not exist' error toast ([#&#8203;101949](https://redirect.github.com/grafana/grafana/issues/101949)) [#&#8203;114766](https://redirect.github.com/grafana/grafana/pull/114766), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Fix Alerts page filtering [#&#8203;115178](https://redirect.github.com/grafana/grafana/pull/115178), [@&#8203;laurenashleigh](https://redirect.github.com/laurenashleigh)
- **Alerting:** Fix NotificationPreview permission checking [#&#8203;114303](https://redirect.github.com/grafana/grafana/pull/114303), [@&#8203;konrad147](https://redirect.github.com/konrad147)
- **Alerting:** Fix RuleEditorCloudRules test flakiness in CI [#&#8203;114695](https://redirect.github.com/grafana/grafana/pull/114695), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Fix a race condition panic in ResetStateByRuleUID [#&#8203;115662](https://redirect.github.com/grafana/grafana/pull/115662), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Fix advanced filter not preserving freewords filter in the list view [#&#8203;114651](https://redirect.github.com/grafana/grafana/pull/114651), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Fix alert rule last evaluation duration units [#&#8203;117814](https://redirect.github.com/grafana/grafana/pull/117814), [@&#8203;JacobsonMT](https://redirect.github.com/JacobsonMT)
- **Alerting:** Fix alert rule last evaluation time including scheduling delays [#&#8203;117819](https://redirect.github.com/grafana/grafana/pull/117819), [@&#8203;JacobsonMT](https://redirect.github.com/JacobsonMT)
- **Alerting:** Fix creating a new alert rule vesion when only keep\_firing\_for changes [#&#8203;114926](https://redirect.github.com/grafana/grafana/pull/114926), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Fix data source recording rules editor [#&#8203;113363](https://redirect.github.com/grafana/grafana/pull/113363), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Fix error when updating Alertmanager config with autogenerated receivers [#&#8203;113710](https://redirect.github.com/grafana/grafana/pull/113710), [@&#8203;moustafab](https://redirect.github.com/moustafab)
- **Alerting:** Fix expression queries when coming from a panel [#&#8203;114095](https://redirect.github.com/grafana/grafana/pull/114095), [@&#8203;konrad147](https://redirect.github.com/konrad147)
- **Alerting:** Fix file import/export of recording rules with target datasource uid [#&#8203;115663](https://redirect.github.com/grafana/grafana/pull/115663), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Fix for fetching evaluation group in new filter [#&#8203;113694](https://redirect.github.com/grafana/grafana/pull/113694), [@&#8203;laurenashleigh](https://redirect.github.com/laurenashleigh)
- **Alerting:** Fix groupBy in simplified routing UI [#&#8203;117076](https://redirect.github.com/grafana/grafana/pull/117076), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Fix ignored filters when paginating alert rules in the API [#&#8203;114710](https://redirect.github.com/grafana/grafana/pull/114710), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Fix label 

</details>

---

### Configuration

📅 **Schedule**: Branch creation - Between 03:00 AM and 10:59 AM, only on Monday ( * 3-10 * * 1 ) in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Enabled.

♻ **Rebasing**: Whenever PR is behind base branch, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

---
### Documentation

Find out how to configure dependency updates in [MintMaker documentation](https://konflux-ci.dev/docs/mintmaker/user/) or see all available configuration options in [Renovate documentation](https://docs.renovatebot.com/configuration-options/).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0Mi4yNi41LXJwbSIsInVwZGF0ZWRJblZlciI6IjQyLjI2LjUtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @github-actions on March 02, 2026 at 05:22 AM UTC

<!-- sc-environment-impact-check -->
## SC Environment Impact Assessment

**Overall Impact:** ⚪ **NONE**

No SC Environment-specific impacts detected in this PR.

<details>
<summary>What was checked</summary>

This PR was automatically scanned for:
- Database migrations
- ClowdApp configuration changes
- Kessel integration changes
- AWS service integrations (S3, RDS, ElastiCache)
- Kafka topic changes
- Secrets management changes
- External dependencies
</details>

### Comment by @codecov-commenter on March 02, 2026 at 05:27 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2081?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.45%. Comparing base ([`c28e17b`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/c28e17b01b0e0cc2dc83b9be3a0ff3b3547a120a?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`3372d78`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/3372d78a61c6a957657f4d95580b93db185c90c5?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2081   +/-   ##
=======================================
  Coverage   59.45%   59.45%           
=======================================
  Files         134      134           
  Lines        8699     8699           
=======================================
  Hits         5172     5172           
  Misses       2981     2981           
  Partials      546      546           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2081/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2081/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.45% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2081?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2081*
