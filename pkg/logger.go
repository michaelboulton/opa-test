package pkg

import "go.uber.org/zap"

var (
	Logger *zap.SugaredLogger
)

func init() {
	zapper, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	Logger = zapper.Sugar()
}
