package example

import "github.com/pkg/errors"

type Example struct {
	Key   string
	Value string
}

type FactoryConfig struct {
	MinKeyLength   int
	MaxKeyLength   int
	MinValueLength int
	MaxValueLength int
}

func (f FactoryConfig) Validate() error {
	if f.MinKeyLength < 0 {
		return errors.New("min key length must be greater than or equal to 0")
	}
	if f.MaxKeyLength < 0 {
		return errors.New("max key length must be greater than or equal to 0")
	}
	if f.MinKeyLength > f.MaxKeyLength {
		return errors.New("min key length must be less than or equal to max key length")
	}
	if f.MinValueLength < 0 {
		return errors.New("min value length must be greater than or equal to 0")
	}
	if f.MaxValueLength < 0 {
		return errors.New("max value length must be greater than or equal to 0")
	}
	if f.MinValueLength > f.MaxValueLength {
		return errors.New("min value length must be less than or equal to max value length")
	}
	return nil
}

type Factory struct {
	fc FactoryConfig
}

func NewFactory(fc FactoryConfig) (Factory, error) {
	if err := fc.Validate(); err != nil {
		return Factory{}, errors.Wrap(err, "invalid factory config")
	}
	return Factory{fc: fc}, nil
}

func MustNewFactory(fc FactoryConfig) Factory {
	f, err := NewFactory(fc)
	if err != nil {
		panic(err)
	}
	return f
}

func (f Factory) Config() FactoryConfig {
	return f.fc
}

func (f Factory) IsZero() bool {
	return f == Factory{}
}

func (f Factory) NewExample(key string, value string) (*Example, error) {
	if err := f.validateKey(key); err != nil {
		return nil, errors.Wrap(err, "invalid key")
	}
	if err := f.validateValue(value); err != nil {
		return nil, errors.Wrap(err, "invalid value")
	}
	return &Example{
		Key:   key,
		Value: value,
	}, nil
}

func (f Factory) UnmarshalExample(key string, value string) (*Example, error) {
	if err := f.validateKey(key); err != nil {
		return nil, errors.Wrap(err, "invalid key")
	}
	if err := f.validateValue(value); err != nil {
		return nil, errors.Wrap(err, "invalid value")
	}
	return &Example{
		Key:   key,
		Value: value,
	}, nil
}

func (f Factory) validateKey(key string) error {
	if len(key) < f.fc.MinKeyLength {
		return errors.Errorf("key length must be greater than or equal to %d", f.fc.MinKeyLength)
	}
	if len(key) > f.fc.MaxKeyLength {
		return errors.Errorf("key length must be less than or equal to %d", f.fc.MaxKeyLength)
	}
	return nil
}

func (f Factory) validateValue(value string) error {
	if len(value) < f.fc.MinValueLength {
		return errors.Errorf("value length must be greater than or equal to %d", f.fc.MinValueLength)
	}
	if len(value) > f.fc.MaxValueLength {
		return errors.Errorf("value length must be less than or equal to %d", f.fc.MaxValueLength)
	}
	return nil
}

func (f Factory) NewNotFoundError(key string) error {
	return errors.Errorf("example with key %q not found", key)
}
