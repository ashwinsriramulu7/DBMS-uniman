DELIMITER $$

CREATE TRIGGER trg_check_fee_amount
BEFORE INSERT ON fee_payment
FOR EACH ROW
BEGIN
  IF NEW.amount IS NULL OR NEW.amount <= 0 THEN
    SIGNAL SQLSTATE '45000'
    SET MESSAGE_TEXT = 'Fee amount must be positive.';
  END IF;
END$$

DELIMITER ;

