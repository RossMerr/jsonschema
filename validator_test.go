package jsonschema

import "testing"

func Test_validator_ValidateSchema(t *testing.T) {
	type args struct {
		schema *Schema
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Random Schema",
			args:    args{schema: &Schema{Schema: "http://www.sample.com"}},
			wantErr: true,
		},
		{
			name:    "Draft 2019-09",
			args:    args{schema: &Schema{Schema: "https://json-schema.org/2019-09/schema"}},
			wantErr: false,
		},
		{
			name:    "Draft 08 (Old name)",
			args:    args{schema: &Schema{Schema: "http://json-schema.org/draft-08/schema#"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &validator{}
			if err := s.ValidateSchema(tt.args.schema); (err != nil) != tt.wantErr {
				t.Errorf("ValidateSchema() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
