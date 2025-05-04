CREATE VIEW class_schedule_view AS
SELECT cs.id, cs.day_of_week, cs.start_time, cs.end_time, cs.location,
       c.course_code, f.name AS faculty_name
FROM class_schedule cs
JOIN course c ON cs.course_id = c.id
JOIN faculty f ON cs.faculty_id = f.id;

