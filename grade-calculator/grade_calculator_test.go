package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 55, Assignment)
	gradeCalculator.AddGrade("exam 1", 55, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 55, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestInvalidGradeTypeIgnored(t *testing.T) {
    gc := NewGradeCalculator()

    gc.AddGrade("invalid", 100, GradeType(99))

    if got := gc.GetFinalGrade(); got != "F" {
        t.Errorf("Expected F for invalid grade type, got %s", got)
    }
}

func TestGetGradeCAndD(t *testing.T) {
    tests := []struct {
        name     string
        score    int
        expected string
    }{
        {name: "C", score: 75, expected: "C"},
        {name: "D", score: 65, expected: "D"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gc := NewGradeCalculator()
            gc.AddGrade("assignment", tt.score, Assignment)
            gc.AddGrade("exam", tt.score, Exam)
            gc.AddGrade("essay", tt.score, Essay)

            if got := gc.GetFinalGrade(); got != tt.expected {
                t.Errorf("GetFinalGrade() = %q; want %q", got, tt.expected)
            }
        })
    }
}

func TestGradeTypeString(t *testing.T) {
    tests := []struct {
        gradeType GradeType
        expected  string
    }{
        {gradeType: Assignment, expected: "assignment"},
        {gradeType: Exam, expected: "exam"},
        {gradeType: Essay, expected: "essay"},
    }

    for _, tt := range tests {
        if got := tt.gradeType.String(); got != tt.expected {
            t.Errorf("%v.String() = %q; want %q", tt.gradeType, got, tt.expected)
        }
    }
}
