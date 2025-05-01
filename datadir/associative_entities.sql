CREATE TABLE IF NOT EXISTS `teaches`(
	faculty_id INT,
	course_id INT,
	semester VARCHAR(10),
	year YEAR,
	PRIMARY KEY (faculty_id, course_id, semester, year),
	FOREIGN KEY (faculty_id) REFERENCES faculty(id),
	FOREIGN KEY (course_id) REFERENCES course(id)
);
CREATE TABLE IF NOT EXISTS `enrollment`(
	student_id INT,
	course_id INT,
	semester VARCHAR(10),
	year YEAR,
	grade VARCHAR(2),
	PRIMARY KEY(student_id, course_id, semester, year),
	FOREIGN KEY(student_id) REFERENCES student(id),
	FOREIGN KEY(course_id) REFERENCES course(id)
);
CREATE TABLE IF NOT EXISTS class_schedule (
    id INT AUTO_INCREMENT PRIMARY KEY,
    course_id INT,
    faculty_id INT,
    day_of_week ENUM('Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'),
    start_time TIME,
    end_time TIME,
    location VARCHAR(100),
    FOREIGN KEY (course_id) REFERENCES course(id),
    FOREIGN KEY (faculty_id) REFERENCES faculty(id)
);
CREATE TABLE IF NOT EXISTS exam (
    id INT AUTO_INCREMENT PRIMARY KEY,
    course_id INT,
    exam_type ENUM('MIDTERM', 'FINAL', 'QUIZ', 'PROJECT'),
    date DATE,
    total_marks INT,
    FOREIGN KEY (course_id) REFERENCES course(id)
);
CREATE TABLE IF NOT EXISTS exam_result (
    exam_id INT,
    student_id INT,
    marks_obtained INT,
    PRIMARY KEY (exam_id, student_id),
    FOREIGN KEY (exam_id) REFERENCES exam(id),
    FOREIGN KEY (student_id) REFERENCES student(id)
);
CREATE TABLE IF NOT EXISTS fee_payment (
    id INT AUTO_INCREMENT PRIMARY KEY,
    student_id INT,
    amount DECIMAL(10,2),
    date_paid DATE,
    payment_mode ENUM('CASH', 'CARD', 'ONLINE'),
    FOREIGN KEY (student_id) REFERENCES student(id)
);
CREATE TABLE IF NOT EXISTS program (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    level ENUM('UG', 'PG', 'PhD'),
    department_id INT,
    FOREIGN KEY (department_id) REFERENCES department(id)
);

