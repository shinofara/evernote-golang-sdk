// Autogenerated by Thrift Compiler (0.9.1)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"UserStore"
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  bool checkVersion(string clientName, i16 edamVersionMajor, i16 edamVersionMinor)")
	fmt.Fprintln(os.Stderr, "  BootstrapInfo getBootstrapInfo(string locale)")
	fmt.Fprintln(os.Stderr, "  AuthenticationResult authenticate(string username, string password, string consumerKey, string consumerSecret, bool supportsTwoFactor)")
	fmt.Fprintln(os.Stderr, "  AuthenticationResult authenticateLongSession(string username, string password, string consumerKey, string consumerSecret, string deviceIdentifier, string deviceDescription, bool supportsTwoFactor)")
	fmt.Fprintln(os.Stderr, "  AuthenticationResult completeTwoFactorAuthentication(string authenticationToken, string oneTimeCode, string deviceIdentifier, string deviceDescription)")
	fmt.Fprintln(os.Stderr, "  void revokeLongSession(string authenticationToken)")
	fmt.Fprintln(os.Stderr, "  AuthenticationResult authenticateToBusiness(string authenticationToken)")
	fmt.Fprintln(os.Stderr, "  AuthenticationResult refreshAuthentication(string authenticationToken)")
	fmt.Fprintln(os.Stderr, "  User getUser(string authenticationToken)")
	fmt.Fprintln(os.Stderr, "  PublicUserInfo getPublicUserInfo(string username)")
	fmt.Fprintln(os.Stderr, "  PremiumInfo getPremiumInfo(string authenticationToken)")
	fmt.Fprintln(os.Stderr, "  string getNoteStoreUrl(string authenticationToken)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = math.MinInt32 // will become unneeded eventually
	_ = strconv.Atoi
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := UserStore.NewUserStoreClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "checkVersion":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "CheckVersion requires 3 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		tmp1, err52 := (strconv.Atoi(flag.Arg(2)))
		if err52 != nil {
			Usage()
			return
		}
		argvalue1 := byte(tmp1)
		value1 := argvalue1
		tmp2, err53 := (strconv.Atoi(flag.Arg(3)))
		if err53 != nil {
			Usage()
			return
		}
		argvalue2 := byte(tmp2)
		value2 := argvalue2
		fmt.Print(client.CheckVersion(value0, value1, value2))
		fmt.Print("\n")
		break
	case "getBootstrapInfo":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetBootstrapInfo requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetBootstrapInfo(value0))
		fmt.Print("\n")
		break
	case "authenticate":
		if flag.NArg()-1 != 5 {
			fmt.Fprintln(os.Stderr, "Authenticate requires 5 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		argvalue3 := flag.Arg(4)
		value3 := argvalue3
		argvalue4 := flag.Arg(5) == "true"
		value4 := argvalue4
		fmt.Print(client.Authenticate(value0, value1, value2, value3, value4))
		fmt.Print("\n")
		break
	case "authenticateLongSession":
		if flag.NArg()-1 != 7 {
			fmt.Fprintln(os.Stderr, "AuthenticateLongSession requires 7 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		argvalue3 := flag.Arg(4)
		value3 := argvalue3
		argvalue4 := flag.Arg(5)
		value4 := argvalue4
		argvalue5 := flag.Arg(6)
		value5 := argvalue5
		argvalue6 := flag.Arg(7) == "true"
		value6 := argvalue6
		fmt.Print(client.AuthenticateLongSession(value0, value1, value2, value3, value4, value5, value6))
		fmt.Print("\n")
		break
	case "completeTwoFactorAuthentication":
		if flag.NArg()-1 != 4 {
			fmt.Fprintln(os.Stderr, "CompleteTwoFactorAuthentication requires 4 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		argvalue3 := flag.Arg(4)
		value3 := argvalue3
		fmt.Print(client.CompleteTwoFactorAuthentication(value0, value1, value2, value3))
		fmt.Print("\n")
		break
	case "revokeLongSession":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "RevokeLongSession requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.RevokeLongSession(value0))
		fmt.Print("\n")
		break
	case "authenticateToBusiness":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "AuthenticateToBusiness requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.AuthenticateToBusiness(value0))
		fmt.Print("\n")
		break
	case "refreshAuthentication":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "RefreshAuthentication requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.RefreshAuthentication(value0))
		fmt.Print("\n")
		break
	case "getUser":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetUser requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetUser(value0))
		fmt.Print("\n")
		break
	case "getPublicUserInfo":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetPublicUserInfo requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetPublicUserInfo(value0))
		fmt.Print("\n")
		break
	case "getPremiumInfo":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetPremiumInfo requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetPremiumInfo(value0))
		fmt.Print("\n")
		break
	case "getNoteStoreUrl":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetNoteStoreUrl requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetNoteStoreUrl(value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}