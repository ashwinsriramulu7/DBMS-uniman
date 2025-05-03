CREATE TABLE `class_schedule` (
  `id` int NOT NULL AUTO_INCREMENT,
  `course_id` int DEFAULT NULL,
  `faculty_id` int DEFAULT NULL,
  `day_of_week` enum('Monday','Tuesday','Wednesday','Thursday','Friday','Saturday') DEFAULT NULL,
  `start_time` time DEFAULT NULL,
  `end_time` time DEFAULT NULL,
  `location` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `course_id` (`course_id`),
  KEY `faculty_id` (`faculty_id`),
  CONSTRAINT `class_schedule_ibfk_1` FOREIGN KEY (`course_id`) REFERENCES `course` (`id`),
  CONSTRAINT `class_schedule_ibfk_2` FOREIGN KEY (`faculty_id`) REFERENCES `faculty` (`id`)
);

CREATE TABLE `college` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `location` varchar(100) DEFAULT NULL,
  `estd` year DEFAULT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `course` (
  `id` int NOT NULL AUTO_INCREMENT,
  `department_id` int NOT NULL,
  `course_code` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `course_code` (`course_code`),
  KEY `department_id` (`department_id`),
  CONSTRAINT `course_ibfk_1` FOREIGN KEY (`department_id`) REFERENCES `department` (`id`)
);

CREATE TABLE `department` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `head_of_department` int DEFAULT NULL,
  `college` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `head_of_department` (`head_of_department`),
  KEY `college` (`college`),
  CONSTRAINT `department_ibfk_1` FOREIGN KEY (`head_of_department`) REFERENCES `faculty` (`id`),
  CONSTRAINT `department_ibfk_2` FOREIGN KEY (`college`) REFERENCES `college` (`id`)
);

CREATE TABLE `enrollment` (
  `student_id` int NOT NULL,
  `course_id` int NOT NULL,
  `semester` varchar(10) NOT NULL,
  `year` year NOT NULL,
  `grade` varchar(2) DEFAULT NULL,
  PRIMARY KEY (`student_id`,`course_id`,`semester`,`year`),
  KEY `course_id` (`course_id`),
  CONSTRAINT `enrollment_ibfk_1` FOREIGN KEY (`student_id`) REFERENCES `student` (`id`),
  CONSTRAINT `enrollment_ibfk_2` FOREIGN KEY (`course_id`) REFERENCES `course` (`id`)
);

CREATE TABLE `exam` (
  `id` int NOT NULL AUTO_INCREMENT,
  `course_id` int DEFAULT NULL,
  `exam_type` enum('MIDTERM','FINAL','QUIZ','PROJECT') DEFAULT NULL,
  `date` date DEFAULT NULL,
  `total_marks` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `course_id` (`course_id`),
  CONSTRAINT `exam_ibfk_1` FOREIGN KEY (`course_id`) REFERENCES `course` (`id`)
);

CREATE TABLE `exam_result` (
  `exam_id` int NOT NULL,
  `student_id` int NOT NULL,
  `marks_obtained` int DEFAULT NULL,
  PRIMARY KEY (`exam_id`,`student_id`),
  KEY `student_id` (`student_id`),
  CONSTRAINT `exam_result_ibfk_1` FOREIGN KEY (`exam_id`) REFERENCES `exam` (`id`),
  CONSTRAINT `exam_result_ibfk_2` FOREIGN KEY (`student_id`) REFERENCES `student` (`id`)
);

CREATE TABLE `faculty` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `mobile_number` varchar(20) NOT NULL,
  `email` varchar(255) NOT NULL,
  `address` text,
  `type` enum('ACADEMIC','NON ACADEMIC') DEFAULT NULL,
  `title` varchar(300) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `mobile_number` (`mobile_number`),
  UNIQUE KEY `email` (`email`)
);

CREATE TABLE `fee_payment` (
  `id` int NOT NULL AUTO_INCREMENT,
  `student_id` int DEFAULT NULL,
  `amount` decimal(10,2) DEFAULT NULL,
  `date_paid` date DEFAULT NULL,
  `payment_mode` enum('CASH','CARD','ONLINE') DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `student_id` (`student_id`),
  CONSTRAINT `fee_payment_ibfk_1` FOREIGN KEY (`student_id`) REFERENCES `student` (`id`)
);

CREATE TABLE `program` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `level` enum('UG','PG','PhD') DEFAULT NULL,
  `department_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `department_id` (`department_id`),
  CONSTRAINT `program_ibfk_1` FOREIGN KEY (`department_id`) REFERENCES `department` (`id`)
);

CREATE TABLE `student` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `mobile_number` varchar(20) NOT NULL,
  `email` varchar(255) NOT NULL,
  `program_enrolled` varchar(300) DEFAULT NULL,
  `type` enum('UG','PG','Phd') DEFAULT NULL,
  PRIMARY KEY (`id`)
);

