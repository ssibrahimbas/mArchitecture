package example

import "testing"

func TestFactoryConfig_Validate(t *testing.T) {
	type fields struct {
		MinFieldLength   int
		MaxFieldLength   int
		MinContentLength int
		MaxContentLength int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid config",
			fields: fields{
				MinFieldLength:   1,
				MaxFieldLength:   100,
				MinContentLength: 1,
				MaxContentLength: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := FactoryConfig{
				MinFieldLength:   tt.fields.MinFieldLength,
				MaxFieldLength:   tt.fields.MaxFieldLength,
				MinContentLength: tt.fields.MinContentLength,
				MaxContentLength: tt.fields.MaxContentLength,
			}
			if err := f.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("FactoryConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
