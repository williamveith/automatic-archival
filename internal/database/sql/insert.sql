INSERT INTO invoices (date, time, prof, user, tool, usage, rate, cost, applied_cost) 
VALUES %s
ON CONFLICT DO NOTHING