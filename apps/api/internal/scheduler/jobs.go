package scheduler

import "context"

type ReminderJob struct{}

func (j *ReminderJob) Name() string { return "reminder" }
func (j *ReminderJob) Execute(_ context.Context) error {
	return nil
}

type WeeklySummaryJob struct{}

func (j *WeeklySummaryJob) Name() string { return "weekly_summary" }
func (j *WeeklySummaryJob) Execute(_ context.Context) error {
	return nil
}

type ClimateAlertJob struct{}

func (j *ClimateAlertJob) Name() string { return "climate_alert" }
func (j *ClimateAlertJob) Execute(_ context.Context) error {
	return nil
}
