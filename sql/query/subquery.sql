-- SUBQUERY - MEMBER DIATAS RATA-RATA TRANSAKSI
SELECT m.id, m.name
FROM members m
WHERE (
  SELECT AVG(total_amount)
  FROM transactions t
  WHERE t.member_id = m.id
) > (
  SELECT AVG(total_amount) FROM transactions
);