package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gopheracademy/congo/client"
	"gopkg.in/alecthomas/kingpin.v2"
)

// PrettyPrint is true if the tool output should be formatted for human consumption.
var PrettyPrint bool

func main() {
	// Create command line parser
	app := kingpin.New("congo-cli", "CLI client for the congo service (https://gopheracademy.com/congo/getting-started.html)")
	c := client.New()
	c.UserAgent = "congo-cli/1.0"
	app.Flag("scheme", "Set the requests scheme").Short('s').Default("http").StringVar(&c.Scheme)
	app.Flag("host", "API hostname").Short('h').Default("api.gopheracademy.com").StringVar(&c.Host)
	app.Flag("timeout", "Set the request timeout, defaults to 20s").Short('t').Default("20s").DurationVar(&c.Timeout)
	app.Flag("dump", "Dump HTTP request and response.").BoolVar(&c.Dump)
	app.Flag("pp", "Pretty print response body").BoolVar(&PrettyPrint)
	commands := RegisterCommands(app)
	// Make "client-cli <action> [<resource>] --help" equivalent to
	// "client-cli help <action> [<resource>]"
	if os.Args[len(os.Args)-1] == "--help" {
		args := append([]string{os.Args[0], "help"}, os.Args[1:len(os.Args)-1]...)
		os.Args = args
	}
	cmdName, err := app.Parse(os.Args[1:])
	if err != nil {
		kingpin.Fatalf(err.Error())
	}
	cmd, ok := commands[cmdName]
	if !ok {
		kingpin.Fatalf("unknown command %s", cmdName)
	}
	resp, err := cmd.Run(c)
	if err != nil {
		kingpin.Fatalf("request failed: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		kingpin.Fatalf("failed to read body: %s", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		// Let user know if something went wrong
		var sbody string
		if len(body) > 0 {
			sbody = ": " + string(body)
		}
		fmt.Printf("error: %d%s", resp.StatusCode, sbody)
	} else if !c.Dump && len(body) > 0 {
		var out string
		if PrettyPrint {
			var jbody interface{}
			err = json.Unmarshal(body, &jbody)
			if err != nil {
				out = string(body)
			} else {
				var b []byte
				b, err = json.MarshalIndent(jbody, "", "    ")
				if err == nil {
					out = string(b)
				} else {
					out = string(body)
				}
			}
		} else {
			out = string(body)
		}
		fmt.Print(out)
	}

	// Figure out exit code
	exitStatus := 0
	switch {
	case resp.StatusCode == 401:
		exitStatus = 1
	case resp.StatusCode == 403:
		exitStatus = 3
	case resp.StatusCode == 404:
		exitStatus = 4
	case resp.StatusCode > 399 && resp.StatusCode < 500:
		exitStatus = 2
	case resp.StatusCode > 499:
		exitStatus = 5
	}
	os.Exit(exitStatus)
}

// RegisterCommands all the resource action subcommands to the application command line.
func RegisterCommands(app *kingpin.Application) map[string]client.ActionCommand {
	res := make(map[string]client.ActionCommand)
	var command, sub *kingpin.CmdClause
	command = app.Command("callback", "OAUTH2 callback endpoint")
	tmp1 := new(CallbackAuthCommand)
	sub = command.Command("auth", "OAUTH2 callback endpoint")
	tmp1.RegisterFlags(sub)
	res["callback auth"] = tmp1
	command = app.Command("create", "create action")
	tmp2 := new(CreateProposalCommand)
	sub = command.Command("proposal", "Create a new proposal")
	tmp2.RegisterFlags(sub)
	res["create proposal"] = tmp2
	tmp3 := new(CreateReviewCommand)
	sub = command.Command("review", "Create a new review")
	tmp3.RegisterFlags(sub)
	res["create review"] = tmp3
	tmp4 := new(CreateUserCommand)
	sub = command.Command("user", "Record new user")
	tmp4.RegisterFlags(sub)
	res["create user"] = tmp4
	command = app.Command("delete", "delete action")
	tmp5 := new(DeleteProposalCommand)
	sub = command.Command("proposal", "")
	tmp5.RegisterFlags(sub)
	res["delete proposal"] = tmp5
	tmp6 := new(DeleteReviewCommand)
	sub = command.Command("review", "")
	tmp6.RegisterFlags(sub)
	res["delete review"] = tmp6
	tmp7 := new(DeleteUserCommand)
	sub = command.Command("user", "")
	tmp7.RegisterFlags(sub)
	res["delete user"] = tmp7
	command = app.Command("list", "list action")
	tmp8 := new(ListProposalCommand)
	sub = command.Command("proposal", "List all proposals for a user")
	tmp8.RegisterFlags(sub)
	res["list proposal"] = tmp8
	tmp9 := new(ListReviewCommand)
	sub = command.Command("review", "List all reviews for a proposal")
	tmp9.RegisterFlags(sub)
	res["list review"] = tmp9
	tmp10 := new(ListUserCommand)
	sub = command.Command("user", "List all users in account")
	tmp10.RegisterFlags(sub)
	res["list user"] = tmp10
	command = app.Command("oauth", "OAUTH2 login endpoint")
	tmp11 := new(OauthAuthCommand)
	sub = command.Command("auth", "OAUTH2 login endpoint")
	tmp11.RegisterFlags(sub)
	res["oauth auth"] = tmp11
	command = app.Command("refresh", "Obtain a refreshed access token")
	tmp12 := new(RefreshAuthCommand)
	sub = command.Command("auth", "Obtain a refreshed access token")
	tmp12.RegisterFlags(sub)
	res["refresh auth"] = tmp12
	command = app.Command("show", "show action")
	tmp13 := new(ShowProposalCommand)
	sub = command.Command("proposal", "Retrieve proposal with given id")
	tmp13.RegisterFlags(sub)
	res["show proposal"] = tmp13
	tmp14 := new(ShowReviewCommand)
	sub = command.Command("review", "Retrieve review with given id")
	tmp14.RegisterFlags(sub)
	res["show review"] = tmp14
	tmp15 := new(ShowUserCommand)
	sub = command.Command("user", "Retrieve user with given id")
	tmp15.RegisterFlags(sub)
	res["show user"] = tmp15
	command = app.Command("token", "Obtain an access token")
	tmp16 := new(TokenAuthCommand)
	sub = command.Command("auth", "Obtain an access token")
	tmp16.RegisterFlags(sub)
	res["token auth"] = tmp16
	command = app.Command("update", "update action")
	tmp17 := new(UpdateProposalCommand)
	sub = command.Command("proposal", "")
	tmp17.RegisterFlags(sub)
	res["update proposal"] = tmp17
	tmp18 := new(UpdateReviewCommand)
	sub = command.Command("review", "")
	tmp18.RegisterFlags(sub)
	res["update review"] = tmp18
	tmp19 := new(UpdateUserCommand)
	sub = command.Command("user", "")
	tmp19.RegisterFlags(sub)
	res["update user"] = tmp19

	return res
}
