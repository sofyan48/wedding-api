module github.com/orn-id/wedding-api

go 1.14

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/Baozisoftware/qrcode-terminal-go v0.0.0-20170407111555-c0650d8dff0f // indirect
	github.com/Rhymen/go-whatsapp v0.1.1
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/aws/aws-sdk-go v1.31.9
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/getsentry/sentry-go v0.6.1
	github.com/go-playground/validator/v10 v10.4.1
	github.com/go-redis/redis v6.15.8+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/uuid v1.1.2
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/schema v1.2.0
	github.com/jmoiron/sqlx v1.2.0
	github.com/joho/godotenv v1.3.0
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pressly/goose v2.7.0-rc4.0.20200419115850-2769ad7b550e+incompatible
	github.com/sirupsen/logrus v1.6.0
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e // indirect
	github.com/spf13/cast v1.3.0
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.6.0
	github.com/swaggo/http-swagger v0.0.0-20200308142732-58ac5e232fba
	github.com/swaggo/swag v1.6.3
	github.com/ziutek/mymysql v1.5.4
	golang.org/x/net v0.0.0-20200528225125-3c3fba18258b // indirect
	golang.org/x/sys v0.0.0-20210927094055-39ccf1dd6fa6 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
)

replace github.com/Rhymen/go-whatsapp => ./go-whatsapp
