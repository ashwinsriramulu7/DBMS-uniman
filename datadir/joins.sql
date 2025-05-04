SELECT s.name AS student_name, c.course_code, e.exam_type, er.marks_obtained
FROM exam_result er
JOIN student s ON er.student_id = s.id
JOIN exam e ON er.exam_id = e.id
JOIN course c ON e.course_id = c.id;

