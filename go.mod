module github.com/hpb-project/votedapp-server

go 1.15

require (
	github.com/ethereum/go-ethereum v1.10.9
	github.com/gin-gonic/gin v1.7.4
	github.com/gorilla/websocket v1.4.2
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	gorm.io/driver/mysql v1.1.2
	gorm.io/gorm v1.21.16
	gotest.tools v2.2.0+incompatible
)
