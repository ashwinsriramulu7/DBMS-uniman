SELECT s.id, s.name
FROM student s
WHERE NOT EXISTS (
  SELECT cs.course_id
  FROM class_schedule cs
  JOIN faculty f ON cs.faculty_id = f.id
  WHERE f.name = 'Dr. Smith'
  EXCEPT
  SELECT e.course_id
  FROM enrollment e
  WHERE e.student_id = s.id
);

