-- JSON — Menyimpan Log Aktivitas
INSERT INTO activity_log (staff_id, action, details)
VALUES (
  1,
  'add_menu',
  '{"menu_name": "Latte", "price": 35000, "category": "Coffee"}'
);
