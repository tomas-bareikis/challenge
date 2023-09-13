# Beat the IT exam of Lithuania

Have you ever wondered if you could pass `IT` National-level Maturity Exam of Lithuania?

We grabbed one of the programming tasks from the latest exam
and tweaked it a bit by adding just little bit more of a challenge.

## How

 * Read the [rules](requirements/rules.md) and [requirements](requirements/requirements.md).
 * [Fork this repo](github.com/tomas-bareikis/challenge/fork).
 * Write code in you favorite language (see [`go`](go/main.go), [`js`](js/main.js), [`php`](php/main.php) folders).
 * Create a Pull request and see automated tests check your results.

## Running tests locally

Assuming you have [`docker`](https://docs.docker.com/engine/install/) installed.

### PHP 8

```shell
docker run -v ${PWD}/php:/opt/php aurelijusbanelis/challenge:php
```
Where your code is in `php` folder with `php/main.php` file.

### JavaScript (Node 18)

```shell
docker run -v ${PWD}/js:/opt/js aurelijusbanelis/challenge:js
```
Where your code is in `js` folder with `js/main.js` file.

### Go 1.21

```shell
docker run -v ${PWD}/go:/opt/go aurelijusbanelis/challenge:go
```
Where your code is in `go` folder with `go/main.go` file.

### C# (dotnet 5)

```shell
docker run -v ${PWD}/cs:/opt/cs aurelijusbanelis/challenge:cs
```
Where your code is in `cs` folder with `cs/Program.cs` file.

### Shell

```shell
docker run -v ${PWD}/sh:/opt/sh aurelijusbanelis/challenge:sh
```
Where your code is in `sh` folder with `sh/main.sh` file.


### Not trusting docker?

Build it yourself:
```shell
docker build --target go . -t aurelijusbanelis/challenge:go
docker build --target php . -t aurelijusbanelis/challenge:php
docker build --target js . -t aurelijusbanelis/challenge:js
docker build --target cs . -t aurelijusbanelis/challenge:cs
docker build --target sh . -t aurelijusbanelis/challenge:sh
```

See [Dockerfile](Dockerfile) and [GitHub Actions](.github/workflows/infrastructure.yml) for details.

## References

* https://www.nsa.smm.lt/wp-content/uploads/2023/06/2023_IT_VBE_pg-web.pdf - original exam example
