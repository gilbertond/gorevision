package contracts

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pact-foundation/pact-go/dsl"
)

func setup() *http.Response {
	pact := dsl.Pact{
		Consumer:                 "api-consumer-3",
		Provider:                 "api-producer",
		LogLevel:                 "DEBUG",
		LogDir:                   "logs",
		PactDir:                  "packs", // can be a path to s3 bucket object
		PactFileWriteMode:        "merge", // if tests are in different files which is usually the case
		SpecificationVersion:     2,       // 1 or 2
		Host:                     "127.0.0.1",
		DisableToolValidityCheck: true,
		ClientTimeout:            time.Second * 10,
		AllowedMockServerPorts:   "64672",
	}

	defer pact.Teardown()
	pact.Setup(true)

	pact.
		AddInteraction().
		Given("some state").
		UponReceiving("name of test case").
		WithRequest(dsl.Request{
			Method:  "GET",
			Path:    dsl.String("/path"),
			Query:   dsl.MapMatcher{"param1": dsl.String("556"), "param2": dsl.String("abc")},
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json"), "Authorization": dsl.String("Bearer user.token.string")},
			Body:    nil,
		}).
		WillRespondWith(dsl.Response{
			Status:  200,
			Headers: dsl.MapMatcher{"Context-Type": dsl.String("application/json")},
			//Body: dsl.Match(&models.MyStruct{
			//	Field1: "value1",
			//	Field3: 0.1,
			//	Field4: 10,
			//}),
		})

	var res *http.Response
	makeCall := func() error {
		u := fmt.Sprintf("http://127.0.0.1:%d/path?param1=556&param2=abc", pact.Server.Port)
		request, err := http.NewRequest("GET", u, nil /*strings.NewReader(`{"field":"value"}`)*/)

		if err != nil {
			return err
		}
		request.Header.Set("Authorization", "Bearer user.token.string")
		request.Header.Set("Content-Type", "application/json")
		// request.RequestURI = "param1=556&param2=abc"
		//if res, err := http.DefaultClient.Do(request); err != nil {
		res, err = http.DefaultClient.Do(request)
		if err != nil {
			return err
		}
		return nil
	}
	if err := pact.Verify(makeCall); err != nil {
		log.Fatalf("Error occured: %v", err)
	}
	return res
}
