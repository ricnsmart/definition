package defination

type (
	Resource struct {
		Mongodb  mongodb
		RabbitMQ rabbitMQ
		Redis    redis
		Postgres postgres
		AliSms   aliSms
		AliDns   aliDns
		AliVms   aliVms
	}

	mongodb struct {
		Address string
		DbName  string
	}

	rabbitMQ struct {
		Address  string
		UserName string
		Password string
	}

	redis struct {
		Address  string
		Password string
	}

	postgres struct {
		Host     string
		Port     int
		User     string
		Password string
		DbName   string
	}

	aliSms struct {
		AccessKeyID     string
		AccessKeySecret string
		SignName        string
	}

	aliDns struct {
		AccessKeyID     string
		AccessKeySecret string
	}

	aliVms struct {
		AccessKeyID     string
		AccessKeySecret string
	}
)
