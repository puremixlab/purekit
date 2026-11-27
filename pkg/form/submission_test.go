package form

import (
	"context"
	"testing"
	"testing/synctest"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/field"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

func TestNamedFieldRendersWithNilValue(t *testing.T) {
	f := New(profile{})
	tt := ui.NewTester(func(c *ui.Context) {
		theme.Apply(c)
		View(c, f, Props[profile]{}, func() {
			Field(c, f, FieldProps{Name: "optional", Value: nil}, func() {
				ui.Text(c, "No selection")
			})
		})
	}, 400, 200)
	if !tt.HasText("No selection") {
		t.Fatalf("nil-valued field did not render its control: %q", tt.Texts())
	}
}

func TestEditingInvalidatesPendingSubmission(t *testing.T) {
	for _, tc := range []struct {
		name        string
		resultReady bool
	}{
		{name: "validation running"},
		{name: "result queued", resultReady: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				f := New(profile{})
				validated := make(chan string, 2)
				finishFirstCheck := make(chan struct{}, 1)
				defer close(finishFirstCheck)
				var submitted []string
				tt := ui.NewTester(func(c *ui.Context) {
					theme.Apply(c)
					View(c, f, Props[profile]{
						ValidateOn: ValidateSubmit,
						OnSubmit: func(v profile) error {
							submitted = append(submitted, v.Name)
							return nil
						},
					}, func() {
						Field(c, f, FieldProps{
							Name:  "name",
							Value: f.Values.Name,
							ValidateAsync: func(ctx context.Context, value any) string {
								name := value.(string)
								validated <- name
								if name == "" {
									select {
									case <-finishFirstCheck:
									case <-ctx.Done():
									}
								}
								return ""
							},
						}, func() {
							field.Input(c, &f.Values.Name, field.InputProps{}).Label("name input")
						})
						SubmitFooter(c, f, SubmitFooterProps{})
					})
				}, 480, 240)

				click := func(label string) {
					t.Helper()
					if err := tt.Click(label); err != nil {
						t.Fatal(err)
					}
				}
				expectValidation := func(want string) {
					t.Helper()
					synctest.Wait()
					select {
					case got := <-validated:
						if got != want {
							t.Fatalf("validated %q, want %q", got, want)
						}
					default:
						t.Fatalf("no validation for %q", want)
					}
				}

				click("Submit")
				expectValidation("")
				click("name input")
				if tc.resultReady {
					finishFirstCheck <- struct{}{}
					synctest.Wait()
				}
				tt.Type("Ada")
				if !tc.resultReady {
					finishFirstCheck <- struct{}{}
					synctest.Wait()
				}
				tt.Frame()
				if f.Values.Name != "Ada" {
					t.Fatalf("editing did not change the field: %q", f.Values.Name)
				}
				if len(submitted) != 0 {
					t.Fatalf("edit did not invalidate earlier submission, submitted %q", submitted)
				}

				click("Submit")
				expectValidation("Ada")
				tt.Frame()
				if len(submitted) != 1 || submitted[0] != "Ada" {
					t.Fatalf("resubmission got %q, want [Ada]", submitted)
				}
			})
		})
	}
}
