package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/orto-core/server/auth-service/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestRegisterRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	user := models.User{
		Fullname: "Nathan Kulewoshie",
		Email:    "nate.ku24@gmail.com",
		Password: "12345",
	}
	router := setupRouter()
	w := httptest.NewRecorder()

	userJson, _ := json.Marshal(user)

	req, _ := http.NewRequest(http.MethodPost, "/register", strings.NewReader(string(userJson)))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	expectedResponse := `{"fullname":"Nathan Kulewoshie","email":"nate.ku24@gmail.com","message":"registration successful"}`
	assert.JSONEq(t, expectedResponse, w.Body.String())
}
