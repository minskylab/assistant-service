package assistservice

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type API struct {
	e       *gin.Engine
	Port    string
	prefix  string
	payload *DiseasesPayload
	records *PatientRecord
	profile *PatientProfile
	repo    *Repository
}

func newAPI(prefix string) *API {
	c, _ := extractConfigFromEnv()
	return &API{
		e:      gin.Default(),
		Port:   c.Port,
		prefix: prefix,
	}
}

func (api *API) registerEndpoints() {
	r := api.e.Group(api.prefix)
	api.registerAPI(r)

}

func (api *API) registerAPI(r *gin.RouterGroup) {

	r.GET("/create-user", func(c *gin.Context) {

	})

	r.POST("/typeform-webhook", func(c *gin.Context) {
		newAutoGen := new(WebHookRequest)
		if err := c.BindJSON(newAutoGen); err != nil {
			panic(err)
		}

		if len(newAutoGen.FormResponse.Answers) >= 1 {
			for _,elements := range newAutoGen.FormResponse.Answers {
				
			}
		}
		fmt.Printf("%+v\n", newAutoGen)

	})
}
