package middlewares

import (
	"net/http"
	"setad/api/configs"
	"setad/api/models"
	"setad/api/services"
	"setad/api/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func IfLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ifLoggedInErr := findUserIfLoggedIn(c)
		if utils.CheckErrorNotNil(c, ifLoggedInErr, http.StatusBadRequest) {
			c.Abort()
		}
		setUserToRequestBody(c, user)
		c.Next()
	}
}
func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		findUserIfLoggedIn(c)
		ifAdmin(c)
	}
}

func findUserIfLoggedIn(c *gin.Context) (*models.User, error) {
	token, headerErr := checkHeader(c)
	if headerErr != nil {
		return nil, headerErr
	}
	jwtBody, extractingJWTErr := extractFromJWTToken(token)
	if extractingJWTErr != nil {
		return nil, extractingJWTErr
	}
	user, noUserFoundedErr := checkIfUserExits(*jwtBody)
	if noUserFoundedErr != nil {
		return nil, noUserFoundedErr
	}
	return user, nil
}

func ifAdmin(c *gin.Context) bool {
	return true
}

func checkHeader(c *gin.Context) (string, error) {
	HEADER_NAME := "Token"
	if len(c.Request.Header[HEADER_NAME]) == 0 {
		return "", utils.NoAuthHeaderError
	}
	token := c.Request.Header[HEADER_NAME][0]
	return token, nil
}

func extractFromJWTToken(token string) (*models.JWT, error) {
	jwtCodedBody, parsingErr := getJWTCodedBody(token)
	if parsingErr != nil {
		return nil, parsingErr
	}
	return decodeJWTBody(jwtCodedBody)
}

func getJWTCodedBody(token string) (*jwt.Token, error) {
	var JWT_SECRET = configs.JWT_SECRET
	jwtCodedBody, decodingError := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, utils.JWTParsingError
		}
		return []byte(JWT_SECRET), nil
	})
	return jwtCodedBody, decodingError
}

func decodeJWTBody(codedBody *jwt.Token) (*models.JWT, error) {
	claims, ok := codedBody.Claims.(jwt.MapClaims)
	if !ok || !codedBody.Valid {
		return nil, utils.JWTBodyDecodingError
	}
	jwtBody := models.NewJWT()
	jwtBody.Depth = utils.ToInt(claims["depth"])
	jwtBody.PhoneNumber = utils.ToString(claims["phoneNumber"])
	jwtBody.ID = utils.ToObjectID(claims["_id"])
	return &jwtBody, nil
}

func checkIfUserExits(jwtBody models.JWT) (*models.User, error) {
	return services.FindOneUserByPhoneNumber(jwtBody.PhoneNumber)
}

func setUserToRequestBody(c *gin.Context, user *models.User) {
	c.Set("_id", user.ID.Hex())
	c.Set("depth", user.Depth)
}
