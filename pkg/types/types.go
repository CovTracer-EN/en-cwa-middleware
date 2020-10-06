package types

import v1 "github.com/google/exposure-notifications-server/pkg/api/v1"

type CWAPublishBody struct {
	Keys                 []v1.ExposureKey `json:"temporaryExposureKeys"`
	HealthAuthorityID    string           `json:"healthAuthorityID"`
	VerificationPayload  string           `json:"verificationPayload,omitempty"`
	HMACKey              string           `json:"hmacKey,omitempty"`
	SymptomOnsetInterval int32            `json:"symptomOnsetInterval,omitempty"`
	Traveler             bool             `json:"traveler,omitempty"`
	RevisionToken        string           `json:"revisionToken"`

	Padding string `json:"padding"`
}

type PathCheckConfig struct {
	MinimumRiskScore                 int64     `json:"minimum_risk_score"`
	AttenuationDurationThresholds    []int64   `json:"attenuation_duration_thresholds"`
	AttenuationLevelValues           []int64   `json:"attenuation_level_values"`
	DaysSinceLastExposureLevelValues []int64   `json:"days_since_last_exposure_level_values"`
	DurationLevelValues              []int64   `json:"duration_level_values"`
	TransmissionRiskLevelValues      []int64   `json:"transmission_risk_level_values"`
	AttenuationBucketWeights         []float64 `json:"attenuation_bucket_weights"`
	TriggerThresholdWeightedDuration int64     `json:"trigger_threshold_weighted_duration"`
}

type VerifiedCodeResponse struct {
	Error    string `json:"error"`
	TestDate string `json:"test_date"`
	TestType string `json:"test_type"`
	Token    string `json:"token"`
}

type TokenVerificationResponse struct {
	Certificate string `json:"certificate"`
	Error       string `json:"error"`
}

type OTPResponse struct {
	Code        string `json:"code"`
	Token       string `json:"token"`
	IsActive    bool   `json:"isActive"`
	IsUsed      bool   `json:"isUsed"`
	TestDate    string `json:"testDate"`
	SymptomDate string `json:"symptomDate"`
	ExpiresAt   string `json:"expiresAt"`
}
