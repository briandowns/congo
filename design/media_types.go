package design

import (
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
)

// User is the user resource media type.
var User = MediaType("application/vnd.user+json", func() {
	Description("A user belonging to a tenant account")
	Metadata("mediatype", "123")
	Reference(UserModel)
	Attributes(func() {
		Attribute("id", Integer, "ID of user")
		Attribute("href", String, "API href of user")
		Attribute("firstname", String, "First name of user")
		Attribute("lastname", String, "Last name of user")
		Attribute("city", String, "City of residence")
		Attribute("state", String, "State of residence")
		Attribute("country", String, "Country of residence")
		Attribute("bio", String, "Biography of user")
		Attribute("email", String, "Email address of user", func() {
			Format("email")
		})
	})
	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("firstname")
		Attribute("lastname")
		Attribute("email")
		Attribute("city")
		Attribute("state")
		Attribute("country")
		Attribute("bio")
	})
	View("link", func() {
		Attribute("id")
		Attribute("href")
		Attribute("email")
	})
})

// Authorize is the authorize resource media type.
var Authorize = MediaType("application/vnd.authorize+json", func() {
	Description("Token authorization response")
	Attributes(func() {
		Attribute("access_token", String, "access token", func() {
		})
		Attribute("expires_in", Integer, "Time to expiration in seconds", func() {
		})
		Attribute("token_type", String, "type of token", func() {
		})

	})

	View("default", func() {
		Attribute("access_token")
		Attribute("expires_in")
		Attribute("token_type")
	})

})

// Login is the Login resource media type.
var Login = MediaType("application/vnd.login+json", func() {
	Description("")
	Attributes(func() {
		Attribute("email", String, "email", func() {
		})
		Attribute("password", String, "password", func() {
		})
		Attribute("application", String, "UUID of requesting application", func() {
		})

	})

	View("default", func() {
		Attribute("email")
		Attribute("password")
		Attribute("application")
	})

})

// Review is the review resource mediatype
var Review = MediaType("application/vnd.review+json", func() {
	Description("A review is submitted by a reviewer")
	Reference(ReviewModel)
	Attributes(func() {
		Attribute("id", Integer, "ID of user")
		Attribute("href", String, "API href of user")
		Attribute("comment", String, "Review comments")
		Attribute("rating", Integer, "Rating of proposal, from 1-5")
	})
	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("comment")
		Attribute("rating")
	})
	View("link", func() {
		Attribute("id")
		Attribute("href")
	})
})

// Proposal is the proposal resource mediatype
var Proposal = MediaType("application/vnd.proposal+json", func() {
	Description("A response to a CFP")
	Reference(ProposalModel)
	Attributes(func() {
		Attribute("id", Integer, "ID of user")
		Attribute("href", String, "API href of user")
		Attribute("title", String, "Response title")
		Attribute("abstract", String, "Response abstract")
		Attribute("detail", String, "Response detail")
	})
	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("title")
		Attribute("abstract")
		Attribute("detail")
	})
	View("link", func() {
		Attribute("id")
		Attribute("href")
		Attribute("title")
	})
})
