package JenkinsRoutes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bndr/gojenkins"
)

type Config struct {
	JAL *gojenkins.Jenkins
}

func (c *Config) Test(w http.ResponseWriter, r *http.Request) {
	jobs, err := c.JAL.GetAllJobs()
	if err != nil {
		log.Fatal(err)
	}
	for _, j := range jobs {
		fmt.Printf("Job: %s\n", j.GetName())
	}

}
