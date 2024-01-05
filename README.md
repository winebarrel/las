# las

Retrieves a list of all email addresses that are on the suppression list for Amazon SES.

## Usage

```
Usage: las

Flags:
  -h, --help             Show help.
  -r, --region=STRING    The region to use ($AWS_REGION).
      --version
```

```
$ las
{"EmailAddress":"foo@example.co","LastUpdateTime":"2020-12-23T01:23:45.111Z","Reason":"BOUNCE"}
{"EmailAddress":"bar@example.co","LastUpdateTime":"2020-12-23T01:23:46.22Z","Reason":"BOUNCE"}
{"EmailAddress":"zoo@example.co","LastUpdateTime":"2020-12-23T01:23:47.3Z","Reason":"BOUNCE"}
...
```
