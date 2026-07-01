---
type: pull_request
number: 1197
title: "[Snyk] Upgrade axios from 1.7.2 to 1.7.7"
state: closed
author: jiridostal
created: 2024-09-28T07:03:30Z
updated: 2024-10-25T20:43:05Z
closed: 2024-10-25T20:43:05Z
base_branch: master
head_branch: snyk-upgrade-82cab48c069cec6dbbca5c6a27b93741
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1197
---

# Pull Request #1197: [Snyk] Upgrade axios from 1.7.2 to 1.7.7

**Author**: @jiridostal
**Created**: September 28, 2024 at 07:03 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `snyk-upgrade-82cab48c069cec6dbbca5c6a27b93741`

## Description

![snyk-top-banner](https://github.com/andygongea/OWASP-Benchmark/assets/818805/c518c423-16fe-447e-b67f-ad5a49b5d123)


<h3>Snyk has created this PR to upgrade axios from 1.7.2 to 1.7.7.</h3>

:information_source: Keep your dependencies up-to-date. This makes it easier to fix existing vulnerabilities and to more quickly identify and fix newly disclosed vulnerabilities when they affect your project.

<hr/>


- The recommended version is **5 versions** ahead of your current version.

- The recommended version was released on **a month ago**.

#### Issues fixed by the recommended upgrade:

|  | Issue | Score | Exploit Maturity |
:-------------------------:|:-------------------------|:-------------------------|:-------------------------
![high severity](https://res.cloudinary.com/snyk/image/upload/w_20,h_20/v1561977819/icon/h.png 'high severity') | Server-side Request Forgery (SSRF)<br/>[SNYK-JS-AXIOS-7361793](https://snyk.io/vuln/SNYK-JS-AXIOS-7361793) | **761** | Proof of Concept 



<details>
<summary><b>Release notes</b></summary>
<br/>
  <details>
    <summary>Package name: <b>axios</b></summary>
    <ul>
      <li>
        <b>1.7.7</b> - <a href="https://github.com/axios/axios/releases/tag/v1.7.7">2024-08-31</a></br><h2>Release notes:</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>fetch:</strong> fix stream handling in Safari by fallback to using a stream reader instead of an async iterator; (<a href="https://github.com/axios/axios/issues/6584" data-hovercard-type="pull_request" data-hovercard-url="/axios/axios/pull/6584/hovercard">#6584</a>) (<a href="https://github.com/axios/axios/commit/d1980854fee1765cd02fa0787adf5d6e34dd9dcf">d198085</a>)</li>
<li><strong>http:</strong> fixed support for IPv6 literal strings in url (<a href="https://github.com/axios/axios/issues/5731" data-hovercard-type="pull_request" data-hovercard-url="/axios/axios/pull/5731/hovercard">#5731</a>) (<a href="https://github.com/axios/axios/commit/364993f0d8bc6e0e06f76b8a35d2d0a35cab054c">364993f</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><a target="_blank" rel="noopener noreferrer nofollow" href="https://avatars.githubusercontent.com/u/10539109?v=4&amp;s=18"><img src="https://avatars.githubusercontent.com/u/10539109?v=4&amp;s=18" alt="avatar" width="18" style="max-width: 100%;"></a> <a href="https://github.com/Rishi556" title="+39/-1 (#5731 )">Rishi556</a></li>
<li><a target="_blank" rel="noopener noreferrer nofollow" href="https://avatars.githubusercontent.com/u/12586868?v=4&amp;s=18"><img src="https://avatars.githubusercontent.com/u/12586868?v=4&amp;s=18" alt="avatar" width="18" style="max-width: 100%;"></a> <a href="https://github.com/DigitalBrainJS" title="+27/-7 (#6584 )">Dmitriy Mozgovoy</a></li>
</ul>
      </li>
      <li>
        <b>1.7.6</b> - <a href="https://github.com/axios/axios/releases/tag/v1.7.6">2024-08-30</a></br><h2>Release notes:</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>fetch:</strong> fix content length calculation for FormData payload; (<a href="https://github.com/axios/axios/issues/6524" data-hovercard-type="pull_request" data-hovercard-url="/axios/axios/pull/6524/hovercard">#6524</a>) (<a href="https://github.com/axios/axios/commit/085f56861a83e9ac02c140ad9d68dac540dfeeaa">085f568</a>)</li>
<li><strong>fetch:</strong> optimize signals composing logic; (<a href="https://github.com/axios/axios/issues/6582" data-hovercard-type="pull_request" data-hovercard-url="/axios/axios/pull/6582/hovercard">#6582</a>) (<a href="https://github.com/axios/axios/commit/df9889b83c2cc37e9e6189675a73ab70c60f031f">df9889b</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><a target="_blank" rel="noopener noreferrer nofollow" href="https://avatars.githubusercontent.com/u/12586868?v=4&amp;s=18"><img src="https://avatars.githubusercontent.com/u/12586868?v=4&amp;s=18" alt="avatar" width="18" style="max-width: 100%;"></a> <a href="https://github.com/DigitalBrainJS" title="+98/-46 (#6582 )">Dmitriy Mozgovoy</a></li>
<li><a target="_blank" rel="noopener noreferrer nofollow" href="https://avatars.githubusercontent.com/u/3534453?v=4&amp;s=18"><img src="https://avatars.githubusercontent.com/u/3534453?v=4&amp;s=18" alt="avatar" width="18" style="max-width: 100%;"></a> <a href="https://github.com/jacquesg" title="+5/-1 (#6524 )">Jacques Germishuys</a></li>
<li><a target="_blank" rel="noopener noreferrer nofollow" href="https://avatars.githubusercontent.com/u/53894505?v=4&amp;s=18"><img src="https://avatars.githubusercontent.com/u/53894505?v=4&amp;s=18" alt="avatar" width="18" style="max-width: 100%;"></a> <a href="https://github.com/kuroino721" title="+3/-1 (#6575 )">kuroino721</a></li>
</ul>
      </li>
      <li>
        <b>1.7.5</b> - <a href="https://github.com/axios/axios/releases/tag/v1.7.5">2024-08-23</a></br><h2>Release notes:</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>adapter:</strong> fix undefined reference to hasBrowserEnv (<a href="https://github.com/axios/axios/issues/6572" data-hovercard-type="pull_request" data-hovercard-url="/axios/axios/pull/6572/hovercard">#6572</a>) (<a href="https://github.com/axios/axios/commit/7004707c4180b416341863bd86913fe4fc2f1df1">7004707</a>)</li>
<li><strong>core:</strong> add the missed implementation of AxiosError#status property; (<a href="https://github.com/axios/axios/issues/6573" data-hovercard-type="pull_request" data-hovercard-url="/axios/axios/pull/6573/hovercard">#6573</a>) (<a href="https://github.com/axios/axios/commit/6700a8adac06942205f6a7a21421ecb36c4e0852">6700a8a</a>)</li>
<li><strong>core:</strong> fix <code>ReferenceError: navigator is not defined</code> for custom environments; (<a href="https://github.com/axios/axios/issues/6567" data-hovercard-type="pull_request" data-hovercard-url="/axios/axios/pull/6567/hovercard">#6567</a>) (<a href="https://github.com/axios/axios/commit/fed1a4b2d78ed4a588c84e09d32749ed01dc2794">fed1a4b</a>)</li>
<li><strong>fetch:</strong> fix credentials handling in Cloudflare workers (<a href="https://github.com/axios/axios/issues/6533" data-hovercard-type="pull_request" data-hovercard-url="/axios/axios/pull/6533/hovercard">#6533</a>) (<a href="https://github.com/axios/axios/commit/550d885eb90fd156add7b93bbdc54d30d2f9a98d">550d885</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><a target="_blank" rel="noopener noreferrer nofollow" href="https://avatars.githubusercontent.com/u/12586868?v=4&amp;s=18"><img src="https://avatars.githubusercontent.com/u/12586868?v=4&amp;s=18" alt="avatar" width="18" style="max-width: 100%;"></a> <a href="https://github.com/DigitalBrainJS" title="+187/-83 (#6573 #6567 #6566 #6564 #6563 #6557 #6556 #6555 #6554 #6552 )">Dmitriy Mozgovoy</a></li>
<li><a target="_blank" rel="noopener noreferrer nofollow" href="https://avatars.githubusercontent.com/u/2495809?v=4&amp;s=18"><img src="https://avatars.githubusercontent.com/u/2495809?v=4&amp;s=18" alt="avatar" width="18" style="max-width: 100%;"></a> <a href="https://github.com/antoninbas" title="+6/-6 (#6572 )">Antonin Bas</a></li>
<li><a target="_blank" rel="noopener noreferrer nofollow" href="https://avatars.githubusercontent.com/u/5406212?v=4&amp;s=18"><img src="https://avatars.githubusercontent.com/u/5406212?v=4&amp;s=18" alt="avatar" width="18" style="max-width: 100%;"></a> <a href="https://github.com/hansottowirtz" title="+4/-1 (#6533 )">Hans Otto Wirtz</a></li>
</ul>
      </li>
      <li>
        <b>1.7.4</b> - <a href="https://github.com/axios/axios/releases/tag/v1.7.4">2024-08-13</a></br><h2>Release notes:</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>sec:</strong> <a title="CVE-2024-39338" data-hovercard-type="advisory" data-hovercard-url="/advisories/GHSA-8hc4-vh64-cxmj/hovercard" href="https://github.com/advisories/GHSA-8hc4-vh64-cxmj">CVE-2024-39338</a> (<a href="https://github.com/axios/axios/issues/6539" data-hovercard-type="pull_request" data-hovercard-url="/axios/axios/pull/6539/hovercard">#6539</a>) (<a href="https://github.com/axios/axios/issues/6543" data-hovercard-type="pull_request" data-hovercard-url="/axios/axios/pull/6543/hovercard">#6543</a>) (<a href="https://github.com/axios/axios/commit/6b6b605eaf73852fb2dae033f1e786155959de3a">6b6b605</a>)</li>
<li><strong>sec:</strong> disregard protocol-relative URL to remediate SSRF (<a href="https://github.com/axios/axios/issues/6539" data-hovercard-type="pull_request" data-hovercard-url="/axios/axios/pull/6539/hovercard">#6539</a>) (<a href="https://github.com/axios/axios/commit/07a661a2a6b9092c4aa640dcc7f724ec5e65bdda">07a661a</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><a target="_blank" rel="noopener noreferrer nofollow" href="https://avatars.githubusercontent.com/u/31389480?v=4&amp;s=18"><img src="https://avatars.githubusercontent.com/u/31389480?v=4&amp;s=18" alt="avatar" width="18" style="max-width: 100%;"></a> <a href="https://github.com/levpachmanov" title="+47/-11 (#6543 )">Lev Pachmanov</a></li>
<li><a target="_blank" rel="noopener noreferrer nofollow" href="https://avatars.githubusercontent.com/u/41283691?v=4&amp;s=18"><img src="https://avatars.githubusercontent.com/u/41283691?v=4&amp;s=18" alt="avatar" width="18" style="max-width: 100%;"></a> <a href="https://github.com/hainenber" title="+49/-4 (#6539 )">Đỗ Trọng Hải</a></li>
</ul>
      </li>
      <li>
        <b>1.7.3</b> - <a href="https://github.com/axios/axios/releases/tag/v1.7.3">2024-08-01</a></br><h2>Release notes:</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>adapter:</strong> fix progress event emitting; (<a href="https://github.com/axios/axios/issues/6518" data-hovercard-type="pull_request" data-hovercard-url="/axios/axios/pull/6518/hovercard">#6518</a>) (<a href="https://github.com/axios/axios/commit/e3c76fc9bdd03aa4d98afaf211df943e2031453f">e3c76fc</a>)</li>
<li><strong>fetch:</strong> fix withCredentials request config (<a href="https://github.com/axios/axios/issues/6505" data-hovercard-type="pull_request" data-hovercard-url="/axios/axios/pull/6505/hovercard">#6505</a>) (<a href="https://github.com/axios/axios/commit/85d4d0ea0aae91082f04e303dec46510d1b4e787">85d4d0e</a>)</li>
<li><strong>xhr:</strong> return original config on errors from XHR adapter (<a href="https://github.com/axios/axios/issues/6515" data-hovercard-type="pull_request" data-hovercard-url="/axios/axios/pull/6515/hovercard">#6515</a>) (<a href="https://github.com/axios/axios/commit/8966ee7ea62ecbd6cfb39a905939bcdab5cf6388">8966ee7</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><a target="_blank" rel="noopener noreferrer nofollow" href="https://avatars.githubusercontent.com/u/12586868?v=4&amp;s=18"><img src="https://avatars.githubusercontent.com/u/12586868?v=4&amp;s=18" alt="avatar" width="18" style="max-width: 100%;"></a> <a href="https://github.com/DigitalBrainJS" title="+211/-159 (#6518 #6519 )">Dmitriy Mozgovoy</a></li>
<li><a target="_blank" rel="noopener noreferrer nofollow" href="https://avatars.githubusercontent.com/u/10867286?v=4&amp;s=18"><img src="https://avatars.githubusercontent.com/u/10867286?v=4&amp;s=18" alt="avatar" width="18" style="max-width: 100%;"></a> <a href="https://github.com/ValeraS" title="+3/-3 (#6515 )">Valerii Sidorenko</a></li>
<li><a target="_blank" rel="noopener noreferrer nofollow" href="https://avatars.githubusercontent.com/u/8599535?v=4&amp;s=18"><img src="https://avatars.githubusercontent.com/u/8599535?v=4&amp;s=18" alt="avatar" width="18" style="max-width: 100%;"></a> <a href="https://github.com/prianyu" title="+2/-2 (#6505 )">prianYu</a></li>
</ul>
      </li>
      <li>
        <b>1.7.2</b> - <a href="https://github.com/axios/axios/releases/tag/v1.7.2">2024-05-21</a></br><h2>Release notes:</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>fetch:</strong> enhance fetch API detection; (<a href="https://github.com/axios/axios/issues/6413" data-hovercard-type="pull_request" data-hovercard-url="/axios/axios/pull/6413/hovercard">#6413</a>) (<a href="https://github.com/axios/axios/commit/4f79aef81b7c4644328365bfc33acf0a9ef595bc">4f79aef</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><a target="_blank" rel="noopener noreferrer nofollow" href="https://avatars.githubusercontent.com/u/12586868?v=4&amp;s=18"><img src="https://avatars.githubusercontent.com/u/12586868?v=4&amp;s=18" alt="avatar" width="18" style="max-width: 100%;"></a> <a href="https://github.com/DigitalBrainJS" title="+3/-3 (#6413 )">Dmitriy Mozgovoy</a></li>
</ul>
      </li>
    </ul>
    from <a href="https://github.com/axios/axios/releases">axios GitHub release notes</a>
  </details>
</details>

---

> [!IMPORTANT]
>
> - Check the changes in this PR to ensure they won't cause issues with your project.
> - This PR was automatically created by Snyk using the credentials of a real user.
> - Max score is 1000. Note that the real score may have changed since the PR was raised.

---

**Note:** _You are seeing this because you or someone else with access to this repository has authorized Snyk to open upgrade PRs._

**For more information:** <img src="https://api.segment.io/v1/pixel/track?data=eyJ3cml0ZUtleSI6InJyWmxZcEdHY2RyTHZsb0lYd0dUcVg4WkFRTnNCOUEwIiwiYW5vbnltb3VzSWQiOiI4MjkwNjY4OS1lNWVmLTQyZjgtYmViMC02OTM5YzYzZWI3YWYiLCJldmVudCI6IlBSIHZpZXdlZCIsInByb3BlcnRpZXMiOnsicHJJZCI6IjgyOTA2Njg5LWU1ZWYtNDJmOC1iZWIwLTY5MzljNjNlYjdhZiJ9fQ==" width="0" height="0"/>

> - 🧐 [View latest project report](https://app.snyk.io/org/hybrid-cloud-experience-red-hat-insights/project/abdec9e9-4fb8-4f10-8528-02a6b74b6b33?utm_source&#x3D;github&amp;utm_medium&#x3D;referral&amp;page&#x3D;upgrade-pr)
> - 📜 [Customise PR templates](https://docs.snyk.io/scan-using-snyk/pull-requests/snyk-fix-pull-or-merge-requests/customize-pr-templates)
> - 🛠 [Adjust upgrade PR settings](https://app.snyk.io/org/hybrid-cloud-experience-red-hat-insights/project/abdec9e9-4fb8-4f10-8528-02a6b74b6b33/settings/integration?utm_source&#x3D;github&amp;utm_medium&#x3D;referral&amp;page&#x3D;upgrade-pr)
> - 🔕 [Ignore this dependency or unsubscribe from future upgrade PRs](https://app.snyk.io/org/hybrid-cloud-experience-red-hat-insights/project/abdec9e9-4fb8-4f10-8528-02a6b74b6b33/settings/integration?pkg&#x3D;axios&amp;utm_source&#x3D;github&amp;utm_medium&#x3D;referral&amp;page&#x3D;upgrade-pr#auto-dep-upgrades)

[//]: # 'snyk:metadata:{"customTemplate":{"variablesUsed":[],"fieldsUsed":[]},"dependencies":[{"name":"axios","from":"1.7.2","to":"1.7.7"}],"env":"prod","hasFixes":true,"isBreakingChange":false,"isMajorUpgrade":false,"issuesToFix":[{"exploit_maturity":"proof-of-concept","id":"SNYK-JS-AXIOS-7361793","issue_id":"SNYK-JS-AXIOS-7361793","priority_score":761,"priority_score_factors":[{"type":"exploit","label":"Proof of Concept","score":107},{"type":"fixability","label":true,"score":214},{"type":"cvssScore","label":"8.8","score":440},{"type":"scoreVersion","label":"v1","score":1}],"severity":"high","title":"Server-side Request Forgery (SSRF)"}],"prId":"82906689-e5ef-42f8-beb0-6939c63eb7af","prPublicId":"82906689-e5ef-42f8-beb0-6939c63eb7af","packageManager":"npm","priorityScoreList":[761],"projectPublicId":"abdec9e9-4fb8-4f10-8528-02a6b74b6b33","projectUrl":"https://app.snyk.io/org/hybrid-cloud-experience-red-hat-insights/project/abdec9e9-4fb8-4f10-8528-02a6b74b6b33?utm_source=github&utm_medium=referral&page=upgrade-pr","prType":"upgrade","templateFieldSources":{"branchName":"default","commitMessage":"default","description":"default","title":"default"},"templateVariants":["priorityScore"],"type":"auto","upgrade":["SNYK-JS-AXIOS-7361793"],"upgradeInfo":{"versionsDiff":5,"publishedDate":"2024-08-31T22:02:08.862Z"},"vulns":["SNYK-JS-AXIOS-7361793"]}'


---

## Discussion

### Comment by @codecov-commenter on September 28, 2024 at 09:21 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1197?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 64.01%. Comparing base [(`c1e56c0`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/c1e56c0efa157e88f5effef10b730a36e0607f89?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`a226a7b`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/a226a7bda88f523a1708ae5b000dfc25678007fa?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1197   +/-   ##
=======================================
  Coverage   64.01%   64.01%           
=======================================
  Files         124      124           
  Lines        3235     3235           
  Branches      831      831           
=======================================
  Hits         2071     2071           
  Misses       1164     1164           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1197/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1197/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `64.01% <ø> (?)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1197?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @Andrewgdewar on October 17, 2024 at 08:57 PM UTC

Can be closed. 
Changes added [here](https://github.com/RedHatInsights/patchman-ui/pull/1206).

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1197*
