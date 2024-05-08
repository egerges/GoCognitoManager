package error

import "log"

func (e GoCognitoManagerError) Error() {
	log.Fatalf("ConfigError: status %v, description: %s", e.Status, e.Description)
}
