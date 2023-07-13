package contracts

import (
	"fmt"
	. "github.com/onsi/gomega"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"log"
	"net/http"
	"os"
)

var dataRepository = "" // User doesn't exist

// 1. start provider api
func startServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/path?param1=1556&param2=abc", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Mako", "'{\"user_id\": 76661}'")
		w.Header().Add("Authorization", "user.token.string")
		fmt.Fprintf(w, fmt.Sprintf(`{"dataRepository":"%s"}`, dataRepository))
	})

	//mux.HandleFunc("/path?param1=556&param2=abc", func(w http.ResponseWriter, req *http.Request) {
	//	w.Header().Add("Content-Type", "application/json")
	//	w.Header().Add("Mako", "'{\"user_id\": 76661}'")
	//	w.Header().Add("Authorization", "user.token.string")
	//	fmt.Fprintf(w, fmt.Sprintf(`{"dataRepository":"%s"}`, dataRepository))
	//})
	log.Fatal(http.ListenAndServe(":8000", mux))
}

// 2. verify provider api
func TestSetupProducer() {
	// create a pact to connect to pact daemon
	pact := &dsl.Pact{
		Provider: "api-producer",
	}
	go startServer()

	//packUrl := "./packs" // This can be a path to s3 bucket
	xx, err := pact.VerifyProviderRaw(types.VerifyRequest{
		ProviderBaseURL: "http://localhost:8000",

		//PactURLs:  []string{filepath.ToSlash(fmt.Sprintf("%s/api-consumer-api-producer.json", packUrl))}, // specific pacts
		BrokerURL:                  "http://localhost",       // Instead of above, 			this will pick pacts for all consumers
		Tags:                       []string{"dev", "stage"}, //use tags to pick pacts with specific tags
		ProviderVersion:            "1.0.0",
		BrokerUsername:             "",
		BrokerPassword:             "",
		PublishVerificationResults: true,
		FailIfNoPactsFound:         false,
		StateHandlers: types.StateHandlers{
			"some state": func() error {
				dataRepository = "ndenzi"
				return nil
			},
			"some state1": func() error {
				dataRepository = "state 1"
				return nil
			},
			"some state2": func() error {
				dataRepository = "does not exist"
				return nil
			},
			"some state3": func() error {
				dataRepository = "some other mocked up state"
				return nil
			},
		},
		BeforeEach: func() error {
			fmt.Println("before hook to run before verification")
			return nil
		},
		AfterEach: func() error {
			fmt.Println("after hook to run after verification")
			return nil
		},
		CustomProviderHeaders: []string{"Authorization: Bearer user.token.string", /*"Mako: '{\"user_id\": 76661}'",  "Content-Type: application/json" */}, // custom headers from consumers
		RequestFilter: func(handler http.Handler) http.Handler { // change req/resp payload capability. handy for things that you cannot or don't want to persist in a pact file
			return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				// 	do something
				// 	e.g if you wanna modify a token
				// request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "some.locally.made.token"))
				// handler.ServeHTTP(writer, request)
			})
		},
		// ConsumerVersionSelectors:   nil,
		// ProviderTags:               nil,
		// Provider:                   "",
		// BrokerUsername:             "",
		// BrokerPassword:             "",
		// BrokerToken:                "",
		// FailIfNoPactsFound:         false,
		// PublishVerificationResults: false,
		// ProviderVersion:            "",
		// CustomTLSConfig:      nil,
		// EnablePending:        false,
		// IncludeWIPPactsSince: nil,
		// PactLogDir:           "",
		// PactLogLevel: "",
	})
	Expect(xx).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())
	pub := dsl.Publisher{}
	fmt.Println("Publishing Pact files to broker", os.Getenv("PACT_DIR"), os.Getenv("PACT_BROKER_URL"))
	err = pub.Publish(types.PublishRequest{
		PactURLs: []string{"./packs"},
		//BrokerUsername:  "ryumugil@gmail.com",
		//BrokerPassword:  "Gillygilly2020*",
		//PactBroker:      "http://zagadat.pactflow.io",
		//BrokerToken:     "aCY3pe7uIKJLLVh-I28R8A",
		PactBroker:      "http://localhost",
		BrokerUsername:  "",
		BrokerPassword:  "",
		ConsumerVersion: "1.0.0",
		Tags:            []string{"dev", "stage"},
	})
}
