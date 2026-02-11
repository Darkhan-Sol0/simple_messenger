package usecase

type (
	chatUsecase struct {
	}

	ChatUsecase interface {
	}
)

func New() ChatUsecase {
	return &chatUsecase{}
}
