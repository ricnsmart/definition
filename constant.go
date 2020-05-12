package definition

const (
	/*Regexp*/
	OnlyDomainRecordRegexp = `^[-a-zA-Z0-9.]+.[a-z]+$`

	HasDomainRecordRegexp = `[-a-zA-Z0-9.]+.[a-z]+`

	HasNumRegexp = `[0-9]+`

	HasIp = `(1\d\d|2[0-5][0-5]|[1-9]\d|[1-9]).((1\d\d|2[0-5][0-5]|[1-9]\d|\d).){2}(1\d\d|2[0-5][0-5]|[1-9]\d|\d)`

	OnlyIP = `^(1\d\d|2[0-5][0-5]|[1-9]\d|[1-9]).((1\d\d|2[0-5][0-5]|[1-9]\d|\d).){2}(1\d\d|2[0-5][0-5]|[1-9]\d|\d)$`

	HasDomainRecordOrIP = `[-a-zA-Z0-9.]+.[a-z]+|(1\d\d|2[0-5][0-5]|[1-9]\d|[1-9]).((1\d\d|2[0-5][0-5]|[1-9]\d|\d).){2}(1\d\d|2[0-5][0-5]|[1-9]\d|\d)`

	OnlyIPOrDomainRecord = `^([-a-zA-Z0-9.]+.[a-z]+|(1\d\d|2[0-5][0-5]|[1-9]\d|[1-9]).((1\d\d|2[0-5][0-5]|[1-9]\d|\d).){2}(1\d\d|2[0-5][0-5]|[1-9]\d|\d))$`
)
