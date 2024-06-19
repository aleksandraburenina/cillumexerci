import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"example.com/GoogleCloudPlatform/golang-samples/run/functions-framework-go/functions"
)

// HelloFirestore is an example of a Cloud Function that uses Firestore.
func HelloFirestore(ctx context.Context, e functions.Event) error {
	client, err := firestore.NewClient(ctx, "project-id")
	if err != nil {
		log.Fatalf("firestore.NewClient: %v", err)
	}
	defer client.Close()

	_, _, err = client.Collection("messages").Add(ctx, map[string]interface{}{
		"message":  "Hello, world!",
		"timestamp": time.Now(),
	})
	if err != nil {
		log.Fatalf("Add: %v", err)
	}

	fmt.Fprint(e, "Hello, Firestore!")
	return nil
}
  
