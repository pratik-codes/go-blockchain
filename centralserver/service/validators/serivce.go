package validators

import (
	"centralserver/internal/datatypes"
	log "centralserver/pkg/logger"
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	log *log.Logger
}

var validate *validator.Validate

func NewValidatorService() *Validator {
	return &Validator{}
}

func (v *Validator) ValidateTransactionPayload(payload map[string]interface{}) (bool, error) {
	// Convert map to JSON and then to the struct
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return false, err
	}

	var userPayload *datatypes.UserTransactionPayload
	if err := json.Unmarshal(jsonData, &userPayload); err != nil {
		return false, err
	}

	// Validate the struct
	validateErr := validate.Struct(userPayload)
	if validateErr != nil {
		return false, validateErr
	}

	return true, nil
}
