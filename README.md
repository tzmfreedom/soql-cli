[![Build Status](https://travis-ci.org/tzmfreedom/soql-cli.svg?branch=master)](https://travis-ci.org/tzmfreedom/soql-cli)

# soql-cli

Salesforce SOQL cli tool

## Installation

For Golang User
```bash
$ go get -u github.com/tzmfreedom/soql-cli
```

## Usage

```bash
$ soql -u {USERNAME} -h {ENDPOINT} [-d] [-e {COMMMAND}] [-v]
```

|option|description|alternative environment variable|
|-|-|-|
|-u|specify salesforce username|SALESFORCE_USERNAME|
|-d|dryrun mode(no write, read only)||
|-h|hostname (e.g. test.salesforce.com)|SALESFORCE_ENDPOINT|
|-e|specify command(batch mode, no interactive)||
|-v|print version||

Select Query
```
>> SELECT Id, Name FROM Account WHERE Name = 'foo'
```

Insert
```
>> INSERT INTO Account(Name, Website) VALUES ("foo", "https://github.com/tzmfreedom/soql-cli")
```

Update
```
>> UPDATE Account SET Name = 'foo', Website = 'https://github.com/tzmfreedom/soql-cli'
```

Delete
```
>> DELETE FROM Account WHERE Name = 'foo'
```

Exit console
```
>> exit # or quit
```

Batch mode
```
$ soql -e "SELECT * FROM Account"
```

DryRun mode
```
$ soql -d -e "INSERT INTO Account(Name) VALUES ("foo");
> no change database
```

## Contribute

Just send pull request if needed or fill an issue!

## License

The MIT License See [LICENSE](https://github.com/tzmfreedom/soql-cli/blob/master/LICENSE) file.

