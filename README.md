> # 🔗 Click!
>
> Link Manager as a Service &mdash; your personal link storage and URL shortener.

[![Build][build.icon]][build.page]
[![Template][template.icon]][template.page]

## Roadmap

- [x] v1: [MVP][project_v1]
  - [**May 31, 2018**][project_v1_dl]
  - Main concepts and working prototype.
- [x] v2: [Accounts and CLI CRUD][project_v2]
  - [**August 31, 2018**][project_v2_dl]
  - Command line interface for create, read, update and delete operations above gRPC.
- [ ] v3: [URL shortener and RESTful API][project_v3]
  - [**Someday, 20xx**][project_v3_dl]
  - URL shortener functionality.
  - Integrate gRPC gateway.
  - Improve gRPC layer.
- [ ] v4: [DSL for rules and CSI][project_v4]
  - [**Sometime, 20xx**][project_v4_dl]
  - Domain-specific language to define target rules.
  - Client-side integration.
  - Graphical user interface and admin panel to perform create, read, update and delete operations.

## Motivation

- We need better integration with static sites built with [Hugo](https://gohugo.io/).
- We want better products than [Bitly](https://bitly.com/) or [Ow.ly](http://ow.ly/).
- We have to full control over our users' data and protect it from third parties.

## Quick start

Requirements:

- Docker 18.06.0-ce or above
- Docker Compose 1.22.0 or above
- Go 1.9.2 or above
- GNU Make 3.81 or above

```bash
$ make up demo status

     Name                    Command               State                          Ports
---------------------------------------------------------------------------------------------------------------
click_db_1        docker-entrypoint.sh postgres    Up      0.0.0.0:5432->5432/tcp
click_server_1    /bin/sh -c echo $BASIC_USE ...   Up      0.0.0.0:443->443/tcp, 0.0.0.0:80->80/tcp
click_service_1   service run --with-profili ...   Up      0.0.0.0:8080->80/tcp, 0.0.0.0:8090->8090/tcp,
                                                           0.0.0.0:8091->8091/tcp, 0.0.0.0:8092->8092/tcp

$ open http://127.0.0.1.xip.io/github/click

$ make help
```

<details>
<summary><strong>GET curl /api/v1/UUID</strong></summary>

```bash
$ curl http://127.0.0.1.xip.io/api/v1/10000000-2000-4000-8000-160000000005 | jq
# {
#   "id": "10000000-2000-4000-8000-160000000005",
#   "name": "Click! - Link Manager as a Service",
#   "aliases": [
#     {
#       "id": "10000000-2000-4000-8000-160000000008",
#       "namespace": "10000000-2000-4000-8000-160000000001",
#       "urn": "github/click"
#     },
#     {
#       "id": "10000000-2000-4000-8000-160000000007",
#       "namespace": "10000000-2000-4000-8000-160000000001",
#       "urn": "github/click!"
#     },
#     {
#       "id": "10000000-2000-4000-8000-160000000006",
#       "namespace": "10000000-2000-4000-8000-160000000004",
#       "urn": "github/click"
#     }
#   ],
#   "targets": [
#     {
#       "id": "10000000-2000-4000-8000-160000000011",
#       "rule": {
#         "description": "Project's source code",
#         "tags": [
#           "src"
#         ]
#       },
#       "url": "https://github.com/kamilsk/click"
#     },
#     {
#       "id": "10000000-2000-4000-8000-160000000009",
#       "rule": {
#         "description": "Project's bug tracker",
#         "alias": "10000000-2000-4000-8000-160000000006"
#       },
#       "url": "https://github.com/kamilsk/click/issues/new"
#     },
#     {
#       "id": "10000000-2000-4000-8000-160000000010",
#       "rule": {
#         "description": "Project's promo page",
#         "alias": "10000000-2000-4000-8000-160000000007",
#         "tags": [
#           "promo"
#         ],
#         "match": 1
#       },
#       "url": "https://kamilsk.github.io/click/"
#     }
#   ]
# }

$ curl -H "X-Click-Namespace: 10000000-2000-4000-8000-160000000001" -v http://127.0.0.1.xip.io/github/click!
# > GET /github/click! HTTP/1.1
# > Host: 127.0.0.1.xip.io
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
<summary><strong>Service command-line interface</strong></summary>

```bash
$ make install

$ click help
Click!

Usage:
  click [command]

Available Commands:
  completion  Print Bash or Zsh completion
  ctl         Click! Service Control
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
$ export REQ_VER=2.0.0  # all available versions are on https://github.com/kamilsk/click/releases/
$ export REQ_OS=Linux   # macOS and Windows are also available
$ export REQ_ARCH=64bit # 32bit is also available
# wget -q -O click.tar.gz
$ curl -sL -o click.tar.gz \
       https://github.com/kamilsk/click/releases/download/"${REQ_VER}/click_${REQ_VER}_${REQ_OS}-${REQ_ARCH}".tar.gz
$ tar xf click.tar.gz -C "${GOPATH}"/bin/ && rm click.tar.gz
```

### Docker Hub

```bash
$ docker pull kamilsk/click:2.x
# or use mirror
$ docker pull quay.io/kamilsk/click:2.x
```

### From source code

```bash
$ egg github.com/kamilsk/click@^2.0.0 -- make test install
# or use mirror
$ egg bitbucket.org/kamilsk/click@^2.0.0 -- make test install
```

> [egg](https://github.com/kamilsk/egg)<sup id="anchor-egg">[1](#egg)</sup> is an `extended go get`.

<sup id="egg">1</sup> The project is still in prototyping.[↩](#anchor-egg)

---

made with ❤️ for everyone

[build.page]:       https://travis-ci.com/octopot/click
[build.icon]:       https://travis-ci.com/octopot/click.svg?branch=master
[design.page]:      https://www.notion.so/octolab/Click-e376b1f4efb34a188dfe210bffc1b112?r=0b753cbf767346f5a6fd51194829a2f3
[promo.page]:       https://kamilsk.github.io/click/
[template.page]:    https://github.com/octomation/go-service
[template.icon]:    https://img.shields.io/badge/template-go--service-blue

[egg]:              https://github.com/kamilsk/egg

[project_v1]:       https://github.com/kamilsk/click/projects/1
[project_v1_dl]:    https://github.com/kamilsk/click/milestone/1
[project_v2]:       https://github.com/kamilsk/click/projects/2
[project_v2_dl]:    https://github.com/kamilsk/click/milestone/2
[project_v3]:       https://github.com/kamilsk/click/projects/3
[project_v3_dl]:    https://github.com/kamilsk/click/milestone/3
[project_v4]:       https://github.com/kamilsk/click/projects/4
[project_v4_dl]:    https://github.com/kamilsk/click/milestone/4
