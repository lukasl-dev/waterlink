package waterlink

import (
	"fmt"
	"log"
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

	events := NewChannelEventBus()
	opts := ConnectionOptions{EventBus: events}

	conn, err := Open("ws://localhost:2333", creds, opts)
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		for {
			select {
			case e := <-events.Stats():
				fmt.Println("Stats:", e)
			case e := <-events.TrackEnds():
				fmt.Println("TrackEnd:", e)
			}
		}
	}()

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
