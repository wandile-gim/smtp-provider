package application

type AdminConfigUseCase interface {
	// CheckAllConfigAvailable 등록된 설정들이 모두 유효한지 검사한다.
	CheckAllConfigAvailable() []error
}
