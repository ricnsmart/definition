package definition

type (
	Resources struct {
		Mongodb  Mongodb
		RabbitMQ RabbitMQ
		Redis    Redis
		Postgres Postgres
		AliSms   AliSms
		AliDns   AliDns
		AliVms   AliVms
	}

	Mongodb struct {
		Address string
		DbName  string
	}

	RabbitMQ struct {
		Address string
	}

	Redis struct {
		Address  string
		Password string
	}

	Postgres struct {
		Address string
	}

	AliSms struct {
		AccessKeyID     string
		AccessKeySecret string
		SignName        string
	}

	AliDns struct {
		AccessKeyID     string
		AccessKeySecret string
	}

	AliVms struct {
		AccessKeyID     string
		AccessKeySecret string
	}
)
