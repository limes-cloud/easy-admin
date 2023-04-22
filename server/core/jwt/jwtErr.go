package jwt

import "time"

type jwtErrOption func(err *jwtErr)

func withVerify(val bool) jwtErrOption {
	return func(err *jwtErr) {
		err.isVerify = val
	}
}

func withExpired(val bool) jwtErrOption {
	return func(err *jwtErr) {
		err.isExpired = val
	}
}

func withExpireUnix(val int64) jwtErrOption {
	return func(err *jwtErr) {
		err.exp = val
	}
}

func withRenewalUnix(val int64) jwtErrOption {
	return func(err *jwtErr) {
		err.renewal = val
	}
}

type jwtErr struct {
	err       error
	isVerify  bool
	isExpired bool
	exp       int64
	renewal   int64
}

func (e jwtErr) Error() string {
	return e.err.Error()
}

// IsVerify 是否验证失败
func (e jwtErr) IsVerify() bool {
	return e.isVerify
}

// IsExpired 是否过期
func (e jwtErr) IsExpired() bool {
	return e.isExpired
}

// CanRenewal 是否能够续期
func (e jwtErr) CanRenewal() bool {
	return time.Now().Unix()-e.exp <= e.renewal
}
