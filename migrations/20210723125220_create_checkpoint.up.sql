CREATE TABLE IF NOT EXISTS public.checkpoints
(
    time TIMESTAMPTZ NOT NULL,
    name TEXT NOT NULL,
    address TEXT NOT NULL,
    balance DECIMAL(22, 12)
);

SELECT create_hypertable('public.checkpoints', 'time');
