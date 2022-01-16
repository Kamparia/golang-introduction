CREATE TABLE IF NOT EXISTS visitors (
  id SERIAL CONSTRAINT visitors_pk PRIMARY KEY,
  user_agent varchar(255) NOT NULL,
  datetime timestamp with time zone NOT NULL
);

CREATE INDEX IF NOT EXISTS visitors_user_agent_idx ON visitors (user_agent);