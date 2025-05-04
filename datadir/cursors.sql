DELIMITER $$

CREATE PROCEDURE ListStudentFees()
BEGIN
  DECLARE done INT DEFAULT FALSE;
  DECLARE s_id INT;
  DECLARE s_name VARCHAR(100);
  DECLARE total_fee DECIMAL(10,2);
  DECLARE cur CURSOR FOR 
    SELECT s.id, s.name FROM student s;

  DECLARE CONTINUE HANDLER FOR NOT FOUND SET done = TRUE;

  OPEN cur;
  read_loop: LOOP
    FETCH cur INTO s_id, s_name;
    IF done THEN
      LEAVE read_loop;
    END IF;

    SELECT SUM(amount) INTO total_fee
    FROM fee_payment
    WHERE student_id = s_id;

    SELECT CONCAT('Student: ', s_name, ' | Total Fee Paid: ', IFNULL(total_fee, 0)) AS Result;
  END LOOP;

  CLOSE cur;
END$$

DELIMITER ;

