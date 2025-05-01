CREATE TABLE IF NOT EXISTS `college`(
	`id` INT AUTO_INCREMENT PRIMARY KEY,
	`name` VARCHAR(100) NOT NULL,
	`location` VARCHAR(100),
	`estd.` YEAR
);
CREATE TABLE IF NOT EXISTS `faculty`(
	`id` INT AUTO_INCREMENT PRIMARY KEY,
	`name` VARCHAR(100) NOT NULL,
	`mobile_number` VARCHAR(20) UNIQUE NOT NULL,
	`email` VARCHAR(255) UNIQUE NOT NULL,
	`address` TEXT,
	`type` ENUM('ACADEMIC', 'NON ACADEMIC'),
	`title` VARCHAR(300)
);
CREATE TABLE IF NOT EXISTS `student`(
	`id` INT AUTO_INCREMENT PRIMARY KEY,
	`name` VARCHAR(100) NOT NULL,
	`mobile_number` VARCHAR(20) UNIQUE NOT NULL,
	`email` VARCHAR(255) UNIQUE NOT NULL,
	`program_enrolled` VARCHAR(300),
	`type` ENUM('UG', 'PG', 'Phd')

);
CREATE TABLE IF NOT EXISTS `department`(
	`id` INT AUTO_INCREMENT PRIMARY KEY,
	`name` VARCHAR(100) NOT NULL,
	`head_of_department` INT,
	FOREIGN KEY(head_of_department) REFERENCES faculty(id),
	`college` INT,
	FOREIGN KEY(college) REFERENCES college(id)
);
CREATE TABLE IF NOT EXISTS `course`(
	`id` INT AUTO_INCREMENT PRIMARY KEY,
	`department_id` INT NOT NULL,
	`course_code` VARCHAR(10) UNIQUE,
	FOREIGN KEY (department_id) references department(id)
);

