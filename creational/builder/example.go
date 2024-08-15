package main

import "fmt"

func main() {
	bldr := newNotificationBuilder()
	bldr.SetTitle("New notification")
	bldr.SetSubTitle("example of builder")
	bldr.SetIcon("icon.png")
	bldr.SetImage("image.png")
	bldr.SetPriority(1)
	bldr.SetMessage("Basic notification")
	bldr.SetType("alert")

	notif, err := bldr.Build()
	if err != nil {
		fmt.Println("Error creating notification:", err)
		return
	}
	fmt.Printf("Notification: %+v\n", notif)
}
