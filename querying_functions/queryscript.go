package querying_functions
import(
	"context"
	"time"
	"database/sql"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
)
func GetStudentsByCourseID(ctx context.Context, courseID int) ([]models.Student, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `
        SELECT s.id, s.name, s.mobile_number, s.email, s.program_enrolled, s.type
        FROM enrollment e
        JOIN student s ON e.student_id = s.id
        WHERE e.course_id = ?
    `

    rows, err := db.QueryContext(ctx, query, courseID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var students []models.Student
    for rows.Next() {
        var s models.Student
        if err := rows.Scan(&s.ID, &s.Name, &s.MobileNumber, &s.Email, &s.ProgramEnrolled, &s.Type); err != nil {
            return nil, err
        }
        students = append(students, s)
    }

    return students, nil
}

func GetScheduleByCourseID(ctx context.Context, courseID int) ([]models.ClassSchedule, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `
        SELECT id, course_id, faculty_id, day_of_week, start_time, end_time, location
        FROM class_schedule
        WHERE course_id = ?
    `
    rows, err := db.QueryContext(ctx, query, courseID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var schedules []models.ClassSchedule
    for rows.Next() {
        var cs models.ClassSchedule
        if err := rows.Scan(&cs.ID, &cs.CourseID, &cs.FacultyID, &cs.DayOfWeek, &cs.StartTime, &cs.EndTime, &cs.Location); err != nil {
            return nil, err
        }
        schedules = append(schedules, cs)
    }

    return schedules, nil
}
func GetExamResultsByCourseID(ctx context.Context, courseID int) ([]models.ExamResult, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `
        SELECT er.exam_id, er.student_id, er.marks_obtained
        FROM exam_result er
        JOIN exam e ON er.exam_id = e.id
        WHERE e.course_id = ?
    `
    rows, err := db.QueryContext(ctx, query, courseID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []models.ExamResult
    for rows.Next() {
        var r models.ExamResult
        if err := rows.Scan(&r.ExamID, &r.StudentID, &r.MarksObtained); err != nil {
            return nil, err
        }
        results = append(results, r)
    }

    return results, nil
}
func GetCoursesByDepartmentID(ctx context.Context, departmentID int) ([]models.Course, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `
        SELECT id, department_id, course_code
        FROM course
        WHERE department_id = ?
    `
    rows, err := db.QueryContext(ctx, query, departmentID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var courses []models.Course
    for rows.Next() {
        var c models.Course
        if err := rows.Scan(&c.ID, &c.DepartmentID, &c.CourseCode); err != nil {
            return nil, err
        }
        courses = append(courses, c)
    }

    return courses, nil
}
func GetExamsByCourseID(ctx context.Context, courseID int) ([]models.Exam, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `
        SELECT id, course_id, exam_type, date, total_marks
        FROM exam
        WHERE course_id = ?
    `
    rows, err := db.QueryContext(ctx, query, courseID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var exams []models.Exam
    for rows.Next() {
        var e models.Exam
        if err := rows.Scan(&e.ID, &e.CourseID, &e.ExamType, &e.Date, &e.TotalMarks); err != nil {
            return nil, err
        }
        exams = append(exams, e)
    }

    return exams, nil
}
func GetClassScheduleByFacultyID(ctx context.Context, facultyID int) ([]models.ClassSchedule, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `
        SELECT id, course_id, faculty_id, day_of_week, start_time, end_time, location
        FROM class_schedule
        WHERE faculty_id = ?
    `
    rows, err := db.QueryContext(ctx, query, facultyID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var schedules []models.ClassSchedule
    for rows.Next() {
        var cs models.ClassSchedule
        if err := rows.Scan(&cs.ID, &cs.CourseID, &cs.FacultyID, &cs.DayOfWeek, &cs.StartTime, &cs.EndTime, &cs.Location); err != nil {
            return nil, err
        }
        schedules = append(schedules, cs)
    }

    return schedules, nil
}
func GetCoursesByFacultyID(ctx context.Context, facultyID int) ([]models.Course, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `
        SELECT DISTINCT c.id, c.department_id, c.course_code
        FROM course c
        JOIN class_schedule cs ON cs.course_id = c.id
        WHERE cs.faculty_id = ?
    `
    rows, err := db.QueryContext(ctx, query, facultyID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var courses []models.Course
    for rows.Next() {
        var c models.Course
        if err := rows.Scan(&c.ID, &c.DepartmentID, &c.CourseCode); err != nil {
            return nil, err
        }
        courses = append(courses, c)
    }

    return courses, nil
}
func GetDepartmentByHOD(ctx context.Context, facultyID int) (*models.Department, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `
        SELECT id, name, head_of_department, college
        FROM department
        WHERE head_of_department = ?
    `
    var dept models.Department
    err := db.QueryRowContext(ctx, query, facultyID).Scan(&dept.ID, &dept.Name, &dept.HeadOfDepartment, &dept.College)
    if err != nil {
        return nil, err
    }

    return &dept, nil
}
func GetStudentsByFacultyID(ctx context.Context, facultyID int) ([]models.Student, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `
        SELECT DISTINCT s.id, s.name, s.mobile_number, s.email, s.program_enrolled, s.type
        FROM student s
        JOIN enrollment e ON s.id = e.student_id
        JOIN class_schedule cs ON cs.course_id = e.course_id
        WHERE cs.faculty_id = ?
    `

    rows, err := db.QueryContext(ctx, query, facultyID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var students []models.Student
    for rows.Next() {
        var s models.Student
        if err := rows.Scan(&s.ID, &s.Name, &s.MobileNumber, &s.Email, &s.ProgramEnrolled, &s.Type); err != nil {
            return nil, err
        }
        students = append(students, s)
    }

    return students, nil
}
func GetFacultyByEmailOrMobile(ctx context.Context, identifier string) (*models.Faculty, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `
        SELECT id, name, mobile_number, email, address, type, title
        FROM faculty
        WHERE email = ? OR mobile_number = ?
    `

    var f models.Faculty
    err := db.QueryRowContext(ctx, query, identifier, identifier).Scan(&f.ID, &f.Name, &f.MobileNumber, &f.Email, &f.Address, &f.Type, &f.Title)
    if err != nil {
        return nil, err
    }

    return &f, nil
}
func GetStudentByID(ctx context.Context, studentID int) (*models.Student, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `
        SELECT id, name, mobile_number, email, program_enrolled, type
        FROM student
        WHERE id = ?
    `

    var s models.Student
    err := db.QueryRowContext(ctx, query, studentID).Scan(&s.ID, &s.Name, &s.MobileNumber, &s.Email, &s.ProgramEnrolled, &s.Type)
    if err != nil {
        return nil, err
    }

    return &s, nil
}
func GetEnrolledCoursesByStudentID(ctx context.Context, studentID int) ([]models.Course, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `
        SELECT DISTINCT c.id, c.department_id, c.course_code
        FROM course c
        JOIN enrollment e ON e.course_id = c.id
        WHERE e.student_id = ?
    `
    rows, err := db.QueryContext(ctx, query, studentID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var courses []models.Course
    for rows.Next() {
        var c models.Course
        if err := rows.Scan(&c.ID, &c.DepartmentID, &c.CourseCode); err != nil {
            return nil, err
        }
        courses = append(courses, c)
    }

    return courses, nil
}
func GetExamResultsByStudentID(ctx context.Context, studentID int) ([]models.ExamResult, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `
        SELECT exam_id, student_id, marks_obtained
        FROM exam_result
        WHERE student_id = ?
    `
    rows, err := db.QueryContext(ctx, query, studentID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []models.ExamResult
    for rows.Next() {
        var r models.ExamResult
        if err := rows.Scan(&r.ExamID, &r.StudentID, &r.MarksObtained); err != nil {
            return nil, err
        }
        results = append(results, r)
    }

    return results, nil
}
func GetScheduleByStudentID(ctx context.Context, studentID int) ([]models.ClassSchedule, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `
        SELECT DISTINCT cs.id, cs.course_id, cs.faculty_id, cs.day_of_week, cs.start_time, cs.end_time, cs.location
        FROM class_schedule cs
        JOIN enrollment e ON cs.course_id = e.course_id
        WHERE e.student_id = ?
    `
    rows, err := db.QueryContext(ctx, query, studentID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var schedules []models.ClassSchedule
    for rows.Next() {
        var cs models.ClassSchedule
        if err := rows.Scan(&cs.ID, &cs.CourseID, &cs.FacultyID, &cs.DayOfWeek, &cs.StartTime, &cs.EndTime, &cs.Location); err != nil {
            return nil, err
        }
        schedules = append(schedules, cs)
    }

    return schedules, nil
}
func GetFeePaymentsByStudentID(ctx context.Context, studentID int) ([]models.FeePayment, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `
        SELECT id, student_id, amount, date_paid, payment_mode
        FROM fee_payment
        WHERE student_id = ?
    `
    rows, err := db.QueryContext(ctx, query, studentID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var payments []models.FeePayment
    for rows.Next() {
        var p models.FeePayment
        if err := rows.Scan(&p.ID, &p.StudentID, &p.Amount, &p.DatePaid, &p.PaymentMode); err != nil {
            return nil, err
        }
        payments = append(payments, p)
    }

    return payments, nil
}
func GetDepartmentsByCollegeID(ctx context.Context, collegeID int) ([]models.Department, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `
        SELECT id, name, head_of_department, college
        FROM department
        WHERE college = ?
    `
    rows, err := db.QueryContext(ctx, query, collegeID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var departments []models.Department
    for rows.Next() {
        var d models.Department
        if err := rows.Scan(&d.ID, &d.Name, &d.HeadOfDepartment, &d.College); err != nil {
            return nil, err
        }
        departments = append(departments, d)
    }

    return departments, nil
}
func GetStudentsByProgram(ctx context.Context, program string) ([]models.Student, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `
        SELECT id, name, mobile_number, email, program_enrolled, type
        FROM student
        WHERE program_enrolled = ?
    `
    rows, err := db.QueryContext(ctx, query, program)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var students []models.Student
    for rows.Next() {
        var s models.Student
        if err := rows.Scan(&s.ID, &s.Name, &s.MobileNumber, &s.Email, &s.ProgramEnrolled, &s.Type); err != nil {
            return nil, err
        }
        students = append(students, s)
    }

    return students, nil
}
func GetAllColleges(ctx context.Context) ([]models.College, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `SELECT id, name, location, estd FROM college`

    rows, err := db.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var colleges []models.College
    for rows.Next() {
        var c models.College
        if err := rows.Scan(&c.ID, &c.Name, &c.Location, &c.Estd); err != nil {
            return nil, err
        }
        colleges = append(colleges, c)
    }

    return colleges, nil
}
func GetFacultyByDepartmentID(ctx context.Context, departmentID int) ([]models.Faculty, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `
        SELECT DISTINCT f.id, f.name, f.mobile_number, f.email, f.address, f.type, f.title
        FROM faculty f
        JOIN department d ON d.head_of_department = f.id
        WHERE d.id = ?
    `
    rows, err := db.QueryContext(ctx, query, departmentID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var faculty []models.Faculty
    for rows.Next() {
        var f models.Faculty
        if err := rows.Scan(&f.ID, &f.Name, &f.MobileNumber, &f.Email, &f.Address, &f.Type, &f.Title); err != nil {
            return nil, err
        }
        faculty = append(faculty, f)
    }

    return faculty, nil
}
type CourseEnrollmentStat struct {
    CourseID    int
    TotalEnrolled int
}

func GetEnrollmentCountPerCourse(ctx context.Context) ([]CourseEnrollmentStat, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `
        SELECT course_id, COUNT(student_id) as total_enrolled
        FROM enrollment
        GROUP BY course_id
    `
    rows, err := db.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var stats []CourseEnrollmentStat
    for rows.Next() {
        var stat CourseEnrollmentStat
        if err := rows.Scan(&stat.CourseID, &stat.TotalEnrolled); err != nil {
            return nil, err
        }
        stats = append(stats, stat)
    }

    return stats, nil
}
func GetTotalFeePaidByStudent(ctx context.Context, studentID int) (float64, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `
        SELECT COALESCE(SUM(amount), 0)
        FROM fee_payment
        WHERE student_id = ?
    `
    var total float64
    err := db.QueryRowContext(ctx, query, studentID).Scan(&total)
    return total, err
}
func GetFeePaymentsByDate(ctx context.Context, date string) ([]models.FeePayment, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `
        SELECT id, student_id, amount, date_paid, payment_mode
        FROM fee_payment
        WHERE date_paid = ?
    `
    rows, err := db.QueryContext(ctx, query, date)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var payments []models.FeePayment
    for rows.Next() {
        var p models.FeePayment
        if err := rows.Scan(&p.ID, &p.StudentID, &p.Amount, &p.DatePaid, &p.PaymentMode); err != nil {
            return nil, err
        }
        payments = append(payments, p)
    }

    return payments, nil
}
type StudentPerformance struct {
    StudentID int
    AvgMarks  float64
}

func GetTopPerformers(ctx context.Context) ([]StudentPerformance, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `
        SELECT student_id, AVG(marks_obtained) as avg_marks
        FROM exam_result
        GROUP BY student_id
        ORDER BY avg_marks DESC
        LIMIT 5
    `
    rows, err := db.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []StudentPerformance
    for rows.Next() {
        var sp StudentPerformance
        if err := rows.Scan(&sp.StudentID, &sp.AvgMarks); err != nil {
            return nil, err
        }
        results = append(results, sp)
    }

    return results, nil
}
func GetFailedStudents(ctx context.Context) ([]int, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `
        SELECT student_id
        FROM exam_result
        GROUP BY student_id
        HAVING AVG(marks_obtained) < 40
    `
    rows, err := db.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var failed []int
    for rows.Next() {
        var id int
        if err := rows.Scan(&id); err != nil {
            return nil, err
        }
        failed = append(failed, id)
    }

    return failed, nil
}
type ProgramStat struct {
    Program string
    Count   int
}

func GetStudentCountPerProgram(ctx context.Context) ([]ProgramStat, error) {
    db := includes.InitDB()
    defer db.Close()

    query := `
        SELECT program_enrolled, COUNT(*) as count
        FROM student
        GROUP BY program_enrolled
    `
    rows, err := db.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var stats []ProgramStat
    for rows.Next() {
        var ps ProgramStat
        if err := rows.Scan(&ps.Program, &ps.Count); err != nil {
            return nil, err
        }
        stats = append(stats, ps)
    }

    return stats, nil
}

