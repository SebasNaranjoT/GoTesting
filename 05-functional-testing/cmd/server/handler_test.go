package server

import (
	"bytes"
	"encoding/json"
	"functional/prey"
	"functional/shark"
	"functional/simulator"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	r := gin.Default()
	sim := simulator.NewCatchSimulator(35.4)

	whiteShark := shark.CreateWhiteShark(sim)
	tuna := prey.CreateTuna()

	handler := NewHandler(whiteShark, tuna)

	g := r.Group("/v1")

	g.PUT("/shark", handler.ConfigureShark())
	g.PUT("/prey", handler.ConfigurePrey())
	g.POST("/simulate", handler.SimulateHunt())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))

	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}
func TestConfigureShark(t *testing.T){
	type request struct {
		XPosition float64 `json:"x_position"`
		YPosition float64 `json:"y_position"`
		Speed     float64 `json:"speed"`
	}
	type response struct {
		Success bool    `json:"success"`
		Data    request `json:"data"`
	}

	t.Run("Should configure a shark with HTTP 200 response", func(t *testing.T){
		//Arrange
		server := createServer()
		req, res := createRequestTest(http.MethodPut, "/v1/shark",
		`{"x_position": 5.0, "y_position": 5.0, "speed": 5.0}`)
		
		expectedResponse := response{
			Success: true,
			Data: request{
				XPosition: 5.0,
				YPosition: 5.0,
				Speed: 5.0,
			},
		}

		//Act
		server.ServeHTTP(res, req)
		var r response
		err := json.Unmarshal(res.Body.Bytes(), &r)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, expectedResponse, r)
	})

	t.Run("Should return nil struct with HTTP 400 response", func(t *testing.T){
		//arrange
		expectedResponse := response{
			Success: false,
			Data: request{
			},
		}
		server := createServer()
		req, res := createRequestTest(http.MethodPut, "/v1/shark",`{`)
		
		//act
		server.ServeHTTP(res, req)
		var r response

		//assert
		assert.Equal(t, 400, res.Code)
		assert.Equal(t, expectedResponse, r)
	})
}

func TestConfigurePrey(t *testing.T){
	type request struct {
		Speed float64 `json:"speed"`
	}
	type response struct {
		Success bool `json:"success"`
	}
	t.Run("should configure a prey with HTTP 200 respponse", func(t *testing.T){
		//arrange
		expectedResponse := response{
			Success: true,
		}

		server := createServer()
		req, res := createRequestTest(http.MethodPut, "/v1/prey",`{"speed": 14.0}`)

		//act
		server.ServeHTTP(res, req)
		var r response
		err := json.Unmarshal(res.Body.Bytes(), &r)
		
		//assert
		assert.NoError(t, err)
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, expectedResponse, r)
	})

	t.Run("should return HTTP 400 response", func(t *testing.T){
		//arrange
		expectedResponse := response{
		}

		server := createServer()
		req, res := createRequestTest(http.MethodPut, "/v1/prey",`{`)
	
		//act
		server.ServeHTTP(res, req)
		//assert
		var r response
		err := json.Unmarshal(res.Body.Bytes(), &r)
		
		//assert
		assert.NoError(t, err)
		assert.Equal(t, 400, res.Code)
		assert.Equal(t, expectedResponse, r)
	})
}