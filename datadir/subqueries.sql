SELECT s.id, s.name
FROM student s
WHERE s.id IN (
  SELECT er.student_id
  FROM exam_result er
  JOIN exam e ON er.exam_id = e.id
  WHERE e.exam_type = 'FINAL'
  GROUP BY er.student_id
  HAVING AVG(er.marks_obtained) > (
    SELECT AVG(marks_obtained)
    FROM exam_result er2
    JOIN exam e2 ON er2.exam_id = e2.id
    WHERE e2.exam_type = 'FINAL'
  )
);

