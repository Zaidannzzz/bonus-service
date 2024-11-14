package helpers

import (
	"backend-api-test/config"
	api_response "backend-api-test/pkg/api-response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

/*
   |--------------------------------------------------------------------------
   | Return Error Message For Validation
   |--------------------------------------------------------------------------
   |
   | This function is for return error message for validation,
   | Its can return array string depends on the errors.
   | This function need go-playground-validator package.
*/

func ErrorMessage(err interface{}) []string {
	errorMessages := []string{}
	for _, e := range err.(validator.ValidationErrors) {
		fmt.Println(e.ActualTag())
		switch e.ActualTag() {
		case "min":
			errorMessage := fmt.Sprintf("Error on field %s, condition: Should Be At Least %v Character", e.Field(), e.Param())
			errorMessages = append(errorMessages, errorMessage)
		case "e164":
			errorMessage := fmt.Sprintf("Error on field %s, condition: Must Use Country Code Like: +62", e.Field())
			errorMessages = append(errorMessages, errorMessage)
		case "email":
			errorMessage := fmt.Sprintf("Error on field %s, condition: Must Use The Correct Email Format", e.Field())
			errorMessages = append(errorMessages, errorMessage)
		case "gte":
			errorMessage := fmt.Sprintf("Error on field %s, condition: Must Grather Than Equals %v", e.Field(), e.Param())
			errorMessages = append(errorMessages, errorMessage)
		default:
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
	}
	return errorMessages
}

/*
	|--------------------------------------------------------------------------
	| to remove all special character on query string
	|--------------------------------------------------------------------------
	|
	| This function is for return new string from query string to avoid XSS and sql injection
	|
*/

func Escape(letter *string) *string {
	var newString string
	re := regexp.MustCompile(`[%\w\d\s]+`)

	for _, match := range re.FindAllString(*letter, -1) {
		if len(match) > 1 {
			newString += match
		} else {
			newString = match
		}
	}
	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	finalString := re_leadclose_whtsp.ReplaceAllString(newString, "")
	finalString = re_inside_whtsp.ReplaceAllString(finalString, " ")
	return &finalString
}

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func ResponseError(statusCode int32, err error, ctx *gin.Context) {
	if err != nil {
		ctx.Error(err)
		var pc uintptr
		var line int
		_, exists := ctx.Get("panic")
		if exists == true {
			pc, _, line, _ = runtime.Caller(3)
			ctx.Set("panic_pc", pc)
			ctx.Set("panic_line", line)
		} else {
			pc, _, line, _ = runtime.Caller(1)
		}
		ctx.Set("error_path", runtime.FuncForPC(pc).Name())
		ctx.Set("error_line", line)
		ctx.Set("error_time", time.Now().Format(time.RFC850))
		ctx.Set("error_message", err)
		ctx.Set("error_status_code", statusCode)
		if statusCode == 404 {
			ctx.SecureJSON(http.StatusNotFound, api_response.NotFound(nil, err.Error()))
			return
		} else if statusCode == 400 {
			ctx.SecureJSON(http.StatusBadRequest, api_response.BadRequest(err.Error()))
			return
		} else if statusCode == 401 {
			ctx.SecureJSON(http.StatusUnauthorized, api_response.UnAuthorized(nil, err.Error()))
			return
		} else if statusCode == 403 {
			ctx.SecureJSON(http.StatusForbidden, api_response.Forbidden(nil, err.Error()))
			return
		} else if statusCode == 501 {
			ctx.SecureJSON(http.StatusNotImplemented, api_response.ServerNotImplemented(err.Error()))
			return
		}
		ctx.SecureJSON(http.StatusInternalServerError, api_response.ServerError(err.Error()))
		return
	}
}

func GetUserEmailFromMiddleware(c *gin.Context) string {
	emailValue, exists := c.Get("email")
	if exists != true {
		c.JSON(http.StatusInternalServerError, api_response.ServerError("Email authentication requested but middleware not registered on routes"))
	}
	return emailValue.(string)
}

func GetUserIdFromMiddleware(c *gin.Context) int {
	// c.Get("user_id") will return uint based on models.Users
	// and from usecases function we use int therefore
	// we convert to int
	uidValue, exists := c.Get("user_id")
	if exists != true {
		c.JSON(http.StatusInternalServerError, api_response.ServerError("UserId authentication requested but middleware not registered on routes"))
	}
	uidString := fmt.Sprintf("%v", uidValue)
	uidInt, _ := strconv.Atoi(uidString)

	return uidInt
}

func GetUserUUIDFromMiddleware(c *gin.Context) string {
	uidValue, exists := c.Get("user_uuid")
	if exists != true {
		c.JSON(http.StatusInternalServerError, api_response.ServerError("UserUUID authentication requested but middleware not registered on routes"))
	}
	userUUID := fmt.Sprintf("%v", uidValue)
	return userUUID
}

func IsDevelopment(conf *config.Config) bool {
	return conf.Server.Mode != "staging" && conf.Server.Mode != "production"
}

// EpochTimeValidator validates that the value is a valid epoch timestamp.
func EpochTimeChecker(epoch int64) bool {
	// Validate that it's a valid epoch timestamp
	t := time.Unix(epoch, 0)
	fmt.Println(fmt.Sprintf("--- Helper Validation Epoch: %v ---", t.Year()))
	return t.Year() >= 2018
}

func CreateUUIDV7() *uuid.UUID {
	UUIDV7, errUUIDV7 := uuid.NewV7()
	if errUUIDV7 != nil {
		log.Fatalln(errUUIDV7)
	}
	return &UUIDV7
}

func ValidateFile(fileSizeMB int, validExt, filetitle string, fileData *multipart.FileHeader) (err error) {
	// CHECK VALID EXT
	fileExt := filepath.Ext(fileData.Filename)
	fileExt = strings.ToLower(fileExt)
	fmt.Println("fileExt: ", fileExt)
	fmt.Println("fileData.Filename: ", fileData.Filename)
	fmt.Println("fileData.Filename: ", fileData.Filename)
	if !strings.Contains(validExt, fileExt) || fileExt == "" {
		return fmt.Errorf("Invalid file extensions")
	}

	// VALID SIZE
	fz := fileSizeMB * 1024 * 1024
	if fileData.Size > int64(fz) {
		return fmt.Errorf("The maximum file size for %s is %dMB.", filetitle, fileSizeMB)
	}

	return err
}
