> # 🔗 Click! [![Tweet][icon_twitter]][twitter_publish] <img align="right" width="126" src=".github/character.png">
> [![Analytics][analytics_pixel]][page_promo]
> Link Manager as a Service &mdash; your personal link storage and URL shortener.

[![Patreon][icon_patreon]](https://www.patreon.com/octolab)
[![Build Status][icon_build]][page_build]
[![Code Quality][icon_quality]][page_quality]
[![Go version][icon_go_min]][page_build]
[![Code Coverage][icon_coverage]][page_quality]
[![Research][icon_research]](../../tree/research)
[![License][icon_license]](LICENSE)

## Roadmap

- [x] v1: [MVP][project_v1]
  - [**May 31, 2018**][project_v1_dl]
  - Main concepts and working prototype.
- [ ] v2: [Accounts and CLI CRUD][project_v2]
  - [**August 31, 2018**][project_v2_dl]
  - Command line interface for create, read, update and delete operations above gRPC.
- [ ] v3: [URL shortener and RESTful API][project_v3]
  - [**September 30, 2018**][project_v3_dl]
  - URL shortener functionality.
  - Integrate gRPC gateway.
  - Improve gRPC layer.
- [ ] v4: [DSL for rules, CSI, and GUI][project_v4]
  - [**October 31, 2018**][project_v4_dl]
  - Domain-specific language to define target rules.
  - Client-side integration.
  - Graphical user interface and admin panel to perform create, read, update and delete operations.
- [ ] Click!, SaaS
  - **December 31, 2018**
  - Ready to apply on Cloud.
  - Move to [OctoLab](https://github.com/octolab/) organization.

## Motivation

- We need better integration with static sites built with [Hugo](https://gohugo.io/).
- We want better products than [Bitly](https://bitly.com/) or [Ow.ly](http://ow.ly/).
- We have to full control over our users' data and protect it from third parties.

## Quick start

Requirements:

- Docker 17.09.0-ce or above
- Docker Compose 1.16.1 or above
- Go 1.9.2 or above
- GNU Make 3.81 or above

```bash
$ make up demo status

     Name                    Command               State                                  Ports
-------------------------------------------------------------------------------------------------------------------------------
click_db_1        docker-entrypoint.sh postgres    Up      0.0.0.0:5432->5432/tcp
click_server_1    /bin/sh -c envsubst '$SERV ...   Up      80/tcp, 0.0.0.0:80->8080/tcp
click_service_1   click run --with-profiling ...   Up      0.0.0.0:8080->80/tcp, 0.0.0.0:8090->8090/tcp, 0.0.0.0:8091->8091/tcp
```

<details>
<summary><strong>GET curl /api/v1/UUID</strong></summary>

```bash
$ curl http://localhost:8080/api/v1/10000000-2000-4000-8000-160000000005
# {
#   "id": "10000000-2000-4000-8000-160000000005",
#   "name": "Click! - Link Manager as a Service",
#   "status": "active",
#   "aliases": [
#     {
#       "id": 1,
#       "namespace": "global",
#       "urn": "github/click"
#     },
#     {
#       "id": 7,
#       "namespace": "global",
#       "urn": "github/click!"
#     }
#   ],
#   "targets": [
#     {
#       "id": 1,
#       "uri": "https://github.com/kamilsk/click",
#       "rule": {
#         "description": "Project location",
#         "tags": ["src"]
#       }
#     },
#     {
#       "id": 2,
#       "uri": "https://kamilsk.github.io/click/",
#       "rule": {
#         "description": "Promotion page",
#         "alias": 7,
#         "tags": ["promo"],
#         "match": 1
#       }
#     }
#   ]
# }
$ curl -v http://localhost:8080/github/click!
# > GET /github/click! HTTP/1.1
# > Host: localhost:8080
# > User-Agent: curl/7.54.0
# > Accept: */*
# >
# < HTTP/1.1 302 Found
# < Location: https://kamilsk.github.io/click/
# < Date: Sat, 05 May 2018 09:34:47 GMT
# < Content-Length: 0
# <
```
</details>

## Specification

### API

You can find API specification [here](env/client/rest.http). Also, we recommend using [Insomnia](https://insomnia.rest/)
HTTP client to work with the API - you can import data for it from the [file](env/client/insomnia.json).
Or you can choose [Postman](https://www.getpostman.com/) - its import data is [here](env/client/postman.json) and
[here](env/client/postman.env.json).

### CLI

You can use CLI not only to start the HTTP server but also to execute
[CRUD](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) operations.

<details>
<summary><strong>CLI interface</strong></summary>

```bash
$ click --help
Click!

Usage:
  click [command]

Available Commands:
  completion  Print Bash or Zsh completion
  ctl         Communicate with Click! server via gRPC
  help        Help about any command
  migrate     Apply database migration
  run         Start HTTP server
  version     Show application version

Flags:
  -h, --help   help for click

Use "click [command] --help" for more information about a command.
```
</details>

#### Bash and Zsh completions

You can find completion files [here](https://github.com/kamilsk/shared/tree/dotfiles/bash_completion.d) or
build your own using these commands

```bash
$ click completion -f bash > /path/to/bash_completion.d/click.sh
$ click completion -f zsh  > /path/to/zsh-completions/_click.zsh
```

## Installation

### Brew

```bash
$ brew install kamilsk/tap/click
```

### Binary

```bash
$ export VER=1.0.0      # all available versions are on https://github.com/kamilsk/click/releases/
$ export REQ_OS=Linux   # macOS and Windows are also available
$ export REQ_ARCH=64bit # 32bit is also available
$ wget -q -O click.tar.gz \
       https://github.com/kamilsk/click/releases/download/"${VER}/click_${VER}_${REQ_OS}-${REQ_ARCH}".tar.gz
$ tar xf click.tar.gz -C "${GOPATH}"/bin/ && rm click.tar.gz
```

### Docker Hub

```bash
$ docker pull kamilsk/click:1.x
```

### From source code

```bash
$ egg github.com/kamilsk/click@^1.0.0 -- make test install
```

#### Mirror

```bash
$ egg bitbucket.org/kamilsk/click@^1.0.0 -- make test install
```

> [egg](https://github.com/kamilsk/egg)<sup id="anchor-egg">[1](#egg)</sup> is an `extended go get`.

## Update

This application is in a state of [MVP](https://en.wikipedia.org/wiki/Minimum_viable_product) and under active
development. [SemVer](https://semver.org/) is used for releases, and you can easily be updated within minor versions,
but major versions can be not [BC](https://en.wikipedia.org/wiki/Backward_compatibility)-safe.

<sup id="egg">1</sup> The project is still in prototyping. [↩](#anchor-egg)

---

[![Gitter][icon_gitter]](https://gitter.im/kamilsk/click)
[![@kamilsk][icon_tw_author]](https://twitter.com/ikamilsk)
[![@octolab][icon_tw_sponsor]](https://twitter.com/octolab_inc)

made with ❤️ by [OctoLab](https://www.octolab.org/)

[analytics_pixel]: https://ga-beacon.appspot.com/UA-109817251-20/click/readme?pixel

[icon_build]:      https://travis-ci.org/kamilsk/click.svg?branch=master
[icon_coverage]:   https://scrutinizer-ci.com/g/kamilsk/click/badges/coverage.png?b=master
[icon_gitter]:     https://badges.gitter.im/Join%20Chat.svg
[icon_go_min]:     https://img.shields.io/badge/Go-%3E%3D%201.9.2-green.svg
[icon_license]:    https://img.shields.io/badge/license-MIT-blue.svg
[icon_patreon]:    https://img.shields.io/badge/patreon-donate-orange.svg
[icon_quality]:    https://scrutinizer-ci.com/g/kamilsk/click/badges/quality-score.png?b=master
[icon_research]:   https://img.shields.io/badge/research-in%20progress-yellow.svg
[icon_tw_author]:  https://img.shields.io/badge/author-%40kamilsk-blue.svg
[icon_tw_sponsor]: https://img.shields.io/badge/sponsor-%40octolab-blue.svg
[icon_twitter]:    https://img.shields.io/twitter/url/http/shields.io.svg?style=social

[page_build]:      https://travis-ci.org/kamilsk/click
[page_promo]:      https://kamilsk.github.io/click/
[page_quality]:    https://scrutinizer-ci.com/g/kamilsk/click/?branch=master

[project_v1]:      https://github.com/kamilsk/click/projects/1
[project_v1_dl]:   https://github.com/kamilsk/click/milestone/1
[project_v2]:      https://github.com/kamilsk/click/projects/2
[project_v2_dl]:   https://github.com/kamilsk/click/milestone/2
[project_v3]:      https://github.com/kamilsk/click/projects/3
[project_v3_dl]:   https://github.com/kamilsk/click/milestone/3
[project_v4]:      https://github.com/kamilsk/click/projects/4
[project_v4_dl]:   https://github.com/kamilsk/click/milestone/4

[twitter_publish]: https://twitter.com/intent/tweet?text=Link%20Manager%20as%20a%20Service&url=https://kamilsk.github.io/click/&via=ikamilsk&hashtags=go,service,link-manager,link-storage,link-shortener,url-shortener
