package waterlink

import (
	"fmt"
	"github.com/lukasl-dev/waterlink/v2/event"
	"log"
	"time"
)

func ExampleOpen() {
	creds := Credentials{
		Authorization: "youshallnotpass",  // passphrase defined in the server's application.yml
		UserID:        820112919224124416, // the user ID of the bot user
	}

	conn, err := Open("ws://localhost:2333", creds)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("The server is running on version %q.", conn.APIVersion())

	// Output:
	// The server is running on version "3".
}

func ExampleOpen_listenToEvents() {
	creds := Credentials{
		Authorization: "youshallnotpass",  // passphrase defined in the server's application.yml
		UserID:        820112919224124416, // the user ID of the bot user
	}

	events := EventHandlerFunc(func(evt interface{}) {
		switch evt := evt.(type) {
		case event.Stats:
			fmt.Println("Stats received:", evt)
		case event.TrackEnd:
			fmt.Println("Track ended:", evt)
		}
	})
	opts := ConnectionOptions{EventHandler: events}

	conn, err := Open("ws://localhost:2333", creds, opts)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("The server is running on version %q.", conn.APIVersion())

	// Output:
	// The server is running on version "3".
}

func ExampleOpen_resumeSession() {
	creds := Credentials{
		Authorization: "youshallnotpass",  // passphrase defined in the server's application.yml
		UserID:        820112919224124416, // the user ID of the bot user
		ResumeKey:     "myResumeKey",      // the former configured resume key
	}

	conn, err := Open("ws://localhost:2333", creds)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("The server is running on version %q.", conn.APIVersion())

	// Output:
	// The server is running on version "3".
}

func ExampleConnection_ConfigureResuming() {
	var conn Connection // TODO: open connection

	err := conn.ConfigureResuming("myResumeKey", 1*time.Minute)
	if err != nil {
		log.Fatalln(err)
	}
}
