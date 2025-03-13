# Stonks App

A full-stack application that retrieves, displays, and analyzes stock information from an external API.

## Project Structure

- `/api` - Golang backend service with Echo framework and CockroachDB
- `/stonks-ui` - Vue 3 frontend with TypeScript, Pinia, and Tailwind CSS

## Quick Start

1. Start the backend API service in one terminal (see `/api/README.md`)
2. Start the frontend UI service in another terminal (see `/stonks-ui/README.md`)

Note: Each service must be started in separate terminal sessions.

## Features

- View and search stock data
- Get recommendations for the best stocks to invest in
- Store historical stock data in CockroachDB

## Tech Stack

- **Backend**: Golang with Echo framework
- **Frontend**: Vue 3, TypeScript, Pinia, Tailwind CSS
- **Database**: CockroachDB

Sergio Pietri