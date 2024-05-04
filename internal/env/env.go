package env

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
	"golang.org/x/xerrors"
)

type Values struct {
	Env                string        `required:"true" split_words:"true"`
	Debug              bool          `default:"false" split_workds:"true"`
	Port               string        `default:"8080" split_words:"true"`
	ServiceName        string        `default:"project-name" split_words:"true"`
	ShutdownTimeout    time.Duration `default:"5s" split_words:"true"`
	DatabaseHost       string        `required:"true" split_words:"true"`
	DatabaseUsername   string        `required:"true" split_words:"true"`
	DatabasePassword   string        `required:"true" split_words:"true"`
	DatabaseHostRo     string        `required:"true" split_words:"true"`
	DatabaseUsernameRo string        `required:"true" split_words:"true"`
	DatabasePasswordRo string        `required:"true" split_words:"true"`
	Database           string        `required:"true" split_words:"true"`
	DatabasePort       string        `default:"3306" split_words:"true"`
}

func NewValue() (*Values, error) {
	var v Values
	err := envconfig.Process("app", &v)
	if err != nil {
		s := fmt.Sprintf("need to set all env value %+v", v)
		return nil, xerrors.Errorf(s+": %w", err)
	}

	return &v, nil
}
