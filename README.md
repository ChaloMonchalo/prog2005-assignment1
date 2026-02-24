# Assignment 1 -- Augmented Country Information Service

## Overview

This project implements a REST web service written in Go.

The service provides:

-   General country information
-   Currency exchange rates for neighbouring countries
-   Diagnostics information about upstream services

The service does not store any data. It retrieves data in real time from
external APIs, processes the information, and returns structured JSON
responses.

------------------------------------------------------------------------

## External Services

REST Countries API\
http://129.241.150.113:8080/v3.1/

Currency API\
http://129.241.150.113:9090/currency/

------------------------------------------------------------------------

## API Endpoints

All endpoints start with:

    /countryinfo/v1/

------------------------------------------------------------------------

### 1. Status

    GET /countryinfo/v1/status/

Returns:

-   HTTP status codes of upstream services
-   Service version
-   Uptime in seconds

------------------------------------------------------------------------

### 2. Country Info

    GET /countryinfo/v1/info/{two_letter_country_code}

Example:

    /countryinfo/v1/info/no

Returns general country information including:

-   name
-   continents
-   population
-   area
-   languages
-   borders
-   flag
-   capital

------------------------------------------------------------------------

### 3. Exchange Rates

    GET /countryinfo/v1/exchange/{two_letter_country_code}

Example:

    /countryinfo/v1/exchange/no

Returns:

-   country name
-   base currency
-   exchange rates for neighbouring countries

------------------------------------------------------------------------

## Error Handling

If an invalid country code is provided or an upstream service fails, the
service returns an appropriate HTTP status code and a JSON error
message.

------------------------------------------------------------------------

## Running Locally

Requirements:

-   Go 1.22 or higher

Run the service with:

    go run .

The server runs on port 8080 by default, or on the port defined by the
`PORT` environment variable.

Local base URL:

    http://localhost:8080

------------------------------------------------------------------------

## Deployment

The service is deployed on Render.

Render URL:

`<insert deployed URL here>`
