package smile_face_test

import (
	"beneversitup/smile_face"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmileFace(t *testing.T) {
	t.Run("test [:-), ;~D, :~) should return 2", func(t *testing.T) {
		smileFace := smile_face.NewSmileFace()
		actual := smileFace.CountSmileFaces([]string{":)", ";(", ";}", ":-D"})
		expected := 2
		assert.Equal(t, expected, actual)
	})

	t.Run("test ;D :-( :-) ;~) should return 3", func(t *testing.T) {
		smileFace := smile_face.NewSmileFace()
		actual := smileFace.CountSmileFaces([]string{";D", ":-(", ":-)", ";~)"})
		expected := 3
		assert.Equal(t, expected, actual)
	})

	t.Run("test ;] :[ ;* :$ ;-D should return 1", func(t *testing.T) {
		smileFace := smile_face.NewSmileFace()
		actual := smileFace.CountSmileFaces([]string{";]", ":[", ";*", ":$", ";-D"})
		expected := 1
		assert.Equal(t, expected, actual)
	})

}
