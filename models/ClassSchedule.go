package models

type ClassSchedule struct {
    ID         int
    CourseID   int
    FacultyID  int
    DayOfWeek  string
    StartTime  string
    EndTime    string
    Location   string
}

