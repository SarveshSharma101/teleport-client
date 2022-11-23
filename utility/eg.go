package utility

import (
	"context"
	"fmt"
	"log"

	"github.com/gravitational/teleport/api/breaker"
	"github.com/gravitational/teleport/api/client"
	"github.com/gravitational/teleport/api/types"
	"google.golang.org/grpc"
)

func main1() {
	ctx := context.Background()
	// client.NewDialer()
	clt, err := client.New(ctx, client.Config{
		Addrs: []string{"fcftport.northeurope.cloudapp.azure.com:443"},
		Credentials: []client.Credentials{client.LoadIdentityFileFromString(`
		-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAq9kmw0xeK7Sk+u6AP6JFzRRGFGMszV3dTaRJWd4nOBDWyxeZ
cO5Xcs41WG1j5RuZ6i0xAeY/ilr2Fh+VqEmV9+7x4ZwBpFnEPf1GNkP8cE18kMxc
CIIJ5jLe2CaP3lUMfYC2Vfrn5TuYpkh42q+iocaST/VyG0I//BGBq16bCYQItZor
aTramKvEKwvUFZniqWNmn9BmxZY+JkRd0o3+6vcR77CUOV1QFtp5VF+1NEq32xyY
LsAfSsBC9ors62spvauJDDE2qpC969GyZtMeHl+yeZY6BfIRvpIonSyn1rIGNaMo
thOjhd+lhutow8oeXVfrRJp2+mVFPGtkJXwgwQIDAQABAoIBACvZAoz+ZWDdfOMS
C+MwsoK7U45MJ9hWxOrUqmqlxngFw+iaIuqvxdxjRTVE5CJHQqR/12tWpovU3cmT
UYpZDEzwyQL53WlkBjCS+WFLQ5QcEVaY7jq1g3UbMcURQdBG/yLevqd9l8HjzPVQ
tJWIDwUcN6TzSaQu6UiV5vIk94YboVP+BWOmv2LOK+njn4b/KXBK3fRYFrVnbqgb
Uonc3+5xFNOmN9lwZbdpf339zVmhqbSDi48hek+PKtND1Tt5w9VNnGmMr04ca37L
lneB25JP2OJAJZqEFqKsBbfsDmYASnsVZvIUw20260CO7Hx6iBCixICvPcZTPV+B
E3e4FmECgYEA1uA1qZCFlm+/Wdg1FFSa9Ws3JlxXTFx8iqDbNEMpv/CMcMBLfzRr
AazH/PTLGszByoYMLJaZ2l7CyipYIXDTagNIq7Bh3/ztuci+ckCTBAnno3gj76/w
OeM8T95gz5/zqYJgrOuk/1Y474Q3CiPwsVBiPMrnODWrqdPYkc/Va5UCgYEAzLzQ
59db2ESEdHzqYmOZa5Uf0PKEeiGaud2+uxIr6VGKdArhvLDNsx76OOX5p11apG1i
91tTWRVdDnQZb2uLPJLGSWxzEopOswKqFn3gCR1jR/RLnyN93EjoiH7WhoXxVZNj
Rn37jfUiXw/9gDjNkX/L3r0A0zZ8vg6cJTTN9X0CgYEAqItpzD2Oa3fv8jtDN1U9
zy9wLOjVmRIapmqSRqZAA0xd/Lr4/ShSnxe2Ltac0cK2Z3NJ7VecCsu1ovof8usb
CdbVk8Zgn1834ThzGg9Iwiw6w+Ci34lztxRk5IkcCv/+EoIv7rNP0vEA6+8bdRrH
222gAOtu/ooqwqVnBMivMRkCgYATAz3LKd6nmMVMJAWIGYO3z+NifgL7bH933zXF
HYziX+YhnJkV8r1Hcwr9ma2zcyHlvxq/qcV1slwv6WwrQqttdpvfWajeAeYJDahJ
6mqRrh74IaGcJ6maeVLOyiiV2X5t2gAQHcbsieUlOtDpxVyhtGJ0Tszy0E6JP8YU
VnEB6QKBgQC/tDa4xP77b9SGDtJnR7lAzDuFdGXjr0v0wwbi6szCgPxSipwiWDPX
VeHYnTSY8t+aWNrv2QNEi3Zvo4sQ+a7ojc1YPi6iJVv22Z/gQWMXht0fZ3XWwrNw
Tm+WARLGdxhvLLn+OT10TZiPuHPslLP25PpYhYEeX1MleB0etanm0Q==
-----END RSA PRIVATE KEY-----
ssh-rsa-cert-v01@openssh.com AAAAHHNzaC1yc2EtY2VydC12MDFAb3BlbnNzaC5jb20AAAAgprHvmO6sCLQugqd7gAjy7nO49vIrHxCytBgGq+rhMFAAAAADAQABAAABAQCr2SbDTF4rtKT67oA/okXNFEYUYyzNXd1NpElZ3ic4ENbLF5lw7ldyzjVYbWPlG5nqLTEB5j+KWvYWH5WoSZX37vHhnAGkWcQ9/UY2Q/xwTXyQzFwIggnmMt7YJo/eVQx9gLZV+uflO5imSHjar6KhxpJP9XIbQj/8EYGrXpsJhAi1mitpOtqYq8QrC9QVmeKpY2af0GbFlj4mRF3Sjf7q9xHvsJQ5XVAW2nlUX7U0SrfbHJguwB9KwEL2iuzraym9q4kMMTaqkL3r0bJm0x4eX7J5ljoF8hG+kiidLKfWsgY1oyi2E6OF36WG62jDyh5dV+tEmnb6ZUU8a2QlfCDBAAAAAAAAAAAAAAABAAAAB3NhcnZlc2gAAAAKAAAABm52aWRpYQAAAABjP/tnAAAAAGNApGMAAAAAAAAB8QAAAAxpbXBlcnNvbmF0b3IAAABQAAAATDJlMDc2NTBjLWRkNzAtNDZhMS1hZGY1LTM4NDMzMDAzNGM2MS5mY2Z0cG9ydC5ub3J0aGV1cm9wZS5jbG91ZGFwcC5henVyZS5jb20AAAAXcGVybWl0LWFnZW50LWZvcndhcmRpbmcAAAAAAAAAFnBlcm1pdC1wb3J0LWZvcndhcmRpbmcAAAAAAAAACnBlcm1pdC1wdHkAAAAAAAAADnRlbGVwb3J0LXJvbGVzAAAAMAAAACx7InZlcnNpb24iOiJ2MSIsInJvbGVzIjpbImVkaXRvciIsImFjY2VzcyJdfQAAABl0ZWxlcG9ydC1yb3V0ZS10by1jbHVzdGVyAAAAKwAAACdmY2Z0cG9ydC5ub3J0aGV1cm9wZS5jbG91ZGFwcC5henVyZS5jb20AAAAPdGVsZXBvcnQtdHJhaXRzAAAAlQAAAJF7ImF3c19yb2xlX2FybnMiOm51bGwsImRiX25hbWVzIjpudWxsLCJkYl91c2VycyI6bnVsbCwia3ViZXJuZXRlc19ncm91cHMiOm51bGwsImt1YmVybmV0ZXNfdXNlcnMiOm51bGwsImxvZ2lucyI6WyJudmlkaWEiXSwid2luZG93c19sb2dpbnMiOm51bGx9AAAAAAAAARcAAAAHc3NoLXJzYQAAAAMBAAEAAAEBAKEhTZVstvFLYz9t9Ug+v2kiFK1jUi5o+CBlIHAwpyGDclKiW2iVQSK9XdJ2HVZGEOjCyI+RCr2hN5/vaYDX43DtzHOg95Yuz3uCG2MpSe2hV7VVOXTJpinis7ocfyqk5BuFEQdzgCB42OvXwHP8fXDxnI5RbWQBXwG5rJFi/JT+/HwZ7ERLKnCNxzzzEl7qizr9HUIxJ59BU75rSUKwkrgJl6ksGl5sawcGzCDnD93FfRNjZpb/eUjYlPFpqiI+bpTvOIQYlIWoBX1kjuWKG2xBhIEVf38F5a3MqKKlGSBbnRIuGZ8lcndyO5uw1NaFa+v1THgh/wjUF0ur7eEX66UAAAEUAAAADHJzYS1zaGEyLTUxMgAAAQCTbpR3s1SAy8mmAZuPW8KnTvks/877wToFI5UvMxnzUrMzo0Dq4yfCs9srpfRnDFXr380jz0Yybll4MaCiRqMLquSAqAIojSv0jrLV8nIX5Mr52b1cU8vTBJFbUdcvtpGirZ2sMaV8LaDM3ClGj19lzjxBrBL9XSz5R1qc7FMdThIPljCO4yWx78TYBjfaAXbGRLE/Ej7/SPdruxXcSJt/vH7ei93jNpPihkLDYALho8fmGCTWGTIHO0AipTIt42CcRRKsHVuvMLcSDKPRuFqWrjIYxLkFj5v6g0zdNy41fy7Q6i0lnBL/zHa+bYooSA9ZMsZeYjoOIoktRWP7eE/8
-----BEGIN CERTIFICATE-----
MIIFJDCCBAygAwIBAgIQe+ELfATCFglVwXaJANqD4zANBgkqhkiG9w0BAQsFADCB
ljEwMC4GA1UEChMnZmNmdHBvcnQubm9ydGhldXJvcGUuY2xvdWRhcHAuYXp1cmUu
Y29tMTAwLgYDVQQDEydmY2Z0cG9ydC5ub3J0aGV1cm9wZS5jbG91ZGFwcC5henVy
ZS5jb20xMDAuBgNVBAUTJzIwMDQ3Mjk3Njg2MDYzMjM5ODU5NDc3OTY1MDUyODUx
NjY0Mzk4NTAeFw0yMjEwMDcxMDExNTFaFw0yMjEwMDcyMjEyNTFaMIIBoTEPMA0G
A1UEBxMGbnZpZGlhMTAwLgYDVQQJEydmY2Z0cG9ydC5ub3J0aGV1cm9wZS5jbG91
ZGFwcC5henVyZS5jb20xgZwwgZkGA1UEEQyBkXsiYXdzX3JvbGVfYXJucyI6bnVs
bCwiZGJfbmFtZXMiOm51bGwsImRiX3VzZXJzIjpudWxsLCJrdWJlcm5ldGVzX2dy
b3VwcyI6bnVsbCwia3ViZXJuZXRlc191c2VycyI6bnVsbCwibG9naW5zIjpbIm52
aWRpYSJdLCJ3aW5kb3dzX2xvZ2lucyI6bnVsbH0xHjANBgNVBAoTBmFjY2VzczAN
BgNVBAoTBmVkaXRvcjEQMA4GA1UEAxMHc2FydmVzaDEyMDAGBSvODwEHEydmY2Z0
cG9ydC5ub3J0aGV1cm9wZS5jbG91ZGFwcC5henVyZS5jb20xVzBVBgUrzg8CBxNM
MmUwNzY1MGMtZGQ3MC00NmExLWFkZjUtMzg0MzMwMDM0YzYxLmZjZnRwb3J0Lm5v
cnRoZXVyb3BlLmNsb3VkYXBwLmF6dXJlLmNvbTCCASIwDQYJKoZIhvcNAQEBBQAD
ggEPADCCAQoCggEBAKvZJsNMXiu0pPrugD+iRc0URhRjLM1d3U2kSVneJzgQ1ssX
mXDuV3LONVhtY+UbmeotMQHmP4pa9hYflahJlffu8eGcAaRZxD39RjZD/HBNfJDM
XAiCCeYy3tgmj95VDH2AtlX65+U7mKZIeNqvoqHGkk/1chtCP/wRgatemwmECLWa
K2k62pirxCsL1BWZ4qljZp/QZsWWPiZEXdKN/ur3Ee+wlDldUBbaeVRftTRKt9sc
mC7AH0rAQvaK7OtrKb2riQwxNqqQvevRsmbTHh5fsnmWOgXyEb6SKJ0sp9ayBjWj
KLYTo4XfpYbraMPKHl1X60SadvplRTxrZCV8IMECAwEAAaNgMF4wDgYDVR0PAQH/
BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMBAf8E
AjAAMB8GA1UdIwQYMBaAFJV+uIpCqL4TStiqjLPIntvqi2lCMA0GCSqGSIb3DQEB
CwUAA4IBAQDdHE/Dy6EkgFJMJXyKDmHPdqcIRIkif7fxdxtlwGd1n2c2LJ5LaH7y
UyPvoXZ8seH8c6ynfbhWge6u97Vg9R69kV8WceObGtm0+yteQA/IXxwEKEK5e7Rj
U7nmeCWAKmvPWqqdZRnX7RibHq5YWb5kFNWVXVsKAszUojDiqKprneDcGZ/qsCeS
ssHlgQxWiu8ETsqg493t9h2Eg0Cu/ZNTXZiJb/Raw9meU7XqbfF07HXPHlv0adL1
NWBrrNhPsZz0EDHSprDaCbuDN2Q4t3aY+JRJGZnmj2/9B5raS5sGTLH6RvdrCf7u
cPAgJlRuhqJXr+JQmP6ku2AqYHAi3A4Q
-----END CERTIFICATE-----
@cert-authority fcftport.northeurope.cloudapp.azure.com,*.fcftport.northeurope.cloudapp.azure.com ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDQC0DNtM34vs0gt4AQAZkSJaAelchYqaHbmxHisr/kEZNRnOdAVPIg7pP3TZyicSMoQqUXl6SlHb7qnBIvIeuotj4Aw4yxKppYShL9QfkFvGDuPSVnqX0K9LTrCTu6rQGeklEqBmBx2S45liDUMPPuBwJZrzyn9ZK04pDXAA2RpuMIDGyOA+gbR0W8KfcLE0jkLXY/ZwQczUMII04kjx4/nyio+nWQjwhr3u2HwmdQESFHCY+RKJMsKJgXP1uYR8823x3cF+8+KIJmgo1s65lVbEWp4phIPb3kajQkCU9HhR/YdAoGniijFE7+r43yoYZ8gYab+GjcYi8igiYxNXF5 type=host
-----BEGIN CERTIFICATE-----
MIID+zCCAuOgAwIBAgIRALQQuAsHH9/NvWUSTDh9KoYwDQYJKoZIhvcNAQELBQAw
gZYxMDAuBgNVBAoTJ2ZjZnRwb3J0Lm5vcnRoZXVyb3BlLmNsb3VkYXBwLmF6dXJl
LmNvbTEwMC4GA1UEAxMnZmNmdHBvcnQubm9ydGhldXJvcGUuY2xvdWRhcHAuYXp1
cmUuY29tMTAwLgYDVQQFEycyMzkzNDc4NDg4MjgxMDMyMDI1ODExMjA0OTcwMDYx
NzA0MTc3OTgwHhcNMjIwMzE1MDIzMTUzWhcNMzIwMzEyMDIzMTUzWjCBljEwMC4G
A1UEChMnZmNmdHBvcnQubm9ydGhldXJvcGUuY2xvdWRhcHAuYXp1cmUuY29tMTAw
LgYDVQQDEydmY2Z0cG9ydC5ub3J0aGV1cm9wZS5jbG91ZGFwcC5henVyZS5jb20x
MDAuBgNVBAUTJzIzOTM0Nzg0ODgyODEwMzIwMjU4MTEyMDQ5NzAwNjE3MDQxNzc5
ODCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAOsv2O4F/YC0FbWasMJW
iF+S5+L8d091M3sGaH5faauc2MIdsGjoQ8Ybamqnu+5Bbnq7MG3U8KoeVzPS2nPC
Bf84JZlbCZ93YhImmiKkMx8vmEvn4Bxt3j9UjZXkXXYpacCP5BCix0S57oBL3ayI
fklL03Wz//FhMNrKjk/RSttTWc6uZ8lvGIeWn+Ggv8RhJm8ifthADFHLAksEQVec
9S4Y0CB/L2Rk/VdjKabZKXi/HUtemkONHyEFVYQcMQm6SZCnsqc7j+MS/moTZbhV
oHJ2hmCpYS3Ac1H7E2X+kO5a1q6gRYdAfGPyAl1n4eeO0m7jBdV+urX483kr5mu7
rDUCAwEAAaNCMEAwDgYDVR0PAQH/BAQDAgGmMA8GA1UdEwEB/wQFMAMBAf8wHQYD
VR0OBBYEFH3YCfsBVhW4/CV5ReKPweCknNPqMA0GCSqGSIb3DQEBCwUAA4IBAQBZ
0Vc/+y2sKMdN+ytkqpZlwr0cekx8hCDrR4UyPCXOXooDhKPXOxEkx5eO1Jt3mjj5
t50OqmxQA2qSi3j2KohlGWxul8w/gOiV3OLi6M3C1cf3zzoFZ9tXTzYmslSm09vt
riv5PDyPjuhqM5q+orSv9hcezwwQ7ctIBJzqS8OMXLg7WQDfVjjsIqnWSKvvFxX+
w+CAPpiF3I2rrhn9ioc7XcYAem9GRx/fHF0781YRLiI0K125RhsPi5S1C3/YkDdn
De+3om84URV1HmN4LX38fjaEndwx82+kQoJckqyxU/5f6k8cSvqrIUiZYD8OpMR0
RGvsK3ChxQpeNIMoQ9O0
-----END CERTIFICATE-----`)},
		Dialer:                     nil,
		DialOpts:                   []grpc.DialOption{},
		DialInBackground:           false,
		DialTimeout:                0,
		KeepAlivePeriod:            0,
		KeepAliveCount:             0,
		InsecureAddressDiscovery:   false,
		ALPNSNIAuthDialClusterName: "",
		CircuitBreakerConfig:       breaker.Config{},
		Context:                    nil,
	})

	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	defer clt.Close()

	resp, err := clt.Ping(ctx)
	if err != nil {
		log.Fatalf("failed to ping server: %v", err)
	}

	// clt.
	log.Printf("Example success!")
	log.Printf("Example server response: %s", resp.ClusterName)
	log.Printf("Server version: %s", resp.ServerVersion)
	l, _ := clt.GetNodes(ctx, "")
	log.Println("---------->", len(l))
	fmt.Println("")
	// for i, s := range l {
	// 	fmt.Print(i, " --> ", s.GetHostname())
	// 	if check(s.GetHostname()) {
	// 		fmt.Println(" present")
	// 	} else {
	// 		fmt.Println(" not present")
	// 	}

	// 	fmt.Println("------------------------------------------")
	// }
	// check(l)
	for i, s := range l {
		fmt.Println(i, " -> ", s.GetHostname())
	}
}

func check(s []types.Server) {
	a := []string{"DE002013", "DE002014", "DE002015", "DE002051", "DE002333", "DE002782", "DE002783", "DE002784", "DE002796", "DE003321", "DE003325", "DE002497", "DE003323", "DE003497", "DE003324", "DE003285", "DE003543", "DE002781", "DE003435"}
	i := 0
	for _, v := range a {
		i = 0
		fmt.Print("--> ", v)
		for _, s2 := range s {
			if s2.GetHostname() == v {
				i = 1
				fmt.Println(" Present")
			}
		}
		if i == 0 {
			fmt.Println(" Not present")
		}
		fmt.Println("------------------------------")
	}

}
