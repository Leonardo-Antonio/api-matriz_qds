package matrix

import "errors"

func ValidateMatrixNxN(arr [][]uint) error {
	for _, sub := range arr {
		if len(sub) != len(arr) {
			return errors.New("la matriz no es valida (NxN)")
		}
	}

	return nil
}

func ValidateSlice() {

}
