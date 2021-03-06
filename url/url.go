package url

import (
	"fmt"
	pkgurl "net/url"

	"github.com/kristinjeanna/redact"
)

// URLRedactor enables redacting a password and, optionally
// a username, from a URL string.
type URLRedactor struct {
	passwordReplacement string
	usernameReplacement *string
}

// New returns a new URLRedactor.
func New(passwordReplacement string, usernameReplacement *string) redact.Redactor {
	return URLRedactor{
		passwordReplacement: passwordReplacement,
		usernameReplacement: usernameReplacement,
	}
}

// Redact simply returns the replacement text for any string passed to it.
func (r URLRedactor) Redact(url string) (string, error) {
	u, err := pkgurl.ParseRequestURI(url)
	if err != nil {
		return "", fmt.Errorf("failed to parse uri: %s", err)
	}

	if u.User != nil {
		_, hasPW := u.User.Password()

		switch hasPW {
		case true:
			if r.usernameReplacement == nil {
				u.User = pkgurl.UserPassword(u.User.Username(), r.passwordReplacement)
			} else if r.usernameReplacement != nil {
				u.User = pkgurl.UserPassword(*r.usernameReplacement, r.passwordReplacement)
			}
		case false:
			if r.usernameReplacement != nil {
				u.User = pkgurl.User(*r.usernameReplacement)
			}
		}
	}

	return u.String(), nil
}

// String returns a text representation of the redactor.
func (r URLRedactor) String() string {
	return fmt.Sprintf("{passwordReplacement=%q, usernameReplacement=%v}",
		r.passwordReplacement, r.usernameReplacement)
}
