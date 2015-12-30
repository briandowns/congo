//************************************************************************//
// congo Swagger Spec
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/gopheracademy/congo
// --design=github.com/gopheracademy/congo/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package swagger

import (
	"github.com/julienschmidt/httprouter"
	"github.com/raphael/goa"
)

// MountController mounts the swagger spec controller under "/swagger.json".
func MountController(service goa.Service) {
	ctrl := service.NewController("Swagger")
	service.Info("mount", "ctrl", "Swagger", "action", "Show", "route", "GET /swagger.json")
	h := ctrl.NewHTTPRouterHandle("Show", getSwagger)
	service.HTTPHandler().(*httprouter.Router).Handle("GET", "/swagger.json", h)
}

// getSwagger is the httprouter handle that returns the Swagger spec.
// func getSwagger(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
func getSwagger(ctx *goa.Context) error {
	ctx.Header().Set("Content-Type", "application/swagger+json")
	ctx.Header().Set("Cache-Control", "public, max-age=3600")
	return ctx.Respond(200, []byte(spec))
}

// Generated spec
const spec = `{"swagger":"2.0","info":{"title":"Congo - Conference Management Made Easy","description":"Multi-tenant conference management application","contact":{"name":"Gopher Academy","email":"social@gopheracademy.com","url":"https://gopheracademy.com"},"license":{"name":"MIT","url":"https://github.com/gopheracademy/congo/blob/master/LICENSE"},"version":""},"host":"api.gopheracademy.com","basePath":"/api","schemes":["http"],"consumes":["application/json"],"produces":["application/json"],"paths":{"/auth/refresh":{"post":{"description":"Obtain a refreshed access token","operationId":"auth#refresh","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"payload","in":"body","required":true,"schema":{"$ref":"#/definitions/RefreshAuthPayload"}}],"responses":{"201":{"description":""}},"schemes":["https"]}},"/auth/token":{"post":{"description":"Obtain an access token","operationId":"auth#token","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"payload","in":"body","required":true,"schema":{"$ref":"#/definitions/TokenAuthPayload"}}],"responses":{"201":{"description":""}},"schemes":["https"]}},"/auth/{provider}":{"get":{"description":"OAUTH2 login endpoint","operationId":"auth#oauth","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"provider","in":"path","required":true,"type":"string"}],"responses":{"200":{"description":"","schema":{"$ref":"#/definitions/Authorize"}}},"schemes":["https"]}},"/auth/{provider}/callback":{"get":{"description":"OAUTH2 callback endpoint","operationId":"auth#callback","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"provider","in":"path","required":true,"type":"string"}],"responses":{"200":{"description":""}},"schemes":["https"]}},"/users":{"get":{"description":"List all users in account","operationId":"user#list","consumes":["application/json"],"produces":["application/json"],"responses":{"200":{"description":""}},"schemes":["https"]},"post":{"description":"Record new user","operationId":"user#create","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"payload","in":"body","required":true,"schema":{"$ref":"#/definitions/CreateUserPayload"}}],"responses":{"201":{"description":"Resource created","headers":{"Location":{"description":"href to created resource","type":"string","pattern":"^/accounts/[0-9]+/users/[0-9]+$"}}}},"schemes":["https"]}},"/users/{userID}":{"get":{"description":"Retrieve user with given id","operationId":"user#show","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"userID","in":"path","required":true,"type":"integer"}],"responses":{"200":{"description":"","schema":{"$ref":"#/definitions/User"}},"404":{"description":""}},"schemes":["https"]},"delete":{"operationId":"user#delete","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"userID","in":"path","description":"User ID","required":true,"type":"integer"}],"responses":{"204":{"description":""},"404":{"description":""}},"schemes":["https"]},"patch":{"operationId":"user#update","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"userID","in":"path","required":true,"type":"integer"},{"name":"payload","in":"body","required":true,"schema":{"$ref":"#/definitions/UpdateUserPayload"}}],"responses":{"204":{"description":""},"404":{"description":""}},"schemes":["https"]}},"/users/{userID}/proposals":{"get":{"description":"List all proposals for a user","operationId":"proposal#list","consumes":["application/json"],"produces":["application/json"],"responses":{"200":{"description":""}},"schemes":["https"]},"post":{"description":"Create a new proposal","operationId":"proposal#create","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"payload","in":"body","required":true,"schema":{"$ref":"#/definitions/CreateProposalPayload"}}],"responses":{"201":{"description":"Resource created","headers":{"Location":{"description":"href to created resource","type":"string","pattern":"^/users/[0-9]+/proposals/[0-9]+$"}}}},"schemes":["https"]}},"/users/{userID}/proposals/{proposalID}":{"get":{"description":"Retrieve proposal with given id","operationId":"proposal#show","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"proposalID","in":"path","required":true,"type":"integer"}],"responses":{"200":{"description":"","schema":{"$ref":"#/definitions/Proposal"}},"404":{"description":""}},"schemes":["https"]},"delete":{"operationId":"proposal#delete","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"proposalID","in":"path","description":"Proposal ID","required":true,"type":"integer"}],"responses":{"204":{"description":""},"404":{"description":""}},"schemes":["https"]},"patch":{"operationId":"proposal#update","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"proposalID","in":"path","required":true,"type":"integer"},{"name":"payload","in":"body","required":true,"schema":{"$ref":"#/definitions/UpdateProposalPayload"}}],"responses":{"204":{"description":""},"404":{"description":""}},"schemes":["https"]}},"/users/{userID}/proposals/{proposalID}/review":{"get":{"description":"List all reviews for a proposal","operationId":"review#list","consumes":["application/json"],"produces":["application/json"],"responses":{"200":{"description":""}},"schemes":["https"]},"post":{"description":"Create a new review","operationId":"review#create","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"payload","in":"body","required":true,"schema":{"$ref":"#/definitions/CreateReviewPayload"}}],"responses":{"201":{"description":"Resource created","headers":{"Location":{"description":"href to created resource","type":"string","pattern":"^/users/[0-9]+/proposals/[0-9]+/reviews/[0-9]+$"}}}},"schemes":["https"]}},"/users/{userID}/proposals/{proposalID}/review/{reviewID}":{"get":{"description":"Retrieve review with given id","operationId":"review#show","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"reviewID","in":"path","required":true,"type":"integer"}],"responses":{"200":{"description":"","schema":{"$ref":"#/definitions/Review"}},"404":{"description":""}},"schemes":["https"]},"delete":{"operationId":"review#delete","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"reviewID","in":"path","description":"Review ID","required":true,"type":"integer"}],"responses":{"204":{"description":""},"404":{"description":""}},"schemes":["https"]},"patch":{"operationId":"review#update","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"reviewID","in":"path","required":true,"type":"integer"},{"name":"payload","in":"body","required":true,"schema":{"$ref":"#/definitions/UpdateReviewPayload"}}],"responses":{"204":{"description":""},"404":{"description":""}},"schemes":["https"]}}},"definitions":{"Authorize":{"title":"Mediatype identifier: application/vnd.authorize+json","type":"object","properties":{"access_token":{"type":"string","description":"access token"},"expires_in":{"type":"integer","description":"Time to expiration in seconds"},"token_type":{"type":"string","description":"type of token"}},"description":"Token authorization response"},"CreateProposalPayload":{"title":"CreateProposalPayload","type":"object","properties":{"abstract":{"type":"string","minLength":50,"maxLength":500},"detail":{"type":"string","minLength":100,"maxLength":2000},"firstname":{"type":"string","minLength":2},"title":{"type":"string","minLength":10,"maxLength":200},"withdrawn":{"type":"boolean"}},"required":["detail"]},"CreateReviewPayload":{"title":"CreateReviewPayload","type":"object","properties":{"comment":{"type":"string","minLength":10,"maxLength":200},"rating":{"type":"integer","minimum":1,"maximum":5}},"required":["rating"]},"CreateUserPayload":{"title":"CreateUserPayload","type":"object","properties":{"bio":{"type":"string","maxLength":500},"city":{"type":"string"},"country":{"type":"string"},"email":{"type":"string"},"firstname":{"type":"string"},"lastname":{"type":"string"},"role":{"type":"string"},"state":{"type":"string"}},"required":["email"]},"Proposal":{"title":"Mediatype identifier: application/vnd.proposal+json","type":"object","properties":{"abstract":{"type":"string","description":"Response abstract","minLength":50,"maxLength":500},"detail":{"type":"string","description":"Response detail","minLength":100,"maxLength":2000},"href":{"type":"string","description":"API href of user"},"id":{"type":"integer","description":"ID of user"},"title":{"type":"string","description":"Response title","minLength":10,"maxLength":200}},"description":"A response to a CFP"},"RefreshAuthPayload":{"title":"RefreshAuthPayload","type":"object","properties":{"application":{"type":"string","description":"UUID of requesting application"},"email":{"type":"string","description":"email"},"password":{"type":"string","description":"password"}}},"Review":{"title":"Mediatype identifier: application/vnd.review+json","type":"object","properties":{"comment":{"type":"string","description":"Review comments","minLength":10,"maxLength":200},"href":{"type":"string","description":"API href of user"},"id":{"type":"integer","description":"ID of user"},"rating":{"type":"integer","description":"Rating of proposal, from 1-5","minimum":1,"maximum":5}},"description":"A review is submitted by a reviewer"},"TokenAuthPayload":{"title":"TokenAuthPayload","type":"object","properties":{"application":{"type":"string","description":"UUID of requesting application"},"email":{"type":"string","description":"email"},"password":{"type":"string","description":"password"}}},"UpdateProposalPayload":{"title":"UpdateProposalPayload","type":"object","properties":{"abstract":{"type":"string","minLength":50,"maxLength":500},"detail":{"type":"string","minLength":100,"maxLength":2000},"firstname":{"type":"string","minLength":2},"title":{"type":"string","minLength":10,"maxLength":200},"withdrawn":{"type":"boolean"}}},"UpdateReviewPayload":{"title":"UpdateReviewPayload","type":"object","properties":{"comment":{"type":"string","minLength":10,"maxLength":200},"rating":{"type":"integer","minimum":1,"maximum":5}}},"UpdateUserPayload":{"title":"UpdateUserPayload","type":"object","properties":{"bio":{"type":"string","maxLength":500},"city":{"type":"string"},"country":{"type":"string"},"email":{"type":"string"},"firstname":{"type":"string"},"lastname":{"type":"string"},"role":{"type":"string"},"state":{"type":"string"}}},"User":{"title":"Mediatype identifier: application/vnd.user+json","type":"object","properties":{"bio":{"type":"string","description":"Biography of user","maxLength":500},"city":{"type":"string","description":"City of residence"},"country":{"type":"string","description":"Country of residence"},"email":{"type":"string","description":"Email address of user","format":"email"},"firstname":{"type":"string","description":"First name of user"},"href":{"type":"string","description":"API href of user"},"id":{"type":"integer","description":"ID of user"},"lastname":{"type":"string","description":"Last name of user"},"state":{"type":"string","description":"State of residence"}},"description":"A user belonging to a tenant account"}},"externalDocs":{"description":"congo guide","url":"https://gopheracademy.com/congo/getting-started.html"}} `
