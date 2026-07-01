---
type: pull_request
number: 175
title: "Fixes 540: Add docker compose"
state: merged
author: jlsherrill
created: 2023-01-12T03:56:10Z
updated: 2023-01-24T20:42:22Z
closed: 2023-01-24T20:42:18Z
merged: 2023-01-24T20:42:18Z
base_branch: main
head_branch: 328_1
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/175
---

# Pull Request #175: Fixes 540: Add docker compose

**Author**: @jlsherrill
**Created**: January 12, 2023 at 03:56 AM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `328_1`

## Description

This converts our docker containers to a podman/docker-compose file and introduces these commands:
* make compose-up
* make compose-down
* make compose-clean
* make compose-build

and removes dedicated kafka and db commands.  In addition this:
* Changes kafka container build to use a quay container that does not require login
* Changes kafka container build to determine newest version at build time.  This enables building via podman-compose directly and simplifies makefile
* Moves all storage to container volumes (it seems like zookeeper only needs the volume, not kafka itself)
  * this simplifies cleanup and the makefiles

You can spinup a pulp by setting `DEPLOY_PULP=true`:

```
DEPLOY_PULP=true make compose-up
DEPLOY_PULP=true make compose-down
```
This will be made permanent in the future.  Then you can interact with pulp after a couple of minutes:
```
curl localhost:8080/pulp/api/v3/status/
```


This compose doesn't include prometheus bits, and since thats not needed for all deployments i lean towards keeping it that way.

- [x] requires https://github.com/pulp/pulp-oci-images/pull/411 to be merged

For testing, you'll want to check:
* compose up/down with pulp
* compose up/down without pulp 
* both of the above with docker 

---

## Discussion

### Comment by @avisiedo on January 17, 2023 at 06:27 AM UTC

I have not checked with pulp; I will do a second round.

### Comment by @avisiedo on January 18, 2023 at 02:58 PM UTC

I have to update with the results from macos system! I get back with them in a few minutes!

### Comment by @avisiedo on January 18, 2023 at 03:24 PM UTC

Results after run in macOS m1 using docker-compose (it only change on the change about declaring the `volumes:` section).

```raw
make compose-up DEPLOY_PULP=true   
info:Trying to load DATABASE configuration from '/Users/avisiedo/github/content-services/content-sources-backend/configs/config.yaml'
CONTENT_DATABASE_USER=content CONTENT_DATABASE_PASSWORD=content CONTENT_DATABASE_DATABASE_NAME=content CONTENT_DATABASE_PORT=5434 \
	KAFKA_CONFIG_DIR=/Users/avisiedo/github/content-services/content-sources-backend/compose_files/kafka/config KAFKA_DATA_DIR=/Users/avisiedo/github/content-services/content-sources-backend/compose_files/kafka/data ZOOKEEPER_CLIENT_PORT=2181 KAFKA_TOPICS=platform.content-sources.introspect  \
	docker-compose --project-name=cs -f "deployments/docker-compose.yml" up --detach
WARN[0000] mount of type `volume` should not define `bind` option 
[+] Running 6/6
 ⠿ Network cs_default               Created                                                                                                                                                                                                     0.0s
 ⠿ Volume "cs_zookeeper"            Created                                                                                                                                                                                                     0.0s
 ⠿ Volume "cs_database"             Created                                                                                                                                                                                                     0.0s
 ⠿ Container cs-postgres-content-1  Started                                                                                                                                                                                                     0.3s
 ⠿ Container cs-zookeeper-1         Started                                                                                                                                                                                                     0.4s
 ⠿ Container cs-kafka-1             Started                                                                                                                                                                                                     0.6s
PULP_POSTGRES_PATH="pulp_db" PULP_STORAGE_PATH="pulp_storage" docker-compose  -f "compose_files/pulp/pulp-oci-images/images/compose/docker-compose.yml" up --detach
WARN[0000] Found orphan containers ([cs-kafka-1 cs-zookeeper-1 cs-postgres-content-1]) for this project. If you removed or renamed this service in your compose file, you can run this command with the --remove-orphans flag to clean it up. 
[+] Running 12/16
 ⠿ Volume "pulpdev"                                                                                                                                            Created                                                                          0.0s
 ⠿ Volume "pg_datadev"                                                                                                                                         Created                                                                          0.0s
 ⠿ Volume "redis_datadev"                                                                                                                                      Created                                                                          0.0s
 ⠿ Container cs-redis-1                                                                                                                                        Started                                                                          0.4s
 ⠿ Container cs-postgres-1                                                                                                                                     Started                                                                          0.4s
 ⠿ Container cs-pulp_api-2                                                                                                                                     Started                                                                          1.3s
 ⠿ Container cs-pulp_content-2                                                                                                                                 Started                                                                          0.9s
 ⠿ Container cs-pulp_worker-2                                                                                                                                  Started                                                                          0.9s
 ⠿ Container cs-pulp_worker-1                                                                                                                                  Started                                                                          1.0s
 ⠿ Container cs-pulp_api-1                                                                                                                                     Started                                                                          1.3s
 ⠿ Container cs-pulp_content-1                                                                                                                                 Started                                                                          1.1s
 ⠧ pulp_content The requested image's platform (linux/amd64) does not match the detected host platform (linux/arm64/v8) and no specific platform was requested                                                                                  0.0s
 ⠧ pulp_api The requested image's platform (linux/amd64) does not match the detected host platform (linux/arm64/v8) and no specific platform was requested                                                                                      0.0s
 ⠧ pulp_worker The requested image's platform (linux/amd64) does not match the detected host platform (linux/arm64/v8) and no specific platform was requested                                                                                   0.0s
 ⠿ Container cs-pulp_web-1                                                                                                                                     Started                                                                          1.6s
 ⠧ pulp_web The requested image's platform (linux/amd64) does not match the detected host platform (linux/arm64/v8) and no specific platform was requested                                                                                      0.0s
/Applications/Xcode.app/Contents/Developer/usr/bin/make .db-health-wait
info:Trying to load DATABASE configuration from '/Users/avisiedo/github/content-services/content-sources-backend/configs/config.yaml'
.../Applications/Xcode.app/Contents/Developer/usr/bin/make db-migrate-up
info:Trying to load DATABASE configuration from '/Users/avisiedo/github/content-services/content-sources-backend/configs/config.yaml'
/Users/avisiedo/github/content-services/content-sources-backend/release/dbmigrate up
{"level":"warn","time":"2023-01-18T16:20:36+01:00","message":"No Red Hat CDN cert pair configured."}
0x140002b19a8
{"level":"debug","time":"2023-01-18T16:20:37+01:00","message":"Successfully migrated up"}
Run 'make db-migrate-seed' to seed the database
```

And the list of containers:

```raw
docker container ls
CONTAINER ID   IMAGE                           COMMAND                  CREATED              STATUS                        PORTS                                            NAMES
8997c85cb2d7   pulp/pulp-web:latest            "container-entrypoin…"   About a minute ago   Up About a minute             0.0.0.0:8080->8080/tcp, 8443/tcp                 cs-pulp_web-1
7eae4347ca10   pulp/pulp-minimal:latest        "/pulp-common-entryp…"   About a minute ago   Up About a minute                                                              cs-pulp_worker-1
c86c767fd919   pulp/pulp-minimal:latest        "/pulp-common-entryp…"   About a minute ago   Up About a minute                                                              cs-pulp_worker-2
c67374cab6f4   pulp/pulp-minimal:latest        "/pulp-common-entryp…"   About a minute ago   Up About a minute                                                              cs-pulp_content-1
15f078bcb34d   pulp/pulp-minimal:latest        "/pulp-common-entryp…"   About a minute ago   Up About a minute                                                              cs-pulp_api-2
3b29160b02a3   pulp/pulp-minimal:latest        "/pulp-common-entryp…"   About a minute ago   Up About a minute                                                              cs-pulp_api-1
8ce7cf7c66f0   pulp/pulp-minimal:latest        "/pulp-common-entryp…"   About a minute ago   Up About a minute                                                              cs-pulp_content-2
061500ba5d47   postgres:13                     "docker-entrypoint.s…"   About a minute ago   Up About a minute             0.0.0.0:5432->5432/tcp                           cs-postgres-1
83447c91b07a   redis:latest                    "docker-entrypoint.s…"   About a minute ago   Up About a minute             6379/tcp                                         cs-redis-1
7a126163e48c   localhost/kafka:latest          "/opt/kafka/scripts/…"   2 minutes ago        Up About a minute             0.0.0.0:9092->9092/tcp                           cs-kafka-1
4a4d1f43aa80   localhost/kafka:latest          "/opt/kafka/scripts/…"   2 minutes ago        Up About a minute (healthy)   0.0.0.0:2181->2181/tcp, 0.0.0.0:8778->8778/tcp   cs-zookeeper-1
d10f4105fd9c   postgres:14                     "docker-entrypoint.s…"   2 minutes ago        Up About a minute (healthy)   0.0.0.0:5434->5432/tcp                           cs-postgres-content-1
fe38c333fa52   moby/buildkit:buildx-stable-1   "buildkitd"              5 weeks ago          Up 5 hours                                                                     buildx_buildkit_mybuilder0
```



### Comment by @avisiedo on January 19, 2023 at 12:51 PM UTC

I have tried the last changes in macos apple m1 environment using docker and docker-compose and fedora 36 linux x86 using podman and podman-compose and the commands `make compose-up DEPLOY_PULP=true` and `make compose-down compose-clean DEPLOY_PULP=true` have worked fine! nice change!

I see a warning from docker indicating that the architecture specified does not match the host (no multiarch container images I guess); but the **containers are up and running** (rosetta2? no it is not), so it is working.

To be more specific:

```raw
make compose-up DEPLOY_PULP=true  
info:Trying to load DATABASE configuration from '/Users/avisiedo/github/content-services/content-sources-backend/configs/config.yaml'
CONTENT_DATABASE_USER=content CONTENT_DATABASE_PASSWORD=content CONTENT_DATABASE_DATABASE_NAME=content CONTENT_DATABASE_PORT=5434 \
	KAFKA_CONFIG_DIR=/Users/avisiedo/github/content-services/content-sources-backend/compose_files/kafka/config KAFKA_DATA_DIR=/Users/avisiedo/github/content-services/content-sources-backend/compose_files/kafka/data ZOOKEEPER_CLIENT_PORT=2181 KAFKA_TOPICS=platform.content-sources.introspect  \
	docker-compose --project-name=cs -f "deployments/docker-compose.yml" up --detach
WARN[0000] mount of type `volume` should not define `bind` option 
[+] Running 4/4
 ⠿ Network cs_default               Created                                                                                                                                                                                                     0.0s
 ⠿ Container cs-postgres-content-1  Started                                                                                                                                                                                                     0.3s
 ⠿ Container cs-zookeeper-1         Started                                                                                                                                                                                                     0.4s
 ⠿ Container cs-kafka-1             Started                                                                                                                                                                                                     0.6s
PULP_POSTGRES_PATH="pulp_db" PULP_STORAGE_PATH="pulp_storage" docker-compose --project-name=cs  -f "compose_files/pulp/pulp-oci-images/images/compose/docker-compose.yml" up --detach
WARN[0000] Found orphan containers ([cs-kafka-1 cs-zookeeper-1 cs-postgres-content-1]) for this project. If you removed or renamed this service in your compose file, you can run this command with the --remove-orphans flag to clean it up. 
[+] Running 9/13
 ⠿ Container cs-redis-1                                                                                                                                        Started                                                                          0.4s
 ⠿ Container cs-postgres-1                                                                                                                                     Started                                                                          0.4s
 ⠿ Container cs-pulp_worker-2                                                                                                                                  Started                                                                          0.9s
 ⠿ Container cs-pulp_api-2                                                                                                                                     Started                                                                          1.0s
 ⠿ Container cs-pulp_content-2                                                                                                                                 Started                                                                          0.9s
 ⠿ Container cs-pulp_api-1                                                                                                                                     Started                                                                          1.0s
 ⠿ Container cs-pulp_worker-1                                                                                                                                  Started                                                                          0.6s
 ⠿ Container cs-pulp_content-1                                                                                                                                 Started                                                                          1.0s
 ⠼ pulp_api The requested image's platform (linux/amd64) does not match the detected host platform (linux/arm64/v8) and no specific platform was requested                                                                                      0.0s
 ⠼ pulp_worker The requested image's platform (linux/amd64) does not match the detected host platform (linux/arm64/v8) and no specific platform was requested                                                                                   0.0s
 ⠼ pulp_content The requested image's platform (linux/amd64) does not match the detected host platform (linux/arm64/v8) and no specific platform was requested                                                                                  0.0s
 ⠿ Container cs-pulp_web-1                                                                                                                                     Started                                                                          1.4s
 ⠸ pulp_web The requested image's platform (linux/amd64) does not match the detected host platform (linux/arm64/v8) and no specific platform was requested                                                                                      0.0s
/Applications/Xcode.app/Contents/Developer/usr/bin/make .db-health-wait
info:Trying to load DATABASE configuration from '/Users/avisiedo/github/content-services/content-sources-backend/configs/config.yaml'
```

### Comment by @avisiedo on January 19, 2023 at 05:27 PM UTC

lgtm

### Comment by @avisiedo on January 21, 2023 at 04:57 PM UTC

UPDATE: linux/amd64 are supported in apple m1 architectures because the docker desktop is using a linux vm which has support for qemu, and binfmt (more info here: https://docs.docker.com/build/building/multi-platform/#support-on-docker-desktop); that is what make the magic, not rossetta2; this means that no multiarch container images, will run in macOS systems using apple m1; so the virtualisation in this case is emulating the architecture; this has a big impact into the performance, so it is expected bigger time processing for that containers when using apple m1.

### Comment by @jlsherrill on January 24, 2023 at 05:25 PM UTC

https://issues.redhat.com/browse/HMS-540

---

## Reviews

### Review by @avisiedo - Changes Requested on January 17, 2023 at 05:46 AM UTC

I am enthusiastic of containers, so I like this change a lot which enhance the local infrastructure definition. Nice change :+1:

I suggest to add prometheus container definition to the compose file.

### Review by @avisiedo - Commented on January 17, 2023 at 03:25 PM UTC

### Review by @avisiedo - Commented on January 17, 2023 at 05:54 PM UTC

Second round!

I updated the comment about `PULP_COMPOSE_*_COMMAND` to add the project option to prefix pulp resources with `pulp_` prefix string.

:+1: tested in linux with podman-compose and all the containers started, stopped and cleaned. Below the result after start all the containers by `make compose-up DEPLOY_PULP=true`.

```raw
$ podman ps
CONTAINER ID  IMAGE                             COMMAND               CREATED         STATUS                        PORTS                                           NAMES
a1c569b41af5  docker.io/library/postgres:14     postgres              32 seconds ago  Up 32 seconds ago (healthy)   0.0.0.0:5433->5432/tcp                          cs_postgres-content_1
fa96ce50d853  localhost/kafka:latest            /opt/kafka/script...  30 seconds ago  Up 31 seconds ago (starting)  0.0.0.0:2181->2181/tcp, 0.0.0.0:8778->8778/tcp  cs_zookeeper_1
bf70bf888878  localhost/kafka:latest            /opt/kafka/script...  29 seconds ago  Up 29 seconds ago             0.0.0.0:9092->9092/tcp                          cs_kafka_1
2d794bc23eb7  docker.io/library/postgres:13     postgres              27 seconds ago  Up 28 seconds ago             0.0.0.0:5432->5432/tcp                          compose_postgres_1
a3f8f8b2c75b  docker.io/library/redis:latest    redis-server          25 seconds ago  Up 26 seconds ago                                                             compose_redis_1
be89bd3c534c  quay.io/pulp/pulp-minimal:latest  pulp-api              24 seconds ago  Up 25 seconds ago                                                             compose_pulp_api_1
9f0aca5cd0b5  quay.io/pulp/pulp-minimal:latest  pulp-api              23 seconds ago  Up 23 seconds ago                                                             compose_pulp_api_2
97bdd86cf41a  quay.io/pulp/pulp-minimal:latest  pulp-content          21 seconds ago  Up 22 seconds ago                                                             compose_pulp_content_1
5080f97f5528  quay.io/pulp/pulp-minimal:latest  pulp-content          20 seconds ago  Up 21 seconds ago                                                             compose_pulp_content_2
f476266f4ccd  quay.io/pulp/pulp-minimal:latest  pulp-worker           19 seconds ago  Up 19 seconds ago                                                             compose_pulp_worker_1
656aad2f7a1e  quay.io/pulp/pulp-minimal:latest  pulp-worker           17 seconds ago  Up 18 seconds ago                                                             compose_pulp_worker_2
ff507a679871  quay.io/pulp/pulp-web:latest      /usr/bin/nginx.sh     16 seconds ago  Up 17 seconds ago             0.0.0.0:8080->8080/tcp                          compose_pulp_web_1
```


### Review by @jlsherrill - Commented on January 17, 2023 at 08:55 PM UTC

### Review by @avisiedo - Commented on January 17, 2023 at 08:58 PM UTC

### Review by @avisiedo - Changes Requested on January 18, 2023 at 02:23 PM UTC

After this change, I got the containers up and running with no variables override from the environment. Sorry for the mistake setting the variable `DATABASE_CONTAINER_NAME`.

### Review by @avisiedo - Commented on January 18, 2023 at 03:19 PM UTC

### Review by @jlsherrill - Commented on January 18, 2023 at 06:48 PM UTC

### Review by @jlsherrill - Commented on January 18, 2023 at 07:16 PM UTC

### Review by @jlsherrill - Commented on January 18, 2023 at 07:36 PM UTC

### Review by @avisiedo - Commented on January 19, 2023 at 08:22 PM UTC

### Review by @avisiedo - Commented on January 19, 2023 at 08:42 PM UTC

### Review by @avisiedo - Approved on January 19, 2023 at 08:43 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/175*
