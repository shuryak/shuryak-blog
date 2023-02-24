package errors

import "fmt"

type ValidationError struct {
	Source       string
	Descriptions []string
}

func (v *ValidationError) Error() string {
	var descriptions string
	for _, desc := range v.Descriptions {
		descriptions += desc
	}
	return fmt.Sprintf("validation error in %s:\n%s", v.Source, descriptions)
}

type ServerError string

const (
	serverNoStart     ServerError = "Unable to start %s server. Error: %v \n"
	serverNoHandler   ServerError = "Unable to register service handlers. Error %v\n"
	dbNoConnection    ServerError = "Unable to connect to DB %s. Error: %v\n"
	missingField      ServerError = "%s must not be empty\n"
	authNoMetadata    ServerError = "Unable to read metadata for endpoint: %s\n"
	authNoToken       ServerError = "No token\n"
	authInvalidToken  ServerError = "Invalid token\n"
	authInvalidClaim  ServerError = "Invalid %s claim\n"
	authNoUserInToken ServerError = "Unable to get logged in user from access token. Error: %v\n"
)

func (se *ServerError) ServerNoStart(serviceName string, err error) string {
	return fmt.Sprintf(string(serverNoStart), serviceName, err)
}

func (se *ServerError) ServerNoHandler(err error) string {
	return fmt.Sprintf(string(serverNoHandler), err)
}

func (se *ServerError) DbNoConnection(dbName string, err error) string {
	return fmt.Sprintf(string(dbNoConnection), dbName, err)
}

func (se *ServerError) MissingField(fieldName string) string {
	return fmt.Sprintf(string(missingField), fieldName)
}

func (se *ServerError) AuthNoMetadata(endpoint string) string {
	return fmt.Sprintf(string(authNoMetadata), endpoint)
}

func (se *ServerError) AuthNoToken() string {
	return fmt.Sprintf(string(authNoToken))
}

func (se *ServerError) AuthInvalidToken() string {
	return fmt.Sprintf(string(authInvalidToken))
}

func (se *ServerError) AuthInvalidClaim(claimType string) string {
	return fmt.Sprintf(string(authInvalidClaim), claimType)
}

func (se *ServerError) AuthNoUserInToken(err error) string {
	return fmt.Sprintf(string(authNoUserInToken), err)
}
