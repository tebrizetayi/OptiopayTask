# Tabriz-Atayi-coding-challenge

## Getting Started 


### 1. Run Tests

```bash
./bin/test.sh 
```

### 2. Run web API 

Navigate to the project root where `docker-compose.yml` exists. There is only 
one `docker-compose.yml` file. 

```bash
docker-compose up -d 
```
Use the `docker-compose.yml` to run all the services. The first time to run this 
command, Docker will download the required images.


### 3. FindCommonManager

e.x http://localhost:8080/findcommonmanager?id1=2&id2=3

## Structure of application

### 1. Corporate Library 

The "corporate" Go package provides functionality for finding the closest common manager of two employees in a corporate directory.
The package defines an Employee struct with an ID, name, and a slice of direct reports (other Employee structs).
The Corparate struct contains a pointer to the CEO Employee and implements the Directory interface, which includes a method for finding the closest common manager.
The package also defines an error variable, ErrEmployeeNotFound, for use when an employee is not found in the directory.
The NewCorporate function creates a new Corparate struct by getting employee data from a storage object and using it to create a directory hierarchy starting with the CEO.


### 4. Docker
Docker is a critical component and required to run this project.
Use the `docker-compose.yml` file to configure the services differently using environment variables when necessary. 


## What is missing

### 1. Testing 

Testing in domain and BlockChainClient layer is missing. And in http(api) layer, I've only added 2 tests cases with wrong NetworkCode.
Given the time, I would add using testing for each layer, and would make use of a FakeBlockChainClient during using testing.

### 2. Caching

Each time when we request transaction info and block info, we send a request to the client. It overloads our server with requests. To respond immediately, we can use caching. Caching is a technique used to create high-performance services. Cache and key-value store using Redis would be a nice option.

### 3. Logging

Datadog - Provides metrics, logging, and tracing.
It gathers system metrics, integrates with key software we use, and provides a standard interface to which our applications can send custom metrics. Datadog has prebuilt integrations to pull data from almost every important service we use. Through the integrations, datadog generates a consolidated event stream that we can filter and search as needed.
Datadog lets us build dashboards that combine metrics from many sources. We can combine and transform metrics to make them more useful. It also provides a powerful interface for the interactive exploration of metrics.

### 4. API Documentation

Swagger API documentation is a technical content deliverable, containing instructions about how to effectively use and integrate with an API. It’s a concise reference manual containing all the information required to work with the API, with details about the functions, classes, return types, arguments, and more, supported by tutorials and examples.

### 5. Scale

Because of our service is stateless, it can easily be scaled horizontally

### 6. Timeouts

Connection Timeout - The time it takes to open a network connection to the
server
Client Request Timeout - The time it takes for a server to process a request

The client request timeout is almost always going to be the longest duration of the two, and I 
recommend the timeout is defined in the configuration of the service. While you might
initially set it to an arbitrary value of, say 10 seconds, you can modify this after the system
has been running in production, and you have a decent data set of transaction times to look
at.

We can  use the deadline package from eapache (https://github.com/eapache/go-resiliency/tree/master/deadline)

### 7. Workflow automation

To Automate the workflow,GitHub Actions makes it easy to automate all your software workflows.

1.Control when the action will run. Triggers the workflow on push or pull request events but only for the master branch

2.Run testing on the code in predefined enviroments(Win, Unix, MacOS),with different golang versions.

3.If test succeed then, merge to main branch(master)

4.Send notification(ex. slack)

5.Deploy to docker hub.







