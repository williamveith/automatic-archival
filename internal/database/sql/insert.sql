INSERT INTO invoices (date, time, prof, user, tool, usage, rate, cost, applied) 
VALUES %s
ON CONFLICT DO NOTHING