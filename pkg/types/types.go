package types

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
