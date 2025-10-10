-- STAFFS
INSERT INTO staffs (name, email, password, role) VALUES
('Rina Wibowo', 'rina@mail.com', '123', 'Manager'),
('Budi Santoso', 'budi@mail.com', '123', 'Cashier'),
('Siti Aisyah', 'siti@mail.com', '123', 'Barista'),
('admin', 'admin@mail.com', '123', 'admin');

-- STAFF PROFILES (One-to-One)
INSERT INTO staff_profiles (staff_id, address, phone, date_of_birth, emergency_contact) VALUES
(1, 'Jakarta', '081234567890', '1988-05-15', '081111111111'),
(2, 'Bandung', '081234567891', '1992-08-20', '081222222222'),
(3, 'Surabaya', '081234567892', '1995-02-10', '081333333333');

-- MENUS
INSERT INTO menus (name, category, price, stock) VALUES
('Espresso', 'Coffee', 25000.00, 50),
('Latte', 'Coffee', 30000.00, 40),
('Green Tea', 'Tea', 20000.00, 30),
('Chocolate Frappe', 'Blended', 35000.00, 25),
('Croissant', 'Pastry', 18000.00, 20);

-- MEMBERS
INSERT INTO members (name, phone, email, reward_points) VALUES
('Andi Setiawan', '081234567100', 'andi@example.com', 120),
('Sinta Dewi', '081234567101', 'sinta@example.com', 80),
('Rafi Pratama', '081234567102', 'rafi@example.com', 200);

-- REWARDS
INSERT INTO rewards (name, description, required_points) VALUES
('Free Coffee', 'Get one free coffee of your choice', 100),
('20% Discount', 'Discount for total transaction above Rp100.000', 150),
('Free Pastry', 'Get a free pastry of your choice', 200);

-- MEMBER REWARDS (M:N)
INSERT INTO member_rewards (member_id, reward_id, status) VALUES
(1, 1, TRUE),
(2, 2, FALSE),
(3, 3, TRUE);

-- RESERVATIONS
INSERT INTO reservations (member_id, reservation_date, time_slot, total_people, note) VALUES
(1, '2025-10-10', '10:00-11:00', 2, 'Anniversary brunch'),
(2, '2025-10-12', '18:00-19:00', 4, 'Team dinner'),
(3, '2025-10-13', '15:00-16:00', 1, 'Solo reading time');

-- TRANSACTIONS
INSERT INTO transactions (staff_id, member_id, transaction_date, total_amount) VALUES
(2, 1, '2025-10-08 10:15:00', 55000.00),
(3, 2, '2025-10-08 14:30:00', 72000.00),
(2, 3, '2025-10-09 11:45:00', 88000.00);

-- TRANSACTION ITEMS
INSERT INTO transaction_items (transaction_id, menu_id, quantity, subtotal) VALUES
(1, 1, 1, 25000.00),
(1, 5, 1, 18000.00),
(1, 3, 1, 12000.00),
(2, 2, 2, 60000.00),
(3, 4, 2, 70000.00);

-- ACTIVITY LOG (JSON)
INSERT INTO activity_log (staff_id, action, details) VALUES
(1, 'CREATE_MENU', '{"menu_name": "Matcha Latte", "price": 32000}'),
(2, 'PROCESS_TRANSACTION', '{"transaction_id": 1, "amount": 55000}'),
(3, 'UPDATE_STOCK', '{"menu_id": 4, "new_stock": 23}');
