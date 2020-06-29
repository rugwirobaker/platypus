package platypus_test

import (
	"context"
	"testing"

	"github.com/rugwirobaker/platypus"
)

const prefix = "*662*104#"

func TestMux(t *testing.T) {

	mux := platypus.New(prefix, platypus.NotFoundHandler())

	mux.Handle("*662*104", platypus.HandlerFunc(h), platypus.TrimTrailHash)
	mux.Handle("*662*104*1", platypus.HandlerFunc(h1), platypus.TrimTrailHash)
	mux.Handle("*662*104*1*:phone#", platypus.HandlerFunc(h2), nil)

	mux.Handle("*662*104*2", platypus.HandlerFunc(h1), platypus.TrimTrailHash)
	mux.Handle("*662*104*2*:name*1#", platypus.HandlerFunc(h3), nil)
	mux.Handle("*662*104*2*:name", platypus.HandlerFunc(h3), platypus.TrimTrailHash)

	mux.HandlerFunc("*662*104*3#", h1, nil)

	cases := []struct {
		desc string
		cmd  string
		res  string
		end  bool
	}{
		{desc: "1", cmd: "*662*104#", res: "main", end: false},
		{desc: "2", cmd: "*662*104*1", res: "*662*104*1", end: false},
		{desc: "3", cmd: "*662*104*1*0784675205#", res: "0784675205", end: true},
		{desc: "4", cmd: "*662*104*2*james#", res: "james", end: false},
		{desc: "5", cmd: "*662*104*2*james*1#", res: "james", end: true},
		{desc: "6", cmd: "*662*104*3#", res: "*662*104*3#", end: true},
		{desc: "6", cmd: "*662*100*4#", res: "undefined", end: false},
	}

	for _, tc := range cases {
		cmd := &platypus.Command{tc.cmd}

		res, err := mux.Process(context.Background(), cmd)
		if err != nil {
			t.Error(err)
		}
		if tc.res != res.String() {
			t.Errorf("desc '%s': output-> expected '%s' got '%s'", tc.desc, tc.res, res)
		}

		if res.Tail() != tc.end {
			t.Errorf("desc '%s': leafness-> expected '%v' got '%v'", tc.desc, tc.end, res.Tail())
		}
	}
}

func TestParamsFromContext(t *testing.T) {

	cases := []struct {
		desc string
		val  interface{}
		err  error
	}{
		{desc: "string", val: "value", err: nil},
		{desc: "int", val: 10000, err: nil},
		{desc: "bool", val: false, err: nil},
	}

	for _, tc := range cases {
		ctx := context.Background()

		ctx = platypus.ContextWithParams(ctx, map[string]interface{}{tc.desc: tc.val})

		params := platypus.ParamsFromContext(ctx)

		switch tc.desc {
		case "string":
			val, err := params.GetString(tc.desc)
			if err != tc.err {
				t.Errorf("desc '%s': unexpected error:'%v' expected: '%v'", tc.desc, err, tc.err)
			}
			if val != tc.val.(string) {
				t.Errorf("desc '%s': expected '%s' got '%s'", tc.desc, tc.val.(string), val)
			}
		case "int":
			val, err := params.GetInt(tc.desc)
			if err != tc.err {
				t.Errorf("desc '%s': unexpected error:'%v' expected: '%v'", tc.desc, err, tc.err)
			}
			if val != tc.val.(int) {
				t.Errorf("desc '%s': expected '%d' got '%d'", tc.desc, tc.val.(int), val)
			}
		case "bool":
			val, err := params.GetBool(tc.desc)
			if err != tc.err {
				t.Errorf("desc '%s': unexpected error:'%v' expected: '%v'", tc.desc, err, tc.err)
			}
			if val != tc.val.(bool) {
				t.Errorf("desc '%s': expected '%v' got '%v'", tc.desc, tc.val.(bool), val)
			}
		}
	}
}

var h = func(ctx context.Context, cmd *platypus.Command) (platypus.Result, error) {
	params := platypus.ParamsFromContext(ctx)

	leafness, _ := params.GetBool("isleaf")

	return platypus.Result{
		Out:  "main",
		Leaf: leafness,
	}, nil
}

var h1 = func(ctx context.Context, cmd *platypus.Command) (platypus.Result, error) {
	params := platypus.ParamsFromContext(ctx)

	leafness, _ := params.GetBool("isleaf")
	return platypus.Result{
		Out:  cmd.Pattern,
		Leaf: leafness,
	}, nil
}

var h2 = func(ctx context.Context, cmd *platypus.Command) (platypus.Result, error) {
	params := platypus.ParamsFromContext(ctx)

	phone, _ := params.GetString("phone")
	leafness, _ := params.GetBool("isleaf")

	return platypus.Result{
		Out:  phone,
		Leaf: leafness,
	}, nil
}

var h3 = func(ctx context.Context, cmd *platypus.Command) (platypus.Result, error) {
	params := platypus.ParamsFromContext(ctx)

	name, _ := params.GetString("name")
	leafness, _ := params.GetBool("isleaf")

	return platypus.Result{
		Out:  name,
		Leaf: leafness,
	}, nil
}
