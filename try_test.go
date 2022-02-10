package try

import (
	"errors"
	"testing"
)

func TestRoutine_Catch(t *testing.T) {
	type args struct {
		catch func(v interface{})
	}
	tests := []struct {
		name    string
		routine Routine
		args    args
	}{
		{
			name: "int",
			routine: func() {
				panic(int(1))
			},
			args: args{catch: func(v interface{}) {
				if i, ok := v.(int); ok && i == 1 {
					return
				}
				t.Fatalf("expect: 1, but: %v", v)
			}},
		},
		{
			name: "error",
			routine: func() {
				panic(errors.New("panic"))
			},
			args: args{catch: func(v interface{}) {
				if e, ok := v.(error); ok && e.Error() == "panic" {
					return
				}
				t.Fatalf("expect error: 'panic', but error: %v", v)
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.routine.Catch(tt.args.catch)
		})
	}
}
