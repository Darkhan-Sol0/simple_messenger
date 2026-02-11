package usecase

type (
	userUsecase struct {
	}

	UserUsecase interface {
	}
)

func New() UserUsecase {
	return &userUsecase{}
}
