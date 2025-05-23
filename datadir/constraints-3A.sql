-- CONSTRAINT: MAKE SURE THAT MARKS ARE IN A LOGICAL RANGE
ALTER TABLE RESULT
ADD CONSTRAINT chk_marks CHECK (marks_obtained BETWEEN 0 AND 100);
-- PREVENT NEGATIVE FEE AMOUNTS
ALTER TABLE FEE PAYMENT
ADD CONSTRAINT chk_fee_positive CHECK (amount>=0);
-- VALIDATE ENUM VALUES
ALTER TABLE exam
ADD CONSTRAINT chk_exam_type CHECK (exam_type IN ('MIDTERM', 'FINAL', 'QUIZ', 'PROJECT'));
-- ENSURE COURSE START TIME IS BEFORE END TIME 
ALTER TABLE class_schedule
ADD CONSTRAINT chk_schedule_time CHECK (start_time < end_time);
-- PREVENT SETTING YEAR OF ESTABLISHMENT IN FUTURE FOR COLLEGES 
ALTER TABLE college
ADD CONSTRAINT chk_college_estd CHECK (estd <= YEAR(CURDATE()));
-- ENSURE ENROLLMENT YEAR IS NOT IN THE FUTURE
ALTER TABLE enrollment
ADD CONSTRAINT chk_enrollment_year CHECK (year <= YEAR(CURDATE()));
-- ENSURE THAT A SCHEDULED EXAM IS NOT IN THE PAST
ALTER TABLE exam
ADD CONSTRAINT chk_exam_date_future CHECK (date >= CURDATE());
-- ENSURE PHONE NUMBERS ARE THE RIGHT LENGTH
ALTER TABLE student
ADD CONSTRAINT chk_student_mobile CHECK (LENGTH(mobile_number) BETWEEN 10 AND 15);

ALTER TABLE faculty
ADD CONSTRAINT chk_faculty_mobile CHECK (LENGTH(mobile_number) BETWEEN 10 AND 15);
-- ENSURE THAT COURSE CODE SYNTAX IS VALID 
ALTER TABLE course
ADD CONSTRAINT chk_course_code_format CHECK (course_code REGEXP '^[A-Z]{3}[0-9]{3}$');
-- ENSURE THAT TITLES FOR FACULTY ARE NOT BLANK IF TYPE IS ACADEMIC
ALTER TABLE faculty
ADD CONSTRAINT chk_title_required_for_academic CHECK (
    (type = 'ACADEMIC' AND title IS NOT NULL AND title <> '') OR type = 'NON ACADEMIC'
);


