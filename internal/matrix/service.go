package matrix

type iService interface {
	Rotate90(*matrixDTO)
}

type service struct{}

func newService() iService {
	return &service{}
}

func (s *service) Rotate90(arr *matrixDTO) {
	var rotatedMatrix [][]uint
	for i := 0; i < len(*arr); i++ {
		var tmp []uint
		for x := 0; x < len(*arr); x++ {
			tmp = append(tmp, (*arr)[x][i])
		}
		rotatedMatrix = append(rotatedMatrix, tmp)
	}
	*arr = nil
	for v := range rotatedMatrix {
		*arr = append(*arr, rotatedMatrix[len(rotatedMatrix)-1-v])
	}
}
