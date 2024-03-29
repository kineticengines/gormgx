# Gormgx

<a href="https://github.com/kineticengines/gormgx/actions?query=workflow%3ATests+branch%3Amain">
    <img src="https://github.com/kineticengines/gormgx/workflows/Tests/badge.svg" alt="Test Status" />
</a>

<a href="https://badgen.net/github/watchers/kineticengines/gormgx">
    <img src="https://badgen.net/github/watchers/kineticengines/gormgx" alt="Watchers"/>
</a>

<a href="https://badgen.net/github/branches/kineticengines/gormgx">
    <img src="https://badgen.net/github/branches/kineticengines/gormgx" alt="Branches" />
</a>

<a href="https://badgen.net/github/issues/kineticengines/gormgx">
    <img src="https://badgen.net/github/issues/kineticengines/gormgx" alt="Issues" />
</a>

<a href="https://badgen.net/github/open-issues/kineticengines/gormgx">
    <img src="https://badgen.net/github/open-issues/kineticengines/gormgx" alt="Open Issues" />
</a>

<a href="https://badgen.net/github/closed-issues/kineticengines/gormgx">
    <img src="https://badgen.net/github/closed-issues/kineticengines/gormgx" alt="Closed Issues" />
</a>

<a href="https://badgen.net/github/license/kineticengines/gormgx">
    <img src="https://badgen.net/github/license/kineticengines/gormgx" alt="License" />
</a>

[![Go Report Card](https://goreportcard.com/badge/github.com/kineticengines/gormgx)](https://goreportcard.com/report/github.com/kineticengines/gormgx)

<br>

### A simple database migrations utility for Gorm

Active development in progress

## Example

```sh
gormgx.go -v make-migrations
```

## Gotchas

- Models specifications must be in one (1) go file. Preferably `models.go`
- Override Foreign Key must be of the form `ModelNameRefer`.
- Foreign key must be `Interface{}`. Gormgx will extract the extact model from the tags

## Environment

```sh
export DATABASE_DSN="host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Africa/Nairobi"
```

## Notice

Gormgx is still in development and not ready for productive use yet
