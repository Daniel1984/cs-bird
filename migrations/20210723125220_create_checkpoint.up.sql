CREATE TABLE IF NOT EXISTS public.checkpoints
(
    time TIMESTAMPTZ NOT NULL,
    coin TEXT NOT NULL,
    address TEXT NOT NULL,
    balance TEXT NOT NULL,
    nonce INT
);

SELECT create_hypertable('public.checkpoints', 'time');
