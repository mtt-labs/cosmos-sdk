package errors

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestABCInfo(t *testing.T) {
	cases := map[string]struct {
		err       error
		debug     bool
		wantCode  uint32
		wantSpace string
		wantLog   string
	}{
		"plain SDK error": {
			err:       ErrUnauthorized,
			debug:     false,
			wantLog:   "unauthorized",
			wantCode:  ErrUnauthorized.code,
			wantSpace: testCodespace,
		},
		"wrapped SDK error": {
			err:       Wrap(Wrap(ErrUnauthorized, "foo"), "bar"),
			debug:     false,
			wantLog:   "bar: foo: unauthorized",
			wantCode:  ErrUnauthorized.code,
			wantSpace: testCodespace,
		},
		"nil is empty message": {
			err:       nil,
			debug:     false,
			wantLog:   "",
			wantCode:  0,
			wantSpace: "",
		},
		"nil SDK error is not an error": {
			err:       (*Error)(nil),
			debug:     false,
			wantLog:   "",
			wantCode:  0,
			wantSpace: "",
		},
		"stdlib returns error message in debug mode": {
			err:       io.EOF,
			debug:     true,
			wantLog:   "EOF",
			wantCode:  1,
			wantSpace: UndefinedCodespace,
		},
		// This is hard to test because of attached stacktrace. This
		// case is tested in an another test.
		// "wrapped stdlib is a full message in debug mode": {
		//	err:      Wrap(io.EOF, "cannot read file"),
		//	debug:    true,
		//	wantLog:  "cannot read file: EOF",
		//	wantCode: 1,
		// },
		"custom error": {
			err:       customErr{},
			debug:     false,
			wantLog:   "custom",
			wantCode:  999,
			wantSpace: "extern",
		},
		"custom error in debug mode": {
			err:       customErr{},
			debug:     true,
			wantLog:   "custom",
			wantCode:  999,
			wantSpace: "extern",
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			space, code, log := ABCIInfo(tc.err, tc.debug)
			if space != tc.wantSpace {
				t.Errorf("%s: expected space %s, got %s", testName, tc.wantSpace, space)
			}
			if code != tc.wantCode {
				t.Errorf("%s: expected code %d, got %d", testName, tc.wantCode, code)
			}
			if log != tc.wantLog {
				t.Errorf("%s: expected log %s, got %s", testName, tc.wantLog, log)
			}
		})
	}
}

func TestABCIInfoStacktrace(t *testing.T) {
	cases := map[string]struct {
		err            error
		debug          bool
		wantStacktrace bool
		wantErrMsg     string
	}{
		"wrapped SDK error in debug mode provides stacktrace": {
			err:            Wrap(ErrUnauthorized, "wrapped"),
			debug:          true,
			wantStacktrace: true,
			wantErrMsg:     "wrapped: unauthorized",
		},
		"wrapped SDK error in non-debug mode does not have stacktrace": {
			err:            Wrap(ErrUnauthorized, "wrapped"),
			debug:          false,
			wantStacktrace: false,
			wantErrMsg:     "wrapped: unauthorized",
		},
		"wrapped stdlib error in debug mode provides stacktrace": {
			err:            Wrap(fmt.Errorf("stdlib"), "wrapped"),
			debug:          true,
			wantStacktrace: true,
			wantErrMsg:     "wrapped: stdlib",
		},
	}

	const thisTestSrc = "cosmossdk.io/errors.(*abciTestSuite).TestABCIInfoStacktrace"

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			_, _, log := ABCIInfo(tc.err, tc.debug)
			if !tc.wantStacktrace {
				if log != tc.wantErrMsg {
					t.Errorf("%s: expected log %s, got %s", testName, tc.wantErrMsg, log)
				}
			} else {
				if !strings.Contains(log, thisTestSrc) {
					t.Errorf("%s: expected log to contain %s", testName, thisTestSrc)
				}
				if !strings.Contains(log, tc.wantErrMsg) {
					t.Errorf("%s: expected log to contain %s", testName, tc.wantErrMsg)
				}
			}
		})
	}
}

func TestABCIInfoHidesStacktrace(t *testing.T) {
	err := Wrap(ErrUnauthorized, "wrapped")
	_, _, log := ABCIInfo(err, false)
	if log != "wrapped: unauthorized" {
		t.Errorf("expected log %s, got %s", "wrapped: unauthorized", log)
	}
}

func TestABCIInfoSerializeErr(t *testing.T) {
	var (
		// Create errors with stacktrace for equal comparison.
		myErrDecode = Wrap(ErrTxDecode, "test")
		myErrAddr   = Wrap(ErrInvalidAddress, "tester")
		myPanic     = ErrPanic
	)

	specs := map[string]struct {
		src   error
		debug bool
		exp   string
	}{
		"single error": {
			src:   myErrDecode,
			debug: false,
			exp:   "test: tx parse error",
		},
		"second error": {
			src:   myErrAddr,
			debug: false,
			exp:   "tester: invalid address",
		},
		"single error with debug": {
			src:   myErrDecode,
			debug: true,
			exp:   fmt.Sprintf("%+v", myErrDecode),
		},
		"do not redact in debug encoder": {
			src:   myPanic,
			debug: true,
			exp:   fmt.Sprintf("%+v", myPanic),
		},
	}
	for msg, spec := range specs {
		spec := spec
		_, _, log := ABCIInfo(spec.src, spec.debug)
		if log != spec.exp {
			t.Errorf("%s: expected log %s, got %s", msg, spec.exp, log)
		}
	}
}

// customErr is a custom implementation of an error that provides an ABCICode
// method.
type customErr struct{}

func (customErr) Codespace() string { return "extern" }

func (customErr) ABCICode() uint32 { return 999 }

func (customErr) Error() string { return "custom" }
