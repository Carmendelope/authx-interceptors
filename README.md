# authx-interceptor

The authx-interceptor repository contains gRPC interceptors used throughout the Nalej platform
in order to validate incoming requests into the system.

## Getting Started

This repository makes use of the client gRPC interface of authx in order to send the required
requests to validate different types of token and/or token information.

The existing interceptors in this package are:

* **apikey**: The API Key gRPC interceptor expects an API Key as the method to authorize the request,
and connects to a generic backend for the API key check implementation. Components using this
interceptor are expected to provide a structure implementing the `APIKeyAccess` interface.

```go
	cfg := config.NewConfig(&config.AuthorizationConfig{
		AllowsAll: false,
		Permissions: map[string]config.Permission{
			"/ping.Ping/Ping": {Must: []string{KeyAccessPrimitive}},
		}}, "globalSecret", "authorization")

	tokenProvider = APIKeyAccessImpl()

	grpcServer = grpc.NewServer(WithAPIKeyInterceptor(tokenProvider, cfg))
```

### Prerequisites

The following components are required for the interceptors to work

* [authx](https://github.com/nalej/authx)

### Build and compile

In order to build and compile this repository use the provided Makefile:

```
make all
```

This operation generates the binaries for this repo, download dependencies,
run existing tests and generate ready-to-deploy Kubernetes files.

### Run tests

Tests are executed using Ginkgo. To run all the available tests:

```
make test
```

### Update dependencies

Dependencies are managed using Godep. For an automatic dependencies download use:

```
make dep
```

In order to have all dependencies up-to-date run:

```
dep ensure -update -v
```

## Known Issues

* Consider this repository as WIP, some of the interceptors have not been migrated to the
repository but they will be in future releases.


## Contributing

Please read [contributing.md](contributing.md) for details on our code of conduct, and the process for submitting pull requests to us.


## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/nalej/authx-interceptors/tags). 

## Authors

See also the list of [contributors](https://github.com/nalej/authx-interceptors/contributors) who participated in this project.

## License
This project is licensed under the Apache 2.0 License - see the [LICENSE-2.0.txt](LICENSE-2.0.txt) file for details.