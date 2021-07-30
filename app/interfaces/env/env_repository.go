package env

import (
	"d-meeting.com/d-server/domain"
)

type EnvRepository struct {
	EnvHandler
}

func (repo *EnvRepository) ReadEnvValue(key string) (envs domain.Env, err error) {
	result, err := repo.ReadEnv(key)
	if err != nil {
		return domain.ERROR_RESPONSE, err
	}

	env := domain.Env{
		Value: result,
	}
	return env, nil
}
