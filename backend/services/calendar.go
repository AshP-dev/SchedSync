package services

import (
    "context"
    "log"
    "google.golang.org/api/calendar/v3"
    "google.golang.org/api/option"
	"golang.org/x/oauth2"
	//"ankified_planner/utils"
)

// GetGoogleCalendarEvents retrieves events from the user's Google Calendar
func GetGoogleCalendarEvents(ctx context.Context, token string) ([]*calendar.Event, error) {
    srv, err := calendar.NewService(ctx, option.WithTokenSource(oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})))
    if err != nil {
        log.Fatalf("Unable to retrieve Calendar client: %v", err)
        return nil, err
    }

    events, err := srv.Events.List("primary").Do()
    if err != nil {
        log.Fatalf("Unable to retrieve events: %v", err)
        return nil, err
    }

    return events.Items, nil
}