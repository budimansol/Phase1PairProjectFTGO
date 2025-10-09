-- ROLLUP - LAPORAN PENJUALAN PER KATEGORI DAN TOTAL AKHIR
SELECT 
  category,
  SUM(ti.subtotal) AS total_sales
FROM transaction_items ti
JOIN menus m ON ti.menu_id = m.id
GROUP BY ROLLUP(category);