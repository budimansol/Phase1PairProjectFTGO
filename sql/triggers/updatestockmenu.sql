-- Function untuk mengurangi stok menu setelah transaksi
CREATE OR REPLACE FUNCTION update_menu_stock()
RETURNS TRIGGER AS $$
BEGIN
  -- Cegah stok negatif
  IF (SELECT stock FROM menus WHERE id = NEW.menu_id) < NEW.quantity THEN
    RAISE EXCEPTION 'Stok menu tidak mencukupi untuk menu_id=%', NEW.menu_id;
  END IF;

  UPDATE menus
  SET stock = stock - NEW.quantity
  WHERE id = NEW.menu_id;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger untuk jalankan function setelah insert transaction_items
CREATE TRIGGER trg_update_stock
AFTER INSERT ON transaction_items
FOR EACH ROW
EXECUTE FUNCTION update_menu_stock();
