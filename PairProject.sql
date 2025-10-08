-- ============================================
-- BEVERAGE CLI PROJECT - POSTGRESQL SCHEMA
-- ============================================

CREATE DATABASE pairproject;
\c pairproject

CREATE TABLE staffs (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100),
  email VARCHAR(100) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  role VARCHAR(50),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE staff_profiles (
  id SERIAL PRIMARY KEY,
  staff_id INT UNIQUE NOT NULL,
  address TEXT,
  phone VARCHAR(20),
  date_of_birth DATE,
  emergency_contact VARCHAR(100),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_staff_profile_staff
    FOREIGN KEY (staff_id) REFERENCES staffs (id)
);

CREATE TABLE menus (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  category VARCHAR(50),
  price DECIMAL(10,2),
  stock INT DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE members (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100),
  phone VARCHAR(20),
  email VARCHAR(100),
  reward_points INT DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE rewards (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100),
  description TEXT,
  required_points INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE member_rewards (
  id SERIAL PRIMARY KEY,
  member_id INT NOT NULL,
  reward_id INT NOT NULL,
  redeemed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  status BOOLEAN DEFAULT FALSE,
  CONSTRAINT fk_member_rewards_member
    FOREIGN KEY (member_id) REFERENCES members (id),
  CONSTRAINT fk_member_rewards_reward
    FOREIGN KEY (reward_id) REFERENCES rewards (id)
);

CREATE TABLE reservations (
  id SERIAL PRIMARY KEY,
  member_id INT NOT NULL,
  reservation_date DATE,
  time_slot VARCHAR(20),
  total_people INT,
  note TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_reservation_member
    FOREIGN KEY (member_id) REFERENCES members (id)
);

CREATE TABLE transactions (
  id SERIAL PRIMARY KEY,
  staff_id INT NOT NULL,
  member_id INT,
  transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  total_amount DECIMAL(10,2),
  CONSTRAINT fk_transaction_staff
    FOREIGN KEY (staff_id) REFERENCES staffs (id),
  CONSTRAINT fk_transaction_member
    FOREIGN KEY (member_id) REFERENCES members (id)
);

CREATE TABLE transaction_items (
  id SERIAL PRIMARY KEY,
  transaction_id INT NOT NULL,
  menu_id INT NOT NULL,
  quantity INT DEFAULT 1,
  subtotal DECIMAL(10,2),
  CONSTRAINT fk_transaction_item_transaction
    FOREIGN KEY (transaction_id) REFERENCES transactions (id),
  CONSTRAINT fk_transaction_item_menu
    FOREIGN KEY (menu_id) REFERENCES menus (id)
);

CREATE TABLE activity_log (
  id SERIAL PRIMARY KEY,
  staff_id INT NOT NULL,
  action VARCHAR(50),
  details JSON,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_activity_log_staff
    FOREIGN KEY (staff_id) REFERENCES staffs (id)
);
