package chronos

import (
	"errors"
)

var (
	ErrWrongTianGanTypes = errors.New("[chronos] wrong tiangan types")
	ErrWrongDiZhiTypes = errors.New("[chronos] wrong dizhi types")
	ErrWrongGanZhiTypes = errors.New("[chronos] wrong ganzhi types")
	ErrWrongZodiacTypes = errors.New("[chronos] wrong zodiac types")
	ErrWrongConstellationTypes = errors.New("[chronos] wrong constellation types")
	ErrWrongSolarTermFormat = errors.New("[chronos] wrong solar term format")
	ErrWrongSolarTermIndex = errors.New("[chronos] wrong solar term index")
)
