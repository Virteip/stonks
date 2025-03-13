# Stonks UI

Frontend application for the Stonks project, built with Vue 3, TypeScript, Pinia, and Tailwind CSS.

## Requirements

- Node.js 18+
- npm or yarn

## Setup

1. Install dependencies:
```bash
npm install
# or
yarn
```

2. Configure environment variables:
Create a `.env` file with:
```
VITE_API_BASE_URL=http://localhost:8080/api/v1/stonks-api
VITE_API_KEY=your_api_key_here
```
And export before startup with this command in your terminal
```
$ export $(cat .env.local | grep -v ^# | xargs)
```

## Running the Application

```bash
npm run dev
# or
yarn dev
```

The application will be available at `http://localhost:3000`.

## Building for Production

```bash
npm run build
# or
yarn build
```

## Features

- View stocks with pagination
- Search stocks by ticker
- View stock recommendations based on advanced analysis

## Views

- **Stocks View**: Browse and search for stocks
- **Recommendations View**: Get recommended stocks for investment

Sergio Pietri