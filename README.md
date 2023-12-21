# go-serverless-template

Template for Golang serverless applications.

Deploys an [AWS Lambda application](https://docs.aws.amazon.com/lambda/latest/dg/deploying-lambda-apps.html) by utilizing [Serverless Framework](https://www.serverless.com). Env file (i.e., `.env`) is used for configuration.

Since this tries to follow the [3 Musketeers](https://3musketeersdev.netlify.app) methodology, the following are heavily used:

- Docker
- Docker Compose
- Make

## Usage

#### configure

```bash
$ make .env
```

- see generated `.env` file for configuration

**Note:** For deployment via CI/CD, `CICD_MODE` environment variable should be set to `true` in the build server. All configuration should also be set in the build server's environment variables. They will automatically be used if `.env` is generated from `.env.cicd`.

#### tidy dependencies

```bash
$ make deps
```

#### run tests

```bash
$ make test
```

#### build serverless functions

```bash
$ make build
```

- this generates `bin` directory to be used in deployment

#### deploy serverless application

```bash
$ make deploy
```

### Helpers during development:

#### format all .go files in project (using go fmt)

```bash
$ make fmt
```

#### generate test mocks (to be used with [stretchr/testify](https://github.com/stretchr/testify)) for all interfaces in project

```bash
$ make mocks
```

- can be configured in `.mockery.yaml`
