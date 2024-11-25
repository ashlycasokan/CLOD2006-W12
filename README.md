# Introducing Security Flaws
we intentionally designed the application with several vulnerabilities to demonstrate typical security risks:
•	Weak input validation: There’s no check on the input being sent to the server, making it prone to injection attacks.
•	Missing authentication: The endpoint is open to everyone, meaning there’s no control over who can access it.
•	Poor secret management: Sensitive information like secrets is handled insecurely, without any encryption.
•	Container issues: The application uses a large, unoptimized Docker base image, which unnecessarily increases the risk of exploitation.
•	Cloud environment threats: we configured the environment poorly, with things like overly permissive IAM roles, open ports, and secrets stored in plaintext.

a.Code-Level Issues

Insecure Application (main.go)
go
Copy code
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/echo", echoHandler)
	http.ListenAndServe(":8080", nil)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// Insecure: No input validation
	query := r.URL.Query().Get("input")
	fmt.Fprintf(w, "Echo: %s", query)
}

// Insecure: Exposing sensitive data via environment variable
func getEnvVariable() string {
	return os.Getenv("SECRET_KEY") // No validation or encryption
}


Flaws in the Code

1.	No Input Validation:
•	The echoHandler function directly takes user input without validating or sanitizing it.
•	Risk: This opens the door for injection attacks, such as XSS or command injection, which can compromise the application.
2.	No Authentication or Authorization:
•	The /echo endpoint is accessible to everyone, with no restrictions.
•	Risk: Sensitive data or operations could be accessed and abused by unauthorized users.
3.	Sensitive Data Handling:
•	The SECRET_KEY environment variable is directly accessed without any encryption or obfuscation.
•	Risk: If logs or memory dumps are exposed, attackers could easily steal this sensitive information.

b. Container Security

Insecure Dockerfile

dockerfile
Copy code
# Insecure: Using a large, vulnerable base image
FROM golang:1.20

WORKDIR /app
COPY . .

RUN go build -o app .
CMD ["./app"]




Flaws in the Dockerfile

1.	Large Base Image:
•	The golang:1.20 image contains unnecessary development tools and libraries.
•	Risk: This increases the attack surface, making it more likely that vulnerabilities exist within the image.
2.	No Vulnerability Scanning:
•	There’s no process in place to check for known issues in the base image or dependencies.
•	Risk: Unpatched vulnerabilities in the image or libraries could be exploited by attackers.
3.	Lack of Multi-Stage Builds:
•	Development tools and dependencies are included in the final production image.
•	Risk: This bloats the image unnecessarily and makes it less secure.

