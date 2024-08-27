CREATE TABLE IF NOT EXISTS notifications  (
    id uuid PRIMARY KEY,
    kind smallint NOT NULL DEFAULT 0,
    type text NULL,
    title text NOT NULL,
    content text NOT NULL,
    images jsonb NOT NULL DEFAULT '{}'::jsonb,
    receiver uuid[] NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now()
  );