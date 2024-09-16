package features

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/ibiscum/Building-Microservices-with-Go/chapter04/data"
)

var criteria interface{}
var response *http.Response
var err error

func iHaveNoSearchCriteria() error {
	if criteria != nil {
		return fmt.Errorf("Criteria should be nil")
	}

	return nil
}

func iCallTheSearchEndpoint() error {
	var request []byte

	if criteria != nil {
		request = []byte(criteria.(string))
	}

	response, err = http.Post("http://localhost:8323", "application/json", bytes.NewReader(request))
	return err
}

func iShouldReceiveABadRequestMessage() error {
	if response.StatusCode != http.StatusBadRequest {
		return fmt.Errorf("Should have recieved a bad response")
	}

	return nil
}

func iHaveAValidSearchCriteria() error {
	criteria = `{ "query": "Fat Freddy's Cat" }`

	return nil
}

func iShouldReceiveAListOfKittens() error {
	var body []byte
	body, err := io.ReadAll(response.Body)

	if len(body) < 1 || err != nil {
		return fmt.Errorf("Should have received a list of kittens")
	}

	return nil
}

// func FeatureContext(s *godog.ScenarioContext) {
// 	s.Step(`^I have no search criteria$`, iHaveNoSearchCriteria)
// 	s.Step(`^I call the search endpoint$`, iCallTheSearchEndpoint)
// 	s.Step(`^I should receive a bad request message$`, iShouldReceiveABadRequestMessage)
// 	s.Step(`^I have a valid search criteria$`, iHaveAValidSearchCriteria)
// 	s.Step(`^I should receive a list of kittens$`, iShouldReceiveAListOfKittens)

// s.Before(func(interface{}) {
// 	clearDB()
// 	setupData()
// 	startServer()
// })

// s.After(func(interface{}, error) {
// 	server.Process.Signal(syscall.SIGINT)
// })

// 	waitForDB()
// }

var server *exec.Cmd
var store *data.MongoStore

func startServer() {
	server = exec.Command("go", "build", "../main.go")
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}

	server = exec.Command("./main")
	go func() {
		err := server.Run()
		if err != nil {
			log.Fatal(err)
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Printf("Server running with pid: %v", server.Process.Pid)
}

func waitForDB() {
	var err error

	serverURI := "localhost"
	if os.Getenv("DOCKER_IP") != "" {
		serverURI = os.Getenv("DOCKER_IP")
	}

	for i := 0; i < 10; i++ {
		store, err = data.NewMongoStore(serverURI)
		if err == nil {
			break
		}

		time.Sleep(1 * time.Second)
	}
}

func clearDB() {
	store.DeleteAllKittens()
}

func setupData() {
	store.InsertKittens(
		[]data.Kitten{
			{
				Id:     "1",
				Name:   "Felix",
				Weight: 12.3,
			},
			{
				Id:     "2",
				Name:   "Fat Freddy's Cat",
				Weight: 20.0,
			},
			{
				Id:     "3",
				Name:   "Garfield",
				Weight: 35.0,
			},
		})
}

func TestSearchFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"./"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I have no search criteria$`, iHaveNoSearchCriteria)
	ctx.Step(`^I call the search endpoint$`, iCallTheSearchEndpoint)
	ctx.Step(`^I should receive a bad request message$`, iShouldReceiveABadRequestMessage)
	ctx.Step(`^I have a valid search criteria$`, iHaveAValidSearchCriteria)
	ctx.Step(`^I should receive a list of kittens$`, iShouldReceiveAListOfKittens)

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		clearDB()
		setupData()
		startServer()
		return ctx, nil
	})

	//ctx.After(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	// 	server.Process.Signal(syscall.SIGINT)
	//})

	waitForDB()
}
