package validators

import (
	"centralserver/internal/datatypes"
	log "centralserver/pkg/logger"
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	log *log.Logger
  validate *validator.Validate
}

var validate *validator.Validate

func NewValidatorService() *Validator {
	log := log.NewLogger()
	return &Validator{
		log: log,
    validate: validator.New(),
	}
}

func (v *Validator) ValidateTransactionPayload(payload []byte) (bool, error) {
	// Directly unmarshal payload into the struct
	var userPayload *datatypes.UserTransactionPayload
	if err := json.Unmarshal(payload, &userPayload); err != nil {
		return false, err
	}

	// Validate the struct
	validateErr := v.validate.Struct(userPayload)
	if validateErr != nil {
		v.log.Error("Error validating transaction payload: %s", validateErr)
		return false, validateErr
	}

	v.log.Info("Transaction payload validated successfully")
	return true, nil
}
