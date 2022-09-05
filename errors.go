package chronos

import (
	"errors"
)

var (
	// ErrYearNotHaveLeapMonth returns an error
	ErrYearNotHaveLeapMonth = errors.New("[chronos] year not have leap month")
	// ErrWrongTianGanTypes returns an error
	ErrWrongTianGanTypes = errors.New("[chronos] wrong tiangan types")
	// ErrWrongDiZhiTypes returns an error
	ErrWrongDiZhiTypes = errors.New("[chronos] wrong dizhi types")
	// ErrWrongGanZhiTypes returns an error
	ErrWrongGanZhiTypes = errors.New("[chronos] wrong ganzhi types")
	// ErrWrongZodiacTypes returns an error
	ErrWrongZodiacTypes = errors.New("[chronos] wrong zodiac types")
	// ErrWrongConstellationTypes returns an error
	ErrWrongConstellationTypes = errors.New("[chronos] wrong constellation types")
	// ErrWrongSolarTermFormat returns an error
	ErrWrongSolarTermFormat = errors.New("[chronos] wrong solar term format")
	// ErrWrongSolarTermIndex returns an error
	ErrWrongSolarTermIndex = errors.New("[chronos] wrong solar term index")
)
