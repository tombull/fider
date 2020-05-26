package email_test

import (
	"strings"
	"testing"

	"github.com/tombull/teamdream/app/models/dto"
	"github.com/tombull/teamdream/app/services/email"

	. "github.com/tombull/teamdream/app/pkg/assert"
)

func ReplaceSpaces(inString string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(inString, "\t", ""), " ", ""), "\n", ""), "/>", ">")
}
func TestRenderMessage(t *testing.T) {
	RegisterT(t)

	message := email.RenderMessage("echo_test", dto.Props{
		"name": "Teamdream",
	})
	Expect(message.Subject).Equals("Message to: Teamdream")
	Expect(ReplaceSpaces(message.Body)).Equals(ReplaceSpaces(`<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
		<meta name="viewport" content="width=device-width">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
	</head>
	<body bgcolor="#F7F7F7" style="font-size:16px">
		<table width="100%" bgcolor="#F7F7F7" cellpadding="0" cellspacing="0" border="0" style="text-align:center;font-size:14px;">
			<tr>
				<td height="40">&nbsp;</td>
			</tr>

			<tr>
				<td align="center">
					<table bgcolor="#FFFFFF" cellpadding="0" cellspacing="0" border="0" style="text-align:left;padding:20px;margin:10px;border-radius:5px;color:#1c262d;border:1px solid #ECECEC;min-width:320px;max-width:660px;">
						Hello World Teamdream!
					</table>
				</td>
			</tr>
			<tr>
				<td>
					<span style="color:#666;font-size:11px">This email was sent from a notification-only address that cannot accept incoming email. Please do not reply to this message.</span>
				</td>
			</tr>
			<tr>
				<td height="40">&nbsp;</td>
			</tr>
		</table>
	</body>
</html>`))
}

func TestCanSendTo(t *testing.T) {
	RegisterT(t)

	testCases := []struct {
		whitelist string
		blacklist string
		input     []string
		canSend   bool
	}{
		{
			whitelist: "(^.+@teamdream.co.uk$)|(^darthvader\\.teamdream(\\+.*)?@gmail\\.com$)",
			blacklist: "",
			input:     []string{"me@teamdream.co.uk", "me+123@teamdream.co.uk", "darthvader.teamdream@gmail.com", "darthvader.teamdream+434@gmail.com"},
			canSend:   true,
		},
		{
			whitelist: "(^.+@teamdream.co.uk$)|(^darthvader\\.teamdream(\\+.*)?@gmail\\.com$)",
			blacklist: "",
			input:     []string{"me+123@teamdream.co.ukd", "me@teamdreamo.io", "darthvader.teamdreama@gmail.com", "@teamdream.co.uk"},
			canSend:   false,
		},
		{
			whitelist: "(^.+@teamdream.co.uk$)|(^darthvader\\.teamdream(\\+.*)?@gmail\\.com$)",
			blacklist: "(^.+@teamdream.co.uk$)",
			input:     []string{"me@teamdream.co.uk"},
			canSend:   true,
		},
		{
			whitelist: "",
			blacklist: "(^.+@teamdream.co.uk$)",
			input:     []string{"me@teamdream.co.uk", "abc@teamdream.co.uk"},
			canSend:   false,
		},
		{
			whitelist: "",
			blacklist: "(^.+@teamdream.co.uk$)",
			input:     []string{"me@teamdream.com", "abc@teamdreamio.io"},
			canSend:   true,
		},
		{
			whitelist: "",
			blacklist: "",
			input:     []string{"me@teamdream.co.uk"},
			canSend:   true,
		},
		{
			whitelist: "",
			blacklist: "",
			input:     []string{"", " "},
			canSend:   false,
		},
	}

	for _, testCase := range testCases {
		email.SetWhitelist(testCase.whitelist)
		email.SetBlacklist(testCase.blacklist)
		for _, input := range testCase.input {
			Expect(email.CanSendTo(input)).Equals(testCase.canSend)
		}
	}
}

func TestRecipient_String(t *testing.T) {
	RegisterT(t)

	testCases := []struct {
		name     string
		email    string
		expected string
	}{
		{
			name:     "Jon",
			email:    "jon@got.com",
			expected: `"Jon" <jon@got.com>`,
		},
		{
			name:     "Snow, Jon",
			email:    "jon@got.com",
			expected: `"Snow, Jon" <jon@got.com>`,
		},
		{
			name:     "",
			email:    "jon@got.com",
			expected: "<jon@got.com>",
		},
		{
			name:     "Jon's Home Account",
			email:    "jon@got.com",
			expected: `"Jon's Home Account" <jon@got.com>`,
		},
		{
			name:     `Jon "Great" Snow`,
			email:    "jon@got.com",
			expected: `"Jon \"Great\" Snow" <jon@got.com>`,
		},
		{
			name:     "Jon",
			email:    "",
			expected: "",
		},
	}

	for _, testCase := range testCases {
		r := dto.NewRecipient(testCase.name, testCase.email, dto.Props{})
		Expect(r.String()).Equals(testCase.expected)
	}
}
