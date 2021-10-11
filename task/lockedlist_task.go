package task

type LockedListRefresh struct{}

func newLockedListTask() (*LockedListRefresh, error) {
	return &LockedListRefresh{}, nil
}

func (b *LockedListRefresh) Start() error {
	return nil
}

func (b *LockedListRefresh) Stop() error {
	return nil
}
