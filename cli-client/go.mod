module gitlab.com/Prosp3r/ori/cli-client

go 1.13

replace gitlab.com/Prosp3r/ori/pb => ../pb

require (
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v0.0.6
	github.com/spf13/viper v1.6.2
	gitlab.com/Prosp3r/ori/pb v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.28.0
)
