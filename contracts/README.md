## Consumer driven tests with pack

Alternative to integration testing Tests only integration points instead of the whole system together. Use case:
Services 1,2,...,n depend on an API X for data resources via an endpoint. If API X changes it's endpoint response with a
breaking changes, dependent services will fail. There is no way for us to know all those services that depend on our API
X. With consumer driven contracts, the producer API has contracts(PACT) with all consumers, and a graph that shows all
dependencies is created in [pack broker server](http://localhost/?tags=true)

### Flow of a consumer contract test

### Why not integration tests

Currently, before we deploy an app to prod, the app has to first have successfully connected to the external apps with
which it integrates. i.e. we connect to live deployed apps

### pros

more confidence with prod deploys

### cons

1. creates dependencies
1. slow running process, sometimes with http polling required
1. flakiness if external systems fail/slow
1. costs running tests which invoke AWS resources like lambdas

### Consumer driven testing with Pact

NB: Pactflow is a commercial version of our Pact Broker

* https://github.com/pact-foundation/pact-go#installation
* https://www.rea-group.com/blog/enter-the-pact-matrix-or-how-to-decouple-the-release-cycles-of-your-microservices/
* https://docs.pact.io/pact_broker/can_i_deploy/

### Run own broker. Recommended using a hosted Broker

### Hosted pact broker

Company : ZagadaT Email:    ryumugil@gmail.com Subdomain for broker: zagadat.pactflow.io Pass: G..g..2020*