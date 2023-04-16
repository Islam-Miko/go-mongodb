# REST API with Golang and MongoDB

This is a repository for a REST API written in Go language that uses MongoDB as the database. The API includes base authentication endpoints.

Prerequisites

Before you can run the REST API, you need to have the following tools installed on your machine:

- Golang
- Docker


API Endpoints

The REST API includes the following base authentication endpoints:

    POST api/v1/auth/register: Registers a new user
    POST api/v1/auth/login: Logs in a user
    GET api/v1/auth/logout: Logs out a user
    GET api/v1/auth/refresh: Refresh access token
    GET api/v1/users/me: Get user info
