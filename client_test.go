package waterlink

import (
	"fmt"
	"log"
)

func ExampleNewClient() {
	creds := Credentials{
		Authorization: "youshallnotpass", // passphrase defined in the server's application.yml
	}

	cl, err := NewClient("http://localhost:2333", creds)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Client dispatches actions to %q.", cl.url.String())

	// Output:
	// Client dispatches actions to "http://localhost:2333".
}
