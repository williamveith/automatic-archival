PRAGMA journal_mode=WAL;

CREATE TABLE IF NOT EXISTS invoices (
  date TEXT NOT NULL,
  time TEXT NOT NULL,
  prof TEXT NOT NULL,
  user TEXT NOT NULL,
  tool TEXT NOT NULL,
  usage NUMERIC NOT NULL,
  rate NUMERIC NOT NULL,
  cost NUMERIC NOT NULL,
  applied NUMERIC NOT NULL,
  UNIQUE(date, time, prof, user, tool, usage, rate, cost, applied)
);