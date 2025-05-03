SET FOREIGN_KEY_CHECKS = 0;

TRUNCATE TABLE exam_result;
TRUNCATE TABLE fee_payment;
TRUNCATE TABLE enrollment;
TRUNCATE TABLE teaches;
TRUNCATE TABLE class_schedule;
TRUNCATE TABLE exam;
TRUNCATE TABLE course;
TRUNCATE TABLE program;
TRUNCATE TABLE department;
TRUNCATE TABLE faculty;
TRUNCATE TABLE student;
TRUNCATE TABLE college;

SET FOREIGN_KEY_CHECKS = 1;
ALTER TABLE class_schedule AUTO_INCREMENT=1;
ALTER TABLE college AUTO_INCREMENT=1;
ALTER TABLE course AUTO_INCREMENT=1;
ALTER TABLE department AUTO_INCREMENT=1;
ALTER TABLE enrollment AUTO_INCREMENT=1;
ALTER TABLE exam AUTO_INCREMENT=1;
ALTER TABLE exam_result AUTO_INCREMENT=1;
ALTER TABLE faculty AUTO_INCREMENT=1;
ALTER TABLE fee_payment AUTO_INCREMENT=1;
ALTER TABLE program AUTO_INCREMENT=1;
ALTER TABLE student AUTO_INCREMENT=1;
ALTER TABLE teaches AUTO_INCREMENT=1;
-- Insert into college
INSERT INTO college (id, name, location, estd) VALUES
(1, 'Tech University', 'New York', 1995),
(2, 'Science College', 'California', 1980),
(3, 'Arts Institute', 'Texas', 2000);

-- Insert into faculty
INSERT INTO faculty (id, name, mobile_number, email, address, type, title) VALUES
(1, 'Dr. Alice Smith', '1111111111', 'alice@example.com', '123 Elm St', 'ACADEMIC', 'Professor'),
(2, 'Dr. Bob Johnson', '2222222222', 'bob@example.com', '456 Oak St', 'ACADEMIC', 'Associate Professor'),
(3, 'Mr. Charlie Brown', '3333333333', 'charlie@example.com', '789 Pine St', 'NON ACADEMIC', 'Administrator');

-- Insert into department
INSERT INTO department (id, name, head_of_department, college) VALUES
(1, 'Computer Science', 1, 1),
(2, 'Physics', 2, 2),
(3, 'Literature', NULL, 3);

-- Insert into program
INSERT INTO program (id, name, level, department_id) VALUES
(1, 'B.Tech CSE', 'UG', 1),
(2, 'M.Sc Physics', 'PG', 2),
(3, 'PhD Literature', 'PhD', 3);

-- Insert into student
INSERT INTO student (id, name, mobile_number, email, program_enrolled, type) VALUES
(1, 'Eve Adams', '4444444444', 'eve@example.com', 'B.Tech CSE', 'UG'),
(2, 'Frank Barnes', '5555555555', 'frank@example.com', 'M.Sc Physics', 'PG'),
(3, 'Grace Lee', '6666666666', 'grace@example.com', 'PhD Literature', 'Phd');

-- Insert into course
INSERT INTO course (id, department_id, course_code) VALUES
(1, 1, 'CSE101'),
(2, 2, 'PHY201'),
(3, 3, 'LIT301');

-- Insert into class_schedule
INSERT INTO class_schedule (id, course_id, faculty_id, day_of_week, start_time, end_time, location) VALUES
(1, 1, 1, 'Monday', '09:00:00', '10:30:00', 'Room 101'),
(2, 2, 2, 'Wednesday', '11:00:00', '12:30:00', 'Lab 202'),
(3, 3, 1, 'Friday', '13:00:00', '14:30:00', 'Room 303');

-- Insert into enrollment
INSERT INTO enrollment (student_id, course_id, semester, year, grade) VALUES
(1, 1, 'Fall', 2024, 'A'),
(2, 2, 'Spring', 2024, 'B'),
(3, 3, 'Fall', 2024, 'A');

-- Insert into exam
INSERT INTO exam (id, course_id, exam_type, date, total_marks) VALUES
(1, 1, 'MIDTERM', '2024-10-15', 100),
(2, 2, 'FINAL', '2024-12-10', 100),
(3, 3, 'QUIZ', '2024-11-05', 20);

-- Insert into exam_result
INSERT INTO exam_result (exam_id, student_id, marks_obtained) VALUES
(1, 1, 85),
(2, 2, 78),
(3, 3, 18);

-- Insert into fee_payment
INSERT INTO fee_payment (id, student_id, amount, date_paid, payment_mode) VALUES
(1, 1, 1500.00, '2024-09-01', 'CASH'),
(2, 2, 2000.00, '2024-09-02', 'CARD'),
(3, 3, 1800.00, '2024-09-03', 'ONLINE');

