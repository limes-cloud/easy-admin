/**
 * @Author: 1280291001@qq.com
 * @Description: jwt验证失败error
 * @File: jwtErr
 * @Version: 1.0.0
 * @Date: 2023/4/18 22:30
 */

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

type jwtErr struct {
	err       error
	isVerify  bool
	isExpired bool
	exp       int64
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
func (e jwtErr) CanRenewal(t int64) bool {
	return time.Now().Unix()-e.exp <= t
}
