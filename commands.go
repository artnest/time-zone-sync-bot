package main

import (
	"fmt"
)

const help = `Available commands:
/add_timezone label timezone — Add time zone (TZ database name).
/update_timezone label new_timezone — Update time zone.
/remove_timezone label — Remove time zone.
/clear_timezones — Remove all time zones.
/time [format] — Show time in time zones. Available formats: full, short.
/timezones — Show time zones.
/help — Show help message.`

func HelpCommand() string {
	return help
}

func AddTimeZoneCommand(chatID int64, args AddTimeZoneArguments) string {
	if _, err := GetLocation(args.Location); err != nil {
		return fmt.Sprintf("Bad location: %v\n", err)
	}

	timeZone := TimeZone{
		Label:    args.Label,
		Location: args.Location,
	}

	if err := InsertTimeZone(chatID, timeZone); err != nil {
		return err.Error()
	}

	return fmt.Sprintf("New time zone with label %q was added.\n", timeZone.Label)
}

func UpdateTimeZoneCommand(chatID int64, args UpdateTimeZoneArguments) string {
	if _, err := GetLocation(args.Location); err != nil {
		return fmt.Sprintf("Bad location: %v\n", err)
	}

	timeZone := TimeZone{
		Label:    args.Label,
		Location: args.Location,
	}

	if err := UpdateTimeZone(chatID, timeZone); err != nil {
		return err.Error()
	}

	return fmt.Sprintf("Time zone with label %q was updated.\n", timeZone.Label)
}

func RemoveTimeZoneCommand(chatID int64, args RemoveTimeZoneArguments) string {
	if err := DeleteTimeZone(chatID, args.Label); err != nil {
		return err.Error()
	}

	return fmt.Sprintf("Time zone with label %q was removed.\n", args.Label)
}

func ClearTimeZonesCommand(chatID int64) string {
	if err := DeleteTimeZones(chatID); err != nil {
		return err.Error()
	}

	return "All time zones were cleared."
}

func TimeCommand(chatID int64, args TimeArguments) string {
	timeZones, err := SelectTimeZones(chatID)
	if err != nil {
		return err.Error()
	}

	if len(timeZones) == 0 {
		return "No time zones to show time."
	}

	return DisplayTime(timeZones, args.Format)
}

func TimeZonesCommand(chatID int64) string {
	timeZones, err := SelectTimeZones(chatID)
	if err != nil {
		return err.Error()
	}

	if len(timeZones) == 0 {
		return "No time zones to show."
	}

	return DisplayTimeZones(timeZones)
}

func UnknownCommand() string {
	return "Unknown command. User /help command."
}
