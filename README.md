# Microbank

This is a fake bank service that is being built for me to learn micro serivces and modular monolith.

This service will consist of three micro services.

- main
  - works as an api gateway
  - also contains other services(modules)
    - user service
      - manages CRUD for user information
- account
  - responsible of money transaction between accounts.

tech stacks
- golang
- grpc
- envoy
- sqs

