// Import the required packages.
package main

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func main() {
	// Connect to the Firebase project using the project's credentials.
	ctx := context.Background()
	// The Firebase project's credentials.
	opt := option.WithCredentialsFile("payx-233fd-firebase-adminsdk-ugnsx-5b951ea23c.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	// Get the FCM service from the Firebase project.
	fcm, err := app.Messaging(ctx)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	// The registration token for the device to which you want to send the push notification.
	token := "your-device-registration-token"

	// The push notification payload.
	payload := map[string]string{
		"title": "Hello, World!",
		"body":  "This is a test push notification.",
	}

	// Send push notification to multiple device
	response1, err1 := fcm.SendMulticast(ctx, &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: "Congratulations!!",
			Body:  "You have just implemented push notification",
		},
		// Token: deviceTokens,
		// it's an array of device tokens
	})

	if err1 != nil {
		fmt.Printf("error: %v\n", err1.Error())
		return
	}

	// Send the push notification to the device.
	response, err := fcm.Send(ctx, &messaging.Message{Token: token,
		Data: payload})
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("response: %v\n", response1)
	fmt.Printf("response: %v\n", response)
}
