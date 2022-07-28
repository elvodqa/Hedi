package api

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/elvodqa/hedi/database"

	"github.com/gin-gonic/gin"
)

type ApiLoginResponse struct {
	HediUsername string `json:"HediUsername"`
	HediToken    string `json:"HediToken"`
	HediUserId   int64  `json:"HediUserId"`
}

func ApiHandleWaffleLogin(ctx *gin.Context) {
	formUsername := ctx.PostForm("username")
	formPassword := ctx.PostForm("password")

	ctx.Header("Access-Control-Allow-Origin", "127.0.0.1")
	ctx.Header("Access-Control-Allow-Origin", "http://localhost:9467")

	passwordHashed := md5.Sum([]byte(formPassword))
	passwordHashedString := hex.EncodeToString(passwordHashed[:])

	userId, authSuccess := database.AuthenticateUser(formUsername, passwordHashedString)

	loginResponse := ApiLoginResponse{}

	if userId == -2 {
		ctx.String(http.StatusInternalServerError, "")
		return
	}

	if userId == -1 {
		loginResponse.HediToken = ""
		loginResponse.HediUsername = ""
		loginResponse.HediUserId = -1

		data, marshalErr := json.Marshal(loginResponse)

		if marshalErr != nil {
			ctx.String(http.StatusInternalServerError, "")
			return
		}

		ctx.Data(http.StatusOK, "hedi/blob", data)
		return
	}

	userQueryResult, user := database.UserFromDatabaseById(uint64(userId))

	if userQueryResult == -2 {
		ctx.String(http.StatusInternalServerError, "")
		return
	}

	if userQueryResult == -1 {
		loginResponse.HediToken = ""
		loginResponse.HediUsername = ""
		loginResponse.HediUserId = -1

		data, marshalErr := json.Marshal(loginResponse)

		if marshalErr != nil {
			ctx.String(http.StatusInternalServerError, "")
			return
		}

		ctx.Data(http.StatusOK, "waffle/blob", data)
		return
	}

	if authSuccess {
		loginResponse.HediToken = database.TokensCreateNewToken(user)
		loginResponse.HediUserId = int64(userId)
		loginResponse.HediUsername = user.Username

		data, marshalErr := json.Marshal(loginResponse)

		if marshalErr != nil {
			ctx.String(http.StatusInternalServerError, "")
			return
		}

		ctx.Data(http.StatusOK, "waffle/blob", data)
	} else {
		loginResponse.HediToken = ""
		loginResponse.HediUsername = ""
		loginResponse.HediUserId = -1

		data, marshalErr := json.Marshal(loginResponse)

		if marshalErr != nil {
			ctx.String(http.StatusInternalServerError, "")
			return
		}

		ctx.Data(http.StatusOK, "waffle/blob", data)
		return
	}
}
