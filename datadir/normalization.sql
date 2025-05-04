-- Step 0: Drop if exists
DROP TABLE IF EXISTS enrollment_3nf, course, enrollment, student, student_course_1nf, student_course_data;

-- Step 1: Create UNNORMALIZED TABLE
CREATE TABLE student_course_data (
  student_id INT,
  student_name VARCHAR(100),
  student_mobile VARCHAR(20),
  course1 VARCHAR(100),
  course2 VARCHAR(100),
  faculty1 VARCHAR(100),
  faculty2 VARCHAR(100),
  fee_paid DECIMAL(10,2)
);

-- Sample Data
INSERT INTO student_course_data VALUES
(1, 'Ravi Kumar', '9876543210', 'Maths', 'Physics', 'Dr. Sen', 'Dr. Bose', 5000.00),
(2, 'Anita Verma', '9123456780', 'Chemistry', 'Biology', 'Dr. Rao', 'Dr. Naik', 4800.00);

-- Step 2: Convert to 1NF - Remove Repeating Groups
CREATE TABLE student_course_1nf (
  student_id INT,
  student_name VARCHAR(100),
  student_mobile VARCHAR(20),
  course VARCHAR(100),
  faculty VARCHAR(100),
  fee_paid DECIMAL(10,2)
);

-- Insert flattened data
INSERT INTO student_course_1nf
SELECT student_id, student_name, student_mobile, course1, faculty1, fee_paid FROM student_course_data
UNION
SELECT student_id, student_name, student_mobile, course2, faculty2, fee_paid FROM student_course_data;

-- Step 3: Convert to 2NF - Remove Partial Dependencies
CREATE TABLE student (
  student_id INT PRIMARY KEY,
  student_name VARCHAR(100),
  student_mobile VARCHAR(20)
);

CREATE TABLE enrollment (
  student_id INT,
  course VARCHAR(100),
  faculty VARCHAR(100),
  fee_paid DECIMAL(10,2),
  PRIMARY KEY (student_id, course),
  FOREIGN KEY (student_id) REFERENCES student(student_id)
);

-- Insert student data
INSERT INTO student (student_id, student_name, student_mobile)
SELECT DISTINCT student_id, student_name, student_mobile FROM student_course_1nf;

-- Insert enrollment data
INSERT INTO enrollment (student_id, course, faculty, fee_paid)
SELECT student_id, course, faculty, fee_paid FROM student_course_1nf;

-- Step 4: Convert to 3NF - Remove Transitive Dependencies
CREATE TABLE course (
  course VARCHAR(100) PRIMARY KEY,
  faculty VARCHAR(100)
);

-- Populate course-faculty mapping
INSERT INTO course (course, faculty)
SELECT DISTINCT course, faculty FROM enrollment;

-- Final normalized enrollment table without faculty
CREATE TABLE enrollment_3nf (
  student_id INT,
  course VARCHAR(100),
  fee_paid DECIMAL(10,2),
  PRIMARY KEY (student_id, course),
  FOREIGN KEY (student_id) REFERENCES student(student_id),
  FOREIGN KEY (course) REFERENCES course(course)
);

-- Populate final enrollment table
INSERT INTO enrollment_3nf (student_id, course, fee_paid)
SELECT student_id, course, fee_paid FROM enrollment;

-- Optionally clean up intermediate tables
-- DROP TABLE enrollment;
-- DROP TABLE student_course_1nf;
-- DROP TABLE student_course_data;

