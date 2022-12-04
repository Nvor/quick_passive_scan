package run_settings

type Settings struct {
	Mode Mode
}

type Mode struct {
	Description string
	PingCount int
	MaxPort int
}
