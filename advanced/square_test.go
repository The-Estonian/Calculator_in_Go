package advanced

import "testing"

func TestSquare(t *testing.T) {
	t.Run("square of one number", func(t *testing.T) {
        sum := Square(2)
        expected := 4
        
        if sum != expected {
            t.Errorf("Expected %d; but got %d", expected, sum)
        }
    })
}