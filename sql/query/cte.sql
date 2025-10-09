-- CTE - TOTAL PENJUALAN PER STAFF
WITH staff_sales AS (
  SELECT 
    s.id AS staff_id,
    s.name AS staff_name,
    SUM(t.total_amount) AS total_sales
  FROM staffs s
  JOIN transactions t ON t.staff_id = s.id
  GROUP BY s.id, s.name
)
SELECT * FROM staff_sales ORDER BY total_sales DESC;