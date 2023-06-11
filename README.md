# Clean Architecture Practice Repository (Gin + Go)
This repository serves as a practice project for implementing the Clean Architecture using the Gin web framework and the Go programming language. It demonstrates how to structure a web application using the principles of Clean Architecture to achieve separation of concerns and maintainability.

## Architecture Overview
The project follows the Clean Architecture principles, which consist of the following layers:
1. Domain Layer: Contains the core business logic and entities of the application. It should be independent of any external frameworks or libraries.
2. Use Case Layer: Defines the application-specific use cases or application services. It orchestrates the interactions between the Domain Layer and the Interface Adapters.
3. Interface Adapters: Adapts the application to the external world, such as the web framework or database. It includes the delivery mechanisms (controllers, routers) and data access implementations.
4. Frameworks & Drivers: Contains the external frameworks, libraries, and tools used in the application. In this project, we use the Gin web framework as the HTTP delivery mechanism.


## Project Structure
The repository follows a standard Go project structure with the following directories:
```
.
├── di
│   └── ...
├── domain
│   └── ...
├── driver
│   └── ...
├── gateway
│   └── ...
├── rest
│   ├── handlers
│   │   └── ...
│   └── ...
└── usecase
    └── ...
```
- di: Contains the dependency injection setup and configuration.
- domain: Defines the core business logic and entities of the application.
- driver: Contains the implementation of external drivers or frameworks.
- gateway: Implements the data access layer (repositories) that interact with the database or external services.
- rest: Contains the delivery mechanism (controllers, routers) that adapt the application to the web framework.
- rest/handlers: Includes the handler functions for the web endpoints.
- usecase: Implements the application-specific use cases or services.
