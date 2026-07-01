---
type: pull_request
number: 1117
title: "HMS-5827: add community EPEL"
state: merged
author: rverdile
created: 2025-05-29T15:53:50Z
updated: 2025-07-02T17:52:30Z
closed: 2025-07-02T17:52:29Z
merged: 2025-07-02T17:52:29Z
base_branch: main
head_branch: community-repos
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1117
---

# Pull Request #1117: HMS-5827: add community EPEL

**Author**: @rverdile
**Created**: May 29, 2025 at 03:53 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `community-repos`

## Description

## Summary
- Adds a shared instance of the EPEL repositories, called "community". Loaded in from community_repos.json. Any org is able to list/fetch/search these repositories and snapshots of these repositories. This does not include support in templates.
- Changes the unique constraint on the repositories table from url to url_origin so that community EPEL repos and existing custom EPEL repos can both be listed/fetched/searched. Updates multiple places to support filtering by origin and URL instead of just URL.
- Migrates imported red hat repos to the correct origin of "red_hat".

## Testing steps
On main branch:
1. Add all 3 popular EPEL repositories with snapshots. You'll want to have these to test to make sure both custom and community EPEL repos are working.

On PR:
1. `make db-migrate-up`
2. `make repos-import`
3. `make run`
4. `./release/external-repos process-repos`
5. List the repositories. Confirm community (org_id -2) and custom (your org_id) EPEL repos are listed.
6. Refer to the JIRA card requirements for more things to test.
7. Check custom EPEL repos still work as before.



---

## Discussion

### Comment by @jlsherrill on May 29, 2025 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-5827

### Comment by @rverdile on June 06, 2025 at 07:02 PM UTC

@coderabbitai review

### Comment by @coderabbitai on June 06, 2025 at 07:07 PM UTC

<!-- This is an auto-generated comment: summarize by coderabbit.ai -->
<!-- walkthrough_start -->

## Walkthrough

This update introduces comprehensive support for a new "community" repository origin across the codebase. It adds constants and logic to recognize community-origin repositories, updates database migrations to handle uniqueness constraints for URLs and origins, and embeds a set of community repository definitions via a new JSON file. The data access layer, handlers, middleware, and background tasks are all revised to accommodate the new community organization and origin, including changes to method signatures, filtering, and validation logic. Extensive new and updated tests ensure correct handling of community repositories alongside existing Red Hat and external origins.

## Poem

> 🎉  
> Three origins now, the code does sing—  
> Red Hat, External, and Community’s spring!  
> With migrations neat and JSON bright,  
> New repos join the backend’s might.  
> Tests abound to keep things true,  
> The codebase welcomes something new!  
> 🌱  


<!-- walkthrough_end -->
<!-- internal state start -->


<!-- DwQgtGAEAqAWCWBnSTIEMB26CuAXA9mAOYCmGJATmriQCaQDG+Ats2bgFyQAOFk+AIwBWJBrngA3EsgEBPRvlqU0AgfFwA6NPEgQAfACgjoCEYDEZyAAUASpETZWaCrKNwSPbABsvkCiQBHbGlcSHFcLzpIACIACQBZAGUwAFYADgAmAHYuNFp6JlZsDHV5AFErMoAZaMg5SGxESj8pClp4SJQMXApFbAZpdHtYZyj4DERcTAH+ADMw2A8K6r8SbnxEdXwKeEGMNDZ6aMLmYtKNWvcm1fXNgh3B0cgvfDyo2d7mIdmOjwY0HxRE5nXCyAD6/luGiEiHwGA0MEWjBGGFI6B8+AA7shMPJtkRMPAAF7UeBwsL4Z5IXAAGkgsxIuAYsDpmHoTWczIW0j+LFOJVBkGWVRuGy2DxxGHouEW8D4iH23EQsHwuEQCNiWJIrTpDm46wooVm22510hYvuu2Q4zCJGY3C81EGqAwqq6DC82CU9BtMtQ2G4tCdGiMAEF7PAiCUfv9upAg1MBGhrohmXa0MjMGjUMw3nV5ACaDtUdyGiUgryJj1tHHyTKPOa7tsrWEVJ0Piw6iRxkQ6smorDIl48Vh65AAKo2EUEBqbEvYCheMHNojjBFwVB5ISNXBsONkNuDASq2AKIoC8qVEWN8UttmQEgAD2pPcYO87wtFTYlFIUT+p6BSl26AMAMiBNOy8CnF4UzkPgjTDgiABqzhko0PDOGqcylkwShJtcTwBgmUQznqBpGh0RavvUx4yvwOyrlg96TiK4yTCQeTYdukwTlOIYGKG+TqGS+w+LIdJQeRUQ2FEsTUF+t6PP4kDMJGVA0NKlIjFIOHbP4Yj0ZGNoSACwTYdE/i0GCIyaNEbhIjQkyvuxSr8HgXjjGMGAel6Hh5O0JbrNw3jOEKV4KZagyYuop4KmgSoqlhdZIrmNoCFQ3mnjOZAOMptGngw75fPewIXmFKw3pFyCzMUYgiQoFD6REsgIqGsxFug+rDq+Y62HSCYqP2KlqaScI4spFDFHSlXNkpHiSdsGmskBY6ILI7FfKgkyYXQ65Ik+RaiRF2zyLw+BgXOvYnPeqBPqIeB0Mt0pIjNP5PB57GaQoGA/BQXwyvJ+VnvypSQAAFPiYLwPQYAZAAlIBBRFeDjSUAA5MgkPQwjn6vXe+ocRQCIAGILvWfCOeIJaoEQ2DQ1E9RjsmsIMPATr0AAUgAkjY4b/G0qxBHKdrsIgT0oKEqCtPAPwkTZb6TB+4V44MTDdOMZkzpi2wANboDIJDGv4/EWJAADCfKi/Yji5i49m8qwVvJR4wU+ILwQ8U8qklLmXh7R4v08Sc7D5pAQheMqlA7G7Z0SPTOJUhgeszmO/iOhpkDc7zKDgWZCTJOk2T0ia/gMv43kkP79iiHCBSW3GqAdv9SJoHgKoUNNrTtJE02BB7VO9hm/hxyQmL0p8pao3wuHKKo6jaAiABylL+cJcIAvGSCFeBIl0gydBJgwOt0iaFZOWN6DKbwPIN6OSLB906pGPoxjgFAZD0Pg8ytwQxBkMoGcH6cB4HwQQIg6pSBkPIGeVA56aG0LoMAhgTBQA3NaSUOBf6kHIOpIE9dgFUDHg4JwLhQ4wJUGoeBOgX6v1MAYWgAgAD0qkiDqREuqdOIQOAGGiLwgwptQxcz/jg9m1sSF4nmMyLM0h7bDVYaNLArRNjknpurWWfBfRIgAAYMOYSNcQY0NCcMmFo+kvxICYmTA0QMoim6QC0RkAADBkFIjiAAsGQACcABGNx3iADM/i0imJnA45xriUgZAyAElIWQXFaKXpSE8zQpGokGCne+cIaC1j4C8VcDBsIsLYRfTElAPC5iUPxcwlhQwwQAew38Y4lAemcAozG8wnxSU/nwYKAgPIFPYMJGRBhl7kBDLwuyBgIBgCMLoopbTGFOJcY4yJ0T/GxJcRCA8bAwRrBIEuFc4wwSlXOLQLE8JEABC8NwiZ/CalCOwQAgcNtnASMzGkxA9kcz6PqqmHY3BQjcEoEbZgCdeg+EPnrfAQLil33kkRJ0koCiREwMgAMx1SEDQRFzUIAgSCMQTtWCYaA6rknvAixypYtGHIwKY2WXgvpowOpQUSaNi58ABL4V68h2A/kxIsZSY4WLfSmGxSAaNYC4FwEqDgjDGEMFoPCSyNkNCFDZfeMchVGohxpTncVllrLUDRti0I9YsBKEiJS3lLY7FaO5ccuEPxaawsQKY/lZTSzM3OmzDO3KJzji5gAEQav4RA6wpTIBnCrZA0U6K6rRic0EbKAahG0h4WgAZ+lOl4lUZAYMWI4gJqhEszBtgeABlgOEAw4Yk3wBiaKJY/RiwlvGfZjJ0kvTWBaWakbFigv2ZAixMVDKMXFYm2Qyb5ZpvjJm+A/waA5vVJAYm4xOXiRbaW9oste2VkmFQcYSU74eDtV278VoQmHlDrQXo+pXy4gfM+Jy85yxmTVnumsh77FgwXF4E+DFxhw1MfeckWif1AZeOQOkZr0D5ATuQMeZwKzfXfQe/gR77FgYUJ6ZglaMCIUROWjKiASUGKwKgaDJx1AaSqXcyAtSixtMaUiZpjoXXYU6YtKIJpen9IfOrcQwzRmVx4XwqZdC5k/LGos8JKyokxLiSkLZ+wdl7IOf+jADrzznADBoS51yRN2QEQ8/+uD2QvNIV/d5pBPkGDQXI2F9gGD/NCJumWLYxwntuIpV1rY+keHqIHUI/hS1x0bdrB9L5n3wCQ2+6sB7EBcE892iUYJEPBFS4uIDQEktnukKll9JAMtqaMrSk13ImKwaGPBss0XX1jTi7WdDhQ1D7FI9hDzmH7zUvU6Ypg2GJitXas0LcO4eqZKrPux+UGkR/PgAC6xxEd32JpXS3YDLfxaLRgamyaM6Umk5X4LEMaBXlu0ZhtWoqJilgcAIPdr5NuSulQluVCqlV0BVWq0xYN/hNDAGxHKwkpAIw1ffBc5dQjdZK6Y1Am3mUUFZQkgjfHxDKWhfUlRMaqAEx9ExMIRGSP1WyhMBcfkCCqVZqCGjAi6lsYyem0QrHGOWY44aLjPTsB9LncjoZNmoDLys+kykLOM7cY57xwZAnEAAG4FioE3PGagg1rjzLa7N+brnZiyFfLFybWF7wDS6FS6NF6/PjNE9Mow3AdZEHlY6yMjCTKekKzrj96oiD4BuXwozwinlmfEdhVJ1nZGB/mt0XoGawJVdHiO30sggXIdghDgA8upi2WnQRuuHWOR3ZktHHD5CCWQ0RwNwiIJsJQEWn29l1bFzAaoEQW21WG2uPZhxQfC9V2vj97Fp5BqCFPRBM90Rz8e6IsNi+Ix7wXi8gaWA1kXgcEgQ/Twj/sccRAYBx0T6eP5XaBGCJCVI8gZlK0ZuMmwvpfAUZiRs5K0i+iBISgklIxvVR4h1ExpivBFzigZZa5LPev+FXs8FflziaBduCvSC8JiFTjUjToxnTq2i0mxszo+F0vRJ4JzgMvxlaM/ObPVnXjBt6IliniVr3oXvYD0K+AALwxBb7Q5YBaJW425qxOoO6mTO71a65u74BaJTL4FViEG760CJbkEXgD6UHFi9i0Fj4ZAT42hMHW624/T24j4OoTau4aDu58FQAWyCFxjCGiHT6lCz6pQYAL5sCSE0F0Eb70GG6KEsF2425qEu7xZaG8EGbPx0LMGMJBj4C+Fz5sTuGe6Gb3I+6mZiK2xvIh42buA97+BOimJsAyiKCG4eZnJmGBpoD4Bcz2heCmJ7r9ChCWLIAa67BfT+RR5jxqztAv6+DpTTCnhGzVxSpjbpqBFYDKaVzWDDxoSIBt6lgZE1iQBdEWJWJNCmorx+BsidiLCPhWGBSlzwCPhRCxqnjr5gC1DFCRDgRUr4hcy0BJHUBpj0BaKsGRgaAyS0ByS4AD5aISRYD8pzoFRDTqBjHIATEbbnFECXF0A3GmHz6L6I5xEh5EEYKrx1HIiiA6xcAyzLYUBEAHGmJ9wAjIBnFOEaBiGlB3HTbtFmEjGL56qfGhLfGYnGH94InAkIDH6PjZK0BLZDE2ijHMyRgYB7ihB5Jc4zgoi0CdAZjVb4iEjP71Rv5ubNCdL9LqCITVJ0bwFH5Mb07IFM4dJoGcbdKYHi44HDJ86UgxEKkRhRjUAk7tIProGi5YHc6S4WIeoVLCa3JiYzIGA+F+GMKLBeAwo8EhG0aCLhGiLEJREB4ohB62b3xBnlK/7bqlgVikI1FrxkboZaJIBJ6YjkC0AyReb3CyB0q1RtYzjjA+QV6YBEFxkbyClP4KKQBBppHaJYz0BcyLy9aOiowIhWC9HwT9HrpjjRnyA/A05RB+rfELiMburKTdZEBQyHHvEPgxTNBjixz0ynH7GBqmImjokqE/FXE3F3GDYdSaphkSShComUj5mehKBomklYkUmD76xDCO7QwP5CkVlVnNG9lUSoh7S3RoFshLapho4X6nrih4gpmUDKhzZQlHyRq6ljT0x8DRr8DAX0CMywD+AeBlnEiMZsQ0CcSWa4DaxVyhqhA4VIg1TeRtackMAPEFmvhRwmg8keQAFAT+C4ALiKIcHNrBY1hoqZTSK0CwGykMbymIEsatLymoFmns4WkS64H8H856lCUM4iUNImibCGnMWhrsaqms7qk8Zc5SXSBm6TIW5OlKEukHrAokr6Xu5ene6PIRH+mvKBnSKxHaLpnJayB6FOpZG8FdBFizAWVTkn7ehDp0T8nR7JEqinFczdAsoAhJ54bggyQfDSCwCXmyCuUhLywkoDAAoJxaKuXfhpV9whBZb0BMUsUJzrBmUUxTF5UAWZkyTN4TBL6T5Fk0VEyIibiVYkWkqiSgxpyiJRqGyhoIAljjoYpWj8RxG1UZknReUMG+UWUYwqSMgRX2LEyMjMjEzbDjiZZTllEkTHneSnl+QVa1EiQbzcCtLJHNCQ6MQrnzCgjx6VXRXVULG9hgxaIABUd2qIWiNaHVyAclKUK1qRKlrWJOv4WVawh5laAK51vguql1VA11HKEGZe9MpYLEdI6w2S4ga6XQswTu3kbR9IG1p4eK2kZIRMMp9G6OV28lSpolKp4lGpulWpvOkA8QIN9Ahh9iUVh0sV8VYIiVw1qVrl32uA8xF2B0ZJ0VtJvcZ8h53A8AvxM1LgMkitCMYMn18UKt+VgFDV4aTQdIbVgGhuVVflMw01blHlkYc19hzp2RjCFtFlPBOhnN3NBp4NykFKdAiW61TIsAW1FAO1XgEtUtWSMtehNA8tDQi4b1Wt+tmZJtjU2wZtgASYRrVk3B2h3h0ipR2R2x0/pvV/olaQDfVUGoiJ11UnQp29AUBm02gu1W1J2zXZEMH2KO3+HN2WUeH2lGVd2MLRrBGeE2UmZ+nmbRFhnOXHoB2bXbW7XhWg2slGnKQlEC4VGgTQ2ATFmQnQqQm3XjCmJI2L4dRgzPW7mUg/VED/VxHdkgH5JTm+1fSxl1HDhmI06hweYrZYaOBkYPU1k9aXy0zsl6rznegnzkzRTXBvFwjv2vkgVf33woqpk5pVzbBKBSGhzCrJgDBSivgcViq1Rhm8VGBRVUpVDUjB3829BhqiCkZJHc24kf0dTOzwlQ6/04YBVflSgHXujHU70Qnw0x6sVO72KkmkGMSpUl6ojl6nYeBAEDziMYmSPjBlC0kxX5GT5rlOoaCqMYCbnUBUmoDpSvBKD01n6EXzBwVlXw4MzyBNqQDdn8T87JJ8DdWQl6mlLKS2l8W0206UhNIKUoHM1qkYE6XYHhDSVQBc0pGQSqUQ3P2JbdXgzcAYqyBeW5EOgIxz1B0L1h1iAR1y2aDR0HR0jF3X3V1q3rqm2mKZ1aIpPn3pOZN5E5PZ35N53S20my0x20hx2+DX2l2joV1SFVNuV11p0d05aAUtMOhTOD3D3aGeEOmW4mVO3RpgilpHwj23Jj0iIZz2UWaSLT2yJL2nG5M527WaLHrxDnQ6yt0uD21bN6yFEGTr3P2/gnm+SCOH7CMn0o3sN3XtZx7HojM9hUnhlHyMAHauhjydL0O9ooXAOOCiyJaFP3EYaZbiyH2lb4FN7hoBREADFjiPV+TgSUDynNGpyMgsX0g5m/Ir1qVRQeofODWE30OljVb/PtodxWavh2IeMNJQ1w0lhaLfaF0lMSu6iV030hKUhitdOSvFPStSF0hgtV0QsP1c7NEHRUB1SNqdpqWiMezjydhjjPPQtuyEMWMeDEZsB0gCB4Doi+DjASDnSMYaogv6xNCGgNKwv8PfPf1APOAgPsAmxwECUNIM2M5M2mlhPmmalRPake1xNe2r0eBJP2KNObPl23NHwPMZPZFtOB2XMFOS353dOlNF3x2VPgwFsTMN11NZu1Tgw5ufV5v3M12PNFvLrtMh2LidMSs9NlP9Ml0iPl21tgz1sPip2NsO1rP+EbPPPuF8H93eELtD1dtWhgiUw7Ne5hG2UT3+6WYxGyKUyWsRwbZ+EaAXP5Nuoss2IZx5lHXfNFlCPry+BaIlBaMhsou3zOtdDbQVyf5+gJnQAhC3v9s/vZbgeTDjiPskAFshIhB0ukXsIfmA1hnoCzAMhiA4huz/A+CQVZ0lt3uhyrwAGy4CzcsdRrFDBfsdCmIj6tS/MfsDHXP2Kwe4CUOTAPNWjUNh4bBAo9XIc8SCtwisgK6JhDTEZSAX0EmHBr5cc8e4B8fSACc9BCf0MiTyGA0JEZzySrRTCGh7x1rQH2N461QVmWbQTiAOjlqHhoqXT2KUyBpto0DQCOcACKwQLgiOoYpq+0QERF5GIQkn4KXgkK+psnXaqGqAMKIKfDBFZagxiu+EtrUwNA/s1wepB64e/QCjaBkppq+OPV9gaAcn+AqGuY+wpAoD94OUENHoHEGA6KaAQ2FMDkYXFiCAnQHmkHodvWnKyA/rb61I7A79l1uxrVtJVA9i37QGCJf7mgRgNNcpUbgTzGwTypcbWl4TYubNSbNmsl091pPjeY0bilF8yljLxpGlLNETlpLYmiqAW6swBlXhjpCzXb4I3xnpo9B749Bzk9jlHywep3eXfQMwdFr4lmfquqNotnc2nQA06XTjvnd4wXMKCBlI7QVedMyo41R48gQMODwX6mCIiQwnbmhHAx0Z7m8s18cc7Z79uYgd/ZW7gwg4+yJP8gwq/rR5AbFeY4uq9KtAJtxOWDaW/803QETA2qBkXqrMFZzRcFdHq0hJwqjroQb3ZSOSd+6D6mgN0FmDXkBZHgVxkANxdI6jAtv6k+8G79wh+B6eeIFPYYIxYVntWi/NmjcVw4wtQ1yVYtXa0OOI+QfDZEi0jA+n9UJovtAe5JsgYAuq0a2NrSgI9FvYY4ijr4lvNxqwSVhPZzdQTrv5QKX0mqSf47wGCJj5pGk12iiQFXJAVgB3DA+VjDqbOYEZ5Rv4TQJ9C6BaPllI+f8kDXGj8OG8rC8ESorIO81+BrFvsk8kuqEpc6Up8g7qWATmHESjfql+bQza5Kj7JNqPQ0rB/SoQLSqMv4Y3mDdQJ4Oak+NKjfx6CDlkgaAgwdKnXfq1XWcUBKKqETCRBXKttMvNmTQ5kofaCHL6F82F5IhUqNfNGnI0gBj8Ya9AW3poxEbHYyALDSgAQ055PwQy7/MmmAKcJ/9FAS1GjsNiAjkVL4GbWAb+FLiRADIXzZRJWnmDKhRgpVIgXWxX4YCXefeWQAjHUgCoFgRZOBvIHQFv9n+d5BMPVHoEKF5BToQbnqjIAogBgL9RYFCx1aPp9+fA9XjNk15ThDc3YcmKWCQEmh0BD5cskfl7jgIlGGaB0HOkRST5bGJQBijO3rrLVwIaAbMPg3nSvh3UFgqvj9Cv56o6OGYGwSaCQHcp+ISeVoJymYZ6lAq6TGvpYnLjSBEW8EIgCqCdYDlY+EndHpQH/xEBymp/VEOLBUGkZWQQkEsOvwpzWx9Q0fZouOhT7qYie+HUvKgNz4lgbBE/O3jgL8brcL4l3EJrtxFwSVE2POPArE1Wpg102C2dmMkxbZgxYKP3cAbMy8AIwzY+nEgIO2KbDtY68GVyhrU6h60fuGtfuFrV1qq03KhtMaCQAbZm1AAKAQjEsQJQn8H6nqBk9ukJWEMDE09qLCmWywjSKsO8jgwNh1TLYdkSyY7DS+XgHWHsL34HDy2n1d3H9A0CBoAAQnSFOF8CAA2gAF07hBba4SEC1okiyRVw6QEbWeGQASRtTSAO8P9b09BgPwvnqYNr4Aj+C8w5egkxgHEQRCzbSEesPSawicirTfAZ/2/7bAVOYMAfFWUGYyi6AuI8upiOYDYi8RMorIlMAuE3tKIlAPUWgFZBb0d4fmdaum2QAkjJ2GI7YFqNxEvDTErIz4fAMGBWDOhNoD/j2EBEpsFht3IUSsNFEFJxR3KSUfCIRgMhA65A9cocN6bHC+mS5MdtgDpjsgZWlFY6okBGCWR8qj/OtFrU3T7J1QBbcAUOVqHeDJmLI1YF0XoA0COUdA0AgUmfbm8hBFBV/nyM9q80Gmaw6ETbScLbCEYvvKfl4H94JUg+yoEPusHjEF1lW7sFDjSOqYUjJgWtHWsrXuEFVHhTVZ0dWOqwl82Ww1YGBQWHqdjU2IIxJrAIhGhi+xBVCMdKOb5SA2+WBfKjOMrZSt+mNo4kbWzaouiPhRCNYCJUGDD96guqE/otm+iE0501/JstcBbECMOxQIs8YGMYHCirxUIiUQOLhHSjhxokMcYHyL6wBDGqnLtK+KVa9MFa/cA0eSKKqTA6QjoPFAMwzGk102FhDLqM3BhrjLhS4ukU8J3HvCPml/KCYwBgnlpDqrYhCf6IFHe0UJwYnsWKJvGAU7x2TSADUJIC51CmFbMiSO32LpjVWo7FUaGnpHl1FxblFCB5AUFwgtxxtB9Ob1oDjgA0gaRIIyK/EysEYP43cZ8LTAQV2UM6FwfOg8DD8bQSA+8DEIRL34PBmMdTGADoaswYwlYvgGwD8HWYPuKzYyjbhdL2o/umzO5nu1CJ0ZjM+zZ5Me2OZOUPee4u5stVTY+8XquEoWiLWD5J90qU5Z3ogS0QdtSxmE7yq800AA1Kp//LejlXo6KtExJVHelohMmbiaJuAIDMgHrHhSaW8OCqlVxeobYJpBtHiU1VGmtVZ2VcEvrTyWzPMlqWiM2JyjoCQCyuRgl2L0DjhBVf27JZtJdRKAMBrQ8wV0KsCNYqS2KDA+wNT1lji9J80GagDQHtBYRsos3EjKWA8GfSnc1UCeOawqm08ysUfQ0MgChkj4cQx+Gcu43pbXc+A7QJqNDI9h0gYeJYIGNS0aoEQgIbVECPL0Jb4Y7M77LAMaHrRLY+hvYS6lKhZSzgLObjFSBVLObWh0MvZYTKt2pyRsxhm3RUjGyUqhM9uCbQ7rMNPH/8I+IouSaGLbbtTNhnUocTVMFoB96pk4xqSRI0nDSq2fTfwIrSom0jNaHEtafVQ2nWTmRChb7tU3UJOpEA2U7ZkszXZfcN2mUpwh7N3ZWUAe+U30sD2KkC4Z6toHiA4HUABwS4P3CCZGHLENJ16gVVYlnnCz4lGZCWIwN4mxRgcQgKnfKtAHwCJBFQyoVUKJz6a4VKQ1Wc9j9mZbKQWpUFVoI5gPA7ANgXQN1l4FCxXRq+aFYUuSEAGVy1Q3CSABPKgBxU/gQ0SmLgNxxM8+iR4sqCPMSj0htAkQRCk6zHB+VfgdcYoJLANjQFdI65EnPQA8iqRBBO8jeRDXOhapLIzwbIRINxxBgCw7uGXNS2qbqC0C9DPhniiF67QDAE83QHRldAWDG50crCNvyGCLz2yy80GKvLdC7zN5wVU8BmGQV4ID56AY8DpDHAXy3i94HYr2jfYK435+AD+Z2i/mflhOT7SkP/KfC2T+IGQVqJ7zHjnsS+WiLjgbJSpGz1gbqKxMIQw6QKVJBAyMmOApoVcqa7WbRDhL1njiCJU47yiXyIoNhE53xSAIGlDBJ4ysu/NwRmDGoH8ppdIfaZDInGngzm4sZmBS2IWBdVFRrDnl/MMlPDIAIwODHCDAB2gAUfPBycw11TEkoKrvcWCLzr52D6oyo8/IgTGr4gysgvOeWCJJp+pmSQEaWGIqRAfNRiHHJxU1RcZJILB4nI6OASySQFCank07t43KRvAq4Oc4uNvDQwPo6SPUFDkwCSFog2hSfDoWXTgog5uw8oCuWvJJk/FVu3pUYfTUllIFpZ13WWVMNZqRNFZiE5WcQRDHgxY5C6T6h1PXKJA6YNABGFwrMWKKwYjdRgi7LcpuzIwgckICu2WYD0/Z3AUFLlO9IFTfckRByiexOYkCBcfMrdAzw8CJAvOIoIQMtO+hnUVECZBIkGD8wFsfOpQ86W1nqCdI2QJNT4pZkHlPlA0TnLyKWABXei4Z4WDJChRbJtkEInZJEFiqZlGiH50zTMqcudRtINAtZOVsBAUIQwESQaYxRiSIkD5AMVcBPv5HPIYlUqdxJjKgA+h9NA8JND/htmZWIlA0bK9cr8WuLUAB8sq3RgKoRJcrep7XXDqDKRCRB2oYcQFXilwokA8BlKntKNNNUuBqVKcsaCuUYL2pUx0MUacEsfzoU2sz5clfhRQ4qKvh8gegday4oh5/p8A8VeSth5YBh+zEByfflWhtoyucPROfZKDTEC1u4s0ZYMW26xthcbOGZU92TYncnK+pYSigXlDISTSWa7Su31zUxobSlSq5eu3SnrNbl5yyYPcr2ZPLDmU9UqQJFYXCKOFXC25SpzEKgg1OPmdei3KpQ2BblmyuOdXOthxyhFDc/YQnAMWJyboWEQcg5kum2DXVYS4NP4uUZyrVVg+MrP9hRlhBwsNgKwPEFWC0z78SvH1IMDhAW8r1qC3SK73SHdL+AFgqIasFpisZ0mVcc9tBhSUM95IIqvPleuqgmgq+76gck4WtVYAIpD6EjO/SfVoCX1YMEXuQETg6w+GMG4QaKFEH4pnAvJJ+d6vxCVlg0U3CCEEvljgb+hkGnydSz/WhQ4Np8hzEhpF4WDL18QOkDUO1x6QOWPo0mW717BVl7wbGp1Ahv9RBpANC48lmereltVMYoEBcADPlhjgeNCnQYGzy5DwsxAUQdGcqrlDexs+SICRcz3caFLYNfAzAACCvxwLECbM4Rc0RsEnjhlqa4jkE0ZoyzJh2ax7npQ5r8ieaKstCWDGQBrKp1WykgDspCCTrmAg66fMOqIEHL5mNy0FDuwuXezzc9a3wk7UQX15g5uzQHoVL9wBlXlXakEqd3hWn4A4oa0ajsCogZhyQKKt1WisNxI97OUnJXB4HZGoz9kA1K+n0tVDzSv5G62lZACp6iAaea6Zhr8pFAAB1WIGUBsBlBhJrcJoDYpATag+i8DclVxlHDhZWtu6vNGOSDTmrSSHKykgjH9buihgfoAWMdvJCsqD1KqpPtuV6kJ875C4QWaWGZnQEBNxTXOfwXIYeZf+fUxQMw0k3JyHMijMgDMC8l6xbtL7M8q9ouJHqO6QbccveXrKNkNtwmKAKDu0RgrDw5c+KKPJ8xukYUqHHqsw1JVAq4ydSi1b93g0uo6VCJCcqYmR3m8+Vh697ZST9FE7Z6ZNMnUALVC3MLUOIwtjQFDBSgkOEO/6WTKRDRA8URsEgLUHvDRB2uRYWoAVq+HPcUdx6C8vzuvIccxyE5Ssg2XW3NkjArjCweRQbYuK2QWfP9CKmKXHy9STwSHhHj3xxEyiQQi+GcBBTv00hq0JgPHmRUhKd15IYTdnzEkCM32rHApVHqHlYAqyNoA7GnG1CEE9dEmtRazorJ9aRhnmwtRmt83lr9ukldmrbt1JlKa1wvMZeePUpiV420whWVaRe7IA3uKU65Q2v8IrLCsQcj3CHJ9KHtw55WkqWD27V7jveiQY1bQFbLGr9BPYRRfwvD5BVWpXladTQFnUD6hFJfJoHQCWzntz+1wH9VEpT0VlodNKisRmAK2WKzw9nA5sNtNTJhj4EYVSP+qc3L7RWc+ugIvrZk2BYgTUkvtr1bAIlGQJNC/S6tT0v8KeSOLRHPuYrcAuOlA+gO8yYEzhaesuEbtHkP2EsFdhuRmfqWc1ubE5+BnsBL1yhjZAYT/AYXL2r5dKr4/gaKfPoZiGwUuHELkLuxTV00vNW3HzZMr80Vqq9R3OYV2NC1LLwt5dLfdFoRh/6F9KFH/UQH2WHLO6G7AfZlpbXZbDKuWl0lMEQA6woYP0fAK2pK3tqQeFWqfVVoLU1bnotrcPahXmBPa097W+/mUhxznrKQuG+QKft60Y8ooMUM3QYZ1hRVjQ2wtA2iVybmrf+lPX6a4LEjMN+ic6Jw9utgNJruZnhscPTtfpx87VP3K1S6lGnej6t2fN/TeowboJ6lH8d4BPGDUlgpBaOjcn8UVWUl1NV0/ADdKiBnblykNCOPHu+Y6N0dJu4xjiBw4Is/tZnLED2G4Qg7GC0R5hj0Z8zJGZg3O46rzre2u9BV9m2RhjSGPNGFVtxNo11iXKjGwRA4fpPNHRVZGSVgK3I2Slq32IltK2soHjtv4zgJVIR0bW5STk370Ocx+xL/0WP4hEAwdKFX53sCXGP9HQZwO/XdEbHhjWxto9Aq0TqNzeV203dVFRIMjP19OkDP7PY1tIw+5xuAYbpwO71hGFG9PQmVx3W6mgOSr9eIId3xSndUoF3RgQgJ1ooCbo/DkWF92hkC1d2MQO/VMZvB0MSK5w5fvlLuGH51za4H1vI71Dew+xskkiePVDKxZfB0vYIauwt65Zbe2ZZLhr2RyzuFShvemp1ORS02TLMtZpWmUBb2a1ZV7jLHe51rfZfexhCEeMPGgtDmgIrfu1Dlj6ipE+yOWexQ6cKi51INLR6a9PjAfTQ+tfXxk0F8M7tXx71t6lEQ/qBSUpkUuY3fy7A+AKpjHVXE935AFAB82HiqUixlHDDUGzrqgGgP196oJUfYUMDGohHJ86gfrWNoL0N8kcVi31td3QwEApgvgJgFgsIVkAiAcaABnYpYokRyjIqv8QArfUEaXDrYQw+Wf/bEHECsColQAvoChSYDFZcc90Dk2ex+jIinYKksM6ZJ31HZ+KECmcC/ac+P8gzXWItAil0MHgqIEuZBzyQuzMm4NHnp7OEm2sia4NLptgBVxxgSgeYqXD16R4qWXXPItmgfMKb5S5SmDNxFoU0z4IcYZC35CT0bx2z5RjjiKpYXVYqdzQDhbopoCpVPOhh1Lc1JVl39WzaZwi/AsFAuGJIeREWNFUQrQJOUJNZzWcX2GMWdYwdQVXRyLMjHw2/FLU+MJ24V75Zhp6JpJJC2LK1Zyy6LeXQkthGy5chmPqiIYtv6Dl4MIsRHA0D6WTDdISyyWK1lgWRIahhwnlv8KxmTDvpy5T7NWYxnPFxiMw4GaB7BmXlk+4MlNVUkFsILPmEvgoWBkOgnQERvar3z4Z2Hrsjh7COuYyMynViiwURtDFJAlg/U0VglTtpZ7ErykntdkfYe61o9Vehgk7HsQ52OqVIxxRYPQHMHiCxwq4KQIwR0krlOuu6RPE0flVbkBdSOUEpCEdCR5HGP6++imnAo6wXp8JbHZOTly44lzF2GsCTR6smr+rdIFUxiYxb3hGZ+NWS2qbOMjZJgoDTcF1D75jd2I3QOE1gCV28gsF99FJWUNLBw6K40igKVGvaPWxbsRVOMKedCAfXRFX1ha5nr4HTp+S7i1rm7C0SOhJgYIAraljTGatqZAy9wQtKwB+r0wNoYhjxWL2KWxlRa5S3af82VrAteBfNR8mi6lr7urenNYFplyNGqqR0egZZi0TxXOE2wjQJFZ+7RXEz+1Uhj5bSmMI4cokLZLcFtxOAI0zte0ItECuj7grZW0K6GfeX5L7EC0Q0J32StfKUzZJxPcCqOguRsOHUF4HkDvRARZOefAQV0MovR5MK0wVI8dN4UbBiTdFqII0BJqaKk8v66kC4HFiAWtEVQMxrLtoAPjmqsVwGqdN4qVkAGDZmG0awwRtVmGOtqCIcAfXv1ON8sNqkIuIPfaKABsNXd9eUMXsr27nEgMvAjtpIKAi8Ng2mVPSasCKf1mnaRiWqcnfAJSseA11nasneShA/EsTbSSkNRZEbMmxaYmW6mpl1N0Q3MsgD03rMjNwUY+vnsiGZhVpLC74zdO+Wpbk/GW42Hls1d6SztQTjFL9PD7itQV0rc8qOZa2QSB2UJNLYBCy2NgGgGhlpzEBqC1iZuyHlfdDo+Ydb69MW58xNunUSyvgB7TzSW7skTaRXDfs1Awg7xRWC3Uq8zyJWZ2cZjxKxLTwzl0Ra5yLe6VwDCG9NAbxVviMvaxAHk9UBD9A8OgzCwOMIyNHli0Q2wYOaHjJlJKd0QLMnsbzuygxyaKVcm+7pp1q5Ug1NT2AmM9q7nPeEOV7t70lITD3ty1v2lwJ98dB/f+6321b99jtaD3CvUke1ws1tD8HIDVRcHAIUGM0Rq7+CBNNm0C1JrYx2g8UrF/++hkfNZo+zOKB8MwHccJwOYiQJPIvA/oeBaxdBJPjo+hCwgMAtQLa+MEIFfysUCB1fZISKKfLixXQdoDdOwAbw/U1qQI3RDF4fF+gaC5AFjW02sgKAzIOOWIBJz9Rg78AbXvVCUS7xJ8AAcSsCdPIAPhk9YboTj7ioKGAVoJLDVD7J5gB6FeFgFDBWAhEhQDmS086DW1JpitAoj0CKK6hGQSjU673eWLvA1s9JKkLhsql5BFcV5mbRWPv0iNxY0tOMKS1kEe3XerlLJpxgoAbPJoBkFUAygTgB2g7e6eQAhfLgzB7w4DHTdzWI7W2arbjugEFRXVq0bUcMpEME9CfhPJ8MXbkJtG6CN71owMlSWzA0VaKL2S6KahHbyBR2Y7aB1YHkCWxwU7EOfAJ3C8Oc9xvoYz4/CSiyhTFDFitcWKXGGodo7Q1ZUPEMOv3Sa/bS/Il9oqFCD2cb8XYFA6KiDyQuDsUGgNwF7hGtqKsJkcOvI6Ak4q4ZjsFbbc5ghOwnxQW2NwJ8B3plI8O8nSFAzgZ6QEkgbNBwuhfExPgK6SIAkiMCJDlAPgFIYoA6NdGE4TQLwLME3xZIawUQNgKkiQD/QhcTL/6UjWslguQKAEBFz8fG3ylpnpYSFB/HsB4u7QpN2R95tnu2mHuNN6vfwUEhBVupU+F56H1fWlPEsodN6hiy0SsS23B10MLU9VdSF23gaZp607hBIQ03w8mVu286fcAiAAAaRIDyBr60Zw+3bx0en3C8MTpZlABrfRvvejTJzG2IvCuUdl+AaiYrXMt2yToy46afYXSd1vnZShDR6u+0eNhLlW71i3W+efCDXnyt1nB86bdHP7EfhGSL2E+rXsvKwH4O1mXnY25H3Wj6Jy+83d0ZWLOuIoiaB1vaWm7mIRRYgDBhAf8U5dMD9kQg+Avdhntt57+6XewfT0a7i8Bu48JvugqHCvd3wE+qKLyPRYBGOS9oCUuW+pExMW5NnZTM2PP7osJ88yf3uYPR99+3B9d50fV2OW908u80ZPv4Pp6WJ3CGsrmG7KlhsK8MnDDVZUXYTsx+iVk8vuYQcIUxK4q7B4DnexQB/q5ao9y2zjZjpQBY8eBUgeIRFJFho/SFZuGkzRT8GDFt5zcrAJKHWP4MGABeXqvAJAB4EoatdHwCMdpxfDSB0hPE4sbxI4lZBuxjOpAJRmOEfBpAAAbGCCK9uJL4dTmgA0+NhChOX6Q3lPIHhOmKv5XROkFU7x5UFh3iiMdxgCad3YuvFXhAFV6ZaWKmI4EDMxnG6e9PHuPhuoC8Chaox6AzRS6kfAi/WmIaIGgPfCCRxTe+n878aAo0Te3TSi3gcQGAHopsSewv2np1YBzWze+kFUkFNQANfmIzkewN0Ik6YgYAeUKxQqCAN5AV4TQ9AsAz60HQCLKCpIApAblhALgZgeg0V4nOSLnOpgxbnHnI4mEqWDTVaum9HjMfdjn3an8z7Smg9KeRxq73rJG6SclhDPl8KgG8j9RI+DcYCCY3R2bfzcgSB1sDAdY687AuvYIZL7Sm59Du8AIkMEJyFgDHXssRAad2CB8PeuJbPhRzxsAvuacr7qtx5dp4jmnttbuDvmpfZoV/2rE4DmcMKxhoUmP2rD0+jdR/o4UvWX1Rd0Iq5ZXV2H5GPK6g4gj6kdE2RDcYBRvZ9s1Be0g7GxD2Pf2gHmWT1Z59nOd2RIS1Am2YQDUkNMxXoairK6EcMVeBH09GfJf8Zo/S38j8tyzYdNiH+CpMKAVgCb0yTwRSy0Pwb4HYmyh2Zs8pjW1cngwD0ZXhtjuPqYpMa/9DdSeW1NnviKmTE3VOqxvqt/ug7f+KTuIk+k/j71HwBzQu8sKeD7dFSgIwiCghQKAOj7zOr7DkhXH72vmw2kk+ViklssemPnHJ2AZgzdvtLmLMHUbUgYrntb1VoisDQpN/I6uSGycoCI5F9WDjsswwl81Vs16ZuvZg0j1AHIL25hq79FupUOVQDuSzkApmkjMM99E16m20DjKI/S02prg0GtiukK6opwDxD/yKpvoxYCI4mcYEUOwHhzXYhMNwaUgcFDSiFWvYBo70ynaJH6lgZzLH5NiSfgQaCObJsa7xKkrqibKGrEuaoQWq2PsjHOcJBmBQWyTj8aNw+Fv9L+qZYIGqo+glOTZl6Qhpj6s2VbjqQmmSlqJQlq69gX76mOgUdwSOe9gr5KEq/hQDr+7/qxjb+VUF5b+meUvo4WGWvm8pxEDcv2C/aTBPYHOAjgT2jOB3lFhassAxoWRQOkJBKpi89iPoxcAEjOphkBokJjraIdwp65FgJootyhscYNRp4aKUHczdG4OsorzAqzopIUCmDrtoVWgJlGaEG69Aw6vq40uuLpBxoorgU+3QNtYNG8VFShz64vhIEMoiSJiAm0rQI16jO7rG1g824OhxxxKTXh5j6MfQeyARK8rAkElYSQQCAyM6NAgKFc1ZvYg9BvbvMFCKoJJqoTGMNpEAmQDzihzPML5tohccb/q7DOAnfAdY3B/gRQD5UuwcyALaMUA/5V47bk8F3BLwaehvBsAFLosQkvqcS/BH/gCG0BQIbICiBCuqS5sBVjKWDnszJgOb2Ck+D3bcmY8NawqBJDGoEbc6PpTYVui9kaYjIT6mo6KeNgZuyIuvdJp5327gSGaH+SAavaVEn8j8YI8Jhn9AVkKcFnKe0Jiv1S4WfqAv7accIHEbYBCRuxwJkQoWIBoGKQi/bysA3LtRrGvkEtJVU+pNyhLUoEuCTEWvgHdJhsk2lBAwmi4Ouhm6ZABWDBAPfmIBlAvVte4WKpYFKG5k5RpdSyA0Lt/J0kxtq2LBKZdDEHg+F9JQDZc4ZF8pbes0jej+CaFpQpshnQoqwj8IAQVSk0gdJ2b6+woaODlGdFvVBnQF0PpTcOvMgI6p+/AVUIiOgnL3bHym4H2QT2vBiW4CGZbszamBRfkvYr2guNdhM2epvaaVu5gbva1qVgTbiUhGzEPq0hbgZr4MhngV1xBwgbiT7dhRAsEGJmYQSuazBQDDEFbqq+CQGJBUnloyI8p3sjwuw01lFCNW3KGAD8hUQHM5cw84p54bCFMqhTCAxwcpAoe75hgTPMe+FNqxS4oVUFaIcwWYiSBLFhvrysTodC7qgFobgChezoWYwpBx6M8wSWZsB5DsAZQBgBmhJAH+EwqHThqiUgJkJ3LoQjQVxIPCDskvgHWF7i4DseVklhHaMOEbIDwcxEFe4rkF4dQG+gTSj4E+ShQnvwdOjrEiLGWCiA6zeAesHrZ9MSNHjS+AvtKHaJhPVCxxm2s2iS76kd4bwJfy0YlyAcKCoVoyhBmBuEHyMsQUAz1i6fGg6DwPAMtKX0VKK+ExBDLuGGxhWSvSb74fwNVqS8p2LebHor4UDK8+D0PQ7QUj1ig6fE5KB75URkwILI9yfcukL2hHTqmHFCvtMfxAQjEVCjY8R+HiESyBIZmpU2W9u3rqW9YcRzaBtYYKAU2mFvXokAMuAdj+6HrGvRBGo4Chz5K9+JcHkhK/s7pr+BWro4Bm/YUeyDhlWuNhm2wrrraIAbnJagkAtSOZyTkJfGA4pWX0KHqhk0JHUqshBVBqE5mz2oGgVB5VoAGe0P5jzSQAbiO4gsmcJANGAUS1EqLBoackEAbwJJOyotGRxseoTWWHLyo/MQkWOYt4uZCvCXmUMhmCzR5XtTILRSICtEPg60ZeznWwgh9p26TJk2IYEfAcPYlgY6qWFhRaann4Y+UUUo4xRearXoFqBgUpRGB0kiYEthxIS2DthUjp2GukJUbYEeWxoLv5BmGtgf5DhJkdIgn+qSjPIVwPgBWRKCCZLGADAWjMUFUoNlsaBf+pGn+7dSY0dg4cBntEzzsAQIG7Ykx4wV0G3RDnIYZDRx5m1qtW7PKcSXa20R9pxECfIFS9oqAJQFQShOGdGwgM4W/rHYGwKhTDRrhiLFpgCJj8TFmSdkiHlGDEj0K/gnVs0AqK1wC4bNo5rJNF42CcLNGOIkADiKcQV7iyZbqiUsRhog0QCdIYAYCtCzExQwDYISWFwEjjt23qhTH7I6cPVB92PASn4+CONpxYRx3MfVAWy/cL3Dx+OISTbSOClhWFSy+ftWFwxyjmDH6BGgZaYYEFfrDEL2RcU/DIxAKrdjyo+wmCDGIuAGCA8+LTiL6GILgQ8p7+2MZ2rWGTIURbHO6AZ4wnRisaWCi85KlqwFInjlSgoiToFUCIouAIO79eHcRMDwRIoT0RlWLMXzHxgHRILFNmI0VOSHWEsZSS+KnQhKrX0nUECi8MYsSozqYREpqwJ83ugVwYIWiMWLzQswKYgNEmUBIJK018cc5LhZBEnwMqnoaOgXxMrD1x4CTSHvGRSQsbuoPR+Tk9HG6F1vvopWZ/qUYX+TWpkaQJ+NtsFiubGPUD1G/cq7wdKo6PS4mgjJFdj3qmZlnh3ma5prHZ+IyvwZ5xQMUSHVxNemMj72ktqpD5AkQJkIkAUthyEDAE5BLgtQXcW2oDhmtoyHHoUEUbADASou1HQwpGvwnR+5IGnIaMN8ZDRlmw8cIyI6/UUiBZWwaBxyiklOJvH/+7HIRRdB/gI4Ic8itLgI2uWMnr528ejCyq9Gq5CtHIkj0d0bixhxoKpgwUQlgDUyHsWt7RAEhFWSxgfsf/JgA3iNEC30SIJ3ijhB0XfoGh/6tYlJhtEVNL2JCjI4nVSziR4kYEL4a4meJiCbrGqmL0afHvSLFHeibwT6AZBBJ2QiElhJwaBElugUSXIQFyj/HRCNy1sSlBKJfCU8BZKxznRyxA0ANAC3es0f4jLo2wGoAR8X3tKCUA3sG4LUslsjjb1AHgvdhcOd/p7yOJ9xmRilEjIFDpTSICOdDz6zCZ7y0kQ9pEDU0k9jnG5+lYfnHNhVcaDEc0cUWvbSSHdhXroesidsDyJLKrQBgw5ABcmiAKoBoD0xVyaX4MAbksyCmGYKZQAQpS7jwnKJowIIlyJhWCYliJvBAiBjiPlNgL0CY6irJFRktl8Yq+tDIv7iJWnlVFSJuMSabPxoLj2r70wjJfExBhSRwytSf4SJASWYnorxAQZ0JdQEgC6I4wxBMoL0DYAeQnaH8RbWNrAUAOsH3aihj4bTzPhr4avifk1YDeF6R/MXrDfhZjC1QL8qDNAqphXgrkmkpSYZymAcsEDMAKEf4bCkfOVcCylAsvoRog7oWALkE80WoUdEkOIcK1IyR1Lt6p+o/zohFUof4WoJZ26GD55wUK3uF6kAqCUbZ0uicpJEjUg8EBDeR5IPQJYGJvIgGWRykc75Fg2NKqCDI+NIWAECJYCqBYhtmhNCMgDwFIA80SaRKkMw/YJ/B3wLoNHjDS2cTn7qBEUeXrAxqltj7VurFsyn6ME7BO76kj3B+7spcIKakz+xKfaErsfoiRG2IuvhX7YQRqT/a4A1qXx6N+G5nrCfU9ltZZv6BlnSBgg5dDBEaAUKsEACe9dEu5TpEqSuwQOrYhMT5e8SdHh2pR9G+HrYIGGOkYAnKX6IS6YpAUDDcV7N74FsfvqRxQcDKndo6JFvvb5Dp9YkpFQ436V1EiRr9iuExOQabtRwSr7JEHCMUGf26W+ALM+mlYXCT4TEpWqbS5UhJytOnkpdIZIk4xNUceh/hAEdC5cpqaFYh4oeAsb4KRVRPwBw0Fvv2lAMtvk9SaRs5FfQQJH1KP5/UVcM5qlO9iGoJdYwdAMCmI2IWPakAHVo+DZURoCaDewaHg6JAyd6B5BRg7JP9FnJyUZ2lsJzyXgQro74d2IDpYmYbijpEqQxlARJPsRloAgEaRn2oFGX3TL+RKarFkZBVGjYv6wQX2Ea+lKTRn9x5aE0qjhk6T5n2o6NgmaG2v6b+AJ8JiogRiR9iEBnephBpohIs57PlGJYXHKLqjypMD4A/BIQAVmJQDVLIDeQoIZxylZL+hVneQHwTKAp43ACiCjupdiJC/+hKgAGsxqbHUEJ2r6sQ66hj8LkCaSLVGgx7R+MUklUccDtkF9MGYA6me+OkQB4v+5oOALkRjglBhC4dXojIaqtADhb1cemTax8yULPpqZcqckw70pfAB8wH6yEr/HuwwsLGlfymoYdEYB9Yokg8OuSNwFu6YjtASu6X0VnzNSpYcK6wSIQMmoea09oDGEhhfq2F1h4MQzZgypmYaa+GmgVdiIxaUWhjv07CvzK8hB2ENlQKD7MKJqOhgAYAoIfGJ/DfweAIQBYxwMOwBcAhCA/a88CgJgwUI6gFoDUISCCTlvwwMOoATkHsr0SjwdAGjbGcoQDQik5tAEV6OIAgJkBuIAgI4hoAGQAwBZAKQLMACAniFoKeIJAGkACAKQAIAMA7iGgDK5aQN4gZADILQDeIKQIgjE5pOZRgtx0MHzk7aAuVZD5u+gEAA= -->

<!-- internal state end -->
<!-- finishing_touch_checkbox_start -->

<details open="true">
<summary>✨ Finishing Touches</summary>

- [ ] <!-- {"checkboxId": "7962f53c-55bc-4827-bfbf-6a18da830691"} --> 📝 Generate Docstrings

</details>

<!-- finishing_touch_checkbox_end -->
<!-- tips_start -->

---



<details>
<summary>🪧 Tips</summary>

### Chat

There are 3 ways to chat with [CodeRabbit](https://coderabbit.ai?utm_source=oss&utm_medium=github&utm_campaign=content-services/content-sources-backend&utm_content=1117):

- Review comments: Directly reply to a review comment made by CodeRabbit. Example:
  - `I pushed a fix in commit <commit_id>, please review it.`
  - `Explain this complex logic.`
  - `Open a follow-up GitHub issue for this discussion.`
- Files and specific lines of code (under the "Files changed" tab): Tag `@coderabbitai` in a new review comment at the desired location with your query. Examples:
  - `@coderabbitai explain this code block.`
  -	`@coderabbitai modularize this function.`
- PR comments: Tag `@coderabbitai` in a new PR comment to ask questions about the PR branch. For the best results, please provide a very specific query, as very limited context is provided in this mode. Examples:
  - `@coderabbitai gather interesting stats about this repository and render them as a table. Additionally, render a pie chart showing the language distribution in the codebase.`
  - `@coderabbitai read src/utils.ts and explain its main purpose.`
  - `@coderabbitai read the files in the src/scheduler package and generate a class diagram using mermaid and a README in the markdown format.`
  - `@coderabbitai help me debug CodeRabbit configuration file.`

### Support

Need help? Create a ticket on our [support page](https://www.coderabbit.ai/contact-us/support) for assistance with any issues or questions.

Note: Be mindful of the bot's finite context window. It's strongly recommended to break down tasks such as reading entire modules into smaller chunks. For a focused discussion, use review comments to chat about specific files and their changes, instead of using the PR comments.

### CodeRabbit Commands (Invoked using PR comments)

- `@coderabbitai pause` to pause the reviews on a PR.
- `@coderabbitai resume` to resume the paused reviews.
- `@coderabbitai review` to trigger an incremental review. This is useful when automatic reviews are disabled for the repository.
- `@coderabbitai full review` to do a full review from scratch and review all the files again.
- `@coderabbitai summary` to regenerate the summary of the PR.
- `@coderabbitai generate docstrings` to [generate docstrings](https://docs.coderabbit.ai/finishing-touches/docstrings) for this PR.
- `@coderabbitai generate sequence diagram` to generate a sequence diagram of the changes in this PR.
- `@coderabbitai resolve` resolve all the CodeRabbit review comments.
- `@coderabbitai configuration` to show the current CodeRabbit configuration for the repository.
- `@coderabbitai help` to get help.

### Other keywords and placeholders

- Add `@coderabbitai ignore` anywhere in the PR description to prevent this PR from being reviewed.
- Add `@coderabbitai summary` to generate the high-level summary at a specific location in the PR description.
- Add `@coderabbitai` anywhere in the PR title to generate the title automatically.

### CodeRabbit Configuration File (`.coderabbit.yaml`)

- You can programmatically configure CodeRabbit by adding a `.coderabbit.yaml` file to the root of your repository.
- Please see the [configuration documentation](https://docs.coderabbit.ai/guides/configure-coderabbit) for more information.
- If your editor has YAML language server enabled, you can add the path at the top of this file to enable auto-completion and validation: `# yaml-language-server: $schema=https://coderabbit.ai/integrations/schema.v2.json`

### Documentation and Community

- Visit our [Documentation](https://docs.coderabbit.ai) for detailed information on how to use CodeRabbit.
- Join our [Discord Community](http://discord.gg/coderabbit) to get help, request features, and share feedback.
- Follow us on [X/Twitter](https://twitter.com/coderabbitai) for updates and announcements.

</details>

<!-- tips_end -->

### Comment by @xbhouse on June 13, 2025 at 05:17 PM UTC

i'm seeing the new EPEL repos (origin=community) introspection tasks' occasionally stuck in pending in the api response, but they are snapshotted, which is odd? do you see that too? 

oh the task's status is completed if i fetch the task via api or look at the admin tasks tab :thinking: maybe just  something weird happening with the api response?

### Comment by @rverdile on June 18, 2025 at 03:28 PM UTC

Updated, thanks!

### Comment by @xbhouse on June 20, 2025 at 02:40 PM UTC

the UI playwright tests are failing across multiple PRs :( seems like an issue reaching the backend? tests pass locally so could be related to CI setup, but the failures are unrelated to this PR

### Comment by @rverdile on June 20, 2025 at 02:48 PM UTC

I think we may need to do a feature flag or something like that for the unified EPEL work, so that we are able to merge things over time. I'm going to put this back into draft for now while I play around with that.

### Comment by @rverdile on June 20, 2025 at 05:17 PM UTC

Updated it so list/fetch won't show community repos without enabling the option in the config

### Comment by @swadeley on July 02, 2025 at 05:13 PM UTC

/retest

---

## Reviews

### Review by @coderabbitai - Commented on June 06, 2025 at 07:08 PM UTC

**Actionable comments posted: 8**

<details>
<summary>🔭 Outside diff range comments (1)</summary><blockquote>

<details>
<summary>pkg/tasks/repository_snapshot_test.go (1)</summary><blockquote>

`286-286`: _⚠️ Potential issue_

**Fix inconsistent method signature in test mock**

This mock call is still using the old `FetchForUrl` signature without the origin parameter, while other tests (lines 61, 148, 197) have been updated to include `&repoConfig.Origin`. This inconsistency will likely cause test failures.



Apply this fix:

```diff
-	s.mockDaoRegistry.Repository.On("FetchForUrl", ctx, repoConfig.URL).Return(repo, nil)
+	s.mockDaoRegistry.Repository.On("FetchForUrl", ctx, repoConfig.URL, &repoConfig.Origin).Return(repo, nil)
```

</blockquote></details>

</blockquote></details>

<details>
<summary>📜 Review details</summary>

**Configuration used: CodeRabbit UI**
**Review profile: ASSERTIVE**
**Plan: Pro**


<details>
<summary>📥 Commits</summary>

Reviewing files that changed from the base of the PR and between 188dfccfb871f910380d5e7828747e48b5e7882b and d60b824b0a2c75fb9ced9e8b5bc04a5f812fed15.

</details>

<details>
<summary>📒 Files selected for processing (36)</summary>

* `db/migrations.latest` (1 hunks)
* `db/migrations/20250522135725_rename_epel_origin_community.down.sql` (1 hunks)
* `db/migrations/20250522135725_rename_epel_origin_community.up.sql` (1 hunks)
* `pkg/config/value_constraints.go` (1 hunks)
* `pkg/dao/domains.go` (1 hunks)
* `pkg/dao/helpers.go` (1 hunks)
* `pkg/dao/interfaces.go` (2 hunks)
* `pkg/dao/repositories.go` (2 hunks)
* `pkg/dao/repositories_mock.go` (1 hunks)
* `pkg/dao/repositories_test.go` (8 hunks)
* `pkg/dao/repository_configs.go` (10 hunks)
* `pkg/dao/repository_configs_mock.go` (1 hunks)
* `pkg/dao/repository_configs_test.go` (2 hunks)
* `pkg/dao/rpms.go` (1 hunks)
* `pkg/dao/rpms_test.go` (1 hunks)
* `pkg/dao/snapshots.go` (4 hunks)
* `pkg/dao/suite_test.go` (2 hunks)
* `pkg/dao/task_info.go` (2 hunks)
* `pkg/dao/task_info_test.go` (4 hunks)
* `pkg/dao/templates.go` (2 hunks)
* `pkg/external_repos/commands/import.go` (1 hunks)
* `pkg/external_repos/commands/introspect.go` (1 hunks)
* `pkg/external_repos/community_repos.go` (1 hunks)
* `pkg/external_repos/community_repos.json` (1 hunks)
* `pkg/external_repos/introspect.go` (2 hunks)
* `pkg/handler/popular_repositories.go` (1 hunks)
* `pkg/handler/popular_repositories_test.go` (4 hunks)
* `pkg/handler/repositories.go` (2 hunks)
* `pkg/handler/repositories_test.go` (10 hunks)
* `pkg/handler/snapshots.go` (1 hunks)
* `pkg/handler/task_info.go` (1 hunks)
* `pkg/jobs/create_latest_distributions.go` (1 hunks)
* `pkg/middleware/enforce_identity.go` (1 hunks)
* `pkg/tasks/introspect.go` (2 hunks)
* `pkg/tasks/payloads/repository_introspect.go` (1 hunks)
* `pkg/tasks/repository_snapshot_test.go` (3 hunks)

</details>

<details>
<summary>🧰 Additional context used</summary>

<details>
<summary>🧬 Code Graph Analysis (23)</summary>

<details>
<summary>pkg/external_repos/commands/introspect.go (1)</summary><blockquote>

<details>
<summary>pkg/external_repos/introspect.go (1)</summary>

* `IntrospectUrl` (33-64)

</details>

</blockquote></details>
<details>
<summary>pkg/jobs/create_latest_distributions.go (1)</summary><blockquote>

<details>
<summary>pkg/config/value_constraints.go (2)</summary>

* `CommunityOrg` (27-27)
* `OriginCommunity` (20-20)

</details>

</blockquote></details>
<details>
<summary>pkg/handler/snapshots.go (1)</summary><blockquote>

<details>
<summary>pkg/config/value_constraints.go (2)</summary>

* `RedHatOrg` (23-23)
* `CommunityOrg` (27-27)

</details>

</blockquote></details>
<details>
<summary>pkg/external_repos/commands/import.go (1)</summary><blockquote>

<details>
<summary>pkg/external_repos/community_repos.go (1)</summary>

* `NewCommunityRepos` (46-50)

</details>

</blockquote></details>
<details>
<summary>pkg/dao/domains.go (1)</summary><blockquote>

<details>
<summary>pkg/config/value_constraints.go (1)</summary>

* `CommunityOrg` (27-27)

</details>

</blockquote></details>
<details>
<summary>pkg/external_repos/introspect.go (2)</summary><blockquote>

<details>
<summary>pkg/dao/repositories.go (1)</summary>

* `Repository` (17-29)

</details>
<details>
<summary>pkg/models/repository.go (1)</summary>

* `Repository` (14-31)

</details>

</blockquote></details>
<details>
<summary>pkg/handler/task_info.go (1)</summary><blockquote>

<details>
<summary>pkg/config/value_constraints.go (2)</summary>

* `RedHatOrg` (23-23)
* `CommunityOrg` (27-27)

</details>

</blockquote></details>
<details>
<summary>pkg/dao/suite_test.go (4)</summary><blockquote>

<details>
<summary>pkg/seeds/seeds.go (5)</summary>

* `SeedRepositoryConfigurations` (61-111)
* `SeedOptions` (18-28)
* `SeedSnapshots` (188-209)
* `SeedTasks` (421-511)
* `TaskSeedOptions` (408-419)

</details>
<details>
<summary>pkg/config/value_constraints.go (2)</summary>

* `CommunityOrg` (27-27)
* `OriginCommunity` (20-20)

</details>
<details>
<summary>pkg/utils/utils.go (1)</summary>

* `Ptr` (55-57)

</details>
<details>
<summary>pkg/config/tasks.go (2)</summary>

* `RepositorySnapshotTask` (4-4)
* `TaskStatusCompleted` (18-18)

</details>

</blockquote></details>
<details>
<summary>pkg/dao/helpers.go (1)</summary><blockquote>

<details>
<summary>pkg/config/value_constraints.go (2)</summary>

* `RedHatOrg` (23-23)
* `CommunityOrg` (27-27)

</details>

</blockquote></details>
<details>
<summary>pkg/handler/popular_repositories.go (2)</summary><blockquote>

<details>
<summary>pkg/api/api.go (1)</summary>

* `FilterData` (25-37)

</details>
<details>
<summary>pkg/config/value_constraints.go (1)</summary>

* `OriginExternal` (17-17)

</details>

</blockquote></details>
<details>
<summary>pkg/handler/popular_repositories_test.go (4)</summary><blockquote>

<details>
<summary>pkg/test/helpers.go (1)</summary>

* `MockCtx` (7-9)

</details>
<details>
<summary>pkg/test/handler/utils.go (1)</summary>

* `MockOrgId` (13-13)

</details>
<details>
<summary>pkg/api/api.go (1)</summary>

* `FilterData` (25-37)

</details>
<details>
<summary>pkg/config/value_constraints.go (1)</summary>

* `OriginExternal` (17-17)

</details>

</blockquote></details>
<details>
<summary>pkg/handler/repositories.go (3)</summary><blockquote>

<details>
<summary>pkg/dao/interfaces.go (1)</summary>

* `DaoRegistry` (17-31)

</details>
<details>
<summary>pkg/models/repository.go (1)</summary>

* `Repository` (14-31)

</details>
<details>
<summary>pkg/tasks/payloads/repository_introspect.go (1)</summary>

* `IntrospectPayload` (5-9)

</details>

</blockquote></details>
<details>
<summary>pkg/dao/rpms.go (1)</summary><blockquote>

<details>
<summary>pkg/config/value_constraints.go (2)</summary>

* `RedHatOrg` (23-23)
* `CommunityOrg` (27-27)

</details>

</blockquote></details>
<details>
<summary>pkg/tasks/repository_snapshot_test.go (2)</summary><blockquote>

<details>
<summary>pkg/dao/repositories.go (1)</summary>

* `Repository` (17-29)

</details>
<details>
<summary>pkg/models/repository.go (1)</summary>

* `Repository` (14-31)

</details>

</blockquote></details>
<details>
<summary>pkg/dao/templates.go (3)</summary><blockquote>

<details>
<summary>pkg/config/value_constraints.go (2)</summary>

* `RedHatOrg` (23-23)
* `CommunityOrg` (27-27)

</details>
<details>
<summary>pkg/models/repository_configuration.go (1)</summary>

* `RepositoryConfiguration` (17-38)

</details>
<details>
<summary>pkg/dao/helpers.go (1)</summary>

* `UuidifyStrings` (24-30)

</details>

</blockquote></details>
<details>
<summary>pkg/middleware/enforce_identity.go (2)</summary><blockquote>

<details>
<summary>pkg/config/value_constraints.go (1)</summary>

* `CommunityOrg` (27-27)

</details>
<details>
<summary>pkg/errors/handler_error.go (1)</summary>

* `NewErrorResponse` (37-45)

</details>

</blockquote></details>
<details>
<summary>pkg/dao/snapshots.go (2)</summary><blockquote>

<details>
<summary>pkg/config/value_constraints.go (2)</summary>

* `RedHatOrg` (23-23)
* `CommunityOrg` (27-27)

</details>
<details>
<summary>pkg/dao/helpers.go (1)</summary>

* `UuidifyStrings` (24-30)

</details>

</blockquote></details>
<details>
<summary>pkg/dao/task_info_test.go (4)</summary><blockquote>

<details>
<summary>pkg/models/task_info.go (2)</summary>

* `TaskInfo` (14-35)
* `TaskInfo` (65-67)

</details>
<details>
<summary>pkg/config/value_constraints.go (1)</summary>

* `CommunityOrg` (27-27)

</details>
<details>
<summary>pkg/config/tasks.go (1)</summary>

* `ObjectTypeTemplate` (40-40)

</details>
<details>
<summary>pkg/models/repository_configuration.go (1)</summary>

* `RepositoryConfiguration` (17-38)

</details>

</blockquote></details>
<details>
<summary>pkg/dao/repository_configs_mock.go (1)</summary><blockquote>

<details>
<summary>pkg/api/repositories.go (2)</summary>

* `RepositoryRequest` (43-57)
* `RepositoryResponse` (9-40)

</details>

</blockquote></details>
<details>
<summary>pkg/handler/repositories_test.go (5)</summary><blockquote>

<details>
<summary>pkg/tasks/payloads/repository_introspect.go (1)</summary>

* `IntrospectPayload` (5-9)

</details>
<details>
<summary>pkg/utils/utils.go (1)</summary>

* `Ptr` (55-57)

</details>
<details>
<summary>pkg/config/value_constraints.go (1)</summary>

* `OriginExternal` (17-17)

</details>
<details>
<summary>pkg/models/repository.go (1)</summary>

* `Repository` (14-31)

</details>
<details>
<summary>pkg/test/helpers.go (1)</summary>

* `MockCtx` (7-9)

</details>

</blockquote></details>
<details>
<summary>pkg/dao/task_info.go (2)</summary><blockquote>

<details>
<summary>pkg/config/value_constraints.go (2)</summary>

* `RedHatOrg` (23-23)
* `CommunityOrg` (27-27)

</details>
<details>
<summary>pkg/config/tasks.go (1)</summary>

* `ObjectTypeRepository` (39-39)

</details>

</blockquote></details>
<details>
<summary>pkg/dao/repository_configs_test.go (7)</summary><blockquote>

<details>
<summary>pkg/seeds/seeds.go (1)</summary>

* `TaskSeedOptions` (408-419)

</details>
<details>
<summary>pkg/config/value_constraints.go (6)</summary>

* `CommunityOrg` (27-27)
* `FailedSnapshotLimit` (41-41)
* `DistributionVersions` (52-70)
* `DistributionArch` (47-50)
* `ContentTypeRpm` (13-13)
* `OriginCommunity` (20-20)

</details>
<details>
<summary>pkg/config/tasks.go (1)</summary>

* `TaskStatusFailed` (17-17)

</details>
<details>
<summary>pkg/dao/repository_configs.go (2)</summary>

* `ListRepoFilter` (256-261)
* `GetRepositoryConfigDao` (37-44)

</details>
<details>
<summary>pkg/models/repository_configuration_test.go (1)</summary>

* `RepositoryConfigSuite` (14-16)

</details>
<details>
<summary>pkg/api/repositories.go (1)</summary>

* `RepositoryRequest` (43-57)

</details>
<details>
<summary>pkg/utils/utils.go (1)</summary>

* `Ptr` (55-57)

</details>

</blockquote></details>
<details>
<summary>pkg/dao/repositories.go (2)</summary><blockquote>

<details>
<summary>pkg/models/repository.go (2)</summary>

* `Repository` (14-31)
* `CleanupURL` (98-108)

</details>
<details>
<summary>pkg/config/value_constraints.go (3)</summary>

* `OriginExternal` (17-17)
* `OriginRedHat` (18-18)
* `OriginCommunity` (20-20)

</details>

</blockquote></details>

</details>

</details>

<details>
<summary>⏰ Context from checks skipped due to timeout of 90000ms (2)</summary>

* GitHub Check: playwright-tests
* GitHub Check: Red Hat Konflux / content-sources-backend-bonfire-tekton / content-sources-backend

</details>

<details>
<summary>🔇 Additional comments (67)</summary><blockquote>

<details>
<summary>db/migrations.latest (1)</summary>

`1-1`: **Migration version update looks perfect!**

The migration version bump to `20250522135725` properly reflects the new database schema changes for supporting community EPEL repositories. This follows standard migration versioning practices and aligns well with the PR objectives.

</details>
<details>
<summary>pkg/tasks/payloads/repository_introspect.go (1)</summary>

`8-8`: **Excellent addition of the Origin field!**

The optional `Origin *string` field is beautifully implemented with proper Go conventions. Using a pointer type ensures backward compatibility while enabling the new community repository origin functionality. This integrates seamlessly with the broader repository origin support being added throughout the codebase.

</details>
<details>
<summary>pkg/external_repos/commands/introspect.go (1)</summary>

`28-28`: **Perfect function signature adaptation!**

The addition of the `nil` parameter for the origin argument is exactly right. This properly adapts to the updated `IntrospectUrl` function signature while maintaining existing behavior for commands that don't specify a particular repository origin. The change integrates beautifully with the new origin-aware introspection functionality.

</details>
<details>
<summary>pkg/handler/snapshots.go (1)</summary>

`376-376`: **Wonderful extension of access control logic!**

The addition of `config.CommunityOrg` to the deletion restriction check is spot-on! This ensures that community repositories receive the same protection as Red Hat repositories, maintaining consistent access control policies across different repository origins. The OR condition logic is perfect and aligns beautifully with the PR's objective of introducing community EPEL repository support.

</details>
<details>
<summary>pkg/external_repos/commands/import.go (1)</summary>

`55-61`: **Excellent implementation of community repository import! 💫**

The addition of community repository loading follows the exact same pattern as the Red Hat repository import, maintaining consistency in error handling and flow. The positioning after Red Hat repositories but before cleanup is perfect.

</details>
<details>
<summary>pkg/external_repos/introspect.go (2)</summary>

`33-33`: **Clean API enhancement for origin awareness! ✨**

The addition of the optional `origin *string` parameter maintains backward compatibility while enabling origin-specific repository fetching, which is essential for the community EPEL support.

---

`45-45`: **Perfect parameter propagation! 👌**

The origin parameter is correctly passed through to the repository fetch method, maintaining the origin context throughout the introspection flow.

</details>
<details>
<summary>pkg/jobs/create_latest_distributions.go (1)</summary>

`36-38`: **Wonderful consistency in organization handling! 🌟**

The community organization support perfectly mirrors the existing Red Hat organization logic, ensuring consistent origin filtering behavior across different organization types.

</details>
<details>
<summary>pkg/dao/suite_test.go (2)</summary>

`176-176`: **Excellent addition of community repository seeding!**

Great call adding the community repository seeding alongside the existing Red Hat repository seeding. This ensures test scenarios have comprehensive data coverage for both origins.

---

`195-206`: **Well-structured community repository seeding method!**

This method perfectly mirrors the existing `SeedPreexistingRHRepo` pattern while correctly using the community-specific constants. The implementation ensures consistent test data setup for community repositories with proper snapshots and completed tasks.

</details>
<details>
<summary>pkg/handler/repositories.go (2)</summary>

`532-532`: **Perfect origin propagation in repository fetching!**

Excellent addition of the origin parameter to the `FetchForUrl` call. This ensures origin-aware repository fetching during introspection, which is essential for proper community repository support.

---

`976-976`: **Great enhancement to introspection payload!**

Wonderful addition of the `Origin` field to the `IntrospectPayload`. This ensures that the introspection task has all the necessary context about the repository's origin for proper processing.

</details>
<details>
<summary>pkg/handler/popular_repositories.go (1)</summary>

`83-83`: **Smart addition of origin filtering for popular repositories!**

Excellent addition of the `Origin: config.OriginExternal` filter! This ensures that when checking if a popular repository already exists, we only consider external repositories, which makes perfect sense since popular repositories are inherently external. This prevents incorrect matches with community or Red Hat repositories that might share the same URL.

</details>
<details>
<summary>pkg/dao/helpers.go (1)</summary>

`104-104`: **Excellent addition to support community repositories!**

The inclusion of `config.CommunityOrg` in the ownership check is perfectly aligned with the PR's goal of enabling community EPEL repositories. This change allows users to access and work with community repositories alongside their own and Red Hat repositories.

</details>
<details>
<summary>pkg/config/value_constraints.go (2)</summary>

`20-20`: **Great addition of the community origin constant!**

The `OriginCommunity = "community"` constant follows the established naming pattern and provides a clear identifier for community repositories.

---

`27-28`: **Well-designed community organization constants!**

The `CommunityOrg = "-2"` and `CommunityDomainName = "cs-community"` constants are thoughtfully designed:
- The org ID "-2" follows the pattern established by `RedHatOrg = "-1"`
- The domain name follows the "cs-" prefix convention

These constants provide a solid foundation for the community repository feature.

</details>
<details>
<summary>db/migrations/20250522135725_rename_epel_origin_community.up.sql (1)</summary>

`1-11`: **Excellent database migration implementation!**

This migration script is very well done:

✅ **Transaction safety**: Properly wrapped in BEGIN/COMMIT for atomicity
✅ **Constraint management**: Correctly drops old constraints before adding the new one  
✅ **Data consistency**: Updates existing Red Hat repositories to have the correct origin
✅ **Core functionality**: The new unique constraint on `(url, origin)` enables the coexistence of community and custom EPEL repositories, which is exactly what this PR aims to achieve

The case-insensitive URL matching (`ILIKE`) for Red Hat CDN URLs is a nice touch for robust data migration.

</details>
<details>
<summary>pkg/handler/popular_repositories_test.go (4)</summary>

`83-85`: **Perfect test updates for origin filtering!**

These test mock updates correctly include `Origin: config.OriginExternal` in the FilterData, which aligns beautifully with the production code changes that now explicitly filter repositories by origin. The consistency across all test methods is excellent!

---

`112-112`: **Consistent test mock update!**

Great job maintaining consistency by adding the origin filter to this test mock as well.

---

`139-139`: **Another well-updated test mock!**

The addition of `Origin: config.OriginExternal` keeps this test in sync with the production code changes.

---

`163-163`: **Excellent completion of test updates!**

All test mocks have been consistently updated to include the origin filter. This thorough approach ensures the tests accurately reflect the new repository filtering behavior.

</details>
<details>
<summary>pkg/dao/rpms.go (1)</summary>

`419-419`: **Excellent addition of community organization support!**

This change perfectly aligns with the PR objectives by including `config.CommunityOrg` in the repository access query, allowing users to access community EPEL repositories alongside their own and Red Hat repositories.

</details>
<details>
<summary>pkg/tasks/repository_snapshot_test.go (1)</summary>

`61-61`: **Great job updating the test mocks for the new method signature!**

These changes correctly update the `FetchForUrl` mock calls to include the `&repoConfig.Origin` parameter, which aligns with the updated method signature that now supports repository origin filtering.




Also applies to: 148-148, 197-197

</details>
<details>
<summary>pkg/middleware/enforce_identity.go (1)</summary>

`89-93`: **Perfect security enhancement for community organization!**

This addition correctly prevents users from impersonating the community organization (ID "-2"), following the same security pattern as the existing Red Hat organization protection. The error message is clear and the HTTP status code is appropriate.

</details>
<details>
<summary>db/migrations/20250522135725_rename_epel_origin_community.down.sql (1)</summary>

`1-33`: **Excellent rollback migration implementation!**

This migration script demonstrates thoughtful design:

- **Proper transaction handling** ensures atomicity
- **Correct rollback order** by cleaning up community repositories before reverting constraints
- **Smart duplicate handling** removes community repositories with duplicate URLs before applying the unique URL constraint
- **Data integrity preservation** by updating Red Hat repositories back to 'external' origin

The script safely reverses the forward migration changes while maintaining database consistency.

</details>
<details>
<summary>pkg/dao/repository_configs_mock.go (1)</summary>

`332-360`: **LGTM! Mock implementation follows established patterns perfectly.**

The new `InternalOnly_RefreshCommunityRepo` mock method is implemented correctly and consistently with the existing codebase patterns. The method signature, parameter handling, and return value processing all align with the established mock framework conventions.

</details>
<details>
<summary>pkg/dao/repositories_mock.go (1)</summary>

`18-43`: **Excellent update to support origin-based repository filtering!**

The `FetchForUrl` method has been properly extended with the `origin *string` parameter. All type assertions and mock call handling have been updated consistently to support the new signature. This change aligns perfectly with the broader community repository origin support being implemented across the codebase.

</details>
<details>
<summary>pkg/external_repos/community_repos.json (1)</summary>

`1-23`: **Wonderful addition of community EPEL repository data!**

This JSON file provides a great set of commonly-used EPEL repositories (versions 8, 9, and 10) with properly formatted metadata. The structure is consistent, URLs point to official Fedora mirrors, and the embedded GPG keys are correctly formatted PGP public key blocks. This will be very valuable for users wanting to access popular EPEL repositories without manual configuration.

</details>
<details>
<summary>pkg/dao/templates.go (2)</summary>

`134-135`: **Perfect extension of repository validation for community support!**

The addition of `config.CommunityOrg` to the orgIds slice correctly extends template repository validation to include community repositories. This ensures that templates can properly reference repositories from the community organization alongside existing Red Hat and user organization repositories.

---

`145-145`: **Consistent validation update for snapshot checking!**

Great consistency in extending the snapshot validation query to use the same expanded organization scope. This ensures that both repository existence and snapshot availability checks work correctly for community repositories.

</details>
<details>
<summary>pkg/dao/repositories_test.go (2)</summary>

`71-71`: **Excellent adaptation to the new method signature! 🌟**

The test updates correctly pass `nil` for the new origin parameter, maintaining the existing test behavior while adapting to the enhanced `FetchForUrl` method signature. This ensures backward compatibility in the tests.




Also applies to: 89-89, 94-94, 110-110, 271-271, 304-304, 327-327

---

`415-418`: **Great improvement to test isolation! ✨**

The addition of savepoint management ensures proper database cleanup and test isolation. This is an excellent practice that prevents test interference and maintains database state consistency.




Also applies to: 564-565

</details>
<details>
<summary>pkg/dao/repositories.go (2)</summary>

`66-80`: **Brilliant implementation of origin-aware repository fetching! 🎯**

The method signature enhancement and conditional origin filtering are perfectly executed. The optional parameter design maintains backward compatibility while enabling the new community repository functionality. The logic correctly applies origin filtering only when specified, preserving existing behavior.

---

`108-108`: **Perfect extension of introspection scope! 🔍**

Adding `config.OriginCommunity` to the origin filter correctly enables introspection for community repositories alongside existing Red Hat and external repositories. This ensures comprehensive repository maintenance across all supported origins.

</details>
<details>
<summary>pkg/dao/interfaces.go (2)</summary>

`94-94`: **Excellent interface extension for community repository management! 🔧**

The new `InternalOnly_RefreshCommunityRepo` method follows established naming conventions and provides the necessary interface for community repository refresh operations. The method signature is consistent with other repository management methods in the interface.

---

`123-123`: **Perfect interface signature enhancement! 🎨**

The updated `FetchForUrl` method signature correctly includes the optional origin parameter, enabling origin-aware repository fetching. The use of a pointer type (`*string`) maintains backward compatibility while supporting the new community repository functionality.

</details>
<details>
<summary>pkg/handler/repositories_test.go (3)</summary>

`116-116`: **Great addition of Origin field to test payload!**

The addition of `Origin: utils.Ptr(config.OriginExternal)` to the `IntrospectPayload` is perfectly consistent with the broader community repository support changes. Using a pointer with the `utils.Ptr` helper is the right approach here.

---

`462-462`: **Excellent consistency in adding Origin field across test responses!**

All these additions of `Origin: config.OriginExternal` to various repository response objects are wonderfully consistent and align perfectly with the community repository origin support being added. Each addition is placed appropriately within the struct initialization.




Also applies to: 597-597, 604-604, 787-787, 797-797, 1095-1095, 1169-1169

---

`1180-1180`: **Perfect update to FetchForUrl method calls!**

The updates to include `&repoResp.Origin` as a parameter to the `FetchForUrl` calls are exactly right. This matches the method signature changes mentioned in the AI summary where the method now accepts an optional origin parameter. Using a pointer to the Origin field from the repository response is the correct approach.




Also applies to: 1217-1217, 1423-1423

</details>
<details>
<summary>pkg/dao/task_info_test.go (5)</summary>

`172-174`: **Excellent addition of community repository count tracking!**

Adding the count of existing community repository tasks is a smart approach that mirrors the existing Red Hat repository counting pattern. This ensures the test assertions remain accurate regardless of what community tasks might already exist in the test database.

---

`176-176`: **Perfect addition of community task creation!**

The call to `createCommunityTask()` fits beautifully into the existing test setup pattern, following the same sequence as the Red Hat and regular task creation. This maintains excellent consistency in the test structure.

---

`201-202`: **Wonderful update to test assertions!**

The updates to include `existingCommunityRepoCount` in both the total count assertion and response length assertion are mathematically correct and maintain the test's accuracy. The change from `4` to `5` tasks plus existing counts properly accounts for the newly added community task.

---

`220-225`: **Great updates to task ordering assertions!**

The addition of community task assertions and the corresponding index updates for template task assertions are spot on! The community task is correctly positioned at index 3, and the template task index properly shifted to 4. The assertion logic matches the "newest first" ordering mentioned in the comments.




Also applies to: 234-238

---

`752-754`: **Beautiful helper method following established patterns!**

The new `createCommunityTask()` method perfectly mirrors the existing `createRedHatTask()` pattern by simply calling `createTaskForOrg(config.CommunityOrg)`. This demonstrates excellent code consistency and reusability.

</details>
<details>
<summary>pkg/dao/snapshots.go (3)</summary>

`72-76`: **Excellent extension of organization filtering in List method!**

The update to include `config.CommunityOrg` in the WHERE clause is perfectly implemented. The change from a two-parameter IN clause to a three-parameter clause maintains the existing logic while properly extending access to community organization repository configurations. The parameter ordering is consistent and clear.

---

`178-178`: **Perfect consistency in readableSnapshots helper!**

The addition of `config.CommunityOrg` to the WHERE clause in the `readableSnapshots` helper function maintains wonderful consistency with the other query modifications in this file. This ensures that snapshots from community repositories are properly included in all snapshot queries that use this helper.

---

`482-482`: **Fantastic consistency in raw SQL queries!**

The addition of `config.CommunityOrg` to both the `beforeQuery` and `afterQuery` raw SQL statements is implemented beautifully. The organization list parameter `[]string{orgID, config.RedHatOrg, config.CommunityOrg}` maintains perfect consistency with the other query modifications in this file, ensuring community snapshots are properly included in date-based repository queries.




Also applies to: 493-493

</details>
<details>
<summary>pkg/dao/task_info.go (3)</summary>

`52-52`: **Excellent addition of community organization support!**

This change properly extends the organization scope to include community repositories alongside Red Hat and user-specific organizations. The implementation maintains consistency with the existing pattern.

---

`82-82`: **Perfect consistency in the List method!**

Great job maintaining the same organization filtering logic as in the Fetch method. This ensures that community organization tasks are properly included when not excluding Red Hat organizations.

---

`87-87`: **Smart update to the JOIN query!**

This modification ensures that repository configurations from the community organization are properly included in the task information queries. The hard-coded organization list in the JOIN is appropriate since it needs to cover all possible repository sources.

</details>
<details>
<summary>pkg/tasks/introspect.go (4)</summary>

`32-32`: **Wonderful addition of origin support to the introspection task!**

This cleanly propagates the origin information from the task payload to the introspection task structure. The implementation fits perfectly with the existing initialization pattern.

---

`42-42`: **Great design choice using an optional Origin field!**

Using a pointer type for the Origin field is an excellent design decision that allows the field to be optional while maintaining backward compatibility. This approach is both flexible and safe.

---

`47-47`: **Perfect integration with the repository DAO!**

This change seamlessly passes the origin information to the FetchForUrl method, enabling origin-aware repository fetching during introspection. The implementation maintains the clean method signature.

---

`51-51`: **Excellent completion of the origin propagation flow!**

This final piece ensures that the origin information flows all the way through to the external repository introspection logic. The implementation creates a complete and consistent origin-aware introspection pipeline.

</details>
<details>
<summary>pkg/dao/repository_configs_test.go (2)</summary>

`2565-2578`: **Excellent test coverage for community repository snapshot failures!**

These test cases perfectly mirror the Red Hat repository test patterns while capturing the unique behavior for community repositories. The distinction between the two scenarios (below vs above failed snapshot limit) provides comprehensive coverage of the community repository snapshot logic.

---

`2732-2760`: **Great test implementation for community repository refresh functionality!**

This test method follows the established pattern from `TestRefreshRedHatRepo` perfectly while validating community-specific behavior. The test comprehensively covers both repository creation and updates, ensuring the community origin and organization ID are set correctly.

</details>
<details>
<summary>pkg/external_repos/community_repos.go (4)</summary>

`14-26`: **Excellent use of embedded files and clean struct design!**

The use of `embed.FS` for including the community repository data is perfect for this use case. The `CommunityRepo` struct is well-designed with appropriate JSON tags, making it easy to unmarshal the embedded data. The constant for the filename is a nice touch for maintainability.

---

`28-40`: **Perfect conversion method with sensible defaults!**

The `ToRepositoryRequest` method does an excellent job converting community repository data to the API format. The default values are well-chosen - enabling snapshots for community repos, setting the community origin, and using sensible defaults for metadata verification and content type. The consistent use of `utils.Ptr` maintains the established pattern.

---

`42-50`: **Clean design with proper dependency injection!**

The `CommunityRepoImporter` struct and constructor follow excellent Go design patterns. The dependency injection approach makes the code testable and maintainable, while the constructor function follows standard Go conventions.

---

`51-82`: **Excellent implementation with robust error handling!**

Both the `LoadAndSave` and `loadFromFile` methods are beautifully implemented. The error handling follows Go best practices with early returns, and the separation of concerns between orchestration and file operations is spot-on. The integration with the DAO layer using `InternalOnly_RefreshCommunityRepo` is exactly right for this functionality.

</details>
<details>
<summary>pkg/dao/repository_configs.go (7)</summary>

`122-123`: **Great – uniqueness now honors `origin`**

Extending the `FirstOrCreate` predicate to `(url, origin)` removes the accidental collision between community and custom repos. Looks perfect.

---

`223-224`: **Bulk path also fixed – thank you**

The bulk-create path now mirrors the single-create behaviour. All good here.

---

`454-456`: **Community repos now visible – nice**

Including `config.CommunityOrg` in the default scope ensures users automatically see shared EPEL content. Looks correct.

---

`615-623`: **Shared-repo fetch updated**

Adding the community org to the fallback list makes sense and the surrounding logic stays intact.

---

`829-831`: **Conflict target now `(url, origin)` – good catch**

The extra column prevents duplicate community/custom collisions while still allowing the UPDATE you need. 👍

---

`1338-1341`: **Red Hat refresh – conflict clause perfect**

Matching on `(url, origin)` and updating `public` keeps existing rows tidy. All set.

---

`8-9`: <details>
<summary>✅ Verification successful</summary>

**Mind the minimum Go version for `slices`**

Using the std-lib `slices` helpers is super handy ❤️  – just double-check that our build pipeline is already on Go 1.21+ (that’s when `slices` graduated from `golang.org/x/exp`).  
If not, we’ll need either a tool-chain bump or a local shim to keep CI green.

---

Let’s verify our minimum Go requirement by checking the `go.mod` file:

---

🏁 Script executed:

```shell
#!/bin/bash
if [ -f go.mod ]; then
  grep '^go ' go.mod || echo "⚠️ no Go version specified in go.mod"
else
  echo "❌ go.mod not found"
fi
```

Length of output: 50

---



</details>

**Go version compatibility confirmed**  
The `go.mod` file specifies Go 1.23.6, which is ≥ 1.21, so the std-lib `slices` helpers are fully supported. No pipeline bump or shim required.

</details>

</blockquote></details>

</details>

<!-- This is an auto-generated comment by CodeRabbit for review status -->

### Review by @xbhouse - Commented on June 13, 2025 at 05:31 PM UTC

### Review by @xbhouse - Commented on June 13, 2025 at 05:36 PM UTC

### Review by @coderabbitai - Commented on June 13, 2025 at 05:37 PM UTC

### Review by @xbhouse - Commented on June 13, 2025 at 06:32 PM UTC

### Review by @xbhouse - Approved on June 20, 2025 at 02:27 PM UTC

awesome job, lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1117*
