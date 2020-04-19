DELIMITER $$

DROP PROCEDURE IF EXISTS `db_blog`.`nextID` $$
CREATE PROCEDURE `db_blog`.`nextID` (IN field varchar(255))
BEGIN
	DECLARE counter INT;
    
	SELECT current_value FROM `db_blog`.`counters` WHERE counter_name = field INTO counter;
    
    IF counter IS NULL THEN
		INSERT INTO `db_blog`.`counters` VALUES(field, 1);
        SET counter = 1;
	ELSE 
		SET counter = counter + 1;
		UPDATE `db_blog`.`counters` SET current_value = counter WHERE counter_name = field;
    END IF;
    
    SELECT counter;
	
END $$

DELIMITER ;