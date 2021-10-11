package task

type VotedListRefresh struct{}

func newVotedListTask() (*VotedListRefresh, error) {
	return &VotedListRefresh{}, nil
}

func (b *VotedListRefresh) Start() error {
	return nil
}

func (b *VotedListRefresh) Stop() error {
	return nil
}
