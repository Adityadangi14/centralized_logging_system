-- Table: parsed_logs
CREATE TABLE IF NOT EXISTS parsed_logs (
    id SERIAL PRIMARY KEY,
    timestamp TIMESTAMPTZ NOT NULL,
    event_category TEXT NOT NULL,
    source_type TEXT NOT NULL,
    username TEXT NOT NULL,
    hostname TEXT NOT NULL,
    severity TEXT NOT NULL,
    raw_message TEXT NOT NULL
    is_blacklisted BOOLEAN NOT NULL
);
