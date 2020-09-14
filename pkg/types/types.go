package types

import v1 "github.com/google/exposure-notifications-server/pkg/api/v1"

type CWAPublishBody struct {
	Keys                 []v1.ExposureKey `json:"temporaryExposureKeys"`
	HealthAuthorityID    string        `json:"healthAuthorityID"`
	VerificationPayload  string        `json:"verificationPayload,omitempty"`
	HMACKey              string        `json:"hmacKey,omitempty"`
	SymptomOnsetInterval int32         `json:"symptomOnsetInterval,omitempty"`
	Traveler             bool          `json:"traveler,omitempty"`
	RevisionToken        string        `json:"revisionToken"`

	Padding string `json:"padding"`
}
