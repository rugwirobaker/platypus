package platypus_test

import (
	"context"
	"testing"

	"github.com/rugwirobaker/platypus"
)

type result struct {
	out string
	end bool
}

func (res result) String() string {
	return res.out
}

func (res result) Tail() bool {
	return res.end
}

func TestMux(t *testing.T) {
	h := func(ctx context.Context, cmd *platypus.Command) (platypus.Result, error) {
		return result{out: "not found"}, nil
	}

	h1 := func(ctx context.Context, cmd *platypus.Command) (platypus.Result, error) {
		return result{out: cmd.Pattern}, nil
	}

	h2 := func(ctx context.Context, cmd *platypus.Command) (platypus.Result, error) {
		params := platypus.ParamsFromContext(ctx)
		return result{out: params.GetString("phone")}, nil
	}

	h3 := func(ctx context.Context, cmd *platypus.Command) (platypus.Result, error) {
		params := platypus.ParamsFromContext(ctx)

		return result{
			out: "ok",
			end: params.GetBool("isleaf"),
		}, nil
	}

	mux := platypus.NewMux("*662", platypus.HandlerFunc(h))

	mux.Handle("*662*1", platypus.HandlerFunc(h1))
	mux.Handle("*662*2", platypus.HandlerFunc(h1))
	mux.Handle("*662*2*:phone", platypus.HandlerFunc(h2))
	mux.Handle("*662*3", platypus.HandlerFunc(h3))

	cases := []struct {
		desc string
		cmd  string
		res  string
		end  bool
	}{
		{desc: "1", cmd: "*662*1#", res: "*662*1", end: false},
		{desc: "2", cmd: "*662*2*0784675205#", res: "0784675205", end: false},
		{desc: "3", cmd: "*662*3#", res: "ok", end: true},
	}

	for _, tc := range cases {
		cmd := &platypus.Command{tc.cmd}

		res, err := mux.Process(context.Background(), cmd)
		if err != nil {
			t.Error(err)
		}
		if tc.res != res.String() {
			t.Errorf("desc '%s': expected '%s' got '%s'", tc.desc, tc.res, res)
		}

		if res.Tail() != tc.end {
			t.Errorf("desc '%s': expected '%v' got '%v'", tc.desc, tc.end, res.Tail())
		}
	}
}

func TestParams(t *testing.T) {

	cases := []struct {
		desc string
		val  interface{}
	}{
		{desc: "string", val: "value"},
		{desc: "int", val: 10000},
		{desc: "bool", val: false},
	}

	for _, tc := range cases {
		ctx := context.Background()

		ctx = platypus.ContextWithParams(ctx, map[string]interface{}{tc.desc: tc.val})

		params := platypus.ParamsFromContext(ctx)

		switch tc.desc {
		case "string":
			val := params.GetString(tc.desc)
			if val != tc.val.(string) {
				t.Errorf("desc '%s': expected '%s' got '%s'", tc.desc, tc.val.(string), val)
			}
		case "int":
			val := params.GetInt(tc.desc)
			if val != tc.val.(int) {
				t.Errorf("desc '%s': expected '%d' got '%d'", tc.desc, tc.val.(int), val)
			}
		case "bool":
			val := params.GetBool(tc.desc)
			if val != tc.val.(bool) {
				t.Errorf("desc '%s': expected '%v' got '%v'", tc.desc, tc.val.(bool), val)
			}
		}
	}
}
