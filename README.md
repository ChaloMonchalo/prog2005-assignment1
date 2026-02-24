# Assignment 1 --- Augmented Country Information Service

## Overview

This project implements a RESTful web service written in Go.

The service provides:

-   General country information
-   Currency exchange rates for neighbouring countries
-   Diagnostics information about upstream services

The service does not store any data locally. All information is
retrieved in real time from external APIs, processed, and returned as
structured JSON responses.

------------------------------------------------------------------------

## External Services

The following self-hosted upstream services are used:

**REST Countries API**\
http://129.241.150.113:8080/v3.1/

**Currency API**\
http://129.241.150.113:9090/currency/

------------------------------------------------------------------------

## Base Path

All endpoints start with:

/countryinfo/v1/

------------------------------------------------------------------------

## API Endpoints

### 1. Status

GET /countryinfo/v1/status/

Returns:

-   HTTP status codes of upstream services
-   Service version
-   Uptime (in seconds)

------------------------------------------------------------------------

### 2. Country Info

GET /countryinfo/v1/info/{two_letter_country_code}

Example:

/countryinfo/v1/info/no

Returns:

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

If:

-   An invalid country code is provided
-   An upstream service fails
-   A request method is not allowed

The service returns an appropriate HTTP status code together with a JSON
error message.

------------------------------------------------------------------------

## Running Locally

### Requirements

-   Go 1.22 or higher

### Run

go run .

The server runs on:

-   Port 8080 by default
-   Or the port defined by the PORT environment variable

Local base URL:

http://localhost:8080

------------------------------------------------------------------------

## Deployment

The service is deployed on Render.

Render URL:

https://prog2005-assignment1-8209.onrender.com

Example deployed endpoints:

https://prog2005-assignment1-8209.onrender.com/countryinfo/v1/status/
https://prog2005-assignment1-8209.onrender.com/countryinfo/v1/info/no
https://prog2005-assignment1-8209.onrender.com/countryinfo/v1/exchange/no

------------------------------------------------------------------------

## AI Usage

AI was used as a supporting tool to clarify parts of the assignment
description and improve English phrasing in documentation.
