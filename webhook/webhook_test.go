package webhook

import (
	"fmt"
	"testing"
)

func TestCheckPayload(t *testing.T) {
	cases := []struct {
		data              []byte
		secret, signature string
	}{
		{
			data:      []byte(``),
			secret:    `secret`,
			signature: `+eZuF5tnR65UEI+C+K3os8Jddv0wr95sOVgixTAZYWk=`,
		},
		{
			data:      []byte(`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus congue diam eget est volutpat accumsan. Aenean sed aliquam metus, at tincidunt ligula. Proin at felis ultricies, euismod nulla in, vulputate eros. Nam tincidunt augue vitae urna semper, in auctor enim mollis. Aliquam vitae feugiat felis, sed congue nisi. Curabitur velit justo, laoreet quis neque a, imperdiet ullamcorper diam. Vestibulum sit amet nunc ipsum. Suspendisse commodo sit amet dui in ultricies. Vestibulum tellus mi, tempor in eleifend eu, semper sit amet ligula. Donec pretium augue quis odio suscipit maximus. Nulla mattis lorem orci, a fringilla libero sagittis in. Etiam ac elementum dui. Nam tempor fermentum ante, sed luctus nunc dapibus malesuada. `),
			secret:    `even bigger secret`,
			signature: `r5VKSl4Ej73XC1QLjy+9yOEjbvlVFK5gBksHjdX0d2Y=`,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			same, err := CheckPayload(c.data, c.signature, c.secret)
			if err != nil {
				t.Fatalf("could not check payload: %v", err)
			}

			if !same {
				t.Fatalf("Signatures are not the same")
			}
		})
	}
}
