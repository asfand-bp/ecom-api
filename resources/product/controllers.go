package product

type Controller struct {
	Service ServiceInterface
}

func New() *Controller {
	return &Controller{
		Service: NewService(),
	}
}
