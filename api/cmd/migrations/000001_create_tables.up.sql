-- Create stocks table with explicit types for CockroachDB compatibility
CREATE TABLE IF NOT EXISTS stocks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ticker VARCHAR(10) NOT NULL,
    company VARCHAR(255) NOT NULL,
    brokerage VARCHAR(255) NOT NULL,
    action VARCHAR(50) NOT NULL,
    rating_from VARCHAR(50),
    rating_to VARCHAR(50),
    target_from DECIMAL(10, 2),
    target_to DECIMAL(10, 2),
    time TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create index for faster lookups by ticker and time
CREATE UNIQUE INDEX IF NOT EXISTS idx_stocks_ticker_time ON stocks(ticker, time);