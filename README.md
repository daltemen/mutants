# mutants 🧟‍♂️🧟‍♀️🦹‍♀️🦸‍♂️

An example of how to use Go + Echo, AppEngine, MySql, Redis and an approach of clean architecture.

<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-85%25-brightgreen.svg?longCache=true&style=flat)</a>

- The DNA analyzed are saved on Mysql.
- While a DNA is analyzed the counters are saved on Redis. Why ? see: [storing-counters-in-redis](https://redislabs.com/ebook/part-2-core-concepts/chapter-5-using-redis-for-application-support/5-2-counters-and-statistics/5-2-1-storing-counters-in-redis/).
- The stats are retrieved from Redis.
- The project was deployed on App Engine Standard.
- [Post Mutant Sequence Diagram ⏱](diagrams/post-mutant-sequence-diagram.pdf)
- [Get Stats Sequence Diagram ⏱](diagrams/stats-mutants-sequence-diagram.pdf)
- [Internal Architecture Diagram 🏠](diagrams/internal-architecture.pdf)
- [Cloud Architecture Diagram 🏛](diagrams/cloud-architecture.pdf)

## Prerequisites

to run this project just you need:

* Docker and docker-compose
* Go 1.15*

the reason, this project uses MySql, Redis and Go, with docker and docker-compose you don't need manual installations

* Go is necessary in case you want to run the tests

### Run the project with docker

```bash
docker-compose up
```

### How to run the tests

Note: You need Go and the dependencies of the project.

TODO: run tests with docker exec

```console
go test ./... -v
```

## Overview
Mutants is a service for looking for if a human either is a mutant or not. This is a Rest project, is not restful.

### Test yourself !

You could test either in local or directly in production.

`Api Production URL: https://mutants-fer.uc.r.appspot.com/`
*Service Off*

Just use the Api URL instead of localhost in the below examples.

### How to know if a human is a mutant ?

### Post Mutant

`POST http://localhost:5007/api/v1/mutant`

**Auth required** : NO

**Request** POST Example

This is a DNA of a Mutant because it has at least 1 sequence of 4 equal letters.
These sequences can be horizontally, vertically or in the diagonal.

```json
{
  "dna": [
    "ATGCGA",
    "CAGTGC",
    "TTATGT",
    "AGAAGG",
    "CCCCTA",
    "TCACTG"
  ]
}
```

**Response** HttpStatus If the human is a mutant

```json
200
```

**Response** HttpStatus If the human is a real human

```json
403
```

### How to know how many humans/mutants are there ?

### Get Stats

`GET http://localhost:5007/api/v1/stats`

**Auth required** : NO

**Request** No Body

**Response** Number of humans and mutants.

```json
{
  "count_mutant_dna": 2,
  "count_human_dna": 3,
  "ratio": 0.6666667
}
```

## Some considerations

If you want to re-generate the mocks. 
[See Dependency](https://github.com/vektra/mockery)

```console
mockery --all --keeptree
```

If you want to deploy.

```console
gcloud app deploy
```

Rename the `app.yaml.example` to `app.yaml` and fill out the environment variables.
