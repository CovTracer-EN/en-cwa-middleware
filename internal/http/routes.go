package http

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/google/exposure-notifications-server/pkg/api/v1"
	_ "github.com/lib/pq"
	"github.com/rise-center/en-cwa-middleware/internal/utils"
	"github.com/rise-center/en-cwa-middleware/pkg/types"
	"github.com/rise-center/en-cwa-middleware/protocols"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type PathCheckConfig struct {
	DailySummariesConfig             DailySummariesConfig `json:"DailySummariesConfig"`
	TriggerThresholdWeightedDuration int64                `json:"triggerThresholdWeightedDuration"`
}

type DailySummariesConfig struct {
	AttenuationDurationThresholds           []int64   `json:"attenuationDurationThresholds"`
	AttenuationBucketWeights                []float64 `json:"attenuationBucketWeights"`
	ReportTypeWeights                       []float64 `json:"reportTypeWeights"`
	ReportTypeWhenMissing                   int64     `json:"reportTypeWhenMissing"`
	InfectiousnessWeights                   []float64 `json:"infectiousnessWeights"`
	DaysSinceOnsetToInfectiousness          [][]int64 `json:"daysSinceOnsetToInfectiousness"`
	InfectiousnessWhenDaysSinceOnsetMissing int64     `json:"infectiousnessWhenDaysSinceOnsetMissing"`
}

func GetHealth(c *gin.Context) {
	c.Status(http.StatusOK)
}

func PostPublish(c *gin.Context) {
	var body v1.Publish
	if err := c.BindJSON(&body); err != nil {
		log.Print(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	reportType := protocols.ReportType_CONFIRMED_TEST

	//TODO: Find how we will get this info
	daysSinceOnsetOfSymptoms := int32(4000)

	temporaryExposureKeys := make([]*protocols.TemporaryExposureKey, len(body.Keys))
	for i := 0; i < len(body.Keys); i++ {
		transmissionRisk := int32(body.Keys[i].TransmissionRisk)

		keyData, err := base64.StdEncoding.DecodeString(body.Keys[i].Key)
		if err != nil {
			log.Print(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		temporaryExposureKey := protocols.TemporaryExposureKey{
			KeyData:                    keyData,
			TransmissionRiskLevel:      &transmissionRisk,
			RollingStartIntervalNumber: &body.Keys[i].IntervalNumber,
			RollingPeriod:              &body.Keys[i].IntervalCount,
			ReportType:                 &reportType,
			DaysSinceOnsetOfSymptoms:   &daysSinceOnsetOfSymptoms,
		}
		temporaryExposureKeys[i] = &temporaryExposureKey
	}

	origin := "CY"
	consentToFederation := true
	oneEU := []string{
		"AT", "BE", "BG", "CZ", "DE", "DK", "EE", "ES", "FI",
		"FR", "GB", "GR", "HU", "HR", "IE", "IT", "LT", "LU",
		"LV", "MT", "NL", "PL", "PT", "RO", "SE", "SI", "SK",
	}

	submissionPayload := protocols.SubmissionPayload{
		Keys:                temporaryExposureKeys,
		RequestPadding:      []byte(body.Padding),
		VisitedCountries:    oneEU,
		Origin:              &origin,
		ConsentToFederation: &consentToFederation,
	}

	cwaBody, err := proto.Marshal(&submissionPayload)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	cwaHeaders := map[string]string{
		"cwa-fake":          "0",
		"Content-Type":      "application/x-protobuf",
		"cwa-authorization": body.VerificationPayload,
	}

	res, err := utils.HttpPost(os.Getenv("CWAPublishUrl"), cwaBody, cwaHeaders)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var message string
	if res.StatusCode() == 200 {
		message = "Submission payload processed successfully."
	} else if res.StatusCode() == 400 {
		message = "Invalid payload or missing CWA headers."
	} else if res.StatusCode() == 403 {
		message = "Specified OTP invalid."
	}

	log.Println(res.Status(), message)
	c.JSON(res.StatusCode(), v1.PublishResponse{
		// TODO: Change to a real revisionToken to revoke tests
		RevisionToken:     fmt.Sprintf("%d", rand.Int()),
		InsertedExposures: 0,
		ErrorMessage:      message,
		Code:              res.Status(),
		Padding:           "",
	})
}

func PostVerify(c *gin.Context) {
	var otpResponse types.OTPResponse
	if err := c.BindJSON(&otpResponse); err != nil {
		log.Print(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	otpHeaders := map[string]string{
		"Token":        os.Getenv("token"),
		"Content-Type": "application/json",
	}

	res, err := utils.HttpGet(os.Getenv("CWAOtpUrl")+"?code="+otpResponse.Code, otpHeaders)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if res.StatusCode() != 200 {
		log.Println(res.StatusCode(), res.Status())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(res.Body(), &otpResponse)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, types.VerifiedCodeResponse{
		Error:    strconv.FormatBool(otpResponse.IsUsed),
		TestDate: otpResponse.TestDate,
		TestType: "confirmed",
		Token:    otpResponse.Code,
	})
}

func PostCertificate(c *gin.Context) {
	var otpResponse types.OTPResponse
	if err := c.BindJSON(&otpResponse); err != nil {
		log.Print(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, types.TokenVerificationResponse{
		Certificate: otpResponse.Token,
		Error:       "",
	})
}

func GetDownload(c *gin.Context) {
	filename := c.Param("filename")
	if filename == "index.txt" {
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DatabaseHost"),
			os.Getenv("DatabasePort"),
			os.Getenv("DatabaseUser"),
			os.Getenv("DatabasePass"),
			os.Getenv("DatabaseName"))

		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Print(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		err = db.Ping()
		if err != nil {
			log.Print(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var content string
		err = db.
			QueryRow("SELECT content FROM index_file").
			Scan(&content)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		file := strings.Split(strings.TrimSuffix(content, "\n"), "\n")
		for i := 0; i < len(file); i++ {
			file[i] = fmt.Sprintf("cyprus/teks/%s-00001.zip", file[i])
		}

		defer db.Close()
		c.JSON(http.StatusOK, file)
		return
	}

	fromStr := strings.Split(filename, "-")[0]
	fromInt, err := strconv.ParseInt(fromStr, 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	from := time.Unix(fromInt, 0)
	link := fmt.Sprintf(os.Getenv("CWADownloadUrl")+"/date/"+from.Format("2006-01-02")+"/hour/%d", from.Hour())
	res, err := http.Get(link)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	zipFile, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	_, err = c.Writer.Write(zipFile)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func GetConfiguration(c *gin.Context) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DatabaseHost"),
		os.Getenv("DatabasePort"),
		os.Getenv("DatabaseUser"),
		os.Getenv("DatabasePass"),
		os.Getenv("DatabaseName"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Print(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = db.Ping()
	if err != nil {
		log.Print(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	version := os.Getenv("version")
	var pathCheckConfig PathCheckConfig

	err = db.
		QueryRow("SELECT * FROM exposure_configuration WHERE version = $1", version).
		Scan(&version, &pathCheckConfig)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	defer db.Close()
	c.JSON(http.StatusOK, pathCheckConfig)
}

func (a *PathCheckConfig) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
