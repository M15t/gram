package firebase

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	auth "firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

// Config represents the configuration
type Config struct {
	FirebaseCredentials string
}

// Service represents the firebase service
type Service struct {
	cfg *Config
	cr  Crypter
	ctx context.Context
	app *firebase.App
	ac  *auth.Client
	fsc *firestore.Client
}

// New initializes firebase service with default config
func New(cfg *Config, cr Crypter) *Service {
	// * Initialize Firebase Admin SDK with the fetched credentials
	opt := option.WithCredentialsJSON([]byte(cfg.FirebaseCredentials))
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// * Initialize Auth client
	ac, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error initializing Auth client: %v\n", err)
	}

	// * Initialize Firestore client
	fsc, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error initializing Firestore client: %v\n", err)
	}
	// ! gonna monitor if have more issues
	// defer fsc.Close()

	return &Service{
		cfg: cfg,
		cr:  cr,
		ctx: ctx,
		app: app,
		ac:  ac,
		fsc: fsc,
	}
}

// Crypter represents security and stuff interface
type Crypter interface {
	UID() string
	GetCurrentUnixTime() int64
	ScheduleToReset() int64
}
