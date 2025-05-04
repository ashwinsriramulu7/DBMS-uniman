ALTER TABLE `student`
ADD CONSTRAINT chk_student_type CHECK (`type` IN ('UG', 'PG', 'Phd'));

ALTER TABLE `faculty`
ADD CONSTRAINT chk_faculty_type CHECK (`type` IN ('ACADEMIC', 'NON ACADEMIC'));

ALTER TABLE `fee_payment`
ADD CONSTRAINT chk_amount CHECK (`amount` >= 0);

SELECT DISTINCT f.name
FROM faculty f
JOIN class_schedule cs ON f.id = cs.faculty_id
WHERE cs.day_of_week = 'Monday';

