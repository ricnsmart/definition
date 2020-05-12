package definition

const (
	/*Pattern*/
	OnlyDomainRecordPattern = `^[-a-zA-Z0-9.]+.[a-z]+$`

	HasDomainRecordPattern = `[-a-zA-Z0-9.]+.[a-z]+`

	HasNumPattern = `[0-9]+`

	HasIpPattern = `(1\d\d|2[0-5][0-5]|[1-9]\d|[1-9]).((1\d\d|2[0-5][0-5]|[1-9]\d|\d).){2}(1\d\d|2[0-5][0-5]|[1-9]\d|\d)`

	OnlyIpPattern = `^(1\d\d|2[0-5][0-5]|[1-9]\d|[1-9]).((1\d\d|2[0-5][0-5]|[1-9]\d|\d).){2}(1\d\d|2[0-5][0-5]|[1-9]\d|\d)$`

	HasDomainRecordOrIpPattern = `[-a-zA-Z0-9.]+.[a-z]+|(1\d\d|2[0-5][0-5]|[1-9]\d|[1-9]).((1\d\d|2[0-5][0-5]|[1-9]\d|\d).){2}(1\d\d|2[0-5][0-5]|[1-9]\d|\d)`

	OnlyIpOrDomainRecordPattern = `^([-a-zA-Z0-9.]+.[a-z]+|(1\d\d|2[0-5][0-5]|[1-9]\d|[1-9]).((1\d\d|2[0-5][0-5]|[1-9]\d|\d).){2}(1\d\d|2[0-5][0-5]|[1-9]\d|\d))$`
)
