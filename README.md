# Jobico

[![Go Report Card](https://goreportcard.com/badge/github.com/andrescosta/jobico)](https://goreportcard.com/report/github.com/andrescosta/jobico)

## Introduction

Welcome to Jobico – An Experimental, Multi-Tenancy asynchronous computing Platform!

Jobico is a Go project designed for experimental development, with a focus on exploring asynchronous computing using WebAssembly (WASM) in a multi-tenancy environment. It prioritizes flexibility and customization, allowing tenants to define events, validate incoming data using JSON Schema, and execute custom programs in any WASM-compatible language.

## Key Characteristics

- **Exploratory Nature**: Jobico serves as an exploratory project, providing a platform for investigating innovative approaches to asynchronous computing technologies.

- **Multi-Tenancy Focus**: Jobico's architecture is specifically designed to facilitate multi-tenancy, enabling the simultaneous operation of multiple isolated tenants on the platform.

- **Event Definition with JSON Schema**: Tenants can define events through JSON Schema, allowing for structured and dynamic event handling. Incoming requests undergo validation against the specified schema.

- **WASM-Compatible Language Support**: Jobico offers support for custom program creation in any WASM-compatible language, fostering flexibility and diversity in the execution of jobs.

# Platform Setup and Administration Guide

Explore how to work with the Jobico platform by checking out our comprehensive tutotial. Learn how to set up new jobs, upload wasm and schema files, and utilize administrative tools. If you're interested in diving deeper, click the link below to open the manual.

### [Learn more with our In-Depth Guide](/GUIDE.md)


# Architecture
![alt](docs/img/Jobico.svg?)

## Components

### Listener
The **Listener** component serves as the entry point for external events, providing a REST API that can function as a webhook. Its primary responsibilities include receiving events, validating them against pre-defined JSON schemas, and subsequently enqueueing them for further processing. This component acts as the bridge between external sources triggering events and the internal processing pipeline.

### Queue
The **Queue** component acts as a chronological buffer for events, temporarily storing them until they can be processed by the Job Executors. Events are maintained in the queue in the order they are received, ensuring a sequential flow of processing. This component plays a crucial role in decoupling the event reception from the actual event processing, allowing for scalability and efficient handling of bursts of incoming events.

### Job Executors
**Job Executors** are responsible for consuming events from the Queue and providing a controlled environment for the execution of WebAssembly (WASM) programs that process these events. This component manages the execution context, ensuring isolation and security for running custom WASM programs written by tenants. It plays a key role in the dynamic and scalable execution of programmed jobs in response to events.

### Control Service
The **Control Service** is a centralized hub where Job definitions are stored and can be queried by other components in the system. It acts as the authoritative source for managing job configurations, allowing dynamic adjustments to the processing logic without interrupting the overall system operation. This service facilitates coordination and control over the execution of jobs across the entire platform.

### Job Repository
The **Job Repository** serves as a storage facility for WebAssembly (WASM) programs and JSON schema files. It provides a dedicated API for storing and retrieving these essential files, ensuring accessibility for the Job Executors and enabling tenants to manage their custom program logic efficiently. This component acts as a repository for the building blocks required for event processing.

### Executions Recorder
The **Executions Recorder** is a service designed to capture and store log information generated by the Job Executors during the execution of jobs. This component acts as a centralized logging system, allowing for post-execution analysis, troubleshooting, and performance monitoring. The recorded information can be queried using both the Command Line Interface (CLI) and Dashboard tools, providing visibility into the historical execution details of jobs.

## Stack

![alt](docs/img/stack.svg?)


## Goico: The Jobico Framework

### Overview:

**Goico** is a specialized framework  crafted to support the development and evolution of Jobico. 

### Key Features:

#### 1. Service Creation and API Exposure:

Goico simplifies the development of microservices within the Jobico ecosystem. It supports creating services that expose REST or gRPC APIs, fostering a modular and scalable architecture. 

#### 2. WASM Runtime Based on WAZERO:

A core strength of Goico lies in its WebAssembly (WASM) runtime, built on the robust foundation of WAZERO. This runtime facilitates the execution of custom logic written in any WASM-supported programming language. 

#### 3. Key/Value Embedded Database:

Goico integrates an embedded database based on [Pebble](https://github.com/cockroachdb/pebble), offering a key/value store for data management. This embedded database serves as the backbone for storing critical information, supporting the reliable and fast retrieval of data essential for the operation of Jobico.

#### 4. Streaming Capabilities for Data Updates:

Goico provides streaming capabilities for database updates based on Grpc. This feature enables real-time monitoring and reaction to changes within the embedded database, facilitating dynamic adjustments, and enhancing the responsiveness of Jobico to evolving requirements.

## Data
[TODO]

## Caching and Streaming
[TODO]

# Management Tools

## CLI
The **Command Line Tool** serves as the primary management interface, providing a comprehensive set of commands to interact with the Jobico platform. This tool provides commands to deploy job definitions, upload associated WASM and JSON schema files, and query the Executions Recorder for log information.

## Dashboard
The **Dashboard** is a terminal-based application designed for visualizing Job definitions and execution information. The Dashboard provides a user-friendly experience to monitor and analyze the status of their jobs. 

# Getting Started with Jobico Development
To start developing follow these steps:

## Checkout from GitHub
Clone the Jobico repository from GitHub to your local machine:

``` bash
git clone https://github.com/andrescosta/jobico.git
```
## Running Docker Compose
Spin up a local Jobico environment using Docker Compose:
[TODO]
1. 
``` bash
make up
```

2. 
``` bash
make up
```

## Building
Ensure the codebase integrity by running the test cases:

### Release
[TODO]

``` bash
make release
```

### Local build

[TODO]

## Running Test Cases
Ensure the codebase integrity by running the test cases:

``` bash
cd Jobico
make test
```

##  Executing Performance tests

In this section, we'll demonstrate how to use k6 to conduct performance tests on Jobico. The following steps will guide you through the process of running these tests, allowing you to measure the platform's throughput, latency, and resource utilization under various conditions.

### 1. Compiling k6 bundled with extensions

``` bash
make k6
```
### 2. Deploy and start
[TODO]

#### 2.1. Local
[TODO]

#### 2.2. Docker
[TODO]

### 3. Testing sending events 

``` bash
make perf1
```

### 4. Testing sending events and streaming

``` bash
make perf2
```

# OpenTelemetry

# Configuration

[todo into]
## Enviroment variables

## .env files

https://github.com/bkeepers/dotenv#what-other-env-files-can-i-use

development
production
test

## Command line

## Parameters

### General
| Parameter | Description |
| --- | --- |
| workdir | |
| basedir ||
| APP_ENV ||
| [svc].addr | List all new or modified files |
| [svc].host | Show file differences that haven't been staged |
| [svc].dir||
| metadata.enabled | |
| grpc.healthcheck ||
| dial.timeout ||

### Profiling
| prof.enabled | List all new or modified files |
| pprof.addr | Show file differences that haven't been staged |

### Observability

#### OpenTelemetry
| Parameter | Description |
| --- | --- |
| obs.enabled | List all new or modified files |
| obs.exporter.trace.grpc.host | List all new or modified files |
| obs.exporter.metrics.http.host | List all new or modified files |
| obs.exporter.metrics.host.path ||
| obs.metrics.host | List all new or modified files |
| obs.metrics.runtime | List all new or modified files |

#### Logging
| Parameter | Description |
| --- | --- |
| log.level | List all new or modified files |
| log.caller ||
| log.console.enabled | List all new or modified files |
| log.console.exclude.timestamp ||
| log.file.enabled | List all new or modified files |
| log.file.name | List all new or modified files |
| log.file.max.size ||
| log.file.max.backups ||
| log.file.max.age ||

### Service specifics

#### Executor
| Parameter | Description |
| --- | --- |
|executor.delay||
|executor.timeout||
|executor.maxproc||

# Roadmap

The roadmap can be accessed or queried at this location:

https://github.com/users/andrescosta/projects/3/views/1


### Short Term
- Complete demo of a Jobicolet
- Extend the capabilities of the Testing framework
- Errors management

### Mid term
- Improvements to the Wasm runtime

### Long Term
- Distributed storage for the queue and control services
- Durable computing exploration

# Contact

For questions, feedback reach out to us at jobicowasm@gmail.com
