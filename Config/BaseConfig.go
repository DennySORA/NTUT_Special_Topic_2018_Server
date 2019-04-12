package Config

import "time"

const GraphQLDefaultPort = "8081"

const StartGraphQLServer = true

const DatabaseURL = "127.0.0.1:27017"

const DatabaseName = "Test"

const DebugLevel = 0

const TokenNumber = 5

const AccountSetTime int64 = 60 * 60 * 24

const TokenCheckTime time.Duration = 5 * time.Second

const TokenTime = 30
