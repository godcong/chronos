package chronos

import (
	"testing"
	"time"
)

func TestGetFateData(t *testing.T) {
	birthDate := time.Date(1990, 6, 15, 12, 0, 0, 0, time.Local)

	t.Run("nil input returns error", func(t *testing.T) {
		_, err := GetFateData(nil)
		if err == nil {
			t.Error("expected error for nil input")
		}
		fe, ok := err.(*FateError)
		if !ok {
			t.Errorf("expected *FateError, got %T", err)
		}
		if fe.Code != ErrCodeInputInvalid {
			t.Errorf("expected code %d, got %d", ErrCodeInputInvalid, fe.Code)
		}
	})

	t.Run("zero birth date returns error", func(t *testing.T) {
		_, err := GetFateData(&FateInput{BirthDate: time.Time{}})
		if err == nil {
			t.Error("expected error for zero birth date")
		}
		fe, ok := err.(*FateError)
		if !ok {
			t.Errorf("expected *FateError, got %T", err)
		}
		if fe.Code != ErrCodeInputInvalid {
			t.Errorf("expected code %d, got %d", ErrCodeInputInvalid, fe.Code)
		}
	})

	t.Run("year out of range returns error", func(t *testing.T) {
		_, err := GetFateData(&FateInput{BirthDate: time.Date(1800, 1, 1, 0, 0, 0, 0, time.Local)})
		if err == nil {
			t.Error("expected error for year 1800")
		}
		fe, ok := err.(*FateError)
		if !ok {
			t.Errorf("expected *FateError, got %T", err)
		}
		if fe.Code != ErrCodeDateRange {
			t.Errorf("expected code %d, got %d", ErrCodeDateRange, fe.Code)
		}
	})

	t.Run("valid input returns fate data with balance method", func(t *testing.T) {
		input := &FateInput{
			BirthDate: birthDate,
			Gender:    1,
			IsLunar:   false,
			Surname:   "张",
			Method:    XiYongMethodBalance,
		}
		data, err := GetFateData(input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if data.Bazi == nil {
			t.Error("expected non-nil Bazi")
		}
		if data.WuxingXiji == nil {
			t.Error("expected non-nil WuxingXiji")
		}
		if len(data.Bazi.FourPillars) != 4 {
			t.Errorf("expected 4 pillars, got %d", len(data.Bazi.FourPillars))
		}
		if data.WuxingXiji.UsefulElement == "" {
			t.Error("expected non-empty UsefulElement")
		}
		if data.WuxingXiji.MethodName != "平衡用神法" {
			t.Errorf("expected method name '平衡用神法', got '%s'", data.WuxingXiji.MethodName)
		}
	})

	t.Run("valid input returns fate data with geju method", func(t *testing.T) {
		input := &FateInput{
			BirthDate: birthDate,
			Gender:    1,
			IsLunar:   false,
			Surname:   "张",
			Method:    XiYongMethodGeJu,
		}
		data, err := GetFateData(input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if data.WuxingXiji.MethodName != "格局用神法" {
			t.Errorf("expected method name '格局用神法', got '%s'", data.WuxingXiji.MethodName)
		}
		if data.WuxingXiji.GeJu == nil {
			t.Error("expected non-nil GeJu for geju method")
		}
	})

	t.Run("invalid method defaults to balance", func(t *testing.T) {
		input := &FateInput{
			BirthDate: birthDate,
			Gender:    1,
			Method:    XiYongMethod(99),
		}
		data, err := GetFateData(input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if data.WuxingXiji.Method != XiYongMethodBalance {
			t.Errorf("expected method XiYongMethodBalance, got %d", data.WuxingXiji.Method)
		}
	})
}

func TestFateError_Error(t *testing.T) {
	fe := &FateError{Code: 1001, Message: "test error", Module: "fate"}
	expected := "[fate] error 1001: test error"
	if got := fe.Error(); got != expected {
		t.Errorf("FateError.Error() = %q, want %q", got, expected)
	}
}

func TestCalculateBazi(t *testing.T) {
	calendar := ParseSolarTime(time.Date(2023, 2, 5, 12, 0, 0, 0, time.Local))
	bazi, err := calculateBazi(calendar)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(bazi.FourPillars) != 4 {
		t.Errorf("expected 4 pillars, got %d", len(bazi.FourPillars))
	}
	if bazi.Zodiac == "" {
		t.Error("expected non-empty zodiac")
	}
	if bazi.Constellation == "" {
		t.Error("expected non-empty constellation")
	}
}

func TestCalculateWuxingXiji(t *testing.T) {
	calendar := ParseSolarTime(time.Date(1990, 6, 15, 12, 0, 0, 0, time.Local))
	bazi, _ := calculateBazi(calendar)

	t.Run("balance method", func(t *testing.T) {
		result := calculateWuxingXiji(bazi, XiYongMethodBalance)
		if result.UsefulElement == "" {
			t.Error("expected non-empty UsefulElement")
		}
		if result.Method != XiYongMethodBalance {
			t.Errorf("expected XiYongMethodBalance, got %d", result.Method)
		}
	})

	t.Run("geju method", func(t *testing.T) {
		result := calculateWuxingXiji(bazi, XiYongMethodGeJu)
		if result.UsefulElement == "" {
			t.Error("expected non-empty UsefulElement")
		}
		if result.Method != XiYongMethodGeJu {
			t.Errorf("expected XiYongMethodGeJu, got %d", result.Method)
		}
	})
}
