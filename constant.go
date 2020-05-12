package definition

const (
	/*Pattern*/
	OnlyDomainRecordPattern = `^[-a-zA-Z0-9.]+.[a-z]+$`

	ExistDomainRecordPattern = `[-a-zA-Z0-9.]+.[a-z]+`

	ExistNumPattern = `[0-9]+`

	ExistIpPattern = `(1\d\d|2[0-5][0-5]|[1-9]\d|[1-9]).((1\d\d|2[0-5][0-5]|[1-9]\d|\d).){2}(1\d\d|2[0-5][0-5]|[1-9]\d|\d)`

	OnlyIpPattern = `^(1\d\d|2[0-5][0-5]|[1-9]\d|[1-9]).((1\d\d|2[0-5][0-5]|[1-9]\d|\d).){2}(1\d\d|2[0-5][0-5]|[1-9]\d|\d)$`

	ExistDomainRecordOrIpPattern = `[-a-zA-Z0-9.]+.[a-z]+|(1\d\d|2[0-5][0-5]|[1-9]\d|[1-9]).((1\d\d|2[0-5][0-5]|[1-9]\d|\d).){2}(1\d\d|2[0-5][0-5]|[1-9]\d|\d)`

	OnlyIpOrDomainRecordPattern = `^([-a-zA-Z0-9.]+.[a-z]+|(1\d\d|2[0-5][0-5]|[1-9]\d|[1-9]).((1\d\d|2[0-5][0-5]|[1-9]\d|\d).){2}(1\d\d|2[0-5][0-5]|[1-9]\d|\d))$`
)
