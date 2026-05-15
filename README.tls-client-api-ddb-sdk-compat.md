# tls-client-api-ddb-sdk-compat branch

Single-commit branch on top of the Solem8s/gosoline tls-client-api fork
(commit `4cf4b1d`) that bridges a Go type-signature change in newer
`aws-sdk-go-v2/service/dynamodb` releases.

## The change

Newer DDB SDK versions (~v1.20+) changed
`DescribeTableOutput.Table.ItemCount` from `int64` to `*int64`.
`pkg/ddb/service.go` line 63 reads this field directly. Without the
patch, downstream consumers that elevate the SDK via Go module MVS hit
a compile error on this line.

The patch is one line: wrap the field access in `aws.ToInt64`, which
dereferences the pointer safely (returns 0 if nil).

## Stability

This branch is tagged with an immutable Git tag for downstream `go.mod`
`replace` pins. See the tag list on GitHub for the exact tag name.
Force-pushes to the branch are not expected; tag the commit before
pinning.

## Standalone build

The fork's `go.mod` continues to pin the older SDK, so building this
module *in isolation* will fail the `aws.ToInt64` call (the wrapped
operand is still `int64`, not `*int64`). This is intentional — the
patch is consumed only via MVS-elevated builds from downstream
consumers that require a newer DDB SDK. Bumping the fork's own
`go.mod` would cascade into every AWS service the fork pins; the
single-line trade-off avoids that.
