/*
Sniperkit-Bot
- Status: analyzed
*/

package aggregation

import (
	"time"

	"github.com/pingcap/check"

	"github.com/sniperkit/snk.fork.pingcap-tidb/sessionctx/stmtctx"
	"github.com/sniperkit/snk.fork.pingcap-tidb/types"
	"github.com/sniperkit/snk.fork.pingcap-tidb/util/testleak"
)

var _ = check.Suite(&testUtilSuite{})

type testUtilSuite struct {
}

func (s *testUtilSuite) TestDistinct(c *check.C) {
	defer testleak.AfterTest(c)()
	sc := &stmtctx.StatementContext{TimeZone: time.Local}
	dc := createDistinctChecker(sc)
	tests := []struct {
		vals   []interface{}
		expect bool
	}{
		{[]interface{}{1, 1}, true},
		{[]interface{}{1, 1}, false},
		{[]interface{}{1, 2}, true},
		{[]interface{}{1, 2}, false},
		{[]interface{}{1, nil}, true},
		{[]interface{}{1, nil}, false},
	}
	for _, tt := range tests {
		d, err := dc.Check(types.MakeDatums(tt.vals...))
		c.Assert(err, check.IsNil)
		c.Assert(d, check.Equals, tt.expect)
	}
}
